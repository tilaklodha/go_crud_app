package main

import (
	"go_crud_app/appcontext"
	"go_crud_app/config"
	"go_crud_app/server"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func main() {
	config.Load()
	appcontext.Initiate()

	clientApp := cli.NewApp()
	clientApp.Name = "go-crud-app"
	clientApp.Commands = []cli.Command{
		{
			Name:        "start",
			Description: "Start HTTP Server",
			Action: func(c *cli.Context) error {
				router := server.Router()
				log.Fatal(http.ListenAndServe(":8632", router))
				return nil
			},
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
