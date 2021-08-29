//go:build !go1.17
// +build !go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package uint

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = false

func CopyUintSlice(dst, src []uint) {
	copy(dst, src)
}
