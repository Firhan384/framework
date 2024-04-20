package facades

import (
	"github.com/Firhan384/framework/contracts/queue"
)

func Queue() queue.Queue {
	return App().MakeQueue()
}
