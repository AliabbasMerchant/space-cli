package main

import (
  "log"
  // "fmt"
  "os"
  "errors"
  "strings"
	"path/filepath"

  "github.com/urfave/cli"
  "gopkg.in/AlecAivazis/survey.v1"

  "github.com/spaceuptech/space-cli/utils"
  "github.com/spaceuptech/space-cli/config"
)

var confg *config.Deploy

func actionInit(c *cli.Context) error {
  confg = &config.Deploy{}
  err := survey.AskOne(&survey.Input{Message: "name:"}, &confg.Name, survey.Required)
	if err != nil {
		return err
  }
  err = survey.AskOne(&survey.Input{Message: "project name:"}, &confg.Project, survey.Required)
	if err != nil {
		return err
  }
  err = survey.AskOne(&survey.Input{Message: "working directory:", Default:"."}, &confg.WorkingDir, survey.Required)
	if err != nil {
		return err
  }
  err = survey.AskOne(&survey.Input{Message: "ignore file:", Default:".gitignore"}, &confg.Ignore, survey.Required)
	if err != nil {
		return err
  }
  confg.Ignore = strings.TrimSpace(confg.Ignore)
  runtime := config.Runtime{}
  confg.Runtime = &runtime
  err = survey.AskOne(&survey.Select{
		Message: "runtime type:",
		Options: []string{utils.Python3, utils.Java11, utils.Golang, utils.NodeJS},
	}, &runtime.Name, survey.Required)
	if err != nil {
		return err
  }
  switch(runtime.Name) {
  case utils.Python3:
    err = survey.AskOne(&survey.Input{Message: "command to install:", Default:"pip install requirements.txt"}, &runtime.Install, survey.Required)
    if err != nil {
      return err
    }
    err = survey.AskOne(&survey.Input{Message: "command to run:", Default:"python main.py"}, &runtime.Run, survey.Required)
    if err != nil {
      return err
    }
  case utils.NodeJS:
    err = survey.AskOne(&survey.Input{Message: "command to install:", Default:"npm install"}, &runtime.Install, survey.Required)
    if err != nil {
      return err
    }
    err = survey.AskOne(&survey.Input{Message: "command to run:", Default:"npm start"}, &runtime.Run, survey.Required)
    if err != nil {
      return err
    }
  case utils.Java11:
    // err = survey.AskOne(&survey.Input{Message: "command to install:", Default:"pip install requirements.txt"}, &runtime.Install, survey.Required)
    // if err != nil {
    //   return err
    // }
    runtime.Install = ""
    err = survey.AskOne(&survey.Input{Message: "command to run (please make a jar file):", Default:"java jarfile.jar"}, &runtime.Run, survey.Required)
    if err != nil {
      return err
    }
  case utils.Golang:
    err = survey.AskOne(&survey.Input{Message: "command to setup:", Default:"go build"}, &runtime.Install, survey.Required)
    if err != nil {
      return err
    }
    err = survey.AskOne(&survey.Input{Message: "command to run:", Default:"./executable"}, &runtime.Run, survey.Required)
    if err != nil {
      return err
    }
  }
  // ports := ""
  // err = survey.AskOne(&survey.Multiline{Message: "ports to expose (eg 80:80/tcp)",}, &ports, survey.MinLength(0))
  // if err != nil {
  //   return err
  // }
  // if ports != "" {
  //   confg.Ports = strings.Split(ports, "\n")
  // }
  constraints := config.Constraints{}
  confg.Constraints = &constraints
  err = survey.AskOne(&survey.Input{Message: "replicas", Default: "1"}, &constraints.Replicas, survey.Required)
  if err != nil {
    return err
  }
  var cpu float32 = 1.0
  err = survey.AskOne(&survey.Input{Message: "cpu limit", Default: "1.0"}, &cpu, survey.Required)
  if err != nil {
    return err
  }
  constraints.CPU = &cpu

  var memory int64 = 500
  err = survey.AskOne(&survey.Input{Message: "memory limit", Default: "500"}, &memory, survey.Required)
  if err != nil {
    return err
  }
  constraints.Memory = &memory
  confg.Env = map[string]string{"exampleEnvVariable":"exampleValue"}
  // confg.Clusters = make(map[string]string)

  return config.StoreConfigToFile(confg, utils.DefaultConfigFilePath)
}

