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
	jbin "github.com/yingzhuo/jing/bin"
	jos "github.com/yingzhuo/jing/os"
)

// 安静模式
var GlobalQuietMode bool

// Debug模式
var GlobalDebugMode bool

// 当前目录
var GlobalPWD string

// app安装目录
var GlobalBinDir string

func init() {
	GlobalQuietMode = false
	GlobalDebugMode = false
	GlobalPWD = jos.GetPWD()
	GlobalBinDir = jbin.GetBinaryDir()
}
