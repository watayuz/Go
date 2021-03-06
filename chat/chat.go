package main

import (
	"bytes"
	"net/http"
	"fmt"
	"io/ioutil"

	// config package path:"config/config.go"
	"./config"
)

func main() {

	url_string := "https://chatbot-api.userlocal.jp/api/chat?"
	key_string := config.Key

	error := HttpPost(url_string, key_string, "こんにちわ")

	if error != nil {
		fmt.Println(error)
	}
}

func HttpPost(url_string, key_string, message_string string) error {
	jsonStr := []byte(`{"key":"` + key_string + `","message":"` + message_string + `"}`)

	req, error := http.NewRequest(
		"POST",
		url_string,
		bytes.NewBuffer(jsonStr),
		)
	if error != nil {
		panic(error)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, error := client.Do(req)
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return error
}