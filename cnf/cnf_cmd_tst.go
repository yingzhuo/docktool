/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
*	 ____   ___   ____ _  _______ ___   ___  _
*	|  _ \ / _ \ / ___| |/ /_   _/ _ \ / _ \| |
*	| | | | | | | |   | ' /  | || | | | | | | |
*	| |_| | |_| | |___| . \  | || |_| | |_| | |___
*	|____/ \___/ \____|_|\_\ |_| \___/ \___/|_____|
*
*	https://github.com/yingzhuo/docktool
* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package cnf

import (
	jstr "github.com/yingzhuo/jing/str"
)

var TestEnvNames jstr.StringSlice
var TestFilenames jstr.StringSlice
var TestDirnames jstr.StringSlice
var TestTcpAddrs jstr.StringSlice

func init() {
	TestEnvNames = jstr.StringSlice{}
	TestFilenames = jstr.StringSlice{}
	TestDirnames = jstr.StringSlice{}
	TestTcpAddrs = jstr.StringSlice{}
}
