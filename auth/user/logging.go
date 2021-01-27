package user

import (
	"context"
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

func (mw logmw) ValidateUser(ctx context.Context, email, password string) (token string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "validateUser",
			"input", email,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	token, err = mw.Service.ValidateUser(ctx, email, password)
	return
}

func (mw logmw) ValidateToken(ctx context.Context, token string) (email string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "validateToken",
			"input", token,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	email, err = mw.Service.ValidateToken(ctx, token)
	return
}
