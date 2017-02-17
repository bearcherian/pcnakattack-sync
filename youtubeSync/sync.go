package youtubeSync

import (
	"github.com/bearcherian/pcnakattackSync/config"
	"github.com/google/google-api-go-client/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"log"
	"net/http"
	"strings"
	"time"
)

var appCfg config.Config

const MAX_RESULTS = 50
const TIME_LAYOUT = "2006-01-02T03:04:05Z"

func sync(ignoreLast bool, cfg config.Config) {
	appCfg = cfg

	client := &http.Client{
		Transport: &transport.APIKey{Key: appCfg.YouTube.ApiKey},
	}

	ytService, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	var hasNextPage bool = true
	var nextPageToken string

	for hasNextPage {
		response := getSearchResponse(ytService, strings.Join(appCfg.Tags, "|"), ignoreLast, nextPageToken)
		for i, item := range response.Items {
			log.Printf("%v:\t%v\n", i, item)
			AddNewYoutube(item, getProfileUrl(item, client), appCfg)
		}

		if response.NextPageToken != "" {
			nextPageToken = response.NextPageToken
		} else {
			hasNextPage = false
		}
	}

}

func getSearchResponse(ytService *youtube.Service, query string, ignoreLast bool, nextPageToken string) *youtube.SearchListResponse {
	call := ytService.Search.List("id,snippet").
		Q(query).
		MaxResults(MAX_RESULTS).
		Order("date").
		Type("video")

	if !ignoreLast {
		publishedAtTime := GetLatestPublishedDate(appCfg)
		call.PublishedAfter(publishedAtTime.Add(time.Second).Format(TIME_LAYOUT))
	}

	if nextPageToken != "" {
		call.PageToken(nextPageToken)
	}

	log.Println(call)

	response, err := call.Do()
	if err != nil {
		log.Fatalf("Could not retrieve YouTube videos.\n%v", err)
	}

	return response
}

func getProfileUrl(item *youtube.SearchResult, client *http.Client) string {
	ytService, err := youtube.New(client)
	if err != nil {
		log.Printf("Unable to retrieve profile image.\n%v", err)
		return ""
	}

	call := ytService.Channels.List("snippet").Id(item.Snippet.ChannelId)

	response, err := call.Do()
	if err != nil {
		log.Printf("Unable to retrieve profile image.\n%v", err)
		return ""
	}

	return response.Items[0].Snippet.Thumbnails.Default.Url
}

func SyncAll(cfg config.Config) {
	sync(true, cfg)
}

func SyncLatest(cfg config.Config) {
	sync(false, cfg)
}
