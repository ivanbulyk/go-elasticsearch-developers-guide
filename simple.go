package main

import (
	"log"

	"github.com/elastic/go-elasticsearch"
	"github.com/ivanbulyk/go-elasticsearch-developers-guide/my_elastic"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating elasticsearch client: %v", err)
	}
	// log.Println(es.Info)

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting elasticsearch response: %v", err)
	}
	defer res.Body.Close()

	my_elastic.Start()
	// log.Println(res)
}
