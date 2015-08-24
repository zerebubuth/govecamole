package govecamole

/*
#cgo pkg-config: vecamole
#include <vecamole/create.h>
extern int goWriterFunc(void *, char *, int);
*/
import "C"

import (
	"errors"
	"io"
	"unsafe"
)

//export goWriterFunc
func goWriterFunc(ptr unsafe.Pointer, bytes *C.char, len C.int) C.int {
	var writer *io.Writer = (*io.Writer)(ptr);
	var b []byte = C.GoBytes(unsafe.Pointer(bytes), len)

	n, err := (*writer).Write(b)

	if err != nil {
		return -1
	}
	return C.int(n)
}

func Render(buf io.Writer, m *VecMap, req *Request, opts *Options) error {
	var ptr unsafe.Pointer = unsafe.Pointer(&buf)
	var writerFuncPtr = C.vecamole_tile_writer_func_t(C.goWriterFunc)
	length := C.vecamole_create_tile(ptr, writerFuncPtr, m.ptr, req.ptr, opts.ptr)
	if length < 0 {
		return errors.New("Failed to create a tile.")
	}
	return nil
}
