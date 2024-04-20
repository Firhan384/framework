package facades

import (
	"github.com/Firhan384/framework/contracts/console"
)

func Artisan() console.Artisan {
	return App().MakeArtisan()
}
