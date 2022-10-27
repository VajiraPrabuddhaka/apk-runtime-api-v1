package common

type Pagination struct {
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Total    int    `json:"total"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}
