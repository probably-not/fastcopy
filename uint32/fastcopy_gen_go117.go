//go:build go1.17
// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package uint32

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyUint32Slice(dst, src []uint32) {
	// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
	// in order to not panic by getting an array that is bigger than len(dst)
	if len(dst) < len(src) {
		switch len(dst) {
		
		case 0:
			copyUint32Slice0(dst, src)
			return
		
		case 1:
			copyUint32Slice1(dst, src)
			return
		
		case 2:
			copyUint32Slice2(dst, src)
			return
		
		case 3:
			copyUint32Slice3(dst, src)
			return
		
		case 4:
			copyUint32Slice4(dst, src)
			return
		
		case 5:
			copyUint32Slice5(dst, src)
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
		copyUint32Slice0(dst, src)
		return
	
	case 1:
		copyUint32Slice1(dst, src)
		return
	
	case 2:
		copyUint32Slice2(dst, src)
		return
	
	case 3:
		copyUint32Slice3(dst, src)
		return
	
	case 4:
		copyUint32Slice4(dst, src)
		return
	
	case 5:
		copyUint32Slice5(dst, src)
		return
	
	default:
		// If len(dst) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
		copy(dst, src)
		return
	}
}

func copyUint32Slice0(dst, src []uint32) {
	*(*[0]uint32)(dst) = *(*[0]uint32)(src)
}

func copyUint32Slice1(dst, src []uint32) {
	*(*[1]uint32)(dst) = *(*[1]uint32)(src)
}

func copyUint32Slice2(dst, src []uint32) {
	*(*[2]uint32)(dst) = *(*[2]uint32)(src)
}

func copyUint32Slice3(dst, src []uint32) {
	*(*[3]uint32)(dst) = *(*[3]uint32)(src)
}

func copyUint32Slice4(dst, src []uint32) {
	*(*[4]uint32)(dst) = *(*[4]uint32)(src)
}

func copyUint32Slice5(dst, src []uint32) {
	*(*[5]uint32)(dst) = *(*[5]uint32)(src)
}
