package utils

import "syscall"

func StringUTF16(s string) *uint16 {
	result, _ := syscall.UTF16PtrFromString(s)

	return result
}
