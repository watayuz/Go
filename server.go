package main

import (
	"net/http"
	"log"
	"html/template"
)

type WeatherData struct {
	Name string 'json:"name"'
	Main struct {
		Kelvin float64 'json:temp'
		} 'json:"Main"'
	}

	func handler(w http.ResponseWriter, r *http.Request) {
		temp := template.New("fieldname example")
		temp, _ = temp.Parse("hello {{.UserName}}")
		p := IndexPage{UserName : "in user name"}
		temp.Execute(w, p)
	}

	func main() {
		http.HandleFunc("/", handler)
		http>HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request){
			
		})
		log.Printf("Start Go Server localhost:8080")

		http.ListenAndServe(":8080", nil)
	}