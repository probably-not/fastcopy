// +build !go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package byte

func CopyByteSlice(dst, src []byte) {
	copy(dst, src)
}
