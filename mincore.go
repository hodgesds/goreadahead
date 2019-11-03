package readahead

import (
	"syscall"
	"unsafe"
)

// Mincore performs readahead on a file.
func Mincore(addr uintptr, length int) ([]byte, error) {
	// The vec argument must point to an array containing at least
	// (length+PAGE_SIZE-1) / PAGE_SIZE bytes.
	pageSize := syscall.Getpagesize()
	vec := make([]byte, (length+pageSize-1)/pageSize)

	var err error
	_, _, errno := syscall.Syscall6(
		syscall.SYS_MINCORE,
		addr,
		uintptr(length),
		uintptr(unsafe.Pointer(&vec[0])),
		uintptr(0),
		uintptr(0),
		uintptr(0),
	)

	if errno != 0 {
		err = errno
		return nil, err
	}

	return vec, nil
}
