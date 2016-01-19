package natsort

const (
	Number = iota + 1
	Spliter
	Alpha
	Other
)

var VersionSpliter byte = '.'

func Kind(a byte) int {
	switch {
	case isNumber(a):
		return Number
	case isAlpha(a):
		return Alpha
	case isSpliter(a):
		return Spliter
	}

	return Other
}

func isNumber(a byte) bool {
	return '0' <= a && a <= '9'
}

func isAlpha(a byte) bool {
	return ('a' <= a && a <= 'z') || ('A' <= a && a <= 'Z')
}

func isSpliter(a byte) bool {
	return a == VersionSpliter
}
