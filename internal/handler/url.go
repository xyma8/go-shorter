package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
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
	short := strings.TrimPrefix(r.URL.Path, "/")
	if short == "" {
		http.NotFound(w, r)
		return
	}

	// запрос к api
	apiURL := "http://localhost:8080/api/get_orig?short_url=" + url.QueryEscape(short)
	res, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "api unavailable", http.StatusBadGateway)
		return
	}
	defer res.Body.Close()

	var urlData models.OrigUrl
	if err := json.NewDecoder(res.Body).Decode(&urlData); err != nil {
		http.Error(w, "bad api response", http.StatusBadGateway)
		return
	}

	if urlData.Original_url == "" {
		http.NotFound(w, r)
		return
	}

	if !strings.HasPrefix(urlData.Original_url, "https:/") {
		urlData.Original_url = strings.Join([]string{"https://", urlData.Original_url}, "")
	}

	http.Redirect(w, r, urlData.Original_url, http.StatusFound) // 302
	//fmt.Println(strings.TrimLeft(r.URL.Path, "/"))
}
