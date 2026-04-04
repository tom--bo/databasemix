# DatabaseMix

MySQL データベースの包括的な情報を取得し、Markdown/XML/Plaintext 形式で出力する Go CLI ツール。

## プロジェクト構成

```
main.go        - エントリポイント、CLI引数解析、Config構造体、データ型定義
version.go     - MySQLバージョン検出、バージョン別機能判定 (5.7/8.0/8.4, MariaDB, Percona)
collector.go   - MySQLからのデータ収集ロジック (テーブル、ユーザー、変数、ルーチン等)
formatter.go   - 出力フォーマッタ (Markdown, XML, Plaintext)
go.mod         - Go 1.22, 依存: go-sql-driver/mysql v1.9.3
```

## ビルド & 実行

```bash
go build -o databasemix .
./databasemix -host=localhost -port=3306 -user=root -password=yourpassword
```

環境変数 `MYSQL_HOST`, `MYSQL_PORT`, `MYSQL_USER`, `MYSQL_PASSWORD`, `MYSQL_DATABASE` も使用可能。

## テスト

Docker コンテナによる手動テスト（Go テストフレームワークは未使用）:

```bash
# MySQL 8.0 の例
cd test_containers/mysql-8.0 && ./run.sh start
./databasemix -host=localhost -port=3380 -user=root -password=rootpass -database=testdb
cd test_containers/mysql-8.0 && ./run.sh stop
```

### テストコンテナ
- `test_containers/mysql-5.7/` - ポート 3357
- `test_containers/mysql-8.0/` - ポート 3380
- `test_containers/mysql-8.4/` - ポート 3384
- `test_containers/common/` - 共通 SQL 初期化スクリプト (テーブル、ビュー、プロシージャ、テストデータ)

接続情報: root/rootpass, testuser/testpass, readonly/readpass, admin/adminpass

## アーキテクチャ

1. `main.go`: CLI 解析 → MySQL 接続
2. `collector.go`: `NewMySQLCollector()` でバージョン検出 → `CollectAll()` で全データ収集
3. `formatter.go`: `NewFormatter()` で出力形式に応じたフォーマッタ生成 → ファイル出力

### バージョン対応
- `version.go` の `MySQLVersion` 構造体で 5.7/8.0/8.4 の機能差を吸収
- ロール、コンポーネントは 8.0+ のみ
- 変数取得は performance_schema → information_schema → SHOW VARIABLES の3段階フォールバック

## コーディング規約

- 標準ライブラリ + `go-sql-driver/mysql` のみ使用
- エラーは `log.Printf` で警告し処理を継続（致命的でないエラーはスキップ）
- SQL クエリはバージョンに応じて `version.go` のメソッドで切り替え
