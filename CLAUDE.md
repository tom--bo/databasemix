# DatabaseMix

MySQL / PostgreSQL データベースの包括的な情報を取得し、Markdown/XML/Plaintext 形式で出力する Go CLI ツール。

## プロジェクト構成

```
main.go                 - エントリポイント、CLI引数解析、Config構造体、共通データ型定義
collector.go            - Collector インターフェース定義
mysql_collector.go      - MySQL 固有のデータ収集ロジック
mysql_version.go        - MySQL バージョン検出・機能判定 (5.7/8.0/8.4, MariaDB, Percona)
postgres_collector.go   - PostgreSQL 固有のデータ収集ロジック
postgres_version.go     - PostgreSQL バージョン検出
formatter.go            - 出力フォーマッタ (Markdown, XML, Plaintext) - DB種別に応じた表示切替
go.mod                  - Go 1.22, 依存: go-sql-driver/mysql, lib/pq
```

## ビルド & 実行

```bash
go build -o databasemix .

# MySQL
./databasemix -type=mysql -host=localhost -port=3306 -user=root -password=yourpassword

# PostgreSQL
./databasemix -type=postgres -host=localhost -port=5432 -user=postgres -password=yourpassword -database=mydb
```

`-type` 未指定時はポート番号で自動判定 (3306→MySQL, 5432→PostgreSQL)。

環境変数: MySQL は `MYSQL_HOST` 等、PostgreSQL は `PGHOST` 等をサポート。

## テスト

Docker コンテナによる手動テスト（Go テストフレームワークは未使用）:

```bash
# MySQL 8.0
cd test_containers/mysql-8.0 && ./run.sh start && cd ../..
./databasemix -type=mysql -host=localhost -port=3380 -user=root -password=rootpass -database=testdb
cd test_containers/mysql-8.0 && ./run.sh stop

# PostgreSQL 16
cd test_containers/postgres-16 && ./run.sh start && cd ../..
./databasemix -type=postgres -host=localhost -port=5416 -user=postgres -password=rootpass -database=testdb
cd test_containers/postgres-16 && ./run.sh stop
```

### テストコンテナ
- `test_containers/mysql-5.7/` - ポート 3357
- `test_containers/mysql-8.0/` - ポート 3380
- `test_containers/mysql-8.4/` - ポート 3384
- `test_containers/postgres-16/` - ポート 5416
- `test_containers/common/` - MySQL 共通 SQL 初期化スクリプト
- `test_containers/postgres-common/` - PostgreSQL 共通 SQL 初期化スクリプト

接続情報: root(postgres)/rootpass, testuser/testpass, readonly/readpass, admin/adminpass

## アーキテクチャ

```
main.go: parseFlags() → DB種別判定 → connect{MySQL,PostgreSQL}() → New{MySQL,PostgreSQL}Collector()
         → collector.CollectAll() → *DatabaseInfo → NewFormatter() → Format() → ファイル出力
```

### Collector インターフェース
```go
type Collector interface {
    CollectAll() (*DatabaseInfo, error)
}
```
`MySQLCollector` と `PostgreSQLCollector` がこのインターフェースを実装。

### MySQL 固有
- `mysql_version.go` の `MySQLVersion` で 5.7/8.0/8.4 の機能差を吸収
- ロール、コンポーネントは 8.0+ のみ
- 変数取得は performance_schema → information_schema → SHOW VARIABLES の3段階フォールバック

### PostgreSQL 固有
- テーブル DDL は SQL で組み立て (`SHOW CREATE TABLE` が無いため)
- ビュー DDL は `pg_get_viewdef()`
- ルーチンは `pg_proc` + `pg_get_functiondef()`
- 変数は `pg_settings`
- ユーザー/ロールは `pg_roles` の `rolcanlogin` で区別
- Extensions サポート (`pg_extension`)

## コーディング規約

- 標準ライブラリ + `go-sql-driver/mysql` + `lib/pq` のみ使用
- エラーは `log.Printf` で警告し処理を継続（致命的でないエラーはスキップ）
- DB 固有ロジックは `{mysql,postgres}_collector.go` に閉じ込め、`main.go` と `formatter.go` は共通
