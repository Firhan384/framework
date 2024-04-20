package testing

import (
	"github.com/Firhan384/framework/contracts/foundation"
	"github.com/Firhan384/framework/contracts/testing"
	"github.com/Firhan384/framework/testing/docker"
)

type Application struct {
	app foundation.Application
}

func NewApplication(app foundation.Application) *Application {
	return &Application{
		app: app,
	}
}

func (receiver *Application) Docker() testing.Docker {
	return docker.NewDocker(receiver.app)
}
