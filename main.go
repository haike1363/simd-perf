package main

import (
    "github.com/intel-go/cpuid"
    "fmt"
    "C"
)

/*
 #cgo CFLAGS: -mavx
 #include "simd/gorilla_avx.c"
*/
func main() {
    fmt.Println("gorilla_avx version", C.gorilla_avx_version())

    fmt.Printf("VendorString:   %s\n", cpuid.VendorIdentificatorString)

    fmt.Printf("Features: ")
    for i := uint64(0); i < 64; i++ {
        if cpuid.HasFeature(1 << i) {
            fmt.Printf("%s ", cpuid.FeatureNames[1<<i])
        }
    }
    fmt.Printf("\n")

    fmt.Printf("ExtendedFeatures: ")
    for i := uint64(0); i < 64; i++ {
        if cpuid.HasExtendedFeature(1 << i) {
            fmt.Printf("%s ", cpuid.ExtendedFeatureNames[1<<i])
        }
    }
    fmt.Printf("\n")

    fmt.Printf("ExtraFeatures: ")
    for i := uint64(0); i < 64; i++ {
        if cpuid.HasExtraFeature(1 << i) {
            fmt.Printf("%s ", cpuid.ExtraFeatureNames[1<<i])
        }
    }
    fmt.Printf("\n")
}