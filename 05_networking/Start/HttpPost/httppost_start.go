package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func postRequestTest() {
	const httpbin = "https://httpbin.org/post"

	// TODO: POST operation using Post

	reqBody := strings.NewReader(`
		{
			"field1": "this is field1",
			"field2: "2"
				}
		`)

	resp, err := http.Post(httpbin, "application/json", reqBody)
	if err != nil {
		return
	}

	content, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Printf("%s\n", content)
	// TODO: POST operation using PostForm
	data := url.Values{}
	data.Add("field1", "test")
	data.Add("field2", "250")
	respa, err := http.PostForm(httpbin, data)
	content, _ = ioutil.ReadAll(respa.Body)
	defer respa.Body.Close()

	fmt.Printf("%s\n", content)
}

func main() {
	// Execute a POST
	postRequestTest()
}
