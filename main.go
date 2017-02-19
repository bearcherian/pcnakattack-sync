package main

import (
	"github.com/bearcherian/pcnakattackSync/config"
	"github.com/bearcherian/pcnakattackSync/db"
	"github.com/bearcherian/pcnakattackSync/twitterSync"
	"github.com/bearcherian/pcnakattackSync/youtubeSync"
	"log"
	"os"
	"fmt"
)

func main() {
	cfg := config.GetConfig()
	// init db connection
	db.GetClient(cfg)
	defer db.Close()

	logFile, err := os.OpenFile("pcnakattack.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Unable to open file")
		fmt.Println(err)
	}

	log.SetOutput(logFile)

	log.Println("Syncing Twitter...")
	twitterSync.SyncLatest(cfg)

	log.Println("Syncing YouTube...")
	youtubeSync.SyncLatest(cfg)
}
