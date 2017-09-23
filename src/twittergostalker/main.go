package main

import (
	"flag"
	"time"

	"twittergostalker/stalker"
)

const (
	pollingInterval = 1 * time.Second
)

func main() {
	username := flag.String("u", "", "user to stalker")
	flag.Parse()

	stalker := stalker.NewStalker(*username, pollingInterval)
	stalker.Init()

	stalker.Loop()
}
