// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package float64

func CopyFloat64Slice(dst, src []float64) {
	// If len(src) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
	if len(src) > 5 {
		copy(dst, src)
		return
	}

	if len(dst) < len(src) {
		// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
		// in order to not panic by getting an array that is bigger than len(dst)
		copyFloat64SliceIdx[len(dst)](dst, src)
		return
	}

	// If len(src) is within our limits and greater than len(dst), then we need to copy with the
	// size equal to len(src) in order to not panic by getting an array that is bigger than len(src)
	copyFloat64SliceIdx[len(src)](dst, src)
}

var copyFloat64SliceIdx = [6]func([]float64, []float64){
	
	0: copyFloat64Slice0,
	
	1: copyFloat64Slice1,
	
	2: copyFloat64Slice2,
	
	3: copyFloat64Slice3,
	
	4: copyFloat64Slice4,
	
	5: copyFloat64Slice5,
	
}

func copyFloat64Slice0(dst, src []float64) {
	*(*[0]float64)(dst) = *(*[0]float64)(src)
}

func copyFloat64Slice1(dst, src []float64) {
	*(*[1]float64)(dst) = *(*[1]float64)(src)
}

func copyFloat64Slice2(dst, src []float64) {
	*(*[2]float64)(dst) = *(*[2]float64)(src)
}

func copyFloat64Slice3(dst, src []float64) {
	*(*[3]float64)(dst) = *(*[3]float64)(src)
}

func copyFloat64Slice4(dst, src []float64) {
	*(*[4]float64)(dst) = *(*[4]float64)(src)
}

func copyFloat64Slice5(dst, src []float64) {
	*(*[5]float64)(dst) = *(*[5]float64)(src)
}
