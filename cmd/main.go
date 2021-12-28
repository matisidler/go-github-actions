package main

import (
	"log"

	"github.com/matisidler/go-github-actions/apiserver"
)

func main() {
	config := apiserver.NewConfig()
	s := apiserver.New(config)
	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
