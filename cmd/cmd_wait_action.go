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
	"os"
	"strings"
	"time"

	"github.com/yingzhuo/docktool/cnf"
	"github.com/yingzhuo/docktool/value"
	"github.com/yingzhuo/go-cli/v2"
	jtcp "github.com/yingzhuo/jing/tcp"
)

var ch chan string

func ActionWait(c *cli.Context) {

	if c.NArg() != 0 {
		panic("too many parameters for sub-command wait")
	}

	list := getList()

	if list.IsEmpty() {
		return
	}

	count := len(list)
	timeoutFlag := false
	ch = make(chan string, count)

	for id, addr := range list {
		go doWait(fmt.Sprintf("%d", id), addr, &timeoutFlag)
	}

	go doTimeout(cnf.WaitTimeout.Get(), &timeoutFlag)

	n := 0
	for {
		select {
		case <-ch:
			n++

			if cnf.WaitLogic == "ANY" {
				goto t1
			}

			if n == count {
				goto t1
			}
		}
	}

t1:
	close(ch)
}

func getList() value.WaitList {
	list := cnf.WaitList

	// 读取环境变量
	for _, v := range os.Environ() {
		i := strings.Index(v, "=")
		if i > -1 {
			name := v[:i]
			val := v[i+1:]
			if name != "" && strings.HasPrefix(name, cnf.WaitEnvPrefix) {
				list.Add(val)
			}
		}
	}

	ret := value.WaitList{}
	dict := make(map[string]bool)

	// 去除重复
	for _, it := range list {
		if !dict[it] {
			ret.Add(it)
			dict[it] = true
		}
	}

	return ret
}

func doWait(threadId string, addr string, timeoutFlag *bool) {
	defer func() {
		recover()
	}()

	for {
		if jtcp.IsReachable(addr) {
			Printf("ok      : %s\n", addr)
			ch <- threadId
			return
		} else {
			if *timeoutFlag {
				Printf("timeout : %s\n", addr)
				ch <- threadId
				return
			} else {
				nap()
			}
		}
	}
}

func nap() {
	time.Sleep(time.Millisecond * 200)
}

func doTimeout(timeout time.Duration, quitVar *bool) {
	if timeout <= 0 || *quitVar {
		return
	}

	select {
	case <-time.After(timeout):
		*quitVar = true
	}
}
