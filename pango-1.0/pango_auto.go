package pango

import "github.com/electricface/go-gir/glib-2.0"
import "github.com/electricface/go-gir/gobject-2.0"
import "github.com/electricface/go-gir3/gi"
import "log"
import "unsafe"

var _I = gi.NewInvokerCache("Pango")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("Pango", "1.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Enum Alignment
type AlignmentEnum int

const (
	AlignmentLeft   AlignmentEnum = 0
	AlignmentCenter AlignmentEnum = 1
	AlignmentRight  AlignmentEnum = 2
)

func AlignmentGetType() gi.GType {
	ret := _I.GetGType(0, "Alignment")
	return ret
}

// Struct Analysis
type Analysis struct {
	P unsafe.Pointer
}

const SizeOfStructAnalysis = 48

func AnalysisGetType() gi.GType {
	ret := _I.GetGType(1, "Analysis")
	return ret
}

// Struct AttrClass
type AttrClass struct {
	P unsafe.Pointer
}

const SizeOfStructAttrClass = 32

func AttrClassGetType() gi.GType {
	ret := _I.GetGType(2, "AttrClass")
	return ret
}

// Struct AttrColor
type AttrColor struct {
	P unsafe.Pointer
}

const SizeOfStructAttrColor = 24

func AttrColorGetType() gi.GType {
	ret := _I.GetGType(3, "AttrColor")
	return ret
}

// Struct AttrFloat
type AttrFloat struct {
	P unsafe.Pointer
}

const SizeOfStructAttrFloat = 24

func AttrFloatGetType() gi.GType {
	ret := _I.GetGType(4, "AttrFloat")
	return ret
}

// Struct AttrFontDesc
type AttrFontDesc struct {
	P unsafe.Pointer
}

const SizeOfStructAttrFontDesc = 24

func AttrFontDescGetType() gi.GType {
	ret := _I.GetGType(5, "AttrFontDesc")
	return ret
}

// Struct AttrFontFeatures
type AttrFontFeatures struct {
	P unsafe.Pointer
}

const SizeOfStructAttrFontFeatures = 24

func AttrFontFeaturesGetType() gi.GType {
	ret := _I.GetGType(6, "AttrFontFeatures")
	return ret
}

// Struct AttrInt
type AttrInt struct {
	P unsafe.Pointer
}

const SizeOfStructAttrInt = 24

func AttrIntGetType() gi.GType {
	ret := _I.GetGType(7, "AttrInt")
	return ret
}

// Struct AttrIterator
type AttrIterator struct {
	P unsafe.Pointer
}

func AttrIteratorGetType() gi.GType {
	ret := _I.GetGType(8, "AttrIterator")
	return ret
}

// pango_attr_iterator_get_attrs
// container is not nil, container is AttrIterator
// is method
func (v AttrIterator) GetAttrs() (result glib.SList) {
	iv, err := _I.Get(0, "AttrIterator", "get_attrs")
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

// pango_attr_iterator_get_font
// container is not nil, container is AttrIterator
// is method
func (v AttrIterator) GetFont(desc FontDescription, language Language, extra_attrs glib.SList) {
	iv, err := _I.Get(1, "AttrIterator", "get_font")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_desc := gi.NewPointerArgument(desc.P)
	arg_language := gi.NewPointerArgument(language.P)
	arg_extra_attrs := gi.NewPointerArgument(extra_attrs.P)
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

const SizeOfStructAttrLanguage = 24

func AttrLanguageGetType() gi.GType {
	ret := _I.GetGType(9, "AttrLanguage")
	return ret
}

// Struct AttrList
type AttrList struct {
	P unsafe.Pointer
}

func AttrListGetType() gi.GType {
	ret := _I.GetGType(10, "AttrList")
	return ret
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

const SizeOfStructAttrShape = 72

func AttrShapeGetType() gi.GType {
	ret := _I.GetGType(11, "AttrShape")
	return ret
}

// Struct AttrSize
type AttrSize struct {
	P unsafe.Pointer
}

const SizeOfStructAttrSize = 24

func AttrSizeGetType() gi.GType {
	ret := _I.GetGType(12, "AttrSize")
	return ret
}

// Struct AttrString
type AttrString struct {
	P unsafe.Pointer
}

const SizeOfStructAttrString = 24

func AttrStringGetType() gi.GType {
	ret := _I.GetGType(13, "AttrString")
	return ret
}

// Enum AttrType
type AttrTypeEnum int

const (
	AttrTypeInvalid            AttrTypeEnum = 0
	AttrTypeLanguage           AttrTypeEnum = 1
	AttrTypeFamily             AttrTypeEnum = 2
	AttrTypeStyle              AttrTypeEnum = 3
	AttrTypeWeight             AttrTypeEnum = 4
	AttrTypeVariant            AttrTypeEnum = 5
	AttrTypeStretch            AttrTypeEnum = 6
	AttrTypeSize               AttrTypeEnum = 7
	AttrTypeFontDesc           AttrTypeEnum = 8
	AttrTypeForeground         AttrTypeEnum = 9
	AttrTypeBackground         AttrTypeEnum = 10
	AttrTypeUnderline          AttrTypeEnum = 11
	AttrTypeStrikethrough      AttrTypeEnum = 12
	AttrTypeRise               AttrTypeEnum = 13
	AttrTypeShape              AttrTypeEnum = 14
	AttrTypeScale              AttrTypeEnum = 15
	AttrTypeFallback           AttrTypeEnum = 16
	AttrTypeLetterSpacing      AttrTypeEnum = 17
	AttrTypeUnderlineColor     AttrTypeEnum = 18
	AttrTypeStrikethroughColor AttrTypeEnum = 19
	AttrTypeAbsoluteSize       AttrTypeEnum = 20
	AttrTypeGravity            AttrTypeEnum = 21
	AttrTypeGravityHint        AttrTypeEnum = 22
	AttrTypeFontFeatures       AttrTypeEnum = 23
	AttrTypeForegroundAlpha    AttrTypeEnum = 24
	AttrTypeBackgroundAlpha    AttrTypeEnum = 25
)

func AttrTypeGetType() gi.GType {
	ret := _I.GetGType(14, "AttrType")
	return ret
}

// Struct Attribute
type Attribute struct {
	P unsafe.Pointer
}

const SizeOfStructAttribute = 16

func AttributeGetType() gi.GType {
	ret := _I.GetGType(15, "Attribute")
	return ret
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

// Enum BidiType
type BidiTypeEnum int

const (
	BidiTypeL   BidiTypeEnum = 0
	BidiTypeLre BidiTypeEnum = 1
	BidiTypeLro BidiTypeEnum = 2
	BidiTypeR   BidiTypeEnum = 3
	BidiTypeAl  BidiTypeEnum = 4
	BidiTypeRle BidiTypeEnum = 5
	BidiTypeRlo BidiTypeEnum = 6
	BidiTypePdf BidiTypeEnum = 7
	BidiTypeEn  BidiTypeEnum = 8
	BidiTypeEs  BidiTypeEnum = 9
	BidiTypeEt  BidiTypeEnum = 10
	BidiTypeAn  BidiTypeEnum = 11
	BidiTypeCs  BidiTypeEnum = 12
	BidiTypeNsm BidiTypeEnum = 13
	BidiTypeBn  BidiTypeEnum = 14
	BidiTypeB   BidiTypeEnum = 15
	BidiTypeS   BidiTypeEnum = 16
	BidiTypeWs  BidiTypeEnum = 17
	BidiTypeOn  BidiTypeEnum = 18
)

func BidiTypeGetType() gi.GType {
	ret := _I.GetGType(16, "BidiType")
	return ret
}

// Struct Color
type Color struct {
	P unsafe.Pointer
}

const SizeOfStructColor = 6

func ColorGetType() gi.GType {
	ret := _I.GetGType(17, "Color")
	return ret
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
	gi.Free(c_spec)
	result = ret.Bool()
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

type IContext interface{ P_Context() unsafe.Pointer }

func (v Context) P_Context() unsafe.Pointer { return v.P }
func ContextGetType() gi.GType {
	ret := _I.GetGType(18, "Context")
	return ret
}

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
func (v Context) GetBaseDir() (result DirectionEnum) {
	iv, err := _I.Get(22, "Context", "get_base_dir")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = DirectionEnum(ret.Int())
	return
}

// pango_context_get_base_gravity
// container is not nil, container is Context
// is method
func (v Context) GetBaseGravity() (result GravityEnum) {
	iv, err := _I.Get(23, "Context", "get_base_gravity")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = GravityEnum(ret.Int())
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
func (v Context) GetGravity() (result GravityEnum) {
	iv, err := _I.Get(26, "Context", "get_gravity")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = GravityEnum(ret.Int())
	return
}

// pango_context_get_gravity_hint
// container is not nil, container is Context
// is method
func (v Context) GetGravityHint() (result GravityHintEnum) {
	iv, err := _I.Get(27, "Context", "get_gravity_hint")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = GravityHintEnum(ret.Int())
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
// arg 0 families lenArgIdx 1
func (v Context) ListFamilies() (families gi.PointerArray) {
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
	var n_families int32
	_ = n_families
	families.P = outArgs[0].Pointer()
	n_families = outArgs[1].Int32()
	families.Len = int(n_families)
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
func (v Context) SetBaseDir(direction DirectionEnum) {
	iv, err := _I.Get(35, "Context", "set_base_dir")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_direction := gi.NewIntArgument(int(direction))
	args := []gi.Argument{arg_v, arg_direction}
	iv.Call(args, nil, nil)
}

// pango_context_set_base_gravity
// container is not nil, container is Context
// is method
func (v Context) SetBaseGravity(gravity GravityEnum) {
	iv, err := _I.Get(36, "Context", "set_base_gravity")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_gravity := gi.NewIntArgument(int(gravity))
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
func (v Context) SetFontMap(font_map IFontMap) {
	iv, err := _I.Get(38, "Context", "set_font_map")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if font_map != nil {
		tmp = font_map.P_FontMap()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_font_map := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_font_map}
	iv.Call(args, nil, nil)
}

// pango_context_set_gravity_hint
// container is not nil, container is Context
// is method
func (v Context) SetGravityHint(hint GravityHintEnum) {
	iv, err := _I.Get(39, "Context", "set_gravity_hint")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_hint := gi.NewIntArgument(int(hint))
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

func CoverageGetType() gi.GType {
	ret := _I.GetGType(19, "Coverage")
	return ret
}

// pango_coverage_get
// container is not nil, container is Coverage
// is method
func (v Coverage) Get(index_ int32) (result CoverageLevelEnum) {
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
	result = CoverageLevelEnum(ret.Int())
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
func (v Coverage) Set(index_ int32, level CoverageLevelEnum) {
	iv, err := _I.Get(44, "Coverage", "set")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewInt32Argument(index_)
	arg_level := gi.NewIntArgument(int(level))
	args := []gi.Argument{arg_v, arg_index_, arg_level}
	iv.Call(args, nil, nil)
}

// pango_coverage_to_bytes
// container is not nil, container is Coverage
// is method
// arg 0 bytes lenArgIdx 1
func (v Coverage) ToBytes() (bytes gi.Uint8Array) {
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
	var n_bytes int32
	_ = n_bytes
	bytes.P = outArgs[0].Pointer()
	n_bytes = outArgs[1].Int32()
	bytes.Len = int(n_bytes)
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

// Enum CoverageLevel
type CoverageLevelEnum int

const (
	CoverageLevelNone        CoverageLevelEnum = 0
	CoverageLevelFallback    CoverageLevelEnum = 1
	CoverageLevelApproximate CoverageLevelEnum = 2
	CoverageLevelExact       CoverageLevelEnum = 3
)

func CoverageLevelGetType() gi.GType {
	ret := _I.GetGType(20, "CoverageLevel")
	return ret
}

// Enum Direction
type DirectionEnum int

const (
	DirectionLtr     DirectionEnum = 0
	DirectionRtl     DirectionEnum = 1
	DirectionTtbLtr  DirectionEnum = 2
	DirectionTtbRtl  DirectionEnum = 3
	DirectionWeakLtr DirectionEnum = 4
	DirectionWeakRtl DirectionEnum = 5
	DirectionNeutral DirectionEnum = 6
)

func DirectionGetType() gi.GType {
	ret := _I.GetGType(21, "Direction")
	return ret
}

// Enum EllipsizeMode
type EllipsizeModeEnum int

const (
	EllipsizeModeNone   EllipsizeModeEnum = 0
	EllipsizeModeStart  EllipsizeModeEnum = 1
	EllipsizeModeMiddle EllipsizeModeEnum = 2
	EllipsizeModeEnd    EllipsizeModeEnum = 3
)

func EllipsizeModeGetType() gi.GType {
	ret := _I.GetGType(22, "EllipsizeMode")
	return ret
}

// Object Engine
type Engine struct {
	gobject.Object
}

func WrapEngine(p unsafe.Pointer) (r Engine) { r.P = p; return }

type IEngine interface{ P_Engine() unsafe.Pointer }

func (v Engine) P_Engine() unsafe.Pointer { return v.P }
func EngineGetType() gi.GType {
	ret := _I.GetGType(23, "Engine")
	return ret
}

// ignore GType struct EngineClass
// Struct EngineInfo
type EngineInfo struct {
	P unsafe.Pointer
}

const SizeOfStructEngineInfo = 40

func EngineInfoGetType() gi.GType {
	ret := _I.GetGType(24, "EngineInfo")
	return ret
}

// Object EngineLang
type EngineLang struct {
	Engine
}

func WrapEngineLang(p unsafe.Pointer) (r EngineLang) { r.P = p; return }

type IEngineLang interface{ P_EngineLang() unsafe.Pointer }

func (v EngineLang) P_EngineLang() unsafe.Pointer { return v.P }
func EngineLangGetType() gi.GType {
	ret := _I.GetGType(25, "EngineLang")
	return ret
}

// ignore GType struct EngineLangClass
// Struct EngineScriptInfo
type EngineScriptInfo struct {
	P unsafe.Pointer
}

const SizeOfStructEngineScriptInfo = 16

func EngineScriptInfoGetType() gi.GType {
	ret := _I.GetGType(26, "EngineScriptInfo")
	return ret
}

// Object EngineShape
type EngineShape struct {
	Engine
}

func WrapEngineShape(p unsafe.Pointer) (r EngineShape) { r.P = p; return }

type IEngineShape interface{ P_EngineShape() unsafe.Pointer }

func (v EngineShape) P_EngineShape() unsafe.Pointer { return v.P }
func EngineShapeGetType() gi.GType {
	ret := _I.GetGType(27, "EngineShape")
	return ret
}

// ignore GType struct EngineShapeClass
// Object Font
type Font struct {
	gobject.Object
}

func WrapFont(p unsafe.Pointer) (r Font) { r.P = p; return }

type IFont interface{ P_Font() unsafe.Pointer }

func (v Font) P_Font() unsafe.Pointer { return v.P }
func FontGetType() gi.GType {
	ret := _I.GetGType(28, "Font")
	return ret
}

// pango_font_descriptions_free
// container is not nil, container is Font
// is method
// arg0Type tag: array, isPtr: true
// arg 0 descs lenArgIdx 1
func FontDescriptionsFree1(descs gi.PointerArray, n_descs int32) {
	iv, err := _I.Get(47, "Font", "descriptions_free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_descs := gi.NewPointerArgument(descs.P)
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
func (v Font) GetGlyphExtents(glyph uint32, ink_rect Rectangle, logical_rect Rectangle) {
	iv, err := _I.Get(52, "Font", "get_glyph_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_glyph := gi.NewUint32Argument(glyph)
	arg_ink_rect := gi.NewPointerArgument(ink_rect.P)
	arg_logical_rect := gi.NewPointerArgument(logical_rect.P)
	args := []gi.Argument{arg_v, arg_glyph, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, nil)
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

func FontDescriptionGetType() gi.GType {
	ret := _I.GetGType(29, "FontDescription")
	return ret
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
func (v FontDescription) GetGravity() (result GravityEnum) {
	iv, err := _I.Get(61, "FontDescription", "get_gravity")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = GravityEnum(ret.Int())
	return
}

// pango_font_description_get_set_fields
// container is not nil, container is FontDescription
// is method
func (v FontDescription) GetSetFields() (result FontMaskFlags) {
	iv, err := _I.Get(62, "FontDescription", "get_set_fields")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FontMaskFlags(ret.Int())
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
func (v FontDescription) GetStretch() (result StretchEnum) {
	iv, err := _I.Get(65, "FontDescription", "get_stretch")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StretchEnum(ret.Int())
	return
}

// pango_font_description_get_style
// container is not nil, container is FontDescription
// is method
func (v FontDescription) GetStyle() (result StyleEnum) {
	iv, err := _I.Get(66, "FontDescription", "get_style")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = StyleEnum(ret.Int())
	return
}

// pango_font_description_get_variant
// container is not nil, container is FontDescription
// is method
func (v FontDescription) GetVariant() (result VariantEnum) {
	iv, err := _I.Get(67, "FontDescription", "get_variant")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = VariantEnum(ret.Int())
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
func (v FontDescription) GetWeight() (result WeightEnum) {
	iv, err := _I.Get(69, "FontDescription", "get_weight")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = WeightEnum(ret.Int())
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
func (v FontDescription) SetGravity(gravity GravityEnum) {
	iv, err := _I.Get(76, "FontDescription", "set_gravity")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_gravity := gi.NewIntArgument(int(gravity))
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
func (v FontDescription) SetStretch(stretch StretchEnum) {
	iv, err := _I.Get(78, "FontDescription", "set_stretch")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_stretch := gi.NewIntArgument(int(stretch))
	args := []gi.Argument{arg_v, arg_stretch}
	iv.Call(args, nil, nil)
}

// pango_font_description_set_style
// container is not nil, container is FontDescription
// is method
func (v FontDescription) SetStyle(style StyleEnum) {
	iv, err := _I.Get(79, "FontDescription", "set_style")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_style := gi.NewIntArgument(int(style))
	args := []gi.Argument{arg_v, arg_style}
	iv.Call(args, nil, nil)
}

// pango_font_description_set_variant
// container is not nil, container is FontDescription
// is method
func (v FontDescription) SetVariant(variant VariantEnum) {
	iv, err := _I.Get(80, "FontDescription", "set_variant")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_variant := gi.NewIntArgument(int(variant))
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
func (v FontDescription) SetWeight(weight WeightEnum) {
	iv, err := _I.Get(83, "FontDescription", "set_weight")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_weight := gi.NewIntArgument(int(weight))
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
func (v FontDescription) UnsetFields(to_unset FontMaskFlags) {
	iv, err := _I.Get(86, "FontDescription", "unset_fields")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_to_unset := gi.NewIntArgument(int(to_unset))
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
	gi.Free(c_str)
	result.P = ret.Pointer()
	return
}

// Object FontFace
type FontFace struct {
	gobject.Object
}

func WrapFontFace(p unsafe.Pointer) (r FontFace) { r.P = p; return }

type IFontFace interface{ P_FontFace() unsafe.Pointer }

func (v FontFace) P_FontFace() unsafe.Pointer { return v.P }
func FontFaceGetType() gi.GType {
	ret := _I.GetGType(30, "FontFace")
	return ret
}

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
// arg 0 sizes lenArgIdx 1
func (v FontFace) ListSizes() (sizes gi.Int32Array) {
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
	var n_sizes int32
	_ = n_sizes
	sizes.P = outArgs[0].Pointer()
	n_sizes = outArgs[1].Int32()
	sizes.Len = int(n_sizes)
	return
}

// ignore GType struct FontFaceClass
// Object FontFamily
type FontFamily struct {
	gobject.Object
}

func WrapFontFamily(p unsafe.Pointer) (r FontFamily) { r.P = p; return }

type IFontFamily interface{ P_FontFamily() unsafe.Pointer }

func (v FontFamily) P_FontFamily() unsafe.Pointer { return v.P }
func FontFamilyGetType() gi.GType {
	ret := _I.GetGType(31, "FontFamily")
	return ret
}

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
// arg 0 faces lenArgIdx 1
func (v FontFamily) ListFaces() (faces gi.PointerArray) {
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
	var n_faces int32
	_ = n_faces
	faces.P = outArgs[0].Pointer()
	n_faces = outArgs[1].Int32()
	faces.Len = int(n_faces)
	return
}

// ignore GType struct FontFamilyClass
// Object FontMap
type FontMap struct {
	gobject.Object
}

func WrapFontMap(p unsafe.Pointer) (r FontMap) { r.P = p; return }

type IFontMap interface{ P_FontMap() unsafe.Pointer }

func (v FontMap) P_FontMap() unsafe.Pointer { return v.P }
func FontMapGetType() gi.GType {
	ret := _I.GetGType(32, "FontMap")
	return ret
}

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
// arg 0 families lenArgIdx 1
func (v FontMap) ListFamilies() (families gi.PointerArray) {
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
	var n_families int32
	_ = n_families
	families.P = outArgs[0].Pointer()
	n_families = outArgs[1].Int32()
	families.Len = int(n_families)
	return
}

// pango_font_map_load_font
// container is not nil, container is FontMap
// is method
func (v FontMap) LoadFont(context IContext, desc FontDescription) (result Font) {
	iv, err := _I.Get(100, "FontMap", "load_font")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_Context()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_context := gi.NewPointerArgument(tmp)
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
func (v FontMap) LoadFontset(context IContext, desc FontDescription, language Language) (result Fontset) {
	iv, err := _I.Get(101, "FontMap", "load_fontset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_Context()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_context := gi.NewPointerArgument(tmp)
	arg_desc := gi.NewPointerArgument(desc.P)
	arg_language := gi.NewPointerArgument(language.P)
	args := []gi.Argument{arg_v, arg_context, arg_desc, arg_language}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct FontMapClass
// Flags FontMask
type FontMaskFlags int

const (
	FontMaskFamily     FontMaskFlags = 1
	FontMaskStyle      FontMaskFlags = 2
	FontMaskVariant    FontMaskFlags = 4
	FontMaskWeight     FontMaskFlags = 8
	FontMaskStretch    FontMaskFlags = 16
	FontMaskSize       FontMaskFlags = 32
	FontMaskGravity    FontMaskFlags = 64
	FontMaskVariations FontMaskFlags = 128
)

func FontMaskGetType() gi.GType {
	ret := _I.GetGType(33, "FontMask")
	return ret
}

// Struct FontMetrics
type FontMetrics struct {
	P unsafe.Pointer
}

const SizeOfStructFontMetrics = 36

func FontMetricsGetType() gi.GType {
	ret := _I.GetGType(34, "FontMetrics")
	return ret
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

type IFontset interface{ P_Fontset() unsafe.Pointer }

func (v Fontset) P_Fontset() unsafe.Pointer { return v.P }
func FontsetGetType() gi.GType {
	ret := _I.GetGType(35, "Fontset")
	return ret
}

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

type IFontsetSimple interface{ P_FontsetSimple() unsafe.Pointer }

func (v FontsetSimple) P_FontsetSimple() unsafe.Pointer { return v.P }
func FontsetSimpleGetType() gi.GType {
	ret := _I.GetGType(36, "FontsetSimple")
	return ret
}

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
func (v FontsetSimple) Append(font IFont) {
	iv, err := _I.Get(117, "FontsetSimple", "append")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if font != nil {
		tmp = font.P_Font()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_font := gi.NewPointerArgument(tmp)
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

const SizeOfStructGlyphGeometry = 12

func GlyphGeometryGetType() gi.GType {
	ret := _I.GetGType(37, "GlyphGeometry")
	return ret
}

// Struct GlyphInfo
type GlyphInfo struct {
	P unsafe.Pointer
}

const SizeOfStructGlyphInfo = 20

func GlyphInfoGetType() gi.GType {
	ret := _I.GetGType(38, "GlyphInfo")
	return ret
}

// Struct GlyphItem
type GlyphItem struct {
	P unsafe.Pointer
}

const SizeOfStructGlyphItem = 16

func GlyphItemGetType() gi.GType {
	ret := _I.GetGType(39, "GlyphItem")
	return ret
}

// pango_glyph_item_apply_attrs
// container is not nil, container is GlyphItem
// is method
func (v GlyphItem) ApplyAttrs(text string, list AttrList) (result glib.SList) {
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
	gi.Free(c_text)
	result.P = ret.Pointer()
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
func (v GlyphItem) GetLogicalWidths(text string, logical_widths gi.Int32Array) {
	iv, err := _I.Get(122, "GlyphItem", "get_logical_widths")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_logical_widths := gi.NewPointerArgument(logical_widths.P)
	args := []gi.Argument{arg_v, arg_text, arg_logical_widths}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// pango_glyph_item_letter_space
// container is not nil, container is GlyphItem
// is method
func (v GlyphItem) LetterSpace(text string, log_attrs unsafe.Pointer, letter_spacing int32) {
	iv, err := _I.Get(123, "GlyphItem", "letter_space")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_log_attrs := gi.NewPointerArgument(log_attrs)
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
	gi.Free(c_text)
	result.P = ret.Pointer()
	return
}

// Struct GlyphItemIter
type GlyphItemIter struct {
	P unsafe.Pointer
}

const SizeOfStructGlyphItemIter = 40

func GlyphItemIterGetType() gi.GType {
	ret := _I.GetGType(40, "GlyphItemIter")
	return ret
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
	gi.Free(c_text)
	result = ret.Bool()
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
	gi.Free(c_text)
	result = ret.Bool()
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

const SizeOfStructGlyphString = 32

func GlyphStringGetType() gi.GType {
	ret := _I.GetGType(41, "GlyphString")
	return ret
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
func (v GlyphString) Extents(font IFont, ink_rect Rectangle, logical_rect Rectangle) {
	iv, err := _I.Get(133, "GlyphString", "extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if font != nil {
		tmp = font.P_Font()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_font := gi.NewPointerArgument(tmp)
	arg_ink_rect := gi.NewPointerArgument(ink_rect.P)
	arg_logical_rect := gi.NewPointerArgument(logical_rect.P)
	args := []gi.Argument{arg_v, arg_font, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, nil)
}

// pango_glyph_string_extents_range
// container is not nil, container is GlyphString
// is method
func (v GlyphString) ExtentsRange(start int32, end int32, font IFont, ink_rect Rectangle, logical_rect Rectangle) {
	iv, err := _I.Get(134, "GlyphString", "extents_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if font != nil {
		tmp = font.P_Font()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewInt32Argument(start)
	arg_end := gi.NewInt32Argument(end)
	arg_font := gi.NewPointerArgument(tmp)
	arg_ink_rect := gi.NewPointerArgument(ink_rect.P)
	arg_logical_rect := gi.NewPointerArgument(logical_rect.P)
	args := []gi.Argument{arg_v, arg_start, arg_end, arg_font, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, nil)
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
func (v GlyphString) GetLogicalWidths(text string, length int32, embedding_level int32, logical_widths gi.Int32Array) {
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
	arg_logical_widths := gi.NewPointerArgument(logical_widths.P)
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

const SizeOfStructGlyphVisAttr = 4

func GlyphVisAttrGetType() gi.GType {
	ret := _I.GetGType(42, "GlyphVisAttr")
	return ret
}

// Enum Gravity
type GravityEnum int

const (
	GravitySouth GravityEnum = 0
	GravityEast  GravityEnum = 1
	GravityNorth GravityEnum = 2
	GravityWest  GravityEnum = 3
	GravityAuto  GravityEnum = 4
)

func GravityGetType() gi.GType {
	ret := _I.GetGType(43, "Gravity")
	return ret
}

// Enum GravityHint
type GravityHintEnum int

const (
	GravityHintNatural GravityHintEnum = 0
	GravityHintStrong  GravityHintEnum = 1
	GravityHintLine    GravityHintEnum = 2
)

func GravityHintGetType() gi.GType {
	ret := _I.GetGType(44, "GravityHint")
	return ret
}

// Struct IncludedModule
type IncludedModule struct {
	P unsafe.Pointer
}

const SizeOfStructIncludedModule = 32

func IncludedModuleGetType() gi.GType {
	ret := _I.GetGType(45, "IncludedModule")
	return ret
}

// Struct Item
type Item struct {
	P unsafe.Pointer
}

const SizeOfStructItem = 64

func ItemGetType() gi.GType {
	ret := _I.GetGType(46, "Item")
	return ret
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

func LanguageGetType() gi.GType {
	ret := _I.GetGType(47, "Language")
	return ret
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
// ret lenArgIdx 0
func (v Language) GetScripts() (result unsafe.Pointer) {
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
	var num_scripts int32
	_ = num_scripts
	num_scripts = outArgs[0].Int32()
	result = ret.Pointer()
	return
}

// pango_language_includes_script
// container is not nil, container is Language
// is method
func (v Language) IncludesScript(script ScriptEnum) (result bool) {
	iv, err := _I.Get(147, "Language", "includes_script")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_script := gi.NewIntArgument(int(script))
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
	gi.Free(c_range_list)
	result = ret.Bool()
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
	gi.Free(c_language)
	result.P = ret.Pointer()
	return
}

// Object Layout
type Layout struct {
	gobject.Object
}

func WrapLayout(p unsafe.Pointer) (r Layout) { r.P = p; return }

type ILayout interface{ P_Layout() unsafe.Pointer }

func (v Layout) P_Layout() unsafe.Pointer { return v.P }
func LayoutGetType() gi.GType {
	ret := _I.GetGType(48, "Layout")
	return ret
}

// pango_layout_new
// container is not nil, container is Layout
// is constructor
func NewLayout(context IContext) (result Layout) {
	iv, err := _I.Get(152, "Layout", "new")
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
func (v Layout) GetAlignment() (result AlignmentEnum) {
	iv, err := _I.Get(155, "Layout", "get_alignment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = AlignmentEnum(ret.Int())
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
func (v Layout) GetCursorPos(index_ int32, strong_pos Rectangle, weak_pos Rectangle) {
	iv, err := _I.Get(161, "Layout", "get_cursor_pos")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewInt32Argument(index_)
	arg_strong_pos := gi.NewPointerArgument(strong_pos.P)
	arg_weak_pos := gi.NewPointerArgument(weak_pos.P)
	args := []gi.Argument{arg_v, arg_index_, arg_strong_pos, arg_weak_pos}
	iv.Call(args, nil, nil)
}

// pango_layout_get_ellipsize
// container is not nil, container is Layout
// is method
func (v Layout) GetEllipsize() (result EllipsizeModeEnum) {
	iv, err := _I.Get(162, "Layout", "get_ellipsize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = EllipsizeModeEnum(ret.Int())
	return
}

// pango_layout_get_extents
// container is not nil, container is Layout
// is method
func (v Layout) GetExtents(ink_rect Rectangle, logical_rect Rectangle) {
	iv, err := _I.Get(163, "Layout", "get_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(ink_rect.P)
	arg_logical_rect := gi.NewPointerArgument(logical_rect.P)
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, nil)
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
func (v Layout) GetLines() (result glib.SList) {
	iv, err := _I.Get(172, "Layout", "get_lines")
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

// pango_layout_get_lines_readonly
// container is not nil, container is Layout
// is method
func (v Layout) GetLinesReadonly() (result glib.SList) {
	iv, err := _I.Get(173, "Layout", "get_lines_readonly")
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

// pango_layout_get_log_attrs
// container is not nil, container is Layout
// is method
// arg 0 attrs lenArgIdx 1
func (v Layout) GetLogAttrs() (attrs unsafe.Pointer) {
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
	var n_attrs int32
	_ = n_attrs
	attrs = outArgs[0].Pointer()
	n_attrs = outArgs[1].Int32()
	return
}

// pango_layout_get_log_attrs_readonly
// container is not nil, container is Layout
// is method
// ret lenArgIdx 0
func (v Layout) GetLogAttrsReadonly() (result unsafe.Pointer) {
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
	var n_attrs int32
	_ = n_attrs
	n_attrs = outArgs[0].Int32()
	result = ret.Pointer()
	return
}

// pango_layout_get_pixel_extents
// container is not nil, container is Layout
// is method
func (v Layout) GetPixelExtents(ink_rect Rectangle, logical_rect Rectangle) {
	iv, err := _I.Get(176, "Layout", "get_pixel_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(ink_rect.P)
	arg_logical_rect := gi.NewPointerArgument(logical_rect.P)
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, nil)
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
func (v Layout) GetWrap() (result WrapModeEnum) {
	iv, err := _I.Get(186, "Layout", "get_wrap")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = WrapModeEnum(ret.Int())
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
func (v Layout) IndexToPos(index_ int32, pos Rectangle) {
	iv, err := _I.Get(188, "Layout", "index_to_pos")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewInt32Argument(index_)
	arg_pos := gi.NewPointerArgument(pos.P)
	args := []gi.Argument{arg_v, arg_index_, arg_pos}
	iv.Call(args, nil, nil)
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
func (v Layout) SetAlignment(alignment AlignmentEnum) {
	iv, err := _I.Get(192, "Layout", "set_alignment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_alignment := gi.NewIntArgument(int(alignment))
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
func (v Layout) SetEllipsize(ellipsize EllipsizeModeEnum) {
	iv, err := _I.Get(195, "Layout", "set_ellipsize")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ellipsize := gi.NewIntArgument(int(ellipsize))
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
func (v Layout) SetMarkupWithAccel(markup string, length int32, accel_marker rune) (accel_char rune) {
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
	arg_accel_marker := gi.NewUint32Argument(uint32(accel_marker))
	arg_accel_char := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_markup, arg_length, arg_accel_marker, arg_accel_char}
	iv.Call(args, nil, &outArgs[0])
	gi.Free(c_markup)
	accel_char = rune(outArgs[0].Uint32())
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
func (v Layout) SetWrap(wrap WrapModeEnum) {
	iv, err := _I.Get(207, "Layout", "set_wrap")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_wrap := gi.NewIntArgument(int(wrap))
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
	index_ = outArgs[0].Int32()
	trailing = outArgs[1].Int32()
	result = ret.Bool()
	return
}

// ignore GType struct LayoutClass
// Struct LayoutIter
type LayoutIter struct {
	P unsafe.Pointer
}

func LayoutIterGetType() gi.GType {
	ret := _I.GetGType(49, "LayoutIter")
	return ret
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
func (v LayoutIter) GetCharExtents(logical_rect Rectangle) {
	iv, err := _I.Get(213, "LayoutIter", "get_char_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_logical_rect := gi.NewPointerArgument(logical_rect.P)
	args := []gi.Argument{arg_v, arg_logical_rect}
	iv.Call(args, nil, nil)
}

// pango_layout_iter_get_cluster_extents
// container is not nil, container is LayoutIter
// is method
func (v LayoutIter) GetClusterExtents(ink_rect Rectangle, logical_rect Rectangle) {
	iv, err := _I.Get(214, "LayoutIter", "get_cluster_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(ink_rect.P)
	arg_logical_rect := gi.NewPointerArgument(logical_rect.P)
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, nil)
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
func (v LayoutIter) GetLayoutExtents(ink_rect Rectangle, logical_rect Rectangle) {
	iv, err := _I.Get(217, "LayoutIter", "get_layout_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(ink_rect.P)
	arg_logical_rect := gi.NewPointerArgument(logical_rect.P)
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, nil)
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
func (v LayoutIter) GetLineExtents(ink_rect Rectangle, logical_rect Rectangle) {
	iv, err := _I.Get(219, "LayoutIter", "get_line_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(ink_rect.P)
	arg_logical_rect := gi.NewPointerArgument(logical_rect.P)
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, nil)
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
func (v LayoutIter) GetRunExtents(ink_rect Rectangle, logical_rect Rectangle) {
	iv, err := _I.Get(223, "LayoutIter", "get_run_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(ink_rect.P)
	arg_logical_rect := gi.NewPointerArgument(logical_rect.P)
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, nil)
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

const SizeOfStructLayoutLine = 32

func LayoutLineGetType() gi.GType {
	ret := _I.GetGType(50, "LayoutLine")
	return ret
}

// pango_layout_line_get_extents
// container is not nil, container is LayoutLine
// is method
func (v LayoutLine) GetExtents(ink_rect Rectangle, logical_rect Rectangle) {
	iv, err := _I.Get(229, "LayoutLine", "get_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(ink_rect.P)
	arg_logical_rect := gi.NewPointerArgument(logical_rect.P)
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, nil)
}

// pango_layout_line_get_pixel_extents
// container is not nil, container is LayoutLine
// is method
func (v LayoutLine) GetPixelExtents(ink_rect Rectangle, logical_rect Rectangle) {
	iv, err := _I.Get(230, "LayoutLine", "get_pixel_extents")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ink_rect := gi.NewPointerArgument(ink_rect.P)
	arg_logical_rect := gi.NewPointerArgument(logical_rect.P)
	args := []gi.Argument{arg_v, arg_ink_rect, arg_logical_rect}
	iv.Call(args, nil, nil)
}

// pango_layout_line_get_x_ranges
// container is not nil, container is LayoutLine
// is method
// arg 2 ranges lenArgIdx 3
func (v LayoutLine) GetXRanges(start_index int32, end_index int32) (ranges gi.Int32Array) {
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
	var n_ranges int32
	_ = n_ranges
	ranges.P = outArgs[0].Pointer()
	n_ranges = outArgs[1].Int32()
	ranges.Len = int(n_ranges)
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
	index_ = outArgs[0].Int32()
	trailing = outArgs[1].Int32()
	result = ret.Bool()
	return
}

// Struct LogAttr
type LogAttr struct {
	P unsafe.Pointer
}

const SizeOfStructLogAttr = 52

func LogAttrGetType() gi.GType {
	ret := _I.GetGType(51, "LogAttr")
	return ret
}

// Struct Map
type Map struct {
	P unsafe.Pointer
}

func MapGetType() gi.GType {
	ret := _I.GetGType(52, "Map")
	return ret
}

// Struct MapEntry
type MapEntry struct {
	P unsafe.Pointer
}

func MapEntryGetType() gi.GType {
	ret := _I.GetGType(53, "MapEntry")
	return ret
}

// Struct Matrix
type Matrix struct {
	P unsafe.Pointer
}

const SizeOfStructMatrix = 48

func MatrixGetType() gi.GType {
	ret := _I.GetGType(54, "Matrix")
	return ret
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
func (v Matrix) TransformDistance(dx int /*TODO:TYPE*/, dy int /*TODO:TYPE*/) {
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
func (v Matrix) TransformPixelRectangle(rect int /*TODO:TYPE*/) {
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
func (v Matrix) TransformPoint(x int /*TODO:TYPE*/, y int /*TODO:TYPE*/) {
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
func (v Matrix) TransformRectangle(rect int /*TODO:TYPE*/) {
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

const SizeOfStructRectangle = 16

func RectangleGetType() gi.GType {
	ret := _I.GetGType(55, "Rectangle")
	return ret
}

// Enum RenderPart
type RenderPartEnum int

const (
	RenderPartForeground    RenderPartEnum = 0
	RenderPartBackground    RenderPartEnum = 1
	RenderPartUnderline     RenderPartEnum = 2
	RenderPartStrikethrough RenderPartEnum = 3
)

func RenderPartGetType() gi.GType {
	ret := _I.GetGType(56, "RenderPart")
	return ret
}

// Object Renderer
type Renderer struct {
	gobject.Object
}

func WrapRenderer(p unsafe.Pointer) (r Renderer) { r.P = p; return }

type IRenderer interface{ P_Renderer() unsafe.Pointer }

func (v Renderer) P_Renderer() unsafe.Pointer { return v.P }
func RendererGetType() gi.GType {
	ret := _I.GetGType(57, "Renderer")
	return ret
}

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
func (v Renderer) DrawGlyph(font IFont, glyph uint32, x float64, y float64) {
	iv, err := _I.Get(251, "Renderer", "draw_glyph")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if font != nil {
		tmp = font.P_Font()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_font := gi.NewPointerArgument(tmp)
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
func (v Renderer) DrawGlyphs(font IFont, glyphs GlyphString, x int32, y int32) {
	iv, err := _I.Get(253, "Renderer", "draw_glyphs")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if font != nil {
		tmp = font.P_Font()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_font := gi.NewPointerArgument(tmp)
	arg_glyphs := gi.NewPointerArgument(glyphs.P)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	args := []gi.Argument{arg_v, arg_font, arg_glyphs, arg_x, arg_y}
	iv.Call(args, nil, nil)
}

// pango_renderer_draw_layout
// container is not nil, container is Renderer
// is method
func (v Renderer) DrawLayout(layout ILayout, x int32, y int32) {
	iv, err := _I.Get(254, "Renderer", "draw_layout")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if layout != nil {
		tmp = layout.P_Layout()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_layout := gi.NewPointerArgument(tmp)
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
func (v Renderer) DrawRectangle(part RenderPartEnum, x int32, y int32, width int32, height int32) {
	iv, err := _I.Get(256, "Renderer", "draw_rectangle")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_part := gi.NewIntArgument(int(part))
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
func (v Renderer) DrawTrapezoid(part RenderPartEnum, y1_ float64, x11 float64, x21 float64, y2 float64, x12 float64, x22 float64) {
	iv, err := _I.Get(257, "Renderer", "draw_trapezoid")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_part := gi.NewIntArgument(int(part))
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
func (v Renderer) GetAlpha(part RenderPartEnum) (result uint16) {
	iv, err := _I.Get(258, "Renderer", "get_alpha")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_part := gi.NewIntArgument(int(part))
	args := []gi.Argument{arg_v, arg_part}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint16()
	return
}

// pango_renderer_get_color
// container is not nil, container is Renderer
// is method
func (v Renderer) GetColor(part RenderPartEnum) (result Color) {
	iv, err := _I.Get(259, "Renderer", "get_color")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_part := gi.NewIntArgument(int(part))
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
func (v Renderer) PartChanged(part RenderPartEnum) {
	iv, err := _I.Get(263, "Renderer", "part_changed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_part := gi.NewIntArgument(int(part))
	args := []gi.Argument{arg_v, arg_part}
	iv.Call(args, nil, nil)
}

// pango_renderer_set_alpha
// container is not nil, container is Renderer
// is method
func (v Renderer) SetAlpha(part RenderPartEnum, alpha uint16) {
	iv, err := _I.Get(264, "Renderer", "set_alpha")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_part := gi.NewIntArgument(int(part))
	arg_alpha := gi.NewUint16Argument(alpha)
	args := []gi.Argument{arg_v, arg_part, arg_alpha}
	iv.Call(args, nil, nil)
}

// pango_renderer_set_color
// container is not nil, container is Renderer
// is method
func (v Renderer) SetColor(part RenderPartEnum, color Color) {
	iv, err := _I.Get(265, "Renderer", "set_color")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_part := gi.NewIntArgument(int(part))
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

func RendererPrivateGetType() gi.GType {
	ret := _I.GetGType(58, "RendererPrivate")
	return ret
}

// Enum Script
type ScriptEnum int

const (
	ScriptInvalidCode          ScriptEnum = -1
	ScriptCommon               ScriptEnum = 0
	ScriptInherited            ScriptEnum = 1
	ScriptArabic               ScriptEnum = 2
	ScriptArmenian             ScriptEnum = 3
	ScriptBengali              ScriptEnum = 4
	ScriptBopomofo             ScriptEnum = 5
	ScriptCherokee             ScriptEnum = 6
	ScriptCoptic               ScriptEnum = 7
	ScriptCyrillic             ScriptEnum = 8
	ScriptDeseret              ScriptEnum = 9
	ScriptDevanagari           ScriptEnum = 10
	ScriptEthiopic             ScriptEnum = 11
	ScriptGeorgian             ScriptEnum = 12
	ScriptGothic               ScriptEnum = 13
	ScriptGreek                ScriptEnum = 14
	ScriptGujarati             ScriptEnum = 15
	ScriptGurmukhi             ScriptEnum = 16
	ScriptHan                  ScriptEnum = 17
	ScriptHangul               ScriptEnum = 18
	ScriptHebrew               ScriptEnum = 19
	ScriptHiragana             ScriptEnum = 20
	ScriptKannada              ScriptEnum = 21
	ScriptKatakana             ScriptEnum = 22
	ScriptKhmer                ScriptEnum = 23
	ScriptLao                  ScriptEnum = 24
	ScriptLatin                ScriptEnum = 25
	ScriptMalayalam            ScriptEnum = 26
	ScriptMongolian            ScriptEnum = 27
	ScriptMyanmar              ScriptEnum = 28
	ScriptOgham                ScriptEnum = 29
	ScriptOldItalic            ScriptEnum = 30
	ScriptOriya                ScriptEnum = 31
	ScriptRunic                ScriptEnum = 32
	ScriptSinhala              ScriptEnum = 33
	ScriptSyriac               ScriptEnum = 34
	ScriptTamil                ScriptEnum = 35
	ScriptTelugu               ScriptEnum = 36
	ScriptThaana               ScriptEnum = 37
	ScriptThai                 ScriptEnum = 38
	ScriptTibetan              ScriptEnum = 39
	ScriptCanadianAboriginal   ScriptEnum = 40
	ScriptYi                   ScriptEnum = 41
	ScriptTagalog              ScriptEnum = 42
	ScriptHanunoo              ScriptEnum = 43
	ScriptBuhid                ScriptEnum = 44
	ScriptTagbanwa             ScriptEnum = 45
	ScriptBraille              ScriptEnum = 46
	ScriptCypriot              ScriptEnum = 47
	ScriptLimbu                ScriptEnum = 48
	ScriptOsmanya              ScriptEnum = 49
	ScriptShavian              ScriptEnum = 50
	ScriptLinearB              ScriptEnum = 51
	ScriptTaiLe                ScriptEnum = 52
	ScriptUgaritic             ScriptEnum = 53
	ScriptNewTaiLue            ScriptEnum = 54
	ScriptBuginese             ScriptEnum = 55
	ScriptGlagolitic           ScriptEnum = 56
	ScriptTifinagh             ScriptEnum = 57
	ScriptSylotiNagri          ScriptEnum = 58
	ScriptOldPersian           ScriptEnum = 59
	ScriptKharoshthi           ScriptEnum = 60
	ScriptUnknown              ScriptEnum = 61
	ScriptBalinese             ScriptEnum = 62
	ScriptCuneiform            ScriptEnum = 63
	ScriptPhoenician           ScriptEnum = 64
	ScriptPhagsPa              ScriptEnum = 65
	ScriptNko                  ScriptEnum = 66
	ScriptKayahLi              ScriptEnum = 67
	ScriptLepcha               ScriptEnum = 68
	ScriptRejang               ScriptEnum = 69
	ScriptSundanese            ScriptEnum = 70
	ScriptSaurashtra           ScriptEnum = 71
	ScriptCham                 ScriptEnum = 72
	ScriptOlChiki              ScriptEnum = 73
	ScriptVai                  ScriptEnum = 74
	ScriptCarian               ScriptEnum = 75
	ScriptLycian               ScriptEnum = 76
	ScriptLydian               ScriptEnum = 77
	ScriptBatak                ScriptEnum = 78
	ScriptBrahmi               ScriptEnum = 79
	ScriptMandaic              ScriptEnum = 80
	ScriptChakma               ScriptEnum = 81
	ScriptMeroiticCursive      ScriptEnum = 82
	ScriptMeroiticHieroglyphs  ScriptEnum = 83
	ScriptMiao                 ScriptEnum = 84
	ScriptSharada              ScriptEnum = 85
	ScriptSoraSompeng          ScriptEnum = 86
	ScriptTakri                ScriptEnum = 87
	ScriptBassaVah             ScriptEnum = 88
	ScriptCaucasianAlbanian    ScriptEnum = 89
	ScriptDuployan             ScriptEnum = 90
	ScriptElbasan              ScriptEnum = 91
	ScriptGrantha              ScriptEnum = 92
	ScriptKhojki               ScriptEnum = 93
	ScriptKhudawadi            ScriptEnum = 94
	ScriptLinearA              ScriptEnum = 95
	ScriptMahajani             ScriptEnum = 96
	ScriptManichaean           ScriptEnum = 97
	ScriptMendeKikakui         ScriptEnum = 98
	ScriptModi                 ScriptEnum = 99
	ScriptMro                  ScriptEnum = 100
	ScriptNabataean            ScriptEnum = 101
	ScriptOldNorthArabian      ScriptEnum = 102
	ScriptOldPermic            ScriptEnum = 103
	ScriptPahawhHmong          ScriptEnum = 104
	ScriptPalmyrene            ScriptEnum = 105
	ScriptPauCinHau            ScriptEnum = 106
	ScriptPsalterPahlavi       ScriptEnum = 107
	ScriptSiddham              ScriptEnum = 108
	ScriptTirhuta              ScriptEnum = 109
	ScriptWarangCiti           ScriptEnum = 110
	ScriptAhom                 ScriptEnum = 111
	ScriptAnatolianHieroglyphs ScriptEnum = 112
	ScriptHatran               ScriptEnum = 113
	ScriptMultani              ScriptEnum = 114
	ScriptOldHungarian         ScriptEnum = 115
	ScriptSignwriting          ScriptEnum = 116
)

func ScriptGetType() gi.GType {
	ret := _I.GetGType(59, "Script")
	return ret
}

// Struct ScriptIter
type ScriptIter struct {
	P unsafe.Pointer
}

func ScriptIterGetType() gi.GType {
	ret := _I.GetGType(60, "ScriptIter")
	return ret
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
func (v ScriptIter) GetRange() (start string, end string, script ScriptEnum) {
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
	script = ScriptEnum(outArgs[2].Int())
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

// Enum Stretch
type StretchEnum int

const (
	StretchUltraCondensed StretchEnum = 0
	StretchExtraCondensed StretchEnum = 1
	StretchCondensed      StretchEnum = 2
	StretchSemiCondensed  StretchEnum = 3
	StretchNormal         StretchEnum = 4
	StretchSemiExpanded   StretchEnum = 5
	StretchExpanded       StretchEnum = 6
	StretchExtraExpanded  StretchEnum = 7
	StretchUltraExpanded  StretchEnum = 8
)

func StretchGetType() gi.GType {
	ret := _I.GetGType(61, "Stretch")
	return ret
}

// Enum Style
type StyleEnum int

const (
	StyleNormal  StyleEnum = 0
	StyleOblique StyleEnum = 1
	StyleItalic  StyleEnum = 2
)

func StyleGetType() gi.GType {
	ret := _I.GetGType(62, "Style")
	return ret
}

// Enum TabAlign
type TabAlignEnum int

const (
	TabAlignLeft TabAlignEnum = 0
)

func TabAlignGetType() gi.GType {
	ret := _I.GetGType(63, "TabAlign")
	return ret
}

// Struct TabArray
type TabArray struct {
	P unsafe.Pointer
}

func TabArrayGetType() gi.GType {
	ret := _I.GetGType(64, "TabArray")
	return ret
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
func (v TabArray) GetTab(tab_index int32) (alignment TabAlignEnum, location int32) {
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
	alignment = TabAlignEnum(outArgs[0].Int())
	location = outArgs[1].Int32()
	return
}

// pango_tab_array_get_tabs
// container is not nil, container is TabArray
// is method
func (v TabArray) GetTabs() (alignments int /*TODO_TYPE tag: ifc, biType: enum*/, locations gi.Int32Array) {
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
	alignments = outArgs[0].Int() /*TODO tagIfc biType: enum*/
	locations.P = outArgs[1].Pointer()
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
func (v TabArray) SetTab(tab_index int32, alignment TabAlignEnum, location int32) {
	iv, err := _I.Get(278, "TabArray", "set_tab")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tab_index := gi.NewInt32Argument(tab_index)
	arg_alignment := gi.NewIntArgument(int(alignment))
	arg_location := gi.NewInt32Argument(location)
	args := []gi.Argument{arg_v, arg_tab_index, arg_alignment, arg_location}
	iv.Call(args, nil, nil)
}

// Enum Underline
type UnderlineEnum int

const (
	UnderlineNone   UnderlineEnum = 0
	UnderlineSingle UnderlineEnum = 1
	UnderlineDouble UnderlineEnum = 2
	UnderlineLow    UnderlineEnum = 3
	UnderlineError  UnderlineEnum = 4
)

func UnderlineGetType() gi.GType {
	ret := _I.GetGType(65, "Underline")
	return ret
}

// Enum Variant
type VariantEnum int

const (
	VariantNormal    VariantEnum = 0
	VariantSmallCaps VariantEnum = 1
)

func VariantGetType() gi.GType {
	ret := _I.GetGType(66, "Variant")
	return ret
}

// Enum Weight
type WeightEnum int

const (
	WeightThin       WeightEnum = 100
	WeightUltralight WeightEnum = 200
	WeightLight      WeightEnum = 300
	WeightSemilight  WeightEnum = 350
	WeightBook       WeightEnum = 380
	WeightNormal     WeightEnum = 400
	WeightMedium     WeightEnum = 500
	WeightSemibold   WeightEnum = 600
	WeightBold       WeightEnum = 700
	WeightUltrabold  WeightEnum = 800
	WeightHeavy      WeightEnum = 900
	WeightUltraheavy WeightEnum = 1000
)

func WeightGetType() gi.GType {
	ret := _I.GetGType(67, "Weight")
	return ret
}

// Enum WrapMode
type WrapModeEnum int

const (
	WrapModeWord     WrapModeEnum = 0
	WrapModeChar     WrapModeEnum = 1
	WrapModeWordChar WrapModeEnum = 2
)

func WrapModeGetType() gi.GType {
	ret := _I.GetGType(68, "WrapMode")
	return ret
}

// pango_attr_type_get_name
// container is nil
func AttrTypeGetName(type1 AttrTypeEnum) (result string) {
	iv, err := _I.Get(279, "attr_type_get_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// pango_attr_type_register
// container is nil
func AttrTypeRegister(name string) (result AttrTypeEnum) {
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
	gi.Free(c_name)
	result = AttrTypeEnum(ret.Int())
	return
}

// pango_bidi_type_for_unichar
// container is nil
func BidiTypeForUnichar(ch rune) (result BidiTypeEnum) {
	iv, err := _I.Get(281, "bidi_type_for_unichar", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewUint32Argument(uint32(ch))
	args := []gi.Argument{arg_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = BidiTypeEnum(ret.Int())
	return
}

// pango_break
// container is nil
// arg 3 attrs lenArgIdx 4
func Break(text string, length int32, analysis Analysis, attrs unsafe.Pointer, attrs_len int32) {
	iv, err := _I.Get(282, "break", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_text := gi.NewStringArgument(c_text)
	arg_length := gi.NewInt32Argument(length)
	arg_analysis := gi.NewPointerArgument(analysis.P)
	arg_attrs := gi.NewPointerArgument(attrs)
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
	gi.Free(c_key)
	result = ret.String().Take()
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
	gi.Free(c_key)
	result = ret.String().Take()
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
func FindBaseDir(text string, length int32) (result DirectionEnum) {
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
	gi.Free(c_text)
	result = DirectionEnum(ret.Int())
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
	gi.Free(c_str)
	result.P = ret.Pointer()
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
// arg 4 log_attrs lenArgIdx 5
func GetLogAttrs(text string, length int32, level int32, language Language, log_attrs unsafe.Pointer, attrs_len int32) {
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
	arg_log_attrs := gi.NewPointerArgument(log_attrs)
	arg_attrs_len := gi.NewInt32Argument(attrs_len)
	args := []gi.Argument{arg_text, arg_length, arg_level, arg_language, arg_log_attrs, arg_attrs_len}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// pango_get_mirror_char
// container is nil
func GetMirrorChar(ch rune, mirrored_ch rune) (result bool) {
	iv, err := _I.Get(292, "get_mirror_char", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewUint32Argument(uint32(ch))
	arg_mirrored_ch := gi.NewUint32Argument(uint32(mirrored_ch))
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
func GravityGetForMatrix(matrix Matrix) (result GravityEnum) {
	iv, err := _I.Get(294, "gravity_get_for_matrix", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_matrix := gi.NewPointerArgument(matrix.P)
	args := []gi.Argument{arg_matrix}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = GravityEnum(ret.Int())
	return
}

// pango_gravity_get_for_script
// container is nil
func GravityGetForScript(script ScriptEnum, base_gravity GravityEnum, hint GravityHintEnum) (result GravityEnum) {
	iv, err := _I.Get(295, "gravity_get_for_script", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_script := gi.NewIntArgument(int(script))
	arg_base_gravity := gi.NewIntArgument(int(base_gravity))
	arg_hint := gi.NewIntArgument(int(hint))
	args := []gi.Argument{arg_script, arg_base_gravity, arg_hint}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = GravityEnum(ret.Int())
	return
}

// pango_gravity_get_for_script_and_width
// container is nil
func GravityGetForScriptAndWidth(script ScriptEnum, wide bool, base_gravity GravityEnum, hint GravityHintEnum) (result GravityEnum) {
	iv, err := _I.Get(296, "gravity_get_for_script_and_width", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_script := gi.NewIntArgument(int(script))
	arg_wide := gi.NewBoolArgument(wide)
	arg_base_gravity := gi.NewIntArgument(int(base_gravity))
	arg_hint := gi.NewIntArgument(int(hint))
	args := []gi.Argument{arg_script, arg_wide, arg_base_gravity, arg_hint}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = GravityEnum(ret.Int())
	return
}

// pango_gravity_to_rotation
// container is nil
func GravityToRotation(gravity GravityEnum) (result float64) {
	iv, err := _I.Get(297, "gravity_to_rotation", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_gravity := gi.NewIntArgument(int(gravity))
	args := []gi.Argument{arg_gravity}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// pango_is_zero_width
// container is nil
func IsZeroWidth(ch rune) (result bool) {
	iv, err := _I.Get(298, "is_zero_width", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewUint32Argument(uint32(ch))
	args := []gi.Argument{arg_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// pango_itemize
// container is nil
func Itemize(context IContext, text string, start_index int32, length int32, attrs AttrList, cached_iter AttrIterator) (result glib.List) {
	iv, err := _I.Get(299, "itemize", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_Context()
	}
	c_text := gi.CString(text)
	arg_context := gi.NewPointerArgument(tmp)
	arg_text := gi.NewStringArgument(c_text)
	arg_start_index := gi.NewInt32Argument(start_index)
	arg_length := gi.NewInt32Argument(length)
	arg_attrs := gi.NewPointerArgument(attrs.P)
	arg_cached_iter := gi.NewPointerArgument(cached_iter.P)
	args := []gi.Argument{arg_context, arg_text, arg_start_index, arg_length, arg_attrs, arg_cached_iter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_text)
	result.P = ret.Pointer()
	return
}

// pango_itemize_with_base_dir
// container is nil
func ItemizeWithBaseDir(context IContext, base_dir DirectionEnum, text string, start_index int32, length int32, attrs AttrList, cached_iter AttrIterator) (result glib.List) {
	iv, err := _I.Get(300, "itemize_with_base_dir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_Context()
	}
	c_text := gi.CString(text)
	arg_context := gi.NewPointerArgument(tmp)
	arg_base_dir := gi.NewIntArgument(int(base_dir))
	arg_text := gi.NewStringArgument(c_text)
	arg_start_index := gi.NewInt32Argument(start_index)
	arg_length := gi.NewInt32Argument(length)
	arg_attrs := gi.NewPointerArgument(attrs.P)
	arg_cached_iter := gi.NewPointerArgument(cached_iter.P)
	args := []gi.Argument{arg_context, arg_base_dir, arg_text, arg_start_index, arg_length, arg_attrs, arg_cached_iter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_text)
	result.P = ret.Pointer()
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
	gi.Free(c_language)
	result.P = ret.Pointer()
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

// pango_lookup_aliases
// container is nil
// arg 1 families lenArgIdx 2
func LookupAliases(fontname string) (families gi.CStrArray) {
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
	var n_families int32
	_ = n_families
	families.P = outArgs[0].Pointer()
	n_families = outArgs[1].Int32()
	return
}

// pango_markup_parser_finish
// container is nil
func MarkupParserFinish(context glib.MarkupParseContext) (result bool, attr_list AttrList, text string, accel_char rune, err error) {
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
	err = gi.ToError(outArgs[3].Pointer())
	attr_list.P = outArgs[0].Pointer()
	text = outArgs[1].String().Take()
	accel_char = rune(outArgs[2].Uint32())
	result = ret.Bool()
	return
}

// pango_markup_parser_new
// container is nil
func MarkupParserNew(accel_marker rune) (result glib.MarkupParseContext) {
	iv, err := _I.Get(305, "markup_parser_new", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_accel_marker := gi.NewUint32Argument(uint32(accel_marker))
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
func ParseEnum(type1 gi.GType, str string, warn bool) (result bool, value int32, possible_values string) {
	iv, err := _I.Get(307, "parse_enum", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	c_str := gi.CString(str)
	arg_type1 := gi.NewUintArgument(uint(type1))
	arg_str := gi.NewStringArgument(c_str)
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_warn := gi.NewBoolArgument(warn)
	arg_possible_values := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_type1, arg_str, arg_value, arg_warn, arg_possible_values}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_str)
	value = outArgs[0].Int32()
	possible_values = outArgs[1].String().Take()
	result = ret.Bool()
	return
}

// pango_parse_markup
// container is nil
func ParseMarkup(markup_text string, length int32, accel_marker rune) (result bool, attr_list AttrList, text string, accel_char rune, err error) {
	iv, err := _I.Get(308, "parse_markup", "")
	if err != nil {
		return
	}
	var outArgs [4]gi.Argument
	c_markup_text := gi.CString(markup_text)
	arg_markup_text := gi.NewStringArgument(c_markup_text)
	arg_length := gi.NewInt32Argument(length)
	arg_accel_marker := gi.NewUint32Argument(uint32(accel_marker))
	arg_attr_list := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_text := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_accel_char := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_markup_text, arg_length, arg_accel_marker, arg_attr_list, arg_text, arg_accel_char, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_markup_text)
	err = gi.ToError(outArgs[3].Pointer())
	attr_list.P = outArgs[0].Pointer()
	text = outArgs[1].String().Take()
	accel_char = rune(outArgs[2].Uint32())
	result = ret.Bool()
	return
}

// pango_parse_stretch
// container is nil
func ParseStretch(str string, warn bool) (result bool, stretch StretchEnum) {
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
	gi.Free(c_str)
	stretch = StretchEnum(outArgs[0].Int())
	result = ret.Bool()
	return
}

// pango_parse_style
// container is nil
func ParseStyle(str string, warn bool) (result bool, style StyleEnum) {
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
	gi.Free(c_str)
	style = StyleEnum(outArgs[0].Int())
	result = ret.Bool()
	return
}

// pango_parse_variant
// container is nil
func ParseVariant(str string, warn bool) (result bool, variant VariantEnum) {
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
	gi.Free(c_str)
	variant = VariantEnum(outArgs[0].Int())
	result = ret.Bool()
	return
}

// pango_parse_weight
// container is nil
func ParseWeight(str string, warn bool) (result bool, weight WeightEnum) {
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
	gi.Free(c_str)
	weight = WeightEnum(outArgs[0].Int())
	result = ret.Bool()
	return
}

// pango_quantize_line_geometry
// container is nil
func QuantizeLineGeometry(thickness int /*TODO:TYPE*/, position int /*TODO:TYPE*/) {
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
func ReadLine(stream unsafe.Pointer, str glib.String) (result int32) {
	iv, err := _I.Get(314, "read_line", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_stream := gi.NewPointerArgument(stream)
	arg_str := gi.NewPointerArgument(str.P)
	args := []gi.Argument{arg_stream, arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// pango_reorder_items
// container is nil
func ReorderItems(logical_items glib.List) (result glib.List) {
	iv, err := _I.Get(315, "reorder_items", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_logical_items := gi.NewPointerArgument(logical_items.P)
	args := []gi.Argument{arg_logical_items}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// pango_scan_int
// container is nil
func ScanInt(pos int /*TODO:TYPE*/) (result bool, out int32) {
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
	out = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// pango_scan_string
// container is nil
func ScanString(pos int /*TODO:TYPE*/, out glib.String) (result bool) {
	iv, err := _I.Get(317, "scan_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_out := gi.NewPointerArgument(out.P)
	args := []gi.Argument{arg_out}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	return
}

// pango_scan_word
// container is nil
func ScanWord(pos int /*TODO:TYPE*/, out glib.String) (result bool) {
	iv, err := _I.Get(318, "scan_word", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_out := gi.NewPointerArgument(out.P)
	args := []gi.Argument{arg_out}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	result = ret.Bool()
	return
}

// pango_script_for_unichar
// container is nil
func ScriptForUnichar(ch rune) (result ScriptEnum) {
	iv, err := _I.Get(319, "script_for_unichar", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewUint32Argument(uint32(ch))
	args := []gi.Argument{arg_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ScriptEnum(ret.Int())
	return
}

// pango_script_get_sample_language
// container is nil
func ScriptGetSampleLanguage(script ScriptEnum) (result Language) {
	iv, err := _I.Get(320, "script_get_sample_language", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_script := gi.NewIntArgument(int(script))
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
func SkipSpace(pos int /*TODO:TYPE*/) (result bool) {
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
func SplitFileList(str string) (result gi.CStrArray) {
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
	gi.Free(c_str)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
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
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// pango_unichar_direction
// container is nil
func UnicharDirection(ch rune) (result DirectionEnum) {
	iv, err := _I.Get(326, "unichar_direction", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewUint32Argument(uint32(ch))
	args := []gi.Argument{arg_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = DirectionEnum(ret.Int())
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
