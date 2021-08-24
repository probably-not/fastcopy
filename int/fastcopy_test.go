package int

import (
	"math/rand"
	"reflect"
	"testing"
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

func Test_FastCopyInt_Simple(t *testing.T) {
	t.Parallel()
	for _, tC := range testCases {
		t.Run(tC.desc, func(sub *testing.T) {
			sub.Parallel()
			original := make([]int, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			cp := make([]int, len(original))
			CopyIntSlice(cp, original)

			if !reflect.DeepEqual(original, cp) {
				t.Errorf("Original and Copy are not Equal. Expected=%v, Got=%v", original, cp)
			}
		})
	}
}

func Benchmark_FastCopyInt_Simple(b *testing.B) {
	for _, tC := range testCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]int, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			cp := make([]int, len(original))

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				CopyIntSlice(cp, original)
			}
		})
	}
}

func Benchmark_FastCopyInt_Complex(b *testing.B) {
	for _, tC := range testCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]int, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			var cps [128][]int
			for i := range cps {
				cps[i] = make([]int, tC.size)
			}

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				CopyIntSlice(cps[i&127], original)
			}
		})
	}
}