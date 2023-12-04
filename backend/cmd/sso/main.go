package main

import (
	"fmt"
	"offline_project/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
