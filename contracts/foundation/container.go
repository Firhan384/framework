package foundation

import (
	"github.com/Firhan384/framework/contracts/auth"
	"github.com/Firhan384/framework/contracts/auth/access"
	"github.com/Firhan384/framework/contracts/cache"
	"github.com/Firhan384/framework/contracts/config"
	"github.com/Firhan384/framework/contracts/console"
	"github.com/Firhan384/framework/contracts/crypt"
	"github.com/Firhan384/framework/contracts/database/orm"
	"github.com/Firhan384/framework/contracts/database/seeder"
	"github.com/Firhan384/framework/contracts/event"
	"github.com/Firhan384/framework/contracts/filesystem"
	"github.com/Firhan384/framework/contracts/grpc"
	"github.com/Firhan384/framework/contracts/hash"
	"github.com/Firhan384/framework/contracts/http"
	"github.com/Firhan384/framework/contracts/log"
	"github.com/Firhan384/framework/contracts/mail"
	"github.com/Firhan384/framework/contracts/queue"
	"github.com/Firhan384/framework/contracts/route"
	"github.com/Firhan384/framework/contracts/schedule"
	"github.com/Firhan384/framework/contracts/testing"
	"github.com/Firhan384/framework/contracts/validation"
)

type Container interface {
	Bind(key any, callback func(app Application) (any, error))
	BindWith(key any, callback func(app Application, parameters map[string]any) (any, error))
	Instance(key, instance any)
	Make(key any) (any, error)
	MakeArtisan() console.Artisan
	MakeAuth() auth.Auth
	MakeCache() cache.Cache
	MakeConfig() config.Config
	MakeCrypt() crypt.Crypt
	MakeEvent() event.Instance
	MakeGate() access.Gate
	MakeGrpc() grpc.Grpc
	MakeHash() hash.Hash
	MakeLog() log.Log
	MakeMail() mail.Mail
	MakeOrm() orm.Orm
	MakeQueue() queue.Queue
	MakeRateLimiter() http.RateLimiter
	MakeRoute() route.Route
	MakeSchedule() schedule.Schedule
	MakeStorage() filesystem.Storage
	MakeTesting() testing.Testing
	MakeValidation() validation.Validation
	MakeView() http.View
	MakeSeeder() seeder.Facade
	MakeWith(key any, parameters map[string]any) (any, error)
	Singleton(key any, callback func(app Application) (any, error))
}
