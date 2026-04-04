# PostgreSQL 対応設計書

## 概要

MySQL と PostgreSQL を共通インターフェースで抽象化し、`main.go` からは DB エンジンを意識せずに呼び出せるようにする。

## 現状の構成

```
main.go       → Config解析 → connectToMySQL() → NewMySQLCollector() → CollectAll() → Formatter → 出力
collector.go  → MySQLCollector (MySQL固有のSQL)
version.go    → MySQLVersion (MySQL固有のバージョン判定)
formatter.go  → Formatter interface (DB非依存だが、テンプレート内に "MySQL" ハードコード箇所あり)
```

**問題点**: `main.go` が MySQL に直接依存している。`DatabaseInfo` の構造体は汎用的だが、`connectToMySQL()` や `NewMySQLCollector()` が MySQL 固定。

## 設計方針

### 1. Collector インターフェースの導入

```go
// collector.go (新規 or 既存を改修)
type Collector interface {
    CollectAll() (*DatabaseInfo, error)
}
```

`MySQLCollector` と `PostgreSQLCollector` がこのインターフェースを実装する。

### 2. ファイル構成

```
main.go                  → Config解析、DB種別判定、Collector/Formatter呼び出し
collector.go             → Collector interface 定義、DatabaseInfo 等の共通データ構造
mysql_collector.go       → MySQLCollector (現 collector.go の MySQL 固有ロジック)
mysql_version.go         → MySQLVersion (現 version.go をリネーム)
postgres_collector.go    → PostgreSQLCollector (新規)
postgres_version.go      → PostgreSQLVersion (新規)
formatter.go             → Formatter interface + 各フォーマッタ (DB種別に応じた表示切替)
go.mod                   → github.com/lib/pq を追加
```

### 3. Config の変更

```go
type Config struct {
    // 既存フィールド
    Host                   string
    Port                   string
    User                   string
    Password               string
    Database               string
    Format                 string
    OutputFile             string

    // 新規: DB種別
    DBType                 string  // "mysql" or "postgres" (CLIフラグ: -type)

    // MySQL 固有
    Replication            bool
    ExceptPlugins          bool
    OnlyModifiedVariables  bool

    // 共通 except フラグ
    ExceptTables           bool
    ExceptStoredProcedures bool
    ExceptVariables        bool
    ExceptUsers            bool
    ExceptRoles            bool
}
```

**DB 種別の判定**:
1. `-type` フラグが明示的に指定された場合はそれを使用
2. 未指定の場合、ポート番号で推定: `3306` → MySQL, `5432` → PostgreSQL
3. それでも判定できない場合はエラー

**デフォルトポート**: `-type=postgres` の場合、`-port` のデフォルトを `5432` に変更

**環境変数**: PostgreSQL 用に `PGHOST`, `PGPORT`, `PGUSER`, `PGPASSWORD`, `PGDATABASE` を追加

### 4. main.go の変更

```go
func main() {
    config := parseFlags()

    // DB種別に応じて接続・Collector生成
    var collector Collector
    switch config.DBType {
    case "mysql":
        db, err := connectToMySQL(config)
        // ...
        collector, err = NewMySQLCollector(db, config)
    case "postgres":
        db, err := connectToPostgreSQL(config)
        // ...
        collector, err = NewPostgreSQLCollector(db, config)
    }

    info, err := collector.CollectAll()
    // ... (以降は現状と同じ)
}
```

### 5. DatabaseInfo の拡張

現在の `DatabaseInfo` はほぼ汎用的であり、大きな変更は不要。以下のフィールドを追加:

