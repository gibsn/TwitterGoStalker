package stalker

import (
	"log"

	"twittergostalker/config"
	"twittergostalker/telegram"
	"twittergostalker/twitter"
)

type Stalker struct {
	username string

	poller *twitter.Poller
	bot    *telegram.Bot
}

func NewStalker(cfg *config.Config) *Stalker {
	return &Stalker{
		username: cfg.UserName,
		poller:   twitter.NewPoller(cfg),
		bot:      telegram.NewBot(cfg),
	}
}

func (this *Stalker) Init() {
	this.poller.Init()
	this.bot.Init()
}

func (this *Stalker) Routine() {
	go this.poller.Routine()
	go this.bot.Routine()

	for newTweet := range this.poller.TweetsChan {
		this.bot.Broadcast(newTweet)
	}

	log.Println("stalker: the new tweets channel has been closed, quittig")
}
