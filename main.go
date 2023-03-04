package main

import (
	"fmt"
	"log"
	"test-bpjs/config/env"
	"test-bpjs/src/api/v1/router"
)

func main() {
	env.NewEnv(".env")
	cfg := env.Config

	fmt.Println(cfg.Redis.Host, "host redis")

	router := router.NewRouter()
	if err := router.Run(cfg.Host + ":" + cfg.Port); err != nil {
		log.Fatal("Error running router : ", err)
	}
}
