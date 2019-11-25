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

func NewCommandTest() *cli.Command {

	examples := `docktool test --env="JAVA_HOME"
docktool test --file="/tmp/test.sh"
docktool test --dir="/tmp"
docktool test --tcp="localhost:8080"`

	return &cli.Command{
		Name:        "test",
		Usage:       "test env/file/tcp",
		UsageText:   "[options]",
		Description: "test env/file/tcp",
		Examples:    examples,
		SeeAlso:     "https://github.com/yingzhuo/docktool/tree/master/.github/test.md",
		Flags: []*cli.Flag{
			{
				Name:        "e, env",
				Usage:       "env name for test, can be passed multiple times",
				Placeholder: "<env-name>",
				Value:       &cnf.TestEnvNames,
			}, {
				Name:        "f, file",
				Usage:       "file name for test, can be passed multiple times",
				Placeholder: "<file-name>",
				Value:       &cnf.TestFilenames,
			}, {
				Name:        "d, dir",
				Usage:       "dir name for test, can be passed multiple times",
				Placeholder: "<dir-name>",
				Value:       &cnf.TestDirnames,
			}, {
				Name:        "tcp",
				Usage:       "tcp addr for test, can be passed multiple times",
				Placeholder: "<tcp-addr>",
				Value:       &cnf.TestTcpAddrs,
			},
		},
		Commands: nil,
		Action:   ActionTest,
	}
}
