package twitterSync

import (
	"github.com/bearcherian/pcnakattackSync/config"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"strings"
)

var appCfg config.AppConfig

func sync(ignoreLast bool, cfg config.AppConfig) {

	appCfg = cfg

	// init twitter
	twitterConfig := oauth1.NewConfig(appCfg.Twitter.ConsumerKey, appCfg.Twitter.ConsumerSecret)
	token := oauth1.NewToken(appCfg.Twitter.Token, appCfg.Twitter.Secret)
	httpClient := twitterConfig.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	// query twitter
	searchParams := &twitter.SearchTweetParams{
		Query: strings.Join(appCfg.Tags, " OR "),
	}

	if !ignoreLast {
		searchParams.SinceID = getLatestId()
		log.Printf("SinceID: %d\n", searchParams.SinceID)
	}

	search, _, err := client.Search.Tweets(searchParams)
	if err != nil {
		log.Fatalf("Could not get tweets. %s", err)
		return
	}

	// load to db
	for _, tweet := range search.Statuses {
		log.Printf("Adding tweet %s\n", tweet)
		AddNewTweet(tweet)
	}
}

func SyncAll(cfg config.AppConfig) {
	sync(true, cfg)
}

func SyncLatest(cfg config.AppConfig) {
	sync(false, cfg)
}

func getLatestId() int64 {
	return GetLatestTwitterId()
}
