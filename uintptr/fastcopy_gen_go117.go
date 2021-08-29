//go:build go1.17
// +build go1.17

// CODE GENERATED AUTOMATICALLY WITH github.com/probably-not/fastcopy/cmd/gen.go

package uintptr

// `isOptimized` is a constant used to ensure that the build constraint works appropriately.
// It's not that I don't trust the build constraint... but... I don't trust them...
const isOptimized = true

func CopyUintptrSlice(dst, src []uintptr) {
	// If len(dst) is less than len(src), then we need to copy with the size equal to len(dst)
	// in order to not panic by getting an array that is bigger than len(dst)
	if len(dst) < len(src) {
		switch len(dst) {
		
		case 0:
			copyUintptrSlice0(dst, src)
			return
		
		case 1:
			copyUintptrSlice1(dst, src)
			return
		
		case 2:
			copyUintptrSlice2(dst, src)
			return
		
		case 3:
			copyUintptrSlice3(dst, src)
			return
		
		case 4:
			copyUintptrSlice4(dst, src)
			return
		
		case 5:
			copyUintptrSlice5(dst, src)
			return
		
		default:
			// If len(dst) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
			copy(dst, src)
			return
		}
	}

	// If len(src) is within our limits and greater than len(dst), then we need to copy with the
	// size equal to len(src) in order to not panic by getting an array that is bigger than len(src)
	switch len(src) {
	
	case 0:
		copyUintptrSlice0(dst, src)
		return
	
	case 1:
		copyUintptrSlice1(dst, src)
		return
	
	case 2:
		copyUintptrSlice2(dst, src)
		return
	
	case 3:
		copyUintptrSlice3(dst, src)
		return
	
	case 4:
		copyUintptrSlice4(dst, src)
		return
	
	case 5:
		copyUintptrSlice5(dst, src)
		return
	
	default:
		// If len(dst) is greater than the maximum that we have generated for, then we utilize the built-in copy function.
		copy(dst, src)
		return
	}
}

func copyUintptrSlice0(dst, src []uintptr) {
	*(*[0]uintptr)(dst) = *(*[0]uintptr)(src)
}

func copyUintptrSlice1(dst, src []uintptr) {
	*(*[1]uintptr)(dst) = *(*[1]uintptr)(src)
}

func copyUintptrSlice2(dst, src []uintptr) {
	*(*[2]uintptr)(dst) = *(*[2]uintptr)(src)
}

func copyUintptrSlice3(dst, src []uintptr) {
	*(*[3]uintptr)(dst) = *(*[3]uintptr)(src)
}

func copyUintptrSlice4(dst, src []uintptr) {
	*(*[4]uintptr)(dst) = *(*[4]uintptr)(src)
}

func copyUintptrSlice5(dst, src []uintptr) {
	*(*[5]uintptr)(dst) = *(*[5]uintptr)(src)
}
