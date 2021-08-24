// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package string

func CopyStringSlice(dst, src []string) {
	// If len(src) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
	if len(src) > 5 {
		copy(dst, src)
		return
	}

	if len(dst) < len(src) {
		// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
		// in order to not panic by getting an array that is bigger than len(dst)
		copyStringSliceIdx[len(dst)](dst, src)
		return
	}

	// If len(src) is within our limits and greater than len(dst), then we need to copy with the
	// size equal to len(src) in order to not panic by getting an array that is bigger than len(src)
	copyStringSliceIdx[len(src)](dst, src)
}

var copyStringSliceIdx = [6]func([]string, []string){
	
	0: copyStringSlice0,
	
	1: copyStringSlice1,
	
	2: copyStringSlice2,
	
	3: copyStringSlice3,
	
	4: copyStringSlice4,
	
	5: copyStringSlice5,
	
}

func copyStringSlice0(dst, src []string) {
	*(*[0]string)(dst) = *(*[0]string)(src)
}

func copyStringSlice1(dst, src []string) {
	*(*[1]string)(dst) = *(*[1]string)(src)
}

func copyStringSlice2(dst, src []string) {
	*(*[2]string)(dst) = *(*[2]string)(src)
}

func copyStringSlice3(dst, src []string) {
	*(*[3]string)(dst) = *(*[3]string)(src)
}

func copyStringSlice4(dst, src []string) {
	*(*[4]string)(dst) = *(*[4]string)(src)
}

func copyStringSlice5(dst, src []string) {
	*(*[5]string)(dst) = *(*[5]string)(src)
}
