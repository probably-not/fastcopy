// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package int

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyIntSlice(dst, src []int) {
	// If len(src) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
	if len(src) > 5 {
		copy(dst, src)
		return
	}

	if len(dst) < len(src) {
		// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
		// in order to not panic by getting an array that is bigger than len(dst)
		copyIntSliceIdx[len(dst)](dst, src)
		return
	}

	// If len(src) is within our limits and greater than len(dst), then we need to copy with the
	// size equal to len(src) in order to not panic by getting an array that is bigger than len(src)
	copyIntSliceIdx[len(src)](dst, src)
}

var copyIntSliceIdx = [6]func([]int, []int){
	
	0: copyIntSlice0,
	
	1: copyIntSlice1,
	
	2: copyIntSlice2,
	
	3: copyIntSlice3,
	
	4: copyIntSlice4,
	
	5: copyIntSlice5,
	
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
