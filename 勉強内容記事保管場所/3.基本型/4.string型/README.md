# string型

文字列型である
文字列型はダブルクオーテーションでラッピングする

```go:main.go
    var s string = "Hello Golang"
```

### 数値を文字列として扱う

```go:main.go
    var s1 string = "300"
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)

```

### 複数行にわたる文字列を扱う方法

バッククオーテーションを使用する

```go:main.go
	fmt.Println(`test
	test
		text
	`)
```

### ダブルクオーテーションを文字列内で使用する方法（２通り）

以下の通り

```go:main.go
	fmt.Println("\"")
	fmt.Println(`"`)
```

### 文字列はそもそもバイトの集合体

以下のコードが参考になる

```go:main.go
	fmt.Println(s[0]) // 本来であれば[H]が取れるはず、結
```