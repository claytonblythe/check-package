package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func makeUSPSRequest(url string, packageNumStr string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "tools.usps.com")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("sec-fetch-site", "cross-site")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("accept-language", "en-US,en;q=0.9,es-US;q=0.8,es;q=0.7,zh-CN;q=0.6,zh;q=0.5")
	q := req.URL.Query()
	q.Add("qtc_tLabels1", packageNumStr)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	expectedDelivery := doc.Find("#tracked-numbers > div > div > div > div > div.product_summary > div.expected_delivery > h2 > span").Text()
	fmt.Println(expectedDelivery)
	weekdayStr := doc.Find("#tracked-numbers > div > div > div > div > div.product_summary > div.expected_delivery > h2 > span > span:nth-child(1) > em").Text()
	dayStr := doc.Find("#tracked-numbers > div > div > div > div > div.product_summary > div.expected_delivery > h2 > span > span:nth-child(1) > strong").Text()
	timeStr := doc.Find("#tracked-numbers > div > div > div > div > div.product_summary > div.expected_delivery > h2 > span > span:nth-child(2) > span > strong").Text()
	timeStr = strings.ReplaceAll(timeStr, "\n", "")
	timeStr = strings.ReplaceAll(timeStr, "\t", "")
	finalTimeStr := strings.Fields(timeStr)[0]
	monthStr := doc.Find("#tracked-numbers > div > div > div > div > div.product_summary > div.expected_delivery > h2 > span > span:nth-child(1) > span > span:nth-child(1)").Text()
	statusStr := doc.Find("#tracked-numbers > div > div > div > div > div.product_summary > div.delivery_status > h2 > strong").Text()
	statusFeed := doc.Find("#tracked-numbers > div > div > div > div > div.product_summary > div.delivery_status > div").ChildrenFiltered("p").Text()
	statusFeed = strings.TrimSpace(statusFeed)
	parts := strings.Split(statusFeed, " ")
	var parts2 []string
	for _, part := range parts {
		if part != "" {
			parts2 = append(parts2, part)
		}
	}
	statusFeed = strings.Join(parts2, " ")
	fmt.Printf("\nPackage %v %v\nExpected Delivery by %v %v %v at %v\n", packageNumStr, statusStr, weekdayStr, monthStr, dayStr, finalTimeStr)
	fmt.Printf("LAST UPDATE:\n%v\n\n", statusFeed)

}

func main() {
	var packageNumStr string = os.Args[1]
	url := "https://tools.usps.com/go/TrackConfirmAction"
	makeUSPSRequest(url, packageNumStr)

}
