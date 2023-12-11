package main

import (
	"admin-panel/v1/cmd"
	"admin-panel/v1/utils"
	"fmt"
	"github.com/gin-gonic/gin"
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
	gin.SetMode(gin.DebugMode)
}

func main() {
	client.Commands = []cli.Command{
		{
			Name:  "worker",
			Usage: "launch machinery worker",
			Action: func(c *cli.Context) error {
				log.Printf("start %s\n", c.Args().First())
				consume := fmt.Sprintf("core-api-consume:%s", utils.NewID())
				log.Printf("consume %s\n", consume)
				if err := cmd.WorkerExecute(c.Args().First(), consume, 12); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
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
