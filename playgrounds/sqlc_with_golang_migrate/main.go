package main

import (
	"context"
	"log"
	"reflect"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/aazw/vscode-go/playgrounds/sqlc_with_golang_migrate/pkg/db/users"
)

const dsn = "postgresql://postgres:hogehoge@host.docker.internal:5432/test_db?sslmode=disable"

func run() error {
	ctx := context.Background()

	// // 1. without connection pool
	// // この実装的にはコネクションは１つだけ張られ、全リクエストから同じ *pgx.Conn が使われるが、並行呼び出しに対して安全ではない(スレッドセーフではない)
	// conn, err := pgx.Connect(ctx, dsn)
	// if err != nil {
	// 	return err
	// }
	// defer conn.Close(ctx)
	//
	// queries := users.New(conn)

	// 2. with connection pool
	// プールは内部で複数接続を管理し、かつ並行安全
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("pgxpool.ParseConfig: %v", err)
	}
	cfg.MinConns = 2
	cfg.MaxConns = 10
	cfg.MaxConnLifetime = time.Hour
	cfg.HealthCheckPeriod = time.Minute

	dbPool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("pgxpool.NewWithConfig: %v", err)
	}
	defer dbPool.Close()

	queries := users.New(dbPool)

	// list all users
	items, err := queries.ListUsers(ctx)
	if err != nil {
		return err
	}
	log.Println(items)

	// create an user
	insertedUser, err := queries.CreateUser(ctx, users.CreateUserParams{
		Name:  "Brian Kernighan",
		Email: "a@example.com",
	})
	if err != nil {
		return err
	}
	log.Println(insertedUser)

	// get the user we just inserted
	fetchedUser, err := queries.GetUser(ctx, insertedUser.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedUser, fetchedUser))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
