package instagramSync

import (
	"github.com/bearcherian/pcnakattackSync/config"
	//"github.com/bearcherian/pcnakattackSync/db"
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"bytes"
)

const IG_TAG_URL_FMT string = "https://www.instagram.com/explore/tags/%v/?__a=1&__b=1&max_id=%v"

func sync(ignoreLast bool, cfg config.AppConfig) {

	for _, tag := range cfg.Tags {
		log.Printf("Processing %v", tag)
		loadMediaForTag(tag, ignoreLast)
	}
}

func SyncLatest() {
	sync(false, config.GetConfig())
}

func loadMediaForTag(tag string, ignoreLast bool) {

	var hasMorePages = true
	var nextCursor = ""

	for hasMorePages {
		tagResponse := getTagResponse(tag, nextCursor)

		for _, mediaNode := range tagResponse.Tag.Media.Nodes {
			postString := getPostDetailsString(mediaNode.Code)

			err := json.Unmarshal([]byte(postString), &mediaNode.PostData)
			if err != nil {
				log.Printf("Json Error: %v\n\n%v\n\n", err, postString)
				continue
			}

			postPage := mediaNode.PostData.EntryData.PostPages[0]

			if postPage.Graphql.Media.Id != "" && postPage.Graphql.Media.Owner.Id != "" {
				mediaNode.MediaId = postPage.Graphql.Media.Id + "_" + postPage.Graphql.Media.Owner.Id
			} else {
				continue
			}

			// If we find something that exists, then we stop importing
			if !ignoreLast && MediaExistsInDb(mediaNode) {
				hasMorePages = false
				break
			} else {
				log.Printf("Adding %v to db\n", mediaNode.Code)
				AddToDb(mediaNode)
			}

		}

		// if hasMorePages == true, we check the tagResponse and the nextCursor
		if hasMorePages {
			hasMorePages = tagResponse.Tag.Media.PageInfo.HasNextPage
			nextCursor = tagResponse.Tag.Media.PageInfo.EndCursor
			if nextCursor == "" {
				hasMorePages = false;
			}
			log.Printf("Next: %v\n", nextCursor)
		}

		log.Printf("Has More Pages: %v\n", hasMorePages)
	}

}
func getTagResponse(tag string, nextCursor string) TagResponse {
	tagResponse := TagResponse{}
	log.Printf("Retrieving %v\n", fmt.Sprintf(IG_TAG_URL_FMT, tag, nextCursor))
	resp, err := http.Get(fmt.Sprintf(IG_TAG_URL_FMT, tag, nextCursor))
	if err != nil {
		log.Printf("Unable to get instagram tag. %v\n", err)
		return tagResponse
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Cannot read respnose. %v", err)
		return tagResponse
	}

	if resp.StatusCode != 200 {
		log.Printf("%v: %v\n%v", resp.StatusCode, resp.Status, string(responseBody))
		return tagResponse
	}

	json.Unmarshal([]byte(string(responseBody)), &tagResponse)

	return tagResponse
}

func getPostDetailsString(code string) string {

	resp, err := http.Get(fmt.Sprintf("https://www.instagram.com/p/%v/", code))
	if err != nil {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()

	tokens := html.NewTokenizer(resp.Body)

	// tokenize and iterate over HTML
	var scriptTag = false
	var endOfDoc = false
	for {
		tokenType := tokens.Next()

		switch tokenType {
		case html.ErrorToken:
			endOfDoc = true
		case html.StartTagToken:
			token := tokens.Token()
			if token.DataAtom == atom.Script {
				scriptTag = true
			}
		case html.EndTagToken:
			token := tokens.Token()
			if token.DataAtom == atom.Script {
				scriptTag = false
			}
		case html.TextToken:
			token := tokens.Token()
			var jsonBuffer bytes.Buffer

			// <script> content is window._sharedData = { ...Post Data Object... }
			if scriptTag && strings.Contains(token.String(), "window._sharedData") {
				splitString := strings.SplitN(token.String(), "=", 2)
				for i, jsonString := range splitString {
					if i == 0 {
						continue
					} else {
						jsonBuffer.WriteString(jsonString)
					}
				}

				postJsonString := jsonBuffer.String()
				postJsonString = strings.TrimRight(postJsonString, ";")

				return html.UnescapeString(postJsonString)

			}
		}

		if endOfDoc {
			break
		}
	}

	return ""

}

//func getLatestId(cfg config.AppConfig) {
//	_ := db.GetClient(cfg)
//}
