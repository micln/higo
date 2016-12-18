package middlewares

import (
	"log"
	"time"

	"github.com/micln/higo/context"
)

type LogEntity struct {
	startAt time.Time
	endAt   time.Time
}

func (lg *LogEntity) Before(ctx *context.Context) {
	lg.startAt = time.Now()
}

func (lg *LogEntity) After(ctx *context.Context) {
	lg.endAt = time.Now()
	log.Printf(
		"%v %vms %s",
		ctx.Resp.GetStatusCode(),
		float64(lg.endAt.Sub(lg.startAt).Nanoseconds())/1000,
		ctx.Req.URL,
	)
}
