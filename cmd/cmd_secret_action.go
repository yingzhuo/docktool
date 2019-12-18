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
	"github.com/yingzhuo/docktool/cnf"
	"github.com/yingzhuo/go-cli/v2"
	"github.com/yingzhuo/jing/io"
	jstr "github.com/yingzhuo/jing/str"
)

func ActionSecretBase64(c *cli.Context) {

	s := getString(c)
	Debugf("raw: %v\n", s)

	if !cnf.SecretBase64Decode {
		s = jstr.Base64URLEncode(s)
	} else {
		s = jstr.Base64URLDecode(s)
	}

	if cnf.SecretNoNewLine {
		Printf("%s", s)
	} else {
		Printf("%s\n", s)
	}
}

func ActionSecretMD5(c *cli.Context) {

	s := getString(c)
	Debugf("raw: %v\n", s)
	s = jstr.MD5(s)

	if cnf.SecretNoNewLine {
		Printf("%s", s)
	} else {
		Printf("%s\n", s)
	}
}

func ActionSecretMD4(c *cli.Context) {

	s := getString(c)
	Debugf("raw: %v\n", s)
	s = jstr.MD4(s)

	if cnf.SecretNoNewLine {
		Printf("%s", s)
	} else {
		Printf("%s\n", s)
	}
}

func getString(c *cli.Context) string {

	if cnf.SecretStdin {
		return io.LoadStdin()
	}

	if c.NArg() == 0 {
		return ""
	}

	if c.NArg() >= 2 {
		panic("too many parameters for the command")
	}

	return c.Arg(0)
}
