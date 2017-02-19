package main

import (
	"github.com/bearcherian/pcnakattackSync/config"
	"github.com/bearcherian/pcnakattackSync/db"
	"github.com/bearcherian/pcnakattackSync/twitterSync"
	"github.com/bearcherian/pcnakattackSync/youtubeSync"
	"log"
)

func main() {
	cfg := config.GetConfig()
	// init db connection
	db.GetClient(cfg)
	defer db.Close()

	log.Println("Syncing Twitter...")
	twitterSync.SyncLatest(cfg)

	log.Println("Syncing YouTube...")
	youtubeSync.SyncLatest(cfg)
}
