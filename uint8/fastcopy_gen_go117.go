//go:build go1.17
// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package uint8

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyUint8Slice(dst, src []uint8) {
	// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
	// in order to not panic by getting an array that is bigger than len(dst)
	if len(dst) < len(src) {
		switch len(dst) {
		
		case 0:
			copyUint8Slice0(dst, src)
			return
		
		case 1:
			copyUint8Slice1(dst, src)
			return
		
		case 2:
			copyUint8Slice2(dst, src)
			return
		
		case 3:
			copyUint8Slice3(dst, src)
			return
		
		case 4:
			copyUint8Slice4(dst, src)
			return
		
		case 5:
			copyUint8Slice5(dst, src)
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
		copyUint8Slice0(dst, src)
		return
	
	case 1:
		copyUint8Slice1(dst, src)
		return
	
	case 2:
		copyUint8Slice2(dst, src)
		return
	
	case 3:
		copyUint8Slice3(dst, src)
		return
	
	case 4:
		copyUint8Slice4(dst, src)
		return
	
	case 5:
		copyUint8Slice5(dst, src)
		return
	
	default:
		// If len(dst) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
		copy(dst, src)
		return
	}
}

func copyUint8Slice0(dst, src []uint8) {
	*(*[0]uint8)(dst) = *(*[0]uint8)(src)
}

func copyUint8Slice1(dst, src []uint8) {
	*(*[1]uint8)(dst) = *(*[1]uint8)(src)
}

func copyUint8Slice2(dst, src []uint8) {
	*(*[2]uint8)(dst) = *(*[2]uint8)(src)
}

func copyUint8Slice3(dst, src []uint8) {
	*(*[3]uint8)(dst) = *(*[3]uint8)(src)
}

func copyUint8Slice4(dst, src []uint8) {
	*(*[4]uint8)(dst) = *(*[4]uint8)(src)
}

func copyUint8Slice5(dst, src []uint8) {
	*(*[5]uint8)(dst) = *(*[5]uint8)(src)
}
