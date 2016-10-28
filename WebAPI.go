package main

import (
	"fmt"
	"net/http"
)

/* json define */
type Input struct {
	Id string `json:"id"`
	Token string `json:"token"`
}
type Output struct {
	Message string `json:"message"`
	Code int `json:"code"`
}


/* Controller */
func doRequest(rw http.ResponseWriter, request *http.Request) {

	method := request.Method
	request.ParseForm()
	values := request.Form
	
	fmt.Println(method)
	fmt.Println(values)
}

func main() {
	http.HandleFunc("/api/", doRequest)
	error := http.ListenAndServe(":8080", nil)

	if error != nil {
		fmt.Println(error)
		return
	}
}