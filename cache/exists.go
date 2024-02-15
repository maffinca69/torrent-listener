package cache

import (
	"context"
	"torrent-listener/infrastructure"
)

func Exists(hash *string) bool {
	ctx := context.Background()

	client := infrastructure.RedisClient()

	key := *hash
	exists, err := client.Exists(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return exists == 1
}
