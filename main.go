package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

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

	filesDir, is := os.LookupEnv("UI_FILES_DIS")
	if !is {
		filesDir = "./public"
	}
	fs := http.FileServer(http.Dir(filesDir))

	http.Handle("/ui", http.StripPrefix("/ui", fs))
	http.Handle("/ui/", http.StripPrefix("/ui", fs))
	http.HandleFunc("/api/get_short", urlHandler.ShortUrl)
	http.HandleFunc("/api/get_orig", urlHandler.GetOrigUrl)
	http.HandleFunc("/", urlHandler.ShortRedirect)

	appPort, is := os.LookupEnv("APP_PORT")
	if !is {
		appPort = "8080"
	}

	log.Printf("go shorter started on :%s ", appPort)
	log.Fatal(http.ListenAndServe(strings.Join([]string{appPort}, ""), nil))
}
