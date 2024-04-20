package facades

import (
	"github.com/Firhan384/framework/contracts/route"
)

func Route() route.Route {
	return App().MakeRoute()
}
