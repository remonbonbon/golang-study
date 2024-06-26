GolangでWebアプリ作る
--------------------------

#### VS Codeでフォーマット

```json
// settings.json
{
  "editor.formatOnSave": true,
  "[go]": {
    "editor.defaultFormatter": "golang.go"
  }
}
```

#### ソース分割

参考: <https://zenn.dev/nobonobo/articles/4fb018a24f9ee9>

以下のようなディレクトリ構造にする。

- <プロジェクトルートディレクトリ>/
  - go.mod
  - main.go
  - foo/
    - bar/
      - baz.go (package bar)

ルートにある`go.mod`に書かれた**モジュール名からのパス**をimportに書く。

```go
// go.mod
module <モジュール名>

go 1.22
```

```go
import "<モジュール名>/foo/bar"
```


#### ルーティング

<https://github.com/go-chi/chi>を使う。


#### logger

slogをカスタマイズして使う。


#### 設定ファイル

参考:
- <https://zenn.dev/heartrails/articles/7899565fc04673>
- <https://github.com/mrk21/go-web-config-sample>
