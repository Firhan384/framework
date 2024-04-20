package facades

import (
	"github.com/Firhan384/framework/contracts/testing"
)

func Testing() testing.Testing {
	return App().MakeTesting()
}
