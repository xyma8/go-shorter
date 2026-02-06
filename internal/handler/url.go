package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/xyma8/go-shorter/internal/models"
	"github.com/xyma8/go-shorter/internal/service"
)

type UrlHandler struct {
	service *service.UrlService
}

type ShortUrl struct {
	Original_url string
}

type GetOrigUrl struct {
}

func NewUrlHandler(service *service.UrlService) *UrlHandler {
	return &UrlHandler{service: service}
}

func (h *UrlHandler) ShortUrl(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	jsonDecoder := json.NewDecoder(req.Body)
	var bodyContent ShortUrl
	err := jsonDecoder.Decode(&bodyContent)
	if err != nil {
		log.Fatal(err)
	}

	var creatingUrlModel models.UrlModel
	creatingUrlModel.Original_url = bodyContent.Original_url
	//creatingUrlModel.Short_url = ""

	result, err := h.service.ShortenUrl(ctx, &creatingUrlModel)
	if err != nil {
		io.WriteString(w, err.Error())
	}
	io.WriteString(w, result)
}

func (h *UrlHandler) GetOrigUrl(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	shortUrl := r.URL.Query().Get("short_url")

	result, err := h.service.GetOrigUrl(ctx, shortUrl)
	if err != nil {
		io.WriteString(w, err.Error())
	}
	io.WriteString(w, result)
}

func (h *UrlHandler) RedirectOrig(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	//fmt.Println(strings.TrimLeft(r.URL.Path, "/"))
}
