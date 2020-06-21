package g

/*
#cgo pkg-config: glib-2.0 gobject-2.0
#include <glib.h>
#include <glib-object.h>
#include <stdlib.h>
*/
import "C"
import gi "github.com/electricface/go-gir3/gi-lite"

const (
	TYPE_INVALID   gi.GType = C.G_TYPE_INVALID
	TYPE_NONE      gi.GType = C.G_TYPE_NONE
	TYPE_INTERFACE gi.GType = C.G_TYPE_INTERFACE
	TYPE_CHAR      gi.GType = C.G_TYPE_CHAR
	TYPE_UCHAR     gi.GType = C.G_TYPE_UCHAR
	TYPE_BOOLEAN   gi.GType = C.G_TYPE_BOOLEAN
	TYPE_INT       gi.GType = C.G_TYPE_INT   // int32
	TYPE_UINT      gi.GType = C.G_TYPE_UINT  // uint32
	TYPE_LONG      gi.GType = C.G_TYPE_LONG  // int64
	TYPE_ULONG     gi.GType = C.G_TYPE_ULONG // uint64
	TYPE_INT64     gi.GType = C.G_TYPE_INT64
	TYPE_UINT64    gi.GType = C.G_TYPE_UINT64
	TYPE_ENUM      gi.GType = C.G_TYPE_ENUM
	TYPE_FLAGS     gi.GType = C.G_TYPE_FLAGS
	TYPE_FLOAT     gi.GType = C.G_TYPE_FLOAT
	TYPE_DOUBLE    gi.GType = C.G_TYPE_DOUBLE
	TYPE_STRING    gi.GType = C.G_TYPE_STRING
	TYPE_POINTER   gi.GType = C.G_TYPE_POINTER
	TYPE_BOXED     gi.GType = C.G_TYPE_BOXED
	TYPE_PARAM     gi.GType = C.G_TYPE_PARAM
	TYPE_OBJECT    gi.GType = C.G_TYPE_OBJECT
	TYPE_VARIANT   gi.GType = C.G_TYPE_VARIANT
)

// Name is a wrapper around g_type_name().
//func (t Type) Name() string {
//	return C.GoString((*C.char)(C.g_type_name(C.GType(t))))
//}
//
//func (t Type) String() string {
//	return t.Name()
//}
//
//// Depth is a wrapper around g_type_depth().
//func (t Type) Depth() uint {
//	return uint(C.g_type_depth(C.GType(t)))
//}

// Parent is a wrapper around g_type_parent().
//func (t Type) Parent() Type {
//	return Type(C.g_type_parent(C.GType(t)))
//}
