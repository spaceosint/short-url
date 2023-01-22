package main

import (
	"ShortURL/internal/app/pkg/app"
	"log"
)

func main() {

	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Server started: %v", err)
	}
}
