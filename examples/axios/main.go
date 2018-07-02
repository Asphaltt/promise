package main

import (
	"fmt"
	"net/http"

	"github.com/Asphaltt/promise/axios"
)

func axiosGet(url string) {
	axios.Get(url).Then(func(v interface{}) (interface{}, error) {
		resp := v.(*http.Response)
		fmt.Println(resp.Status)
		return nil, nil
	}).Catch(func(e error) {
		fmt.Println(e)
	})
}

func main() {
	axiosGet("https://www.bing.com")
	axiosGet("https://www.bingggggg.com")
}
