package util

import (
	"runtime"
	"strings"
)

// `ParseGoVersion` takes the go version from `runtime.Version()`
// and parses it to return major, minor, and patch int64 values.
func ParseGoVersion() (int64, int64, int64) {
	str := strings.TrimPrefix(runtime.Version(), "go")
	var major, minor, patch int64
	if len(str) == 0 || str[0] < '0' || str[0] > '9' {
		return major, minor, patch
	}
	for i := 0; i < 3; i++ {
		empty := true
		val := 0
		l := len(str) - 1

		for k, c := range str {
			if c >= '0' && c <= '9' {
				if empty {
					val = int(c) - 48
					empty = false
					if k == l {
						str = str[:0]
					}
					continue
				}

				if val == 0 {
					if c == '0' {
						if k == l {
							str = str[:0]
						}
						continue
					}
					str = str[k:]
					break
				}

				val = 10*val + int(c) - 48
				if k == l {
					str = str[:0]
				}
				continue
			}
			str = str[k+1:]
			break
		}

		switch i {
		case 0:
			major = int64(val)

		case 1:
			minor = int64(val)

		case 2:
			patch = int64(val)
		}
	}
	return major, minor, patch
}
