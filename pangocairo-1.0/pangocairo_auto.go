package pangocairo

/*
#cgo pkg-config: pangocairo
#include <pango/pangocairo.h>
extern void myPangoCairoShapeRendererFunc(cairo_t* cr, PangoAttrShape* attr, gboolean do_path, gpointer data);
static void* getPointer_myPangoCairoShapeRendererFunc() {
return (void*)(myPangoCairoShapeRendererFunc);
}
*/
import "C"
import "github.com/electricface/go-gir/cairo-1.0"
import "github.com/electricface/go-gir/g-2.0"
import "github.com/electricface/go-gir/pango-1.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("PangoCairo")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("PangoCairo", "1.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Interface Font
type Font struct {
	FontIfc
	P unsafe.Pointer
}
type FontIfc struct{}
type IFont interface{ P_Font() unsafe.Pointer }

func (v Font) P_Font() unsafe.Pointer { return v.P }
func FontGetType() gi.GType {
	ret := _I.GetGType(0, "Font")
	return ret
}

// pango_cairo_font_get_scaled_font
//
// [ result ] trans: everything
//
func (v *FontIfc) GetScaledFont() (result cairo.ScaledFont) {
	iv, err := _I.Get(0, "Font", "get_scaled_font", 0, 0, gi.INFO_TYPE_INTERFACE, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Interface FontMap
type FontMap struct {
	FontMapIfc
	P unsafe.Pointer
}
type FontMapIfc struct{}
type IFontMap interface{ P_FontMap() unsafe.Pointer }

func (v FontMap) P_FontMap() unsafe.Pointer { return v.P }
func FontMapGetType() gi.GType {
	ret := _I.GetGType(1, "FontMap")
	return ret
}

// pango_cairo_font_map_new_for_font_type
//
// [ fonttype ] trans: nothing
//
// [ result ] trans: everything
//
func FontMapNewForFontType1(fonttype cairo.FontTypeEnum) (result pango.FontMap) {
	iv, err := _I.Get(3, "FontMap", "new_for_font_type", 1, 2, gi.INFO_TYPE_INTERFACE, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_fonttype := gi.NewIntArgument(int(fonttype))
	args := []gi.Argument{arg_fonttype}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_cairo_font_map_get_font_type
//
// [ result ] trans: nothing
//
func (v *FontMapIfc) GetFontType() (result cairo.FontTypeEnum) {
	iv, err := _I.Get(4, "FontMap", "get_font_type", 1, 3, gi.INFO_TYPE_INTERFACE, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = cairo.FontTypeEnum(ret.Int())
	return
}

// pango_cairo_font_map_get_resolution
//
// [ result ] trans: nothing
//
func (v *FontMapIfc) GetResolution() (result float64) {
	iv, err := _I.Get(5, "FontMap", "get_resolution", 1, 4, gi.INFO_TYPE_INTERFACE, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// pango_cairo_font_map_set_default
//
func (v *FontMapIfc) SetDefault() {
	iv, err := _I.Get(6, "FontMap", "set_default", 1, 5, gi.INFO_TYPE_INTERFACE, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_cairo_font_map_set_resolution
//
// [ dpi ] trans: nothing
//
func (v *FontMapIfc) SetResolution(dpi float64) {
	iv, err := _I.Get(7, "FontMap", "set_resolution", 1, 6, gi.INFO_TYPE_INTERFACE, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_dpi := gi.NewDoubleArgument(dpi)
	args := []gi.Argument{arg_v, arg_dpi}
	iv.Call(args, nil, nil)
}

type ShapeRendererFuncStruct struct {
	F_cr      cairo.Context
	F_attr    pango.AttrShape
	F_do_path bool
	F_data    unsafe.Pointer
}

func GetPointer_myShapeRendererFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myPangoCairoShapeRendererFunc())
}

//export myPangoCairoShapeRendererFunc
func myPangoCairoShapeRendererFunc(cr *C.cairo_t, attr *C.PangoAttrShape, do_path C.gboolean, data C.gpointer) {
	// TODO: not found user_data
}

// pango_cairo_context_get_font_options
//
// [ context ] trans: nothing
//
// [ result ] trans: nothing
//
func ContextGetFontOptions(context pango.IContext) (result cairo.FontOptions) {
	iv, err := _I.Get(8, "context_get_font_options", "", 3, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_Context()
	}
	arg_context := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_context}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_cairo_context_get_resolution
//
// [ context ] trans: nothing
//
// [ result ] trans: nothing
//
func ContextGetResolution(context pango.IContext) (result float64) {
	iv, err := _I.Get(9, "context_get_resolution", "", 4, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_Context()
	}
	arg_context := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_context}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// pango_cairo_context_set_font_options
//
// [ context ] trans: nothing
//
// [ options ] trans: nothing
//
func ContextSetFontOptions(context pango.IContext, options cairo.FontOptions) {
	iv, err := _I.Get(10, "context_set_font_options", "", 5, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_Context()
	}
	arg_context := gi.NewPointerArgument(tmp)
	arg_options := gi.NewPointerArgument(options.P)
	args := []gi.Argument{arg_context, arg_options}
	iv.Call(args, nil, nil)
}

// pango_cairo_context_set_resolution
//
// [ context ] trans: nothing
//
// [ dpi ] trans: nothing
//
func ContextSetResolution(context pango.IContext, dpi float64) {
	iv, err := _I.Get(11, "context_set_resolution", "", 6, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_Context()
	}
	arg_context := gi.NewPointerArgument(tmp)
	arg_dpi := gi.NewDoubleArgument(dpi)
	args := []gi.Argument{arg_context, arg_dpi}
	iv.Call(args, nil, nil)
}

// pango_cairo_context_set_shape_renderer
//
// [ context ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ data ] trans: nothing
//
// [ dnotify ] trans: nothing
//
func ContextSetShapeRenderer(context pango.IContext, func1 int /*TODO_TYPE CALLBACK*/, data unsafe.Pointer, dnotify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(12, "context_set_shape_renderer", "", 7, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_Context()
	}
	arg_context := gi.NewPointerArgument(tmp)
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myShapeRendererFunc()))
	arg_data := gi.NewPointerArgument(data)
	arg_dnotify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_context, arg_func1, arg_data, arg_dnotify}
	iv.Call(args, nil, nil)
}

// pango_cairo_create_context
//
// [ cr ] trans: nothing
//
// [ result ] trans: everything
//
func CreateContext(cr cairo.Context) (result pango.Context) {
	iv, err := _I.Get(13, "create_context", "", 8, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	args := []gi.Argument{arg_cr}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_cairo_create_layout
//
// [ cr ] trans: nothing
//
// [ result ] trans: everything
//
func CreateLayout(cr cairo.Context) (result pango.Layout) {
	iv, err := _I.Get(14, "create_layout", "", 9, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	args := []gi.Argument{arg_cr}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_cairo_error_underline_path
//
// [ cr ] trans: nothing
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
func ErrorUnderlinePath(cr cairo.Context, x float64, y float64, width float64, height float64) {
	iv, err := _I.Get(15, "error_underline_path", "", 10, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_x := gi.NewDoubleArgument(x)
	arg_y := gi.NewDoubleArgument(y)
	arg_width := gi.NewDoubleArgument(width)
	arg_height := gi.NewDoubleArgument(height)
	args := []gi.Argument{arg_cr, arg_x, arg_y, arg_width, arg_height}
	iv.Call(args, nil, nil)
}

// pango_cairo_font_map_get_default
//
// [ result ] trans: nothing
//
func FontMapGetDefault() (result pango.FontMap) {
	iv, err := _I.Get(16, "font_map_get_default", "", 11, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_cairo_font_map_new
//
// [ result ] trans: everything
//
func FontMapNew() (result pango.FontMap) {
	iv, err := _I.Get(17, "font_map_new", "", 12, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_cairo_font_map_new_for_font_type
//
// [ fonttype ] trans: nothing
//
// [ result ] trans: everything
//
func FontMapNewForFontType(fonttype cairo.FontTypeEnum) (result pango.FontMap) {
	iv, err := _I.Get(18, "font_map_new_for_font_type", "", 13, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_fonttype := gi.NewIntArgument(int(fonttype))
	args := []gi.Argument{arg_fonttype}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_cairo_glyph_string_path
//
// [ cr ] trans: nothing
//
// [ font ] trans: nothing
//
// [ glyphs ] trans: nothing
//
func GlyphStringPath(cr cairo.Context, font pango.IFont, glyphs pango.GlyphString) {
	iv, err := _I.Get(19, "glyph_string_path", "", 14, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if font != nil {
		tmp = font.P_Font()
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_font := gi.NewPointerArgument(tmp)
	arg_glyphs := gi.NewPointerArgument(glyphs.P)
	args := []gi.Argument{arg_cr, arg_font, arg_glyphs}
	iv.Call(args, nil, nil)
}

// pango_cairo_layout_line_path
//
// [ cr ] trans: nothing
//
// [ line ] trans: nothing
//
func LayoutLinePath(cr cairo.Context, line pango.LayoutLine) {
	iv, err := _I.Get(20, "layout_line_path", "", 15, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_line := gi.NewPointerArgument(line.P)
	args := []gi.Argument{arg_cr, arg_line}
	iv.Call(args, nil, nil)
}

// pango_cairo_layout_path
//
// [ cr ] trans: nothing
//
// [ layout ] trans: nothing
//
func LayoutPath(cr cairo.Context, layout pango.ILayout) {
	iv, err := _I.Get(21, "layout_path", "", 16, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if layout != nil {
		tmp = layout.P_Layout()
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_layout := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_cr, arg_layout}
	iv.Call(args, nil, nil)
}

// pango_cairo_show_error_underline
//
// [ cr ] trans: nothing
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
func ShowErrorUnderline(cr cairo.Context, x float64, y float64, width float64, height float64) {
	iv, err := _I.Get(22, "show_error_underline", "", 17, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_x := gi.NewDoubleArgument(x)
	arg_y := gi.NewDoubleArgument(y)
	arg_width := gi.NewDoubleArgument(width)
	arg_height := gi.NewDoubleArgument(height)
	args := []gi.Argument{arg_cr, arg_x, arg_y, arg_width, arg_height}
	iv.Call(args, nil, nil)
}

// pango_cairo_show_glyph_item
//
// [ cr ] trans: nothing
//
// [ text ] trans: nothing
//
// [ glyph_item ] trans: nothing
//
func ShowGlyphItem(cr cairo.Context, text string, glyph_item pango.GlyphItem) {
	iv, err := _I.Get(23, "show_glyph_item", "", 18, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_glyph_item := gi.NewPointerArgument(glyph_item.P)
	args := []gi.Argument{arg_cr, arg_text, arg_glyph_item}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// pango_cairo_show_glyph_string
//
// [ cr ] trans: nothing
//
// [ font ] trans: nothing
//
// [ glyphs ] trans: nothing
//
func ShowGlyphString(cr cairo.Context, font pango.IFont, glyphs pango.GlyphString) {
	iv, err := _I.Get(24, "show_glyph_string", "", 19, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if font != nil {
		tmp = font.P_Font()
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_font := gi.NewPointerArgument(tmp)
	arg_glyphs := gi.NewPointerArgument(glyphs.P)
	args := []gi.Argument{arg_cr, arg_font, arg_glyphs}
	iv.Call(args, nil, nil)
}

// pango_cairo_show_layout
//
// [ cr ] trans: nothing
//
// [ layout ] trans: nothing
//
func ShowLayout(cr cairo.Context, layout pango.ILayout) {
	iv, err := _I.Get(25, "show_layout", "", 20, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if layout != nil {
		tmp = layout.P_Layout()
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_layout := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_cr, arg_layout}
	iv.Call(args, nil, nil)
}

// pango_cairo_show_layout_line
//
// [ cr ] trans: nothing
//
// [ line ] trans: nothing
//
func ShowLayoutLine(cr cairo.Context, line pango.LayoutLine) {
	iv, err := _I.Get(26, "show_layout_line", "", 21, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_line := gi.NewPointerArgument(line.P)
	args := []gi.Argument{arg_cr, arg_line}
	iv.Call(args, nil, nil)
}

// pango_cairo_update_context
//
// [ cr ] trans: nothing
//
// [ context ] trans: nothing
//
func UpdateContext(cr cairo.Context, context pango.IContext) {
	iv, err := _I.Get(27, "update_context", "", 22, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_Context()
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_context := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_cr, arg_context}
	iv.Call(args, nil, nil)
}

// pango_cairo_update_layout
//
// [ cr ] trans: nothing
//
// [ layout ] trans: nothing
//
func UpdateLayout(cr cairo.Context, layout pango.ILayout) {
	iv, err := _I.Get(28, "update_layout", "", 23, 0, gi.INFO_TYPE_FUNCTION, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if layout != nil {
		tmp = layout.P_Layout()
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_layout := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_cr, arg_layout}
	iv.Call(args, nil, nil)
}

// constants
const ()
