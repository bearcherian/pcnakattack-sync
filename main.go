package main

import (
	"github.com/bearcherian/pcnakattackSync/config"
	"github.com/bearcherian/pcnakattackSync/twitterSync"
	"github.com/bearcherian/pcnakattackSync/youtubeSync"
	"github.com/bearcherian/pcnakattackSync/db"
)

func main() {
	cfg := config.GetConfig()
	// init db connection
	db.GetClient(cfg)
	defer db.Close()

	twitterSync.SyncLatest(cfg)

	youtubeSync.SyncLatest(cfg)
}
