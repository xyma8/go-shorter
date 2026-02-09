package models

type CreatingUrl struct {
	Original_url string `json:"original_url"`
	//Short_url    string
}

type OrigUrl struct {
	Original_url string `json:"original_url"`
}

type Url struct {
	Original_url string `json:"original_url"`
	Short_url    string `json:"short_url"`
}
