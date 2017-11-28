package main

import (
	"fork/tools/log"
	"os"
	"os/signal"
	"syscall"

	"github.com/FlyCynomys/nunit/server"
)

func main() {
	log.Info("nginx-unit module")
	server.Ng()
	sc := make(chan os.Signal, 1)

	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-sc
}
