package facades

import "github.com/Firhan384/framework/contracts/event"

func Event() event.Instance {
	return App().MakeEvent()
}
