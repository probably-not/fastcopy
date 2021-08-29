//go:build go1.17
// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package string

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyStringSlice(dst, src []string) {
	// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
	// in order to not panic by getting an array that is bigger than len(dst)
	if len(dst) < len(src) {
		switch len(dst) {
		
		case 0:
			copyStringSlice0(dst, src)
			return
		
		case 1:
			copyStringSlice1(dst, src)
			return
		
		case 2:
			copyStringSlice2(dst, src)
			return
		
		case 3:
			copyStringSlice3(dst, src)
			return
		
		case 4:
			copyStringSlice4(dst, src)
			return
		
		case 5:
			copyStringSlice5(dst, src)
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
		copyStringSlice0(dst, src)
		return
	
	case 1:
		copyStringSlice1(dst, src)
		return
	
	case 2:
		copyStringSlice2(dst, src)
		return
	
	case 3:
		copyStringSlice3(dst, src)
		return
	
	case 4:
		copyStringSlice4(dst, src)
		return
	
	case 5:
		copyStringSlice5(dst, src)
		return
	
	default:
		// If len(dst) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
		copy(dst, src)
		return
	}
}

func copyStringSlice0(dst, src []string) {
	*(*[0]string)(dst) = *(*[0]string)(src)
}

func copyStringSlice1(dst, src []string) {
	*(*[1]string)(dst) = *(*[1]string)(src)
}

func copyStringSlice2(dst, src []string) {
	*(*[2]string)(dst) = *(*[2]string)(src)
}

func copyStringSlice3(dst, src []string) {
	*(*[3]string)(dst) = *(*[3]string)(src)
}

func copyStringSlice4(dst, src []string) {
	*(*[4]string)(dst) = *(*[4]string)(src)
}

func copyStringSlice5(dst, src []string) {
	*(*[5]string)(dst) = *(*[5]string)(src)
}
