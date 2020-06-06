package gdkpixbuf

import "github.com/electricface/go-gir/gio-2.0"
import "github.com/electricface/go-gir/glib-2.0"
import "github.com/electricface/go-gir/gobject-2.0"
import "github.com/electricface/go-gir3/gi"
import "log"
import "unsafe"

var _I = gi.NewInvokerCache("GdkPixbuf")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GdkPixbuf", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Enum Colorspace
type ColorspaceEnum int

const (
	ColorspaceRgb ColorspaceEnum = 0
)

// Enum InterpType
type InterpTypeEnum int

const (
	InterpTypeNearest  InterpTypeEnum = 0
	InterpTypeTiles    InterpTypeEnum = 1
	InterpTypeBilinear InterpTypeEnum = 2
	InterpTypeHyper    InterpTypeEnum = 3
)

// Object Pixbuf
type Pixbuf struct {
	gio.IconIfc
	gio.LoadableIconIfc
	gobject.Object
}

func WrapPixbuf(p unsafe.Pointer) (r Pixbuf) { r.P = p; return }

type IPixbuf interface{ P_Pixbuf() unsafe.Pointer }

func (v Pixbuf) P_Pixbuf() unsafe.Pointer { return v.P }

// gdk_pixbuf_new
// container is not nil, container is Pixbuf
// is constructor
func NewPixbuf(colorspace ColorspaceEnum, has_alpha bool, bits_per_sample int32, width int32, height int32) (result Pixbuf) {
	iv, err := _I.Get(0, "Pixbuf", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_colorspace := gi.NewIntArgument(int(colorspace))
	arg_has_alpha := gi.NewBoolArgument(has_alpha)
	arg_bits_per_sample := gi.NewInt32Argument(bits_per_sample)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	args := []gi.Argument{arg_colorspace, arg_has_alpha, arg_bits_per_sample, arg_width, arg_height}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_new_from_bytes
// container is not nil, container is Pixbuf
// is constructor
func NewPixbufFromBytes(data glib.Bytes, colorspace ColorspaceEnum, has_alpha bool, bits_per_sample int32, width int32, height int32, rowstride int32) (result Pixbuf) {
	iv, err := _I.Get(1, "Pixbuf", "new_from_bytes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data.P)
	arg_colorspace := gi.NewIntArgument(int(colorspace))
	arg_has_alpha := gi.NewBoolArgument(has_alpha)
	arg_bits_per_sample := gi.NewInt32Argument(bits_per_sample)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_rowstride := gi.NewInt32Argument(rowstride)
	args := []gi.Argument{arg_data, arg_colorspace, arg_has_alpha, arg_bits_per_sample, arg_width, arg_height, arg_rowstride}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_new_from_data
// container is not nil, container is Pixbuf
// is constructor
func NewPixbufFromData(data gi.Uint8Array, colorspace ColorspaceEnum, has_alpha bool, bits_per_sample int32, width int32, height int32, rowstride int32, destroy_fn int /*TODO_TYPE isPtr: false, tag: interface*/, destroy_fn_data unsafe.Pointer) (result Pixbuf) {
	iv, err := _I.Get(2, "Pixbuf", "new_from_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data.P)
	arg_colorspace := gi.NewIntArgument(int(colorspace))
	arg_has_alpha := gi.NewBoolArgument(has_alpha)
	arg_bits_per_sample := gi.NewInt32Argument(bits_per_sample)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_rowstride := gi.NewInt32Argument(rowstride)
	arg_destroy_fn := gi.NewIntArgument(destroy_fn) /*TODO*/
	arg_destroy_fn_data := gi.NewPointerArgument(destroy_fn_data)
	args := []gi.Argument{arg_data, arg_colorspace, arg_has_alpha, arg_bits_per_sample, arg_width, arg_height, arg_rowstride, arg_destroy_fn, arg_destroy_fn_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_new_from_file
// container is not nil, container is Pixbuf
// is constructor
func NewPixbufFromFile(filename string) (result Pixbuf, err error) {
	iv, err := _I.Get(3, "Pixbuf", "new_from_file")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_filename, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_new_from_file_at_scale
// container is not nil, container is Pixbuf
// is constructor
func NewPixbufFromFileAtScale(filename string, width int32, height int32, preserve_aspect_ratio bool) (result Pixbuf, err error) {
	iv, err := _I.Get(4, "Pixbuf", "new_from_file_at_scale")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_preserve_aspect_ratio := gi.NewBoolArgument(preserve_aspect_ratio)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_filename, arg_width, arg_height, arg_preserve_aspect_ratio, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_new_from_file_at_size
// container is not nil, container is Pixbuf
// is constructor
func NewPixbufFromFileAtSize(filename string, width int32, height int32) (result Pixbuf, err error) {
	iv, err := _I.Get(5, "Pixbuf", "new_from_file_at_size")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_filename, arg_width, arg_height, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_new_from_inline
// container is not nil, container is Pixbuf
// is constructor
// arg 1 data lenArgIdx 0
func NewPixbufFromInline(data_length int32, data gi.Uint8Array, copy_pixels bool) (result Pixbuf, err error) {
	iv, err := _I.Get(6, "Pixbuf", "new_from_inline")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_data_length := gi.NewInt32Argument(data_length)
	arg_data := gi.NewPointerArgument(data.P)
	arg_copy_pixels := gi.NewBoolArgument(copy_pixels)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_data_length, arg_data, arg_copy_pixels, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_new_from_resource
// container is not nil, container is Pixbuf
// is constructor
func NewPixbufFromResource(resource_path string) (result Pixbuf, err error) {
	iv, err := _I.Get(7, "Pixbuf", "new_from_resource")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_resource_path := gi.CString(resource_path)
	arg_resource_path := gi.NewStringArgument(c_resource_path)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_resource_path, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_resource_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_new_from_resource_at_scale
// container is not nil, container is Pixbuf
// is constructor
func NewPixbufFromResourceAtScale(resource_path string, width int32, height int32, preserve_aspect_ratio bool) (result Pixbuf, err error) {
	iv, err := _I.Get(8, "Pixbuf", "new_from_resource_at_scale")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_resource_path := gi.CString(resource_path)
	arg_resource_path := gi.NewStringArgument(c_resource_path)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_preserve_aspect_ratio := gi.NewBoolArgument(preserve_aspect_ratio)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_resource_path, arg_width, arg_height, arg_preserve_aspect_ratio, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_resource_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_new_from_stream
// container is not nil, container is Pixbuf
// is constructor
func NewPixbufFromStream(stream gio.IInputStream, cancellable gio.ICancellable) (result Pixbuf, err error) {
	iv, err := _I.Get(9, "Pixbuf", "new_from_stream")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_stream := gi.NewPointerArgument(stream.P_InputStream())
	arg_cancellable := gi.NewPointerArgument(cancellable.P_Cancellable())
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_stream, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_new_from_stream_at_scale
// container is not nil, container is Pixbuf
// is constructor
func NewPixbufFromStreamAtScale(stream gio.IInputStream, width int32, height int32, preserve_aspect_ratio bool, cancellable gio.ICancellable) (result Pixbuf, err error) {
	iv, err := _I.Get(10, "Pixbuf", "new_from_stream_at_scale")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_stream := gi.NewPointerArgument(stream.P_InputStream())
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_preserve_aspect_ratio := gi.NewBoolArgument(preserve_aspect_ratio)
	arg_cancellable := gi.NewPointerArgument(cancellable.P_Cancellable())
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_stream, arg_width, arg_height, arg_preserve_aspect_ratio, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_new_from_stream_finish
// container is not nil, container is Pixbuf
// is constructor
func NewPixbufFromStreamFinish(async_result gio.AsyncResult) (result Pixbuf, err error) {
	iv, err := _I.Get(11, "Pixbuf", "new_from_stream_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_async_result := gi.NewPointerArgument(async_result.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_async_result, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_new_from_xpm_data
// container is not nil, container is Pixbuf
// is constructor
func NewPixbufFromXpmData(data gi.CStrArray) (result Pixbuf) {
	iv, err := _I.Get(12, "Pixbuf", "new_from_xpm_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data.P)
	args := []gi.Argument{arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_calculate_rowstride
// container is not nil, container is Pixbuf
// is method
// arg0Type tag: interface, isPtr: false
func PixbufCalculateRowstride1(colorspace ColorspaceEnum, has_alpha bool, bits_per_sample int32, width int32, height int32) (result int32) {
	iv, err := _I.Get(13, "Pixbuf", "calculate_rowstride")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_colorspace := gi.NewIntArgument(int(colorspace))
	arg_has_alpha := gi.NewBoolArgument(has_alpha)
	arg_bits_per_sample := gi.NewInt32Argument(bits_per_sample)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	args := []gi.Argument{arg_colorspace, arg_has_alpha, arg_bits_per_sample, arg_width, arg_height}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gdk_pixbuf_get_file_info
// container is not nil, container is Pixbuf
// is method
// arg0Type tag: filename, isPtr: true
func PixbufGetFileInfo1(filename string) (result PixbufFormat, width int32, height int32) {
	iv, err := _I.Get(14, "Pixbuf", "get_file_info")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_filename, arg_width, arg_height}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	width = outArgs[0].Int32()
	height = outArgs[1].Int32()
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_get_file_info_async
// container is not nil, container is Pixbuf
// is method
// arg0Type tag: filename, isPtr: true
func PixbufGetFileInfoAsync1(filename string, cancellable gio.ICancellable, callback int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(15, "Pixbuf", "get_file_info_async")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_cancellable := gi.NewPointerArgument(cancellable.P_Cancellable())
	arg_callback := gi.NewIntArgument(callback) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_filename, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_filename)
}

// gdk_pixbuf_get_file_info_finish
// container is not nil, container is Pixbuf
// is method
// arg0Type tag: interface, isPtr: true
func PixbufGetFileInfoFinish1(async_result gio.AsyncResult) (result PixbufFormat, width int32, height int32, err error) {
	iv, err := _I.Get(16, "Pixbuf", "get_file_info_finish")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_async_result := gi.NewPointerArgument(async_result.P)
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_async_result, arg_width, arg_height, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	width = outArgs[0].Int32()
	height = outArgs[1].Int32()
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_new_from_stream_async
// container is not nil, container is Pixbuf
// is method
// arg0Type tag: interface, isPtr: true
func PixbufNewFromStreamAsync1(stream gio.IInputStream, cancellable gio.ICancellable, callback int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(18, "Pixbuf", "new_from_stream_async")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_stream := gi.NewPointerArgument(stream.P_InputStream())
	arg_cancellable := gi.NewPointerArgument(cancellable.P_Cancellable())
	arg_callback := gi.NewIntArgument(callback) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_stream, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// gdk_pixbuf_new_from_stream_at_scale_async
// container is not nil, container is Pixbuf
// is method
// arg0Type tag: interface, isPtr: true
func PixbufNewFromStreamAtScaleAsync1(stream gio.IInputStream, width int32, height int32, preserve_aspect_ratio bool, cancellable gio.ICancellable, callback int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(19, "Pixbuf", "new_from_stream_at_scale_async")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_stream := gi.NewPointerArgument(stream.P_InputStream())
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_preserve_aspect_ratio := gi.NewBoolArgument(preserve_aspect_ratio)
	arg_cancellable := gi.NewPointerArgument(cancellable.P_Cancellable())
	arg_callback := gi.NewIntArgument(callback) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_stream, arg_width, arg_height, arg_preserve_aspect_ratio, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// gdk_pixbuf_save_to_stream_finish
// container is not nil, container is Pixbuf
// is method
// arg0Type tag: interface, isPtr: true
func PixbufSaveToStreamFinish1(async_result gio.AsyncResult) (result bool, err error) {
	iv, err := _I.Get(20, "Pixbuf", "save_to_stream_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_async_result := gi.NewPointerArgument(async_result.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_async_result, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// gdk_pixbuf_add_alpha
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) AddAlpha(substitute_color bool, r uint8, g uint8, b uint8) (result Pixbuf) {
	iv, err := _I.Get(21, "Pixbuf", "add_alpha")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_substitute_color := gi.NewBoolArgument(substitute_color)
	arg_r := gi.NewUint8Argument(r)
	arg_g := gi.NewUint8Argument(g)
	arg_b := gi.NewUint8Argument(b)
	args := []gi.Argument{arg_v, arg_substitute_color, arg_r, arg_g, arg_b}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_apply_embedded_orientation
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) ApplyEmbeddedOrientation() (result Pixbuf) {
	iv, err := _I.Get(22, "Pixbuf", "apply_embedded_orientation")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_composite
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) Composite(dest IPixbuf, dest_x int32, dest_y int32, dest_width int32, dest_height int32, offset_x float64, offset_y float64, scale_x float64, scale_y float64, interp_type InterpTypeEnum, overall_alpha int32) {
	iv, err := _I.Get(23, "Pixbuf", "composite")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dest := gi.NewPointerArgument(dest.P_Pixbuf())
	arg_dest_x := gi.NewInt32Argument(dest_x)
	arg_dest_y := gi.NewInt32Argument(dest_y)
	arg_dest_width := gi.NewInt32Argument(dest_width)
	arg_dest_height := gi.NewInt32Argument(dest_height)
	arg_offset_x := gi.NewDoubleArgument(offset_x)
	arg_offset_y := gi.NewDoubleArgument(offset_y)
	arg_scale_x := gi.NewDoubleArgument(scale_x)
	arg_scale_y := gi.NewDoubleArgument(scale_y)
	arg_interp_type := gi.NewIntArgument(int(interp_type))
	arg_overall_alpha := gi.NewInt32Argument(overall_alpha)
	args := []gi.Argument{arg_v, arg_dest, arg_dest_x, arg_dest_y, arg_dest_width, arg_dest_height, arg_offset_x, arg_offset_y, arg_scale_x, arg_scale_y, arg_interp_type, arg_overall_alpha}
	iv.Call(args, nil, nil)
}

// gdk_pixbuf_composite_color
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) CompositeColor(dest IPixbuf, dest_x int32, dest_y int32, dest_width int32, dest_height int32, offset_x float64, offset_y float64, scale_x float64, scale_y float64, interp_type InterpTypeEnum, overall_alpha int32, check_x int32, check_y int32, check_size int32, color1 uint32, color2 uint32) {
	iv, err := _I.Get(24, "Pixbuf", "composite_color")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dest := gi.NewPointerArgument(dest.P_Pixbuf())
	arg_dest_x := gi.NewInt32Argument(dest_x)
	arg_dest_y := gi.NewInt32Argument(dest_y)
	arg_dest_width := gi.NewInt32Argument(dest_width)
	arg_dest_height := gi.NewInt32Argument(dest_height)
	arg_offset_x := gi.NewDoubleArgument(offset_x)
	arg_offset_y := gi.NewDoubleArgument(offset_y)
	arg_scale_x := gi.NewDoubleArgument(scale_x)
	arg_scale_y := gi.NewDoubleArgument(scale_y)
	arg_interp_type := gi.NewIntArgument(int(interp_type))
	arg_overall_alpha := gi.NewInt32Argument(overall_alpha)
	arg_check_x := gi.NewInt32Argument(check_x)
	arg_check_y := gi.NewInt32Argument(check_y)
	arg_check_size := gi.NewInt32Argument(check_size)
	arg_color1 := gi.NewUint32Argument(color1)
	arg_color2 := gi.NewUint32Argument(color2)
	args := []gi.Argument{arg_v, arg_dest, arg_dest_x, arg_dest_y, arg_dest_width, arg_dest_height, arg_offset_x, arg_offset_y, arg_scale_x, arg_scale_y, arg_interp_type, arg_overall_alpha, arg_check_x, arg_check_y, arg_check_size, arg_color1, arg_color2}
	iv.Call(args, nil, nil)
}

// gdk_pixbuf_composite_color_simple
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) CompositeColorSimple(dest_width int32, dest_height int32, interp_type InterpTypeEnum, overall_alpha int32, check_size int32, color1 uint32, color2 uint32) (result Pixbuf) {
	iv, err := _I.Get(25, "Pixbuf", "composite_color_simple")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dest_width := gi.NewInt32Argument(dest_width)
	arg_dest_height := gi.NewInt32Argument(dest_height)
	arg_interp_type := gi.NewIntArgument(int(interp_type))
	arg_overall_alpha := gi.NewInt32Argument(overall_alpha)
	arg_check_size := gi.NewInt32Argument(check_size)
	arg_color1 := gi.NewUint32Argument(color1)
	arg_color2 := gi.NewUint32Argument(color2)
	args := []gi.Argument{arg_v, arg_dest_width, arg_dest_height, arg_interp_type, arg_overall_alpha, arg_check_size, arg_color1, arg_color2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_copy
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) Copy() (result Pixbuf) {
	iv, err := _I.Get(26, "Pixbuf", "copy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_copy_area
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) CopyArea(src_x int32, src_y int32, width int32, height int32, dest_pixbuf IPixbuf, dest_x int32, dest_y int32) {
	iv, err := _I.Get(27, "Pixbuf", "copy_area")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_src_x := gi.NewInt32Argument(src_x)
	arg_src_y := gi.NewInt32Argument(src_y)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_dest_pixbuf := gi.NewPointerArgument(dest_pixbuf.P_Pixbuf())
	arg_dest_x := gi.NewInt32Argument(dest_x)
	arg_dest_y := gi.NewInt32Argument(dest_y)
	args := []gi.Argument{arg_v, arg_src_x, arg_src_y, arg_width, arg_height, arg_dest_pixbuf, arg_dest_x, arg_dest_y}
	iv.Call(args, nil, nil)
}

// gdk_pixbuf_copy_options
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) CopyOptions(dest_pixbuf IPixbuf) (result bool) {
	iv, err := _I.Get(28, "Pixbuf", "copy_options")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dest_pixbuf := gi.NewPointerArgument(dest_pixbuf.P_Pixbuf())
	args := []gi.Argument{arg_v, arg_dest_pixbuf}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_pixbuf_fill
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) Fill(pixel uint32) {
	iv, err := _I.Get(29, "Pixbuf", "fill")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pixel := gi.NewUint32Argument(pixel)
	args := []gi.Argument{arg_v, arg_pixel}
	iv.Call(args, nil, nil)
}

// gdk_pixbuf_flip
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) Flip(horizontal bool) (result Pixbuf) {
	iv, err := _I.Get(30, "Pixbuf", "flip")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_horizontal := gi.NewBoolArgument(horizontal)
	args := []gi.Argument{arg_v, arg_horizontal}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_get_bits_per_sample
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) GetBitsPerSample() (result int32) {
	iv, err := _I.Get(31, "Pixbuf", "get_bits_per_sample")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gdk_pixbuf_get_byte_length
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) GetByteLength() (result uint64) {
	iv, err := _I.Get(32, "Pixbuf", "get_byte_length")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// gdk_pixbuf_get_colorspace
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) GetColorspace() (result ColorspaceEnum) {
	iv, err := _I.Get(33, "Pixbuf", "get_colorspace")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ColorspaceEnum(ret.Int())
	return
}

// gdk_pixbuf_get_has_alpha
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) GetHasAlpha() (result bool) {
	iv, err := _I.Get(34, "Pixbuf", "get_has_alpha")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_pixbuf_get_height
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) GetHeight() (result int32) {
	iv, err := _I.Get(35, "Pixbuf", "get_height")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gdk_pixbuf_get_n_channels
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) GetNChannels() (result int32) {
	iv, err := _I.Get(36, "Pixbuf", "get_n_channels")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gdk_pixbuf_get_option
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) GetOption(key string) (result string) {
	iv, err := _I.Get(37, "Pixbuf", "get_option")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	args := []gi.Argument{arg_v, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_key)
	result = ret.String().Take()
	return
}

// gdk_pixbuf_get_options
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) GetOptions() (result glib.HashTable) {
	iv, err := _I.Get(38, "Pixbuf", "get_options")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_get_pixels_with_length
// container is not nil, container is Pixbuf
// is method
// ret lenArgIdx 0
func (v Pixbuf) GetPixels() (result gi.Uint8Array) {
	iv, err := _I.Get(39, "Pixbuf", "get_pixels")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var length uint32
	_ = length
	length = outArgs[0].Uint32()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(length)}
	return
}

// gdk_pixbuf_get_rowstride
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) GetRowstride() (result int32) {
	iv, err := _I.Get(40, "Pixbuf", "get_rowstride")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gdk_pixbuf_get_width
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) GetWidth() (result int32) {
	iv, err := _I.Get(41, "Pixbuf", "get_width")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gdk_pixbuf_new_subpixbuf
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) NewSubpixbuf(src_x int32, src_y int32, width int32, height int32) (result Pixbuf) {
	iv, err := _I.Get(42, "Pixbuf", "new_subpixbuf")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_src_x := gi.NewInt32Argument(src_x)
	arg_src_y := gi.NewInt32Argument(src_y)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	args := []gi.Argument{arg_v, arg_src_x, arg_src_y, arg_width, arg_height}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_read_pixel_bytes
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) ReadPixelBytes() (result glib.Bytes) {
	iv, err := _I.Get(43, "Pixbuf", "read_pixel_bytes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_read_pixels
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) ReadPixels() (result uint8) {
	iv, err := _I.Get(44, "Pixbuf", "read_pixels")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// gdk_pixbuf_remove_option
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) RemoveOption(key string) (result bool) {
	iv, err := _I.Get(45, "Pixbuf", "remove_option")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	args := []gi.Argument{arg_v, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_key)
	result = ret.Bool()
	return
}

// gdk_pixbuf_rotate_simple
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) RotateSimple(angle PixbufRotationEnum) (result Pixbuf) {
	iv, err := _I.Get(46, "Pixbuf", "rotate_simple")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_angle := gi.NewIntArgument(int(angle))
	args := []gi.Argument{arg_v, arg_angle}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_saturate_and_pixelate
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) SaturateAndPixelate(dest IPixbuf, saturation float32, pixelate bool) {
	iv, err := _I.Get(47, "Pixbuf", "saturate_and_pixelate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dest := gi.NewPointerArgument(dest.P_Pixbuf())
	arg_saturation := gi.NewFloatArgument(saturation)
	arg_pixelate := gi.NewBoolArgument(pixelate)
	args := []gi.Argument{arg_v, arg_dest, arg_saturation, arg_pixelate}
	iv.Call(args, nil, nil)
}

// gdk_pixbuf_save_to_bufferv
// container is not nil, container is Pixbuf
// is method
// arg 0 buffer lenArgIdx 1
func (v Pixbuf) SaveToBufferv(type1 string, option_keys gi.CStrArray, option_values gi.CStrArray) (result bool, buffer gi.Uint8Array, err error) {
	iv, err := _I.Get(48, "Pixbuf", "save_to_bufferv")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_type1 := gi.CString(type1)
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_buffer_size := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_type1 := gi.NewStringArgument(c_type1)
	arg_option_keys := gi.NewPointerArgument(option_keys.P)
	arg_option_values := gi.NewPointerArgument(option_values.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_buffer, arg_buffer_size, arg_type1, arg_option_keys, arg_option_values, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var buffer_size uint64
	_ = buffer_size
	gi.Free(c_type1)
	err = gi.ToError(outArgs[2].Pointer())
	buffer.P = outArgs[0].Pointer()
	buffer_size = outArgs[1].Uint64()
	result = ret.Bool()
	buffer.Len = int(buffer_size)
	return
}

// gdk_pixbuf_save_to_callbackv
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) SaveToCallbackv(save_func int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer, type1 string, option_keys gi.CStrArray, option_values gi.CStrArray) (result bool, err error) {
	iv, err := _I.Get(49, "Pixbuf", "save_to_callbackv")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_type1 := gi.CString(type1)
	arg_v := gi.NewPointerArgument(v.P)
	arg_save_func := gi.NewIntArgument(save_func) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_type1 := gi.NewStringArgument(c_type1)
	arg_option_keys := gi.NewPointerArgument(option_keys.P)
	arg_option_values := gi.NewPointerArgument(option_values.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_save_func, arg_user_data, arg_type1, arg_option_keys, arg_option_values, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_type1)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// gdk_pixbuf_save_to_streamv
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) SaveToStreamv(stream gio.IOutputStream, type1 string, option_keys gi.CStrArray, option_values gi.CStrArray, cancellable gio.ICancellable) (result bool, err error) {
	iv, err := _I.Get(50, "Pixbuf", "save_to_streamv")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_type1 := gi.CString(type1)
	arg_v := gi.NewPointerArgument(v.P)
	arg_stream := gi.NewPointerArgument(stream.P_OutputStream())
	arg_type1 := gi.NewStringArgument(c_type1)
	arg_option_keys := gi.NewPointerArgument(option_keys.P)
	arg_option_values := gi.NewPointerArgument(option_values.P)
	arg_cancellable := gi.NewPointerArgument(cancellable.P_Cancellable())
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_stream, arg_type1, arg_option_keys, arg_option_values, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_type1)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// gdk_pixbuf_save_to_streamv_async
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) SaveToStreamvAsync(stream gio.IOutputStream, type1 string, option_keys gi.CStrArray, option_values gi.CStrArray, cancellable gio.ICancellable, callback int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(51, "Pixbuf", "save_to_streamv_async")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_type1 := gi.CString(type1)
	arg_v := gi.NewPointerArgument(v.P)
	arg_stream := gi.NewPointerArgument(stream.P_OutputStream())
	arg_type1 := gi.NewStringArgument(c_type1)
	arg_option_keys := gi.NewPointerArgument(option_keys.P)
	arg_option_values := gi.NewPointerArgument(option_values.P)
	arg_cancellable := gi.NewPointerArgument(cancellable.P_Cancellable())
	arg_callback := gi.NewIntArgument(callback) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_stream, arg_type1, arg_option_keys, arg_option_values, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
	gi.Free(c_type1)
}

// gdk_pixbuf_savev
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) Savev(filename string, type1 string, option_keys gi.CStrArray, option_values gi.CStrArray) (result bool, err error) {
	iv, err := _I.Get(52, "Pixbuf", "savev")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	c_type1 := gi.CString(type1)
	arg_v := gi.NewPointerArgument(v.P)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_type1 := gi.NewStringArgument(c_type1)
	arg_option_keys := gi.NewPointerArgument(option_keys.P)
	arg_option_values := gi.NewPointerArgument(option_values.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_filename, arg_type1, arg_option_keys, arg_option_values, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	gi.Free(c_type1)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// gdk_pixbuf_scale
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) Scale(dest IPixbuf, dest_x int32, dest_y int32, dest_width int32, dest_height int32, offset_x float64, offset_y float64, scale_x float64, scale_y float64, interp_type InterpTypeEnum) {
	iv, err := _I.Get(53, "Pixbuf", "scale")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dest := gi.NewPointerArgument(dest.P_Pixbuf())
	arg_dest_x := gi.NewInt32Argument(dest_x)
	arg_dest_y := gi.NewInt32Argument(dest_y)
	arg_dest_width := gi.NewInt32Argument(dest_width)
	arg_dest_height := gi.NewInt32Argument(dest_height)
	arg_offset_x := gi.NewDoubleArgument(offset_x)
	arg_offset_y := gi.NewDoubleArgument(offset_y)
	arg_scale_x := gi.NewDoubleArgument(scale_x)
	arg_scale_y := gi.NewDoubleArgument(scale_y)
	arg_interp_type := gi.NewIntArgument(int(interp_type))
	args := []gi.Argument{arg_v, arg_dest, arg_dest_x, arg_dest_y, arg_dest_width, arg_dest_height, arg_offset_x, arg_offset_y, arg_scale_x, arg_scale_y, arg_interp_type}
	iv.Call(args, nil, nil)
}

// gdk_pixbuf_scale_simple
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) ScaleSimple(dest_width int32, dest_height int32, interp_type InterpTypeEnum) (result Pixbuf) {
	iv, err := _I.Get(54, "Pixbuf", "scale_simple")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dest_width := gi.NewInt32Argument(dest_width)
	arg_dest_height := gi.NewInt32Argument(dest_height)
	arg_interp_type := gi.NewIntArgument(int(interp_type))
	args := []gi.Argument{arg_v, arg_dest_width, arg_dest_height, arg_interp_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_set_option
// container is not nil, container is Pixbuf
// is method
func (v Pixbuf) SetOption(key string, value string) (result bool) {
	iv, err := _I.Get(55, "Pixbuf", "set_option")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	c_value := gi.CString(value)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewStringArgument(c_value)
	args := []gi.Argument{arg_v, arg_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_key)
	gi.Free(c_value)
	result = ret.Bool()
	return
}

// Enum PixbufAlphaMode
type PixbufAlphaModeEnum int

const (
	PixbufAlphaModeBilevel PixbufAlphaModeEnum = 0
	PixbufAlphaModeFull    PixbufAlphaModeEnum = 1
)

// Object PixbufAnimation
type PixbufAnimation struct {
	gobject.Object
}

func WrapPixbufAnimation(p unsafe.Pointer) (r PixbufAnimation) { r.P = p; return }

type IPixbufAnimation interface{ P_PixbufAnimation() unsafe.Pointer }

func (v PixbufAnimation) P_PixbufAnimation() unsafe.Pointer { return v.P }

// gdk_pixbuf_animation_new_from_file
// container is not nil, container is PixbufAnimation
// is constructor
func NewPixbufAnimationFromFile(filename string) (result PixbufAnimation, err error) {
	iv, err := _I.Get(56, "PixbufAnimation", "new_from_file")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_filename, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_animation_new_from_resource
// container is not nil, container is PixbufAnimation
// is constructor
func NewPixbufAnimationFromResource(resource_path string) (result PixbufAnimation, err error) {
	iv, err := _I.Get(57, "PixbufAnimation", "new_from_resource")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_resource_path := gi.CString(resource_path)
	arg_resource_path := gi.NewStringArgument(c_resource_path)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_resource_path, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_resource_path)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_animation_new_from_stream
// container is not nil, container is PixbufAnimation
// is constructor
func NewPixbufAnimationFromStream(stream gio.IInputStream, cancellable gio.ICancellable) (result PixbufAnimation, err error) {
	iv, err := _I.Get(58, "PixbufAnimation", "new_from_stream")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_stream := gi.NewPointerArgument(stream.P_InputStream())
	arg_cancellable := gi.NewPointerArgument(cancellable.P_Cancellable())
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_stream, arg_cancellable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_animation_new_from_stream_finish
// container is not nil, container is PixbufAnimation
// is constructor
func NewPixbufAnimationFromStreamFinish(async_result gio.AsyncResult) (result PixbufAnimation, err error) {
	iv, err := _I.Get(59, "PixbufAnimation", "new_from_stream_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_async_result := gi.NewPointerArgument(async_result.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_async_result, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_animation_new_from_stream_async
// container is not nil, container is PixbufAnimation
// is method
// arg0Type tag: interface, isPtr: true
func PixbufAnimationNewFromStreamAsync1(stream gio.IInputStream, cancellable gio.ICancellable, callback int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(60, "PixbufAnimation", "new_from_stream_async")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_stream := gi.NewPointerArgument(stream.P_InputStream())
	arg_cancellable := gi.NewPointerArgument(cancellable.P_Cancellable())
	arg_callback := gi.NewIntArgument(callback) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_stream, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// gdk_pixbuf_animation_get_height
// container is not nil, container is PixbufAnimation
// is method
func (v PixbufAnimation) GetHeight() (result int32) {
	iv, err := _I.Get(61, "PixbufAnimation", "get_height")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gdk_pixbuf_animation_get_iter
// container is not nil, container is PixbufAnimation
// is method
func (v PixbufAnimation) GetIter(start_time glib.TimeVal) (result PixbufAnimationIter) {
	iv, err := _I.Get(62, "PixbufAnimation", "get_iter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start_time := gi.NewPointerArgument(start_time.P)
	args := []gi.Argument{arg_v, arg_start_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_animation_get_static_image
// container is not nil, container is PixbufAnimation
// is method
func (v PixbufAnimation) GetStaticImage() (result Pixbuf) {
	iv, err := _I.Get(63, "PixbufAnimation", "get_static_image")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_animation_get_width
// container is not nil, container is PixbufAnimation
// is method
func (v PixbufAnimation) GetWidth() (result int32) {
	iv, err := _I.Get(64, "PixbufAnimation", "get_width")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gdk_pixbuf_animation_is_static_image
// container is not nil, container is PixbufAnimation
// is method
func (v PixbufAnimation) IsStaticImage() (result bool) {
	iv, err := _I.Get(65, "PixbufAnimation", "is_static_image")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Object PixbufAnimationIter
type PixbufAnimationIter struct {
	gobject.Object
}

func WrapPixbufAnimationIter(p unsafe.Pointer) (r PixbufAnimationIter) { r.P = p; return }

type IPixbufAnimationIter interface{ P_PixbufAnimationIter() unsafe.Pointer }

func (v PixbufAnimationIter) P_PixbufAnimationIter() unsafe.Pointer { return v.P }

// gdk_pixbuf_animation_iter_advance
// container is not nil, container is PixbufAnimationIter
// is method
func (v PixbufAnimationIter) Advance(current_time glib.TimeVal) (result bool) {
	iv, err := _I.Get(66, "PixbufAnimationIter", "advance")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_current_time := gi.NewPointerArgument(current_time.P)
	args := []gi.Argument{arg_v, arg_current_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_pixbuf_animation_iter_get_delay_time
// container is not nil, container is PixbufAnimationIter
// is method
func (v PixbufAnimationIter) GetDelayTime() (result int32) {
	iv, err := _I.Get(67, "PixbufAnimationIter", "get_delay_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gdk_pixbuf_animation_iter_get_pixbuf
// container is not nil, container is PixbufAnimationIter
// is method
func (v PixbufAnimationIter) GetPixbuf() (result Pixbuf) {
	iv, err := _I.Get(68, "PixbufAnimationIter", "get_pixbuf")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_animation_iter_on_currently_loading_frame
// container is not nil, container is PixbufAnimationIter
// is method
func (v PixbufAnimationIter) OnCurrentlyLoadingFrame() (result bool) {
	iv, err := _I.Get(69, "PixbufAnimationIter", "on_currently_loading_frame")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Enum PixbufError
type PixbufErrorEnum int

const (
	PixbufErrorCorruptImage         PixbufErrorEnum = 0
	PixbufErrorInsufficientMemory   PixbufErrorEnum = 1
	PixbufErrorBadOption            PixbufErrorEnum = 2
	PixbufErrorUnknownType          PixbufErrorEnum = 3
	PixbufErrorUnsupportedOperation PixbufErrorEnum = 4
	PixbufErrorFailed               PixbufErrorEnum = 5
	PixbufErrorIncompleteAnimation  PixbufErrorEnum = 6
)

// Struct PixbufFormat
type PixbufFormat struct {
	P unsafe.Pointer
}

// gdk_pixbuf_format_copy
// container is not nil, container is PixbufFormat
// is method
func (v PixbufFormat) Copy() (result PixbufFormat) {
	iv, err := _I.Get(70, "PixbufFormat", "copy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_format_free
// container is not nil, container is PixbufFormat
// is method
func (v PixbufFormat) Free() {
	iv, err := _I.Get(71, "PixbufFormat", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_pixbuf_format_get_description
// container is not nil, container is PixbufFormat
// is method
func (v PixbufFormat) GetDescription() (result string) {
	iv, err := _I.Get(72, "PixbufFormat", "get_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// gdk_pixbuf_format_get_extensions
// container is not nil, container is PixbufFormat
// is method
func (v PixbufFormat) GetExtensions() (result gi.CStrArray) {
	iv, err := _I.Get(73, "PixbufFormat", "get_extensions")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// gdk_pixbuf_format_get_license
// container is not nil, container is PixbufFormat
// is method
func (v PixbufFormat) GetLicense() (result string) {
	iv, err := _I.Get(74, "PixbufFormat", "get_license")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// gdk_pixbuf_format_get_mime_types
// container is not nil, container is PixbufFormat
// is method
func (v PixbufFormat) GetMimeTypes() (result gi.CStrArray) {
	iv, err := _I.Get(75, "PixbufFormat", "get_mime_types")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// gdk_pixbuf_format_get_name
// container is not nil, container is PixbufFormat
// is method
func (v PixbufFormat) GetName() (result string) {
	iv, err := _I.Get(76, "PixbufFormat", "get_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// gdk_pixbuf_format_is_disabled
// container is not nil, container is PixbufFormat
// is method
func (v PixbufFormat) IsDisabled() (result bool) {
	iv, err := _I.Get(77, "PixbufFormat", "is_disabled")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_pixbuf_format_is_save_option_supported
// container is not nil, container is PixbufFormat
// is method
func (v PixbufFormat) IsSaveOptionSupported(option_key string) (result bool) {
	iv, err := _I.Get(78, "PixbufFormat", "is_save_option_supported")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_option_key := gi.CString(option_key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_option_key := gi.NewStringArgument(c_option_key)
	args := []gi.Argument{arg_v, arg_option_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_option_key)
	result = ret.Bool()
	return
}

// gdk_pixbuf_format_is_scalable
// container is not nil, container is PixbufFormat
// is method
func (v PixbufFormat) IsScalable() (result bool) {
	iv, err := _I.Get(79, "PixbufFormat", "is_scalable")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_pixbuf_format_is_writable
// container is not nil, container is PixbufFormat
// is method
func (v PixbufFormat) IsWritable() (result bool) {
	iv, err := _I.Get(80, "PixbufFormat", "is_writable")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_pixbuf_format_set_disabled
// container is not nil, container is PixbufFormat
// is method
func (v PixbufFormat) SetDisabled(disabled bool) {
	iv, err := _I.Get(81, "PixbufFormat", "set_disabled")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_disabled := gi.NewBoolArgument(disabled)
	args := []gi.Argument{arg_v, arg_disabled}
	iv.Call(args, nil, nil)
}

// Object PixbufLoader
type PixbufLoader struct {
	gobject.Object
}

func WrapPixbufLoader(p unsafe.Pointer) (r PixbufLoader) { r.P = p; return }

type IPixbufLoader interface{ P_PixbufLoader() unsafe.Pointer }

func (v PixbufLoader) P_PixbufLoader() unsafe.Pointer { return v.P }

// gdk_pixbuf_loader_new
// container is not nil, container is PixbufLoader
// is constructor
func NewPixbufLoader() (result PixbufLoader) {
	iv, err := _I.Get(82, "PixbufLoader", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_loader_new_with_mime_type
// container is not nil, container is PixbufLoader
// is constructor
func NewPixbufLoaderWithMimeType(mime_type string) (result PixbufLoader, err error) {
	iv, err := _I.Get(83, "PixbufLoader", "new_with_mime_type")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_mime_type := gi.CString(mime_type)
	arg_mime_type := gi.NewStringArgument(c_mime_type)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_mime_type, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_mime_type)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_loader_new_with_type
// container is not nil, container is PixbufLoader
// is constructor
func NewPixbufLoaderWithType(image_type string) (result PixbufLoader, err error) {
	iv, err := _I.Get(84, "PixbufLoader", "new_with_type")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_image_type := gi.CString(image_type)
	arg_image_type := gi.NewStringArgument(c_image_type)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_image_type, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_image_type)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_loader_close
// container is not nil, container is PixbufLoader
// is method
func (v PixbufLoader) Close() (result bool, err error) {
	iv, err := _I.Get(85, "PixbufLoader", "close")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// gdk_pixbuf_loader_get_animation
// container is not nil, container is PixbufLoader
// is method
func (v PixbufLoader) GetAnimation() (result PixbufAnimation) {
	iv, err := _I.Get(86, "PixbufLoader", "get_animation")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_loader_get_format
// container is not nil, container is PixbufLoader
// is method
func (v PixbufLoader) GetFormat() (result PixbufFormat) {
	iv, err := _I.Get(87, "PixbufLoader", "get_format")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_loader_get_pixbuf
// container is not nil, container is PixbufLoader
// is method
func (v PixbufLoader) GetPixbuf() (result Pixbuf) {
	iv, err := _I.Get(88, "PixbufLoader", "get_pixbuf")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_loader_set_size
// container is not nil, container is PixbufLoader
// is method
func (v PixbufLoader) SetSize(width int32, height int32) {
	iv, err := _I.Get(89, "PixbufLoader", "set_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	args := []gi.Argument{arg_v, arg_width, arg_height}
	iv.Call(args, nil, nil)
}

// gdk_pixbuf_loader_write
// container is not nil, container is PixbufLoader
// is method
// arg 0 buf lenArgIdx 1
func (v PixbufLoader) Write(buf gi.Uint8Array, count uint64) (result bool, err error) {
	iv, err := _I.Get(90, "PixbufLoader", "write")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_count := gi.NewUint64Argument(count)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_buf, arg_count, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// gdk_pixbuf_loader_write_bytes
// container is not nil, container is PixbufLoader
// is method
func (v PixbufLoader) WriteBytes(buffer glib.Bytes) (result bool, err error) {
	iv, err := _I.Get(91, "PixbufLoader", "write_bytes")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer := gi.NewPointerArgument(buffer.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_buffer, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// ignore GType struct PixbufLoaderClass
// Enum PixbufRotation
type PixbufRotationEnum int

const (
	PixbufRotationNone             PixbufRotationEnum = 0
	PixbufRotationCounterclockwise PixbufRotationEnum = 90
	PixbufRotationUpsidedown       PixbufRotationEnum = 180
	PixbufRotationClockwise        PixbufRotationEnum = 270
)

// Object PixbufSimpleAnim
type PixbufSimpleAnim struct {
	PixbufAnimation
}

func WrapPixbufSimpleAnim(p unsafe.Pointer) (r PixbufSimpleAnim) { r.P = p; return }

type IPixbufSimpleAnim interface{ P_PixbufSimpleAnim() unsafe.Pointer }

func (v PixbufSimpleAnim) P_PixbufSimpleAnim() unsafe.Pointer { return v.P }

// gdk_pixbuf_simple_anim_new
// container is not nil, container is PixbufSimpleAnim
// is constructor
func NewPixbufSimpleAnim(width int32, height int32, rate float32) (result PixbufSimpleAnim) {
	iv, err := _I.Get(92, "PixbufSimpleAnim", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_rate := gi.NewFloatArgument(rate)
	args := []gi.Argument{arg_width, arg_height, arg_rate}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_simple_anim_add_frame
// container is not nil, container is PixbufSimpleAnim
// is method
func (v PixbufSimpleAnim) AddFrame(pixbuf IPixbuf) {
	iv, err := _I.Get(93, "PixbufSimpleAnim", "add_frame")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pixbuf := gi.NewPointerArgument(pixbuf.P_Pixbuf())
	args := []gi.Argument{arg_v, arg_pixbuf}
	iv.Call(args, nil, nil)
}

// gdk_pixbuf_simple_anim_get_loop
// container is not nil, container is PixbufSimpleAnim
// is method
func (v PixbufSimpleAnim) GetLoop() (result bool) {
	iv, err := _I.Get(94, "PixbufSimpleAnim", "get_loop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_pixbuf_simple_anim_set_loop
// container is not nil, container is PixbufSimpleAnim
// is method
func (v PixbufSimpleAnim) SetLoop(loop bool) {
	iv, err := _I.Get(95, "PixbufSimpleAnim", "set_loop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_loop := gi.NewBoolArgument(loop)
	args := []gi.Argument{arg_v, arg_loop}
	iv.Call(args, nil, nil)
}

// ignore GType struct PixbufSimpleAnimClass
// Object PixbufSimpleAnimIter
type PixbufSimpleAnimIter struct {
	PixbufAnimationIter
}

func WrapPixbufSimpleAnimIter(p unsafe.Pointer) (r PixbufSimpleAnimIter) { r.P = p; return }

type IPixbufSimpleAnimIter interface{ P_PixbufSimpleAnimIter() unsafe.Pointer }

func (v PixbufSimpleAnimIter) P_PixbufSimpleAnimIter() unsafe.Pointer { return v.P }

// gdk_pixbuf_error_quark
// container is nil
func PixbufErrorQuark() (result uint32) {
	iv, err := _I.Get(96, "pixbuf_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// constants
const (
	PIXBUF_FEATURES_H = 1
	PIXBUF_MAJOR      = 2
	PIXBUF_MICRO      = 1
	PIXBUF_MINOR      = 38
	PIXBUF_VERSION    = "2.38.1"
)
