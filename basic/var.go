package basic

import "fmt"

func VarSample() {
	// 宣言した後、値を代入パターン
	var msg string
	msg = "宣言してから代入"
	fmt.Printf("msg: %s\n", msg)

	// 宣言と代入を一緒にするパターン
	var msg2 string = "Hello World"
	fmt.Printf("msg2: %s\n", msg2)

	// 宣言と代入を一緒にするパターン (型省略可能)
	var msg3 = "Hello Hello"
	fmt.Printf("msg3: %s\n", msg3)

	// 宣言と代入を一緒にするパターン (var省略)
	msg4 := "Super Hello"
	fmt.Printf("msg4: %s\n", msg4)

	// 複数の同じ型の変数を同時に定義
	var a, b int
	a, b = 10, 15
	fmt.Printf("a: %d, b:%d\n", a, b)

	// 複数の同じ型の変数を同時に定義 (var省略)
	a2, b2 := 10, 14
	fmt.Printf("a2: %d, b2:%d\n", a2, b2)

	// 複数の型違いの変数を同時に定義
	var (
		c int
		d string
	)
	c = 20
	d = "hoge"
	fmt.Printf("c: %d, d:%s\n", c, d)

	// 複数の型違いの変数を同時に定義 (型省略パターン)
	var (
		c2 = 20
		d2 = "hoge"
	)
	fmt.Printf("c2: %d, d2:%s\n", c2, d2)

	var aa int
	var bb float64
	var cc string
	var dd bool
	fmt.Printf("aa: %d, bb:%f, cc:%s, dd:%t\n", aa, bb, cc, dd)

	// 連番定数
	const (
		sun = iota // 0
		mon // 1
		tue // 2
	)
	fmt.Printf("sun:%d, mon:%d, tue:%d\n", sun, mon, tue)
}
