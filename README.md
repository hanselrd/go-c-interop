# go-c-interop

The purpose of this repo is to show how Go and C can be used together very easily. This example can be built using the included Makefile:

```
$ make
```

or using the following command

```
$ go build
```

## Explanation

For this example I created a very simple math library that only exposes two functions: `math_add` and `math_sub`. I then call these two functions from Go using the cgo "C" module to bridge the gap between the two languages.
