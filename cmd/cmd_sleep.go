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
	"github.com/yingzhuo/go-cli/v2"
)

func NewCommandSleep() *cli.Command {

	examples := `docktool sleep
docktool sleep 5s
docktool sleep 1m30s
docktool sleep '2019-11-08 12:00:00 +08:00'`

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
