package user

import (
	"time"

	"github.com/go-kit/kit/log"
)

func NewLoggingMiddleware(logger log.Logger, next Service) logmw {
	return logmw{logger, next}
}

type logmw struct {
	logger log.Logger
	Service
}

func (mw logmw) ValidateUser(email, password string) (err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "validateUser",
			"input", email,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Service.ValidateUser(email, password)
	return
}
