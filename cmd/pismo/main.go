package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"pismo-back-teste/internal"
	"pismo-back-teste/internal/api"
	"pismo-back-teste/internal/config"
	"pismo-back-teste/internal/database"
	"syscall"

	"github.com/spf13/viper"
)

func main() {
	config.Load()
	config.InitLog()

	db, err := config.InitDatabase()
	if err != nil {
		panic(err)
	}

	repoProvider := database.NewRepositoryProvider()
	serviceProvider := internal.NewServiceProvider(repoProvider)
	e := api.NewServer()
	opts := api.Options{
		ServiceProvider:  serviceProvider,
		DatabaseProvider: db,
	}

	api.InitRoutes(e, opts)

	errc := make(chan error, 3)
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := e.Start(viper.GetString("http.listen")); err != nil {
			errc <- err
		}
	}()

	fmt.Println("starting")
	select {
	case <-quit:
		ctx := context.Background()
		if err := e.Shutdown(ctx); err != nil {
			fmt.Println(err)
		}
	case err := <-errc:
		fmt.Println(err)
		os.Exit(1)
	}
}
