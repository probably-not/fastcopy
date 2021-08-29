//go:build go1.17
// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package int

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyIntSlice(dst, src []int) {
	// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
	// in order to not panic by getting an array that is bigger than len(dst)
	if len(dst) < len(src) {
		switch len(dst) {
		
		case 0:
			copyIntSlice0(dst, src)
			return
		
		case 1:
			copyIntSlice1(dst, src)
			return
		
		case 2:
			copyIntSlice2(dst, src)
			return
		
		case 3:
			copyIntSlice3(dst, src)
			return
		
		case 4:
			copyIntSlice4(dst, src)
			return
		
		case 5:
			copyIntSlice5(dst, src)
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
		copyIntSlice0(dst, src)
		return
	
	case 1:
		copyIntSlice1(dst, src)
		return
	
	case 2:
		copyIntSlice2(dst, src)
		return
	
	case 3:
		copyIntSlice3(dst, src)
		return
	
	case 4:
		copyIntSlice4(dst, src)
		return
	
	case 5:
		copyIntSlice5(dst, src)
		return
	
	default:
		// If len(dst) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
		copy(dst, src)
		return
	}
}

func copyIntSlice0(dst, src []int) {
	*(*[0]int)(dst) = *(*[0]int)(src)
}

func copyIntSlice1(dst, src []int) {
	*(*[1]int)(dst) = *(*[1]int)(src)
}

func copyIntSlice2(dst, src []int) {
	*(*[2]int)(dst) = *(*[2]int)(src)
}

func copyIntSlice3(dst, src []int) {
	*(*[3]int)(dst) = *(*[3]int)(src)
}

func copyIntSlice4(dst, src []int) {
	*(*[4]int)(dst) = *(*[4]int)(src)
}

func copyIntSlice5(dst, src []int) {
	*(*[5]int)(dst) = *(*[5]int)(src)
}
