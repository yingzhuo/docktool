/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
*	 ____   ___   ____ _  _______ ___   ___  _
*	|  _ \ / _ \ / ___| |/ /_   _/ _ \ / _ \| |
*	| | | | | | | |   | ' /  | || | | | | | | |
*	| |_| | |_| | |___| . \  | || |_| | |_| | |___
*	|____/ \___/ \____|_|\_\ |_| \___/ \___/|_____|
*
*	https://github.com/yingzhuo/docktool
* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package main

import (
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/yingzhuo/docktool/cmd"
	"github.com/yingzhuo/docktool/cnf"
	"github.com/yingzhuo/go-cli/v2"
)

// 构建信息
var (
	BuildVersion   string
	BuildGitBranch string
	BuildGitRev    string
	BuildGitCommit string
	BuildDate      string
)

func main() {

	// 基本信息
	app := cli.NewApp()
	app.Name = "docktool"
	app.Usage = "tools for docker container"
	app.Version = BuildVersion
	app.Authors = `yingzhuo <yingzhor@gmail.com>`

	app.BuildInfo = &cli.BuildInfo{
		GitBranch:   BuildGitBranch,
		GitCommit:   BuildGitCommit,
		GitRevCount: BuildGitRev,
		Timestamp:   BuildDate,
	}

	app.SeeAlso = `https://github.com/yingzhuo/docktool`

	// 全局Flag
	app.Flags = []*cli.Flag{
		{
			Name:          "q, quiet",
			Usage:         "quiet mode",
			DefValue:      "false",
			NoOptDefValue: "true",
			IsBool:        true,
			Value:         &cnf.GlobalQuietMode,
		},
	}

	// 处理所有错误
	app.OnActionPanic = func(context *cli.Context, e error) {
		msg := e.Error()

		if msg == "" {
			os.Exit(1)
		}

		if !cnf.GlobalQuietMode {
			if !strings.HasSuffix(msg, "\n") {
				msg += "\n"
			}
			os.Stderr.WriteString("error: " + msg)
		}
		os.Exit(1)
	}

	// 子命令
	app.Commands = []*cli.Command{
		cmd.NewCommandFilegen(),
		cmd.NewCommandWait(),
		cmd.NewCommandSleep(),
	}

	app.Run(os.Args)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
