


```bash
migrate create -ext sql -dir db/migrations -seq create_users_table
```

```bash
sqlc generate
```

```bash
# migrate -database "postgresql://postgres:hogehoge@postgres:5432/test_db?sslmode=disable" -path db/migrations up

migrate -database "postgresql://postgres:hogehoge@host.docker.internal:5432/test_db?sslmode=disable" -path db/migrations up
```