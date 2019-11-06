/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
*	 ____   ___   ____ _  _______ ___   ___  _
*	|  _ \ / _ \ / ___| |/ /_   _/ _ \ / _ \| |
*	| | | | | | | |   | ' /  | || | | | | | | |
*	| |_| | |_| | |___| . \  | || |_| | |_| | |___
*	|____/ \___/ \____|_|\_\ |_| \___/ \___/|_____|
*
*	https://github.com/yingzhuo/docktool
* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package fn

import (
	"fmt"
	"strconv"
	"strings"
)

func toString(n interface{}) string {
	switch s := n.(type) {
	case string:
		return s
	case fmt.Stringer:
		return s.String()
	case fmt.GoStringer:
		return s.GoString()
	case []rune:
		return string(s)
	default:
		return fmt.Sprintf("%v", s)
	}
}

func toFloat64(input interface{}) float64 {
	switch n := input.(type) {
	case uint:
		return float64(n)
	case uint8:
		return float64(n)
	case uint16:
		return float64(n)
	case uint32:
		return float64(n)
	case uint64:
		return float64(n)
	case int:
		return float64(n)
	case int8:
		return float64(n)
	case int16:
		return float64(n)
	case int32:
		return float64(n)
	case int64:
		return float64(n)
	case float32:
		return float64(n)
	case float64:
		return n
	case string:
		if r, err := strconv.ParseFloat(n, 64); err != nil {
			panic("operation not supported")
		} else {
			return r
		}
	default:
		panic("operation not supported")
	}
}

func toInt64(input interface{}) int64 {
	switch n := input.(type) {
	case uint:
		return int64(n)
	case uint8:
		return int64(n)
	case uint16:
		return int64(n)
	case uint32:
		return int64(n)
	case uint64:
		return int64(n)
	case int:
		return int64(n)
	case int8:
		return int64(n)
	case int16:
		return int64(n)
	case int32:
		return int64(n)
	case int64:
		return n
	case float32:
		return int64(n)
	case float64:
		return int64(n)
	case string:
		if r, err := strconv.ParseInt(n, 10, 64); err != nil {
			panic("operation not supported")
		} else {
			return r
		}
	default:
		panic("operation not supported")
	}
}

func toBool(s string) bool {
	s = strings.ToLower(s)

	switch s {
	case "t", "true", "1", "on", "yes":
		return true
	case "f", "false", "0", "off", "no":
		return false
	default:
		panic("can't convert to bool ('" + s + "')")
	}
}
