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
	"github.com/yingzhuo/docktool/cnf"
	"github.com/yingzhuo/go-cli/v2"
)

func NewCommandWait() *cli.Command {

	examples := `docktool wait -w localhost:3306 -w localhost:6379
docktool wait -e WAIT_
docktool wait -w localhost:8080 -t "2019-11-02 14:34:40 +0800
docktool wait -w localhost:3306 -w localhost:6379 -w localhost:8080 --logic=ANY
WAIT_MYSQL=localhost:3306 WAIT_REDIS=localhost:6379 docktool wait`

	return &cli.Command{
		Name:        "wait",
		Usage:       "wait until tcp reachable or timeout",
		UsageText:   "[options ...]",
		Description: "wait until tcp reachable or timeout",
		Examples:    examples,
		SeeAlso:     "https://github.com/yingzhuo/docktool/tree/master/.github/wait.md",
		Flags: []*cli.Flag{
			{
				Name:        "e, environment-prefix",
				Usage:       "environment name prefix of waiting list",
				DefValue:    "WAIT_",
				Value:       &cnf.WaitEnvPrefix,
				Placeholder: "<prefix>",
			}, {
				Name:        "w, wait",
				Usage:       "addr to wait, can be passed multiple times",
				Value:       &cnf.WaitList,
				Placeholder: "<addr>",
			}, {
				Name:        "t, timeout",
				Usage:       "timeout",
				Value:       &cnf.WaitTimeout,
				Placeholder: "<time>",
			}, {
				Name:        "s, shell",
				Usage:       "run shell after waiting",
				Value:       &cnf.WaitShell,
				Placeholder: "<shell>",
			}, {
				Name:        "logic",
				Usage:       "logic of testing (ALL | ANY)",
				Placeholder: "<logic>",
				DefValue:    "ALL",
				Value:       &cnf.WaitLogic,
			},
		},
		Action: ActionWait,
	}
}
