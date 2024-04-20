package cache

import (
	"github.com/Firhan384/framework/contracts/config"
)

func prefix(config config.Config) string {
	return config.GetString("cache.prefix") + ":"
}
