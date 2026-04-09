# ACL (ロール・ユーザアカウント) 情報取得 設計書

## 1. 現状分析

### 1.1 現在の実装の問題点

#### MySQL

| 問題 | 詳細 |
|------|------|
| ロール判定が不正確 | `account_locked='Y'` で判定しているが、手動でロックされたユーザもロール扱いになる |
| ロールと一般ユーザが User List に混在 | `analyst`(ロール)と `analyst_user`(ユーザ)が同じリストに表示される |
| ロール階層が不明 | ロール間の継承関係(role → role)が表示されない |
| デフォルトロール未表示 | `SET DEFAULT ROLE` の設定が出力に反映されない |
| リソース制限未表示 | `MAX_QUERIES_PER_HOUR` 等の制限が表示されない |
| テーブル/カラム単位の権限未取得 | `mysql.tables_priv`, `mysql.columns_priv` を参照していない |
| SSL要件未表示 | `ssl_type`, `ssl_cipher` の情報が欠落 |

#### PostgreSQL

| 問題 | 詳細 |
|------|------|
| ロールの Grants/Members が空 | `User Roles` セクションにロールの権限・メンバーが表示されない |
| ACL文字列が未解析 | `{=Tc/postgres,testuser=CTc/postgres}` のような生の aclitem 形式で表示される |
| スキーマレベル権限未取得 | スキーマに対する USAGE/CREATE 権限を収集していない |
| テーブル/カラム単位の権限未取得 | `information_schema.role_table_grants` 等を参照していない |
| デフォルト権限未取得 | `pg_default_acl` の情報が欠落 |
| ロール階層が不明確 | ロール間の継承関係が Members としてしか表示されない |
| RLS ポリシー未取得 | Row Level Security の設定が出力されない |

---

## 2. MySQL ACL 情報源の調査

### 2.1 利用可能なシステムテーブル/ビュー

#### mysql.user — ユーザ・ロールの基本情報

```sql
SELECT User, Host, plugin, ssl_type, ssl_cipher,
       account_locked, password_expired, password_lifetime,
       max_questions, max_updates, max_connections, max_user_connections,
       User_attributes
FROM mysql.user
ORDER BY User, Host;
```

| カラム | 内容 |
|--------|------|
| `User`, `Host` | アカウント識別子 |
| `plugin` | 認証プラグイン (mysql_native_password, caching_sha2_password 等) |
| `ssl_type` | SSL要件 ('', ANY, X509, SPECIFIED) |
| `ssl_cipher` | 指定されたSSL暗号 |
| `account_locked` | アカウントロック状態 (Y/N) |
| `password_expired` | パスワード有効期限切れ (Y/N) |
| `password_lifetime` | パスワード有効期間 (日数, NULL=デフォルト) |
| `max_questions` | 時間あたりのクエリ上限 |
| `max_updates` | 時間あたりの更新上限 |
| `max_connections` | 時間あたりの接続上限 |
| `max_user_connections` | 同時接続上限 |
| `User_attributes` | JSON属性 (8.0+, コメント等) |

#### mysql.role_edges — ロール割り当て (8.0+)

```sql
SELECT FROM_USER, FROM_HOST, TO_USER, TO_HOST, WITH_ADMIN_OPTION
FROM mysql.role_edges;
```

| カラム | 内容 |
|--------|------|
| `FROM_USER`, `FROM_HOST` | 付与されるロール |
| `TO_USER`, `TO_HOST` | ロールを受けるユーザ/ロール |
| `WITH_ADMIN_OPTION` | ADMIN OPTION 付きか (Y/N) |

#### mysql.default_roles — デフォルトロール (8.0+)

```sql
SELECT USER, HOST, DEFAULT_ROLE_USER, DEFAULT_ROLE_HOST
FROM mysql.default_roles;
```

| カラム | 内容 |
|--------|------|
| `USER`, `HOST` | 対象ユーザ |
| `DEFAULT_ROLE_USER`, `DEFAULT_ROLE_HOST` | デフォルトロール |

#### mysql.tables_priv — テーブルレベル権限

```sql
SELECT Host, Db, User, Table_name, Grantor, Table_priv, Column_priv
FROM mysql.tables_priv
ORDER BY User, Host, Db, Table_name;
```

#### mysql.columns_priv — カラムレベル権限

```sql
SELECT Host, Db, User, Table_name, Column_name, Column_priv
FROM mysql.columns_priv
ORDER BY User, Host, Db, Table_name, Column_name;
```

