package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/learning-microservice/account/api"
	"github.com/learning-microservice/account/infrastructure/memory"
)

const (
	defaultPort = "18080"
)

func main() {
	var (
		addr     = envString("PORT", defaultPort)
		httpAddr = flag.String("http.addr", ":"+addr, "HTTP listen address")
	)
	flag.Parse()

	// Logger
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	}

	// Repository
	var (
		accounts = memory.NewAccountRepository()
	)

	// Service
	var service api.Service
	{
		service = api.NewService(accounts)
		service = api.LoggingMiddleware(logger)(service)
	}

	// Http-logger
	httpLogger := log.With(logger, "component", "http")

	// Http-handler
	var router http.Handler
	{
		router = api.MakeHandler(service, httpLogger)
	}

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, router)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
