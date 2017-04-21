package instagramSync

type PageInfo struct {
	EndCursor   string `json:"end_cursor"`
	HasNextPage bool   `json:"has_next_page"`
}

type Owner struct {
	Id string `json:"id"`
}

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Comments struct {
	Count int `json:"count"`
}

type Likes struct {
	Count int `json:"count"`
}

type MediaNode struct {
	Owner            Owner      `json:"owner"`
	Dimensions       Dimensions `json:"dimensions"`
	Comments         Comments   `json:"comments"`
	CommentsDisabled bool       `json:"comments_disabled"`
	DisplaySrc       string     `json:"display_src"`
	IsVideo          bool       `json:"is_video"`
	Date             int        `json:"date"`
	Likes            Likes      `json:"likes"`
	Id               string     `json:"id"`
	ThumbnailSrc     string     `json:"thumbnail_src"`
	Caption          string     `json:"caption"`
	Code             string     `json:"code"`
	PostData         PostData
	MediaId          string
}

type Media struct {
	PageInfo PageInfo    `json:"page_info"`
	Nodes    []MediaNode `json:"nodes"`
}

type Tags struct {
	Media Media `json:"media"`
}

type TagResponse struct {
	Tag Tags `json:"tag"`
}
