package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sourcecode081017/auth-service-go/config"
	"github.com/sourcecode081017/auth-service-go/internal/db"
	"github.com/sourcecode081017/auth-service-go/internal/rest"
)

func main() {
	// start http server
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	mongoConn, err := db.NewMongo(
		ctx,
		cfg.Mongo.URI,
		cfg.Mongo.Database,
	)
	if err != nil {
		log.Fatalf("failed to connect to mongo: %v", err)
	}
	defer mongoConn.Close(ctx)
	restHandler := rest.New(mongoConn)
	router := restHandler.NewRouter(&cfg.App)
	go func() {
		if err := router.StartHttpServer(); err != nil {
			log.Fatalf("error starting http server: %v", err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = router.ShutDown(ctx)
}
