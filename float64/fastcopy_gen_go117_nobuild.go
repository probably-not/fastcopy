//go:build !go1.17
// +build !go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package float64

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = false

func CopyFloat64Slice(dst, src []float64) {
	copy(dst, src)
}
