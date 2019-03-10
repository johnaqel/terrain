package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	resFlag := true
	outFlag := "text"

	app := cli.NewApp()
	app.Name = "terrain"
	app.Usage = "A terraform documentation utility"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "resources, r",
			Usage:       "List the Terraform resources declared",
			Destination: &resFlag,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "Output format: [text, json, markdown]",
			Value:       "text",
			Destination: &outFlag,
		},
	}

	app.Action = func(c *cli.Context) error {
		if resFlag {
			fmt.Println("Resources will be listed")
		}

		switch outFlag {
		case "json":
			fmt.Println("Gonna output json")
		case "markdown":
			fmt.Println("Gonna output markdown")
		default: // Todo: should this error, or just default to text?
			fmt.Println("Gonna output text")
		}

		return nil
	}

	_ = app.Run(os.Args)

}
