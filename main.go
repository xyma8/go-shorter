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

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello")
	//fmt.Println(*req)
	//body := req.Body
	//body, _ := req.GetBody()
	fmt.Println(req.Body)
	json_decoder := json.NewDecoder(req.Body)

	var body_content Message
	err := json_decoder.Decode(&body_content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(body_content.Original_url)
	//longBuf := make([]byte, 1)
	//fmt.Println(body.Read(longBuf))

	//if _, err := body.Read(longBuf); err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%s\n", longBuf)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
