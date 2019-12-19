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
	"os"

	"github.com/yingzhuo/docktool/cnf"
	"github.com/yingzhuo/go-cli/v2"
	jio "github.com/yingzhuo/jing/io"
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

func ActionSecretSHA1(c *cli.Context) {
	s := getString(c)
	Debugf("raw: %v\n", s)
	s = jstr.SHA1(s)

	if cnf.SecretNoNewLine {
		Printf("%s", s)
	} else {
		Printf("%s\n", s)
	}
}

func ActionSecretSHA256(c *cli.Context) {
	s := getString(c)
	Debugf("raw: %v\n", s)
	s = jstr.SHA256(s)

	if cnf.SecretNoNewLine {
		Printf("%s", s)
	} else {
		Printf("%s\n", s)
	}
}

func ActionSecretSHA512(c *cli.Context) {
	s := getString(c)
	Debugf("raw: %v\n", s)
	s = jstr.SHA512(s)

	if cnf.SecretNoNewLine {
		Printf("%s", s)
	} else {
		Printf("%s\n", s)
	}
}

func ActionSecretSHA384(c *cli.Context) {
	s := getString(c)
	Debugf("raw: %v\n", s)
	s = jstr.SHA384(s)

	if cnf.SecretNoNewLine {
		Printf("%s", s)
	} else {
		Printf("%s\n", s)
	}
}

func ActionSecretBcrypt(c *cli.Context) {

	if cnf.SecretBcryptChecking {
		// 验证

		if c.NArg() != 2 {
			os.Exit(1)
		}

		raw := c.Arg(0)
		hashed := c.Arg(1)

		if jstr.CheckBCrypt(raw, hashed) {
			Println("match")
		} else {
			Println("not match")
			os.Exit(1)
		}

	} else {
		// 生成
		s := getString(c)
		Debugf("raw: %v\n", s)

		if s == "" {
			return
		}

		s = jstr.BCrypt(s)

		if cnf.SecretNoNewLine {
			Printf("%s", s)
		} else {
			Printf("%s\n", s)
		}
	}
}

func getString(c *cli.Context) string {

	if cnf.SecretStdin {
		return jio.LoadStdin()
	}

	if c.NArg() == 0 {
		return ""
	}

	if c.NArg() >= 2 {
		panic("too many parameters for the command")
	}

	return c.Arg(0)
}
