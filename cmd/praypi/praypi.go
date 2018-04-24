package main

import (
	"github.com/bbriggs/praypi"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {

	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Usage = "PrayPI: Prayer apps, redefined."
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Bren \"fraq\" Briggs",
			Email: "code@fraq.io",
		},
	}
	app.Action = func(c *cli.Context) error {
		var s *praypi.Server
		s.Run()
		waitForCtrlC()
		return nil
	}

	app.Run(os.Args)
}

func waitForCtrlC() {
	var end_waiter sync.WaitGroup
	end_waiter.Add(1)
	var signal_channel chan os.Signal
	signal_channel = make(chan os.Signal, 1)
	signal.Notify(signal_channel, os.Interrupt)
	go func() {
		<-signal_channel
		end_waiter.Done()
	}()
	end_waiter.Wait()
}
