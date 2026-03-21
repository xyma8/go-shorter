package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/xyma8/go-shorter/db"
	"github.com/xyma8/go-shorter/internal/handler"
	"github.com/xyma8/go-shorter/internal/postgresrepo"
	"github.com/xyma8/go-shorter/internal/service"
	"github.com/xyma8/go-shorter/internal/sqliterepo"
)

func main() {
	dbType := os.ExpandEnv("DB_TYPE")
	var database *db.DB
	if dbType == "postgres" {
		postgres := db.NewPostgres()
		database = db.NewDB(postgres)
	} else {
		sqlite := db.NewSqlite()
		database = db.NewDB(sqlite)
	}

	dbConn, err := database.Database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	if err := database.Database.InitDB(context.Background(), dbConn); err != nil {
		log.Fatal(err)
	}

	var urlService *service.UrlService
	if dbType == "postgres" {
		urlRepo := postgresrepo.NewUrlRepository(dbConn)
		urlService = service.NewUrlService(urlRepo)
	} else {
		urlRepo := sqliterepo.NewUrlRepository(dbConn)
		urlService = service.NewUrlService(urlRepo)
	}

	urlHandler := handler.NewUrlHandler(urlService)

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/ui", http.StripPrefix("/ui", fs))
	http.Handle("/ui/", http.StripPrefix("/ui", fs))
	http.HandleFunc("/api/get_short", urlHandler.ShortUrl)
	http.HandleFunc("/api/get_orig", urlHandler.GetOrigUrl)
	http.HandleFunc("/", urlHandler.ShortRedirect)

	addr := ":8007"
	log.Printf("go shorter started on %s ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