func actionDeploy(c *cli.Context) error {
	if confg == nil {
    return errors.New("Space CLI not yet initialized!")
  }
  var ignore *utils.Ignore
  var err error
  if confg.Ignore == "" {
    ignore, err = utils.InitIgnoreFromText("")
    if err != nil {
      log.Println(err)
      return err
    }
  } else {
    ignore, err = utils.InitIgnore(confg.Ignore)
    if err != nil {
      log.Println(err)
      return err
    }
  }

  var files []string
  os.Chdir(confg.WorkingDir)
  err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    if !info.IsDir() {
      if ignore.ToBeIncluded(path) {
        files = append(files, path)
        // log.Println(path)
      }
    }
    return nil
  })
  if err != nil {
    return err
  }
  ignore.Close()

  err = utils.ZipFiles(utils.ZipName, files)
  if err != nil {
    return err
  }

  for name, url := range confg.Clusters {
    json, err := config.GetJSON(confg)
    if err != nil {
      return err
    }
    err = utils.SendToCluster(url+utils.Temporary, utils.ZipName, json)
    if err != nil {
      log.Println(err)
    } else {
      log.Println("Successfully deployed to cluster: " + name)
    }
  }
  os.Remove(utils.ZipName)
  return nil
}

func main() {
	app := cli.NewApp()
	app.Version = utils.BuildVersion
	app.Name = "space-cli"
	app.Usage = "Runs the Space CLI"
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "creates a config file with sensible defaults",
			Action: actionInit,
		},
		{
			Name:   "deploy",
			Usage:  "deploys the space cloud instance",
			Action: actionDeploy,
		},
		{
			Name:   "cluster",
			Usage:  "allows to add, remove and edit cluster urls",
			Subcommands: []cli.Command{
				{
				  Name:  "add",
				  Usage: "add a new cluster",
          UsageText: "space-cli cluster add [name - the new cluster name] [url - the cluster url]",
				  Action: func(c *cli.Context) error {
            if confg == nil {
              return errors.New("Space CLI not yet initialized!")
            }
            if len(c.Args().Get(0)) == 0 {
              return errors.New("Cluster Name Cannot be Empty")
            }
            if len(c.Args().Get(1)) == 0 {
              return errors.New("Cluster URL Cannot be Empty")
            }
            return confg.AddCluster(c.Args().Get(0), c.Args().Get(1))
				  },
				},
				{
				  Name: "remove",
          Usage: "remove an existing cluster",
          UsageText: "space-cli cluster remove [name - the cluster to remove]",
				  Action: func(c *cli.Context) error {
            if confg == nil {
              return errors.New("Space CLI not yet initialized!")
            }
            if len(c.Args().Get(0)) == 0 {
              return errors.New("Cluster Name Cannot be Empty")
            }
            return confg.RemoveCluster(c.Args().Get(0))
				  },
				},
				{
				  Name:  "set",
				  Usage: "set the url of an existing cluster",
          UsageText: "space-cli cluster set [name - the cluster name] [url - the new cluster url]",
				  Action: func(c *cli.Context) error {
            if confg == nil {
              return errors.New("Space CLI not yet initialized!")
            }
            if len(c.Args().Get(0)) == 0 {
              return errors.New("Cluster Name Cannot be Empty")
            }
            if len(c.Args().Get(1)) == 0 {
              return errors.New("Cluster URL Cannot be Empty")
            }
            return confg.SetClusterURL(c.Args().Get(0), c.Args().Get(1))
				  },
				},
		  },
		},
  }
  
  var err error
  confg, err = config.LoadConfigFromFile(utils.DefaultConfigFilePath)
  if err != nil {
    log.Fatal(err)
  }

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}