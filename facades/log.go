package facades

import (
	"github.com/Firhan384/framework/contracts/log"
)

func Log() log.Log {
	return App().MakeLog()
}
