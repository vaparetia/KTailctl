package main

// #include <stdlib.h>
// #include "logging.h"
// #cgo LDFLAGS: -L/home/fkoehler/Code/KTailctl/build/bin -ltailwrap_logging_wrapper -Wl,-rpath=/home/fkoehler/Code/KTailctl/build/bin
import "C"
import (
	"fmt"
	"unsafe"
)

func log_critical(message string) {
	msg := C.CString(message)
	defer C.free(unsafe.Pointer(msg))
	C.ktailctl_critical(msg)
}

func log_debug(message string) {
	msg := C.CString(message)
	defer C.free(unsafe.Pointer(msg))
	C.ktailctl_debug(msg)
}

func log_info(message string) {
	msg := C.CString(message)
	defer C.free(unsafe.Pointer(msg))
	C.ktailctl_info(msg)
}

func log_warning(message string) {
	msg := C.CString(message)
	defer C.free(unsafe.Pointer(msg))
	C.ktailctl_warning(msg)
}

func log_critical_error(err error) {
	msg := C.CString(fmt.Sprintln(err))
	defer C.free(unsafe.Pointer(msg))
	C.ktailctl_critical(msg)
}
