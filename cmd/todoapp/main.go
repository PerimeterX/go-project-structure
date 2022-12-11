package main

import (
	"github.com/caarlos0/env"
	"os"
	"os/signal"
	"syscall"
	"todoapp/internal/core/todo"
	"todoapp/internal/infra/http"
	"todoapp/internal/infra/sqlite"
)

var envVars = struct {
	DBFile   string `env:"DB_FILE" envDefault:"todo.db"`
	HTTPAddr string `env:"HTTP_ADDR" envDefault:":8080"`
}{}

func main() {
	// listen to os signals
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	// read env vars
	err := env.Parse(&envVars)

	// init adapter layer
	db, err := sqlite.NewTodoDB(envVars.DBFile)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// init core layer
	todoService := todo.NewService(db)

	// init port layer
	server, err := http.NewServer(envVars.HTTPAddr, todoService)
	if err != nil {
		panic(err)
	}
	defer server.Close()
	go server.Start()

	// wait for os signal
	<-c
}
