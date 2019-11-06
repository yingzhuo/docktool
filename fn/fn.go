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
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"strings"
	"text/template"
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
	f["randAlphabetic"] = randAlphabetic
	f["randNumeric"] = randNumeric
	f["randAlphanumeric"] = randAlphanumeric
	f["randString"] = randString
	f["randBool"] = randBool

	// base64
	f["base64Enc"] = base64URLEncode
	f["base64Dec"] = base64URLDecode

	// uuid
	f["uuid"] = uuid36
	f["uuid36"] = uuid36
	f["uuid32"] = uuid32

	// crypto
	f["md5"] = _md5
	f["sha1"] = _sha1
	f["sha256"] = _sha256
	f["sha384"] = _sha384
	f["sha512"] = _sha512

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

func base64URLEncode(s string) string {
	input := []byte(s)
	return base64.URLEncoding.EncodeToString(input)
}

func base64URLDecode(s string) string {
	data, _ := base64.URLEncoding.DecodeString(s)
	return string(data)
}

func uuid36() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func uuid32() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func _md5(s string) string {
	h := md5.New()
	_, _ = io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func _sha1(s string) string {
	h := sha1.New()
	_, _ = io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func _sha256(s string) string {
	h := sha256.New()
	_, _ = io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func _sha384(s string) string {
	h := sha512.New384()
	_, _ = io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func _sha512(s string) string {
	h := sha512.New()
	_, _ = io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

var (
	alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	alphabetic   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numeric      = "0123456789"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func randAlphabetic(n int) string {
	return randString(n, alphabetic)
}

func randNumeric(n int) string {
	return randString(n, numeric)
}

func randAlphanumeric(n int) string {
	return randString(n, alphanumeric)
}

func randString(n int, charset string) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(charset) {
			b[i] = charset[idx]
			i++
		}
	}
	return string(b)
}

func randBool() bool {
	return rand.Intn(1000)%2 == 0
}
