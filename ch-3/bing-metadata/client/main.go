package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/shnupta/blackhatgo/ch-3/bing-metadata/metadata"
)

func handler(i int, s *goquery.Selection, wg *sync.WaitGroup) {
	defer wg.Done()
	url, ok := s.Find("a").Attr("href")
	if !ok {
		return
	}

	var ret string

	ret = fmt.Sprintf("%d: %s\n", i, url)
	res, err := http.Get(url)
	if err != nil {
		return
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	defer res.Body.Close()

	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return
	}

	cp, ap, err := metadata.NewProperties(r)
	if err != nil {
		return
	}

	ret = ret + fmt.Sprintf(
		"%25s %25s - %s %s\n",
		cp.Creator,
		cp.LastModifiedBy,
		ap.Application,
		ap.GetMajorVersion())

	fmt.Print(ret)
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("Missing required argument. Usage: main.go domain ext")
	}
	domain := os.Args[1]
	filetype := os.Args[2]

	q := fmt.Sprintf(
		"site:%s && filetype:%s && instreamset:(url title):%s",
		domain,
		filetype,
		filetype)
	search := fmt.Sprintf("http://www.bing.com/search?q=%s", url.QueryEscape(q))
	doc, err := goquery.NewDocument(search)
	if err != nil {
		log.Panicln(err)
	}

	//fmt.Println(doc.Html())

	s := "html body div#b_content main ol#b_results li.b_algo h2"
	// Make file downloading concurrent
	// Should probably try to organise the results...
	var wg sync.WaitGroup
	doc.Find(s).Each(func(i int, s *goquery.Selection) {
		wg.Add(1)
		go handler(i, s, &wg)
	})
	wg.Wait()
}
