package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

type Message struct {
	Original_url string
}

func encodeXOR(id uint) (uint, error) {
	const MASK uint = 0x2A5B8D3F
	id = id ^ MASK
	return id, nil
}

func encodeBase62(id uint) (string, error) {
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	binary_result := []byte{0, 0, 0, 0, 0}

	if id >= uintPow(62, 5) {
		//return "", err
	}

	for i := 1; id > 0; i++ {
		mod := id % 62
		id = id / 62
		binary_result[len(binary_result)-i] = alphabet[mod]
	}
	return string(binary_result), nil
}

func shortURLHandler(w http.ResponseWriter, req *http.Request) {
	jsonDecoder := json.NewDecoder(req.Body)

	var bodyContent Message
	err := jsonDecoder.Decode(&bodyContent)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bodyContent.Original_url)
	var id uint = uintPow(62, 5) - 2
	fmt.Println(id)
	obfID, _ := encodeXOR(id)
	encodedID, _ := encodeBase62(obfID)
	fmt.Println(encodedID)

	io.WriteString(w, encodedID)
}

func uintPow(base, exponent uint) uint {
	result := uint(1)
	for i := uint(0); i < exponent; i++ {
		result *= base
	}

	return result
}

func main() {
	db, err := sql.Open("sqlite", "./go-shorter.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected sucessfully")

	//init_database(db)

	var test string
	errQuery := db.QueryRow(`INSERT INTO urls (original_url, short_url) VALUES ("asdasd", "sdad")`).Scan(&test)
	if errQuery != nil {
		log.Fatal(errQuery)
	}
	fmt.Println(test)

	errQuery = db.QueryRow("SELECT name FROM urls").Scan(&test)
	if errQuery != nil {
		log.Fatal(errQuery)
	}
	fmt.Println(test)

	http.HandleFunc("/hello", shortURLHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init_database(db *sql.DB) {
	var talbeUrls string

	err := db.QueryRow(`CREATE TABLE urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		original_url TEXT NOT NULL,
		short_url TEXT UNIQUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`).Scan(&talbeUrls)

	if err != nil {
		log.Fatal(err)
	}
	print("Table 'urls' created")
}
