//go:build go1.17
// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package uint16

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyUint16Slice(dst, src []uint16) {
	// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
	// in order to not panic by getting an array that is bigger than len(dst)
	if len(dst) < len(src) {
		switch len(dst) {
		
		case 0:
			copyUint16Slice0(dst, src)
			return
		
		case 1:
			copyUint16Slice1(dst, src)
			return
		
		case 2:
			copyUint16Slice2(dst, src)
			return
		
		case 3:
			copyUint16Slice3(dst, src)
			return
		
		case 4:
			copyUint16Slice4(dst, src)
			return
		
		case 5:
			copyUint16Slice5(dst, src)
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
		copyUint16Slice0(dst, src)
		return
	
	case 1:
		copyUint16Slice1(dst, src)
		return
	
	case 2:
		copyUint16Slice2(dst, src)
		return
	
	case 3:
		copyUint16Slice3(dst, src)
		return
	
	case 4:
		copyUint16Slice4(dst, src)
		return
	
	case 5:
		copyUint16Slice5(dst, src)
		return
	
	default:
		// If len(dst) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
		copy(dst, src)
		return
	}
}

func copyUint16Slice0(dst, src []uint16) {
	*(*[0]uint16)(dst) = *(*[0]uint16)(src)
}

func copyUint16Slice1(dst, src []uint16) {
	*(*[1]uint16)(dst) = *(*[1]uint16)(src)
}

func copyUint16Slice2(dst, src []uint16) {
	*(*[2]uint16)(dst) = *(*[2]uint16)(src)
}

func copyUint16Slice3(dst, src []uint16) {
	*(*[3]uint16)(dst) = *(*[3]uint16)(src)
}

func copyUint16Slice4(dst, src []uint16) {
	*(*[4]uint16)(dst) = *(*[4]uint16)(src)
}

func copyUint16Slice5(dst, src []uint16) {
	*(*[5]uint16)(dst) = *(*[5]uint16)(src)
}
