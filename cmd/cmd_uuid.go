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

func NewCommandUUID() *cli.Command {

	examples := `docktool uuid
docktool uuid -s
docktool uuid -s -n`

	return &cli.Command{
		Name:        "uuid",
		Usage:       "create random uuid",
		UsageText:   "[options]",
		Description: "create random uuid",
		Examples:    examples,
		SeeAlso:     "https://github.com/yingzhuo/docktool/tree/master/.github/uuid.md",
		Flags: []*cli.Flag{
			{
				Name:          "s, short",
				Usage:         "short version of uuid",
				IsBool:        true,
				DefValue:      "false",
				NoOptDefValue: "true",
				Value:         &cnf.UUIDShortVersion,
			}, {
				Name:          "n",
				Usage:         "do not print the trailing newline character",
				IsBool:        true,
				DefValue:      "false",
				NoOptDefValue: "true",
				Value:         &cnf.UUIDNoNewLine,
			},
		},
		Action: ActionUUID,
	}

}
