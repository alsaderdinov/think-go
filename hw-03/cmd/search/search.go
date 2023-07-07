package main

import (
	"flag"
	"fmt"
	"strings"
	"think-go/hw-03/pkg/crawler"
	"think-go/hw-03/pkg/crawler/spider"
	"think-go/hw-03/pkg/index"
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

	idx := index.New()
	s := spider.New()

	docs := scan(s, urls, depth)
	index := indexed(idx, docs)

	ids := idx.Find(strings.ToLower(*query))

	for _, id := range ids {
		doc := index[id]
		fmt.Printf("%s: %s\n", doc.Title, doc.URL)
	}
}

// indexed индексирует документы и возвращает проиндексированные документы
func indexed(idx *index.Service, docs []crawler.Document) []crawler.Document {
	var res []crawler.Document

	for i, doc := range docs {
		doc.ID = i
		idx.Add(doc.Title, doc.ID)
		res = append(res, doc)
	}
	return res
}

// scan совершает обход ссылок сайта по заданным URL-адресам
// и возвращает найденные документы
func scan(s *spider.Service, urls []string, depth int) []crawler.Document {
	var res []crawler.Document

	for _, url := range urls {
		docs, err := s.Scan(url, depth)
		if err != nil {
			fmt.Println("Error scanning URL", url, "Error:", err)
			continue
		}
		res = append(res, docs...)
	}
	return res
}
