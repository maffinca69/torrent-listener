package cache

import (
	"context"
	"strconv"
	"torrent-listener/infrastructure"
)

func CompletePercent(hash *string) uint8 {
	ctx := context.Background()

	client := infrastructure.RedisClient()
	key := *hash
	var value, _ = client.Get(ctx, key).Result()

	uint64Value, _ := strconv.ParseUint(value, 10, 8)

	return uint8(uint64Value)
}
