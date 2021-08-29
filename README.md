# fastcopy

Just because I can...

There was a tweet recently (link here: https://twitter.com/go100and1/status/1429800152278081543) which mentioned a new faster way to copy ever since go 1.17. Me being the troll that I am, I figured I could create a simple package to make this easy for people to use.

You probably shouldn't use it... honestly it's probably a huge mistake. But, I felt like doing it, just to check it out.

## Strategies

I have a couple of strategies that I'd like to benchmark. The first is utilizing a gigantic switch statement. My thoughts are that a switch would be faster because the compiler could possibly compile it into a nice little jump table, especially since the switch cases are all integers (since we are switching on the length).

### Uh-Ohs

Well, it seems that this will never work past 3072 cases, due to the way the compiler currently handles the size of the switch cases in the function. For the interested, the following is the output when attempting to generate the test binary using just the fastcopy/byte package:

```shell
# github.com/probably-not/fastcopy/byte [github.com/probably-not/fastcopy/byte.test]
byte/fastcopy_test.go:101:19: prepwrite: bad off=1073741825 siz=8 s=
byte/fastcopy_test.go:101:19: prepwrite: bad off=1073741833 siz=7 s=
byte/fastcopy_test.go:101:19: prepwrite: bad off=1073741840 siz=8 s=
byte/fastcopy_test.go:101:19: prepwrite: bad off=1073741848 siz=8 s=
byte/fastcopy_test.go:101:19: prepwrite: bad off=1073741856 siz=7 s=
byte/fastcopy_test.go:101:19: prepwrite: bad off=1073741863 siz=8 s=
byte/fastcopy_test.go:101:19: prepwrite: bad off=1073741871 siz=8 s=
byte/fastcopy_test.go:101:19: prepwrite: bad off=1073741879 siz=7 s=
byte/fastcopy_test.go:101:19: prepwrite: bad off=1073741886 siz=8 s=
byte/fastcopy_test.go:101:19: prepwrite: bad off=1073741894 siz=8 s=
byte/fastcopy_test.go:101:19: too many errors
```

My guess? The function is simply too large and the compiler is attempting too much.

### Using Slice Indexes

Welp, the switch case didn't work... so, here's my next attempt. A much simpler function, but worse on memory. Utilizing a Slice pointing to each of the functions based on the length as a slice index. Murder on the memory usage, since we need to hold a gigantic array in memory, but you know what? It's worth it so that we can compile more that 3072 cases.

#### Performance

All of the benchmarks were taken with the following specs:
```
goos: darwin
goarch: amd64
pkg: github.com/probably-not/fastcopy/byte
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
```

The first function is very simple:
```go
func Benchmark_FastCopyByte_Simple(b *testing.B) {
	for _, tC := range benchCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]byte, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			cp := make([]byte, len(original))

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				CopyByteSlice(cp, original)
			}
		})
	}
}
```
As compared to:
```go
func Benchmark_CopyByte_Simple(b *testing.B) {
	for _, tC := range benchCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]byte, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			cp := make([]byte, len(original))

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				copy(cp, original)
			}
		})
	}
}
```

Results:

```
Benchmark_FastCopyByte_Simple/0-16      331438114                3.352 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/0-16                  1000000000               0.2675 ns/op          0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/2-16      346514058                3.514 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/2-16                  635587597                1.877 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/4-16      338763511                3.853 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/4-16                  609830407                1.866 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/8-16      329836778                3.958 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/8-16                  548220964                2.207 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/16-16     268806114                3.768 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/16-16                 455642914                2.814 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/32-16     253866597                4.153 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/32-16                 475931059                2.755 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/64-16     317877523                3.690 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/64-16                 318210625                7.703 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/128-16    226545384                5.689 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/128-16                339175380                3.442 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/256-16    154479928                7.464 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/256-16                207067909                5.592 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/512-16    100000000               11.14 ns/op            0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/512-16                140265024                8.912 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/1024-16           60288210                19.33 ns/op            0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/1024-16               86477966                12.59 ns/op            0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/2048-16           49276530                26.10 ns/op            0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/2048-16               53489262                19.45 ns/op            0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/4096-16           26011794                49.61 ns/op            0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/4096-16               30486660                36.77 ns/op            0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/8192-16           13099939                83.61 ns/op            0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/8192-16               15829779                73.24 ns/op            0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/16384-16           7597996               158.2 ns/op             0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/16384-16               5435125               234.9 ns/op             0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/32768-16           1997217               724.7 ns/op             0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/32768-16               1799995               668.6 ns/op             0 B/op          0 allocs/op
Benchmark_FastCopyByte_Simple/65536-16            849024              1611 ns/op               0 B/op          0 allocs/op
Benchmark_CopyByte_Simple/65536-16                794526              1654 ns/op               0 B/op          0 allocs/op
```

The next function is a bit trickier, but might mimic the copying a little better (according to some random googling I've found):
```go
func Benchmark_FastCopyByte_Complex(b *testing.B) {
	for _, tC := range benchCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]byte, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			var cps [128][]byte
			for i := range cps {
				cps[i] = make([]byte, tC.size)
			}

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				CopyByteSlice(cps[i&127], original)
			}
		})
	}
}
```
As compared to:
```go
func Benchmark_CopyByte_Complex(b *testing.B) {
	for _, tC := range benchCases {
		b.Run(tC.desc, func(sub *testing.B) {
			original := make([]byte, tC.size)
			for i := 0; i < len(original); i++ {
				if rand.Intn(2) > 0 {
					original[i] = 1
				}
			}

			var cps [128][]byte
			for i := range cps {
				cps[i] = make([]byte, tC.size)
			}

			sub.ResetTimer()
			sub.ReportAllocs()
			for i := 0; i < sub.N; i++ {
				copy(cps[i&127], original)
			}
		})
	}
}
```

Results:

```
Benchmark_FastCopyByte_Complex/0-16             321482926                3.360 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/0-16                 1000000000               0.6879 ns/op          0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/2-16             331151377                3.845 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/2-16                 588958662                2.041 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/4-16             328108588                3.982 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/4-16                 540487888                2.190 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/8-16             317882390                3.435 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/8-16                 499007793                2.335 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/16-16            352054576                3.994 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/16-16                451692368                2.755 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/32-16            331322299                3.671 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/32-16                360414014                2.827 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/64-16            307835654                3.922 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/64-16                376982901                3.327 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/128-16           193465417                5.842 ns/op           0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/128-16               290560802                3.943 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/256-16           120189603               10.57 ns/op            0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/256-16               164566411                7.530 ns/op           0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/512-16           72027721                18.62 ns/op            0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/512-16               76892673                14.19 ns/op            0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/1024-16          42415664                30.99 ns/op            0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/1024-16              43460444                24.12 ns/op            0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/2048-16          22409776                54.43 ns/op            0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/2048-16              20837419                51.28 ns/op            0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/4096-16          11989815               108.1 ns/op             0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/4096-16              10347876               114.1 ns/op             0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/8192-16           5310859               236.8 ns/op             0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/8192-16               5288917               220.6 ns/op             0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/16384-16          3073270               388.3 ns/op             0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/16384-16              2466318               469.6 ns/op             0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/32768-16          1302118               996.1 ns/op             0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/32768-16              1000000              1029 ns/op               0 B/op          0 allocs/op
Benchmark_FastCopyByte_Complex/65536-16           512608              2271 ns/op               0 B/op          0 allocs/op
Benchmark_CopyByte_Complex/65536-16               442358              2319 ns/op               0 B/op          0 allocs/op
```