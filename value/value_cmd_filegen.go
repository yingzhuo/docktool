/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
*	 ____   ___   ____ _  _______ ___   ___  _
*	|  _ \ / _ \ / ___| |/ /_   _/ _ \ / _ \| |
*	| | | | | | | |   | ' /  | || | | | | | | |
*	| |_| | |_| | |___| . \  | || |_| | |_| | |___
*	|____/ \___/ \____|_|\_\ |_| \___/ \___/|_____|
*
*	https://github.com/yingzhuo/docktool
* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package value

import (
	"fmt"
	"strings"
)

type FilegenVars map[string]interface{}

func (e *FilegenVars) String() string {
	return fmt.Sprintf("%v", *e)
}

func (e *FilegenVars) Set(value string) error {
	if parts := strings.Split(value, "="); len(parts) == 2 {
		(*e)[parts[0]] = parts[1]
	}
	return nil
}

// -----------------

type FilegenDelims struct {
	Left  string
	Right string
}

func (e *FilegenDelims) String() string {
	return e.Left + ":" + e.Right
}

func (e *FilegenDelims) Set(value string) error {
	if parts := strings.Split(value, ":"); len(parts) == 2 {
		e.Left, e.Right = "{{", "}}"
	} else {
		e.Left, e.Right = parts[0], parts[1]
	}
	return nil
}
