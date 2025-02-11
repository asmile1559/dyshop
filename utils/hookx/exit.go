package hookx

import (
	"os"
	"os/signal"
	"syscall"
)

func GracefulExit(hooks ...func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	<-sigs
	for _, hook := range hooks {
		hook()
	}
}
