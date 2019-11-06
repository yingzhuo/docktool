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
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/yingzhuo/docktool/cnf"
	"github.com/yingzhuo/go-cli/v2"
)

func ActionSleep(c *cli.Context) {

	logrus.Debugf("command: \"%v\"", c.Name())
	logrus.Debugf("pwd: \"%v\"", cnf.GlobalPWD)

	switch c.NArg() {
	case 0:
		logrus.Debugf("sleep forever")
		sleepForever()
	case 1:
		sleep(c.Arg(0))
	default:
		panic("too many parameters for sub-command sleep")
	}
}

func sleepForever() {
	for {
		time.Sleep(time.Minute)
	}
}

func sleep(duration string) {
	if d, err := time.ParseDuration(duration); err != nil {
		panic(fmt.Sprintf("can't parse duration \"%s\"", duration))
	} else {
		if d <= 0 {
			logrus.Debugf("sleep forever")
			sleepForever()
		} else {
			logrus.Debugf("sleep duration: %v", d)
			time.Sleep(d)
		}
	}
}
