package api

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/learning-microservice/account/domain"
)

type Middleware func(Service) Service

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

func (mw *loggingMiddleware) Health() string {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Health",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Health()
}

func (mw *loggingMiddleware) Login(username, password string) (a domain.Account, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Login",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return mw.next.Login(username, password)
}

func (mw *loggingMiddleware) Register(username, email, password string) (a domain.Account, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Register",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return mw.next.Register(username, email, password)
}
