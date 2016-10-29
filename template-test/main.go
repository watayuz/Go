package main

import (
	"text/template"
	"fmt"
	"net/http"
	"time"
)

type MainPage struct {
	Title string
	Message string
	List []string
	Link string
}

func main() {
	page := MainPage{
		Title: "Hello Go",
		Message: "Go World",
		List: []string{
			"item1",
			"item2",
			"item3",
		},
		Link: "<a href=\"/\">hoge</a>",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " " + req.URL.Path)

		temp, error1 := template.ParseFiles("templates/main.html")
		if error1 != nil {
			panic(error1)
		}

		//render template
		error2 := temp.Execute(w, page)
		if error2 != nil {
			panic(error2)
		}
	})
	http.ListenAndServe(":8080", nil)
}