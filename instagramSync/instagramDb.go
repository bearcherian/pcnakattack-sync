package instagramSync

import (
	"database/sql"
	"fmt"
	"github.com/bearcherian/pcnakattackSync/db"
	"log"
	"time"
	"bytes"
)

const SELECT_INSTAGRAM_BY_ID string = "SELECT * FROM instagrams WHERE id = ?"
const SELECT_LAST_INSTAGRAM_ID string = "SELECT id FROM instagrams ORDER BY id DESC LIMIT 1"
const SELECT_BY_CODE string = "SELECT * FROM instagrams WHERE link LIKE ?"
const INSERT_INSTAGRAM string = "INSERT INTO instagrams(id,type,link,caption_text,created_time,user_id,user_username,user_full_name,user_profile_picture,images_thumbnail,images_standard_resolution,images_low_resolution,videos_low_bandwidth,videos_low_resolution,videos_standard_resolution) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

// Adds a post to the data base. Returns false if
func AddToDb(media MediaNode)  {

	dbConn := db.GetClient()

	var mediaType string
	if media.IsVideo {
		mediaType = "video"
	} else {
		mediaType = "image"
	}

	var link = fmt.Sprintf("http://www.instagram.com/p/%v/", media.Code)

	postData := media.PostData.EntryData.PostPages[0]

	log.Printf("Adding instagram %v", media.MediaId)
	_, err := dbConn.Exec(INSERT_INSTAGRAM, media.MediaId, mediaType, link, postData.Graphql.Media.Caption, time.Unix(int64(media.Date), 0),
		postData.Graphql.Media.Owner.Id, postData.Graphql.Media.Owner.Username, postData.Graphql.Media.Owner.FullName,
		postData.Graphql.Media.Owner.ProfilePicUrl, sql.NullString{"", false}, postData.Graphql.Media.DisplayUrl, sql.NullString{"", false},
		sql.NullString{"", false}, sql.NullString{postData.Graphql.Media.VideoUrl, postData.Graphql.Media.IsVideo}, sql.NullString{"", false})
	if err != nil {
		log.Printf("Unable to add new Instagram. %v\n", err)
	}

}

func MediaExistsInDb(media MediaNode) bool {
	inDb := codeIsAlreadyInDb(media.Code)
	return inDb
}

func codeIsAlreadyInDb(code string) bool {
	dbConn := db.GetClient()
	codeBuffer := bytes.Buffer{}
	codeBuffer.WriteString("%")
	codeBuffer.WriteString(code)
	codeBuffer.WriteString("%")
	
	rows, err := dbConn.Query(SELECT_BY_CODE, codeBuffer.String())
	if err != nil {
		log.Printf("Unable to query DB for code %v - %v\n", code, err)
		return false
	}
	defer rows.Close()
	return rows.Next()
}
