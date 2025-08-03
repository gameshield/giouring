// MIT License
//
// Copyright (c) 2023 Paweł Gaczyński
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
// SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package giouring

import (
	"math"
	"syscall"
	"unsafe"
)

func sysMmap(addr, length uintptr, prot, flags, fd int, offset int64) (unsafe.Pointer, error) {
	r0, _, errno := syscall.Syscall6(
		syscall.SYS_MMAP,
		addr,
		length,
		uintptr(prot),
		uintptr(flags),
		uintptr(fd),
		uintptr(offset),
	)
	if errno != 0 {
		return nil, errno
	}
	return unsafe.Pointer(r0), nil
}

func sysMunmap(addr, length uintptr) error {
	_, _, errno := syscall.Syscall(
		syscall.SYS_MUNMAP,
		addr,
		length,
		0,
	)
	if errno != 0 {
		return errno
	}
	return nil
}

func sysMadvise(addr, length, advice uintptr) error {
	_, _, errno := syscall.Syscall(
		syscall.SYS_MADVISE,
		addr,
		length,
		advice,
	)
	if errno != 0 {
		return errno
	}
	return nil
}

const liburingUdataTimeout uint64 = math.MaxUint64
