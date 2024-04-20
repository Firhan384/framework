package facades

import (
	"github.com/Firhan384/framework/contracts/database/orm"
)

func Orm() orm.Orm {
	return App().MakeOrm()
}
