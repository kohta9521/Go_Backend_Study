package main

import "fmt"

// for
// 繰り返し処理

func main() {
	// 無限ループ
	// i := 0
	// for {
	// 	i ++
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println("Looooop");
	// }

	// 条件付きfor文
	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point ++
	// }

	// 古典的for文
	// for i := 0; i < 10; i ++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// 配列で使用するfor文
	arr := [3]int{1, 2, 3}
	// for i := 0; i < len(arr); i ++ {
	// 	fmt.Println(arr[i])
	// }

	for i, v := range arr {
		fmt.Println(i, v)
	}

	// スライスで使用するfor文
	sl := []string{"Python", "PHP", "Golang"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	// map
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
}