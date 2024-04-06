package main

import (
	"fmt"
	"log"

	"github.com/alwindoss/manna/internal/engine"
)

func main() {
	cfg := engine.Config{
		Port: 3030,
	}
	fmt.Println("starting manna on port ", cfg.Port)
	log.Fatal(engine.Run(&cfg))
}
