package main

import (
	"log"

	"github.com/alfonsovgs/go-hexagonal-architecture/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
