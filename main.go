package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"dailydo.fe1.xyz/internal/app"
	"dailydo.fe1.xyz/internal/globals"
	"dailydo.fe1.xyz/internal/migration"

	"go.uber.org/zap"
)

func main() {
	// init logging
	globals.InitLogging()
	// load configuration
	globals.MustLoad()

	// init database
	globals.InitDatabase()

	// migration database
	migration.MustMigration()
	// init redis
	globals.InitRedis()
	// register router
	srv := app.NewServer(globals.CONF)
	// startup
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		if err := srv.Listen(fmt.Sprintf(":%d", globals.CONF.Server.Port)); err != nil {
			globals.LOG.Fatal("shutting down the server error", zap.Error(err))
		}
	}()

	<-ctx.Done()
	if err := srv.Shutdown(); err != nil {
		globals.LOG.Fatal("shutdown server error", zap.Error(err))
	}
}
