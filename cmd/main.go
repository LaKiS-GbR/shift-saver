package main

import (
	"log"

	"github.com/LaKiS-GbR/shift-saver/pkg/config"
	"github.com/LaKiS-GbR/shift-saver/pkg/database"
	"github.com/LaKiS-GbR/shift-saver/pkg/router"
)

func main() {
	config.Init()
	if config.Running != nil {
		log.Printf("[main-1] started with config: %+v", config.Running)
	}
	err := database.Init()
	if err != nil {
		log.Fatalf("[main-2] failed to initialize database: %s", err)
	}
	err = router.Init()
	if err != nil {
		log.Fatalf("[main-3] failed to initialize router: %s", err)
	}
}
