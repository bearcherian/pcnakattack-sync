package twitterSync

import (
	"github.com/bearcherian/pcnakattackSync/config"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"strings"
)

var appCfg config.Config

func sync(ignoreLast bool, cfg config.Config) {

	appCfg = cfg

	// init twitter
	twitterConfig := oauth1.NewConfig(appCfg.Twitter.ConsumerKey, appCfg.Twitter.ConsumerSecret)
	token := oauth1.NewToken(appCfg.Twitter.Token, appCfg.Twitter.Secret)
	httpCLient := twitterConfig.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpCLient)
	fmt.Println(client)

	// query twitter
	searchParams := &twitter.SearchTweetParams{
		Query: strings.Join(appCfg.Tags, " OR "),
	}

	if !ignoreLast {
		searchParams.SinceID = getLatestId()
		fmt.Printf("SinceID: %d\n", searchParams.SinceID)
	}

	search, _, err := client.Search.Tweets(searchParams)
	if err != nil {
		log.Fatalf("Could not get tweets. %s", err)
		return
	}

	// load to db
	for _, tweet := range search.Statuses {
		fmt.Printf("Adding tweet %s\n", tweet)
		AddNewTweet(tweet, appCfg)
	}
}

func SyncAll(cfg config.Config) {
	sync(true, cfg)
}

func SyncLatest(cfg config.Config) {
	sync(false, cfg)
}

func getLatestId() int64 {
	return GetLatestTwitterId(appCfg)
}
