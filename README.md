# TwitterGoStalker
Get push notifications in Telegram when a twitter user posts a new tweet. User MUST be public

# Building
```
git clone git@github.com:gibsn/TwitterGoStalker.git
cd TwitterGoStalker
make
```

# Configuration
Configuration file is a JSON file
```
Usage of ./bin/twittergostalker:
  -c string
    	path to config (default "cfg.json")
```

Configuration file params:
```
{
    "UserName": "",              // twitter user to poll
    "PollingInterval": 1,        // in seconds
    "TwitterConsumerKey": "",    // twitter application consumer key
    "TwitterConsumerSecret": "", // twitter application consumer secret
    "TwitterAccessToken": "",    // your twitter access token
    "TwitterAccessSecret": "",   // your twitter access secret
    "TelegramBotToken": ""       // telegram bot token
}
```

# Usage
TwitterGoStalker polls the provided user and sends a new message to all subscribers in Telegram.

You will need:
* A Twitter App (https://apps.twitter.com)
* Twitter access token and secret (you can get it in your twitter developer account) 
* A Telegram Bot (https://core.telegram.org)

Telegram Bot commands:
```
\stalk   -- subscribe to notifications
\unstalk -- unsubscribe
```
