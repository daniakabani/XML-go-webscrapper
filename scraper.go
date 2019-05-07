package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// make a funcion that will parse through html elements and select those needed
func processHtmlElements(index int, element *goquery.Selection) {e
	link := element.Find("loc").Text()
	content := strings.Split(link, "/")
	strippedContent := content[5]
	strippedContent = strings.Replace(strippedContent, "-", " ", -1)
	merchant := strings.Replace(strippedContent, "https:", " ", -1)

	merchantArray := []string{merchant}

	fmt.Println(merchantArray)

	f, err := os.OpenFile("merchants.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := csv.NewWriter(f)
	w.Write([]string{merchant})
	w.Flush()

}

func main() {
	// by default the timeout for connection is forever, just to stay safe it is better to set a timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// create an http request and check if there is an error
	request, err := http.NewRequest("GET", "https://myfave.com/kuala-lumpur/sitemap/partners.xml", nil)
	// check for errors in the url
	// nil is like null for empty pointers
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("User-Agent", "this is a scraping test as requested kindly ignore me, i shall mean no harm =)")

	// request a connection
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	// in case of an error close the connection body
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find("url").Each(processHtmlElements)

}
