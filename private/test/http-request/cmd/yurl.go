package main

import (
	"fmt"
	"http-request/utils"
	"io"
	"net/http"
)

var url string

func HttpRequest(url string) *http.Response {
	resp, err := http.Get(url)
	utils.HandleError(err)
	//defer resp.Body.Close()
	return resp
}

func BodyReader(body *http.Response) []byte {
	data, err := io.ReadAll(body.Body)
	utils.HandleError(err)
	return data
}

func main() {

	res := HttpRequest("http://localhost:8887/hello")
	fmt.Println(res)

	data := BodyReader(res)
	utils.Log(string(data))
}
