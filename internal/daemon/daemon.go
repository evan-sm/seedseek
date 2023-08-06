package daemon

import (
	"context"
	"time"

	"seedseek/pkg/logger"
)

type daemon struct {
	log logger.Logger
}

type Daemon interface {
	Run(context.Context)
}

func New(log logger.Logger) Daemon {
	return &daemon{
		log: log,
	}
}

func (d *daemon) Run(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)

loop:
	for {
		select {
		case <-ticker.C:
			d.log.InfoContext(ctx, "daemon run tick")
		case <-ctx.Done():
			d.log.ErrorContext(ctx, "daemon close: %s", context.Cause(ctx).Error())

			break loop
		}
	}
}
