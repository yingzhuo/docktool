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
)

func NewCommandSecret() *cli.Command {

	examples := `docktool secret base64 "secret"
docktool secret base64 "secret" -n
docktool secret base64 --decode "c2VjcmV0"
docktool secret md5 "secret"
docktool secret md4 "secret"`

	return &cli.Command{
		Name:        "secret",
		Usage:       "encode/decode a string",
		Description: "encode/decode a string",
		Examples:    examples,
		SeeAlso:     "https://github.com/yingzhuo/docktool/tree/master/.github/secret.md", // TODO
		Commands: []*cli.Command{
			{
				Name: "base64",
				Flags: []*cli.Flag{
					{
						Name:          "n",
						Usage:         "do not print the trailing newline character",
						IsBool:        true,
						DefValue:      "false",
						NoOptDefValue: "true",
						Value:         &cnf.SecretNoNewLine,
					}, {
						Name:          "stdin",
						Usage:         "take string from stdin",
						IsBool:        true,
						DefValue:      "false",
						NoOptDefValue: "true",
						Value:         &cnf.SecretStdin,
					}, {
						Name:          "d, decode",
						Usage:         "decoding",
						IsBool:        true,
						DefValue:      "false",
						NoOptDefValue: "true",
						Value:         &cnf.SecretBase64Decode,
					},
				},
				Action: ActionSecretBase64,
			}, {
				Name: "md5",
				Flags: []*cli.Flag{
					{
						Name:          "n",
						Usage:         "do not print the trailing newline character",
						IsBool:        true,
						DefValue:      "false",
						NoOptDefValue: "true",
						Value:         &cnf.SecretNoNewLine,
					}, {
						Name:          "stdin",
						Usage:         "take string from stdin",
						IsBool:        true,
						DefValue:      "false",
						NoOptDefValue: "true",
						Value:         &cnf.SecretStdin,
					},
				},
				Action: ActionSecretMD5,
			}, {
				Name: "md4",
				Flags: []*cli.Flag{
					{
						Name:          "n",
						Usage:         "do not print the trailing newline character",
						IsBool:        true,
						DefValue:      "false",
						NoOptDefValue: "true",
						Value:         &cnf.SecretNoNewLine,
					}, {
						Name:          "stdin",
						Usage:         "take string from stdin",
						IsBool:        true,
						DefValue:      "false",
						NoOptDefValue: "true",
						Value:         &cnf.SecretStdin,
					},
				},
				Action: ActionSecretMD4,
			},
		},
	}

}
