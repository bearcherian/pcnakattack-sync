package main

import (
	"fmt"
	"github.com/bearcherian/pcnakattackSync/config"
	"github.com/bearcherian/pcnakattackSync/db"
	"github.com/bearcherian/pcnakattackSync/instagramSync"
	"github.com/bearcherian/pcnakattackSync/twitterSync"
	"github.com/bearcherian/pcnakattackSync/youtubeSync"
	"log"
	"os"
)

func main() {
	cfg := config.GetConfig()
	// init db connection
	db.GetClient()
	defer db.Close()

	logFile, err := os.OpenFile("pcnakattack.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Unable to open file")
		fmt.Println(err)
	}

	log.SetOutput(logFile)

	log.Println("Syncing Instagram...")
	instagramSync.SyncLatest()

	log.Println("Syncing Twitter...")
	twitterSync.SyncLatest(cfg)

	log.Println("Syncing YouTube...")
	youtubeSync.SyncLatest(cfg)

	log.Println("Done!")

}
