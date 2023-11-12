package main

import (
	"compartamos/customers/bootstrap"
	"log"
)

func main() {
	err := bootstrap.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
