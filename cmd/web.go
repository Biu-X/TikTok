package cmd

import (
	"biu-x.org/TikTok/modules/config"
	"biu-x.org/TikTok/routers"
	"fmt"
	"github.com/urfave/cli/v2"
)

// CmdWeb api 子命令
var CmdWeb = &cli.Command{ //nolint:typecheck
	Name:        "server",
	Usage:       "Start TikTok api server",
	Description: `Star TikTok api server`,
	Action:      runWeb,
	Flags: []cli.Flag{
		&cli.StringFlag{ //nolint:typecheck
			Name:    "port",
			Aliases: []string{"p"},
			Value:   "3000",
			Usage:   "Temporary port number to prevent conflict",
		},
	},
}

func runWeb(ctx *cli.Context) error { //nolint:typecheck
	config.InitConfig()
	fmt.Println(config.Get("redis"))
	routers.Init()
	return nil
}
