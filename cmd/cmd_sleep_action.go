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

	"github.com/yingzhuo/go-cli/v2"
	jtime "github.com/yingzhuo/jing/time"
)

func ActionSleep(c *cli.Context) {

	switch c.NArg() {
	case 0:
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
	d, err := jtime.ParseDurationWildly(duration)

	if err != nil {
		panic(fmt.Sprintf("can't parse duration \"%s\"", duration))
	}

	if d <= 0 {
		sleepForever()
	} else {
		time.Sleep(d)
	}
}
