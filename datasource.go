package govecamole

/*
#cgo pkg-config: vecamole
#include <vecamole/datasource.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

func RegisterDefaultDatasources() error {
	status := C.vecamole_register_default_datasources()
	if status != 0 {
		return errors.New("Unable to register default data sources.")
	}
	return nil
}

func RegisterDatasources(dir string, recurse bool) error {
	var cstring = C.CString(dir)
	defer C.free(unsafe.Pointer(cstring))

	var crecurse C.int = 0
	if recurse {
		crecurse = 1
	}

	status := C.vecamole_register_datasources(cstring, crecurse)
	if status != 0 {
		return fmt.Errorf("Unable to register data sources at %v.", dir)
	}

	return nil
}