#### mysql.procs_priv — ルーチンレベル権限

```sql
SELECT Host, Db, User, Routine_name, Routine_type, Proc_priv, Grantor
FROM mysql.procs_priv
ORDER BY User, Host, Db, Routine_name;
```

#### mysql.global_grants — 動的権限 (8.0+)

```sql
SELECT USER, HOST, PRIV, WITH_GRANT_OPTION
FROM mysql.global_grants
ORDER BY USER, HOST, PRIV;
```

#### SHOW GRANTS FOR — 統合された権限表示

```sql
SHOW GRANTS FOR 'user'@'host';
SHOW GRANTS FOR 'user'@'host' USING 'role1'@'%', 'role2'@'%';  -- ロール適用後の実効権限
```

### 2.2 MySQL ロール判定の改善

MySQL にはロールを明示的に識別するシステムカラムが存在しない。`CREATE ROLE` は内部的に `CREATE USER ... ACCOUNT LOCK` を実行するだけである。

**改善案: `mysql.role_edges` を活用した判定**

```sql
-- role_edges の FROM_USER 側に現れ、かつ rolcanlogin 的な判定ができるもの
-- 方針: role_edges に FROM_USER として存在するアカウントをロール候補とする
SELECT DISTINCT u.User, u.Host
FROM mysql.user u
INNER JOIN mysql.role_edges re ON u.User = re.FROM_USER AND u.Host = re.FROM_HOST
WHERE u.account_locked = 'Y'
ORDER BY u.User, u.Host;
```

ただし、この方法でも完全ではない（メンバーが未割当のロールは検出できない）。
現実的には現在の `account_locked='Y'` による推定 + `SHOW GRANTS` の結果からロール付与の有無を判定する併用方式が妥当。

### 2.3 MySQL バージョン別の機能差異

| 機能 | 5.7 | 8.0 | 8.4 |
|------|-----|-----|-----|
| ロール (CREATE ROLE) | x | o | o |
| mysql.role_edges | x | o | o |
| mysql.default_roles | x | o | o |
| mysql.global_grants | x | o | o |
| User_attributes (JSON) | x | o | o |
| password_lifetime | o | o | o |
| max_questions 等リソース制限 | o | o | o |
| ssl_type / ssl_cipher | o | o | o |
| SHOW GRANTS ... USING (実効権限) | x | o | o |

---

## 3. PostgreSQL ACL 情報源の調査

### 3.1 利用可能なシステムカタログ/ビュー

#### pg_roles — ロール基本情報

```sql
SELECT rolname, rolsuper, rolinherit, rolcreaterole, rolcreatedb,
       rolcanlogin, rolreplication, rolconnlimit, rolvaliduntil, rolbypassrls,
       rolconfig
FROM pg_catalog.pg_roles
WHERE rolname NOT LIKE 'pg_%'
ORDER BY rolcanlogin DESC, rolname;
```

| カラム | 内容 |
|--------|------|
| `rolname` | ロール名 |
| `rolsuper` | スーパーユーザか |
| `rolinherit` | 権限を自動継承するか |
| `rolcreaterole` | ロール作成権限 |
| `rolcreatedb` | DB作成権限 |
| `rolcanlogin` | ログイン可能か (true=ユーザ, false=ロール) |
| `rolreplication` | レプリケーション権限 |
| `rolconnlimit` | 接続数上限 (-1=無制限) |
| `rolvaliduntil` | パスワード有効期限 |
| `rolbypassrls` | RLSバイパス権限 |
| `rolconfig` | ロール固有のGUC設定 |

#### pg_auth_members — ロールメンバーシップ

```sql
-- PostgreSQL 16+: admin_option, inherit_option, set_option が追加
SELECT r.rolname AS role, m.rolname AS member,
       am.admin_option, am.inherit_option, am.set_option,
       g.rolname AS grantor
FROM pg_catalog.pg_auth_members am
JOIN pg_catalog.pg_roles r ON am.roleid = r.oid
JOIN pg_catalog.pg_roles m ON am.member = m.oid
JOIN pg_catalog.pg_roles g ON am.grantor = g.oid
ORDER BY r.rolname, m.rolname;
```

| カラム | 内容 | バージョン |
|--------|------|------------|
| `roleid` | 付与されるロールのOID | 全バージョン |
| `member` | ロールを受けるメンバーのOID | 全バージョン |
| `admin_option` | ADMIN OPTION 付きか | 全バージョン |
| `inherit_option` | INHERIT OPTION (16+) | 16+ |
| `set_option` | SET OPTION (16+) | 16+ |
| `grantor` | 付与者のOID | 16+ |

