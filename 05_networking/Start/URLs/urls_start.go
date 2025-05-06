package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Define a URL
	s := "https://www.example.com:8000/user?username=joemarini"

	// // TODO: Use the Parse function to parse the URL content

	// result, _ := url.Parse(s)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	// // TODO: Extract the query components into a Values struct

	// vals := result.Query()
	// fmt.Println(vals["username"])

	// TODO: create a URL from components

	newUrl := &url.URL{
		Scheme:   "https",
		Host:     "ww.exmaple.com",
		Path:     "/args",
		RawQuery: "x=1&y=2",
	}

	s = newUrl.String()
	fmt.Println(s)

	newUrl.Host = "test.com"
	fmt.Println(newUrl.String())

	// TODO: create a new Values struct and encode it as a query string

	newVals := url.Values{}
	newVals.Add("x", "10")

	newVals.Add("z", "some string")

	newUrl.RawQuery = newVals.Encode()

	s = newUrl.String()
	fmt.Println(s)
}
