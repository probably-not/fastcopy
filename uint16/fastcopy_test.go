package uint16

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

func Test_FastCopyUint16(t *testing.T) {
	t.Parallel()
	for _, tC := range testCases {
		t.Run(tC.desc, func(sub *testing.T) {
			sub.Parallel()
			original := make([]uint16, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			cp := make([]uint16, len(original))
			CopyUint16Slice(cp, original)

			if !reflect.DeepEqual(original, cp) {
				t.Errorf("Original and Copy are not Equal. Expected=%v, Got=%v", original, cp)
			}
		})
	}
}

func Benchmark_FastCopyUint16_Simple(b *testing.B) {
	for _, tC := range testCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]uint16, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			cp := make([]uint16, len(original))

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				CopyUint16Slice(cp, original)
			}
		})
	}
}

func Benchmark_FastCopyUint16_Complex(b *testing.B) {
	for _, tC := range testCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]uint16, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			var cps [128][]uint16
			for i := range cps {
				cps[i] = make([]uint16, tC.size)
			}

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				CopyUint16Slice(cps[i&127], original)
			}
		})
	}
}
