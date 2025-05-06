package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getRequestTest() {
	const httpbin = "https://httpbin.org/get"

	// TODO: Perform a GET operation
	resp, err := http.Get(httpbin)
	if err != nil {
		panic(err)
	}

	// TODO: The caller is responsible for closing the response
	defer resp.Body.Close()

	// TODO: We can access parts of the response to get information:
	fmt.Println("status", resp.Status)
	fmt.Println("status code ", resp.StatusCode)
	fmt.Println("proto", resp.Proto)
	fmt.Println("content length", resp.ContentLength)

	// TODO: Use a String Builder to build the content from bytes
	var sb strings.Builder

	// TODO: Read the content and write it to the builder
	content, _ := ioutil.ReadAll(resp.Body)
	bytecount, _ := sb.Write(content)
	fmt.Println("bytecount", bytecount, sb.String())

	// TODO: Format the output
}

func main() {
	// Execute a GET request
	getRequestTest()
}
