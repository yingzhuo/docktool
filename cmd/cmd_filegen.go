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

func NewCommandFilegen() *cli.Command {

	examples := `docktool filegen -t ./filegen.tmpl -v name=bill
docktool filegen -t ./filegen.tmpl --json in.json --yaml in.yaml --toml in.toml
docktool filegen -t ./filegen.tmpl -o ./gen.txt
echo "{{ .foo }}" | docktool filegen --stdin -v foo=foo
echo "{{ .ENV.foo }}" | foo=foo docktool filegen --stdin
echo "{{ .ENV.HOME }}" | docktool filegen --stdin`

	return &cli.Command{
		Name:        "filegen",
		Usage:       "generate file using template",
		UsageText:   "[options]",
		Description: "generate file using template from environment, arguments, json/yaml/toml config files",
		Examples:    examples,
		SeeAlso:     "https://github.com/yingzhuo/docktool/tree/master/.github/filegen.md",
		Flags: []*cli.Flag{
			{
				Name:        "t, template",
				Usage:       "template file",
				Placeholder: "<file>",
				Value:       &cnf.FilegenTemplateFile,
			}, {
				Name:        "json",
				Usage:       "json file for input",
				Placeholder: "<file>",
				Value:       &cnf.FilegenJsonFile,
			}, {
				Name:        "yaml",
				Usage:       "yaml file for input",
				Placeholder: "<file>",
				Value:       &cnf.FilegenYamlFile,
			}, {
				Name:        "toml",
				Usage:       "toml file for input",
				Placeholder: "<file>",
				Value:       &cnf.FilegenTomlFile,
			}, {
				Name:        "v, var",
				Usage:       "vars for input, can be passed multiple times",
				Placeholder: "<key>=<value>",
				Value:       &cnf.FilegenVars,
			}, {
				Name:          "delims",
				Usage:         "template tag delimiters",
				Placeholder:   "<left>:<right>",
				DefValue:      "{{:}}",
				NoOptDefValue: "{{:}}",
				Value:         &cnf.FilegenDelims,
			}, {
				Name:   "stdin",
				Usage:  "get template content from console",
				IsBool: true,
				Value:  &cnf.FilegenStdin,
			}, {
				Name:        "o, output",
				Usage:       "file for output",
				Placeholder: "<file>",
				Value:       &cnf.FilegenOutputFile,
			}, {
				Name:          "a, append",
				Usage:         "append output to file",
				NoOptDefValue: "true",
				IsBool:        true,
				Value:         &cnf.FilegenOutputAppend,
			}, {
				Name:        "perm",
				Usage:       "perm of output file",
				DefValue:    "0660",
				Placeholder: "<value>",
				Value:       &cnf.FilegenOutputPerm,
			},
		},
		Action: ActionFilegen,
	}
}
