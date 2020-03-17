Fuzz.go contains to methods that can be fuzzed using go-fuzz, https://github.com/dvyukov/go-fuzz.

(1) Build and Fuzz ReadFrom method

```
go-fuzz-build -func=FuzzReadFrom
go-fuzz -bin=GoFuzz-fuzz.zip -workdir=workdir -procs=1
```

(2) Build and Fuzz UnmarshalBinary method

```
go-fuzz-build -func=FuzzUnmarshalBinary
go-fuzz -bin=GoFuzz-fuzz.zip -workdir=workdir -procs=1
```