package main

import "C"
import (
    "fmt"
    "github.com/haike1363/simd-perf/simd/gorilla_avx"
    "github.com/intel-go/cpuid"
)

func main() {
    fmt.Println("avx supported: ", cpuid.HasFeature(cpuid.AVX))
    fmt.Println("avx enabled: ", cpuid.EnabledAVX)
    fmt.Println("gorilla_avx go version", gorilla_avx.Version())
}
