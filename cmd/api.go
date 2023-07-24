package cmd

import (
	"github.com/urfave/cli/v2"
)

// CmdApi api 子命令
var CmdApi = &cli.Command{
	Name:        "api",
	Usage:       "Start TikTok api server",
	Description: `Star TikTok api server`,
	Action:      runApi,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Value:   "3000",
			Usage:   "Temporary port number to prevent conflict",
		},
	},
}

func runApi(ctx *cli.Context) error {
	return nil
}
