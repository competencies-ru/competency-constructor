package main

import (
	"context"
	"flag"
	"log"
	"os/signal"
	"syscall"

	"github.com/competencies-ru/competency-constructor/internal/runner"
)

const defaultConfigPath = "configs"

func main() {
	flag.Parse()
	configPath := flag.Arg(0)

	if configPath == "" {
		configPath = defaultConfigPath
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := runner.New(configPath)
	go r.StartServer()
	<-ctx.Done()

	r.Stop()
	log.Println("Stop server")
}
