package main

import (
	"github.com/spaceosint/short-url/internal/app/pkg/app"
	"log"
)

func main() {

	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}
