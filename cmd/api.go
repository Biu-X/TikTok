package cmd

import (
	"github.com/urfave/cli/v2"
)

// CmdAPI api 子命令
var CmdAPI = &cli.Command{
	Name:        "api",
	Usage:       "Start TikTok api server",
	Description: `Star TikTok api server`,
	Action:      runAPI,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Value:   "3000",
			Usage:   "Temporary port number to prevent conflict",
		},
	},
}

func runAPI(ctx *cli.Context) error {
	return nil
}
