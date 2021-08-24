// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package uint64

func CopyUint64Slice(dst, src []uint64) {
	// If len(src) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
	if len(src) > 5 {
		copy(dst, src)
		return
	}

	if len(dst) < len(src) {
		// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
		// in order to not panic by getting an array that is bigger than len(dst)
		copyUint64SliceIdx[len(dst)](dst, src)
		return
	}

	// If len(src) is within our limits and greater than len(dst), then we need to copy with the
	// size equal to len(src) in order to not panic by getting an array that is bigger than len(src)
	copyUint64SliceIdx[len(src)](dst, src)
}

var copyUint64SliceIdx = [6]func([]uint64, []uint64){
	
	0: copyUint64Slice0,
	
	1: copyUint64Slice1,
	
	2: copyUint64Slice2,
	
	3: copyUint64Slice3,
	
	4: copyUint64Slice4,
	
	5: copyUint64Slice5,
	
}

func copyUint64Slice0(dst, src []uint64) {
	*(*[0]uint64)(dst) = *(*[0]uint64)(src)
}

func copyUint64Slice1(dst, src []uint64) {
	*(*[1]uint64)(dst) = *(*[1]uint64)(src)
}

func copyUint64Slice2(dst, src []uint64) {
	*(*[2]uint64)(dst) = *(*[2]uint64)(src)
}

func copyUint64Slice3(dst, src []uint64) {
	*(*[3]uint64)(dst) = *(*[3]uint64)(src)
}

func copyUint64Slice4(dst, src []uint64) {
	*(*[4]uint64)(dst) = *(*[4]uint64)(src)
}

func copyUint64Slice5(dst, src []uint64) {
	*(*[5]uint64)(dst) = *(*[5]uint64)(src)
}
