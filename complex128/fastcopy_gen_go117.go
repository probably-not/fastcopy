// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package complex128

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyComplex128Slice(dst, src []complex128) {
	// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
	// in order to not panic by getting an array that is bigger than len(dst)
	if len(dst) < len(src) {
		switch len(dst) {
		
		case 0:
			copyComplex128Slice0(dst, src)
			return
		
		case 1:
			copyComplex128Slice1(dst, src)
			return
		
		case 2:
			copyComplex128Slice2(dst, src)
			return
		
		case 3:
			copyComplex128Slice3(dst, src)
			return
		
		case 4:
			copyComplex128Slice4(dst, src)
			return
		
		case 5:
			copyComplex128Slice5(dst, src)
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
		copyComplex128Slice0(dst, src)
		return
	
	case 1:
		copyComplex128Slice1(dst, src)
		return
	
	case 2:
		copyComplex128Slice2(dst, src)
		return
	
	case 3:
		copyComplex128Slice3(dst, src)
		return
	
	case 4:
		copyComplex128Slice4(dst, src)
		return
	
	case 5:
		copyComplex128Slice5(dst, src)
		return
	
	default:
		// If len(dst) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
		copy(dst, src)
		return
	}
}

func copyComplex128Slice0(dst, src []complex128) {
	*(*[0]complex128)(dst) = *(*[0]complex128)(src)
}

func copyComplex128Slice1(dst, src []complex128) {
	*(*[1]complex128)(dst) = *(*[1]complex128)(src)
}

func copyComplex128Slice2(dst, src []complex128) {
	*(*[2]complex128)(dst) = *(*[2]complex128)(src)
}

func copyComplex128Slice3(dst, src []complex128) {
	*(*[3]complex128)(dst) = *(*[3]complex128)(src)
}

func copyComplex128Slice4(dst, src []complex128) {
	*(*[4]complex128)(dst) = *(*[4]complex128)(src)
}

func copyComplex128Slice5(dst, src []complex128) {
	*(*[5]complex128)(dst) = *(*[5]complex128)(src)
}
