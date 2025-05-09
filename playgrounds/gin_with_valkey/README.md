# playgrounds/gin_with_valkey

## 概要

### Valkey

Redis互換KVS.  
というか、Redisのライセンスが変わったことがきっかけで、Redisからフォークされたプロジェクト.

* https://valkey.io/
* https://hub.docker.com/r/valkey/valkey
* https://github.com/valkey-io/valkey

### Gin middleware for session management

* https://github.com/gin-contrib/sessions
    * Valkeyに明示的に対応はしてない
    * Redisは対応
    * RedisStoreで、Valkeyが使えるかどうか

## 検証

### Valkey起動 (Docker Compose)

```bash
$ docker compose up valkey
[+] Running 2/2
 ✔ Network gin_with_valkey_default     Created                                                                                       0.0s 
 ✔ Container gin_with_valkey-valkey-1  Created                                                                                       0.0s 
Attaching to valkey-1
valkey-1  | 1:M 09 May 2025 17:25:16.227 # WARNING Memory overcommit must be enabled! Without it, a background save or replication may fail under low memory condition. Being disabled, it can also cause failures without low memory condition, see https://github.com/jemalloc/jemalloc/issues/1328. To fix this issue add 'vm.overcommit_memory = 1' to /etc/sysctl.conf and then reboot or run the command 'sysctl vm.overcommit_memory=1' for this to take effect.
valkey-1  | 1:M 09 May 2025 17:25:16.228 * oO0OoO0OoO0Oo Valkey is starting oO0OoO0OoO0Oo
valkey-1  | 1:M 09 May 2025 17:25:16.228 * Valkey version=8.1.1, bits=64, commit=00000000, modified=0, pid=1, just started
valkey-1  | 1:M 09 May 2025 17:25:16.228 # Warning: no config file specified, using the default config. In order to specify a config file use valkey-server /path/to/valkey.conf
valkey-1  | 1:M 09 May 2025 17:25:16.229 * monotonic clock: POSIX clock_gettime
valkey-1  | 1:M 09 May 2025 17:25:16.229 * Running mode=standalone, port=6379.
valkey-1  | 1:M 09 May 2025 17:25:16.229 * Server initialized
valkey-1  | 1:M 09 May 2025 17:25:16.229 * Ready to accept connections tcp
```

### Gin Server 起動

```bash
$ go run -v main.go 
command-line-arguments
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (4 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on 0.0.0.0:8080
```

### cURLで検証

* cURL
    * `-c cookie.txt`: localhost から受け付けたCookieを cookie.txt として保存
    * `-b cookie.txt`: localhost に保存した cookie.txt を参照してCookieを送信
    * `-w '\n'`: 最後に改行を付ける
    * `--verbose`: 仔細を出す

