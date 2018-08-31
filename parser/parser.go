package parser

import (
	"net/http"
	"strings"
	"log"
	"github.com/PuerkitoBio/goquery"
	"encoding/json"
) 

func GetDataJSON(barcode string) []byte{
	pageObject := GetData(barcode)
	obj, err := json.Marshal(pageObject)
	if err != nil {
		log.Fatal(err)
	}
	return obj
}

func GetData(barcode string) *BarcodeResult {
	upcURLFormat := `https://www.barcodable.com/upc/{0}`
	eanURLFormat := `https://www.barcodable.com/ean/{0}`
	url := ""

	if(len(barcode) > 12){
		url = strings.Replace(eanURLFormat, "{0}", barcode, 1)
	}else{
		url = strings.Replace(upcURLFormat, "{0}", barcode, 1)
	}

	pageResult := downloadUpcPage(url)
	pageObject := parseDocument(pageResult)

	return pageObject
}


func parseDocument(response *http.Response) *BarcodeResult{
	
	result := new(BarcodeResult)
	result.Defines = make(map[string]string)

	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		log.Fatal(err)
	}

	result.Title = doc.Find("h2").First().Text()

	counter := 0
	lastKey := ""

	doc.Find("tbody tr > td").Each(func(i int, td *goquery.Selection) {	
		text := td.Text()

		if(counter == 0){
			lastKey = text
			counter = counter + 1
		} else {
			if(len(lastKey) > 0){
				result.Defines[lastKey] = text
				counter = 0
			}
		}		
	})

	return result
}

func downloadUpcPage(url string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8`)
	req.Header.Add("UserAgent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36`)

	resp, err := client.Do(req)
	_check(err)

	return resp
}

func _check(err error) {
	if err != nil {
		panic(err)
	}
}