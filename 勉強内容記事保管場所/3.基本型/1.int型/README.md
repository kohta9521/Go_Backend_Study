# int型について

```go:main.go
    func main() {
	var i int = 100
	fmt.Println(i)
    }
```

上記のint指定は環境依存
32ビットなど、、、

## 明示的な指定

```go:main.go
    var i2 int64 = 267
```

## 計算でのありがちなミス

```go:main.go
    	fmt.Println(i + i2)
        // 上記のコードはエラーとなる（厳密に方が違うため）
    	fmt.Println(i + int(i2))
```

## 型変換（重要）

まずは現在の変数の型を調べる方法について
- 書式指定子を使用する

現時点の知識だと%Tというものが数値型を表すものという認識で大丈夫

```go:main.go
    	fmt.Printf("%T\n", i2)
	fmt.Printf("%T\n", int32(i2))
```

作成日：2023/12/13