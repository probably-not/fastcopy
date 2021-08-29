//go:build go1.17
// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package float32

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyFloat32Slice(dst, src []float32) {
	// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
	// in order to not panic by getting an array that is bigger than len(dst)
	if len(dst) < len(src) {
		switch len(dst) {
		
		case 0:
			copyFloat32Slice0(dst, src)
			return
		
		case 1:
			copyFloat32Slice1(dst, src)
			return
		
		case 2:
			copyFloat32Slice2(dst, src)
			return
		
		case 3:
			copyFloat32Slice3(dst, src)
			return
		
		case 4:
			copyFloat32Slice4(dst, src)
			return
		
		case 5:
			copyFloat32Slice5(dst, src)
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
		copyFloat32Slice0(dst, src)
		return
	
	case 1:
		copyFloat32Slice1(dst, src)
		return
	
	case 2:
		copyFloat32Slice2(dst, src)
		return
	
	case 3:
		copyFloat32Slice3(dst, src)
		return
	
	case 4:
		copyFloat32Slice4(dst, src)
		return
	
	case 5:
		copyFloat32Slice5(dst, src)
		return
	
	default:
		// If len(dst) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
		copy(dst, src)
		return
	}
}

func copyFloat32Slice0(dst, src []float32) {
	*(*[0]float32)(dst) = *(*[0]float32)(src)
}

func copyFloat32Slice1(dst, src []float32) {
	*(*[1]float32)(dst) = *(*[1]float32)(src)
}

func copyFloat32Slice2(dst, src []float32) {
	*(*[2]float32)(dst) = *(*[2]float32)(src)
}

func copyFloat32Slice3(dst, src []float32) {
	*(*[3]float32)(dst) = *(*[3]float32)(src)
}

func copyFloat32Slice4(dst, src []float32) {
	*(*[4]float32)(dst) = *(*[4]float32)(src)
}

func copyFloat32Slice5(dst, src []float32) {
	*(*[5]float32)(dst) = *(*[5]float32)(src)
}
