package main

import (
	"crudProj/server"
	"log"
)

func main() {
	if err := server.Run();
	err != nil {
		log.Fatal(err)
	}
}

