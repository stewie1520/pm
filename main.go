package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stewie1520/pm/data"
	"github.com/stewie1520/pm/helpers"
	"github.com/urfave/cli/v2"
	_ "gopkg.in/mgo.v2"
	"log"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "new",
				Usage:  "Create new password for a site",
				Action: newPassword,
				Before: requireForMasterPassword,
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

func requireForMasterPassword(c *cli.Context) error {
	require := helpers.RequireForMasterPassword()
	if !require {
		return nil
	}

	password, err := helpers.PromptForPassword("Please enter your master key to continue\n")
	if err != nil {
		return err
	}

	if !data.CheckMasterKey(password) {
		return fmt.Errorf("master key is incorrect")
	}

	return nil
}

func newPassword(c *cli.Context) error {
	siteName := c.String("site")
	fmt.Println(siteName)
	return nil
}
