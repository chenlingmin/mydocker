package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"mydocker/cgroups/subsystems"
	"mydocker/container"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: `create a container with namespace and cgroups limit mydocker run -ti [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		cli.BoolFlag{
			Name:  "d",
			Usage: "detach container",
		},
		cli.StringFlag{
			Name:  "v",
			Usage: "volume",
		},
		cli.StringFlag{
			Name: "name",
			Usage: "container name",
		},
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		var cmdArray []string
		for _, arg := range context.Args() {
			cmdArray = append(cmdArray, arg)
		}
		createTty := context.Bool("ti")
		detach := context.Bool("d")

		if createTty && detach {
			return fmt.Errorf("ti and d paramter can not both provided")
		}

		resConf := &subsystems.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuSet:      context.String("cpuset"),
			CpuShare:    context.String("cpushare"),
		}

		volume := context.String("v")
		containerName := context.String("name")
		Run(createTty, cmdArray, volume, containerName, resConf)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: `Init container process run user's process in container. Do not call it outside`,

	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		err := container.RunContainerInitProcess()
		return err
	},
}

var commitCommand = cli.Command{
	Name:  "commit",
	Usage: "commit a container into image",
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missage container name")
		}
		imageName := context.Args().Get(0)
		commitContainer(imageName)
		return nil
	},
}

var listCommand = cli.Command{
	Name: "ps",
	Usage: "list all the containers",
	Action: func(context *cli.Context) error {
		ListContainers()
		return nil
	},
}
