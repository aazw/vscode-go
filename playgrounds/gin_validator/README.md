# playgrounds/gin_validator

## 目的

### 1. Ginのバリデーターのタグは？

Ginは [github.com/go-playground/validator](https://github.com/go-playground/validator) を内部的に利用しているそう.

* 直接 github.com/go-playground/validator を使う場合は `validate` タグを利用
* Ginから github.com/go-playground/validator を使う場合は、 `binding` タグを利用

本当か？

### 2. Ginの BindJSON や ShouldBindJSON の違いは？

BindJSONとか、ShouldBindJSONとか、MustBindJSON とかあるけど、その違いとは？

### 3. Ginのバリデーター使えば、バリデーション結果の個別のエラーメッセージってよしなにしてくれるのか？

以下みたいな、外部のユーザーフレンドリー？、セーフ？なメッセージを出せるか？

```json
{
    "errors": [
        {
            "param": "name",
            "message": "name is not empty, and the length is greater than or equal to 10, and less then or equal to 100."
        },
        ...
    ]
}
```

## 検証結果

```bash
$ go test -v *.go
=== RUN   TestBindHandler_NoTest_OutputOnly
POST /users/bind_json_by_binding_tag_with_json_resp HTTP/1.1
Host: 127.0.0.1:35569
Content-Type: application/json

{}

HTTP/1.1 400 Bad Request
Content-Length: 228
Content-Type: text/plain; charset=utf-8
Date: Fri, 09 May 2025 16:09:16 GMT

{"error":"Key: 'CreateUserRequestWithBindingTag.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'CreateUserRequestWithBindingTag.Email' Error:Field validation for 'Email' failed on the 'required' tag"}

==========
POST /users/bind_json_by_binding_tag_with_text_resp HTTP/1.1
Host: 127.0.0.1:37987
Content-Type: application/json

{}

HTTP/1.1 400 Bad Request
Content-Length: 215
Content-Type: text/plain; charset=utf-8
Date: Fri, 09 May 2025 16:09:16 GMT

Key: 'CreateUserRequestWithBindingTag.Name' Error:Field validation for 'Name' failed on the 'required' tag
Key: 'CreateUserRequestWithBindingTag.Email' Error:Field validation for 'Email' failed on the 'required' tag

==========
POST /users/bind_json_by_binding_tag_without_resp HTTP/1.1
Host: 127.0.0.1:38981
Content-Type: application/json

{}

HTTP/1.1 400 Bad Request
Content-Length: 0
Date: Fri, 09 May 2025 16:09:16 GMT



==========
POST /users/bind_json_by_validate_tag_with_json_resp HTTP/1.1
Host: 127.0.0.1:43945
Content-Type: application/json

{}

HTTP/1.1 201 Created
Content-Length: 15
Content-Type: application/json; charset=utf-8
Date: Fri, 09 May 2025 16:09:16 GMT

{"status":"ok"}

==========
POST /users/bind_json_by_validate_tag_with_text_resp HTTP/1.1
Host: 127.0.0.1:33675
Content-Type: application/json

{}

HTTP/1.1 201 Created
Content-Length: 2
Content-Type: text/plain; charset=utf-8
Date: Fri, 09 May 2025 16:09:16 GMT

ok

==========
POST /users/bind_json_by_validate_tag_without_resp HTTP/1.1
Host: 127.0.0.1:34075
Content-Type: application/json

{}

HTTP/1.1 201 Created
Content-Length: 0
Date: Fri, 09 May 2025 16:09:16 GMT



==========
POST /users/should_bind_json_by_binding_tag_with_json_resp HTTP/1.1
Host: 127.0.0.1:39933
Content-Type: application/json

{}

HTTP/1.1 400 Bad Request
Content-Length: 228
Content-Type: application/json; charset=utf-8
Date: Fri, 09 May 2025 16:09:16 GMT

{"error":"Key: 'CreateUserRequestWithBindingTag.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'CreateUserRequestWithBindingTag.Email' Error:Field validation for 'Email' failed on the 'required' tag"}

==========
POST /users/should_bind_json_by_binding_tag_with_text_resp HTTP/1.1
Host: 127.0.0.1:39291
Content-Type: application/json

{}

HTTP/1.1 400 Bad Request
Content-Length: 215
Content-Type: text/plain; charset=utf-8
Date: Fri, 09 May 2025 16:09:16 GMT

Key: 'CreateUserRequestWithBindingTag.Name' Error:Field validation for 'Name' failed on the 'required' tag
Key: 'CreateUserRequestWithBindingTag.Email' Error:Field validation for 'Email' failed on the 'required' tag

==========
POST /users/should_bind_json_by_binding_tag_without_resp HTTP/1.1
Host: 127.0.0.1:36629
Content-Type: application/json

{}

HTTP/1.1 200 OK
Content-Length: 0
Date: Fri, 09 May 2025 16:09:16 GMT



==========
POST /users/should_bind_json_by_validate_tag_with_json_resp HTTP/1.1
Host: 127.0.0.1:35147
Content-Type: application/json

{}

HTTP/1.1 201 Created
Content-Length: 15
Content-Type: application/json; charset=utf-8
Date: Fri, 09 May 2025 16:09:16 GMT

{"status":"ok"}

==========
POST /users/should_bind_json_by_validate_tag_with_text_resp HTTP/1.1
Host: 127.0.0.1:36463
Content-Type: application/json

{}

HTTP/1.1 201 Created
Content-Length: 2
Content-Type: text/plain; charset=utf-8
Date: Fri, 09 May 2025 16:09:16 GMT

ok

==========
POST /users/should_bind_json_by_validate_tag_without_resp HTTP/1.1
Host: 127.0.0.1:43979
Content-Type: application/json

{}

HTTP/1.1 201 Created
Content-Length: 0
Date: Fri, 09 May 2025 16:09:16 GMT


--- PASS: TestBindHandler_NoTest_OutputOnly (0.01s)
PASS
ok      command-line-arguments  0.008s
```

### 1. バリデーションのタグ

* `validate` ... 効かない
* `binding` ... 効く

→ `binding` を使う

### 2. BindJSON vs ShouldBindJSON

バリデーションエラーを含む、エラー時の挙動について.

| メソッド       | ステータスコード                     | レスポンスヘッダ                            | レスポンスボディ |
|----------------|--------------------------------------|---------------------------------------------|------------------|
| BindJSON       | BindJSON内で`400`に設定<br/>変更不可 | BindJSON内で`text/plain`に設定<br/>変更不可 | 未設定           |
| ShouldBindJSON | 未設定                               | 未設定                                      | 未設定           |

* `/users/bind_json_by_binding_tag_with_json_resp`
    * JSON型を返しているのに、 `Content-Type`が`text/plain`
        * → BindJSONではレスポンスヘッダの上書きはできないことを意味する
* `/users/bind_json_by_binding_tag_without_resp`
    * デフォルトではレスポンスボディに値は入っていないことを意味する
* `/users/bind_json_by_validate_tag_with_json_resp`  
  `/users/bind_json_by_validate_tag_with_text_resp`  
  `/users/bind_json_by_validate_tag_without_resp`
    * go-playground/validatorのデフォルトのタグの`validate`タグは効かないことを意味する
    * ginはgo-playground/validatorのタグを変更している
        * → `binding`
* `/users/should_bind_json_by_binding_tag_with_json_resp`
    * ShouldBindJSONではレスポンスヘッダの上書きができることを意味する
* `/users/should_bind_json_by_binding_tag_without_resp`
    * 200を返す
        * 何もしなければステータスコードのデフォルト値の200を返す
        * ShouldBindJSONはBindJSONと違い、ステータスコードも設定していないことを意味する
    * BindJSON同様、デフォルトではレスポンスボディに値は入っていないことを意味する

# 3. エラーの内容

* BindJSON、ShouldBindJSONともに、バリデーションエラーの内容をそのまま返すと、システムの実装上のエラーメッセージとなる
    * API仕様上の`name`ではなく、`CreateUserRequestWithBindingTag.Name`と表記されている
    * `... on the 'required' tag`のような、実装上のタグが表記されている
* ある意味情報漏洩である
* 実装上のタグ情報を、外部向けのAPIのシステムメッセージに書き換える必要がある
    * `gte=x`タグ → `... is greater than or equal to x` みたいなメッセージに変える
    * `CreateUserRequestWithBindingTag.Name`のようなGoフィールド名 → `name` みたいなAPI仕様上のパラメータ名に変える
