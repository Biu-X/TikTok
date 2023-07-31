package cmd

import "github.com/urfave/cli/v2"

func NewApp() *cli.App { //nolint:typecheck
	app := cli.NewApp() //nolint:typecheck
	app.EnableBashCompletion = true

	// 子命令集
	subCmdWithConfig := []*cli.Command{ //nolint:typecheck
		CmdWeb,
		CmdGen,
	}

	app.Commands = append(app.Commands, subCmdWithConfig...)
	return app
}
