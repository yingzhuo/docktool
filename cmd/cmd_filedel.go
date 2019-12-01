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

func NewCommandFiledel() *cli.Command {

	examples := `docktool filedel --dir=/my/dir -p='*.yml' -p='*yaml'
docktool filedel --dir=/my/dir -p='data[0-9].csv'
docktool filedel --dir=/my/dir -p='*.txt' -p='*.md' -p=LICENSE
docktool filedel --dir=/my/dir -p='sub-dir-*/'
docktool filedel --dir=/my/dir --delete-empty-dir`

	return &cli.Command{
		Name:        "filedel",
		Usage:       "delete files/dirs using wildcard",
		UsageText:   "[options]",
		Description: "delete files/dirs using wildcard",
		Examples:    examples,
		SeeAlso:     "https://github.com/yingzhuo/docktool/tree/master/.github/filedel.md",
		Flags: []*cli.Flag{
			{
				Name:          "dir",
				Usage:         "start dir",
				Placeholder:   "<dir>",
				DefValue:      ".",
				NoOptDefValue: ".",
				Value:         &cnf.FiledelDir,
			},
			{
				Name:        "p, pattern",
				Usage:       "wildcard pattern, can be passed multiple times",
				Placeholder: "<wildcard>",
				Value:       &cnf.FiledelPatterns,
			}, {
				Name:          "delete-empty-dir",
				Usage:         "delete empty dirs at last",
				DefValue:      "false",
				NoOptDefValue: "true",
				Value:         &cnf.FiledelDelEmptyDir,
			},
		},
		Action: ActionFiledel,
	}
}
