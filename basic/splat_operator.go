package basic

import (
	"fmt"
)

func SplatOperator() {
	process("hoge", "fuga", "foo")
}

func process(args ...string) {
	fmt.Println(args)
}
