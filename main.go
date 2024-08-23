package main

import (
	"log"

	"github.com/hellskater/udhaar-backend/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
