package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/xyma8/go-shorter/internal/models"
	"github.com/xyma8/go-shorter/internal/service"
)

type UrlHandler struct {
	service *service.UrlService
}

type Message struct {
	Original_url string
}

func NewUrlHandler(service *service.UrlService) *UrlHandler {
	return &UrlHandler{service: service}
}

func (h *UrlHandler) ShortUrl(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	jsonDecoder := json.NewDecoder(req.Body)
	var bodyContent Message
	err := jsonDecoder.Decode(&bodyContent)
	if err != nil {
		log.Fatal(err)
	}
	var creatingUrlModel models.UrlModel
	creatingUrlModel.Original_url = bodyContent.Original_url
	creatingUrlModel.Short_url = ""

	h.service.ShortUrl(ctx, &creatingUrlModel)
}
