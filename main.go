package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	_ "gopkg.in/mgo.v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "new",
				Usage:  "Create new password for a site",
				Action: newPassword,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "site",
						Aliases:  []string{"s"},
						Usage:    "the password used in site",
						Required: true,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func requireForMasterPassword() bool {
	return false
}

func newPassword(c *cli.Context) error {
	siteName := c.String("site")
	fmt.Println(siteName)
	return nil
}
