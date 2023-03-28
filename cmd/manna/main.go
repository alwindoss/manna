package main

import (
	"log"

	"github.com/alwindoss/manna/internal/engine"
)

func main() {
	cfg := engine.Config{
		Port: 3030,
	}
	log.Fatal(engine.Run(&cfg))
}
