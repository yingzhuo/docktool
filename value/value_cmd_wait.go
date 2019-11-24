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
	"fmt"
	"strings"
	"time"

	jtime "github.com/yingzhuo/jing/time"
)

type WaitList []string

func (e *WaitList) String() string {
	return strings.Join(*e, ",")
}

func (e *WaitList) Set(value string) error {
	e.Add(value)
	return nil
}

func (e *WaitList) Add(value string) {
	if !strings.Contains(value, ",") {
		value = strings.TrimSpace(value)
		*e = append(*e, value)
	} else {
		for _, v := range strings.Split(value, ",") {
			v = strings.TrimSpace(v)
			if v != "" {
				*e = append(*e, strings.TrimSpace(v))
			}
		}
	}
}

func (e *WaitList) IsEmpty() bool {
	return len(*e) == 0
}

// -----------------

type WaitTimeout time.Duration

func (e *WaitTimeout) String() string {
	return fmt.Sprintf("%v", *e)
}

func (e *WaitTimeout) Set(value string) error {
	value = strings.TrimSpace(value)

	if d, err := time.ParseDuration(value); err == nil {
		*e = WaitTimeout(d)
		return nil
	}

	futureTime, err := jtime.ParseTimeWildly(value)
	if err != nil {
		return err
	}

	*e = WaitTimeout(futureTime.Sub(time.Now().Local()))
	return nil
}

func (e *WaitTimeout) Get() time.Duration {
	return time.Duration(*e)
}

// -----------------

type WaitLogic string

func (e *WaitLogic) String() string {
	return string(*e)
}

func (e *WaitLogic) Set(value string) error {
	value = strings.ToUpper(strings.TrimSpace(value))

	switch value {
	case "ALL", "ANY":
		*e = WaitLogic(value)
		return nil
	default:
		return errors.New("invalid flag: logic\nmust be one of 'ALL', 'ANY'")
	}
}
