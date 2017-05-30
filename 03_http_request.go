package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.jp/")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	display(resp)
}

func display(resp *http.Response) {
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(b))
	}
}
