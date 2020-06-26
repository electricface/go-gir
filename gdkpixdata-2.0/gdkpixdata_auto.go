package gdkpixdata

/*
#cgo pkg-config: gdk-pixbuf-2.0
#include <gdk-pixbuf/gdk-pixdata.h>
*/
import "C"
import "github.com/electricface/go-gir/g-2.0"
import "github.com/electricface/go-gir/gdkpixbuf-2.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("GdkPixdata")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GdkPixdata", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Struct Pixdata
type Pixdata struct {
	P unsafe.Pointer
}

const SizeOfStructPixdata = 32

func PixdataGetType() gi.GType {
	ret := _I.GetGType(0, "Pixdata")
	return ret
}

// Deprecated
//
// gdk_pixdata_deserialize
//
// [ stream_length ] trans: nothing
//
// [ stream ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pixdata) Deserialize(stream_length uint32, stream gi.Uint8Array) (result bool, err error) {
	iv, err := _I.Get(0, "Pixdata", "deserialize", 2, 0, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_stream_length := gi.NewUint32Argument(stream_length)
	arg_stream := gi.NewPointerArgument(stream.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_stream_length, arg_stream, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// Deprecated
//
// gdk_pixdata_serialize
//
// [ stream_length_p ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v Pixdata) Serialize() (result gi.Uint8Array) {
	iv, err := _I.Get(1, "Pixdata", "serialize", 2, 1, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_stream_length_p := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_stream_length_p}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var stream_length_p uint32
	_ = stream_length_p
	stream_length_p = outArgs[0].Uint32()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(stream_length_p)}
	return
}

// Deprecated
//
// gdk_pixdata_to_csource
//
// [ name ] trans: nothing
//
// [ dump_type ] trans: nothing
//
// [ result ] trans: everything
//
func (v Pixdata) ToCsource(name string, dump_type PixdataDumpTypeFlags) (result g.String) {
	iv, err := _I.Get(2, "Pixdata", "to_csource", 2, 2, gi.INFO_TYPE_STRUCT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	arg_dump_type := gi.NewIntArgument(int(dump_type))
	args := []gi.Argument{arg_v, arg_name, arg_dump_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// Flags PixdataDumpType
type PixdataDumpTypeFlags int

const (
	PixdataDumpTypePixdataStream PixdataDumpTypeFlags = 0
	PixdataDumpTypePixdataStruct PixdataDumpTypeFlags = 1
	PixdataDumpTypeMacros        PixdataDumpTypeFlags = 2
	PixdataDumpTypeGtypes        PixdataDumpTypeFlags = 0
	PixdataDumpTypeCtypes        PixdataDumpTypeFlags = 256
	PixdataDumpTypeStatic        PixdataDumpTypeFlags = 512
	PixdataDumpTypeConst         PixdataDumpTypeFlags = 1024
	PixdataDumpTypeRleDecoder    PixdataDumpTypeFlags = 65536
)

func PixdataDumpTypeGetType() gi.GType {
	ret := _I.GetGType(1, "PixdataDumpType")
	return ret
}

// Flags PixdataType
type PixdataTypeFlags int

const (
	PixdataTypeColorTypeRgb    PixdataTypeFlags = 1
	PixdataTypeColorTypeRgba   PixdataTypeFlags = 2
	PixdataTypeColorTypeMask   PixdataTypeFlags = 255
	PixdataTypeSampleWidth8    PixdataTypeFlags = 65536
	PixdataTypeSampleWidthMask PixdataTypeFlags = 983040
	PixdataTypeEncodingRaw     PixdataTypeFlags = 16777216
	PixdataTypeEncodingRle     PixdataTypeFlags = 33554432
	PixdataTypeEncodingMask    PixdataTypeFlags = 251658240
)

func PixdataTypeGetType() gi.GType {
	ret := _I.GetGType(2, "PixdataType")
	return ret
}

// Deprecated
//
// gdk_pixbuf_from_pixdata
//
// [ pixdata ] trans: nothing
//
// [ copy_pixels ] trans: nothing
//
// [ result ] trans: everything
//
func PixbufFromPixdata(pixdata Pixdata, copy_pixels bool) (result gdkpixbuf.Pixbuf, err error) {
	iv, err := _I.Get(3, "pixbuf_from_pixdata", "", 5, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_pixdata := gi.NewPointerArgument(pixdata.P)
	arg_copy_pixels := gi.NewBoolArgument(copy_pixels)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_pixdata, arg_copy_pixels, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// constants
const (
	PIXBUF_MAGIC_NUMBER   = 1197763408
	PIXDATA_HEADER_LENGTH = 24
)
