package facades

import (
	"github.com/Firhan384/framework/contracts/auth"
	"github.com/Firhan384/framework/contracts/auth/access"
)

func Auth() auth.Auth {
	return App().MakeAuth()
}

func Gate() access.Gate {
	return App().MakeGate()
}
