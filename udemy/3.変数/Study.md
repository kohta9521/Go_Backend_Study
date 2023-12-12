# Go言語の変数について

## 明示的な定義

```go:main.go
    // var 変数名 型 = 値
    var i int = 100
```

現状型についてはまだ深く学習をできていないためここではサンプルコードとして認識

- 数値型（int）
```go:main.go
    var i int = 100
```

- 文字列型（string）
```go:main.go
    var s string = "Hello Golang"
```

- 真偽ちの定義（boolean）
```go:main.go
    var t, f = true, false
```

- 複数の型の定義
```golang:main.go
    var (
        i2 int = 200
        s2 string = "Golang"
    )
```

### 初期化をしない場合

```go:main.go
    var i3 int
    var s3 string
    fmt.Println(i3, s3)

    // 上記の出力は初期値が出力
    // この場合 0
```

### 変数の上書き

```go:main.go
	i = 150
	fmt.Println(i)
```


## 暗黙的な定義

```go:main.go
	// 変数名 := 値
	i4 := 400
	fmt.Println(i4)
```


## まとめメモ（注意点含む）
- 暗黙的な定義は関数内でしか使用不可能
- 明示的な定義と暗黙的な定義は混在可能
- しかし普段は明示的な定義を使用する方が好ましい
    - 理由としては他の人がコードを見た時に分かりやすくなるから

- 関数内での変数は関数内でのみ使用可能
- 現時点での勉強だとGo内では定義した変数は必ず使用しなくてはならない

作成日：2023/12/13