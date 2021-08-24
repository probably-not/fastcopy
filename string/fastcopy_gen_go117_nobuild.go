// +build !go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package string

func CopyStringSlice(dst, src []string) {
	copy(dst, src)
}
