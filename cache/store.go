package cache

import (
	"context"
	"torrent-listener/infrastructure"
)

func Store(percentComplete uint8, hash *string) {
	ctx := context.Background()

	client := infrastructure.RedisClient()

	key := *hash
	if err := client.Set(ctx, key, percentComplete, 0).Err(); err != nil {
		panic("Error save version to cache: " + err.Error())
	}
}
