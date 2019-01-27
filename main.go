package main

import (
	"fmt"

	"github.com/t1732/go-practice/basic"
)

func main() {
	basic.VarSample()

	t, err := basic.TimeParse("2020-04-01")
	if err == nil {
		fmt.Println(t)
	}
	fmt.Println(basic.AdvanceOneHourNow())

	basic.Crawl()
}