**注意**: PostgreSQL 15以前は `admin_option` のみ。16+で `inherit_option`, `set_option`, `grantor` が追加された。

#### pg_database — データベースレベル権限

```sql
SELECT datname, datacl, datdba,
       pg_catalog.pg_get_userbyid(datdba) AS owner
FROM pg_catalog.pg_database
ORDER BY datname;
```

`datacl` は aclitem 配列で、各要素の形式:
```
grantee=privileges/grantor
```

権限文字の意味:

| 文字 | 権限 |
|------|------|
| `C` | CREATE |
| `c` | CONNECT |
| `T` | TEMPORARY |
| `*` (接尾) | WITH GRANT OPTION |

#### pg_namespace — スキーマレベル権限

```sql
SELECT nspname, nspacl,
       pg_catalog.pg_get_userbyid(nspowner) AS owner
FROM pg_catalog.pg_namespace
WHERE nspname NOT LIKE 'pg_%' AND nspname != 'information_schema'
ORDER BY nspname;
```

権限文字:

| 文字 | 権限 |
|------|------|
| `U` | USAGE |
| `C` | CREATE |

#### pg_class — テーブル/ビュー/シーケンスレベル権限

```sql
SELECT n.nspname AS schema, c.relname AS name,
       c.relkind, c.relacl,
       pg_catalog.pg_get_userbyid(c.relowner) AS owner
FROM pg_catalog.pg_class c
JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
WHERE c.relacl IS NOT NULL
  AND n.nspname NOT LIKE 'pg_%'
  AND n.nspname != 'information_schema'
  AND c.relkind IN ('r', 'v', 'm', 'S', 'f')  -- table, view, matview, sequence, foreign table
ORDER BY n.nspname, c.relname;
```

テーブル権限文字:

| 文字 | 権限 |
|------|------|
| `r` | SELECT (read) |
| `w` | UPDATE (write) |
| `a` | INSERT (append) |
| `d` | DELETE |
| `D` | TRUNCATE |
| `x` | REFERENCES |
| `t` | TRIGGER |

シーケンス権限文字:

| 文字 | 権限 |
|------|------|
| `r` | SELECT (currval) |
| `w` | UPDATE (nextval, setval) |
| `U` | USAGE |

#### information_schema.role_table_grants — テーブル権限 (読みやすい形式)

```sql
SELECT grantor, grantee, table_schema, table_name, privilege_type, is_grantable
FROM information_schema.role_table_grants
WHERE table_schema NOT IN ('pg_catalog', 'information_schema')
ORDER BY grantee, table_schema, table_name, privilege_type;
```

#### pg_default_acl — デフォルト権限

```sql
SELECT pg_catalog.pg_get_userbyid(d.defaclrole) AS owner,
       n.nspname AS schema,
       d.defaclobjtype,
       d.defaclacl
FROM pg_catalog.pg_default_acl d
LEFT JOIN pg_catalog.pg_namespace n ON n.oid = d.defaclnamespace;
```

`defaclobjtype` の値:

| 値 | 対象 |
|----|------|
| `r` | テーブル |
| `S` | シーケンス |
| `f` | 関数 |
| `T` | 型 |
| `n` | スキーマ (15+) |

#### pg_policy — Row Level Security ポリシー

```sql
SELECT pol.polname, c.relname AS table_name, n.nspname AS schema,
       pol.polpermissive, pol.polcmd,
       pg_catalog.pg_get_expr(pol.polqual, pol.polrelid) AS using_expr,
       pg_catalog.pg_get_expr(pol.polwithcheck, pol.polrelid) AS with_check_expr,
       ARRAY(SELECT rolname FROM pg_roles WHERE oid = ANY(pol.polroles)) AS roles
FROM pg_catalog.pg_policy pol
JOIN pg_catalog.pg_class c ON c.oid = pol.polrelid
JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
ORDER BY n.nspname, c.relname, pol.polname;
```

#### aclitem のパース方法

PostgreSQL の ACL 文字列 `grantee=privileges/grantor` をパースする SQL:

```sql
SELECT
    (aclexplode(relacl)).grantor,
    (aclexplode(relacl)).grantee,
    (aclexplode(relacl)).privilege_type,
    (aclexplode(relacl)).is_grantable
FROM pg_class
WHERE relname = 'target_table';
```

