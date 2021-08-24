// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package bool

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyBoolSlice(dst, src []bool) {
	// If len(src) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
	if len(src) > 5 {
		copy(dst, src)
		return
	}

	if len(dst) < len(src) {
		// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
		// in order to not panic by getting an array that is bigger than len(dst)
		copyBoolSliceIdx[len(dst)](dst, src)
		return
	}

	// If len(src) is within our limits and greater than len(dst), then we need to copy with the
	// size equal to len(src) in order to not panic by getting an array that is bigger than len(src)
	copyBoolSliceIdx[len(src)](dst, src)
}

var copyBoolSliceIdx = [6]func([]bool, []bool){
	
	0: copyBoolSlice0,
	
	1: copyBoolSlice1,
	
	2: copyBoolSlice2,
	
	3: copyBoolSlice3,
	
	4: copyBoolSlice4,
	
	5: copyBoolSlice5,
	
}

func copyBoolSlice0(dst, src []bool) {
	*(*[0]bool)(dst) = *(*[0]bool)(src)
}

func copyBoolSlice1(dst, src []bool) {
	*(*[1]bool)(dst) = *(*[1]bool)(src)
}

func copyBoolSlice2(dst, src []bool) {
	*(*[2]bool)(dst) = *(*[2]bool)(src)
}

func copyBoolSlice3(dst, src []bool) {
	*(*[3]bool)(dst) = *(*[3]bool)(src)
}

func copyBoolSlice4(dst, src []bool) {
	*(*[4]bool)(dst) = *(*[4]bool)(src)
}

func copyBoolSlice5(dst, src []bool) {
	*(*[5]bool)(dst) = *(*[5]bool)(src)
}