```go
type DatabaseInfo struct {
    DBType          string          // "mysql" or "postgres" (新規)
    ConnectionInfo  *ConnectionInfo
    Tables          []TableInfo
    Users           []UserAccount
    Routines        []RoutineInfo
    Variables       []Variable
    Roles           []UserRole
    Plugins         []Plugin        // MySQL のみ
    Components      []Component     // MySQL のみ
    ReplicationInfo *ReplicationInfo // MySQL のみ (将来的に PostgreSQL のストリーミングレプリケーションも対応可能)
    Extensions      []Extension     // PostgreSQL のみ (新規)
    Schemas         []string        // PostgreSQL のみ: 複数スキーマ対応 (新規)
}
```

**新規型**:

```go
// PostgreSQL 拡張情報
type Extension struct {
    Name           string
    Version        string
    Description    string
}
```

### 6. PostgreSQLCollector の設計

#### 6.1 バージョン検出 (`postgres_version.go`)

```go
type PostgreSQLVersion struct {
    FullVersion string
    Major       int
    Minor       int
}

func DetectPostgreSQLVersion(db *sql.DB) (*PostgreSQLVersion, error) {
    // SELECT version();            → "PostgreSQL 16.3 on ..."
    // SHOW server_version_num;     → "160003"
}
```

バージョンによる機能差: PostgreSQL はバージョン間の差異が MySQL ほど大きくないが、以下を考慮:
- PostgreSQL 14+: ロールの `pg_read_all_settings` 等の定義済みロール
- PostgreSQL 15+: `MERGE` 文、`pg_publication_tables` の変更
- PostgreSQL 16+: `pg_stat_io` 等の新システムビュー

#### 6.2 データ収集メソッド

| 収集対象 | MySQL クエリソース | PostgreSQL クエリソース |
|----------|-------------------|----------------------|
| テーブル一覧 | `information_schema.TABLES` | `information_schema.tables` + `pg_catalog.pg_class` |
| テーブル DDL | `SHOW CREATE TABLE` | `pg_dump --schema-only` 相当の組み立て ※後述 |
| ビュー DDL | `SHOW CREATE VIEW` | `pg_get_viewdef()` |
| ユーザー | `mysql.user` | `pg_catalog.pg_roles` |
| ユーザー権限 | `SHOW GRANTS FOR` | `information_schema.role_table_grants` + `\dp` 相当クエリ |
| ロール | `mysql.user` + `mysql.role_edges` | `pg_catalog.pg_roles` + `pg_catalog.pg_auth_members` |
| 変数 | `performance_schema.global_variables` | `pg_catalog.pg_settings` |
| ストアドプロシージャ | `information_schema.ROUTINES` | `pg_catalog.pg_proc` + `pg_get_functiondef()` |
| プラグイン | `information_schema.PLUGINS` | N/A (代わりに Extensions) |
| 拡張 (Extensions) | N/A | `pg_catalog.pg_extension` |

#### 6.3 PostgreSQL 固有の考慮事項

**テーブル DDL の組み立て**:
PostgreSQL には `SHOW CREATE TABLE` がないため、以下の情報を組み合わせて DDL を構築する:

```go
func (c *PostgreSQLCollector) getTableDDL(schema, table string) (string, error) {
    // 方法: pg_dump を外部コマンドとして呼ぶか、SQLで組み立てるか
    //
    // 推奨: SQL で組み立て (外部コマンド依存を避ける)
    // 1. pg_catalog.pg_attribute + pg_catalog.pg_type → カラム定義
    // 2. pg_catalog.pg_constraint → 制約 (PK, FK, UNIQUE, CHECK)
    // 3. pg_catalog.pg_index → インデックス
    // 4. pg_catalog.pg_class → テーブルオプション
    //
    // もしくは、pg_dump -t <table> --schema-only を exec.Command で呼ぶ方法もある。
    // ただし pg_dump がインストールされていない環境もあるため、
    // SQL ベースの組み立てを基本とし、pg_dump は将来のオプションとする。
}
```

**スキーマ対応**:
PostgreSQL はスキーマ (namespace) の概念があるため、`public` 以外のスキーマも走査する:

