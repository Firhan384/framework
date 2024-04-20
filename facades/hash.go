package facades

import (
	"github.com/Firhan384/framework/contracts/hash"
)

func Hash() hash.Hash {
	return App().MakeHash()
}