`aclexplode()` 関数(9.0+)を使えば aclitem 配列を分解できる。

### 3.2 PostgreSQL バージョン別の機能差異

| 機能 | 16 | 17 | 18 |
|------|-----|-----|-----|
| pg_auth_members.inherit_option | o | o | o |
| pg_auth_members.set_option | o | o | o |
| pg_auth_members.grantor | o | o | o |
| pg_roles.rolbypassrls | o | o | o |
| aclexplode() | o | o | o |
| pg_policy (RLS) | o | o | o |
| pg_default_acl | o | o | o |

対象バージョン (16/17/18) では全機能が利用可能。

---

## 4. データ構造の設計

### 4.1 拡張する構造体

```go
// UserAccount — 既存構造体の拡張
type UserAccount struct {
    User            string
    Host            string       // MySQL のみ
    SSLType         string       // MySQL: '', ANY, X509, SPECIFIED
    Plugin          string       // MySQL: 認証プラグイン / PostgreSQL: 属性文字列
    AccountLocked   string       // MySQL のみ
    PasswordExpired string       // MySQL のみ
    Grants          []string     // SHOW GRANTS / ロールメンバーシップ + DB権限

    // PostgreSQL specific (既存)
    IsSuperuser   bool
    CanCreateDB   bool
    CanCreateRole bool
    ConnLimit     int
    ValidUntil    string

    // === 新規追加フィールド ===

    // MySQL specific (新規)
    PasswordLifetime  *int         // パスワード有効期間 (日数, nil=デフォルト)
    MaxQuestions      int          // MAX_QUERIES_PER_HOUR
    MaxUpdates        int          // MAX_UPDATES_PER_HOUR
    MaxConnections    int          // MAX_CONNECTIONS_PER_HOUR
    MaxUserConnections int         // MAX_USER_CONNECTIONS
    DefaultRoles      []string     // デフォルトロール一覧
    UserAttributes    string       // User_attributes JSON (8.0+)

    // PostgreSQL specific (新規)
    Inherit       bool             // rolinherit
    CanBypassRLS  bool             // rolbypassrls
    IsReplication bool             // rolreplication
    RoleConfig    []string         // rolconfig (GUC設定)
    MemberOf      []RoleMembership // 所属ロール (詳細)

    // 共通 (新規)
    ObjectPrivileges []ObjectPrivilege // テーブル/カラム/ルーチン単位の権限
}

// RoleMembership — ロール所属の詳細情報
type RoleMembership struct {
    RoleName     string
    AdminOption  bool
    InheritOption bool   // PostgreSQL 16+ のみ
    SetOption    bool    // PostgreSQL 16+ のみ
    Grantor      string  // 付与者
}

// ObjectPrivilege — オブジェクト単位の権限
type ObjectPrivilege struct {
    ObjectType string   // TABLE, VIEW, SEQUENCE, FUNCTION, PROCEDURE, SCHEMA, DATABASE, COLUMN
    Schema     string
    ObjectName string
    ColumnName string   // カラム権限の場合のみ
    Privileges []string // SELECT, INSERT, UPDATE, DELETE, etc.
    Grantor    string
    IsGrantable bool
}

// UserRole — 既存構造体の拡張
type UserRole struct {
    RoleName string
    RoleHost string       // MySQL のみ
    Grants   []string
    Members  []RoleMember // []string から変更

    // === 新規追加フィールド ===
    // PostgreSQL specific
    Inherit       bool
    CanBypassRLS  bool
    ObjectPrivileges []ObjectPrivilege
}

// RoleMember — ロールメンバーの詳細情報
type RoleMember struct {
    MemberName   string
    MemberHost   string   // MySQL のみ
    AdminOption  bool
    InheritOption bool    // PostgreSQL 16+
    SetOption    bool     // PostgreSQL 16+
}

// DefaultPrivilege — PostgreSQL デフォルト権限
type DefaultPrivilege struct {
    Owner      string
    Schema     string    // 空の場合はグローバル
    ObjectType string    // TABLE, SEQUENCE, FUNCTION, TYPE, SCHEMA
    Privileges []ParsedACL
}

// ParsedACL — パース済み ACL エントリ
type ParsedACL struct {
    Grantee     string
    Privileges  []string
    Grantor     string
    IsGrantable bool
}

// RLSPolicy — Row Level Security ポリシー
type RLSPolicy struct {
    PolicyName    string
    TableSchema   string
    TableName     string
    Permissive    bool     // true=PERMISSIVE, false=RESTRICTIVE
    Command       string   // ALL, SELECT, INSERT, UPDATE, DELETE
    UsingExpr     string   // USING式
    WithCheckExpr string   // WITH CHECK式
    Roles         []string // 適用対象ロール
}

// DatabaseInfo への追加フィールド
type DatabaseInfo struct {
    // ... 既存フィールド ...
    DefaultPrivileges []DefaultPrivilege  // PostgreSQL のみ
    RLSPolicies       []RLSPolicy         // PostgreSQL のみ
}
```

