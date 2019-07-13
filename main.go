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
				// Load config from file
				conf, err := config.LoadConfigFromFile(utils.DefaultConfigFilePath)
				if err != nil {
					return err
				}
				return utils.Deploy(conf)
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
		{
			Name:  "env",
			Usage: "allows to add, remove and edit environment variables",
			Subcommands: []cli.Command{
				{
					Name:      "set",
					Usage:     "set an environment variable",
					UsageText: "space-cli env set [name - the environment variable name] [value - the variable value]",
					Action: func(c *cli.Context) error {
						return config.SetEnvVar(c.Args().Get(0), c.Args().Get(1))
					},
				},
				{
					Name:      "remove",
					Usage:     "remove an existing environment variable",
					UsageText: "space-cli env remove [name - the environment variable to remove]",
					Action: func(c *cli.Context) error {
						return config.RemoveEnvVar(c.Args().Get(0))
					},
				},
			},
		},
		{
			Name:  "port",
			Usage: "allows to add, remove and edit port details",
			Subcommands: []cli.Command{
				{
					Name:      "add",
					Usage:     "add new port information",
					UsageText: "space-cli port add [name - a name to give to the port]",
					Action: func(c *cli.Context) error {
						return config.AddPort(c.Args().Get(0))
					},
				},
				{
					Name:      "remove",
					Usage:     "remove existing port information",
					UsageText: "space-cli env remove [name - the port name]",
					Action: func(c *cli.Context) error {
						return config.RemovePort(c.Args().Get(0))
					},
				},
				{
					Name:      "edit",
					Usage:     "edit existing port information",
					UsageText: "space-cli env edit [name - the port name]",
					Action: func(c *cli.Context) error {
						return config.EditPort(c.Args().Get(0))
					},
				},
			},
		},
		{
			Name:  "constraints",
			Usage: "allows to set the runtime constraints",
			Subcommands: []cli.Command{
				{
					Name:      "set-replicas",
					Usage:     "set the number of replicas",
					UsageText: "space-cli constraints set replicas [replicas - the number of replicas]",
					Action: func(c *cli.Context) error {
						return config.SetReplicas(c.Args().Get(0))
					},
				},
				{
					Name:      "set-cpu",
					Usage:     "set the cpu limit",
					UsageText: "space-cli constraints set cpu [cpu - the cpu limit]",
					Action: func(c *cli.Context) error {
						return config.SetCPU(c.Args().Get(0))
					},
				},
				{
					Name:      "set-memory",
					Usage:     "set the memory limit",
					UsageText: "space-cli constraints set memory [memory - the memory limit]",
					Action: func(c *cli.Context) error {
						return config.SetMemory(c.Args().Get(0))
					},
				},
			},
		},
		{ Name:  "expose",
			Usage: "allows to add an exposed port and its proxy information",
			Subcommands: []cli.Command{
				{
					Name:      "add",
					Usage:     "add an exposed port and its proxy information",
					UsageText: "space-cli expose add",
					Action: func(c *cli.Context) error {
						return config.AddExpose()
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
