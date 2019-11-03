package readahead

import "syscall"

// Readahead performs readahead on a file.
func Readahead(fd int, offset int, size int) error {
	var err error
	_, _, errno := syscall.Syscall6(
		syscall.SYS_READAHEAD,
		uintptr(fd),
		uintptr(offset),
		uintptr(size),
		uintptr(0),
		uintptr(0),
		uintptr(0),
	)

	if errno != 0 {
		err = errno
	}

	return err
}
