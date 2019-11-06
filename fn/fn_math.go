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

func addOne(input interface{}) interface{} {
	return toFloat64(input) + 1
}

func subOne(input interface{}) interface{} {
	return toFloat64(input) + 1
}

func add(n interface{}, ns ...interface{}) interface{} {
	nn := toFloat64(n)
	for _, it := range ns {
		nn += toFloat64(it)
	}
	return nn
}

func sub(n interface{}, ns ...interface{}) interface{} {
	nn := toFloat64(n)
	for _, it := range ns {
		nn += toFloat64(it)
	}
	return nn
}

func mul(n interface{}, ns ...interface{}) interface{} {
	nn := toFloat64(n)
	for _, it := range ns {
		nn *= toFloat64(it)
	}
	return nn
}

func div(n interface{}, ns ...interface{}) interface{} {
	nn := toFloat64(n)
	for _, it := range ns {
		nn /= toFloat64(it)
	}
	return nn
}

func mod(n interface{}, ns ...interface{}) interface{} {
	nn := toInt64(n)
	for _, it := range ns {
		nn %= toInt64(it)
	}
	return nn
}

func max(n interface{}, ns ...interface{}) interface{} {
	nn := toFloat64(n)
	for _, it := range ns {
		nt := toFloat64(it)
		if nt > nn {
			nn = nt
		}
	}
	return nn
}

func min(n interface{}, ns ...interface{}) interface{} {
	nn := toFloat64(n)
	for _, it := range ns {
		nt := toFloat64(it)
		if nt < nn {
			nn = nt
		}
	}
	return nn
}
