package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:8080/posts"
	method := "POST"

	payload := strings.NewReader("{\n    \"title\": \"jonny\",\n    \"description\": \"any description\",\n    \"url\": \"any url\"\n}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Basic am9ubnk6bGVuaXA3NzA=")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
