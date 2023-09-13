package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"think-go/hw-13/pkg/api"
	"think-go/hw-13/pkg/crawler"
	"think-go/hw-13/pkg/crawler/spider"
	"think-go/hw-13/pkg/index"
)

const (
	depth = 2
	file  = "docs.json"
	addr  = ":8080"
)

var urls = []string{"https://go.dev", "https://practical-go-lessons.com"}

func main() {
	s := spider.New()
	idx := index.New()

	var docs []crawler.Document
	var err error

	if fileExists(file) {
		docs, err = loadDocs(file)
		if err != nil {
			fmt.Println("Error loading documents from file:", err)
			return
		}
	} else {
		docs = scan(s, urls, depth)
		err = saveDocs(file, docs)
		if err != nil {
			fmt.Println("Error saving documents:", err)
			return
		}
	}

	idxDocs := indexDocs(idx, docs)

	api := api.New(idxDocs, idx)
	if err = api.Serve(addr); err != nil {
		log.Fatal(err)
	}
}

// saveDocs создает файл и записывает документы
func saveDocs(file string, docs []crawler.Document) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	return writeDocs(f, docs)
}

// loadDocs загружает документы из файла
func loadDocs(file string) ([]crawler.Document, error) {
	docs, err := readDocs(file)
	if err != nil {
		return nil, err
	}
	return docs, nil
}

// fileExists возвращает true если файл существует
func fileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

// readDocs читает документы из файла
func readDocs(file string) ([]crawler.Document, error) {
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
