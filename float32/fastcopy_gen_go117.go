// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package float32

func CopyFloat32Slice(dst, src []float32) {
	// If len(src) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
	if len(src) > 5 {
		copy(dst, src)
		return
	}

	if len(dst) < len(src) {
		// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
		// in order to not panic by getting an array that is bigger than len(dst)
		copyFloat32SliceIdx[len(dst)](dst, src)
		return
	}

	// If len(src) is within our limits and greater than len(dst), then we need to copy with the
	// size equal to len(src) in order to not panic by getting an array that is bigger than len(src)
	copyFloat32SliceIdx[len(src)](dst, src)
}

var copyFloat32SliceIdx = [6]func([]float32, []float32){
	
	0: copyFloat32Slice0,
	
	1: copyFloat32Slice1,
	
	2: copyFloat32Slice2,
	
	3: copyFloat32Slice3,
	
	4: copyFloat32Slice4,
	
	5: copyFloat32Slice5,
	
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
