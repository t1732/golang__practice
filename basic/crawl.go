package basic

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Crawl() {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	Xprint(resp)
}

func Xprint(resp *http.Response) {
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(b))
	}
}
