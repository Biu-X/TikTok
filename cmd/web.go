package cmd

import (
	"biu-x.org/TikTok/module/config"
	"biu-x.org/TikTok/module/db"
	"biu-x.org/TikTok/module/log"
	"biu-x.org/TikTok/module/middleware/cache"
	"biu-x.org/TikTok/module/oss"
	"biu-x.org/TikTok/router"
	"github.com/urfave/cli/v2"
)

// CmdWeb api 子命令
var CmdWeb = &cli.Command{ //nolint:typecheck
	Name:        "server",
	Usage:       "Start TikTok api server",
	Description: `Star TikTok api server`,
	Action:      runWeb, //nolint:typecheck
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
	config.Init()
	log.Init()
	db.Init()
	oss.Init()
	cache.Init()
	router.Init()
	return nil
}