### 4.2 構造体設計の方針

- `ObjectPrivilege` を共通型として MySQL/PostgreSQL の両方で使用
- MySQL の `SHOW GRANTS` 結果はそのまま `Grants []string` に保持（現状維持）
- PostgreSQL の aclitem は `aclexplode()` でパースし `ObjectPrivilege` に変換
- `RoleMember` を詳細情報付きの構造体に変更（`[]string` → `[]RoleMember`）

---

## 5. 収集クエリの設計

### 5.1 MySQL — 追加収集クエリ

#### ユーザ基本情報の拡張 (全バージョン)

```sql
-- 5.7
SELECT User, Host, plugin, ssl_type,
       account_locked, password_expired, password_lifetime,
       max_questions, max_updates, max_connections, max_user_connections
FROM mysql.user
ORDER BY User, Host;

-- 8.0+
SELECT User, Host, plugin, ssl_type,
       account_locked, password_expired, password_lifetime,
       max_questions, max_updates, max_connections, max_user_connections,
       User_attributes
FROM mysql.user
ORDER BY User, Host;
```

#### デフォルトロールの取得 (8.0+)

```sql
SELECT USER, HOST, DEFAULT_ROLE_USER, DEFAULT_ROLE_HOST
FROM mysql.default_roles
ORDER BY USER, HOST;
```

#### ロール階層の取得 (8.0+)

```sql
SELECT FROM_USER, FROM_HOST, TO_USER, TO_HOST, WITH_ADMIN_OPTION
FROM mysql.role_edges
ORDER BY FROM_USER, FROM_HOST;
```

#### テーブル/カラム/ルーチン権限 (全バージョン)

```sql
-- テーブル権限
SELECT Host, Db, User, Table_name, Grantor, Table_priv, Column_priv
FROM mysql.tables_priv
WHERE Db NOT IN ('sys', 'mysql', 'information_schema', 'performance_schema')
ORDER BY User, Host, Db, Table_name;

-- カラム権限
SELECT Host, Db, User, Table_name, Column_name, Column_priv
FROM mysql.columns_priv
WHERE Db NOT IN ('sys', 'mysql', 'information_schema', 'performance_schema')
ORDER BY User, Host, Db, Table_name, Column_name;

-- ルーチン権限
SELECT Host, Db, User, Routine_name, Routine_type, Proc_priv, Grantor
FROM mysql.procs_priv
WHERE Db NOT IN ('sys', 'mysql', 'information_schema', 'performance_schema')
ORDER BY User, Host, Db, Routine_name;
```

### 5.2 PostgreSQL — 追加収集クエリ

#### ロール基本情報の拡張

```sql
SELECT rolname, rolsuper, rolinherit, rolcreaterole, rolcreatedb,
       rolcanlogin, rolreplication, rolconnlimit, rolvaliduntil,
       rolbypassrls, rolconfig
FROM pg_catalog.pg_roles
WHERE rolname NOT LIKE 'pg_%'
ORDER BY rolcanlogin DESC, rolname;
```

#### ロールメンバーシップ詳細 (PostgreSQL 16+)

```sql
SELECT r.rolname AS role_name,
       m.rolname AS member_name,
       am.admin_option,
       am.inherit_option,
       am.set_option,
       g.rolname AS grantor
FROM pg_catalog.pg_auth_members am
JOIN pg_catalog.pg_roles r ON am.roleid = r.oid
JOIN pg_catalog.pg_roles m ON am.member = m.oid
JOIN pg_catalog.pg_roles g ON am.grantor = g.oid
ORDER BY r.rolname, m.rolname;
```

#### スキーマ権限

```sql
SELECT nspname,
       pg_catalog.pg_get_userbyid(nspowner) AS owner,
       nspacl
FROM pg_catalog.pg_namespace
WHERE nspname NOT LIKE 'pg_%'
  AND nspname != 'information_schema'
ORDER BY nspname;
```

