package facades

import "github.com/Firhan384/framework/contracts/filesystem"

func Storage() filesystem.Storage {
	return App().MakeStorage()
}
