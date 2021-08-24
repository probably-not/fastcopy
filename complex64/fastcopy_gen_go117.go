// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package complex64

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyComplex64Slice(dst, src []complex64) {
	// If len(src) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
	if len(src) > 5 {
		copy(dst, src)
		return
	}

	if len(dst) < len(src) {
		// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
		// in order to not panic by getting an array that is bigger than len(dst)
		copyComplex64SliceIdx[len(dst)](dst, src)
		return
	}

	// If len(src) is within our limits and greater than len(dst), then we need to copy with the
	// size equal to len(src) in order to not panic by getting an array that is bigger than len(src)
	copyComplex64SliceIdx[len(src)](dst, src)
}

var copyComplex64SliceIdx = [6]func([]complex64, []complex64){
	
	0: copyComplex64Slice0,
	
	1: copyComplex64Slice1,
	
	2: copyComplex64Slice2,
	
	3: copyComplex64Slice3,
	
	4: copyComplex64Slice4,
	
	5: copyComplex64Slice5,
	
}

func copyComplex64Slice0(dst, src []complex64) {
	*(*[0]complex64)(dst) = *(*[0]complex64)(src)
}

func copyComplex64Slice1(dst, src []complex64) {
	*(*[1]complex64)(dst) = *(*[1]complex64)(src)
}

func copyComplex64Slice2(dst, src []complex64) {
	*(*[2]complex64)(dst) = *(*[2]complex64)(src)
}

func copyComplex64Slice3(dst, src []complex64) {
	*(*[3]complex64)(dst) = *(*[3]complex64)(src)
}

func copyComplex64Slice4(dst, src []complex64) {
	*(*[4]complex64)(dst) = *(*[4]complex64)(src)
}

func copyComplex64Slice5(dst, src []complex64) {
	*(*[5]complex64)(dst) = *(*[5]complex64)(src)
}
