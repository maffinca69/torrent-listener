package main

import (
	"context"
	"fmt"
	"github.com/hekmon/transmissionrpc/v3"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"math"
	"sync"
	"time"
	"torrent-listener/cache"
	"torrent-listener/infrastructure"
	"torrent-listener/rate_limiter"
	"torrent-listener/telegram"
	"torrent-listener/transmission"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {
	setupCron()
	telegram.LongPulling()

	fmt.Scanln()
}

func setupCron() {
	c := cron.New()

	if _, err := c.AddFunc(infrastructure.Config().CronExpression, func() { checkTorrents() }); err != nil {
		panic("Error start schedule function")
	}

	c.Start()
}

func checkTorrents() {
	var wg sync.WaitGroup

	rateLimit := rate_limiter.NewLimiter(1*time.Second, 30)
	torrents, _ := transmission.Client().TorrentGetAll(context.TODO())
	wg.Add(len(torrents))

	for _, torrent := range torrents {
		go func(torrent transmissionrpc.Torrent) {
			defer wg.Done()
			notifyIfNeeded(torrent, rateLimit)
		}(torrent)
	}

	wg.Wait()
}

func notifyIfNeeded(torrent transmissionrpc.Torrent, limiter rate_limiter.Limiter) {
	fmt.Println("Checking hash:", *torrent.HashString)
	percentDone := uint8(math.RoundToEven(*torrent.PercentDone * 100))

	if !cache.Exists(torrent.HashString) {
		fmt.Println(fmt.Sprintf("Saving hash: %s", *torrent.HashString))
		cache.Store(percentDone, torrent.HashString)
		return
	}

	currentPercentComplete := cache.CompletePercent(torrent.HashString)
	if currentPercentComplete < 100 && percentDone == 100 {
		limiter.Wait()
		telegram.Notify(torrent)
		cache.Store(percentDone, torrent.HashString)
	}
}
