// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package uint

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyUintSlice(dst, src []uint) {
	// If len(src) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
	if len(src) > 5 {
		copy(dst, src)
		return
	}

	if len(dst) < len(src) {
		// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
		// in order to not panic by getting an array that is bigger than len(dst)
		copyUintSliceIdx[len(dst)](dst, src)
		return
	}

	// If len(src) is within our limits and greater than len(dst), then we need to copy with the
	// size equal to len(src) in order to not panic by getting an array that is bigger than len(src)
	copyUintSliceIdx[len(src)](dst, src)
}

var copyUintSliceIdx = [6]func([]uint, []uint){
	
	0: copyUintSlice0,
	
	1: copyUintSlice1,
	
	2: copyUintSlice2,
	
	3: copyUintSlice3,
	
	4: copyUintSlice4,
	
	5: copyUintSlice5,
	
}

func copyUintSlice0(dst, src []uint) {
	*(*[0]uint)(dst) = *(*[0]uint)(src)
}

func copyUintSlice1(dst, src []uint) {
	*(*[1]uint)(dst) = *(*[1]uint)(src)
}

func copyUintSlice2(dst, src []uint) {
	*(*[2]uint)(dst) = *(*[2]uint)(src)
}

func copyUintSlice3(dst, src []uint) {
	*(*[3]uint)(dst) = *(*[3]uint)(src)
}

func copyUintSlice4(dst, src []uint) {
	*(*[4]uint)(dst) = *(*[4]uint)(src)
}

func copyUintSlice5(dst, src []uint) {
	*(*[5]uint)(dst) = *(*[5]uint)(src)
}
