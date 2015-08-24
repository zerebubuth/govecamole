package vecmap

/*
#cgo pkg-config: vecamole
#include <vecamole/create.h>
*/
import "C"

import (
	"errors"
)

type Options struct {
	ptr *C.vecamole_create_options_t
}

func DefaultOptions() (*Options, error) {
	ptr := C.vecamole_create_options_default()
	if ptr == nil {
		return nil, errors.New("Failed to allocate a new vecamole options object.")
	}
	return &Options{ptr: ptr}, nil
}

func (self *Options) Close() {
	var status = C.vecamole_create_options_delete(self.ptr)
	if status != 0 {
		panic("Unable to deallocate vecamole options object. This might indicate memory corruption.")
	}

}
