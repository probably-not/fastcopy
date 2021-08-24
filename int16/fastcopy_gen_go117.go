// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package int16

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyInt16Slice(dst, src []int16) {
	// If len(src) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
	if len(src) > 5 {
		copy(dst, src)
		return
	}

	if len(dst) < len(src) {
		// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
		// in order to not panic by getting an array that is bigger than len(dst)
		copyInt16SliceIdx[len(dst)](dst, src)
		return
	}

	// If len(src) is within our limits and greater than len(dst), then we need to copy with the
	// size equal to len(src) in order to not panic by getting an array that is bigger than len(src)
	copyInt16SliceIdx[len(src)](dst, src)
}

var copyInt16SliceIdx = [6]func([]int16, []int16){
	
	0: copyInt16Slice0,
	
	1: copyInt16Slice1,
	
	2: copyInt16Slice2,
	
	3: copyInt16Slice3,
	
	4: copyInt16Slice4,
	
	5: copyInt16Slice5,
	
}

func copyInt16Slice0(dst, src []int16) {
	*(*[0]int16)(dst) = *(*[0]int16)(src)
}

func copyInt16Slice1(dst, src []int16) {
	*(*[1]int16)(dst) = *(*[1]int16)(src)
}

func copyInt16Slice2(dst, src []int16) {
	*(*[2]int16)(dst) = *(*[2]int16)(src)
}

func copyInt16Slice3(dst, src []int16) {
	*(*[3]int16)(dst) = *(*[3]int16)(src)
}

func copyInt16Slice4(dst, src []int16) {
	*(*[4]int16)(dst) = *(*[4]int16)(src)
}

func copyInt16Slice5(dst, src []int16) {
	*(*[5]int16)(dst) = *(*[5]int16)(src)
}