#### テーブル/ビュー/シーケンス権限 (aclexplode 使用)

```sql
SELECT n.nspname AS schema_name,
       c.relname AS object_name,
       CASE c.relkind
           WHEN 'r' THEN 'TABLE'
           WHEN 'v' THEN 'VIEW'
           WHEN 'm' THEN 'MATERIALIZED VIEW'
           WHEN 'S' THEN 'SEQUENCE'
           WHEN 'f' THEN 'FOREIGN TABLE'
       END AS object_type,
       pg_catalog.pg_get_userbyid(c.relowner) AS owner,
       (aclexplode(c.relacl)).grantor::regrole::text AS grantor,
       (aclexplode(c.relacl)).grantee::regrole::text AS grantee,
       (aclexplode(c.relacl)).privilege_type,
       (aclexplode(c.relacl)).is_grantable
FROM pg_catalog.pg_class c
JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
WHERE c.relacl IS NOT NULL
  AND n.nspname NOT LIKE 'pg_%'
  AND n.nspname != 'information_schema'
  AND c.relkind IN ('r', 'v', 'm', 'S', 'f')
ORDER BY n.nspname, c.relname;
```

#### データベース権限 (aclexplode 使用)

```sql
SELECT datname,
       pg_catalog.pg_get_userbyid(datdba) AS owner,
       (aclexplode(datacl)).grantor::regrole::text AS grantor,
       (aclexplode(datacl)).grantee::regrole::text AS grantee,
       (aclexplode(datacl)).privilege_type,
       (aclexplode(datacl)).is_grantable
FROM pg_catalog.pg_database
WHERE datacl IS NOT NULL
ORDER BY datname;
```

#### デフォルト権限

```sql
SELECT pg_catalog.pg_get_userbyid(d.defaclrole) AS owner,
       n.nspname AS schema_name,
       CASE d.defaclobjtype
           WHEN 'r' THEN 'TABLE'
           WHEN 'S' THEN 'SEQUENCE'
           WHEN 'f' THEN 'FUNCTION'
           WHEN 'T' THEN 'TYPE'
           WHEN 'n' THEN 'SCHEMA'
       END AS object_type,
       (aclexplode(d.defaclacl)).grantee::regrole::text AS grantee,
       (aclexplode(d.defaclacl)).privilege_type,
       (aclexplode(d.defaclacl)).is_grantable
FROM pg_catalog.pg_default_acl d
LEFT JOIN pg_catalog.pg_namespace n ON n.oid = d.defaclnamespace
ORDER BY owner, schema_name;
```

#### RLS ポリシー

```sql
SELECT pol.polname,
       n.nspname AS schema_name,
       c.relname AS table_name,
       CASE WHEN pol.polpermissive THEN 'PERMISSIVE' ELSE 'RESTRICTIVE' END AS policy_type,
       CASE pol.polcmd
           WHEN 'r' THEN 'SELECT'
           WHEN 'a' THEN 'INSERT'
           WHEN 'w' THEN 'UPDATE'
           WHEN 'd' THEN 'DELETE'
           WHEN '*' THEN 'ALL'
       END AS command,
       pg_catalog.pg_get_expr(pol.polqual, pol.polrelid) AS using_expr,
       pg_catalog.pg_get_expr(pol.polwithcheck, pol.polrelid) AS with_check_expr,
       ARRAY(SELECT rolname FROM pg_catalog.pg_roles WHERE oid = ANY(pol.polroles)) AS roles
FROM pg_catalog.pg_policy pol
JOIN pg_catalog.pg_class c ON c.oid = pol.polrelid
JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
WHERE n.nspname NOT LIKE 'pg_%'
ORDER BY n.nspname, c.relname, pol.polname;
```

---

## 6. 出力フォーマットの設計

### 6.1 セクション構成の変更

**現在:**
```
# User Roles (MySQL 8.0+)
# User List
```

**変更後:**
```
# User Roles
# User Accounts
# Object Privileges        (新規)
# Default Privileges        (新規, PostgreSQL のみ)
# Row Level Security        (新規, PostgreSQL のみ)
```

### 6.2 User Roles セクション — Markdown 出力例

#### MySQL 8.0+

