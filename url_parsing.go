package main

import (
	"fmt"
	"net/url"
)

func main() {
	// URL Parsing

	// URL Parsing
	urlString := "http://kalipare.com:80/hello?name=john wick&age=27"
	u, err := url.Parse(urlString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme) //http
	fmt.Printf("host: %s\n", u.Host)       // kalipare.com:80
	fmt.Printf("path: %s\n", u.Path)       // /helo

	name := u.Query()["name"][0]
	age := u.Query()["age"][0]
	fmt.Printf("name: %s, age: %s\n", name, age)
}
