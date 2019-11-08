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
	"math/rand"
	"reflect"
	"strings"
	"text/template"

	jstr "github.com/yingzhuo/jing/str"
)

func TxtFuncMap() template.FuncMap {
	f := template.FuncMap{}

	// for debug
	f["type"] = func(i interface{}) string {
		return fmt.Sprintf("%T", i)
	}

	// common
	f["size"] = size
	f["join"] = join
	f["default"] = defaultValue

	// to xxx
	f["toBool"] = toBool
	f["toString"] = toString
	f["toInt64"] = toInt64
	f["toInt"] = toInt64
	f["toFloat"] = toFloat64
	f["toFloat64"] = toFloat64

	// math
	f["increase"] = addOne
	f["addOne"] = addOne
	f["decrease"] = subOne
	f["subOne"] = subOne
	f["add"] = add
	f["sub"] = sub
	f["mul"] = mul
	f["div"] = div
	f["mod"] = mod
	f["max"] = max
	f["min"] = min

	// strings
	f["cat"] = cat
	f["upper"] = strings.ToUpper
	f["lower"] = strings.ToLower
	f["title"] = strings.ToTitle
	f["trim"] = strings.TrimSpace
	f["repeat"] = repeat
	f["replace"] = replaceAll
	f["quote"] = quote
	f["singleQuote"] = singleQuote

	// randXxx
	f["randAlphabetic"] = jstr.RandAlphabetic
	f["randNumeric"] = jstr.RandNumeric
	f["randAlphanumeric"] = jstr.RandAlphanumeric
	f["randBool"] = func() bool {
		return rand.Intn(1000)%2 == 0
	}

	// base64
	f["base64Enc"] = jstr.Base64URLEncode
	f["base64Dec"] = jstr.Base64URLDecode

	// uuid
	f["uuid"] = jstr.NewUUID36
	f["uuid36"] = jstr.NewUUID36
	f["uuid32"] = jstr.NewUUID32

	// crypto
	f["md4"] = jstr.MD4
	f["md5"] = jstr.MD5
	f["sha1"] = jstr.SHA1
	f["sha256"] = jstr.SHA256
	f["sha384"] = jstr.SHA384
	f["sha512"] = jstr.SHA512

	return f
}

func size(value interface{}) int {
	if value == nil {
		return 0
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return len([]rune(v.String()))
	case reflect.Slice, reflect.Array, reflect.Map:
		return v.Len()
	default:
		return 1
	}
}

func defaultValue(defaultValue interface{}, value interface{}) interface{} {
	if value == nil {
		return defaultValue
	}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		if v.Len() == 0 {
			return defaultValue
		}
	case reflect.Bool:
		if !v.Bool() {
			return defaultValue
		}
	default:
		return value
	}
	return value
}

func join(separator string, value []interface{}) string {

	sb := &strings.Builder{}
	for i, v := range value {
		_, _ = fmt.Fprintf(sb, "%v", v)
		if i != len(value)-1 {
			_, _ = fmt.Fprintf(sb, separator)
		}
	}
	return sb.String()
}

func singleQuote(s string) string {
	return "'" + s + "'"
}

func quote(s string) string {
	return "\"" + s + "\""
}

func repeat(n int, s string) string {
	return strings.Repeat(s, n)
}

func replaceAll(old, new string, s string) string {
	return strings.ReplaceAll(s, old, new)
}

func cat(s string, ss ...string) string {
	sb := &strings.Builder{}
	sb.WriteString(s)
	for _, it := range ss {
		sb.WriteString(it)
	}
	return sb.String()
}
