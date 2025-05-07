package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	crErrors "github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "server stopped with error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("server exited cleanly")
}

func run() error {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}

	errCh := make(chan error, 1)
	defer close(errCh)
	go func() {
		fmt.Printf("server listening on %s\n", "0.0.0.0:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- crErrors.Wrapf(err, "listen error on %s", "0.0.0.0:8080")
		} else {
			errCh <- nil
		}
		fmt.Printf("server shutdown\n")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	select {
	case sig := <-quit:
		// 終了シグナルを受け取ったら graceful shutdown
		fmt.Printf("shutdown signal received: %v\n", sig)

		// context.Background() を親に、最大 5 秒後に自動的にキャンセルされる子コンテキスト ctx を生成
		ctx, cancelTimeout := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelTimeout()

		// 新規接続の受付停止、既存接続の完了待ち
		if err := srv.Shutdown(ctx); err != nil {
			return crErrors.Wrap(err, "shutdown error")
		}
		return nil
	case err := <-errCh:
		if err != nil {
			return crErrors.Wrap(err, "unknown error")
		}
		return nil
	}
}
