package twitterSync

import (
	"database/sql"
	"github.com/bearcherian/pcnakattackSync/db"
	"github.com/dghubble/go-twitter/twitter"
	"log"
	"time"
)

const SELECT_LAST_TWEET_ID string = "SELECT id FROM tweets ORDER BY id DESC LIMIT 1"
const INSERT_NEW_TWEET string = "INSERT INTO tweets(id,created_at,text,user_id,user_name,user_screen_name,user_profile_image_url_https,media_type,media_url_https) VALUES(?,?,?,?,?,?,?,?,?)"
const TIME_LAYOUT string = "Mon Jan 02 15:04:05 -0700 2006"

func GetLatestTwitterId() int64 {
	dbConn := db.GetClient()

	rows, err := dbConn.Query(SELECT_LAST_TWEET_ID)
	if err != nil {
		log.Fatal(err)
		return 0
	}

	rows.Next()
	var id int64
	rows.Scan(&id)

	return id
}

func AddNewTweet(tweet twitter.Tweet) {
	dbConn := db.GetClient()

	var mediaType sql.NullString
	var mediaUrl sql.NullString

	// id,created_at,text,user_id,user_name,user_screen_name,user_profile_image_url_https,media_type,media_url_https
	if len(tweet.Entities.Media) > 0 {
		mediaType = sql.NullString{String: tweet.Entities.Media[0].Type, Valid: true}
		mediaUrl = sql.NullString{String: tweet.Entities.Media[0].MediaURLHttps, Valid: true}
	}

	log.Println(INSERT_NEW_TWEET)
	stmt, prepareErr := dbConn.Prepare(INSERT_NEW_TWEET)
	if prepareErr != nil {
		log.Fatal("Could not prepare insert")
		log.Fatal(prepareErr)
		return
	}

	createdTime, timeParseErr := time.Parse(TIME_LAYOUT, tweet.CreatedAt)
	if timeParseErr != nil {
		log.Fatalln(timeParseErr)
	}

	_, stmtErr := stmt.Exec(tweet.ID, createdTime, tweet.Text, tweet.User.ID, tweet.User.Name, tweet.User.ScreenName, tweet.User.ProfileImageURLHttps, mediaType, mediaUrl)

	if stmtErr != nil {
		log.Println("Could not insert tweet.")
		log.Fatal(stmtErr)
	}
}
