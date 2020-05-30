package pango

import "github.com/electricface/go-gir/glib-2.0"
import "github.com/electricface/go-gir/gobject-2.0"
import "github.com/electricface/go-gir3/gi"
import "log"
import "unsafe"

var _I = gi.NewInvokerCache("Pango")
var _ unsafe.Pointer

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("Pango", "1.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

type AlignmentEnum int

const (
	AlignmentLeft   AlignmentEnum = 0
	AlignmentCenter               = 1
	AlignmentRight                = 2
)

// Struct Analysis
type Analysis struct {
	P unsafe.Pointer
}

// Struct AttrClass
type AttrClass struct {
	P unsafe.Pointer
}

// Struct AttrColor
type AttrColor struct {
	P unsafe.Pointer
}

// Struct AttrFloat
type AttrFloat struct {
	P unsafe.Pointer
}

// Struct AttrFontDesc
type AttrFontDesc struct {
	P unsafe.Pointer
}

// Struct AttrFontFeatures
type AttrFontFeatures struct {
	P unsafe.Pointer
}

// Struct AttrInt
type AttrInt struct {
	P unsafe.Pointer
}

// Struct AttrIterator
type AttrIterator struct {
	P unsafe.Pointer
}

// black function AttrIterator.destroy

// pango_attr_iterator_get_attrs
// container is not nil, container is AttrIterator
// is method
func (v AttrIterator) GetAttrs() (result int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(0, "AttrIterator", "get_attrs")
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

// pango_attr_iterator_get_font
// container is not nil, container is AttrIterator
// is method
func (v AttrIterator) GetFont(desc FontDescription, language Language, extra_attrs int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(1, "AttrIterator", "get_font")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_desc := gi.NewPointerArgument(desc.P)
	arg_language := gi.NewPointerArgument(language.P)
	arg_extra_attrs := gi.NewIntArgument(extra_attrs) /*TODO*/
	args := []gi.Argument{arg_v, arg_desc, arg_language, arg_extra_attrs}
	iv.Call(args, nil, nil)
}

// pango_attr_iterator_next
// container is not nil, container is AttrIterator
// is method
func (v AttrIterator) Next() (result bool) {
	iv, err := _I.Get(2, "AttrIterator", "next")
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

// pango_attr_iterator_range
// container is not nil, container is AttrIterator
// is method
func (v AttrIterator) Range() (start int32, end int32) {
	iv, err := _I.Get(3, "AttrIterator", "range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_start, arg_end}
	iv.Call(args, nil, &outArgs[0])
	start = outArgs[0].Int32()
	end = outArgs[1].Int32()
	return
}

// Struct AttrLanguage
type AttrLanguage struct {
	P unsafe.Pointer
}

// Struct AttrList
type AttrList struct {
	P unsafe.Pointer
}

// pango_attr_list_new
// container is not nil, container is AttrList
// is constructor
func NewAttrList() (result AttrList) {
	iv, err := _I.Get(4, "AttrList", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_attr_list_change
// container is not nil, container is AttrList
// is method
func (v AttrList) Change(attr Attribute) {
	iv, err := _I.Get(5, "AttrList", "change")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_attr := gi.NewPointerArgument(attr.P)
	args := []gi.Argument{arg_v, arg_attr}
	iv.Call(args, nil, nil)
}

// pango_attr_list_copy
// container is not nil, container is AttrList
// is method
func (v AttrList) Copy() (result AttrList) {
	iv, err := _I.Get(6, "AttrList", "copy")
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

// pango_attr_list_filter
// container is not nil, container is AttrList
// is method
func (v AttrList) Filter(func1 int /*TODO_TYPE isPtr: false, tag: interface*/, data unsafe.Pointer) (result AttrList) {
	iv, err := _I.Get(7, "AttrList", "filter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewIntArgument(func1) /*TODO*/
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_func1, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_attr_list_insert
// container is not nil, container is AttrList
// is method
func (v AttrList) Insert(attr Attribute) {
	iv, err := _I.Get(8, "AttrList", "insert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_attr := gi.NewPointerArgument(attr.P)
	args := []gi.Argument{arg_v, arg_attr}
	iv.Call(args, nil, nil)
}

// pango_attr_list_insert_before
// container is not nil, container is AttrList
// is method
func (v AttrList) InsertBefore(attr Attribute) {
	iv, err := _I.Get(9, "AttrList", "insert_before")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_attr := gi.NewPointerArgument(attr.P)
	args := []gi.Argument{arg_v, arg_attr}
	iv.Call(args, nil, nil)
}

// pango_attr_list_ref
// container is not nil, container is AttrList
// is method
func (v AttrList) Ref() (result AttrList) {
	iv, err := _I.Get(10, "AttrList", "ref")
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

// pango_attr_list_splice
// container is not nil, container is AttrList
// is method
func (v AttrList) Splice(other AttrList, pos int32, len1 int32) {
	iv, err := _I.Get(11, "AttrList", "splice")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_other := gi.NewPointerArgument(other.P)
	arg_pos := gi.NewInt32Argument(pos)
	arg_len1 := gi.NewInt32Argument(len1)
	args := []gi.Argument{arg_v, arg_other, arg_pos, arg_len1}
	iv.Call(args, nil, nil)
}

// pango_attr_list_unref
// container is not nil, container is AttrList
// is method
func (v AttrList) Unref() {
	iv, err := _I.Get(12, "AttrList", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct AttrShape
type AttrShape struct {
	P unsafe.Pointer
}

// Struct AttrSize
type AttrSize struct {
	P unsafe.Pointer
}

// Struct AttrString
type AttrString struct {
	P unsafe.Pointer
}
type AttrTypeEnum int

const (
	AttrTypeInvalid            AttrTypeEnum = 0
	AttrTypeLanguage                        = 1
	AttrTypeFamily                          = 2
	AttrTypeStyle                           = 3
	AttrTypeWeight                          = 4
	AttrTypeVariant                         = 5
	AttrTypeStretch                         = 6
	AttrTypeSize                            = 7
	AttrTypeFontDesc                        = 8
	AttrTypeForeground                      = 9
	AttrTypeBackground                      = 10
	AttrTypeUnderline                       = 11
	AttrTypeStrikethrough                   = 12
	AttrTypeRise                            = 13
	AttrTypeShape                           = 14
	AttrTypeScale                           = 15
	AttrTypeFallback                        = 16
	AttrTypeLetterSpacing                   = 17
	AttrTypeUnderlineColor                  = 18
	AttrTypeStrikethroughColor              = 19
	AttrTypeAbsoluteSize                    = 20
	AttrTypeGravity                         = 21
	AttrTypeGravityHint                     = 22
	AttrTypeFontFeatures                    = 23
	AttrTypeForegroundAlpha                 = 24
	AttrTypeBackgroundAlpha                 = 25
)

// Struct Attribute
type Attribute struct {
	P unsafe.Pointer
}

// pango_attribute_destroy
// container is not nil, container is Attribute
// is method
func (v Attribute) Destroy() {
	iv, err := _I.Get(13, "Attribute", "destroy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_attribute_equal
// container is not nil, container is Attribute
// is method
func (v Attribute) Equal(attr2 Attribute) (result bool) {
	iv, err := _I.Get(14, "Attribute", "equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_attr2 := gi.NewPointerArgument(attr2.P)
	args := []gi.Argument{arg_v, arg_attr2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// pango_attribute_init
// container is not nil, container is Attribute
// is method
func (v Attribute) Init(klass AttrClass) {
	iv, err := _I.Get(15, "Attribute", "init")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_klass := gi.NewPointerArgument(klass.P)
	args := []gi.Argument{arg_v, arg_klass}
	iv.Call(args, nil, nil)
}

type BidiTypeEnum int

const (
	BidiTypeL   BidiTypeEnum = 0
	BidiTypeLre              = 1
	BidiTypeLro              = 2
	BidiTypeR                = 3
	BidiTypeAl               = 4
	BidiTypeRle              = 5
	BidiTypeRlo              = 6
	BidiTypePdf              = 7
	BidiTypeEn               = 8
	BidiTypeEs               = 9
	BidiTypeEt               = 10
	BidiTypeAn               = 11
	BidiTypeCs               = 12
	BidiTypeNsm              = 13
	BidiTypeBn               = 14
	BidiTypeB                = 15
	BidiTypeS                = 16
	BidiTypeWs               = 17
	BidiTypeOn               = 18
)

// Struct Color
type Color struct {
	P unsafe.Pointer
}

// pango_color_copy
// container is not nil, container is Color
// is method
func (v Color) Copy() (result Color) {
	iv, err := _I.Get(16, "Color", "copy")
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

// pango_color_free
// container is not nil, container is Color
// is method
func (v Color) Free() {
	iv, err := _I.Get(17, "Color", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_color_parse
// container is not nil, container is Color
// is method
func (v Color) Parse(spec string) (result bool) {
	iv, err := _I.Get(18, "Color", "parse")
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
	result = ret.Bool()
	gi.Free(c_spec)
	return
}

// pango_color_to_string
// container is not nil, container is Color
// is method
func (v Color) ToString() (result string) {
	iv, err := _I.Get(19, "Color", "to_string")
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

// Object Context
type Context struct {
	gobject.Object
}

func WrapContext(p unsafe.Pointer) (r Context) { r.P = p; return }

// pango_context_new
// container is not nil, container is Context
// is constructor
func NewContext() (result Context) {
	iv, err := _I.Get(20, "Context", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_context_changed
// container is not nil, container is Context
// is method
func (v Context) Changed() {
	iv, err := _I.Get(21, "Context", "changed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_context_get_base_dir
// container is not nil, container is Context
// is method
func (v Context) GetBaseDir() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(22, "Context", "get_base_dir")
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

// pango_context_get_base_gravity
// container is not nil, container is Context
// is method
func (v Context) GetBaseGravity() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(23, "Context", "get_base_gravity")
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

// pango_context_get_font_description
// container is not nil, container is Context
// is method
func (v Context) GetFontDescription() (result FontDescription) {
	iv, err := _I.Get(24, "Context", "get_font_description")
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

// pango_context_get_font_map
// container is not nil, container is Context
// is method
func (v Context) GetFontMap() (result FontMap) {
	iv, err := _I.Get(25, "Context", "get_font_map")
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

// pango_context_get_gravity
// container is not nil, container is Context
// is method
func (v Context) GetGravity() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(26, "Context", "get_gravity")
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

// pango_context_get_gravity_hint
// container is not nil, container is Context
// is method
func (v Context) GetGravityHint() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(27, "Context", "get_gravity_hint")
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

// pango_context_get_language
// container is not nil, container is Context
// is method
func (v Context) GetLanguage() (result Language) {
	iv, err := _I.Get(28, "Context", "get_language")
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

// pango_context_get_matrix
// container is not nil, container is Context
// is method
func (v Context) GetMatrix() (result Matrix) {
	iv, err := _I.Get(29, "Context", "get_matrix")
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

// pango_context_get_metrics
// container is not nil, container is Context
// is method
func (v Context) GetMetrics(desc FontDescription, language Language) (result FontMetrics) {
	iv, err := _I.Get(30, "Context", "get_metrics")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_desc := gi.NewPointerArgument(desc.P)
	arg_language := gi.NewPointerArgument(language.P)
	args := []gi.Argument{arg_v, arg_desc, arg_language}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_context_get_serial
// container is not nil, container is Context
// is method
func (v Context) GetSerial() (result uint32) {
	iv, err := _I.Get(31, "Context", "get_serial")
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

// pango_context_list_families
// container is not nil, container is Context
// is method
func (v Context) ListFamilies() (families int /*TODO_TYPE*/, n_families int32) {
	iv, err := _I.Get(32, "Context", "list_families")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_families := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_n_families := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_families, arg_n_families}
	iv.Call(args, nil, &outArgs[0])
	families = outArgs[0].Int() /*TODO*/
	n_families = outArgs[1].Int32()
	return
}

// pango_context_load_font
// container is not nil, container is Context
// is method
func (v Context) LoadFont(desc FontDescription) (result Font) {
	iv, err := _I.Get(33, "Context", "load_font")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_desc := gi.NewPointerArgument(desc.P)
	args := []gi.Argument{arg_v, arg_desc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_context_load_fontset
// container is not nil, container is Context
// is method
func (v Context) LoadFontset(desc FontDescription, language Language) (result Fontset) {
	iv, err := _I.Get(34, "Context", "load_fontset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_desc := gi.NewPointerArgument(desc.P)
	arg_language := gi.NewPointerArgument(language.P)
	args := []gi.Argument{arg_v, arg_desc, arg_language}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_context_set_base_dir
// container is not nil, container is Context
// is method
func (v Context) SetBaseDir(direction int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(35, "Context", "set_base_dir")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_direction := gi.NewIntArgument(direction) /*TODO*/
	args := []gi.Argument{arg_v, arg_direction}
	iv.Call(args, nil, nil)
}

// pango_context_set_base_gravity
// container is not nil, container is Context
// is method
func (v Context) SetBaseGravity(gravity int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(36, "Context", "set_base_gravity")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_gravity := gi.NewIntArgument(gravity) /*TODO*/
	args := []gi.Argument{arg_v, arg_gravity}
	iv.Call(args, nil, nil)
}

// pango_context_set_font_description
// container is not nil, container is Context
// is method
func (v Context) SetFontDescription(desc FontDescription) {
	iv, err := _I.Get(37, "Context", "set_font_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_desc := gi.NewPointerArgument(desc.P)
	args := []gi.Argument{arg_v, arg_desc}
	iv.Call(args, nil, nil)
}

// pango_context_set_font_map
// container is not nil, container is Context
// is method
func (v Context) SetFontMap(font_map FontMap) {
	iv, err := _I.Get(38, "Context", "set_font_map")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_font_map := gi.NewPointerArgument(font_map.P)
	args := []gi.Argument{arg_v, arg_font_map}
	iv.Call(args, nil, nil)
}

// pango_context_set_gravity_hint
// container is not nil, container is Context
// is method
func (v Context) SetGravityHint(hint int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(39, "Context", "set_gravity_hint")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_hint := gi.NewIntArgument(hint) /*TODO*/
	args := []gi.Argument{arg_v, arg_hint}
	iv.Call(args, nil, nil)
}

// pango_context_set_language
// container is not nil, container is Context
// is method
func (v Context) SetLanguage(language Language) {
	iv, err := _I.Get(40, "Context", "set_language")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_language := gi.NewPointerArgument(language.P)
	args := []gi.Argument{arg_v, arg_language}
	iv.Call(args, nil, nil)
}

// pango_context_set_matrix
// container is not nil, container is Context
// is method
func (v Context) SetMatrix(matrix Matrix) {
	iv, err := _I.Get(41, "Context", "set_matrix")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_matrix := gi.NewPointerArgument(matrix.P)
	args := []gi.Argument{arg_v, arg_matrix}
	iv.Call(args, nil, nil)
}

// ignore GType struct ContextClass
// Struct Coverage
type Coverage struct {
	P unsafe.Pointer
}

// pango_coverage_get
// container is not nil, container is Coverage
// is method
func (v Coverage) Get(index_ int32) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(42, "Coverage", "get")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewInt32Argument(index_)
	args := []gi.Argument{arg_v, arg_index_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// pango_coverage_max
// container is not nil, container is Coverage
// is method
func (v Coverage) Max(other Coverage) {
	iv, err := _I.Get(43, "Coverage", "max")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_other := gi.NewPointerArgument(other.P)
	args := []gi.Argument{arg_v, arg_other}
	iv.Call(args, nil, nil)
}

// pango_coverage_set
// container is not nil, container is Coverage
// is method
func (v Coverage) Set(index_ int32, level int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(44, "Coverage", "set")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewInt32Argument(index_)
	arg_level := gi.NewIntArgument(level) /*TODO*/
	args := []gi.Argument{arg_v, arg_index_, arg_level}
	iv.Call(args, nil, nil)
}

// pango_coverage_to_bytes
// container is not nil, container is Coverage
// is method
func (v Coverage) ToBytes() (bytes int /*TODO_TYPE*/, n_bytes int32) {
	iv, err := _I.Get(45, "Coverage", "to_bytes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_bytes := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_n_bytes := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_bytes, arg_n_bytes}
	iv.Call(args, nil, &outArgs[0])
	bytes = outArgs[0].Int() /*TODO*/
	n_bytes = outArgs[1].Int32()
	return
}

// pango_coverage_unref
// container is not nil, container is Coverage
// is method
func (v Coverage) Unref() {
	iv, err := _I.Get(46, "Coverage", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

type CoverageLevelEnum int

const (
	CoverageLevelNone        CoverageLevelEnum = 0
	CoverageLevelFallback                      = 1
	CoverageLevelApproximate                   = 2
	CoverageLevelExact                         = 3
)

type DirectionEnum int

const (
	DirectionLtr     DirectionEnum = 0
	DirectionRtl                   = 1
	DirectionTtbLtr                = 2
	DirectionTtbRtl                = 3
	DirectionWeakLtr               = 4
	DirectionWeakRtl               = 5
	DirectionNeutral               = 6
)

type EllipsizeModeEnum int

const (
	EllipsizeModeNone   EllipsizeModeEnum = 0
	EllipsizeModeStart                    = 1
	EllipsizeModeMiddle                   = 2
	EllipsizeModeEnd                      = 3
)

// Object Engine
type Engine struct {
	gobject.Object
}

func WrapEngine(p unsafe.Pointer) (r Engine) { r.P = p; return }

// ignore GType struct EngineClass
// Struct EngineInfo
type EngineInfo struct {
	P unsafe.Pointer
}

// Object EngineLang
type EngineLang struct {
	Engine
}

func WrapEngineLang(p unsafe.Pointer) (r EngineLang) { r.P = p; return }

// ignore GType struct EngineLangClass
// Struct EngineScriptInfo
type EngineScriptInfo struct {
	P unsafe.Pointer
}

// Object EngineShape
type EngineShape struct {
	Engine
}

func WrapEngineShape(p unsafe.Pointer) (r EngineShape) { r.P = p; return }

// ignore GType struct EngineShapeClass
// Object Font
type Font struct {
	gobject.Object
}

func WrapFont(p unsafe.Pointer) (r Font) { r.P = p; return }

// pango_font_descriptions_free
// container is not nil, container is Font
// is method
// arg0Type tag: array, isPtr: true
func FontDescriptionsFree1(descs int /*TODO_TYPE isPtr: true, tag: array*/, n_descs int32) {
	iv, err := _I.Get(47, "Font", "descriptions_free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_descs := gi.NewIntArgument(descs) /*TODO*/
	arg_n_descs := gi.NewInt32Argument(n_descs)
	args := []gi.Argument{arg_descs, arg_n_descs}
	iv.Call(args, nil, nil)
}

// pango_font_describe
// container is not nil, container is Font
// is method
func (v Font) Describe() (result FontDescription) {
	iv, err := _I.Get(48, "Font", "describe")
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

// pango_font_describe_with_absolute_size
// container is not nil, container is Font
// is method
func (v Font) DescribeWithAbsoluteSize() (result FontDescription) {
	iv, err := _I.Get(49, "Font", "describe_with_absolute_size")
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

// pango_font_find_shaper
// container is not nil, container is Font
// is method
func (v Font) FindShaper(language Language, ch uint32) (result EngineShape) {
	iv, err := _I.Get(50, "Font", "find_shaper")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_language := gi.NewPointerArgument(language.P)
	arg_ch := gi.NewUint32Argument(ch)
	args := []gi.Argument{arg_v, arg_language, arg_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_font_get_font_map
// container is not nil, container is Font
// is method
func (v Font) GetFontMap() (result FontMap) {
	iv, err := _I.Get(51, "Font", "get_font_map")
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

// pango_font_get_glyph_extents
// container is not nil, container is Font
// is method
func (v Font) GetGlyphExtents(glyph uint32) (ink_rect int /*TODO_TYPE*/, logical_rect int /*TODO_TYPE*/) {
	iv, err := _I.Get(52, "Font", "get_glyph_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_glyph := gi.NewUint32Argument(glyph)
	arg_ink_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_logical_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_glyph, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, &outArgs[0])
	ink_rect = outArgs[0].Int()     /*TODO*/
	logical_rect = outArgs[1].Int() /*TODO*/
	return
}

// pango_font_get_metrics
// container is not nil, container is Font
// is method
func (v Font) GetMetrics(language Language) (result FontMetrics) {
	iv, err := _I.Get(53, "Font", "get_metrics")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_language := gi.NewPointerArgument(language.P)
	args := []gi.Argument{arg_v, arg_language}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct FontClass
// Struct FontDescription
type FontDescription struct {
	P unsafe.Pointer
}

// pango_font_description_new
// container is not nil, container is FontDescription
// is constructor
func NewFontDescription() (result FontDescription) {
	iv, err := _I.Get(54, "FontDescription", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_font_description_better_match
// container is not nil, container is FontDescription
// is method
func (v FontDescription) BetterMatch(old_match FontDescription, new_match FontDescription) (result bool) {
	iv, err := _I.Get(55, "FontDescription", "better_match")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_old_match := gi.NewPointerArgument(old_match.P)
	arg_new_match := gi.NewPointerArgument(new_match.P)
	args := []gi.Argument{arg_v, arg_old_match, arg_new_match}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// pango_font_description_copy
// container is not nil, container is FontDescription
// is method
func (v FontDescription) Copy() (result FontDescription) {
	iv, err := _I.Get(56, "FontDescription", "copy")
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

// pango_font_description_copy_static
// container is not nil, container is FontDescription
// is method
func (v FontDescription) CopyStatic() (result FontDescription) {
	iv, err := _I.Get(57, "FontDescription", "copy_static")
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

// pango_font_description_equal
// container is not nil, container is FontDescription
// is method
func (v FontDescription) Equal(desc2 FontDescription) (result bool) {
	iv, err := _I.Get(58, "FontDescription", "equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_desc2 := gi.NewPointerArgument(desc2.P)
	args := []gi.Argument{arg_v, arg_desc2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// pango_font_description_free
// container is not nil, container is FontDescription
// is method
func (v FontDescription) Free() {
	iv, err := _I.Get(59, "FontDescription", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_font_description_get_family
// container is not nil, container is FontDescription
// is method
func (v FontDescription) GetFamily() (result string) {
	iv, err := _I.Get(60, "FontDescription", "get_family")
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

// pango_font_description_get_gravity
// container is not nil, container is FontDescription
// is method
func (v FontDescription) GetGravity() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(61, "FontDescription", "get_gravity")
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

// pango_font_description_get_set_fields
// container is not nil, container is FontDescription
// is method
func (v FontDescription) GetSetFields() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(62, "FontDescription", "get_set_fields")
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

// pango_font_description_get_size
// container is not nil, container is FontDescription
// is method
func (v FontDescription) GetSize() (result int32) {
	iv, err := _I.Get(63, "FontDescription", "get_size")
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

// pango_font_description_get_size_is_absolute
// container is not nil, container is FontDescription
// is method
func (v FontDescription) GetSizeIsAbsolute() (result bool) {
	iv, err := _I.Get(64, "FontDescription", "get_size_is_absolute")
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

// pango_font_description_get_stretch
// container is not nil, container is FontDescription
// is method
func (v FontDescription) GetStretch() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(65, "FontDescription", "get_stretch")
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

// pango_font_description_get_style
// container is not nil, container is FontDescription
// is method
func (v FontDescription) GetStyle() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(66, "FontDescription", "get_style")
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

// pango_font_description_get_variant
// container is not nil, container is FontDescription
// is method
func (v FontDescription) GetVariant() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(67, "FontDescription", "get_variant")
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

// pango_font_description_get_variations
// container is not nil, container is FontDescription
// is method
func (v FontDescription) GetVariations() (result string) {
	iv, err := _I.Get(68, "FontDescription", "get_variations")
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

// pango_font_description_get_weight
// container is not nil, container is FontDescription
// is method
func (v FontDescription) GetWeight() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(69, "FontDescription", "get_weight")
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

// pango_font_description_hash
// container is not nil, container is FontDescription
// is method
func (v FontDescription) Hash() (result uint32) {
	iv, err := _I.Get(70, "FontDescription", "hash")
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

// pango_font_description_merge
// container is not nil, container is FontDescription
// is method
func (v FontDescription) Merge(desc_to_merge FontDescription, replace_existing bool) {
	iv, err := _I.Get(71, "FontDescription", "merge")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_desc_to_merge := gi.NewPointerArgument(desc_to_merge.P)
	arg_replace_existing := gi.NewBoolArgument(replace_existing)
	args := []gi.Argument{arg_v, arg_desc_to_merge, arg_replace_existing}
	iv.Call(args, nil, nil)
}

// pango_font_description_merge_static
// container is not nil, container is FontDescription
// is method
func (v FontDescription) MergeStatic(desc_to_merge FontDescription, replace_existing bool) {
	iv, err := _I.Get(72, "FontDescription", "merge_static")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_desc_to_merge := gi.NewPointerArgument(desc_to_merge.P)
	arg_replace_existing := gi.NewBoolArgument(replace_existing)
	args := []gi.Argument{arg_v, arg_desc_to_merge, arg_replace_existing}
	iv.Call(args, nil, nil)
}

// pango_font_description_set_absolute_size
// container is not nil, container is FontDescription
// is method
func (v FontDescription) SetAbsoluteSize(size float64) {
	iv, err := _I.Get(73, "FontDescription", "set_absolute_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewDoubleArgument(size)
	args := []gi.Argument{arg_v, arg_size}
	iv.Call(args, nil, nil)
}

// pango_font_description_set_family
// container is not nil, container is FontDescription
// is method
func (v FontDescription) SetFamily(family string) {
	iv, err := _I.Get(74, "FontDescription", "set_family")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_family := gi.CString(family)
	arg_v := gi.NewPointerArgument(v.P)
	arg_family := gi.NewStringArgument(c_family)
	args := []gi.Argument{arg_v, arg_family}
	iv.Call(args, nil, nil)
	gi.Free(c_family)
}

// pango_font_description_set_family_static
// container is not nil, container is FontDescription
// is method
func (v FontDescription) SetFamilyStatic(family string) {
	iv, err := _I.Get(75, "FontDescription", "set_family_static")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_family := gi.CString(family)
	arg_v := gi.NewPointerArgument(v.P)
	arg_family := gi.NewStringArgument(c_family)
	args := []gi.Argument{arg_v, arg_family}
	iv.Call(args, nil, nil)
	gi.Free(c_family)
}

// pango_font_description_set_gravity
// container is not nil, container is FontDescription
// is method
func (v FontDescription) SetGravity(gravity int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(76, "FontDescription", "set_gravity")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_gravity := gi.NewIntArgument(gravity) /*TODO*/
	args := []gi.Argument{arg_v, arg_gravity}
	iv.Call(args, nil, nil)
}

// pango_font_description_set_size
// container is not nil, container is FontDescription
// is method
func (v FontDescription) SetSize(size int32) {
	iv, err := _I.Get(77, "FontDescription", "set_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewInt32Argument(size)
	args := []gi.Argument{arg_v, arg_size}
	iv.Call(args, nil, nil)
}

// pango_font_description_set_stretch
// container is not nil, container is FontDescription
// is method
func (v FontDescription) SetStretch(stretch int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(78, "FontDescription", "set_stretch")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_stretch := gi.NewIntArgument(stretch) /*TODO*/
	args := []gi.Argument{arg_v, arg_stretch}
	iv.Call(args, nil, nil)
}

// pango_font_description_set_style
// container is not nil, container is FontDescription
// is method
func (v FontDescription) SetStyle(style int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(79, "FontDescription", "set_style")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_style := gi.NewIntArgument(style) /*TODO*/
	args := []gi.Argument{arg_v, arg_style}
	iv.Call(args, nil, nil)
}

// pango_font_description_set_variant
// container is not nil, container is FontDescription
// is method
func (v FontDescription) SetVariant(variant int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(80, "FontDescription", "set_variant")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_variant := gi.NewIntArgument(variant) /*TODO*/
	args := []gi.Argument{arg_v, arg_variant}
	iv.Call(args, nil, nil)
}

// pango_font_description_set_variations
// container is not nil, container is FontDescription
// is method
func (v FontDescription) SetVariations(settings string) {
	iv, err := _I.Get(81, "FontDescription", "set_variations")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_settings := gi.CString(settings)
	arg_v := gi.NewPointerArgument(v.P)
	arg_settings := gi.NewStringArgument(c_settings)
	args := []gi.Argument{arg_v, arg_settings}
	iv.Call(args, nil, nil)
	gi.Free(c_settings)
}

// pango_font_description_set_variations_static
// container is not nil, container is FontDescription
// is method
func (v FontDescription) SetVariationsStatic(settings string) {
	iv, err := _I.Get(82, "FontDescription", "set_variations_static")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_settings := gi.CString(settings)
	arg_v := gi.NewPointerArgument(v.P)
	arg_settings := gi.NewStringArgument(c_settings)
	args := []gi.Argument{arg_v, arg_settings}
	iv.Call(args, nil, nil)
	gi.Free(c_settings)
}

// pango_font_description_set_weight
// container is not nil, container is FontDescription
// is method
func (v FontDescription) SetWeight(weight int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(83, "FontDescription", "set_weight")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_weight := gi.NewIntArgument(weight) /*TODO*/
	args := []gi.Argument{arg_v, arg_weight}
	iv.Call(args, nil, nil)
}

// pango_font_description_to_filename
// container is not nil, container is FontDescription
// is method
func (v FontDescription) ToFilename() (result string) {
	iv, err := _I.Get(84, "FontDescription", "to_filename")
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

// pango_font_description_to_string
// container is not nil, container is FontDescription
// is method
func (v FontDescription) ToString() (result string) {
	iv, err := _I.Get(85, "FontDescription", "to_string")
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

// pango_font_description_unset_fields
// container is not nil, container is FontDescription
// is method
func (v FontDescription) UnsetFields(to_unset int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(86, "FontDescription", "unset_fields")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_to_unset := gi.NewIntArgument(to_unset) /*TODO*/
	args := []gi.Argument{arg_v, arg_to_unset}
	iv.Call(args, nil, nil)
}

// pango_font_description_from_string
// container is not nil, container is FontDescription
// is method
// arg0Type tag: utf8, isPtr: true
func FontDescriptionFromString1(str string) (result FontDescription) {
	iv, err := _I.Get(87, "FontDescription", "from_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	gi.Free(c_str)
	return
}

// Object FontFace
type FontFace struct {
	gobject.Object
}

func WrapFontFace(p unsafe.Pointer) (r FontFace) { r.P = p; return }

// pango_font_face_describe
// container is not nil, container is FontFace
// is method
func (v FontFace) Describe() (result FontDescription) {
	iv, err := _I.Get(88, "FontFace", "describe")
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

// pango_font_face_get_face_name
// container is not nil, container is FontFace
// is method
func (v FontFace) GetFaceName() (result string) {
	iv, err := _I.Get(89, "FontFace", "get_face_name")
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

// pango_font_face_is_synthesized
// container is not nil, container is FontFace
// is method
func (v FontFace) IsSynthesized() (result bool) {
	iv, err := _I.Get(90, "FontFace", "is_synthesized")
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

// pango_font_face_list_sizes
// container is not nil, container is FontFace
// is method
func (v FontFace) ListSizes() (sizes int /*TODO_TYPE*/, n_sizes int32) {
	iv, err := _I.Get(91, "FontFace", "list_sizes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_sizes := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_n_sizes := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_sizes, arg_n_sizes}
	iv.Call(args, nil, &outArgs[0])
	sizes = outArgs[0].Int() /*TODO*/
	n_sizes = outArgs[1].Int32()
	return
}

// ignore GType struct FontFaceClass
// Object FontFamily
type FontFamily struct {
	gobject.Object
}

func WrapFontFamily(p unsafe.Pointer) (r FontFamily) { r.P = p; return }

// pango_font_family_get_name
// container is not nil, container is FontFamily
// is method
func (v FontFamily) GetName() (result string) {
	iv, err := _I.Get(92, "FontFamily", "get_name")
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

// pango_font_family_is_monospace
// container is not nil, container is FontFamily
// is method
func (v FontFamily) IsMonospace() (result bool) {
	iv, err := _I.Get(93, "FontFamily", "is_monospace")
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

// pango_font_family_list_faces
// container is not nil, container is FontFamily
// is method
func (v FontFamily) ListFaces() (faces int /*TODO_TYPE*/, n_faces int32) {
	iv, err := _I.Get(94, "FontFamily", "list_faces")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_faces := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_n_faces := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_faces, arg_n_faces}
	iv.Call(args, nil, &outArgs[0])
	faces = outArgs[0].Int() /*TODO*/
	n_faces = outArgs[1].Int32()
	return
}

// ignore GType struct FontFamilyClass
// Object FontMap
type FontMap struct {
	gobject.Object
}

func WrapFontMap(p unsafe.Pointer) (r FontMap) { r.P = p; return }

// pango_font_map_changed
// container is not nil, container is FontMap
// is method
func (v FontMap) Changed() {
	iv, err := _I.Get(95, "FontMap", "changed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_font_map_create_context
// container is not nil, container is FontMap
// is method
func (v FontMap) CreateContext() (result Context) {
	iv, err := _I.Get(96, "FontMap", "create_context")
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

// pango_font_map_get_serial
// container is not nil, container is FontMap
// is method
func (v FontMap) GetSerial() (result uint32) {
	iv, err := _I.Get(97, "FontMap", "get_serial")
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

// pango_font_map_get_shape_engine_type
// container is not nil, container is FontMap
// is method
func (v FontMap) GetShapeEngineType() (result string) {
	iv, err := _I.Get(98, "FontMap", "get_shape_engine_type")
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

// pango_font_map_list_families
// container is not nil, container is FontMap
// is method
func (v FontMap) ListFamilies() (families int /*TODO_TYPE*/, n_families int32) {
	iv, err := _I.Get(99, "FontMap", "list_families")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_families := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_n_families := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_families, arg_n_families}
	iv.Call(args, nil, &outArgs[0])
	families = outArgs[0].Int() /*TODO*/
	n_families = outArgs[1].Int32()
	return
}

// pango_font_map_load_font
// container is not nil, container is FontMap
// is method
func (v FontMap) LoadFont(context Context, desc FontDescription) (result Font) {
	iv, err := _I.Get(100, "FontMap", "load_font")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_context := gi.NewPointerArgument(context.P)
	arg_desc := gi.NewPointerArgument(desc.P)
	args := []gi.Argument{arg_v, arg_context, arg_desc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_font_map_load_fontset
// container is not nil, container is FontMap
// is method
func (v FontMap) LoadFontset(context Context, desc FontDescription, language Language) (result Fontset) {
	iv, err := _I.Get(101, "FontMap", "load_fontset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_context := gi.NewPointerArgument(context.P)
	arg_desc := gi.NewPointerArgument(desc.P)
	arg_language := gi.NewPointerArgument(language.P)
	args := []gi.Argument{arg_v, arg_context, arg_desc, arg_language}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct FontMapClass
type FontMaskFlags int

const (
	FontMaskFamily     FontMaskFlags = 1
	FontMaskStyle                    = 2
	FontMaskVariant                  = 4
	FontMaskWeight                   = 8
	FontMaskStretch                  = 16
	FontMaskSize                     = 32
	FontMaskGravity                  = 64
	FontMaskVariations               = 128
)

// Struct FontMetrics
type FontMetrics struct {
	P unsafe.Pointer
}

// pango_font_metrics_new
// container is not nil, container is FontMetrics
// is constructor
func NewFontMetrics() (result FontMetrics) {
	iv, err := _I.Get(102, "FontMetrics", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_font_metrics_get_approximate_char_width
// container is not nil, container is FontMetrics
// is method
func (v FontMetrics) GetApproximateCharWidth() (result int32) {
	iv, err := _I.Get(103, "FontMetrics", "get_approximate_char_width")
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

// pango_font_metrics_get_approximate_digit_width
// container is not nil, container is FontMetrics
// is method
func (v FontMetrics) GetApproximateDigitWidth() (result int32) {
	iv, err := _I.Get(104, "FontMetrics", "get_approximate_digit_width")
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

// pango_font_metrics_get_ascent
// container is not nil, container is FontMetrics
// is method
func (v FontMetrics) GetAscent() (result int32) {
	iv, err := _I.Get(105, "FontMetrics", "get_ascent")
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

// pango_font_metrics_get_descent
// container is not nil, container is FontMetrics
// is method
func (v FontMetrics) GetDescent() (result int32) {
	iv, err := _I.Get(106, "FontMetrics", "get_descent")
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

// pango_font_metrics_get_strikethrough_position
// container is not nil, container is FontMetrics
// is method
func (v FontMetrics) GetStrikethroughPosition() (result int32) {
	iv, err := _I.Get(107, "FontMetrics", "get_strikethrough_position")
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

// pango_font_metrics_get_strikethrough_thickness
// container is not nil, container is FontMetrics
// is method
func (v FontMetrics) GetStrikethroughThickness() (result int32) {
	iv, err := _I.Get(108, "FontMetrics", "get_strikethrough_thickness")
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

// pango_font_metrics_get_underline_position
// container is not nil, container is FontMetrics
// is method
func (v FontMetrics) GetUnderlinePosition() (result int32) {
	iv, err := _I.Get(109, "FontMetrics", "get_underline_position")
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

// pango_font_metrics_get_underline_thickness
// container is not nil, container is FontMetrics
// is method
func (v FontMetrics) GetUnderlineThickness() (result int32) {
	iv, err := _I.Get(110, "FontMetrics", "get_underline_thickness")
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

// pango_font_metrics_ref
// container is not nil, container is FontMetrics
// is method
func (v FontMetrics) Ref() (result FontMetrics) {
	iv, err := _I.Get(111, "FontMetrics", "ref")
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

// pango_font_metrics_unref
// container is not nil, container is FontMetrics
// is method
func (v FontMetrics) Unref() {
	iv, err := _I.Get(112, "FontMetrics", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Object Fontset
type Fontset struct {
	gobject.Object
}

func WrapFontset(p unsafe.Pointer) (r Fontset) { r.P = p; return }

// pango_fontset_foreach
// container is not nil, container is Fontset
// is method
func (v Fontset) Foreach(func1 int /*TODO_TYPE isPtr: false, tag: interface*/, data unsafe.Pointer) {
	iv, err := _I.Get(113, "Fontset", "foreach")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewIntArgument(func1) /*TODO*/
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_func1, arg_data}
	iv.Call(args, nil, nil)
}

// pango_fontset_get_font
// container is not nil, container is Fontset
// is method
func (v Fontset) GetFont(wc uint32) (result Font) {
	iv, err := _I.Get(114, "Fontset", "get_font")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_wc := gi.NewUint32Argument(wc)
	args := []gi.Argument{arg_v, arg_wc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_fontset_get_metrics
// container is not nil, container is Fontset
// is method
func (v Fontset) GetMetrics() (result FontMetrics) {
	iv, err := _I.Get(115, "Fontset", "get_metrics")
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

// ignore GType struct FontsetClass
// Object FontsetSimple
type FontsetSimple struct {
	Fontset
}

func WrapFontsetSimple(p unsafe.Pointer) (r FontsetSimple) { r.P = p; return }

// pango_fontset_simple_new
// container is not nil, container is FontsetSimple
// is constructor
func NewFontsetSimple(language Language) (result FontsetSimple) {
	iv, err := _I.Get(116, "FontsetSimple", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_language := gi.NewPointerArgument(language.P)
	args := []gi.Argument{arg_language}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_fontset_simple_append
// container is not nil, container is FontsetSimple
// is method
func (v FontsetSimple) Append(font Font) {
	iv, err := _I.Get(117, "FontsetSimple", "append")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_font := gi.NewPointerArgument(font.P)
	args := []gi.Argument{arg_v, arg_font}
	iv.Call(args, nil, nil)
}

// pango_fontset_simple_size
// container is not nil, container is FontsetSimple
// is method
func (v FontsetSimple) Size() (result int32) {
	iv, err := _I.Get(118, "FontsetSimple", "size")
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

// ignore GType struct FontsetSimpleClass
// Struct GlyphGeometry
type GlyphGeometry struct {
	P unsafe.Pointer
}

// Struct GlyphInfo
type GlyphInfo struct {
	P unsafe.Pointer
}

// Struct GlyphItem
type GlyphItem struct {
	P unsafe.Pointer
}

// pango_glyph_item_apply_attrs
// container is not nil, container is GlyphItem
// is method
func (v GlyphItem) ApplyAttrs(text string, list AttrList) (result int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(119, "GlyphItem", "apply_attrs")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_list := gi.NewPointerArgument(list.P)
	args := []gi.Argument{arg_v, arg_text, arg_list}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	gi.Free(c_text)
	return
}

// pango_glyph_item_copy
// container is not nil, container is GlyphItem
// is method
func (v GlyphItem) Copy() (result GlyphItem) {
	iv, err := _I.Get(120, "GlyphItem", "copy")
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

// pango_glyph_item_free
// container is not nil, container is GlyphItem
// is method
func (v GlyphItem) Free() {
	iv, err := _I.Get(121, "GlyphItem", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_glyph_item_get_logical_widths
// container is not nil, container is GlyphItem
// is method
func (v GlyphItem) GetLogicalWidths(text string, logical_widths int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(122, "GlyphItem", "get_logical_widths")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_logical_widths := gi.NewIntArgument(logical_widths) /*TODO*/
	args := []gi.Argument{arg_v, arg_text, arg_logical_widths}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// pango_glyph_item_letter_space
// container is not nil, container is GlyphItem
// is method
func (v GlyphItem) LetterSpace(text string, log_attrs int /*TODO_TYPE isPtr: true, tag: array*/, letter_spacing int32) {
	iv, err := _I.Get(123, "GlyphItem", "letter_space")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_log_attrs := gi.NewIntArgument(log_attrs) /*TODO*/
	arg_letter_spacing := gi.NewInt32Argument(letter_spacing)
	args := []gi.Argument{arg_v, arg_text, arg_log_attrs, arg_letter_spacing}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// pango_glyph_item_split
// container is not nil, container is GlyphItem
// is method
func (v GlyphItem) Split(text string, split_index int32) (result GlyphItem) {
	iv, err := _I.Get(124, "GlyphItem", "split")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_split_index := gi.NewInt32Argument(split_index)
	args := []gi.Argument{arg_v, arg_text, arg_split_index}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	gi.Free(c_text)
	return
}

// Struct GlyphItemIter
type GlyphItemIter struct {
	P unsafe.Pointer
}

// pango_glyph_item_iter_copy
// container is not nil, container is GlyphItemIter
// is method
func (v GlyphItemIter) Copy() (result GlyphItemIter) {
	iv, err := _I.Get(125, "GlyphItemIter", "copy")
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

// pango_glyph_item_iter_free
// container is not nil, container is GlyphItemIter
// is method
func (v GlyphItemIter) Free() {
	iv, err := _I.Get(126, "GlyphItemIter", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_glyph_item_iter_init_end
// container is not nil, container is GlyphItemIter
// is method
func (v GlyphItemIter) InitEnd(glyph_item GlyphItem, text string) (result bool) {
	iv, err := _I.Get(127, "GlyphItemIter", "init_end")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_glyph_item := gi.NewPointerArgument(glyph_item.P)
	arg_text := gi.NewStringArgument(c_text)
	args := []gi.Argument{arg_v, arg_glyph_item, arg_text}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	gi.Free(c_text)
	return
}

// pango_glyph_item_iter_init_start
// container is not nil, container is GlyphItemIter
// is method
func (v GlyphItemIter) InitStart(glyph_item GlyphItem, text string) (result bool) {
	iv, err := _I.Get(128, "GlyphItemIter", "init_start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_glyph_item := gi.NewPointerArgument(glyph_item.P)
	arg_text := gi.NewStringArgument(c_text)
	args := []gi.Argument{arg_v, arg_glyph_item, arg_text}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	gi.Free(c_text)
	return
}

// pango_glyph_item_iter_next_cluster
// container is not nil, container is GlyphItemIter
// is method
func (v GlyphItemIter) NextCluster() (result bool) {
	iv, err := _I.Get(129, "GlyphItemIter", "next_cluster")
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

// pango_glyph_item_iter_prev_cluster
// container is not nil, container is GlyphItemIter
// is method
func (v GlyphItemIter) PrevCluster() (result bool) {
	iv, err := _I.Get(130, "GlyphItemIter", "prev_cluster")
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

// Struct GlyphString
type GlyphString struct {
	P unsafe.Pointer
}

// pango_glyph_string_new
// container is not nil, container is GlyphString
// is constructor
func NewGlyphString() (result GlyphString) {
	iv, err := _I.Get(131, "GlyphString", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_glyph_string_copy
// container is not nil, container is GlyphString
// is method
func (v GlyphString) Copy() (result GlyphString) {
	iv, err := _I.Get(132, "GlyphString", "copy")
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

// pango_glyph_string_extents
// container is not nil, container is GlyphString
// is method
func (v GlyphString) Extents(font Font) (ink_rect int /*TODO_TYPE*/, logical_rect int /*TODO_TYPE*/) {
	iv, err := _I.Get(133, "GlyphString", "extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_font := gi.NewPointerArgument(font.P)
	arg_ink_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_logical_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_font, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, &outArgs[0])
	ink_rect = outArgs[0].Int()     /*TODO*/
	logical_rect = outArgs[1].Int() /*TODO*/
	return
}

// pango_glyph_string_extents_range
// container is not nil, container is GlyphString
// is method
func (v GlyphString) ExtentsRange(start int32, end int32, font Font) (ink_rect int /*TODO_TYPE*/, logical_rect int /*TODO_TYPE*/) {
	iv, err := _I.Get(134, "GlyphString", "extents_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewInt32Argument(start)
	arg_end := gi.NewInt32Argument(end)
	arg_font := gi.NewPointerArgument(font.P)
	arg_ink_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_logical_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_start, arg_end, arg_font, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, &outArgs[0])
	ink_rect = outArgs[0].Int()     /*TODO*/
	logical_rect = outArgs[1].Int() /*TODO*/
	return
}

// pango_glyph_string_free
// container is not nil, container is GlyphString
// is method
func (v GlyphString) Free() {
	iv, err := _I.Get(135, "GlyphString", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_glyph_string_get_logical_widths
// container is not nil, container is GlyphString
// is method
func (v GlyphString) GetLogicalWidths(text string, length int32, embedding_level int32, logical_widths int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(136, "GlyphString", "get_logical_widths")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_length := gi.NewInt32Argument(length)
	arg_embedding_level := gi.NewInt32Argument(embedding_level)
	arg_logical_widths := gi.NewIntArgument(logical_widths) /*TODO*/
	args := []gi.Argument{arg_v, arg_text, arg_length, arg_embedding_level, arg_logical_widths}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// pango_glyph_string_get_width
// container is not nil, container is GlyphString
// is method
func (v GlyphString) GetWidth() (result int32) {
	iv, err := _I.Get(137, "GlyphString", "get_width")
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

// pango_glyph_string_index_to_x
// container is not nil, container is GlyphString
// is method
func (v GlyphString) IndexToX(text string, length int32, analysis Analysis, index_ int32, trailing bool) (x_pos int32) {
	iv, err := _I.Get(138, "GlyphString", "index_to_x")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_length := gi.NewInt32Argument(length)
	arg_analysis := gi.NewPointerArgument(analysis.P)
	arg_index_ := gi.NewInt32Argument(index_)
	arg_trailing := gi.NewBoolArgument(trailing)
	arg_x_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_text, arg_length, arg_analysis, arg_index_, arg_trailing, arg_x_pos}
	iv.Call(args, nil, &outArgs[0])
	gi.Free(c_text)
	x_pos = outArgs[0].Int32()
	return
}

// pango_glyph_string_set_size
// container is not nil, container is GlyphString
// is method
func (v GlyphString) SetSize(new_len int32) {
	iv, err := _I.Get(139, "GlyphString", "set_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_new_len := gi.NewInt32Argument(new_len)
	args := []gi.Argument{arg_v, arg_new_len}
	iv.Call(args, nil, nil)
}

// pango_glyph_string_x_to_index
// container is not nil, container is GlyphString
// is method
func (v GlyphString) XToIndex(text string, length int32, analysis Analysis, x_pos int32) (index_ int32, trailing int32) {
	iv, err := _I.Get(140, "GlyphString", "x_to_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_length := gi.NewInt32Argument(length)
	arg_analysis := gi.NewPointerArgument(analysis.P)
	arg_x_pos := gi.NewInt32Argument(x_pos)
	arg_index_ := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_trailing := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_text, arg_length, arg_analysis, arg_x_pos, arg_index_, arg_trailing}
	iv.Call(args, nil, &outArgs[0])
	gi.Free(c_text)
	index_ = outArgs[0].Int32()
	trailing = outArgs[1].Int32()
	return
}

// Struct GlyphVisAttr
type GlyphVisAttr struct {
	P unsafe.Pointer
}
type GravityEnum int

const (
	GravitySouth GravityEnum = 0
	GravityEast              = 1
	GravityNorth             = 2
	GravityWest              = 3
	GravityAuto              = 4
)

type GravityHintEnum int

const (
	GravityHintNatural GravityHintEnum = 0
	GravityHintStrong                  = 1
	GravityHintLine                    = 2
)

// Struct IncludedModule
type IncludedModule struct {
	P unsafe.Pointer
}

// Struct Item
type Item struct {
	P unsafe.Pointer
}

// pango_item_new
// container is not nil, container is Item
// is constructor
func NewItem() (result Item) {
	iv, err := _I.Get(141, "Item", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_item_copy
// container is not nil, container is Item
// is method
func (v Item) Copy() (result Item) {
	iv, err := _I.Get(142, "Item", "copy")
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

// pango_item_free
// container is not nil, container is Item
// is method
func (v Item) Free() {
	iv, err := _I.Get(143, "Item", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_item_split
// container is not nil, container is Item
// is method
func (v Item) Split(split_index int32, split_offset int32) (result Item) {
	iv, err := _I.Get(144, "Item", "split")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_split_index := gi.NewInt32Argument(split_index)
	arg_split_offset := gi.NewInt32Argument(split_offset)
	args := []gi.Argument{arg_v, arg_split_index, arg_split_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Struct Language
type Language struct {
	P unsafe.Pointer
}

// pango_language_get_sample_string
// container is not nil, container is Language
// is method
func (v Language) GetSampleString() (result string) {
	iv, err := _I.Get(145, "Language", "get_sample_string")
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

// pango_language_get_scripts
// container is not nil, container is Language
// is method
func (v Language) GetScripts() (result int /*TODO_TYPE isPtr: true, tag: array*/, num_scripts int32) {
	iv, err := _I.Get(146, "Language", "get_scripts")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_num_scripts := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_num_scripts}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Int() /*TODO*/
	num_scripts = outArgs[0].Int32()
	return
}

// pango_language_includes_script
// container is not nil, container is Language
// is method
func (v Language) IncludesScript(script int /*TODO_TYPE isPtr: false, tag: interface*/) (result bool) {
	iv, err := _I.Get(147, "Language", "includes_script")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_script := gi.NewIntArgument(script) /*TODO*/
	args := []gi.Argument{arg_v, arg_script}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// pango_language_matches
// container is not nil, container is Language
// is method
func (v Language) Matches(range_list string) (result bool) {
	iv, err := _I.Get(148, "Language", "matches")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_range_list := gi.CString(range_list)
	arg_v := gi.NewPointerArgument(v.P)
	arg_range_list := gi.NewStringArgument(c_range_list)
	args := []gi.Argument{arg_v, arg_range_list}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	gi.Free(c_range_list)
	return
}

// pango_language_to_string
// container is not nil, container is Language
// is method
func (v Language) ToString() (result string) {
	iv, err := _I.Get(149, "Language", "to_string")
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

// pango_language_from_string
// container is not nil, container is Language
// is method
// arg0Type tag: utf8, isPtr: true
func LanguageFromString1(language string) (result Language) {
	iv, err := _I.Get(150, "Language", "from_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_language := gi.CString(language)
	arg_language := gi.NewStringArgument(c_language)
	args := []gi.Argument{arg_language}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	gi.Free(c_language)
	return
}

// pango_language_get_default
// container is not nil, container is Language
// num arg is 0
// Object Layout
type Layout struct {
	gobject.Object
}

func WrapLayout(p unsafe.Pointer) (r Layout) { r.P = p; return }

// pango_layout_new
// container is not nil, container is Layout
// is constructor
func NewLayout(context Context) (result Layout) {
	iv, err := _I.Get(152, "Layout", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_context := gi.NewPointerArgument(context.P)
	args := []gi.Argument{arg_context}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_layout_context_changed
// container is not nil, container is Layout
// is method
func (v Layout) ContextChanged() {
	iv, err := _I.Get(153, "Layout", "context_changed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_layout_copy
// container is not nil, container is Layout
// is method
func (v Layout) Copy() (result Layout) {
	iv, err := _I.Get(154, "Layout", "copy")
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

// pango_layout_get_alignment
// container is not nil, container is Layout
// is method
func (v Layout) GetAlignment() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(155, "Layout", "get_alignment")
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

// pango_layout_get_attributes
// container is not nil, container is Layout
// is method
func (v Layout) GetAttributes() (result AttrList) {
	iv, err := _I.Get(156, "Layout", "get_attributes")
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

// pango_layout_get_auto_dir
// container is not nil, container is Layout
// is method
func (v Layout) GetAutoDir() (result bool) {
	iv, err := _I.Get(157, "Layout", "get_auto_dir")
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

// pango_layout_get_baseline
// container is not nil, container is Layout
// is method
func (v Layout) GetBaseline() (result int32) {
	iv, err := _I.Get(158, "Layout", "get_baseline")
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

// pango_layout_get_character_count
// container is not nil, container is Layout
// is method
func (v Layout) GetCharacterCount() (result int32) {
	iv, err := _I.Get(159, "Layout", "get_character_count")
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

// pango_layout_get_context
// container is not nil, container is Layout
// is method
func (v Layout) GetContext() (result Context) {
	iv, err := _I.Get(160, "Layout", "get_context")
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

// pango_layout_get_cursor_pos
// container is not nil, container is Layout
// is method
func (v Layout) GetCursorPos(index_ int32) (strong_pos int /*TODO_TYPE*/, weak_pos int /*TODO_TYPE*/) {
	iv, err := _I.Get(161, "Layout", "get_cursor_pos")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewInt32Argument(index_)
	arg_strong_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_weak_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_index_, arg_strong_pos, arg_weak_pos}
	iv.Call(args, nil, &outArgs[0])
	strong_pos = outArgs[0].Int() /*TODO*/
	weak_pos = outArgs[1].Int()   /*TODO*/
	return
}

// pango_layout_get_ellipsize
// container is not nil, container is Layout
// is method
func (v Layout) GetEllipsize() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(162, "Layout", "get_ellipsize")
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

// pango_layout_get_extents
// container is not nil, container is Layout
// is method
func (v Layout) GetExtents() (ink_rect int /*TODO_TYPE*/, logical_rect int /*TODO_TYPE*/) {
	iv, err := _I.Get(163, "Layout", "get_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_logical_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, &outArgs[0])
	ink_rect = outArgs[0].Int()     /*TODO*/
	logical_rect = outArgs[1].Int() /*TODO*/
	return
}

// pango_layout_get_font_description
// container is not nil, container is Layout
// is method
func (v Layout) GetFontDescription() (result FontDescription) {
	iv, err := _I.Get(164, "Layout", "get_font_description")
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

// pango_layout_get_height
// container is not nil, container is Layout
// is method
func (v Layout) GetHeight() (result int32) {
	iv, err := _I.Get(165, "Layout", "get_height")
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

// pango_layout_get_indent
// container is not nil, container is Layout
// is method
func (v Layout) GetIndent() (result int32) {
	iv, err := _I.Get(166, "Layout", "get_indent")
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

// pango_layout_get_iter
// container is not nil, container is Layout
// is method
func (v Layout) GetIter() (result LayoutIter) {
	iv, err := _I.Get(167, "Layout", "get_iter")
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

// pango_layout_get_justify
// container is not nil, container is Layout
// is method
func (v Layout) GetJustify() (result bool) {
	iv, err := _I.Get(168, "Layout", "get_justify")
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

// pango_layout_get_line
// container is not nil, container is Layout
// is method
func (v Layout) GetLine(line int32) (result LayoutLine) {
	iv, err := _I.Get(169, "Layout", "get_line")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_line := gi.NewInt32Argument(line)
	args := []gi.Argument{arg_v, arg_line}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_layout_get_line_count
// container is not nil, container is Layout
// is method
func (v Layout) GetLineCount() (result int32) {
	iv, err := _I.Get(170, "Layout", "get_line_count")
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

// pango_layout_get_line_readonly
// container is not nil, container is Layout
// is method
func (v Layout) GetLineReadonly(line int32) (result LayoutLine) {
	iv, err := _I.Get(171, "Layout", "get_line_readonly")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_line := gi.NewInt32Argument(line)
	args := []gi.Argument{arg_v, arg_line}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_layout_get_lines
// container is not nil, container is Layout
// is method
func (v Layout) GetLines() (result int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(172, "Layout", "get_lines")
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

// pango_layout_get_lines_readonly
// container is not nil, container is Layout
// is method
func (v Layout) GetLinesReadonly() (result int /*TODO_TYPE isPtr: true, tag: gslist*/) {
	iv, err := _I.Get(173, "Layout", "get_lines_readonly")
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

// pango_layout_get_log_attrs
// container is not nil, container is Layout
// is method
func (v Layout) GetLogAttrs() (attrs int /*TODO_TYPE*/, n_attrs int32) {
	iv, err := _I.Get(174, "Layout", "get_log_attrs")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_attrs := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_n_attrs := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_attrs, arg_n_attrs}
	iv.Call(args, nil, &outArgs[0])
	attrs = outArgs[0].Int() /*TODO*/
	n_attrs = outArgs[1].Int32()
	return
}

// pango_layout_get_log_attrs_readonly
// container is not nil, container is Layout
// is method
func (v Layout) GetLogAttrsReadonly() (result int /*TODO_TYPE isPtr: true, tag: array*/, n_attrs int32) {
	iv, err := _I.Get(175, "Layout", "get_log_attrs_readonly")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_attrs := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_n_attrs}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Int() /*TODO*/
	n_attrs = outArgs[0].Int32()
	return
}

// pango_layout_get_pixel_extents
// container is not nil, container is Layout
// is method
func (v Layout) GetPixelExtents() (ink_rect int /*TODO_TYPE*/, logical_rect int /*TODO_TYPE*/) {
	iv, err := _I.Get(176, "Layout", "get_pixel_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_logical_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, &outArgs[0])
	ink_rect = outArgs[0].Int()     /*TODO*/
	logical_rect = outArgs[1].Int() /*TODO*/
	return
}

// pango_layout_get_pixel_size
// container is not nil, container is Layout
// is method
func (v Layout) GetPixelSize() (width int32, height int32) {
	iv, err := _I.Get(177, "Layout", "get_pixel_size")
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
	width = outArgs[0].Int32()
	height = outArgs[1].Int32()
	return
}

// pango_layout_get_serial
// container is not nil, container is Layout
// is method
func (v Layout) GetSerial() (result uint32) {
	iv, err := _I.Get(178, "Layout", "get_serial")
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

// pango_layout_get_single_paragraph_mode
// container is not nil, container is Layout
// is method
func (v Layout) GetSingleParagraphMode() (result bool) {
	iv, err := _I.Get(179, "Layout", "get_single_paragraph_mode")
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

// pango_layout_get_size
// container is not nil, container is Layout
// is method
func (v Layout) GetSize() (width int32, height int32) {
	iv, err := _I.Get(180, "Layout", "get_size")
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
	width = outArgs[0].Int32()
	height = outArgs[1].Int32()
	return
}

// pango_layout_get_spacing
// container is not nil, container is Layout
// is method
func (v Layout) GetSpacing() (result int32) {
	iv, err := _I.Get(181, "Layout", "get_spacing")
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

// pango_layout_get_tabs
// container is not nil, container is Layout
// is method
func (v Layout) GetTabs() (result TabArray) {
	iv, err := _I.Get(182, "Layout", "get_tabs")
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

// pango_layout_get_text
// container is not nil, container is Layout
// is method
func (v Layout) GetText() (result string) {
	iv, err := _I.Get(183, "Layout", "get_text")
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

// pango_layout_get_unknown_glyphs_count
// container is not nil, container is Layout
// is method
func (v Layout) GetUnknownGlyphsCount() (result int32) {
	iv, err := _I.Get(184, "Layout", "get_unknown_glyphs_count")
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

// pango_layout_get_width
// container is not nil, container is Layout
// is method
func (v Layout) GetWidth() (result int32) {
	iv, err := _I.Get(185, "Layout", "get_width")
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

// pango_layout_get_wrap
// container is not nil, container is Layout
// is method
func (v Layout) GetWrap() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(186, "Layout", "get_wrap")
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

// pango_layout_index_to_line_x
// container is not nil, container is Layout
// is method
func (v Layout) IndexToLineX(index_ int32, trailing bool) (line int32, x_pos int32) {
	iv, err := _I.Get(187, "Layout", "index_to_line_x")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewInt32Argument(index_)
	arg_trailing := gi.NewBoolArgument(trailing)
	arg_line := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_x_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_index_, arg_trailing, arg_line, arg_x_pos}
	iv.Call(args, nil, &outArgs[0])
	line = outArgs[0].Int32()
	x_pos = outArgs[1].Int32()
	return
}

// pango_layout_index_to_pos
// container is not nil, container is Layout
// is method
func (v Layout) IndexToPos(index_ int32) (pos int /*TODO_TYPE*/) {
	iv, err := _I.Get(188, "Layout", "index_to_pos")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewInt32Argument(index_)
	arg_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_index_, arg_pos}
	iv.Call(args, nil, &outArgs[0])
	pos = outArgs[0].Int() /*TODO*/
	return
}

// pango_layout_is_ellipsized
// container is not nil, container is Layout
// is method
func (v Layout) IsEllipsized() (result bool) {
	iv, err := _I.Get(189, "Layout", "is_ellipsized")
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

// pango_layout_is_wrapped
// container is not nil, container is Layout
// is method
func (v Layout) IsWrapped() (result bool) {
	iv, err := _I.Get(190, "Layout", "is_wrapped")
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

// pango_layout_move_cursor_visually
// container is not nil, container is Layout
// is method
func (v Layout) MoveCursorVisually(strong bool, old_index int32, old_trailing int32, direction int32) (new_index int32, new_trailing int32) {
	iv, err := _I.Get(191, "Layout", "move_cursor_visually")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_strong := gi.NewBoolArgument(strong)
	arg_old_index := gi.NewInt32Argument(old_index)
	arg_old_trailing := gi.NewInt32Argument(old_trailing)
	arg_direction := gi.NewInt32Argument(direction)
	arg_new_index := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_new_trailing := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_strong, arg_old_index, arg_old_trailing, arg_direction, arg_new_index, arg_new_trailing}
	iv.Call(args, nil, &outArgs[0])
	new_index = outArgs[0].Int32()
	new_trailing = outArgs[1].Int32()
	return
}

// pango_layout_set_alignment
// container is not nil, container is Layout
// is method
func (v Layout) SetAlignment(alignment int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(192, "Layout", "set_alignment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_alignment := gi.NewIntArgument(alignment) /*TODO*/
	args := []gi.Argument{arg_v, arg_alignment}
	iv.Call(args, nil, nil)
}

// pango_layout_set_attributes
// container is not nil, container is Layout
// is method
func (v Layout) SetAttributes(attrs AttrList) {
	iv, err := _I.Get(193, "Layout", "set_attributes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_attrs := gi.NewPointerArgument(attrs.P)
	args := []gi.Argument{arg_v, arg_attrs}
	iv.Call(args, nil, nil)
}

// pango_layout_set_auto_dir
// container is not nil, container is Layout
// is method
func (v Layout) SetAutoDir(auto_dir bool) {
	iv, err := _I.Get(194, "Layout", "set_auto_dir")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_auto_dir := gi.NewBoolArgument(auto_dir)
	args := []gi.Argument{arg_v, arg_auto_dir}
	iv.Call(args, nil, nil)
}

// pango_layout_set_ellipsize
// container is not nil, container is Layout
// is method
func (v Layout) SetEllipsize(ellipsize int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(195, "Layout", "set_ellipsize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ellipsize := gi.NewIntArgument(ellipsize) /*TODO*/
	args := []gi.Argument{arg_v, arg_ellipsize}
	iv.Call(args, nil, nil)
}

// pango_layout_set_font_description
// container is not nil, container is Layout
// is method
func (v Layout) SetFontDescription(desc FontDescription) {
	iv, err := _I.Get(196, "Layout", "set_font_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_desc := gi.NewPointerArgument(desc.P)
	args := []gi.Argument{arg_v, arg_desc}
	iv.Call(args, nil, nil)
}

// pango_layout_set_height
// container is not nil, container is Layout
// is method
func (v Layout) SetHeight(height int32) {
	iv, err := _I.Get(197, "Layout", "set_height")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_height := gi.NewInt32Argument(height)
	args := []gi.Argument{arg_v, arg_height}
	iv.Call(args, nil, nil)
}

// pango_layout_set_indent
// container is not nil, container is Layout
// is method
func (v Layout) SetIndent(indent int32) {
	iv, err := _I.Get(198, "Layout", "set_indent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_indent := gi.NewInt32Argument(indent)
	args := []gi.Argument{arg_v, arg_indent}
	iv.Call(args, nil, nil)
}

// pango_layout_set_justify
// container is not nil, container is Layout
// is method
func (v Layout) SetJustify(justify bool) {
	iv, err := _I.Get(199, "Layout", "set_justify")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_justify := gi.NewBoolArgument(justify)
	args := []gi.Argument{arg_v, arg_justify}
	iv.Call(args, nil, nil)
}

// pango_layout_set_markup
// container is not nil, container is Layout
// is method
func (v Layout) SetMarkup(markup string, length int32) {
	iv, err := _I.Get(200, "Layout", "set_markup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_markup := gi.CString(markup)
	arg_v := gi.NewPointerArgument(v.P)
	arg_markup := gi.NewStringArgument(c_markup)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_v, arg_markup, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_markup)
}

// pango_layout_set_markup_with_accel
// container is not nil, container is Layout
// is method
func (v Layout) SetMarkupWithAccel(markup string, length int32, accel_marker int /*TODO_TYPE isPtr: false, tag: gunichar*/) (accel_char int /*TODO_TYPE*/) {
	iv, err := _I.Get(201, "Layout", "set_markup_with_accel")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_markup := gi.CString(markup)
	arg_v := gi.NewPointerArgument(v.P)
	arg_markup := gi.NewStringArgument(c_markup)
	arg_length := gi.NewInt32Argument(length)
	arg_accel_marker := gi.NewIntArgument(accel_marker) /*TODO*/
	arg_accel_char := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_markup, arg_length, arg_accel_marker, arg_accel_char}
	iv.Call(args, nil, &outArgs[0])
	gi.Free(c_markup)
	accel_char = outArgs[0].Int() /*TODO*/
	return
}

// pango_layout_set_single_paragraph_mode
// container is not nil, container is Layout
// is method
func (v Layout) SetSingleParagraphMode(setting bool) {
	iv, err := _I.Get(202, "Layout", "set_single_paragraph_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_setting := gi.NewBoolArgument(setting)
	args := []gi.Argument{arg_v, arg_setting}
	iv.Call(args, nil, nil)
}

// pango_layout_set_spacing
// container is not nil, container is Layout
// is method
func (v Layout) SetSpacing(spacing int32) {
	iv, err := _I.Get(203, "Layout", "set_spacing")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_spacing := gi.NewInt32Argument(spacing)
	args := []gi.Argument{arg_v, arg_spacing}
	iv.Call(args, nil, nil)
}

// pango_layout_set_tabs
// container is not nil, container is Layout
// is method
func (v Layout) SetTabs(tabs TabArray) {
	iv, err := _I.Get(204, "Layout", "set_tabs")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tabs := gi.NewPointerArgument(tabs.P)
	args := []gi.Argument{arg_v, arg_tabs}
	iv.Call(args, nil, nil)
}

// pango_layout_set_text
// container is not nil, container is Layout
// is method
func (v Layout) SetText(text string, length int32) {
	iv, err := _I.Get(205, "Layout", "set_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_v, arg_text, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// pango_layout_set_width
// container is not nil, container is Layout
// is method
func (v Layout) SetWidth(width int32) {
	iv, err := _I.Get(206, "Layout", "set_width")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_width := gi.NewInt32Argument(width)
	args := []gi.Argument{arg_v, arg_width}
	iv.Call(args, nil, nil)
}

// pango_layout_set_wrap
// container is not nil, container is Layout
// is method
func (v Layout) SetWrap(wrap int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(207, "Layout", "set_wrap")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_wrap := gi.NewIntArgument(wrap) /*TODO*/
	args := []gi.Argument{arg_v, arg_wrap}
	iv.Call(args, nil, nil)
}

// pango_layout_xy_to_index
// container is not nil, container is Layout
// is method
func (v Layout) XyToIndex(x int32, y int32) (result bool, index_ int32, trailing int32) {
	iv, err := _I.Get(208, "Layout", "xy_to_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_index_ := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_trailing := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_x, arg_y, arg_index_, arg_trailing}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	index_ = outArgs[0].Int32()
	trailing = outArgs[1].Int32()
	return
}

// ignore GType struct LayoutClass
// Struct LayoutIter
type LayoutIter struct {
	P unsafe.Pointer
}

// pango_layout_iter_at_last_line
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) AtLastLine() (result bool) {
	iv, err := _I.Get(209, "LayoutIter", "at_last_line")
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

// pango_layout_iter_copy
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) Copy() (result LayoutIter) {
	iv, err := _I.Get(210, "LayoutIter", "copy")
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

// pango_layout_iter_free
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) Free() {
	iv, err := _I.Get(211, "LayoutIter", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_layout_iter_get_baseline
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) GetBaseline() (result int32) {
	iv, err := _I.Get(212, "LayoutIter", "get_baseline")
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

// pango_layout_iter_get_char_extents
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) GetCharExtents() (logical_rect int /*TODO_TYPE*/) {
	iv, err := _I.Get(213, "LayoutIter", "get_char_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_logical_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_logical_rect}
	iv.Call(args, nil, &outArgs[0])
	logical_rect = outArgs[0].Int() /*TODO*/
	return
}

// pango_layout_iter_get_cluster_extents
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) GetClusterExtents() (ink_rect int /*TODO_TYPE*/, logical_rect int /*TODO_TYPE*/) {
	iv, err := _I.Get(214, "LayoutIter", "get_cluster_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_logical_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, &outArgs[0])
	ink_rect = outArgs[0].Int()     /*TODO*/
	logical_rect = outArgs[1].Int() /*TODO*/
	return
}

// pango_layout_iter_get_index
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) GetIndex() (result int32) {
	iv, err := _I.Get(215, "LayoutIter", "get_index")
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

// pango_layout_iter_get_layout
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) GetLayout() (result Layout) {
	iv, err := _I.Get(216, "LayoutIter", "get_layout")
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

// pango_layout_iter_get_layout_extents
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) GetLayoutExtents() (ink_rect int /*TODO_TYPE*/, logical_rect int /*TODO_TYPE*/) {
	iv, err := _I.Get(217, "LayoutIter", "get_layout_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_logical_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, &outArgs[0])
	ink_rect = outArgs[0].Int()     /*TODO*/
	logical_rect = outArgs[1].Int() /*TODO*/
	return
}

// pango_layout_iter_get_line
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) GetLine() (result LayoutLine) {
	iv, err := _I.Get(218, "LayoutIter", "get_line")
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

// pango_layout_iter_get_line_extents
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) GetLineExtents() (ink_rect int /*TODO_TYPE*/, logical_rect int /*TODO_TYPE*/) {
	iv, err := _I.Get(219, "LayoutIter", "get_line_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_logical_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, &outArgs[0])
	ink_rect = outArgs[0].Int()     /*TODO*/
	logical_rect = outArgs[1].Int() /*TODO*/
	return
}

// pango_layout_iter_get_line_readonly
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) GetLineReadonly() (result LayoutLine) {
	iv, err := _I.Get(220, "LayoutIter", "get_line_readonly")
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

// pango_layout_iter_get_line_yrange
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) GetLineYrange() (y0_ int32, y1_ int32) {
	iv, err := _I.Get(221, "LayoutIter", "get_line_yrange")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_y0_ := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_y1_ := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_y0_, arg_y1_}
	iv.Call(args, nil, &outArgs[0])
	y0_ = outArgs[0].Int32()
	y1_ = outArgs[1].Int32()
	return
}

// pango_layout_iter_get_run
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) GetRun() (result GlyphItem) {
	iv, err := _I.Get(222, "LayoutIter", "get_run")
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

// pango_layout_iter_get_run_extents
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) GetRunExtents() (ink_rect int /*TODO_TYPE*/, logical_rect int /*TODO_TYPE*/) {
	iv, err := _I.Get(223, "LayoutIter", "get_run_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_logical_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, &outArgs[0])
	ink_rect = outArgs[0].Int()     /*TODO*/
	logical_rect = outArgs[1].Int() /*TODO*/
	return
}

// pango_layout_iter_get_run_readonly
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) GetRunReadonly() (result GlyphItem) {
	iv, err := _I.Get(224, "LayoutIter", "get_run_readonly")
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

// pango_layout_iter_next_char
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) NextChar() (result bool) {
	iv, err := _I.Get(225, "LayoutIter", "next_char")
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

// pango_layout_iter_next_cluster
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) NextCluster() (result bool) {
	iv, err := _I.Get(226, "LayoutIter", "next_cluster")
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

// pango_layout_iter_next_line
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) NextLine() (result bool) {
	iv, err := _I.Get(227, "LayoutIter", "next_line")
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

// pango_layout_iter_next_run
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) NextRun() (result bool) {
	iv, err := _I.Get(228, "LayoutIter", "next_run")
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

// Struct LayoutLine
type LayoutLine struct {
	P unsafe.Pointer
}

// pango_layout_line_get_extents
// container is not nil, container is LayoutLine
// is method
func (v LayoutLine) GetExtents() (ink_rect int /*TODO_TYPE*/, logical_rect int /*TODO_TYPE*/) {
	iv, err := _I.Get(229, "LayoutLine", "get_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_logical_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, &outArgs[0])
	ink_rect = outArgs[0].Int()     /*TODO*/
	logical_rect = outArgs[1].Int() /*TODO*/
	return
}

// pango_layout_line_get_pixel_extents
// container is not nil, container is LayoutLine
// is method
func (v LayoutLine) GetPixelExtents() (ink_rect int /*TODO_TYPE*/, logical_rect int /*TODO_TYPE*/) {
	iv, err := _I.Get(230, "LayoutLine", "get_pixel_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_logical_rect := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, &outArgs[0])
	ink_rect = outArgs[0].Int()     /*TODO*/
	logical_rect = outArgs[1].Int() /*TODO*/
	return
}

// pango_layout_line_get_x_ranges
// container is not nil, container is LayoutLine
// is method
func (v LayoutLine) GetXRanges(start_index int32, end_index int32) (ranges int /*TODO_TYPE*/, n_ranges int32) {
	iv, err := _I.Get(231, "LayoutLine", "get_x_ranges")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_start_index := gi.NewInt32Argument(start_index)
	arg_end_index := gi.NewInt32Argument(end_index)
	arg_ranges := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_n_ranges := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_start_index, arg_end_index, arg_ranges, arg_n_ranges}
	iv.Call(args, nil, &outArgs[0])
	ranges = outArgs[0].Int() /*TODO*/
	n_ranges = outArgs[1].Int32()
	return
}

// pango_layout_line_index_to_x
// container is not nil, container is LayoutLine
// is method
func (v LayoutLine) IndexToX(index_ int32, trailing bool) (x_pos int32) {
	iv, err := _I.Get(232, "LayoutLine", "index_to_x")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewInt32Argument(index_)
	arg_trailing := gi.NewBoolArgument(trailing)
	arg_x_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_index_, arg_trailing, arg_x_pos}
	iv.Call(args, nil, &outArgs[0])
	x_pos = outArgs[0].Int32()
	return
}

// pango_layout_line_ref
// container is not nil, container is LayoutLine
// is method
func (v LayoutLine) Ref() (result LayoutLine) {
	iv, err := _I.Get(233, "LayoutLine", "ref")
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

// pango_layout_line_unref
// container is not nil, container is LayoutLine
// is method
func (v LayoutLine) Unref() {
	iv, err := _I.Get(234, "LayoutLine", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_layout_line_x_to_index
// container is not nil, container is LayoutLine
// is method
func (v LayoutLine) XToIndex(x_pos int32) (result bool, index_ int32, trailing int32) {
	iv, err := _I.Get(235, "LayoutLine", "x_to_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_x_pos := gi.NewInt32Argument(x_pos)
	arg_index_ := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_trailing := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_x_pos, arg_index_, arg_trailing}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	index_ = outArgs[0].Int32()
	trailing = outArgs[1].Int32()
	return
}

// Struct LogAttr
type LogAttr struct {
	P unsafe.Pointer
}

// Struct Map
type Map struct {
	P unsafe.Pointer
}

// Struct MapEntry
type MapEntry struct {
	P unsafe.Pointer
}

// Struct Matrix
type Matrix struct {
	P unsafe.Pointer
}

// pango_matrix_concat
// container is not nil, container is Matrix
// is method
func (v Matrix) Concat(new_matrix Matrix) {
	iv, err := _I.Get(236, "Matrix", "concat")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_new_matrix := gi.NewPointerArgument(new_matrix.P)
	args := []gi.Argument{arg_v, arg_new_matrix}
	iv.Call(args, nil, nil)
}

// pango_matrix_copy
// container is not nil, container is Matrix
// is method
func (v Matrix) Copy() (result Matrix) {
	iv, err := _I.Get(237, "Matrix", "copy")
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

// pango_matrix_free
// container is not nil, container is Matrix
// is method
func (v Matrix) Free() {
	iv, err := _I.Get(238, "Matrix", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_matrix_get_font_scale_factor
// container is not nil, container is Matrix
// is method
func (v Matrix) GetFontScaleFactor() (result float64) {
	iv, err := _I.Get(239, "Matrix", "get_font_scale_factor")
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

// pango_matrix_get_font_scale_factors
// container is not nil, container is Matrix
// is method
func (v Matrix) GetFontScaleFactors() (xscale float64, yscale float64) {
	iv, err := _I.Get(240, "Matrix", "get_font_scale_factors")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_xscale := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_yscale := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_xscale, arg_yscale}
	iv.Call(args, nil, &outArgs[0])
	xscale = outArgs[0].Double()
	yscale = outArgs[1].Double()
	return
}

// pango_matrix_rotate
// container is not nil, container is Matrix
// is method
func (v Matrix) Rotate(degrees float64) {
	iv, err := _I.Get(241, "Matrix", "rotate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_degrees := gi.NewDoubleArgument(degrees)
	args := []gi.Argument{arg_v, arg_degrees}
	iv.Call(args, nil, nil)
}

// pango_matrix_scale
// container is not nil, container is Matrix
// is method
func (v Matrix) Scale(scale_x float64, scale_y float64) {
	iv, err := _I.Get(242, "Matrix", "scale")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_scale_x := gi.NewDoubleArgument(scale_x)
	arg_scale_y := gi.NewDoubleArgument(scale_y)
	args := []gi.Argument{arg_v, arg_scale_x, arg_scale_y}
	iv.Call(args, nil, nil)
}

// pango_matrix_transform_distance
// container is not nil, container is Matrix
// is method
func (v Matrix) TransformDistance(dx int, dy int) {
	iv, err := _I.Get(243, "Matrix", "transform_distance")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, &outArgs[0])
}

// pango_matrix_transform_pixel_rectangle
// container is not nil, container is Matrix
// is method
func (v Matrix) TransformPixelRectangle(rect int) {
	iv, err := _I.Get(244, "Matrix", "transform_pixel_rectangle")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, &outArgs[0])
}

// pango_matrix_transform_point
// container is not nil, container is Matrix
// is method
func (v Matrix) TransformPoint(x int, y int) {
	iv, err := _I.Get(245, "Matrix", "transform_point")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, &outArgs[0])
}

// pango_matrix_transform_rectangle
// container is not nil, container is Matrix
// is method
func (v Matrix) TransformRectangle(rect int) {
	iv, err := _I.Get(246, "Matrix", "transform_rectangle")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, &outArgs[0])
}

// pango_matrix_translate
// container is not nil, container is Matrix
// is method
func (v Matrix) Translate(tx float64, ty float64) {
	iv, err := _I.Get(247, "Matrix", "translate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tx := gi.NewDoubleArgument(tx)
	arg_ty := gi.NewDoubleArgument(ty)
	args := []gi.Argument{arg_v, arg_tx, arg_ty}
	iv.Call(args, nil, nil)
}

// Struct Rectangle
type Rectangle struct {
	P unsafe.Pointer
}
type RenderPartEnum int

const (
	RenderPartForeground    RenderPartEnum = 0
	RenderPartBackground                   = 1
	RenderPartUnderline                    = 2
	RenderPartStrikethrough                = 3
)

// Object Renderer
type Renderer struct {
	gobject.Object
}

func WrapRenderer(p unsafe.Pointer) (r Renderer) { r.P = p; return }

// pango_renderer_activate
// container is not nil, container is Renderer
// is method
func (v Renderer) Activate() {
	iv, err := _I.Get(248, "Renderer", "activate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_renderer_deactivate
// container is not nil, container is Renderer
// is method
func (v Renderer) Deactivate() {
	iv, err := _I.Get(249, "Renderer", "deactivate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_renderer_draw_error_underline
// container is not nil, container is Renderer
// is method
func (v Renderer) DrawErrorUnderline(x int32, y int32, width int32, height int32) {
	iv, err := _I.Get(250, "Renderer", "draw_error_underline")
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

// pango_renderer_draw_glyph
// container is not nil, container is Renderer
// is method
func (v Renderer) DrawGlyph(font Font, glyph uint32, x float64, y float64) {
	iv, err := _I.Get(251, "Renderer", "draw_glyph")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_font := gi.NewPointerArgument(font.P)
	arg_glyph := gi.NewUint32Argument(glyph)
	arg_x := gi.NewDoubleArgument(x)
	arg_y := gi.NewDoubleArgument(y)
	args := []gi.Argument{arg_v, arg_font, arg_glyph, arg_x, arg_y}
	iv.Call(args, nil, nil)
}

// pango_renderer_draw_glyph_item
// container is not nil, container is Renderer
// is method
func (v Renderer) DrawGlyphItem(text string, glyph_item GlyphItem, x int32, y int32) {
	iv, err := _I.Get(252, "Renderer", "draw_glyph_item")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_glyph_item := gi.NewPointerArgument(glyph_item.P)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	args := []gi.Argument{arg_v, arg_text, arg_glyph_item, arg_x, arg_y}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// pango_renderer_draw_glyphs
// container is not nil, container is Renderer
// is method
func (v Renderer) DrawGlyphs(font Font, glyphs GlyphString, x int32, y int32) {
	iv, err := _I.Get(253, "Renderer", "draw_glyphs")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_font := gi.NewPointerArgument(font.P)
	arg_glyphs := gi.NewPointerArgument(glyphs.P)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	args := []gi.Argument{arg_v, arg_font, arg_glyphs, arg_x, arg_y}
	iv.Call(args, nil, nil)
}

// pango_renderer_draw_layout
// container is not nil, container is Renderer
// is method
func (v Renderer) DrawLayout(layout Layout, x int32, y int32) {
	iv, err := _I.Get(254, "Renderer", "draw_layout")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_layout := gi.NewPointerArgument(layout.P)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	args := []gi.Argument{arg_v, arg_layout, arg_x, arg_y}
	iv.Call(args, nil, nil)
}

// pango_renderer_draw_layout_line
// container is not nil, container is Renderer
// is method
func (v Renderer) DrawLayoutLine(line LayoutLine, x int32, y int32) {
	iv, err := _I.Get(255, "Renderer", "draw_layout_line")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_line := gi.NewPointerArgument(line.P)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	args := []gi.Argument{arg_v, arg_line, arg_x, arg_y}
	iv.Call(args, nil, nil)
}

// pango_renderer_draw_rectangle
// container is not nil, container is Renderer
// is method
func (v Renderer) DrawRectangle(part int /*TODO_TYPE isPtr: false, tag: interface*/, x int32, y int32, width int32, height int32) {
	iv, err := _I.Get(256, "Renderer", "draw_rectangle")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_part := gi.NewIntArgument(part) /*TODO*/
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_width := gi.NewInt32Argument(width)
	arg_height := gi.NewInt32Argument(height)
	args := []gi.Argument{arg_v, arg_part, arg_x, arg_y, arg_width, arg_height}
	iv.Call(args, nil, nil)
}

// pango_renderer_draw_trapezoid
// container is not nil, container is Renderer
// is method
func (v Renderer) DrawTrapezoid(part int /*TODO_TYPE isPtr: false, tag: interface*/, y1_ float64, x11 float64, x21 float64, y2 float64, x12 float64, x22 float64) {
	iv, err := _I.Get(257, "Renderer", "draw_trapezoid")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_part := gi.NewIntArgument(part) /*TODO*/
	arg_y1_ := gi.NewDoubleArgument(y1_)
	arg_x11 := gi.NewDoubleArgument(x11)
	arg_x21 := gi.NewDoubleArgument(x21)
	arg_y2 := gi.NewDoubleArgument(y2)
	arg_x12 := gi.NewDoubleArgument(x12)
	arg_x22 := gi.NewDoubleArgument(x22)
	args := []gi.Argument{arg_v, arg_part, arg_y1_, arg_x11, arg_x21, arg_y2, arg_x12, arg_x22}
	iv.Call(args, nil, nil)
}

// pango_renderer_get_alpha
// container is not nil, container is Renderer
// is method
func (v Renderer) GetAlpha(part int /*TODO_TYPE isPtr: false, tag: interface*/) (result uint16) {
	iv, err := _I.Get(258, "Renderer", "get_alpha")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_part := gi.NewIntArgument(part) /*TODO*/
	args := []gi.Argument{arg_v, arg_part}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint16()
	return
}

// pango_renderer_get_color
// container is not nil, container is Renderer
// is method
func (v Renderer) GetColor(part int /*TODO_TYPE isPtr: false, tag: interface*/) (result Color) {
	iv, err := _I.Get(259, "Renderer", "get_color")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_part := gi.NewIntArgument(part) /*TODO*/
	args := []gi.Argument{arg_v, arg_part}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_renderer_get_layout
// container is not nil, container is Renderer
// is method
func (v Renderer) GetLayout() (result Layout) {
	iv, err := _I.Get(260, "Renderer", "get_layout")
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

// pango_renderer_get_layout_line
// container is not nil, container is Renderer
// is method
func (v Renderer) GetLayoutLine() (result LayoutLine) {
	iv, err := _I.Get(261, "Renderer", "get_layout_line")
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

// pango_renderer_get_matrix
// container is not nil, container is Renderer
// is method
func (v Renderer) GetMatrix() (result Matrix) {
	iv, err := _I.Get(262, "Renderer", "get_matrix")
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

// pango_renderer_part_changed
// container is not nil, container is Renderer
// is method
func (v Renderer) PartChanged(part int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(263, "Renderer", "part_changed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_part := gi.NewIntArgument(part) /*TODO*/
	args := []gi.Argument{arg_v, arg_part}
	iv.Call(args, nil, nil)
}

// pango_renderer_set_alpha
// container is not nil, container is Renderer
// is method
func (v Renderer) SetAlpha(part int /*TODO_TYPE isPtr: false, tag: interface*/, alpha uint16) {
	iv, err := _I.Get(264, "Renderer", "set_alpha")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_part := gi.NewIntArgument(part) /*TODO*/
	arg_alpha := gi.NewUint16Argument(alpha)
	args := []gi.Argument{arg_v, arg_part, arg_alpha}
	iv.Call(args, nil, nil)
}

// pango_renderer_set_color
// container is not nil, container is Renderer
// is method
func (v Renderer) SetColor(part int /*TODO_TYPE isPtr: false, tag: interface*/, color Color) {
	iv, err := _I.Get(265, "Renderer", "set_color")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_part := gi.NewIntArgument(part) /*TODO*/
	arg_color := gi.NewPointerArgument(color.P)
	args := []gi.Argument{arg_v, arg_part, arg_color}
	iv.Call(args, nil, nil)
}

// pango_renderer_set_matrix
// container is not nil, container is Renderer
// is method
func (v Renderer) SetMatrix(matrix Matrix) {
	iv, err := _I.Get(266, "Renderer", "set_matrix")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_matrix := gi.NewPointerArgument(matrix.P)
	args := []gi.Argument{arg_v, arg_matrix}
	iv.Call(args, nil, nil)
}

// ignore GType struct RendererClass
// Struct RendererPrivate
type RendererPrivate struct {
	P unsafe.Pointer
}
type ScriptEnum int

const (
	ScriptInvalidCode          ScriptEnum = -1
	ScriptCommon                          = 0
	ScriptInherited                       = 1
	ScriptArabic                          = 2
	ScriptArmenian                        = 3
	ScriptBengali                         = 4
	ScriptBopomofo                        = 5
	ScriptCherokee                        = 6
	ScriptCoptic                          = 7
	ScriptCyrillic                        = 8
	ScriptDeseret                         = 9
	ScriptDevanagari                      = 10
	ScriptEthiopic                        = 11
	ScriptGeorgian                        = 12
	ScriptGothic                          = 13
	ScriptGreek                           = 14
	ScriptGujarati                        = 15
	ScriptGurmukhi                        = 16
	ScriptHan                             = 17
	ScriptHangul                          = 18
	ScriptHebrew                          = 19
	ScriptHiragana                        = 20
	ScriptKannada                         = 21
	ScriptKatakana                        = 22
	ScriptKhmer                           = 23
	ScriptLao                             = 24
	ScriptLatin                           = 25
	ScriptMalayalam                       = 26
	ScriptMongolian                       = 27
	ScriptMyanmar                         = 28
	ScriptOgham                           = 29
	ScriptOldItalic                       = 30
	ScriptOriya                           = 31
	ScriptRunic                           = 32
	ScriptSinhala                         = 33
	ScriptSyriac                          = 34
	ScriptTamil                           = 35
	ScriptTelugu                          = 36
	ScriptThaana                          = 37
	ScriptThai                            = 38
	ScriptTibetan                         = 39
	ScriptCanadianAboriginal              = 40
	ScriptYi                              = 41
	ScriptTagalog                         = 42
	ScriptHanunoo                         = 43
	ScriptBuhid                           = 44
	ScriptTagbanwa                        = 45
	ScriptBraille                         = 46
	ScriptCypriot                         = 47
	ScriptLimbu                           = 48
	ScriptOsmanya                         = 49
	ScriptShavian                         = 50
	ScriptLinearB                         = 51
	ScriptTaiLe                           = 52
	ScriptUgaritic                        = 53
	ScriptNewTaiLue                       = 54
	ScriptBuginese                        = 55
	ScriptGlagolitic                      = 56
	ScriptTifinagh                        = 57
	ScriptSylotiNagri                     = 58
	ScriptOldPersian                      = 59
	ScriptKharoshthi                      = 60
	ScriptUnknown                         = 61
	ScriptBalinese                        = 62
	ScriptCuneiform                       = 63
	ScriptPhoenician                      = 64
	ScriptPhagsPa                         = 65
	ScriptNko                             = 66
	ScriptKayahLi                         = 67
	ScriptLepcha                          = 68
	ScriptRejang                          = 69
	ScriptSundanese                       = 70
	ScriptSaurashtra                      = 71
	ScriptCham                            = 72
	ScriptOlChiki                         = 73
	ScriptVai                             = 74
	ScriptCarian                          = 75
	ScriptLycian                          = 76
	ScriptLydian                          = 77
	ScriptBatak                           = 78
	ScriptBrahmi                          = 79
	ScriptMandaic                         = 80
	ScriptChakma                          = 81
	ScriptMeroiticCursive                 = 82
	ScriptMeroiticHieroglyphs             = 83
	ScriptMiao                            = 84
	ScriptSharada                         = 85
	ScriptSoraSompeng                     = 86
	ScriptTakri                           = 87
	ScriptBassaVah                        = 88
	ScriptCaucasianAlbanian               = 89
	ScriptDuployan                        = 90
	ScriptElbasan                         = 91
	ScriptGrantha                         = 92
	ScriptKhojki                          = 93
	ScriptKhudawadi                       = 94
	ScriptLinearA                         = 95
	ScriptMahajani                        = 96
	ScriptManichaean                      = 97
	ScriptMendeKikakui                    = 98
	ScriptModi                            = 99
	ScriptMro                             = 100
	ScriptNabataean                       = 101
	ScriptOldNorthArabian                 = 102
	ScriptOldPermic                       = 103
	ScriptPahawhHmong                     = 104
	ScriptPalmyrene                       = 105
	ScriptPauCinHau                       = 106
	ScriptPsalterPahlavi                  = 107
	ScriptSiddham                         = 108
	ScriptTirhuta                         = 109
	ScriptWarangCiti                      = 110
	ScriptAhom                            = 111
	ScriptAnatolianHieroglyphs            = 112
	ScriptHatran                          = 113
	ScriptMultani                         = 114
	ScriptOldHungarian                    = 115
	ScriptSignwriting                     = 116
)

// Struct ScriptIter
type ScriptIter struct {
	P unsafe.Pointer
}

// pango_script_iter_free
// container is not nil, container is ScriptIter
// is method
func (v ScriptIter) Free() {
	iv, err := _I.Get(267, "ScriptIter", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_script_iter_get_range
// container is not nil, container is ScriptIter
// is method
func (v ScriptIter) GetRange() (start string, end string, script int /*TODO_TYPE*/) {
	iv, err := _I.Get(268, "ScriptIter", "get_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_script := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_start, arg_end, arg_script}
	iv.Call(args, nil, &outArgs[0])
	start = outArgs[0].String().Take()
	end = outArgs[1].String().Take()
	script = outArgs[2].Int() /*TODO*/
	return
}

// pango_script_iter_next
// container is not nil, container is ScriptIter
// is method
func (v ScriptIter) Next() (result bool) {
	iv, err := _I.Get(269, "ScriptIter", "next")
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

type StretchEnum int

const (
	StretchUltraCondensed StretchEnum = 0
	StretchExtraCondensed             = 1
	StretchCondensed                  = 2
	StretchSemiCondensed              = 3
	StretchNormal                     = 4
	StretchSemiExpanded               = 5
	StretchExpanded                   = 6
	StretchExtraExpanded              = 7
	StretchUltraExpanded              = 8
)

type StyleEnum int

const (
	StyleNormal  StyleEnum = 0
	StyleOblique           = 1
	StyleItalic            = 2
)

type TabAlignEnum int

const (
	TabAlignLeft TabAlignEnum = 0
)

// Struct TabArray
type TabArray struct {
	P unsafe.Pointer
}

// pango_tab_array_new
// container is not nil, container is TabArray
// is constructor
func NewTabArray(initial_size int32, positions_in_pixels bool) (result TabArray) {
	iv, err := _I.Get(270, "TabArray", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_initial_size := gi.NewInt32Argument(initial_size)
	arg_positions_in_pixels := gi.NewBoolArgument(positions_in_pixels)
	args := []gi.Argument{arg_initial_size, arg_positions_in_pixels}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_tab_array_copy
// container is not nil, container is TabArray
// is method
func (v TabArray) Copy() (result TabArray) {
	iv, err := _I.Get(271, "TabArray", "copy")
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

// pango_tab_array_free
// container is not nil, container is TabArray
// is method
func (v TabArray) Free() {
	iv, err := _I.Get(272, "TabArray", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// pango_tab_array_get_positions_in_pixels
// container is not nil, container is TabArray
// is method
func (v TabArray) GetPositionsInPixels() (result bool) {
	iv, err := _I.Get(273, "TabArray", "get_positions_in_pixels")
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

// pango_tab_array_get_size
// container is not nil, container is TabArray
// is method
func (v TabArray) GetSize() (result int32) {
	iv, err := _I.Get(274, "TabArray", "get_size")
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

// pango_tab_array_get_tab
// container is not nil, container is TabArray
// is method
func (v TabArray) GetTab(tab_index int32) (alignment int /*TODO_TYPE*/, location int32) {
	iv, err := _I.Get(275, "TabArray", "get_tab")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_tab_index := gi.NewInt32Argument(tab_index)
	arg_alignment := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_location := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_tab_index, arg_alignment, arg_location}
	iv.Call(args, nil, &outArgs[0])
	alignment = outArgs[0].Int() /*TODO*/
	location = outArgs[1].Int32()
	return
}

// pango_tab_array_get_tabs
// container is not nil, container is TabArray
// is method
func (v TabArray) GetTabs() (alignments int /*TODO_TYPE*/, locations int /*TODO_TYPE*/) {
	iv, err := _I.Get(276, "TabArray", "get_tabs")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_alignments := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_locations := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_alignments, arg_locations}
	iv.Call(args, nil, &outArgs[0])
	alignments = outArgs[0].Int() /*TODO*/
	locations = outArgs[1].Int()  /*TODO*/
	return
}

// pango_tab_array_resize
// container is not nil, container is TabArray
// is method
func (v TabArray) Resize(new_size int32) {
	iv, err := _I.Get(277, "TabArray", "resize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_new_size := gi.NewInt32Argument(new_size)
	args := []gi.Argument{arg_v, arg_new_size}
	iv.Call(args, nil, nil)
}

// pango_tab_array_set_tab
// container is not nil, container is TabArray
// is method
func (v TabArray) SetTab(tab_index int32, alignment int /*TODO_TYPE isPtr: false, tag: interface*/, location int32) {
	iv, err := _I.Get(278, "TabArray", "set_tab")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tab_index := gi.NewInt32Argument(tab_index)
	arg_alignment := gi.NewIntArgument(alignment) /*TODO*/
	arg_location := gi.NewInt32Argument(location)
	args := []gi.Argument{arg_v, arg_tab_index, arg_alignment, arg_location}
	iv.Call(args, nil, nil)
}

type UnderlineEnum int

const (
	UnderlineNone   UnderlineEnum = 0
	UnderlineSingle               = 1
	UnderlineDouble               = 2
	UnderlineLow                  = 3
	UnderlineError                = 4
)

type VariantEnum int

const (
	VariantNormal    VariantEnum = 0
	VariantSmallCaps             = 1
)

type WeightEnum int

const (
	WeightThin       WeightEnum = 100
	WeightUltralight            = 200
	WeightLight                 = 300
	WeightSemilight             = 350
	WeightBook                  = 380
	WeightNormal                = 400
	WeightMedium                = 500
	WeightSemibold              = 600
	WeightBold                  = 700
	WeightUltrabold             = 800
	WeightHeavy                 = 900
	WeightUltraheavy            = 1000
)

type WrapModeEnum int

const (
	WrapModeWord     WrapModeEnum = 0
	WrapModeChar                  = 1
	WrapModeWordChar              = 2
)

// pango_attr_type_get_name
// container is nil
func AttrTypeGetName(type1 int /*TODO_TYPE isPtr: false, tag: interface*/) (result string) {
	iv, err := _I.Get(279, "attr_type_get_name", "")
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

// pango_attr_type_register
// container is nil
func AttrTypeRegister(name string) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(280, "attr_type_register", "")
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

// pango_bidi_type_for_unichar
// container is nil
func BidiTypeForUnichar(ch int /*TODO_TYPE isPtr: false, tag: gunichar*/) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(281, "bidi_type_for_unichar", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewIntArgument(ch) /*TODO*/
	args := []gi.Argument{arg_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// pango_break
// container is nil
func Break(text string, length int32, analysis Analysis, attrs int /*TODO_TYPE isPtr: true, tag: array*/, attrs_len int32) {
	iv, err := _I.Get(282, "break", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_text := gi.NewStringArgument(c_text)
	arg_length := gi.NewInt32Argument(length)
	arg_analysis := gi.NewPointerArgument(analysis.P)
	arg_attrs := gi.NewIntArgument(attrs) /*TODO*/
	arg_attrs_len := gi.NewInt32Argument(attrs_len)
	args := []gi.Argument{arg_text, arg_length, arg_analysis, arg_attrs, arg_attrs_len}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// pango_config_key_get
// container is nil
func ConfigKeyGet(key string) (result string) {
	iv, err := _I.Get(283, "config_key_get", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	arg_key := gi.NewStringArgument(c_key)
	args := []gi.Argument{arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	gi.Free(c_key)
	return
}

// pango_config_key_get_system
// container is nil
func ConfigKeyGetSystem(key string) (result string) {
	iv, err := _I.Get(284, "config_key_get_system", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	arg_key := gi.NewStringArgument(c_key)
	args := []gi.Argument{arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	gi.Free(c_key)
	return
}

// pango_default_break
// container is nil
func DefaultBreak(text string, length int32, analysis Analysis, attrs LogAttr, attrs_len int32) {
	iv, err := _I.Get(285, "default_break", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_text := gi.NewStringArgument(c_text)
	arg_length := gi.NewInt32Argument(length)
	arg_analysis := gi.NewPointerArgument(analysis.P)
	arg_attrs := gi.NewPointerArgument(attrs.P)
	arg_attrs_len := gi.NewInt32Argument(attrs_len)
	args := []gi.Argument{arg_text, arg_length, arg_analysis, arg_attrs, arg_attrs_len}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// pango_extents_to_pixels
// container is nil
func ExtentsToPixels(inclusive Rectangle, nearest Rectangle) {
	iv, err := _I.Get(286, "extents_to_pixels", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_inclusive := gi.NewPointerArgument(inclusive.P)
	arg_nearest := gi.NewPointerArgument(nearest.P)
	args := []gi.Argument{arg_inclusive, arg_nearest}
	iv.Call(args, nil, nil)
}

// pango_find_base_dir
// container is nil
func FindBaseDir(text string, length int32) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(287, "find_base_dir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_text := gi.NewStringArgument(c_text)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_text, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	gi.Free(c_text)
	return
}

// pango_find_paragraph_boundary
// container is nil
func FindParagraphBoundary(text string, length int32) (paragraph_delimiter_index int32, next_paragraph_start int32) {
	iv, err := _I.Get(288, "find_paragraph_boundary", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	c_text := gi.CString(text)
	arg_text := gi.NewStringArgument(c_text)
	arg_length := gi.NewInt32Argument(length)
	arg_paragraph_delimiter_index := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_next_paragraph_start := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_text, arg_length, arg_paragraph_delimiter_index, arg_next_paragraph_start}
	iv.Call(args, nil, &outArgs[0])
	gi.Free(c_text)
	paragraph_delimiter_index = outArgs[0].Int32()
	next_paragraph_start = outArgs[1].Int32()
	return
}

// pango_font_description_from_string
// container is nil
func FontDescriptionFromString(str string) (result FontDescription) {
	iv, err := _I.Get(289, "font_description_from_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	gi.Free(c_str)
	return
}

// pango_get_lib_subdirectory
// container is nil
func GetLibSubdirectory() (result string) {
	iv, err := _I.Get(290, "get_lib_subdirectory", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// pango_get_log_attrs
// container is nil
func GetLogAttrs(text string, length int32, level int32, language Language, log_attrs int /*TODO_TYPE isPtr: true, tag: array*/, attrs_len int32) {
	iv, err := _I.Get(291, "get_log_attrs", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_text := gi.NewStringArgument(c_text)
	arg_length := gi.NewInt32Argument(length)
	arg_level := gi.NewInt32Argument(level)
	arg_language := gi.NewPointerArgument(language.P)
	arg_log_attrs := gi.NewIntArgument(log_attrs) /*TODO*/
	arg_attrs_len := gi.NewInt32Argument(attrs_len)
	args := []gi.Argument{arg_text, arg_length, arg_level, arg_language, arg_log_attrs, arg_attrs_len}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// pango_get_mirror_char
// container is nil
func GetMirrorChar(ch int /*TODO_TYPE isPtr: false, tag: gunichar*/, mirrored_ch int /*TODO_TYPE isPtr: true, tag: gunichar*/) (result bool) {
	iv, err := _I.Get(292, "get_mirror_char", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewIntArgument(ch)                   /*TODO*/
	arg_mirrored_ch := gi.NewIntArgument(mirrored_ch) /*TODO*/
	args := []gi.Argument{arg_ch, arg_mirrored_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// pango_get_sysconf_subdirectory
// container is nil
func GetSysconfSubdirectory() (result string) {
	iv, err := _I.Get(293, "get_sysconf_subdirectory", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// pango_gravity_get_for_matrix
// container is nil
func GravityGetForMatrix(matrix Matrix) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(294, "gravity_get_for_matrix", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_matrix := gi.NewPointerArgument(matrix.P)
	args := []gi.Argument{arg_matrix}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// pango_gravity_get_for_script
// container is nil
func GravityGetForScript(script int /*TODO_TYPE isPtr: false, tag: interface*/, base_gravity int /*TODO_TYPE isPtr: false, tag: interface*/, hint int /*TODO_TYPE isPtr: false, tag: interface*/) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(295, "gravity_get_for_script", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_script := gi.NewIntArgument(script)             /*TODO*/
	arg_base_gravity := gi.NewIntArgument(base_gravity) /*TODO*/
	arg_hint := gi.NewIntArgument(hint)                 /*TODO*/
	args := []gi.Argument{arg_script, arg_base_gravity, arg_hint}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// pango_gravity_get_for_script_and_width
// container is nil
func GravityGetForScriptAndWidth(script int /*TODO_TYPE isPtr: false, tag: interface*/, wide bool, base_gravity int /*TODO_TYPE isPtr: false, tag: interface*/, hint int /*TODO_TYPE isPtr: false, tag: interface*/) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(296, "gravity_get_for_script_and_width", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_script := gi.NewIntArgument(script) /*TODO*/
	arg_wide := gi.NewBoolArgument(wide)
	arg_base_gravity := gi.NewIntArgument(base_gravity) /*TODO*/
	arg_hint := gi.NewIntArgument(hint)                 /*TODO*/
	args := []gi.Argument{arg_script, arg_wide, arg_base_gravity, arg_hint}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// pango_gravity_to_rotation
// container is nil
func GravityToRotation(gravity int /*TODO_TYPE isPtr: false, tag: interface*/) (result float64) {
	iv, err := _I.Get(297, "gravity_to_rotation", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_gravity := gi.NewIntArgument(gravity) /*TODO*/
	args := []gi.Argument{arg_gravity}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// pango_is_zero_width
// container is nil
func IsZeroWidth(ch int /*TODO_TYPE isPtr: false, tag: gunichar*/) (result bool) {
	iv, err := _I.Get(298, "is_zero_width", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewIntArgument(ch) /*TODO*/
	args := []gi.Argument{arg_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// pango_itemize
// container is nil
func Itemize(context Context, text string, start_index int32, length int32, attrs AttrList, cached_iter AttrIterator) (result int /*TODO_TYPE isPtr: true, tag: glist*/) {
	iv, err := _I.Get(299, "itemize", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_context := gi.NewPointerArgument(context.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_start_index := gi.NewInt32Argument(start_index)
	arg_length := gi.NewInt32Argument(length)
	arg_attrs := gi.NewPointerArgument(attrs.P)
	arg_cached_iter := gi.NewPointerArgument(cached_iter.P)
	args := []gi.Argument{arg_context, arg_text, arg_start_index, arg_length, arg_attrs, arg_cached_iter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	gi.Free(c_text)
	return
}

// pango_itemize_with_base_dir
// container is nil
func ItemizeWithBaseDir(context Context, base_dir int /*TODO_TYPE isPtr: false, tag: interface*/, text string, start_index int32, length int32, attrs AttrList, cached_iter AttrIterator) (result int /*TODO_TYPE isPtr: true, tag: glist*/) {
	iv, err := _I.Get(300, "itemize_with_base_dir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_context := gi.NewPointerArgument(context.P)
	arg_base_dir := gi.NewIntArgument(base_dir) /*TODO*/
	arg_text := gi.NewStringArgument(c_text)
	arg_start_index := gi.NewInt32Argument(start_index)
	arg_length := gi.NewInt32Argument(length)
	arg_attrs := gi.NewPointerArgument(attrs.P)
	arg_cached_iter := gi.NewPointerArgument(cached_iter.P)
	args := []gi.Argument{arg_context, arg_base_dir, arg_text, arg_start_index, arg_length, arg_attrs, arg_cached_iter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	gi.Free(c_text)
	return
}

// pango_language_from_string
// container is nil
func LanguageFromString(language string) (result Language) {
	iv, err := _I.Get(301, "language_from_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_language := gi.CString(language)
	arg_language := gi.NewStringArgument(c_language)
	args := []gi.Argument{arg_language}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	gi.Free(c_language)
	return
}

// pango_language_get_default
// container is nil
func LanguageGetDefault() (result Language) {
	iv, err := _I.Get(302, "language_get_default", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// black function log2vis_get_embedding_levels

// pango_lookup_aliases
// container is nil
func LookupAliases(fontname string) (families int /*TODO_TYPE*/, n_families int32) {
	iv, err := _I.Get(303, "lookup_aliases", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	c_fontname := gi.CString(fontname)
	arg_fontname := gi.NewStringArgument(c_fontname)
	arg_families := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_n_families := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_fontname, arg_families, arg_n_families}
	iv.Call(args, nil, &outArgs[0])
	gi.Free(c_fontname)
	families = outArgs[0].Int() /*TODO*/
	n_families = outArgs[1].Int32()
	return
}

// pango_markup_parser_finish
// container is nil
func MarkupParserFinish(context glib.MarkupParseContext) (result bool, attr_list int /*TODO_TYPE*/, text string, accel_char int /*TODO_TYPE*/, err error) {
	iv, err := _I.Get(304, "markup_parser_finish", "")
	if err != nil {
		return
	}
	var outArgs [4]gi.Argument
	arg_context := gi.NewPointerArgument(context.P)
	arg_attr_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_text := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_accel_char := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_context, arg_attr_list, arg_text, arg_accel_char, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	attr_list = outArgs[0].Int() /*TODO*/
	text = outArgs[1].String().Take()
	accel_char = outArgs[2].Int() /*TODO*/
	err = gi.ToError(outArgs[3].Pointer())
	return
}

// pango_markup_parser_new
// container is nil
func MarkupParserNew(accel_marker int /*TODO_TYPE isPtr: false, tag: gunichar*/) (result glib.MarkupParseContext) {
	iv, err := _I.Get(305, "markup_parser_new", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_accel_marker := gi.NewIntArgument(accel_marker) /*TODO*/
	args := []gi.Argument{arg_accel_marker}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_module_register
// container is nil
func ModuleRegister(module IncludedModule) {
	iv, err := _I.Get(306, "module_register", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_module := gi.NewPointerArgument(module.P)
	args := []gi.Argument{arg_module}
	iv.Call(args, nil, nil)
}

// pango_parse_enum
// container is nil
func ParseEnum(type1 int /*TODO_TYPE isPtr: false, tag: GType*/, str string, warn bool) (result bool, value int32, possible_values string) {
	iv, err := _I.Get(307, "parse_enum", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	c_str := gi.CString(str)
	arg_type1 := gi.NewIntArgument(type1) /*TODO*/
	arg_str := gi.NewStringArgument(c_str)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_warn := gi.NewBoolArgument(warn)
	arg_possible_values := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_type1, arg_str, arg_value, arg_warn, arg_possible_values}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	gi.Free(c_str)
	value = outArgs[0].Int32()
	possible_values = outArgs[1].String().Take()
	return
}

// pango_parse_markup
// container is nil
func ParseMarkup(markup_text string, length int32, accel_marker int /*TODO_TYPE isPtr: false, tag: gunichar*/) (result bool, attr_list int /*TODO_TYPE*/, text string, accel_char int /*TODO_TYPE*/, err error) {
	iv, err := _I.Get(308, "parse_markup", "")
	if err != nil {
		return
	}
	var outArgs [4]gi.Argument
	c_markup_text := gi.CString(markup_text)
	arg_markup_text := gi.NewStringArgument(c_markup_text)
	arg_length := gi.NewInt32Argument(length)
	arg_accel_marker := gi.NewIntArgument(accel_marker) /*TODO*/
	arg_attr_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_text := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_accel_char := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_markup_text, arg_length, arg_accel_marker, arg_attr_list, arg_text, arg_accel_char, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	gi.Free(c_markup_text)
	attr_list = outArgs[0].Int() /*TODO*/
	text = outArgs[1].String().Take()
	accel_char = outArgs[2].Int() /*TODO*/
	err = gi.ToError(outArgs[3].Pointer())
	return
}

// pango_parse_stretch
// container is nil
func ParseStretch(str string, warn bool) (result bool, stretch int /*TODO_TYPE*/) {
	iv, err := _I.Get(309, "parse_stretch", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_stretch := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_warn := gi.NewBoolArgument(warn)
	args := []gi.Argument{arg_str, arg_stretch, arg_warn}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	gi.Free(c_str)
	stretch = outArgs[0].Int() /*TODO*/
	return
}

// pango_parse_style
// container is nil
func ParseStyle(str string, warn bool) (result bool, style int /*TODO_TYPE*/) {
	iv, err := _I.Get(310, "parse_style", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_style := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_warn := gi.NewBoolArgument(warn)
	args := []gi.Argument{arg_str, arg_style, arg_warn}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	gi.Free(c_str)
	style = outArgs[0].Int() /*TODO*/
	return
}

// pango_parse_variant
// container is nil
func ParseVariant(str string, warn bool) (result bool, variant int /*TODO_TYPE*/) {
	iv, err := _I.Get(311, "parse_variant", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_variant := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_warn := gi.NewBoolArgument(warn)
	args := []gi.Argument{arg_str, arg_variant, arg_warn}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	gi.Free(c_str)
	variant = outArgs[0].Int() /*TODO*/
	return
}

// pango_parse_weight
// container is nil
func ParseWeight(str string, warn bool) (result bool, weight int /*TODO_TYPE*/) {
	iv, err := _I.Get(312, "parse_weight", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_weight := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_warn := gi.NewBoolArgument(warn)
	args := []gi.Argument{arg_str, arg_weight, arg_warn}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	gi.Free(c_str)
	weight = outArgs[0].Int() /*TODO*/
	return
}

// pango_quantize_line_geometry
// container is nil
func QuantizeLineGeometry(thickness int, position int) {
	iv, err := _I.Get(313, "quantize_line_geometry", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	iv.Call(nil, nil, &outArgs[0])
}

// pango_read_line
// container is nil
func ReadLine(stream unsafe.Pointer) (result int32, str int /*TODO_TYPE*/) {
	iv, err := _I.Get(314, "read_line", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_stream := gi.NewPointerArgument(stream)
	arg_str := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_stream, arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Int32()
	str = outArgs[0].Int() /*TODO*/
	return
}

// pango_reorder_items
// container is nil
func ReorderItems(logical_items int /*TODO_TYPE isPtr: true, tag: glist*/) (result int /*TODO_TYPE isPtr: true, tag: glist*/) {
	iv, err := _I.Get(315, "reorder_items", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_logical_items := gi.NewIntArgument(logical_items) /*TODO*/
	args := []gi.Argument{arg_logical_items}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// pango_scan_int
// container is nil
func ScanInt(pos int) (result bool, out int32) {
	iv, err := _I.Get(316, "scan_int", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_out := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_out}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	out = outArgs[0].Int32()
	return
}

// pango_scan_string
// container is nil
func ScanString(pos int) (result bool, out int /*TODO_TYPE*/) {
	iv, err := _I.Get(317, "scan_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_out := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_out}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	out = outArgs[0].Int() /*TODO*/
	return
}

// pango_scan_word
// container is nil
func ScanWord(pos int) (result bool, out int /*TODO_TYPE*/) {
	iv, err := _I.Get(318, "scan_word", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_out := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_out}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	out = outArgs[0].Int() /*TODO*/
	return
}

// pango_script_for_unichar
// container is nil
func ScriptForUnichar(ch int /*TODO_TYPE isPtr: false, tag: gunichar*/) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(319, "script_for_unichar", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewIntArgument(ch) /*TODO*/
	args := []gi.Argument{arg_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// pango_script_get_sample_language
// container is nil
func ScriptGetSampleLanguage(script int /*TODO_TYPE isPtr: false, tag: interface*/) (result Language) {
	iv, err := _I.Get(320, "script_get_sample_language", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_script := gi.NewIntArgument(script) /*TODO*/
	args := []gi.Argument{arg_script}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_shape
// container is nil
func Shape(text string, length int32, analysis Analysis, glyphs GlyphString) {
	iv, err := _I.Get(321, "shape", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_text := gi.NewStringArgument(c_text)
	arg_length := gi.NewInt32Argument(length)
	arg_analysis := gi.NewPointerArgument(analysis.P)
	arg_glyphs := gi.NewPointerArgument(glyphs.P)
	args := []gi.Argument{arg_text, arg_length, arg_analysis, arg_glyphs}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// pango_shape_full
// container is nil
func ShapeFull(item_text string, item_length int32, paragraph_text string, paragraph_length int32, analysis Analysis, glyphs GlyphString) {
	iv, err := _I.Get(322, "shape_full", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_item_text := gi.CString(item_text)
	c_paragraph_text := gi.CString(paragraph_text)
	arg_item_text := gi.NewStringArgument(c_item_text)
	arg_item_length := gi.NewInt32Argument(item_length)
	arg_paragraph_text := gi.NewStringArgument(c_paragraph_text)
	arg_paragraph_length := gi.NewInt32Argument(paragraph_length)
	arg_analysis := gi.NewPointerArgument(analysis.P)
	arg_glyphs := gi.NewPointerArgument(glyphs.P)
	args := []gi.Argument{arg_item_text, arg_item_length, arg_paragraph_text, arg_paragraph_length, arg_analysis, arg_glyphs}
	iv.Call(args, nil, nil)
	gi.Free(c_item_text)
	gi.Free(c_paragraph_text)
}

// pango_skip_space
// container is nil
func SkipSpace(pos int) (result bool) {
	iv, err := _I.Get(323, "skip_space", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	var ret gi.Argument
	iv.Call(nil, &ret, &outArgs[0])
	result = ret.Bool()
	return
}

// pango_split_file_list
// container is nil
func SplitFileList(str string) (result int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(324, "split_file_list", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	gi.Free(c_str)
	return
}

// pango_trim_string
// container is nil
func TrimString(str string) (result string) {
	iv, err := _I.Get(325, "trim_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	gi.Free(c_str)
	return
}

// pango_unichar_direction
// container is nil
func UnicharDirection(ch int /*TODO_TYPE isPtr: false, tag: gunichar*/) (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(326, "unichar_direction", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewIntArgument(ch) /*TODO*/
	args := []gi.Argument{arg_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int() /*TODO*/
	return
}

// pango_units_from_double
// container is nil
func UnitsFromDouble(d float64) (result int32) {
	iv, err := _I.Get(327, "units_from_double", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_d := gi.NewDoubleArgument(d)
	args := []gi.Argument{arg_d}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// pango_units_to_double
// container is nil
func UnitsToDouble(i int32) (result float64) {
	iv, err := _I.Get(328, "units_to_double", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_i := gi.NewInt32Argument(i)
	args := []gi.Argument{arg_i}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// pango_version
// container is nil
func Version() (result int32) {
	iv, err := _I.Get(329, "version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int32()
	return
}

// pango_version_check
// container is nil
func VersionCheck(required_major int32, required_minor int32, required_micro int32) (result string) {
	iv, err := _I.Get(330, "version_check", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_required_major := gi.NewInt32Argument(required_major)
	arg_required_minor := gi.NewInt32Argument(required_minor)
	arg_required_micro := gi.NewInt32Argument(required_micro)
	args := []gi.Argument{arg_required_major, arg_required_minor, arg_required_micro}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// pango_version_string
// container is nil
func VersionString() (result string) {
	iv, err := _I.Get(331, "version_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// constants
const (
	ANALYSIS_FLAG_CENTERED_BASELINE = 1
	ANALYSIS_FLAG_IS_ELLIPSIS       = 2
	ATTR_INDEX_FROM_TEXT_BEGINNING  = 0
	ENGINE_TYPE_LANG                = "PangoEngineLang"
	ENGINE_TYPE_SHAPE               = "PangoEngineShape"
	GLYPH_EMPTY                     = 268435455
	GLYPH_INVALID_INPUT             = 4294967295
	GLYPH_UNKNOWN_FLAG              = 268435456
	RENDER_TYPE_NONE                = "PangoRenderNone"
	SCALE                           = 1024
	UNKNOWN_GLYPH_HEIGHT            = 14
	UNKNOWN_GLYPH_WIDTH             = 10
	VERSION_MIN_REQUIRED            = 2
)