package stalker

import (
	"log"
	"time"

	"twittergostalker/telegram"
	"twittergostalker/twitter"
)

type Stalker struct {
	username string

	poller *twitter.Poller
	bot    *telegram.Bot
}

func NewStalker(username string, pollingInterval time.Duration) *Stalker {
	return &Stalker{
		username: username,
		bot:      telegram.NewBot(),
	}
}

func (this *Stalker) Init() {
	this.poller.Init()
	this.bot.Init()
}

func (this *Stalker) Loop() {
	go this.poller.Routine()
	go this.bot.Routine()

	for newTweet := range this.poller.TweetsChan {
		this.bot.Broadcast(newTweet)
	}

	log.Println("stalker: the new tweets channel has been closed, quittig")
}
