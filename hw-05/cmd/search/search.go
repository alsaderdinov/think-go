package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"think-go/hw-05/pkg/crawler"
	"think-go/hw-05/pkg/crawler/spider"
	"think-go/hw-05/pkg/index"
)

const (
	depth = 2
	file  = "docs.json"
)

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

	docs, err := processDocs(s)
	if err != nil {
		fmt.Println("err", err)
	}

	index := indexDocs(idx, docs)
	ids := idx.Find(strings.ToLower(*query))

	for _, id := range ids {
		doc := index[id]
		fmt.Printf("%s: %s\n", doc.Title, doc.URL)
	}
}

// processDocs загружает документы из файла если он существует
// если файла не существует, то совершает обход ссылок сайта записывает их в файл
// возвращает найденные документы
func processDocs(s *spider.Service) ([]crawler.Document, error) {
	var docs []crawler.Document
	var err error

	if fileExists() {
		docs, err = readDocs()
		if err != nil {
			return nil, err
		}
	} else {
		f, err := os.Create(file)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		docs = scan(s, urls, depth)

		err = writeDocs(f, docs)
		if err != nil {
			return nil, err
		}
	}
	return docs, nil
}

// fileExists возвращает true если файл существует
func fileExists() bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

// readDocs загружает документы из файла
func readDocs() ([]crawler.Document, error) {
	var res []crawler.Document

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	docs, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(docs, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// writeDocs записывает документы в файл в формате json
func writeDocs(w io.Writer, docs []crawler.Document) error {
	b, err := json.MarshalIndent(docs, "", "\t")
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

// indexDocs индексирует документы и возвращает проиндексированные документы
func indexDocs(idx *index.Service, docs []crawler.Document) []crawler.Document {
	res := make([]crawler.Document, len(docs))

	for i, doc := range docs {
		doc.ID = i
		idx.Add(doc.Title, doc.ID)
		res[doc.ID] = doc
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
