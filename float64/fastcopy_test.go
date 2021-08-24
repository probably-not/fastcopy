package float64

import (
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

func Test_FastCopyFloat64_Simple(t *testing.T) {
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
			original := make([]float64, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			cp := make([]float64, len(original))
			CopyFloat64Slice(cp, original)

			if !reflect.DeepEqual(original, cp) {
				t.Errorf("Original and Copy are not Equal. Expected=%v, Got=%v", original, cp)
			}
		})
	}
}

func Benchmark_FastCopyFloat64_Simple(b *testing.B) {
	for _, tC := range testCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]float64, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			cp := make([]float64, len(original))

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				CopyFloat64Slice(cp, original)
			}
		})
	}
}

func Benchmark_FastCopyFloat64_Complex(b *testing.B) {
	for _, tC := range testCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]float64, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			var cps [128][]float64
			for i := range cps {
				cps[i] = make([]float64, tC.size)
			}

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				CopyFloat64Slice(cps[i&127], original)
			}
		})
	}
}

func Benchmark_CopyFloat64_Simple(b *testing.B) {
	for _, tC := range testCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]float64, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			cp := make([]float64, len(original))

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				copy(cp, original)
			}
		})
	}
}

func Benchmark_CopyFloat64_Complex(b *testing.B) {
	for _, tC := range testCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]float64, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			var cps [128][]float64
			for i := range cps {
				cps[i] = make([]float64, tC.size)
			}

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				copy(cps[i&127], original)
			}
		})
	}
}
