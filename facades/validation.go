package facades

import (
	"github.com/Firhan384/framework/contracts/validation"
)

func Validation() validation.Validation {
	return App().MakeValidation()
}
