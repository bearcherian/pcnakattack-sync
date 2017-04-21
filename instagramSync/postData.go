package instagramSync

type OwnerData struct {
	FullName      string `json:"full_name"`
	Id            string `json:"id"`
	ProfilePicUrl string `json:"profile_pic_url"`
	Username      string `json:"username"`
}

type MediaData struct {
	Caption    string    `json:"caption"`
	Code       string    `json:"code"`
	Date       int       `json:"date"`
	DisplayUrl string    `json:"display_url"`
	Id         string    `json:"id"`
	IsVideo    bool      `json:"is_video"`
	Owner      OwnerData `json:"owner"`
	VideoUrl   string    `json:"video_url"`
	MediaId    string
}

type Graphql struct {
	Media 	MediaData	`json:"shortcode_media"`
}

type PostPage struct {
	Graphql Graphql `json:"graphql"`
}

type EntryData struct {
	PostPages []PostPage `json:"PostPage"`
}

type PostData struct {
	EntryData EntryData `json:"entry_data"`
}
