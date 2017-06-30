package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// JSON Configuration Options
type DrunConfig struct {
	Image          string `json:"image"`
	DefaultCommand string `json:"defaultCommand"`
	Net            string `json:"net"`
}

func main() {
	app := cli.NewApp()
	app.Name = "drun"
	app.Usage = "Docker Runner"
	app.Version = "0.0.2"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "image",
			Value: "",
			Usage: "Docker Image to run",
		},
		cli.StringFlag{
			Name:  "net",
			Value: "",
			Usage: "Network setting (see --net in docker run), defaults to host",
		},
	}
	app.Action = func(c *cli.Context) error {
		ex, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error finding path", err)
			os.Exit(1)
		}

		raw, err := ioutil.ReadFile("./drun.json")
		// Config file exists, parse it out
		var config DrunConfig
		if err == nil {
			json.Unmarshal(raw, &config)
		}

		volume := ex + ":/app"
		args := c.Args()
		cmdInput := []string{}
		if len(config.DefaultCommand) > 0 {
			cmdInput = strings.Split(config.DefaultCommand, " ")
		}
		// Override default command if specified
		if len(args) > 0 {
			cmdInput = args
		}
		if len(cmdInput) == 0 {
			fmt.Fprintln(os.Stderr, "No command specified")
			os.Exit(1)
		}
		command := cmdInput

		var image string
		if len(config.Image) > 0 {
			image = config.Image
		}
		// Override config if specified in flag
		if len(c.String("image")) > 0 {
			image = c.String("image")
		}

		net := "host"
		if len(config.Net) > 0 {
			net = config.Net
		}
		// Override config if specified in flag
		if len(c.String("net")) > 0 {
			net = c.String("net")
		}
		cmdName := "docker"
		cmdArgs := []string{
			"run",
			"--rm",
			"-w", "/app",
			"--net", net,
			"-v", volume,
			image}

		cmdArgs = append(cmdArgs, command...)

		cmd := exec.Command(cmdName, cmdArgs...)
		cmdReader, err := cmd.StdoutPipe()
		cmdReaderErr, err := cmd.StderrPipe()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
			os.Exit(1)
		}

		scanner := bufio.NewScanner(cmdReader)
		go func() {
			for scanner.Scan() {
				fmt.Printf("%s\n", scanner.Text())
			}
		}()

		errScanner := bufio.NewScanner(cmdReaderErr)
		go func() {
			for errScanner.Scan() {
				fmt.Printf("%s\n", errScanner.Text())
			}
		}()

		err = cmd.Start()
		if err != nil {
			fmt.Printf("Error with start: %s\n", err)
			os.Exit(1)
		}

		err = cmd.Wait()
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		return nil
	}
	app.Run(os.Args)
}
