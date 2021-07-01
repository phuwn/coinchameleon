package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/phuwn/coinchameleon/util"
)

func init() {
	env := util.Getenv("RUN_MODE", "")
	if env == "local" || env == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}
}
