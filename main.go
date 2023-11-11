package main

import (
	"compartamos/customers/bootstrap"
	"log"
	"os"
)

func main() {
	err := bootstrap.Run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
