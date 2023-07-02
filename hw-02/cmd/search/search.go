package main

import (
	"flag"
	"fmt"
	"strings"
	"think-go/hw-02/pkg/crawler"
	"think-go/hw-02/pkg/crawler/spider"
	"think-go/hw-02/pkg/index"
)

const depth = 2

var urls = []string{"https://go.dev", "https://practical-go-lessons.com"}

func main() {
	query := flag.String("s", "", "query to find URLs")
	flag.Parse()

	if *query == "" {
		fmt.Println("The query is empty. Please, read the help: -help")
		return
	}

	idx, s := index.New(), spider.New()

	docs := search(idx, s, urls, depth)
	ids := idx.Find(strings.ToLower(*query))

	for _, id := range ids {
		doc := docs[id]
		fmt.Printf("%s: %s\n", doc.Title, doc.URL)
	}
}

// search обход ссылок сайта по заданным URL-адресам для сбора документов и индексирует их
// возвращает фрагмент crawler.Document, содержащий
// отсканированные документы.
func search(idx *index.Service, s *spider.Service, urls []string, depth int) []crawler.Document {
	var res []crawler.Document
	var count int

	for _, url := range urls {
		docs, err := s.Scan(url, depth)
		if err != nil {
			fmt.Println("Error scanning URL", url, "Error:", err)
			continue
		}

		for _, doc := range docs {
			doc.ID = count
			idx.Add(doc.Title, doc.ID)
			res = append(res, doc)
			count++
		}
	}

	return res
}
