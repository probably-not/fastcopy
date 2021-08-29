//go:build go1.17
// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package bool

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyBoolSlice(dst, src []bool) {
	// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
	// in order to not panic by getting an array that is bigger than len(dst)
	if len(dst) < len(src) {
		switch len(dst) {
		
		case 0:
			copyBoolSlice0(dst, src)
			return
		
		case 1:
			copyBoolSlice1(dst, src)
			return
		
		case 2:
			copyBoolSlice2(dst, src)
			return
		
		case 3:
			copyBoolSlice3(dst, src)
			return
		
		case 4:
			copyBoolSlice4(dst, src)
			return
		
		case 5:
			copyBoolSlice5(dst, src)
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
		copyBoolSlice0(dst, src)
		return
	
	case 1:
		copyBoolSlice1(dst, src)
		return
	
	case 2:
		copyBoolSlice2(dst, src)
		return
	
	case 3:
		copyBoolSlice3(dst, src)
		return
	
	case 4:
		copyBoolSlice4(dst, src)
		return
	
	case 5:
		copyBoolSlice5(dst, src)
		return
	
	default:
		// If len(dst) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
		copy(dst, src)
		return
	}
}

func copyBoolSlice0(dst, src []bool) {
	*(*[0]bool)(dst) = *(*[0]bool)(src)
}

func copyBoolSlice1(dst, src []bool) {
	*(*[1]bool)(dst) = *(*[1]bool)(src)
}

func copyBoolSlice2(dst, src []bool) {
	*(*[2]bool)(dst) = *(*[2]bool)(src)
}

func copyBoolSlice3(dst, src []bool) {
	*(*[3]bool)(dst) = *(*[3]bool)(src)
}

func copyBoolSlice4(dst, src []bool) {
	*(*[4]bool)(dst) = *(*[4]bool)(src)
}

func copyBoolSlice5(dst, src []bool) {
	*(*[5]bool)(dst) = *(*[5]bool)(src)
}
