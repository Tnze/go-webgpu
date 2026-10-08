package spec

import (
	"strings"
	"unsafe"
)

func (s *Spec) SizeOf(typ string) (size int) {
	switch typ {
	case "uint8", "int8":
		size = 1
	case "uint16", "int16":
		size = 2
	case "uint32", "int32", "float":
		size = 4
	case "uint64", "int64", "double":
		size = 8
	case "usize":
		size = int(unsafe.Sizeof(uintptr(0)))
	case "c_void",
		"c_void_data_ptr",
		"c_void_mapped_range_ptr",
		"c_void_a_native_window",
		"c_void_ca_metal_layer",
		"c_void_h_instance",
		"c_void_h_wnd",
		"c_void_wl_display",
		"c_void_wl_surface",
		"c_void_x11_display",
		"c_void_xcb_connection":
		size = int(unsafe.Sizeof(unsafe.Pointer(nil)))
	default:
		if strings.HasPrefix(typ, "bitflag.") {
			size = 8
		} else if strings.HasPrefix(typ, "enum.") {
			size = 4
		} else if strings.HasPrefix(typ, "struct.") {
			biggestField := 0
			for _, v := range s.GetDef(typ).(*Struct).Members {
				fieldSize := s.SizeOf(v.Type)
				if fieldSize > biggestField {
					biggestField = fieldSize
				}
				if padding := size % fieldSize; padding != 0 {
					size += padding
				}
				size += fieldSize
			}
			if padding := size % biggestField; padding != 0 {
				size += padding
			}
		} else {
			panic("Unknown type: " + typ)
		}
	}
	return
}
