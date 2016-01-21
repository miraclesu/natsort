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

func getPart(version string, i int, ipart []byte) (part []byte, kind, current int, ok bool) {
	a := version[i]
	kind, part = Kind(a), ipart
	part = append(part, a)

	i++
	ok = (kind == Spliter)
	if kind != Number {
		current = i
		return
	}

	for count := len(version); i < count; i++ {
		a = version[i]
		if Kind(a) != Number {
			break
		}

		part = append(part, a)
	}

	current = i
	return
}
