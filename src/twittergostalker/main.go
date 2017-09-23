package main

import (
	"flag"
	"time"

	"twittergostalker/config"
	"twittergostalker/stalker"
)

const (
	pollingInterval = 1 * time.Second
)

func main() {
	pathToCfg := flag.String("c", "cfg.json", "path to config")
	flag.Parse()

	cfg := config.Parse(*pathToCfg)

	stalker := stalker.NewStalker(cfg)
	stalker.Init()

	stalker.Routine()
}
