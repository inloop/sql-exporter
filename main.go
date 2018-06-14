package main

import (
	"os"

	"github.com/inloop/sql-exporter/cmd"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "sql-exporter"
	app.Usage = "Command line utility to export data from SQL database"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		cmd.Dump(),
	}

	app.Run(os.Args)
}
