//go:build go1.17
// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package byte

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyByteSlice(dst, src []byte) {
	// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
	// in order to not panic by getting an array that is bigger than len(dst)
	if len(dst) < len(src) {
		switch len(dst) {
		
		case 0:
			copyByteSlice0(dst, src)
			return
		
		case 1:
			copyByteSlice1(dst, src)
			return
		
		case 2:
			copyByteSlice2(dst, src)
			return
		
		case 3:
			copyByteSlice3(dst, src)
			return
		
		case 4:
			copyByteSlice4(dst, src)
			return
		
		case 5:
			copyByteSlice5(dst, src)
			return
		
		default:
			// If len(dst) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
			copy(dst, src)
			return
		}
	}

	// If len(src) is within our limits and greater than len(dst), then we need to copy with the
	// size equal to len(src) in order to not panic by getting an array that is bigger than len(src)
	switch len(src) {
	
	case 0:
		copyByteSlice0(dst, src)
		return
	
	case 1:
		copyByteSlice1(dst, src)
		return
	
	case 2:
		copyByteSlice2(dst, src)
		return
	
	case 3:
		copyByteSlice3(dst, src)
		return
	
	case 4:
		copyByteSlice4(dst, src)
		return
	
	case 5:
		copyByteSlice5(dst, src)
		return
	
	default:
		// If len(dst) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
		copy(dst, src)
		return
	}
}

func copyByteSlice0(dst, src []byte) {
	*(*[0]byte)(dst) = *(*[0]byte)(src)
}

func copyByteSlice1(dst, src []byte) {
	*(*[1]byte)(dst) = *(*[1]byte)(src)
}

func copyByteSlice2(dst, src []byte) {
	*(*[2]byte)(dst) = *(*[2]byte)(src)
}

func copyByteSlice3(dst, src []byte) {
	*(*[3]byte)(dst) = *(*[3]byte)(src)
}

func copyByteSlice4(dst, src []byte) {
	*(*[4]byte)(dst) = *(*[4]byte)(src)
}

func copyByteSlice5(dst, src []byte) {
	*(*[5]byte)(dst) = *(*[5]byte)(src)
}
