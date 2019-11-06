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

	"github.com/yingzhuo/docktool/cnf"
)

// print
// ---------------------

func Print(a ...interface{}) {
	if cnf.GlobalQuietMode {
		return
	}
	fmt.Print(a...)
}

func Println(a ...interface{}) {
	if cnf.GlobalQuietMode {
		return
	}
	fmt.Println(a...)
}

func Printf(format string, a ...interface{}) {
	if cnf.GlobalQuietMode {
		return
	}
	fmt.Printf(format, a...)
}
