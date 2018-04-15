package main

import "fmt"
import "net/url"

func main() {
	var urlString = "http://developer.com:80/hello?name=john wick&age=27"
	var u, err = url.Parse(urlString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("url: %s", urlString)

	fmt.Printf("protocols: %s\n", u.Scheme) // http
	fmt.Printf("host: %s\n", u.Host)        // developer.com:80
	fmt.Printf("path: %s\n", u.Path)        // hello

	var name = u.Query()["name"][0] // john wick
	var age = u.Query()["age"][0]   // 27

	fmt.Printf("name: %s, age: %s", name, age)

}
