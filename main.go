package main

import (
	"admin-panel/v1/cmd"
	"github.com/urfave/cli"
	"log"
	"os"
)

var (
	client *cli.App
)

func init() {
	// Initialise a CLI app
	client = cli.NewApp()
	client.Name = "lido core api"
	client.Usage = "lido core api worker and handler"
	client.Version = "0.0.1"
}

func main() {

	client.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "start api server",
			Action: func(c *cli.Context) error {
				log.Printf("start %s\n", c.Command.Name)
				cmd.StartServer()
				return nil
			},
		},
	}

	// Run the CLI app
	client.Run(os.Args)
}
