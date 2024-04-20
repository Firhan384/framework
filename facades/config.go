package facades

import (
	"github.com/Firhan384/framework/contracts/config"
)

func Config() config.Config {
	return App().MakeConfig()
}
