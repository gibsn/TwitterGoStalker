package twitter

import (
	"log"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

type Poller struct {
	username string

	api             *anaconda.TwitterApi
	pollingInterval time.Duration

	consumerKey    string
	consumerSecret string

	accessToken  string
	accessSecret string

	TweetsChan chan string
}

func NewPoller(username string, pollingInterval time.Duration) *Poller {
	return &Poller{
		username:        username,
		pollingInterval: pollingInterval,
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
	log.Println("twitter: polling @%s", this.username)
	currTweetId := int64(0)

	for {
		user, err := this.api.GetUsersShow(this.username, nil)
		if err != nil {
			log.Printf("twitter: could not get an update for @%s: %s\n", this.username, err)
			close(this.TweetsChan)
			return
		}

		if currTweetId != user.Status.Id {
			currTweetId = user.Status.Id
			log.Printf("twitter: got a new tweet from %s: %s\n",
				this.username, user.Status.FullText)

			msg := this.username + ":\n" + user.Status.FullText
			this.TweetsChan <- msg
		}

		time.Sleep(this.pollingInterval)
	}
}
