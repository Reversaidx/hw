package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Reversaidx/hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/Reversaidx/hw/hw12_13_14_15_calendar/internal/server/http"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "configs/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}
	config := NewConfig(configFile)
	logg, err := logger.New(config.Logger.Level)
	if err != nil {
		panic(err)
	}
	// var storage interface{}
	// if config.Storage.Type == "sql" {
	//	storage = sqlstorage.New()
	//	storage.Cone
	// } else if config.Storage.Type == "memory" {
	//	storage = memorystorage.New()
	// } else {
	//	logg.Error("Unsupported type")
	// }
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	// storage := memorystorage.New()
	// calendar := app.New(logg, storage)

	server := internalhttp.Server{}
	server.Start(ctx)
	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