```markdown
# User Roles

## app_read@%

- Grants:
  - GRANT SELECT ON *.* TO `app_read`@`%`
- Members:
  - app_user1@% (ADMIN: NO)
  - app_user2@% (ADMIN: NO)
  - limited_user@% (ADMIN: NO)

## app_write@%

- Grants:
  - GRANT SELECT, INSERT, UPDATE, DELETE ON `testdb`.* TO `app_write`@`%`
- Members:
  - app_user2@% (ADMIN: NO)

## developer@%

- Grants:
  - GRANT SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, INDEX, ALTER ON `testdb`.* TO `developer`@`%`
- Members:
  - dev_user@% (ADMIN: NO)
```

#### PostgreSQL

```markdown
# User Roles

## app_read

- Attributes: INHERIT
- Members:
  - readonly (ADMIN: NO, INHERIT: YES, SET: YES, Grantor: postgres)
  - app_write (ADMIN: NO, INHERIT: YES, SET: YES, Grantor: postgres)
  - analyst (ADMIN: NO, INHERIT: YES, SET: YES, Grantor: postgres)
- Object Privileges:
  - TABLE public.employees: SELECT (Grantor: postgres)
  - TABLE public.departments: SELECT (Grantor: postgres)

## app_write

- Attributes: INHERIT
- Member Of: app_read
- Members:
  - testuser (ADMIN: NO, INHERIT: YES, SET: YES, Grantor: postgres)
  - app_admin (ADMIN: NO, INHERIT: YES, SET: YES, Grantor: postgres)
- Object Privileges:
  - TABLE public.employees: INSERT, UPDATE, DELETE (Grantor: postgres)

## developer

- Attributes: INHERIT
- Member Of: app_admin
- Members:
  - admin (ADMIN: NO, INHERIT: YES, SET: YES, Grantor: postgres)
```

### 6.3 User Accounts セクション — Markdown 出力例

#### MySQL

```markdown
# User Accounts

## admin@%

- Authentication: mysql_native_password
- Account Locked: NO
- SSL Required: NONE
- Resource Limits: (none)
- Default Roles: (none)
- Grants:
  - GRANT ALL PRIVILEGES ON *.* TO `admin`@`%` WITH GRANT OPTION
  - GRANT APPLICATION_PASSWORD_ADMIN, ... ON *.* TO `admin`@`%` WITH GRANT OPTION

## limited_user@%

- Authentication: mysql_native_password
- Account Locked: NO
- SSL Required: NONE
- Password Lifetime: (default)
- Resource Limits:
  - MAX_QUERIES_PER_HOUR: 1000
  - MAX_CONNECTIONS_PER_HOUR: 100
  - MAX_USER_CONNECTIONS: 5
- Default Roles: app_read@%
- Grants:
  - GRANT USAGE ON *.* TO `limited_user`@`%`
  - GRANT `app_read`@`%` TO `limited_user`@`%`
```

#### PostgreSQL

```markdown
# User Accounts

## admin

- Attributes: CREATEDB, CREATEROLE, INHERIT
- Connection Limit: unlimited
- Member Of:
  - developer (ADMIN: NO, INHERIT: YES, SET: YES)
- Database Privileges:
  - testdb: CONNECT, CREATE, TEMPORARY (Grantor: postgres)
- Schema Privileges:
  - public: USAGE, CREATE (Grantor: postgres)

## readonly

- Attributes: INHERIT
- Connection Limit: unlimited
- Member Of:
  - app_read (ADMIN: NO, INHERIT: YES, SET: YES)
- Database Privileges:
  - testdb: CONNECT (Grantor: postgres)
- Schema Privileges:
  - public: USAGE (Grantor: postgres)
```

### 6.4 Object Privileges セクション — Markdown 出力例

テーブル/カラム/ルーチン単位の個別権限設定がある場合のみ出力する。

#### MySQL

```markdown
# Object Privileges

## Table Privileges

| User | Database | Table | Privileges | Grantor |
|------|----------|-------|------------|---------|
| testuser@% | testdb | employees | SELECT, INSERT, UPDATE | root@localhost |

## Column Privileges

| User | Database | Table | Column | Privileges |
|------|----------|-------|--------|------------|
| analyst_user@% | testdb | employees | salary | SELECT |

## Routine Privileges

| User | Database | Routine | Type | Privileges | Grantor |
|------|----------|---------|------|------------|---------|
| testuser@% | testdb | calc_bonus | FUNCTION | EXECUTE | root@localhost |
```

#### PostgreSQL

