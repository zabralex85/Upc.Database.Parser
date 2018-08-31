package parser

type BarcodeResult struct {
	Title string `json:"title"`
	Defines map[string]string
}