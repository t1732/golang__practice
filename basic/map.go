package basic

import (
	"fmt"
)

func Map() {
	m := map[string]string{}
	fmt.Println(m)

	// 指定要素文のメモリを確保
	m1 := make(map[string]string, 10)
	fmt.Println(m1)

	m2 := map[string]string{
		"hoge": "fuga",
		"moga": "fuge",
	}
	fmt.Println(m2)

	m3 := make(map[string]string)
	m3["hoge"] = "fuga"
	m3["moga"] = "fuge"
	fmt.Println(m3)
}
