# playgrounds/scs_with_redis

Gin用の`gin-contrib/sessions`ではなく、もっと汎用のセッション管理はないか、と調査した.  
バックエンドはRedis/Valkeyを想定.

その中で、ここでは`alexedwards/scs`を動作検証.

## セッション管理パッケージ

* https://github.com/gin-contrib/sessions
    * Gin用

* https://github.com/gorilla/sessions
    * gorillaなパッケージ

* https://github.com/alexedwards/scs
    * https://github.com/alexedwards/scs/tree/master/redisstore

## 検証結果

```bash
docker compose up valkey
```

```bash
go run -v main.go
```
