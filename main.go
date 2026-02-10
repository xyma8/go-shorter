package main

import (
	"context"
	"log"
	"net/http"

	"github.com/xyma8/go-shorter/db"
	"github.com/xyma8/go-shorter/internal/handler"
	"github.com/xyma8/go-shorter/internal/repository"
	"github.com/xyma8/go-shorter/internal/service"
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

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/ui", http.StripPrefix("/ui", fs))
	http.Handle("/ui/", http.StripPrefix("/ui", fs))
	http.HandleFunc("/api/get_short", urlHandler.ShortUrl)
	http.HandleFunc("/api/get_orig", urlHandler.GetOrigUrl)
	http.HandleFunc("/", urlHandler.ShortRedirect)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
