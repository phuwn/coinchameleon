package main

import (
	"log"

	"github.com/phuwn/coinchameleon/data/db"
)

func main() {
	db.Start()
	defer db.Close()

	if err := crawl(); err != nil {
		log.Fatal(err)
	}
}
