package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		panic(fmt.Sprintf("Error: %v", res.Status))
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", body[:100])
	defer res.Body.Close()
}
