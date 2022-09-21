package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

var (
	version = "unknown"
)

func main() {
	// Load env-file if it exists first
	if env := os.Getenv("PLUGIN_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := cli.NewApp()
	app.Name = "docker-compose plugin"
	app.Usage = "docker-compose plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "docker.registry",
			Usage:  "docker registry",
			Value:  "https://index.docker.io/v1/",
			EnvVar: "PLUGIN_REGISTRY,DOCKER_REGISTRY",
		},
		cli.StringFlag{
			Name:   "docker.username",
			Usage:  "docker username",
			EnvVar: "PLUGIN_USERNAME,DOCKER_USERNAME",
		},
		cli.StringFlag{
			Name:   "docker.password",
			Usage:  "docker password",
			EnvVar: "PLUGIN_PASSWORD,DOCKER_PASSWORD",
		},
		cli.StringFlag{
			Name:   "docker.email",
			Usage:  "docker email",
			EnvVar: "PLUGIN_EMAIL,DOCKER_EMAIL",
		},
		cli.StringFlag{
			Name:   "docker.config",
			Usage:  "docker json dockerconfig content",
			EnvVar: "PLUGIN_CONFIG,DOCKER_PLUGIN_CONFIG",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Login: Login{
			Registry: c.String("docker.registry"),
			Username: c.String("docker.username"),
			Password: c.String("docker.password"),
			Email:    c.String("docker.email"),
			Config:   c.String("docker.config"),
		},
		Daemon: Daemon{
			Registry: c.String("docker.registry"),
		},
	}

	return plugin.Exec()
}
