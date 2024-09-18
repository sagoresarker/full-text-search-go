package main

import (
	"flag"
	"log"
	"time"

	"github.com/sagoresarker/full-text-search-go/utils"
)

func main() {
	var dummppath, query string
	flag.StringVar(&dummppath, "p", "enwiki-latest-abstract1.xml.gz", "Path to the dump file")
	flag.StringVar(&query, "q", "Small wild cat", "Query string")
	flag.Parse()
	log.Println("Full text search on", dummppath, "for", query)
	start := time.Now()
	docs, err := utils.LoadDoccuments(dummppath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Loaded", len(docs), "documents in", time.Since(start))

	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Println("Indexed", len(idx), "documents in", time.Since(start))

	start = time.Now()
	matchedIDs := idx.Search(query)
	log.Println("Found", len(matchedIDs), "documents in", time.Since(start))

	for _, id := range matchedIDs {
		docs := docs[id]
		log.Println("%d\t%s\n", id, docs.Text)

	}
}
