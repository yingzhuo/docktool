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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"github.com/yingzhuo/docktool/cnf"
	"github.com/yingzhuo/docktool/fn"
	"github.com/yingzhuo/go-cli/v2"
	jfile "github.com/yingzhuo/jing/file"
	jio "github.com/yingzhuo/jing/io"
	"gopkg.in/yaml.v2"
)

func ActionFilegen(c *cli.Context) {

	initConfig()

	logrus.Debugf("command: \"%v\"", c.Name())
	logrus.Debugf("bin dir: \"%v\"", cnf.GlobalBinDir)
	logrus.Debugf("pwd: \"%v\"", cnf.GlobalPWD)
	logrus.Debugf("template file: \"%v\"", cnf.FilegenTemplateFile)
	logrus.Debugf("json file: \"%v\"", cnf.FilegenJsonFile)
	logrus.Debugf("yaml file: \"%v\"", cnf.FilegenYamlFile)
	logrus.Debugf("toml file: \"%v\"", cnf.FilegenTomlFile)
	logrus.Debugf("output file: \"%v\"", cnf.FilegenOutputFile)

	templateContent := loadTemplateContent()
	rootModel := getRootModel()

	engine := getTemplate(templateContent)
	stringBuilder := &strings.Builder{}

	if err := engine.Execute(stringBuilder, rootModel); err != nil {
		panic(err)
	}

	writeOutput(stringBuilder.String())
}

func initConfig() {
	cnf.FilegenTemplateFile = jfile.BetterFilename(cnf.FilegenTemplateFile)
	cnf.FilegenYamlFile = jfile.BetterFilename(cnf.FilegenYamlFile)
	cnf.FilegenJsonFile = jfile.BetterFilename(cnf.FilegenJsonFile)
	cnf.FilegenTomlFile = jfile.BetterFilename(cnf.FilegenTomlFile)
	cnf.FilegenOutputFile = jfile.BetterFilename(cnf.FilegenOutputFile)

	backupFile(&cnf.FilegenTemplateFile, "./filegen.tmpl", "./docktool/filegen.tmpl")
	backupFile(&cnf.FilegenJsonFile, "./filegen.in.json", "./docktool/filegen.in.json")
	backupFile(&cnf.FilegenYamlFile, "./filegen.in.yaml", "./filegen.in.yml", "./docktool/filegen.in.yaml", "./docktool/filegen.in.yml")
	backupFile(&cnf.FilegenTomlFile, "./filegen.in.toml", "./docktool/filegen.in.toml")
}

func backupFile(filename *string, backupFilenames ...string) {
	if filename == nil || *filename != "" {
		return
	}

	if len(backupFilenames) == 0 {
		return
	}

	for _, it := range backupFilenames {
		if jfile.IsFileExists(it) {
			*filename = it
			return
		}
	}
}

func loadTemplateContent() string {

	if cnf.FilegenStdin {
		return jio.LoadStdin()
	}

	if cnf.FilegenTemplateFile == "" {
		panic("template file is not specified")
	}

	if !jfile.IsFileExists(cnf.FilegenTemplateFile) {
		panic(fmt.Sprintf("\"%s\" not exists", cnf.FilegenTemplateFile))
	}

	return jio.LoadFileString(cnf.FilegenTemplateFile)
}

func getRootModel() map[string]interface{} {
	root := make(map[string]interface{})
	root["YAML"] = getYamlModel()
	root["JSON"] = getJsonModel()
	root["TOML"] = getTomlModel()
	root["ENV"] = getEnvModel()

	// vars
	for k, v := range getVarsModel() {
		root[k] = v
	}

	return root
}

func getVarsModel() map[string]interface{} {

	defer func() {
		if cnf.GlobalDebugMode {
			logrus.Debugf("count of vars: %v", len(cnf.FilegenVars))
		}
	}()

	return cnf.FilegenVars
}

func getJsonModel() map[string]interface{} {
	dict := make(map[string]interface{})

	defer func() {
		if cnf.GlobalDebugMode {
			logrus.Debugf("count of json model: %v", len(dict))
		}
	}()

	// 指定了文件但是又不存在
	if cnf.FilegenJsonFile != "" && !jfile.IsFileExists(cnf.FilegenJsonFile) {
		panic(fmt.Sprintf("\"%s\" not exists", cnf.FilegenJsonFile))
	}

	if cnf.FilegenJsonFile != "" {
		data, _ := ioutil.ReadFile(cnf.FilegenJsonFile)
		if err := json.Unmarshal(data, &dict); err != nil {
			panic("invalid json format")
		}
	}

	return dict
}

func getYamlModel() map[string]interface{} {
	dict := make(map[string]interface{})

	defer func() {
		if cnf.GlobalDebugMode {
			logrus.Debugf("count of yaml model: %v", len(dict))
		}
	}()

	// 指定了文件但是又不存在
	if cnf.FilegenYamlFile != "" && !jfile.IsFileExists(cnf.FilegenYamlFile) {
		panic(fmt.Sprintf("\"%s\" not exists", cnf.FilegenYamlFile))
	}

	if cnf.FilegenYamlFile != "" {
		data, _ := ioutil.ReadFile(cnf.FilegenYamlFile)
		if err := yaml.Unmarshal(data, &dict); err != nil {
			panic("invalid yaml format")
		}
	}

	return dict
}

func getTomlModel() map[string]interface{} {
	dict := make(map[string]interface{})

	defer func() {
		if cnf.GlobalDebugMode {
			logrus.Debugf("count of toml model: %v", len(dict))
		}
	}()

	// 指定了文件但是又不存在
	if cnf.FilegenTomlFile != "" && !jfile.IsFileExists(cnf.FilegenTomlFile) {
		panic(fmt.Sprintf("\"%s\" not exists", cnf.FilegenTomlFile))
	}

	if cnf.FilegenTomlFile != "" {
		if _, err := toml.DecodeFile(cnf.FilegenTomlFile, &dict); err != nil {
			panic("invalid toml format")
		}
	}

	return dict
}

func getEnvModel() map[string]interface{} {
	dict := make(map[string]interface{})

	defer func() {
		if cnf.GlobalDebugMode {
			logrus.Debugf("count of env model: %v", len(dict))
		}
	}()

	for _, v := range os.Environ() {
		kv := strings.Split(v, "=")
		dict[kv[0]] = kv[1]
	}
	return dict
}

func getTemplate(templateContent string) *template.Template {
	var t *template.Template

	t = template.New("docktool-filegen").
		Delims(cnf.FilegenDelims.Left, cnf.FilegenDelims.Right).
		Funcs(fn.TxtFuncMap())

	return template.Must(t.Parse(templateContent))
}

func writeOutput(content string) {
	if cnf.FilegenOutputFile == "" {
		Print(content)
		return
	}

	fg := os.O_CREATE | os.O_WRONLY

	if cnf.FilegenOutputAppend {
		fg |= os.O_APPEND // 追加写
	} else {
		fg |= os.O_TRUNC // 覆盖写
	}

	filename := cnf.FilegenOutputFile
	dirname := filepath.Dir(filename)

	if _, err := os.Stat(dirname); err != nil {
		err = os.MkdirAll(dirname, 0777)
		if err != nil {
			panic(fmt.Sprintf("Can not create dir \"%s\"", dirname))
		}
	}

	f, err := os.OpenFile(filename, fg, os.FileMode(cnf.FilegenOutputPerm))
	if err != nil {
		panic(fmt.Sprintf("Can not open file \"%s\"", filename))
	}

	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		panic(fmt.Sprintf("Can not write \"%s\"", filename))
	}
}
