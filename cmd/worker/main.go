package main

import (
	"goDistributedSystem/internal/worker"
	"log"
)

func main() {

	if err := worker.Run("localhost:50051"); err != nil {
		log.Fatal(err)
	}
}
