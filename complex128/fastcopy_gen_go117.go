// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package complex128

func CopyComplex128Slice(dst, src []complex128) {
	// If len(src) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
	if len(src) > 5 {
		copy(dst, src)
		return
	}

	if len(dst) < len(src) {
		// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
		// in order to not panic by getting an array that is bigger than len(dst)
		copyComplex128SliceIdx[len(dst)](dst, src)
		return
	}

	// If len(src) is within our limits and greater than len(dst), then we need to copy with the
	// size equal to len(src) in order to not panic by getting an array that is bigger than len(src)
	copyComplex128SliceIdx[len(src)](dst, src)
}

var copyComplex128SliceIdx = [6]func([]complex128, []complex128){
	
	0: copyComplex128Slice0,
	
	1: copyComplex128Slice1,
	
	2: copyComplex128Slice2,
	
	3: copyComplex128Slice3,
	
	4: copyComplex128Slice4,
	
	5: copyComplex128Slice5,
	
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
