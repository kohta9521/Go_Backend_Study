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
	for i := 0; i < 10; i ++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}