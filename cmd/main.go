package main

import (
	"log"
	"net/http"
	"time"

	"github.com/phuwn/coinchameleon/data/db"
	"github.com/phuwn/coinchameleon/handlers"

	coingecko "github.com/superoo7/go-gecko/v3"
)

func crawl() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	cg := coingecko.NewClient(httpClient)
	handlers.CrawlMarket(cg)
}

func main() {
	db.Start()
	defer db.Close()

	addr := ":8080"
	log.Printf("listening on port%s\n", addr)

	err := http.ListenAndServe(addr, handlers.Router())
	if err != nil {
		log.Printf("server got terminated, err: %s\n", err.Error())
	}
}
