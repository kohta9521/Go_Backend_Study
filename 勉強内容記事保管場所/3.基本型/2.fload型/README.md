# float型

## 浮動小数点型

sample
```go:main.go
    var fl64 float64 = 2.4
    fmt.Println(fl64)
```

## 暗黙的な定義
```go:main.go
    fl := 3.2
```

注意点としてint型と違い環境依存ではない
そのため以下のような計算が可能

```go:main.go
    fmt.Println(fl64 + fl)
```

実際に型を調べるとどちらともfload64である

```go:main.go
    fmt.Printf("%T, %T\n", fl64, fl)
```

## 型指定について
以下のように行う
実際のところfloat32はあまり使用しないが参考程度に、、

```go:.main.go
    var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)
```

## 型変換

```go:main.go
    fmt.Printf("%T\n", float64(fl32))
```


## 演算が不能な特殊な型？

```go:main.go
	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf) // +Inf

	ninf := -1.0 / zero
	fmt.Println(ninf) // -Inf

	nan := zero / zero
	fmt.Println(nan) // NaN(Not a Number)

```

作成日：2023/12/13