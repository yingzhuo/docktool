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
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/yingzhuo/docktool/cnf"
	"github.com/yingzhuo/go-cli/v2"
	"github.com/yingzhuo/jing/tcp"
)

func ActionPing(c *cli.Context) {

	logrus.Debugf("command: \"%v\"", c.Name())
	logrus.Debugf("pwd: \"%v\"", cnf.GlobalPWD)

	eff := make([]string, 0)
	dict := make(map[string]bool)

	// 去除重复参数
	for _, it := range c.Args() {
		it = strings.TrimSpace(it)

		if !dict[it] {
			eff = append(eff, it)
			dict[it] = true
		}
	}

	if len(eff) == 0 {
		panic("require at least one parameter")
	}

	if cnf.GlobalDebugMode {
		for i, v := range eff {
			logrus.Debugf("(%v) ping: %v", i+1, v)
		}
	}

	switch cnf.PingLogic {
	case "ALL":

		for _, addr := range eff {
			if tcp.IsNotReachable(addr) {
				os.Exit(1)
			}
		}

	case "ANY":

		anyOK := false

		for _, addr := range eff {
			if tcp.IsReachable(addr) {
				anyOK = true
				break
			}
		}

		if !anyOK {
			os.Exit(1)
		}

	case "NONE":

		noneOK := true

		for _, addr := range eff {
			if tcp.IsReachable(addr) {
				noneOK = false
				break
			}
		}

		if !noneOK {
			os.Exit(1)
		}
	}
}
