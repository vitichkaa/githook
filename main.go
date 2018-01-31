package main

import (
	"context"
	"flag"
	"githook/api"
	"githook/conf"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env"
)

func main() {

	var (
		cnf     conf.Conf
		err     error
		dbSetup bool
	)

	flag.BoolVar(&dbSetup, "setup", false, "initial setup db")
	flag.Parse()
	if err = env.Parse(&cnf); err != nil {
		log.Fatal("configuration parsed with err", err)
	}

	//Init API
	ctx, cancel := context.WithCancel(context.Background())
	api := api.New(cnf)
	go api.Start()

	//System signal handling for graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	select {
	case <-c:
		log.Println("Interrupted", api.Stop())
	case <-ctx.Done():
		log.Println("Exited ", ctx.Err())
	}

	// will use for all cases for satisfy vet linter
	cancel()
}
