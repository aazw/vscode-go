# playgrounds/gin_graceful_shutdown

GinでGraceful Shutdownをどう実現するかの検証.

## Logs

```bash
$ go run -v main.go 
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> main.run.func1 (3 handlers)
server listening on 0.0.0.0:8080
^Cshutdown signal received: interrupt
server shutdown
server exited cleanly
```

* `^C`でCtrl+Cでプログラムを停止させるシグナルを送信している
* それをプログラムで明示的にキャッチし、終了処理を行って、終了
