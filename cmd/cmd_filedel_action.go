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
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/yingzhuo/docktool/cnf"
	"github.com/yingzhuo/go-cli/v2"
	jfile "github.com/yingzhuo/jing/file"
)

func ActionFiledel(_ *cli.Context) {

	startDir := jfile.BetterFilename(cnf.FiledelDir)
	if startDir == "" {
		startDir = "."
	}

	if jfile.IsDirNotExists(startDir) {
		panic(fmt.Errorf("\"%v\" not exists", startDir))
	}

	patterns := cnf.FiledelPatterns.
		Unique()

	fs := make([]string, 0)

	walker := jfile.FileWalker{
		Start: startDir,
		OnDir: func(dirname string) {
			basename := filepath.Base(dirname)
			for _, pattern := range *patterns {

				if !strings.HasSuffix(pattern, "/") {
					continue
				} else {
					pattern = pattern[0 : len(pattern)-1]
				}

				if matched, _ := filepath.Match(pattern, basename); matched {
					fs = append(fs, dirname)
				}
			}
		},
		OnFile: func(filename string) {
			basename := filepath.Base(filename)
			for _, pattern := range *patterns {
				if matched, _ := filepath.Match(pattern, basename); matched {
					fs = append(fs, filename)
				}
			}
		},
		OnError: func(_ string, _ error) {
			// NOP
		},
	}
	walker.Run()

	for _, f := range fs {
		Printf("del: \"%v\"\n", f)
		_ = jfile.RemoveFileOrDir(f)
	}

	// 最后删除所有空目录
	if cnf.FiledelDelEmptyDir {

		w := jfile.FileWalker{
			Start: startDir,
			OnDir: func(dirname string) {
				if jfile.IsDirExists(dirname ) && isEmptyDir(dirname) {
					_ = jfile.RemoveFileOrDir(dirname)
				}
			},
		}
		w.Run()
	}

}

func isEmptyDir(dirname string) bool {
	f, err := os.Open(dirname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true
	}
	return false
}
