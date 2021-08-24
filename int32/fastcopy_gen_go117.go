// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package int32

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyInt32Slice(dst, src []int32) {
	// If len(src) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
	if len(src) > 5 {
		copy(dst, src)
		return
	}

	if len(dst) < len(src) {
		// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
		// in order to not panic by getting an array that is bigger than len(dst)
		copyInt32SliceIdx[len(dst)](dst, src)
		return
	}

	// If len(src) is within our limits and greater than len(dst), then we need to copy with the
	// size equal to len(src) in order to not panic by getting an array that is bigger than len(src)
	copyInt32SliceIdx[len(src)](dst, src)
}

var copyInt32SliceIdx = [6]func([]int32, []int32){
	
	0: copyInt32Slice0,
	
	1: copyInt32Slice1,
	
	2: copyInt32Slice2,
	
	3: copyInt32Slice3,
	
	4: copyInt32Slice4,
	
	5: copyInt32Slice5,
	
}

func copyInt32Slice0(dst, src []int32) {
	*(*[0]int32)(dst) = *(*[0]int32)(src)
}

func copyInt32Slice1(dst, src []int32) {
	*(*[1]int32)(dst) = *(*[1]int32)(src)
}

func copyInt32Slice2(dst, src []int32) {
	*(*[2]int32)(dst) = *(*[2]int32)(src)
}

func copyInt32Slice3(dst, src []int32) {
	*(*[3]int32)(dst) = *(*[3]int32)(src)
}

func copyInt32Slice4(dst, src []int32) {
	*(*[4]int32)(dst) = *(*[4]int32)(src)
}

func copyInt32Slice5(dst, src []int32) {
	*(*[5]int32)(dst) = *(*[5]int32)(src)
}
