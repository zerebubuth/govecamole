package govecamole

/*
#cgo pkg-config: vecamole
#include <vecamole/request.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

type Request struct {
	ptr unsafe.Pointer
}

func NewRequestZXY(w, h, z, x, y int) (*Request, error) {
	ptr := C.vecamole_request_zxy(C.int(w), C.int(h), C.int(z), C.int(x), C.int(y))
	if ptr == nil {
		return nil, errors.New("Failed to allocate a new vecamole request object.")
	}
	return &Request{ptr: ptr}, nil
}

func (self *Request) Close() {
	var status = C.vecamole_request_delete(self.ptr)
	if status != 0 {
		panic("Unable to deallocate vecamole request object. This might indicate memory corruption.")
	}
	self.ptr = nil
}

func (self *Request) SetBufferSize(size int) {
	C.vecamole_request_set_buffer_size(self.ptr, C.int(size))
}
