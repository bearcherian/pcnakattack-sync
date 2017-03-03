package youtubeSync

import (
	"github.com/bearcherian/pcnakattackSync/db"
	"google.golang.org/api/youtube/v3"
	"log"
	"time"
)

const SELECT_LATEST_DATE = "SELECT publishedAt FROM youtube ORDER BY publishedAt DESC LIMIT 1"
const INSERT_NEW = "INSERT INTO youtube(id,publishedAt,title,description,thumbnail_default,thumbnail_medium,thumbnail_high,channel_id,channelTitle,authorName,link,profile_picture) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
const VIDEO_LINK_PREFIX = "https://www.youtube.com/watch?v="
const YT_DATE_LAYOUT = "2006-01-02T15:04:05.000Z"

func GetLatestPublishedDate() time.Time {
	dbConn := db.GetClient()

	rows, err := dbConn.Query(SELECT_LATEST_DATE)
	if err != nil {
		log.Println("Could not retrieve latest YouTube")
	}

	rows.Next()
	var publishedAt time.Time
	rows.Scan(&publishedAt)

	return publishedAt

}

func AddNewYoutube(searchResponse *youtube.SearchResult, profileImageUrl string) {
	// id,publishedAt,title,description,thumbnail_default,thumbnail_medium,thumbnail_high,
	// channel_id,channelTitle,authorName,link,profile_picture

	publishedTime, timeErr := time.Parse(YT_DATE_LAYOUT, searchResponse.Snippet.PublishedAt)
	if timeErr != nil {
		log.Printf("Unable to format time. %v", timeErr)
		return
	}

	dbConn := db.GetClient()

	var videoLinkUrl = VIDEO_LINK_PREFIX + searchResponse.Id.VideoId

	_, err := dbConn.Exec(INSERT_NEW, searchResponse.Id.VideoId, publishedTime, searchResponse.Snippet.Title, searchResponse.Snippet.Description,
		searchResponse.Snippet.Thumbnails.Default.Url, searchResponse.Snippet.Thumbnails.Medium.Url, searchResponse.Snippet.Thumbnails.High.Url,
		searchResponse.Snippet.ChannelId, searchResponse.Snippet.ChannelTitle, searchResponse.Snippet.ChannelTitle,
		videoLinkUrl, profileImageUrl)
	if err != nil {
		log.Printf("Unable to insert new YouTube video data.\n%v", err)
	}
}
