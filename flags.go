package main

import (
	"os"
	"os/user"

	"github.com/urfave/cli/v2"
)

const version = "v1.0.0"

// NewApp creates an app with sane defaults.
func NewApp(usage string) *cli.App {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = version
	app.Usage = usage
	app.Copyright = "Copyright 2013-2026 The go-ethereum Authors"
	return app
}

func HomeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}
