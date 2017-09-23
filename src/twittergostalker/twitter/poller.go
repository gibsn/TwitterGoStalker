package twitter

import (
	"log"
	"time"

	"github.com/ChimeraCoder/anaconda"

	"twittergostalker/config"
)

type Poller struct {
	username string

	api             *anaconda.TwitterApi
	currTweetId     int64
	pollingInterval time.Duration

	consumerKey    string
	consumerSecret string

	accessToken  string
	accessSecret string

	TweetsChan chan string
}

func NewPoller(cfg *config.Config) *Poller {
	return &Poller{
		username:        cfg.UserName,
		pollingInterval: cfg.PollingInterval,
		consumerKey:     cfg.TwitterConsumerKey,
		consumerSecret:  cfg.TwitterConsumerSecret,
		accessToken:     cfg.TwitterAccessToken,
		accessSecret:    cfg.TwitterAccessSecret,
		TweetsChan:      make(chan string),
	}
}

func (this *Poller) Init() {
	log.Println("twitter: initialising")

	anaconda.SetConsumerKey(this.consumerKey)
	anaconda.SetConsumerSecret(this.consumerSecret)

	this.api = anaconda.NewTwitterApi(this.accessToken, this.accessSecret)
	if this.api == nil {
		log.Fatalln("could not create twitter API")
	}

	log.Println("twitter: initialised")
}

func (this *Poller) Routine() {
	log.Printf("twitter: polling @%s\n", this.username)

	for {
		user, err := this.api.GetUsersShow(this.username, nil)
		if err != nil {
			log.Printf("twitter: could not get an update for @%s: %s\n", this.username, err)
			close(this.TweetsChan)
			return
		}

		if user.Status == nil {
			continue
		}

		if this.currTweetId != user.Status.Id {
			this.currTweetId = user.Status.Id
			log.Printf("twitter: got a new tweet from @%s: %s\n",
				this.username, user.Status.FullText)

			msg := "@" + this.username + ":\n" + user.Status.FullText
			this.TweetsChan <- msg
		}

		time.Sleep(time.Second * this.pollingInterval)
	}
}
