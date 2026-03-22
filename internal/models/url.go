package models

type CreatingUrl struct {
	Original_url string `json:"orig_url"`
	//Short_url    string
}

type OrigUrl struct {
	Original_url string `json:"orig_url"`
}

type ShortUrl struct {
	Short_url string `json:"s_url"`
}

type Url struct {
	Original_url string `json:"orig_url"`
	Short_url    string `json:"short_url"`
}
