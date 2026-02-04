package main

import (
	"context"
	"log"
	"net/http"

	"github.com/xyma8/go-shorter/db"
	"github.com/xyma8/go-shorter/internal/handler"
	"github.com/xyma8/go-shorter/internal/repository"
	"github.com/xyma8/go-shorter/internal/service"
	_ "modernc.org/sqlite"
)

func main() {
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	if err := db.Init(context.Background(), dbConn.DB); err != nil {
		log.Fatal(err)
	}

	urlRepo := repository.NewUrlRepository(dbConn.DB)
	urlService := service.NewUrlService(urlRepo)
	urlHandler := handler.NewUrlHandler(urlService)
	/*

		errQuery = db.QueryRow("SELECT name FROM urls").Scan(&test)
		if errQuery != nil {
			log.Fatal(errQuery)
		}
		fmt.Println(test)
	*/

	http.HandleFunc("/get_short", urlHandler.ShortUrl)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
