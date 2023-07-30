package cmd

import (
	"biu-x.org/TikTok/modules/config"
	"biu-x.org/TikTok/routers"
	"fmt"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

// CmdWeb api 子命令
var CmdWeb = &cli.Command{
	Name:        "server",
	Usage:       "Start TikTok api server",
	Description: `Star TikTok api server`,
	Action:      runWeb,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Value:   "3000",
			Usage:   "Temporary port number to prevent conflict",
		},
	},
}

func runWeb(ctx *cli.Context) error {
	config.InitConfig()
	fmt.Println(viper.Get("server.port"))
	routers.Init()
	return nil
}
