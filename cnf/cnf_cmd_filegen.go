/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
*	 ____   ___   ____ _  _______ ___   ___  _
*	|  _ \ / _ \ / ___| |/ /_   _/ _ \ / _ \| |
*	| | | | | | | |   | ' /  | || | | | | | | |
*	| |_| | |_| | |___| . \  | || |_| | |_| | |___
*	|____/ \___/ \____|_|\_\ |_| \___/ \___/|_____|
*
*	https://github.com/yingzhuo/docktool
* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package cnf

import (
	"github.com/yingzhuo/docktool/value"
)

var FilegenTemplateFile string
var FilegenJsonFile string
var FilegenYamlFile string
var FilegenTomlFile string
var FilegenVars value.FilegenVars
var FilegenDelims value.FilegenDelims
var FilegenStdin bool
var FilegenOutputFile string
var FilegenOutputAppend bool
var FilegenOutputPerm int32

func init() {
	FilegenVars = make(map[string]interface{})
}
