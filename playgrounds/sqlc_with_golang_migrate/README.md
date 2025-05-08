# playgrounds/sqlc_with_golang_migrate

sqlc x pgx(pgxpool) x golang-migrate の実装検証.

* sqlc  
  https://github.com/sqlc-dev/sqlc
* pgx  
  https://github.com/jackc/pgx
* golang-migrate  
  https://github.com/golang-migrate/migrate

## How to use

### 1. マイグレーションファイルの作成、テーブル定義SQLを書く

```bash
migrate create -ext sql -dir db/migrations -seq create_users_table
```

### 2. db/queries/\*/\*.sql にクエリSQLを書く

### 3. sqlc.yaml を更新する

### 4. sqlc でGoのコードを生成する 

```bash
sqlc generate
```

### 5. マイグレーションしてPostgreSQLにテーブルを作成する

事前にPostgreSQLを起動しておくこと.

```bash
# migrate -database "postgresql://postgres:hogehoge@postgres:5432/test_db?sslmode=disable" -path db/migrations up

migrate -database "postgresql://postgres:hogehoge@host.docker.internal:5432/test_db?sslmode=disable" -path db/migrations up
```

### 6. Goのコードを実行する

```bash
go run -v main.go
```

```bash
$ go run -v main.go
command-line-arguments
2025/05/07 16:19:04 [{1 Brian Kernighan a@example.com} {2 Brian Kernighan a@example.com} {3 Brian Kernighan a@example.com}]
2025/05/07 16:19:04 {4 Brian Kernighan a@example.com}
2025/05/07 16:19:04 true
```
