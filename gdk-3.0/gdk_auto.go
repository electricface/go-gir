package gdk

/*
#cgo pkg-config: gdk-3.0
#include <gdk/gdk.h>
extern void myGdkEventFunc(GdkEvent* event, gpointer data);
static void* getPointer_myGdkEventFunc() {
return (void*)(myGdkEventFunc);
}
extern void myGdkFilterFunc(gpointer xevent, GdkEvent* event, gpointer data);
static void* getPointer_myGdkFilterFunc() {
return (void*)(myGdkFilterFunc);
}
extern void myGdkSeatGrabPrepareFunc(GdkSeat* seat, GdkWindow* window, gpointer user_data);
static void* getPointer_myGdkSeatGrabPrepareFunc() {
return (void*)(myGdkSeatGrabPrepareFunc);
}
extern void myGdkWindowChildFunc(GdkWindow* window, gpointer user_data);
static void* getPointer_myGdkWindowChildFunc() {
return (void*)(myGdkWindowChildFunc);
}
extern void myGdkWindowInvalidateHandlerFunc(GdkWindow* window, cairo_region_t* region);
static void* getPointer_myGdkWindowInvalidateHandlerFunc() {
return (void*)(myGdkWindowInvalidateHandlerFunc);
}
*/
import "C"
import "github.com/electricface/go-gir/cairo-1.0"
import "github.com/electricface/go-gir/g-2.0"
import "github.com/electricface/go-gir/gdkpixbuf-2.0"
import "github.com/electricface/go-gir/pango-1.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("Gdk")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("Gdk", "3.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Flags AnchorHints
type AnchorHintsFlags int

const (
	AnchorHintsFlipX   AnchorHintsFlags = 1
	AnchorHintsFlipY   AnchorHintsFlags = 2
	AnchorHintsSlideX  AnchorHintsFlags = 4
	AnchorHintsSlideY  AnchorHintsFlags = 8
	AnchorHintsResizeX AnchorHintsFlags = 16
	AnchorHintsResizeY AnchorHintsFlags = 32
	AnchorHintsFlip    AnchorHintsFlags = 3
	AnchorHintsSlide   AnchorHintsFlags = 12
	AnchorHintsResize  AnchorHintsFlags = 48
)

func AnchorHintsGetType() gi.GType {
	ret := _I.GetGType(0, "AnchorHints")
	return ret
}

// Object AppLaunchContext
type AppLaunchContext struct {
	g.AppLaunchContext
}

func WrapAppLaunchContext(p unsafe.Pointer) (r AppLaunchContext) { r.P = p; return }

type IAppLaunchContext interface{ P_AppLaunchContext() unsafe.Pointer }

func (v AppLaunchContext) P_AppLaunchContext() unsafe.Pointer { return v.P }
func AppLaunchContextGetType() gi.GType {
	ret := _I.GetGType(1, "AppLaunchContext")
	return ret
}

// Deprecated
//
// gdk_app_launch_context_new
//
// [ result ] trans: everything
//
func NewAppLaunchContext() (result AppLaunchContext) {
	iv, err := _I.Get(0, "AppLaunchContext", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_app_launch_context_set_desktop
//
// [ desktop ] trans: nothing
//
func (v AppLaunchContext) SetDesktop(desktop int32) {
	iv, err := _I.Get(1, "AppLaunchContext", "set_desktop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_desktop := gi.NewInt32Argument(desktop)
	args := []gi.Argument{arg_v, arg_desktop}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_app_launch_context_set_display
//
// [ display ] trans: nothing
//
func (v AppLaunchContext) SetDisplay(display IDisplay) {
	iv, err := _I.Get(2, "AppLaunchContext", "set_display")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if display != nil {
		tmp = display.P_Display()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_display := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_display}
	iv.Call(args, nil, nil)
}

// gdk_app_launch_context_set_icon
//
// [ icon ] trans: nothing
//
func (v AppLaunchContext) SetIcon(icon g.IIcon) {
	iv, err := _I.Get(3, "AppLaunchContext", "set_icon")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if icon != nil {
		tmp = icon.P_Icon()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_icon := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_icon}
	iv.Call(args, nil, nil)
}

// gdk_app_launch_context_set_icon_name
//
// [ icon_name ] trans: nothing
//
func (v AppLaunchContext) SetIconName(icon_name string) {
	iv, err := _I.Get(4, "AppLaunchContext", "set_icon_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_icon_name := gi.CString(icon_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_icon_name := gi.NewStringArgument(c_icon_name)
	args := []gi.Argument{arg_v, arg_icon_name}
	iv.Call(args, nil, nil)
	gi.Free(c_icon_name)
}

// gdk_app_launch_context_set_screen
//
// [ screen ] trans: nothing
//
func (v AppLaunchContext) SetScreen(screen IScreen) {
	iv, err := _I.Get(5, "AppLaunchContext", "set_screen")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if screen != nil {
		tmp = screen.P_Screen()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_screen := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_screen}
	iv.Call(args, nil, nil)
}

// gdk_app_launch_context_set_timestamp
//
// [ timestamp ] trans: nothing
//
func (v AppLaunchContext) SetTimestamp(timestamp uint32) {
	iv, err := _I.Get(6, "AppLaunchContext", "set_timestamp")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timestamp := gi.NewUint32Argument(timestamp)
	args := []gi.Argument{arg_v, arg_timestamp}
	iv.Call(args, nil, nil)
}

// Struct Atom
type Atom struct {
	P unsafe.Pointer
}

func AtomGetType() gi.GType {
	ret := _I.GetGType(2, "Atom")
	return ret
}

// gdk_atom_name
//
// [ result ] trans: everything
//
func (v Atom) Name() (result string) {
	iv, err := _I.Get(7, "Atom", "name")
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

// gdk_atom_intern
//
// [ atom_name ] trans: nothing
//
// [ only_if_exists ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomIntern1(atom_name string, only_if_exists bool) (result Atom) {
	iv, err := _I.Get(8, "Atom", "intern")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_atom_name := gi.CString(atom_name)
	arg_atom_name := gi.NewStringArgument(c_atom_name)
	arg_only_if_exists := gi.NewBoolArgument(only_if_exists)
	args := []gi.Argument{arg_atom_name, arg_only_if_exists}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_atom_name)
	result.P = ret.Pointer()
	return
}

// gdk_atom_intern_static_string
//
// [ atom_name ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomInternStaticString1(atom_name string) (result Atom) {
	iv, err := _I.Get(9, "Atom", "intern_static_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_atom_name := gi.CString(atom_name)
	arg_atom_name := gi.NewStringArgument(c_atom_name)
	args := []gi.Argument{arg_atom_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_atom_name)
	result.P = ret.Pointer()
	return
}

// Flags AxisFlags
type AxisFlags int

const (
	AxisFlagsX        AxisFlags = 2
	AxisFlagsY        AxisFlags = 4
	AxisFlagsPressure AxisFlags = 8
	AxisFlagsXtilt    AxisFlags = 16
	AxisFlagsYtilt    AxisFlags = 32
	AxisFlagsWheel    AxisFlags = 64
	AxisFlagsDistance AxisFlags = 128
	AxisFlagsRotation AxisFlags = 256
	AxisFlagsSlider   AxisFlags = 512
)

func AxisFlagsGetType() gi.GType {
	ret := _I.GetGType(3, "AxisFlags")
	return ret
}

// Enum AxisUse
type AxisUseEnum int

const (
	AxisUseIgnore   AxisUseEnum = 0
	AxisUseX        AxisUseEnum = 1
	AxisUseY        AxisUseEnum = 2
	AxisUsePressure AxisUseEnum = 3
	AxisUseXtilt    AxisUseEnum = 4
	AxisUseYtilt    AxisUseEnum = 5
	AxisUseWheel    AxisUseEnum = 6
	AxisUseDistance AxisUseEnum = 7
	AxisUseRotation AxisUseEnum = 8
	AxisUseSlider   AxisUseEnum = 9
	AxisUseLast     AxisUseEnum = 10
)

func AxisUseGetType() gi.GType {
	ret := _I.GetGType(4, "AxisUse")
	return ret
}

// Enum ByteOrder
type ByteOrderEnum int

const (
	ByteOrderLsbFirst ByteOrderEnum = 0
	ByteOrderMsbFirst ByteOrderEnum = 1
)

func ByteOrderGetType() gi.GType {
	ret := _I.GetGType(5, "ByteOrder")
	return ret
}

// Deprecated
//
// Struct Color
type Color struct {
	P unsafe.Pointer
}

const SizeOfStructColor = 12

func ColorGetType() gi.GType {
	ret := _I.GetGType(6, "Color")
	return ret
}

// Deprecated
//
// gdk_color_copy
//
// [ result ] trans: everything
//
func (v Color) Copy() (result Color) {
	iv, err := _I.Get(10, "Color", "copy")
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

// Deprecated
//
// gdk_color_equal
//
// [ colorb ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Color) Equal(colorb Color) (result bool) {
	iv, err := _I.Get(11, "Color", "equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_colorb := gi.NewPointerArgument(colorb.P)
	args := []gi.Argument{arg_v, arg_colorb}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Deprecated
//
// gdk_color_free
//
func (v Color) Free() {
	iv, err := _I.Get(12, "Color", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_color_hash
//
// [ result ] trans: nothing
//
func (v Color) Hash() (result uint32) {
	iv, err := _I.Get(13, "Color", "hash")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// Deprecated
//
// gdk_color_to_string
//
// [ result ] trans: everything
//
func (v Color) ToString() (result string) {
	iv, err := _I.Get(14, "Color", "to_string")
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

// Deprecated
//
// gdk_color_parse
//
// [ spec ] trans: nothing
//
// [ color ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func ColorParse1(spec string, color Color) (result bool) {
	iv, err := _I.Get(15, "Color", "parse")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_spec := gi.CString(spec)
	arg_spec := gi.NewStringArgument(c_spec)
	arg_color := gi.NewPointerArgument(color.P)
	args := []gi.Argument{arg_spec, arg_color}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_spec)
	result = ret.Bool()
	return
}

// Enum CrossingMode
type CrossingModeEnum int

const (
	CrossingModeNormal       CrossingModeEnum = 0
	CrossingModeGrab         CrossingModeEnum = 1
	CrossingModeUngrab       CrossingModeEnum = 2
	CrossingModeGtkGrab      CrossingModeEnum = 3
	CrossingModeGtkUngrab    CrossingModeEnum = 4
	CrossingModeStateChanged CrossingModeEnum = 5
	CrossingModeTouchBegin   CrossingModeEnum = 6
	CrossingModeTouchEnd     CrossingModeEnum = 7
	CrossingModeDeviceSwitch CrossingModeEnum = 8
)

func CrossingModeGetType() gi.GType {
	ret := _I.GetGType(7, "CrossingMode")
	return ret
}

// Object Cursor
type Cursor struct {
	g.Object
}

func WrapCursor(p unsafe.Pointer) (r Cursor) { r.P = p; return }

type ICursor interface{ P_Cursor() unsafe.Pointer }

func (v Cursor) P_Cursor() unsafe.Pointer { return v.P }
func CursorGetType() gi.GType {
	ret := _I.GetGType(8, "Cursor")
	return ret
}

// Deprecated
//
// gdk_cursor_new
//
// [ cursor_type ] trans: nothing
//
// [ result ] trans: everything
//
func NewCursor(cursor_type CursorTypeEnum) (result Cursor) {
	iv, err := _I.Get(16, "Cursor", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_cursor_type := gi.NewIntArgument(int(cursor_type))
	args := []gi.Argument{arg_cursor_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_cursor_new_for_display
//
// [ display ] trans: nothing
//
// [ cursor_type ] trans: nothing
//
// [ result ] trans: everything
//
func NewCursorForDisplay(display IDisplay, cursor_type CursorTypeEnum) (result Cursor) {
	iv, err := _I.Get(17, "Cursor", "new_for_display")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if display != nil {
		tmp = display.P_Display()
	}
	arg_display := gi.NewPointerArgument(tmp)
	arg_cursor_type := gi.NewIntArgument(int(cursor_type))
	args := []gi.Argument{arg_display, arg_cursor_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_cursor_new_from_name
//
// [ display ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func NewCursorFromName(display IDisplay, name string) (result Cursor) {
	iv, err := _I.Get(18, "Cursor", "new_from_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if display != nil {
		tmp = display.P_Display()
	}
	c_name := gi.CString(name)
	arg_display := gi.NewPointerArgument(tmp)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_display, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gdk_cursor_new_from_pixbuf
//
// [ display ] trans: nothing
//
// [ pixbuf ] trans: nothing
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ result ] trans: everything
//
func NewCursorFromPixbuf(display IDisplay, pixbuf gdkpixbuf.IPixbuf, x int32, y int32) (result Cursor) {
	iv, err := _I.Get(19, "Cursor", "new_from_pixbuf")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if display != nil {
		tmp = display.P_Display()
	}
	var tmp1 unsafe.Pointer
	if pixbuf != nil {
		tmp1 = pixbuf.P_Pixbuf()
	}
	arg_display := gi.NewPointerArgument(tmp)
	arg_pixbuf := gi.NewPointerArgument(tmp1)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	args := []gi.Argument{arg_display, arg_pixbuf, arg_x, arg_y}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_cursor_new_from_surface
//
// [ display ] trans: nothing
//
// [ surface ] trans: nothing
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ result ] trans: everything
//
func NewCursorFromSurface(display IDisplay, surface cairo.Surface, x float64, y float64) (result Cursor) {
	iv, err := _I.Get(20, "Cursor", "new_from_surface")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if display != nil {
		tmp = display.P_Display()
	}
	arg_display := gi.NewPointerArgument(tmp)
	arg_surface := gi.NewPointerArgument(surface.P)
	arg_x := gi.NewDoubleArgument(x)
	arg_y := gi.NewDoubleArgument(y)
	args := []gi.Argument{arg_display, arg_surface, arg_x, arg_y}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_cursor_get_cursor_type
//
// [ result ] trans: nothing
//
func (v Cursor) GetCursorType() (result CursorTypeEnum) {
	iv, err := _I.Get(21, "Cursor", "get_cursor_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = CursorTypeEnum(ret.Int())
	return
}

// gdk_cursor_get_display
//
// [ result ] trans: nothing
//
func (v Cursor) GetDisplay() (result Display) {
	iv, err := _I.Get(22, "Cursor", "get_display")
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

// gdk_cursor_get_image
//
// [ result ] trans: everything
//
func (v Cursor) GetImage() (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(23, "Cursor", "get_image")
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

// gdk_cursor_get_surface
//
// [ x_hot ] trans: everything, dir: out
//
// [ y_hot ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func (v Cursor) GetSurface() (result cairo.Surface, x_hot float64, y_hot float64) {
	iv, err := _I.Get(24, "Cursor", "get_surface")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_x_hot := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y_hot := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_x_hot, arg_y_hot}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	x_hot = outArgs[0].Double()
	y_hot = outArgs[1].Double()
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// gdk_cursor_ref
//
// [ result ] trans: everything
//
func (v Cursor) Ref() (result Cursor) {
	iv, err := _I.Get(25, "Cursor", "ref")
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

// Deprecated
//
// gdk_cursor_unref
//
func (v Cursor) Unref() {
	iv, err := _I.Get(26, "Cursor", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum CursorType
type CursorTypeEnum int

const (
	CursorTypeXCursor           CursorTypeEnum = 0
	CursorTypeArrow             CursorTypeEnum = 2
	CursorTypeBasedArrowDown    CursorTypeEnum = 4
	CursorTypeBasedArrowUp      CursorTypeEnum = 6
	CursorTypeBoat              CursorTypeEnum = 8
	CursorTypeBogosity          CursorTypeEnum = 10
	CursorTypeBottomLeftCorner  CursorTypeEnum = 12
	CursorTypeBottomRightCorner CursorTypeEnum = 14
	CursorTypeBottomSide        CursorTypeEnum = 16
	CursorTypeBottomTee         CursorTypeEnum = 18
	CursorTypeBoxSpiral         CursorTypeEnum = 20
	CursorTypeCenterPtr         CursorTypeEnum = 22
	CursorTypeCircle            CursorTypeEnum = 24
	CursorTypeClock             CursorTypeEnum = 26
	CursorTypeCoffeeMug         CursorTypeEnum = 28
	CursorTypeCross             CursorTypeEnum = 30
	CursorTypeCrossReverse      CursorTypeEnum = 32
	CursorTypeCrosshair         CursorTypeEnum = 34
	CursorTypeDiamondCross      CursorTypeEnum = 36
	CursorTypeDot               CursorTypeEnum = 38
	CursorTypeDotbox            CursorTypeEnum = 40
	CursorTypeDoubleArrow       CursorTypeEnum = 42
	CursorTypeDraftLarge        CursorTypeEnum = 44
	CursorTypeDraftSmall        CursorTypeEnum = 46
	CursorTypeDrapedBox         CursorTypeEnum = 48
	CursorTypeExchange          CursorTypeEnum = 50
	CursorTypeFleur             CursorTypeEnum = 52
	CursorTypeGobbler           CursorTypeEnum = 54
	CursorTypeGumby             CursorTypeEnum = 56
	CursorTypeHand1             CursorTypeEnum = 58
	CursorTypeHand2             CursorTypeEnum = 60
	CursorTypeHeart             CursorTypeEnum = 62
	CursorTypeIcon              CursorTypeEnum = 64
	CursorTypeIronCross         CursorTypeEnum = 66
	CursorTypeLeftPtr           CursorTypeEnum = 68
	CursorTypeLeftSide          CursorTypeEnum = 70
	CursorTypeLeftTee           CursorTypeEnum = 72
	CursorTypeLeftbutton        CursorTypeEnum = 74
	CursorTypeLlAngle           CursorTypeEnum = 76
	CursorTypeLrAngle           CursorTypeEnum = 78
	CursorTypeMan               CursorTypeEnum = 80
	CursorTypeMiddlebutton      CursorTypeEnum = 82
	CursorTypeMouse             CursorTypeEnum = 84
	CursorTypePencil            CursorTypeEnum = 86
	CursorTypePirate            CursorTypeEnum = 88
	CursorTypePlus              CursorTypeEnum = 90
	CursorTypeQuestionArrow     CursorTypeEnum = 92
	CursorTypeRightPtr          CursorTypeEnum = 94
	CursorTypeRightSide         CursorTypeEnum = 96
	CursorTypeRightTee          CursorTypeEnum = 98
	CursorTypeRightbutton       CursorTypeEnum = 100
	CursorTypeRtlLogo           CursorTypeEnum = 102
	CursorTypeSailboat          CursorTypeEnum = 104
	CursorTypeSbDownArrow       CursorTypeEnum = 106
	CursorTypeSbHDoubleArrow    CursorTypeEnum = 108
	CursorTypeSbLeftArrow       CursorTypeEnum = 110
	CursorTypeSbRightArrow      CursorTypeEnum = 112
	CursorTypeSbUpArrow         CursorTypeEnum = 114
	CursorTypeSbVDoubleArrow    CursorTypeEnum = 116
	CursorTypeShuttle           CursorTypeEnum = 118
	CursorTypeSizing            CursorTypeEnum = 120
	CursorTypeSpider            CursorTypeEnum = 122
	CursorTypeSpraycan          CursorTypeEnum = 124
	CursorTypeStar              CursorTypeEnum = 126
	CursorTypeTarget            CursorTypeEnum = 128
	CursorTypeTcross            CursorTypeEnum = 130
	CursorTypeTopLeftArrow      CursorTypeEnum = 132
	CursorTypeTopLeftCorner     CursorTypeEnum = 134
	CursorTypeTopRightCorner    CursorTypeEnum = 136
	CursorTypeTopSide           CursorTypeEnum = 138
	CursorTypeTopTee            CursorTypeEnum = 140
	CursorTypeTrek              CursorTypeEnum = 142
	CursorTypeUlAngle           CursorTypeEnum = 144
	CursorTypeUmbrella          CursorTypeEnum = 146
	CursorTypeUrAngle           CursorTypeEnum = 148
	CursorTypeWatch             CursorTypeEnum = 150
	CursorTypeXterm             CursorTypeEnum = 152
	CursorTypeLastCursor        CursorTypeEnum = 153
	CursorTypeBlankCursor       CursorTypeEnum = -2
	CursorTypeCursorIsPixmap    CursorTypeEnum = -1
)

func CursorTypeGetType() gi.GType {
	ret := _I.GetGType(9, "CursorType")
	return ret
}

// Object Device
type Device struct {
	g.Object
}

func WrapDevice(p unsafe.Pointer) (r Device) { r.P = p; return }

type IDevice interface{ P_Device() unsafe.Pointer }

func (v Device) P_Device() unsafe.Pointer { return v.P }
func DeviceGetType() gi.GType {
	ret := _I.GetGType(10, "Device")
	return ret
}

// Deprecated
//
// gdk_device_grab_info_libgtk_only
//
// [ display ] trans: nothing
//
// [ device ] trans: nothing
//
// [ grab_window ] trans: nothing, dir: out
//
// [ owner_events ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func DeviceGrabInfoLibgtkOnly1(display IDisplay, device IDevice) (result bool, grab_window Window, owner_events bool) {
	iv, err := _I.Get(27, "Device", "grab_info_libgtk_only")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if display != nil {
		tmp = display.P_Display()
	}
	var tmp1 unsafe.Pointer
	if device != nil {
		tmp1 = device.P_Device()
	}
	arg_display := gi.NewPointerArgument(tmp)
	arg_device := gi.NewPointerArgument(tmp1)
	arg_grab_window := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_owner_events := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_display, arg_device, arg_grab_window, arg_owner_events}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	grab_window.P = outArgs[0].Pointer()
	owner_events = outArgs[1].Bool()
	result = ret.Bool()
	return
}

// gdk_device_get_associated_device
//
// [ result ] trans: nothing
//
func (v Device) GetAssociatedDevice() (result Device) {
	iv, err := _I.Get(28, "Device", "get_associated_device")
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

// gdk_device_get_axes
//
// [ result ] trans: nothing
//
func (v Device) GetAxes() (result AxisFlags) {
	iv, err := _I.Get(29, "Device", "get_axes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = AxisFlags(ret.Int())
	return
}

// gdk_device_get_axis_use
//
// [ index_ ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) GetAxisUse(index_ uint32) (result AxisUseEnum) {
	iv, err := _I.Get(30, "Device", "get_axis_use")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewUint32Argument(index_)
	args := []gi.Argument{arg_v, arg_index_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = AxisUseEnum(ret.Int())
	return
}

// gdk_device_get_device_type
//
// [ result ] trans: nothing
//
func (v Device) GetDeviceType() (result DeviceTypeEnum) {
	iv, err := _I.Get(31, "Device", "get_device_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = DeviceTypeEnum(ret.Int())
	return
}

// gdk_device_get_display
//
// [ result ] trans: nothing
//
func (v Device) GetDisplay() (result Display) {
	iv, err := _I.Get(32, "Device", "get_display")
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

// gdk_device_get_has_cursor
//
// [ result ] trans: nothing
//
func (v Device) GetHasCursor() (result bool) {
	iv, err := _I.Get(33, "Device", "get_has_cursor")
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

// gdk_device_get_key
//
// [ index_ ] trans: nothing
//
// [ keyval ] trans: everything, dir: out
//
// [ modifiers ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Device) GetKey(index_ uint32) (result bool, keyval uint32, modifiers ModifierTypeFlags) {
	iv, err := _I.Get(34, "Device", "get_key")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewUint32Argument(index_)
	arg_keyval := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_modifiers := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_index_, arg_keyval, arg_modifiers}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	keyval = outArgs[0].Uint32()
	modifiers = ModifierTypeFlags(outArgs[1].Int())
	result = ret.Bool()
	return
}

// gdk_device_get_last_event_window
//
// [ result ] trans: nothing
//
func (v Device) GetLastEventWindow() (result Window) {
	iv, err := _I.Get(35, "Device", "get_last_event_window")
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

// gdk_device_get_mode
//
// [ result ] trans: nothing
//
func (v Device) GetMode() (result InputModeEnum) {
	iv, err := _I.Get(36, "Device", "get_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = InputModeEnum(ret.Int())
	return
}

// gdk_device_get_n_axes
//
// [ result ] trans: nothing
//
func (v Device) GetNAxes() (result int32) {
	iv, err := _I.Get(37, "Device", "get_n_axes")
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

// gdk_device_get_n_keys
//
// [ result ] trans: nothing
//
func (v Device) GetNKeys() (result int32) {
	iv, err := _I.Get(38, "Device", "get_n_keys")
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

// gdk_device_get_name
//
// [ result ] trans: nothing
//
func (v Device) GetName() (result string) {
	iv, err := _I.Get(39, "Device", "get_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gdk_device_get_position
//
// [ screen ] trans: nothing, dir: out
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
func (v Device) GetPosition() (screen Screen, x int32, y int32) {
	iv, err := _I.Get(40, "Device", "get_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_screen := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_screen, arg_x, arg_y}
	iv.Call(args, nil, &outArgs[0])
	screen.P = outArgs[0].Pointer()
	x = outArgs[1].Int32()
	y = outArgs[2].Int32()
	return
}

// gdk_device_get_position_double
//
// [ screen ] trans: nothing, dir: out
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
func (v Device) GetPositionDouble() (screen Screen, x float64, y float64) {
	iv, err := _I.Get(41, "Device", "get_position_double")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_screen := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_screen, arg_x, arg_y}
	iv.Call(args, nil, &outArgs[0])
	screen.P = outArgs[0].Pointer()
	x = outArgs[1].Double()
	y = outArgs[2].Double()
	return
}

// gdk_device_get_product_id
//
// [ result ] trans: nothing
//
func (v Device) GetProductId() (result string) {
	iv, err := _I.Get(42, "Device", "get_product_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gdk_device_get_seat
//
// [ result ] trans: nothing
//
func (v Device) GetSeat() (result Seat) {
	iv, err := _I.Get(43, "Device", "get_seat")
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

// gdk_device_get_source
//
// [ result ] trans: nothing
//
func (v Device) GetSource() (result InputSourceEnum) {
	iv, err := _I.Get(44, "Device", "get_source")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = InputSourceEnum(ret.Int())
	return
}

// gdk_device_get_vendor_id
//
// [ result ] trans: nothing
//
func (v Device) GetVendorId() (result string) {
	iv, err := _I.Get(45, "Device", "get_vendor_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gdk_device_get_window_at_position
//
// [ win_x ] trans: everything, dir: out
//
// [ win_y ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Device) GetWindowAtPosition() (result Window, win_x int32, win_y int32) {
	iv, err := _I.Get(46, "Device", "get_window_at_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_win_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_win_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_win_x, arg_win_y}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	win_x = outArgs[0].Int32()
	win_y = outArgs[1].Int32()
	result.P = ret.Pointer()
	return
}

// gdk_device_get_window_at_position_double
//
// [ win_x ] trans: everything, dir: out
//
// [ win_y ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Device) GetWindowAtPositionDouble() (result Window, win_x float64, win_y float64) {
	iv, err := _I.Get(47, "Device", "get_window_at_position_double")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_win_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_win_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_win_x, arg_win_y}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	win_x = outArgs[0].Double()
	win_y = outArgs[1].Double()
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// gdk_device_grab
//
// [ window ] trans: nothing
//
// [ grab_ownership ] trans: nothing
//
// [ owner_events ] trans: nothing
//
// [ event_mask ] trans: nothing
//
// [ cursor ] trans: nothing
//
// [ time_ ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) Grab(window IWindow, grab_ownership GrabOwnershipEnum, owner_events bool, event_mask EventMaskFlags, cursor ICursor, time_ uint32) (result GrabStatusEnum) {
	iv, err := _I.Get(48, "Device", "grab")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	var tmp1 unsafe.Pointer
	if cursor != nil {
		tmp1 = cursor.P_Cursor()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_window := gi.NewPointerArgument(tmp)
	arg_grab_ownership := gi.NewIntArgument(int(grab_ownership))
	arg_owner_events := gi.NewBoolArgument(owner_events)
	arg_event_mask := gi.NewIntArgument(int(event_mask))
	arg_cursor := gi.NewPointerArgument(tmp1)
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_v, arg_window, arg_grab_ownership, arg_owner_events, arg_event_mask, arg_cursor, arg_time_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = GrabStatusEnum(ret.Int())
	return
}

// gdk_device_list_axes
//
// [ result ] trans: container
//
func (v Device) ListAxes() (result g.List) {
	iv, err := _I.Get(49, "Device", "list_axes")
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

// gdk_device_list_slave_devices
//
// [ result ] trans: container
//
func (v Device) ListSlaveDevices() (result g.List) {
	iv, err := _I.Get(50, "Device", "list_slave_devices")
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

// gdk_device_set_axis_use
//
// [ index_ ] trans: nothing
//
// [ use ] trans: nothing
//
func (v Device) SetAxisUse(index_ uint32, use AxisUseEnum) {
	iv, err := _I.Get(51, "Device", "set_axis_use")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewUint32Argument(index_)
	arg_use := gi.NewIntArgument(int(use))
	args := []gi.Argument{arg_v, arg_index_, arg_use}
	iv.Call(args, nil, nil)
}

// gdk_device_set_key
//
// [ index_ ] trans: nothing
//
// [ keyval ] trans: nothing
//
// [ modifiers ] trans: nothing
//
func (v Device) SetKey(index_ uint32, keyval uint32, modifiers ModifierTypeFlags) {
	iv, err := _I.Get(52, "Device", "set_key")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewUint32Argument(index_)
	arg_keyval := gi.NewUint32Argument(keyval)
	arg_modifiers := gi.NewIntArgument(int(modifiers))
	args := []gi.Argument{arg_v, arg_index_, arg_keyval, arg_modifiers}
	iv.Call(args, nil, nil)
}

// gdk_device_set_mode
//
// [ mode ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) SetMode(mode InputModeEnum) (result bool) {
	iv, err := _I.Get(53, "Device", "set_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_v, arg_mode}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Deprecated
//
// gdk_device_ungrab
//
// [ time_ ] trans: nothing
//
func (v Device) Ungrab(time_ uint32) {
	iv, err := _I.Get(54, "Device", "ungrab")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_v, arg_time_}
	iv.Call(args, nil, nil)
}

// gdk_device_warp
//
// [ screen ] trans: nothing
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
func (v Device) Warp(screen IScreen, x int32, y int32) {
	iv, err := _I.Get(55, "Device", "warp")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if screen != nil {
		tmp = screen.P_Screen()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_screen := gi.NewPointerArgument(tmp)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	args := []gi.Argument{arg_v, arg_screen, arg_x, arg_y}
	iv.Call(args, nil, nil)
}

// Object DeviceManager
type DeviceManager struct {
	g.Object
}

func WrapDeviceManager(p unsafe.Pointer) (r DeviceManager) { r.P = p; return }

type IDeviceManager interface{ P_DeviceManager() unsafe.Pointer }

func (v DeviceManager) P_DeviceManager() unsafe.Pointer { return v.P }
func DeviceManagerGetType() gi.GType {
	ret := _I.GetGType(11, "DeviceManager")
	return ret
}

// Deprecated
//
// gdk_device_manager_get_client_pointer
//
// [ result ] trans: nothing
//
func (v DeviceManager) GetClientPointer() (result Device) {
	iv, err := _I.Get(56, "DeviceManager", "get_client_pointer")
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

// gdk_device_manager_get_display
//
// [ result ] trans: nothing
//
func (v DeviceManager) GetDisplay() (result Display) {
	iv, err := _I.Get(57, "DeviceManager", "get_display")
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

// Deprecated
//
// gdk_device_manager_list_devices
//
// [ type1 ] trans: nothing
//
// [ result ] trans: container
//
func (v DeviceManager) ListDevices(type1 DeviceTypeEnum) (result g.List) {
	iv, err := _I.Get(58, "DeviceManager", "list_devices")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Interface DevicePad
type DevicePad struct {
	DevicePadIfc
	P unsafe.Pointer
}
type DevicePadIfc struct{}
type IDevicePad interface{ P_DevicePad() unsafe.Pointer }

func (v DevicePad) P_DevicePad() unsafe.Pointer { return v.P }
func DevicePadGetType() gi.GType {
	ret := _I.GetGType(12, "DevicePad")
	return ret
}

// gdk_device_pad_get_feature_group
//
// [ feature ] trans: nothing
//
// [ feature_idx ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DevicePadIfc) GetFeatureGroup(feature DevicePadFeatureEnum, feature_idx int32) (result int32) {
	iv, err := _I.Get(59, "DevicePad", "get_feature_group")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_feature := gi.NewIntArgument(int(feature))
	arg_feature_idx := gi.NewInt32Argument(feature_idx)
	args := []gi.Argument{arg_v, arg_feature, arg_feature_idx}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gdk_device_pad_get_group_n_modes
//
// [ group_idx ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DevicePadIfc) GetGroupNModes(group_idx int32) (result int32) {
	iv, err := _I.Get(60, "DevicePad", "get_group_n_modes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_group_idx := gi.NewInt32Argument(group_idx)
	args := []gi.Argument{arg_v, arg_group_idx}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gdk_device_pad_get_n_features
//
// [ feature ] trans: nothing
//
// [ result ] trans: nothing
//
func (v *DevicePadIfc) GetNFeatures(feature DevicePadFeatureEnum) (result int32) {
	iv, err := _I.Get(61, "DevicePad", "get_n_features")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_feature := gi.NewIntArgument(int(feature))
	args := []gi.Argument{arg_v, arg_feature}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gdk_device_pad_get_n_groups
//
// [ result ] trans: nothing
//
func (v *DevicePadIfc) GetNGroups() (result int32) {
	iv, err := _I.Get(62, "DevicePad", "get_n_groups")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// Enum DevicePadFeature
type DevicePadFeatureEnum int

const (
	DevicePadFeatureButton DevicePadFeatureEnum = 0
	DevicePadFeatureRing   DevicePadFeatureEnum = 1
	DevicePadFeatureStrip  DevicePadFeatureEnum = 2
)

func DevicePadFeatureGetType() gi.GType {
	ret := _I.GetGType(13, "DevicePadFeature")
	return ret
}

// ignore GType struct DevicePadInterface

// Object DeviceTool
type DeviceTool struct {
	g.Object
}

func WrapDeviceTool(p unsafe.Pointer) (r DeviceTool) { r.P = p; return }

type IDeviceTool interface{ P_DeviceTool() unsafe.Pointer }

func (v DeviceTool) P_DeviceTool() unsafe.Pointer { return v.P }
func DeviceToolGetType() gi.GType {
	ret := _I.GetGType(14, "DeviceTool")
	return ret
}

// gdk_device_tool_get_hardware_id
//
// [ result ] trans: nothing
//
func (v DeviceTool) GetHardwareId() (result uint64) {
	iv, err := _I.Get(63, "DeviceTool", "get_hardware_id")
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

// gdk_device_tool_get_serial
//
// [ result ] trans: nothing
//
func (v DeviceTool) GetSerial() (result uint64) {
	iv, err := _I.Get(64, "DeviceTool", "get_serial")
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

// gdk_device_tool_get_tool_type
//
// [ result ] trans: nothing
//
func (v DeviceTool) GetToolType() (result DeviceToolTypeEnum) {
	iv, err := _I.Get(65, "DeviceTool", "get_tool_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = DeviceToolTypeEnum(ret.Int())
	return
}

// Enum DeviceToolType
type DeviceToolTypeEnum int

const (
	DeviceToolTypeUnknown  DeviceToolTypeEnum = 0
	DeviceToolTypePen      DeviceToolTypeEnum = 1
	DeviceToolTypeEraser   DeviceToolTypeEnum = 2
	DeviceToolTypeBrush    DeviceToolTypeEnum = 3
	DeviceToolTypePencil   DeviceToolTypeEnum = 4
	DeviceToolTypeAirbrush DeviceToolTypeEnum = 5
	DeviceToolTypeMouse    DeviceToolTypeEnum = 6
	DeviceToolTypeLens     DeviceToolTypeEnum = 7
)

func DeviceToolTypeGetType() gi.GType {
	ret := _I.GetGType(15, "DeviceToolType")
	return ret
}

// Enum DeviceType
type DeviceTypeEnum int

const (
	DeviceTypeMaster   DeviceTypeEnum = 0
	DeviceTypeSlave    DeviceTypeEnum = 1
	DeviceTypeFloating DeviceTypeEnum = 2
)

func DeviceTypeGetType() gi.GType {
	ret := _I.GetGType(16, "DeviceType")
	return ret
}

// Object Display
type Display struct {
	g.Object
}

func WrapDisplay(p unsafe.Pointer) (r Display) { r.P = p; return }

type IDisplay interface{ P_Display() unsafe.Pointer }

func (v Display) P_Display() unsafe.Pointer { return v.P }
func DisplayGetType() gi.GType {
	ret := _I.GetGType(17, "Display")
	return ret
}

// gdk_display_open
//
// [ display_name ] trans: nothing
//
// [ result ] trans: nothing
//
func DisplayOpen1(display_name string) (result Display) {
	iv, err := _I.Get(67, "Display", "open")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_display_name := gi.CString(display_name)
	arg_display_name := gi.NewStringArgument(c_display_name)
	args := []gi.Argument{arg_display_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_display_name)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// gdk_display_beep
//
func (v Display) Beep() {
	iv, err := _I.Get(69, "Display", "beep")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_display_close
//
func (v Display) Close() {
	iv, err := _I.Get(70, "Display", "close")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_display_device_is_grabbed
//
// [ device ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Display) DeviceIsGrabbed(device IDevice) (result bool) {
	iv, err := _I.Get(71, "Display", "device_is_grabbed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if device != nil {
		tmp = device.P_Device()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_device := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_device}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_display_flush
//
func (v Display) Flush() {
	iv, err := _I.Get(72, "Display", "flush")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_display_get_app_launch_context
//
// [ result ] trans: everything
//
func (v Display) GetAppLaunchContext() (result AppLaunchContext) {
	iv, err := _I.Get(73, "Display", "get_app_launch_context")
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

// gdk_display_get_default_cursor_size
//
// [ result ] trans: nothing
//
func (v Display) GetDefaultCursorSize() (result uint32) {
	iv, err := _I.Get(74, "Display", "get_default_cursor_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gdk_display_get_default_group
//
// [ result ] trans: nothing
//
func (v Display) GetDefaultGroup() (result Window) {
	iv, err := _I.Get(75, "Display", "get_default_group")
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

// gdk_display_get_default_screen
//
// [ result ] trans: nothing
//
func (v Display) GetDefaultScreen() (result Screen) {
	iv, err := _I.Get(76, "Display", "get_default_screen")
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

// gdk_display_get_default_seat
//
// [ result ] trans: nothing
//
func (v Display) GetDefaultSeat() (result Seat) {
	iv, err := _I.Get(77, "Display", "get_default_seat")
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

// Deprecated
//
// gdk_display_get_device_manager
//
// [ result ] trans: nothing
//
func (v Display) GetDeviceManager() (result DeviceManager) {
	iv, err := _I.Get(78, "Display", "get_device_manager")
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

// gdk_display_get_event
//
// [ result ] trans: everything
//
func (v Display) GetEvent() (result Event) {
	iv, err := _I.Get(79, "Display", "get_event")
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

// gdk_display_get_maximal_cursor_size
//
// [ width ] trans: everything, dir: out
//
// [ height ] trans: everything, dir: out
//
func (v Display) GetMaximalCursorSize() (width uint32, height uint32) {
	iv, err := _I.Get(80, "Display", "get_maximal_cursor_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_width, arg_height}
	iv.Call(args, nil, &outArgs[0])
	width = outArgs[0].Uint32()
	height = outArgs[1].Uint32()
	return
}

// gdk_display_get_monitor
//
// [ monitor_num ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Display) GetMonitor(monitor_num int32) (result Monitor) {
	iv, err := _I.Get(81, "Display", "get_monitor")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_monitor_num := gi.NewInt32Argument(monitor_num)
	args := []gi.Argument{arg_v, arg_monitor_num}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_display_get_monitor_at_point
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Display) GetMonitorAtPoint(x int32, y int32) (result Monitor) {
	iv, err := _I.Get(82, "Display", "get_monitor_at_point")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	args := []gi.Argument{arg_v, arg_x, arg_y}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_display_get_monitor_at_window
//
// [ window ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Display) GetMonitorAtWindow(window IWindow) (result Monitor) {
	iv, err := _I.Get(83, "Display", "get_monitor_at_window")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_window := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_window}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_display_get_n_monitors
//
// [ result ] trans: nothing
//
func (v Display) GetNMonitors() (result int32) {
	iv, err := _I.Get(84, "Display", "get_n_monitors")
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

// Deprecated
//
// gdk_display_get_n_screens
//
// [ result ] trans: nothing
//
func (v Display) GetNScreens() (result int32) {
	iv, err := _I.Get(85, "Display", "get_n_screens")
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

// gdk_display_get_name
//
// [ result ] trans: nothing
//
func (v Display) GetName() (result string) {
	iv, err := _I.Get(86, "Display", "get_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// Deprecated
//
// gdk_display_get_pointer
//
// [ screen ] trans: nothing, dir: out
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
// [ mask ] trans: everything, dir: out
//
func (v Display) GetPointer() (screen Screen, x int32, y int32, mask ModifierTypeFlags) {
	iv, err := _I.Get(87, "Display", "get_pointer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_screen := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_mask := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_screen, arg_x, arg_y, arg_mask}
	iv.Call(args, nil, &outArgs[0])
	screen.P = outArgs[0].Pointer()
	x = outArgs[1].Int32()
	y = outArgs[2].Int32()
	mask = ModifierTypeFlags(outArgs[3].Int())
	return
}

// gdk_display_get_primary_monitor
//
// [ result ] trans: nothing
//
func (v Display) GetPrimaryMonitor() (result Monitor) {
	iv, err := _I.Get(88, "Display", "get_primary_monitor")
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

// Deprecated
//
// gdk_display_get_screen
//
// [ screen_num ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Display) GetScreen(screen_num int32) (result Screen) {
	iv, err := _I.Get(89, "Display", "get_screen")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_screen_num := gi.NewInt32Argument(screen_num)
	args := []gi.Argument{arg_v, arg_screen_num}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// gdk_display_get_window_at_pointer
//
// [ win_x ] trans: everything, dir: out
//
// [ win_y ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Display) GetWindowAtPointer() (result Window, win_x int32, win_y int32) {
	iv, err := _I.Get(90, "Display", "get_window_at_pointer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_win_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_win_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_win_x, arg_win_y}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	win_x = outArgs[0].Int32()
	win_y = outArgs[1].Int32()
	result.P = ret.Pointer()
	return
}

// gdk_display_has_pending
//
// [ result ] trans: nothing
//
func (v Display) HasPending() (result bool) {
	iv, err := _I.Get(91, "Display", "has_pending")
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

// gdk_display_is_closed
//
// [ result ] trans: nothing
//
func (v Display) IsClosed() (result bool) {
	iv, err := _I.Get(92, "Display", "is_closed")
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

// Deprecated
//
// gdk_display_keyboard_ungrab
//
// [ time_ ] trans: nothing
//
func (v Display) KeyboardUngrab(time_ uint32) {
	iv, err := _I.Get(93, "Display", "keyboard_ungrab")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_v, arg_time_}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_display_list_devices
//
// [ result ] trans: nothing
//
func (v Display) ListDevices() (result g.List) {
	iv, err := _I.Get(94, "Display", "list_devices")
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

// gdk_display_list_seats
//
// [ result ] trans: container
//
func (v Display) ListSeats() (result g.List) {
	iv, err := _I.Get(95, "Display", "list_seats")
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

// gdk_display_notify_startup_complete
//
// [ startup_id ] trans: nothing
//
func (v Display) NotifyStartupComplete(startup_id string) {
	iv, err := _I.Get(96, "Display", "notify_startup_complete")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_startup_id := gi.CString(startup_id)
	arg_v := gi.NewPointerArgument(v.P)
	arg_startup_id := gi.NewStringArgument(c_startup_id)
	args := []gi.Argument{arg_v, arg_startup_id}
	iv.Call(args, nil, nil)
	gi.Free(c_startup_id)
}

// gdk_display_peek_event
//
// [ result ] trans: everything
//
func (v Display) PeekEvent() (result Event) {
	iv, err := _I.Get(97, "Display", "peek_event")
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

// Deprecated
//
// gdk_display_pointer_is_grabbed
//
// [ result ] trans: nothing
//
func (v Display) PointerIsGrabbed() (result bool) {
	iv, err := _I.Get(98, "Display", "pointer_is_grabbed")
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

// Deprecated
//
// gdk_display_pointer_ungrab
//
// [ time_ ] trans: nothing
//
func (v Display) PointerUngrab(time_ uint32) {
	iv, err := _I.Get(99, "Display", "pointer_ungrab")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_v, arg_time_}
	iv.Call(args, nil, nil)
}

// gdk_display_put_event
//
// [ event ] trans: nothing
//
func (v Display) PutEvent(event Event) {
	iv, err := _I.Get(100, "Display", "put_event")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_event := gi.NewPointerArgument(event.P)
	args := []gi.Argument{arg_v, arg_event}
	iv.Call(args, nil, nil)
}

// gdk_display_request_selection_notification
//
// [ selection ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Display) RequestSelectionNotification(selection Atom) (result bool) {
	iv, err := _I.Get(101, "Display", "request_selection_notification")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_selection := gi.NewPointerArgument(selection.P)
	args := []gi.Argument{arg_v, arg_selection}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_display_set_double_click_distance
//
// [ distance ] trans: nothing
//
func (v Display) SetDoubleClickDistance(distance uint32) {
	iv, err := _I.Get(102, "Display", "set_double_click_distance")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_distance := gi.NewUint32Argument(distance)
	args := []gi.Argument{arg_v, arg_distance}
	iv.Call(args, nil, nil)
}

// gdk_display_set_double_click_time
//
// [ msec ] trans: nothing
//
func (v Display) SetDoubleClickTime(msec uint32) {
	iv, err := _I.Get(103, "Display", "set_double_click_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_msec := gi.NewUint32Argument(msec)
	args := []gi.Argument{arg_v, arg_msec}
	iv.Call(args, nil, nil)
}

// gdk_display_store_clipboard
//
// [ clipboard_window ] trans: nothing
//
// [ time_ ] trans: nothing
//
// [ targets ] trans: nothing
//
// [ n_targets ] trans: nothing
//
func (v Display) StoreClipboard(clipboard_window IWindow, time_ uint32, targets gi.PointerArray, n_targets int32) {
	iv, err := _I.Get(104, "Display", "store_clipboard")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if clipboard_window != nil {
		tmp = clipboard_window.P_Window()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_clipboard_window := gi.NewPointerArgument(tmp)
	arg_time_ := gi.NewUint32Argument(time_)
	arg_targets := gi.NewPointerArgument(targets.P)
	arg_n_targets := gi.NewInt32Argument(n_targets)
	args := []gi.Argument{arg_v, arg_clipboard_window, arg_time_, arg_targets, arg_n_targets}
	iv.Call(args, nil, nil)
}

// gdk_display_supports_clipboard_persistence
//
// [ result ] trans: nothing
//
func (v Display) SupportsClipboardPersistence() (result bool) {
	iv, err := _I.Get(105, "Display", "supports_clipboard_persistence")
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

// Deprecated
//
// gdk_display_supports_composite
//
// [ result ] trans: nothing
//
func (v Display) SupportsComposite() (result bool) {
	iv, err := _I.Get(106, "Display", "supports_composite")
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

// gdk_display_supports_cursor_alpha
//
// [ result ] trans: nothing
//
func (v Display) SupportsCursorAlpha() (result bool) {
	iv, err := _I.Get(107, "Display", "supports_cursor_alpha")
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

// gdk_display_supports_cursor_color
//
// [ result ] trans: nothing
//
func (v Display) SupportsCursorColor() (result bool) {
	iv, err := _I.Get(108, "Display", "supports_cursor_color")
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

// gdk_display_supports_input_shapes
//
// [ result ] trans: nothing
//
func (v Display) SupportsInputShapes() (result bool) {
	iv, err := _I.Get(109, "Display", "supports_input_shapes")
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

// gdk_display_supports_selection_notification
//
// [ result ] trans: nothing
//
func (v Display) SupportsSelectionNotification() (result bool) {
	iv, err := _I.Get(110, "Display", "supports_selection_notification")
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

// gdk_display_supports_shapes
//
// [ result ] trans: nothing
//
func (v Display) SupportsShapes() (result bool) {
	iv, err := _I.Get(111, "Display", "supports_shapes")
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

// gdk_display_sync
//
func (v Display) Sync() {
	iv, err := _I.Get(112, "Display", "sync")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_display_warp_pointer
//
// [ screen ] trans: nothing
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
func (v Display) WarpPointer(screen IScreen, x int32, y int32) {
	iv, err := _I.Get(113, "Display", "warp_pointer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if screen != nil {
		tmp = screen.P_Screen()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_screen := gi.NewPointerArgument(tmp)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	args := []gi.Argument{arg_v, arg_screen, arg_x, arg_y}
	iv.Call(args, nil, nil)
}

// Object DisplayManager
type DisplayManager struct {
	g.Object
}

func WrapDisplayManager(p unsafe.Pointer) (r DisplayManager) { r.P = p; return }

type IDisplayManager interface{ P_DisplayManager() unsafe.Pointer }

func (v DisplayManager) P_DisplayManager() unsafe.Pointer { return v.P }
func DisplayManagerGetType() gi.GType {
	ret := _I.GetGType(18, "DisplayManager")
	return ret
}

// gdk_display_manager_get_default_display
//
// [ result ] trans: nothing
//
func (v DisplayManager) GetDefaultDisplay() (result Display) {
	iv, err := _I.Get(115, "DisplayManager", "get_default_display")
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

// gdk_display_manager_list_displays
//
// [ result ] trans: container
//
func (v DisplayManager) ListDisplays() (result g.SList) {
	iv, err := _I.Get(116, "DisplayManager", "list_displays")
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

// gdk_display_manager_open_display
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v DisplayManager) OpenDisplay(name string) (result Display) {
	iv, err := _I.Get(117, "DisplayManager", "open_display")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_v, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gdk_display_manager_set_default_display
//
// [ display ] trans: nothing
//
func (v DisplayManager) SetDefaultDisplay(display IDisplay) {
	iv, err := _I.Get(118, "DisplayManager", "set_default_display")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if display != nil {
		tmp = display.P_Display()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_display := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_display}
	iv.Call(args, nil, nil)
}

// Flags DragAction
type DragActionFlags int

const (
	DragActionDefault DragActionFlags = 1
	DragActionCopy    DragActionFlags = 2
	DragActionMove    DragActionFlags = 4
	DragActionLink    DragActionFlags = 8
	DragActionPrivate DragActionFlags = 16
	DragActionAsk     DragActionFlags = 32
)

func DragActionGetType() gi.GType {
	ret := _I.GetGType(19, "DragAction")
	return ret
}

// Enum DragCancelReason
type DragCancelReasonEnum int

const (
	DragCancelReasonNoTarget      DragCancelReasonEnum = 0
	DragCancelReasonUserCancelled DragCancelReasonEnum = 1
	DragCancelReasonError         DragCancelReasonEnum = 2
)

func DragCancelReasonGetType() gi.GType {
	ret := _I.GetGType(20, "DragCancelReason")
	return ret
}

// Object DragContext
type DragContext struct {
	g.Object
}

func WrapDragContext(p unsafe.Pointer) (r DragContext) { r.P = p; return }

type IDragContext interface{ P_DragContext() unsafe.Pointer }

func (v DragContext) P_DragContext() unsafe.Pointer { return v.P }
func DragContextGetType() gi.GType {
	ret := _I.GetGType(21, "DragContext")
	return ret
}

// gdk_drag_context_get_actions
//
// [ result ] trans: nothing
//
func (v DragContext) GetActions() (result DragActionFlags) {
	iv, err := _I.Get(119, "DragContext", "get_actions")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = DragActionFlags(ret.Int())
	return
}

// gdk_drag_context_get_dest_window
//
// [ result ] trans: nothing
//
func (v DragContext) GetDestWindow() (result Window) {
	iv, err := _I.Get(120, "DragContext", "get_dest_window")
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

// gdk_drag_context_get_device
//
// [ result ] trans: nothing
//
func (v DragContext) GetDevice() (result Device) {
	iv, err := _I.Get(121, "DragContext", "get_device")
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

// gdk_drag_context_get_drag_window
//
// [ result ] trans: nothing
//
func (v DragContext) GetDragWindow() (result Window) {
	iv, err := _I.Get(122, "DragContext", "get_drag_window")
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

// gdk_drag_context_get_protocol
//
// [ result ] trans: nothing
//
func (v DragContext) GetProtocol() (result DragProtocolEnum) {
	iv, err := _I.Get(123, "DragContext", "get_protocol")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = DragProtocolEnum(ret.Int())
	return
}

// gdk_drag_context_get_selected_action
//
// [ result ] trans: nothing
//
func (v DragContext) GetSelectedAction() (result DragActionFlags) {
	iv, err := _I.Get(124, "DragContext", "get_selected_action")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = DragActionFlags(ret.Int())
	return
}

// gdk_drag_context_get_source_window
//
// [ result ] trans: nothing
//
func (v DragContext) GetSourceWindow() (result Window) {
	iv, err := _I.Get(125, "DragContext", "get_source_window")
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

// gdk_drag_context_get_suggested_action
//
// [ result ] trans: nothing
//
func (v DragContext) GetSuggestedAction() (result DragActionFlags) {
	iv, err := _I.Get(126, "DragContext", "get_suggested_action")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = DragActionFlags(ret.Int())
	return
}

// gdk_drag_context_list_targets
//
// [ result ] trans: nothing
//
func (v DragContext) ListTargets() (result g.List) {
	iv, err := _I.Get(127, "DragContext", "list_targets")
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

// gdk_drag_context_manage_dnd
//
// [ ipc_window ] trans: nothing
//
// [ actions ] trans: nothing
//
// [ result ] trans: nothing
//
func (v DragContext) ManageDnd(ipc_window IWindow, actions DragActionFlags) (result bool) {
	iv, err := _I.Get(128, "DragContext", "manage_dnd")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if ipc_window != nil {
		tmp = ipc_window.P_Window()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ipc_window := gi.NewPointerArgument(tmp)
	arg_actions := gi.NewIntArgument(int(actions))
	args := []gi.Argument{arg_v, arg_ipc_window, arg_actions}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_drag_context_set_device
//
// [ device ] trans: nothing
//
func (v DragContext) SetDevice(device IDevice) {
	iv, err := _I.Get(129, "DragContext", "set_device")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if device != nil {
		tmp = device.P_Device()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_device := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_device}
	iv.Call(args, nil, nil)
}

// gdk_drag_context_set_hotspot
//
// [ hot_x ] trans: nothing
//
// [ hot_y ] trans: nothing
//
func (v DragContext) SetHotspot(hot_x int32, hot_y int32) {
	iv, err := _I.Get(130, "DragContext", "set_hotspot")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_hot_x := gi.NewInt32Argument(hot_x)
	arg_hot_y := gi.NewInt32Argument(hot_y)
	args := []gi.Argument{arg_v, arg_hot_x, arg_hot_y}
	iv.Call(args, nil, nil)
}

// Enum DragProtocol
type DragProtocolEnum int

const (
	DragProtocolNone           DragProtocolEnum = 0
	DragProtocolMotif          DragProtocolEnum = 1
	DragProtocolXdnd           DragProtocolEnum = 2
	DragProtocolRootwin        DragProtocolEnum = 3
	DragProtocolWin32Dropfiles DragProtocolEnum = 4
	DragProtocolOle2           DragProtocolEnum = 5
	DragProtocolLocal          DragProtocolEnum = 6
	DragProtocolWayland        DragProtocolEnum = 7
)

func DragProtocolGetType() gi.GType {
	ret := _I.GetGType(22, "DragProtocol")
	return ret
}

// Object DrawingContext
type DrawingContext struct {
	g.Object
}

func WrapDrawingContext(p unsafe.Pointer) (r DrawingContext) { r.P = p; return }

type IDrawingContext interface{ P_DrawingContext() unsafe.Pointer }

func (v DrawingContext) P_DrawingContext() unsafe.Pointer { return v.P }
func DrawingContextGetType() gi.GType {
	ret := _I.GetGType(23, "DrawingContext")
	return ret
}

// gdk_drawing_context_get_cairo_context
//
// [ result ] trans: nothing
//
func (v DrawingContext) GetCairoContext() (result cairo.Context) {
	iv, err := _I.Get(131, "DrawingContext", "get_cairo_context")
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

// gdk_drawing_context_get_clip
//
// [ result ] trans: everything
//
func (v DrawingContext) GetClip() (result cairo.Region) {
	iv, err := _I.Get(132, "DrawingContext", "get_clip")
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

// gdk_drawing_context_get_window
//
// [ result ] trans: nothing
//
func (v DrawingContext) GetWindow() (result Window) {
	iv, err := _I.Get(133, "DrawingContext", "get_window")
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

// gdk_drawing_context_is_valid
//
// [ result ] trans: nothing
//
func (v DrawingContext) IsValid() (result bool) {
	iv, err := _I.Get(134, "DrawingContext", "is_valid")
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

// ignore GType struct DrawingContextClass

// Union Event
type Event struct {
	P unsafe.Pointer
}

const SizeOfUnionEvent = 96

func EventGetType() gi.GType {
	ret := _I.GetGType(24, "Event")
	return ret
}

// gdk_event_new
//
// [ type1 ] trans: nothing
//
// [ result ] trans: everything
//
func NewEvent(type1 EventTypeEnum) (result Event) {
	iv, err := _I.Get(135, "Event", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_events_get_angle
//
// [ event2 ] trans: nothing
//
// [ angle ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Event) _GetAngle(event2 Event) (result bool, angle float64) {
	iv, err := _I.Get(136, "Event", "_get_angle")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_event2 := gi.NewPointerArgument(event2.P)
	arg_angle := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_event2, arg_angle}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	angle = outArgs[0].Double()
	result = ret.Bool()
	return
}

// gdk_events_get_center
//
// [ event2 ] trans: nothing
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Event) _GetCenter(event2 Event) (result bool, x float64, y float64) {
	iv, err := _I.Get(137, "Event", "_get_center")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_event2 := gi.NewPointerArgument(event2.P)
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_event2, arg_x, arg_y}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	x = outArgs[0].Double()
	y = outArgs[1].Double()
	result = ret.Bool()
	return
}

// gdk_events_get_distance
//
// [ event2 ] trans: nothing
//
// [ distance ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Event) _GetDistance(event2 Event) (result bool, distance float64) {
	iv, err := _I.Get(138, "Event", "_get_distance")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_event2 := gi.NewPointerArgument(event2.P)
	arg_distance := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_event2, arg_distance}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	distance = outArgs[0].Double()
	result = ret.Bool()
	return
}

// gdk_event_copy
//
// [ result ] trans: everything
//
func (v Event) Copy() (result Event) {
	iv, err := _I.Get(139, "Event", "copy")
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

// gdk_event_free
//
func (v Event) Free() {
	iv, err := _I.Get(140, "Event", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_event_get_axis
//
// [ axis_use ] trans: nothing
//
// [ value ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Event) GetAxis(axis_use AxisUseEnum) (result bool, value float64) {
	iv, err := _I.Get(141, "Event", "get_axis")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_axis_use := gi.NewIntArgument(int(axis_use))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_axis_use, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	value = outArgs[0].Double()
	result = ret.Bool()
	return
}

// gdk_event_get_button
//
// [ button ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Event) GetButton() (result bool, button uint32) {
	iv, err := _I.Get(142, "Event", "get_button")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_button := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_button}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	button = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gdk_event_get_click_count
//
// [ click_count ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Event) GetClickCount() (result bool, click_count uint32) {
	iv, err := _I.Get(143, "Event", "get_click_count")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_click_count := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_click_count}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	click_count = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gdk_event_get_coords
//
// [ x_win ] trans: everything, dir: out
//
// [ y_win ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Event) GetCoords() (result bool, x_win float64, y_win float64) {
	iv, err := _I.Get(144, "Event", "get_coords")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_x_win := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y_win := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_x_win, arg_y_win}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	x_win = outArgs[0].Double()
	y_win = outArgs[1].Double()
	result = ret.Bool()
	return
}

// gdk_event_get_device
//
// [ result ] trans: nothing
//
func (v Event) GetDevice() (result Device) {
	iv, err := _I.Get(145, "Event", "get_device")
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

// gdk_event_get_device_tool
//
// [ result ] trans: nothing
//
func (v Event) GetDeviceTool() (result DeviceTool) {
	iv, err := _I.Get(146, "Event", "get_device_tool")
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

// gdk_event_get_event_sequence
//
// [ result ] trans: nothing
//
func (v Event) GetEventSequence() (result EventSequence) {
	iv, err := _I.Get(147, "Event", "get_event_sequence")
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

// gdk_event_get_event_type
//
// [ result ] trans: nothing
//
func (v Event) GetEventType() (result EventTypeEnum) {
	iv, err := _I.Get(148, "Event", "get_event_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = EventTypeEnum(ret.Int())
	return
}

// gdk_event_get_keycode
//
// [ keycode ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Event) GetKeycode() (result bool, keycode uint16) {
	iv, err := _I.Get(149, "Event", "get_keycode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_keycode := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_keycode}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	keycode = outArgs[0].Uint16()
	result = ret.Bool()
	return
}

// gdk_event_get_keyval
//
// [ keyval ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Event) GetKeyval() (result bool, keyval uint32) {
	iv, err := _I.Get(150, "Event", "get_keyval")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_keyval := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_keyval}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	keyval = outArgs[0].Uint32()
	result = ret.Bool()
	return
}

// gdk_event_get_pointer_emulated
//
// [ result ] trans: nothing
//
func (v Event) GetPointerEmulated() (result bool) {
	iv, err := _I.Get(151, "Event", "get_pointer_emulated")
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

// gdk_event_get_root_coords
//
// [ x_root ] trans: everything, dir: out
//
// [ y_root ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Event) GetRootCoords() (result bool, x_root float64, y_root float64) {
	iv, err := _I.Get(152, "Event", "get_root_coords")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_x_root := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y_root := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_x_root, arg_y_root}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	x_root = outArgs[0].Double()
	y_root = outArgs[1].Double()
	result = ret.Bool()
	return
}

// gdk_event_get_scancode
//
// [ result ] trans: nothing
//
func (v Event) GetScancode() (result int32) {
	iv, err := _I.Get(153, "Event", "get_scancode")
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

// gdk_event_get_screen
//
// [ result ] trans: nothing
//
func (v Event) GetScreen() (result Screen) {
	iv, err := _I.Get(154, "Event", "get_screen")
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

// gdk_event_get_scroll_deltas
//
// [ delta_x ] trans: everything, dir: out
//
// [ delta_y ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Event) GetScrollDeltas() (result bool, delta_x float64, delta_y float64) {
	iv, err := _I.Get(155, "Event", "get_scroll_deltas")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_delta_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_delta_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_delta_x, arg_delta_y}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	delta_x = outArgs[0].Double()
	delta_y = outArgs[1].Double()
	result = ret.Bool()
	return
}

// gdk_event_get_scroll_direction
//
// [ direction ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Event) GetScrollDirection() (result bool, direction ScrollDirectionEnum) {
	iv, err := _I.Get(156, "Event", "get_scroll_direction")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_direction := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_direction}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	direction = ScrollDirectionEnum(outArgs[0].Int())
	result = ret.Bool()
	return
}

// gdk_event_get_seat
//
// [ result ] trans: nothing
//
func (v Event) GetSeat() (result Seat) {
	iv, err := _I.Get(157, "Event", "get_seat")
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

// gdk_event_get_source_device
//
// [ result ] trans: nothing
//
func (v Event) GetSourceDevice() (result Device) {
	iv, err := _I.Get(158, "Event", "get_source_device")
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

// gdk_event_get_state
//
// [ state ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Event) GetState() (result bool, state ModifierTypeFlags) {
	iv, err := _I.Get(159, "Event", "get_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_state := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_state}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	state = ModifierTypeFlags(outArgs[0].Int())
	result = ret.Bool()
	return
}

// gdk_event_get_time
//
// [ result ] trans: nothing
//
func (v Event) GetTime() (result uint32) {
	iv, err := _I.Get(160, "Event", "get_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gdk_event_get_window
//
// [ result ] trans: nothing
//
func (v Event) GetWindow() (result Window) {
	iv, err := _I.Get(161, "Event", "get_window")
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

// gdk_event_is_scroll_stop_event
//
// [ result ] trans: nothing
//
func (v Event) IsScrollStopEvent() (result bool) {
	iv, err := _I.Get(162, "Event", "is_scroll_stop_event")
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

// gdk_event_put
//
func (v Event) Put() {
	iv, err := _I.Get(163, "Event", "put")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_event_set_device
//
// [ device ] trans: nothing
//
func (v Event) SetDevice(device IDevice) {
	iv, err := _I.Get(164, "Event", "set_device")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if device != nil {
		tmp = device.P_Device()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_device := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_device}
	iv.Call(args, nil, nil)
}

// gdk_event_set_device_tool
//
// [ tool ] trans: nothing
//
func (v Event) SetDeviceTool(tool IDeviceTool) {
	iv, err := _I.Get(165, "Event", "set_device_tool")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if tool != nil {
		tmp = tool.P_DeviceTool()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tool := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_tool}
	iv.Call(args, nil, nil)
}

// gdk_event_set_screen
//
// [ screen ] trans: nothing
//
func (v Event) SetScreen(screen IScreen) {
	iv, err := _I.Get(166, "Event", "set_screen")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if screen != nil {
		tmp = screen.P_Screen()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_screen := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_screen}
	iv.Call(args, nil, nil)
}

// gdk_event_set_source_device
//
// [ device ] trans: nothing
//
func (v Event) SetSourceDevice(device IDevice) {
	iv, err := _I.Get(167, "Event", "set_source_device")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if device != nil {
		tmp = device.P_Device()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_device := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_device}
	iv.Call(args, nil, nil)
}

// gdk_event_triggers_context_menu
//
// [ result ] trans: nothing
//
func (v Event) TriggersContextMenu() (result bool) {
	iv, err := _I.Get(168, "Event", "triggers_context_menu")
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

// gdk_event_handler_set
//
// [ func1 ] trans: nothing
//
// [ data ] trans: nothing
//
// [ notify ] trans: nothing
//
func EventHandlerSet1(func1 int /*TODO_TYPE CALLBACK*/, data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(170, "Event", "handler_set")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myEventFunc()))
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_func1, arg_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gdk_event_request_motions
//
// [ event ] trans: nothing
//
func EventRequestMotions1(event EventMotion) {
	iv, err := _I.Get(172, "Event", "request_motions")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_event := gi.NewPointerArgument(event.P)
	args := []gi.Argument{arg_event}
	iv.Call(args, nil, nil)
}

// Struct EventAny
type EventAny struct {
	P unsafe.Pointer
}

const SizeOfStructEventAny = 24

func EventAnyGetType() gi.GType {
	ret := _I.GetGType(25, "EventAny")
	return ret
}

// Struct EventButton
type EventButton struct {
	P unsafe.Pointer
}

const SizeOfStructEventButton = 80

func EventButtonGetType() gi.GType {
	ret := _I.GetGType(26, "EventButton")
	return ret
}

// Struct EventConfigure
type EventConfigure struct {
	P unsafe.Pointer
}

const SizeOfStructEventConfigure = 40

func EventConfigureGetType() gi.GType {
	ret := _I.GetGType(27, "EventConfigure")
	return ret
}

// Struct EventCrossing
type EventCrossing struct {
	P unsafe.Pointer
}

const SizeOfStructEventCrossing = 88

func EventCrossingGetType() gi.GType {
	ret := _I.GetGType(28, "EventCrossing")
	return ret
}

// Struct EventDND
type EventDND struct {
	P unsafe.Pointer
}

const SizeOfStructEventDND = 40

func EventDNDGetType() gi.GType {
	ret := _I.GetGType(29, "EventDND")
	return ret
}

// Struct EventExpose
type EventExpose struct {
	P unsafe.Pointer
}

const SizeOfStructEventExpose = 56

func EventExposeGetType() gi.GType {
	ret := _I.GetGType(30, "EventExpose")
	return ret
}

// Struct EventFocus
type EventFocus struct {
	P unsafe.Pointer
}

const SizeOfStructEventFocus = 24

func EventFocusGetType() gi.GType {
	ret := _I.GetGType(31, "EventFocus")
	return ret
}

type EventFuncStruct struct {
	F_event Event
	F_data  unsafe.Pointer
}

func GetPointer_myEventFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGdkEventFunc())
}

//export myGdkEventFunc
func myGdkEventFunc(event *C.GdkEvent, data C.gpointer) {
	// TODO: not found user_data
}

// Struct EventGrabBroken
type EventGrabBroken struct {
	P unsafe.Pointer
}

const SizeOfStructEventGrabBroken = 40

func EventGrabBrokenGetType() gi.GType {
	ret := _I.GetGType(32, "EventGrabBroken")
	return ret
}

// Struct EventKey
type EventKey struct {
	P unsafe.Pointer
}

const SizeOfStructEventKey = 56

func EventKeyGetType() gi.GType {
	ret := _I.GetGType(33, "EventKey")
	return ret
}

// Flags EventMask
type EventMaskFlags int

const (
	EventMaskExposureMask          EventMaskFlags = 2
	EventMaskPointerMotionMask     EventMaskFlags = 4
	EventMaskPointerMotionHintMask EventMaskFlags = 8
	EventMaskButtonMotionMask      EventMaskFlags = 16
	EventMaskButton1MotionMask     EventMaskFlags = 32
	EventMaskButton2MotionMask     EventMaskFlags = 64
	EventMaskButton3MotionMask     EventMaskFlags = 128
	EventMaskButtonPressMask       EventMaskFlags = 256
	EventMaskButtonReleaseMask     EventMaskFlags = 512
	EventMaskKeyPressMask          EventMaskFlags = 1024
	EventMaskKeyReleaseMask        EventMaskFlags = 2048
	EventMaskEnterNotifyMask       EventMaskFlags = 4096
	EventMaskLeaveNotifyMask       EventMaskFlags = 8192
	EventMaskFocusChangeMask       EventMaskFlags = 16384
	EventMaskStructureMask         EventMaskFlags = 32768
	EventMaskPropertyChangeMask    EventMaskFlags = 65536
	EventMaskVisibilityNotifyMask  EventMaskFlags = 131072
	EventMaskProximityInMask       EventMaskFlags = 262144
	EventMaskProximityOutMask      EventMaskFlags = 524288
	EventMaskSubstructureMask      EventMaskFlags = 1048576
	EventMaskScrollMask            EventMaskFlags = 2097152
	EventMaskTouchMask             EventMaskFlags = 4194304
	EventMaskSmoothScrollMask      EventMaskFlags = 8388608
	EventMaskTouchpadGestureMask   EventMaskFlags = 16777216
	EventMaskTabletPadMask         EventMaskFlags = 33554432
	EventMaskAllEventsMask         EventMaskFlags = 67108862
)

func EventMaskGetType() gi.GType {
	ret := _I.GetGType(34, "EventMask")
	return ret
}

// Struct EventMotion
type EventMotion struct {
	P unsafe.Pointer
}

const SizeOfStructEventMotion = 80

func EventMotionGetType() gi.GType {
	ret := _I.GetGType(35, "EventMotion")
	return ret
}

// Struct EventOwnerChange
type EventOwnerChange struct {
	P unsafe.Pointer
}

const SizeOfStructEventOwnerChange = 56

func EventOwnerChangeGetType() gi.GType {
	ret := _I.GetGType(36, "EventOwnerChange")
	return ret
}

// Struct EventPadAxis
type EventPadAxis struct {
	P unsafe.Pointer
}

const SizeOfStructEventPadAxis = 48

func EventPadAxisGetType() gi.GType {
	ret := _I.GetGType(37, "EventPadAxis")
	return ret
}

// Struct EventPadButton
type EventPadButton struct {
	P unsafe.Pointer
}

const SizeOfStructEventPadButton = 40

func EventPadButtonGetType() gi.GType {
	ret := _I.GetGType(38, "EventPadButton")
	return ret
}

// Struct EventPadGroupMode
type EventPadGroupMode struct {
	P unsafe.Pointer
}

const SizeOfStructEventPadGroupMode = 32

func EventPadGroupModeGetType() gi.GType {
	ret := _I.GetGType(39, "EventPadGroupMode")
	return ret
}

// Struct EventProperty
type EventProperty struct {
	P unsafe.Pointer
}

const SizeOfStructEventProperty = 40

func EventPropertyGetType() gi.GType {
	ret := _I.GetGType(40, "EventProperty")
	return ret
}

// Struct EventProximity
type EventProximity struct {
	P unsafe.Pointer
}

const SizeOfStructEventProximity = 32

func EventProximityGetType() gi.GType {
	ret := _I.GetGType(41, "EventProximity")
	return ret
}

// Struct EventScroll
type EventScroll struct {
	P unsafe.Pointer
}

const SizeOfStructEventScroll = 96

func EventScrollGetType() gi.GType {
	ret := _I.GetGType(42, "EventScroll")
	return ret
}

// Struct EventSelection
type EventSelection struct {
	P unsafe.Pointer
}

const SizeOfStructEventSelection = 64

func EventSelectionGetType() gi.GType {
	ret := _I.GetGType(43, "EventSelection")
	return ret
}

// Struct EventSequence
type EventSequence struct {
	P unsafe.Pointer
}

func EventSequenceGetType() gi.GType {
	ret := _I.GetGType(44, "EventSequence")
	return ret
}

// Struct EventSetting
type EventSetting struct {
	P unsafe.Pointer
}

const SizeOfStructEventSetting = 32

func EventSettingGetType() gi.GType {
	ret := _I.GetGType(45, "EventSetting")
	return ret
}

// Struct EventTouch
type EventTouch struct {
	P unsafe.Pointer
}

const SizeOfStructEventTouch = 96

func EventTouchGetType() gi.GType {
	ret := _I.GetGType(46, "EventTouch")
	return ret
}

// Struct EventTouchpadPinch
type EventTouchpadPinch struct {
	P unsafe.Pointer
}

const SizeOfStructEventTouchpadPinch = 96

func EventTouchpadPinchGetType() gi.GType {
	ret := _I.GetGType(47, "EventTouchpadPinch")
	return ret
}

// Struct EventTouchpadSwipe
type EventTouchpadSwipe struct {
	P unsafe.Pointer
}

const SizeOfStructEventTouchpadSwipe = 80

func EventTouchpadSwipeGetType() gi.GType {
	ret := _I.GetGType(48, "EventTouchpadSwipe")
	return ret
}

// Enum EventType
type EventTypeEnum int

const (
	EventTypeNothing           EventTypeEnum = -1
	EventTypeDelete            EventTypeEnum = 0
	EventTypeDestroy           EventTypeEnum = 1
	EventTypeExpose            EventTypeEnum = 2
	EventTypeMotionNotify      EventTypeEnum = 3
	EventTypeButtonPress       EventTypeEnum = 4
	EventType2buttonPress      EventTypeEnum = 5
	EventTypeDoubleButtonPress EventTypeEnum = 5
	EventType3buttonPress      EventTypeEnum = 6
	EventTypeTripleButtonPress EventTypeEnum = 6
	EventTypeButtonRelease     EventTypeEnum = 7
	EventTypeKeyPress          EventTypeEnum = 8
	EventTypeKeyRelease        EventTypeEnum = 9
	EventTypeEnterNotify       EventTypeEnum = 10
	EventTypeLeaveNotify       EventTypeEnum = 11
	EventTypeFocusChange       EventTypeEnum = 12
	EventTypeConfigure         EventTypeEnum = 13
	EventTypeMap               EventTypeEnum = 14
	EventTypeUnmap             EventTypeEnum = 15
	EventTypePropertyNotify    EventTypeEnum = 16
	EventTypeSelectionClear    EventTypeEnum = 17
	EventTypeSelectionRequest  EventTypeEnum = 18
	EventTypeSelectionNotify   EventTypeEnum = 19
	EventTypeProximityIn       EventTypeEnum = 20
	EventTypeProximityOut      EventTypeEnum = 21
	EventTypeDragEnter         EventTypeEnum = 22
	EventTypeDragLeave         EventTypeEnum = 23
	EventTypeDragMotion        EventTypeEnum = 24
	EventTypeDragStatus        EventTypeEnum = 25
	EventTypeDropStart         EventTypeEnum = 26
	EventTypeDropFinished      EventTypeEnum = 27
	EventTypeClientEvent       EventTypeEnum = 28
	EventTypeVisibilityNotify  EventTypeEnum = 29
	EventTypeScroll            EventTypeEnum = 31
	EventTypeWindowState       EventTypeEnum = 32
	EventTypeSetting           EventTypeEnum = 33
	EventTypeOwnerChange       EventTypeEnum = 34
	EventTypeGrabBroken        EventTypeEnum = 35
	EventTypeDamage            EventTypeEnum = 36
	EventTypeTouchBegin        EventTypeEnum = 37
	EventTypeTouchUpdate       EventTypeEnum = 38
	EventTypeTouchEnd          EventTypeEnum = 39
	EventTypeTouchCancel       EventTypeEnum = 40
	EventTypeTouchpadSwipe     EventTypeEnum = 41
	EventTypeTouchpadPinch     EventTypeEnum = 42
	EventTypePadButtonPress    EventTypeEnum = 43
	EventTypePadButtonRelease  EventTypeEnum = 44
	EventTypePadRing           EventTypeEnum = 45
	EventTypePadStrip          EventTypeEnum = 46
	EventTypePadGroupMode      EventTypeEnum = 47
	EventTypeEventLast         EventTypeEnum = 48
)

func EventTypeGetType() gi.GType {
	ret := _I.GetGType(49, "EventType")
	return ret
}

// Deprecated
//
// Struct EventVisibility
type EventVisibility struct {
	P unsafe.Pointer
}

const SizeOfStructEventVisibility = 24

func EventVisibilityGetType() gi.GType {
	ret := _I.GetGType(50, "EventVisibility")
	return ret
}

// Struct EventWindowState
type EventWindowState struct {
	P unsafe.Pointer
}

const SizeOfStructEventWindowState = 32

func EventWindowStateGetType() gi.GType {
	ret := _I.GetGType(51, "EventWindowState")
	return ret
}

type FilterFuncStruct struct {
	F_xevent unsafe.Pointer
	F_event  Event
	F_data   unsafe.Pointer
}

func GetPointer_myFilterFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGdkFilterFunc())
}

//export myGdkFilterFunc
func myGdkFilterFunc(xevent C.gpointer, event *C.GdkEvent, data C.gpointer) {
	// TODO: not found user_data
}

// Enum FilterReturn
type FilterReturnEnum int

const (
	FilterReturnContinue  FilterReturnEnum = 0
	FilterReturnTranslate FilterReturnEnum = 1
	FilterReturnRemove    FilterReturnEnum = 2
)

func FilterReturnGetType() gi.GType {
	ret := _I.GetGType(52, "FilterReturn")
	return ret
}

// Object FrameClock
type FrameClock struct {
	g.Object
}

func WrapFrameClock(p unsafe.Pointer) (r FrameClock) { r.P = p; return }

type IFrameClock interface{ P_FrameClock() unsafe.Pointer }

func (v FrameClock) P_FrameClock() unsafe.Pointer { return v.P }
func FrameClockGetType() gi.GType {
	ret := _I.GetGType(53, "FrameClock")
	return ret
}

// gdk_frame_clock_begin_updating
//
func (v FrameClock) BeginUpdating() {
	iv, err := _I.Get(173, "FrameClock", "begin_updating")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_frame_clock_end_updating
//
func (v FrameClock) EndUpdating() {
	iv, err := _I.Get(174, "FrameClock", "end_updating")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_frame_clock_get_current_timings
//
// [ result ] trans: nothing
//
func (v FrameClock) GetCurrentTimings() (result FrameTimings) {
	iv, err := _I.Get(175, "FrameClock", "get_current_timings")
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

// gdk_frame_clock_get_frame_counter
//
// [ result ] trans: nothing
//
func (v FrameClock) GetFrameCounter() (result int64) {
	iv, err := _I.Get(176, "FrameClock", "get_frame_counter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// gdk_frame_clock_get_frame_time
//
// [ result ] trans: nothing
//
func (v FrameClock) GetFrameTime() (result int64) {
	iv, err := _I.Get(177, "FrameClock", "get_frame_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// gdk_frame_clock_get_history_start
//
// [ result ] trans: nothing
//
func (v FrameClock) GetHistoryStart() (result int64) {
	iv, err := _I.Get(178, "FrameClock", "get_history_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// gdk_frame_clock_get_refresh_info
//
// [ base_time ] trans: nothing
//
// [ refresh_interval_return ] trans: everything, dir: out
//
// [ presentation_time_return ] trans: everything, dir: out
//
func (v FrameClock) GetRefreshInfo(base_time int64) (refresh_interval_return int64, presentation_time_return int64) {
	iv, err := _I.Get(179, "FrameClock", "get_refresh_info")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_base_time := gi.NewInt64Argument(base_time)
	arg_refresh_interval_return := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_presentation_time_return := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_base_time, arg_refresh_interval_return, arg_presentation_time_return}
	iv.Call(args, nil, &outArgs[0])
	refresh_interval_return = outArgs[0].Int64()
	presentation_time_return = outArgs[1].Int64()
	return
}

// gdk_frame_clock_get_timings
//
// [ frame_counter ] trans: nothing
//
// [ result ] trans: nothing
//
func (v FrameClock) GetTimings(frame_counter int64) (result FrameTimings) {
	iv, err := _I.Get(180, "FrameClock", "get_timings")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_frame_counter := gi.NewInt64Argument(frame_counter)
	args := []gi.Argument{arg_v, arg_frame_counter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_frame_clock_request_phase
//
// [ phase ] trans: nothing
//
func (v FrameClock) RequestPhase(phase FrameClockPhaseFlags) {
	iv, err := _I.Get(181, "FrameClock", "request_phase")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_phase := gi.NewIntArgument(int(phase))
	args := []gi.Argument{arg_v, arg_phase}
	iv.Call(args, nil, nil)
}

// ignore GType struct FrameClockClass

// Flags FrameClockPhase
type FrameClockPhaseFlags int

const (
	FrameClockPhaseNone         FrameClockPhaseFlags = 0
	FrameClockPhaseFlushEvents  FrameClockPhaseFlags = 1
	FrameClockPhaseBeforePaint  FrameClockPhaseFlags = 2
	FrameClockPhaseUpdate       FrameClockPhaseFlags = 4
	FrameClockPhaseLayout       FrameClockPhaseFlags = 8
	FrameClockPhasePaint        FrameClockPhaseFlags = 16
	FrameClockPhaseResumeEvents FrameClockPhaseFlags = 32
	FrameClockPhaseAfterPaint   FrameClockPhaseFlags = 64
)

func FrameClockPhaseGetType() gi.GType {
	ret := _I.GetGType(54, "FrameClockPhase")
	return ret
}

// Struct FrameClockPrivate
type FrameClockPrivate struct {
	P unsafe.Pointer
}

func FrameClockPrivateGetType() gi.GType {
	ret := _I.GetGType(55, "FrameClockPrivate")
	return ret
}

// Struct FrameTimings
type FrameTimings struct {
	P unsafe.Pointer
}

func FrameTimingsGetType() gi.GType {
	ret := _I.GetGType(56, "FrameTimings")
	return ret
}

// gdk_frame_timings_get_complete
//
// [ result ] trans: nothing
//
func (v FrameTimings) GetComplete() (result bool) {
	iv, err := _I.Get(182, "FrameTimings", "get_complete")
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

// gdk_frame_timings_get_frame_counter
//
// [ result ] trans: nothing
//
func (v FrameTimings) GetFrameCounter() (result int64) {
	iv, err := _I.Get(183, "FrameTimings", "get_frame_counter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// gdk_frame_timings_get_frame_time
//
// [ result ] trans: nothing
//
func (v FrameTimings) GetFrameTime() (result int64) {
	iv, err := _I.Get(184, "FrameTimings", "get_frame_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// gdk_frame_timings_get_predicted_presentation_time
//
// [ result ] trans: nothing
//
func (v FrameTimings) GetPredictedPresentationTime() (result int64) {
	iv, err := _I.Get(185, "FrameTimings", "get_predicted_presentation_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// gdk_frame_timings_get_presentation_time
//
// [ result ] trans: nothing
//
func (v FrameTimings) GetPresentationTime() (result int64) {
	iv, err := _I.Get(186, "FrameTimings", "get_presentation_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// gdk_frame_timings_get_refresh_interval
//
// [ result ] trans: nothing
//
func (v FrameTimings) GetRefreshInterval() (result int64) {
	iv, err := _I.Get(187, "FrameTimings", "get_refresh_interval")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// gdk_frame_timings_ref
//
// [ result ] trans: everything
//
func (v FrameTimings) Ref() (result FrameTimings) {
	iv, err := _I.Get(188, "FrameTimings", "ref")
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

// gdk_frame_timings_unref
//
func (v FrameTimings) Unref() {
	iv, err := _I.Get(189, "FrameTimings", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum FullscreenMode
type FullscreenModeEnum int

const (
	FullscreenModeCurrentMonitor FullscreenModeEnum = 0
	FullscreenModeAllMonitors    FullscreenModeEnum = 1
)

func FullscreenModeGetType() gi.GType {
	ret := _I.GetGType(57, "FullscreenMode")
	return ret
}

// Object GLContext
type GLContext struct {
	g.Object
}

func WrapGLContext(p unsafe.Pointer) (r GLContext) { r.P = p; return }

type IGLContext interface{ P_GLContext() unsafe.Pointer }

func (v GLContext) P_GLContext() unsafe.Pointer { return v.P }
func GLContextGetType() gi.GType {
	ret := _I.GetGType(58, "GLContext")
	return ret
}

// gdk_gl_context_get_debug_enabled
//
// [ result ] trans: nothing
//
func (v GLContext) GetDebugEnabled() (result bool) {
	iv, err := _I.Get(192, "GLContext", "get_debug_enabled")
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

// gdk_gl_context_get_display
//
// [ result ] trans: nothing
//
func (v GLContext) GetDisplay() (result Display) {
	iv, err := _I.Get(193, "GLContext", "get_display")
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

// gdk_gl_context_get_forward_compatible
//
// [ result ] trans: nothing
//
func (v GLContext) GetForwardCompatible() (result bool) {
	iv, err := _I.Get(194, "GLContext", "get_forward_compatible")
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

// gdk_gl_context_get_required_version
//
// [ major ] trans: everything, dir: out
//
// [ minor ] trans: everything, dir: out
//
func (v GLContext) GetRequiredVersion() (major int32, minor int32) {
	iv, err := _I.Get(195, "GLContext", "get_required_version")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_major := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_minor := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_major, arg_minor}
	iv.Call(args, nil, &outArgs[0])
	major = outArgs[0].Int32()
	minor = outArgs[1].Int32()
	return
}

// gdk_gl_context_get_shared_context
//
// [ result ] trans: nothing
//
func (v GLContext) GetSharedContext() (result GLContext) {
	iv, err := _I.Get(196, "GLContext", "get_shared_context")
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

// gdk_gl_context_get_use_es
//
// [ result ] trans: nothing
//
func (v GLContext) GetUseEs() (result bool) {
	iv, err := _I.Get(197, "GLContext", "get_use_es")
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

// gdk_gl_context_get_version
//
// [ major ] trans: everything, dir: out
//
// [ minor ] trans: everything, dir: out
//
func (v GLContext) GetVersion() (major int32, minor int32) {
	iv, err := _I.Get(198, "GLContext", "get_version")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_major := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_minor := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_major, arg_minor}
	iv.Call(args, nil, &outArgs[0])
	major = outArgs[0].Int32()
	minor = outArgs[1].Int32()
	return
}

// gdk_gl_context_get_window
//
// [ result ] trans: nothing
//
func (v GLContext) GetWindow() (result Window) {
	iv, err := _I.Get(199, "GLContext", "get_window")
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

// gdk_gl_context_is_legacy
//
// [ result ] trans: nothing
//
func (v GLContext) IsLegacy() (result bool) {
	iv, err := _I.Get(200, "GLContext", "is_legacy")
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

// gdk_gl_context_make_current
//
func (v GLContext) MakeCurrent() {
	iv, err := _I.Get(201, "GLContext", "make_current")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_gl_context_realize
//
// [ result ] trans: nothing
//
func (v GLContext) Realize() (result bool, err error) {
	iv, err := _I.Get(202, "GLContext", "realize")
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

// gdk_gl_context_set_debug_enabled
//
// [ enabled ] trans: nothing
//
func (v GLContext) SetDebugEnabled(enabled bool) {
	iv, err := _I.Get(203, "GLContext", "set_debug_enabled")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_enabled := gi.NewBoolArgument(enabled)
	args := []gi.Argument{arg_v, arg_enabled}
	iv.Call(args, nil, nil)
}

// gdk_gl_context_set_forward_compatible
//
// [ compatible ] trans: nothing
//
func (v GLContext) SetForwardCompatible(compatible bool) {
	iv, err := _I.Get(204, "GLContext", "set_forward_compatible")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_compatible := gi.NewBoolArgument(compatible)
	args := []gi.Argument{arg_v, arg_compatible}
	iv.Call(args, nil, nil)
}

// gdk_gl_context_set_required_version
//
// [ major ] trans: nothing
//
// [ minor ] trans: nothing
//
func (v GLContext) SetRequiredVersion(major int32, minor int32) {
	iv, err := _I.Get(205, "GLContext", "set_required_version")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_major := gi.NewInt32Argument(major)
	arg_minor := gi.NewInt32Argument(minor)
	args := []gi.Argument{arg_v, arg_major, arg_minor}
	iv.Call(args, nil, nil)
}

// gdk_gl_context_set_use_es
//
// [ use_es ] trans: nothing
//
func (v GLContext) SetUseEs(use_es int32) {
	iv, err := _I.Get(206, "GLContext", "set_use_es")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_use_es := gi.NewInt32Argument(use_es)
	args := []gi.Argument{arg_v, arg_use_es}
	iv.Call(args, nil, nil)
}

// Enum GLError
type GLErrorEnum int

const (
	GLErrorNotAvailable       GLErrorEnum = 0
	GLErrorUnsupportedFormat  GLErrorEnum = 1
	GLErrorUnsupportedProfile GLErrorEnum = 2
)

func GLErrorGetType() gi.GType {
	ret := _I.GetGType(59, "GLError")
	return ret
}

// Struct Geometry
type Geometry struct {
	P unsafe.Pointer
}

const SizeOfStructGeometry = 56

func GeometryGetType() gi.GType {
	ret := _I.GetGType(60, "Geometry")
	return ret
}

// Enum GrabOwnership
type GrabOwnershipEnum int

const (
	GrabOwnershipNone        GrabOwnershipEnum = 0
	GrabOwnershipWindow      GrabOwnershipEnum = 1
	GrabOwnershipApplication GrabOwnershipEnum = 2
)

func GrabOwnershipGetType() gi.GType {
	ret := _I.GetGType(61, "GrabOwnership")
	return ret
}

// Enum GrabStatus
type GrabStatusEnum int

const (
	GrabStatusSuccess        GrabStatusEnum = 0
	GrabStatusAlreadyGrabbed GrabStatusEnum = 1
	GrabStatusInvalidTime    GrabStatusEnum = 2
	GrabStatusNotViewable    GrabStatusEnum = 3
	GrabStatusFrozen         GrabStatusEnum = 4
	GrabStatusFailed         GrabStatusEnum = 5
)

func GrabStatusGetType() gi.GType {
	ret := _I.GetGType(62, "GrabStatus")
	return ret
}

// Enum Gravity
type GravityEnum int

const (
	GravityNorthWest GravityEnum = 1
	GravityNorth     GravityEnum = 2
	GravityNorthEast GravityEnum = 3
	GravityWest      GravityEnum = 4
	GravityCenter    GravityEnum = 5
	GravityEast      GravityEnum = 6
	GravitySouthWest GravityEnum = 7
	GravitySouth     GravityEnum = 8
	GravitySouthEast GravityEnum = 9
	GravityStatic    GravityEnum = 10
)

func GravityGetType() gi.GType {
	ret := _I.GetGType(63, "Gravity")
	return ret
}

// Enum InputMode
type InputModeEnum int

const (
	InputModeDisabled InputModeEnum = 0
	InputModeScreen   InputModeEnum = 1
	InputModeWindow   InputModeEnum = 2
)

func InputModeGetType() gi.GType {
	ret := _I.GetGType(64, "InputMode")
	return ret
}

// Enum InputSource
type InputSourceEnum int

const (
	InputSourceMouse       InputSourceEnum = 0
	InputSourcePen         InputSourceEnum = 1
	InputSourceEraser      InputSourceEnum = 2
	InputSourceCursor      InputSourceEnum = 3
	InputSourceKeyboard    InputSourceEnum = 4
	InputSourceTouchscreen InputSourceEnum = 5
	InputSourceTouchpad    InputSourceEnum = 6
	InputSourceTrackpoint  InputSourceEnum = 7
	InputSourceTabletPad   InputSourceEnum = 8
)

func InputSourceGetType() gi.GType {
	ret := _I.GetGType(65, "InputSource")
	return ret
}

// Object Keymap
type Keymap struct {
	g.Object
}

func WrapKeymap(p unsafe.Pointer) (r Keymap) { r.P = p; return }

type IKeymap interface{ P_Keymap() unsafe.Pointer }

func (v Keymap) P_Keymap() unsafe.Pointer { return v.P }
func KeymapGetType() gi.GType {
	ret := _I.GetGType(66, "Keymap")
	return ret
}

// Deprecated
//
// gdk_keymap_get_for_display
//
// [ display ] trans: nothing
//
// [ result ] trans: nothing
//
func KeymapGetForDisplay1(display IDisplay) (result Keymap) {
	iv, err := _I.Get(208, "Keymap", "get_for_display")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if display != nil {
		tmp = display.P_Display()
	}
	arg_display := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_display}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_keymap_add_virtual_modifiers
//
// [ state ] trans: everything, dir: inout
//
func (v Keymap) AddVirtualModifiers(state int /*TODO:TYPE*/) {
	iv, err := _I.Get(209, "Keymap", "add_virtual_modifiers")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, &outArgs[0])
}

// gdk_keymap_get_caps_lock_state
//
// [ result ] trans: nothing
//
func (v Keymap) GetCapsLockState() (result bool) {
	iv, err := _I.Get(210, "Keymap", "get_caps_lock_state")
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

// gdk_keymap_get_direction
//
// [ result ] trans: nothing
//
func (v Keymap) GetDirection() (result pango.DirectionEnum) {
	iv, err := _I.Get(211, "Keymap", "get_direction")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = pango.DirectionEnum(ret.Int())
	return
}

// gdk_keymap_get_entries_for_keycode
//
// [ hardware_keycode ] trans: nothing
//
// [ keys ] trans: everything, dir: out
//
// [ keyvals ] trans: everything, dir: out
//
// [ n_entries ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Keymap) GetEntriesForKeycode(hardware_keycode uint32) (result bool, keys unsafe.Pointer, keyvals gi.Uint32Array) {
	iv, err := _I.Get(212, "Keymap", "get_entries_for_keycode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_hardware_keycode := gi.NewUint32Argument(hardware_keycode)
	arg_keys := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_keyvals := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_n_entries := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_hardware_keycode, arg_keys, arg_keyvals, arg_n_entries}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_entries int32
	_ = n_entries
	keys = outArgs[0].Pointer()
	keyvals.P = outArgs[1].Pointer()
	n_entries = outArgs[2].Int32()
	result = ret.Bool()
	keyvals.Len = int(n_entries)
	return
}

// gdk_keymap_get_entries_for_keyval
//
// [ keyval ] trans: nothing
//
// [ keys ] trans: everything, dir: out
//
// [ n_keys ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Keymap) GetEntriesForKeyval(keyval uint32) (result bool, keys unsafe.Pointer) {
	iv, err := _I.Get(213, "Keymap", "get_entries_for_keyval")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_keyval := gi.NewUint32Argument(keyval)
	arg_keys := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_n_keys := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_keyval, arg_keys, arg_n_keys}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_keys int32
	_ = n_keys
	keys = outArgs[0].Pointer()
	n_keys = outArgs[1].Int32()
	result = ret.Bool()
	return
}

// gdk_keymap_get_modifier_mask
//
// [ intent ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Keymap) GetModifierMask(intent ModifierIntentEnum) (result ModifierTypeFlags) {
	iv, err := _I.Get(214, "Keymap", "get_modifier_mask")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_intent := gi.NewIntArgument(int(intent))
	args := []gi.Argument{arg_v, arg_intent}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ModifierTypeFlags(ret.Int())
	return
}

// gdk_keymap_get_modifier_state
//
// [ result ] trans: nothing
//
func (v Keymap) GetModifierState() (result uint32) {
	iv, err := _I.Get(215, "Keymap", "get_modifier_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gdk_keymap_get_num_lock_state
//
// [ result ] trans: nothing
//
func (v Keymap) GetNumLockState() (result bool) {
	iv, err := _I.Get(216, "Keymap", "get_num_lock_state")
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

// gdk_keymap_get_scroll_lock_state
//
// [ result ] trans: nothing
//
func (v Keymap) GetScrollLockState() (result bool) {
	iv, err := _I.Get(217, "Keymap", "get_scroll_lock_state")
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

// gdk_keymap_have_bidi_layouts
//
// [ result ] trans: nothing
//
func (v Keymap) HaveBidiLayouts() (result bool) {
	iv, err := _I.Get(218, "Keymap", "have_bidi_layouts")
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

// gdk_keymap_lookup_key
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Keymap) LookupKey(key KeymapKey) (result uint32) {
	iv, err := _I.Get(219, "Keymap", "lookup_key")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewPointerArgument(key.P)
	args := []gi.Argument{arg_v, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gdk_keymap_map_virtual_modifiers
//
// [ state ] trans: everything, dir: inout
//
// [ result ] trans: nothing
//
func (v Keymap) MapVirtualModifiers(state int /*TODO:TYPE*/) (result bool) {
	iv, err := _I.Get(220, "Keymap", "map_virtual_modifiers")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	return
}

// gdk_keymap_translate_keyboard_state
//
// [ hardware_keycode ] trans: nothing
//
// [ state ] trans: nothing
//
// [ group ] trans: nothing
//
// [ keyval ] trans: everything, dir: out
//
// [ effective_group ] trans: everything, dir: out
//
// [ level ] trans: everything, dir: out
//
// [ consumed_modifiers ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Keymap) TranslateKeyboardState(hardware_keycode uint32, state ModifierTypeFlags, group int32) (result bool, keyval uint32, effective_group int32, level int32, consumed_modifiers ModifierTypeFlags) {
	iv, err := _I.Get(221, "Keymap", "translate_keyboard_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_hardware_keycode := gi.NewUint32Argument(hardware_keycode)
	arg_state := gi.NewIntArgument(int(state))
	arg_group := gi.NewInt32Argument(group)
	arg_keyval := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_effective_group := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_level := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_consumed_modifiers := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_hardware_keycode, arg_state, arg_group, arg_keyval, arg_effective_group, arg_level, arg_consumed_modifiers}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	keyval = outArgs[0].Uint32()
	effective_group = outArgs[1].Int32()
	level = outArgs[2].Int32()
	consumed_modifiers = ModifierTypeFlags(outArgs[3].Int())
	result = ret.Bool()
	return
}

// Struct KeymapKey
type KeymapKey struct {
	P unsafe.Pointer
}

const SizeOfStructKeymapKey = 12

func KeymapKeyGetType() gi.GType {
	ret := _I.GetGType(67, "KeymapKey")
	return ret
}

// Enum ModifierIntent
type ModifierIntentEnum int

const (
	ModifierIntentPrimaryAccelerator ModifierIntentEnum = 0
	ModifierIntentContextMenu        ModifierIntentEnum = 1
	ModifierIntentExtendSelection    ModifierIntentEnum = 2
	ModifierIntentModifySelection    ModifierIntentEnum = 3
	ModifierIntentNoTextInput        ModifierIntentEnum = 4
	ModifierIntentShiftGroup         ModifierIntentEnum = 5
	ModifierIntentDefaultModMask     ModifierIntentEnum = 6
)

func ModifierIntentGetType() gi.GType {
	ret := _I.GetGType(68, "ModifierIntent")
	return ret
}

// Flags ModifierType
type ModifierTypeFlags int

const (
	ModifierTypeShiftMask              ModifierTypeFlags = 1
	ModifierTypeLockMask               ModifierTypeFlags = 2
	ModifierTypeControlMask            ModifierTypeFlags = 4
	ModifierTypeMod1Mask               ModifierTypeFlags = 8
	ModifierTypeMod2Mask               ModifierTypeFlags = 16
	ModifierTypeMod3Mask               ModifierTypeFlags = 32
	ModifierTypeMod4Mask               ModifierTypeFlags = 64
	ModifierTypeMod5Mask               ModifierTypeFlags = 128
	ModifierTypeButton1Mask            ModifierTypeFlags = 256
	ModifierTypeButton2Mask            ModifierTypeFlags = 512
	ModifierTypeButton3Mask            ModifierTypeFlags = 1024
	ModifierTypeButton4Mask            ModifierTypeFlags = 2048
	ModifierTypeButton5Mask            ModifierTypeFlags = 4096
	ModifierTypeModifierReserved13Mask ModifierTypeFlags = 8192
	ModifierTypeModifierReserved14Mask ModifierTypeFlags = 16384
	ModifierTypeModifierReserved15Mask ModifierTypeFlags = 32768
	ModifierTypeModifierReserved16Mask ModifierTypeFlags = 65536
	ModifierTypeModifierReserved17Mask ModifierTypeFlags = 131072
	ModifierTypeModifierReserved18Mask ModifierTypeFlags = 262144
	ModifierTypeModifierReserved19Mask ModifierTypeFlags = 524288
	ModifierTypeModifierReserved20Mask ModifierTypeFlags = 1048576
	ModifierTypeModifierReserved21Mask ModifierTypeFlags = 2097152
	ModifierTypeModifierReserved22Mask ModifierTypeFlags = 4194304
	ModifierTypeModifierReserved23Mask ModifierTypeFlags = 8388608
	ModifierTypeModifierReserved24Mask ModifierTypeFlags = 16777216
	ModifierTypeModifierReserved25Mask ModifierTypeFlags = 33554432
	ModifierTypeSuperMask              ModifierTypeFlags = 67108864
	ModifierTypeHyperMask              ModifierTypeFlags = 134217728
	ModifierTypeMetaMask               ModifierTypeFlags = 268435456
	ModifierTypeModifierReserved29Mask ModifierTypeFlags = 536870912
	ModifierTypeReleaseMask            ModifierTypeFlags = 1073741824
	ModifierTypeModifierMask           ModifierTypeFlags = 1543512063
)

func ModifierTypeGetType() gi.GType {
	ret := _I.GetGType(69, "ModifierType")
	return ret
}

// Object Monitor
type Monitor struct {
	g.Object
}

func WrapMonitor(p unsafe.Pointer) (r Monitor) { r.P = p; return }

type IMonitor interface{ P_Monitor() unsafe.Pointer }

func (v Monitor) P_Monitor() unsafe.Pointer { return v.P }
func MonitorGetType() gi.GType {
	ret := _I.GetGType(70, "Monitor")
	return ret
}

// gdk_monitor_get_display
//
// [ result ] trans: nothing
//
func (v Monitor) GetDisplay() (result Display) {
	iv, err := _I.Get(222, "Monitor", "get_display")
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

// gdk_monitor_get_geometry
//
// [ geometry ] trans: nothing, dir: out
//
func (v Monitor) GetGeometry(geometry Rectangle) {
	iv, err := _I.Get(223, "Monitor", "get_geometry")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_geometry := gi.NewPointerArgument(geometry.P)
	args := []gi.Argument{arg_v, arg_geometry}
	iv.Call(args, nil, nil)
}

// gdk_monitor_get_height_mm
//
// [ result ] trans: nothing
//
func (v Monitor) GetHeightMm() (result int32) {
	iv, err := _I.Get(224, "Monitor", "get_height_mm")
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

// gdk_monitor_get_manufacturer
//
// [ result ] trans: nothing
//
func (v Monitor) GetManufacturer() (result string) {
	iv, err := _I.Get(225, "Monitor", "get_manufacturer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gdk_monitor_get_model
//
// [ result ] trans: nothing
//
func (v Monitor) GetModel() (result string) {
	iv, err := _I.Get(226, "Monitor", "get_model")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gdk_monitor_get_refresh_rate
//
// [ result ] trans: nothing
//
func (v Monitor) GetRefreshRate() (result int32) {
	iv, err := _I.Get(227, "Monitor", "get_refresh_rate")
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

// gdk_monitor_get_scale_factor
//
// [ result ] trans: nothing
//
func (v Monitor) GetScaleFactor() (result int32) {
	iv, err := _I.Get(228, "Monitor", "get_scale_factor")
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

// gdk_monitor_get_subpixel_layout
//
// [ result ] trans: nothing
//
func (v Monitor) GetSubpixelLayout() (result SubpixelLayoutEnum) {
	iv, err := _I.Get(229, "Monitor", "get_subpixel_layout")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = SubpixelLayoutEnum(ret.Int())
	return
}

// gdk_monitor_get_width_mm
//
// [ result ] trans: nothing
//
func (v Monitor) GetWidthMm() (result int32) {
	iv, err := _I.Get(230, "Monitor", "get_width_mm")
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

// gdk_monitor_get_workarea
//
// [ workarea ] trans: nothing, dir: out
//
func (v Monitor) GetWorkarea(workarea Rectangle) {
	iv, err := _I.Get(231, "Monitor", "get_workarea")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_workarea := gi.NewPointerArgument(workarea.P)
	args := []gi.Argument{arg_v, arg_workarea}
	iv.Call(args, nil, nil)
}

// gdk_monitor_is_primary
//
// [ result ] trans: nothing
//
func (v Monitor) IsPrimary() (result bool) {
	iv, err := _I.Get(232, "Monitor", "is_primary")
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

// ignore GType struct MonitorClass

// Enum NotifyType
type NotifyTypeEnum int

const (
	NotifyTypeAncestor         NotifyTypeEnum = 0
	NotifyTypeVirtual          NotifyTypeEnum = 1
	NotifyTypeInferior         NotifyTypeEnum = 2
	NotifyTypeNonlinear        NotifyTypeEnum = 3
	NotifyTypeNonlinearVirtual NotifyTypeEnum = 4
	NotifyTypeUnknown          NotifyTypeEnum = 5
)

func NotifyTypeGetType() gi.GType {
	ret := _I.GetGType(71, "NotifyType")
	return ret
}

// Enum OwnerChange
type OwnerChangeEnum int

const (
	OwnerChangeNewOwner OwnerChangeEnum = 0
	OwnerChangeDestroy  OwnerChangeEnum = 1
	OwnerChangeClose    OwnerChangeEnum = 2
)

func OwnerChangeGetType() gi.GType {
	ret := _I.GetGType(72, "OwnerChange")
	return ret
}

// Struct Point
type Point struct {
	P unsafe.Pointer
}

const SizeOfStructPoint = 8

func PointGetType() gi.GType {
	ret := _I.GetGType(73, "Point")
	return ret
}

// Enum PropMode
type PropModeEnum int

const (
	PropModeReplace PropModeEnum = 0
	PropModePrepend PropModeEnum = 1
	PropModeAppend  PropModeEnum = 2
)

func PropModeGetType() gi.GType {
	ret := _I.GetGType(74, "PropMode")
	return ret
}

// Enum PropertyState
type PropertyStateEnum int

const (
	PropertyStateNewValue PropertyStateEnum = 0
	PropertyStateDelete   PropertyStateEnum = 1
)

func PropertyStateGetType() gi.GType {
	ret := _I.GetGType(75, "PropertyState")
	return ret
}

// Struct RGBA
type RGBA struct {
	P unsafe.Pointer
}

const SizeOfStructRGBA = 32

func RGBAGetType() gi.GType {
	ret := _I.GetGType(76, "RGBA")
	return ret
}

// gdk_rgba_copy
//
// [ result ] trans: everything
//
func (v RGBA) Copy() (result RGBA) {
	iv, err := _I.Get(233, "RGBA", "copy")
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

// gdk_rgba_equal
//
// [ p2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v RGBA) Equal(p2 RGBA) (result bool) {
	iv, err := _I.Get(234, "RGBA", "equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_p2 := gi.NewPointerArgument(p2.P)
	args := []gi.Argument{arg_v, arg_p2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_rgba_free
//
func (v RGBA) Free() {
	iv, err := _I.Get(235, "RGBA", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_rgba_hash
//
// [ result ] trans: nothing
//
func (v RGBA) Hash() (result uint32) {
	iv, err := _I.Get(236, "RGBA", "hash")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gdk_rgba_parse
//
// [ spec ] trans: nothing
//
// [ result ] trans: nothing
//
func (v RGBA) Parse(spec string) (result bool) {
	iv, err := _I.Get(237, "RGBA", "parse")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_spec := gi.CString(spec)
	arg_v := gi.NewPointerArgument(v.P)
	arg_spec := gi.NewStringArgument(c_spec)
	args := []gi.Argument{arg_v, arg_spec}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_spec)
	result = ret.Bool()
	return
}

// gdk_rgba_to_string
//
// [ result ] trans: everything
//
func (v RGBA) ToString() (result string) {
	iv, err := _I.Get(238, "RGBA", "to_string")
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

// Struct Rectangle
type Rectangle struct {
	P unsafe.Pointer
}

const SizeOfStructRectangle = 16

func RectangleGetType() gi.GType {
	ret := _I.GetGType(77, "Rectangle")
	return ret
}

// gdk_rectangle_equal
//
// [ rect2 ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Rectangle) Equal(rect2 Rectangle) (result bool) {
	iv, err := _I.Get(239, "Rectangle", "equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_rect2 := gi.NewPointerArgument(rect2.P)
	args := []gi.Argument{arg_v, arg_rect2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_rectangle_intersect
//
// [ src2 ] trans: nothing
//
// [ dest ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func (v Rectangle) Intersect(src2 Rectangle, dest Rectangle) (result bool) {
	iv, err := _I.Get(240, "Rectangle", "intersect")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_src2 := gi.NewPointerArgument(src2.P)
	arg_dest := gi.NewPointerArgument(dest.P)
	args := []gi.Argument{arg_v, arg_src2, arg_dest}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_rectangle_union
//
// [ src2 ] trans: nothing
//
// [ dest ] trans: nothing, dir: out
//
func (v Rectangle) Union(src2 Rectangle, dest Rectangle) {
	iv, err := _I.Get(241, "Rectangle", "union")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_src2 := gi.NewPointerArgument(src2.P)
	arg_dest := gi.NewPointerArgument(dest.P)
	args := []gi.Argument{arg_v, arg_src2, arg_dest}
	iv.Call(args, nil, nil)
}

// Object Screen
type Screen struct {
	g.Object
}

func WrapScreen(p unsafe.Pointer) (r Screen) { r.P = p; return }

type IScreen interface{ P_Screen() unsafe.Pointer }

func (v Screen) P_Screen() unsafe.Pointer { return v.P }
func ScreenGetType() gi.GType {
	ret := _I.GetGType(78, "Screen")
	return ret
}

// Deprecated
//
// Deprecated
//
// Deprecated
//
// Deprecated
//
// Deprecated
//
// gdk_screen_get_active_window
//
// [ result ] trans: everything
//
func (v Screen) GetActiveWindow() (result Window) {
	iv, err := _I.Get(247, "Screen", "get_active_window")
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

// gdk_screen_get_display
//
// [ result ] trans: nothing
//
func (v Screen) GetDisplay() (result Display) {
	iv, err := _I.Get(248, "Screen", "get_display")
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

// gdk_screen_get_font_options
//
// [ result ] trans: nothing
//
func (v Screen) GetFontOptions() (result cairo.FontOptions) {
	iv, err := _I.Get(249, "Screen", "get_font_options")
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

// Deprecated
//
// gdk_screen_get_height
//
// [ result ] trans: nothing
//
func (v Screen) GetHeight() (result int32) {
	iv, err := _I.Get(250, "Screen", "get_height")
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

// Deprecated
//
// gdk_screen_get_height_mm
//
// [ result ] trans: nothing
//
func (v Screen) GetHeightMm() (result int32) {
	iv, err := _I.Get(251, "Screen", "get_height_mm")
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

// Deprecated
//
// gdk_screen_get_monitor_at_point
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Screen) GetMonitorAtPoint(x int32, y int32) (result int32) {
	iv, err := _I.Get(252, "Screen", "get_monitor_at_point")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	args := []gi.Argument{arg_v, arg_x, arg_y}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// Deprecated
//
// gdk_screen_get_monitor_at_window
//
// [ window ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Screen) GetMonitorAtWindow(window IWindow) (result int32) {
	iv, err := _I.Get(253, "Screen", "get_monitor_at_window")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_window := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_window}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// Deprecated
//
// gdk_screen_get_monitor_geometry
//
// [ monitor_num ] trans: nothing
//
// [ dest ] trans: nothing, dir: out
//
func (v Screen) GetMonitorGeometry(monitor_num int32, dest Rectangle) {
	iv, err := _I.Get(254, "Screen", "get_monitor_geometry")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_monitor_num := gi.NewInt32Argument(monitor_num)
	arg_dest := gi.NewPointerArgument(dest.P)
	args := []gi.Argument{arg_v, arg_monitor_num, arg_dest}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_screen_get_monitor_height_mm
//
// [ monitor_num ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Screen) GetMonitorHeightMm(monitor_num int32) (result int32) {
	iv, err := _I.Get(255, "Screen", "get_monitor_height_mm")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_monitor_num := gi.NewInt32Argument(monitor_num)
	args := []gi.Argument{arg_v, arg_monitor_num}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// Deprecated
//
// gdk_screen_get_monitor_plug_name
//
// [ monitor_num ] trans: nothing
//
// [ result ] trans: everything
//
func (v Screen) GetMonitorPlugName(monitor_num int32) (result string) {
	iv, err := _I.Get(256, "Screen", "get_monitor_plug_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_monitor_num := gi.NewInt32Argument(monitor_num)
	args := []gi.Argument{arg_v, arg_monitor_num}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// Deprecated
//
// gdk_screen_get_monitor_scale_factor
//
// [ monitor_num ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Screen) GetMonitorScaleFactor(monitor_num int32) (result int32) {
	iv, err := _I.Get(257, "Screen", "get_monitor_scale_factor")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_monitor_num := gi.NewInt32Argument(monitor_num)
	args := []gi.Argument{arg_v, arg_monitor_num}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// Deprecated
//
// gdk_screen_get_monitor_width_mm
//
// [ monitor_num ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Screen) GetMonitorWidthMm(monitor_num int32) (result int32) {
	iv, err := _I.Get(258, "Screen", "get_monitor_width_mm")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_monitor_num := gi.NewInt32Argument(monitor_num)
	args := []gi.Argument{arg_v, arg_monitor_num}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// Deprecated
//
// gdk_screen_get_monitor_workarea
//
// [ monitor_num ] trans: nothing
//
// [ dest ] trans: nothing, dir: out
//
func (v Screen) GetMonitorWorkarea(monitor_num int32, dest Rectangle) {
	iv, err := _I.Get(259, "Screen", "get_monitor_workarea")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_monitor_num := gi.NewInt32Argument(monitor_num)
	arg_dest := gi.NewPointerArgument(dest.P)
	args := []gi.Argument{arg_v, arg_monitor_num, arg_dest}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_screen_get_n_monitors
//
// [ result ] trans: nothing
//
func (v Screen) GetNMonitors() (result int32) {
	iv, err := _I.Get(260, "Screen", "get_n_monitors")
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

// Deprecated
//
// gdk_screen_get_number
//
// [ result ] trans: nothing
//
func (v Screen) GetNumber() (result int32) {
	iv, err := _I.Get(261, "Screen", "get_number")
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

// Deprecated
//
// gdk_screen_get_primary_monitor
//
// [ result ] trans: nothing
//
func (v Screen) GetPrimaryMonitor() (result int32) {
	iv, err := _I.Get(262, "Screen", "get_primary_monitor")
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

// gdk_screen_get_resolution
//
// [ result ] trans: nothing
//
func (v Screen) GetResolution() (result float64) {
	iv, err := _I.Get(263, "Screen", "get_resolution")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// gdk_screen_get_rgba_visual
//
// [ result ] trans: nothing
//
func (v Screen) GetRgbaVisual() (result Visual) {
	iv, err := _I.Get(264, "Screen", "get_rgba_visual")
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

// gdk_screen_get_root_window
//
// [ result ] trans: nothing
//
func (v Screen) GetRootWindow() (result Window) {
	iv, err := _I.Get(265, "Screen", "get_root_window")
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

// gdk_screen_get_setting
//
// [ name ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Screen) GetSetting(name string, value g.Value) (result bool) {
	iv, err := _I.Get(266, "Screen", "get_setting")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_name, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = ret.Bool()
	return
}

// gdk_screen_get_system_visual
//
// [ result ] trans: nothing
//
func (v Screen) GetSystemVisual() (result Visual) {
	iv, err := _I.Get(267, "Screen", "get_system_visual")
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

// gdk_screen_get_toplevel_windows
//
// [ result ] trans: container
//
func (v Screen) GetToplevelWindows() (result g.List) {
	iv, err := _I.Get(268, "Screen", "get_toplevel_windows")
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

// Deprecated
//
// gdk_screen_get_width
//
// [ result ] trans: nothing
//
func (v Screen) GetWidth() (result int32) {
	iv, err := _I.Get(269, "Screen", "get_width")
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

// Deprecated
//
// gdk_screen_get_width_mm
//
// [ result ] trans: nothing
//
func (v Screen) GetWidthMm() (result int32) {
	iv, err := _I.Get(270, "Screen", "get_width_mm")
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

// gdk_screen_get_window_stack
//
// [ result ] trans: everything
//
func (v Screen) GetWindowStack() (result g.List) {
	iv, err := _I.Get(271, "Screen", "get_window_stack")
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

// gdk_screen_is_composited
//
// [ result ] trans: nothing
//
func (v Screen) IsComposited() (result bool) {
	iv, err := _I.Get(272, "Screen", "is_composited")
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

// gdk_screen_list_visuals
//
// [ result ] trans: container
//
func (v Screen) ListVisuals() (result g.List) {
	iv, err := _I.Get(273, "Screen", "list_visuals")
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

// Deprecated
//
// gdk_screen_make_display_name
//
// [ result ] trans: everything
//
func (v Screen) MakeDisplayName() (result string) {
	iv, err := _I.Get(274, "Screen", "make_display_name")
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

// gdk_screen_set_font_options
//
// [ options ] trans: nothing
//
func (v Screen) SetFontOptions(options cairo.FontOptions) {
	iv, err := _I.Get(275, "Screen", "set_font_options")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_options := gi.NewPointerArgument(options.P)
	args := []gi.Argument{arg_v, arg_options}
	iv.Call(args, nil, nil)
}

// gdk_screen_set_resolution
//
// [ dpi ] trans: nothing
//
func (v Screen) SetResolution(dpi float64) {
	iv, err := _I.Get(276, "Screen", "set_resolution")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dpi := gi.NewDoubleArgument(dpi)
	args := []gi.Argument{arg_v, arg_dpi}
	iv.Call(args, nil, nil)
}

// Enum ScrollDirection
type ScrollDirectionEnum int

const (
	ScrollDirectionUp     ScrollDirectionEnum = 0
	ScrollDirectionDown   ScrollDirectionEnum = 1
	ScrollDirectionLeft   ScrollDirectionEnum = 2
	ScrollDirectionRight  ScrollDirectionEnum = 3
	ScrollDirectionSmooth ScrollDirectionEnum = 4
)

func ScrollDirectionGetType() gi.GType {
	ret := _I.GetGType(79, "ScrollDirection")
	return ret
}

// Object Seat
type Seat struct {
	g.Object
}

func WrapSeat(p unsafe.Pointer) (r Seat) { r.P = p; return }

type ISeat interface{ P_Seat() unsafe.Pointer }

func (v Seat) P_Seat() unsafe.Pointer { return v.P }
func SeatGetType() gi.GType {
	ret := _I.GetGType(80, "Seat")
	return ret
}

// gdk_seat_get_capabilities
//
// [ result ] trans: nothing
//
func (v Seat) GetCapabilities() (result SeatCapabilitiesFlags) {
	iv, err := _I.Get(277, "Seat", "get_capabilities")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = SeatCapabilitiesFlags(ret.Int())
	return
}

// gdk_seat_get_display
//
// [ result ] trans: nothing
//
func (v Seat) GetDisplay() (result Display) {
	iv, err := _I.Get(278, "Seat", "get_display")
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

// gdk_seat_get_keyboard
//
// [ result ] trans: nothing
//
func (v Seat) GetKeyboard() (result Device) {
	iv, err := _I.Get(279, "Seat", "get_keyboard")
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

// gdk_seat_get_pointer
//
// [ result ] trans: nothing
//
func (v Seat) GetPointer() (result Device) {
	iv, err := _I.Get(280, "Seat", "get_pointer")
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

// gdk_seat_get_slaves
//
// [ capabilities ] trans: nothing
//
// [ result ] trans: container
//
func (v Seat) GetSlaves(capabilities SeatCapabilitiesFlags) (result g.List) {
	iv, err := _I.Get(281, "Seat", "get_slaves")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_capabilities := gi.NewIntArgument(int(capabilities))
	args := []gi.Argument{arg_v, arg_capabilities}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_seat_grab
//
// [ window ] trans: nothing
//
// [ capabilities ] trans: nothing
//
// [ owner_events ] trans: nothing
//
// [ cursor ] trans: nothing
//
// [ event ] trans: nothing
//
// [ prepare_func ] trans: nothing
//
// [ prepare_func_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Seat) Grab(window IWindow, capabilities SeatCapabilitiesFlags, owner_events bool, cursor ICursor, event Event, prepare_func int /*TODO_TYPE CALLBACK*/, prepare_func_data unsafe.Pointer) (result GrabStatusEnum) {
	iv, err := _I.Get(282, "Seat", "grab")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	var tmp1 unsafe.Pointer
	if cursor != nil {
		tmp1 = cursor.P_Cursor()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_window := gi.NewPointerArgument(tmp)
	arg_capabilities := gi.NewIntArgument(int(capabilities))
	arg_owner_events := gi.NewBoolArgument(owner_events)
	arg_cursor := gi.NewPointerArgument(tmp1)
	arg_event := gi.NewPointerArgument(event.P)
	arg_prepare_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySeatGrabPrepareFunc()))
	arg_prepare_func_data := gi.NewPointerArgument(prepare_func_data)
	args := []gi.Argument{arg_v, arg_window, arg_capabilities, arg_owner_events, arg_cursor, arg_event, arg_prepare_func, arg_prepare_func_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = GrabStatusEnum(ret.Int())
	return
}

// gdk_seat_ungrab
//
func (v Seat) Ungrab() {
	iv, err := _I.Get(283, "Seat", "ungrab")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Flags SeatCapabilities
type SeatCapabilitiesFlags int

const (
	SeatCapabilitiesNone         SeatCapabilitiesFlags = 0
	SeatCapabilitiesPointer      SeatCapabilitiesFlags = 1
	SeatCapabilitiesTouch        SeatCapabilitiesFlags = 2
	SeatCapabilitiesTabletStylus SeatCapabilitiesFlags = 4
	SeatCapabilitiesKeyboard     SeatCapabilitiesFlags = 8
	SeatCapabilitiesAllPointing  SeatCapabilitiesFlags = 7
	SeatCapabilitiesAll          SeatCapabilitiesFlags = 15
)

func SeatCapabilitiesGetType() gi.GType {
	ret := _I.GetGType(81, "SeatCapabilities")
	return ret
}

type SeatGrabPrepareFuncStruct struct {
	F_seat   Seat
	F_window Window
}

func GetPointer_mySeatGrabPrepareFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGdkSeatGrabPrepareFunc())
}

//export myGdkSeatGrabPrepareFunc
func myGdkSeatGrabPrepareFunc(seat *C.GdkSeat, window *C.GdkWindow, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &SeatGrabPrepareFuncStruct{
		F_seat:   WrapSeat(unsafe.Pointer(seat)),
		F_window: WrapWindow(unsafe.Pointer(window)),
	}
	fn(args)
}

// Enum SettingAction
type SettingActionEnum int

const (
	SettingActionNew     SettingActionEnum = 0
	SettingActionChanged SettingActionEnum = 1
	SettingActionDeleted SettingActionEnum = 2
)

func SettingActionGetType() gi.GType {
	ret := _I.GetGType(82, "SettingAction")
	return ret
}

// Enum Status
type StatusEnum int

const (
	StatusOk         StatusEnum = 0
	StatusError      StatusEnum = -1
	StatusErrorParam StatusEnum = -2
	StatusErrorFile  StatusEnum = -3
	StatusErrorMem   StatusEnum = -4
)

func StatusGetType() gi.GType {
	ret := _I.GetGType(83, "Status")
	return ret
}

// Enum SubpixelLayout
type SubpixelLayoutEnum int

const (
	SubpixelLayoutUnknown       SubpixelLayoutEnum = 0
	SubpixelLayoutNone          SubpixelLayoutEnum = 1
	SubpixelLayoutHorizontalRgb SubpixelLayoutEnum = 2
	SubpixelLayoutHorizontalBgr SubpixelLayoutEnum = 3
	SubpixelLayoutVerticalRgb   SubpixelLayoutEnum = 4
	SubpixelLayoutVerticalBgr   SubpixelLayoutEnum = 5
)

func SubpixelLayoutGetType() gi.GType {
	ret := _I.GetGType(84, "SubpixelLayout")
	return ret
}

// Struct TimeCoord
type TimeCoord struct {
	P unsafe.Pointer
}

const SizeOfStructTimeCoord = 1032

func TimeCoordGetType() gi.GType {
	ret := _I.GetGType(85, "TimeCoord")
	return ret
}

// Enum TouchpadGesturePhase
type TouchpadGesturePhaseEnum int

const (
	TouchpadGesturePhaseBegin  TouchpadGesturePhaseEnum = 0
	TouchpadGesturePhaseUpdate TouchpadGesturePhaseEnum = 1
	TouchpadGesturePhaseEnd    TouchpadGesturePhaseEnum = 2
	TouchpadGesturePhaseCancel TouchpadGesturePhaseEnum = 3
)

func TouchpadGesturePhaseGetType() gi.GType {
	ret := _I.GetGType(86, "TouchpadGesturePhase")
	return ret
}

// Enum VisibilityState
type VisibilityStateEnum int

const (
	VisibilityStateUnobscured    VisibilityStateEnum = 0
	VisibilityStatePartial       VisibilityStateEnum = 1
	VisibilityStateFullyObscured VisibilityStateEnum = 2
)

func VisibilityStateGetType() gi.GType {
	ret := _I.GetGType(87, "VisibilityState")
	return ret
}

// Object Visual
type Visual struct {
	g.Object
}

func WrapVisual(p unsafe.Pointer) (r Visual) { r.P = p; return }

type IVisual interface{ P_Visual() unsafe.Pointer }

func (v Visual) P_Visual() unsafe.Pointer { return v.P }
func VisualGetType() gi.GType {
	ret := _I.GetGType(88, "Visual")
	return ret
}

// Deprecated
//
// Deprecated
//
// Deprecated
//
// Deprecated
//
// gdk_visual_get_best_with_both
//
// [ depth ] trans: nothing
//
// [ visual_type ] trans: nothing
//
// [ result ] trans: nothing
//
func VisualGetBestWithBoth1(depth int32, visual_type VisualTypeEnum) (result Visual) {
	iv, err := _I.Get(287, "Visual", "get_best_with_both")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_depth := gi.NewInt32Argument(depth)
	arg_visual_type := gi.NewIntArgument(int(visual_type))
	args := []gi.Argument{arg_depth, arg_visual_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// gdk_visual_get_best_with_depth
//
// [ depth ] trans: nothing
//
// [ result ] trans: nothing
//
func VisualGetBestWithDepth1(depth int32) (result Visual) {
	iv, err := _I.Get(288, "Visual", "get_best_with_depth")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_depth := gi.NewInt32Argument(depth)
	args := []gi.Argument{arg_depth}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// gdk_visual_get_best_with_type
//
// [ visual_type ] trans: nothing
//
// [ result ] trans: nothing
//
func VisualGetBestWithType1(visual_type VisualTypeEnum) (result Visual) {
	iv, err := _I.Get(289, "Visual", "get_best_with_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_visual_type := gi.NewIntArgument(int(visual_type))
	args := []gi.Argument{arg_visual_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// Deprecated
//
// gdk_visual_get_bits_per_rgb
//
// [ result ] trans: nothing
//
func (v Visual) GetBitsPerRgb() (result int32) {
	iv, err := _I.Get(291, "Visual", "get_bits_per_rgb")
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

// gdk_visual_get_blue_pixel_details
//
// [ mask ] trans: everything, dir: out
//
// [ shift ] trans: everything, dir: out
//
// [ precision ] trans: everything, dir: out
//
func (v Visual) GetBluePixelDetails() (mask uint32, shift int32, precision int32) {
	iv, err := _I.Get(292, "Visual", "get_blue_pixel_details")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_mask := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_shift := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_precision := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_mask, arg_shift, arg_precision}
	iv.Call(args, nil, &outArgs[0])
	mask = outArgs[0].Uint32()
	shift = outArgs[1].Int32()
	precision = outArgs[2].Int32()
	return
}

// Deprecated
//
// gdk_visual_get_byte_order
//
// [ result ] trans: nothing
//
func (v Visual) GetByteOrder() (result ByteOrderEnum) {
	iv, err := _I.Get(293, "Visual", "get_byte_order")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ByteOrderEnum(ret.Int())
	return
}

// Deprecated
//
// gdk_visual_get_colormap_size
//
// [ result ] trans: nothing
//
func (v Visual) GetColormapSize() (result int32) {
	iv, err := _I.Get(294, "Visual", "get_colormap_size")
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

// gdk_visual_get_depth
//
// [ result ] trans: nothing
//
func (v Visual) GetDepth() (result int32) {
	iv, err := _I.Get(295, "Visual", "get_depth")
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

// gdk_visual_get_green_pixel_details
//
// [ mask ] trans: everything, dir: out
//
// [ shift ] trans: everything, dir: out
//
// [ precision ] trans: everything, dir: out
//
func (v Visual) GetGreenPixelDetails() (mask uint32, shift int32, precision int32) {
	iv, err := _I.Get(296, "Visual", "get_green_pixel_details")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_mask := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_shift := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_precision := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_mask, arg_shift, arg_precision}
	iv.Call(args, nil, &outArgs[0])
	mask = outArgs[0].Uint32()
	shift = outArgs[1].Int32()
	precision = outArgs[2].Int32()
	return
}

// gdk_visual_get_red_pixel_details
//
// [ mask ] trans: everything, dir: out
//
// [ shift ] trans: everything, dir: out
//
// [ precision ] trans: everything, dir: out
//
func (v Visual) GetRedPixelDetails() (mask uint32, shift int32, precision int32) {
	iv, err := _I.Get(297, "Visual", "get_red_pixel_details")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_mask := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_shift := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_precision := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_mask, arg_shift, arg_precision}
	iv.Call(args, nil, &outArgs[0])
	mask = outArgs[0].Uint32()
	shift = outArgs[1].Int32()
	precision = outArgs[2].Int32()
	return
}

// gdk_visual_get_screen
//
// [ result ] trans: nothing
//
func (v Visual) GetScreen() (result Screen) {
	iv, err := _I.Get(298, "Visual", "get_screen")
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

// gdk_visual_get_visual_type
//
// [ result ] trans: nothing
//
func (v Visual) GetVisualType() (result VisualTypeEnum) {
	iv, err := _I.Get(299, "Visual", "get_visual_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = VisualTypeEnum(ret.Int())
	return
}

// Enum VisualType
type VisualTypeEnum int

const (
	VisualTypeStaticGray  VisualTypeEnum = 0
	VisualTypeGrayscale   VisualTypeEnum = 1
	VisualTypeStaticColor VisualTypeEnum = 2
	VisualTypePseudoColor VisualTypeEnum = 3
	VisualTypeTrueColor   VisualTypeEnum = 4
	VisualTypeDirectColor VisualTypeEnum = 5
)

func VisualTypeGetType() gi.GType {
	ret := _I.GetGType(89, "VisualType")
	return ret
}

// Flags WMDecoration
type WMDecorationFlags int

const (
	WMDecorationAll      WMDecorationFlags = 1
	WMDecorationBorder   WMDecorationFlags = 2
	WMDecorationResizeh  WMDecorationFlags = 4
	WMDecorationTitle    WMDecorationFlags = 8
	WMDecorationMenu     WMDecorationFlags = 16
	WMDecorationMinimize WMDecorationFlags = 32
	WMDecorationMaximize WMDecorationFlags = 64
)

func WMDecorationGetType() gi.GType {
	ret := _I.GetGType(90, "WMDecoration")
	return ret
}

// Flags WMFunction
type WMFunctionFlags int

const (
	WMFunctionAll      WMFunctionFlags = 1
	WMFunctionResize   WMFunctionFlags = 2
	WMFunctionMove     WMFunctionFlags = 4
	WMFunctionMinimize WMFunctionFlags = 8
	WMFunctionMaximize WMFunctionFlags = 16
	WMFunctionClose    WMFunctionFlags = 32
)

func WMFunctionGetType() gi.GType {
	ret := _I.GetGType(91, "WMFunction")
	return ret
}

// Object Window
type Window struct {
	g.Object
}

func WrapWindow(p unsafe.Pointer) (r Window) { r.P = p; return }

type IWindow interface{ P_Window() unsafe.Pointer }

func (v Window) P_Window() unsafe.Pointer { return v.P }
func WindowGetType() gi.GType {
	ret := _I.GetGType(92, "Window")
	return ret
}

// gdk_window_new
//
// [ parent ] trans: nothing
//
// [ attributes ] trans: nothing
//
// [ attributes_mask ] trans: nothing
//
// [ result ] trans: everything
//
func NewWindow(parent IWindow, attributes WindowAttr, attributes_mask WindowAttributesTypeFlags) (result Window) {
	iv, err := _I.Get(300, "Window", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if parent != nil {
		tmp = parent.P_Window()
	}
	arg_parent := gi.NewPointerArgument(tmp)
	arg_attributes := gi.NewPointerArgument(attributes.P)
	arg_attributes_mask := gi.NewIntArgument(int(attributes_mask))
	args := []gi.Argument{arg_parent, arg_attributes, arg_attributes_mask}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// gdk_window_at_pointer
//
// [ win_x ] trans: everything, dir: out
//
// [ win_y ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func WindowAtPointer1() (result Window, win_x int32, win_y int32) {
	iv, err := _I.Get(301, "Window", "at_pointer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_win_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_win_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_win_x, arg_win_y}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	win_x = outArgs[0].Int32()
	win_y = outArgs[1].Int32()
	result.P = ret.Pointer()
	return
}

// gdk_window_constrain_size
//
// [ geometry ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
// [ new_width ] trans: everything, dir: out
//
// [ new_height ] trans: everything, dir: out
//
func WindowConstrainSize1(geometry Geometry, flags WindowHintsFlags, width int32, height int32) (new_width int32, new_height int32) {
	iv, err := _I.Get(302, "Window", "constrain_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_geometry := gi.NewPointerArgument(geometry.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_new_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_new_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_geometry, arg_flags, arg_width, arg_height, arg_new_width, arg_new_height}
	iv.Call(args, nil, &outArgs[0])
	new_width = outArgs[0].Int32()
	new_height = outArgs[1].Int32()
	return
}

// Deprecated
//
// Deprecated
//
// gdk_window_set_debug_updates
//
// [ setting ] trans: nothing
//
func WindowSetDebugUpdates1(setting bool) {
	iv, err := _I.Get(304, "Window", "set_debug_updates")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_setting := gi.NewBoolArgument(setting)
	args := []gi.Argument{arg_setting}
	iv.Call(args, nil, nil)
}

// gdk_window_beep
//
func (v Window) Beep() {
	iv, err := _I.Get(305, "Window", "beep")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_begin_draw_frame
//
// [ region ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Window) BeginDrawFrame(region cairo.Region) (result DrawingContext) {
	iv, err := _I.Get(306, "Window", "begin_draw_frame")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_region := gi.NewPointerArgument(region.P)
	args := []gi.Argument{arg_v, arg_region}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_window_begin_move_drag
//
// [ button ] trans: nothing
//
// [ root_x ] trans: nothing
//
// [ root_y ] trans: nothing
//
// [ timestamp ] trans: nothing
//
func (v Window) BeginMoveDrag(button int32, root_x int32, root_y int32, timestamp uint32) {
	iv, err := _I.Get(307, "Window", "begin_move_drag")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_button := gi.NewInt32Argument(button)
	arg_root_x := gi.NewInt32Argument(root_x)
	arg_root_y := gi.NewInt32Argument(root_y)
	arg_timestamp := gi.NewUint32Argument(timestamp)
	args := []gi.Argument{arg_v, arg_button, arg_root_x, arg_root_y, arg_timestamp}
	iv.Call(args, nil, nil)
}

// gdk_window_begin_move_drag_for_device
//
// [ device ] trans: nothing
//
// [ button ] trans: nothing
//
// [ root_x ] trans: nothing
//
// [ root_y ] trans: nothing
//
// [ timestamp ] trans: nothing
//
func (v Window) BeginMoveDragForDevice(device IDevice, button int32, root_x int32, root_y int32, timestamp uint32) {
	iv, err := _I.Get(308, "Window", "begin_move_drag_for_device")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if device != nil {
		tmp = device.P_Device()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_device := gi.NewPointerArgument(tmp)
	arg_button := gi.NewInt32Argument(button)
	arg_root_x := gi.NewInt32Argument(root_x)
	arg_root_y := gi.NewInt32Argument(root_y)
	arg_timestamp := gi.NewUint32Argument(timestamp)
	args := []gi.Argument{arg_v, arg_device, arg_button, arg_root_x, arg_root_y, arg_timestamp}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_window_begin_paint_rect
//
// [ rectangle ] trans: nothing
//
func (v Window) BeginPaintRect(rectangle Rectangle) {
	iv, err := _I.Get(309, "Window", "begin_paint_rect")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_rectangle := gi.NewPointerArgument(rectangle.P)
	args := []gi.Argument{arg_v, arg_rectangle}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_window_begin_paint_region
//
// [ region ] trans: nothing
//
func (v Window) BeginPaintRegion(region cairo.Region) {
	iv, err := _I.Get(310, "Window", "begin_paint_region")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_region := gi.NewPointerArgument(region.P)
	args := []gi.Argument{arg_v, arg_region}
	iv.Call(args, nil, nil)
}

// gdk_window_begin_resize_drag
//
// [ edge ] trans: nothing
//
// [ button ] trans: nothing
//
// [ root_x ] trans: nothing
//
// [ root_y ] trans: nothing
//
// [ timestamp ] trans: nothing
//
func (v Window) BeginResizeDrag(edge WindowEdgeEnum, button int32, root_x int32, root_y int32, timestamp uint32) {
	iv, err := _I.Get(311, "Window", "begin_resize_drag")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_edge := gi.NewIntArgument(int(edge))
	arg_button := gi.NewInt32Argument(button)
	arg_root_x := gi.NewInt32Argument(root_x)
	arg_root_y := gi.NewInt32Argument(root_y)
	arg_timestamp := gi.NewUint32Argument(timestamp)
	args := []gi.Argument{arg_v, arg_edge, arg_button, arg_root_x, arg_root_y, arg_timestamp}
	iv.Call(args, nil, nil)
}

// gdk_window_begin_resize_drag_for_device
//
// [ edge ] trans: nothing
//
// [ device ] trans: nothing
//
// [ button ] trans: nothing
//
// [ root_x ] trans: nothing
//
// [ root_y ] trans: nothing
//
// [ timestamp ] trans: nothing
//
func (v Window) BeginResizeDragForDevice(edge WindowEdgeEnum, device IDevice, button int32, root_x int32, root_y int32, timestamp uint32) {
	iv, err := _I.Get(312, "Window", "begin_resize_drag_for_device")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if device != nil {
		tmp = device.P_Device()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_edge := gi.NewIntArgument(int(edge))
	arg_device := gi.NewPointerArgument(tmp)
	arg_button := gi.NewInt32Argument(button)
	arg_root_x := gi.NewInt32Argument(root_x)
	arg_root_y := gi.NewInt32Argument(root_y)
	arg_timestamp := gi.NewUint32Argument(timestamp)
	args := []gi.Argument{arg_v, arg_edge, arg_device, arg_button, arg_root_x, arg_root_y, arg_timestamp}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_window_configure_finished
//
func (v Window) ConfigureFinished() {
	iv, err := _I.Get(313, "Window", "configure_finished")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_coords_from_parent
//
// [ parent_x ] trans: nothing
//
// [ parent_y ] trans: nothing
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
func (v Window) CoordsFromParent(parent_x float64, parent_y float64) (x float64, y float64) {
	iv, err := _I.Get(314, "Window", "coords_from_parent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_parent_x := gi.NewDoubleArgument(parent_x)
	arg_parent_y := gi.NewDoubleArgument(parent_y)
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_parent_x, arg_parent_y, arg_x, arg_y}
	iv.Call(args, nil, &outArgs[0])
	x = outArgs[0].Double()
	y = outArgs[1].Double()
	return
}

// gdk_window_coords_to_parent
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ parent_x ] trans: everything, dir: out
//
// [ parent_y ] trans: everything, dir: out
//
func (v Window) CoordsToParent(x float64, y float64) (parent_x float64, parent_y float64) {
	iv, err := _I.Get(315, "Window", "coords_to_parent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_x := gi.NewDoubleArgument(x)
	arg_y := gi.NewDoubleArgument(y)
	arg_parent_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_parent_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_parent_x, arg_parent_y}
	iv.Call(args, nil, &outArgs[0])
	parent_x = outArgs[0].Double()
	parent_y = outArgs[1].Double()
	return
}

// gdk_window_create_gl_context
//
// [ result ] trans: everything
//
func (v Window) CreateGlContext() (result GLContext, err error) {
	iv, err := _I.Get(316, "Window", "create_gl_context")
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
	result.P = ret.Pointer()
	return
}

// gdk_window_create_similar_image_surface
//
// [ format ] trans: nothing
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
// [ scale ] trans: nothing
//
// [ result ] trans: everything
//
func (v Window) CreateSimilarImageSurface(format int32, width int32, height int32, scale int32) (result cairo.Surface) {
	iv, err := _I.Get(317, "Window", "create_similar_image_surface")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewInt32Argument(format)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_scale := gi.NewInt32Argument(scale)
	args := []gi.Argument{arg_v, arg_format, arg_width, arg_height, arg_scale}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_window_create_similar_surface
//
// [ content ] trans: nothing
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
// [ result ] trans: everything
//
func (v Window) CreateSimilarSurface(content cairo.ContentEnum, width int32, height int32) (result cairo.Surface) {
	iv, err := _I.Get(318, "Window", "create_similar_surface")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_content := gi.NewIntArgument(int(content))
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	args := []gi.Argument{arg_v, arg_content, arg_width, arg_height}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_window_deiconify
//
func (v Window) Deiconify() {
	iv, err := _I.Get(319, "Window", "deiconify")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_destroy
//
func (v Window) Destroy() {
	iv, err := _I.Get(320, "Window", "destroy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_destroy_notify
//
func (v Window) DestroyNotify() {
	iv, err := _I.Get(321, "Window", "destroy_notify")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_window_enable_synchronized_configure
//
func (v Window) EnableSynchronizedConfigure() {
	iv, err := _I.Get(322, "Window", "enable_synchronized_configure")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_end_draw_frame
//
// [ context ] trans: nothing
//
func (v Window) EndDrawFrame(context IDrawingContext) {
	iv, err := _I.Get(323, "Window", "end_draw_frame")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_DrawingContext()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_context := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_context}
	iv.Call(args, nil, nil)
}

// gdk_window_end_paint
//
func (v Window) EndPaint() {
	iv, err := _I.Get(324, "Window", "end_paint")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_ensure_native
//
// [ result ] trans: nothing
//
func (v Window) EnsureNative() (result bool) {
	iv, err := _I.Get(325, "Window", "ensure_native")
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

// Deprecated
//
// gdk_window_flush
//
func (v Window) Flush() {
	iv, err := _I.Get(326, "Window", "flush")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_focus
//
// [ timestamp ] trans: nothing
//
func (v Window) Focus(timestamp uint32) {
	iv, err := _I.Get(327, "Window", "focus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timestamp := gi.NewUint32Argument(timestamp)
	args := []gi.Argument{arg_v, arg_timestamp}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_window_freeze_toplevel_updates_libgtk_only
//
func (v Window) FreezeToplevelUpdatesLibgtkOnly() {
	iv, err := _I.Get(328, "Window", "freeze_toplevel_updates_libgtk_only")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_freeze_updates
//
func (v Window) FreezeUpdates() {
	iv, err := _I.Get(329, "Window", "freeze_updates")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_fullscreen
//
func (v Window) Fullscreen() {
	iv, err := _I.Get(330, "Window", "fullscreen")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_fullscreen_on_monitor
//
// [ monitor ] trans: nothing
//
func (v Window) FullscreenOnMonitor(monitor int32) {
	iv, err := _I.Get(331, "Window", "fullscreen_on_monitor")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_monitor := gi.NewInt32Argument(monitor)
	args := []gi.Argument{arg_v, arg_monitor}
	iv.Call(args, nil, nil)
}

// gdk_window_geometry_changed
//
func (v Window) GeometryChanged() {
	iv, err := _I.Get(332, "Window", "geometry_changed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_get_accept_focus
//
// [ result ] trans: nothing
//
func (v Window) GetAcceptFocus() (result bool) {
	iv, err := _I.Get(333, "Window", "get_accept_focus")
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

// Deprecated
//
// gdk_window_get_background_pattern
//
// [ result ] trans: nothing
//
func (v Window) GetBackgroundPattern() (result cairo.Pattern) {
	iv, err := _I.Get(334, "Window", "get_background_pattern")
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

// gdk_window_get_children
//
// [ result ] trans: container
//
func (v Window) GetChildren() (result g.List) {
	iv, err := _I.Get(335, "Window", "get_children")
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

// gdk_window_get_children_with_user_data
//
// [ user_data ] trans: nothing
//
// [ result ] trans: container
//
func (v Window) GetChildrenWithUserData(user_data unsafe.Pointer) (result g.List) {
	iv, err := _I.Get(336, "Window", "get_children_with_user_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_window_get_clip_region
//
// [ result ] trans: everything
//
func (v Window) GetClipRegion() (result cairo.Region) {
	iv, err := _I.Get(337, "Window", "get_clip_region")
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

// Deprecated
//
// gdk_window_get_composited
//
// [ result ] trans: nothing
//
func (v Window) GetComposited() (result bool) {
	iv, err := _I.Get(338, "Window", "get_composited")
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

// gdk_window_get_cursor
//
// [ result ] trans: nothing
//
func (v Window) GetCursor() (result Cursor) {
	iv, err := _I.Get(339, "Window", "get_cursor")
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

// gdk_window_get_decorations
//
// [ decorations ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Window) GetDecorations() (result bool, decorations WMDecorationFlags) {
	iv, err := _I.Get(340, "Window", "get_decorations")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_decorations := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_decorations}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	decorations = WMDecorationFlags(outArgs[0].Int())
	result = ret.Bool()
	return
}

// gdk_window_get_device_cursor
//
// [ device ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Window) GetDeviceCursor(device IDevice) (result Cursor) {
	iv, err := _I.Get(341, "Window", "get_device_cursor")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if device != nil {
		tmp = device.P_Device()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_device := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_device}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_window_get_device_events
//
// [ device ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Window) GetDeviceEvents(device IDevice) (result EventMaskFlags) {
	iv, err := _I.Get(342, "Window", "get_device_events")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if device != nil {
		tmp = device.P_Device()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_device := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_device}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = EventMaskFlags(ret.Int())
	return
}

// gdk_window_get_device_position
//
// [ device ] trans: nothing
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
// [ mask ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Window) GetDevicePosition(device IDevice) (result Window, x int32, y int32, mask ModifierTypeFlags) {
	iv, err := _I.Get(343, "Window", "get_device_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	var tmp unsafe.Pointer
	if device != nil {
		tmp = device.P_Device()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_device := gi.NewPointerArgument(tmp)
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_mask := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_device, arg_x, arg_y, arg_mask}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	x = outArgs[0].Int32()
	y = outArgs[1].Int32()
	mask = ModifierTypeFlags(outArgs[2].Int())
	result.P = ret.Pointer()
	return
}

// gdk_window_get_device_position_double
//
// [ device ] trans: nothing
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
// [ mask ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Window) GetDevicePositionDouble(device IDevice) (result Window, x float64, y float64, mask ModifierTypeFlags) {
	iv, err := _I.Get(344, "Window", "get_device_position_double")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	var tmp unsafe.Pointer
	if device != nil {
		tmp = device.P_Device()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_device := gi.NewPointerArgument(tmp)
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_mask := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_device, arg_x, arg_y, arg_mask}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	x = outArgs[0].Double()
	y = outArgs[1].Double()
	mask = ModifierTypeFlags(outArgs[2].Int())
	result.P = ret.Pointer()
	return
}

// gdk_window_get_display
//
// [ result ] trans: nothing
//
func (v Window) GetDisplay() (result Display) {
	iv, err := _I.Get(345, "Window", "get_display")
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

// gdk_window_get_drag_protocol
//
// [ target ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Window) GetDragProtocol() (result DragProtocolEnum, target Window) {
	iv, err := _I.Get(346, "Window", "get_drag_protocol")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_target := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	target.P = outArgs[0].Pointer()
	result = DragProtocolEnum(ret.Int())
	return
}

// gdk_window_get_effective_parent
//
// [ result ] trans: nothing
//
func (v Window) GetEffectiveParent() (result Window) {
	iv, err := _I.Get(347, "Window", "get_effective_parent")
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

// gdk_window_get_effective_toplevel
//
// [ result ] trans: nothing
//
func (v Window) GetEffectiveToplevel() (result Window) {
	iv, err := _I.Get(348, "Window", "get_effective_toplevel")
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

// gdk_window_get_event_compression
//
// [ result ] trans: nothing
//
func (v Window) GetEventCompression() (result bool) {
	iv, err := _I.Get(349, "Window", "get_event_compression")
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

// gdk_window_get_events
//
// [ result ] trans: nothing
//
func (v Window) GetEvents() (result EventMaskFlags) {
	iv, err := _I.Get(350, "Window", "get_events")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = EventMaskFlags(ret.Int())
	return
}

// gdk_window_get_focus_on_map
//
// [ result ] trans: nothing
//
func (v Window) GetFocusOnMap() (result bool) {
	iv, err := _I.Get(351, "Window", "get_focus_on_map")
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

// gdk_window_get_frame_clock
//
// [ result ] trans: nothing
//
func (v Window) GetFrameClock() (result FrameClock) {
	iv, err := _I.Get(352, "Window", "get_frame_clock")
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

// gdk_window_get_frame_extents
//
// [ rect ] trans: nothing, dir: out
//
func (v Window) GetFrameExtents(rect Rectangle) {
	iv, err := _I.Get(353, "Window", "get_frame_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_rect := gi.NewPointerArgument(rect.P)
	args := []gi.Argument{arg_v, arg_rect}
	iv.Call(args, nil, nil)
}

// gdk_window_get_fullscreen_mode
//
// [ result ] trans: nothing
//
func (v Window) GetFullscreenMode() (result FullscreenModeEnum) {
	iv, err := _I.Get(354, "Window", "get_fullscreen_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FullscreenModeEnum(ret.Int())
	return
}

// gdk_window_get_geometry
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
// [ width ] trans: everything, dir: out
//
// [ height ] trans: everything, dir: out
//
func (v Window) GetGeometry() (x int32, y int32, width int32, height int32) {
	iv, err := _I.Get(355, "Window", "get_geometry")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_width, arg_height}
	iv.Call(args, nil, &outArgs[0])
	x = outArgs[0].Int32()
	y = outArgs[1].Int32()
	width = outArgs[2].Int32()
	height = outArgs[3].Int32()
	return
}

// gdk_window_get_group
//
// [ result ] trans: nothing
//
func (v Window) GetGroup() (result Window) {
	iv, err := _I.Get(356, "Window", "get_group")
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

// gdk_window_get_height
//
// [ result ] trans: nothing
//
func (v Window) GetHeight() (result int32) {
	iv, err := _I.Get(357, "Window", "get_height")
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

// gdk_window_get_modal_hint
//
// [ result ] trans: nothing
//
func (v Window) GetModalHint() (result bool) {
	iv, err := _I.Get(358, "Window", "get_modal_hint")
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

// gdk_window_get_origin
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Window) GetOrigin() (result int32, x int32, y int32) {
	iv, err := _I.Get(359, "Window", "get_origin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_x, arg_y}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	x = outArgs[0].Int32()
	y = outArgs[1].Int32()
	result = ret.Int32()
	return
}

// gdk_window_get_parent
//
// [ result ] trans: nothing
//
func (v Window) GetParent() (result Window) {
	iv, err := _I.Get(360, "Window", "get_parent")
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

// gdk_window_get_pass_through
//
// [ result ] trans: nothing
//
func (v Window) GetPassThrough() (result bool) {
	iv, err := _I.Get(361, "Window", "get_pass_through")
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

// Deprecated
//
// gdk_window_get_pointer
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
// [ mask ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func (v Window) GetPointer() (result Window, x int32, y int32, mask ModifierTypeFlags) {
	iv, err := _I.Get(362, "Window", "get_pointer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_mask := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_mask}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	x = outArgs[0].Int32()
	y = outArgs[1].Int32()
	mask = ModifierTypeFlags(outArgs[2].Int())
	result.P = ret.Pointer()
	return
}

// gdk_window_get_position
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
func (v Window) GetPosition() (x int32, y int32) {
	iv, err := _I.Get(363, "Window", "get_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_x, arg_y}
	iv.Call(args, nil, &outArgs[0])
	x = outArgs[0].Int32()
	y = outArgs[1].Int32()
	return
}

// gdk_window_get_root_coords
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ root_x ] trans: everything, dir: out
//
// [ root_y ] trans: everything, dir: out
//
func (v Window) GetRootCoords(x int32, y int32) (root_x int32, root_y int32) {
	iv, err := _I.Get(364, "Window", "get_root_coords")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_root_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_root_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_root_x, arg_root_y}
	iv.Call(args, nil, &outArgs[0])
	root_x = outArgs[0].Int32()
	root_y = outArgs[1].Int32()
	return
}

// gdk_window_get_root_origin
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
func (v Window) GetRootOrigin() (x int32, y int32) {
	iv, err := _I.Get(365, "Window", "get_root_origin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_x, arg_y}
	iv.Call(args, nil, &outArgs[0])
	x = outArgs[0].Int32()
	y = outArgs[1].Int32()
	return
}

// gdk_window_get_scale_factor
//
// [ result ] trans: nothing
//
func (v Window) GetScaleFactor() (result int32) {
	iv, err := _I.Get(366, "Window", "get_scale_factor")
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

// gdk_window_get_screen
//
// [ result ] trans: nothing
//
func (v Window) GetScreen() (result Screen) {
	iv, err := _I.Get(367, "Window", "get_screen")
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

// gdk_window_get_source_events
//
// [ source ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Window) GetSourceEvents(source InputSourceEnum) (result EventMaskFlags) {
	iv, err := _I.Get(368, "Window", "get_source_events")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_source := gi.NewIntArgument(int(source))
	args := []gi.Argument{arg_v, arg_source}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = EventMaskFlags(ret.Int())
	return
}

// gdk_window_get_state
//
// [ result ] trans: nothing
//
func (v Window) GetState() (result WindowStateFlags) {
	iv, err := _I.Get(369, "Window", "get_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = WindowStateFlags(ret.Int())
	return
}

// gdk_window_get_support_multidevice
//
// [ result ] trans: nothing
//
func (v Window) GetSupportMultidevice() (result bool) {
	iv, err := _I.Get(370, "Window", "get_support_multidevice")
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

// gdk_window_get_toplevel
//
// [ result ] trans: nothing
//
func (v Window) GetToplevel() (result Window) {
	iv, err := _I.Get(371, "Window", "get_toplevel")
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

// gdk_window_get_type_hint
//
// [ result ] trans: nothing
//
func (v Window) GetTypeHint() (result WindowTypeHintEnum) {
	iv, err := _I.Get(372, "Window", "get_type_hint")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = WindowTypeHintEnum(ret.Int())
	return
}

// gdk_window_get_update_area
//
// [ result ] trans: everything
//
func (v Window) GetUpdateArea() (result cairo.Region) {
	iv, err := _I.Get(373, "Window", "get_update_area")
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

// gdk_window_get_user_data
//
// [ data ] trans: everything, dir: out
//
func (v Window) GetUserData() (data unsafe.Pointer) {
	iv, err := _I.Get(374, "Window", "get_user_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_data}
	iv.Call(args, nil, &outArgs[0])
	data = outArgs[0].Pointer()
	return
}

// gdk_window_get_visible_region
//
// [ result ] trans: everything
//
func (v Window) GetVisibleRegion() (result cairo.Region) {
	iv, err := _I.Get(375, "Window", "get_visible_region")
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

// gdk_window_get_visual
//
// [ result ] trans: nothing
//
func (v Window) GetVisual() (result Visual) {
	iv, err := _I.Get(376, "Window", "get_visual")
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

// gdk_window_get_width
//
// [ result ] trans: nothing
//
func (v Window) GetWidth() (result int32) {
	iv, err := _I.Get(377, "Window", "get_width")
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

// gdk_window_get_window_type
//
// [ result ] trans: nothing
//
func (v Window) GetWindowType() (result WindowTypeEnum) {
	iv, err := _I.Get(378, "Window", "get_window_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = WindowTypeEnum(ret.Int())
	return
}

// gdk_window_has_native
//
// [ result ] trans: nothing
//
func (v Window) HasNative() (result bool) {
	iv, err := _I.Get(379, "Window", "has_native")
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

// gdk_window_hide
//
func (v Window) Hide() {
	iv, err := _I.Get(380, "Window", "hide")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_iconify
//
func (v Window) Iconify() {
	iv, err := _I.Get(381, "Window", "iconify")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_input_shape_combine_region
//
// [ shape_region ] trans: nothing
//
// [ offset_x ] trans: nothing
//
// [ offset_y ] trans: nothing
//
func (v Window) InputShapeCombineRegion(shape_region cairo.Region, offset_x int32, offset_y int32) {
	iv, err := _I.Get(382, "Window", "input_shape_combine_region")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_shape_region := gi.NewPointerArgument(shape_region.P)
	arg_offset_x := gi.NewInt32Argument(offset_x)
	arg_offset_y := gi.NewInt32Argument(offset_y)
	args := []gi.Argument{arg_v, arg_shape_region, arg_offset_x, arg_offset_y}
	iv.Call(args, nil, nil)
}

// gdk_window_invalidate_maybe_recurse
//
// [ region ] trans: nothing
//
// [ child_func ] trans: nothing
//
// [ user_data ] trans: nothing
//
func (v Window) InvalidateMaybeRecurse(region cairo.Region, child_func int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(383, "Window", "invalidate_maybe_recurse")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_region := gi.NewPointerArgument(region.P)
	arg_child_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myWindowChildFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_region, arg_child_func, arg_user_data}
	iv.Call(args, nil, nil)
}

// gdk_window_invalidate_rect
//
// [ rect ] trans: nothing
//
// [ invalidate_children ] trans: nothing
//
func (v Window) InvalidateRect(rect Rectangle, invalidate_children bool) {
	iv, err := _I.Get(384, "Window", "invalidate_rect")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_rect := gi.NewPointerArgument(rect.P)
	arg_invalidate_children := gi.NewBoolArgument(invalidate_children)
	args := []gi.Argument{arg_v, arg_rect, arg_invalidate_children}
	iv.Call(args, nil, nil)
}

// gdk_window_invalidate_region
//
// [ region ] trans: nothing
//
// [ invalidate_children ] trans: nothing
//
func (v Window) InvalidateRegion(region cairo.Region, invalidate_children bool) {
	iv, err := _I.Get(385, "Window", "invalidate_region")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_region := gi.NewPointerArgument(region.P)
	arg_invalidate_children := gi.NewBoolArgument(invalidate_children)
	args := []gi.Argument{arg_v, arg_region, arg_invalidate_children}
	iv.Call(args, nil, nil)
}

// gdk_window_is_destroyed
//
// [ result ] trans: nothing
//
func (v Window) IsDestroyed() (result bool) {
	iv, err := _I.Get(386, "Window", "is_destroyed")
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

// gdk_window_is_input_only
//
// [ result ] trans: nothing
//
func (v Window) IsInputOnly() (result bool) {
	iv, err := _I.Get(387, "Window", "is_input_only")
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

// gdk_window_is_shaped
//
// [ result ] trans: nothing
//
func (v Window) IsShaped() (result bool) {
	iv, err := _I.Get(388, "Window", "is_shaped")
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

// gdk_window_is_viewable
//
// [ result ] trans: nothing
//
func (v Window) IsViewable() (result bool) {
	iv, err := _I.Get(389, "Window", "is_viewable")
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

// gdk_window_is_visible
//
// [ result ] trans: nothing
//
func (v Window) IsVisible() (result bool) {
	iv, err := _I.Get(390, "Window", "is_visible")
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

// gdk_window_lower
//
func (v Window) Lower() {
	iv, err := _I.Get(391, "Window", "lower")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_mark_paint_from_clip
//
// [ cr ] trans: nothing
//
func (v Window) MarkPaintFromClip(cr cairo.Context) {
	iv, err := _I.Get(392, "Window", "mark_paint_from_clip")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_cr := gi.NewPointerArgument(cr.P)
	args := []gi.Argument{arg_v, arg_cr}
	iv.Call(args, nil, nil)
}

// gdk_window_maximize
//
func (v Window) Maximize() {
	iv, err := _I.Get(393, "Window", "maximize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_merge_child_input_shapes
//
func (v Window) MergeChildInputShapes() {
	iv, err := _I.Get(394, "Window", "merge_child_input_shapes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_merge_child_shapes
//
func (v Window) MergeChildShapes() {
	iv, err := _I.Get(395, "Window", "merge_child_shapes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_move
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
func (v Window) Move(x int32, y int32) {
	iv, err := _I.Get(396, "Window", "move")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	args := []gi.Argument{arg_v, arg_x, arg_y}
	iv.Call(args, nil, nil)
}

// gdk_window_move_region
//
// [ region ] trans: nothing
//
// [ dx ] trans: nothing
//
// [ dy ] trans: nothing
//
func (v Window) MoveRegion(region cairo.Region, dx int32, dy int32) {
	iv, err := _I.Get(397, "Window", "move_region")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_region := gi.NewPointerArgument(region.P)
	arg_dx := gi.NewInt32Argument(dx)
	arg_dy := gi.NewInt32Argument(dy)
	args := []gi.Argument{arg_v, arg_region, arg_dx, arg_dy}
	iv.Call(args, nil, nil)
}

// gdk_window_move_resize
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
func (v Window) MoveResize(x int32, y int32, width int32, height int32) {
	iv, err := _I.Get(398, "Window", "move_resize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_width, arg_height}
	iv.Call(args, nil, nil)
}

// gdk_window_move_to_rect
//
// [ rect ] trans: nothing
//
// [ rect_anchor ] trans: nothing
//
// [ window_anchor ] trans: nothing
//
// [ anchor_hints ] trans: nothing
//
// [ rect_anchor_dx ] trans: nothing
//
// [ rect_anchor_dy ] trans: nothing
//
func (v Window) MoveToRect(rect Rectangle, rect_anchor GravityEnum, window_anchor GravityEnum, anchor_hints AnchorHintsFlags, rect_anchor_dx int32, rect_anchor_dy int32) {
	iv, err := _I.Get(399, "Window", "move_to_rect")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_rect := gi.NewPointerArgument(rect.P)
	arg_rect_anchor := gi.NewIntArgument(int(rect_anchor))
	arg_window_anchor := gi.NewIntArgument(int(window_anchor))
	arg_anchor_hints := gi.NewIntArgument(int(anchor_hints))
	arg_rect_anchor_dx := gi.NewInt32Argument(rect_anchor_dx)
	arg_rect_anchor_dy := gi.NewInt32Argument(rect_anchor_dy)
	args := []gi.Argument{arg_v, arg_rect, arg_rect_anchor, arg_window_anchor, arg_anchor_hints, arg_rect_anchor_dx, arg_rect_anchor_dy}
	iv.Call(args, nil, nil)
}

// gdk_window_peek_children
//
// [ result ] trans: nothing
//
func (v Window) PeekChildren() (result g.List) {
	iv, err := _I.Get(400, "Window", "peek_children")
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

// Deprecated
//
// gdk_window_process_updates
//
// [ update_children ] trans: nothing
//
func (v Window) ProcessUpdates(update_children bool) {
	iv, err := _I.Get(401, "Window", "process_updates")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_update_children := gi.NewBoolArgument(update_children)
	args := []gi.Argument{arg_v, arg_update_children}
	iv.Call(args, nil, nil)
}

// gdk_window_raise
//
func (v Window) Raise() {
	iv, err := _I.Get(402, "Window", "raise")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_register_dnd
//
func (v Window) RegisterDnd() {
	iv, err := _I.Get(403, "Window", "register_dnd")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_reparent
//
// [ new_parent ] trans: nothing
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
func (v Window) Reparent(new_parent IWindow, x int32, y int32) {
	iv, err := _I.Get(404, "Window", "reparent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if new_parent != nil {
		tmp = new_parent.P_Window()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_new_parent := gi.NewPointerArgument(tmp)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	args := []gi.Argument{arg_v, arg_new_parent, arg_x, arg_y}
	iv.Call(args, nil, nil)
}

// gdk_window_resize
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
func (v Window) Resize(width int32, height int32) {
	iv, err := _I.Get(405, "Window", "resize")
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

// gdk_window_restack
//
// [ sibling ] trans: nothing
//
// [ above ] trans: nothing
//
func (v Window) Restack(sibling IWindow, above bool) {
	iv, err := _I.Get(406, "Window", "restack")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if sibling != nil {
		tmp = sibling.P_Window()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_sibling := gi.NewPointerArgument(tmp)
	arg_above := gi.NewBoolArgument(above)
	args := []gi.Argument{arg_v, arg_sibling, arg_above}
	iv.Call(args, nil, nil)
}

// gdk_window_scroll
//
// [ dx ] trans: nothing
//
// [ dy ] trans: nothing
//
func (v Window) Scroll(dx int32, dy int32) {
	iv, err := _I.Get(407, "Window", "scroll")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dx := gi.NewInt32Argument(dx)
	arg_dy := gi.NewInt32Argument(dy)
	args := []gi.Argument{arg_v, arg_dx, arg_dy}
	iv.Call(args, nil, nil)
}

// gdk_window_set_accept_focus
//
// [ accept_focus ] trans: nothing
//
func (v Window) SetAcceptFocus(accept_focus bool) {
	iv, err := _I.Get(408, "Window", "set_accept_focus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_accept_focus := gi.NewBoolArgument(accept_focus)
	args := []gi.Argument{arg_v, arg_accept_focus}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_window_set_background
//
// [ color ] trans: nothing
//
func (v Window) SetBackground(color Color) {
	iv, err := _I.Get(409, "Window", "set_background")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_color := gi.NewPointerArgument(color.P)
	args := []gi.Argument{arg_v, arg_color}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_window_set_background_pattern
//
// [ pattern ] trans: nothing
//
func (v Window) SetBackgroundPattern(pattern cairo.Pattern) {
	iv, err := _I.Get(410, "Window", "set_background_pattern")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pattern := gi.NewPointerArgument(pattern.P)
	args := []gi.Argument{arg_v, arg_pattern}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_window_set_background_rgba
//
// [ rgba ] trans: nothing
//
func (v Window) SetBackgroundRgba(rgba RGBA) {
	iv, err := _I.Get(411, "Window", "set_background_rgba")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_rgba := gi.NewPointerArgument(rgba.P)
	args := []gi.Argument{arg_v, arg_rgba}
	iv.Call(args, nil, nil)
}

// gdk_window_set_child_input_shapes
//
func (v Window) SetChildInputShapes() {
	iv, err := _I.Get(412, "Window", "set_child_input_shapes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_set_child_shapes
//
func (v Window) SetChildShapes() {
	iv, err := _I.Get(413, "Window", "set_child_shapes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_window_set_composited
//
// [ composited ] trans: nothing
//
func (v Window) SetComposited(composited bool) {
	iv, err := _I.Get(414, "Window", "set_composited")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_composited := gi.NewBoolArgument(composited)
	args := []gi.Argument{arg_v, arg_composited}
	iv.Call(args, nil, nil)
}

// gdk_window_set_cursor
//
// [ cursor ] trans: nothing
//
func (v Window) SetCursor(cursor ICursor) {
	iv, err := _I.Get(415, "Window", "set_cursor")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cursor != nil {
		tmp = cursor.P_Cursor()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_cursor := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_cursor}
	iv.Call(args, nil, nil)
}

// gdk_window_set_decorations
//
// [ decorations ] trans: nothing
//
func (v Window) SetDecorations(decorations WMDecorationFlags) {
	iv, err := _I.Get(416, "Window", "set_decorations")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_decorations := gi.NewIntArgument(int(decorations))
	args := []gi.Argument{arg_v, arg_decorations}
	iv.Call(args, nil, nil)
}

// gdk_window_set_device_cursor
//
// [ device ] trans: nothing
//
// [ cursor ] trans: nothing
//
func (v Window) SetDeviceCursor(device IDevice, cursor ICursor) {
	iv, err := _I.Get(417, "Window", "set_device_cursor")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if device != nil {
		tmp = device.P_Device()
	}
	var tmp1 unsafe.Pointer
	if cursor != nil {
		tmp1 = cursor.P_Cursor()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_device := gi.NewPointerArgument(tmp)
	arg_cursor := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_v, arg_device, arg_cursor}
	iv.Call(args, nil, nil)
}

// gdk_window_set_device_events
//
// [ device ] trans: nothing
//
// [ event_mask ] trans: nothing
//
func (v Window) SetDeviceEvents(device IDevice, event_mask EventMaskFlags) {
	iv, err := _I.Get(418, "Window", "set_device_events")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if device != nil {
		tmp = device.P_Device()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_device := gi.NewPointerArgument(tmp)
	arg_event_mask := gi.NewIntArgument(int(event_mask))
	args := []gi.Argument{arg_v, arg_device, arg_event_mask}
	iv.Call(args, nil, nil)
}

// gdk_window_set_event_compression
//
// [ event_compression ] trans: nothing
//
func (v Window) SetEventCompression(event_compression bool) {
	iv, err := _I.Get(419, "Window", "set_event_compression")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_event_compression := gi.NewBoolArgument(event_compression)
	args := []gi.Argument{arg_v, arg_event_compression}
	iv.Call(args, nil, nil)
}

// gdk_window_set_events
//
// [ event_mask ] trans: nothing
//
func (v Window) SetEvents(event_mask EventMaskFlags) {
	iv, err := _I.Get(420, "Window", "set_events")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_event_mask := gi.NewIntArgument(int(event_mask))
	args := []gi.Argument{arg_v, arg_event_mask}
	iv.Call(args, nil, nil)
}

// gdk_window_set_focus_on_map
//
// [ focus_on_map ] trans: nothing
//
func (v Window) SetFocusOnMap(focus_on_map bool) {
	iv, err := _I.Get(421, "Window", "set_focus_on_map")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_focus_on_map := gi.NewBoolArgument(focus_on_map)
	args := []gi.Argument{arg_v, arg_focus_on_map}
	iv.Call(args, nil, nil)
}

// gdk_window_set_fullscreen_mode
//
// [ mode ] trans: nothing
//
func (v Window) SetFullscreenMode(mode FullscreenModeEnum) {
	iv, err := _I.Get(422, "Window", "set_fullscreen_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_v, arg_mode}
	iv.Call(args, nil, nil)
}

// gdk_window_set_functions
//
// [ functions ] trans: nothing
//
func (v Window) SetFunctions(functions WMFunctionFlags) {
	iv, err := _I.Get(423, "Window", "set_functions")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_functions := gi.NewIntArgument(int(functions))
	args := []gi.Argument{arg_v, arg_functions}
	iv.Call(args, nil, nil)
}

// gdk_window_set_geometry_hints
//
// [ geometry ] trans: nothing
//
// [ geom_mask ] trans: nothing
//
func (v Window) SetGeometryHints(geometry Geometry, geom_mask WindowHintsFlags) {
	iv, err := _I.Get(424, "Window", "set_geometry_hints")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_geometry := gi.NewPointerArgument(geometry.P)
	arg_geom_mask := gi.NewIntArgument(int(geom_mask))
	args := []gi.Argument{arg_v, arg_geometry, arg_geom_mask}
	iv.Call(args, nil, nil)
}

// gdk_window_set_group
//
// [ leader ] trans: nothing
//
func (v Window) SetGroup(leader IWindow) {
	iv, err := _I.Get(425, "Window", "set_group")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if leader != nil {
		tmp = leader.P_Window()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_leader := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_leader}
	iv.Call(args, nil, nil)
}

// gdk_window_set_icon_list
//
// [ pixbufs ] trans: nothing
//
func (v Window) SetIconList(pixbufs g.List) {
	iv, err := _I.Get(426, "Window", "set_icon_list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pixbufs := gi.NewPointerArgument(pixbufs.P)
	args := []gi.Argument{arg_v, arg_pixbufs}
	iv.Call(args, nil, nil)
}

// gdk_window_set_icon_name
//
// [ name ] trans: nothing
//
func (v Window) SetIconName(name string) {
	iv, err := _I.Get(427, "Window", "set_icon_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_v, arg_name}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
}

// gdk_window_set_keep_above
//
// [ setting ] trans: nothing
//
func (v Window) SetKeepAbove(setting bool) {
	iv, err := _I.Get(428, "Window", "set_keep_above")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_setting := gi.NewBoolArgument(setting)
	args := []gi.Argument{arg_v, arg_setting}
	iv.Call(args, nil, nil)
}

// gdk_window_set_keep_below
//
// [ setting ] trans: nothing
//
func (v Window) SetKeepBelow(setting bool) {
	iv, err := _I.Get(429, "Window", "set_keep_below")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_setting := gi.NewBoolArgument(setting)
	args := []gi.Argument{arg_v, arg_setting}
	iv.Call(args, nil, nil)
}

// gdk_window_set_modal_hint
//
// [ modal ] trans: nothing
//
func (v Window) SetModalHint(modal bool) {
	iv, err := _I.Get(430, "Window", "set_modal_hint")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_modal := gi.NewBoolArgument(modal)
	args := []gi.Argument{arg_v, arg_modal}
	iv.Call(args, nil, nil)
}

// gdk_window_set_opacity
//
// [ opacity ] trans: nothing
//
func (v Window) SetOpacity(opacity float64) {
	iv, err := _I.Get(431, "Window", "set_opacity")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_opacity := gi.NewDoubleArgument(opacity)
	args := []gi.Argument{arg_v, arg_opacity}
	iv.Call(args, nil, nil)
}

// gdk_window_set_opaque_region
//
// [ region ] trans: nothing
//
func (v Window) SetOpaqueRegion(region cairo.Region) {
	iv, err := _I.Get(432, "Window", "set_opaque_region")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_region := gi.NewPointerArgument(region.P)
	args := []gi.Argument{arg_v, arg_region}
	iv.Call(args, nil, nil)
}

// gdk_window_set_override_redirect
//
// [ override_redirect ] trans: nothing
//
func (v Window) SetOverrideRedirect(override_redirect bool) {
	iv, err := _I.Get(433, "Window", "set_override_redirect")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_override_redirect := gi.NewBoolArgument(override_redirect)
	args := []gi.Argument{arg_v, arg_override_redirect}
	iv.Call(args, nil, nil)
}

// gdk_window_set_pass_through
//
// [ pass_through ] trans: nothing
//
func (v Window) SetPassThrough(pass_through bool) {
	iv, err := _I.Get(434, "Window", "set_pass_through")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pass_through := gi.NewBoolArgument(pass_through)
	args := []gi.Argument{arg_v, arg_pass_through}
	iv.Call(args, nil, nil)
}

// gdk_window_set_role
//
// [ role ] trans: nothing
//
func (v Window) SetRole(role string) {
	iv, err := _I.Get(435, "Window", "set_role")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_role := gi.CString(role)
	arg_v := gi.NewPointerArgument(v.P)
	arg_role := gi.NewStringArgument(c_role)
	args := []gi.Argument{arg_v, arg_role}
	iv.Call(args, nil, nil)
	gi.Free(c_role)
}

// gdk_window_set_shadow_width
//
// [ left ] trans: nothing
//
// [ right ] trans: nothing
//
// [ top ] trans: nothing
//
// [ bottom ] trans: nothing
//
func (v Window) SetShadowWidth(left int32, right int32, top int32, bottom int32) {
	iv, err := _I.Get(436, "Window", "set_shadow_width")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_left := gi.NewInt32Argument(left)
	arg_right := gi.NewInt32Argument(right)
	arg_top := gi.NewInt32Argument(top)
	arg_bottom := gi.NewInt32Argument(bottom)
	args := []gi.Argument{arg_v, arg_left, arg_right, arg_top, arg_bottom}
	iv.Call(args, nil, nil)
}

// gdk_window_set_skip_pager_hint
//
// [ skips_pager ] trans: nothing
//
func (v Window) SetSkipPagerHint(skips_pager bool) {
	iv, err := _I.Get(437, "Window", "set_skip_pager_hint")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_skips_pager := gi.NewBoolArgument(skips_pager)
	args := []gi.Argument{arg_v, arg_skips_pager}
	iv.Call(args, nil, nil)
}

// gdk_window_set_skip_taskbar_hint
//
// [ skips_taskbar ] trans: nothing
//
func (v Window) SetSkipTaskbarHint(skips_taskbar bool) {
	iv, err := _I.Get(438, "Window", "set_skip_taskbar_hint")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_skips_taskbar := gi.NewBoolArgument(skips_taskbar)
	args := []gi.Argument{arg_v, arg_skips_taskbar}
	iv.Call(args, nil, nil)
}

// gdk_window_set_source_events
//
// [ source ] trans: nothing
//
// [ event_mask ] trans: nothing
//
func (v Window) SetSourceEvents(source InputSourceEnum, event_mask EventMaskFlags) {
	iv, err := _I.Get(439, "Window", "set_source_events")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_source := gi.NewIntArgument(int(source))
	arg_event_mask := gi.NewIntArgument(int(event_mask))
	args := []gi.Argument{arg_v, arg_source, arg_event_mask}
	iv.Call(args, nil, nil)
}

// gdk_window_set_startup_id
//
// [ startup_id ] trans: nothing
//
func (v Window) SetStartupId(startup_id string) {
	iv, err := _I.Get(440, "Window", "set_startup_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_startup_id := gi.CString(startup_id)
	arg_v := gi.NewPointerArgument(v.P)
	arg_startup_id := gi.NewStringArgument(c_startup_id)
	args := []gi.Argument{arg_v, arg_startup_id}
	iv.Call(args, nil, nil)
	gi.Free(c_startup_id)
}

// Deprecated
//
// gdk_window_set_static_gravities
//
// [ use_static ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Window) SetStaticGravities(use_static bool) (result bool) {
	iv, err := _I.Get(441, "Window", "set_static_gravities")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_use_static := gi.NewBoolArgument(use_static)
	args := []gi.Argument{arg_v, arg_use_static}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_window_set_support_multidevice
//
// [ support_multidevice ] trans: nothing
//
func (v Window) SetSupportMultidevice(support_multidevice bool) {
	iv, err := _I.Get(442, "Window", "set_support_multidevice")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_support_multidevice := gi.NewBoolArgument(support_multidevice)
	args := []gi.Argument{arg_v, arg_support_multidevice}
	iv.Call(args, nil, nil)
}

// gdk_window_set_title
//
// [ title ] trans: nothing
//
func (v Window) SetTitle(title string) {
	iv, err := _I.Get(443, "Window", "set_title")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_title := gi.CString(title)
	arg_v := gi.NewPointerArgument(v.P)
	arg_title := gi.NewStringArgument(c_title)
	args := []gi.Argument{arg_v, arg_title}
	iv.Call(args, nil, nil)
	gi.Free(c_title)
}

// gdk_window_set_transient_for
//
// [ parent ] trans: nothing
//
func (v Window) SetTransientFor(parent IWindow) {
	iv, err := _I.Get(444, "Window", "set_transient_for")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if parent != nil {
		tmp = parent.P_Window()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_parent := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_parent}
	iv.Call(args, nil, nil)
}

// gdk_window_set_type_hint
//
// [ hint ] trans: nothing
//
func (v Window) SetTypeHint(hint WindowTypeHintEnum) {
	iv, err := _I.Get(445, "Window", "set_type_hint")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_hint := gi.NewIntArgument(int(hint))
	args := []gi.Argument{arg_v, arg_hint}
	iv.Call(args, nil, nil)
}

// gdk_window_set_urgency_hint
//
// [ urgent ] trans: nothing
//
func (v Window) SetUrgencyHint(urgent bool) {
	iv, err := _I.Get(446, "Window", "set_urgency_hint")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_urgent := gi.NewBoolArgument(urgent)
	args := []gi.Argument{arg_v, arg_urgent}
	iv.Call(args, nil, nil)
}

// gdk_window_set_user_data
//
// [ user_data ] trans: nothing
//
func (v Window) SetUserData(user_data g.IObject) {
	iv, err := _I.Get(447, "Window", "set_user_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if user_data != nil {
		tmp = user_data.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_user_data := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_user_data}
	iv.Call(args, nil, nil)
}

// gdk_window_shape_combine_region
//
// [ shape_region ] trans: nothing
//
// [ offset_x ] trans: nothing
//
// [ offset_y ] trans: nothing
//
func (v Window) ShapeCombineRegion(shape_region cairo.Region, offset_x int32, offset_y int32) {
	iv, err := _I.Get(448, "Window", "shape_combine_region")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_shape_region := gi.NewPointerArgument(shape_region.P)
	arg_offset_x := gi.NewInt32Argument(offset_x)
	arg_offset_y := gi.NewInt32Argument(offset_y)
	args := []gi.Argument{arg_v, arg_shape_region, arg_offset_x, arg_offset_y}
	iv.Call(args, nil, nil)
}

// gdk_window_show
//
func (v Window) Show() {
	iv, err := _I.Get(449, "Window", "show")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_show_unraised
//
func (v Window) ShowUnraised() {
	iv, err := _I.Get(450, "Window", "show_unraised")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_show_window_menu
//
// [ event ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Window) ShowWindowMenu(event Event) (result bool) {
	iv, err := _I.Get(451, "Window", "show_window_menu")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_event := gi.NewPointerArgument(event.P)
	args := []gi.Argument{arg_v, arg_event}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_window_stick
//
func (v Window) Stick() {
	iv, err := _I.Get(452, "Window", "stick")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_window_thaw_toplevel_updates_libgtk_only
//
func (v Window) ThawToplevelUpdatesLibgtkOnly() {
	iv, err := _I.Get(453, "Window", "thaw_toplevel_updates_libgtk_only")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_thaw_updates
//
func (v Window) ThawUpdates() {
	iv, err := _I.Get(454, "Window", "thaw_updates")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_unfullscreen
//
func (v Window) Unfullscreen() {
	iv, err := _I.Get(455, "Window", "unfullscreen")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_unmaximize
//
func (v Window) Unmaximize() {
	iv, err := _I.Get(456, "Window", "unmaximize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_unstick
//
func (v Window) Unstick() {
	iv, err := _I.Get(457, "Window", "unstick")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gdk_window_withdraw
//
func (v Window) Withdraw() {
	iv, err := _I.Get(458, "Window", "withdraw")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct WindowAttr
type WindowAttr struct {
	P unsafe.Pointer
}

const SizeOfStructWindowAttr = 80

func WindowAttrGetType() gi.GType {
	ret := _I.GetGType(93, "WindowAttr")
	return ret
}

// Flags WindowAttributesType
type WindowAttributesTypeFlags int

const (
	WindowAttributesTypeTitle    WindowAttributesTypeFlags = 2
	WindowAttributesTypeX        WindowAttributesTypeFlags = 4
	WindowAttributesTypeY        WindowAttributesTypeFlags = 8
	WindowAttributesTypeCursor   WindowAttributesTypeFlags = 16
	WindowAttributesTypeVisual   WindowAttributesTypeFlags = 32
	WindowAttributesTypeWmclass  WindowAttributesTypeFlags = 64
	WindowAttributesTypeNoredir  WindowAttributesTypeFlags = 128
	WindowAttributesTypeTypeHint WindowAttributesTypeFlags = 256
)

func WindowAttributesTypeGetType() gi.GType {
	ret := _I.GetGType(94, "WindowAttributesType")
	return ret
}

type WindowChildFuncStruct struct {
	F_window Window
}

func GetPointer_myWindowChildFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGdkWindowChildFunc())
}

//export myGdkWindowChildFunc
func myGdkWindowChildFunc(window *C.GdkWindow, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &WindowChildFuncStruct{
		F_window: WrapWindow(unsafe.Pointer(window)),
	}
	fn(args)
}

// ignore GType struct WindowClass

// Enum WindowEdge
type WindowEdgeEnum int

const (
	WindowEdgeNorthWest WindowEdgeEnum = 0
	WindowEdgeNorth     WindowEdgeEnum = 1
	WindowEdgeNorthEast WindowEdgeEnum = 2
	WindowEdgeWest      WindowEdgeEnum = 3
	WindowEdgeEast      WindowEdgeEnum = 4
	WindowEdgeSouthWest WindowEdgeEnum = 5
	WindowEdgeSouth     WindowEdgeEnum = 6
	WindowEdgeSouthEast WindowEdgeEnum = 7
)

func WindowEdgeGetType() gi.GType {
	ret := _I.GetGType(95, "WindowEdge")
	return ret
}

// Flags WindowHints
type WindowHintsFlags int

const (
	WindowHintsPos        WindowHintsFlags = 1
	WindowHintsMinSize    WindowHintsFlags = 2
	WindowHintsMaxSize    WindowHintsFlags = 4
	WindowHintsBaseSize   WindowHintsFlags = 8
	WindowHintsAspect     WindowHintsFlags = 16
	WindowHintsResizeInc  WindowHintsFlags = 32
	WindowHintsWinGravity WindowHintsFlags = 64
	WindowHintsUserPos    WindowHintsFlags = 128
	WindowHintsUserSize   WindowHintsFlags = 256
)

func WindowHintsGetType() gi.GType {
	ret := _I.GetGType(96, "WindowHints")
	return ret
}

type WindowInvalidateHandlerFuncStruct struct {
	F_window Window
	F_region cairo.Region
}

func GetPointer_myWindowInvalidateHandlerFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGdkWindowInvalidateHandlerFunc())
}

//export myGdkWindowInvalidateHandlerFunc
func myGdkWindowInvalidateHandlerFunc(window *C.GdkWindow, region *C.cairo_region_t) {
	// TODO: not found user_data
}

// Struct WindowRedirect
type WindowRedirect struct {
	P unsafe.Pointer
}

func WindowRedirectGetType() gi.GType {
	ret := _I.GetGType(97, "WindowRedirect")
	return ret
}

// Flags WindowState
type WindowStateFlags int

const (
	WindowStateWithdrawn       WindowStateFlags = 1
	WindowStateIconified       WindowStateFlags = 2
	WindowStateMaximized       WindowStateFlags = 4
	WindowStateSticky          WindowStateFlags = 8
	WindowStateFullscreen      WindowStateFlags = 16
	WindowStateAbove           WindowStateFlags = 32
	WindowStateBelow           WindowStateFlags = 64
	WindowStateFocused         WindowStateFlags = 128
	WindowStateTiled           WindowStateFlags = 256
	WindowStateTopTiled        WindowStateFlags = 512
	WindowStateTopResizable    WindowStateFlags = 1024
	WindowStateRightTiled      WindowStateFlags = 2048
	WindowStateRightResizable  WindowStateFlags = 4096
	WindowStateBottomTiled     WindowStateFlags = 8192
	WindowStateBottomResizable WindowStateFlags = 16384
	WindowStateLeftTiled       WindowStateFlags = 32768
	WindowStateLeftResizable   WindowStateFlags = 65536
)

func WindowStateGetType() gi.GType {
	ret := _I.GetGType(98, "WindowState")
	return ret
}

// Enum WindowType
type WindowTypeEnum int

const (
	WindowTypeRoot       WindowTypeEnum = 0
	WindowTypeToplevel   WindowTypeEnum = 1
	WindowTypeChild      WindowTypeEnum = 2
	WindowTypeTemp       WindowTypeEnum = 3
	WindowTypeForeign    WindowTypeEnum = 4
	WindowTypeOffscreen  WindowTypeEnum = 5
	WindowTypeSubsurface WindowTypeEnum = 6
)

func WindowTypeGetType() gi.GType {
	ret := _I.GetGType(99, "WindowType")
	return ret
}

// Enum WindowTypeHint
type WindowTypeHintEnum int

const (
	WindowTypeHintNormal       WindowTypeHintEnum = 0
	WindowTypeHintDialog       WindowTypeHintEnum = 1
	WindowTypeHintMenu         WindowTypeHintEnum = 2
	WindowTypeHintToolbar      WindowTypeHintEnum = 3
	WindowTypeHintSplashscreen WindowTypeHintEnum = 4
	WindowTypeHintUtility      WindowTypeHintEnum = 5
	WindowTypeHintDock         WindowTypeHintEnum = 6
	WindowTypeHintDesktop      WindowTypeHintEnum = 7
	WindowTypeHintDropdownMenu WindowTypeHintEnum = 8
	WindowTypeHintPopupMenu    WindowTypeHintEnum = 9
	WindowTypeHintTooltip      WindowTypeHintEnum = 10
	WindowTypeHintNotification WindowTypeHintEnum = 11
	WindowTypeHintCombo        WindowTypeHintEnum = 12
	WindowTypeHintDnd          WindowTypeHintEnum = 13
)

func WindowTypeHintGetType() gi.GType {
	ret := _I.GetGType(100, "WindowTypeHint")
	return ret
}

// Enum WindowWindowClass
type WindowWindowClassEnum int

const (
	WindowWindowClassInputOutput WindowWindowClassEnum = 0
	WindowWindowClassInputOnly   WindowWindowClassEnum = 1
)

func WindowWindowClassGetType() gi.GType {
	ret := _I.GetGType(101, "WindowWindowClass")
	return ret
}

// Deprecated
//
// gdk_add_option_entries_libgtk_only
//
// [ group ] trans: nothing
//
func AddOptionEntriesLibgtkOnly(group g.OptionGroup) {
	iv, err := _I.Get(459, "add_option_entries_libgtk_only", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_group := gi.NewPointerArgument(group.P)
	args := []gi.Argument{arg_group}
	iv.Call(args, nil, nil)
}

// gdk_atom_intern
//
// [ atom_name ] trans: nothing
//
// [ only_if_exists ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomIntern(atom_name string, only_if_exists bool) (result Atom) {
	iv, err := _I.Get(460, "atom_intern", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_atom_name := gi.CString(atom_name)
	arg_atom_name := gi.NewStringArgument(c_atom_name)
	arg_only_if_exists := gi.NewBoolArgument(only_if_exists)
	args := []gi.Argument{arg_atom_name, arg_only_if_exists}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_atom_name)
	result.P = ret.Pointer()
	return
}

// gdk_atom_intern_static_string
//
// [ atom_name ] trans: nothing
//
// [ result ] trans: nothing
//
func AtomInternStaticString(atom_name string) (result Atom) {
	iv, err := _I.Get(461, "atom_intern_static_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_atom_name := gi.CString(atom_name)
	arg_atom_name := gi.NewStringArgument(c_atom_name)
	args := []gi.Argument{arg_atom_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_atom_name)
	result.P = ret.Pointer()
	return
}

// gdk_beep
//
func Beep() {
	iv, err := _I.Get(462, "beep", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// Deprecated
//
// gdk_cairo_create
//
// [ window ] trans: nothing
//
// [ result ] trans: everything
//
func CairoCreate(window IWindow) (result cairo.Context) {
	iv, err := _I.Get(463, "cairo_create", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_window := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_window}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_cairo_draw_from_gl
//
// [ cr ] trans: nothing
//
// [ window ] trans: nothing
//
// [ source ] trans: nothing
//
// [ source_type ] trans: nothing
//
// [ buffer_scale ] trans: nothing
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
func CairoDrawFromGl(cr cairo.Context, window IWindow, source int32, source_type int32, buffer_scale int32, x int32, y int32, width int32, height int32) {
	iv, err := _I.Get(464, "cairo_draw_from_gl", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_window := gi.NewPointerArgument(tmp)
	arg_source := gi.NewInt32Argument(source)
	arg_source_type := gi.NewInt32Argument(source_type)
	arg_buffer_scale := gi.NewInt32Argument(buffer_scale)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	args := []gi.Argument{arg_cr, arg_window, arg_source, arg_source_type, arg_buffer_scale, arg_x, arg_y, arg_width, arg_height}
	iv.Call(args, nil, nil)
}

// gdk_cairo_get_clip_rectangle
//
// [ cr ] trans: nothing
//
// [ rect ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func CairoGetClipRectangle(cr cairo.Context, rect Rectangle) (result bool) {
	iv, err := _I.Get(465, "cairo_get_clip_rectangle", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_rect := gi.NewPointerArgument(rect.P)
	args := []gi.Argument{arg_cr, arg_rect}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_cairo_get_drawing_context
//
// [ cr ] trans: nothing
//
// [ result ] trans: nothing
//
func CairoGetDrawingContext(cr cairo.Context) (result DrawingContext) {
	iv, err := _I.Get(466, "cairo_get_drawing_context", "")
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

// gdk_cairo_rectangle
//
// [ cr ] trans: nothing
//
// [ rectangle ] trans: nothing
//
func CairoRectangle(cr cairo.Context, rectangle Rectangle) {
	iv, err := _I.Get(467, "cairo_rectangle", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_rectangle := gi.NewPointerArgument(rectangle.P)
	args := []gi.Argument{arg_cr, arg_rectangle}
	iv.Call(args, nil, nil)
}

// gdk_cairo_region
//
// [ cr ] trans: nothing
//
// [ region ] trans: nothing
//
func CairoRegion(cr cairo.Context, region cairo.Region) {
	iv, err := _I.Get(468, "cairo_region", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_region := gi.NewPointerArgument(region.P)
	args := []gi.Argument{arg_cr, arg_region}
	iv.Call(args, nil, nil)
}

// gdk_cairo_region_create_from_surface
//
// [ surface ] trans: nothing
//
// [ result ] trans: everything
//
func CairoRegionCreateFromSurface(surface cairo.Surface) (result cairo.Region) {
	iv, err := _I.Get(469, "cairo_region_create_from_surface", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_surface := gi.NewPointerArgument(surface.P)
	args := []gi.Argument{arg_surface}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// gdk_cairo_set_source_color
//
// [ cr ] trans: nothing
//
// [ color ] trans: nothing
//
func CairoSetSourceColor(cr cairo.Context, color Color) {
	iv, err := _I.Get(470, "cairo_set_source_color", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_color := gi.NewPointerArgument(color.P)
	args := []gi.Argument{arg_cr, arg_color}
	iv.Call(args, nil, nil)
}

// gdk_cairo_set_source_pixbuf
//
// [ cr ] trans: nothing
//
// [ pixbuf ] trans: nothing
//
// [ pixbuf_x ] trans: nothing
//
// [ pixbuf_y ] trans: nothing
//
func CairoSetSourcePixbuf(cr cairo.Context, pixbuf gdkpixbuf.IPixbuf, pixbuf_x float64, pixbuf_y float64) {
	iv, err := _I.Get(471, "cairo_set_source_pixbuf", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pixbuf != nil {
		tmp = pixbuf.P_Pixbuf()
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_pixbuf := gi.NewPointerArgument(tmp)
	arg_pixbuf_x := gi.NewDoubleArgument(pixbuf_x)
	arg_pixbuf_y := gi.NewDoubleArgument(pixbuf_y)
	args := []gi.Argument{arg_cr, arg_pixbuf, arg_pixbuf_x, arg_pixbuf_y}
	iv.Call(args, nil, nil)
}

// gdk_cairo_set_source_rgba
//
// [ cr ] trans: nothing
//
// [ rgba ] trans: nothing
//
func CairoSetSourceRgba(cr cairo.Context, rgba RGBA) {
	iv, err := _I.Get(472, "cairo_set_source_rgba", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_rgba := gi.NewPointerArgument(rgba.P)
	args := []gi.Argument{arg_cr, arg_rgba}
	iv.Call(args, nil, nil)
}

// gdk_cairo_set_source_window
//
// [ cr ] trans: nothing
//
// [ window ] trans: nothing
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
func CairoSetSourceWindow(cr cairo.Context, window IWindow, x float64, y float64) {
	iv, err := _I.Get(473, "cairo_set_source_window", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_window := gi.NewPointerArgument(tmp)
	arg_x := gi.NewDoubleArgument(x)
	arg_y := gi.NewDoubleArgument(y)
	args := []gi.Argument{arg_cr, arg_window, arg_x, arg_y}
	iv.Call(args, nil, nil)
}

// gdk_cairo_surface_create_from_pixbuf
//
// [ pixbuf ] trans: nothing
//
// [ scale ] trans: nothing
//
// [ for_window ] trans: nothing
//
// [ result ] trans: everything
//
func CairoSurfaceCreateFromPixbuf(pixbuf gdkpixbuf.IPixbuf, scale int32, for_window IWindow) (result cairo.Surface) {
	iv, err := _I.Get(474, "cairo_surface_create_from_pixbuf", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pixbuf != nil {
		tmp = pixbuf.P_Pixbuf()
	}
	var tmp1 unsafe.Pointer
	if for_window != nil {
		tmp1 = for_window.P_Window()
	}
	arg_pixbuf := gi.NewPointerArgument(tmp)
	arg_scale := gi.NewInt32Argument(scale)
	arg_for_window := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_pixbuf, arg_scale, arg_for_window}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// gdk_color_parse
//
// [ spec ] trans: nothing
//
// [ color ] trans: nothing, dir: out
//
// [ result ] trans: nothing
//
func ColorParse(spec string, color Color) (result bool) {
	iv, err := _I.Get(475, "color_parse", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_spec := gi.CString(spec)
	arg_spec := gi.NewStringArgument(c_spec)
	arg_color := gi.NewPointerArgument(color.P)
	args := []gi.Argument{arg_spec, arg_color}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_spec)
	result = ret.Bool()
	return
}

// gdk_disable_multidevice
//
func DisableMultidevice() {
	iv, err := _I.Get(476, "disable_multidevice", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// gdk_drag_abort
//
// [ context ] trans: nothing
//
// [ time_ ] trans: nothing
//
func DragAbort(context IDragContext, time_ uint32) {
	iv, err := _I.Get(477, "drag_abort", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_DragContext()
	}
	arg_context := gi.NewPointerArgument(tmp)
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_context, arg_time_}
	iv.Call(args, nil, nil)
}

// gdk_drag_begin
//
// [ window ] trans: nothing
//
// [ targets ] trans: nothing
//
// [ result ] trans: everything
//
func DragBegin(window IWindow, targets g.List) (result DragContext) {
	iv, err := _I.Get(478, "drag_begin", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_window := gi.NewPointerArgument(tmp)
	arg_targets := gi.NewPointerArgument(targets.P)
	args := []gi.Argument{arg_window, arg_targets}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_drag_begin_for_device
//
// [ window ] trans: nothing
//
// [ device ] trans: nothing
//
// [ targets ] trans: nothing
//
// [ result ] trans: everything
//
func DragBeginForDevice(window IWindow, device IDevice, targets g.List) (result DragContext) {
	iv, err := _I.Get(479, "drag_begin_for_device", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	var tmp1 unsafe.Pointer
	if device != nil {
		tmp1 = device.P_Device()
	}
	arg_window := gi.NewPointerArgument(tmp)
	arg_device := gi.NewPointerArgument(tmp1)
	arg_targets := gi.NewPointerArgument(targets.P)
	args := []gi.Argument{arg_window, arg_device, arg_targets}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_drag_begin_from_point
//
// [ window ] trans: nothing
//
// [ device ] trans: nothing
//
// [ targets ] trans: nothing
//
// [ x_root ] trans: nothing
//
// [ y_root ] trans: nothing
//
// [ result ] trans: everything
//
func DragBeginFromPoint(window IWindow, device IDevice, targets g.List, x_root int32, y_root int32) (result DragContext) {
	iv, err := _I.Get(480, "drag_begin_from_point", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	var tmp1 unsafe.Pointer
	if device != nil {
		tmp1 = device.P_Device()
	}
	arg_window := gi.NewPointerArgument(tmp)
	arg_device := gi.NewPointerArgument(tmp1)
	arg_targets := gi.NewPointerArgument(targets.P)
	arg_x_root := gi.NewInt32Argument(x_root)
	arg_y_root := gi.NewInt32Argument(y_root)
	args := []gi.Argument{arg_window, arg_device, arg_targets, arg_x_root, arg_y_root}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_drag_drop
//
// [ context ] trans: nothing
//
// [ time_ ] trans: nothing
//
func DragDrop(context IDragContext, time_ uint32) {
	iv, err := _I.Get(481, "drag_drop", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_DragContext()
	}
	arg_context := gi.NewPointerArgument(tmp)
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_context, arg_time_}
	iv.Call(args, nil, nil)
}

// gdk_drag_drop_done
//
// [ context ] trans: nothing
//
// [ success ] trans: nothing
//
func DragDropDone(context IDragContext, success bool) {
	iv, err := _I.Get(482, "drag_drop_done", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_DragContext()
	}
	arg_context := gi.NewPointerArgument(tmp)
	arg_success := gi.NewBoolArgument(success)
	args := []gi.Argument{arg_context, arg_success}
	iv.Call(args, nil, nil)
}

// gdk_drag_drop_succeeded
//
// [ context ] trans: nothing
//
// [ result ] trans: nothing
//
func DragDropSucceeded(context IDragContext) (result bool) {
	iv, err := _I.Get(483, "drag_drop_succeeded", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_DragContext()
	}
	arg_context := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_context}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_drag_find_window_for_screen
//
// [ context ] trans: nothing
//
// [ drag_window ] trans: nothing
//
// [ screen ] trans: nothing
//
// [ x_root ] trans: nothing
//
// [ y_root ] trans: nothing
//
// [ dest_window ] trans: everything, dir: out
//
// [ protocol ] trans: everything, dir: out
//
func DragFindWindowForScreen(context IDragContext, drag_window IWindow, screen IScreen, x_root int32, y_root int32) (dest_window Window, protocol DragProtocolEnum) {
	iv, err := _I.Get(484, "drag_find_window_for_screen", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_DragContext()
	}
	var tmp1 unsafe.Pointer
	if drag_window != nil {
		tmp1 = drag_window.P_Window()
	}
	var tmp2 unsafe.Pointer
	if screen != nil {
		tmp2 = screen.P_Screen()
	}
	arg_context := gi.NewPointerArgument(tmp)
	arg_drag_window := gi.NewPointerArgument(tmp1)
	arg_screen := gi.NewPointerArgument(tmp2)
	arg_x_root := gi.NewInt32Argument(x_root)
	arg_y_root := gi.NewInt32Argument(y_root)
	arg_dest_window := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_protocol := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_context, arg_drag_window, arg_screen, arg_x_root, arg_y_root, arg_dest_window, arg_protocol}
	iv.Call(args, nil, &outArgs[0])
	dest_window.P = outArgs[0].Pointer()
	protocol = DragProtocolEnum(outArgs[1].Int())
	return
}

// gdk_drag_get_selection
//
// [ context ] trans: nothing
//
// [ result ] trans: nothing
//
func DragGetSelection(context IDragContext) (result Atom) {
	iv, err := _I.Get(485, "drag_get_selection", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_DragContext()
	}
	arg_context := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_context}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_drag_motion
//
// [ context ] trans: nothing
//
// [ dest_window ] trans: nothing
//
// [ protocol ] trans: nothing
//
// [ x_root ] trans: nothing
//
// [ y_root ] trans: nothing
//
// [ suggested_action ] trans: nothing
//
// [ possible_actions ] trans: nothing
//
// [ time_ ] trans: nothing
//
// [ result ] trans: nothing
//
func DragMotion(context IDragContext, dest_window IWindow, protocol DragProtocolEnum, x_root int32, y_root int32, suggested_action DragActionFlags, possible_actions DragActionFlags, time_ uint32) (result bool) {
	iv, err := _I.Get(486, "drag_motion", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_DragContext()
	}
	var tmp1 unsafe.Pointer
	if dest_window != nil {
		tmp1 = dest_window.P_Window()
	}
	arg_context := gi.NewPointerArgument(tmp)
	arg_dest_window := gi.NewPointerArgument(tmp1)
	arg_protocol := gi.NewIntArgument(int(protocol))
	arg_x_root := gi.NewInt32Argument(x_root)
	arg_y_root := gi.NewInt32Argument(y_root)
	arg_suggested_action := gi.NewIntArgument(int(suggested_action))
	arg_possible_actions := gi.NewIntArgument(int(possible_actions))
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_context, arg_dest_window, arg_protocol, arg_x_root, arg_y_root, arg_suggested_action, arg_possible_actions, arg_time_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_drag_status
//
// [ context ] trans: nothing
//
// [ action ] trans: nothing
//
// [ time_ ] trans: nothing
//
func DragStatus(context IDragContext, action DragActionFlags, time_ uint32) {
	iv, err := _I.Get(487, "drag_status", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_DragContext()
	}
	arg_context := gi.NewPointerArgument(tmp)
	arg_action := gi.NewIntArgument(int(action))
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_context, arg_action, arg_time_}
	iv.Call(args, nil, nil)
}

// gdk_drop_finish
//
// [ context ] trans: nothing
//
// [ success ] trans: nothing
//
// [ time_ ] trans: nothing
//
func DropFinish(context IDragContext, success bool, time_ uint32) {
	iv, err := _I.Get(488, "drop_finish", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_DragContext()
	}
	arg_context := gi.NewPointerArgument(tmp)
	arg_success := gi.NewBoolArgument(success)
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_context, arg_success, arg_time_}
	iv.Call(args, nil, nil)
}

// gdk_drop_reply
//
// [ context ] trans: nothing
//
// [ accepted ] trans: nothing
//
// [ time_ ] trans: nothing
//
func DropReply(context IDragContext, accepted bool, time_ uint32) {
	iv, err := _I.Get(489, "drop_reply", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_DragContext()
	}
	arg_context := gi.NewPointerArgument(tmp)
	arg_accepted := gi.NewBoolArgument(accepted)
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_context, arg_accepted, arg_time_}
	iv.Call(args, nil, nil)
}

// gdk_error_trap_pop
//
// [ result ] trans: nothing
//
func ErrorTrapPop() (result int32) {
	iv, err := _I.Get(490, "error_trap_pop", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int32()
	return
}

// gdk_error_trap_pop_ignored
//
func ErrorTrapPopIgnored() {
	iv, err := _I.Get(491, "error_trap_pop_ignored", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// gdk_error_trap_push
//
func ErrorTrapPush() {
	iv, err := _I.Get(492, "error_trap_push", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// gdk_event_get
//
// [ result ] trans: everything
//
func EventGet() (result Event) {
	iv, err := _I.Get(493, "event_get", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_event_handler_set
//
// [ func1 ] trans: nothing
//
// [ data ] trans: nothing
//
// [ notify ] trans: nothing
//
func EventHandlerSet(func1 int /*TODO_TYPE CALLBACK*/, data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) {
	iv, err := _I.Get(494, "event_handler_set", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_func1 := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myEventFunc()))
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_func1, arg_data, arg_notify}
	iv.Call(args, nil, nil)
}

// gdk_event_peek
//
// [ result ] trans: everything
//
func EventPeek() (result Event) {
	iv, err := _I.Get(495, "event_peek", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_event_request_motions
//
// [ event ] trans: nothing
//
func EventRequestMotions(event EventMotion) {
	iv, err := _I.Get(496, "event_request_motions", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_event := gi.NewPointerArgument(event.P)
	args := []gi.Argument{arg_event}
	iv.Call(args, nil, nil)
}

// gdk_events_get_angle
//
// [ event1 ] trans: nothing
//
// [ event2 ] trans: nothing
//
// [ angle ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func EventsGetAngle(event1 Event, event2 Event) (result bool, angle float64) {
	iv, err := _I.Get(497, "events_get_angle", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_event1 := gi.NewPointerArgument(event1.P)
	arg_event2 := gi.NewPointerArgument(event2.P)
	arg_angle := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_event1, arg_event2, arg_angle}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	angle = outArgs[0].Double()
	result = ret.Bool()
	return
}

// gdk_events_get_center
//
// [ event1 ] trans: nothing
//
// [ event2 ] trans: nothing
//
// [ x ] trans: everything, dir: out
//
// [ y ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func EventsGetCenter(event1 Event, event2 Event) (result bool, x float64, y float64) {
	iv, err := _I.Get(498, "events_get_center", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_event1 := gi.NewPointerArgument(event1.P)
	arg_event2 := gi.NewPointerArgument(event2.P)
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_event1, arg_event2, arg_x, arg_y}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	x = outArgs[0].Double()
	y = outArgs[1].Double()
	result = ret.Bool()
	return
}

// gdk_events_get_distance
//
// [ event1 ] trans: nothing
//
// [ event2 ] trans: nothing
//
// [ distance ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func EventsGetDistance(event1 Event, event2 Event) (result bool, distance float64) {
	iv, err := _I.Get(499, "events_get_distance", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_event1 := gi.NewPointerArgument(event1.P)
	arg_event2 := gi.NewPointerArgument(event2.P)
	arg_distance := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_event1, arg_event2, arg_distance}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	distance = outArgs[0].Double()
	result = ret.Bool()
	return
}

// gdk_events_pending
//
// [ result ] trans: nothing
//
func EventsPending() (result bool) {
	iv, err := _I.Get(500, "events_pending", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_flush
//
func Flush() {
	iv, err := _I.Get(501, "flush", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// gdk_get_default_root_window
//
// [ result ] trans: nothing
//
func GetDefaultRootWindow() (result Window) {
	iv, err := _I.Get(502, "get_default_root_window", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// gdk_get_display
//
// [ result ] trans: everything
//
func GetDisplay() (result string) {
	iv, err := _I.Get(503, "get_display", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// gdk_get_display_arg_name
//
// [ result ] trans: nothing
//
func GetDisplayArgName() (result string) {
	iv, err := _I.Get(504, "get_display_arg_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// gdk_get_program_class
//
// [ result ] trans: nothing
//
func GetProgramClass() (result string) {
	iv, err := _I.Get(505, "get_program_class", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Copy()
	return
}

// gdk_get_show_events
//
// [ result ] trans: nothing
//
func GetShowEvents() (result bool) {
	iv, err := _I.Get(506, "get_show_events", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_gl_error_quark
//
// [ result ] trans: nothing
//
func GlErrorQuark() (result uint32) {
	iv, err := _I.Get(507, "gl_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// gdk_init
//
// [ argc ] trans: everything, dir: inout
//
// [ argv ] trans: everything, dir: inout
//
func Init(argc int /*TODO:TYPE*/, argv int /*TODO:TYPE*/) {
	iv, err := _I.Get(508, "init", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	iv.Call(nil, nil, &outArgs[0])
}

// gdk_init_check
//
// [ argc ] trans: everything, dir: inout
//
// [ argv ] trans: everything, dir: inout
//
// [ result ] trans: nothing
//
func InitCheck(argc int /*TODO:TYPE*/, argv int /*TODO:TYPE*/) (result bool) {
	iv, err := _I.Get(509, "init_check", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	var ret gi.Argument
	iv.Call(nil, &ret, &outArgs[0])
	result = ret.Bool()
	return
}

// Deprecated
//
// gdk_keyboard_grab
//
// [ window ] trans: nothing
//
// [ owner_events ] trans: nothing
//
// [ time_ ] trans: nothing
//
// [ result ] trans: nothing
//
func KeyboardGrab(window IWindow, owner_events bool, time_ uint32) (result GrabStatusEnum) {
	iv, err := _I.Get(510, "keyboard_grab", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_window := gi.NewPointerArgument(tmp)
	arg_owner_events := gi.NewBoolArgument(owner_events)
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_window, arg_owner_events, arg_time_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = GrabStatusEnum(ret.Int())
	return
}

// Deprecated
//
// gdk_keyboard_ungrab
//
// [ time_ ] trans: nothing
//
func KeyboardUngrab(time_ uint32) {
	iv, err := _I.Get(511, "keyboard_ungrab", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_time_}
	iv.Call(args, nil, nil)
}

// gdk_keyval_convert_case
//
// [ symbol ] trans: nothing
//
// [ lower ] trans: everything, dir: out
//
// [ upper ] trans: everything, dir: out
//
func KeyvalConvertCase(symbol uint32) (lower uint32, upper uint32) {
	iv, err := _I.Get(512, "keyval_convert_case", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_symbol := gi.NewUint32Argument(symbol)
	arg_lower := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_upper := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_symbol, arg_lower, arg_upper}
	iv.Call(args, nil, &outArgs[0])
	lower = outArgs[0].Uint32()
	upper = outArgs[1].Uint32()
	return
}

// gdk_keyval_from_name
//
// [ keyval_name ] trans: nothing
//
// [ result ] trans: nothing
//
func KeyvalFromName(keyval_name string) (result uint32) {
	iv, err := _I.Get(513, "keyval_from_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_keyval_name := gi.CString(keyval_name)
	arg_keyval_name := gi.NewStringArgument(c_keyval_name)
	args := []gi.Argument{arg_keyval_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_keyval_name)
	result = ret.Uint32()
	return
}

// gdk_keyval_is_lower
//
// [ keyval ] trans: nothing
//
// [ result ] trans: nothing
//
func KeyvalIsLower(keyval uint32) (result bool) {
	iv, err := _I.Get(514, "keyval_is_lower", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_keyval := gi.NewUint32Argument(keyval)
	args := []gi.Argument{arg_keyval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_keyval_is_upper
//
// [ keyval ] trans: nothing
//
// [ result ] trans: nothing
//
func KeyvalIsUpper(keyval uint32) (result bool) {
	iv, err := _I.Get(515, "keyval_is_upper", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_keyval := gi.NewUint32Argument(keyval)
	args := []gi.Argument{arg_keyval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_keyval_name
//
// [ keyval ] trans: nothing
//
// [ result ] trans: nothing
//
func KeyvalName(keyval uint32) (result string) {
	iv, err := _I.Get(516, "keyval_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_keyval := gi.NewUint32Argument(keyval)
	args := []gi.Argument{arg_keyval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// gdk_keyval_to_lower
//
// [ keyval ] trans: nothing
//
// [ result ] trans: nothing
//
func KeyvalToLower(keyval uint32) (result uint32) {
	iv, err := _I.Get(517, "keyval_to_lower", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_keyval := gi.NewUint32Argument(keyval)
	args := []gi.Argument{arg_keyval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gdk_keyval_to_unicode
//
// [ keyval ] trans: nothing
//
// [ result ] trans: nothing
//
func KeyvalToUnicode(keyval uint32) (result uint32) {
	iv, err := _I.Get(518, "keyval_to_unicode", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_keyval := gi.NewUint32Argument(keyval)
	args := []gi.Argument{arg_keyval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gdk_keyval_to_upper
//
// [ keyval ] trans: nothing
//
// [ result ] trans: nothing
//
func KeyvalToUpper(keyval uint32) (result uint32) {
	iv, err := _I.Get(519, "keyval_to_upper", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_keyval := gi.NewUint32Argument(keyval)
	args := []gi.Argument{arg_keyval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// Deprecated
//
// gdk_list_visuals
//
// [ result ] trans: container
//
func ListVisuals() (result g.List) {
	iv, err := _I.Get(520, "list_visuals", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_notify_startup_complete
//
func NotifyStartupComplete() {
	iv, err := _I.Get(521, "notify_startup_complete", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// gdk_notify_startup_complete_with_id
//
// [ startup_id ] trans: nothing
//
func NotifyStartupCompleteWithId(startup_id string) {
	iv, err := _I.Get(522, "notify_startup_complete_with_id", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_startup_id := gi.CString(startup_id)
	arg_startup_id := gi.NewStringArgument(c_startup_id)
	args := []gi.Argument{arg_startup_id}
	iv.Call(args, nil, nil)
	gi.Free(c_startup_id)
}

// gdk_offscreen_window_get_embedder
//
// [ window ] trans: nothing
//
// [ result ] trans: nothing
//
func OffscreenWindowGetEmbedder(window IWindow) (result Window) {
	iv, err := _I.Get(523, "offscreen_window_get_embedder", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_window := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_window}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_offscreen_window_get_surface
//
// [ window ] trans: nothing
//
// [ result ] trans: nothing
//
func OffscreenWindowGetSurface(window IWindow) (result cairo.Surface) {
	iv, err := _I.Get(524, "offscreen_window_get_surface", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_window := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_window}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_offscreen_window_set_embedder
//
// [ window ] trans: nothing
//
// [ embedder ] trans: nothing
//
func OffscreenWindowSetEmbedder(window IWindow, embedder IWindow) {
	iv, err := _I.Get(525, "offscreen_window_set_embedder", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	var tmp1 unsafe.Pointer
	if embedder != nil {
		tmp1 = embedder.P_Window()
	}
	arg_window := gi.NewPointerArgument(tmp)
	arg_embedder := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_window, arg_embedder}
	iv.Call(args, nil, nil)
}

// gdk_pango_context_get
//
// [ result ] trans: everything
//
func PangoContextGet() (result pango.Context) {
	iv, err := _I.Get(526, "pango_context_get", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pango_context_get_for_display
//
// [ display ] trans: nothing
//
// [ result ] trans: everything
//
func PangoContextGetForDisplay(display IDisplay) (result pango.Context) {
	iv, err := _I.Get(527, "pango_context_get_for_display", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if display != nil {
		tmp = display.P_Display()
	}
	arg_display := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_display}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pango_context_get_for_screen
//
// [ screen ] trans: nothing
//
// [ result ] trans: everything
//
func PangoContextGetForScreen(screen IScreen) (result pango.Context) {
	iv, err := _I.Get(528, "pango_context_get_for_screen", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if screen != nil {
		tmp = screen.P_Screen()
	}
	arg_screen := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_screen}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_parse_args
//
// [ argc ] trans: nothing, dir: inout
//
// [ argv ] trans: everything, dir: inout
//
func ParseArgs(argc int /*TODO:TYPE*/, argv int /*TODO:TYPE*/) {
	iv, err := _I.Get(529, "parse_args", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	iv.Call(nil, nil, &outArgs[0])
}

// gdk_pixbuf_get_from_surface
//
// [ surface ] trans: nothing
//
// [ src_x ] trans: nothing
//
// [ src_y ] trans: nothing
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
// [ result ] trans: everything
//
func PixbufGetFromSurface(surface cairo.Surface, src_x int32, src_y int32, width int32, height int32) (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(530, "pixbuf_get_from_surface", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_surface := gi.NewPointerArgument(surface.P)
	arg_src_x := gi.NewInt32Argument(src_x)
	arg_src_y := gi.NewInt32Argument(src_y)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	args := []gi.Argument{arg_surface, arg_src_x, arg_src_y, arg_width, arg_height}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_pixbuf_get_from_window
//
// [ window ] trans: nothing
//
// [ src_x ] trans: nothing
//
// [ src_y ] trans: nothing
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
// [ result ] trans: everything
//
func PixbufGetFromWindow(window IWindow, src_x int32, src_y int32, width int32, height int32) (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(531, "pixbuf_get_from_window", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_window := gi.NewPointerArgument(tmp)
	arg_src_x := gi.NewInt32Argument(src_x)
	arg_src_y := gi.NewInt32Argument(src_y)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	args := []gi.Argument{arg_window, arg_src_x, arg_src_y, arg_width, arg_height}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// gdk_pointer_grab
//
// [ window ] trans: nothing
//
// [ owner_events ] trans: nothing
//
// [ event_mask ] trans: nothing
//
// [ confine_to ] trans: nothing
//
// [ cursor ] trans: nothing
//
// [ time_ ] trans: nothing
//
// [ result ] trans: nothing
//
func PointerGrab(window IWindow, owner_events bool, event_mask EventMaskFlags, confine_to IWindow, cursor ICursor, time_ uint32) (result GrabStatusEnum) {
	iv, err := _I.Get(532, "pointer_grab", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	var tmp1 unsafe.Pointer
	if confine_to != nil {
		tmp1 = confine_to.P_Window()
	}
	var tmp2 unsafe.Pointer
	if cursor != nil {
		tmp2 = cursor.P_Cursor()
	}
	arg_window := gi.NewPointerArgument(tmp)
	arg_owner_events := gi.NewBoolArgument(owner_events)
	arg_event_mask := gi.NewIntArgument(int(event_mask))
	arg_confine_to := gi.NewPointerArgument(tmp1)
	arg_cursor := gi.NewPointerArgument(tmp2)
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_window, arg_owner_events, arg_event_mask, arg_confine_to, arg_cursor, arg_time_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = GrabStatusEnum(ret.Int())
	return
}

// Deprecated
//
// gdk_pointer_is_grabbed
//
// [ result ] trans: nothing
//
func PointerIsGrabbed() (result bool) {
	iv, err := _I.Get(533, "pointer_is_grabbed", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// Deprecated
//
// gdk_pointer_ungrab
//
// [ time_ ] trans: nothing
//
func PointerUngrab(time_ uint32) {
	iv, err := _I.Get(534, "pointer_ungrab", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_time_}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// gdk_pre_parse_libgtk_only
//
func PreParseLibgtkOnly() {
	iv, err := _I.Get(535, "pre_parse_libgtk_only", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// gdk_property_delete
//
// [ window ] trans: nothing
//
// [ property ] trans: nothing
//
func PropertyDelete(window IWindow, property Atom) {
	iv, err := _I.Get(536, "property_delete", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_window := gi.NewPointerArgument(tmp)
	arg_property := gi.NewPointerArgument(property.P)
	args := []gi.Argument{arg_window, arg_property}
	iv.Call(args, nil, nil)
}

// gdk_property_get
//
// [ window ] trans: nothing
//
// [ property ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ offset ] trans: nothing
//
// [ length ] trans: nothing
//
// [ pdelete ] trans: nothing
//
// [ actual_property_type ] trans: nothing, dir: out
//
// [ actual_format ] trans: everything, dir: out
//
// [ actual_length ] trans: everything, dir: out
//
// [ data ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func PropertyGet(window IWindow, property Atom, type1 Atom, offset uint64, length uint64, pdelete int32, actual_property_type Atom) (result bool, actual_format int32, data gi.Uint8Array) {
	iv, err := _I.Get(537, "property_get", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_window := gi.NewPointerArgument(tmp)
	arg_property := gi.NewPointerArgument(property.P)
	arg_type1 := gi.NewPointerArgument(type1.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_length := gi.NewUint64Argument(length)
	arg_pdelete := gi.NewInt32Argument(pdelete)
	arg_actual_property_type := gi.NewPointerArgument(actual_property_type.P)
	arg_actual_format := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_actual_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_data := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_window, arg_property, arg_type1, arg_offset, arg_length, arg_pdelete, arg_actual_property_type, arg_actual_format, arg_actual_length, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var actual_length int32
	_ = actual_length
	actual_format = outArgs[0].Int32()
	actual_length = outArgs[1].Int32()
	data.P = outArgs[2].Pointer()
	result = ret.Bool()
	data.Len = int(actual_length)
	return
}

// Deprecated
//
// gdk_query_depths
//
// [ depths ] trans: nothing, dir: out
//
// [ count ] trans: everything, dir: out
//
func QueryDepths() (depths gi.Int32Array) {
	iv, err := _I.Get(538, "query_depths", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_depths := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_count := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_depths, arg_count}
	iv.Call(args, nil, &outArgs[0])
	var count int32
	_ = count
	depths.P = outArgs[0].Pointer()
	count = outArgs[1].Int32()
	depths.Len = int(count)
	return
}

// Deprecated
//
// gdk_query_visual_types
//
// [ visual_types ] trans: nothing, dir: out
//
// [ count ] trans: everything, dir: out
//
func QueryVisualTypes() (visual_types unsafe.Pointer) {
	iv, err := _I.Get(539, "query_visual_types", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_visual_types := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_count := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_visual_types, arg_count}
	iv.Call(args, nil, &outArgs[0])
	var count int32
	_ = count
	visual_types = outArgs[0].Pointer()
	count = outArgs[1].Int32()
	return
}

// gdk_selection_convert
//
// [ requestor ] trans: nothing
//
// [ selection ] trans: nothing
//
// [ target ] trans: nothing
//
// [ time_ ] trans: nothing
//
func SelectionConvert(requestor IWindow, selection Atom, target Atom, time_ uint32) {
	iv, err := _I.Get(540, "selection_convert", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if requestor != nil {
		tmp = requestor.P_Window()
	}
	arg_requestor := gi.NewPointerArgument(tmp)
	arg_selection := gi.NewPointerArgument(selection.P)
	arg_target := gi.NewPointerArgument(target.P)
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_requestor, arg_selection, arg_target, arg_time_}
	iv.Call(args, nil, nil)
}

// gdk_selection_owner_get
//
// [ selection ] trans: nothing
//
// [ result ] trans: nothing
//
func SelectionOwnerGet(selection Atom) (result Window) {
	iv, err := _I.Get(541, "selection_owner_get", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_selection := gi.NewPointerArgument(selection.P)
	args := []gi.Argument{arg_selection}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_selection_owner_get_for_display
//
// [ display ] trans: nothing
//
// [ selection ] trans: nothing
//
// [ result ] trans: nothing
//
func SelectionOwnerGetForDisplay(display IDisplay, selection Atom) (result Window) {
	iv, err := _I.Get(542, "selection_owner_get_for_display", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if display != nil {
		tmp = display.P_Display()
	}
	arg_display := gi.NewPointerArgument(tmp)
	arg_selection := gi.NewPointerArgument(selection.P)
	args := []gi.Argument{arg_display, arg_selection}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gdk_selection_owner_set
//
// [ owner ] trans: nothing
//
// [ selection ] trans: nothing
//
// [ time_ ] trans: nothing
//
// [ send_event ] trans: nothing
//
// [ result ] trans: nothing
//
func SelectionOwnerSet(owner IWindow, selection Atom, time_ uint32, send_event bool) (result bool) {
	iv, err := _I.Get(543, "selection_owner_set", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if owner != nil {
		tmp = owner.P_Window()
	}
	arg_owner := gi.NewPointerArgument(tmp)
	arg_selection := gi.NewPointerArgument(selection.P)
	arg_time_ := gi.NewUint32Argument(time_)
	arg_send_event := gi.NewBoolArgument(send_event)
	args := []gi.Argument{arg_owner, arg_selection, arg_time_, arg_send_event}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_selection_owner_set_for_display
//
// [ display ] trans: nothing
//
// [ owner ] trans: nothing
//
// [ selection ] trans: nothing
//
// [ time_ ] trans: nothing
//
// [ send_event ] trans: nothing
//
// [ result ] trans: nothing
//
func SelectionOwnerSetForDisplay(display IDisplay, owner IWindow, selection Atom, time_ uint32, send_event bool) (result bool) {
	iv, err := _I.Get(544, "selection_owner_set_for_display", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if display != nil {
		tmp = display.P_Display()
	}
	var tmp1 unsafe.Pointer
	if owner != nil {
		tmp1 = owner.P_Window()
	}
	arg_display := gi.NewPointerArgument(tmp)
	arg_owner := gi.NewPointerArgument(tmp1)
	arg_selection := gi.NewPointerArgument(selection.P)
	arg_time_ := gi.NewUint32Argument(time_)
	arg_send_event := gi.NewBoolArgument(send_event)
	args := []gi.Argument{arg_display, arg_owner, arg_selection, arg_time_, arg_send_event}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_selection_send_notify
//
// [ requestor ] trans: nothing
//
// [ selection ] trans: nothing
//
// [ target ] trans: nothing
//
// [ property ] trans: nothing
//
// [ time_ ] trans: nothing
//
func SelectionSendNotify(requestor IWindow, selection Atom, target Atom, property Atom, time_ uint32) {
	iv, err := _I.Get(545, "selection_send_notify", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if requestor != nil {
		tmp = requestor.P_Window()
	}
	arg_requestor := gi.NewPointerArgument(tmp)
	arg_selection := gi.NewPointerArgument(selection.P)
	arg_target := gi.NewPointerArgument(target.P)
	arg_property := gi.NewPointerArgument(property.P)
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_requestor, arg_selection, arg_target, arg_property, arg_time_}
	iv.Call(args, nil, nil)
}

// gdk_selection_send_notify_for_display
//
// [ display ] trans: nothing
//
// [ requestor ] trans: nothing
//
// [ selection ] trans: nothing
//
// [ target ] trans: nothing
//
// [ property ] trans: nothing
//
// [ time_ ] trans: nothing
//
func SelectionSendNotifyForDisplay(display IDisplay, requestor IWindow, selection Atom, target Atom, property Atom, time_ uint32) {
	iv, err := _I.Get(546, "selection_send_notify_for_display", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if display != nil {
		tmp = display.P_Display()
	}
	var tmp1 unsafe.Pointer
	if requestor != nil {
		tmp1 = requestor.P_Window()
	}
	arg_display := gi.NewPointerArgument(tmp)
	arg_requestor := gi.NewPointerArgument(tmp1)
	arg_selection := gi.NewPointerArgument(selection.P)
	arg_target := gi.NewPointerArgument(target.P)
	arg_property := gi.NewPointerArgument(property.P)
	arg_time_ := gi.NewUint32Argument(time_)
	args := []gi.Argument{arg_display, arg_requestor, arg_selection, arg_target, arg_property, arg_time_}
	iv.Call(args, nil, nil)
}

// gdk_set_allowed_backends
//
// [ backends ] trans: nothing
//
func SetAllowedBackends(backends string) {
	iv, err := _I.Get(547, "set_allowed_backends", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_backends := gi.CString(backends)
	arg_backends := gi.NewStringArgument(c_backends)
	args := []gi.Argument{arg_backends}
	iv.Call(args, nil, nil)
	gi.Free(c_backends)
}

// gdk_set_double_click_time
//
// [ msec ] trans: nothing
//
func SetDoubleClickTime(msec uint32) {
	iv, err := _I.Get(548, "set_double_click_time", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_msec := gi.NewUint32Argument(msec)
	args := []gi.Argument{arg_msec}
	iv.Call(args, nil, nil)
}

// gdk_set_program_class
//
// [ program_class ] trans: nothing
//
func SetProgramClass(program_class string) {
	iv, err := _I.Get(549, "set_program_class", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_program_class := gi.CString(program_class)
	arg_program_class := gi.NewStringArgument(c_program_class)
	args := []gi.Argument{arg_program_class}
	iv.Call(args, nil, nil)
	gi.Free(c_program_class)
}

// gdk_set_show_events
//
// [ show_events ] trans: nothing
//
func SetShowEvents(show_events bool) {
	iv, err := _I.Get(550, "set_show_events", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_show_events := gi.NewBoolArgument(show_events)
	args := []gi.Argument{arg_show_events}
	iv.Call(args, nil, nil)
}

// gdk_setting_get
//
// [ name ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func SettingGet(name string, value g.Value) (result bool) {
	iv, err := _I.Get(551, "setting_get", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_name, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = ret.Bool()
	return
}

// gdk_synthesize_window_state
//
// [ window ] trans: nothing
//
// [ unset_flags ] trans: nothing
//
// [ set_flags ] trans: nothing
//
func SynthesizeWindowState(window IWindow, unset_flags WindowStateFlags, set_flags WindowStateFlags) {
	iv, err := _I.Get(552, "synthesize_window_state", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_window := gi.NewPointerArgument(tmp)
	arg_unset_flags := gi.NewIntArgument(int(unset_flags))
	arg_set_flags := gi.NewIntArgument(int(set_flags))
	args := []gi.Argument{arg_window, arg_unset_flags, arg_set_flags}
	iv.Call(args, nil, nil)
}

// gdk_test_render_sync
//
// [ window ] trans: nothing
//
func TestRenderSync(window IWindow) {
	iv, err := _I.Get(553, "test_render_sync", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_window := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_window}
	iv.Call(args, nil, nil)
}

// gdk_test_simulate_button
//
// [ window ] trans: nothing
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ button ] trans: nothing
//
// [ modifiers ] trans: nothing
//
// [ button_pressrelease ] trans: nothing
//
// [ result ] trans: nothing
//
func TestSimulateButton(window IWindow, x int32, y int32, button uint32, modifiers ModifierTypeFlags, button_pressrelease EventTypeEnum) (result bool) {
	iv, err := _I.Get(554, "test_simulate_button", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_window := gi.NewPointerArgument(tmp)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_button := gi.NewUint32Argument(button)
	arg_modifiers := gi.NewIntArgument(int(modifiers))
	arg_button_pressrelease := gi.NewIntArgument(int(button_pressrelease))
	args := []gi.Argument{arg_window, arg_x, arg_y, arg_button, arg_modifiers, arg_button_pressrelease}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_test_simulate_key
//
// [ window ] trans: nothing
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ keyval ] trans: nothing
//
// [ modifiers ] trans: nothing
//
// [ key_pressrelease ] trans: nothing
//
// [ result ] trans: nothing
//
func TestSimulateKey(window IWindow, x int32, y int32, keyval uint32, modifiers ModifierTypeFlags, key_pressrelease EventTypeEnum) (result bool) {
	iv, err := _I.Get(555, "test_simulate_key", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if window != nil {
		tmp = window.P_Window()
	}
	arg_window := gi.NewPointerArgument(tmp)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_keyval := gi.NewUint32Argument(keyval)
	arg_modifiers := gi.NewIntArgument(int(modifiers))
	arg_key_pressrelease := gi.NewIntArgument(int(key_pressrelease))
	args := []gi.Argument{arg_window, arg_x, arg_y, arg_keyval, arg_modifiers, arg_key_pressrelease}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gdk_text_property_to_utf8_list_for_display
//
// [ display ] trans: nothing
//
// [ encoding ] trans: nothing
//
// [ format ] trans: nothing
//
// [ text ] trans: nothing
//
// [ length ] trans: nothing
//
// [ list ] trans: everything, dir: out
//
// [ result ] trans: nothing
//
func TextPropertyToUtf8ListForDisplay(display IDisplay, encoding Atom, format int32, text gi.Uint8Array, length int32) (result int32, list gi.CStrArray) {
	iv, err := _I.Get(556, "text_property_to_utf8_list_for_display", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if display != nil {
		tmp = display.P_Display()
	}
	arg_display := gi.NewPointerArgument(tmp)
	arg_encoding := gi.NewPointerArgument(encoding.P)
	arg_format := gi.NewInt32Argument(format)
	arg_text := gi.NewPointerArgument(text.P)
	arg_length := gi.NewInt32Argument(length)
	arg_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_display, arg_encoding, arg_format, arg_text, arg_length, arg_list}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	list.P = outArgs[0].Pointer()
	result = ret.Int32()
	return
}

// gdk_threads_add_idle_full
//
// [ priority ] trans: nothing
//
// [ function ] trans: nothing
//
// [ data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: nothing
//
func ThreadsAddIdle(priority int32, function int /*TODO_TYPE CALLBACK*/, data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result uint32) {
	iv, err := _I.Get(557, "threads_add_idle", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_priority := gi.NewInt32Argument(priority)
	arg_function := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_mySourceFunc()))
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_priority, arg_function, arg_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gdk_threads_add_timeout_full
//
// [ priority ] trans: nothing
//
// [ interval ] trans: nothing
//
// [ function ] trans: nothing
//
// [ data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: nothing
//
func ThreadsAddTimeout(priority int32, interval uint32, function int /*TODO_TYPE CALLBACK*/, data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result uint32) {
	iv, err := _I.Get(558, "threads_add_timeout", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_priority := gi.NewInt32Argument(priority)
	arg_interval := gi.NewUint32Argument(interval)
	arg_function := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_mySourceFunc()))
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_priority, arg_interval, arg_function, arg_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gdk_threads_add_timeout_seconds_full
//
// [ priority ] trans: nothing
//
// [ interval ] trans: nothing
//
// [ function ] trans: nothing
//
// [ data ] trans: nothing
//
// [ notify ] trans: nothing
//
// [ result ] trans: nothing
//
func ThreadsAddTimeoutSeconds(priority int32, interval uint32, function int /*TODO_TYPE CALLBACK*/, data unsafe.Pointer, notify int /*TODO_TYPE CALLBACK*/) (result uint32) {
	iv, err := _I.Get(559, "threads_add_timeout_seconds", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_priority := gi.NewInt32Argument(priority)
	arg_interval := gi.NewUint32Argument(interval)
	arg_function := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_mySourceFunc()))
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewPointerArgument(unsafe.Pointer(g.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_priority, arg_interval, arg_function, arg_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// Deprecated
//
// gdk_threads_enter
//
func ThreadsEnter() {
	iv, err := _I.Get(560, "threads_enter", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// Deprecated
//
// gdk_threads_init
//
func ThreadsInit() {
	iv, err := _I.Get(561, "threads_init", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// Deprecated
//
// gdk_threads_leave
//
func ThreadsLeave() {
	iv, err := _I.Get(562, "threads_leave", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// gdk_unicode_to_keyval
//
// [ wc ] trans: nothing
//
// [ result ] trans: nothing
//
func UnicodeToKeyval(wc uint32) (result uint32) {
	iv, err := _I.Get(563, "unicode_to_keyval", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_wc := gi.NewUint32Argument(wc)
	args := []gi.Argument{arg_wc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gdk_utf8_to_string_target
//
// [ str ] trans: nothing
//
// [ result ] trans: everything
//
func Utf8ToStringTarget(str string) (result string) {
	iv, err := _I.Get(564, "utf8_to_string_target", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// constants
const (
	BUTTON_MIDDLE                   = 2
	BUTTON_PRIMARY                  = 1
	BUTTON_SECONDARY                = 3
	CURRENT_TIME                    = 0
	EVENT_PROPAGATE                 = false
	EVENT_STOP                      = true
	KEY_0                           = 48
	KEY_1                           = 49
	KEY_2                           = 50
	KEY_3                           = 51
	KEY_3270_AltCursor              = 64784
	KEY_3270_Attn                   = 64782
	KEY_3270_BackTab                = 64773
	KEY_3270_ChangeScreen           = 64793
	KEY_3270_Copy                   = 64789
	KEY_3270_CursorBlink            = 64783
	KEY_3270_CursorSelect           = 64796
	KEY_3270_DeleteWord             = 64794
	KEY_3270_Duplicate              = 64769
	KEY_3270_Enter                  = 64798
	KEY_3270_EraseEOF               = 64774
	KEY_3270_EraseInput             = 64775
	KEY_3270_ExSelect               = 64795
	KEY_3270_FieldMark              = 64770
	KEY_3270_Ident                  = 64787
	KEY_3270_Jump                   = 64786
	KEY_3270_KeyClick               = 64785
	KEY_3270_Left2                  = 64772
	KEY_3270_PA1                    = 64778
	KEY_3270_PA2                    = 64779
	KEY_3270_PA3                    = 64780
	KEY_3270_Play                   = 64790
	KEY_3270_PrintScreen            = 64797
	KEY_3270_Quit                   = 64777
	KEY_3270_Record                 = 64792
	KEY_3270_Reset                  = 64776
	KEY_3270_Right2                 = 64771
	KEY_3270_Rule                   = 64788
	KEY_3270_Setup                  = 64791
	KEY_3270_Test                   = 64781
	KEY_4                           = 52
	KEY_5                           = 53
	KEY_6                           = 54
	KEY_7                           = 55
	KEY_8                           = 56
	KEY_9                           = 57
	KEY_A                           = 65
	KEY_AE                          = 198
	KEY_Aacute                      = 193
	KEY_Abelowdot                   = 16785056
	KEY_Abreve                      = 451
	KEY_Abreveacute                 = 16785070
	KEY_Abrevebelowdot              = 16785078
	KEY_Abrevegrave                 = 16785072
	KEY_Abrevehook                  = 16785074
	KEY_Abrevetilde                 = 16785076
	KEY_AccessX_Enable              = 65136
	KEY_AccessX_Feedback_Enable     = 65137
	KEY_Acircumflex                 = 194
	KEY_Acircumflexacute            = 16785060
	KEY_Acircumflexbelowdot         = 16785068
	KEY_Acircumflexgrave            = 16785062
	KEY_Acircumflexhook             = 16785064
	KEY_Acircumflextilde            = 16785066
	KEY_AddFavorite                 = 269025081
	KEY_Adiaeresis                  = 196
	KEY_Agrave                      = 192
	KEY_Ahook                       = 16785058
	KEY_Alt_L                       = 65513
	KEY_Alt_R                       = 65514
	KEY_Amacron                     = 960
	KEY_Aogonek                     = 417
	KEY_ApplicationLeft             = 269025104
	KEY_ApplicationRight            = 269025105
	KEY_Arabic_0                    = 16778848
	KEY_Arabic_1                    = 16778849
	KEY_Arabic_2                    = 16778850
	KEY_Arabic_3                    = 16778851
	KEY_Arabic_4                    = 16778852
	KEY_Arabic_5                    = 16778853
	KEY_Arabic_6                    = 16778854
	KEY_Arabic_7                    = 16778855
	KEY_Arabic_8                    = 16778856
	KEY_Arabic_9                    = 16778857
	KEY_Arabic_ain                  = 1497
	KEY_Arabic_alef                 = 1479
	KEY_Arabic_alefmaksura          = 1513
	KEY_Arabic_beh                  = 1480
	KEY_Arabic_comma                = 1452
	KEY_Arabic_dad                  = 1494
	KEY_Arabic_dal                  = 1487
	KEY_Arabic_damma                = 1519
	KEY_Arabic_dammatan             = 1516
	KEY_Arabic_ddal                 = 16778888
	KEY_Arabic_farsi_yeh            = 16778956
	KEY_Arabic_fatha                = 1518
	KEY_Arabic_fathatan             = 1515
	KEY_Arabic_feh                  = 1505
	KEY_Arabic_fullstop             = 16778964
	KEY_Arabic_gaf                  = 16778927
	KEY_Arabic_ghain                = 1498
	KEY_Arabic_ha                   = 1511
	KEY_Arabic_hah                  = 1485
	KEY_Arabic_hamza                = 1473
	KEY_Arabic_hamza_above          = 16778836
	KEY_Arabic_hamza_below          = 16778837
	KEY_Arabic_hamzaonalef          = 1475
	KEY_Arabic_hamzaonwaw           = 1476
	KEY_Arabic_hamzaonyeh           = 1478
	KEY_Arabic_hamzaunderalef       = 1477
	KEY_Arabic_heh                  = 1511
	KEY_Arabic_heh_doachashmee      = 16778942
	KEY_Arabic_heh_goal             = 16778945
	KEY_Arabic_jeem                 = 1484
	KEY_Arabic_jeh                  = 16778904
	KEY_Arabic_kaf                  = 1507
	KEY_Arabic_kasra                = 1520
	KEY_Arabic_kasratan             = 1517
	KEY_Arabic_keheh                = 16778921
	KEY_Arabic_khah                 = 1486
	KEY_Arabic_lam                  = 1508
	KEY_Arabic_madda_above          = 16778835
	KEY_Arabic_maddaonalef          = 1474
	KEY_Arabic_meem                 = 1509
	KEY_Arabic_noon                 = 1510
	KEY_Arabic_noon_ghunna          = 16778938
	KEY_Arabic_peh                  = 16778878
	KEY_Arabic_percent              = 16778858
	KEY_Arabic_qaf                  = 1506
	KEY_Arabic_question_mark        = 1471
	KEY_Arabic_ra                   = 1489
	KEY_Arabic_rreh                 = 16778897
	KEY_Arabic_sad                  = 1493
	KEY_Arabic_seen                 = 1491
	KEY_Arabic_semicolon            = 1467
	KEY_Arabic_shadda               = 1521
	KEY_Arabic_sheen                = 1492
	KEY_Arabic_sukun                = 1522
	KEY_Arabic_superscript_alef     = 16778864
	KEY_Arabic_switch               = 65406
	KEY_Arabic_tah                  = 1495
	KEY_Arabic_tatweel              = 1504
	KEY_Arabic_tcheh                = 16778886
	KEY_Arabic_teh                  = 1482
	KEY_Arabic_tehmarbuta           = 1481
	KEY_Arabic_thal                 = 1488
	KEY_Arabic_theh                 = 1483
	KEY_Arabic_tteh                 = 16778873
	KEY_Arabic_veh                  = 16778916
	KEY_Arabic_waw                  = 1512
	KEY_Arabic_yeh                  = 1514
	KEY_Arabic_yeh_baree            = 16778962
	KEY_Arabic_zah                  = 1496
	KEY_Arabic_zain                 = 1490
	KEY_Aring                       = 197
	KEY_Armenian_AT                 = 16778552
	KEY_Armenian_AYB                = 16778545
	KEY_Armenian_BEN                = 16778546
	KEY_Armenian_CHA                = 16778569
	KEY_Armenian_DA                 = 16778548
	KEY_Armenian_DZA                = 16778561
	KEY_Armenian_E                  = 16778551
	KEY_Armenian_FE                 = 16778582
	KEY_Armenian_GHAT               = 16778562
	KEY_Armenian_GIM                = 16778547
	KEY_Armenian_HI                 = 16778565
	KEY_Armenian_HO                 = 16778560
	KEY_Armenian_INI                = 16778555
	KEY_Armenian_JE                 = 16778571
	KEY_Armenian_KE                 = 16778580
	KEY_Armenian_KEN                = 16778559
	KEY_Armenian_KHE                = 16778557
	KEY_Armenian_LYUN               = 16778556
	KEY_Armenian_MEN                = 16778564
	KEY_Armenian_NU                 = 16778566
	KEY_Armenian_O                  = 16778581
	KEY_Armenian_PE                 = 16778570
	KEY_Armenian_PYUR               = 16778579
	KEY_Armenian_RA                 = 16778572
	KEY_Armenian_RE                 = 16778576
	KEY_Armenian_SE                 = 16778573
	KEY_Armenian_SHA                = 16778567
	KEY_Armenian_TCHE               = 16778563
	KEY_Armenian_TO                 = 16778553
	KEY_Armenian_TSA                = 16778558
	KEY_Armenian_TSO                = 16778577
	KEY_Armenian_TYUN               = 16778575
	KEY_Armenian_VEV                = 16778574
	KEY_Armenian_VO                 = 16778568
	KEY_Armenian_VYUN               = 16778578
	KEY_Armenian_YECH               = 16778549
	KEY_Armenian_ZA                 = 16778550
	KEY_Armenian_ZHE                = 16778554
	KEY_Armenian_accent             = 16778587
	KEY_Armenian_amanak             = 16778588
	KEY_Armenian_apostrophe         = 16778586
	KEY_Armenian_at                 = 16778600
	KEY_Armenian_ayb                = 16778593
	KEY_Armenian_ben                = 16778594
	KEY_Armenian_but                = 16778589
	KEY_Armenian_cha                = 16778617
	KEY_Armenian_da                 = 16778596
	KEY_Armenian_dza                = 16778609
	KEY_Armenian_e                  = 16778599
	KEY_Armenian_exclam             = 16778588
	KEY_Armenian_fe                 = 16778630
	KEY_Armenian_full_stop          = 16778633
	KEY_Armenian_ghat               = 16778610
	KEY_Armenian_gim                = 16778595
	KEY_Armenian_hi                 = 16778613
	KEY_Armenian_ho                 = 16778608
	KEY_Armenian_hyphen             = 16778634
	KEY_Armenian_ini                = 16778603
	KEY_Armenian_je                 = 16778619
	KEY_Armenian_ke                 = 16778628
	KEY_Armenian_ken                = 16778607
	KEY_Armenian_khe                = 16778605
	KEY_Armenian_ligature_ew        = 16778631
	KEY_Armenian_lyun               = 16778604
	KEY_Armenian_men                = 16778612
	KEY_Armenian_nu                 = 16778614
	KEY_Armenian_o                  = 16778629
	KEY_Armenian_paruyk             = 16778590
	KEY_Armenian_pe                 = 16778618
	KEY_Armenian_pyur               = 16778627
	KEY_Armenian_question           = 16778590
	KEY_Armenian_ra                 = 16778620
	KEY_Armenian_re                 = 16778624
	KEY_Armenian_se                 = 16778621
	KEY_Armenian_separation_mark    = 16778589
	KEY_Armenian_sha                = 16778615
	KEY_Armenian_shesht             = 16778587
	KEY_Armenian_tche               = 16778611
	KEY_Armenian_to                 = 16778601
	KEY_Armenian_tsa                = 16778606
	KEY_Armenian_tso                = 16778625
	KEY_Armenian_tyun               = 16778623
	KEY_Armenian_verjaket           = 16778633
	KEY_Armenian_vev                = 16778622
	KEY_Armenian_vo                 = 16778616
	KEY_Armenian_vyun               = 16778626
	KEY_Armenian_yech               = 16778597
	KEY_Armenian_yentamna           = 16778634
	KEY_Armenian_za                 = 16778598
	KEY_Armenian_zhe                = 16778602
	KEY_Atilde                      = 195
	KEY_AudibleBell_Enable          = 65146
	KEY_AudioCycleTrack             = 269025179
	KEY_AudioForward                = 269025175
	KEY_AudioLowerVolume            = 269025041
	KEY_AudioMedia                  = 269025074
	KEY_AudioMicMute                = 269025202
	KEY_AudioMute                   = 269025042
	KEY_AudioNext                   = 269025047
	KEY_AudioPause                  = 269025073
	KEY_AudioPlay                   = 269025044
	KEY_AudioPrev                   = 269025046
	KEY_AudioRaiseVolume            = 269025043
	KEY_AudioRandomPlay             = 269025177
	KEY_AudioRecord                 = 269025052
	KEY_AudioRepeat                 = 269025176
	KEY_AudioRewind                 = 269025086
	KEY_AudioStop                   = 269025045
	KEY_Away                        = 269025165
	KEY_B                           = 66
	KEY_Babovedot                   = 16784898
	KEY_Back                        = 269025062
	KEY_BackForward                 = 269025087
	KEY_BackSpace                   = 65288
	KEY_Battery                     = 269025171
	KEY_Begin                       = 65368
	KEY_Blue                        = 269025190
	KEY_Bluetooth                   = 269025172
	KEY_Book                        = 269025106
	KEY_BounceKeys_Enable           = 65140
	KEY_Break                       = 65387
	KEY_BrightnessAdjust            = 269025083
	KEY_Byelorussian_SHORTU         = 1726
	KEY_Byelorussian_shortu         = 1710
	KEY_C                           = 67
	KEY_CD                          = 269025107
	KEY_CH                          = 65186
	KEY_C_H                         = 65189
	KEY_C_h                         = 65188
	KEY_Cabovedot                   = 709
	KEY_Cacute                      = 454
	KEY_Calculator                  = 269025053
	KEY_Calendar                    = 269025056
	KEY_Cancel                      = 65385
	KEY_Caps_Lock                   = 65509
	KEY_Ccaron                      = 456
	KEY_Ccedilla                    = 199
	KEY_Ccircumflex                 = 710
	KEY_Ch                          = 65185
	KEY_Clear                       = 65291
	KEY_ClearGrab                   = 269024801
	KEY_Close                       = 269025110
	KEY_Codeinput                   = 65335
	KEY_ColonSign                   = 16785569
	KEY_Community                   = 269025085
	KEY_ContrastAdjust              = 269025058
	KEY_Control_L                   = 65507
	KEY_Control_R                   = 65508
	KEY_Copy                        = 269025111
	KEY_CruzeiroSign                = 16785570
	KEY_Cut                         = 269025112
	KEY_CycleAngle                  = 269025180
	KEY_Cyrillic_A                  = 1761
	KEY_Cyrillic_BE                 = 1762
	KEY_Cyrillic_CHE                = 1790
	KEY_Cyrillic_CHE_descender      = 16778422
	KEY_Cyrillic_CHE_vertstroke     = 16778424
	KEY_Cyrillic_DE                 = 1764
	KEY_Cyrillic_DZHE               = 1727
	KEY_Cyrillic_E                  = 1788
	KEY_Cyrillic_EF                 = 1766
	KEY_Cyrillic_EL                 = 1772
	KEY_Cyrillic_EM                 = 1773
	KEY_Cyrillic_EN                 = 1774
	KEY_Cyrillic_EN_descender       = 16778402
	KEY_Cyrillic_ER                 = 1778
	KEY_Cyrillic_ES                 = 1779
	KEY_Cyrillic_GHE                = 1767
	KEY_Cyrillic_GHE_bar            = 16778386
	KEY_Cyrillic_HA                 = 1768
	KEY_Cyrillic_HARDSIGN           = 1791
	KEY_Cyrillic_HA_descender       = 16778418
	KEY_Cyrillic_I                  = 1769
	KEY_Cyrillic_IE                 = 1765
	KEY_Cyrillic_IO                 = 1715
	KEY_Cyrillic_I_macron           = 16778466
	KEY_Cyrillic_JE                 = 1720
	KEY_Cyrillic_KA                 = 1771
	KEY_Cyrillic_KA_descender       = 16778394
	KEY_Cyrillic_KA_vertstroke      = 16778396
	KEY_Cyrillic_LJE                = 1721
	KEY_Cyrillic_NJE                = 1722
	KEY_Cyrillic_O                  = 1775
	KEY_Cyrillic_O_bar              = 16778472
	KEY_Cyrillic_PE                 = 1776
	KEY_Cyrillic_SCHWA              = 16778456
	KEY_Cyrillic_SHA                = 1787
	KEY_Cyrillic_SHCHA              = 1789
	KEY_Cyrillic_SHHA               = 16778426
	KEY_Cyrillic_SHORTI             = 1770
	KEY_Cyrillic_SOFTSIGN           = 1784
	KEY_Cyrillic_TE                 = 1780
	KEY_Cyrillic_TSE                = 1763
	KEY_Cyrillic_U                  = 1781
	KEY_Cyrillic_U_macron           = 16778478
	KEY_Cyrillic_U_straight         = 16778414
	KEY_Cyrillic_U_straight_bar     = 16778416
	KEY_Cyrillic_VE                 = 1783
	KEY_Cyrillic_YA                 = 1777
	KEY_Cyrillic_YERU               = 1785
	KEY_Cyrillic_YU                 = 1760
	KEY_Cyrillic_ZE                 = 1786
	KEY_Cyrillic_ZHE                = 1782
	KEY_Cyrillic_ZHE_descender      = 16778390
	KEY_Cyrillic_a                  = 1729
	KEY_Cyrillic_be                 = 1730
	KEY_Cyrillic_che                = 1758
	KEY_Cyrillic_che_descender      = 16778423
	KEY_Cyrillic_che_vertstroke     = 16778425
	KEY_Cyrillic_de                 = 1732
	KEY_Cyrillic_dzhe               = 1711
	KEY_Cyrillic_e                  = 1756
	KEY_Cyrillic_ef                 = 1734
	KEY_Cyrillic_el                 = 1740
	KEY_Cyrillic_em                 = 1741
	KEY_Cyrillic_en                 = 1742
	KEY_Cyrillic_en_descender       = 16778403
	KEY_Cyrillic_er                 = 1746
	KEY_Cyrillic_es                 = 1747
	KEY_Cyrillic_ghe                = 1735
	KEY_Cyrillic_ghe_bar            = 16778387
	KEY_Cyrillic_ha                 = 1736
	KEY_Cyrillic_ha_descender       = 16778419
	KEY_Cyrillic_hardsign           = 1759
	KEY_Cyrillic_i                  = 1737
	KEY_Cyrillic_i_macron           = 16778467
	KEY_Cyrillic_ie                 = 1733
	KEY_Cyrillic_io                 = 1699
	KEY_Cyrillic_je                 = 1704
	KEY_Cyrillic_ka                 = 1739
	KEY_Cyrillic_ka_descender       = 16778395
	KEY_Cyrillic_ka_vertstroke      = 16778397
	KEY_Cyrillic_lje                = 1705
	KEY_Cyrillic_nje                = 1706
	KEY_Cyrillic_o                  = 1743
	KEY_Cyrillic_o_bar              = 16778473
	KEY_Cyrillic_pe                 = 1744
	KEY_Cyrillic_schwa              = 16778457
	KEY_Cyrillic_sha                = 1755
	KEY_Cyrillic_shcha              = 1757
	KEY_Cyrillic_shha               = 16778427
	KEY_Cyrillic_shorti             = 1738
	KEY_Cyrillic_softsign           = 1752
	KEY_Cyrillic_te                 = 1748
	KEY_Cyrillic_tse                = 1731
	KEY_Cyrillic_u                  = 1749
	KEY_Cyrillic_u_macron           = 16778479
	KEY_Cyrillic_u_straight         = 16778415
	KEY_Cyrillic_u_straight_bar     = 16778417
	KEY_Cyrillic_ve                 = 1751
	KEY_Cyrillic_ya                 = 1745
	KEY_Cyrillic_yeru               = 1753
	KEY_Cyrillic_yu                 = 1728
	KEY_Cyrillic_ze                 = 1754
	KEY_Cyrillic_zhe                = 1750
	KEY_Cyrillic_zhe_descender      = 16778391
	KEY_D                           = 68
	KEY_DOS                         = 269025114
	KEY_Dabovedot                   = 16784906
	KEY_Dcaron                      = 463
	KEY_Delete                      = 65535
	KEY_Display                     = 269025113
	KEY_Documents                   = 269025115
	KEY_DongSign                    = 16785579
	KEY_Down                        = 65364
	KEY_Dstroke                     = 464
	KEY_E                           = 69
	KEY_ENG                         = 957
	KEY_ETH                         = 208
	KEY_EZH                         = 16777655
	KEY_Eabovedot                   = 972
	KEY_Eacute                      = 201
	KEY_Ebelowdot                   = 16785080
	KEY_Ecaron                      = 460
	KEY_Ecircumflex                 = 202
	KEY_Ecircumflexacute            = 16785086
	KEY_Ecircumflexbelowdot         = 16785094
	KEY_Ecircumflexgrave            = 16785088
	KEY_Ecircumflexhook             = 16785090
	KEY_Ecircumflextilde            = 16785092
	KEY_EcuSign                     = 16785568
	KEY_Ediaeresis                  = 203
	KEY_Egrave                      = 200
	KEY_Ehook                       = 16785082
	KEY_Eisu_Shift                  = 65327
	KEY_Eisu_toggle                 = 65328
	KEY_Eject                       = 269025068
	KEY_Emacron                     = 938
	KEY_End                         = 65367
	KEY_Eogonek                     = 458
	KEY_Escape                      = 65307
	KEY_Eth                         = 208
	KEY_Etilde                      = 16785084
	KEY_EuroSign                    = 8364
	KEY_Excel                       = 269025116
	KEY_Execute                     = 65378
	KEY_Explorer                    = 269025117
	KEY_F                           = 70
	KEY_F1                          = 65470
	KEY_F10                         = 65479
	KEY_F11                         = 65480
	KEY_F12                         = 65481
	KEY_F13                         = 65482
	KEY_F14                         = 65483
	KEY_F15                         = 65484
	KEY_F16                         = 65485
	KEY_F17                         = 65486
	KEY_F18                         = 65487
	KEY_F19                         = 65488
	KEY_F2                          = 65471
	KEY_F20                         = 65489
	KEY_F21                         = 65490
	KEY_F22                         = 65491
	KEY_F23                         = 65492
	KEY_F24                         = 65493
	KEY_F25                         = 65494
	KEY_F26                         = 65495
	KEY_F27                         = 65496
	KEY_F28                         = 65497
	KEY_F29                         = 65498
	KEY_F3                          = 65472
	KEY_F30                         = 65499
	KEY_F31                         = 65500
	KEY_F32                         = 65501
	KEY_F33                         = 65502
	KEY_F34                         = 65503
	KEY_F35                         = 65504
	KEY_F4                          = 65473
	KEY_F5                          = 65474
	KEY_F6                          = 65475
	KEY_F7                          = 65476
	KEY_F8                          = 65477
	KEY_F9                          = 65478
	KEY_FFrancSign                  = 16785571
	KEY_Fabovedot                   = 16784926
	KEY_Farsi_0                     = 16778992
	KEY_Farsi_1                     = 16778993
	KEY_Farsi_2                     = 16778994
	KEY_Farsi_3                     = 16778995
	KEY_Farsi_4                     = 16778996
	KEY_Farsi_5                     = 16778997
	KEY_Farsi_6                     = 16778998
	KEY_Farsi_7                     = 16778999
	KEY_Farsi_8                     = 16779000
	KEY_Farsi_9                     = 16779001
	KEY_Farsi_yeh                   = 16778956
	KEY_Favorites                   = 269025072
	KEY_Finance                     = 269025084
	KEY_Find                        = 65384
	KEY_First_Virtual_Screen        = 65232
	KEY_Forward                     = 269025063
	KEY_FrameBack                   = 269025181
	KEY_FrameForward                = 269025182
	KEY_G                           = 71
	KEY_Gabovedot                   = 725
	KEY_Game                        = 269025118
	KEY_Gbreve                      = 683
	KEY_Gcaron                      = 16777702
	KEY_Gcedilla                    = 939
	KEY_Gcircumflex                 = 728
	KEY_Georgian_an                 = 16781520
	KEY_Georgian_ban                = 16781521
	KEY_Georgian_can                = 16781546
	KEY_Georgian_char               = 16781549
	KEY_Georgian_chin               = 16781545
	KEY_Georgian_cil                = 16781548
	KEY_Georgian_don                = 16781523
	KEY_Georgian_en                 = 16781524
	KEY_Georgian_fi                 = 16781558
	KEY_Georgian_gan                = 16781522
	KEY_Georgian_ghan               = 16781542
	KEY_Georgian_hae                = 16781552
	KEY_Georgian_har                = 16781556
	KEY_Georgian_he                 = 16781553
	KEY_Georgian_hie                = 16781554
	KEY_Georgian_hoe                = 16781557
	KEY_Georgian_in                 = 16781528
	KEY_Georgian_jhan               = 16781551
	KEY_Georgian_jil                = 16781547
	KEY_Georgian_kan                = 16781529
	KEY_Georgian_khar               = 16781541
	KEY_Georgian_las                = 16781530
	KEY_Georgian_man                = 16781531
	KEY_Georgian_nar                = 16781532
	KEY_Georgian_on                 = 16781533
	KEY_Georgian_par                = 16781534
	KEY_Georgian_phar               = 16781540
	KEY_Georgian_qar                = 16781543
	KEY_Georgian_rae                = 16781536
	KEY_Georgian_san                = 16781537
	KEY_Georgian_shin               = 16781544
	KEY_Georgian_tan                = 16781527
	KEY_Georgian_tar                = 16781538
	KEY_Georgian_un                 = 16781539
	KEY_Georgian_vin                = 16781525
	KEY_Georgian_we                 = 16781555
	KEY_Georgian_xan                = 16781550
	KEY_Georgian_zen                = 16781526
	KEY_Georgian_zhar               = 16781535
	KEY_Go                          = 269025119
	KEY_Greek_ALPHA                 = 1985
	KEY_Greek_ALPHAaccent           = 1953
	KEY_Greek_BETA                  = 1986
	KEY_Greek_CHI                   = 2007
	KEY_Greek_DELTA                 = 1988
	KEY_Greek_EPSILON               = 1989
	KEY_Greek_EPSILONaccent         = 1954
	KEY_Greek_ETA                   = 1991
	KEY_Greek_ETAaccent             = 1955
	KEY_Greek_GAMMA                 = 1987
	KEY_Greek_IOTA                  = 1993
	KEY_Greek_IOTAaccent            = 1956
	KEY_Greek_IOTAdiaeresis         = 1957
	KEY_Greek_IOTAdieresis          = 1957
	KEY_Greek_KAPPA                 = 1994
	KEY_Greek_LAMBDA                = 1995
	KEY_Greek_LAMDA                 = 1995
	KEY_Greek_MU                    = 1996
	KEY_Greek_NU                    = 1997
	KEY_Greek_OMEGA                 = 2009
	KEY_Greek_OMEGAaccent           = 1963
	KEY_Greek_OMICRON               = 1999
	KEY_Greek_OMICRONaccent         = 1959
	KEY_Greek_PHI                   = 2006
	KEY_Greek_PI                    = 2000
	KEY_Greek_PSI                   = 2008
	KEY_Greek_RHO                   = 2001
	KEY_Greek_SIGMA                 = 2002
	KEY_Greek_TAU                   = 2004
	KEY_Greek_THETA                 = 1992
	KEY_Greek_UPSILON               = 2005
	KEY_Greek_UPSILONaccent         = 1960
	KEY_Greek_UPSILONdieresis       = 1961
	KEY_Greek_XI                    = 1998
	KEY_Greek_ZETA                  = 1990
	KEY_Greek_accentdieresis        = 1966
	KEY_Greek_alpha                 = 2017
	KEY_Greek_alphaaccent           = 1969
	KEY_Greek_beta                  = 2018
	KEY_Greek_chi                   = 2039
	KEY_Greek_delta                 = 2020
	KEY_Greek_epsilon               = 2021
	KEY_Greek_epsilonaccent         = 1970
	KEY_Greek_eta                   = 2023
	KEY_Greek_etaaccent             = 1971
	KEY_Greek_finalsmallsigma       = 2035
	KEY_Greek_gamma                 = 2019
	KEY_Greek_horizbar              = 1967
	KEY_Greek_iota                  = 2025
	KEY_Greek_iotaaccent            = 1972
	KEY_Greek_iotaaccentdieresis    = 1974
	KEY_Greek_iotadieresis          = 1973
	KEY_Greek_kappa                 = 2026
	KEY_Greek_lambda                = 2027
	KEY_Greek_lamda                 = 2027
	KEY_Greek_mu                    = 2028
	KEY_Greek_nu                    = 2029
	KEY_Greek_omega                 = 2041
	KEY_Greek_omegaaccent           = 1979
	KEY_Greek_omicron               = 2031
	KEY_Greek_omicronaccent         = 1975
	KEY_Greek_phi                   = 2038
	KEY_Greek_pi                    = 2032
	KEY_Greek_psi                   = 2040
	KEY_Greek_rho                   = 2033
	KEY_Greek_sigma                 = 2034
	KEY_Greek_switch                = 65406
	KEY_Greek_tau                   = 2036
	KEY_Greek_theta                 = 2024
	KEY_Greek_upsilon               = 2037
	KEY_Greek_upsilonaccent         = 1976
	KEY_Greek_upsilonaccentdieresis = 1978
	KEY_Greek_upsilondieresis       = 1977
	KEY_Greek_xi                    = 2030
	KEY_Greek_zeta                  = 2022
	KEY_Green                       = 269025188
	KEY_H                           = 72
	KEY_Hangul                      = 65329
	KEY_Hangul_A                    = 3775
	KEY_Hangul_AE                   = 3776
	KEY_Hangul_AraeA                = 3830
	KEY_Hangul_AraeAE               = 3831
	KEY_Hangul_Banja                = 65337
	KEY_Hangul_Cieuc                = 3770
	KEY_Hangul_Codeinput            = 65335
	KEY_Hangul_Dikeud               = 3751
	KEY_Hangul_E                    = 3780
	KEY_Hangul_EO                   = 3779
	KEY_Hangul_EU                   = 3793
	KEY_Hangul_End                  = 65331
	KEY_Hangul_Hanja                = 65332
	KEY_Hangul_Hieuh                = 3774
	KEY_Hangul_I                    = 3795
	KEY_Hangul_Ieung                = 3767
	KEY_Hangul_J_Cieuc              = 3818
	KEY_Hangul_J_Dikeud             = 3802
	KEY_Hangul_J_Hieuh              = 3822
	KEY_Hangul_J_Ieung              = 3816
	KEY_Hangul_J_Jieuj              = 3817
	KEY_Hangul_J_Khieuq             = 3819
	KEY_Hangul_J_Kiyeog             = 3796
	KEY_Hangul_J_KiyeogSios         = 3798
	KEY_Hangul_J_KkogjiDalrinIeung  = 3833
	KEY_Hangul_J_Mieum              = 3811
	KEY_Hangul_J_Nieun              = 3799
	KEY_Hangul_J_NieunHieuh         = 3801
	KEY_Hangul_J_NieunJieuj         = 3800
	KEY_Hangul_J_PanSios            = 3832
	KEY_Hangul_J_Phieuf             = 3821
	KEY_Hangul_J_Pieub              = 3812
	KEY_Hangul_J_PieubSios          = 3813
	KEY_Hangul_J_Rieul              = 3803
	KEY_Hangul_J_RieulHieuh         = 3810
	KEY_Hangul_J_RieulKiyeog        = 3804
	KEY_Hangul_J_RieulMieum         = 3805
	KEY_Hangul_J_RieulPhieuf        = 3809
	KEY_Hangul_J_RieulPieub         = 3806
	KEY_Hangul_J_RieulSios          = 3807
	KEY_Hangul_J_RieulTieut         = 3808
	KEY_Hangul_J_Sios               = 3814
	KEY_Hangul_J_SsangKiyeog        = 3797
	KEY_Hangul_J_SsangSios          = 3815
	KEY_Hangul_J_Tieut              = 3820
	KEY_Hangul_J_YeorinHieuh        = 3834
	KEY_Hangul_Jamo                 = 65333
	KEY_Hangul_Jeonja               = 65336
	KEY_Hangul_Jieuj                = 3768
	KEY_Hangul_Khieuq               = 3771
	KEY_Hangul_Kiyeog               = 3745
	KEY_Hangul_KiyeogSios           = 3747
	KEY_Hangul_KkogjiDalrinIeung    = 3827
	KEY_Hangul_Mieum                = 3761
	KEY_Hangul_MultipleCandidate    = 65341
	KEY_Hangul_Nieun                = 3748
	KEY_Hangul_NieunHieuh           = 3750
	KEY_Hangul_NieunJieuj           = 3749
	KEY_Hangul_O                    = 3783
	KEY_Hangul_OE                   = 3786
	KEY_Hangul_PanSios              = 3826
	KEY_Hangul_Phieuf               = 3773
	KEY_Hangul_Pieub                = 3762
	KEY_Hangul_PieubSios            = 3764
	KEY_Hangul_PostHanja            = 65339
	KEY_Hangul_PreHanja             = 65338
	KEY_Hangul_PreviousCandidate    = 65342
	KEY_Hangul_Rieul                = 3753
	KEY_Hangul_RieulHieuh           = 3760
	KEY_Hangul_RieulKiyeog          = 3754
	KEY_Hangul_RieulMieum           = 3755
	KEY_Hangul_RieulPhieuf          = 3759
	KEY_Hangul_RieulPieub           = 3756
	KEY_Hangul_RieulSios            = 3757
	KEY_Hangul_RieulTieut           = 3758
	KEY_Hangul_RieulYeorinHieuh     = 3823
	KEY_Hangul_Romaja               = 65334
	KEY_Hangul_SingleCandidate      = 65340
	KEY_Hangul_Sios                 = 3765
	KEY_Hangul_Special              = 65343
	KEY_Hangul_SsangDikeud          = 3752
	KEY_Hangul_SsangJieuj           = 3769
	KEY_Hangul_SsangKiyeog          = 3746
	KEY_Hangul_SsangPieub           = 3763
	KEY_Hangul_SsangSios            = 3766
	KEY_Hangul_Start                = 65330
	KEY_Hangul_SunkyeongeumMieum    = 3824
	KEY_Hangul_SunkyeongeumPhieuf   = 3828
	KEY_Hangul_SunkyeongeumPieub    = 3825
	KEY_Hangul_Tieut                = 3772
	KEY_Hangul_U                    = 3788
	KEY_Hangul_WA                   = 3784
	KEY_Hangul_WAE                  = 3785
	KEY_Hangul_WE                   = 3790
	KEY_Hangul_WEO                  = 3789
	KEY_Hangul_WI                   = 3791
	KEY_Hangul_YA                   = 3777
	KEY_Hangul_YAE                  = 3778
	KEY_Hangul_YE                   = 3782
	KEY_Hangul_YEO                  = 3781
	KEY_Hangul_YI                   = 3794
	KEY_Hangul_YO                   = 3787
	KEY_Hangul_YU                   = 3792
	KEY_Hangul_YeorinHieuh          = 3829
	KEY_Hangul_switch               = 65406
	KEY_Hankaku                     = 65321
	KEY_Hcircumflex                 = 678
	KEY_Hebrew_switch               = 65406
	KEY_Help                        = 65386
	KEY_Henkan                      = 65315
	KEY_Henkan_Mode                 = 65315
	KEY_Hibernate                   = 269025192
	KEY_Hiragana                    = 65317
	KEY_Hiragana_Katakana           = 65319
	KEY_History                     = 269025079
	KEY_Home                        = 65360
	KEY_HomePage                    = 269025048
	KEY_HotLinks                    = 269025082
	KEY_Hstroke                     = 673
	KEY_Hyper_L                     = 65517
	KEY_Hyper_R                     = 65518
	KEY_I                           = 73
	KEY_ISO_Center_Object           = 65075
	KEY_ISO_Continuous_Underline    = 65072
	KEY_ISO_Discontinuous_Underline = 65073
	KEY_ISO_Emphasize               = 65074
	KEY_ISO_Enter                   = 65076
	KEY_ISO_Fast_Cursor_Down        = 65071
	KEY_ISO_Fast_Cursor_Left        = 65068
	KEY_ISO_Fast_Cursor_Right       = 65069
	KEY_ISO_Fast_Cursor_Up          = 65070
	KEY_ISO_First_Group             = 65036
	KEY_ISO_First_Group_Lock        = 65037
	KEY_ISO_Group_Latch             = 65030
	KEY_ISO_Group_Lock              = 65031
	KEY_ISO_Group_Shift             = 65406
	KEY_ISO_Last_Group              = 65038
	KEY_ISO_Last_Group_Lock         = 65039
	KEY_ISO_Left_Tab                = 65056
	KEY_ISO_Level2_Latch            = 65026
	KEY_ISO_Level3_Latch            = 65028
	KEY_ISO_Level3_Lock             = 65029
	KEY_ISO_Level3_Shift            = 65027
	KEY_ISO_Level5_Latch            = 65042
	KEY_ISO_Level5_Lock             = 65043
	KEY_ISO_Level5_Shift            = 65041
	KEY_ISO_Lock                    = 65025
	KEY_ISO_Move_Line_Down          = 65058
	KEY_ISO_Move_Line_Up            = 65057
	KEY_ISO_Next_Group              = 65032
	KEY_ISO_Next_Group_Lock         = 65033
	KEY_ISO_Partial_Line_Down       = 65060
	KEY_ISO_Partial_Line_Up         = 65059
	KEY_ISO_Partial_Space_Left      = 65061
	KEY_ISO_Partial_Space_Right     = 65062
	KEY_ISO_Prev_Group              = 65034
	KEY_ISO_Prev_Group_Lock         = 65035
	KEY_ISO_Release_Both_Margins    = 65067
	KEY_ISO_Release_Margin_Left     = 65065
	KEY_ISO_Release_Margin_Right    = 65066
	KEY_ISO_Set_Margin_Left         = 65063
	KEY_ISO_Set_Margin_Right        = 65064
	KEY_Iabovedot                   = 681
	KEY_Iacute                      = 205
	KEY_Ibelowdot                   = 16785098
	KEY_Ibreve                      = 16777516
	KEY_Icircumflex                 = 206
	KEY_Idiaeresis                  = 207
	KEY_Igrave                      = 204
	KEY_Ihook                       = 16785096
	KEY_Imacron                     = 975
	KEY_Insert                      = 65379
	KEY_Iogonek                     = 967
	KEY_Itilde                      = 933
	KEY_J                           = 74
	KEY_Jcircumflex                 = 684
	KEY_K                           = 75
	KEY_KP_0                        = 65456
	KEY_KP_1                        = 65457
	KEY_KP_2                        = 65458
	KEY_KP_3                        = 65459
	KEY_KP_4                        = 65460
	KEY_KP_5                        = 65461
	KEY_KP_6                        = 65462
	KEY_KP_7                        = 65463
	KEY_KP_8                        = 65464
	KEY_KP_9                        = 65465
	KEY_KP_Add                      = 65451
	KEY_KP_Begin                    = 65437
	KEY_KP_Decimal                  = 65454
	KEY_KP_Delete                   = 65439
	KEY_KP_Divide                   = 65455
	KEY_KP_Down                     = 65433
	KEY_KP_End                      = 65436
	KEY_KP_Enter                    = 65421
	KEY_KP_Equal                    = 65469
	KEY_KP_F1                       = 65425
	KEY_KP_F2                       = 65426
	KEY_KP_F3                       = 65427
	KEY_KP_F4                       = 65428
	KEY_KP_Home                     = 65429
	KEY_KP_Insert                   = 65438
	KEY_KP_Left                     = 65430
	KEY_KP_Multiply                 = 65450
	KEY_KP_Next                     = 65435
	KEY_KP_Page_Down                = 65435
	KEY_KP_Page_Up                  = 65434
	KEY_KP_Prior                    = 65434
	KEY_KP_Right                    = 65432
	KEY_KP_Separator                = 65452
	KEY_KP_Space                    = 65408
	KEY_KP_Subtract                 = 65453
	KEY_KP_Tab                      = 65417
	KEY_KP_Up                       = 65431
	KEY_Kana_Lock                   = 65325
	KEY_Kana_Shift                  = 65326
	KEY_Kanji                       = 65313
	KEY_Kanji_Bangou                = 65335
	KEY_Katakana                    = 65318
	KEY_KbdBrightnessDown           = 269025030
	KEY_KbdBrightnessUp             = 269025029
	KEY_KbdLightOnOff               = 269025028
	KEY_Kcedilla                    = 979
	KEY_Korean_Won                  = 3839
	KEY_L                           = 76
	KEY_L1                          = 65480
	KEY_L10                         = 65489
	KEY_L2                          = 65481
	KEY_L3                          = 65482
	KEY_L4                          = 65483
	KEY_L5                          = 65484
	KEY_L6                          = 65485
	KEY_L7                          = 65486
	KEY_L8                          = 65487
	KEY_L9                          = 65488
	KEY_Lacute                      = 453
	KEY_Last_Virtual_Screen         = 65236
	KEY_Launch0                     = 269025088
	KEY_Launch1                     = 269025089
	KEY_Launch2                     = 269025090
	KEY_Launch3                     = 269025091
	KEY_Launch4                     = 269025092
	KEY_Launch5                     = 269025093
	KEY_Launch6                     = 269025094
	KEY_Launch7                     = 269025095
	KEY_Launch8                     = 269025096
	KEY_Launch9                     = 269025097
	KEY_LaunchA                     = 269025098
	KEY_LaunchB                     = 269025099
	KEY_LaunchC                     = 269025100
	KEY_LaunchD                     = 269025101
	KEY_LaunchE                     = 269025102
	KEY_LaunchF                     = 269025103
	KEY_Lbelowdot                   = 16784950
	KEY_Lcaron                      = 421
	KEY_Lcedilla                    = 934
	KEY_Left                        = 65361
	KEY_LightBulb                   = 269025077
	KEY_Linefeed                    = 65290
	KEY_LiraSign                    = 16785572
	KEY_LogGrabInfo                 = 269024805
	KEY_LogOff                      = 269025121
	KEY_LogWindowTree               = 269024804
	KEY_Lstroke                     = 419
	KEY_M                           = 77
	KEY_Mabovedot                   = 16784960
	KEY_Macedonia_DSE               = 1717
	KEY_Macedonia_GJE               = 1714
	KEY_Macedonia_KJE               = 1724
	KEY_Macedonia_dse               = 1701
	KEY_Macedonia_gje               = 1698
	KEY_Macedonia_kje               = 1708
	KEY_Mae_Koho                    = 65342
	KEY_Mail                        = 269025049
	KEY_MailForward                 = 269025168
	KEY_Market                      = 269025122
	KEY_Massyo                      = 65324
	KEY_Meeting                     = 269025123
	KEY_Memo                        = 269025054
	KEY_Menu                        = 65383
	KEY_MenuKB                      = 269025125
	KEY_MenuPB                      = 269025126
	KEY_Messenger                   = 269025166
	KEY_Meta_L                      = 65511
	KEY_Meta_R                      = 65512
	KEY_MillSign                    = 16785573
	KEY_ModeLock                    = 269025025
	KEY_Mode_switch                 = 65406
	KEY_MonBrightnessDown           = 269025027
	KEY_MonBrightnessUp             = 269025026
	KEY_MouseKeys_Accel_Enable      = 65143
	KEY_MouseKeys_Enable            = 65142
	KEY_Muhenkan                    = 65314
	KEY_Multi_key                   = 65312
	KEY_MultipleCandidate           = 65341
	KEY_Music                       = 269025170
	KEY_MyComputer                  = 269025075
	KEY_MySites                     = 269025127
	KEY_N                           = 78
	KEY_Nacute                      = 465
	KEY_NairaSign                   = 16785574
	KEY_Ncaron                      = 466
	KEY_Ncedilla                    = 977
	KEY_New                         = 269025128
	KEY_NewSheqelSign               = 16785578
	KEY_News                        = 269025129
	KEY_Next                        = 65366
	KEY_Next_VMode                  = 269024802
	KEY_Next_Virtual_Screen         = 65234
	KEY_Ntilde                      = 209
	KEY_Num_Lock                    = 65407
	KEY_O                           = 79
	KEY_OE                          = 5052
	KEY_Oacute                      = 211
	KEY_Obarred                     = 16777631
	KEY_Obelowdot                   = 16785100
	KEY_Ocaron                      = 16777681
	KEY_Ocircumflex                 = 212
	KEY_Ocircumflexacute            = 16785104
	KEY_Ocircumflexbelowdot         = 16785112
	KEY_Ocircumflexgrave            = 16785106
	KEY_Ocircumflexhook             = 16785108
	KEY_Ocircumflextilde            = 16785110
	KEY_Odiaeresis                  = 214
	KEY_Odoubleacute                = 469
	KEY_OfficeHome                  = 269025130
	KEY_Ograve                      = 210
	KEY_Ohook                       = 16785102
	KEY_Ohorn                       = 16777632
	KEY_Ohornacute                  = 16785114
	KEY_Ohornbelowdot               = 16785122
	KEY_Ohorngrave                  = 16785116
	KEY_Ohornhook                   = 16785118
	KEY_Ohorntilde                  = 16785120
	KEY_Omacron                     = 978
	KEY_Ooblique                    = 216
	KEY_Open                        = 269025131
	KEY_OpenURL                     = 269025080
	KEY_Option                      = 269025132
	KEY_Oslash                      = 216
	KEY_Otilde                      = 213
	KEY_Overlay1_Enable             = 65144
	KEY_Overlay2_Enable             = 65145
	KEY_P                           = 80
	KEY_Pabovedot                   = 16784982
	KEY_Page_Down                   = 65366
	KEY_Page_Up                     = 65365
	KEY_Paste                       = 269025133
	KEY_Pause                       = 65299
	KEY_PesetaSign                  = 16785575
	KEY_Phone                       = 269025134
	KEY_Pictures                    = 269025169
	KEY_Pointer_Accelerate          = 65274
	KEY_Pointer_Button1             = 65257
	KEY_Pointer_Button2             = 65258
	KEY_Pointer_Button3             = 65259
	KEY_Pointer_Button4             = 65260
	KEY_Pointer_Button5             = 65261
	KEY_Pointer_Button_Dflt         = 65256
	KEY_Pointer_DblClick1           = 65263
	KEY_Pointer_DblClick2           = 65264
	KEY_Pointer_DblClick3           = 65265
	KEY_Pointer_DblClick4           = 65266
	KEY_Pointer_DblClick5           = 65267
	KEY_Pointer_DblClick_Dflt       = 65262
	KEY_Pointer_DfltBtnNext         = 65275
	KEY_Pointer_DfltBtnPrev         = 65276
	KEY_Pointer_Down                = 65251
	KEY_Pointer_DownLeft            = 65254
	KEY_Pointer_DownRight           = 65255
	KEY_Pointer_Drag1               = 65269
	KEY_Pointer_Drag2               = 65270
	KEY_Pointer_Drag3               = 65271
	KEY_Pointer_Drag4               = 65272
	KEY_Pointer_Drag5               = 65277
	KEY_Pointer_Drag_Dflt           = 65268
	KEY_Pointer_EnableKeys          = 65273
	KEY_Pointer_Left                = 65248
	KEY_Pointer_Right               = 65249
	KEY_Pointer_Up                  = 65250
	KEY_Pointer_UpLeft              = 65252
	KEY_Pointer_UpRight             = 65253
	KEY_PowerDown                   = 269025057
	KEY_PowerOff                    = 269025066
	KEY_Prev_VMode                  = 269024803
	KEY_Prev_Virtual_Screen         = 65233
	KEY_PreviousCandidate           = 65342
	KEY_Print                       = 65377
	KEY_Prior                       = 65365
	KEY_Q                           = 81
	KEY_R                           = 82
	KEY_R1                          = 65490
	KEY_R10                         = 65499
	KEY_R11                         = 65500
	KEY_R12                         = 65501
	KEY_R13                         = 65502
	KEY_R14                         = 65503
	KEY_R15                         = 65504
	KEY_R2                          = 65491
	KEY_R3                          = 65492
	KEY_R4                          = 65493
	KEY_R5                          = 65494
	KEY_R6                          = 65495
	KEY_R7                          = 65496
	KEY_R8                          = 65497
	KEY_R9                          = 65498
	KEY_Racute                      = 448
	KEY_Rcaron                      = 472
	KEY_Rcedilla                    = 931
	KEY_Red                         = 269025187
	KEY_Redo                        = 65382
	KEY_Refresh                     = 269025065
	KEY_Reload                      = 269025139
	KEY_RepeatKeys_Enable           = 65138
	KEY_Reply                       = 269025138
	KEY_Return                      = 65293
	KEY_Right                       = 65363
	KEY_RockerDown                  = 269025060
	KEY_RockerEnter                 = 269025061
	KEY_RockerUp                    = 269025059
	KEY_Romaji                      = 65316
	KEY_RotateWindows               = 269025140
	KEY_RotationKB                  = 269025142
	KEY_RotationPB                  = 269025141
	KEY_RupeeSign                   = 16785576
	KEY_S                           = 83
	KEY_SCHWA                       = 16777615
	KEY_Sabovedot                   = 16784992
	KEY_Sacute                      = 422
	KEY_Save                        = 269025143
	KEY_Scaron                      = 425
	KEY_Scedilla                    = 426
	KEY_Scircumflex                 = 734
	KEY_ScreenSaver                 = 269025069
	KEY_ScrollClick                 = 269025146
	KEY_ScrollDown                  = 269025145
	KEY_ScrollUp                    = 269025144
	KEY_Scroll_Lock                 = 65300
	KEY_Search                      = 269025051
	KEY_Select                      = 65376
	KEY_SelectButton                = 269025184
	KEY_Send                        = 269025147
	KEY_Serbian_DJE                 = 1713
	KEY_Serbian_DZE                 = 1727
	KEY_Serbian_JE                  = 1720
	KEY_Serbian_LJE                 = 1721
	KEY_Serbian_NJE                 = 1722
	KEY_Serbian_TSHE                = 1723
	KEY_Serbian_dje                 = 1697
	KEY_Serbian_dze                 = 1711
	KEY_Serbian_je                  = 1704
	KEY_Serbian_lje                 = 1705
	KEY_Serbian_nje                 = 1706
	KEY_Serbian_tshe                = 1707
	KEY_Shift_L                     = 65505
	KEY_Shift_Lock                  = 65510
	KEY_Shift_R                     = 65506
	KEY_Shop                        = 269025078
	KEY_SingleCandidate             = 65340
	KEY_Sinh_a                      = 16780677
	KEY_Sinh_aa                     = 16780678
	KEY_Sinh_aa2                    = 16780751
	KEY_Sinh_ae                     = 16780679
	KEY_Sinh_ae2                    = 16780752
	KEY_Sinh_aee                    = 16780680
	KEY_Sinh_aee2                   = 16780753
	KEY_Sinh_ai                     = 16780691
	KEY_Sinh_ai2                    = 16780763
	KEY_Sinh_al                     = 16780746
	KEY_Sinh_au                     = 16780694
	KEY_Sinh_au2                    = 16780766
	KEY_Sinh_ba                     = 16780726
	KEY_Sinh_bha                    = 16780727
	KEY_Sinh_ca                     = 16780704
	KEY_Sinh_cha                    = 16780705
	KEY_Sinh_dda                    = 16780713
	KEY_Sinh_ddha                   = 16780714
	KEY_Sinh_dha                    = 16780719
	KEY_Sinh_dhha                   = 16780720
	KEY_Sinh_e                      = 16780689
	KEY_Sinh_e2                     = 16780761
	KEY_Sinh_ee                     = 16780690
	KEY_Sinh_ee2                    = 16780762
	KEY_Sinh_fa                     = 16780742
	KEY_Sinh_ga                     = 16780700
	KEY_Sinh_gha                    = 16780701
	KEY_Sinh_h2                     = 16780675
	KEY_Sinh_ha                     = 16780740
	KEY_Sinh_i                      = 16780681
	KEY_Sinh_i2                     = 16780754
	KEY_Sinh_ii                     = 16780682
	KEY_Sinh_ii2                    = 16780755
	KEY_Sinh_ja                     = 16780706
	KEY_Sinh_jha                    = 16780707
	KEY_Sinh_jnya                   = 16780709
	KEY_Sinh_ka                     = 16780698
	KEY_Sinh_kha                    = 16780699
	KEY_Sinh_kunddaliya             = 16780788
	KEY_Sinh_la                     = 16780733
	KEY_Sinh_lla                    = 16780741
	KEY_Sinh_lu                     = 16780687
	KEY_Sinh_lu2                    = 16780767
	KEY_Sinh_luu                    = 16780688
	KEY_Sinh_luu2                   = 16780787
	KEY_Sinh_ma                     = 16780728
	KEY_Sinh_mba                    = 16780729
	KEY_Sinh_na                     = 16780721
	KEY_Sinh_ndda                   = 16780716
	KEY_Sinh_ndha                   = 16780723
	KEY_Sinh_ng                     = 16780674
	KEY_Sinh_ng2                    = 16780702
	KEY_Sinh_nga                    = 16780703
	KEY_Sinh_nja                    = 16780710
	KEY_Sinh_nna                    = 16780715
	KEY_Sinh_nya                    = 16780708
	KEY_Sinh_o                      = 16780692
	KEY_Sinh_o2                     = 16780764
	KEY_Sinh_oo                     = 16780693
	KEY_Sinh_oo2                    = 16780765
	KEY_Sinh_pa                     = 16780724
	KEY_Sinh_pha                    = 16780725
	KEY_Sinh_ra                     = 16780731
	KEY_Sinh_ri                     = 16780685
	KEY_Sinh_rii                    = 16780686
	KEY_Sinh_ru2                    = 16780760
	KEY_Sinh_ruu2                   = 16780786
	KEY_Sinh_sa                     = 16780739
	KEY_Sinh_sha                    = 16780737
	KEY_Sinh_ssha                   = 16780738
	KEY_Sinh_tha                    = 16780717
	KEY_Sinh_thha                   = 16780718
	KEY_Sinh_tta                    = 16780711
	KEY_Sinh_ttha                   = 16780712
	KEY_Sinh_u                      = 16780683
	KEY_Sinh_u2                     = 16780756
	KEY_Sinh_uu                     = 16780684
	KEY_Sinh_uu2                    = 16780758
	KEY_Sinh_va                     = 16780736
	KEY_Sinh_ya                     = 16780730
	KEY_Sleep                       = 269025071
	KEY_SlowKeys_Enable             = 65139
	KEY_Spell                       = 269025148
	KEY_SplitScreen                 = 269025149
	KEY_Standby                     = 269025040
	KEY_Start                       = 269025050
	KEY_StickyKeys_Enable           = 65141
	KEY_Stop                        = 269025064
	KEY_Subtitle                    = 269025178
	KEY_Super_L                     = 65515
	KEY_Super_R                     = 65516
	KEY_Support                     = 269025150
	KEY_Suspend                     = 269025191
	KEY_Switch_VT_1                 = 269024769
	KEY_Switch_VT_10                = 269024778
	KEY_Switch_VT_11                = 269024779
	KEY_Switch_VT_12                = 269024780
	KEY_Switch_VT_2                 = 269024770
	KEY_Switch_VT_3                 = 269024771
	KEY_Switch_VT_4                 = 269024772
	KEY_Switch_VT_5                 = 269024773
	KEY_Switch_VT_6                 = 269024774
	KEY_Switch_VT_7                 = 269024775
	KEY_Switch_VT_8                 = 269024776
	KEY_Switch_VT_9                 = 269024777
	KEY_Sys_Req                     = 65301
	KEY_T                           = 84
	KEY_THORN                       = 222
	KEY_Tab                         = 65289
	KEY_Tabovedot                   = 16785002
	KEY_TaskPane                    = 269025151
	KEY_Tcaron                      = 427
	KEY_Tcedilla                    = 478
	KEY_Terminal                    = 269025152
	KEY_Terminate_Server            = 65237
	KEY_Thai_baht                   = 3551
	KEY_Thai_bobaimai               = 3514
	KEY_Thai_chochan                = 3496
	KEY_Thai_chochang               = 3498
	KEY_Thai_choching               = 3497
	KEY_Thai_chochoe                = 3500
	KEY_Thai_dochada                = 3502
	KEY_Thai_dodek                  = 3508
	KEY_Thai_fofa                   = 3517
	KEY_Thai_fofan                  = 3519
	KEY_Thai_hohip                  = 3531
	KEY_Thai_honokhuk               = 3534
	KEY_Thai_khokhai                = 3490
	KEY_Thai_khokhon                = 3493
	KEY_Thai_khokhuat               = 3491
	KEY_Thai_khokhwai               = 3492
	KEY_Thai_khorakhang             = 3494
	KEY_Thai_kokai                  = 3489
	KEY_Thai_lakkhangyao            = 3557
	KEY_Thai_lekchet                = 3575
	KEY_Thai_lekha                  = 3573
	KEY_Thai_lekhok                 = 3574
	KEY_Thai_lekkao                 = 3577
	KEY_Thai_leknung                = 3569
	KEY_Thai_lekpaet                = 3576
	KEY_Thai_leksam                 = 3571
	KEY_Thai_leksi                  = 3572
	KEY_Thai_leksong                = 3570
	KEY_Thai_leksun                 = 3568
	KEY_Thai_lochula                = 3532
	KEY_Thai_loling                 = 3525
	KEY_Thai_lu                     = 3526
	KEY_Thai_maichattawa            = 3563
	KEY_Thai_maiek                  = 3560
	KEY_Thai_maihanakat             = 3537
	KEY_Thai_maihanakat_maitho      = 3550
	KEY_Thai_maitaikhu              = 3559
	KEY_Thai_maitho                 = 3561
	KEY_Thai_maitri                 = 3562
	KEY_Thai_maiyamok               = 3558
	KEY_Thai_moma                   = 3521
	KEY_Thai_ngongu                 = 3495
	KEY_Thai_nikhahit               = 3565
	KEY_Thai_nonen                  = 3507
	KEY_Thai_nonu                   = 3513
	KEY_Thai_oang                   = 3533
	KEY_Thai_paiyannoi              = 3535
	KEY_Thai_phinthu                = 3546
	KEY_Thai_phophan                = 3518
	KEY_Thai_phophung               = 3516
	KEY_Thai_phosamphao             = 3520
	KEY_Thai_popla                  = 3515
	KEY_Thai_rorua                  = 3523
	KEY_Thai_ru                     = 3524
	KEY_Thai_saraa                  = 3536
	KEY_Thai_saraaa                 = 3538
	KEY_Thai_saraae                 = 3553
	KEY_Thai_saraaimaimalai         = 3556
	KEY_Thai_saraaimaimuan          = 3555
	KEY_Thai_saraam                 = 3539
	KEY_Thai_sarae                  = 3552
	KEY_Thai_sarai                  = 3540
	KEY_Thai_saraii                 = 3541
	KEY_Thai_sarao                  = 3554
	KEY_Thai_sarau                  = 3544
	KEY_Thai_saraue                 = 3542
	KEY_Thai_sarauee                = 3543
	KEY_Thai_sarauu                 = 3545
	KEY_Thai_sorusi                 = 3529
	KEY_Thai_sosala                 = 3528
	KEY_Thai_soso                   = 3499
	KEY_Thai_sosua                  = 3530
	KEY_Thai_thanthakhat            = 3564
	KEY_Thai_thonangmontho          = 3505
	KEY_Thai_thophuthao             = 3506
	KEY_Thai_thothahan              = 3511
	KEY_Thai_thothan                = 3504
	KEY_Thai_thothong               = 3512
	KEY_Thai_thothung               = 3510
	KEY_Thai_topatak                = 3503
	KEY_Thai_totao                  = 3509
	KEY_Thai_wowaen                 = 3527
	KEY_Thai_yoyak                  = 3522
	KEY_Thai_yoying                 = 3501
	KEY_Thorn                       = 222
	KEY_Time                        = 269025183
	KEY_ToDoList                    = 269025055
	KEY_Tools                       = 269025153
	KEY_TopMenu                     = 269025186
	KEY_TouchpadOff                 = 269025201
	KEY_TouchpadOn                  = 269025200
	KEY_TouchpadToggle              = 269025193
	KEY_Touroku                     = 65323
	KEY_Travel                      = 269025154
	KEY_Tslash                      = 940
	KEY_U                           = 85
	KEY_UWB                         = 269025174
	KEY_Uacute                      = 218
	KEY_Ubelowdot                   = 16785124
	KEY_Ubreve                      = 733
	KEY_Ucircumflex                 = 219
	KEY_Udiaeresis                  = 220
	KEY_Udoubleacute                = 475
	KEY_Ugrave                      = 217
	KEY_Uhook                       = 16785126
	KEY_Uhorn                       = 16777647
	KEY_Uhornacute                  = 16785128
	KEY_Uhornbelowdot               = 16785136
	KEY_Uhorngrave                  = 16785130
	KEY_Uhornhook                   = 16785132
	KEY_Uhorntilde                  = 16785134
	KEY_Ukrainian_GHE_WITH_UPTURN   = 1725
	KEY_Ukrainian_I                 = 1718
	KEY_Ukrainian_IE                = 1716
	KEY_Ukrainian_YI                = 1719
	KEY_Ukrainian_ghe_with_upturn   = 1709
	KEY_Ukrainian_i                 = 1702
	KEY_Ukrainian_ie                = 1700
	KEY_Ukrainian_yi                = 1703
	KEY_Ukranian_I                  = 1718
	KEY_Ukranian_JE                 = 1716
	KEY_Ukranian_YI                 = 1719
	KEY_Ukranian_i                  = 1702
	KEY_Ukranian_je                 = 1700
	KEY_Ukranian_yi                 = 1703
	KEY_Umacron                     = 990
	KEY_Undo                        = 65381
	KEY_Ungrab                      = 269024800
	KEY_Uogonek                     = 985
	KEY_Up                          = 65362
	KEY_Uring                       = 473
	KEY_User1KB                     = 269025157
	KEY_User2KB                     = 269025158
	KEY_UserPB                      = 269025156
	KEY_Utilde                      = 989
	KEY_V                           = 86
	KEY_VendorHome                  = 269025076
	KEY_Video                       = 269025159
	KEY_View                        = 269025185
	KEY_VoidSymbol                  = 16777215
	KEY_W                           = 87
	KEY_WLAN                        = 269025173
	KEY_WWW                         = 269025070
	KEY_Wacute                      = 16785026
	KEY_WakeUp                      = 269025067
	KEY_Wcircumflex                 = 16777588
	KEY_Wdiaeresis                  = 16785028
	KEY_WebCam                      = 269025167
	KEY_Wgrave                      = 16785024
	KEY_WheelButton                 = 269025160
	KEY_WindowClear                 = 269025109
	KEY_WonSign                     = 16785577
	KEY_Word                        = 269025161
	KEY_X                           = 88
	KEY_Xabovedot                   = 16785034
	KEY_Xfer                        = 269025162
	KEY_Y                           = 89
	KEY_Yacute                      = 221
	KEY_Ybelowdot                   = 16785140
	KEY_Ycircumflex                 = 16777590
	KEY_Ydiaeresis                  = 5054
	KEY_Yellow                      = 269025189
	KEY_Ygrave                      = 16785138
	KEY_Yhook                       = 16785142
	KEY_Ytilde                      = 16785144
	KEY_Z                           = 90
	KEY_Zabovedot                   = 431
	KEY_Zacute                      = 428
	KEY_Zcaron                      = 430
	KEY_Zen_Koho                    = 65341
	KEY_Zenkaku                     = 65320
	KEY_Zenkaku_Hankaku             = 65322
	KEY_ZoomIn                      = 269025163
	KEY_ZoomOut                     = 269025164
	KEY_Zstroke                     = 16777653
	KEY_a                           = 97
	KEY_aacute                      = 225
	KEY_abelowdot                   = 16785057
	KEY_abovedot                    = 511
	KEY_abreve                      = 483
	KEY_abreveacute                 = 16785071
	KEY_abrevebelowdot              = 16785079
	KEY_abrevegrave                 = 16785073
	KEY_abrevehook                  = 16785075
	KEY_abrevetilde                 = 16785077
	KEY_acircumflex                 = 226
	KEY_acircumflexacute            = 16785061
	KEY_acircumflexbelowdot         = 16785069
	KEY_acircumflexgrave            = 16785063
	KEY_acircumflexhook             = 16785065
	KEY_acircumflextilde            = 16785067
	KEY_acute                       = 180
	KEY_adiaeresis                  = 228
	KEY_ae                          = 230
	KEY_agrave                      = 224
	KEY_ahook                       = 16785059
	KEY_amacron                     = 992
	KEY_ampersand                   = 38
	KEY_aogonek                     = 433
	KEY_apostrophe                  = 39
	KEY_approxeq                    = 16785992
	KEY_approximate                 = 2248
	KEY_aring                       = 229
	KEY_asciicircum                 = 94
	KEY_asciitilde                  = 126
	KEY_asterisk                    = 42
	KEY_at                          = 64
	KEY_atilde                      = 227
	KEY_b                           = 98
	KEY_babovedot                   = 16784899
	KEY_backslash                   = 92
	KEY_ballotcross                 = 2804
	KEY_bar                         = 124
	KEY_because                     = 16785973
	KEY_blank                       = 2527
	KEY_botintegral                 = 2213
	KEY_botleftparens               = 2220
	KEY_botleftsqbracket            = 2216
	KEY_botleftsummation            = 2226
	KEY_botrightparens              = 2222
	KEY_botrightsqbracket           = 2218
	KEY_botrightsummation           = 2230
	KEY_bott                        = 2550
	KEY_botvertsummationconnector   = 2228
	KEY_braceleft                   = 123
	KEY_braceright                  = 125
	KEY_bracketleft                 = 91
	KEY_bracketright                = 93
	KEY_braille_blank               = 16787456
	KEY_braille_dot_1               = 65521
	KEY_braille_dot_10              = 65530
	KEY_braille_dot_2               = 65522
	KEY_braille_dot_3               = 65523
	KEY_braille_dot_4               = 65524
	KEY_braille_dot_5               = 65525
	KEY_braille_dot_6               = 65526
	KEY_braille_dot_7               = 65527
	KEY_braille_dot_8               = 65528
	KEY_braille_dot_9               = 65529
	KEY_braille_dots_1              = 16787457
	KEY_braille_dots_12             = 16787459
	KEY_braille_dots_123            = 16787463
	KEY_braille_dots_1234           = 16787471
	KEY_braille_dots_12345          = 16787487
	KEY_braille_dots_123456         = 16787519
	KEY_braille_dots_1234567        = 16787583
	KEY_braille_dots_12345678       = 16787711
	KEY_braille_dots_1234568        = 16787647
	KEY_braille_dots_123457         = 16787551
	KEY_braille_dots_1234578        = 16787679
	KEY_braille_dots_123458         = 16787615
	KEY_braille_dots_12346          = 16787503
	KEY_braille_dots_123467         = 16787567
	KEY_braille_dots_1234678        = 16787695
	KEY_braille_dots_123468         = 16787631
	KEY_braille_dots_12347          = 16787535
	KEY_braille_dots_123478         = 16787663
	KEY_braille_dots_12348          = 16787599
	KEY_braille_dots_1235           = 16787479
	KEY_braille_dots_12356          = 16787511
	KEY_braille_dots_123567         = 16787575
	KEY_braille_dots_1235678        = 16787703
	KEY_braille_dots_123568         = 16787639
	KEY_braille_dots_12357          = 16787543
	KEY_braille_dots_123578         = 16787671
	KEY_braille_dots_12358          = 16787607
	KEY_braille_dots_1236           = 16787495
	KEY_braille_dots_12367          = 16787559
	KEY_braille_dots_123678         = 16787687
	KEY_braille_dots_12368          = 16787623
	KEY_braille_dots_1237           = 16787527
	KEY_braille_dots_12378          = 16787655
	KEY_braille_dots_1238           = 16787591
	KEY_braille_dots_124            = 16787467
	KEY_braille_dots_1245           = 16787483
	KEY_braille_dots_12456          = 16787515
	KEY_braille_dots_124567         = 16787579
	KEY_braille_dots_1245678        = 16787707
	KEY_braille_dots_124568         = 16787643
	KEY_braille_dots_12457          = 16787547
	KEY_braille_dots_124578         = 16787675
	KEY_braille_dots_12458          = 16787611
	KEY_braille_dots_1246           = 16787499
	KEY_braille_dots_12467          = 16787563
	KEY_braille_dots_124678         = 16787691
	KEY_braille_dots_12468          = 16787627
	KEY_braille_dots_1247           = 16787531
	KEY_braille_dots_12478          = 16787659
	KEY_braille_dots_1248           = 16787595
	KEY_braille_dots_125            = 16787475
	KEY_braille_dots_1256           = 16787507
	KEY_braille_dots_12567          = 16787571
	KEY_braille_dots_125678         = 16787699
	KEY_braille_dots_12568          = 16787635
	KEY_braille_dots_1257           = 16787539
	KEY_braille_dots_12578          = 16787667
	KEY_braille_dots_1258           = 16787603
	KEY_braille_dots_126            = 16787491
	KEY_braille_dots_1267           = 16787555
	KEY_braille_dots_12678          = 16787683
	KEY_braille_dots_1268           = 16787619
	KEY_braille_dots_127            = 16787523
	KEY_braille_dots_1278           = 16787651
	KEY_braille_dots_128            = 16787587
	KEY_braille_dots_13             = 16787461
	KEY_braille_dots_134            = 16787469
	KEY_braille_dots_1345           = 16787485
	KEY_braille_dots_13456          = 16787517
	KEY_braille_dots_134567         = 16787581
	KEY_braille_dots_1345678        = 16787709
	KEY_braille_dots_134568         = 16787645
	KEY_braille_dots_13457          = 16787549
	KEY_braille_dots_134578         = 16787677
	KEY_braille_dots_13458          = 16787613
	KEY_braille_dots_1346           = 16787501
	KEY_braille_dots_13467          = 16787565
	KEY_braille_dots_134678         = 16787693
	KEY_braille_dots_13468          = 16787629
	KEY_braille_dots_1347           = 16787533
	KEY_braille_dots_13478          = 16787661
	KEY_braille_dots_1348           = 16787597
	KEY_braille_dots_135            = 16787477
	KEY_braille_dots_1356           = 16787509
	KEY_braille_dots_13567          = 16787573
	KEY_braille_dots_135678         = 16787701
	KEY_braille_dots_13568          = 16787637
	KEY_braille_dots_1357           = 16787541
	KEY_braille_dots_13578          = 16787669
	KEY_braille_dots_1358           = 16787605
	KEY_braille_dots_136            = 16787493
	KEY_braille_dots_1367           = 16787557
	KEY_braille_dots_13678          = 16787685
	KEY_braille_dots_1368           = 16787621
	KEY_braille_dots_137            = 16787525
	KEY_braille_dots_1378           = 16787653
	KEY_braille_dots_138            = 16787589
	KEY_braille_dots_14             = 16787465
	KEY_braille_dots_145            = 16787481
	KEY_braille_dots_1456           = 16787513
	KEY_braille_dots_14567          = 16787577
	KEY_braille_dots_145678         = 16787705
	KEY_braille_dots_14568          = 16787641
	KEY_braille_dots_1457           = 16787545
	KEY_braille_dots_14578          = 16787673
	KEY_braille_dots_1458           = 16787609
	KEY_braille_dots_146            = 16787497
	KEY_braille_dots_1467           = 16787561
	KEY_braille_dots_14678          = 16787689
	KEY_braille_dots_1468           = 16787625
	KEY_braille_dots_147            = 16787529
	KEY_braille_dots_1478           = 16787657
	KEY_braille_dots_148            = 16787593
	KEY_braille_dots_15             = 16787473
	KEY_braille_dots_156            = 16787505
	KEY_braille_dots_1567           = 16787569
	KEY_braille_dots_15678          = 16787697
	KEY_braille_dots_1568           = 16787633
	KEY_braille_dots_157            = 16787537
	KEY_braille_dots_1578           = 16787665
	KEY_braille_dots_158            = 16787601
	KEY_braille_dots_16             = 16787489
	KEY_braille_dots_167            = 16787553
	KEY_braille_dots_1678           = 16787681
	KEY_braille_dots_168            = 16787617
	KEY_braille_dots_17             = 16787521
	KEY_braille_dots_178            = 16787649
	KEY_braille_dots_18             = 16787585
	KEY_braille_dots_2              = 16787458
	KEY_braille_dots_23             = 16787462
	KEY_braille_dots_234            = 16787470
	KEY_braille_dots_2345           = 16787486
	KEY_braille_dots_23456          = 16787518
	KEY_braille_dots_234567         = 16787582
	KEY_braille_dots_2345678        = 16787710
	KEY_braille_dots_234568         = 16787646
	KEY_braille_dots_23457          = 16787550
	KEY_braille_dots_234578         = 16787678
	KEY_braille_dots_23458          = 16787614
	KEY_braille_dots_2346           = 16787502
	KEY_braille_dots_23467          = 16787566
	KEY_braille_dots_234678         = 16787694
	KEY_braille_dots_23468          = 16787630
	KEY_braille_dots_2347           = 16787534
	KEY_braille_dots_23478          = 16787662
	KEY_braille_dots_2348           = 16787598
	KEY_braille_dots_235            = 16787478
	KEY_braille_dots_2356           = 16787510
	KEY_braille_dots_23567          = 16787574
	KEY_braille_dots_235678         = 16787702
	KEY_braille_dots_23568          = 16787638
	KEY_braille_dots_2357           = 16787542
	KEY_braille_dots_23578          = 16787670
	KEY_braille_dots_2358           = 16787606
	KEY_braille_dots_236            = 16787494
	KEY_braille_dots_2367           = 16787558
	KEY_braille_dots_23678          = 16787686
	KEY_braille_dots_2368           = 16787622
	KEY_braille_dots_237            = 16787526
	KEY_braille_dots_2378           = 16787654
	KEY_braille_dots_238            = 16787590
	KEY_braille_dots_24             = 16787466
	KEY_braille_dots_245            = 16787482
	KEY_braille_dots_2456           = 16787514
	KEY_braille_dots_24567          = 16787578
	KEY_braille_dots_245678         = 16787706
	KEY_braille_dots_24568          = 16787642
	KEY_braille_dots_2457           = 16787546
	KEY_braille_dots_24578          = 16787674
	KEY_braille_dots_2458           = 16787610
	KEY_braille_dots_246            = 16787498
	KEY_braille_dots_2467           = 16787562
	KEY_braille_dots_24678          = 16787690
	KEY_braille_dots_2468           = 16787626
	KEY_braille_dots_247            = 16787530
	KEY_braille_dots_2478           = 16787658
	KEY_braille_dots_248            = 16787594
	KEY_braille_dots_25             = 16787474
	KEY_braille_dots_256            = 16787506
	KEY_braille_dots_2567           = 16787570
	KEY_braille_dots_25678          = 16787698
	KEY_braille_dots_2568           = 16787634
	KEY_braille_dots_257            = 16787538
	KEY_braille_dots_2578           = 16787666
	KEY_braille_dots_258            = 16787602
	KEY_braille_dots_26             = 16787490
	KEY_braille_dots_267            = 16787554
	KEY_braille_dots_2678           = 16787682
	KEY_braille_dots_268            = 16787618
	KEY_braille_dots_27             = 16787522
	KEY_braille_dots_278            = 16787650
	KEY_braille_dots_28             = 16787586
	KEY_braille_dots_3              = 16787460
	KEY_braille_dots_34             = 16787468
	KEY_braille_dots_345            = 16787484
	KEY_braille_dots_3456           = 16787516
	KEY_braille_dots_34567          = 16787580
	KEY_braille_dots_345678         = 16787708
	KEY_braille_dots_34568          = 16787644
	KEY_braille_dots_3457           = 16787548
	KEY_braille_dots_34578          = 16787676
	KEY_braille_dots_3458           = 16787612
	KEY_braille_dots_346            = 16787500
	KEY_braille_dots_3467           = 16787564
	KEY_braille_dots_34678          = 16787692
	KEY_braille_dots_3468           = 16787628
	KEY_braille_dots_347            = 16787532
	KEY_braille_dots_3478           = 16787660
	KEY_braille_dots_348            = 16787596
	KEY_braille_dots_35             = 16787476
	KEY_braille_dots_356            = 16787508
	KEY_braille_dots_3567           = 16787572
	KEY_braille_dots_35678          = 16787700
	KEY_braille_dots_3568           = 16787636
	KEY_braille_dots_357            = 16787540
	KEY_braille_dots_3578           = 16787668
	KEY_braille_dots_358            = 16787604
	KEY_braille_dots_36             = 16787492
	KEY_braille_dots_367            = 16787556
	KEY_braille_dots_3678           = 16787684
	KEY_braille_dots_368            = 16787620
	KEY_braille_dots_37             = 16787524
	KEY_braille_dots_378            = 16787652
	KEY_braille_dots_38             = 16787588
	KEY_braille_dots_4              = 16787464
	KEY_braille_dots_45             = 16787480
	KEY_braille_dots_456            = 16787512
	KEY_braille_dots_4567           = 16787576
	KEY_braille_dots_45678          = 16787704
	KEY_braille_dots_4568           = 16787640
	KEY_braille_dots_457            = 16787544
	KEY_braille_dots_4578           = 16787672
	KEY_braille_dots_458            = 16787608
	KEY_braille_dots_46             = 16787496
	KEY_braille_dots_467            = 16787560
	KEY_braille_dots_4678           = 16787688
	KEY_braille_dots_468            = 16787624
	KEY_braille_dots_47             = 16787528
	KEY_braille_dots_478            = 16787656
	KEY_braille_dots_48             = 16787592
	KEY_braille_dots_5              = 16787472
	KEY_braille_dots_56             = 16787504
	KEY_braille_dots_567            = 16787568
	KEY_braille_dots_5678           = 16787696
	KEY_braille_dots_568            = 16787632
	KEY_braille_dots_57             = 16787536
	KEY_braille_dots_578            = 16787664
	KEY_braille_dots_58             = 16787600
	KEY_braille_dots_6              = 16787488
	KEY_braille_dots_67             = 16787552
	KEY_braille_dots_678            = 16787680
	KEY_braille_dots_68             = 16787616
	KEY_braille_dots_7              = 16787520
	KEY_braille_dots_78             = 16787648
	KEY_braille_dots_8              = 16787584
	KEY_breve                       = 418
	KEY_brokenbar                   = 166
	KEY_c                           = 99
	KEY_c_h                         = 65187
	KEY_cabovedot                   = 741
	KEY_cacute                      = 486
	KEY_careof                      = 2744
	KEY_caret                       = 2812
	KEY_caron                       = 439
	KEY_ccaron                      = 488
	KEY_ccedilla                    = 231
	KEY_ccircumflex                 = 742
	KEY_cedilla                     = 184
	KEY_cent                        = 162
	KEY_ch                          = 65184
	KEY_checkerboard                = 2529
	KEY_checkmark                   = 2803
	KEY_circle                      = 3023
	KEY_club                        = 2796
	KEY_colon                       = 58
	KEY_comma                       = 44
	KEY_containsas                  = 16785931
	KEY_copyright                   = 169
	KEY_cr                          = 2532
	KEY_crossinglines               = 2542
	KEY_cuberoot                    = 16785947
	KEY_currency                    = 164
	KEY_cursor                      = 2815
	KEY_d                           = 100
	KEY_dabovedot                   = 16784907
	KEY_dagger                      = 2801
	KEY_dcaron                      = 495
	KEY_dead_A                      = 65153
	KEY_dead_E                      = 65155
	KEY_dead_I                      = 65157
	KEY_dead_O                      = 65159
	KEY_dead_U                      = 65161
	KEY_dead_a                      = 65152
	KEY_dead_abovecomma             = 65124
	KEY_dead_abovedot               = 65110
	KEY_dead_abovereversedcomma     = 65125
	KEY_dead_abovering              = 65112
	KEY_dead_acute                  = 65105
	KEY_dead_belowbreve             = 65131
	KEY_dead_belowcircumflex        = 65129
	KEY_dead_belowcomma             = 65134
	KEY_dead_belowdiaeresis         = 65132
	KEY_dead_belowdot               = 65120
	KEY_dead_belowmacron            = 65128
	KEY_dead_belowring              = 65127
	KEY_dead_belowtilde             = 65130
	KEY_dead_breve                  = 65109
	KEY_dead_capital_schwa          = 65163
	KEY_dead_caron                  = 65114
	KEY_dead_cedilla                = 65115
	KEY_dead_circumflex             = 65106
	KEY_dead_currency               = 65135
	KEY_dead_dasia                  = 65125
	KEY_dead_diaeresis              = 65111
	KEY_dead_doubleacute            = 65113
	KEY_dead_doublegrave            = 65126
	KEY_dead_e                      = 65154
	KEY_dead_grave                  = 65104
	KEY_dead_greek                  = 65164
	KEY_dead_hook                   = 65121
	KEY_dead_horn                   = 65122
	KEY_dead_i                      = 65156
	KEY_dead_invertedbreve          = 65133
	KEY_dead_iota                   = 65117
	KEY_dead_macron                 = 65108
	KEY_dead_o                      = 65158
	KEY_dead_ogonek                 = 65116
	KEY_dead_perispomeni            = 65107
	KEY_dead_psili                  = 65124
	KEY_dead_semivoiced_sound       = 65119
	KEY_dead_small_schwa            = 65162
	KEY_dead_stroke                 = 65123
	KEY_dead_tilde                  = 65107
	KEY_dead_u                      = 65160
	KEY_dead_voiced_sound           = 65118
	KEY_decimalpoint                = 2749
	KEY_degree                      = 176
	KEY_diaeresis                   = 168
	KEY_diamond                     = 2797
	KEY_digitspace                  = 2725
	KEY_dintegral                   = 16785964
	KEY_division                    = 247
	KEY_dollar                      = 36
	KEY_doubbaselinedot             = 2735
	KEY_doubleacute                 = 445
	KEY_doubledagger                = 2802
	KEY_doublelowquotemark          = 2814
	KEY_downarrow                   = 2302
	KEY_downcaret                   = 2984
	KEY_downshoe                    = 3030
	KEY_downstile                   = 3012
	KEY_downtack                    = 3010
	KEY_dstroke                     = 496
	KEY_e                           = 101
	KEY_eabovedot                   = 1004
	KEY_eacute                      = 233
	KEY_ebelowdot                   = 16785081
	KEY_ecaron                      = 492
	KEY_ecircumflex                 = 234
	KEY_ecircumflexacute            = 16785087
	KEY_ecircumflexbelowdot         = 16785095
	KEY_ecircumflexgrave            = 16785089
	KEY_ecircumflexhook             = 16785091
	KEY_ecircumflextilde            = 16785093
	KEY_ediaeresis                  = 235
	KEY_egrave                      = 232
	KEY_ehook                       = 16785083
	KEY_eightsubscript              = 16785544
	KEY_eightsuperior               = 16785528
	KEY_elementof                   = 16785928
	KEY_ellipsis                    = 2734
	KEY_em3space                    = 2723
	KEY_em4space                    = 2724
	KEY_emacron                     = 954
	KEY_emdash                      = 2729
	KEY_emfilledcircle              = 2782
	KEY_emfilledrect                = 2783
	KEY_emopencircle                = 2766
	KEY_emopenrectangle             = 2767
	KEY_emptyset                    = 16785925
	KEY_emspace                     = 2721
	KEY_endash                      = 2730
	KEY_enfilledcircbullet          = 2790
	KEY_enfilledsqbullet            = 2791
	KEY_eng                         = 959
	KEY_enopencircbullet            = 2784
	KEY_enopensquarebullet          = 2785
	KEY_enspace                     = 2722
	KEY_eogonek                     = 490
	KEY_equal                       = 61
	KEY_eth                         = 240
	KEY_etilde                      = 16785085
	KEY_exclam                      = 33
	KEY_exclamdown                  = 161
	KEY_ezh                         = 16777874
	KEY_f                           = 102
	KEY_fabovedot                   = 16784927
	KEY_femalesymbol                = 2808
	KEY_ff                          = 2531
	KEY_figdash                     = 2747
	KEY_filledlefttribullet         = 2780
	KEY_filledrectbullet            = 2779
	KEY_filledrighttribullet        = 2781
	KEY_filledtribulletdown         = 2793
	KEY_filledtribulletup           = 2792
	KEY_fiveeighths                 = 2757
	KEY_fivesixths                  = 2743
	KEY_fivesubscript               = 16785541
	KEY_fivesuperior                = 16785525
	KEY_fourfifths                  = 2741
	KEY_foursubscript               = 16785540
	KEY_foursuperior                = 16785524
	KEY_fourthroot                  = 16785948
	KEY_function                    = 2294
	KEY_g                           = 103
	KEY_gabovedot                   = 757
	KEY_gbreve                      = 699
	KEY_gcaron                      = 16777703
	KEY_gcedilla                    = 955
	KEY_gcircumflex                 = 760
	KEY_grave                       = 96
	KEY_greater                     = 62
	KEY_greaterthanequal            = 2238
	KEY_guillemotleft               = 171
	KEY_guillemotright              = 187
	KEY_h                           = 104
	KEY_hairspace                   = 2728
	KEY_hcircumflex                 = 694
	KEY_heart                       = 2798
	KEY_hebrew_aleph                = 3296
	KEY_hebrew_ayin                 = 3314
	KEY_hebrew_bet                  = 3297
	KEY_hebrew_beth                 = 3297
	KEY_hebrew_chet                 = 3303
	KEY_hebrew_dalet                = 3299
	KEY_hebrew_daleth               = 3299
	KEY_hebrew_doublelowline        = 3295
	KEY_hebrew_finalkaph            = 3306
	KEY_hebrew_finalmem             = 3309
	KEY_hebrew_finalnun             = 3311
	KEY_hebrew_finalpe              = 3315
	KEY_hebrew_finalzade            = 3317
	KEY_hebrew_finalzadi            = 3317
	KEY_hebrew_gimel                = 3298
	KEY_hebrew_gimmel               = 3298
	KEY_hebrew_he                   = 3300
	KEY_hebrew_het                  = 3303
	KEY_hebrew_kaph                 = 3307
	KEY_hebrew_kuf                  = 3319
	KEY_hebrew_lamed                = 3308
	KEY_hebrew_mem                  = 3310
	KEY_hebrew_nun                  = 3312
	KEY_hebrew_pe                   = 3316
	KEY_hebrew_qoph                 = 3319
	KEY_hebrew_resh                 = 3320
	KEY_hebrew_samech               = 3313
	KEY_hebrew_samekh               = 3313
	KEY_hebrew_shin                 = 3321
	KEY_hebrew_taf                  = 3322
	KEY_hebrew_taw                  = 3322
	KEY_hebrew_tet                  = 3304
	KEY_hebrew_teth                 = 3304
	KEY_hebrew_waw                  = 3301
	KEY_hebrew_yod                  = 3305
	KEY_hebrew_zade                 = 3318
	KEY_hebrew_zadi                 = 3318
	KEY_hebrew_zain                 = 3302
	KEY_hebrew_zayin                = 3302
	KEY_hexagram                    = 2778
	KEY_horizconnector              = 2211
	KEY_horizlinescan1              = 2543
	KEY_horizlinescan3              = 2544
	KEY_horizlinescan5              = 2545
	KEY_horizlinescan7              = 2546
	KEY_horizlinescan9              = 2547
	KEY_hstroke                     = 689
	KEY_ht                          = 2530
	KEY_hyphen                      = 173
	KEY_i                           = 105
	KEY_iTouch                      = 269025120
	KEY_iacute                      = 237
	KEY_ibelowdot                   = 16785099
	KEY_ibreve                      = 16777517
	KEY_icircumflex                 = 238
	KEY_identical                   = 2255
	KEY_idiaeresis                  = 239
	KEY_idotless                    = 697
	KEY_ifonlyif                    = 2253
	KEY_igrave                      = 236
	KEY_ihook                       = 16785097
	KEY_imacron                     = 1007
	KEY_implies                     = 2254
	KEY_includedin                  = 2266
	KEY_includes                    = 2267
	KEY_infinity                    = 2242
	KEY_integral                    = 2239
	KEY_intersection                = 2268
	KEY_iogonek                     = 999
	KEY_itilde                      = 949
	KEY_j                           = 106
	KEY_jcircumflex                 = 700
	KEY_jot                         = 3018
	KEY_k                           = 107
	KEY_kana_A                      = 1201
	KEY_kana_CHI                    = 1217
	KEY_kana_E                      = 1204
	KEY_kana_FU                     = 1228
	KEY_kana_HA                     = 1226
	KEY_kana_HE                     = 1229
	KEY_kana_HI                     = 1227
	KEY_kana_HO                     = 1230
	KEY_kana_HU                     = 1228
	KEY_kana_I                      = 1202
	KEY_kana_KA                     = 1206
	KEY_kana_KE                     = 1209
	KEY_kana_KI                     = 1207
	KEY_kana_KO                     = 1210
	KEY_kana_KU                     = 1208
	KEY_kana_MA                     = 1231
	KEY_kana_ME                     = 1234
	KEY_kana_MI                     = 1232
	KEY_kana_MO                     = 1235
	KEY_kana_MU                     = 1233
	KEY_kana_N                      = 1245
	KEY_kana_NA                     = 1221
	KEY_kana_NE                     = 1224
	KEY_kana_NI                     = 1222
	KEY_kana_NO                     = 1225
	KEY_kana_NU                     = 1223
	KEY_kana_O                      = 1205
	KEY_kana_RA                     = 1239
	KEY_kana_RE                     = 1242
	KEY_kana_RI                     = 1240
	KEY_kana_RO                     = 1243
	KEY_kana_RU                     = 1241
	KEY_kana_SA                     = 1211
	KEY_kana_SE                     = 1214
	KEY_kana_SHI                    = 1212
	KEY_kana_SO                     = 1215
	KEY_kana_SU                     = 1213
	KEY_kana_TA                     = 1216
	KEY_kana_TE                     = 1219
	KEY_kana_TI                     = 1217
	KEY_kana_TO                     = 1220
	KEY_kana_TSU                    = 1218
	KEY_kana_TU                     = 1218
	KEY_kana_U                      = 1203
	KEY_kana_WA                     = 1244
	KEY_kana_WO                     = 1190
	KEY_kana_YA                     = 1236
	KEY_kana_YO                     = 1238
	KEY_kana_YU                     = 1237
	KEY_kana_a                      = 1191
	KEY_kana_closingbracket         = 1187
	KEY_kana_comma                  = 1188
	KEY_kana_conjunctive            = 1189
	KEY_kana_e                      = 1194
	KEY_kana_fullstop               = 1185
	KEY_kana_i                      = 1192
	KEY_kana_middledot              = 1189
	KEY_kana_o                      = 1195
	KEY_kana_openingbracket         = 1186
	KEY_kana_switch                 = 65406
	KEY_kana_tsu                    = 1199
	KEY_kana_tu                     = 1199
	KEY_kana_u                      = 1193
	KEY_kana_ya                     = 1196
	KEY_kana_yo                     = 1198
	KEY_kana_yu                     = 1197
	KEY_kappa                       = 930
	KEY_kcedilla                    = 1011
	KEY_kra                         = 930
	KEY_l                           = 108
	KEY_lacute                      = 485
	KEY_latincross                  = 2777
	KEY_lbelowdot                   = 16784951
	KEY_lcaron                      = 437
	KEY_lcedilla                    = 950
	KEY_leftanglebracket            = 2748
	KEY_leftarrow                   = 2299
	KEY_leftcaret                   = 2979
	KEY_leftdoublequotemark         = 2770
	KEY_leftmiddlecurlybrace        = 2223
	KEY_leftopentriangle            = 2764
	KEY_leftpointer                 = 2794
	KEY_leftradical                 = 2209
	KEY_leftshoe                    = 3034
	KEY_leftsinglequotemark         = 2768
	KEY_leftt                       = 2548
	KEY_lefttack                    = 3036
	KEY_less                        = 60
	KEY_lessthanequal               = 2236
	KEY_lf                          = 2533
	KEY_logicaland                  = 2270
	KEY_logicalor                   = 2271
	KEY_lowleftcorner               = 2541
	KEY_lowrightcorner              = 2538
	KEY_lstroke                     = 435
	KEY_m                           = 109
	KEY_mabovedot                   = 16784961
	KEY_macron                      = 175
	KEY_malesymbol                  = 2807
	KEY_maltesecross                = 2800
	KEY_marker                      = 2751
	KEY_masculine                   = 186
	KEY_minus                       = 45
	KEY_minutes                     = 2774
	KEY_mu                          = 181
	KEY_multiply                    = 215
	KEY_musicalflat                 = 2806
	KEY_musicalsharp                = 2805
	KEY_n                           = 110
	KEY_nabla                       = 2245
	KEY_nacute                      = 497
	KEY_ncaron                      = 498
	KEY_ncedilla                    = 1009
	KEY_ninesubscript               = 16785545
	KEY_ninesuperior                = 16785529
	KEY_nl                          = 2536
	KEY_nobreakspace                = 160
	KEY_notapproxeq                 = 16785991
	KEY_notelementof                = 16785929
	KEY_notequal                    = 2237
	KEY_notidentical                = 16786018
	KEY_notsign                     = 172
	KEY_ntilde                      = 241
	KEY_numbersign                  = 35
	KEY_numerosign                  = 1712
	KEY_o                           = 111
	KEY_oacute                      = 243
	KEY_obarred                     = 16777845
	KEY_obelowdot                   = 16785101
	KEY_ocaron                      = 16777682
	KEY_ocircumflex                 = 244
	KEY_ocircumflexacute            = 16785105
	KEY_ocircumflexbelowdot         = 16785113
	KEY_ocircumflexgrave            = 16785107
	KEY_ocircumflexhook             = 16785109
	KEY_ocircumflextilde            = 16785111
	KEY_odiaeresis                  = 246
	KEY_odoubleacute                = 501
	KEY_oe                          = 5053
	KEY_ogonek                      = 434
	KEY_ograve                      = 242
	KEY_ohook                       = 16785103
	KEY_ohorn                       = 16777633
	KEY_ohornacute                  = 16785115
	KEY_ohornbelowdot               = 16785123
	KEY_ohorngrave                  = 16785117
	KEY_ohornhook                   = 16785119
	KEY_ohorntilde                  = 16785121
	KEY_omacron                     = 1010
	KEY_oneeighth                   = 2755
	KEY_onefifth                    = 2738
	KEY_onehalf                     = 189
	KEY_onequarter                  = 188
	KEY_onesixth                    = 2742
	KEY_onesubscript                = 16785537
	KEY_onesuperior                 = 185
	KEY_onethird                    = 2736
	KEY_ooblique                    = 248
	KEY_openrectbullet              = 2786
	KEY_openstar                    = 2789
	KEY_opentribulletdown           = 2788
	KEY_opentribulletup             = 2787
	KEY_ordfeminine                 = 170
	KEY_oslash                      = 248
	KEY_otilde                      = 245
	KEY_overbar                     = 3008
	KEY_overline                    = 1150
	KEY_p                           = 112
	KEY_pabovedot                   = 16784983
	KEY_paragraph                   = 182
	KEY_parenleft                   = 40
	KEY_parenright                  = 41
	KEY_partdifferential            = 16785922
	KEY_partialderivative           = 2287
	KEY_percent                     = 37
	KEY_period                      = 46
	KEY_periodcentered              = 183
	KEY_permille                    = 2773
	KEY_phonographcopyright         = 2811
	KEY_plus                        = 43
	KEY_plusminus                   = 177
	KEY_prescription                = 2772
	KEY_prolongedsound              = 1200
	KEY_punctspace                  = 2726
	KEY_q                           = 113
	KEY_quad                        = 3020
	KEY_question                    = 63
	KEY_questiondown                = 191
	KEY_quotedbl                    = 34
	KEY_quoteleft                   = 96
	KEY_quoteright                  = 39
	KEY_r                           = 114
	KEY_racute                      = 480
	KEY_radical                     = 2262
	KEY_rcaron                      = 504
	KEY_rcedilla                    = 947
	KEY_registered                  = 174
	KEY_rightanglebracket           = 2750
	KEY_rightarrow                  = 2301
	KEY_rightcaret                  = 2982
	KEY_rightdoublequotemark        = 2771
	KEY_rightmiddlecurlybrace       = 2224
	KEY_rightmiddlesummation        = 2231
	KEY_rightopentriangle           = 2765
	KEY_rightpointer                = 2795
	KEY_rightshoe                   = 3032
	KEY_rightsinglequotemark        = 2769
	KEY_rightt                      = 2549
	KEY_righttack                   = 3068
	KEY_s                           = 115
	KEY_sabovedot                   = 16784993
	KEY_sacute                      = 438
	KEY_scaron                      = 441
	KEY_scedilla                    = 442
	KEY_schwa                       = 16777817
	KEY_scircumflex                 = 766
	KEY_script_switch               = 65406
	KEY_seconds                     = 2775
	KEY_section                     = 167
	KEY_semicolon                   = 59
	KEY_semivoicedsound             = 1247
	KEY_seveneighths                = 2758
	KEY_sevensubscript              = 16785543
	KEY_sevensuperior               = 16785527
	KEY_signaturemark               = 2762
	KEY_signifblank                 = 2732
	KEY_similarequal                = 2249
	KEY_singlelowquotemark          = 2813
	KEY_sixsubscript                = 16785542
	KEY_sixsuperior                 = 16785526
	KEY_slash                       = 47
	KEY_soliddiamond                = 2528
	KEY_space                       = 32
	KEY_squareroot                  = 16785946
	KEY_ssharp                      = 223
	KEY_sterling                    = 163
	KEY_stricteq                    = 16786019
	KEY_t                           = 116
	KEY_tabovedot                   = 16785003
	KEY_tcaron                      = 443
	KEY_tcedilla                    = 510
	KEY_telephone                   = 2809
	KEY_telephonerecorder           = 2810
	KEY_therefore                   = 2240
	KEY_thinspace                   = 2727
	KEY_thorn                       = 254
	KEY_threeeighths                = 2756
	KEY_threefifths                 = 2740
	KEY_threequarters               = 190
	KEY_threesubscript              = 16785539
	KEY_threesuperior               = 179
	KEY_tintegral                   = 16785965
	KEY_topintegral                 = 2212
	KEY_topleftparens               = 2219
	KEY_topleftradical              = 2210
	KEY_topleftsqbracket            = 2215
	KEY_topleftsummation            = 2225
	KEY_toprightparens              = 2221
	KEY_toprightsqbracket           = 2217
	KEY_toprightsummation           = 2229
	KEY_topt                        = 2551
	KEY_topvertsummationconnector   = 2227
	KEY_trademark                   = 2761
	KEY_trademarkincircle           = 2763
	KEY_tslash                      = 956
	KEY_twofifths                   = 2739
	KEY_twosubscript                = 16785538
	KEY_twosuperior                 = 178
	KEY_twothirds                   = 2737
	KEY_u                           = 117
	KEY_uacute                      = 250
	KEY_ubelowdot                   = 16785125
	KEY_ubreve                      = 765
	KEY_ucircumflex                 = 251
	KEY_udiaeresis                  = 252
	KEY_udoubleacute                = 507
	KEY_ugrave                      = 249
	KEY_uhook                       = 16785127
	KEY_uhorn                       = 16777648
	KEY_uhornacute                  = 16785129
	KEY_uhornbelowdot               = 16785137
	KEY_uhorngrave                  = 16785131
	KEY_uhornhook                   = 16785133
	KEY_uhorntilde                  = 16785135
	KEY_umacron                     = 1022
	KEY_underbar                    = 3014
	KEY_underscore                  = 95
	KEY_union                       = 2269
	KEY_uogonek                     = 1017
	KEY_uparrow                     = 2300
	KEY_upcaret                     = 2985
	KEY_upleftcorner                = 2540
	KEY_uprightcorner               = 2539
	KEY_upshoe                      = 3011
	KEY_upstile                     = 3027
	KEY_uptack                      = 3022
	KEY_uring                       = 505
	KEY_utilde                      = 1021
	KEY_v                           = 118
	KEY_variation                   = 2241
	KEY_vertbar                     = 2552
	KEY_vertconnector               = 2214
	KEY_voicedsound                 = 1246
	KEY_vt                          = 2537
	KEY_w                           = 119
	KEY_wacute                      = 16785027
	KEY_wcircumflex                 = 16777589
	KEY_wdiaeresis                  = 16785029
	KEY_wgrave                      = 16785025
	KEY_x                           = 120
	KEY_xabovedot                   = 16785035
	KEY_y                           = 121
	KEY_yacute                      = 253
	KEY_ybelowdot                   = 16785141
	KEY_ycircumflex                 = 16777591
	KEY_ydiaeresis                  = 255
	KEY_yen                         = 165
	KEY_ygrave                      = 16785139
	KEY_yhook                       = 16785143
	KEY_ytilde                      = 16785145
	KEY_z                           = 122
	KEY_zabovedot                   = 447
	KEY_zacute                      = 444
	KEY_zcaron                      = 446
	KEY_zerosubscript               = 16785536
	KEY_zerosuperior                = 16785520
	KEY_zstroke                     = 16777654
	MAX_TIMECOORD_AXES              = 128
	PARENT_RELATIVE                 = 1
	PRIORITY_REDRAW                 = 120
)
