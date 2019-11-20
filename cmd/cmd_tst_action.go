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
	"os"

	"github.com/yingzhuo/docktool/cnf"
	"github.com/yingzhuo/go-cli/v2"
	jfile "github.com/yingzhuo/jing/file"
	jtcp "github.com/yingzhuo/jing/tcp"
)

func ActionTest(_ *cli.Context) {

	envNames := cnf.TestEnvNames.Unique()

	for _, it := range *envNames {
		if _, found := os.LookupEnv(it); !found {
			Printf("env name \"%s\" NOT found\n", it)
			os.Exit(-1)
		}
	}

	filenams := cnf.TestFilenames.Unique()
	for _, it := range *filenams {
		if jfile.IsFileNotExists(it) {
			Printf("file \"%s\" NOT exists\n", it)
			os.Exit(-1)
		}
	}

	dirnames := cnf.TestDirnames.Unique()
	for _, it := range *dirnames {
		if jfile.IsDirNotExists(it) {
			Printf("dir \"%s\" NOT exists\n", it)
			os.Exit(-1)
		}
	}

	addrs := cnf.TestTcpAddrs.Unique()
	for _, it := range *addrs {
		if jtcp.IsNotReachable(it) {
			Printf("addr \"%s\" is NOT reachable\n", it)
			os.Exit(-1)
		}
	}
}
