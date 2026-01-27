# playgrounds/go_generate

## go generate とは

ソース中の特別なコメント //go:generate を手がかりに、**任意のコマンド**を実行してコードやリソースを生成する仕組み.  
ビルドやテスト時には自動実行されず、手動で go generate を叩く必要がある.

しばしstringerが例にあがるので、stringerとセットかと思われるが、そうではない.

## stringerでお試し

stringerはパッケージとしてではなく、コマンドとしてインストールする必要がある.

* <https://pkg.go.dev/golang.org/x/tools/cmd/stringer>
* <https://cs.opensource.google/go/x/tools>

```bash
# failed
go get golang.org/x/tools/cmd/stringer@latest

# ok
go install golang.org/x/tools/cmd/stringer@latest
```

stringerをパッケージとしてインストールした場合など、stringerがPATHにないと以下のようなエラーになる.

```bash
$ go generate .
status.go:3: running "stringer": exec: "stringer": executable file not found in $PATH
```

stringerがPATHにあると成功する. 成功時は特にエラーも成功ログも何も表示されない.

```bash
go generate .
```
