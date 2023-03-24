package main

import (
	_ "market/bootstrap"
	"market/router"
	"log"
)

func main() {
	if err := router.Router(); err != nil {
		log.Fatal(err)
	}
}
