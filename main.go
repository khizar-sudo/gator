package main

import (
	"fmt"

	"github.com/khizar-sudo/feed-aggregator/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.CurrentUserName = "khizar"
	cfg.SetUser()
	cfg = config.Read()

	fmt.Println(cfg)
}
