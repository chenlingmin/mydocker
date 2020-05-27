package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

const usage = `hello mydocker`

func main() {

	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = usage
	app.Commands = []cli.Command{
		initCommand,
		runCommand,
		commitCommand,
		listCommand,
		logCommand,
	}

	app.Before = func(context *cli.Context) error {
		log.SetReportCaller(true)
		log.SetFormatter(&log.TextFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
