package main

import (
	"github.com/bearcherian/pcnakattackSync/config"
	"github.com/bearcherian/pcnakattackSync/twitterSync"
	"github.com/bearcherian/pcnakattackSync/youtubeSync"
)

func main() {
	cfg := config.GetConfig()

	twitterSync.SyncLatest(cfg)

	youtubeSync.SyncLatest(cfg)
}
