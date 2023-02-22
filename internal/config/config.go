package config

import (
	"flag"
	"time"
)

type Config struct {
	RunAddress string
	UpdateTime time.Duration
	FieldSize  int
}

func GetConfig() *Config {
	cfg := Config{}

	flag.StringVar(&cfg.RunAddress, "a", ":8080", "Run address for service")
	flag.DurationVar(&cfg.UpdateTime, "t", 1*time.Second, "Field update time for user")
	flag.IntVar(&cfg.FieldSize, "s", 60, "Field size width/height")

	flag.Parse()

	return &cfg
}
