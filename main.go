package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Message struct {
	Original_url string
}

func encodeBase62(id uint) (string, error) {
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	binary_result := []byte{0, 0, 0, 0, 0}

	for i := 1; id > 0; i++ {
		mod := id % 62
		id = id / 62
		binary_result[len(binary_result)-i] = alphabet[mod]
	}
	return string(binary_result), nil
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello")
	json_decoder := json.NewDecoder(req.Body)

	var body_content Message
	err := json_decoder.Decode(&body_content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(body_content.Original_url)

	encodedID, _ := encodeBase62(125)
	fmt.Println(encodedID)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