```markdown
# Object Privileges

## Table Privileges

| Grantee | Schema | Table | Privileges | Grantor | Grantable |
|---------|--------|-------|------------|---------|-----------|
| testuser | public | employees | SELECT, INSERT, UPDATE, DELETE | postgres | NO |
| readonly | public | employees | SELECT | postgres | NO |
| app_read | public | departments | SELECT | postgres | NO |

## Sequence Privileges

| Grantee | Schema | Sequence | Privileges | Grantor | Grantable |
|---------|--------|----------|------------|---------|-----------|
| testuser | public | employees_id_seq | USAGE, SELECT, UPDATE | postgres | NO |
```

### 6.5 Default Privileges セクション (PostgreSQL のみ)

```markdown
# Default Privileges

| Owner | Schema | Object Type | Grantee | Privileges | Grantable |
|-------|--------|-------------|---------|------------|-----------|
| postgres | public | TABLE | app_read | SELECT | NO |
| postgres | public | TABLE | app_write | INSERT, UPDATE, DELETE | NO |
| postgres | public | SEQUENCE | app_write | USAGE | NO |
```

### 6.6 Row Level Security セクション (PostgreSQL のみ)

```markdown
# Row Level Security Policies

## public.employees

### employee_isolation_policy

- Type: PERMISSIVE
- Command: ALL
- Roles: (all)
- USING: (current_user = employee_name)
- WITH CHECK: (current_user = employee_name)
```

---

## 7. 実装計画

### 7.1 段階的実装

#### Phase 1: 既存情報の改善 (最小限の変更)

1. **MySQL**: ユーザの拡張属性取得 (SSL, リソース制限, パスワード有効期間)
2. **MySQL**: デフォルトロール、ロールメンバー詳細 (WITH ADMIN OPTION) の取得
3. **PostgreSQL**: `rolinherit`, `rolbypassrls`, `rolconfig` の取得
4. **PostgreSQL**: `pg_auth_members` の詳細情報取得 (`inherit_option`, `set_option`, `grantor`)
5. **PostgreSQL**: ロールセクションに Grants と Members を正しく表示
6. **PostgreSQL**: データベース権限の ACL を `aclexplode()` でパースして人間が読める形式に変換
7. **フォーマッタ**: User Accounts の表示を拡張 (新規フィールドの出力)

#### Phase 2: オブジェクト権限の追加

1. **MySQL**: `mysql.tables_priv`, `mysql.columns_priv`, `mysql.procs_priv` の収集
2. **PostgreSQL**: `pg_class.relacl` + `aclexplode()` によるテーブル/ビュー/シーケンス権限の収集
3. **PostgreSQL**: スキーマ権限 (`pg_namespace.nspacl`) の収集
4. **フォーマッタ**: Object Privileges セクションの実装

#### Phase 3: PostgreSQL 固有の高度な ACL 情報

1. **PostgreSQL**: `pg_default_acl` によるデフォルト権限の収集
2. **PostgreSQL**: `pg_policy` による RLS ポリシーの収集
3. **フォーマッタ**: Default Privileges, RLS セクションの実装

### 7.2 変更対象ファイル

| ファイル | 変更内容 |
|----------|----------|
| `src/main.go` | 構造体の拡張 (`UserAccount`, `UserRole`, 新規型の追加, `DatabaseInfo` 拡張) |
| `src/mysql_collector.go` | 拡張属性の収集、デフォルトロール取得、テーブル/カラム/ルーチン権限取得 |
| `src/mysql_version.go` | 変更なし (既存の `SupportsRoles()` で十分) |
| `src/postgres_collector.go` | ロール詳細取得、ACLパース、スキーマ/テーブル権限取得、デフォルト権限、RLS |
| `src/postgres_version.go` | 変更なし (対象バージョン16+で全機能利用可能) |
| `src/formatter.go` | 全フォーマット (Markdown/XML/Plaintext) の拡張、新セクション追加 |

### 7.3 考慮事項

- **権限不足時のフォールバック**: テーブル `mysql.tables_priv` 等にアクセスできない場合は警告を出してスキップ (既存方針と同じ)
- **出力量の制御**: オブジェクト権限が大量になる可能性があるため、システムスキーマ/カタログは除外
- **MySQL 5.7 互換性**: ロール関連の機能は 8.0+ のみ収集、基本的なユーザ属性は全バージョンで取得
- **`aclexplode()` の利用**: PostgreSQL の ACL パースは Go 側でなく SQL 側で行い、実装を簡潔に保つ
- **`SHOW GRANTS` との関係**: MySQL では `SHOW GRANTS` が最も包括的な権限表示なので、Grants フィールドは現状維持。追加の `mysql.tables_priv` 等は補足情報として Object Privileges に表示