```go
func (c *PostgreSQLCollector) getSchemas() ([]string, error) {
    query := `
        SELECT schema_name
        FROM information_schema.schemata
        WHERE schema_name NOT IN ('pg_catalog', 'information_schema', 'pg_toast')
        ORDER BY schema_name`
    // ...
}
```

**ユーザーとロールの統合**:
PostgreSQL ではユーザーとロールは同じ仕組み (`pg_roles`)。`rolcanlogin` で区別:
- `rolcanlogin = true` → ログイン可能 (ユーザー相当)
- `rolcanlogin = false` → ロール

```go
func (c *PostgreSQLCollector) collectUsers(info *DatabaseInfo) error {
    query := `
        SELECT rolname, rolsuper, rolcreaterole, rolcreatedb,
               rolcanlogin, rolreplication, rolconnlimit, rolvaliduntil
        FROM pg_catalog.pg_roles
        WHERE rolcanlogin = true
        AND rolname NOT LIKE 'pg_%'
        ORDER BY rolname`
    // ...
}

func (c *PostgreSQLCollector) collectRoles(info *DatabaseInfo) error {
    query := `
        SELECT rolname, rolsuper, rolcreaterole, rolcreatedb,
               rolcanlogin, rolreplication
        FROM pg_catalog.pg_roles
        WHERE rolcanlogin = false
        AND rolname NOT LIKE 'pg_%'
        ORDER BY rolname`
    // ...
}
```

**変数 (設定パラメータ)**:

```go
func (c *PostgreSQLCollector) collectVariables(info *DatabaseInfo) error {
    query := `
        SELECT name, setting, unit, source, boot_val
        FROM pg_catalog.pg_settings
        ORDER BY name`
    // source が 'default' 以外 → 変更済み (OnlyModifiedVariables 対応)
}
```

**権限の取得**:

```go
func (c *PostgreSQLCollector) getUserGrants(rolname string) ([]string, error) {
    // 1. データベースレベル権限
    //    SELECT datname, privilege_type FROM information_schema.role_usage_grants
    // 2. スキーマレベル権限
    //    SELECT ... FROM information_schema.role_usage_grants WHERE object_type = 'SCHEMA'
    // 3. テーブルレベル権限
    //    SELECT ... FROM information_schema.role_table_grants WHERE grantee = $1
    // 4. ロールメンバシップ
    //    SELECT r.rolname FROM pg_auth_members m JOIN pg_roles r ON m.roleid = r.oid
    //    WHERE m.member = (SELECT oid FROM pg_roles WHERE rolname = $1)
}
```

### 7. Formatter の変更

`DatabaseInfo.DBType` フィールドを参照して表示を切り替える:

```go
func (f *MarkdownFormatter) Format(info *DatabaseInfo) (string, error) {
    // ...
    if info.DBType == "postgres" {
        result.WriteString("**Database Type**: PostgreSQL  \n")
    } else {
        result.WriteString("**Database Type**: MySQL  \n")
    }
    // ...

    // Extensions セクション (PostgreSQL のみ)
    if len(info.Extensions) > 0 {
        result.WriteString("# Extensions\n\n")
        f.formatExtensions(&result, info.Extensions)
    }
}
```

### 8. UserAccount 構造体の調整

MySQL と PostgreSQL でユーザー属性が異なるため、汎用的に保持する:

```go
type UserAccount struct {
    User            string
    Host            string   // MySQL: host, PostgreSQL: 空文字 (pg_hba.conf で制御)
    SSLType         string
    Plugin          string   // MySQL: auth plugin, PostgreSQL: 空文字
    AccountLocked   string
    PasswordExpired string
    Grants          []string
    // PostgreSQL 追加属性
    IsSuperuser     bool     // PostgreSQL: rolsuper
    CanCreateDB     bool     // PostgreSQL: rolcreatedb
    CanCreateRole   bool     // PostgreSQL: rolcreaterole
    ConnLimit       int      // PostgreSQL: rolconnlimit
    ValidUntil      string   // PostgreSQL: rolvaliduntil
}
```

