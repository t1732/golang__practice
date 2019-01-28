package basic

import (
	"fmt"
)

const (
	White = iota
	Black
	Chatora
	Kijitora
)

type Animal struct {
	Name string
	Age  int
}

type Cat struct {
	Animal
	Color int
}

func StructSample() {
	cat := Cat{}
	cat.Name = "Shiro"
	cat.Age = 5
	fmt.Println(cat)

	cat2 := Cat{Animal{"Sora", 9}, Kijitora}
	fmt.Println(cat2)

	cat3 := Cat{Animal: Animal{"Chizu", 12}, Color: Chatora}
	fmt.Println(cat3)

	cat4 := Cat{Animal: Animal{Name: "Tama"}}
	fmt.Println(cat4)

	// ポインタのアドレスでの初期化
	a := new(Animal)  // == &Animal{}と同意 初期化変数は指定できない
	a.Name = "Mike"
	a.Age = 5
	fmt.Println(a)

	a2 := &Animal{Name: "Mike", Age: 5}
	a2.Name = "Mike"
	a2.Age = 5
	fmt.Println(a2)
}
