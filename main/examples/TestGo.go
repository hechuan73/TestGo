package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main()  {

	//var array =  [] int {1, 2, 3}
	//
	//fmt.Printf("%d\n", array[0])
	//modify(array)
	//fmt.Printf("%d", array[0])
	//fmt.Println("end")

	//var ptr *int
	//
	//fmt.Println(ptr)
	//fmt.Printf("ptr 的值为 : %x\n", ptr)


	requestBody := make(url.Values)
	requestBody["name"] = []string{"He Chuan"}
	httpPost("http://localhost:8888/hello", requestBody)
}

func modify(array [] int)  {
	array[0]++
}

func httpPost(url string, requestBody url.Values)  {
	request, err := http.PostForm(url, requestBody)

	if err != nil {
		fmt.Println(err)
	}

	defer request.Body.Close()
	responseBody, err := ioutil.ReadAll(request.Body)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(responseBody))
}