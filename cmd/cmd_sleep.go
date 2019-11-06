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
)

func NewCommandSleep() *cli.Command {

	examples := `docktool sleep
docktool sleep 5s
docktool sleep 1m30s`

	return &cli.Command{
		Name:        "sleep",
		Usage:       "make current thread sleep",
		UsageText:   "[argument]",
		Description: "make current thread sleep",
		Examples:    examples,
		SeeAlso:     "https://github.com/yingzhuo/docktool/tree/master/.github/sleep.md",
		Action:      ActionSleep,
	}
}
