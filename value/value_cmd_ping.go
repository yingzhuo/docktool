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
	"errors"
	"strings"
)

type PingLogic string

func (e *PingLogic) String() string {
	return string(*e)
}

func (e *PingLogic) Set(value string) error {
	value = strings.ToUpper(strings.TrimSpace(value))

	switch value {
	case "ALL", "ANY", "NONE":
		*e = PingLogic(value)
		return nil
	default:
		return errors.New("invalid flag: logic\nmust be one of 'ALL', 'ANY', 'NONE'")
	}
}
