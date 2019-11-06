/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
*	 ____   ___   ____ _  _______ ___   ___  _
*	|  _ \ / _ \ / ___| |/ /_   _/ _ \ / _ \| |
*	| | | | | | | |   | ' /  | || | | | | | | |
*	| |_| | |_| | |___| . \  | || |_| | |_| | |___
*	|____/ \___/ \____|_|\_\ |_| \___/ \___/|_____|
*
*	https://github.com/yingzhuo/docktool
* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package cmd

import (
	"github.com/subchen/go-cli"
	"github.com/yingzhuo/docktool/cnf"
)

func NewCommandPing() *cli.Command {

	examples := `docktool ping localhost:8080
docktool ping localhost:3306 localhost:6379
docktool ping --logic=ALL localhost:3306 localhost:6379
docktool ping --logic=ANY localhost:3306 localhost:6379`

	return &cli.Command{
		Name:        "ping",
		Usage:       "test if tcp is reachable",
		UsageText:   "[options ...] arg1 [arg2] [argN]",
		Description: "test if tcp is reachable",
		Examples:    examples,
		SeeAlso:     "https://github.com/yingzhuo/docktool/tree/master/.github/ping.md",
		Action:      ActionPing,
		Flags: []*cli.Flag{
			{
				Name:        "logic",
				Usage:       "logic of testing (ALL | ANY | NONE)",
				Placeholder: "<logic>",
				DefValue:    "ALL",
				Value:       &cnf.PingLogic,
			},
		},
	}
}
