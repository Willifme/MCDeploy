package mcdeploy

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func NewCLI() {

	app := cli.NewApp()

	app.Name = "MCDeploy"

	app.Usage = "A basic Minecraft package manager"

	app.Commands = []cli.Command{
		{
			Name:      "search",
			ShortName: "s",
			Usage:     "Search for a plugin",
			Action: func(c *cli.Context) {
				SearchPlugins(100, c.Args().First())
			},
		},
		{
			Name:      "install",
			ShortName: "i",
			Usage:     "Install a plugin",
			Action: func(c *cli.Context) {
				DownloadPlugin(100, c.Args().First())
			},
		},
		{
			Name:      "new",
			ShortName: "n",
			Usage:     "Create a new plugin",
			Subcommands: []cli.Command{
				{
					Name:  "vanilla",
					Usage: "Make a vanilla server instead of a bukkit one",
					Action: func(c *cli.Context) {
						fmt.Println("something here soon!")
					},
				},
			},
			Action: func(c *cli.Context) {
				NewBukkitServer()
			},
		},
	}

	app.Run(os.Args)

}
