//go:build go1.17
// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package int64

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyInt64Slice(dst, src []int64) {
	// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
	// in order to not panic by getting an array that is bigger than len(dst)
	if len(dst) < len(src) {
		switch len(dst) {
		
		case 0:
			copyInt64Slice0(dst, src)
			return
		
		case 1:
			copyInt64Slice1(dst, src)
			return
		
		case 2:
			copyInt64Slice2(dst, src)
			return
		
		case 3:
			copyInt64Slice3(dst, src)
			return
		
		case 4:
			copyInt64Slice4(dst, src)
			return
		
		case 5:
			copyInt64Slice5(dst, src)
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
		copyInt64Slice0(dst, src)
		return
	
	case 1:
		copyInt64Slice1(dst, src)
		return
	
	case 2:
		copyInt64Slice2(dst, src)
		return
	
	case 3:
		copyInt64Slice3(dst, src)
		return
	
	case 4:
		copyInt64Slice4(dst, src)
		return
	
	case 5:
		copyInt64Slice5(dst, src)
		return
	
	default:
		// If len(dst) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
		copy(dst, src)
		return
	}
}

func copyInt64Slice0(dst, src []int64) {
	*(*[0]int64)(dst) = *(*[0]int64)(src)
}

func copyInt64Slice1(dst, src []int64) {
	*(*[1]int64)(dst) = *(*[1]int64)(src)
}

func copyInt64Slice2(dst, src []int64) {
	*(*[2]int64)(dst) = *(*[2]int64)(src)
}

func copyInt64Slice3(dst, src []int64) {
	*(*[3]int64)(dst) = *(*[3]int64)(src)
}

func copyInt64Slice4(dst, src []int64) {
	*(*[4]int64)(dst) = *(*[4]int64)(src)
}

func copyInt64Slice5(dst, src []int64) {
	*(*[5]int64)(dst) = *(*[5]int64)(src)
}
