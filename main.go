package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var query string
	flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()

	docs, err := loadDocuments("enwiki-latest-abstract1.xml")
	if err != nil {
		panic(err)
	}

	startedAt := time.Now()

	idx := index{}
	idx.add(docs)
	fmt.Println("build index finished:", time.Since(startedAt))

	foundIDs := idx.search(query)
	result := buildSearchResult(docs, foundIDs)

	fmt.Println("query:", query)
	fmt.Println("docs:", len(docs))
	fmt.Println("search finished:", time.Since(startedAt))
	for _, doc := range result {
		fmt.Println(doc.ID, doc.Text)
	}
}

func buildSearchResult(docs []document, foundIDs []int) []document {
	docMap := map[int]document{}
	for _, doc := range docs {
		docMap[doc.ID] = doc
	}

	var r []document
	for _, id := range foundIDs {
		if doc, ok := docMap[id]; ok {
			r = append(r, doc)
		}
	}

	return r
}
