package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/go-errors/errors"
	"github.com/oklog/oklog/pkg/group"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "stratus"

	app.Commands = []cli.Command{{
		Name:   "dev",
		Usage:  "Runs gitpods on you local development machine",
		Action: actionDev,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "addr-ui", Usage: "The address to run the UI on", Value: ":3000"},
			cli.StringFlag{Name: "addr-api", Usage: "The address to run the API on", Value: ":3010"},
			cli.StringFlag{Name: "env", Usage: "Set the env gitpods runs in", Value: "development"},
			cli.StringFlag{Name: "log-level", Usage: "The log level to filter logs with before printing", Value: "debug"},
			cli.BoolFlag{Name: "watch,w", Usage: "Watch files in this project and rebuild binaries if something changes"},
		},
	}}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func actionDev(c *cli.Context) error {
	uiAddrFlag := c.String("addr-ui")
	apiAddrFlag := c.String("addr-api")
	envFlag := c.String("env")
	loglevelFlag := c.String("log-level")
	watch := c.Bool("watch")

	uiRunner := NewGitPodsRunner("ui", []string{
		fmt.Sprintf("GITPODS_ADDR=%s", uiAddrFlag),
		fmt.Sprintf("GITPODS_ADDR_API=%s", apiAddrFlag),
		fmt.Sprintf("GITPODS_ENV=%s", envFlag),
		fmt.Sprintf("GITPODS_LOGLEVEL=%s", loglevelFlag),
	})

	apiRunner := NewGitPodsRunner("api", []string{
		fmt.Sprintf("GITPODS_ADDR=%s", apiAddrFlag),
		fmt.Sprintf("GITPODS_ENV=%s", envFlag),
		fmt.Sprintf("GITPODS_LOGLEVEL=%s", loglevelFlag),
	})

	if watch {
		watcher := &FileWatcher{}
		watcher.Add(uiRunner, apiRunner)

		go watcher.Watch()
	}

	var g group.Group
	{
		g.Add(func() error {
			log.Println("starting ui")
			return uiRunner.Run()
		}, func(err error) {
			log.Println("stopping ui")
			uiRunner.Stop()
		})
	}
	{
		g.Add(func() error {
			log.Println("starting api")
			return apiRunner.Run()
		}, func(err error) {
			log.Println("stopping api")
			apiRunner.Stop()
		})
	}
	{
		g.Add(func() error {
			stop := make(chan os.Signal, 1)
			signal.Notify(stop, os.Interrupt)
			<-stop
			return errors.New("stopping stratus")
		}, func(err error) {
		})
	}

	webpack := &WebpackRunner{}
	if watch {
		g.Add(func() error {
			log.Println("starting webpack")
			return webpack.Run(true)
		}, func(err error) {
			log.Println("stopping webpack")
			webpack.Stop()
		})
	} else {
		webpack.Run(false)
	}

	return g.Run()
}

//// RunAPI runs a development server and restarts it with a new build if files change.
//func RunAPI(env []string) func() error {
//	return func() error {
//		builds := make(chan bool)
//
//		go BuildForever(builds)
//
//		go func() {
//			if err := build(); err == nil {
//				builds <- true
//			}
//		}()
//
//		var cmd *exec.Cmd
//		for {
//			<-builds
//			if cmd != nil {
//				cmd.Process.Kill()
//			}
//
//			cmd = exec.Command("./dist/gitpods")
//			go func() {
//				cmd.Env = env
//				cmd.Stdin = os.Stdin
//				cmd.Stdout = os.Stdout
//				cmd.Stderr = os.Stderr
//				if err := cmd.Run(); err != nil {
//					log.Println(err)
//					return
//				}
//			}()
//		}
//
//		return nil
//	}
//}
