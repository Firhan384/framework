package facades

import (
	"github.com/Firhan384/framework/contracts/cache"
)

func Cache() cache.Cache {
	return App().MakeCache()
}
