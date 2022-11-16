package main

import (
    "fmt"
    "github.com/caarlos0/env/v6"
)

type Config struct {
    User string	 `env:"USER"`
}

func main() {
    var cfg Config
    // допишите код здесь
	err := env.Parse(&cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

    fmt.Printf("Current user is %s\n", cfg.User)
}