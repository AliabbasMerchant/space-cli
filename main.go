package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/spaceuptech/space-cli/config"
	"github.com/spaceuptech/space-cli/utils"
)

func main() {
	app := cli.NewApp()
	app.Version = utils.BuildVersion
	app.Name = "space-cli"
	app.Usage = "Runs the Space CLI"
	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "logs in the user",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "url",
					Usage: "the url of the cluster to login into",
				},
				cli.StringFlag{
					Name:   "user",
					Usage:  "the username to use for logging in",
					EnvVar: "SPACE_CLOUD_USER",
				},
				cli.StringFlag{
					Name:   "pass",
					Usage:  "the password to use for logging in",
					EnvVar: "SPACE_CLOUD_PASS",
				},
			},
			Action: func(c *cli.Context) error {
				return utils.Login(c.String("url"), c.String("user"), c.String("pass"))
			},
		},
		{
			Name:  "init",
			Usage: "creates a config file with sensible defaults",
			Action: func(c *cli.Context) error {
				return config.GenerateConfig()
			},
		},
		{
			Name:  "deploy",
			Usage: "deploys the space cloud instance",
			Action: func(c *cli.Context) error {
				return utils.Deploy()
			},
		},
		{
			Name:  "cluster",
			Usage: "allows to add, remove and edit cluster urls",
			Subcommands: []cli.Command{
				{
					Name:      "add",
					Usage:     "add a new cluster",
					UsageText: "space-cli cluster add [name - the new cluster name] [url - the cluster url]",
					Action: func(c *cli.Context) error {
						clusterName := c.Args().Get(0)
						url := c.Args().Get(1)
						return config.AddCluster(clusterName, url)
					},
				},
				{
					Name:      "remove",
					Usage:     "remove an existing cluster",
					UsageText: "space-cli cluster remove [name - the cluster to remove]",
					Action: func(c *cli.Context) error {
						return config.RemoveCluster(c.Args().Get(0))
					},
				},
				{
					Name:      "set",
					Usage:     "set the url of an existing cluster",
					UsageText: "space-cli cluster set [name - the cluster name] [url - the new cluster url]",
					Action: func(c *cli.Context) error {
						return config.SetClusterURL(c.Args().Get(0), c.Args().Get(1))
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
