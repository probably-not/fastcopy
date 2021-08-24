// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package int8

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyInt8Slice(dst, src []int8) {
	// If len(src) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
	if len(src) > 5 {
		copy(dst, src)
		return
	}

	if len(dst) < len(src) {
		// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
		// in order to not panic by getting an array that is bigger than len(dst)
		copyInt8SliceIdx[len(dst)](dst, src)
		return
	}

	// If len(src) is within our limits and greater than len(dst), then we need to copy with the
	// size equal to len(src) in order to not panic by getting an array that is bigger than len(src)
	copyInt8SliceIdx[len(src)](dst, src)
}

var copyInt8SliceIdx = [6]func([]int8, []int8){
	
	0: copyInt8Slice0,
	
	1: copyInt8Slice1,
	
	2: copyInt8Slice2,
	
	3: copyInt8Slice3,
	
	4: copyInt8Slice4,
	
	5: copyInt8Slice5,
	
}

func copyInt8Slice0(dst, src []int8) {
	*(*[0]int8)(dst) = *(*[0]int8)(src)
}

func copyInt8Slice1(dst, src []int8) {
	*(*[1]int8)(dst) = *(*[1]int8)(src)
}

func copyInt8Slice2(dst, src []int8) {
	*(*[2]int8)(dst) = *(*[2]int8)(src)
}

func copyInt8Slice3(dst, src []int8) {
	*(*[3]int8)(dst) = *(*[3]int8)(src)
}

func copyInt8Slice4(dst, src []int8) {
	*(*[4]int8)(dst) = *(*[4]int8)(src)
}

func copyInt8Slice5(dst, src []int8) {
	*(*[5]int8)(dst) = *(*[5]int8)(src)
}
