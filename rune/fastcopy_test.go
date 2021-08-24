package rune

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/probably-not/fastcopy/util"
)

var testCases = []struct {
	desc string
	size int
}{
	{
		desc: "0",
		size: 0,
	},
	{
		desc: "10",
		size: 10,
	},
	{
		desc: "100",
		size: 100,
	},
	{
		desc: "1000",
		size: 1000,
	},
	{
		desc: "5000",
		size: 5000,
	},
	{
		desc: "8192",
		size: 8192,
	},
	{
		desc: "10000",
		size: 10000,
	},
}

func Test_FastCopyRune_Simple(t *testing.T) {
	t.Parallel()

	// Ensure that when isOptimized is true we are above Go1.17
	// Like I said... it's not that I don't trust build constraints, but, I don't trust build constraints.
	major, minor, patch := util.ParseGoVersion()
	if isOptimized && !(major >= 1 && minor >= 17) {
		t.Fatalf("isOptimized is true and go version is not above 1.17. Expected Go>=1.17, Got: Go=%d.%d.%d", major, minor, patch)
		t.FailNow()
	}
	if !isOptimized && (major >= 1 && minor >= 17) {
		t.Fatalf("isOptimized is false and go version is above 1.17. Expected Go>=1.17, Got: Go=%d.%d.%d", major, minor, patch)
		t.FailNow()
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(sub *testing.T) {
			sub.Parallel()
			original := make([]rune, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			cp := make([]rune, len(original))
			CopyRuneSlice(cp, original)

			if !reflect.DeepEqual(original, cp) {
				t.Errorf("Original and Copy are not Equal. Expected=%v, Got=%v", original, cp)
			}
		})
	}
}

var benchCases = []struct {
	desc string
	size int
}{
	{desc: "0", size: 0},
	{desc: fmt.Sprint(1 << 1), size: 1 << 1},
	{desc: fmt.Sprint(1 << 2), size: 1 << 2},
	{desc: fmt.Sprint(1 << 3), size: 1 << 3},
	{desc: fmt.Sprint(1 << 4), size: 1 << 4},
	{desc: fmt.Sprint(1 << 5), size: 1 << 5},
	{desc: fmt.Sprint(1 << 6), size: 1 << 6},
	{desc: fmt.Sprint(1 << 7), size: 1 << 7},
	{desc: fmt.Sprint(1 << 8), size: 1 << 8},
	{desc: fmt.Sprint(1 << 9), size: 1 << 9},
	{desc: fmt.Sprint(1 << 10), size: 1 << 10},
	{desc: fmt.Sprint(1 << 11), size: 1 << 11},
	{desc: fmt.Sprint(1 << 12), size: 1 << 12},
	{desc: fmt.Sprint(1 << 13), size: 1 << 13},
	{desc: fmt.Sprint(1 << 14), size: 1 << 14},
	{desc: fmt.Sprint(1 << 15), size: 1 << 15},
	{desc: fmt.Sprint(1 << 16), size: 1 << 16},
}

func Benchmark_FastCopyRune_Simple(b *testing.B) {
	for _, tC := range benchCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]rune, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			cp := make([]rune, len(original))

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				CopyRuneSlice(cp, original)
			}
		})
	}
}

func Benchmark_FastCopyRune_Complex(b *testing.B) {
	for _, tC := range benchCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]rune, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			var cps [128][]rune
			for i := range cps {
				cps[i] = make([]rune, tC.size)
			}

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				CopyRuneSlice(cps[i&127], original)
			}
		})
	}
}

func Benchmark_CopyRune_Simple(b *testing.B) {
	for _, tC := range benchCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]rune, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			cp := make([]rune, len(original))

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				copy(cp, original)
			}
		})
	}
}

func Benchmark_CopyRune_Complex(b *testing.B) {
	for _, tC := range benchCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]rune, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			var cps [128][]rune
			for i := range cps {
				cps[i] = make([]rune, tC.size)
			}

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				copy(cps[i&127], original)
			}
		})
	}
}
