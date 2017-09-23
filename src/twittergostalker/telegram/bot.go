package telegram

import (
	"log"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"

	"twittergostalker/config"
)

type Bot struct {
	bot   *telegram.BotAPI
	token string

	updates telegram.UpdatesChannel
	subs    []*Subscriber
}

type Subscriber struct {
	Name   string
	ChatID int64
}

func NewBot(cfg *config.Config) *Bot {
	return &Bot{
		token: cfg.TelegramBotToken,
	}
}

func (this *Bot) Init() {
	log.Println("telegram: initialiasing")

	var err error

	this.bot, err = telegram.NewBotAPI(this.token)
	if err != nil {
		log.Fatalln("could not create telegram API")
	}

	u := telegram.NewUpdate(0)
	u.Timeout = 60

	this.updates, err = this.bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalln("telegram: could not get updates channel:", err)
	}

	log.Println("telegram: initialiased")
}

func (this *Bot) Routine() {
	for update := range this.updates {
		if update.Message == nil {
			continue
		}

		log.Printf("telegram: @%s: %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Text {
		case "/ping":
			this.ping(&update)
		case "/stalk":
			this.stalk(&update)
		case "/unstalk":
			this.unstalk(&update)
		default:
			log.Printf("telegram: got unknown command '%s' from @%s\n",
				update.Message.Text, update.Message.From.UserName)
		}
	}
}

func (this *Bot) ping(update *telegram.Update) {
	log.Printf("telegram: responding 'pong' to @%s\n", update.Message.From.UserName)
	msg := telegram.NewMessage(update.Message.Chat.ID, "pong")
	this.bot.Send(msg)
}

func (this *Bot) stalk(update *telegram.Update) {
	newSubscriber := &Subscriber{
		Name:   update.Message.From.UserName,
		ChatID: update.Message.Chat.ID,
	}

	for _, sub := range this.subs {
		if sub.Name == update.Message.From.UserName {
			log.Printf("telegram: user @%s is already subscribed", sub.Name)

			msg := telegram.NewMessage(update.Message.Chat.ID, "you are already subscribed")
			this.bot.Send(msg)
			return
		}
	}

	this.subs = append(this.subs, newSubscriber)
	log.Printf("telegram: added @%s to subscribers\n", update.Message.From.UserName)

	msg := telegram.NewMessage(newSubscriber.ChatID, "gotcha")
	this.bot.Send(msg)
}

func (this *Bot) unstalk(update *telegram.Update) {
	for i, sub := range this.subs {
		if sub.Name == update.Message.From.UserName {
			this.subs = append(this.subs[:i], this.subs[i+1:]...)
			log.Printf("telegram: removed @%s from subscribers\n", update.Message.From.UserName)

			msg := telegram.NewMessage(update.Message.Chat.ID, "you are free")
			this.bot.Send(msg)
			return
		}
	}

	log.Printf("telegram: user @%s is not subscribed\n", update.Message.From.UserName)

	msg := telegram.NewMessage(update.Message.Chat.ID, "you are not subscribed")
	this.bot.Send(msg)
}

func (this *Bot) Broadcast(msg string) {
	for _, sub := range this.subs {
		newMessage := telegram.NewMessage(sub.ChatID, msg)
		for _, err := this.bot.Send(newMessage); err != nil; {
			log.Println("telegram: could not send:", err)
		}
	}
}