### 9. 依存パッケージ

```bash
go get github.com/lib/pq
```

`go.mod` に追加:
```
require (
    github.com/go-sql-driver/mysql v1.9.3
    github.com/lib/pq v1.10.9
)
```

`main.go` の import:
```go
import (
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"
)
```

### 10. テストコンテナの追加

```
test_containers/
├── common/                    (MySQL 共通)
├── postgres-common/           (PostgreSQL 共通、新規)
│   ├── 01-database-setup.sql
│   ├── 02-tables-views.sql
│   ├── 03-functions.sql
│   ├── 05-sample-data.sql
│   └── 07-limited-user.sql
├── postgres-15/               (新規)
│   ├── Dockerfile
│   ├── docker-compose.yml     (ポート 5415)
│   └── run.sh
├── postgres-16/               (新規)
│   ├── Dockerfile
│   ├── docker-compose.yml     (ポート 5416)
│   └── run.sh
├── mysql-5.7/
├── mysql-8.0/
└── mysql-8.4/
```

### 11. 実装順序

1. **collector.go の分離**: `Collector` インターフェース定義、既存の MySQL ロジックを `mysql_collector.go` に移動
2. **version.go のリネーム**: `mysql_version.go` にリネーム
3. **main.go の改修**: `-type` フラグ追加、DB 種別に応じた分岐
4. **postgres_version.go**: PostgreSQL バージョン検出
5. **postgres_collector.go**: PostgreSQL データ収集の実装
   - 5a. `collectConnectionInfo` + `collectVariables` (最も単純)
   - 5b. `collectUsers` + `collectRoles`
   - 5c. `collectTables` (DDL 組み立てが最も複雑)
   - 5d. `collectRoutines`
   - 5e. `collectExtensions`
6. **formatter.go の改修**: DB 種別の表示切替、Extensions セクション追加
7. **テストコンテナの作成**: PostgreSQL 15/16 の Docker 環境
8. **テスト実行・修正**

## データフロー図

```
main.go
  │
  ├── parseFlags() → Config { DBType: "mysql" | "postgres", ... }
  │
  ├── [MySQL]
  │   ├── connectToMySQL(config)     → *sql.DB
  │   └── NewMySQLCollector(db, config) → Collector
  │
  ├── [PostgreSQL]
  │   ├── connectToPostgreSQL(config) → *sql.DB
  │   └── NewPostgreSQLCollector(db, config) → Collector
  │
  ├── collector.CollectAll()         → *DatabaseInfo
  │
  ├── NewFormatter(format)           → Formatter
  │
  └── formatter.Format(info)         → string → ファイル出力
```

## MySQL との機能対応表

| 機能 | MySQL | PostgreSQL | 備考 |
|------|-------|------------|------|
| テーブル一覧 | `information_schema.TABLES` | `information_schema.tables` | ほぼ同じ |
| テーブル DDL | `SHOW CREATE TABLE` | SQL組み立て | PG は組み立てが必要 |
| ビュー DDL | `SHOW CREATE VIEW` | `pg_get_viewdef()` | |
| ユーザー | `mysql.user` | `pg_roles (rolcanlogin=true)` | |
| ロール | `mysql.user` + `role_edges` | `pg_roles (rolcanlogin=false)` + `pg_auth_members` | |
| 権限 | `SHOW GRANTS FOR` | `information_schema.role_table_grants` 等 | PG は複数ソース |
| 変数 | `performance_schema.global_variables` | `pg_settings` | |
| プロシージャ | `information_schema.ROUTINES` | `pg_proc` + `pg_get_functiondef()` | |
| プラグイン | `information_schema.PLUGINS` | N/A | MySQL のみ |
| 拡張 | N/A | `pg_extension` | PostgreSQL のみ |
| レプリケーション | `SHOW REPLICA STATUS` 等 | `pg_stat_replication` 等 | 将来対応 |
