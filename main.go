package main

import (
	"github.com/pluveto/coin-bot/pkg/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	done := make(chan bool, 1)
	handleSignal(done)
	c := &Config{}
	config.LoadConfigN(c, "./configs", "bot")
	bot := NewBot(c)
	bot.Run()
	<-done
	print("See you\n")
}

func handleSignal(stop chan bool) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		stop <- true
	}()
}
