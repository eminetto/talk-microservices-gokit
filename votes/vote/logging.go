package vote

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
)

func NewLoggingMiddleware(logger log.Logger, next Service) logmw {
	return logmw{logger, next}
}

type logmw struct {
	logger log.Logger
	Service
}

func (mw logmw) Store(v Vote) (id uuid.UUID, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "store",
			"input", v.TalkName,
			"output", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	id, err = mw.Service.Store(v)
	return
}
