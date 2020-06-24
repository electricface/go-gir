package wnck

/*
#cgo pkg-config: libwnck-3.0
#include <libwnck/libwnck.h>
*/
import "C"
import "github.com/electricface/go-gir/atk-1.0"
import "github.com/electricface/go-gir/g-2.0"
import "github.com/electricface/go-gir/gdk-3.0"
import "github.com/electricface/go-gir/gdkpixbuf-2.0"
import "github.com/electricface/go-gir/gtk-3.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("Wnck")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("Wnck", "3.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Object ActionMenu
type ActionMenu struct {
	atk.ImplementorIfaceIfc
	gtk.BuildableIfc
	gtk.Menu
}

func WrapActionMenu(p unsafe.Pointer) (r ActionMenu) { r.P = p; return }

type IActionMenu interface{ P_ActionMenu() unsafe.Pointer }

func (v ActionMenu) P_ActionMenu() unsafe.Pointer       { return v.P }
func (v ActionMenu) P_ImplementorIface() unsafe.Pointer { return v.P }
func (v ActionMenu) P_Buildable() unsafe.Pointer        { return v.P }
func ActionMenuGetType() gi.GType {
	ret := _I.GetGType(0, "ActionMenu")
	return ret
}

// wnck_action_menu_new
//
// [ window ] trans: nothing
//
// [ result ] trans: nothing
//
func NewActionMenu(window IWindow) (result ActionMenu) {
	iv, err := _I.Get(0, "ActionMenu", "new")
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

// ignore GType struct ActionMenuClass

// Struct ActionMenuPrivate
type ActionMenuPrivate struct {
	P unsafe.Pointer
}

func ActionMenuPrivateGetType() gi.GType {
	ret := _I.GetGType(1, "ActionMenuPrivate")
	return ret
}

// Object Application
type Application struct {
	g.Object
}

func WrapApplication(p unsafe.Pointer) (r Application) { r.P = p; return }

type IApplication interface{ P_Application() unsafe.Pointer }

func (v Application) P_Application() unsafe.Pointer { return v.P }
func ApplicationGetType() gi.GType {
	ret := _I.GetGType(2, "Application")
	return ret
}

// wnck_application_get
//
// [ xwindow ] trans: nothing
//
// [ result ] trans: nothing
//
func ApplicationGet1(xwindow uint64) (result Application) {
	iv, err := _I.Get(1, "Application", "get")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_xwindow := gi.NewUint64Argument(xwindow)
	args := []gi.Argument{arg_xwindow}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// wnck_application_get_icon
//
// [ result ] trans: nothing
//
func (v Application) GetIcon() (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(2, "Application", "get_icon")
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

// wnck_application_get_icon_is_fallback
//
// [ result ] trans: nothing
//
func (v Application) GetIconIsFallback() (result bool) {
	iv, err := _I.Get(3, "Application", "get_icon_is_fallback")
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

// wnck_application_get_icon_name
//
// [ result ] trans: nothing
//
func (v Application) GetIconName() (result string) {
	iv, err := _I.Get(4, "Application", "get_icon_name")
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

// wnck_application_get_mini_icon
//
// [ result ] trans: nothing
//
func (v Application) GetMiniIcon() (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(5, "Application", "get_mini_icon")
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

// wnck_application_get_n_windows
//
// [ result ] trans: nothing
//
func (v Application) GetNWindows() (result int32) {
	iv, err := _I.Get(6, "Application", "get_n_windows")
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

// wnck_application_get_name
//
// [ result ] trans: nothing
//
func (v Application) GetName() (result string) {
	iv, err := _I.Get(7, "Application", "get_name")
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

// wnck_application_get_pid
//
// [ result ] trans: nothing
//
func (v Application) GetPid() (result int32) {
	iv, err := _I.Get(8, "Application", "get_pid")
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

// wnck_application_get_startup_id
//
// [ result ] trans: nothing
//
func (v Application) GetStartupId() (result string) {
	iv, err := _I.Get(9, "Application", "get_startup_id")
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

// wnck_application_get_windows
//
// [ result ] trans: nothing
//
func (v Application) GetWindows() (result g.List) {
	iv, err := _I.Get(10, "Application", "get_windows")
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

// wnck_application_get_xid
//
// [ result ] trans: nothing
//
func (v Application) GetXid() (result uint64) {
	iv, err := _I.Get(11, "Application", "get_xid")
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

// ignore GType struct ApplicationClass

// Struct ApplicationPrivate
type ApplicationPrivate struct {
	P unsafe.Pointer
}

func ApplicationPrivateGetType() gi.GType {
	ret := _I.GetGType(3, "ApplicationPrivate")
	return ret
}

// Object ClassGroup
type ClassGroup struct {
	g.Object
}

func WrapClassGroup(p unsafe.Pointer) (r ClassGroup) { r.P = p; return }

type IClassGroup interface{ P_ClassGroup() unsafe.Pointer }

func (v ClassGroup) P_ClassGroup() unsafe.Pointer { return v.P }
func ClassGroupGetType() gi.GType {
	ret := _I.GetGType(4, "ClassGroup")
	return ret
}

// wnck_class_group_get
//
// [ id ] trans: nothing
//
// [ result ] trans: nothing
//
func ClassGroupGet1(id string) (result ClassGroup) {
	iv, err := _I.Get(12, "ClassGroup", "get")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_id := gi.CString(id)
	arg_id := gi.NewStringArgument(c_id)
	args := []gi.Argument{arg_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_id)
	result.P = ret.Pointer()
	return
}

// wnck_class_group_get_icon
//
// [ result ] trans: nothing
//
func (v ClassGroup) GetIcon() (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(13, "ClassGroup", "get_icon")
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

// wnck_class_group_get_id
//
// [ result ] trans: nothing
//
func (v ClassGroup) GetId() (result string) {
	iv, err := _I.Get(14, "ClassGroup", "get_id")
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

// wnck_class_group_get_mini_icon
//
// [ result ] trans: nothing
//
func (v ClassGroup) GetMiniIcon() (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(15, "ClassGroup", "get_mini_icon")
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

// wnck_class_group_get_name
//
// [ result ] trans: nothing
//
func (v ClassGroup) GetName() (result string) {
	iv, err := _I.Get(16, "ClassGroup", "get_name")
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
// wnck_class_group_get_res_class
//
// [ result ] trans: nothing
//
func (v ClassGroup) GetResClass() (result string) {
	iv, err := _I.Get(17, "ClassGroup", "get_res_class")
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

// wnck_class_group_get_windows
//
// [ result ] trans: nothing
//
func (v ClassGroup) GetWindows() (result g.List) {
	iv, err := _I.Get(18, "ClassGroup", "get_windows")
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

// ignore GType struct ClassGroupClass

// Struct ClassGroupPrivate
type ClassGroupPrivate struct {
	P unsafe.Pointer
}

func ClassGroupPrivateGetType() gi.GType {
	ret := _I.GetGType(5, "ClassGroupPrivate")
	return ret
}

// Enum ClientType
type ClientTypeEnum int

const (
	ClientTypeApplication ClientTypeEnum = 1
	ClientTypePager       ClientTypeEnum = 2
)

func ClientTypeGetType() gi.GType {
	ret := _I.GetGType(6, "ClientType")
	return ret
}

// Enum MotionDirection
type MotionDirectionEnum int

const (
	MotionDirectionUp    MotionDirectionEnum = -1
	MotionDirectionDown  MotionDirectionEnum = -2
	MotionDirectionLeft  MotionDirectionEnum = -3
	MotionDirectionRight MotionDirectionEnum = -4
)

func MotionDirectionGetType() gi.GType {
	ret := _I.GetGType(7, "MotionDirection")
	return ret
}

// Object Pager
type Pager struct {
	atk.ImplementorIfaceIfc
	gtk.BuildableIfc
	gtk.Widget
}

func WrapPager(p unsafe.Pointer) (r Pager) { r.P = p; return }

type IPager interface{ P_Pager() unsafe.Pointer }

func (v Pager) P_Pager() unsafe.Pointer            { return v.P }
func (v Pager) P_ImplementorIface() unsafe.Pointer { return v.P }
func (v Pager) P_Buildable() unsafe.Pointer        { return v.P }
func PagerGetType() gi.GType {
	ret := _I.GetGType(8, "Pager")
	return ret
}

// wnck_pager_new
//
// [ result ] trans: nothing
//
func NewPager() (result Pager) {
	iv, err := _I.Get(19, "Pager", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// wnck_pager_get_wrap_on_scroll
//
// [ result ] trans: nothing
//
func (v Pager) GetWrapOnScroll() (result bool) {
	iv, err := _I.Get(20, "Pager", "get_wrap_on_scroll")
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

// wnck_pager_set_display_mode
//
// [ mode ] trans: nothing
//
func (v Pager) SetDisplayMode(mode PagerDisplayModeEnum) {
	iv, err := _I.Get(21, "Pager", "set_display_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_v, arg_mode}
	iv.Call(args, nil, nil)
}

// wnck_pager_set_n_rows
//
// [ n_rows ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pager) SetNRows(n_rows int32) (result bool) {
	iv, err := _I.Get(22, "Pager", "set_n_rows")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_rows := gi.NewInt32Argument(n_rows)
	args := []gi.Argument{arg_v, arg_n_rows}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// wnck_pager_set_orientation
//
// [ orientation ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Pager) SetOrientation(orientation gtk.OrientationEnum) (result bool) {
	iv, err := _I.Get(23, "Pager", "set_orientation")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_orientation := gi.NewIntArgument(int(orientation))
	args := []gi.Argument{arg_v, arg_orientation}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// wnck_pager_set_shadow_type
//
// [ shadow_type ] trans: nothing
//
func (v Pager) SetShadowType(shadow_type gtk.ShadowTypeEnum) {
	iv, err := _I.Get(24, "Pager", "set_shadow_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_shadow_type := gi.NewIntArgument(int(shadow_type))
	args := []gi.Argument{arg_v, arg_shadow_type}
	iv.Call(args, nil, nil)
}

// wnck_pager_set_show_all
//
// [ show_all_workspaces ] trans: nothing
//
func (v Pager) SetShowAll(show_all_workspaces bool) {
	iv, err := _I.Get(25, "Pager", "set_show_all")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_show_all_workspaces := gi.NewBoolArgument(show_all_workspaces)
	args := []gi.Argument{arg_v, arg_show_all_workspaces}
	iv.Call(args, nil, nil)
}

// wnck_pager_set_wrap_on_scroll
//
// [ wrap_on_scroll ] trans: nothing
//
func (v Pager) SetWrapOnScroll(wrap_on_scroll bool) {
	iv, err := _I.Get(26, "Pager", "set_wrap_on_scroll")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_wrap_on_scroll := gi.NewBoolArgument(wrap_on_scroll)
	args := []gi.Argument{arg_v, arg_wrap_on_scroll}
	iv.Call(args, nil, nil)
}

// ignore GType struct PagerClass

// Enum PagerDisplayMode
type PagerDisplayModeEnum int

const (
	PagerDisplayModeName    PagerDisplayModeEnum = 0
	PagerDisplayModeContent PagerDisplayModeEnum = 1
)

func PagerDisplayModeGetType() gi.GType {
	ret := _I.GetGType(9, "PagerDisplayMode")
	return ret
}

// Struct PagerPrivate
type PagerPrivate struct {
	P unsafe.Pointer
}

func PagerPrivateGetType() gi.GType {
	ret := _I.GetGType(10, "PagerPrivate")
	return ret
}

// Struct ResourceUsage
type ResourceUsage struct {
	P unsafe.Pointer
}

const SizeOfStructResourceUsage = 96

func ResourceUsageGetType() gi.GType {
	ret := _I.GetGType(11, "ResourceUsage")
	return ret
}

// Object Screen
type Screen struct {
	g.Object
}

func WrapScreen(p unsafe.Pointer) (r Screen) { r.P = p; return }

type IScreen interface{ P_Screen() unsafe.Pointer }

func (v Screen) P_Screen() unsafe.Pointer { return v.P }
func ScreenGetType() gi.GType {
	ret := _I.GetGType(12, "Screen")
	return ret
}

// Deprecated
//
// wnck_screen_free_workspace_layout
//
// [ layout ] trans: nothing
//
func ScreenFreeWorkspaceLayout1(layout WorkspaceLayout) {
	iv, err := _I.Get(27, "Screen", "free_workspace_layout")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_layout := gi.NewPointerArgument(layout.P)
	args := []gi.Argument{arg_layout}
	iv.Call(args, nil, nil)
}

// wnck_screen_get
//
// [ index ] trans: nothing
//
// [ result ] trans: nothing
//
func ScreenGet1(index int32) (result Screen) {
	iv, err := _I.Get(28, "Screen", "get")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_index := gi.NewInt32Argument(index)
	args := []gi.Argument{arg_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// wnck_screen_get_for_root
//
// [ root_window_id ] trans: nothing
//
// [ result ] trans: nothing
//
func ScreenGetForRoot1(root_window_id uint64) (result Screen) {
	iv, err := _I.Get(30, "Screen", "get_for_root")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_root_window_id := gi.NewUint64Argument(root_window_id)
	args := []gi.Argument{arg_root_window_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Deprecated
//
// wnck_screen_calc_workspace_layout
//
// [ num_workspaces ] trans: nothing
//
// [ space_index ] trans: nothing
//
// [ layout ] trans: nothing
//
func (v Screen) CalcWorkspaceLayout(num_workspaces int32, space_index int32, layout WorkspaceLayout) {
	iv, err := _I.Get(31, "Screen", "calc_workspace_layout")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_num_workspaces := gi.NewInt32Argument(num_workspaces)
	arg_space_index := gi.NewInt32Argument(space_index)
	arg_layout := gi.NewPointerArgument(layout.P)
	args := []gi.Argument{arg_v, arg_num_workspaces, arg_space_index, arg_layout}
	iv.Call(args, nil, nil)
}

// wnck_screen_change_workspace_count
//
// [ count ] trans: nothing
//
func (v Screen) ChangeWorkspaceCount(count int32) {
	iv, err := _I.Get(32, "Screen", "change_workspace_count")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_count := gi.NewInt32Argument(count)
	args := []gi.Argument{arg_v, arg_count}
	iv.Call(args, nil, nil)
}

// wnck_screen_force_update
//
func (v Screen) ForceUpdate() {
	iv, err := _I.Get(33, "Screen", "force_update")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_screen_get_active_window
//
// [ result ] trans: nothing
//
func (v Screen) GetActiveWindow() (result Window) {
	iv, err := _I.Get(34, "Screen", "get_active_window")
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

// wnck_screen_get_active_workspace
//
// [ result ] trans: nothing
//
func (v Screen) GetActiveWorkspace() (result Workspace) {
	iv, err := _I.Get(35, "Screen", "get_active_workspace")
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

// wnck_screen_get_background_pixmap
//
// [ result ] trans: nothing
//
func (v Screen) GetBackgroundPixmap() (result uint64) {
	iv, err := _I.Get(36, "Screen", "get_background_pixmap")
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

// wnck_screen_get_height
//
// [ result ] trans: nothing
//
func (v Screen) GetHeight() (result int32) {
	iv, err := _I.Get(37, "Screen", "get_height")
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

// wnck_screen_get_number
//
// [ result ] trans: nothing
//
func (v Screen) GetNumber() (result int32) {
	iv, err := _I.Get(38, "Screen", "get_number")
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

// wnck_screen_get_previously_active_window
//
// [ result ] trans: nothing
//
func (v Screen) GetPreviouslyActiveWindow() (result Window) {
	iv, err := _I.Get(39, "Screen", "get_previously_active_window")
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

// wnck_screen_get_showing_desktop
//
// [ result ] trans: nothing
//
func (v Screen) GetShowingDesktop() (result bool) {
	iv, err := _I.Get(40, "Screen", "get_showing_desktop")
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

// wnck_screen_get_width
//
// [ result ] trans: nothing
//
func (v Screen) GetWidth() (result int32) {
	iv, err := _I.Get(41, "Screen", "get_width")
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

// wnck_screen_get_window_manager_name
//
// [ result ] trans: nothing
//
func (v Screen) GetWindowManagerName() (result string) {
	iv, err := _I.Get(42, "Screen", "get_window_manager_name")
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

// wnck_screen_get_windows
//
// [ result ] trans: nothing
//
func (v Screen) GetWindows() (result g.List) {
	iv, err := _I.Get(43, "Screen", "get_windows")
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

// wnck_screen_get_windows_stacked
//
// [ result ] trans: nothing
//
func (v Screen) GetWindowsStacked() (result g.List) {
	iv, err := _I.Get(44, "Screen", "get_windows_stacked")
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

// wnck_screen_get_workspace
//
// [ workspace ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Screen) GetWorkspace(workspace int32) (result Workspace) {
	iv, err := _I.Get(45, "Screen", "get_workspace")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_workspace := gi.NewInt32Argument(workspace)
	args := []gi.Argument{arg_v, arg_workspace}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// wnck_screen_get_workspace_count
//
// [ result ] trans: nothing
//
func (v Screen) GetWorkspaceCount() (result int32) {
	iv, err := _I.Get(46, "Screen", "get_workspace_count")
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

// wnck_screen_get_workspaces
//
// [ result ] trans: nothing
//
func (v Screen) GetWorkspaces() (result g.List) {
	iv, err := _I.Get(47, "Screen", "get_workspaces")
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

// wnck_screen_move_viewport
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
func (v Screen) MoveViewport(x int32, y int32) {
	iv, err := _I.Get(48, "Screen", "move_viewport")
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

// wnck_screen_net_wm_supports
//
// [ atom ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Screen) NetWmSupports(atom string) (result bool) {
	iv, err := _I.Get(49, "Screen", "net_wm_supports")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_atom := gi.CString(atom)
	arg_v := gi.NewPointerArgument(v.P)
	arg_atom := gi.NewStringArgument(c_atom)
	args := []gi.Argument{arg_v, arg_atom}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_atom)
	result = ret.Bool()
	return
}

// wnck_screen_release_workspace_layout
//
// [ current_token ] trans: nothing
//
func (v Screen) ReleaseWorkspaceLayout(current_token int32) {
	iv, err := _I.Get(50, "Screen", "release_workspace_layout")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_current_token := gi.NewInt32Argument(current_token)
	args := []gi.Argument{arg_v, arg_current_token}
	iv.Call(args, nil, nil)
}

// wnck_screen_toggle_showing_desktop
//
// [ show ] trans: nothing
//
func (v Screen) ToggleShowingDesktop(show bool) {
	iv, err := _I.Get(51, "Screen", "toggle_showing_desktop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_show := gi.NewBoolArgument(show)
	args := []gi.Argument{arg_v, arg_show}
	iv.Call(args, nil, nil)
}

// wnck_screen_try_set_workspace_layout
//
// [ current_token ] trans: nothing
//
// [ rows ] trans: nothing
//
// [ columns ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Screen) TrySetWorkspaceLayout(current_token int32, rows int32, columns int32) (result int32) {
	iv, err := _I.Get(52, "Screen", "try_set_workspace_layout")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_current_token := gi.NewInt32Argument(current_token)
	arg_rows := gi.NewInt32Argument(rows)
	arg_columns := gi.NewInt32Argument(columns)
	args := []gi.Argument{arg_v, arg_current_token, arg_rows, arg_columns}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// ignore GType struct ScreenClass

// Struct ScreenPrivate
type ScreenPrivate struct {
	P unsafe.Pointer
}

func ScreenPrivateGetType() gi.GType {
	ret := _I.GetGType(13, "ScreenPrivate")
	return ret
}

// Object Selector
type Selector struct {
	atk.ImplementorIfaceIfc
	gtk.BuildableIfc
	gtk.MenuBar
}

func WrapSelector(p unsafe.Pointer) (r Selector) { r.P = p; return }

type ISelector interface{ P_Selector() unsafe.Pointer }

func (v Selector) P_Selector() unsafe.Pointer         { return v.P }
func (v Selector) P_ImplementorIface() unsafe.Pointer { return v.P }
func (v Selector) P_Buildable() unsafe.Pointer        { return v.P }
func SelectorGetType() gi.GType {
	ret := _I.GetGType(14, "Selector")
	return ret
}

// wnck_selector_new
//
// [ result ] trans: nothing
//
func NewSelector() (result Selector) {
	iv, err := _I.Get(53, "Selector", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct SelectorClass

// Struct SelectorPrivate
type SelectorPrivate struct {
	P unsafe.Pointer
}

func SelectorPrivateGetType() gi.GType {
	ret := _I.GetGType(15, "SelectorPrivate")
	return ret
}

// Object Tasklist
type Tasklist struct {
	atk.ImplementorIfaceIfc
	gtk.BuildableIfc
	gtk.Container
}

func WrapTasklist(p unsafe.Pointer) (r Tasklist) { r.P = p; return }

type ITasklist interface{ P_Tasklist() unsafe.Pointer }

func (v Tasklist) P_Tasklist() unsafe.Pointer         { return v.P }
func (v Tasklist) P_ImplementorIface() unsafe.Pointer { return v.P }
func (v Tasklist) P_Buildable() unsafe.Pointer        { return v.P }
func TasklistGetType() gi.GType {
	ret := _I.GetGType(16, "Tasklist")
	return ret
}

// wnck_tasklist_new
//
// [ result ] trans: nothing
//
func NewTasklist() (result Tasklist) {
	iv, err := _I.Get(54, "Tasklist", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// wnck_tasklist_get_scroll_enabled
//
// [ result ] trans: nothing
//
func (v Tasklist) GetScrollEnabled() (result bool) {
	iv, err := _I.Get(55, "Tasklist", "get_scroll_enabled")
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

// wnck_tasklist_get_size_hint_list
//
// [ n_elements ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Tasklist) GetSizeHintList(n_elements int32) (result int32) {
	iv, err := _I.Get(56, "Tasklist", "get_size_hint_list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_elements := gi.NewInt32Argument(n_elements)
	args := []gi.Argument{arg_v, arg_n_elements}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// wnck_tasklist_set_button_relief
//
// [ relief ] trans: nothing
//
func (v Tasklist) SetButtonRelief(relief gtk.ReliefStyleEnum) {
	iv, err := _I.Get(57, "Tasklist", "set_button_relief")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relief := gi.NewIntArgument(int(relief))
	args := []gi.Argument{arg_v, arg_relief}
	iv.Call(args, nil, nil)
}

// wnck_tasklist_set_grouping
//
// [ grouping ] trans: nothing
//
func (v Tasklist) SetGrouping(grouping TasklistGroupingTypeEnum) {
	iv, err := _I.Get(58, "Tasklist", "set_grouping")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_grouping := gi.NewIntArgument(int(grouping))
	args := []gi.Argument{arg_v, arg_grouping}
	iv.Call(args, nil, nil)
}

// wnck_tasklist_set_grouping_limit
//
// [ limit ] trans: nothing
//
func (v Tasklist) SetGroupingLimit(limit int32) {
	iv, err := _I.Get(59, "Tasklist", "set_grouping_limit")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_limit := gi.NewInt32Argument(limit)
	args := []gi.Argument{arg_v, arg_limit}
	iv.Call(args, nil, nil)
}

// wnck_tasklist_set_include_all_workspaces
//
// [ include_all_workspaces ] trans: nothing
//
func (v Tasklist) SetIncludeAllWorkspaces(include_all_workspaces bool) {
	iv, err := _I.Get(60, "Tasklist", "set_include_all_workspaces")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_include_all_workspaces := gi.NewBoolArgument(include_all_workspaces)
	args := []gi.Argument{arg_v, arg_include_all_workspaces}
	iv.Call(args, nil, nil)
}

// wnck_tasklist_set_middle_click_close
//
// [ middle_click_close ] trans: nothing
//
func (v Tasklist) SetMiddleClickClose(middle_click_close bool) {
	iv, err := _I.Get(61, "Tasklist", "set_middle_click_close")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_middle_click_close := gi.NewBoolArgument(middle_click_close)
	args := []gi.Argument{arg_v, arg_middle_click_close}
	iv.Call(args, nil, nil)
}

// wnck_tasklist_set_orientation
//
// [ orient ] trans: nothing
//
func (v Tasklist) SetOrientation(orient gtk.OrientationEnum) {
	iv, err := _I.Get(62, "Tasklist", "set_orientation")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_orient := gi.NewIntArgument(int(orient))
	args := []gi.Argument{arg_v, arg_orient}
	iv.Call(args, nil, nil)
}

// wnck_tasklist_set_scroll_enabled
//
// [ scroll_enabled ] trans: nothing
//
func (v Tasklist) SetScrollEnabled(scroll_enabled bool) {
	iv, err := _I.Get(63, "Tasklist", "set_scroll_enabled")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_scroll_enabled := gi.NewBoolArgument(scroll_enabled)
	args := []gi.Argument{arg_v, arg_scroll_enabled}
	iv.Call(args, nil, nil)
}

// wnck_tasklist_set_switch_workspace_on_unminimize
//
// [ switch_workspace_on_unminimize ] trans: nothing
//
func (v Tasklist) SetSwitchWorkspaceOnUnminimize(switch_workspace_on_unminimize bool) {
	iv, err := _I.Get(64, "Tasklist", "set_switch_workspace_on_unminimize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_switch_workspace_on_unminimize := gi.NewBoolArgument(switch_workspace_on_unminimize)
	args := []gi.Argument{arg_v, arg_switch_workspace_on_unminimize}
	iv.Call(args, nil, nil)
}

// ignore GType struct TasklistClass

// Enum TasklistGroupingType
type TasklistGroupingTypeEnum int

const (
	TasklistGroupingTypeNeverGroup  TasklistGroupingTypeEnum = 0
	TasklistGroupingTypeAutoGroup   TasklistGroupingTypeEnum = 1
	TasklistGroupingTypeAlwaysGroup TasklistGroupingTypeEnum = 2
)

func TasklistGroupingTypeGetType() gi.GType {
	ret := _I.GetGType(17, "TasklistGroupingType")
	return ret
}

// Struct TasklistPrivate
type TasklistPrivate struct {
	P unsafe.Pointer
}

func TasklistPrivateGetType() gi.GType {
	ret := _I.GetGType(18, "TasklistPrivate")
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
	ret := _I.GetGType(19, "Window")
	return ret
}

// wnck_window_get
//
// [ xwindow ] trans: nothing
//
// [ result ] trans: nothing
//
func WindowGet1(xwindow uint64) (result Window) {
	iv, err := _I.Get(65, "Window", "get")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_xwindow := gi.NewUint64Argument(xwindow)
	args := []gi.Argument{arg_xwindow}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// wnck_window_activate
//
// [ timestamp ] trans: nothing
//
func (v Window) Activate(timestamp uint32) {
	iv, err := _I.Get(66, "Window", "activate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timestamp := gi.NewUint32Argument(timestamp)
	args := []gi.Argument{arg_v, arg_timestamp}
	iv.Call(args, nil, nil)
}

// wnck_window_activate_transient
//
// [ timestamp ] trans: nothing
//
func (v Window) ActivateTransient(timestamp uint32) {
	iv, err := _I.Get(67, "Window", "activate_transient")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timestamp := gi.NewUint32Argument(timestamp)
	args := []gi.Argument{arg_v, arg_timestamp}
	iv.Call(args, nil, nil)
}

// wnck_window_close
//
// [ timestamp ] trans: nothing
//
func (v Window) Close(timestamp uint32) {
	iv, err := _I.Get(68, "Window", "close")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timestamp := gi.NewUint32Argument(timestamp)
	args := []gi.Argument{arg_v, arg_timestamp}
	iv.Call(args, nil, nil)
}

// wnck_window_get_actions
//
// [ result ] trans: nothing
//
func (v Window) GetActions() (result WindowActionsFlags) {
	iv, err := _I.Get(69, "Window", "get_actions")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = WindowActionsFlags(ret.Int())
	return
}

// wnck_window_get_application
//
// [ result ] trans: nothing
//
func (v Window) GetApplication() (result Application) {
	iv, err := _I.Get(70, "Window", "get_application")
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

// wnck_window_get_class_group
//
// [ result ] trans: nothing
//
func (v Window) GetClassGroup() (result ClassGroup) {
	iv, err := _I.Get(71, "Window", "get_class_group")
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

// wnck_window_get_class_group_name
//
// [ result ] trans: nothing
//
func (v Window) GetClassGroupName() (result string) {
	iv, err := _I.Get(72, "Window", "get_class_group_name")
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

// wnck_window_get_class_instance_name
//
// [ result ] trans: nothing
//
func (v Window) GetClassInstanceName() (result string) {
	iv, err := _I.Get(73, "Window", "get_class_instance_name")
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

// wnck_window_get_client_window_geometry
//
// [ xp ] trans: everything, dir: out
//
// [ yp ] trans: everything, dir: out
//
// [ widthp ] trans: everything, dir: out
//
// [ heightp ] trans: everything, dir: out
//
func (v Window) GetClientWindowGeometry() (xp int32, yp int32, widthp int32, heightp int32) {
	iv, err := _I.Get(74, "Window", "get_client_window_geometry")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_xp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_yp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_widthp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_heightp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_xp, arg_yp, arg_widthp, arg_heightp}
	iv.Call(args, nil, &outArgs[0])
	xp = outArgs[0].Int32()
	yp = outArgs[1].Int32()
	widthp = outArgs[2].Int32()
	heightp = outArgs[3].Int32()
	return
}

// wnck_window_get_geometry
//
// [ xp ] trans: everything, dir: out
//
// [ yp ] trans: everything, dir: out
//
// [ widthp ] trans: everything, dir: out
//
// [ heightp ] trans: everything, dir: out
//
func (v Window) GetGeometry() (xp int32, yp int32, widthp int32, heightp int32) {
	iv, err := _I.Get(75, "Window", "get_geometry")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_xp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_yp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_widthp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_heightp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_xp, arg_yp, arg_widthp, arg_heightp}
	iv.Call(args, nil, &outArgs[0])
	xp = outArgs[0].Int32()
	yp = outArgs[1].Int32()
	widthp = outArgs[2].Int32()
	heightp = outArgs[3].Int32()
	return
}

// wnck_window_get_group_leader
//
// [ result ] trans: nothing
//
func (v Window) GetGroupLeader() (result uint64) {
	iv, err := _I.Get(76, "Window", "get_group_leader")
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

// wnck_window_get_icon
//
// [ result ] trans: nothing
//
func (v Window) GetIcon() (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(77, "Window", "get_icon")
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

// wnck_window_get_icon_is_fallback
//
// [ result ] trans: nothing
//
func (v Window) GetIconIsFallback() (result bool) {
	iv, err := _I.Get(78, "Window", "get_icon_is_fallback")
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

// wnck_window_get_icon_name
//
// [ result ] trans: nothing
//
func (v Window) GetIconName() (result string) {
	iv, err := _I.Get(79, "Window", "get_icon_name")
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

// wnck_window_get_mini_icon
//
// [ result ] trans: nothing
//
func (v Window) GetMiniIcon() (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(80, "Window", "get_mini_icon")
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

// wnck_window_get_name
//
// [ result ] trans: nothing
//
func (v Window) GetName() (result string) {
	iv, err := _I.Get(81, "Window", "get_name")
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

// wnck_window_get_pid
//
// [ result ] trans: nothing
//
func (v Window) GetPid() (result int32) {
	iv, err := _I.Get(82, "Window", "get_pid")
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

// wnck_window_get_role
//
// [ result ] trans: nothing
//
func (v Window) GetRole() (result string) {
	iv, err := _I.Get(83, "Window", "get_role")
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

// wnck_window_get_screen
//
// [ result ] trans: nothing
//
func (v Window) GetScreen() (result Screen) {
	iv, err := _I.Get(84, "Window", "get_screen")
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

// wnck_window_get_session_id
//
// [ result ] trans: nothing
//
func (v Window) GetSessionId() (result string) {
	iv, err := _I.Get(85, "Window", "get_session_id")
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

// wnck_window_get_session_id_utf8
//
// [ result ] trans: nothing
//
func (v Window) GetSessionIdUtf8() (result string) {
	iv, err := _I.Get(86, "Window", "get_session_id_utf8")
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

// wnck_window_get_sort_order
//
// [ result ] trans: nothing
//
func (v Window) GetSortOrder() (result int32) {
	iv, err := _I.Get(87, "Window", "get_sort_order")
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

// wnck_window_get_state
//
// [ result ] trans: nothing
//
func (v Window) GetState() (result WindowStateFlags) {
	iv, err := _I.Get(88, "Window", "get_state")
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

// wnck_window_get_transient
//
// [ result ] trans: nothing
//
func (v Window) GetTransient() (result Window) {
	iv, err := _I.Get(89, "Window", "get_transient")
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

// wnck_window_get_window_type
//
// [ result ] trans: nothing
//
func (v Window) GetWindowType() (result WindowTypeEnum) {
	iv, err := _I.Get(90, "Window", "get_window_type")
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

// wnck_window_get_workspace
//
// [ result ] trans: nothing
//
func (v Window) GetWorkspace() (result Workspace) {
	iv, err := _I.Get(91, "Window", "get_workspace")
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

// wnck_window_get_xid
//
// [ result ] trans: nothing
//
func (v Window) GetXid() (result uint64) {
	iv, err := _I.Get(92, "Window", "get_xid")
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

// wnck_window_has_icon_name
//
// [ result ] trans: nothing
//
func (v Window) HasIconName() (result bool) {
	iv, err := _I.Get(93, "Window", "has_icon_name")
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

// wnck_window_has_name
//
// [ result ] trans: nothing
//
func (v Window) HasName() (result bool) {
	iv, err := _I.Get(94, "Window", "has_name")
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

// wnck_window_is_above
//
// [ result ] trans: nothing
//
func (v Window) IsAbove() (result bool) {
	iv, err := _I.Get(95, "Window", "is_above")
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

// wnck_window_is_active
//
// [ result ] trans: nothing
//
func (v Window) IsActive() (result bool) {
	iv, err := _I.Get(96, "Window", "is_active")
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

// wnck_window_is_below
//
// [ result ] trans: nothing
//
func (v Window) IsBelow() (result bool) {
	iv, err := _I.Get(97, "Window", "is_below")
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

// wnck_window_is_fullscreen
//
// [ result ] trans: nothing
//
func (v Window) IsFullscreen() (result bool) {
	iv, err := _I.Get(98, "Window", "is_fullscreen")
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

// wnck_window_is_in_viewport
//
// [ workspace ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Window) IsInViewport(workspace IWorkspace) (result bool) {
	iv, err := _I.Get(99, "Window", "is_in_viewport")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if workspace != nil {
		tmp = workspace.P_Workspace()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_workspace := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_workspace}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// wnck_window_is_maximized
//
// [ result ] trans: nothing
//
func (v Window) IsMaximized() (result bool) {
	iv, err := _I.Get(100, "Window", "is_maximized")
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

// wnck_window_is_maximized_horizontally
//
// [ result ] trans: nothing
//
func (v Window) IsMaximizedHorizontally() (result bool) {
	iv, err := _I.Get(101, "Window", "is_maximized_horizontally")
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

// wnck_window_is_maximized_vertically
//
// [ result ] trans: nothing
//
func (v Window) IsMaximizedVertically() (result bool) {
	iv, err := _I.Get(102, "Window", "is_maximized_vertically")
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

// wnck_window_is_minimized
//
// [ result ] trans: nothing
//
func (v Window) IsMinimized() (result bool) {
	iv, err := _I.Get(103, "Window", "is_minimized")
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

// wnck_window_is_most_recently_activated
//
// [ result ] trans: nothing
//
func (v Window) IsMostRecentlyActivated() (result bool) {
	iv, err := _I.Get(104, "Window", "is_most_recently_activated")
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

// wnck_window_is_on_workspace
//
// [ workspace ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Window) IsOnWorkspace(workspace IWorkspace) (result bool) {
	iv, err := _I.Get(105, "Window", "is_on_workspace")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if workspace != nil {
		tmp = workspace.P_Workspace()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_workspace := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_workspace}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// wnck_window_is_pinned
//
// [ result ] trans: nothing
//
func (v Window) IsPinned() (result bool) {
	iv, err := _I.Get(106, "Window", "is_pinned")
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

// wnck_window_is_shaded
//
// [ result ] trans: nothing
//
func (v Window) IsShaded() (result bool) {
	iv, err := _I.Get(107, "Window", "is_shaded")
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

// wnck_window_is_skip_pager
//
// [ result ] trans: nothing
//
func (v Window) IsSkipPager() (result bool) {
	iv, err := _I.Get(108, "Window", "is_skip_pager")
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

// wnck_window_is_skip_tasklist
//
// [ result ] trans: nothing
//
func (v Window) IsSkipTasklist() (result bool) {
	iv, err := _I.Get(109, "Window", "is_skip_tasklist")
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

// wnck_window_is_sticky
//
// [ result ] trans: nothing
//
func (v Window) IsSticky() (result bool) {
	iv, err := _I.Get(110, "Window", "is_sticky")
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

// wnck_window_is_visible_on_workspace
//
// [ workspace ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Window) IsVisibleOnWorkspace(workspace IWorkspace) (result bool) {
	iv, err := _I.Get(111, "Window", "is_visible_on_workspace")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if workspace != nil {
		tmp = workspace.P_Workspace()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_workspace := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_workspace}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// wnck_window_keyboard_move
//
func (v Window) KeyboardMove() {
	iv, err := _I.Get(112, "Window", "keyboard_move")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_keyboard_size
//
func (v Window) KeyboardSize() {
	iv, err := _I.Get(113, "Window", "keyboard_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_make_above
//
func (v Window) MakeAbove() {
	iv, err := _I.Get(114, "Window", "make_above")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_make_below
//
func (v Window) MakeBelow() {
	iv, err := _I.Get(115, "Window", "make_below")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_maximize
//
func (v Window) Maximize() {
	iv, err := _I.Get(116, "Window", "maximize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_maximize_horizontally
//
func (v Window) MaximizeHorizontally() {
	iv, err := _I.Get(117, "Window", "maximize_horizontally")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_maximize_vertically
//
func (v Window) MaximizeVertically() {
	iv, err := _I.Get(118, "Window", "maximize_vertically")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_minimize
//
func (v Window) Minimize() {
	iv, err := _I.Get(119, "Window", "minimize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_move_to_workspace
//
// [ space ] trans: nothing
//
func (v Window) MoveToWorkspace(space IWorkspace) {
	iv, err := _I.Get(120, "Window", "move_to_workspace")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if space != nil {
		tmp = space.P_Workspace()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_space := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_space}
	iv.Call(args, nil, nil)
}

// wnck_window_needs_attention
//
// [ result ] trans: nothing
//
func (v Window) NeedsAttention() (result bool) {
	iv, err := _I.Get(121, "Window", "needs_attention")
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

// wnck_window_or_transient_needs_attention
//
// [ result ] trans: nothing
//
func (v Window) OrTransientNeedsAttention() (result bool) {
	iv, err := _I.Get(122, "Window", "or_transient_needs_attention")
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

// wnck_window_pin
//
func (v Window) Pin() {
	iv, err := _I.Get(123, "Window", "pin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_set_fullscreen
//
// [ fullscreen ] trans: nothing
//
func (v Window) SetFullscreen(fullscreen bool) {
	iv, err := _I.Get(124, "Window", "set_fullscreen")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fullscreen := gi.NewBoolArgument(fullscreen)
	args := []gi.Argument{arg_v, arg_fullscreen}
	iv.Call(args, nil, nil)
}

// wnck_window_set_geometry
//
// [ gravity ] trans: nothing
//
// [ geometry_mask ] trans: nothing
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
func (v Window) SetGeometry(gravity WindowGravityEnum, geometry_mask WindowMoveResizeMaskFlags, x int32, y int32, width int32, height int32) {
	iv, err := _I.Get(125, "Window", "set_geometry")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_gravity := gi.NewIntArgument(int(gravity))
	arg_geometry_mask := gi.NewIntArgument(int(geometry_mask))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	args := []gi.Argument{arg_v, arg_gravity, arg_geometry_mask, arg_x, arg_y, arg_width, arg_height}
	iv.Call(args, nil, nil)
}

// wnck_window_set_icon_geometry
//
// [ x ] trans: nothing
//
// [ y ] trans: nothing
//
// [ width ] trans: nothing
//
// [ height ] trans: nothing
//
func (v Window) SetIconGeometry(x int32, y int32, width int32, height int32) {
	iv, err := _I.Get(126, "Window", "set_icon_geometry")
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

// wnck_window_set_skip_pager
//
// [ skip ] trans: nothing
//
func (v Window) SetSkipPager(skip bool) {
	iv, err := _I.Get(127, "Window", "set_skip_pager")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_skip := gi.NewBoolArgument(skip)
	args := []gi.Argument{arg_v, arg_skip}
	iv.Call(args, nil, nil)
}

// wnck_window_set_skip_tasklist
//
// [ skip ] trans: nothing
//
func (v Window) SetSkipTasklist(skip bool) {
	iv, err := _I.Get(128, "Window", "set_skip_tasklist")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_skip := gi.NewBoolArgument(skip)
	args := []gi.Argument{arg_v, arg_skip}
	iv.Call(args, nil, nil)
}

// wnck_window_set_sort_order
//
// [ order ] trans: nothing
//
func (v Window) SetSortOrder(order int32) {
	iv, err := _I.Get(129, "Window", "set_sort_order")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_order := gi.NewInt32Argument(order)
	args := []gi.Argument{arg_v, arg_order}
	iv.Call(args, nil, nil)
}

// wnck_window_set_window_type
//
// [ wintype ] trans: nothing
//
func (v Window) SetWindowType(wintype WindowTypeEnum) {
	iv, err := _I.Get(130, "Window", "set_window_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_wintype := gi.NewIntArgument(int(wintype))
	args := []gi.Argument{arg_v, arg_wintype}
	iv.Call(args, nil, nil)
}

// wnck_window_shade
//
func (v Window) Shade() {
	iv, err := _I.Get(131, "Window", "shade")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_stick
//
func (v Window) Stick() {
	iv, err := _I.Get(132, "Window", "stick")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_transient_is_most_recently_activated
//
// [ result ] trans: nothing
//
func (v Window) TransientIsMostRecentlyActivated() (result bool) {
	iv, err := _I.Get(133, "Window", "transient_is_most_recently_activated")
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

// wnck_window_unmake_above
//
func (v Window) UnmakeAbove() {
	iv, err := _I.Get(134, "Window", "unmake_above")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_unmake_below
//
func (v Window) UnmakeBelow() {
	iv, err := _I.Get(135, "Window", "unmake_below")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_unmaximize
//
func (v Window) Unmaximize() {
	iv, err := _I.Get(136, "Window", "unmaximize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_unmaximize_horizontally
//
func (v Window) UnmaximizeHorizontally() {
	iv, err := _I.Get(137, "Window", "unmaximize_horizontally")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_unmaximize_vertically
//
func (v Window) UnmaximizeVertically() {
	iv, err := _I.Get(138, "Window", "unmaximize_vertically")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_unminimize
//
// [ timestamp ] trans: nothing
//
func (v Window) Unminimize(timestamp uint32) {
	iv, err := _I.Get(139, "Window", "unminimize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timestamp := gi.NewUint32Argument(timestamp)
	args := []gi.Argument{arg_v, arg_timestamp}
	iv.Call(args, nil, nil)
}

// wnck_window_unpin
//
func (v Window) Unpin() {
	iv, err := _I.Get(140, "Window", "unpin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_unshade
//
func (v Window) Unshade() {
	iv, err := _I.Get(141, "Window", "unshade")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// wnck_window_unstick
//
func (v Window) Unstick() {
	iv, err := _I.Get(142, "Window", "unstick")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Flags WindowActions
type WindowActionsFlags int

const (
	WindowActionsMove                   WindowActionsFlags = 1
	WindowActionsResize                 WindowActionsFlags = 2
	WindowActionsShade                  WindowActionsFlags = 4
	WindowActionsStick                  WindowActionsFlags = 8
	WindowActionsMaximizeHorizontally   WindowActionsFlags = 16
	WindowActionsMaximizeVertically     WindowActionsFlags = 32
	WindowActionsChangeWorkspace        WindowActionsFlags = 64
	WindowActionsClose                  WindowActionsFlags = 128
	WindowActionsUnmaximizeHorizontally WindowActionsFlags = 256
	WindowActionsUnmaximizeVertically   WindowActionsFlags = 512
	WindowActionsUnshade                WindowActionsFlags = 1024
	WindowActionsUnstick                WindowActionsFlags = 2048
	WindowActionsMinimize               WindowActionsFlags = 4096
	WindowActionsUnminimize             WindowActionsFlags = 8192
	WindowActionsMaximize               WindowActionsFlags = 16384
	WindowActionsUnmaximize             WindowActionsFlags = 32768
	WindowActionsFullscreen             WindowActionsFlags = 65536
	WindowActionsAbove                  WindowActionsFlags = 131072
	WindowActionsBelow                  WindowActionsFlags = 262144
)

func WindowActionsGetType() gi.GType {
	ret := _I.GetGType(20, "WindowActions")
	return ret
}

// ignore GType struct WindowClass

// Enum WindowGravity
type WindowGravityEnum int

const (
	WindowGravityCurrent   WindowGravityEnum = 0
	WindowGravityNorthwest WindowGravityEnum = 1
	WindowGravityNorth     WindowGravityEnum = 2
	WindowGravityNortheast WindowGravityEnum = 3
	WindowGravityWest      WindowGravityEnum = 4
	WindowGravityCenter    WindowGravityEnum = 5
	WindowGravityEast      WindowGravityEnum = 6
	WindowGravitySouthwest WindowGravityEnum = 7
	WindowGravitySouth     WindowGravityEnum = 8
	WindowGravitySoutheast WindowGravityEnum = 9
	WindowGravityStatic    WindowGravityEnum = 10
)

func WindowGravityGetType() gi.GType {
	ret := _I.GetGType(21, "WindowGravity")
	return ret
}

// Flags WindowMoveResizeMask
type WindowMoveResizeMaskFlags int

const (
	WindowMoveResizeMaskX      WindowMoveResizeMaskFlags = 1
	WindowMoveResizeMaskY      WindowMoveResizeMaskFlags = 2
	WindowMoveResizeMaskWidth  WindowMoveResizeMaskFlags = 4
	WindowMoveResizeMaskHeight WindowMoveResizeMaskFlags = 8
)

func WindowMoveResizeMaskGetType() gi.GType {
	ret := _I.GetGType(22, "WindowMoveResizeMask")
	return ret
}

// Struct WindowPrivate
type WindowPrivate struct {
	P unsafe.Pointer
}

func WindowPrivateGetType() gi.GType {
	ret := _I.GetGType(23, "WindowPrivate")
	return ret
}

// Flags WindowState
type WindowStateFlags int

const (
	WindowStateMinimized             WindowStateFlags = 1
	WindowStateMaximizedHorizontally WindowStateFlags = 2
	WindowStateMaximizedVertically   WindowStateFlags = 4
	WindowStateShaded                WindowStateFlags = 8
	WindowStateSkipPager             WindowStateFlags = 16
	WindowStateSkipTasklist          WindowStateFlags = 32
	WindowStateSticky                WindowStateFlags = 64
	WindowStateHidden                WindowStateFlags = 128
	WindowStateFullscreen            WindowStateFlags = 256
	WindowStateDemandsAttention      WindowStateFlags = 512
	WindowStateUrgent                WindowStateFlags = 1024
	WindowStateAbove                 WindowStateFlags = 2048
	WindowStateBelow                 WindowStateFlags = 4096
)

func WindowStateGetType() gi.GType {
	ret := _I.GetGType(24, "WindowState")
	return ret
}

// Enum WindowType
type WindowTypeEnum int

const (
	WindowTypeNormal       WindowTypeEnum = 0
	WindowTypeDesktop      WindowTypeEnum = 1
	WindowTypeDock         WindowTypeEnum = 2
	WindowTypeDialog       WindowTypeEnum = 3
	WindowTypeToolbar      WindowTypeEnum = 4
	WindowTypeMenu         WindowTypeEnum = 5
	WindowTypeUtility      WindowTypeEnum = 6
	WindowTypeSplashscreen WindowTypeEnum = 7
)

func WindowTypeGetType() gi.GType {
	ret := _I.GetGType(25, "WindowType")
	return ret
}

// Object Workspace
type Workspace struct {
	g.Object
}

func WrapWorkspace(p unsafe.Pointer) (r Workspace) { r.P = p; return }

type IWorkspace interface{ P_Workspace() unsafe.Pointer }

func (v Workspace) P_Workspace() unsafe.Pointer { return v.P }
func WorkspaceGetType() gi.GType {
	ret := _I.GetGType(26, "Workspace")
	return ret
}

// wnck_workspace_activate
//
// [ timestamp ] trans: nothing
//
func (v Workspace) Activate(timestamp uint32) {
	iv, err := _I.Get(143, "Workspace", "activate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timestamp := gi.NewUint32Argument(timestamp)
	args := []gi.Argument{arg_v, arg_timestamp}
	iv.Call(args, nil, nil)
}

// wnck_workspace_change_name
//
// [ name ] trans: nothing
//
func (v Workspace) ChangeName(name string) {
	iv, err := _I.Get(144, "Workspace", "change_name")
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

// wnck_workspace_get_height
//
// [ result ] trans: nothing
//
func (v Workspace) GetHeight() (result int32) {
	iv, err := _I.Get(145, "Workspace", "get_height")
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

// wnck_workspace_get_layout_column
//
// [ result ] trans: nothing
//
func (v Workspace) GetLayoutColumn() (result int32) {
	iv, err := _I.Get(146, "Workspace", "get_layout_column")
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

// wnck_workspace_get_layout_row
//
// [ result ] trans: nothing
//
func (v Workspace) GetLayoutRow() (result int32) {
	iv, err := _I.Get(147, "Workspace", "get_layout_row")
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

// wnck_workspace_get_name
//
// [ result ] trans: nothing
//
func (v Workspace) GetName() (result string) {
	iv, err := _I.Get(148, "Workspace", "get_name")
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

// wnck_workspace_get_neighbor
//
// [ direction ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Workspace) GetNeighbor(direction MotionDirectionEnum) (result Workspace) {
	iv, err := _I.Get(149, "Workspace", "get_neighbor")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_direction := gi.NewIntArgument(int(direction))
	args := []gi.Argument{arg_v, arg_direction}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// wnck_workspace_get_number
//
// [ result ] trans: nothing
//
func (v Workspace) GetNumber() (result int32) {
	iv, err := _I.Get(150, "Workspace", "get_number")
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

// wnck_workspace_get_screen
//
// [ result ] trans: nothing
//
func (v Workspace) GetScreen() (result Screen) {
	iv, err := _I.Get(151, "Workspace", "get_screen")
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

// wnck_workspace_get_viewport_x
//
// [ result ] trans: nothing
//
func (v Workspace) GetViewportX() (result int32) {
	iv, err := _I.Get(152, "Workspace", "get_viewport_x")
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

// wnck_workspace_get_viewport_y
//
// [ result ] trans: nothing
//
func (v Workspace) GetViewportY() (result int32) {
	iv, err := _I.Get(153, "Workspace", "get_viewport_y")
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

// wnck_workspace_get_width
//
// [ result ] trans: nothing
//
func (v Workspace) GetWidth() (result int32) {
	iv, err := _I.Get(154, "Workspace", "get_width")
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

// wnck_workspace_is_virtual
//
// [ result ] trans: nothing
//
func (v Workspace) IsVirtual() (result bool) {
	iv, err := _I.Get(155, "Workspace", "is_virtual")
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

// ignore GType struct WorkspaceClass

// Deprecated
//
// Struct WorkspaceLayout
type WorkspaceLayout struct {
	P unsafe.Pointer
}

const SizeOfStructWorkspaceLayout = 32

func WorkspaceLayoutGetType() gi.GType {
	ret := _I.GetGType(27, "WorkspaceLayout")
	return ret
}

// Struct WorkspacePrivate
type WorkspacePrivate struct {
	P unsafe.Pointer
}

func WorkspacePrivateGetType() gi.GType {
	ret := _I.GetGType(28, "WorkspacePrivate")
	return ret
}

// Enum _LayoutCorner
type _LayoutCornerEnum int

const (
	_LayoutCornerTopleft     _LayoutCornerEnum = 0
	_LayoutCornerTopright    _LayoutCornerEnum = 1
	_LayoutCornerBottomright _LayoutCornerEnum = 2
	_LayoutCornerBottomleft  _LayoutCornerEnum = 3
)

func _LayoutCornerGetType() gi.GType {
	ret := _I.GetGType(29, "_LayoutCorner")
	return ret
}

// Enum _LayoutOrientation
type _LayoutOrientationEnum int

const (
	_LayoutOrientationHorizontal _LayoutOrientationEnum = 0
	_LayoutOrientationVertical   _LayoutOrientationEnum = 1
)

func _LayoutOrientationGetType() gi.GType {
	ret := _I.GetGType(30, "_LayoutOrientation")
	return ret
}

// wnck_pid_read_resource_usage
//
// [ gdk_display ] trans: nothing
//
// [ pid ] trans: nothing
//
// [ usage ] trans: nothing
//
func PidReadResourceUsage(gdk_display gdk.IDisplay, pid uint64, usage ResourceUsage) {
	iv, err := _I.Get(156, "pid_read_resource_usage", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if gdk_display != nil {
		tmp = gdk_display.P_Display()
	}
	arg_gdk_display := gi.NewPointerArgument(tmp)
	arg_pid := gi.NewUint64Argument(pid)
	arg_usage := gi.NewPointerArgument(usage.P)
	args := []gi.Argument{arg_gdk_display, arg_pid, arg_usage}
	iv.Call(args, nil, nil)
}

// wnck_set_client_type
//
// [ ewmh_sourceindication_client_type ] trans: nothing
//
func SetClientType(ewmh_sourceindication_client_type ClientTypeEnum) {
	iv, err := _I.Get(157, "set_client_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ewmh_sourceindication_client_type := gi.NewIntArgument(int(ewmh_sourceindication_client_type))
	args := []gi.Argument{arg_ewmh_sourceindication_client_type}
	iv.Call(args, nil, nil)
}

// wnck_set_default_icon_size
//
// [ size ] trans: nothing
//
func SetDefaultIconSize(size uint64) {
	iv, err := _I.Get(158, "set_default_icon_size", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_size}
	iv.Call(args, nil, nil)
}

// wnck_set_default_mini_icon_size
//
// [ size ] trans: nothing
//
func SetDefaultMiniIconSize(size uint64) {
	iv, err := _I.Get(159, "set_default_mini_icon_size", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_size}
	iv.Call(args, nil, nil)
}

// wnck_shutdown
//
func Shutdown() {
	iv, err := _I.Get(160, "shutdown", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// wnck_xid_read_resource_usage
//
// [ gdk_display ] trans: nothing
//
// [ xid ] trans: nothing
//
// [ usage ] trans: nothing
//
func XidReadResourceUsage(gdk_display gdk.IDisplay, xid uint64, usage ResourceUsage) {
	iv, err := _I.Get(161, "xid_read_resource_usage", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if gdk_display != nil {
		tmp = gdk_display.P_Display()
	}
	arg_gdk_display := gi.NewPointerArgument(tmp)
	arg_xid := gi.NewUint64Argument(xid)
	arg_usage := gi.NewPointerArgument(usage.P)
	args := []gi.Argument{arg_gdk_display, arg_xid, arg_usage}
	iv.Call(args, nil, nil)
}

// constants
const (
	DEFAULT_ICON_SIZE      = 32
	DEFAULT_MINI_ICON_SIZE = 16
	MAJOR_VERSION          = 3
	MICRO_VERSION          = 0
	MINOR_VERSION          = 30
)
