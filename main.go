package main

import (
	"fmt"
	"log"

	"github.com/anandki1209/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Read config: %+v\n", cfg)

	if err := cfg.SetUser("Anand"); err != nil {
		log.Fatalf("error trying to set current user name: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Read config: %+v\n", cfg)
}
