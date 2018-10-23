package gorilla_avx

/*
 #cgo CFLAGS: -mavx
 #include "c/gorilla_avx.c"
*/
import "C"

func Version() (int32) {
    return (int32)(C.gorilla_avx_version())
}
