package facades

import (
	foundationcontract "github.com/Firhan384/framework/contracts/foundation"
	"github.com/Firhan384/framework/foundation"
)

func App() foundationcontract.Application {
	return foundation.App
}
