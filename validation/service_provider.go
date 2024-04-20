package validation

import (
	consolecontract "github.com/Firhan384/framework/contracts/console"
	"github.com/Firhan384/framework/contracts/foundation"
	"github.com/Firhan384/framework/validation/console"
)

const Binding = "goravel.validation"

type ServiceProvider struct {
}

func (database *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewValidation(), nil
	})
}

func (database *ServiceProvider) Boot(app foundation.Application) {
	database.registerCommands(app)
}

func (database *ServiceProvider) registerCommands(app foundation.Application) {
	app.MakeArtisan().Register([]consolecontract.Command{
		&console.RuleMakeCommand{},
	})
}
