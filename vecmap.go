package govecamole

/*
#cgo pkg-config: vecamole
#include <vecamole.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

type VecMap struct {
	ptr unsafe.Pointer
}

func New(width, height int) (*VecMap, error) {
	var ptr = C.vecamole_map_new(C.int(width), C.int(height))
	if ptr == nil {
		return nil, errors.New("Unable to allocate new vecamole map object.")
	}
	return &VecMap{ptr: ptr}, nil
}

func (self *VecMap) Close() {
	var status = C.vecamole_map_delete(self.ptr)
	if status != 0 {
		panic("Unable to deallocate vecamole map object. This might indicate memory corruption.")
	}
}

func (self *VecMap) RegisterFonts(dir string, recurse bool) error {
	var cstring = C.CString(dir)
	defer C.free(unsafe.Pointer(cstring))

	var crecurse C.int = 0
	if recurse {
		crecurse = 1
	}

	status := C.vecamole_map_register_fonts(self.ptr, cstring, crecurse)
	if status != 0 {
		return fmt.Errorf("Unable to register fonts at %v.", dir)
	}

	return nil
}

func (self *VecMap) LoadString(str string, strict bool, basename string) error {
	var cstr = C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	var cbasename = C.CString(basename)
	defer C.free(unsafe.Pointer(cbasename))

	var cstrict C.int = 0
	if strict {
		cstrict = 1
	}

	status := C.vecamole_map_load_string(self.ptr, cstr, cstrict, cbasename)
	if status != 0 {
		return fmt.Errorf("Unable to load map from string %v.", str)
	}
	return nil
}

func (self *VecMap) LoadFile(filename string, strict bool, basename string) error {
	var cfilename = C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	var cbasename = C.CString(basename)
	defer C.free(unsafe.Pointer(cbasename))

	var cstrict C.int = 0
	if strict {
		cstrict = 1
	}

	status := C.vecamole_map_load_file(self.ptr, cfilename, cstrict, cbasename)
	if status != 0 {
		return fmt.Errorf("Unable to load map from file %v.", filename)
	}
	return nil
}
