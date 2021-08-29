# fastcopy

Just because I can...

There was a tweet recently (link here: https://twitter.com/go100and1/status/1429800152278081543) which mentioned a new faster way to copy ever since go 1.17. Me being the troll that I am, I figured I could create a simple package to make this easy for people to use.

You probably shouldn't use it... honestly it's probably a huge mistake. But, I felt like doing it, just to check it out.

## Uh-Ohs

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