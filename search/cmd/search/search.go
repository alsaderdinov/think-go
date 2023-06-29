package main

import (
	"flag"
	"fmt"
	"strings"

	"think-go/search/pkg/crawler"
	"think-go/search/pkg/crawler/spider"
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

	docs := scan(urls, depth)
	res := search(docs, *query)

	for _, i := range res {
		fmt.Printf("%s: %s\n", i.Title, i.URL)
	}
}

// scan осуществляет обход ссылок сайта, указанного в URL,
// с учётом глубины перехода по ссылкам, переданной в depth.
func scan(urls []string, depth int) []crawler.Document {
	var res []crawler.Document

	s := spider.New()

	for _, url := range urls {
		docs, err := s.Scan(url, depth)
		if err != nil {
			fmt.Println("Error scanning URL:", url, "Error:", err)
			continue
		}

		res = append(res, docs...)
	}

	return res
}

// search ищет ссылки в переданном списке документов по заданному запросу и возвращает совпадающие документы.
// выполняет поиск без учета регистра, сравнивая запрос переданный в query с URL-адресами документов
func search(docs []crawler.Document, query string) []crawler.Document {
	var res []crawler.Document

	for _, doc := range docs {
		if strings.Contains(strings.ToLower(doc.URL), strings.ToLower(query)) {
			res = append(res, doc)
		}
	}
	return res
}
