# playgrounds/openapi_codegen

OpenAPIのSpecからGoのコードを生成して、それを使って実装する方法.

## openapi

* <https://www.openapis.org>
* <https://swagger.io>

## openapi-generator

* <https://openapi-generator.tech>
* <https://github.com/OpenAPITools/openapi-generator>
* <https://github.com/OpenAPITools/openapi-generator-cli>
* <https://github.com/OpenAPITools/openapi-generator-pip>
* <https://hub.docker.com/r/openapitools/openapi-generator-cli>
* <https://hub.docker.com/r/openapitools/openapi-generator-online>

## oapi-codegen

* <https://github.com/oapi-codegen/oapi-codegen>
* 良い実装例
  * <https://github.com/oapi-codegen/oapi-codegen/tree/main/examples/minimal-server/gin/api>

```bash
go generate ./...
```

Ginは github.com/go-playground/validator を内部的に利用.

* 直接 github.com/go-playground/validator を使う場合は `validate` タグを利用
* Ginから github.com/go-playground/validator を使う場合は、 `binding` タグを利用

OpenAPIのSpecに、以下のように `x-oapi-codegen-extra-tags` の下に、生成したGoの構造体のフィールドに埋め込むタグを記載することができる.  
ここに `binding` キーで github.com/go-playground/validator の値を書くことで、 github.com/go-playground/validator でバリデーションを利かすことができる.

```yaml
      ...

      properties:
        id:
          type: string
          format: uuid
          description: Unique identifier for the user (UUIDv7)
          pattern: "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-7[0-9a-fA-F]{3}-[89ABab][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
          minLength: 36
          maxLength: 36
          x-oapi-codegen-extra-tags:
            binding: "required,uuid"

      ...
```

### Ginでの実行例

```bash
$ go run -v main.go 
run gin server
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /users                    --> github.com/aazw/vscode-go/playgrounds/openapi_codegen/pkg/api/openapi.(*ServerInterfaceWrapper).ListUsers-fm (3 handlers)
[GIN-debug] POST   /users                    --> github.com/aazw/vscode-go/playgrounds/openapi_codegen/pkg/api/openapi.(*ServerInterfaceWrapper).CreateUser-fm (3 handlers)
[GIN-debug] DELETE /users/:user_id           --> github.com/aazw/vscode-go/playgrounds/openapi_codegen/pkg/api/openapi.(*ServerInterfaceWrapper).DeleteUserById-fm (3 handlers)
[GIN-debug] GET    /users/:user_id           --> github.com/aazw/vscode-go/playgrounds/openapi_codegen/pkg/api/openapi.(*ServerInterfaceWrapper).GetUserById-fm (3 handlers)
[GIN-debug] PATCH  /users/:user_id           --> github.com/aazw/vscode-go/playgrounds/openapi_codegen/pkg/api/openapi.(*ServerInterfaceWrapper).UpdateUserById-fm (3 handlers)
```

### Strict Serverでの実行例

```bash
$ go run -v main.go -strict
run strict server
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /users                    --> github.com/aazw/vscode-go/playgrounds/openapi_codegen/pkg/api/openapi.(*ServerInterfaceWrapper).ListUsers-fm (3 handlers)
[GIN-debug] POST   /users                    --> github.com/aazw/vscode-go/playgrounds/openapi_codegen/pkg/api/openapi.(*ServerInterfaceWrapper).CreateUser-fm (3 handlers)
[GIN-debug] DELETE /users/:user_id           --> github.com/aazw/vscode-go/playgrounds/openapi_codegen/pkg/api/openapi.(*ServerInterfaceWrapper).DeleteUserById-fm (3 handlers)
[GIN-debug] GET    /users/:user_id           --> github.com/aazw/vscode-go/playgrounds/openapi_codegen/pkg/api/openapi.(*ServerInterfaceWrapper).GetUserById-fm (3 handlers)
[GIN-debug] PATCH  /users/:user_id           --> github.com/aazw/vscode-go/playgrounds/openapi_codegen/pkg/api/openapi.(*ServerInterfaceWrapper).UpdateUserById-fm (3 handlers)
```
