// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package rune

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyRuneSlice(dst, src []rune) {
	// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
	// in order to not panic by getting an array that is bigger than len(dst)
	if len(dst) < len(src) {
		switch len(dst) {
		
		case 0:
			copyRuneSlice0(dst, src)
			return
		
		case 1:
			copyRuneSlice1(dst, src)
			return
		
		case 2:
			copyRuneSlice2(dst, src)
			return
		
		case 3:
			copyRuneSlice3(dst, src)
			return
		
		case 4:
			copyRuneSlice4(dst, src)
			return
		
		case 5:
			copyRuneSlice5(dst, src)
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
		copyRuneSlice0(dst, src)
		return
	
	case 1:
		copyRuneSlice1(dst, src)
		return
	
	case 2:
		copyRuneSlice2(dst, src)
		return
	
	case 3:
		copyRuneSlice3(dst, src)
		return
	
	case 4:
		copyRuneSlice4(dst, src)
		return
	
	case 5:
		copyRuneSlice5(dst, src)
		return
	
	default:
		// If len(dst) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
		copy(dst, src)
		return
	}
}

func copyRuneSlice0(dst, src []rune) {
	*(*[0]rune)(dst) = *(*[0]rune)(src)
}

func copyRuneSlice1(dst, src []rune) {
	*(*[1]rune)(dst) = *(*[1]rune)(src)
}

func copyRuneSlice2(dst, src []rune) {
	*(*[2]rune)(dst) = *(*[2]rune)(src)
}

func copyRuneSlice3(dst, src []rune) {
	*(*[3]rune)(dst) = *(*[3]rune)(src)
}

func copyRuneSlice4(dst, src []rune) {
	*(*[4]rune)(dst) = *(*[4]rune)(src)
}

func copyRuneSlice5(dst, src []rune) {
	*(*[5]rune)(dst) = *(*[5]rune)(src)
}
