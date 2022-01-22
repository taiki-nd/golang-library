package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	/*
		resp, _ := http.Get("https://trade-note.jp/")
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	*/

	/*
		base, err := url.Parse("https://t    rade-note.jp/")
		fmt.Println(base, err)
	*/

	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match", `W/wxxyz`)

	q := req.URL.Query()
	fmt.Println(q)
	q.Add("c", "3")
	fmt.Println(q)          // =>map[a:[1] b:[2] c:[3]]
	fmt.Println(q.Encode()) //=>a=1&b=2&c=3

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
