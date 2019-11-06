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
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

// 安静模式
var GlobalQuietMode bool

// Debug模式
var GlobalDebugMode bool

// 当前目录
var GlobalPWD string

// 二进制文件安装目录
var GlobalBinaryDir string

// 初始化全军设置
func InitGlobalConfig() {

	if GlobalDebugMode {
		// 修改日志级别
		logrus.SetLevel(logrus.DebugLevel)
	}

}

func init() {
	GlobalQuietMode = false
	GlobalDebugMode = false
	if GlobalPWD, _ = os.Getwd(); GlobalPWD == "" {
		GlobalPWD = "."
	}

	GlobalBinaryDir = getBinaryDir()
}

func getBinaryDir() string {
	app := os.Args[0]
	if !strings.ContainsAny(app, `/\`) {
		app, _ = exec.LookPath(app)
	}
	dir, _ := filepath.Abs(filepath.Dir(app))
	if !strings.HasSuffix(dir, string(os.PathSeparator)) {
		dir += string(os.PathSeparator)
	}
	return dir
}
