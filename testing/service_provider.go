package testing

import (
	contractsconsole "github.com/Firhan384/framework/contracts/console"
	"github.com/Firhan384/framework/contracts/foundation"
)

const Binding = "goravel.testing"

var artisanFacades contractsconsole.Artisan

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		return NewApplication(app), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	artisanFacades = app.MakeArtisan()
}
