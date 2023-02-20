package main

import (
	"log"
	"pixelBattle/internal/config"
	"pixelBattle/internal/service"
)

func main() {
	cfg := config.GetConfig()
	s := service.NewService(cfg)
	log.Fatalln(s.Run())
}
