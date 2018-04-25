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
	var (
		dbHost string
		dbName string
		dbPass string
		dbPort string
		dbUser string
	)

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
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "dbUser",
			Value:       "postgres",
			Usage:       "Postgres user",
			Destination: &dbUser,
		},
		cli.StringFlag{
			Name:        "dbPass",
			Usage:       "Postgres password",
			Destination: &dbPass,
		},
		cli.StringFlag{
			Name:        "dbName",
			Value:       "postgres",
			Usage:       "Postgres database name",
			Destination: &dbName,
		},
		cli.StringFlag{
			Name:        "dbPort",
			Value:       "5432",
			Usage:       "Postgres port",
			Destination: &dbPort,
		},
	}
	app.Action = func(c *cli.Context) error {
		var s *praypi.Server
		s.Run(dbUser, dbPass, dbName, dbPort)
		waitForCtrlC()
		return nil
	}
	sort.Sort(cli.FlagsByName(app.Flags))
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
