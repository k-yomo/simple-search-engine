package main

import "strings"

func search(docs []document, query string) []document {
	var r []document
	for _, doc := range docs {
		if strings.Contains(doc.Text, query) {
			r = append(r, doc)
		}
	}
	return r
}
