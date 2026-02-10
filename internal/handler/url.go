package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/xyma8/go-shorter/internal/models"
	"github.com/xyma8/go-shorter/internal/service"
)

type UrlHandler struct {
	service *service.UrlService
}

type ShortUrl struct {
	Original_url string `json:"original_url"`
}

type GetOrigUrl struct {
}

func NewUrlHandler(service *service.UrlService) *UrlHandler {
	return &UrlHandler{service: service}
}

func (h *UrlHandler) ShortUrl(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	jsonDecoder := json.NewDecoder(req.Body)
	var bodyContent models.CreatingUrl
	err := jsonDecoder.Decode(&bodyContent)
	if err != nil {
		log.Fatal(err)
	}

	var creatingUrlModel models.CreatingUrl
	creatingUrlModel.Original_url = bodyContent.Original_url
	//creatingUrlModel.Short_url = ""

	res, err := h.service.ShortenUrl(ctx, &creatingUrlModel)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	fullUrl := []string{os.Getenv("BACKEND_PROTOCOL"), "://", os.Getenv("BACKEND_HOST"), "/", res.Short_url}
	res.Short_url = strings.Join(fullUrl, "")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}

func (h *UrlHandler) GetOrigUrl(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	shortUrl := r.URL.Query().Get("short_url")

	res, err := h.service.GetOrigUrl(ctx, shortUrl)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}

func (h *UrlHandler) ShortRedirect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	shortUrl := strings.TrimPrefix(r.URL.Path, "/")
	if shortUrl == "" {
		http.NotFound(w, r)
		return
	}

	res, err := h.service.GetOrigUrl(ctx, shortUrl)
	if err != nil {
		//http.Error(w, "unavailable", http.StatusBadGateway)
		http.NotFound(w, r)
		return
	}

	if res == nil || res.Original_url == "" {
		http.NotFound(w, r)
		return
	}

	if !strings.HasPrefix(res.Original_url, "https:/") {
		res.Original_url = strings.Join([]string{"https://", res.Original_url}, "")
	}

	http.Redirect(w, r, res.Original_url, http.StatusFound) // 302
}