```bash
root ➜ /workspaces/vscode-go/playgrounds/gin_with_valkey (main) $ curl -c cookie.txt -b cookie.txt localhost:8080 -w '\n' --verbose
* WARNING: failed to open cookie file "cookie.txt"
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.88.1
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
* Added cookie mysession="MTc0NjgxMTcyMHxOd3dBTkROU1VWQkpUelpEVVVGWVVFNVRVa3RFV0ZWUVR6TXpORXBKUVZCUFNqVk1Xa2xTU1RSQlJqVlZOREkwVlVOWlFVdFdRVkU9fMgkIXaiL09bTNQ3ExP64wlye3JWLMf46JLLrBo-eWFd" for domain localhost, path /, expire 1749403720
< Set-Cookie: mysession=MTc0NjgxMTcyMHxOd3dBTkROU1VWQkpUelpEVVVGWVVFNVRVa3RFV0ZWUVR6TXpORXBKUVZCUFNqVk1Xa2xTU1RSQlJqVlZOREkwVlVOWlFVdFdRVkU9fMgkIXaiL09bTNQ3ExP64wlye3JWLMf46JLLrBo-eWFd; Path=/; Expires=Sun, 08 Jun 2025 17:28:40 GMT; Max-Age=2592000
< Date: Fri, 09 May 2025 17:28:40 GMT
< Content-Length: 11
< 
* Connection #0 to host localhost left intact
{"count":0}
* WARNING: failed to open cookie file "cookie.txt"
root ➜ /workspaces/vscode-go/playgrounds/gin_with_valkey (main) $ curl -c cookie.txt -b cookie.txt localhost:8080 -w '\n' --verbose
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.88.1
> Accept: */*
> Cookie: mysession=MTc0NjgxMTcyMHxOd3dBTkROU1VWQkpUelpEVVVGWVVFNVRVa3RFV0ZWUVR6TXpORXBKUVZCUFNqVk1Xa2xTU1RSQlJqVlZOREkwVlVOWlFVdFdRVkU9fMgkIXaiL09bTNQ3ExP64wlye3JWLMf46JLLrBo-eWFd
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
* Replaced cookie mysession="MTc0NjgxMTcyMnxOd3dBTkROU1VWQkpUelpEVVVGWVVFNVRVa3RFV0ZWUVR6TXpORXBKUVZCUFNqVk1Xa2xTU1RSQlJqVlZOREkwVlVOWlFVdFdRVkU9fAuZ3Tg8FRLE1qJ7WGNPws7byk5PHDs9VXVxqPDQzl6N" for domain localhost, path /, expire 1749403722
< Set-Cookie: mysession=MTc0NjgxMTcyMnxOd3dBTkROU1VWQkpUelpEVVVGWVVFNVRVa3RFV0ZWUVR6TXpORXBKUVZCUFNqVk1Xa2xTU1RSQlJqVlZOREkwVlVOWlFVdFdRVkU9fAuZ3Tg8FRLE1qJ7WGNPws7byk5PHDs9VXVxqPDQzl6N; Path=/; Expires=Sun, 08 Jun 2025 17:28:42 GMT; Max-Age=2592000
< Date: Fri, 09 May 2025 17:28:42 GMT
< Content-Length: 11
< 
* Connection #0 to host localhost left intact
{"count":1}
root ➜ /workspaces/vscode-go/playgrounds/gin_with_valkey (main) $ curl -c cookie.txt -b cookie.txt localhost:8080 -w '\n' --verbose
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.88.1
> Accept: */*
> Cookie: mysession=MTc0NjgxMTcyMnxOd3dBTkROU1VWQkpUelpEVVVGWVVFNVRVa3RFV0ZWUVR6TXpORXBKUVZCUFNqVk1Xa2xTU1RSQlJqVlZOREkwVlVOWlFVdFdRVkU9fAuZ3Tg8FRLE1qJ7WGNPws7byk5PHDs9VXVxqPDQzl6N
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
* Replaced cookie mysession="MTc0NjgxMTcyM3xOd3dBTkROU1VWQkpUelpEVVVGWVVFNVRVa3RFV0ZWUVR6TXpORXBKUVZCUFNqVk1Xa2xTU1RSQlJqVlZOREkwVlVOWlFVdFdRVkU9fJ30tXiF86jmroc_wBFL3k9mczxnhmwo9giHi2AT71JJ" for domain localhost, path /, expire 1749403723
< Set-Cookie: mysession=MTc0NjgxMTcyM3xOd3dBTkROU1VWQkpUelpEVVVGWVVFNVRVa3RFV0ZWUVR6TXpORXBKUVZCUFNqVk1Xa2xTU1RSQlJqVlZOREkwVlVOWlFVdFdRVkU9fJ30tXiF86jmroc_wBFL3k9mczxnhmwo9giHi2AT71JJ; Path=/; Expires=Sun, 08 Jun 2025 17:28:43 GMT; Max-Age=2592000
< Date: Fri, 09 May 2025 17:28:43 GMT
< Content-Length: 11
< 
* Connection #0 to host localhost left intact
{"count":2}
root ➜ /workspaces/vscode-go/playgrounds/gin_with_valkey (main) $ curl -c cookie.txt -b cookie.txt localhost:8080 -w '\n' --verbose
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.88.1
> Accept: */*
> Cookie: mysession=MTc0NjgxMTcyM3xOd3dBTkROU1VWQkpUelpEVVVGWVVFNVRVa3RFV0ZWUVR6TXpORXBKUVZCUFNqVk1Xa2xTU1RSQlJqVlZOREkwVlVOWlFVdFdRVkU9fJ30tXiF86jmroc_wBFL3k9mczxnhmwo9giHi2AT71JJ
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
* Replaced cookie mysession="MTc0NjgxMTcyNHxOd3dBTkROU1VWQkpUelpEVVVGWVVFNVRVa3RFV0ZWUVR6TXpORXBKUVZCUFNqVk1Xa2xTU1RSQlJqVlZOREkwVlVOWlFVdFdRVkU9fL3f3pnEeqYXWIKonD6x0_e8WhkWRaeNZiZ2rxlBajMl" for domain localhost, path /, expire 1749403724
< Set-Cookie: mysession=MTc0NjgxMTcyNHxOd3dBTkROU1VWQkpUelpEVVVGWVVFNVRVa3RFV0ZWUVR6TXpORXBKUVZCUFNqVk1Xa2xTU1RSQlJqVlZOREkwVlVOWlFVdFdRVkU9fL3f3pnEeqYXWIKonD6x0_e8WhkWRaeNZiZ2rxlBajMl; Path=/; Expires=Sun, 08 Jun 2025 17:28:44 GMT; Max-Age=2592000
< Date: Fri, 09 May 2025 17:28:44 GMT
< Content-Length: 11
< 
* Connection #0 to host localhost left intact
{"count":3}
root ➜ /workspaces/vscode-go/playgrounds/gin_with_valkey $ curl -c cookie.txt -b cookie.txt localhost:8080 -w '\n' --verbose
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.88.1
> Accept: */*
> Cookie: mysession=MTc0NjgxMTcyNHxOd3dBTkROU1VWQkpUelpEVVVGWVVFNVRVa3RFV0ZWUVR6TXpORXBKUVZCUFNqVk1Xa2xTU1RSQlJqVlZOREkwVlVOWlFVdFdRVkU9fL3f3pnEeqYXWIKonD6x0_e8WhkWRaeNZiZ2rxlBajMl
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
* Replaced cookie mysession="MTc0NjgxMTcyNHxOd3dBTkROU1VWQkpUelpEVVVGWVVFNVRVa3RFV0ZWUVR6TXpORXBKUVZCUFNqVk1Xa2xTU1RSQlJqVlZOREkwVlVOWlFVdFdRVkU9fL3f3pnEeqYXWIKonD6x0_e8WhkWRaeNZiZ2rxlBajMl" for domain localhost, path /, expire 1749403724
< Set-Cookie: mysession=MTc0NjgxMTcyNHxOd3dBTkROU1VWQkpUelpEVVVGWVVFNVRVa3RFV0ZWUVR6TXpORXBKUVZCUFNqVk1Xa2xTU1RSQlJqVlZOREkwVlVOWlFVdFdRVkU9fL3f3pnEeqYXWIKonD6x0_e8WhkWRaeNZiZ2rxlBajMl; Path=/; Expires=Sun, 08 Jun 2025 17:28:44 GMT; Max-Age=2592000
< Date: Fri, 09 May 2025 17:28:44 GMT
< Content-Length: 11
< 
* Connection #0 to host localhost left intact
{"count":4}
```
