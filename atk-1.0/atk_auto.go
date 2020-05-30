package atk

import "github.com/electricface/go-gir/glib-2.0"
import "github.com/electricface/go-gir/gobject-2.0"
import "github.com/electricface/go-gir3/gi"
import "log"
import "unsafe"

var _I = gi.NewInvokerCache("Atk")
var _ unsafe.Pointer

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("Atk", "1.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Interface Action
type Action struct {
	ActionIfc
	P unsafe.Pointer
}
type ActionIfc struct{}

// atk_action_do_action
// container is not nil, container is Action
// is method
func (v *ActionIfc) DoAction(i int32) (result bool) {
	iv, err := _I.Get(0, "Action", "do_action")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_action_get_description
// container is not nil, container is Action
// is method
func (v *ActionIfc) GetDescription(i int32) (result string) {
	iv, err := _I.Get(1, "Action", "get_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_action_get_keybinding
// container is not nil, container is Action
// is method
func (v *ActionIfc) GetKeybinding(i int32) (result string) {
	iv, err := _I.Get(2, "Action", "get_keybinding")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_action_get_localized_name
// container is not nil, container is Action
// is method
func (v *ActionIfc) GetLocalizedName(i int32) (result string) {
	iv, err := _I.Get(3, "Action", "get_localized_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_action_get_n_actions
// container is not nil, container is Action
// is method
func (v *ActionIfc) GetNActions() (result int32) {
	iv, err := _I.Get(4, "Action", "get_n_actions")
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

// atk_action_get_name
// container is not nil, container is Action
// is method
func (v *ActionIfc) GetName(i int32) (result string) {
	iv, err := _I.Get(5, "Action", "get_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_action_set_description
// container is not nil, container is Action
// is method
func (v *ActionIfc) SetDescription(i int32, desc string) (result bool) {
	iv, err := _I.Get(6, "Action", "set_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_desc := gi.CString(desc)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	arg_desc := gi.NewStringArgument(c_desc)
	args := []gi.Argument{arg_v, arg_i, arg_desc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	gi.Free(c_desc)
	return
}

// ignore GType struct ActionIface
// Struct Attribute
type Attribute struct {
	P unsafe.Pointer
}

// atk_attribute_set_free
// container is not nil, container is Attribute
// is method
// arg0Type tag: gslist, isPtr: true
func AttributeSetFree1(attrib_set int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(7, "Attribute", "set_free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_attrib_set := gi.NewIntArgument(attrib_set) /*TODO*/
	args := []gi.Argument{arg_attrib_set}
	iv.Call(args, nil, nil)
}

// Interface Component
type Component struct {
	ComponentIfc
	P unsafe.Pointer
}
type ComponentIfc struct{}

// atk_component_contains
// container is not nil, container is Component
// is method
func (v *ComponentIfc) Contains(x int32, y int32, coord_type int /*TODO_TYPE isPtr: false, tag: interface*/) (result bool) {
	iv, err := _I.Get(8, "Component", "contains")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_coord_type := gi.NewIntArgument(coord_type) /*TODO*/
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_coord_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_component_get_alpha
// container is not nil, container is Component
// is method
func (v *ComponentIfc) GetAlpha() (result float64) {
	iv, err := _I.Get(9, "Component", "get_alpha")
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

// atk_component_get_extents
// container is not nil, container is Component
// is method
func (v *ComponentIfc) GetExtents(coord_type int /*TODO_TYPE isPtr: false, tag: interface*/) (x int32, y int32, width int32, height int32) {
	iv, err := _I.Get(10, "Component", "get_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	arg_coord_type := gi.NewIntArgument(coord_type) /*TODO*/
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_width, arg_height, arg_coord_type}
	iv.Call(args, nil, &outArgs[0])
	x = outArgs[0].Int32()
	y = outArgs[1].Int32()
	width = outArgs[2].Int32()
	height = outArgs[3].Int32()
	return
}

// atk_component_get_layer
// container is not nil, container is Component
// is method
func (v *ComponentIfc) GetLayer() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(11, "Component", "get_layer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_component_get_mdi_zorder
// container is not nil, container is Component
// is method
func (v *ComponentIfc) GetMdiZorder() (result int32) {
	iv, err := _I.Get(12, "Component", "get_mdi_zorder")
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

// atk_component_get_position
// container is not nil, container is Component
// is method
func (v *ComponentIfc) GetPosition(coord_type int /*TODO_TYPE isPtr: false, tag: interface*/) (x int32, y int32) {
	iv, err := _I.Get(13, "Component", "get_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_coord_type := gi.NewIntArgument(coord_type) /*TODO*/
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_coord_type}
	iv.Call(args, nil, &outArgs[0])
	x = outArgs[0].Int32()
	y = outArgs[1].Int32()
	return
}

// atk_component_get_size
// container is not nil, container is Component
// is method
func (v *ComponentIfc) GetSize() (width int32, height int32) {
	iv, err := _I.Get(14, "Component", "get_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_width, arg_height}
	iv.Call(args, nil, &outArgs[0])
	width = outArgs[0].Int32()
	height = outArgs[1].Int32()
	return
}

// atk_component_grab_focus
// container is not nil, container is Component
// is method
func (v *ComponentIfc) GrabFocus() (result bool) {
	iv, err := _I.Get(15, "Component", "grab_focus")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_component_ref_accessible_at_point
// container is not nil, container is Component
// is method
func (v *ComponentIfc) RefAccessibleAtPoint(x int32, y int32, coord_type int /*TODO_TYPE isPtr: false, tag: interface*/) (result Object) {
	iv, err := _I.Get(16, "Component", "ref_accessible_at_point")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_coord_type := gi.NewIntArgument(coord_type) /*TODO*/
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_coord_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_component_remove_focus_handler
// container is not nil, container is Component
// is method
func (v *ComponentIfc) RemoveFocusHandler(handler_id uint32) {
	iv, err := _I.Get(17, "Component", "remove_focus_handler")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_handler_id := gi.NewUint32Argument(handler_id)
	args := []gi.Argument{arg_v, arg_handler_id}
	iv.Call(args, nil, nil)
}

// atk_component_scroll_to
// container is not nil, container is Component
// is method
func (v *ComponentIfc) ScrollTo(type1 int /*TODO_TYPE isPtr: false, tag: interface*/) (result bool) {
	iv, err := _I.Get(18, "Component", "scroll_to")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_type1 := gi.NewIntArgument(type1) /*TODO*/
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_component_scroll_to_point
// container is not nil, container is Component
// is method
func (v *ComponentIfc) ScrollToPoint(coords int /*TODO_TYPE isPtr: false, tag: interface*/, x int32, y int32) (result bool) {
	iv, err := _I.Get(19, "Component", "scroll_to_point")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_coords := gi.NewIntArgument(coords) /*TODO*/
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	args := []gi.Argument{arg_v, arg_coords, arg_x, arg_y}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_component_set_extents
// container is not nil, container is Component
// is method
func (v *ComponentIfc) SetExtents(x int32, y int32, width int32, height int32, coord_type int /*TODO_TYPE isPtr: false, tag: interface*/) (result bool) {
	iv, err := _I.Get(20, "Component", "set_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	arg_coord_type := gi.NewIntArgument(coord_type) /*TODO*/
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_width, arg_height, arg_coord_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_component_set_position
// container is not nil, container is Component
// is method
func (v *ComponentIfc) SetPosition(x int32, y int32, coord_type int /*TODO_TYPE isPtr: false, tag: interface*/) (result bool) {
	iv, err := _I.Get(21, "Component", "set_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_coord_type := gi.NewIntArgument(coord_type) /*TODO*/
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_coord_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_component_set_size
// container is not nil, container is Component
// is method
func (v *ComponentIfc) SetSize(width int32, height int32) (result bool) {
	iv, err := _I.Get(22, "Component", "set_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	args := []gi.Argument{arg_v, arg_width, arg_height}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// ignore GType struct ComponentIface
type CoordTypeEnum int

const (
	CoordTypeScreen CoordTypeEnum = 0
	CoordTypeWindow               = 1
	CoordTypeParent               = 2
)

// Interface Document
type Document struct {
	DocumentIfc
	P unsafe.Pointer
}
type DocumentIfc struct{}

// atk_document_get_attribute_value
// container is not nil, container is Document
// is method
func (v *DocumentIfc) GetAttributeValue(attribute_name string) (result string) {
	iv, err := _I.Get(23, "Document", "get_attribute_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_attribute_name := gi.CString(attribute_name)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_attribute_name := gi.NewStringArgument(c_attribute_name)
	args := []gi.Argument{arg_v, arg_attribute_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	gi.Free(c_attribute_name)
	return
}

// atk_document_get_attributes
// container is not nil, container is Document
// is method
func (v *DocumentIfc) GetAttributes() (result int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(24, "Document", "get_attributes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_document_get_current_page_number
// container is not nil, container is Document
// is method
func (v *DocumentIfc) GetCurrentPageNumber() (result int32) {
	iv, err := _I.Get(25, "Document", "get_current_page_number")
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

// atk_document_get_document
// container is not nil, container is Document
// is method
func (v *DocumentIfc) GetDocument() {
	iv, err := _I.Get(26, "Document", "get_document")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// atk_document_get_document_type
// container is not nil, container is Document
// is method
func (v *DocumentIfc) GetDocumentType() (result string) {
	iv, err := _I.Get(27, "Document", "get_document_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_document_get_locale
// container is not nil, container is Document
// is method
func (v *DocumentIfc) GetLocale() (result string) {
	iv, err := _I.Get(28, "Document", "get_locale")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_document_get_page_count
// container is not nil, container is Document
// is method
func (v *DocumentIfc) GetPageCount() (result int32) {
	iv, err := _I.Get(29, "Document", "get_page_count")
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

// atk_document_set_attribute_value
// container is not nil, container is Document
// is method
func (v *DocumentIfc) SetAttributeValue(attribute_name string, attribute_value string) (result bool) {
	iv, err := _I.Get(30, "Document", "set_attribute_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_attribute_name := gi.CString(attribute_name)
	c_attribute_value := gi.CString(attribute_value)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_attribute_name := gi.NewStringArgument(c_attribute_name)
	arg_attribute_value := gi.NewStringArgument(c_attribute_value)
	args := []gi.Argument{arg_v, arg_attribute_name, arg_attribute_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	gi.Free(c_attribute_name)
	gi.Free(c_attribute_value)
	return
}

// ignore GType struct DocumentIface
// Interface EditableText
type EditableText struct {
	EditableTextIfc
	P unsafe.Pointer
}
type EditableTextIfc struct{}

// atk_editable_text_copy_text
// container is not nil, container is EditableText
// is method
func (v *EditableTextIfc) CopyText(start_pos int32, end_pos int32) {
	iv, err := _I.Get(31, "EditableText", "copy_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_pos := gi.NewInt32Argument(start_pos)
	arg_end_pos := gi.NewInt32Argument(end_pos)
	args := []gi.Argument{arg_v, arg_start_pos, arg_end_pos}
	iv.Call(args, nil, nil)
}

// atk_editable_text_cut_text
// container is not nil, container is EditableText
// is method
func (v *EditableTextIfc) CutText(start_pos int32, end_pos int32) {
	iv, err := _I.Get(32, "EditableText", "cut_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_pos := gi.NewInt32Argument(start_pos)
	arg_end_pos := gi.NewInt32Argument(end_pos)
	args := []gi.Argument{arg_v, arg_start_pos, arg_end_pos}
	iv.Call(args, nil, nil)
}

// atk_editable_text_delete_text
// container is not nil, container is EditableText
// is method
func (v *EditableTextIfc) DeleteText(start_pos int32, end_pos int32) {
	iv, err := _I.Get(33, "EditableText", "delete_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_pos := gi.NewInt32Argument(start_pos)
	arg_end_pos := gi.NewInt32Argument(end_pos)
	args := []gi.Argument{arg_v, arg_start_pos, arg_end_pos}
	iv.Call(args, nil, nil)
}

// atk_editable_text_insert_text
// container is not nil, container is EditableText
// is method
func (v *EditableTextIfc) InsertText(string string, length int32, position int32) {
	iv, err := _I.Get(34, "EditableText", "insert_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_string := gi.NewStringArgument(c_string)
	arg_length := gi.NewInt32Argument(length)
	arg_position := gi.NewInt32Argument(position)
	args := []gi.Argument{arg_v, arg_string, arg_length, arg_position}
	iv.Call(args, nil, nil)
	gi.Free(c_string)
}

// atk_editable_text_paste_text
// container is not nil, container is EditableText
// is method
func (v *EditableTextIfc) PasteText(position int32) {
	iv, err := _I.Get(35, "EditableText", "paste_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_position := gi.NewInt32Argument(position)
	args := []gi.Argument{arg_v, arg_position}
	iv.Call(args, nil, nil)
}

// atk_editable_text_set_run_attributes
// container is not nil, container is EditableText
// is method
func (v *EditableTextIfc) SetRunAttributes(attrib_set int /*TODO_TYPE isPtr: true, tag: gslist*/, start_offset int32, end_offset int32) (result bool) {
	iv, err := _I.Get(36, "EditableText", "set_run_attributes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_attrib_set := gi.NewIntArgument(attrib_set) /*TODO*/
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	args := []gi.Argument{arg_v, arg_attrib_set, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_editable_text_set_text_contents
// container is not nil, container is EditableText
// is method
func (v *EditableTextIfc) SetTextContents(string string) {
	iv, err := _I.Get(37, "EditableText", "set_text_contents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_v, arg_string}
	iv.Call(args, nil, nil)
	gi.Free(c_string)
}

// ignore GType struct EditableTextIface
// Object GObjectAccessible
type GObjectAccessible struct {
	Object
}

// atk_gobject_accessible_for_object
// container is not nil, container is GObjectAccessible
// is method
// arg0Type tag: interface, isPtr: true
func GObjectAccessibleForObject1(obj gobject.Object) (result Object) {
	iv, err := _I.Get(38, "GObjectAccessible", "for_object")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_obj := gi.NewPointerArgument(obj.P)
	args := []gi.Argument{arg_obj}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_gobject_accessible_get_object
// container is not nil, container is GObjectAccessible
// is method
func (v GObjectAccessible) GetObject() (result gobject.Object) {
	iv, err := _I.Get(39, "GObjectAccessible", "get_object")
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

// ignore GType struct GObjectAccessibleClass
// Object Hyperlink
type Hyperlink struct {
	ActionIfc
	gobject.Object
}

// atk_hyperlink_get_end_index
// container is not nil, container is Hyperlink
// is method
func (v Hyperlink) GetEndIndex() (result int32) {
	iv, err := _I.Get(40, "Hyperlink", "get_end_index")
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

// atk_hyperlink_get_n_anchors
// container is not nil, container is Hyperlink
// is method
func (v Hyperlink) GetNAnchors() (result int32) {
	iv, err := _I.Get(41, "Hyperlink", "get_n_anchors")
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

// atk_hyperlink_get_object
// container is not nil, container is Hyperlink
// is method
func (v Hyperlink) GetObject(i int32) (result Object) {
	iv, err := _I.Get(42, "Hyperlink", "get_object")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_hyperlink_get_start_index
// container is not nil, container is Hyperlink
// is method
func (v Hyperlink) GetStartIndex() (result int32) {
	iv, err := _I.Get(43, "Hyperlink", "get_start_index")
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

// atk_hyperlink_get_uri
// container is not nil, container is Hyperlink
// is method
func (v Hyperlink) GetUri(i int32) (result string) {
	iv, err := _I.Get(44, "Hyperlink", "get_uri")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_hyperlink_is_inline
// container is not nil, container is Hyperlink
// is method
func (v Hyperlink) IsInline() (result bool) {
	iv, err := _I.Get(45, "Hyperlink", "is_inline")
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

// atk_hyperlink_is_selected_link
// container is not nil, container is Hyperlink
// is method
func (v Hyperlink) IsSelectedLink() (result bool) {
	iv, err := _I.Get(46, "Hyperlink", "is_selected_link")
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

// atk_hyperlink_is_valid
// container is not nil, container is Hyperlink
// is method
func (v Hyperlink) IsValid() (result bool) {
	iv, err := _I.Get(47, "Hyperlink", "is_valid")
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

// ignore GType struct HyperlinkClass
// Interface HyperlinkImpl
type HyperlinkImpl struct {
	HyperlinkImplIfc
	P unsafe.Pointer
}
type HyperlinkImplIfc struct{}

// atk_hyperlink_impl_get_hyperlink
// container is not nil, container is HyperlinkImpl
// is method
func (v *HyperlinkImplIfc) GetHyperlink() (result Hyperlink) {
	iv, err := _I.Get(48, "HyperlinkImpl", "get_hyperlink")
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

// ignore GType struct HyperlinkImplIface
type HyperlinkStateFlags int

const (
	HyperlinkStateFlagsInline HyperlinkStateFlags = 1
)

// Interface Hypertext
type Hypertext struct {
	HypertextIfc
	P unsafe.Pointer
}
type HypertextIfc struct{}

// atk_hypertext_get_link
// container is not nil, container is Hypertext
// is method
func (v *HypertextIfc) GetLink(link_index int32) (result Hyperlink) {
	iv, err := _I.Get(49, "Hypertext", "get_link")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_link_index := gi.NewInt32Argument(link_index)
	args := []gi.Argument{arg_v, arg_link_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_hypertext_get_link_index
// container is not nil, container is Hypertext
// is method
func (v *HypertextIfc) GetLinkIndex(char_index int32) (result int32) {
	iv, err := _I.Get(50, "Hypertext", "get_link_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_char_index := gi.NewInt32Argument(char_index)
	args := []gi.Argument{arg_v, arg_char_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_hypertext_get_n_links
// container is not nil, container is Hypertext
// is method
func (v *HypertextIfc) GetNLinks() (result int32) {
	iv, err := _I.Get(51, "Hypertext", "get_n_links")
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

// ignore GType struct HypertextIface
// Interface Image
type Image struct {
	ImageIfc
	P unsafe.Pointer
}
type ImageIfc struct{}

// atk_image_get_image_description
// container is not nil, container is Image
// is method
func (v *ImageIfc) GetImageDescription() (result string) {
	iv, err := _I.Get(52, "Image", "get_image_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_image_get_image_locale
// container is not nil, container is Image
// is method
func (v *ImageIfc) GetImageLocale() (result string) {
	iv, err := _I.Get(53, "Image", "get_image_locale")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_image_get_image_position
// container is not nil, container is Image
// is method
func (v *ImageIfc) GetImagePosition(coord_type int /*TODO_TYPE isPtr: false, tag: interface*/) (x int32, y int32) {
	iv, err := _I.Get(54, "Image", "get_image_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_coord_type := gi.NewIntArgument(coord_type) /*TODO*/
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_coord_type}
	iv.Call(args, nil, &outArgs[0])
	x = outArgs[0].Int32()
	y = outArgs[1].Int32()
	return
}

// atk_image_get_image_size
// container is not nil, container is Image
// is method
func (v *ImageIfc) GetImageSize() (width int32, height int32) {
	iv, err := _I.Get(55, "Image", "get_image_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_width, arg_height}
	iv.Call(args, nil, &outArgs[0])
	width = outArgs[0].Int32()
	height = outArgs[1].Int32()
	return
}

// atk_image_set_image_description
// container is not nil, container is Image
// is method
func (v *ImageIfc) SetImageDescription(description string) (result bool) {
	iv, err := _I.Get(56, "Image", "set_image_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_description := gi.CString(description)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_description := gi.NewStringArgument(c_description)
	args := []gi.Argument{arg_v, arg_description}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	gi.Free(c_description)
	return
}

// ignore GType struct ImageIface
// Struct Implementor
type Implementor struct {
	P unsafe.Pointer
}

// atk_implementor_ref_accessible
// container is not nil, container is Implementor
// is method
func (v Implementor) RefAccessible() (result Object) {
	iv, err := _I.Get(57, "Implementor", "ref_accessible")
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

// Interface ImplementorIface
type ImplementorIface struct {
	ImplementorIfaceIfc
	P unsafe.Pointer
}
type ImplementorIfaceIfc struct{}

// Struct KeyEventStruct
type KeyEventStruct struct {
	P unsafe.Pointer
}
type KeyEventTypeEnum int

const (
	KeyEventTypePress       KeyEventTypeEnum = 0
	KeyEventTypeRelease                      = 1
	KeyEventTypeLastDefined                  = 2
)

type LayerEnum int

const (
	LayerInvalid    LayerEnum = 0
	LayerBackground           = 1
	LayerCanvas               = 2
	LayerWidget               = 3
	LayerMdi                  = 4
	LayerPopup                = 5
	LayerOverlay              = 6
	LayerWindow               = 7
)

// Object Misc
type Misc struct {
	gobject.Object
}

// atk_misc_get_instance
// container is not nil, container is Misc
// num arg is 0
// atk_misc_threads_enter
// container is not nil, container is Misc
// is method
func (v Misc) ThreadsEnter() {
	iv, err := _I.Get(59, "Misc", "threads_enter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// atk_misc_threads_leave
// container is not nil, container is Misc
// is method
func (v Misc) ThreadsLeave() {
	iv, err := _I.Get(60, "Misc", "threads_leave")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// ignore GType struct MiscClass
// Object NoOpObject
type NoOpObject struct {
	ActionIfc
	ComponentIfc
	DocumentIfc
	EditableTextIfc
	HypertextIfc
	ImageIfc
	SelectionIfc
	TableIfc
	TableCellIfc
	TextIfc
	ValueIfc
	WindowIfc
	Object
}

// atk_no_op_object_new
// container is not nil, container is NoOpObject
// is constructor
func NewNoOpObject(obj gobject.Object) (result Object) {
	iv, err := _I.Get(61, "NoOpObject", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_obj := gi.NewPointerArgument(obj.P)
	args := []gi.Argument{arg_obj}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct NoOpObjectClass
// Object NoOpObjectFactory
type NoOpObjectFactory struct {
	ObjectFactory
}

// atk_no_op_object_factory_new
// container is not nil, container is NoOpObjectFactory
// is constructor
func NewNoOpObjectFactory() (result ObjectFactory) {
	iv, err := _I.Get(62, "NoOpObjectFactory", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct NoOpObjectFactoryClass
// Object Object
type Object struct {
	gobject.Object
}

// atk_object_add_relationship
// container is not nil, container is Object
// is method
func (v Object) AddRelationship(relationship int /*TODO_TYPE isPtr: false, tag: interface*/, target Object) (result bool) {
	iv, err := _I.Get(63, "Object", "add_relationship")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(relationship) /*TODO*/
	arg_target := gi.NewPointerArgument(target.P)
	args := []gi.Argument{arg_v, arg_relationship, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_object_get_attributes
// container is not nil, container is Object
// is method
func (v Object) GetAttributes() (result int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(64, "Object", "get_attributes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_object_get_description
// container is not nil, container is Object
// is method
func (v Object) GetDescription() (result string) {
	iv, err := _I.Get(65, "Object", "get_description")
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

// atk_object_get_index_in_parent
// container is not nil, container is Object
// is method
func (v Object) GetIndexInParent() (result int32) {
	iv, err := _I.Get(66, "Object", "get_index_in_parent")
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

// atk_object_get_layer
// container is not nil, container is Object
// is method
func (v Object) GetLayer() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(67, "Object", "get_layer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_object_get_mdi_zorder
// container is not nil, container is Object
// is method
func (v Object) GetMdiZorder() (result int32) {
	iv, err := _I.Get(68, "Object", "get_mdi_zorder")
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

// atk_object_get_n_accessible_children
// container is not nil, container is Object
// is method
func (v Object) GetNAccessibleChildren() (result int32) {
	iv, err := _I.Get(69, "Object", "get_n_accessible_children")
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

// atk_object_get_name
// container is not nil, container is Object
// is method
func (v Object) GetName() (result string) {
	iv, err := _I.Get(70, "Object", "get_name")
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

// atk_object_get_object_locale
// container is not nil, container is Object
// is method
func (v Object) GetObjectLocale() (result string) {
	iv, err := _I.Get(71, "Object", "get_object_locale")
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

// atk_object_get_parent
// container is not nil, container is Object
// is method
func (v Object) GetParent() (result Object) {
	iv, err := _I.Get(72, "Object", "get_parent")
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

// atk_object_get_role
// container is not nil, container is Object
// is method
func (v Object) GetRole() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(73, "Object", "get_role")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_object_initialize
// container is not nil, container is Object
// is method
func (v Object) Initialize(data unsafe.Pointer) {
	iv, err := _I.Get(74, "Object", "initialize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	iv.Call(args, nil, nil)
}

// atk_object_notify_state_change
// container is not nil, container is Object
// is method
func (v Object) NotifyStateChange(state uint64, value bool) {
	iv, err := _I.Get(75, "Object", "notify_state_change")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_state := gi.NewUint64Argument(state)
	arg_value := gi.NewBoolArgument(value)
	args := []gi.Argument{arg_v, arg_state, arg_value}
	iv.Call(args, nil, nil)
}

// atk_object_peek_parent
// container is not nil, container is Object
// is method
func (v Object) PeekParent() (result Object) {
	iv, err := _I.Get(76, "Object", "peek_parent")
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

// atk_object_ref_accessible_child
// container is not nil, container is Object
// is method
func (v Object) RefAccessibleChild(i int32) (result Object) {
	iv, err := _I.Get(77, "Object", "ref_accessible_child")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_object_ref_relation_set
// container is not nil, container is Object
// is method
func (v Object) RefRelationSet() (result RelationSet) {
	iv, err := _I.Get(78, "Object", "ref_relation_set")
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

// atk_object_ref_state_set
// container is not nil, container is Object
// is method
func (v Object) RefStateSet() (result StateSet) {
	iv, err := _I.Get(79, "Object", "ref_state_set")
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

// atk_object_remove_property_change_handler
// container is not nil, container is Object
// is method
func (v Object) RemovePropertyChangeHandler(handler_id uint32) {
	iv, err := _I.Get(80, "Object", "remove_property_change_handler")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_handler_id := gi.NewUint32Argument(handler_id)
	args := []gi.Argument{arg_v, arg_handler_id}
	iv.Call(args, nil, nil)
}

// atk_object_remove_relationship
// container is not nil, container is Object
// is method
func (v Object) RemoveRelationship(relationship int /*TODO_TYPE isPtr: false, tag: interface*/, target Object) (result bool) {
	iv, err := _I.Get(81, "Object", "remove_relationship")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(relationship) /*TODO*/
	arg_target := gi.NewPointerArgument(target.P)
	args := []gi.Argument{arg_v, arg_relationship, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_object_set_description
// container is not nil, container is Object
// is method
func (v Object) SetDescription(description string) {
	iv, err := _I.Get(82, "Object", "set_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_description := gi.CString(description)
	arg_v := gi.NewPointerArgument(v.P)
	arg_description := gi.NewStringArgument(c_description)
	args := []gi.Argument{arg_v, arg_description}
	iv.Call(args, nil, nil)
	gi.Free(c_description)
}

// atk_object_set_name
// container is not nil, container is Object
// is method
func (v Object) SetName(name string) {
	iv, err := _I.Get(83, "Object", "set_name")
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

// atk_object_set_parent
// container is not nil, container is Object
// is method
func (v Object) SetParent(parent Object) {
	iv, err := _I.Get(84, "Object", "set_parent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_parent := gi.NewPointerArgument(parent.P)
	args := []gi.Argument{arg_v, arg_parent}
	iv.Call(args, nil, nil)
}

// atk_object_set_role
// container is not nil, container is Object
// is method
func (v Object) SetRole(role int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(85, "Object", "set_role")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_role := gi.NewIntArgument(role) /*TODO*/
	args := []gi.Argument{arg_v, arg_role}
	iv.Call(args, nil, nil)
}

// ignore GType struct ObjectClass
// Object ObjectFactory
type ObjectFactory struct {
	gobject.Object
}

// atk_object_factory_create_accessible
// container is not nil, container is ObjectFactory
// is method
func (v ObjectFactory) CreateAccessible(obj gobject.Object) (result Object) {
	iv, err := _I.Get(86, "ObjectFactory", "create_accessible")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_obj := gi.NewPointerArgument(obj.P)
	args := []gi.Argument{arg_v, arg_obj}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_object_factory_get_accessible_type
// container is not nil, container is ObjectFactory
// is method
func (v ObjectFactory) GetAccessibleType() (result int /*TODO_TYPE isPtr: false, tag: GType*/) {
	iv, err := _I.Get(87, "ObjectFactory", "get_accessible_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_object_factory_invalidate
// container is not nil, container is ObjectFactory
// is method
func (v ObjectFactory) Invalidate() {
	iv, err := _I.Get(88, "ObjectFactory", "invalidate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// ignore GType struct ObjectFactoryClass
// Object Plug
type Plug struct {
	ComponentIfc
	Object
}

// atk_plug_new
// container is not nil, container is Plug
// is constructor
func NewPlug() (result Object) {
	iv, err := _I.Get(89, "Plug", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_plug_get_id
// container is not nil, container is Plug
// is method
func (v Plug) GetId() (result string) {
	iv, err := _I.Get(90, "Plug", "get_id")
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

// ignore GType struct PlugClass
// Struct PropertyValues
type PropertyValues struct {
	P unsafe.Pointer
}

// Struct Range
type Range struct {
	P unsafe.Pointer
}

// atk_range_new
// container is not nil, container is Range
// is constructor
func NewRange(lower_limit float64, upper_limit float64, description string) (result Range) {
	iv, err := _I.Get(91, "Range", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_description := gi.CString(description)
	arg_lower_limit := gi.NewDoubleArgument(lower_limit)
	arg_upper_limit := gi.NewDoubleArgument(upper_limit)
	arg_description := gi.NewStringArgument(c_description)
	args := []gi.Argument{arg_lower_limit, arg_upper_limit, arg_description}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	gi.Free(c_description)
	return
}

// atk_range_copy
// container is not nil, container is Range
// is method
func (v Range) Copy() (result Range) {
	iv, err := _I.Get(92, "Range", "copy")
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

// atk_range_free
// container is not nil, container is Range
// is method
func (v Range) Free() {
	iv, err := _I.Get(93, "Range", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// atk_range_get_description
// container is not nil, container is Range
// is method
func (v Range) GetDescription() (result string) {
	iv, err := _I.Get(94, "Range", "get_description")
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

// atk_range_get_lower_limit
// container is not nil, container is Range
// is method
func (v Range) GetLowerLimit() (result float64) {
	iv, err := _I.Get(95, "Range", "get_lower_limit")
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

// atk_range_get_upper_limit
// container is not nil, container is Range
// is method
func (v Range) GetUpperLimit() (result float64) {
	iv, err := _I.Get(96, "Range", "get_upper_limit")
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

// Struct Rectangle
type Rectangle struct {
	P unsafe.Pointer
}

// Object Registry
type Registry struct {
	gobject.Object
}

// atk_registry_get_factory
// container is not nil, container is Registry
// is method
func (v Registry) GetFactory(type1 int /*TODO_TYPE isPtr: false, tag: GType*/) (result ObjectFactory) {
	iv, err := _I.Get(97, "Registry", "get_factory")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(type1) /*TODO*/
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_registry_get_factory_type
// container is not nil, container is Registry
// is method
func (v Registry) GetFactoryType(type1 int /*TODO_TYPE isPtr: false, tag: GType*/) (result int /*TODO_TYPE isPtr: false, tag: GType*/) {
	iv, err := _I.Get(98, "Registry", "get_factory_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(type1) /*TODO*/
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_registry_set_factory_type
// container is not nil, container is Registry
// is method
func (v Registry) SetFactoryType(type1 int /*TODO_TYPE isPtr: false, tag: GType*/, factory_type int /*TODO_TYPE isPtr: false, tag: GType*/) {
	iv, err := _I.Get(99, "Registry", "set_factory_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(type1)               /*TODO*/
	arg_factory_type := gi.NewIntArgument(factory_type) /*TODO*/
	args := []gi.Argument{arg_v, arg_type1, arg_factory_type}
	iv.Call(args, nil, nil)
}

// ignore GType struct RegistryClass
// Object Relation
type Relation struct {
	gobject.Object
}

// atk_relation_new
// container is not nil, container is Relation
// is constructor
func NewRelation(targets int /*TODO_TYPE isPtr: true, tag: array*/, n_targets int32, relationship int /*TODO_TYPE isPtr: false, tag: interface*/) (result Relation) {
	iv, err := _I.Get(100, "Relation", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_targets := gi.NewIntArgument(targets) /*TODO*/
	arg_n_targets := gi.NewInt32Argument(n_targets)
	arg_relationship := gi.NewIntArgument(relationship) /*TODO*/
	args := []gi.Argument{arg_targets, arg_n_targets, arg_relationship}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_relation_add_target
// container is not nil, container is Relation
// is method
func (v Relation) AddTarget(target Object) {
	iv, err := _I.Get(101, "Relation", "add_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_target := gi.NewPointerArgument(target.P)
	args := []gi.Argument{arg_v, arg_target}
	iv.Call(args, nil, nil)
}

// atk_relation_get_relation_type
// container is not nil, container is Relation
// is method
func (v Relation) GetRelationType() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(102, "Relation", "get_relation_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_relation_get_target
// container is not nil, container is Relation
// is method
func (v Relation) GetTarget() (result int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(103, "Relation", "get_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_relation_remove_target
// container is not nil, container is Relation
// is method
func (v Relation) RemoveTarget(target Object) (result bool) {
	iv, err := _I.Get(104, "Relation", "remove_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_target := gi.NewPointerArgument(target.P)
	args := []gi.Argument{arg_v, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// ignore GType struct RelationClass
// Object RelationSet
type RelationSet struct {
	gobject.Object
}

// atk_relation_set_new
// container is not nil, container is RelationSet
// is constructor
func NewRelationSet() (result RelationSet) {
	iv, err := _I.Get(105, "RelationSet", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_relation_set_add
// container is not nil, container is RelationSet
// is method
func (v RelationSet) Add(relation Relation) {
	iv, err := _I.Get(106, "RelationSet", "add")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relation := gi.NewPointerArgument(relation.P)
	args := []gi.Argument{arg_v, arg_relation}
	iv.Call(args, nil, nil)
}

// atk_relation_set_add_relation_by_type
// container is not nil, container is RelationSet
// is method
func (v RelationSet) AddRelationByType(relationship int /*TODO_TYPE isPtr: false, tag: interface*/, target Object) {
	iv, err := _I.Get(107, "RelationSet", "add_relation_by_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(relationship) /*TODO*/
	arg_target := gi.NewPointerArgument(target.P)
	args := []gi.Argument{arg_v, arg_relationship, arg_target}
	iv.Call(args, nil, nil)
}

// atk_relation_set_contains
// container is not nil, container is RelationSet
// is method
func (v RelationSet) Contains(relationship int /*TODO_TYPE isPtr: false, tag: interface*/) (result bool) {
	iv, err := _I.Get(108, "RelationSet", "contains")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(relationship) /*TODO*/
	args := []gi.Argument{arg_v, arg_relationship}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_relation_set_contains_target
// container is not nil, container is RelationSet
// is method
func (v RelationSet) ContainsTarget(relationship int /*TODO_TYPE isPtr: false, tag: interface*/, target Object) (result bool) {
	iv, err := _I.Get(109, "RelationSet", "contains_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(relationship) /*TODO*/
	arg_target := gi.NewPointerArgument(target.P)
	args := []gi.Argument{arg_v, arg_relationship, arg_target}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_relation_set_get_n_relations
// container is not nil, container is RelationSet
// is method
func (v RelationSet) GetNRelations() (result int32) {
	iv, err := _I.Get(110, "RelationSet", "get_n_relations")
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

// atk_relation_set_get_relation
// container is not nil, container is RelationSet
// is method
func (v RelationSet) GetRelation(i int32) (result Relation) {
	iv, err := _I.Get(111, "RelationSet", "get_relation")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_relation_set_get_relation_by_type
// container is not nil, container is RelationSet
// is method
func (v RelationSet) GetRelationByType(relationship int /*TODO_TYPE isPtr: false, tag: interface*/) (result Relation) {
	iv, err := _I.Get(112, "RelationSet", "get_relation_by_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relationship := gi.NewIntArgument(relationship) /*TODO*/
	args := []gi.Argument{arg_v, arg_relationship}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_relation_set_remove
// container is not nil, container is RelationSet
// is method
func (v RelationSet) Remove(relation Relation) {
	iv, err := _I.Get(113, "RelationSet", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_relation := gi.NewPointerArgument(relation.P)
	args := []gi.Argument{arg_v, arg_relation}
	iv.Call(args, nil, nil)
}

// ignore GType struct RelationSetClass
type RelationTypeEnum int

const (
	RelationTypeNull           RelationTypeEnum = 0
	RelationTypeControlledBy                    = 1
	RelationTypeControllerFor                   = 2
	RelationTypeLabelFor                        = 3
	RelationTypeLabelledBy                      = 4
	RelationTypeMemberOf                        = 5
	RelationTypeNodeChildOf                     = 6
	RelationTypeFlowsTo                         = 7
	RelationTypeFlowsFrom                       = 8
	RelationTypeSubwindowOf                     = 9
	RelationTypeEmbeds                          = 10
	RelationTypeEmbeddedBy                      = 11
	RelationTypePopupFor                        = 12
	RelationTypeParentWindowOf                  = 13
	RelationTypeDescribedBy                     = 14
	RelationTypeDescriptionFor                  = 15
	RelationTypeNodeParentOf                    = 16
	RelationTypeDetails                         = 17
	RelationTypeDetailsFor                      = 18
	RelationTypeErrorMessage                    = 19
	RelationTypeErrorFor                        = 20
	RelationTypeLastDefined                     = 21
)

type RoleEnum int

const (
	RoleInvalid              RoleEnum = 0
	RoleAcceleratorLabel              = 1
	RoleAlert                         = 2
	RoleAnimation                     = 3
	RoleArrow                         = 4
	RoleCalendar                      = 5
	RoleCanvas                        = 6
	RoleCheckBox                      = 7
	RoleCheckMenuItem                 = 8
	RoleColorChooser                  = 9
	RoleColumnHeader                  = 10
	RoleComboBox                      = 11
	RoleDateEditor                    = 12
	RoleDesktopIcon                   = 13
	RoleDesktopFrame                  = 14
	RoleDial                          = 15
	RoleDialog                        = 16
	RoleDirectoryPane                 = 17
	RoleDrawingArea                   = 18
	RoleFileChooser                   = 19
	RoleFiller                        = 20
	RoleFontChooser                   = 21
	RoleFrame                         = 22
	RoleGlassPane                     = 23
	RoleHtmlContainer                 = 24
	RoleIcon                          = 25
	RoleImage                         = 26
	RoleInternalFrame                 = 27
	RoleLabel                         = 28
	RoleLayeredPane                   = 29
	RoleList                          = 30
	RoleListItem                      = 31
	RoleMenu                          = 32
	RoleMenuBar                       = 33
	RoleMenuItem                      = 34
	RoleOptionPane                    = 35
	RolePageTab                       = 36
	RolePageTabList                   = 37
	RolePanel                         = 38
	RolePasswordText                  = 39
	RolePopupMenu                     = 40
	RoleProgressBar                   = 41
	RolePushButton                    = 42
	RoleRadioButton                   = 43
	RoleRadioMenuItem                 = 44
	RoleRootPane                      = 45
	RoleRowHeader                     = 46
	RoleScrollBar                     = 47
	RoleScrollPane                    = 48
	RoleSeparator                     = 49
	RoleSlider                        = 50
	RoleSplitPane                     = 51
	RoleSpinButton                    = 52
	RoleStatusbar                     = 53
	RoleTable                         = 54
	RoleTableCell                     = 55
	RoleTableColumnHeader             = 56
	RoleTableRowHeader                = 57
	RoleTearOffMenuItem               = 58
	RoleTerminal                      = 59
	RoleText                          = 60
	RoleToggleButton                  = 61
	RoleToolBar                       = 62
	RoleToolTip                       = 63
	RoleTree                          = 64
	RoleTreeTable                     = 65
	RoleUnknown                       = 66
	RoleViewport                      = 67
	RoleWindow                        = 68
	RoleHeader                        = 69
	RoleFooter                        = 70
	RoleParagraph                     = 71
	RoleRuler                         = 72
	RoleApplication                   = 73
	RoleAutocomplete                  = 74
	RoleEditBar                       = 75
	RoleEmbedded                      = 76
	RoleEntry                         = 77
	RoleChart                         = 78
	RoleCaption                       = 79
	RoleDocumentFrame                 = 80
	RoleHeading                       = 81
	RolePage                          = 82
	RoleSection                       = 83
	RoleRedundantObject               = 84
	RoleForm                          = 85
	RoleLink                          = 86
	RoleInputMethodWindow             = 87
	RoleTableRow                      = 88
	RoleTreeItem                      = 89
	RoleDocumentSpreadsheet           = 90
	RoleDocumentPresentation          = 91
	RoleDocumentText                  = 92
	RoleDocumentWeb                   = 93
	RoleDocumentEmail                 = 94
	RoleComment                       = 95
	RoleListBox                       = 96
	RoleGrouping                      = 97
	RoleImageMap                      = 98
	RoleNotification                  = 99
	RoleInfoBar                       = 100
	RoleLevelBar                      = 101
	RoleTitleBar                      = 102
	RoleBlockQuote                    = 103
	RoleAudio                         = 104
	RoleVideo                         = 105
	RoleDefinition                    = 106
	RoleArticle                       = 107
	RoleLandmark                      = 108
	RoleLog                           = 109
	RoleMarquee                       = 110
	RoleMath                          = 111
	RoleRating                        = 112
	RoleTimer                         = 113
	RoleDescriptionList               = 114
	RoleDescriptionTerm               = 115
	RoleDescriptionValue              = 116
	RoleStatic                        = 117
	RoleMathFraction                  = 118
	RoleMathRoot                      = 119
	RoleSubscript                     = 120
	RoleSuperscript                   = 121
	RoleFootnote                      = 122
	RoleLastDefined                   = 123
)

type ScrollTypeEnum int

const (
	ScrollTypeTopLeft     ScrollTypeEnum = 0
	ScrollTypeBottomRight                = 1
	ScrollTypeTopEdge                    = 2
	ScrollTypeBottomEdge                 = 3
	ScrollTypeLeftEdge                   = 4
	ScrollTypeRightEdge                  = 5
	ScrollTypeAnywhere                   = 6
)

// Interface Selection
type Selection struct {
	SelectionIfc
	P unsafe.Pointer
}
type SelectionIfc struct{}

// atk_selection_add_selection
// container is not nil, container is Selection
// is method
func (v *SelectionIfc) AddSelection(i int32) (result bool) {
	iv, err := _I.Get(114, "Selection", "add_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_selection_clear_selection
// container is not nil, container is Selection
// is method
func (v *SelectionIfc) ClearSelection() (result bool) {
	iv, err := _I.Get(115, "Selection", "clear_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_selection_get_selection_count
// container is not nil, container is Selection
// is method
func (v *SelectionIfc) GetSelectionCount() (result int32) {
	iv, err := _I.Get(116, "Selection", "get_selection_count")
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

// atk_selection_is_child_selected
// container is not nil, container is Selection
// is method
func (v *SelectionIfc) IsChildSelected(i int32) (result bool) {
	iv, err := _I.Get(117, "Selection", "is_child_selected")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_selection_ref_selection
// container is not nil, container is Selection
// is method
func (v *SelectionIfc) RefSelection(i int32) (result Object) {
	iv, err := _I.Get(118, "Selection", "ref_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_selection_remove_selection
// container is not nil, container is Selection
// is method
func (v *SelectionIfc) RemoveSelection(i int32) (result bool) {
	iv, err := _I.Get(119, "Selection", "remove_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_selection_select_all_selection
// container is not nil, container is Selection
// is method
func (v *SelectionIfc) SelectAllSelection() (result bool) {
	iv, err := _I.Get(120, "Selection", "select_all_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// ignore GType struct SelectionIface
// Object Socket
type Socket struct {
	ComponentIfc
	Object
}

// atk_socket_new
// container is not nil, container is Socket
// is constructor
func NewSocket() (result Object) {
	iv, err := _I.Get(121, "Socket", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_socket_embed
// container is not nil, container is Socket
// is method
func (v Socket) Embed(plug_id string) {
	iv, err := _I.Get(122, "Socket", "embed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_plug_id := gi.CString(plug_id)
	arg_v := gi.NewPointerArgument(v.P)
	arg_plug_id := gi.NewStringArgument(c_plug_id)
	args := []gi.Argument{arg_v, arg_plug_id}
	iv.Call(args, nil, nil)
	gi.Free(c_plug_id)
}

// atk_socket_is_occupied
// container is not nil, container is Socket
// is method
func (v Socket) IsOccupied() (result bool) {
	iv, err := _I.Get(123, "Socket", "is_occupied")
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

// ignore GType struct SocketClass
// Object StateSet
type StateSet struct {
	gobject.Object
}

// atk_state_set_new
// container is not nil, container is StateSet
// is constructor
func NewStateSet() (result StateSet) {
	iv, err := _I.Get(124, "StateSet", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_state_set_add_state
// container is not nil, container is StateSet
// is method
func (v StateSet) AddState(type1 int /*TODO_TYPE isPtr: false, tag: interface*/) (result bool) {
	iv, err := _I.Get(125, "StateSet", "add_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(type1) /*TODO*/
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_state_set_add_states
// container is not nil, container is StateSet
// is method
func (v StateSet) AddStates(types int /*TODO_TYPE isPtr: true, tag: array*/, n_types int32) {
	iv, err := _I.Get(126, "StateSet", "add_states")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_types := gi.NewIntArgument(types) /*TODO*/
	arg_n_types := gi.NewInt32Argument(n_types)
	args := []gi.Argument{arg_v, arg_types, arg_n_types}
	iv.Call(args, nil, nil)
}

// atk_state_set_and_sets
// container is not nil, container is StateSet
// is method
func (v StateSet) AndSets(compare_set StateSet) (result StateSet) {
	iv, err := _I.Get(127, "StateSet", "and_sets")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_compare_set := gi.NewPointerArgument(compare_set.P)
	args := []gi.Argument{arg_v, arg_compare_set}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_state_set_clear_states
// container is not nil, container is StateSet
// is method
func (v StateSet) ClearStates() {
	iv, err := _I.Get(128, "StateSet", "clear_states")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// atk_state_set_contains_state
// container is not nil, container is StateSet
// is method
func (v StateSet) ContainsState(type1 int /*TODO_TYPE isPtr: false, tag: interface*/) (result bool) {
	iv, err := _I.Get(129, "StateSet", "contains_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(type1) /*TODO*/
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_state_set_contains_states
// container is not nil, container is StateSet
// is method
func (v StateSet) ContainsStates(types int /*TODO_TYPE isPtr: true, tag: array*/, n_types int32) (result bool) {
	iv, err := _I.Get(130, "StateSet", "contains_states")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_types := gi.NewIntArgument(types) /*TODO*/
	arg_n_types := gi.NewInt32Argument(n_types)
	args := []gi.Argument{arg_v, arg_types, arg_n_types}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_state_set_is_empty
// container is not nil, container is StateSet
// is method
func (v StateSet) IsEmpty() (result bool) {
	iv, err := _I.Get(131, "StateSet", "is_empty")
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

// atk_state_set_or_sets
// container is not nil, container is StateSet
// is method
func (v StateSet) OrSets(compare_set StateSet) (result StateSet) {
	iv, err := _I.Get(132, "StateSet", "or_sets")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_compare_set := gi.NewPointerArgument(compare_set.P)
	args := []gi.Argument{arg_v, arg_compare_set}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_state_set_remove_state
// container is not nil, container is StateSet
// is method
func (v StateSet) RemoveState(type1 int /*TODO_TYPE isPtr: false, tag: interface*/) (result bool) {
	iv, err := _I.Get(133, "StateSet", "remove_state")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(type1) /*TODO*/
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_state_set_xor_sets
// container is not nil, container is StateSet
// is method
func (v StateSet) XorSets(compare_set StateSet) (result StateSet) {
	iv, err := _I.Get(134, "StateSet", "xor_sets")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_compare_set := gi.NewPointerArgument(compare_set.P)
	args := []gi.Argument{arg_v, arg_compare_set}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct StateSetClass
type StateTypeEnum int

const (
	StateTypeInvalid                StateTypeEnum = 0
	StateTypeActive                               = 1
	StateTypeArmed                                = 2
	StateTypeBusy                                 = 3
	StateTypeChecked                              = 4
	StateTypeDefunct                              = 5
	StateTypeEditable                             = 6
	StateTypeEnabled                              = 7
	StateTypeExpandable                           = 8
	StateTypeExpanded                             = 9
	StateTypeFocusable                            = 10
	StateTypeFocused                              = 11
	StateTypeHorizontal                           = 12
	StateTypeIconified                            = 13
	StateTypeModal                                = 14
	StateTypeMultiLine                            = 15
	StateTypeMultiselectable                      = 16
	StateTypeOpaque                               = 17
	StateTypePressed                              = 18
	StateTypeResizable                            = 19
	StateTypeSelectable                           = 20
	StateTypeSelected                             = 21
	StateTypeSensitive                            = 22
	StateTypeShowing                              = 23
	StateTypeSingleLine                           = 24
	StateTypeStale                                = 25
	StateTypeTransient                            = 26
	StateTypeVertical                             = 27
	StateTypeVisible                              = 28
	StateTypeManagesDescendants                   = 29
	StateTypeIndeterminate                        = 30
	StateTypeTruncated                            = 31
	StateTypeRequired                             = 32
	StateTypeInvalidEntry                         = 33
	StateTypeSupportsAutocompletion               = 34
	StateTypeSelectableText                       = 35
	StateTypeDefault                              = 36
	StateTypeAnimated                             = 37
	StateTypeVisited                              = 38
	StateTypeCheckable                            = 39
	StateTypeHasPopup                             = 40
	StateTypeHasTooltip                           = 41
	StateTypeReadOnly                             = 42
	StateTypeLastDefined                          = 43
)

// Interface StreamableContent
type StreamableContent struct {
	StreamableContentIfc
	P unsafe.Pointer
}
type StreamableContentIfc struct{}

// atk_streamable_content_get_mime_type
// container is not nil, container is StreamableContent
// is method
func (v *StreamableContentIfc) GetMimeType(i int32) (result string) {
	iv, err := _I.Get(135, "StreamableContent", "get_mime_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_v, arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_streamable_content_get_n_mime_types
// container is not nil, container is StreamableContent
// is method
func (v *StreamableContentIfc) GetNMimeTypes() (result int32) {
	iv, err := _I.Get(136, "StreamableContent", "get_n_mime_types")
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

// atk_streamable_content_get_stream
// container is not nil, container is StreamableContent
// is method
func (v *StreamableContentIfc) GetStream(mime_type string) (result glib.IOChannel) {
	iv, err := _I.Get(137, "StreamableContent", "get_stream")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_mime_type := gi.CString(mime_type)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_mime_type := gi.NewStringArgument(c_mime_type)
	args := []gi.Argument{arg_v, arg_mime_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	gi.Free(c_mime_type)
	return
}

// atk_streamable_content_get_uri
// container is not nil, container is StreamableContent
// is method
func (v *StreamableContentIfc) GetUri(mime_type string) (result string) {
	iv, err := _I.Get(138, "StreamableContent", "get_uri")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_mime_type := gi.CString(mime_type)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_mime_type := gi.NewStringArgument(c_mime_type)
	args := []gi.Argument{arg_v, arg_mime_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	gi.Free(c_mime_type)
	return
}

// ignore GType struct StreamableContentIface
// Interface Table
type Table struct {
	TableIfc
	P unsafe.Pointer
}
type TableIfc struct{}

// atk_table_add_column_selection
// container is not nil, container is Table
// is method
func (v *TableIfc) AddColumnSelection(column int32) (result bool) {
	iv, err := _I.Get(139, "Table", "add_column_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_table_add_row_selection
// container is not nil, container is Table
// is method
func (v *TableIfc) AddRowSelection(row int32) (result bool) {
	iv, err := _I.Get(140, "Table", "add_row_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	args := []gi.Argument{arg_v, arg_row}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_table_get_caption
// container is not nil, container is Table
// is method
func (v *TableIfc) GetCaption() (result Object) {
	iv, err := _I.Get(141, "Table", "get_caption")
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

// atk_table_get_column_at_index
// container is not nil, container is Table
// is method
func (v *TableIfc) GetColumnAtIndex(index_ int32) (result int32) {
	iv, err := _I.Get(142, "Table", "get_column_at_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_index_ := gi.NewInt32Argument(index_)
	args := []gi.Argument{arg_v, arg_index_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_get_column_description
// container is not nil, container is Table
// is method
func (v *TableIfc) GetColumnDescription(column int32) (result string) {
	iv, err := _I.Get(143, "Table", "get_column_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_table_get_column_extent_at
// container is not nil, container is Table
// is method
func (v *TableIfc) GetColumnExtentAt(row int32, column int32) (result int32) {
	iv, err := _I.Get(144, "Table", "get_column_extent_at")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_row, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_get_column_header
// container is not nil, container is Table
// is method
func (v *TableIfc) GetColumnHeader(column int32) (result Object) {
	iv, err := _I.Get(145, "Table", "get_column_header")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_table_get_index_at
// container is not nil, container is Table
// is method
func (v *TableIfc) GetIndexAt(row int32, column int32) (result int32) {
	iv, err := _I.Get(146, "Table", "get_index_at")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_row, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_get_n_columns
// container is not nil, container is Table
// is method
func (v *TableIfc) GetNColumns() (result int32) {
	iv, err := _I.Get(147, "Table", "get_n_columns")
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

// atk_table_get_n_rows
// container is not nil, container is Table
// is method
func (v *TableIfc) GetNRows() (result int32) {
	iv, err := _I.Get(148, "Table", "get_n_rows")
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

// atk_table_get_row_at_index
// container is not nil, container is Table
// is method
func (v *TableIfc) GetRowAtIndex(index_ int32) (result int32) {
	iv, err := _I.Get(149, "Table", "get_row_at_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_index_ := gi.NewInt32Argument(index_)
	args := []gi.Argument{arg_v, arg_index_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_get_row_description
// container is not nil, container is Table
// is method
func (v *TableIfc) GetRowDescription(row int32) (result string) {
	iv, err := _I.Get(150, "Table", "get_row_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	args := []gi.Argument{arg_v, arg_row}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_table_get_row_extent_at
// container is not nil, container is Table
// is method
func (v *TableIfc) GetRowExtentAt(row int32, column int32) (result int32) {
	iv, err := _I.Get(151, "Table", "get_row_extent_at")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_row, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_get_row_header
// container is not nil, container is Table
// is method
func (v *TableIfc) GetRowHeader(row int32) (result Object) {
	iv, err := _I.Get(152, "Table", "get_row_header")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	args := []gi.Argument{arg_v, arg_row}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_table_get_selected_columns
// container is not nil, container is Table
// is method
func (v *TableIfc) GetSelectedColumns(selected int32) (result int32) {
	iv, err := _I.Get(153, "Table", "get_selected_columns")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_selected := gi.NewInt32Argument(selected)
	args := []gi.Argument{arg_v, arg_selected}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_get_selected_rows
// container is not nil, container is Table
// is method
func (v *TableIfc) GetSelectedRows(selected int32) (result int32) {
	iv, err := _I.Get(154, "Table", "get_selected_rows")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_selected := gi.NewInt32Argument(selected)
	args := []gi.Argument{arg_v, arg_selected}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_table_get_summary
// container is not nil, container is Table
// is method
func (v *TableIfc) GetSummary() (result Object) {
	iv, err := _I.Get(155, "Table", "get_summary")
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

// atk_table_is_column_selected
// container is not nil, container is Table
// is method
func (v *TableIfc) IsColumnSelected(column int32) (result bool) {
	iv, err := _I.Get(156, "Table", "is_column_selected")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_table_is_row_selected
// container is not nil, container is Table
// is method
func (v *TableIfc) IsRowSelected(row int32) (result bool) {
	iv, err := _I.Get(157, "Table", "is_row_selected")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	args := []gi.Argument{arg_v, arg_row}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_table_is_selected
// container is not nil, container is Table
// is method
func (v *TableIfc) IsSelected(row int32, column int32) (result bool) {
	iv, err := _I.Get(158, "Table", "is_selected")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_row, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_table_ref_at
// container is not nil, container is Table
// is method
func (v *TableIfc) RefAt(row int32, column int32) (result Object) {
	iv, err := _I.Get(159, "Table", "ref_at")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_row, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_table_remove_column_selection
// container is not nil, container is Table
// is method
func (v *TableIfc) RemoveColumnSelection(column int32) (result bool) {
	iv, err := _I.Get(160, "Table", "remove_column_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_table_remove_row_selection
// container is not nil, container is Table
// is method
func (v *TableIfc) RemoveRowSelection(row int32) (result bool) {
	iv, err := _I.Get(161, "Table", "remove_row_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	args := []gi.Argument{arg_v, arg_row}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_table_set_caption
// container is not nil, container is Table
// is method
func (v *TableIfc) SetCaption(caption Object) {
	iv, err := _I.Get(162, "Table", "set_caption")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_caption := gi.NewPointerArgument(caption.P)
	args := []gi.Argument{arg_v, arg_caption}
	iv.Call(args, nil, nil)
}

// atk_table_set_column_description
// container is not nil, container is Table
// is method
func (v *TableIfc) SetColumnDescription(column int32, description string) {
	iv, err := _I.Get(163, "Table", "set_column_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_description := gi.CString(description)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	arg_description := gi.NewStringArgument(c_description)
	args := []gi.Argument{arg_v, arg_column, arg_description}
	iv.Call(args, nil, nil)
	gi.Free(c_description)
}

// atk_table_set_column_header
// container is not nil, container is Table
// is method
func (v *TableIfc) SetColumnHeader(column int32, header Object) {
	iv, err := _I.Get(164, "Table", "set_column_header")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_column := gi.NewInt32Argument(column)
	arg_header := gi.NewPointerArgument(header.P)
	args := []gi.Argument{arg_v, arg_column, arg_header}
	iv.Call(args, nil, nil)
}

// atk_table_set_row_description
// container is not nil, container is Table
// is method
func (v *TableIfc) SetRowDescription(row int32, description string) {
	iv, err := _I.Get(165, "Table", "set_row_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_description := gi.CString(description)
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_description := gi.NewStringArgument(c_description)
	args := []gi.Argument{arg_v, arg_row, arg_description}
	iv.Call(args, nil, nil)
	gi.Free(c_description)
}

// atk_table_set_row_header
// container is not nil, container is Table
// is method
func (v *TableIfc) SetRowHeader(row int32, header Object) {
	iv, err := _I.Get(166, "Table", "set_row_header")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewInt32Argument(row)
	arg_header := gi.NewPointerArgument(header.P)
	args := []gi.Argument{arg_v, arg_row, arg_header}
	iv.Call(args, nil, nil)
}

// atk_table_set_summary
// container is not nil, container is Table
// is method
func (v *TableIfc) SetSummary(accessible Object) {
	iv, err := _I.Get(167, "Table", "set_summary")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_accessible := gi.NewPointerArgument(accessible.P)
	args := []gi.Argument{arg_v, arg_accessible}
	iv.Call(args, nil, nil)
}

// Interface TableCell
type TableCell struct {
	TableCellIfc
	P unsafe.Pointer
}
type TableCellIfc struct{}

// atk_table_cell_get_column_header_cells
// container is not nil, container is TableCell
// is method
func (v *TableCellIfc) GetColumnHeaderCells() (result int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(168, "TableCell", "get_column_header_cells")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_table_cell_get_column_span
// container is not nil, container is TableCell
// is method
func (v *TableCellIfc) GetColumnSpan() (result int32) {
	iv, err := _I.Get(169, "TableCell", "get_column_span")
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

// atk_table_cell_get_position
// container is not nil, container is TableCell
// is method
func (v *TableCellIfc) GetPosition() (result bool, row int32, column int32) {
	iv, err := _I.Get(170, "TableCell", "get_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_column := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_row, arg_column}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	row = outArgs[0].Int32()
	column = outArgs[1].Int32()
	return
}

// atk_table_cell_get_row_column_span
// container is not nil, container is TableCell
// is method
func (v *TableCellIfc) GetRowColumnSpan() (result bool, row int32, column int32, row_span int32, column_span int32) {
	iv, err := _I.Get(171, "TableCell", "get_row_column_span")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_row := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_column := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_row_span := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_column_span := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_row, arg_column, arg_row_span, arg_column_span}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	row = outArgs[0].Int32()
	column = outArgs[1].Int32()
	row_span = outArgs[2].Int32()
	column_span = outArgs[3].Int32()
	return
}

// atk_table_cell_get_row_header_cells
// container is not nil, container is TableCell
// is method
func (v *TableCellIfc) GetRowHeaderCells() (result int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(172, "TableCell", "get_row_header_cells")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_table_cell_get_row_span
// container is not nil, container is TableCell
// is method
func (v *TableCellIfc) GetRowSpan() (result int32) {
	iv, err := _I.Get(173, "TableCell", "get_row_span")
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

// atk_table_cell_get_table
// container is not nil, container is TableCell
// is method
func (v *TableCellIfc) GetTable() (result Object) {
	iv, err := _I.Get(174, "TableCell", "get_table")
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

// ignore GType struct TableCellIface
// ignore GType struct TableIface
// Interface Text
type Text struct {
	TextIfc
	P unsafe.Pointer
}
type TextIfc struct{}

// atk_text_free_ranges
// container is not nil, container is Text
// is method
// arg0Type tag: array, isPtr: true
func TextFreeRanges1(ranges int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(175, "Text", "free_ranges")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ranges := gi.NewIntArgument(ranges) /*TODO*/
	args := []gi.Argument{arg_ranges}
	iv.Call(args, nil, nil)
}

// atk_text_add_selection
// container is not nil, container is Text
// is method
func (v *TextIfc) AddSelection(start_offset int32, end_offset int32) (result bool) {
	iv, err := _I.Get(176, "Text", "add_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	args := []gi.Argument{arg_v, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_text_get_bounded_ranges
// container is not nil, container is Text
// is method
func (v *TextIfc) GetBoundedRanges(rect TextRectangle, coord_type int /*TODO_TYPE isPtr: false, tag: interface*/, x_clip_type int /*TODO_TYPE isPtr: false, tag: interface*/, y_clip_type int /*TODO_TYPE isPtr: false, tag: interface*/) (result int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(177, "Text", "get_bounded_ranges")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_rect := gi.NewPointerArgument(rect.P)
	arg_coord_type := gi.NewIntArgument(coord_type)   /*TODO*/
	arg_x_clip_type := gi.NewIntArgument(x_clip_type) /*TODO*/
	arg_y_clip_type := gi.NewIntArgument(y_clip_type) /*TODO*/
	args := []gi.Argument{arg_v, arg_rect, arg_coord_type, arg_x_clip_type, arg_y_clip_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_text_get_caret_offset
// container is not nil, container is Text
// is method
func (v *TextIfc) GetCaretOffset() (result int32) {
	iv, err := _I.Get(178, "Text", "get_caret_offset")
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

// atk_text_get_character_at_offset
// container is not nil, container is Text
// is method
func (v *TextIfc) GetCharacterAtOffset(offset int32) (result int /*TODO_TYPE isPtr: false, tag: gunichar*/) {
	iv, err := _I.Get(179, "Text", "get_character_at_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	args := []gi.Argument{arg_v, arg_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_text_get_character_count
// container is not nil, container is Text
// is method
func (v *TextIfc) GetCharacterCount() (result int32) {
	iv, err := _I.Get(180, "Text", "get_character_count")
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

// atk_text_get_character_extents
// container is not nil, container is Text
// is method
func (v *TextIfc) GetCharacterExtents(offset int32, coords int /*TODO_TYPE isPtr: false, tag: interface*/) (x int32, y int32, width int32, height int32) {
	iv, err := _I.Get(181, "Text", "get_character_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_x := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	arg_coords := gi.NewIntArgument(coords) /*TODO*/
	args := []gi.Argument{arg_v, arg_offset, arg_x, arg_y, arg_width, arg_height, arg_coords}
	iv.Call(args, nil, &outArgs[0])
	x = outArgs[0].Int32()
	y = outArgs[1].Int32()
	width = outArgs[2].Int32()
	height = outArgs[3].Int32()
	return
}

// atk_text_get_default_attributes
// container is not nil, container is Text
// is method
func (v *TextIfc) GetDefaultAttributes() (result int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(182, "Text", "get_default_attributes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_text_get_n_selections
// container is not nil, container is Text
// is method
func (v *TextIfc) GetNSelections() (result int32) {
	iv, err := _I.Get(183, "Text", "get_n_selections")
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

// atk_text_get_offset_at_point
// container is not nil, container is Text
// is method
func (v *TextIfc) GetOffsetAtPoint(x int32, y int32, coords int /*TODO_TYPE isPtr: false, tag: interface*/) (result int32) {
	iv, err := _I.Get(184, "Text", "get_offset_at_point")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_coords := gi.NewIntArgument(coords) /*TODO*/
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_coords}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// atk_text_get_range_extents
// container is not nil, container is Text
// is method
func (v *TextIfc) GetRangeExtents(start_offset int32, end_offset int32, coord_type int /*TODO_TYPE isPtr: false, tag: interface*/) (rect int /*TODO_TYPE*/) {
	iv, err := _I.Get(185, "Text", "get_range_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	arg_coord_type := gi.NewIntArgument(coord_type) /*TODO*/
	arg_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_start_offset, arg_end_offset, arg_coord_type, arg_rect}
	iv.Call(args, nil, &outArgs[0])
	rect = outArgs[0].Int() /*TODO*/
	return
}

// atk_text_get_run_attributes
// container is not nil, container is Text
// is method
func (v *TextIfc) GetRunAttributes(offset int32) (result int /*TODO_TYPE isPtr: true, tag: gslist*/, start_offset int32, end_offset int32) {
	iv, err := _I.Get(186, "Text", "get_run_attributes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_start_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_offset, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Int() /*TODO*/
	start_offset = outArgs[0].Int32()
	end_offset = outArgs[1].Int32()
	return
}

// atk_text_get_selection
// container is not nil, container is Text
// is method
func (v *TextIfc) GetSelection(selection_num int32) (result string, start_offset int32, end_offset int32) {
	iv, err := _I.Get(187, "Text", "get_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_selection_num := gi.NewInt32Argument(selection_num)
	arg_start_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_selection_num, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.String().Take()
	start_offset = outArgs[0].Int32()
	end_offset = outArgs[1].Int32()
	return
}

// atk_text_get_string_at_offset
// container is not nil, container is Text
// is method
func (v *TextIfc) GetStringAtOffset(offset int32, granularity int /*TODO_TYPE isPtr: false, tag: interface*/) (result string, start_offset int32, end_offset int32) {
	iv, err := _I.Get(188, "Text", "get_string_at_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_granularity := gi.NewIntArgument(granularity) /*TODO*/
	arg_start_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_offset, arg_granularity, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.String().Take()
	start_offset = outArgs[0].Int32()
	end_offset = outArgs[1].Int32()
	return
}

// atk_text_get_text
// container is not nil, container is Text
// is method
func (v *TextIfc) GetText(start_offset int32, end_offset int32) (result string) {
	iv, err := _I.Get(189, "Text", "get_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	args := []gi.Argument{arg_v, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_text_get_text_after_offset
// container is not nil, container is Text
// is method
func (v *TextIfc) GetTextAfterOffset(offset int32, boundary_type int /*TODO_TYPE isPtr: false, tag: interface*/) (result string, start_offset int32, end_offset int32) {
	iv, err := _I.Get(190, "Text", "get_text_after_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_boundary_type := gi.NewIntArgument(boundary_type) /*TODO*/
	arg_start_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_offset, arg_boundary_type, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.String().Take()
	start_offset = outArgs[0].Int32()
	end_offset = outArgs[1].Int32()
	return
}

// atk_text_get_text_at_offset
// container is not nil, container is Text
// is method
func (v *TextIfc) GetTextAtOffset(offset int32, boundary_type int /*TODO_TYPE isPtr: false, tag: interface*/) (result string, start_offset int32, end_offset int32) {
	iv, err := _I.Get(191, "Text", "get_text_at_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_boundary_type := gi.NewIntArgument(boundary_type) /*TODO*/
	arg_start_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_offset, arg_boundary_type, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.String().Take()
	start_offset = outArgs[0].Int32()
	end_offset = outArgs[1].Int32()
	return
}

// atk_text_get_text_before_offset
// container is not nil, container is Text
// is method
func (v *TextIfc) GetTextBeforeOffset(offset int32, boundary_type int /*TODO_TYPE isPtr: false, tag: interface*/) (result string, start_offset int32, end_offset int32) {
	iv, err := _I.Get(192, "Text", "get_text_before_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	arg_boundary_type := gi.NewIntArgument(boundary_type) /*TODO*/
	arg_start_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_offset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_offset, arg_boundary_type, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.String().Take()
	start_offset = outArgs[0].Int32()
	end_offset = outArgs[1].Int32()
	return
}

// atk_text_remove_selection
// container is not nil, container is Text
// is method
func (v *TextIfc) RemoveSelection(selection_num int32) (result bool) {
	iv, err := _I.Get(193, "Text", "remove_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_selection_num := gi.NewInt32Argument(selection_num)
	args := []gi.Argument{arg_v, arg_selection_num}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_text_set_caret_offset
// container is not nil, container is Text
// is method
func (v *TextIfc) SetCaretOffset(offset int32) (result bool) {
	iv, err := _I.Get(194, "Text", "set_caret_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_offset := gi.NewInt32Argument(offset)
	args := []gi.Argument{arg_v, arg_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_text_set_selection
// container is not nil, container is Text
// is method
func (v *TextIfc) SetSelection(selection_num int32, start_offset int32, end_offset int32) (result bool) {
	iv, err := _I.Get(195, "Text", "set_selection")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_selection_num := gi.NewInt32Argument(selection_num)
	arg_start_offset := gi.NewInt32Argument(start_offset)
	arg_end_offset := gi.NewInt32Argument(end_offset)
	args := []gi.Argument{arg_v, arg_selection_num, arg_start_offset, arg_end_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

type TextAttributeEnum int

const (
	TextAttributeInvalid          TextAttributeEnum = 0
	TextAttributeLeftMargin                         = 1
	TextAttributeRightMargin                        = 2
	TextAttributeIndent                             = 3
	TextAttributeInvisible                          = 4
	TextAttributeEditable                           = 5
	TextAttributePixelsAboveLines                   = 6
	TextAttributePixelsBelowLines                   = 7
	TextAttributePixelsInsideWrap                   = 8
	TextAttributeBgFullHeight                       = 9
	TextAttributeRise                               = 10
	TextAttributeUnderline                          = 11
	TextAttributeStrikethrough                      = 12
	TextAttributeSize                               = 13
	TextAttributeScale                              = 14
	TextAttributeWeight                             = 15
	TextAttributeLanguage                           = 16
	TextAttributeFamilyName                         = 17
	TextAttributeBgColor                            = 18
	TextAttributeFgColor                            = 19
	TextAttributeBgStipple                          = 20
	TextAttributeFgStipple                          = 21
	TextAttributeWrapMode                           = 22
	TextAttributeDirection                          = 23
	TextAttributeJustification                      = 24
	TextAttributeStretch                            = 25
	TextAttributeVariant                            = 26
	TextAttributeStyle                              = 27
	TextAttributeLastDefined                        = 28
)

type TextBoundaryEnum int

const (
	TextBoundaryChar          TextBoundaryEnum = 0
	TextBoundaryWordStart                      = 1
	TextBoundaryWordEnd                        = 2
	TextBoundarySentenceStart                  = 3
	TextBoundarySentenceEnd                    = 4
	TextBoundaryLineStart                      = 5
	TextBoundaryLineEnd                        = 6
)

type TextClipTypeEnum int

const (
	TextClipTypeNone TextClipTypeEnum = 0
	TextClipTypeMin                   = 1
	TextClipTypeMax                   = 2
	TextClipTypeBoth                  = 3
)

type TextGranularityEnum int

const (
	TextGranularityChar      TextGranularityEnum = 0
	TextGranularityWord                          = 1
	TextGranularitySentence                      = 2
	TextGranularityLine                          = 3
	TextGranularityParagraph                     = 4
)

// ignore GType struct TextIface
// Struct TextRange
type TextRange struct {
	P unsafe.Pointer
}

// Struct TextRectangle
type TextRectangle struct {
	P unsafe.Pointer
}

// Object Util
type Util struct {
	gobject.Object
}

// ignore GType struct UtilClass
// Interface Value
type Value struct {
	ValueIfc
	P unsafe.Pointer
}
type ValueIfc struct{}

// atk_value_get_current_value
// container is not nil, container is Value
// is method
func (v *ValueIfc) GetCurrentValue() (value int /*TODO_TYPE*/) {
	iv, err := _I.Get(196, "Value", "get_current_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, &outArgs[0])
	value = outArgs[0].Int() /*TODO*/
	return
}

// atk_value_get_increment
// container is not nil, container is Value
// is method
func (v *ValueIfc) GetIncrement() (result float64) {
	iv, err := _I.Get(197, "Value", "get_increment")
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

// atk_value_get_maximum_value
// container is not nil, container is Value
// is method
func (v *ValueIfc) GetMaximumValue() (value int /*TODO_TYPE*/) {
	iv, err := _I.Get(198, "Value", "get_maximum_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, &outArgs[0])
	value = outArgs[0].Int() /*TODO*/
	return
}

// atk_value_get_minimum_increment
// container is not nil, container is Value
// is method
func (v *ValueIfc) GetMinimumIncrement() (value int /*TODO_TYPE*/) {
	iv, err := _I.Get(199, "Value", "get_minimum_increment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, &outArgs[0])
	value = outArgs[0].Int() /*TODO*/
	return
}

// atk_value_get_minimum_value
// container is not nil, container is Value
// is method
func (v *ValueIfc) GetMinimumValue() (value int /*TODO_TYPE*/) {
	iv, err := _I.Get(200, "Value", "get_minimum_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, &outArgs[0])
	value = outArgs[0].Int() /*TODO*/
	return
}

// atk_value_get_range
// container is not nil, container is Value
// is method
func (v *ValueIfc) GetRange() (result Range) {
	iv, err := _I.Get(201, "Value", "get_range")
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

// atk_value_get_sub_ranges
// container is not nil, container is Value
// is method
func (v *ValueIfc) GetSubRanges() (result int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(202, "Value", "get_sub_ranges")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// atk_value_get_value_and_text
// container is not nil, container is Value
// is method
func (v *ValueIfc) GetValueAndText() (value float64, text string) {
	iv, err := _I.Get(203, "Value", "get_value_and_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_text := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_value, arg_text}
	iv.Call(args, nil, &outArgs[0])
	value = outArgs[0].Double()
	text = outArgs[1].String().Take()
	return
}

// atk_value_set_current_value
// container is not nil, container is Value
// is method
func (v *ValueIfc) SetCurrentValue(value gobject.Value) (result bool) {
	iv, err := _I.Get(204, "Value", "set_current_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// atk_value_set_value
// container is not nil, container is Value
// is method
func (v *ValueIfc) SetValue(new_value float64) {
	iv, err := _I.Get(205, "Value", "set_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_new_value := gi.NewDoubleArgument(new_value)
	args := []gi.Argument{arg_v, arg_new_value}
	iv.Call(args, nil, nil)
}

// ignore GType struct ValueIface
type ValueTypeEnum int

const (
	ValueTypeVeryWeak    ValueTypeEnum = 0
	ValueTypeWeak                      = 1
	ValueTypeAcceptable                = 2
	ValueTypeStrong                    = 3
	ValueTypeVeryStrong                = 4
	ValueTypeVeryLow                   = 5
	ValueTypeLow                       = 6
	ValueTypeMedium                    = 7
	ValueTypeHigh                      = 8
	ValueTypeVeryHigh                  = 9
	ValueTypeVeryBad                   = 10
	ValueTypeBad                       = 11
	ValueTypeGood                      = 12
	ValueTypeVeryGood                  = 13
	ValueTypeBest                      = 14
	ValueTypeLastDefined               = 15
)

// Interface Window
type Window struct {
	WindowIfc
	P unsafe.Pointer
}
type WindowIfc struct{}

// ignore GType struct WindowIface
// atk_attribute_set_free
// container is nil
func AttributeSetFree(attrib_set int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(206, "attribute_set_free", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_attrib_set := gi.NewIntArgument(attrib_set) /*TODO*/
	args := []gi.Argument{arg_attrib_set}
	iv.Call(args, nil, nil)
}

// atk_focus_tracker_notify
// container is nil
func FocusTrackerNotify(object Object) {
	iv, err := _I.Get(207, "focus_tracker_notify", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_object := gi.NewPointerArgument(object.P)
	args := []gi.Argument{arg_object}
	iv.Call(args, nil, nil)
}

// atk_get_binary_age
// container is nil
func GetBinaryAge() (result uint32) {
	iv, err := _I.Get(208, "get_binary_age", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// atk_get_default_registry
// container is nil
func GetDefaultRegistry() (result Registry) {
	iv, err := _I.Get(209, "get_default_registry", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_get_focus_object
// container is nil
func GetFocusObject() (result Object) {
	iv, err := _I.Get(210, "get_focus_object", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_get_interface_age
// container is nil
func GetInterfaceAge() (result uint32) {
	iv, err := _I.Get(211, "get_interface_age", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// atk_get_major_version
// container is nil
func GetMajorVersion() (result uint32) {
	iv, err := _I.Get(212, "get_major_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// atk_get_micro_version
// container is nil
func GetMicroVersion() (result uint32) {
	iv, err := _I.Get(213, "get_micro_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// atk_get_minor_version
// container is nil
func GetMinorVersion() (result uint32) {
	iv, err := _I.Get(214, "get_minor_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// atk_get_root
// container is nil
func GetRoot() (result Object) {
	iv, err := _I.Get(215, "get_root", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// atk_get_toolkit_name
// container is nil
func GetToolkitName() (result string) {
	iv, err := _I.Get(216, "get_toolkit_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_get_toolkit_version
// container is nil
func GetToolkitVersion() (result string) {
	iv, err := _I.Get(217, "get_toolkit_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_get_version
// container is nil
func GetVersion() (result string) {
	iv, err := _I.Get(218, "get_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_relation_type_for_name
// container is nil
func RelationTypeForName(name string) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(219, "relation_type_for_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	gi.Free(c_name)
	return
}

// atk_relation_type_get_name
// container is nil
func RelationTypeGetName(type1 int /*TODO_TYPE isPtr: false, tag: interface*/) (result string) {
	iv, err := _I.Get(220, "relation_type_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(type1) /*TODO*/
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_relation_type_register
// container is nil
func RelationTypeRegister(name string) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(221, "relation_type_register", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	gi.Free(c_name)
	return
}

// atk_remove_focus_tracker
// container is nil
func RemoveFocusTracker(tracker_id uint32) {
	iv, err := _I.Get(222, "remove_focus_tracker", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_tracker_id := gi.NewUint32Argument(tracker_id)
	args := []gi.Argument{arg_tracker_id}
	iv.Call(args, nil, nil)
}

// atk_remove_global_event_listener
// container is nil
func RemoveGlobalEventListener(listener_id uint32) {
	iv, err := _I.Get(223, "remove_global_event_listener", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_listener_id := gi.NewUint32Argument(listener_id)
	args := []gi.Argument{arg_listener_id}
	iv.Call(args, nil, nil)
}

// atk_remove_key_event_listener
// container is nil
func RemoveKeyEventListener(listener_id uint32) {
	iv, err := _I.Get(224, "remove_key_event_listener", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_listener_id := gi.NewUint32Argument(listener_id)
	args := []gi.Argument{arg_listener_id}
	iv.Call(args, nil, nil)
}

// atk_role_for_name
// container is nil
func RoleForName(name string) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(225, "role_for_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	gi.Free(c_name)
	return
}

// atk_role_get_localized_name
// container is nil
func RoleGetLocalizedName(role int /*TODO_TYPE isPtr: false, tag: interface*/) (result string) {
	iv, err := _I.Get(226, "role_get_localized_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_role := gi.NewIntArgument(role) /*TODO*/
	args := []gi.Argument{arg_role}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_role_get_name
// container is nil
func RoleGetName(role int /*TODO_TYPE isPtr: false, tag: interface*/) (result string) {
	iv, err := _I.Get(227, "role_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_role := gi.NewIntArgument(role) /*TODO*/
	args := []gi.Argument{arg_role}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_role_register
// container is nil
func RoleRegister(name string) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(228, "role_register", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	gi.Free(c_name)
	return
}

// atk_state_type_for_name
// container is nil
func StateTypeForName(name string) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(229, "state_type_for_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	gi.Free(c_name)
	return
}

// atk_state_type_get_name
// container is nil
func StateTypeGetName(type1 int /*TODO_TYPE isPtr: false, tag: interface*/) (result string) {
	iv, err := _I.Get(230, "state_type_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(type1) /*TODO*/
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_state_type_register
// container is nil
func StateTypeRegister(name string) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(231, "state_type_register", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	gi.Free(c_name)
	return
}

// atk_text_attribute_for_name
// container is nil
func TextAttributeForName(name string) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(232, "text_attribute_for_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	gi.Free(c_name)
	return
}

// atk_text_attribute_get_name
// container is nil
func TextAttributeGetName(attr int /*TODO_TYPE isPtr: false, tag: interface*/) (result string) {
	iv, err := _I.Get(233, "text_attribute_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_attr := gi.NewIntArgument(attr) /*TODO*/
	args := []gi.Argument{arg_attr}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_text_attribute_get_value
// container is nil
func TextAttributeGetValue(attr int /*TODO_TYPE isPtr: false, tag: interface*/, index_ int32) (result string) {
	iv, err := _I.Get(234, "text_attribute_get_value", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_attr := gi.NewIntArgument(attr) /*TODO*/
	arg_index_ := gi.NewInt32Argument(index_)
	args := []gi.Argument{arg_attr, arg_index_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_text_attribute_register
// container is nil
func TextAttributeRegister(name string) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(235, "text_attribute_register", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	gi.Free(c_name)
	return
}

// atk_text_free_ranges
// container is nil
func TextFreeRanges(ranges int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(236, "text_free_ranges", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ranges := gi.NewIntArgument(ranges) /*TODO*/
	args := []gi.Argument{arg_ranges}
	iv.Call(args, nil, nil)
}

// atk_value_type_get_localized_name
// container is nil
func ValueTypeGetLocalizedName(value_type int /*TODO_TYPE isPtr: false, tag: interface*/) (result string) {
	iv, err := _I.Get(237, "value_type_get_localized_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value_type := gi.NewIntArgument(value_type) /*TODO*/
	args := []gi.Argument{arg_value_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// atk_value_type_get_name
// container is nil
func ValueTypeGetName(value_type int /*TODO_TYPE isPtr: false, tag: interface*/) (result string) {
	iv, err := _I.Get(238, "value_type_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value_type := gi.NewIntArgument(value_type) /*TODO*/
	args := []gi.Argument{arg_value_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// constants
const (
	BINARY_AGE           = 23010
	INTERFACE_AGE        = 1
	MAJOR_VERSION        = 2
	MICRO_VERSION        = 0
	MINOR_VERSION        = 30
	VERSION_MIN_REQUIRED = 2
)
