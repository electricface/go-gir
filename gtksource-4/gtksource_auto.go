package gtksource

/*
#cgo pkg-config: gtksourceview-4
#include <gtksourceview/gtksource.h>
*/
import "C"
import "github.com/electricface/go-gir/atk-1.0"
import "github.com/electricface/go-gir/cairo-1.0"
import "github.com/electricface/go-gir/gdk-3.0"
import "github.com/electricface/go-gir/gdkpixbuf-2.0"
import "github.com/electricface/go-gir/gio-2.0"
import "github.com/electricface/go-gir/glib-2.0"
import "github.com/electricface/go-gir/gobject-2.0"
import "github.com/electricface/go-gir/gtk-3.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("GtkSource")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GtkSource", "4", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Enum BackgroundPatternType
type BackgroundPatternTypeEnum int

const (
	BackgroundPatternTypeNone BackgroundPatternTypeEnum = 0
	BackgroundPatternTypeGrid BackgroundPatternTypeEnum = 1
)

func BackgroundPatternTypeGetType() gi.GType {
	ret := _I.GetGType(0, "BackgroundPatternType")
	return ret
}

// Enum BracketMatchType
type BracketMatchTypeEnum int

const (
	BracketMatchTypeNone       BracketMatchTypeEnum = 0
	BracketMatchTypeOutOfRange BracketMatchTypeEnum = 1
	BracketMatchTypeNotFound   BracketMatchTypeEnum = 2
	BracketMatchTypeFound      BracketMatchTypeEnum = 3
)

func BracketMatchTypeGetType() gi.GType {
	ret := _I.GetGType(1, "BracketMatchType")
	return ret
}

// Object Buffer
type Buffer struct {
	gtk.TextBuffer
}

func WrapBuffer(p unsafe.Pointer) (r Buffer) { r.P = p; return }

type IBuffer interface{ P_Buffer() unsafe.Pointer }

func (v Buffer) P_Buffer() unsafe.Pointer { return v.P }
func BufferGetType() gi.GType {
	ret := _I.GetGType(2, "Buffer")
	return ret
}

// gtk_source_buffer_new
// container is not nil, container is Buffer
// is constructor
func NewBuffer(table gtk.ITextTagTable) (result Buffer) {
	iv, err := _I.Get(0, "Buffer", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if table != nil {
		tmp = table.P_TextTagTable()
	}
	arg_table := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_table}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_buffer_new_with_language
// container is not nil, container is Buffer
// is constructor
func NewBufferWithLanguage(language ILanguage) (result Buffer) {
	iv, err := _I.Get(1, "Buffer", "new_with_language")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if language != nil {
		tmp = language.P_Language()
	}
	arg_language := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_language}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_buffer_backward_iter_to_source_mark
// container is not nil, container is Buffer
// is method
func (v Buffer) BackwardIterToSourceMark(iter gtk.TextIter, category string) (result bool) {
	iv, err := _I.Get(2, "Buffer", "backward_iter_to_source_mark")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_category := gi.CString(category)
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_category := gi.NewStringArgument(c_category)
	args := []gi.Argument{arg_v, arg_iter, arg_category}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_category)
	result = ret.Bool()
	return
}

// gtk_source_buffer_begin_not_undoable_action
// container is not nil, container is Buffer
// is method
func (v Buffer) BeginNotUndoableAction() {
	iv, err := _I.Get(3, "Buffer", "begin_not_undoable_action")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_buffer_can_redo
// container is not nil, container is Buffer
// is method
func (v Buffer) CanRedo() (result bool) {
	iv, err := _I.Get(4, "Buffer", "can_redo")
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

// gtk_source_buffer_can_undo
// container is not nil, container is Buffer
// is method
func (v Buffer) CanUndo() (result bool) {
	iv, err := _I.Get(5, "Buffer", "can_undo")
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

// gtk_source_buffer_change_case
// container is not nil, container is Buffer
// is method
func (v Buffer) ChangeCase(case_type ChangeCaseTypeEnum, start gtk.TextIter, end gtk.TextIter) {
	iv, err := _I.Get(6, "Buffer", "change_case")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_case_type := gi.NewIntArgument(int(case_type))
	arg_start := gi.NewPointerArgument(start.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_v, arg_case_type, arg_start, arg_end}
	iv.Call(args, nil, nil)
}

// gtk_source_buffer_create_source_mark
// container is not nil, container is Buffer
// is method
func (v Buffer) CreateSourceMark(name string, category string, where gtk.TextIter) (result Mark) {
	iv, err := _I.Get(7, "Buffer", "create_source_mark")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_category := gi.CString(category)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	arg_category := gi.NewStringArgument(c_category)
	arg_where := gi.NewPointerArgument(where.P)
	args := []gi.Argument{arg_v, arg_name, arg_category, arg_where}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_category)
	result.P = ret.Pointer()
	return
}

// gtk_source_buffer_end_not_undoable_action
// container is not nil, container is Buffer
// is method
func (v Buffer) EndNotUndoableAction() {
	iv, err := _I.Get(8, "Buffer", "end_not_undoable_action")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_buffer_ensure_highlight
// container is not nil, container is Buffer
// is method
func (v Buffer) EnsureHighlight(start gtk.TextIter, end gtk.TextIter) {
	iv, err := _I.Get(9, "Buffer", "ensure_highlight")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewPointerArgument(start.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_v, arg_start, arg_end}
	iv.Call(args, nil, nil)
}

// gtk_source_buffer_forward_iter_to_source_mark
// container is not nil, container is Buffer
// is method
func (v Buffer) ForwardIterToSourceMark(iter gtk.TextIter, category string) (result bool) {
	iv, err := _I.Get(10, "Buffer", "forward_iter_to_source_mark")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_category := gi.CString(category)
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_category := gi.NewStringArgument(c_category)
	args := []gi.Argument{arg_v, arg_iter, arg_category}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_category)
	result = ret.Bool()
	return
}

// gtk_source_buffer_get_context_classes_at_iter
// container is not nil, container is Buffer
// is method
func (v Buffer) GetContextClassesAtIter(iter gtk.TextIter) (result gi.CStrArray) {
	iv, err := _I.Get(11, "Buffer", "get_context_classes_at_iter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	args := []gi.Argument{arg_v, arg_iter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// gtk_source_buffer_get_highlight_matching_brackets
// container is not nil, container is Buffer
// is method
func (v Buffer) GetHighlightMatchingBrackets() (result bool) {
	iv, err := _I.Get(12, "Buffer", "get_highlight_matching_brackets")
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

// gtk_source_buffer_get_highlight_syntax
// container is not nil, container is Buffer
// is method
func (v Buffer) GetHighlightSyntax() (result bool) {
	iv, err := _I.Get(13, "Buffer", "get_highlight_syntax")
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

// gtk_source_buffer_get_implicit_trailing_newline
// container is not nil, container is Buffer
// is method
func (v Buffer) GetImplicitTrailingNewline() (result bool) {
	iv, err := _I.Get(14, "Buffer", "get_implicit_trailing_newline")
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

// gtk_source_buffer_get_language
// container is not nil, container is Buffer
// is method
func (v Buffer) GetLanguage() (result Language) {
	iv, err := _I.Get(15, "Buffer", "get_language")
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

// gtk_source_buffer_get_max_undo_levels
// container is not nil, container is Buffer
// is method
func (v Buffer) GetMaxUndoLevels() (result int32) {
	iv, err := _I.Get(16, "Buffer", "get_max_undo_levels")
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

// gtk_source_buffer_get_source_marks_at_iter
// container is not nil, container is Buffer
// is method
func (v Buffer) GetSourceMarksAtIter(iter gtk.TextIter, category string) (result glib.SList) {
	iv, err := _I.Get(17, "Buffer", "get_source_marks_at_iter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_category := gi.CString(category)
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_category := gi.NewStringArgument(c_category)
	args := []gi.Argument{arg_v, arg_iter, arg_category}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_category)
	result.P = ret.Pointer()
	return
}

// gtk_source_buffer_get_source_marks_at_line
// container is not nil, container is Buffer
// is method
func (v Buffer) GetSourceMarksAtLine(line int32, category string) (result glib.SList) {
	iv, err := _I.Get(18, "Buffer", "get_source_marks_at_line")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_category := gi.CString(category)
	arg_v := gi.NewPointerArgument(v.P)
	arg_line := gi.NewInt32Argument(line)
	arg_category := gi.NewStringArgument(c_category)
	args := []gi.Argument{arg_v, arg_line, arg_category}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_category)
	result.P = ret.Pointer()
	return
}

// gtk_source_buffer_get_style_scheme
// container is not nil, container is Buffer
// is method
func (v Buffer) GetStyleScheme() (result StyleScheme) {
	iv, err := _I.Get(19, "Buffer", "get_style_scheme")
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

// gtk_source_buffer_get_undo_manager
// container is not nil, container is Buffer
// is method
func (v Buffer) GetUndoManager() (result UndoManager) {
	iv, err := _I.Get(20, "Buffer", "get_undo_manager")
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

// gtk_source_buffer_iter_backward_to_context_class_toggle
// container is not nil, container is Buffer
// is method
func (v Buffer) IterBackwardToContextClassToggle(iter gtk.TextIter, context_class string) (result bool) {
	iv, err := _I.Get(21, "Buffer", "iter_backward_to_context_class_toggle")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_context_class := gi.CString(context_class)
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_context_class := gi.NewStringArgument(c_context_class)
	args := []gi.Argument{arg_v, arg_iter, arg_context_class}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_context_class)
	result = ret.Bool()
	return
}

// gtk_source_buffer_iter_forward_to_context_class_toggle
// container is not nil, container is Buffer
// is method
func (v Buffer) IterForwardToContextClassToggle(iter gtk.TextIter, context_class string) (result bool) {
	iv, err := _I.Get(22, "Buffer", "iter_forward_to_context_class_toggle")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_context_class := gi.CString(context_class)
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_context_class := gi.NewStringArgument(c_context_class)
	args := []gi.Argument{arg_v, arg_iter, arg_context_class}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_context_class)
	result = ret.Bool()
	return
}

// gtk_source_buffer_iter_has_context_class
// container is not nil, container is Buffer
// is method
func (v Buffer) IterHasContextClass(iter gtk.TextIter, context_class string) (result bool) {
	iv, err := _I.Get(23, "Buffer", "iter_has_context_class")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_context_class := gi.CString(context_class)
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_context_class := gi.NewStringArgument(c_context_class)
	args := []gi.Argument{arg_v, arg_iter, arg_context_class}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_context_class)
	result = ret.Bool()
	return
}

// gtk_source_buffer_join_lines
// container is not nil, container is Buffer
// is method
func (v Buffer) JoinLines(start gtk.TextIter, end gtk.TextIter) {
	iv, err := _I.Get(24, "Buffer", "join_lines")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewPointerArgument(start.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_v, arg_start, arg_end}
	iv.Call(args, nil, nil)
}

// gtk_source_buffer_redo
// container is not nil, container is Buffer
// is method
func (v Buffer) Redo() {
	iv, err := _I.Get(25, "Buffer", "redo")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_buffer_remove_source_marks
// container is not nil, container is Buffer
// is method
func (v Buffer) RemoveSourceMarks(start gtk.TextIter, end gtk.TextIter, category string) {
	iv, err := _I.Get(26, "Buffer", "remove_source_marks")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_category := gi.CString(category)
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewPointerArgument(start.P)
	arg_end := gi.NewPointerArgument(end.P)
	arg_category := gi.NewStringArgument(c_category)
	args := []gi.Argument{arg_v, arg_start, arg_end, arg_category}
	iv.Call(args, nil, nil)
	gi.Free(c_category)
}

// gtk_source_buffer_set_highlight_matching_brackets
// container is not nil, container is Buffer
// is method
func (v Buffer) SetHighlightMatchingBrackets(highlight bool) {
	iv, err := _I.Get(27, "Buffer", "set_highlight_matching_brackets")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_highlight := gi.NewBoolArgument(highlight)
	args := []gi.Argument{arg_v, arg_highlight}
	iv.Call(args, nil, nil)
}

// gtk_source_buffer_set_highlight_syntax
// container is not nil, container is Buffer
// is method
func (v Buffer) SetHighlightSyntax(highlight bool) {
	iv, err := _I.Get(28, "Buffer", "set_highlight_syntax")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_highlight := gi.NewBoolArgument(highlight)
	args := []gi.Argument{arg_v, arg_highlight}
	iv.Call(args, nil, nil)
}

// gtk_source_buffer_set_implicit_trailing_newline
// container is not nil, container is Buffer
// is method
func (v Buffer) SetImplicitTrailingNewline(implicit_trailing_newline bool) {
	iv, err := _I.Get(29, "Buffer", "set_implicit_trailing_newline")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_implicit_trailing_newline := gi.NewBoolArgument(implicit_trailing_newline)
	args := []gi.Argument{arg_v, arg_implicit_trailing_newline}
	iv.Call(args, nil, nil)
}

// gtk_source_buffer_set_language
// container is not nil, container is Buffer
// is method
func (v Buffer) SetLanguage(language ILanguage) {
	iv, err := _I.Get(30, "Buffer", "set_language")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if language != nil {
		tmp = language.P_Language()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_language := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_language}
	iv.Call(args, nil, nil)
}

// gtk_source_buffer_set_max_undo_levels
// container is not nil, container is Buffer
// is method
func (v Buffer) SetMaxUndoLevels(max_undo_levels int32) {
	iv, err := _I.Get(31, "Buffer", "set_max_undo_levels")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_max_undo_levels := gi.NewInt32Argument(max_undo_levels)
	args := []gi.Argument{arg_v, arg_max_undo_levels}
	iv.Call(args, nil, nil)
}

// gtk_source_buffer_set_style_scheme
// container is not nil, container is Buffer
// is method
func (v Buffer) SetStyleScheme(scheme IStyleScheme) {
	iv, err := _I.Get(32, "Buffer", "set_style_scheme")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if scheme != nil {
		tmp = scheme.P_StyleScheme()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_scheme := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_scheme}
	iv.Call(args, nil, nil)
}

// gtk_source_buffer_set_undo_manager
// container is not nil, container is Buffer
// is method
func (v Buffer) SetUndoManager(manager IUndoManager) {
	iv, err := _I.Get(33, "Buffer", "set_undo_manager")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if manager != nil {
		tmp = manager.P_UndoManager()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_manager := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_manager}
	iv.Call(args, nil, nil)
}

// gtk_source_buffer_sort_lines
// container is not nil, container is Buffer
// is method
func (v Buffer) SortLines(start gtk.TextIter, end gtk.TextIter, flags SortFlags, column int32) {
	iv, err := _I.Get(34, "Buffer", "sort_lines")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewPointerArgument(start.P)
	arg_end := gi.NewPointerArgument(end.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_column := gi.NewInt32Argument(column)
	args := []gi.Argument{arg_v, arg_start, arg_end, arg_flags, arg_column}
	iv.Call(args, nil, nil)
}

// gtk_source_buffer_undo
// container is not nil, container is Buffer
// is method
func (v Buffer) Undo() {
	iv, err := _I.Get(35, "Buffer", "undo")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// ignore GType struct BufferClass
// Struct BufferPrivate
type BufferPrivate struct {
	P unsafe.Pointer
}

func BufferPrivateGetType() gi.GType {
	ret := _I.GetGType(3, "BufferPrivate")
	return ret
}

// Enum ChangeCaseType
type ChangeCaseTypeEnum int

const (
	ChangeCaseTypeLower  ChangeCaseTypeEnum = 0
	ChangeCaseTypeUpper  ChangeCaseTypeEnum = 1
	ChangeCaseTypeToggle ChangeCaseTypeEnum = 2
	ChangeCaseTypeTitle  ChangeCaseTypeEnum = 3
)

func ChangeCaseTypeGetType() gi.GType {
	ret := _I.GetGType(4, "ChangeCaseType")
	return ret
}

// Object Completion
type Completion struct {
	gtk.BuildableIfc
	gobject.Object
}

func WrapCompletion(p unsafe.Pointer) (r Completion) { r.P = p; return }

type ICompletion interface{ P_Completion() unsafe.Pointer }

func (v Completion) P_Completion() unsafe.Pointer { return v.P }
func (v Completion) P_Buildable() unsafe.Pointer  { return v.P }
func CompletionGetType() gi.GType {
	ret := _I.GetGType(5, "Completion")
	return ret
}

// gtk_source_completion_add_provider
// container is not nil, container is Completion
// is method
func (v Completion) AddProvider(provider ICompletionProvider) (result bool, err error) {
	iv, err := _I.Get(36, "Completion", "add_provider")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if provider != nil {
		tmp = provider.P_CompletionProvider()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_provider := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_provider, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// gtk_source_completion_block_interactive
// container is not nil, container is Completion
// is method
func (v Completion) BlockInteractive() {
	iv, err := _I.Get(37, "Completion", "block_interactive")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_completion_create_context
// container is not nil, container is Completion
// is method
func (v Completion) CreateContext(position gtk.TextIter) (result CompletionContext) {
	iv, err := _I.Get(38, "Completion", "create_context")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_position := gi.NewPointerArgument(position.P)
	args := []gi.Argument{arg_v, arg_position}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_completion_get_info_window
// container is not nil, container is Completion
// is method
func (v Completion) GetInfoWindow() (result CompletionInfo) {
	iv, err := _I.Get(39, "Completion", "get_info_window")
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

// gtk_source_completion_get_providers
// container is not nil, container is Completion
// is method
func (v Completion) GetProviders() (result glib.List) {
	iv, err := _I.Get(40, "Completion", "get_providers")
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

// gtk_source_completion_get_view
// container is not nil, container is Completion
// is method
func (v Completion) GetView() (result View) {
	iv, err := _I.Get(41, "Completion", "get_view")
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

// gtk_source_completion_hide
// container is not nil, container is Completion
// is method
func (v Completion) Hide() {
	iv, err := _I.Get(42, "Completion", "hide")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_completion_remove_provider
// container is not nil, container is Completion
// is method
func (v Completion) RemoveProvider(provider ICompletionProvider) (result bool, err error) {
	iv, err := _I.Get(43, "Completion", "remove_provider")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if provider != nil {
		tmp = provider.P_CompletionProvider()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_provider := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_provider, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// gtk_source_completion_start
// container is not nil, container is Completion
// is method
func (v Completion) Start(providers glib.List, context ICompletionContext) (result bool) {
	iv, err := _I.Get(44, "Completion", "start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_CompletionContext()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_providers := gi.NewPointerArgument(providers.P)
	arg_context := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_providers, arg_context}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gtk_source_completion_unblock_interactive
// container is not nil, container is Completion
// is method
func (v Completion) UnblockInteractive() {
	iv, err := _I.Get(45, "Completion", "unblock_interactive")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Flags CompletionActivation
type CompletionActivationFlags int

const (
	CompletionActivationNone          CompletionActivationFlags = 0
	CompletionActivationInteractive   CompletionActivationFlags = 1
	CompletionActivationUserRequested CompletionActivationFlags = 2
)

func CompletionActivationGetType() gi.GType {
	ret := _I.GetGType(6, "CompletionActivation")
	return ret
}

// ignore GType struct CompletionClass
// Object CompletionContext
type CompletionContext struct {
	gobject.InitiallyUnowned
}

func WrapCompletionContext(p unsafe.Pointer) (r CompletionContext) { r.P = p; return }

type ICompletionContext interface{ P_CompletionContext() unsafe.Pointer }

func (v CompletionContext) P_CompletionContext() unsafe.Pointer { return v.P }
func CompletionContextGetType() gi.GType {
	ret := _I.GetGType(7, "CompletionContext")
	return ret
}

// gtk_source_completion_context_add_proposals
// container is not nil, container is CompletionContext
// is method
func (v CompletionContext) AddProposals(provider ICompletionProvider, proposals glib.List, finished bool) {
	iv, err := _I.Get(46, "CompletionContext", "add_proposals")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if provider != nil {
		tmp = provider.P_CompletionProvider()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_provider := gi.NewPointerArgument(tmp)
	arg_proposals := gi.NewPointerArgument(proposals.P)
	arg_finished := gi.NewBoolArgument(finished)
	args := []gi.Argument{arg_v, arg_provider, arg_proposals, arg_finished}
	iv.Call(args, nil, nil)
}

// gtk_source_completion_context_get_activation
// container is not nil, container is CompletionContext
// is method
func (v CompletionContext) GetActivation() (result CompletionActivationFlags) {
	iv, err := _I.Get(47, "CompletionContext", "get_activation")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = CompletionActivationFlags(ret.Int())
	return
}

// gtk_source_completion_context_get_iter
// container is not nil, container is CompletionContext
// is method
func (v CompletionContext) GetIter(iter gtk.TextIter) (result bool) {
	iv, err := _I.Get(48, "CompletionContext", "get_iter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	args := []gi.Argument{arg_v, arg_iter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// ignore GType struct CompletionContextClass
// Struct CompletionContextPrivate
type CompletionContextPrivate struct {
	P unsafe.Pointer
}

func CompletionContextPrivateGetType() gi.GType {
	ret := _I.GetGType(8, "CompletionContextPrivate")
	return ret
}

// Enum CompletionError
type CompletionErrorEnum int

const (
	CompletionErrorAlreadyBound CompletionErrorEnum = 0
	CompletionErrorNotBound     CompletionErrorEnum = 1
)

func CompletionErrorGetType() gi.GType {
	ret := _I.GetGType(9, "CompletionError")
	return ret
}

// Object CompletionInfo
type CompletionInfo struct {
	atk.ImplementorIfaceIfc
	gtk.BuildableIfc
	gtk.Window
}

func WrapCompletionInfo(p unsafe.Pointer) (r CompletionInfo) { r.P = p; return }

type ICompletionInfo interface{ P_CompletionInfo() unsafe.Pointer }

func (v CompletionInfo) P_CompletionInfo() unsafe.Pointer   { return v.P }
func (v CompletionInfo) P_ImplementorIface() unsafe.Pointer { return v.P }
func (v CompletionInfo) P_Buildable() unsafe.Pointer        { return v.P }
func CompletionInfoGetType() gi.GType {
	ret := _I.GetGType(10, "CompletionInfo")
	return ret
}

// gtk_source_completion_info_new
// container is not nil, container is CompletionInfo
// is constructor
func NewCompletionInfo() (result CompletionInfo) {
	iv, err := _I.Get(49, "CompletionInfo", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_completion_info_move_to_iter
// container is not nil, container is CompletionInfo
// is method
func (v CompletionInfo) MoveToIter(view gtk.ITextView, iter gtk.TextIter) {
	iv, err := _I.Get(50, "CompletionInfo", "move_to_iter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if view != nil {
		tmp = view.P_TextView()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_view := gi.NewPointerArgument(tmp)
	arg_iter := gi.NewPointerArgument(iter.P)
	args := []gi.Argument{arg_v, arg_view, arg_iter}
	iv.Call(args, nil, nil)
}

// ignore GType struct CompletionInfoClass
// Struct CompletionInfoPrivate
type CompletionInfoPrivate struct {
	P unsafe.Pointer
}

func CompletionInfoPrivateGetType() gi.GType {
	ret := _I.GetGType(11, "CompletionInfoPrivate")
	return ret
}

// Object CompletionItem
type CompletionItem struct {
	CompletionProposalIfc
	gobject.Object
}

func WrapCompletionItem(p unsafe.Pointer) (r CompletionItem) { r.P = p; return }

type ICompletionItem interface{ P_CompletionItem() unsafe.Pointer }

func (v CompletionItem) P_CompletionItem() unsafe.Pointer     { return v.P }
func (v CompletionItem) P_CompletionProposal() unsafe.Pointer { return v.P }
func CompletionItemGetType() gi.GType {
	ret := _I.GetGType(12, "CompletionItem")
	return ret
}

// gtk_source_completion_item_new
// container is not nil, container is CompletionItem
// is constructor
func NewCompletionItem() (result CompletionItem) {
	iv, err := _I.Get(51, "CompletionItem", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_completion_item_set_gicon
// container is not nil, container is CompletionItem
// is method
func (v CompletionItem) SetGicon(gicon gio.IIcon) {
	iv, err := _I.Get(52, "CompletionItem", "set_gicon")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if gicon != nil {
		tmp = gicon.P_Icon()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_gicon := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_gicon}
	iv.Call(args, nil, nil)
}

// gtk_source_completion_item_set_icon
// container is not nil, container is CompletionItem
// is method
func (v CompletionItem) SetIcon(icon gdkpixbuf.IPixbuf) {
	iv, err := _I.Get(53, "CompletionItem", "set_icon")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if icon != nil {
		tmp = icon.P_Pixbuf()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_icon := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_icon}
	iv.Call(args, nil, nil)
}

// gtk_source_completion_item_set_icon_name
// container is not nil, container is CompletionItem
// is method
func (v CompletionItem) SetIconName(icon_name string) {
	iv, err := _I.Get(54, "CompletionItem", "set_icon_name")
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

// gtk_source_completion_item_set_info
// container is not nil, container is CompletionItem
// is method
func (v CompletionItem) SetInfo(info string) {
	iv, err := _I.Get(55, "CompletionItem", "set_info")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_info := gi.CString(info)
	arg_v := gi.NewPointerArgument(v.P)
	arg_info := gi.NewStringArgument(c_info)
	args := []gi.Argument{arg_v, arg_info}
	iv.Call(args, nil, nil)
	gi.Free(c_info)
}

// gtk_source_completion_item_set_label
// container is not nil, container is CompletionItem
// is method
func (v CompletionItem) SetLabel(label string) {
	iv, err := _I.Get(56, "CompletionItem", "set_label")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_label := gi.CString(label)
	arg_v := gi.NewPointerArgument(v.P)
	arg_label := gi.NewStringArgument(c_label)
	args := []gi.Argument{arg_v, arg_label}
	iv.Call(args, nil, nil)
	gi.Free(c_label)
}

// gtk_source_completion_item_set_markup
// container is not nil, container is CompletionItem
// is method
func (v CompletionItem) SetMarkup(markup string) {
	iv, err := _I.Get(57, "CompletionItem", "set_markup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_markup := gi.CString(markup)
	arg_v := gi.NewPointerArgument(v.P)
	arg_markup := gi.NewStringArgument(c_markup)
	args := []gi.Argument{arg_v, arg_markup}
	iv.Call(args, nil, nil)
	gi.Free(c_markup)
}

// gtk_source_completion_item_set_text
// container is not nil, container is CompletionItem
// is method
func (v CompletionItem) SetText(text string) {
	iv, err := _I.Get(58, "CompletionItem", "set_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	args := []gi.Argument{arg_v, arg_text}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// ignore GType struct CompletionItemClass
// Struct CompletionItemPrivate
type CompletionItemPrivate struct {
	P unsafe.Pointer
}

func CompletionItemPrivateGetType() gi.GType {
	ret := _I.GetGType(13, "CompletionItemPrivate")
	return ret
}

// Struct CompletionPrivate
type CompletionPrivate struct {
	P unsafe.Pointer
}

func CompletionPrivateGetType() gi.GType {
	ret := _I.GetGType(14, "CompletionPrivate")
	return ret
}

// Interface CompletionProposal
type CompletionProposal struct {
	CompletionProposalIfc
	P unsafe.Pointer
}
type CompletionProposalIfc struct{}
type ICompletionProposal interface{ P_CompletionProposal() unsafe.Pointer }

func (v CompletionProposal) P_CompletionProposal() unsafe.Pointer { return v.P }
func CompletionProposalGetType() gi.GType {
	ret := _I.GetGType(15, "CompletionProposal")
	return ret
}

// gtk_source_completion_proposal_changed
// container is not nil, container is CompletionProposal
// is method
func (v *CompletionProposalIfc) Changed() {
	iv, err := _I.Get(59, "CompletionProposal", "changed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_completion_proposal_equal
// container is not nil, container is CompletionProposal
// is method
func (v *CompletionProposalIfc) Equal(other ICompletionProposal) (result bool) {
	iv, err := _I.Get(60, "CompletionProposal", "equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if other != nil {
		tmp = other.P_CompletionProposal()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_other := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_other}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gtk_source_completion_proposal_get_gicon
// container is not nil, container is CompletionProposal
// is method
func (v *CompletionProposalIfc) GetGicon() (result gio.Icon) {
	iv, err := _I.Get(61, "CompletionProposal", "get_gicon")
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

// gtk_source_completion_proposal_get_icon
// container is not nil, container is CompletionProposal
// is method
func (v *CompletionProposalIfc) GetIcon() (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(62, "CompletionProposal", "get_icon")
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

// gtk_source_completion_proposal_get_icon_name
// container is not nil, container is CompletionProposal
// is method
func (v *CompletionProposalIfc) GetIconName() (result string) {
	iv, err := _I.Get(63, "CompletionProposal", "get_icon_name")
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

// gtk_source_completion_proposal_get_info
// container is not nil, container is CompletionProposal
// is method
func (v *CompletionProposalIfc) GetInfo() (result string) {
	iv, err := _I.Get(64, "CompletionProposal", "get_info")
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

// gtk_source_completion_proposal_get_label
// container is not nil, container is CompletionProposal
// is method
func (v *CompletionProposalIfc) GetLabel() (result string) {
	iv, err := _I.Get(65, "CompletionProposal", "get_label")
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

// gtk_source_completion_proposal_get_markup
// container is not nil, container is CompletionProposal
// is method
func (v *CompletionProposalIfc) GetMarkup() (result string) {
	iv, err := _I.Get(66, "CompletionProposal", "get_markup")
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

// gtk_source_completion_proposal_get_text
// container is not nil, container is CompletionProposal
// is method
func (v *CompletionProposalIfc) GetText() (result string) {
	iv, err := _I.Get(67, "CompletionProposal", "get_text")
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

// gtk_source_completion_proposal_hash
// container is not nil, container is CompletionProposal
// is method
func (v *CompletionProposalIfc) Hash() (result uint32) {
	iv, err := _I.Get(68, "CompletionProposal", "hash")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// ignore GType struct CompletionProposalIface
// Interface CompletionProvider
type CompletionProvider struct {
	CompletionProviderIfc
	P unsafe.Pointer
}
type CompletionProviderIfc struct{}
type ICompletionProvider interface{ P_CompletionProvider() unsafe.Pointer }

func (v CompletionProvider) P_CompletionProvider() unsafe.Pointer { return v.P }
func CompletionProviderGetType() gi.GType {
	ret := _I.GetGType(16, "CompletionProvider")
	return ret
}

// gtk_source_completion_provider_activate_proposal
// container is not nil, container is CompletionProvider
// is method
func (v *CompletionProviderIfc) ActivateProposal(proposal ICompletionProposal, iter gtk.TextIter) (result bool) {
	iv, err := _I.Get(69, "CompletionProvider", "activate_proposal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if proposal != nil {
		tmp = proposal.P_CompletionProposal()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_proposal := gi.NewPointerArgument(tmp)
	arg_iter := gi.NewPointerArgument(iter.P)
	args := []gi.Argument{arg_v, arg_proposal, arg_iter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gtk_source_completion_provider_get_activation
// container is not nil, container is CompletionProvider
// is method
func (v *CompletionProviderIfc) GetActivation() (result CompletionActivationFlags) {
	iv, err := _I.Get(70, "CompletionProvider", "get_activation")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = CompletionActivationFlags(ret.Int())
	return
}

// gtk_source_completion_provider_get_gicon
// container is not nil, container is CompletionProvider
// is method
func (v *CompletionProviderIfc) GetGicon() (result gio.Icon) {
	iv, err := _I.Get(71, "CompletionProvider", "get_gicon")
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

// gtk_source_completion_provider_get_icon
// container is not nil, container is CompletionProvider
// is method
func (v *CompletionProviderIfc) GetIcon() (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(72, "CompletionProvider", "get_icon")
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

// gtk_source_completion_provider_get_icon_name
// container is not nil, container is CompletionProvider
// is method
func (v *CompletionProviderIfc) GetIconName() (result string) {
	iv, err := _I.Get(73, "CompletionProvider", "get_icon_name")
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

// gtk_source_completion_provider_get_info_widget
// container is not nil, container is CompletionProvider
// is method
func (v *CompletionProviderIfc) GetInfoWidget(proposal ICompletionProposal) (result gtk.Widget) {
	iv, err := _I.Get(74, "CompletionProvider", "get_info_widget")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if proposal != nil {
		tmp = proposal.P_CompletionProposal()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_proposal := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_proposal}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_completion_provider_get_interactive_delay
// container is not nil, container is CompletionProvider
// is method
func (v *CompletionProviderIfc) GetInteractiveDelay() (result int32) {
	iv, err := _I.Get(75, "CompletionProvider", "get_interactive_delay")
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

// gtk_source_completion_provider_get_name
// container is not nil, container is CompletionProvider
// is method
func (v *CompletionProviderIfc) GetName() (result string) {
	iv, err := _I.Get(76, "CompletionProvider", "get_name")
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

// gtk_source_completion_provider_get_priority
// container is not nil, container is CompletionProvider
// is method
func (v *CompletionProviderIfc) GetPriority() (result int32) {
	iv, err := _I.Get(77, "CompletionProvider", "get_priority")
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

// gtk_source_completion_provider_get_start_iter
// container is not nil, container is CompletionProvider
// is method
func (v *CompletionProviderIfc) GetStartIter(context ICompletionContext, proposal ICompletionProposal, iter gtk.TextIter) (result bool) {
	iv, err := _I.Get(78, "CompletionProvider", "get_start_iter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_CompletionContext()
	}
	var tmp1 unsafe.Pointer
	if proposal != nil {
		tmp1 = proposal.P_CompletionProposal()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_context := gi.NewPointerArgument(tmp)
	arg_proposal := gi.NewPointerArgument(tmp1)
	arg_iter := gi.NewPointerArgument(iter.P)
	args := []gi.Argument{arg_v, arg_context, arg_proposal, arg_iter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gtk_source_completion_provider_match
// container is not nil, container is CompletionProvider
// is method
func (v *CompletionProviderIfc) Match(context ICompletionContext) (result bool) {
	iv, err := _I.Get(79, "CompletionProvider", "match")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_CompletionContext()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_context := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_context}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gtk_source_completion_provider_populate
// container is not nil, container is CompletionProvider
// is method
func (v *CompletionProviderIfc) Populate(context ICompletionContext) {
	iv, err := _I.Get(80, "CompletionProvider", "populate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_CompletionContext()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_context := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_context}
	iv.Call(args, nil, nil)
}

// gtk_source_completion_provider_update_info
// container is not nil, container is CompletionProvider
// is method
func (v *CompletionProviderIfc) UpdateInfo(proposal ICompletionProposal, info ICompletionInfo) {
	iv, err := _I.Get(81, "CompletionProvider", "update_info")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if proposal != nil {
		tmp = proposal.P_CompletionProposal()
	}
	var tmp1 unsafe.Pointer
	if info != nil {
		tmp1 = info.P_CompletionInfo()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_proposal := gi.NewPointerArgument(tmp)
	arg_info := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_v, arg_proposal, arg_info}
	iv.Call(args, nil, nil)
}

// ignore GType struct CompletionProviderIface
// Object CompletionWords
type CompletionWords struct {
	CompletionProviderIfc
	gobject.Object
}

func WrapCompletionWords(p unsafe.Pointer) (r CompletionWords) { r.P = p; return }

type ICompletionWords interface{ P_CompletionWords() unsafe.Pointer }

func (v CompletionWords) P_CompletionWords() unsafe.Pointer    { return v.P }
func (v CompletionWords) P_CompletionProvider() unsafe.Pointer { return v.P }
func CompletionWordsGetType() gi.GType {
	ret := _I.GetGType(17, "CompletionWords")
	return ret
}

// gtk_source_completion_words_new
// container is not nil, container is CompletionWords
// is constructor
func NewCompletionWords(name string, icon gdkpixbuf.IPixbuf) (result CompletionWords) {
	iv, err := _I.Get(82, "CompletionWords", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	var tmp unsafe.Pointer
	if icon != nil {
		tmp = icon.P_Pixbuf()
	}
	arg_name := gi.NewStringArgument(c_name)
	arg_icon := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_name, arg_icon}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// gtk_source_completion_words_register
// container is not nil, container is CompletionWords
// is method
func (v CompletionWords) Register(buffer gtk.ITextBuffer) {
	iv, err := _I.Get(83, "CompletionWords", "register")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if buffer != nil {
		tmp = buffer.P_TextBuffer()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_buffer}
	iv.Call(args, nil, nil)
}

// gtk_source_completion_words_unregister
// container is not nil, container is CompletionWords
// is method
func (v CompletionWords) Unregister(buffer gtk.ITextBuffer) {
	iv, err := _I.Get(84, "CompletionWords", "unregister")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if buffer != nil {
		tmp = buffer.P_TextBuffer()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_buffer}
	iv.Call(args, nil, nil)
}

// ignore GType struct CompletionWordsClass
// Struct CompletionWordsPrivate
type CompletionWordsPrivate struct {
	P unsafe.Pointer
}

func CompletionWordsPrivateGetType() gi.GType {
	ret := _I.GetGType(18, "CompletionWordsPrivate")
	return ret
}

// Enum CompressionType
type CompressionTypeEnum int

const (
	CompressionTypeNone CompressionTypeEnum = 0
	CompressionTypeGzip CompressionTypeEnum = 1
)

func CompressionTypeGetType() gi.GType {
	ret := _I.GetGType(19, "CompressionType")
	return ret
}

// Struct Encoding
type Encoding struct {
	P unsafe.Pointer
}

func EncodingGetType() gi.GType {
	ret := _I.GetGType(20, "Encoding")
	return ret
}

// gtk_source_encoding_copy
// container is not nil, container is Encoding
// is method
func (v Encoding) Copy() (result Encoding) {
	iv, err := _I.Get(85, "Encoding", "copy")
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

// gtk_source_encoding_free
// container is not nil, container is Encoding
// is method
func (v Encoding) Free() {
	iv, err := _I.Get(86, "Encoding", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_encoding_get_charset
// container is not nil, container is Encoding
// is method
func (v Encoding) GetCharset() (result string) {
	iv, err := _I.Get(87, "Encoding", "get_charset")
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

// gtk_source_encoding_get_name
// container is not nil, container is Encoding
// is method
func (v Encoding) GetName() (result string) {
	iv, err := _I.Get(88, "Encoding", "get_name")
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

// gtk_source_encoding_to_string
// container is not nil, container is Encoding
// is method
func (v Encoding) ToString() (result string) {
	iv, err := _I.Get(89, "Encoding", "to_string")
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

// gtk_source_encoding_get_from_charset
// container is not nil, container is Encoding
// is method
// arg0Type tag: utf8, isPtr: true
func EncodingGetFromCharset1(charset string) (result Encoding) {
	iv, err := _I.Get(93, "Encoding", "get_from_charset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_charset := gi.CString(charset)
	arg_charset := gi.NewStringArgument(c_charset)
	args := []gi.Argument{arg_charset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_charset)
	result.P = ret.Pointer()
	return
}

// Object File
type File struct {
	gobject.Object
}

func WrapFile(p unsafe.Pointer) (r File) { r.P = p; return }

type IFile interface{ P_File() unsafe.Pointer }

func (v File) P_File() unsafe.Pointer { return v.P }
func FileGetType() gi.GType {
	ret := _I.GetGType(21, "File")
	return ret
}

// gtk_source_file_new
// container is not nil, container is File
// is constructor
func NewFile() (result File) {
	iv, err := _I.Get(95, "File", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_file_check_file_on_disk
// container is not nil, container is File
// is method
func (v File) CheckFileOnDisk() {
	iv, err := _I.Get(96, "File", "check_file_on_disk")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_file_get_compression_type
// container is not nil, container is File
// is method
func (v File) GetCompressionType() (result CompressionTypeEnum) {
	iv, err := _I.Get(97, "File", "get_compression_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = CompressionTypeEnum(ret.Int())
	return
}

// gtk_source_file_get_encoding
// container is not nil, container is File
// is method
func (v File) GetEncoding() (result Encoding) {
	iv, err := _I.Get(98, "File", "get_encoding")
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

// gtk_source_file_get_location
// container is not nil, container is File
// is method
func (v File) GetLocation() (result gio.File) {
	iv, err := _I.Get(99, "File", "get_location")
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

// gtk_source_file_get_newline_type
// container is not nil, container is File
// is method
func (v File) GetNewlineType() (result NewlineTypeEnum) {
	iv, err := _I.Get(100, "File", "get_newline_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = NewlineTypeEnum(ret.Int())
	return
}

// gtk_source_file_is_deleted
// container is not nil, container is File
// is method
func (v File) IsDeleted() (result bool) {
	iv, err := _I.Get(101, "File", "is_deleted")
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

// gtk_source_file_is_externally_modified
// container is not nil, container is File
// is method
func (v File) IsExternallyModified() (result bool) {
	iv, err := _I.Get(102, "File", "is_externally_modified")
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

// gtk_source_file_is_local
// container is not nil, container is File
// is method
func (v File) IsLocal() (result bool) {
	iv, err := _I.Get(103, "File", "is_local")
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

// gtk_source_file_is_readonly
// container is not nil, container is File
// is method
func (v File) IsReadonly() (result bool) {
	iv, err := _I.Get(104, "File", "is_readonly")
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

// gtk_source_file_set_location
// container is not nil, container is File
// is method
func (v File) SetLocation(location gio.IFile) {
	iv, err := _I.Get(105, "File", "set_location")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if location != nil {
		tmp = location.P_File()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_location := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_location}
	iv.Call(args, nil, nil)
}

// ignore GType struct FileClass
// Object FileLoader
type FileLoader struct {
	gobject.Object
}

func WrapFileLoader(p unsafe.Pointer) (r FileLoader) { r.P = p; return }

type IFileLoader interface{ P_FileLoader() unsafe.Pointer }

func (v FileLoader) P_FileLoader() unsafe.Pointer { return v.P }
func FileLoaderGetType() gi.GType {
	ret := _I.GetGType(22, "FileLoader")
	return ret
}

// gtk_source_file_loader_new
// container is not nil, container is FileLoader
// is constructor
func NewFileLoader(buffer IBuffer, file IFile) (result FileLoader) {
	iv, err := _I.Get(106, "FileLoader", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if buffer != nil {
		tmp = buffer.P_Buffer()
	}
	var tmp1 unsafe.Pointer
	if file != nil {
		tmp1 = file.P_File()
	}
	arg_buffer := gi.NewPointerArgument(tmp)
	arg_file := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_buffer, arg_file}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_file_loader_new_from_stream
// container is not nil, container is FileLoader
// is constructor
func NewFileLoaderFromStream(buffer IBuffer, file IFile, stream gio.IInputStream) (result FileLoader) {
	iv, err := _I.Get(107, "FileLoader", "new_from_stream")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if buffer != nil {
		tmp = buffer.P_Buffer()
	}
	var tmp1 unsafe.Pointer
	if file != nil {
		tmp1 = file.P_File()
	}
	var tmp2 unsafe.Pointer
	if stream != nil {
		tmp2 = stream.P_InputStream()
	}
	arg_buffer := gi.NewPointerArgument(tmp)
	arg_file := gi.NewPointerArgument(tmp1)
	arg_stream := gi.NewPointerArgument(tmp2)
	args := []gi.Argument{arg_buffer, arg_file, arg_stream}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_file_loader_get_buffer
// container is not nil, container is FileLoader
// is method
func (v FileLoader) GetBuffer() (result Buffer) {
	iv, err := _I.Get(108, "FileLoader", "get_buffer")
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

// gtk_source_file_loader_get_compression_type
// container is not nil, container is FileLoader
// is method
func (v FileLoader) GetCompressionType() (result CompressionTypeEnum) {
	iv, err := _I.Get(109, "FileLoader", "get_compression_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = CompressionTypeEnum(ret.Int())
	return
}

// gtk_source_file_loader_get_encoding
// container is not nil, container is FileLoader
// is method
func (v FileLoader) GetEncoding() (result Encoding) {
	iv, err := _I.Get(110, "FileLoader", "get_encoding")
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

// gtk_source_file_loader_get_file
// container is not nil, container is FileLoader
// is method
func (v FileLoader) GetFile() (result File) {
	iv, err := _I.Get(111, "FileLoader", "get_file")
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

// gtk_source_file_loader_get_input_stream
// container is not nil, container is FileLoader
// is method
func (v FileLoader) GetInputStream() (result gio.InputStream) {
	iv, err := _I.Get(112, "FileLoader", "get_input_stream")
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

// gtk_source_file_loader_get_location
// container is not nil, container is FileLoader
// is method
func (v FileLoader) GetLocation() (result gio.File) {
	iv, err := _I.Get(113, "FileLoader", "get_location")
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

// gtk_source_file_loader_get_newline_type
// container is not nil, container is FileLoader
// is method
func (v FileLoader) GetNewlineType() (result NewlineTypeEnum) {
	iv, err := _I.Get(114, "FileLoader", "get_newline_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = NewlineTypeEnum(ret.Int())
	return
}

// gtk_source_file_loader_load_async
// container is not nil, container is FileLoader
// is method
func (v FileLoader) LoadAsync(io_priority int32, cancellable gio.ICancellable, progress_callback int /*TODO_TYPE CALLBACK*/, progress_callback_data unsafe.Pointer, progress_callback_notify int /*TODO_TYPE CALLBACK*/, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(115, "FileLoader", "load_async")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_io_priority := gi.NewInt32Argument(io_priority)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_progress_callback := gi.NewPointerArgument(unsafe.Pointer(gio.GetPointer_myFileProgressCallback()))
	arg_progress_callback_data := gi.NewPointerArgument(progress_callback_data)
	arg_progress_callback_notify := gi.NewPointerArgument(unsafe.Pointer(glib.GetPointer_myDestroyNotify()))
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(gio.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_io_priority, arg_cancellable, arg_progress_callback, arg_progress_callback_data, arg_progress_callback_notify, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// gtk_source_file_loader_load_finish
// container is not nil, container is FileLoader
// is method
func (v FileLoader) LoadFinish(result gio.IAsyncResult) (result1 bool, err error) {
	iv, err := _I.Get(116, "FileLoader", "load_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if result != nil {
		tmp = result.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_result := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_result, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result1 = ret.Bool()
	return
}

// gtk_source_file_loader_set_candidate_encodings
// container is not nil, container is FileLoader
// is method
func (v FileLoader) SetCandidateEncodings(candidate_encodings glib.SList) {
	iv, err := _I.Get(117, "FileLoader", "set_candidate_encodings")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_candidate_encodings := gi.NewPointerArgument(candidate_encodings.P)
	args := []gi.Argument{arg_v, arg_candidate_encodings}
	iv.Call(args, nil, nil)
}

// ignore GType struct FileLoaderClass
// Enum FileLoaderError
type FileLoaderErrorEnum int

const (
	FileLoaderErrorTooBig                      FileLoaderErrorEnum = 0
	FileLoaderErrorEncodingAutoDetectionFailed FileLoaderErrorEnum = 1
	FileLoaderErrorConversionFallback          FileLoaderErrorEnum = 2
)

func FileLoaderErrorGetType() gi.GType {
	ret := _I.GetGType(23, "FileLoaderError")
	return ret
}

// Struct FileLoaderPrivate
type FileLoaderPrivate struct {
	P unsafe.Pointer
}

func FileLoaderPrivateGetType() gi.GType {
	ret := _I.GetGType(24, "FileLoaderPrivate")
	return ret
}

// Struct FilePrivate
type FilePrivate struct {
	P unsafe.Pointer
}

func FilePrivateGetType() gi.GType {
	ret := _I.GetGType(25, "FilePrivate")
	return ret
}

// Object FileSaver
type FileSaver struct {
	gobject.Object
}

func WrapFileSaver(p unsafe.Pointer) (r FileSaver) { r.P = p; return }

type IFileSaver interface{ P_FileSaver() unsafe.Pointer }

func (v FileSaver) P_FileSaver() unsafe.Pointer { return v.P }
func FileSaverGetType() gi.GType {
	ret := _I.GetGType(26, "FileSaver")
	return ret
}

// gtk_source_file_saver_new
// container is not nil, container is FileSaver
// is constructor
func NewFileSaver(buffer IBuffer, file IFile) (result FileSaver) {
	iv, err := _I.Get(118, "FileSaver", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if buffer != nil {
		tmp = buffer.P_Buffer()
	}
	var tmp1 unsafe.Pointer
	if file != nil {
		tmp1 = file.P_File()
	}
	arg_buffer := gi.NewPointerArgument(tmp)
	arg_file := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_buffer, arg_file}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_file_saver_new_with_target
// container is not nil, container is FileSaver
// is constructor
func NewFileSaverWithTarget(buffer IBuffer, file IFile, target_location gio.IFile) (result FileSaver) {
	iv, err := _I.Get(119, "FileSaver", "new_with_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if buffer != nil {
		tmp = buffer.P_Buffer()
	}
	var tmp1 unsafe.Pointer
	if file != nil {
		tmp1 = file.P_File()
	}
	var tmp2 unsafe.Pointer
	if target_location != nil {
		tmp2 = target_location.P_File()
	}
	arg_buffer := gi.NewPointerArgument(tmp)
	arg_file := gi.NewPointerArgument(tmp1)
	arg_target_location := gi.NewPointerArgument(tmp2)
	args := []gi.Argument{arg_buffer, arg_file, arg_target_location}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_file_saver_get_buffer
// container is not nil, container is FileSaver
// is method
func (v FileSaver) GetBuffer() (result Buffer) {
	iv, err := _I.Get(120, "FileSaver", "get_buffer")
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

// gtk_source_file_saver_get_compression_type
// container is not nil, container is FileSaver
// is method
func (v FileSaver) GetCompressionType() (result CompressionTypeEnum) {
	iv, err := _I.Get(121, "FileSaver", "get_compression_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = CompressionTypeEnum(ret.Int())
	return
}

// gtk_source_file_saver_get_encoding
// container is not nil, container is FileSaver
// is method
func (v FileSaver) GetEncoding() (result Encoding) {
	iv, err := _I.Get(122, "FileSaver", "get_encoding")
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

// gtk_source_file_saver_get_file
// container is not nil, container is FileSaver
// is method
func (v FileSaver) GetFile() (result File) {
	iv, err := _I.Get(123, "FileSaver", "get_file")
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

// gtk_source_file_saver_get_flags
// container is not nil, container is FileSaver
// is method
func (v FileSaver) GetFlags() (result FileSaverFlags) {
	iv, err := _I.Get(124, "FileSaver", "get_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FileSaverFlags(ret.Int())
	return
}

// gtk_source_file_saver_get_location
// container is not nil, container is FileSaver
// is method
func (v FileSaver) GetLocation() (result gio.File) {
	iv, err := _I.Get(125, "FileSaver", "get_location")
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

// gtk_source_file_saver_get_newline_type
// container is not nil, container is FileSaver
// is method
func (v FileSaver) GetNewlineType() (result NewlineTypeEnum) {
	iv, err := _I.Get(126, "FileSaver", "get_newline_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = NewlineTypeEnum(ret.Int())
	return
}

// gtk_source_file_saver_save_async
// container is not nil, container is FileSaver
// is method
func (v FileSaver) SaveAsync(io_priority int32, cancellable gio.ICancellable, progress_callback int /*TODO_TYPE CALLBACK*/, progress_callback_data unsafe.Pointer, progress_callback_notify int /*TODO_TYPE CALLBACK*/, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(127, "FileSaver", "save_async")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_io_priority := gi.NewInt32Argument(io_priority)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_progress_callback := gi.NewPointerArgument(unsafe.Pointer(gio.GetPointer_myFileProgressCallback()))
	arg_progress_callback_data := gi.NewPointerArgument(progress_callback_data)
	arg_progress_callback_notify := gi.NewPointerArgument(unsafe.Pointer(glib.GetPointer_myDestroyNotify()))
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(gio.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_io_priority, arg_cancellable, arg_progress_callback, arg_progress_callback_data, arg_progress_callback_notify, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// gtk_source_file_saver_save_finish
// container is not nil, container is FileSaver
// is method
func (v FileSaver) SaveFinish(result gio.IAsyncResult) (result1 bool, err error) {
	iv, err := _I.Get(128, "FileSaver", "save_finish")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	var tmp unsafe.Pointer
	if result != nil {
		tmp = result.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_result := gi.NewPointerArgument(tmp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_result, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result1 = ret.Bool()
	return
}

// gtk_source_file_saver_set_compression_type
// container is not nil, container is FileSaver
// is method
func (v FileSaver) SetCompressionType(compression_type CompressionTypeEnum) {
	iv, err := _I.Get(129, "FileSaver", "set_compression_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_compression_type := gi.NewIntArgument(int(compression_type))
	args := []gi.Argument{arg_v, arg_compression_type}
	iv.Call(args, nil, nil)
}

// gtk_source_file_saver_set_encoding
// container is not nil, container is FileSaver
// is method
func (v FileSaver) SetEncoding(encoding Encoding) {
	iv, err := _I.Get(130, "FileSaver", "set_encoding")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_encoding := gi.NewPointerArgument(encoding.P)
	args := []gi.Argument{arg_v, arg_encoding}
	iv.Call(args, nil, nil)
}

// gtk_source_file_saver_set_flags
// container is not nil, container is FileSaver
// is method
func (v FileSaver) SetFlags(flags FileSaverFlags) {
	iv, err := _I.Get(131, "FileSaver", "set_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_flags}
	iv.Call(args, nil, nil)
}

// gtk_source_file_saver_set_newline_type
// container is not nil, container is FileSaver
// is method
func (v FileSaver) SetNewlineType(newline_type NewlineTypeEnum) {
	iv, err := _I.Get(132, "FileSaver", "set_newline_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_newline_type := gi.NewIntArgument(int(newline_type))
	args := []gi.Argument{arg_v, arg_newline_type}
	iv.Call(args, nil, nil)
}

// ignore GType struct FileSaverClass
// Enum FileSaverError
type FileSaverErrorEnum int

const (
	FileSaverErrorInvalidChars       FileSaverErrorEnum = 0
	FileSaverErrorExternallyModified FileSaverErrorEnum = 1
)

func FileSaverErrorGetType() gi.GType {
	ret := _I.GetGType(27, "FileSaverError")
	return ret
}

// Flags FileSaverFlags
type FileSaverFlags int

const (
	FileSaverFlagsNone                   FileSaverFlags = 0
	FileSaverFlagsIgnoreInvalidChars     FileSaverFlags = 1
	FileSaverFlagsIgnoreModificationTime FileSaverFlags = 2
	FileSaverFlagsCreateBackup           FileSaverFlags = 4
)

func FileSaverFlagsGetType() gi.GType {
	ret := _I.GetGType(28, "FileSaverFlags")
	return ret
}

// Struct FileSaverPrivate
type FileSaverPrivate struct {
	P unsafe.Pointer
}

func FileSaverPrivateGetType() gi.GType {
	ret := _I.GetGType(29, "FileSaverPrivate")
	return ret
}

// Object Gutter
type Gutter struct {
	gobject.Object
}

func WrapGutter(p unsafe.Pointer) (r Gutter) { r.P = p; return }

type IGutter interface{ P_Gutter() unsafe.Pointer }

func (v Gutter) P_Gutter() unsafe.Pointer { return v.P }
func GutterGetType() gi.GType {
	ret := _I.GetGType(30, "Gutter")
	return ret
}

// gtk_source_gutter_get_renderer_at_pos
// container is not nil, container is Gutter
// is method
func (v Gutter) GetRendererAtPos(x int32, y int32) (result GutterRenderer) {
	iv, err := _I.Get(133, "Gutter", "get_renderer_at_pos")
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

// gtk_source_gutter_get_view
// container is not nil, container is Gutter
// is method
func (v Gutter) GetView() (result View) {
	iv, err := _I.Get(134, "Gutter", "get_view")
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

// gtk_source_gutter_get_window_type
// container is not nil, container is Gutter
// is method
func (v Gutter) GetWindowType() (result gtk.TextWindowTypeEnum) {
	iv, err := _I.Get(135, "Gutter", "get_window_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gtk.TextWindowTypeEnum(ret.Int())
	return
}

// gtk_source_gutter_insert
// container is not nil, container is Gutter
// is method
func (v Gutter) Insert(renderer IGutterRenderer, position int32) (result bool) {
	iv, err := _I.Get(136, "Gutter", "insert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if renderer != nil {
		tmp = renderer.P_GutterRenderer()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_renderer := gi.NewPointerArgument(tmp)
	arg_position := gi.NewInt32Argument(position)
	args := []gi.Argument{arg_v, arg_renderer, arg_position}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gtk_source_gutter_queue_draw
// container is not nil, container is Gutter
// is method
func (v Gutter) QueueDraw() {
	iv, err := _I.Get(137, "Gutter", "queue_draw")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_gutter_remove
// container is not nil, container is Gutter
// is method
func (v Gutter) Remove(renderer IGutterRenderer) {
	iv, err := _I.Get(138, "Gutter", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if renderer != nil {
		tmp = renderer.P_GutterRenderer()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_renderer := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_renderer}
	iv.Call(args, nil, nil)
}

// gtk_source_gutter_reorder
// container is not nil, container is Gutter
// is method
func (v Gutter) Reorder(renderer IGutterRenderer, position int32) {
	iv, err := _I.Get(139, "Gutter", "reorder")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if renderer != nil {
		tmp = renderer.P_GutterRenderer()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_renderer := gi.NewPointerArgument(tmp)
	arg_position := gi.NewInt32Argument(position)
	args := []gi.Argument{arg_v, arg_renderer, arg_position}
	iv.Call(args, nil, nil)
}

// ignore GType struct GutterClass
// Struct GutterPrivate
type GutterPrivate struct {
	P unsafe.Pointer
}

func GutterPrivateGetType() gi.GType {
	ret := _I.GetGType(31, "GutterPrivate")
	return ret
}

// Object GutterRenderer
type GutterRenderer struct {
	gobject.InitiallyUnowned
}

func WrapGutterRenderer(p unsafe.Pointer) (r GutterRenderer) { r.P = p; return }

type IGutterRenderer interface{ P_GutterRenderer() unsafe.Pointer }

func (v GutterRenderer) P_GutterRenderer() unsafe.Pointer { return v.P }
func GutterRendererGetType() gi.GType {
	ret := _I.GetGType(32, "GutterRenderer")
	return ret
}

// gtk_source_gutter_renderer_activate
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) Activate(iter gtk.TextIter, area gdk.Rectangle, event gdk.Event) {
	iv, err := _I.Get(140, "GutterRenderer", "activate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_area := gi.NewPointerArgument(area.P)
	arg_event := gi.NewPointerArgument(event.P)
	args := []gi.Argument{arg_v, arg_iter, arg_area, arg_event}
	iv.Call(args, nil, nil)
}

// gtk_source_gutter_renderer_begin
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) Begin(cr cairo.Context, background_area gdk.Rectangle, cell_area gdk.Rectangle, start gtk.TextIter, end gtk.TextIter) {
	iv, err := _I.Get(141, "GutterRenderer", "begin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_background_area := gi.NewPointerArgument(background_area.P)
	arg_cell_area := gi.NewPointerArgument(cell_area.P)
	arg_start := gi.NewPointerArgument(start.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_v, arg_cr, arg_background_area, arg_cell_area, arg_start, arg_end}
	iv.Call(args, nil, nil)
}

// gtk_source_gutter_renderer_draw
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) Draw(cr cairo.Context, background_area gdk.Rectangle, cell_area gdk.Rectangle, start gtk.TextIter, end gtk.TextIter, state GutterRendererStateFlags) {
	iv, err := _I.Get(142, "GutterRenderer", "draw")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_cr := gi.NewPointerArgument(cr.P)
	arg_background_area := gi.NewPointerArgument(background_area.P)
	arg_cell_area := gi.NewPointerArgument(cell_area.P)
	arg_start := gi.NewPointerArgument(start.P)
	arg_end := gi.NewPointerArgument(end.P)
	arg_state := gi.NewIntArgument(int(state))
	args := []gi.Argument{arg_v, arg_cr, arg_background_area, arg_cell_area, arg_start, arg_end, arg_state}
	iv.Call(args, nil, nil)
}

// gtk_source_gutter_renderer_end
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) End() {
	iv, err := _I.Get(143, "GutterRenderer", "end")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_gutter_renderer_get_alignment
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) GetAlignment() (xalign float32, yalign float32) {
	iv, err := _I.Get(144, "GutterRenderer", "get_alignment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_xalign := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_yalign := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_xalign, arg_yalign}
	iv.Call(args, nil, &outArgs[0])
	xalign = outArgs[0].Float()
	yalign = outArgs[1].Float()
	return
}

// gtk_source_gutter_renderer_get_alignment_mode
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) GetAlignmentMode() (result GutterRendererAlignmentModeEnum) {
	iv, err := _I.Get(145, "GutterRenderer", "get_alignment_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = GutterRendererAlignmentModeEnum(ret.Int())
	return
}

// gtk_source_gutter_renderer_get_background
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) GetBackground(color gdk.RGBA) (result bool) {
	iv, err := _I.Get(146, "GutterRenderer", "get_background")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_color := gi.NewPointerArgument(color.P)
	args := []gi.Argument{arg_v, arg_color}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gtk_source_gutter_renderer_get_padding
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) GetPadding() (xpad int32, ypad int32) {
	iv, err := _I.Get(147, "GutterRenderer", "get_padding")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_xpad := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_ypad := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_xpad, arg_ypad}
	iv.Call(args, nil, &outArgs[0])
	xpad = outArgs[0].Int32()
	ypad = outArgs[1].Int32()
	return
}

// gtk_source_gutter_renderer_get_size
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) GetSize() (result int32) {
	iv, err := _I.Get(148, "GutterRenderer", "get_size")
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

// gtk_source_gutter_renderer_get_view
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) GetView() (result gtk.TextView) {
	iv, err := _I.Get(149, "GutterRenderer", "get_view")
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

// gtk_source_gutter_renderer_get_visible
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) GetVisible() (result bool) {
	iv, err := _I.Get(150, "GutterRenderer", "get_visible")
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

// gtk_source_gutter_renderer_get_window_type
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) GetWindowType() (result gtk.TextWindowTypeEnum) {
	iv, err := _I.Get(151, "GutterRenderer", "get_window_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gtk.TextWindowTypeEnum(ret.Int())
	return
}

// gtk_source_gutter_renderer_query_activatable
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) QueryActivatable(iter gtk.TextIter, area gdk.Rectangle, event gdk.Event) (result bool) {
	iv, err := _I.Get(152, "GutterRenderer", "query_activatable")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_area := gi.NewPointerArgument(area.P)
	arg_event := gi.NewPointerArgument(event.P)
	args := []gi.Argument{arg_v, arg_iter, arg_area, arg_event}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gtk_source_gutter_renderer_query_data
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) QueryData(start gtk.TextIter, end gtk.TextIter, state GutterRendererStateFlags) {
	iv, err := _I.Get(153, "GutterRenderer", "query_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewPointerArgument(start.P)
	arg_end := gi.NewPointerArgument(end.P)
	arg_state := gi.NewIntArgument(int(state))
	args := []gi.Argument{arg_v, arg_start, arg_end, arg_state}
	iv.Call(args, nil, nil)
}

// gtk_source_gutter_renderer_query_tooltip
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) QueryTooltip(iter gtk.TextIter, area gdk.Rectangle, x int32, y int32, tooltip gtk.ITooltip) (result bool) {
	iv, err := _I.Get(154, "GutterRenderer", "query_tooltip")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if tooltip != nil {
		tmp = tooltip.P_Tooltip()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_area := gi.NewPointerArgument(area.P)
	arg_x := gi.NewInt32Argument(x)
	arg_y := gi.NewInt32Argument(y)
	arg_tooltip := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_iter, arg_area, arg_x, arg_y, arg_tooltip}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gtk_source_gutter_renderer_queue_draw
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) QueueDraw() {
	iv, err := _I.Get(155, "GutterRenderer", "queue_draw")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_gutter_renderer_set_alignment
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) SetAlignment(xalign float32, yalign float32) {
	iv, err := _I.Get(156, "GutterRenderer", "set_alignment")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_xalign := gi.NewFloatArgument(xalign)
	arg_yalign := gi.NewFloatArgument(yalign)
	args := []gi.Argument{arg_v, arg_xalign, arg_yalign}
	iv.Call(args, nil, nil)
}

// gtk_source_gutter_renderer_set_alignment_mode
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) SetAlignmentMode(mode GutterRendererAlignmentModeEnum) {
	iv, err := _I.Get(157, "GutterRenderer", "set_alignment_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_v, arg_mode}
	iv.Call(args, nil, nil)
}

// gtk_source_gutter_renderer_set_background
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) SetBackground(color gdk.RGBA) {
	iv, err := _I.Get(158, "GutterRenderer", "set_background")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_color := gi.NewPointerArgument(color.P)
	args := []gi.Argument{arg_v, arg_color}
	iv.Call(args, nil, nil)
}

// gtk_source_gutter_renderer_set_padding
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) SetPadding(xpad int32, ypad int32) {
	iv, err := _I.Get(159, "GutterRenderer", "set_padding")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_xpad := gi.NewInt32Argument(xpad)
	arg_ypad := gi.NewInt32Argument(ypad)
	args := []gi.Argument{arg_v, arg_xpad, arg_ypad}
	iv.Call(args, nil, nil)
}

// gtk_source_gutter_renderer_set_size
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) SetSize(size int32) {
	iv, err := _I.Get(160, "GutterRenderer", "set_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewInt32Argument(size)
	args := []gi.Argument{arg_v, arg_size}
	iv.Call(args, nil, nil)
}

// gtk_source_gutter_renderer_set_visible
// container is not nil, container is GutterRenderer
// is method
func (v GutterRenderer) SetVisible(visible bool) {
	iv, err := _I.Get(161, "GutterRenderer", "set_visible")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_visible := gi.NewBoolArgument(visible)
	args := []gi.Argument{arg_v, arg_visible}
	iv.Call(args, nil, nil)
}

// Enum GutterRendererAlignmentMode
type GutterRendererAlignmentModeEnum int

const (
	GutterRendererAlignmentModeCell  GutterRendererAlignmentModeEnum = 0
	GutterRendererAlignmentModeFirst GutterRendererAlignmentModeEnum = 1
	GutterRendererAlignmentModeLast  GutterRendererAlignmentModeEnum = 2
)

func GutterRendererAlignmentModeGetType() gi.GType {
	ret := _I.GetGType(33, "GutterRendererAlignmentMode")
	return ret
}

// ignore GType struct GutterRendererClass
// Object GutterRendererPixbuf
type GutterRendererPixbuf struct {
	GutterRenderer
}

func WrapGutterRendererPixbuf(p unsafe.Pointer) (r GutterRendererPixbuf) { r.P = p; return }

type IGutterRendererPixbuf interface{ P_GutterRendererPixbuf() unsafe.Pointer }

func (v GutterRendererPixbuf) P_GutterRendererPixbuf() unsafe.Pointer { return v.P }
func GutterRendererPixbufGetType() gi.GType {
	ret := _I.GetGType(34, "GutterRendererPixbuf")
	return ret
}

// gtk_source_gutter_renderer_pixbuf_new
// container is not nil, container is GutterRendererPixbuf
// is constructor
func NewGutterRendererPixbuf() (result GutterRendererPixbuf) {
	iv, err := _I.Get(162, "GutterRendererPixbuf", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_gutter_renderer_pixbuf_get_gicon
// container is not nil, container is GutterRendererPixbuf
// is method
func (v GutterRendererPixbuf) GetGicon() (result gio.Icon) {
	iv, err := _I.Get(163, "GutterRendererPixbuf", "get_gicon")
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

// gtk_source_gutter_renderer_pixbuf_get_icon_name
// container is not nil, container is GutterRendererPixbuf
// is method
func (v GutterRendererPixbuf) GetIconName() (result string) {
	iv, err := _I.Get(164, "GutterRendererPixbuf", "get_icon_name")
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

// gtk_source_gutter_renderer_pixbuf_get_pixbuf
// container is not nil, container is GutterRendererPixbuf
// is method
func (v GutterRendererPixbuf) GetPixbuf() (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(165, "GutterRendererPixbuf", "get_pixbuf")
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

// gtk_source_gutter_renderer_pixbuf_set_gicon
// container is not nil, container is GutterRendererPixbuf
// is method
func (v GutterRendererPixbuf) SetGicon(icon gio.IIcon) {
	iv, err := _I.Get(166, "GutterRendererPixbuf", "set_gicon")
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

// gtk_source_gutter_renderer_pixbuf_set_icon_name
// container is not nil, container is GutterRendererPixbuf
// is method
func (v GutterRendererPixbuf) SetIconName(icon_name string) {
	iv, err := _I.Get(167, "GutterRendererPixbuf", "set_icon_name")
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

// gtk_source_gutter_renderer_pixbuf_set_pixbuf
// container is not nil, container is GutterRendererPixbuf
// is method
func (v GutterRendererPixbuf) SetPixbuf(pixbuf gdkpixbuf.IPixbuf) {
	iv, err := _I.Get(168, "GutterRendererPixbuf", "set_pixbuf")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pixbuf != nil {
		tmp = pixbuf.P_Pixbuf()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pixbuf := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_pixbuf}
	iv.Call(args, nil, nil)
}

// ignore GType struct GutterRendererPixbufClass
// Struct GutterRendererPixbufPrivate
type GutterRendererPixbufPrivate struct {
	P unsafe.Pointer
}

func GutterRendererPixbufPrivateGetType() gi.GType {
	ret := _I.GetGType(35, "GutterRendererPixbufPrivate")
	return ret
}

// Struct GutterRendererPrivate
type GutterRendererPrivate struct {
	P unsafe.Pointer
}

func GutterRendererPrivateGetType() gi.GType {
	ret := _I.GetGType(36, "GutterRendererPrivate")
	return ret
}

// Flags GutterRendererState
type GutterRendererStateFlags int

const (
	GutterRendererStateNormal   GutterRendererStateFlags = 0
	GutterRendererStateCursor   GutterRendererStateFlags = 1
	GutterRendererStatePrelit   GutterRendererStateFlags = 2
	GutterRendererStateSelected GutterRendererStateFlags = 4
)

func GutterRendererStateGetType() gi.GType {
	ret := _I.GetGType(37, "GutterRendererState")
	return ret
}

// Object GutterRendererText
type GutterRendererText struct {
	GutterRenderer
}

func WrapGutterRendererText(p unsafe.Pointer) (r GutterRendererText) { r.P = p; return }

type IGutterRendererText interface{ P_GutterRendererText() unsafe.Pointer }

func (v GutterRendererText) P_GutterRendererText() unsafe.Pointer { return v.P }
func GutterRendererTextGetType() gi.GType {
	ret := _I.GetGType(38, "GutterRendererText")
	return ret
}

// gtk_source_gutter_renderer_text_new
// container is not nil, container is GutterRendererText
// is constructor
func NewGutterRendererText() (result GutterRendererText) {
	iv, err := _I.Get(169, "GutterRendererText", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_gutter_renderer_text_measure
// container is not nil, container is GutterRendererText
// is method
func (v GutterRendererText) Measure(text string) (width int32, height int32) {
	iv, err := _I.Get(170, "GutterRendererText", "measure")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_text, arg_width, arg_height}
	iv.Call(args, nil, &outArgs[0])
	gi.Free(c_text)
	width = outArgs[0].Int32()
	height = outArgs[1].Int32()
	return
}

// gtk_source_gutter_renderer_text_measure_markup
// container is not nil, container is GutterRendererText
// is method
func (v GutterRendererText) MeasureMarkup(markup string) (width int32, height int32) {
	iv, err := _I.Get(171, "GutterRendererText", "measure_markup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	c_markup := gi.CString(markup)
	arg_v := gi.NewPointerArgument(v.P)
	arg_markup := gi.NewStringArgument(c_markup)
	arg_width := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_height := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_markup, arg_width, arg_height}
	iv.Call(args, nil, &outArgs[0])
	gi.Free(c_markup)
	width = outArgs[0].Int32()
	height = outArgs[1].Int32()
	return
}

// gtk_source_gutter_renderer_text_set_markup
// container is not nil, container is GutterRendererText
// is method
func (v GutterRendererText) SetMarkup(markup string, length int32) {
	iv, err := _I.Get(172, "GutterRendererText", "set_markup")
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

// gtk_source_gutter_renderer_text_set_text
// container is not nil, container is GutterRendererText
// is method
func (v GutterRendererText) SetText(text string, length int32) {
	iv, err := _I.Get(173, "GutterRendererText", "set_text")
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

// ignore GType struct GutterRendererTextClass
// Struct GutterRendererTextPrivate
type GutterRendererTextPrivate struct {
	P unsafe.Pointer
}

func GutterRendererTextPrivateGetType() gi.GType {
	ret := _I.GetGType(39, "GutterRendererTextPrivate")
	return ret
}

// Object Language
type Language struct {
	gobject.Object
}

func WrapLanguage(p unsafe.Pointer) (r Language) { r.P = p; return }

type ILanguage interface{ P_Language() unsafe.Pointer }

func (v Language) P_Language() unsafe.Pointer { return v.P }
func LanguageGetType() gi.GType {
	ret := _I.GetGType(40, "Language")
	return ret
}

// gtk_source_language_get_globs
// container is not nil, container is Language
// is method
func (v Language) GetGlobs() (result gi.CStrArray) {
	iv, err := _I.Get(174, "Language", "get_globs")
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

// gtk_source_language_get_hidden
// container is not nil, container is Language
// is method
func (v Language) GetHidden() (result bool) {
	iv, err := _I.Get(175, "Language", "get_hidden")
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

// gtk_source_language_get_id
// container is not nil, container is Language
// is method
func (v Language) GetId() (result string) {
	iv, err := _I.Get(176, "Language", "get_id")
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

// gtk_source_language_get_metadata
// container is not nil, container is Language
// is method
func (v Language) GetMetadata(name string) (result string) {
	iv, err := _I.Get(177, "Language", "get_metadata")
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
	result = ret.String().Take()
	return
}

// gtk_source_language_get_mime_types
// container is not nil, container is Language
// is method
func (v Language) GetMimeTypes() (result gi.CStrArray) {
	iv, err := _I.Get(178, "Language", "get_mime_types")
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

// gtk_source_language_get_name
// container is not nil, container is Language
// is method
func (v Language) GetName() (result string) {
	iv, err := _I.Get(179, "Language", "get_name")
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

// gtk_source_language_get_section
// container is not nil, container is Language
// is method
func (v Language) GetSection() (result string) {
	iv, err := _I.Get(180, "Language", "get_section")
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

// gtk_source_language_get_style_fallback
// container is not nil, container is Language
// is method
func (v Language) GetStyleFallback(style_id string) (result string) {
	iv, err := _I.Get(181, "Language", "get_style_fallback")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_style_id := gi.CString(style_id)
	arg_v := gi.NewPointerArgument(v.P)
	arg_style_id := gi.NewStringArgument(c_style_id)
	args := []gi.Argument{arg_v, arg_style_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_style_id)
	result = ret.String().Take()
	return
}

// gtk_source_language_get_style_ids
// container is not nil, container is Language
// is method
func (v Language) GetStyleIds() (result gi.CStrArray) {
	iv, err := _I.Get(182, "Language", "get_style_ids")
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

// gtk_source_language_get_style_name
// container is not nil, container is Language
// is method
func (v Language) GetStyleName(style_id string) (result string) {
	iv, err := _I.Get(183, "Language", "get_style_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_style_id := gi.CString(style_id)
	arg_v := gi.NewPointerArgument(v.P)
	arg_style_id := gi.NewStringArgument(c_style_id)
	args := []gi.Argument{arg_v, arg_style_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_style_id)
	result = ret.String().Take()
	return
}

// ignore GType struct LanguageClass
// Object LanguageManager
type LanguageManager struct {
	gobject.Object
}

func WrapLanguageManager(p unsafe.Pointer) (r LanguageManager) { r.P = p; return }

type ILanguageManager interface{ P_LanguageManager() unsafe.Pointer }

func (v LanguageManager) P_LanguageManager() unsafe.Pointer { return v.P }
func LanguageManagerGetType() gi.GType {
	ret := _I.GetGType(41, "LanguageManager")
	return ret
}

// gtk_source_language_manager_new
// container is not nil, container is LanguageManager
// is constructor
func NewLanguageManager() (result LanguageManager) {
	iv, err := _I.Get(184, "LanguageManager", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_language_manager_get_language
// container is not nil, container is LanguageManager
// is method
func (v LanguageManager) GetLanguage(id string) (result Language) {
	iv, err := _I.Get(186, "LanguageManager", "get_language")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_id := gi.CString(id)
	arg_v := gi.NewPointerArgument(v.P)
	arg_id := gi.NewStringArgument(c_id)
	args := []gi.Argument{arg_v, arg_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_id)
	result.P = ret.Pointer()
	return
}

// gtk_source_language_manager_get_language_ids
// container is not nil, container is LanguageManager
// is method
func (v LanguageManager) GetLanguageIds() (result gi.CStrArray) {
	iv, err := _I.Get(187, "LanguageManager", "get_language_ids")
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

// gtk_source_language_manager_get_search_path
// container is not nil, container is LanguageManager
// is method
func (v LanguageManager) GetSearchPath() (result gi.CStrArray) {
	iv, err := _I.Get(188, "LanguageManager", "get_search_path")
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

// gtk_source_language_manager_guess_language
// container is not nil, container is LanguageManager
// is method
func (v LanguageManager) GuessLanguage(filename string, content_type string) (result Language) {
	iv, err := _I.Get(189, "LanguageManager", "guess_language")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	c_content_type := gi.CString(content_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_content_type := gi.NewStringArgument(c_content_type)
	args := []gi.Argument{arg_v, arg_filename, arg_content_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	gi.Free(c_content_type)
	result.P = ret.Pointer()
	return
}

// gtk_source_language_manager_set_search_path
// container is not nil, container is LanguageManager
// is method
func (v LanguageManager) SetSearchPath(dirs gi.CStrArray) {
	iv, err := _I.Get(190, "LanguageManager", "set_search_path")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dirs := gi.NewPointerArgument(dirs.P)
	args := []gi.Argument{arg_v, arg_dirs}
	iv.Call(args, nil, nil)
}

// ignore GType struct LanguageManagerClass
// Struct LanguageManagerPrivate
type LanguageManagerPrivate struct {
	P unsafe.Pointer
}

func LanguageManagerPrivateGetType() gi.GType {
	ret := _I.GetGType(42, "LanguageManagerPrivate")
	return ret
}

// Struct LanguagePrivate
type LanguagePrivate struct {
	P unsafe.Pointer
}

func LanguagePrivateGetType() gi.GType {
	ret := _I.GetGType(43, "LanguagePrivate")
	return ret
}

// Object Map
type Map struct {
	atk.ImplementorIfaceIfc
	gtk.BuildableIfc
	gtk.ScrollableIfc
	View
}

func WrapMap(p unsafe.Pointer) (r Map) { r.P = p; return }

type IMap interface{ P_Map() unsafe.Pointer }

func (v Map) P_Map() unsafe.Pointer              { return v.P }
func (v Map) P_ImplementorIface() unsafe.Pointer { return v.P }
func (v Map) P_Buildable() unsafe.Pointer        { return v.P }
func (v Map) P_Scrollable() unsafe.Pointer       { return v.P }
func MapGetType() gi.GType {
	ret := _I.GetGType(44, "Map")
	return ret
}

// gtk_source_map_new
// container is not nil, container is Map
// is constructor
func NewMap() (result Map) {
	iv, err := _I.Get(191, "Map", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_map_get_view
// container is not nil, container is Map
// is method
func (v Map) GetView() (result View) {
	iv, err := _I.Get(192, "Map", "get_view")
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

// gtk_source_map_set_view
// container is not nil, container is Map
// is method
func (v Map) SetView(view IView) {
	iv, err := _I.Get(193, "Map", "set_view")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if view != nil {
		tmp = view.P_View()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_view := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_view}
	iv.Call(args, nil, nil)
}

// ignore GType struct MapClass
// Object Mark
type Mark struct {
	gtk.TextMark
}

func WrapMark(p unsafe.Pointer) (r Mark) { r.P = p; return }

type IMark interface{ P_Mark() unsafe.Pointer }

func (v Mark) P_Mark() unsafe.Pointer { return v.P }
func MarkGetType() gi.GType {
	ret := _I.GetGType(45, "Mark")
	return ret
}

// gtk_source_mark_new
// container is not nil, container is Mark
// is constructor
func NewMark(name string, category string) (result Mark) {
	iv, err := _I.Get(194, "Mark", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_category := gi.CString(category)
	arg_name := gi.NewStringArgument(c_name)
	arg_category := gi.NewStringArgument(c_category)
	args := []gi.Argument{arg_name, arg_category}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_category)
	result.P = ret.Pointer()
	return
}

// gtk_source_mark_get_category
// container is not nil, container is Mark
// is method
func (v Mark) GetCategory() (result string) {
	iv, err := _I.Get(195, "Mark", "get_category")
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

// gtk_source_mark_next
// container is not nil, container is Mark
// is method
func (v Mark) Next(category string) (result Mark) {
	iv, err := _I.Get(196, "Mark", "next")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_category := gi.CString(category)
	arg_v := gi.NewPointerArgument(v.P)
	arg_category := gi.NewStringArgument(c_category)
	args := []gi.Argument{arg_v, arg_category}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_category)
	result.P = ret.Pointer()
	return
}

// gtk_source_mark_prev
// container is not nil, container is Mark
// is method
func (v Mark) Prev(category string) (result Mark) {
	iv, err := _I.Get(197, "Mark", "prev")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_category := gi.CString(category)
	arg_v := gi.NewPointerArgument(v.P)
	arg_category := gi.NewStringArgument(c_category)
	args := []gi.Argument{arg_v, arg_category}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_category)
	result.P = ret.Pointer()
	return
}

// Object MarkAttributes
type MarkAttributes struct {
	gobject.Object
}

func WrapMarkAttributes(p unsafe.Pointer) (r MarkAttributes) { r.P = p; return }

type IMarkAttributes interface{ P_MarkAttributes() unsafe.Pointer }

func (v MarkAttributes) P_MarkAttributes() unsafe.Pointer { return v.P }
func MarkAttributesGetType() gi.GType {
	ret := _I.GetGType(46, "MarkAttributes")
	return ret
}

// gtk_source_mark_attributes_new
// container is not nil, container is MarkAttributes
// is constructor
func NewMarkAttributes() (result MarkAttributes) {
	iv, err := _I.Get(198, "MarkAttributes", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_mark_attributes_get_background
// container is not nil, container is MarkAttributes
// is method
func (v MarkAttributes) GetBackground(background gdk.RGBA) (result bool) {
	iv, err := _I.Get(199, "MarkAttributes", "get_background")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_background := gi.NewPointerArgument(background.P)
	args := []gi.Argument{arg_v, arg_background}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gtk_source_mark_attributes_get_gicon
// container is not nil, container is MarkAttributes
// is method
func (v MarkAttributes) GetGicon() (result gio.Icon) {
	iv, err := _I.Get(200, "MarkAttributes", "get_gicon")
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

// gtk_source_mark_attributes_get_icon_name
// container is not nil, container is MarkAttributes
// is method
func (v MarkAttributes) GetIconName() (result string) {
	iv, err := _I.Get(201, "MarkAttributes", "get_icon_name")
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

// gtk_source_mark_attributes_get_pixbuf
// container is not nil, container is MarkAttributes
// is method
func (v MarkAttributes) GetPixbuf() (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(202, "MarkAttributes", "get_pixbuf")
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

// gtk_source_mark_attributes_get_tooltip_markup
// container is not nil, container is MarkAttributes
// is method
func (v MarkAttributes) GetTooltipMarkup(mark IMark) (result string) {
	iv, err := _I.Get(203, "MarkAttributes", "get_tooltip_markup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if mark != nil {
		tmp = mark.P_Mark()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mark := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_mark}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// gtk_source_mark_attributes_get_tooltip_text
// container is not nil, container is MarkAttributes
// is method
func (v MarkAttributes) GetTooltipText(mark IMark) (result string) {
	iv, err := _I.Get(204, "MarkAttributes", "get_tooltip_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if mark != nil {
		tmp = mark.P_Mark()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mark := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_mark}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// gtk_source_mark_attributes_render_icon
// container is not nil, container is MarkAttributes
// is method
func (v MarkAttributes) RenderIcon(widget gtk.IWidget, size int32) (result gdkpixbuf.Pixbuf) {
	iv, err := _I.Get(205, "MarkAttributes", "render_icon")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if widget != nil {
		tmp = widget.P_Widget()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_widget := gi.NewPointerArgument(tmp)
	arg_size := gi.NewInt32Argument(size)
	args := []gi.Argument{arg_v, arg_widget, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_mark_attributes_set_background
// container is not nil, container is MarkAttributes
// is method
func (v MarkAttributes) SetBackground(background gdk.RGBA) {
	iv, err := _I.Get(206, "MarkAttributes", "set_background")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_background := gi.NewPointerArgument(background.P)
	args := []gi.Argument{arg_v, arg_background}
	iv.Call(args, nil, nil)
}

// gtk_source_mark_attributes_set_gicon
// container is not nil, container is MarkAttributes
// is method
func (v MarkAttributes) SetGicon(gicon gio.IIcon) {
	iv, err := _I.Get(207, "MarkAttributes", "set_gicon")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if gicon != nil {
		tmp = gicon.P_Icon()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_gicon := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_gicon}
	iv.Call(args, nil, nil)
}

// gtk_source_mark_attributes_set_icon_name
// container is not nil, container is MarkAttributes
// is method
func (v MarkAttributes) SetIconName(icon_name string) {
	iv, err := _I.Get(208, "MarkAttributes", "set_icon_name")
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

// gtk_source_mark_attributes_set_pixbuf
// container is not nil, container is MarkAttributes
// is method
func (v MarkAttributes) SetPixbuf(pixbuf gdkpixbuf.IPixbuf) {
	iv, err := _I.Get(209, "MarkAttributes", "set_pixbuf")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pixbuf != nil {
		tmp = pixbuf.P_Pixbuf()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pixbuf := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_pixbuf}
	iv.Call(args, nil, nil)
}

// ignore GType struct MarkAttributesClass
// Struct MarkAttributesPrivate
type MarkAttributesPrivate struct {
	P unsafe.Pointer
}

func MarkAttributesPrivateGetType() gi.GType {
	ret := _I.GetGType(47, "MarkAttributesPrivate")
	return ret
}

// ignore GType struct MarkClass
// Struct MarkPrivate
type MarkPrivate struct {
	P unsafe.Pointer
}

func MarkPrivateGetType() gi.GType {
	ret := _I.GetGType(48, "MarkPrivate")
	return ret
}

// Enum NewlineType
type NewlineTypeEnum int

const (
	NewlineTypeLf   NewlineTypeEnum = 0
	NewlineTypeCr   NewlineTypeEnum = 1
	NewlineTypeCrLf NewlineTypeEnum = 2
)

func NewlineTypeGetType() gi.GType {
	ret := _I.GetGType(49, "NewlineType")
	return ret
}

// Object PrintCompositor
type PrintCompositor struct {
	gobject.Object
}

func WrapPrintCompositor(p unsafe.Pointer) (r PrintCompositor) { r.P = p; return }

type IPrintCompositor interface{ P_PrintCompositor() unsafe.Pointer }

func (v PrintCompositor) P_PrintCompositor() unsafe.Pointer { return v.P }
func PrintCompositorGetType() gi.GType {
	ret := _I.GetGType(50, "PrintCompositor")
	return ret
}

// gtk_source_print_compositor_new
// container is not nil, container is PrintCompositor
// is constructor
func NewPrintCompositor(buffer IBuffer) (result PrintCompositor) {
	iv, err := _I.Get(210, "PrintCompositor", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if buffer != nil {
		tmp = buffer.P_Buffer()
	}
	arg_buffer := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_buffer}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_print_compositor_new_from_view
// container is not nil, container is PrintCompositor
// is constructor
func NewPrintCompositorFromView(view IView) (result PrintCompositor) {
	iv, err := _I.Get(211, "PrintCompositor", "new_from_view")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if view != nil {
		tmp = view.P_View()
	}
	arg_view := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_view}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_print_compositor_draw_page
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) DrawPage(context gtk.IPrintContext, page_nr int32) {
	iv, err := _I.Get(212, "PrintCompositor", "draw_page")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_PrintContext()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_context := gi.NewPointerArgument(tmp)
	arg_page_nr := gi.NewInt32Argument(page_nr)
	args := []gi.Argument{arg_v, arg_context, arg_page_nr}
	iv.Call(args, nil, nil)
}

// gtk_source_print_compositor_get_body_font_name
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetBodyFontName() (result string) {
	iv, err := _I.Get(213, "PrintCompositor", "get_body_font_name")
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

// gtk_source_print_compositor_get_bottom_margin
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetBottomMargin(unit gtk.UnitEnum) (result float64) {
	iv, err := _I.Get(214, "PrintCompositor", "get_bottom_margin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_unit := gi.NewIntArgument(int(unit))
	args := []gi.Argument{arg_v, arg_unit}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// gtk_source_print_compositor_get_buffer
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetBuffer() (result Buffer) {
	iv, err := _I.Get(215, "PrintCompositor", "get_buffer")
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

// gtk_source_print_compositor_get_footer_font_name
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetFooterFontName() (result string) {
	iv, err := _I.Get(216, "PrintCompositor", "get_footer_font_name")
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

// gtk_source_print_compositor_get_header_font_name
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetHeaderFontName() (result string) {
	iv, err := _I.Get(217, "PrintCompositor", "get_header_font_name")
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

// gtk_source_print_compositor_get_highlight_syntax
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetHighlightSyntax() (result bool) {
	iv, err := _I.Get(218, "PrintCompositor", "get_highlight_syntax")
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

// gtk_source_print_compositor_get_left_margin
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetLeftMargin(unit gtk.UnitEnum) (result float64) {
	iv, err := _I.Get(219, "PrintCompositor", "get_left_margin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_unit := gi.NewIntArgument(int(unit))
	args := []gi.Argument{arg_v, arg_unit}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// gtk_source_print_compositor_get_line_numbers_font_name
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetLineNumbersFontName() (result string) {
	iv, err := _I.Get(220, "PrintCompositor", "get_line_numbers_font_name")
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

// gtk_source_print_compositor_get_n_pages
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetNPages() (result int32) {
	iv, err := _I.Get(221, "PrintCompositor", "get_n_pages")
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

// gtk_source_print_compositor_get_pagination_progress
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetPaginationProgress() (result float64) {
	iv, err := _I.Get(222, "PrintCompositor", "get_pagination_progress")
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

// gtk_source_print_compositor_get_print_footer
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetPrintFooter() (result bool) {
	iv, err := _I.Get(223, "PrintCompositor", "get_print_footer")
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

// gtk_source_print_compositor_get_print_header
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetPrintHeader() (result bool) {
	iv, err := _I.Get(224, "PrintCompositor", "get_print_header")
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

// gtk_source_print_compositor_get_print_line_numbers
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetPrintLineNumbers() (result uint32) {
	iv, err := _I.Get(225, "PrintCompositor", "get_print_line_numbers")
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

// gtk_source_print_compositor_get_right_margin
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetRightMargin(unit gtk.UnitEnum) (result float64) {
	iv, err := _I.Get(226, "PrintCompositor", "get_right_margin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_unit := gi.NewIntArgument(int(unit))
	args := []gi.Argument{arg_v, arg_unit}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// gtk_source_print_compositor_get_tab_width
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetTabWidth() (result uint32) {
	iv, err := _I.Get(227, "PrintCompositor", "get_tab_width")
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

// gtk_source_print_compositor_get_top_margin
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetTopMargin(unit gtk.UnitEnum) (result float64) {
	iv, err := _I.Get(228, "PrintCompositor", "get_top_margin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_unit := gi.NewIntArgument(int(unit))
	args := []gi.Argument{arg_v, arg_unit}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// gtk_source_print_compositor_get_wrap_mode
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) GetWrapMode() (result gtk.WrapModeEnum) {
	iv, err := _I.Get(229, "PrintCompositor", "get_wrap_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gtk.WrapModeEnum(ret.Int())
	return
}

// gtk_source_print_compositor_paginate
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) Paginate(context gtk.IPrintContext) (result bool) {
	iv, err := _I.Get(230, "PrintCompositor", "paginate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if context != nil {
		tmp = context.P_PrintContext()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_context := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_context}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gtk_source_print_compositor_set_body_font_name
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetBodyFontName(font_name string) {
	iv, err := _I.Get(231, "PrintCompositor", "set_body_font_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_font_name := gi.CString(font_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_font_name := gi.NewStringArgument(c_font_name)
	args := []gi.Argument{arg_v, arg_font_name}
	iv.Call(args, nil, nil)
	gi.Free(c_font_name)
}

// gtk_source_print_compositor_set_bottom_margin
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetBottomMargin(margin float64, unit gtk.UnitEnum) {
	iv, err := _I.Get(232, "PrintCompositor", "set_bottom_margin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_margin := gi.NewDoubleArgument(margin)
	arg_unit := gi.NewIntArgument(int(unit))
	args := []gi.Argument{arg_v, arg_margin, arg_unit}
	iv.Call(args, nil, nil)
}

// gtk_source_print_compositor_set_footer_font_name
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetFooterFontName(font_name string) {
	iv, err := _I.Get(233, "PrintCompositor", "set_footer_font_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_font_name := gi.CString(font_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_font_name := gi.NewStringArgument(c_font_name)
	args := []gi.Argument{arg_v, arg_font_name}
	iv.Call(args, nil, nil)
	gi.Free(c_font_name)
}

// gtk_source_print_compositor_set_footer_format
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetFooterFormat(separator bool, left string, center string, right string) {
	iv, err := _I.Get(234, "PrintCompositor", "set_footer_format")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_left := gi.CString(left)
	c_center := gi.CString(center)
	c_right := gi.CString(right)
	arg_v := gi.NewPointerArgument(v.P)
	arg_separator := gi.NewBoolArgument(separator)
	arg_left := gi.NewStringArgument(c_left)
	arg_center := gi.NewStringArgument(c_center)
	arg_right := gi.NewStringArgument(c_right)
	args := []gi.Argument{arg_v, arg_separator, arg_left, arg_center, arg_right}
	iv.Call(args, nil, nil)
	gi.Free(c_left)
	gi.Free(c_center)
	gi.Free(c_right)
}

// gtk_source_print_compositor_set_header_font_name
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetHeaderFontName(font_name string) {
	iv, err := _I.Get(235, "PrintCompositor", "set_header_font_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_font_name := gi.CString(font_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_font_name := gi.NewStringArgument(c_font_name)
	args := []gi.Argument{arg_v, arg_font_name}
	iv.Call(args, nil, nil)
	gi.Free(c_font_name)
}

// gtk_source_print_compositor_set_header_format
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetHeaderFormat(separator bool, left string, center string, right string) {
	iv, err := _I.Get(236, "PrintCompositor", "set_header_format")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_left := gi.CString(left)
	c_center := gi.CString(center)
	c_right := gi.CString(right)
	arg_v := gi.NewPointerArgument(v.P)
	arg_separator := gi.NewBoolArgument(separator)
	arg_left := gi.NewStringArgument(c_left)
	arg_center := gi.NewStringArgument(c_center)
	arg_right := gi.NewStringArgument(c_right)
	args := []gi.Argument{arg_v, arg_separator, arg_left, arg_center, arg_right}
	iv.Call(args, nil, nil)
	gi.Free(c_left)
	gi.Free(c_center)
	gi.Free(c_right)
}

// gtk_source_print_compositor_set_highlight_syntax
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetHighlightSyntax(highlight bool) {
	iv, err := _I.Get(237, "PrintCompositor", "set_highlight_syntax")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_highlight := gi.NewBoolArgument(highlight)
	args := []gi.Argument{arg_v, arg_highlight}
	iv.Call(args, nil, nil)
}

// gtk_source_print_compositor_set_left_margin
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetLeftMargin(margin float64, unit gtk.UnitEnum) {
	iv, err := _I.Get(238, "PrintCompositor", "set_left_margin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_margin := gi.NewDoubleArgument(margin)
	arg_unit := gi.NewIntArgument(int(unit))
	args := []gi.Argument{arg_v, arg_margin, arg_unit}
	iv.Call(args, nil, nil)
}

// gtk_source_print_compositor_set_line_numbers_font_name
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetLineNumbersFontName(font_name string) {
	iv, err := _I.Get(239, "PrintCompositor", "set_line_numbers_font_name")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_font_name := gi.CString(font_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_font_name := gi.NewStringArgument(c_font_name)
	args := []gi.Argument{arg_v, arg_font_name}
	iv.Call(args, nil, nil)
	gi.Free(c_font_name)
}

// gtk_source_print_compositor_set_print_footer
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetPrintFooter(print1 bool) {
	iv, err := _I.Get(240, "PrintCompositor", "set_print_footer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_print1 := gi.NewBoolArgument(print1)
	args := []gi.Argument{arg_v, arg_print1}
	iv.Call(args, nil, nil)
}

// gtk_source_print_compositor_set_print_header
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetPrintHeader(print1 bool) {
	iv, err := _I.Get(241, "PrintCompositor", "set_print_header")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_print1 := gi.NewBoolArgument(print1)
	args := []gi.Argument{arg_v, arg_print1}
	iv.Call(args, nil, nil)
}

// gtk_source_print_compositor_set_print_line_numbers
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetPrintLineNumbers(interval uint32) {
	iv, err := _I.Get(242, "PrintCompositor", "set_print_line_numbers")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interval := gi.NewUint32Argument(interval)
	args := []gi.Argument{arg_v, arg_interval}
	iv.Call(args, nil, nil)
}

// gtk_source_print_compositor_set_right_margin
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetRightMargin(margin float64, unit gtk.UnitEnum) {
	iv, err := _I.Get(243, "PrintCompositor", "set_right_margin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_margin := gi.NewDoubleArgument(margin)
	arg_unit := gi.NewIntArgument(int(unit))
	args := []gi.Argument{arg_v, arg_margin, arg_unit}
	iv.Call(args, nil, nil)
}

// gtk_source_print_compositor_set_tab_width
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetTabWidth(width uint32) {
	iv, err := _I.Get(244, "PrintCompositor", "set_tab_width")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_width := gi.NewUint32Argument(width)
	args := []gi.Argument{arg_v, arg_width}
	iv.Call(args, nil, nil)
}

// gtk_source_print_compositor_set_top_margin
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetTopMargin(margin float64, unit gtk.UnitEnum) {
	iv, err := _I.Get(245, "PrintCompositor", "set_top_margin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_margin := gi.NewDoubleArgument(margin)
	arg_unit := gi.NewIntArgument(int(unit))
	args := []gi.Argument{arg_v, arg_margin, arg_unit}
	iv.Call(args, nil, nil)
}

// gtk_source_print_compositor_set_wrap_mode
// container is not nil, container is PrintCompositor
// is method
func (v PrintCompositor) SetWrapMode(wrap_mode gtk.WrapModeEnum) {
	iv, err := _I.Get(246, "PrintCompositor", "set_wrap_mode")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_wrap_mode := gi.NewIntArgument(int(wrap_mode))
	args := []gi.Argument{arg_v, arg_wrap_mode}
	iv.Call(args, nil, nil)
}

// ignore GType struct PrintCompositorClass
// Struct PrintCompositorPrivate
type PrintCompositorPrivate struct {
	P unsafe.Pointer
}

func PrintCompositorPrivateGetType() gi.GType {
	ret := _I.GetGType(51, "PrintCompositorPrivate")
	return ret
}

// Object Region
type Region struct {
	gobject.Object
}

func WrapRegion(p unsafe.Pointer) (r Region) { r.P = p; return }

type IRegion interface{ P_Region() unsafe.Pointer }

func (v Region) P_Region() unsafe.Pointer { return v.P }
func RegionGetType() gi.GType {
	ret := _I.GetGType(52, "Region")
	return ret
}

// gtk_source_region_new
// container is not nil, container is Region
// is constructor
func NewRegion(buffer gtk.ITextBuffer) (result Region) {
	iv, err := _I.Get(247, "Region", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if buffer != nil {
		tmp = buffer.P_TextBuffer()
	}
	arg_buffer := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_buffer}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_region_add_region
// container is not nil, container is Region
// is method
func (v Region) AddRegion(region_to_add IRegion) {
	iv, err := _I.Get(248, "Region", "add_region")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if region_to_add != nil {
		tmp = region_to_add.P_Region()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_region_to_add := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_region_to_add}
	iv.Call(args, nil, nil)
}

// gtk_source_region_add_subregion
// container is not nil, container is Region
// is method
func (v Region) AddSubregion(_start gtk.TextIter, _end gtk.TextIter) {
	iv, err := _I.Get(249, "Region", "add_subregion")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg__start := gi.NewPointerArgument(_start.P)
	arg__end := gi.NewPointerArgument(_end.P)
	args := []gi.Argument{arg_v, arg__start, arg__end}
	iv.Call(args, nil, nil)
}

// gtk_source_region_get_bounds
// container is not nil, container is Region
// is method
func (v Region) GetBounds(start gtk.TextIter, end gtk.TextIter) (result bool) {
	iv, err := _I.Get(250, "Region", "get_bounds")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewPointerArgument(start.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_v, arg_start, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gtk_source_region_get_buffer
// container is not nil, container is Region
// is method
func (v Region) GetBuffer() (result gtk.TextBuffer) {
	iv, err := _I.Get(251, "Region", "get_buffer")
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

// gtk_source_region_get_start_region_iter
// container is not nil, container is Region
// is method
func (v Region) GetStartRegionIter(iter RegionIter) {
	iv, err := _I.Get(252, "Region", "get_start_region_iter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	args := []gi.Argument{arg_v, arg_iter}
	iv.Call(args, nil, nil)
}

// gtk_source_region_intersect_region
// container is not nil, container is Region
// is method
func (v Region) IntersectRegion(region2 IRegion) (result Region) {
	iv, err := _I.Get(253, "Region", "intersect_region")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if region2 != nil {
		tmp = region2.P_Region()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_region2 := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_region2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_region_intersect_subregion
// container is not nil, container is Region
// is method
func (v Region) IntersectSubregion(_start gtk.TextIter, _end gtk.TextIter) (result Region) {
	iv, err := _I.Get(254, "Region", "intersect_subregion")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg__start := gi.NewPointerArgument(_start.P)
	arg__end := gi.NewPointerArgument(_end.P)
	args := []gi.Argument{arg_v, arg__start, arg__end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_region_is_empty
// container is not nil, container is Region
// is method
func (v Region) IsEmpty() (result bool) {
	iv, err := _I.Get(255, "Region", "is_empty")
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

// gtk_source_region_subtract_region
// container is not nil, container is Region
// is method
func (v Region) SubtractRegion(region_to_subtract IRegion) {
	iv, err := _I.Get(256, "Region", "subtract_region")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if region_to_subtract != nil {
		tmp = region_to_subtract.P_Region()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_region_to_subtract := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_region_to_subtract}
	iv.Call(args, nil, nil)
}

// gtk_source_region_subtract_subregion
// container is not nil, container is Region
// is method
func (v Region) SubtractSubregion(_start gtk.TextIter, _end gtk.TextIter) {
	iv, err := _I.Get(257, "Region", "subtract_subregion")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg__start := gi.NewPointerArgument(_start.P)
	arg__end := gi.NewPointerArgument(_end.P)
	args := []gi.Argument{arg_v, arg__start, arg__end}
	iv.Call(args, nil, nil)
}

// gtk_source_region_to_string
// container is not nil, container is Region
// is method
func (v Region) ToString() (result string) {
	iv, err := _I.Get(258, "Region", "to_string")
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

// ignore GType struct RegionClass
// Struct RegionIter
type RegionIter struct {
	P unsafe.Pointer
}

const SizeOfStructRegionIter = 24

func RegionIterGetType() gi.GType {
	ret := _I.GetGType(53, "RegionIter")
	return ret
}

// gtk_source_region_iter_get_subregion
// container is not nil, container is RegionIter
// is method
func (v RegionIter) GetSubregion(start gtk.TextIter, end gtk.TextIter) (result bool) {
	iv, err := _I.Get(259, "RegionIter", "get_subregion")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewPointerArgument(start.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_v, arg_start, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// gtk_source_region_iter_is_end
// container is not nil, container is RegionIter
// is method
func (v RegionIter) IsEnd() (result bool) {
	iv, err := _I.Get(260, "RegionIter", "is_end")
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

// gtk_source_region_iter_next
// container is not nil, container is RegionIter
// is method
func (v RegionIter) Next() (result bool) {
	iv, err := _I.Get(261, "RegionIter", "next")
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

// Object SearchContext
type SearchContext struct {
	gobject.Object
}

func WrapSearchContext(p unsafe.Pointer) (r SearchContext) { r.P = p; return }

type ISearchContext interface{ P_SearchContext() unsafe.Pointer }

func (v SearchContext) P_SearchContext() unsafe.Pointer { return v.P }
func SearchContextGetType() gi.GType {
	ret := _I.GetGType(54, "SearchContext")
	return ret
}

// gtk_source_search_context_new
// container is not nil, container is SearchContext
// is constructor
func NewSearchContext(buffer IBuffer, settings ISearchSettings) (result SearchContext) {
	iv, err := _I.Get(262, "SearchContext", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if buffer != nil {
		tmp = buffer.P_Buffer()
	}
	var tmp1 unsafe.Pointer
	if settings != nil {
		tmp1 = settings.P_SearchSettings()
	}
	arg_buffer := gi.NewPointerArgument(tmp)
	arg_settings := gi.NewPointerArgument(tmp1)
	args := []gi.Argument{arg_buffer, arg_settings}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_search_context_backward
// container is not nil, container is SearchContext
// is method
func (v SearchContext) Backward(iter gtk.TextIter, match_start gtk.TextIter, match_end gtk.TextIter) (result bool, has_wrapped_around bool) {
	iv, err := _I.Get(263, "SearchContext", "backward")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_match_start := gi.NewPointerArgument(match_start.P)
	arg_match_end := gi.NewPointerArgument(match_end.P)
	arg_has_wrapped_around := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_iter, arg_match_start, arg_match_end, arg_has_wrapped_around}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	has_wrapped_around = outArgs[0].Bool()
	result = ret.Bool()
	return
}

// gtk_source_search_context_backward_async
// container is not nil, container is SearchContext
// is method
func (v SearchContext) BackwardAsync(iter gtk.TextIter, cancellable gio.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(264, "SearchContext", "backward_async")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(gio.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_iter, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// gtk_source_search_context_backward_finish
// container is not nil, container is SearchContext
// is method
func (v SearchContext) BackwardFinish(result gio.IAsyncResult, match_start gtk.TextIter, match_end gtk.TextIter) (result1 bool, has_wrapped_around bool, err error) {
	iv, err := _I.Get(265, "SearchContext", "backward_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if result != nil {
		tmp = result.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_result := gi.NewPointerArgument(tmp)
	arg_match_start := gi.NewPointerArgument(match_start.P)
	arg_match_end := gi.NewPointerArgument(match_end.P)
	arg_has_wrapped_around := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_result, arg_match_start, arg_match_end, arg_has_wrapped_around, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	has_wrapped_around = outArgs[0].Bool()
	result1 = ret.Bool()
	return
}

// gtk_source_search_context_forward
// container is not nil, container is SearchContext
// is method
func (v SearchContext) Forward(iter gtk.TextIter, match_start gtk.TextIter, match_end gtk.TextIter) (result bool, has_wrapped_around bool) {
	iv, err := _I.Get(266, "SearchContext", "forward")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_match_start := gi.NewPointerArgument(match_start.P)
	arg_match_end := gi.NewPointerArgument(match_end.P)
	arg_has_wrapped_around := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_iter, arg_match_start, arg_match_end, arg_has_wrapped_around}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	has_wrapped_around = outArgs[0].Bool()
	result = ret.Bool()
	return
}

// gtk_source_search_context_forward_async
// container is not nil, container is SearchContext
// is method
func (v SearchContext) ForwardAsync(iter gtk.TextIter, cancellable gio.ICancellable, callback int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(267, "SearchContext", "forward_async")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if cancellable != nil {
		tmp = cancellable.P_Cancellable()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_cancellable := gi.NewPointerArgument(tmp)
	arg_callback := gi.NewPointerArgument(unsafe.Pointer(gio.GetPointer_myAsyncReadyCallback()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_iter, arg_cancellable, arg_callback, arg_user_data}
	iv.Call(args, nil, nil)
}

// gtk_source_search_context_forward_finish
// container is not nil, container is SearchContext
// is method
func (v SearchContext) ForwardFinish(result gio.IAsyncResult, match_start gtk.TextIter, match_end gtk.TextIter) (result1 bool, has_wrapped_around bool, err error) {
	iv, err := _I.Get(268, "SearchContext", "forward_finish")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	var tmp unsafe.Pointer
	if result != nil {
		tmp = result.P_AsyncResult()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_result := gi.NewPointerArgument(tmp)
	arg_match_start := gi.NewPointerArgument(match_start.P)
	arg_match_end := gi.NewPointerArgument(match_end.P)
	arg_has_wrapped_around := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_result, arg_match_start, arg_match_end, arg_has_wrapped_around, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	has_wrapped_around = outArgs[0].Bool()
	result1 = ret.Bool()
	return
}

// gtk_source_search_context_get_buffer
// container is not nil, container is SearchContext
// is method
func (v SearchContext) GetBuffer() (result Buffer) {
	iv, err := _I.Get(269, "SearchContext", "get_buffer")
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

// gtk_source_search_context_get_highlight
// container is not nil, container is SearchContext
// is method
func (v SearchContext) GetHighlight() (result bool) {
	iv, err := _I.Get(270, "SearchContext", "get_highlight")
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

// gtk_source_search_context_get_match_style
// container is not nil, container is SearchContext
// is method
func (v SearchContext) GetMatchStyle() (result Style) {
	iv, err := _I.Get(271, "SearchContext", "get_match_style")
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

// gtk_source_search_context_get_occurrence_position
// container is not nil, container is SearchContext
// is method
func (v SearchContext) GetOccurrencePosition(match_start gtk.TextIter, match_end gtk.TextIter) (result int32) {
	iv, err := _I.Get(272, "SearchContext", "get_occurrence_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_match_start := gi.NewPointerArgument(match_start.P)
	arg_match_end := gi.NewPointerArgument(match_end.P)
	args := []gi.Argument{arg_v, arg_match_start, arg_match_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// gtk_source_search_context_get_occurrences_count
// container is not nil, container is SearchContext
// is method
func (v SearchContext) GetOccurrencesCount() (result int32) {
	iv, err := _I.Get(273, "SearchContext", "get_occurrences_count")
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

// gtk_source_search_context_get_regex_error
// container is not nil, container is SearchContext
// is method
func (v SearchContext) GetRegexError() (result glib.Error) {
	iv, err := _I.Get(274, "SearchContext", "get_regex_error")
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

// gtk_source_search_context_get_settings
// container is not nil, container is SearchContext
// is method
func (v SearchContext) GetSettings() (result SearchSettings) {
	iv, err := _I.Get(275, "SearchContext", "get_settings")
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

// gtk_source_search_context_replace
// container is not nil, container is SearchContext
// is method
func (v SearchContext) Replace(match_start gtk.TextIter, match_end gtk.TextIter, replace string, replace_length int32) (result bool, err error) {
	iv, err := _I.Get(276, "SearchContext", "replace")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_replace := gi.CString(replace)
	arg_v := gi.NewPointerArgument(v.P)
	arg_match_start := gi.NewPointerArgument(match_start.P)
	arg_match_end := gi.NewPointerArgument(match_end.P)
	arg_replace := gi.NewStringArgument(c_replace)
	arg_replace_length := gi.NewInt32Argument(replace_length)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_match_start, arg_match_end, arg_replace, arg_replace_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_replace)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// gtk_source_search_context_replace_all
// container is not nil, container is SearchContext
// is method
func (v SearchContext) ReplaceAll(replace string, replace_length int32) (result uint32, err error) {
	iv, err := _I.Get(277, "SearchContext", "replace_all")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_replace := gi.CString(replace)
	arg_v := gi.NewPointerArgument(v.P)
	arg_replace := gi.NewStringArgument(c_replace)
	arg_replace_length := gi.NewInt32Argument(replace_length)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_replace, arg_replace_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_replace)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Uint32()
	return
}

// gtk_source_search_context_set_highlight
// container is not nil, container is SearchContext
// is method
func (v SearchContext) SetHighlight(highlight bool) {
	iv, err := _I.Get(278, "SearchContext", "set_highlight")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_highlight := gi.NewBoolArgument(highlight)
	args := []gi.Argument{arg_v, arg_highlight}
	iv.Call(args, nil, nil)
}

// gtk_source_search_context_set_match_style
// container is not nil, container is SearchContext
// is method
func (v SearchContext) SetMatchStyle(match_style IStyle) {
	iv, err := _I.Get(279, "SearchContext", "set_match_style")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if match_style != nil {
		tmp = match_style.P_Style()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_match_style := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_match_style}
	iv.Call(args, nil, nil)
}

// ignore GType struct SearchContextClass
// Struct SearchContextPrivate
type SearchContextPrivate struct {
	P unsafe.Pointer
}

func SearchContextPrivateGetType() gi.GType {
	ret := _I.GetGType(55, "SearchContextPrivate")
	return ret
}

// Object SearchSettings
type SearchSettings struct {
	gobject.Object
}

func WrapSearchSettings(p unsafe.Pointer) (r SearchSettings) { r.P = p; return }

type ISearchSettings interface{ P_SearchSettings() unsafe.Pointer }

func (v SearchSettings) P_SearchSettings() unsafe.Pointer { return v.P }
func SearchSettingsGetType() gi.GType {
	ret := _I.GetGType(56, "SearchSettings")
	return ret
}

// gtk_source_search_settings_new
// container is not nil, container is SearchSettings
// is constructor
func NewSearchSettings() (result SearchSettings) {
	iv, err := _I.Get(280, "SearchSettings", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_search_settings_get_at_word_boundaries
// container is not nil, container is SearchSettings
// is method
func (v SearchSettings) GetAtWordBoundaries() (result bool) {
	iv, err := _I.Get(281, "SearchSettings", "get_at_word_boundaries")
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

// gtk_source_search_settings_get_case_sensitive
// container is not nil, container is SearchSettings
// is method
func (v SearchSettings) GetCaseSensitive() (result bool) {
	iv, err := _I.Get(282, "SearchSettings", "get_case_sensitive")
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

// gtk_source_search_settings_get_regex_enabled
// container is not nil, container is SearchSettings
// is method
func (v SearchSettings) GetRegexEnabled() (result bool) {
	iv, err := _I.Get(283, "SearchSettings", "get_regex_enabled")
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

// gtk_source_search_settings_get_search_text
// container is not nil, container is SearchSettings
// is method
func (v SearchSettings) GetSearchText() (result string) {
	iv, err := _I.Get(284, "SearchSettings", "get_search_text")
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

// gtk_source_search_settings_get_wrap_around
// container is not nil, container is SearchSettings
// is method
func (v SearchSettings) GetWrapAround() (result bool) {
	iv, err := _I.Get(285, "SearchSettings", "get_wrap_around")
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

// gtk_source_search_settings_set_at_word_boundaries
// container is not nil, container is SearchSettings
// is method
func (v SearchSettings) SetAtWordBoundaries(at_word_boundaries bool) {
	iv, err := _I.Get(286, "SearchSettings", "set_at_word_boundaries")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_at_word_boundaries := gi.NewBoolArgument(at_word_boundaries)
	args := []gi.Argument{arg_v, arg_at_word_boundaries}
	iv.Call(args, nil, nil)
}

// gtk_source_search_settings_set_case_sensitive
// container is not nil, container is SearchSettings
// is method
func (v SearchSettings) SetCaseSensitive(case_sensitive bool) {
	iv, err := _I.Get(287, "SearchSettings", "set_case_sensitive")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_case_sensitive := gi.NewBoolArgument(case_sensitive)
	args := []gi.Argument{arg_v, arg_case_sensitive}
	iv.Call(args, nil, nil)
}

// gtk_source_search_settings_set_regex_enabled
// container is not nil, container is SearchSettings
// is method
func (v SearchSettings) SetRegexEnabled(regex_enabled bool) {
	iv, err := _I.Get(288, "SearchSettings", "set_regex_enabled")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_regex_enabled := gi.NewBoolArgument(regex_enabled)
	args := []gi.Argument{arg_v, arg_regex_enabled}
	iv.Call(args, nil, nil)
}

// gtk_source_search_settings_set_search_text
// container is not nil, container is SearchSettings
// is method
func (v SearchSettings) SetSearchText(search_text string) {
	iv, err := _I.Get(289, "SearchSettings", "set_search_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_search_text := gi.CString(search_text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_search_text := gi.NewStringArgument(c_search_text)
	args := []gi.Argument{arg_v, arg_search_text}
	iv.Call(args, nil, nil)
	gi.Free(c_search_text)
}

// gtk_source_search_settings_set_wrap_around
// container is not nil, container is SearchSettings
// is method
func (v SearchSettings) SetWrapAround(wrap_around bool) {
	iv, err := _I.Get(290, "SearchSettings", "set_wrap_around")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_wrap_around := gi.NewBoolArgument(wrap_around)
	args := []gi.Argument{arg_v, arg_wrap_around}
	iv.Call(args, nil, nil)
}

// ignore GType struct SearchSettingsClass
// Struct SearchSettingsPrivate
type SearchSettingsPrivate struct {
	P unsafe.Pointer
}

func SearchSettingsPrivateGetType() gi.GType {
	ret := _I.GetGType(57, "SearchSettingsPrivate")
	return ret
}

// Enum SmartHomeEndType
type SmartHomeEndTypeEnum int

const (
	SmartHomeEndTypeDisabled SmartHomeEndTypeEnum = 0
	SmartHomeEndTypeBefore   SmartHomeEndTypeEnum = 1
	SmartHomeEndTypeAfter    SmartHomeEndTypeEnum = 2
	SmartHomeEndTypeAlways   SmartHomeEndTypeEnum = 3
)

func SmartHomeEndTypeGetType() gi.GType {
	ret := _I.GetGType(58, "SmartHomeEndType")
	return ret
}

// Flags SortFlags
type SortFlags int

const (
	SortFlagsNone             SortFlags = 0
	SortFlagsCaseSensitive    SortFlags = 1
	SortFlagsReverseOrder     SortFlags = 2
	SortFlagsRemoveDuplicates SortFlags = 4
)

func SortFlagsGetType() gi.GType {
	ret := _I.GetGType(59, "SortFlags")
	return ret
}

// Object SpaceDrawer
type SpaceDrawer struct {
	gobject.Object
}

func WrapSpaceDrawer(p unsafe.Pointer) (r SpaceDrawer) { r.P = p; return }

type ISpaceDrawer interface{ P_SpaceDrawer() unsafe.Pointer }

func (v SpaceDrawer) P_SpaceDrawer() unsafe.Pointer { return v.P }
func SpaceDrawerGetType() gi.GType {
	ret := _I.GetGType(60, "SpaceDrawer")
	return ret
}

// gtk_source_space_drawer_new
// container is not nil, container is SpaceDrawer
// is constructor
func NewSpaceDrawer() (result SpaceDrawer) {
	iv, err := _I.Get(291, "SpaceDrawer", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_space_drawer_bind_matrix_setting
// container is not nil, container is SpaceDrawer
// is method
func (v SpaceDrawer) BindMatrixSetting(settings gio.ISettings, key string, flags gio.SettingsBindFlags) {
	iv, err := _I.Get(292, "SpaceDrawer", "bind_matrix_setting")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if settings != nil {
		tmp = settings.P_Settings()
	}
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_settings := gi.NewPointerArgument(tmp)
	arg_key := gi.NewStringArgument(c_key)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_settings, arg_key, arg_flags}
	iv.Call(args, nil, nil)
	gi.Free(c_key)
}

// gtk_source_space_drawer_get_enable_matrix
// container is not nil, container is SpaceDrawer
// is method
func (v SpaceDrawer) GetEnableMatrix() (result bool) {
	iv, err := _I.Get(293, "SpaceDrawer", "get_enable_matrix")
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

// gtk_source_space_drawer_get_matrix
// container is not nil, container is SpaceDrawer
// is method
func (v SpaceDrawer) GetMatrix() (result glib.Variant) {
	iv, err := _I.Get(294, "SpaceDrawer", "get_matrix")
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

// gtk_source_space_drawer_get_types_for_locations
// container is not nil, container is SpaceDrawer
// is method
func (v SpaceDrawer) GetTypesForLocations(locations SpaceLocationFlags) (result SpaceTypeFlags) {
	iv, err := _I.Get(295, "SpaceDrawer", "get_types_for_locations")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_locations := gi.NewIntArgument(int(locations))
	args := []gi.Argument{arg_v, arg_locations}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = SpaceTypeFlags(ret.Int())
	return
}

// gtk_source_space_drawer_set_enable_matrix
// container is not nil, container is SpaceDrawer
// is method
func (v SpaceDrawer) SetEnableMatrix(enable_matrix bool) {
	iv, err := _I.Get(296, "SpaceDrawer", "set_enable_matrix")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_enable_matrix := gi.NewBoolArgument(enable_matrix)
	args := []gi.Argument{arg_v, arg_enable_matrix}
	iv.Call(args, nil, nil)
}

// gtk_source_space_drawer_set_matrix
// container is not nil, container is SpaceDrawer
// is method
func (v SpaceDrawer) SetMatrix(matrix glib.Variant) {
	iv, err := _I.Get(297, "SpaceDrawer", "set_matrix")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_matrix := gi.NewPointerArgument(matrix.P)
	args := []gi.Argument{arg_v, arg_matrix}
	iv.Call(args, nil, nil)
}

// gtk_source_space_drawer_set_types_for_locations
// container is not nil, container is SpaceDrawer
// is method
func (v SpaceDrawer) SetTypesForLocations(locations SpaceLocationFlags, types SpaceTypeFlags) {
	iv, err := _I.Get(298, "SpaceDrawer", "set_types_for_locations")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_locations := gi.NewIntArgument(int(locations))
	arg_types := gi.NewIntArgument(int(types))
	args := []gi.Argument{arg_v, arg_locations, arg_types}
	iv.Call(args, nil, nil)
}

// ignore GType struct SpaceDrawerClass
// Struct SpaceDrawerPrivate
type SpaceDrawerPrivate struct {
	P unsafe.Pointer
}

func SpaceDrawerPrivateGetType() gi.GType {
	ret := _I.GetGType(61, "SpaceDrawerPrivate")
	return ret
}

// Flags SpaceLocationFlags
type SpaceLocationFlags int

const (
	SpaceLocationFlagsNone       SpaceLocationFlags = 0
	SpaceLocationFlagsLeading    SpaceLocationFlags = 1
	SpaceLocationFlagsInsideText SpaceLocationFlags = 2
	SpaceLocationFlagsTrailing   SpaceLocationFlags = 4
	SpaceLocationFlagsAll        SpaceLocationFlags = 7
)

func SpaceLocationFlagsGetType() gi.GType {
	ret := _I.GetGType(62, "SpaceLocationFlags")
	return ret
}

// Flags SpaceTypeFlags
type SpaceTypeFlags int

const (
	SpaceTypeFlagsNone    SpaceTypeFlags = 0
	SpaceTypeFlagsSpace   SpaceTypeFlags = 1
	SpaceTypeFlagsTab     SpaceTypeFlags = 2
	SpaceTypeFlagsNewline SpaceTypeFlags = 4
	SpaceTypeFlagsNbsp    SpaceTypeFlags = 8
	SpaceTypeFlagsAll     SpaceTypeFlags = 15
)

func SpaceTypeFlagsGetType() gi.GType {
	ret := _I.GetGType(63, "SpaceTypeFlags")
	return ret
}

// Object Style
type Style struct {
	gobject.Object
}

func WrapStyle(p unsafe.Pointer) (r Style) { r.P = p; return }

type IStyle interface{ P_Style() unsafe.Pointer }

func (v Style) P_Style() unsafe.Pointer { return v.P }
func StyleGetType() gi.GType {
	ret := _I.GetGType(64, "Style")
	return ret
}

// gtk_source_style_apply
// container is not nil, container is Style
// is method
func (v Style) Apply(tag gtk.ITextTag) {
	iv, err := _I.Get(299, "Style", "apply")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if tag != nil {
		tmp = tag.P_TextTag()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_tag}
	iv.Call(args, nil, nil)
}

// gtk_source_style_copy
// container is not nil, container is Style
// is method
func (v Style) Copy() (result Style) {
	iv, err := _I.Get(300, "Style", "copy")
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

// ignore GType struct StyleClass
// Object StyleScheme
type StyleScheme struct {
	gobject.Object
}

func WrapStyleScheme(p unsafe.Pointer) (r StyleScheme) { r.P = p; return }

type IStyleScheme interface{ P_StyleScheme() unsafe.Pointer }

func (v StyleScheme) P_StyleScheme() unsafe.Pointer { return v.P }
func StyleSchemeGetType() gi.GType {
	ret := _I.GetGType(65, "StyleScheme")
	return ret
}

// gtk_source_style_scheme_get_authors
// container is not nil, container is StyleScheme
// is method
func (v StyleScheme) GetAuthors() (result gi.CStrArray) {
	iv, err := _I.Get(301, "StyleScheme", "get_authors")
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

// gtk_source_style_scheme_get_description
// container is not nil, container is StyleScheme
// is method
func (v StyleScheme) GetDescription() (result string) {
	iv, err := _I.Get(302, "StyleScheme", "get_description")
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

// gtk_source_style_scheme_get_filename
// container is not nil, container is StyleScheme
// is method
func (v StyleScheme) GetFilename() (result string) {
	iv, err := _I.Get(303, "StyleScheme", "get_filename")
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

// gtk_source_style_scheme_get_id
// container is not nil, container is StyleScheme
// is method
func (v StyleScheme) GetId() (result string) {
	iv, err := _I.Get(304, "StyleScheme", "get_id")
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

// gtk_source_style_scheme_get_name
// container is not nil, container is StyleScheme
// is method
func (v StyleScheme) GetName() (result string) {
	iv, err := _I.Get(305, "StyleScheme", "get_name")
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

// gtk_source_style_scheme_get_style
// container is not nil, container is StyleScheme
// is method
func (v StyleScheme) GetStyle(style_id string) (result Style) {
	iv, err := _I.Get(306, "StyleScheme", "get_style")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_style_id := gi.CString(style_id)
	arg_v := gi.NewPointerArgument(v.P)
	arg_style_id := gi.NewStringArgument(c_style_id)
	args := []gi.Argument{arg_v, arg_style_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_style_id)
	result.P = ret.Pointer()
	return
}

// Interface StyleSchemeChooser
type StyleSchemeChooser struct {
	StyleSchemeChooserIfc
	P unsafe.Pointer
}
type StyleSchemeChooserIfc struct{}
type IStyleSchemeChooser interface{ P_StyleSchemeChooser() unsafe.Pointer }

func (v StyleSchemeChooser) P_StyleSchemeChooser() unsafe.Pointer { return v.P }
func StyleSchemeChooserGetType() gi.GType {
	ret := _I.GetGType(66, "StyleSchemeChooser")
	return ret
}

// gtk_source_style_scheme_chooser_get_style_scheme
// container is not nil, container is StyleSchemeChooser
// is method
func (v *StyleSchemeChooserIfc) GetStyleScheme() (result StyleScheme) {
	iv, err := _I.Get(307, "StyleSchemeChooser", "get_style_scheme")
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

// gtk_source_style_scheme_chooser_set_style_scheme
// container is not nil, container is StyleSchemeChooser
// is method
func (v *StyleSchemeChooserIfc) SetStyleScheme(scheme IStyleScheme) {
	iv, err := _I.Get(308, "StyleSchemeChooser", "set_style_scheme")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if scheme != nil {
		tmp = scheme.P_StyleScheme()
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_scheme := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_scheme}
	iv.Call(args, nil, nil)
}

// Object StyleSchemeChooserButton
type StyleSchemeChooserButton struct {
	atk.ImplementorIfaceIfc
	gtk.ActionableIfc
	gtk.ActivatableIfc
	gtk.BuildableIfc
	StyleSchemeChooserIfc
	gtk.Button
}

func WrapStyleSchemeChooserButton(p unsafe.Pointer) (r StyleSchemeChooserButton) { r.P = p; return }

type IStyleSchemeChooserButton interface{ P_StyleSchemeChooserButton() unsafe.Pointer }

func (v StyleSchemeChooserButton) P_StyleSchemeChooserButton() unsafe.Pointer { return v.P }
func (v StyleSchemeChooserButton) P_ImplementorIface() unsafe.Pointer         { return v.P }
func (v StyleSchemeChooserButton) P_Actionable() unsafe.Pointer               { return v.P }
func (v StyleSchemeChooserButton) P_Activatable() unsafe.Pointer              { return v.P }
func (v StyleSchemeChooserButton) P_Buildable() unsafe.Pointer                { return v.P }
func (v StyleSchemeChooserButton) P_StyleSchemeChooser() unsafe.Pointer       { return v.P }
func StyleSchemeChooserButtonGetType() gi.GType {
	ret := _I.GetGType(67, "StyleSchemeChooserButton")
	return ret
}

// gtk_source_style_scheme_chooser_button_new
// container is not nil, container is StyleSchemeChooserButton
// is constructor
func NewStyleSchemeChooserButton() (result StyleSchemeChooserButton) {
	iv, err := _I.Get(309, "StyleSchemeChooserButton", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct StyleSchemeChooserButtonClass
// ignore GType struct StyleSchemeChooserInterface
// Object StyleSchemeChooserWidget
type StyleSchemeChooserWidget struct {
	atk.ImplementorIfaceIfc
	gtk.BuildableIfc
	StyleSchemeChooserIfc
	gtk.Bin
}

func WrapStyleSchemeChooserWidget(p unsafe.Pointer) (r StyleSchemeChooserWidget) { r.P = p; return }

type IStyleSchemeChooserWidget interface{ P_StyleSchemeChooserWidget() unsafe.Pointer }

func (v StyleSchemeChooserWidget) P_StyleSchemeChooserWidget() unsafe.Pointer { return v.P }
func (v StyleSchemeChooserWidget) P_ImplementorIface() unsafe.Pointer         { return v.P }
func (v StyleSchemeChooserWidget) P_Buildable() unsafe.Pointer                { return v.P }
func (v StyleSchemeChooserWidget) P_StyleSchemeChooser() unsafe.Pointer       { return v.P }
func StyleSchemeChooserWidgetGetType() gi.GType {
	ret := _I.GetGType(68, "StyleSchemeChooserWidget")
	return ret
}

// gtk_source_style_scheme_chooser_widget_new
// container is not nil, container is StyleSchemeChooserWidget
// is constructor
func NewStyleSchemeChooserWidget() (result StyleSchemeChooserWidget) {
	iv, err := _I.Get(310, "StyleSchemeChooserWidget", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// ignore GType struct StyleSchemeChooserWidgetClass
// ignore GType struct StyleSchemeClass
// Object StyleSchemeManager
type StyleSchemeManager struct {
	gobject.Object
}

func WrapStyleSchemeManager(p unsafe.Pointer) (r StyleSchemeManager) { r.P = p; return }

type IStyleSchemeManager interface{ P_StyleSchemeManager() unsafe.Pointer }

func (v StyleSchemeManager) P_StyleSchemeManager() unsafe.Pointer { return v.P }
func StyleSchemeManagerGetType() gi.GType {
	ret := _I.GetGType(69, "StyleSchemeManager")
	return ret
}

// gtk_source_style_scheme_manager_new
// container is not nil, container is StyleSchemeManager
// is constructor
func NewStyleSchemeManager() (result StyleSchemeManager) {
	iv, err := _I.Get(311, "StyleSchemeManager", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_style_scheme_manager_append_search_path
// container is not nil, container is StyleSchemeManager
// is method
func (v StyleSchemeManager) AppendSearchPath(path string) {
	iv, err := _I.Get(313, "StyleSchemeManager", "append_search_path")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_path := gi.CString(path)
	arg_v := gi.NewPointerArgument(v.P)
	arg_path := gi.NewStringArgument(c_path)
	args := []gi.Argument{arg_v, arg_path}
	iv.Call(args, nil, nil)
	gi.Free(c_path)
}

// gtk_source_style_scheme_manager_force_rescan
// container is not nil, container is StyleSchemeManager
// is method
func (v StyleSchemeManager) ForceRescan() {
	iv, err := _I.Get(314, "StyleSchemeManager", "force_rescan")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_style_scheme_manager_get_scheme
// container is not nil, container is StyleSchemeManager
// is method
func (v StyleSchemeManager) GetScheme(scheme_id string) (result StyleScheme) {
	iv, err := _I.Get(315, "StyleSchemeManager", "get_scheme")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_scheme_id := gi.CString(scheme_id)
	arg_v := gi.NewPointerArgument(v.P)
	arg_scheme_id := gi.NewStringArgument(c_scheme_id)
	args := []gi.Argument{arg_v, arg_scheme_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_scheme_id)
	result.P = ret.Pointer()
	return
}

// gtk_source_style_scheme_manager_get_scheme_ids
// container is not nil, container is StyleSchemeManager
// is method
func (v StyleSchemeManager) GetSchemeIds() (result gi.CStrArray) {
	iv, err := _I.Get(316, "StyleSchemeManager", "get_scheme_ids")
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

// gtk_source_style_scheme_manager_get_search_path
// container is not nil, container is StyleSchemeManager
// is method
func (v StyleSchemeManager) GetSearchPath() (result gi.CStrArray) {
	iv, err := _I.Get(317, "StyleSchemeManager", "get_search_path")
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

// gtk_source_style_scheme_manager_prepend_search_path
// container is not nil, container is StyleSchemeManager
// is method
func (v StyleSchemeManager) PrependSearchPath(path string) {
	iv, err := _I.Get(318, "StyleSchemeManager", "prepend_search_path")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_path := gi.CString(path)
	arg_v := gi.NewPointerArgument(v.P)
	arg_path := gi.NewStringArgument(c_path)
	args := []gi.Argument{arg_v, arg_path}
	iv.Call(args, nil, nil)
	gi.Free(c_path)
}

// gtk_source_style_scheme_manager_set_search_path
// container is not nil, container is StyleSchemeManager
// is method
func (v StyleSchemeManager) SetSearchPath(path gi.CStrArray) {
	iv, err := _I.Get(319, "StyleSchemeManager", "set_search_path")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_path := gi.NewPointerArgument(path.P)
	args := []gi.Argument{arg_v, arg_path}
	iv.Call(args, nil, nil)
}

// ignore GType struct StyleSchemeManagerClass
// Struct StyleSchemeManagerPrivate
type StyleSchemeManagerPrivate struct {
	P unsafe.Pointer
}

func StyleSchemeManagerPrivateGetType() gi.GType {
	ret := _I.GetGType(70, "StyleSchemeManagerPrivate")
	return ret
}

// Struct StyleSchemePrivate
type StyleSchemePrivate struct {
	P unsafe.Pointer
}

func StyleSchemePrivateGetType() gi.GType {
	ret := _I.GetGType(71, "StyleSchemePrivate")
	return ret
}

// Object Tag
type Tag struct {
	gtk.TextTag
}

func WrapTag(p unsafe.Pointer) (r Tag) { r.P = p; return }

type ITag interface{ P_Tag() unsafe.Pointer }

func (v Tag) P_Tag() unsafe.Pointer { return v.P }
func TagGetType() gi.GType {
	ret := _I.GetGType(72, "Tag")
	return ret
}

// gtk_source_tag_new
// container is not nil, container is Tag
// is constructor
func NewTag(name string) (result Tag) {
	iv, err := _I.Get(320, "Tag", "new")
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
	result.P = ret.Pointer()
	return
}

// ignore GType struct TagClass
// Interface UndoManager
type UndoManager struct {
	UndoManagerIfc
	P unsafe.Pointer
}
type UndoManagerIfc struct{}
type IUndoManager interface{ P_UndoManager() unsafe.Pointer }

func (v UndoManager) P_UndoManager() unsafe.Pointer { return v.P }
func UndoManagerGetType() gi.GType {
	ret := _I.GetGType(73, "UndoManager")
	return ret
}

// gtk_source_undo_manager_begin_not_undoable_action
// container is not nil, container is UndoManager
// is method
func (v *UndoManagerIfc) BeginNotUndoableAction() {
	iv, err := _I.Get(321, "UndoManager", "begin_not_undoable_action")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_undo_manager_can_redo
// container is not nil, container is UndoManager
// is method
func (v *UndoManagerIfc) CanRedo() (result bool) {
	iv, err := _I.Get(322, "UndoManager", "can_redo")
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

// gtk_source_undo_manager_can_redo_changed
// container is not nil, container is UndoManager
// is method
func (v *UndoManagerIfc) CanRedoChanged() {
	iv, err := _I.Get(323, "UndoManager", "can_redo_changed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_undo_manager_can_undo
// container is not nil, container is UndoManager
// is method
func (v *UndoManagerIfc) CanUndo() (result bool) {
	iv, err := _I.Get(324, "UndoManager", "can_undo")
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

// gtk_source_undo_manager_can_undo_changed
// container is not nil, container is UndoManager
// is method
func (v *UndoManagerIfc) CanUndoChanged() {
	iv, err := _I.Get(325, "UndoManager", "can_undo_changed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_undo_manager_end_not_undoable_action
// container is not nil, container is UndoManager
// is method
func (v *UndoManagerIfc) EndNotUndoableAction() {
	iv, err := _I.Get(326, "UndoManager", "end_not_undoable_action")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_undo_manager_redo
// container is not nil, container is UndoManager
// is method
func (v *UndoManagerIfc) Redo() {
	iv, err := _I.Get(327, "UndoManager", "redo")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// gtk_source_undo_manager_undo
// container is not nil, container is UndoManager
// is method
func (v *UndoManagerIfc) Undo() {
	iv, err := _I.Get(328, "UndoManager", "undo")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// ignore GType struct UndoManagerIface
// Object View
type View struct {
	atk.ImplementorIfaceIfc
	gtk.BuildableIfc
	gtk.ScrollableIfc
	gtk.TextView
}

func WrapView(p unsafe.Pointer) (r View) { r.P = p; return }

type IView interface{ P_View() unsafe.Pointer }

func (v View) P_View() unsafe.Pointer             { return v.P }
func (v View) P_ImplementorIface() unsafe.Pointer { return v.P }
func (v View) P_Buildable() unsafe.Pointer        { return v.P }
func (v View) P_Scrollable() unsafe.Pointer       { return v.P }
func ViewGetType() gi.GType {
	ret := _I.GetGType(74, "View")
	return ret
}

// gtk_source_view_new
// container is not nil, container is View
// is constructor
func NewView() (result View) {
	iv, err := _I.Get(329, "View", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_view_new_with_buffer
// container is not nil, container is View
// is constructor
func NewViewWithBuffer(buffer IBuffer) (result View) {
	iv, err := _I.Get(330, "View", "new_with_buffer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if buffer != nil {
		tmp = buffer.P_Buffer()
	}
	arg_buffer := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_buffer}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_view_get_auto_indent
// container is not nil, container is View
// is method
func (v View) GetAutoIndent() (result bool) {
	iv, err := _I.Get(331, "View", "get_auto_indent")
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

// gtk_source_view_get_background_pattern
// container is not nil, container is View
// is method
func (v View) GetBackgroundPattern() (result BackgroundPatternTypeEnum) {
	iv, err := _I.Get(332, "View", "get_background_pattern")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = BackgroundPatternTypeEnum(ret.Int())
	return
}

// gtk_source_view_get_completion
// container is not nil, container is View
// is method
func (v View) GetCompletion() (result Completion) {
	iv, err := _I.Get(333, "View", "get_completion")
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

// gtk_source_view_get_gutter
// container is not nil, container is View
// is method
func (v View) GetGutter(window_type gtk.TextWindowTypeEnum) (result Gutter) {
	iv, err := _I.Get(334, "View", "get_gutter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_window_type := gi.NewIntArgument(int(window_type))
	args := []gi.Argument{arg_v, arg_window_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_view_get_highlight_current_line
// container is not nil, container is View
// is method
func (v View) GetHighlightCurrentLine() (result bool) {
	iv, err := _I.Get(335, "View", "get_highlight_current_line")
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

// gtk_source_view_get_indent_on_tab
// container is not nil, container is View
// is method
func (v View) GetIndentOnTab() (result bool) {
	iv, err := _I.Get(336, "View", "get_indent_on_tab")
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

// gtk_source_view_get_indent_width
// container is not nil, container is View
// is method
func (v View) GetIndentWidth() (result int32) {
	iv, err := _I.Get(337, "View", "get_indent_width")
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

// gtk_source_view_get_insert_spaces_instead_of_tabs
// container is not nil, container is View
// is method
func (v View) GetInsertSpacesInsteadOfTabs() (result bool) {
	iv, err := _I.Get(338, "View", "get_insert_spaces_instead_of_tabs")
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

// gtk_source_view_get_mark_attributes
// container is not nil, container is View
// is method
func (v View) GetMarkAttributes(category string, priority int32) (result MarkAttributes) {
	iv, err := _I.Get(339, "View", "get_mark_attributes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_category := gi.CString(category)
	arg_v := gi.NewPointerArgument(v.P)
	arg_category := gi.NewStringArgument(c_category)
	arg_priority := gi.NewInt32Argument(priority)
	args := []gi.Argument{arg_v, arg_category, arg_priority}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_category)
	result.P = ret.Pointer()
	return
}

// gtk_source_view_get_right_margin_position
// container is not nil, container is View
// is method
func (v View) GetRightMarginPosition() (result uint32) {
	iv, err := _I.Get(340, "View", "get_right_margin_position")
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

// gtk_source_view_get_show_line_marks
// container is not nil, container is View
// is method
func (v View) GetShowLineMarks() (result bool) {
	iv, err := _I.Get(341, "View", "get_show_line_marks")
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

// gtk_source_view_get_show_line_numbers
// container is not nil, container is View
// is method
func (v View) GetShowLineNumbers() (result bool) {
	iv, err := _I.Get(342, "View", "get_show_line_numbers")
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

// gtk_source_view_get_show_right_margin
// container is not nil, container is View
// is method
func (v View) GetShowRightMargin() (result bool) {
	iv, err := _I.Get(343, "View", "get_show_right_margin")
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

// gtk_source_view_get_smart_backspace
// container is not nil, container is View
// is method
func (v View) GetSmartBackspace() (result bool) {
	iv, err := _I.Get(344, "View", "get_smart_backspace")
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

// gtk_source_view_get_smart_home_end
// container is not nil, container is View
// is method
func (v View) GetSmartHomeEnd() (result SmartHomeEndTypeEnum) {
	iv, err := _I.Get(345, "View", "get_smart_home_end")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = SmartHomeEndTypeEnum(ret.Int())
	return
}

// gtk_source_view_get_space_drawer
// container is not nil, container is View
// is method
func (v View) GetSpaceDrawer() (result SpaceDrawer) {
	iv, err := _I.Get(346, "View", "get_space_drawer")
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

// gtk_source_view_get_tab_width
// container is not nil, container is View
// is method
func (v View) GetTabWidth() (result uint32) {
	iv, err := _I.Get(347, "View", "get_tab_width")
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

// gtk_source_view_get_visual_column
// container is not nil, container is View
// is method
func (v View) GetVisualColumn(iter gtk.TextIter) (result uint32) {
	iv, err := _I.Get(348, "View", "get_visual_column")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_iter := gi.NewPointerArgument(iter.P)
	args := []gi.Argument{arg_v, arg_iter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// gtk_source_view_indent_lines
// container is not nil, container is View
// is method
func (v View) IndentLines(start gtk.TextIter, end gtk.TextIter) {
	iv, err := _I.Get(349, "View", "indent_lines")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewPointerArgument(start.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_v, arg_start, arg_end}
	iv.Call(args, nil, nil)
}

// gtk_source_view_set_auto_indent
// container is not nil, container is View
// is method
func (v View) SetAutoIndent(enable bool) {
	iv, err := _I.Get(350, "View", "set_auto_indent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_enable := gi.NewBoolArgument(enable)
	args := []gi.Argument{arg_v, arg_enable}
	iv.Call(args, nil, nil)
}

// gtk_source_view_set_background_pattern
// container is not nil, container is View
// is method
func (v View) SetBackgroundPattern(background_pattern BackgroundPatternTypeEnum) {
	iv, err := _I.Get(351, "View", "set_background_pattern")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_background_pattern := gi.NewIntArgument(int(background_pattern))
	args := []gi.Argument{arg_v, arg_background_pattern}
	iv.Call(args, nil, nil)
}

// gtk_source_view_set_highlight_current_line
// container is not nil, container is View
// is method
func (v View) SetHighlightCurrentLine(highlight bool) {
	iv, err := _I.Get(352, "View", "set_highlight_current_line")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_highlight := gi.NewBoolArgument(highlight)
	args := []gi.Argument{arg_v, arg_highlight}
	iv.Call(args, nil, nil)
}

// gtk_source_view_set_indent_on_tab
// container is not nil, container is View
// is method
func (v View) SetIndentOnTab(enable bool) {
	iv, err := _I.Get(353, "View", "set_indent_on_tab")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_enable := gi.NewBoolArgument(enable)
	args := []gi.Argument{arg_v, arg_enable}
	iv.Call(args, nil, nil)
}

// gtk_source_view_set_indent_width
// container is not nil, container is View
// is method
func (v View) SetIndentWidth(width int32) {
	iv, err := _I.Get(354, "View", "set_indent_width")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_width := gi.NewInt32Argument(width)
	args := []gi.Argument{arg_v, arg_width}
	iv.Call(args, nil, nil)
}

// gtk_source_view_set_insert_spaces_instead_of_tabs
// container is not nil, container is View
// is method
func (v View) SetInsertSpacesInsteadOfTabs(enable bool) {
	iv, err := _I.Get(355, "View", "set_insert_spaces_instead_of_tabs")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_enable := gi.NewBoolArgument(enable)
	args := []gi.Argument{arg_v, arg_enable}
	iv.Call(args, nil, nil)
}

// gtk_source_view_set_mark_attributes
// container is not nil, container is View
// is method
func (v View) SetMarkAttributes(category string, attributes IMarkAttributes, priority int32) {
	iv, err := _I.Get(356, "View", "set_mark_attributes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_category := gi.CString(category)
	var tmp unsafe.Pointer
	if attributes != nil {
		tmp = attributes.P_MarkAttributes()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_category := gi.NewStringArgument(c_category)
	arg_attributes := gi.NewPointerArgument(tmp)
	arg_priority := gi.NewInt32Argument(priority)
	args := []gi.Argument{arg_v, arg_category, arg_attributes, arg_priority}
	iv.Call(args, nil, nil)
	gi.Free(c_category)
}

// gtk_source_view_set_right_margin_position
// container is not nil, container is View
// is method
func (v View) SetRightMarginPosition(pos uint32) {
	iv, err := _I.Get(357, "View", "set_right_margin_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewUint32Argument(pos)
	args := []gi.Argument{arg_v, arg_pos}
	iv.Call(args, nil, nil)
}

// gtk_source_view_set_show_line_marks
// container is not nil, container is View
// is method
func (v View) SetShowLineMarks(show bool) {
	iv, err := _I.Get(358, "View", "set_show_line_marks")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_show := gi.NewBoolArgument(show)
	args := []gi.Argument{arg_v, arg_show}
	iv.Call(args, nil, nil)
}

// gtk_source_view_set_show_line_numbers
// container is not nil, container is View
// is method
func (v View) SetShowLineNumbers(show bool) {
	iv, err := _I.Get(359, "View", "set_show_line_numbers")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_show := gi.NewBoolArgument(show)
	args := []gi.Argument{arg_v, arg_show}
	iv.Call(args, nil, nil)
}

// gtk_source_view_set_show_right_margin
// container is not nil, container is View
// is method
func (v View) SetShowRightMargin(show bool) {
	iv, err := _I.Get(360, "View", "set_show_right_margin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_show := gi.NewBoolArgument(show)
	args := []gi.Argument{arg_v, arg_show}
	iv.Call(args, nil, nil)
}

// gtk_source_view_set_smart_backspace
// container is not nil, container is View
// is method
func (v View) SetSmartBackspace(smart_backspace bool) {
	iv, err := _I.Get(361, "View", "set_smart_backspace")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_smart_backspace := gi.NewBoolArgument(smart_backspace)
	args := []gi.Argument{arg_v, arg_smart_backspace}
	iv.Call(args, nil, nil)
}

// gtk_source_view_set_smart_home_end
// container is not nil, container is View
// is method
func (v View) SetSmartHomeEnd(smart_home_end SmartHomeEndTypeEnum) {
	iv, err := _I.Get(362, "View", "set_smart_home_end")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_smart_home_end := gi.NewIntArgument(int(smart_home_end))
	args := []gi.Argument{arg_v, arg_smart_home_end}
	iv.Call(args, nil, nil)
}

// gtk_source_view_set_tab_width
// container is not nil, container is View
// is method
func (v View) SetTabWidth(width uint32) {
	iv, err := _I.Get(363, "View", "set_tab_width")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_width := gi.NewUint32Argument(width)
	args := []gi.Argument{arg_v, arg_width}
	iv.Call(args, nil, nil)
}

// gtk_source_view_unindent_lines
// container is not nil, container is View
// is method
func (v View) UnindentLines(start gtk.TextIter, end gtk.TextIter) {
	iv, err := _I.Get(364, "View", "unindent_lines")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_start := gi.NewPointerArgument(start.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_v, arg_start, arg_end}
	iv.Call(args, nil, nil)
}

// ignore GType struct ViewClass
// Enum ViewGutterPosition
type ViewGutterPositionEnum int

const (
	ViewGutterPositionLines ViewGutterPositionEnum = -30
	ViewGutterPositionMarks ViewGutterPositionEnum = -20
)

func ViewGutterPositionGetType() gi.GType {
	ret := _I.GetGType(75, "ViewGutterPosition")
	return ret
}

// Struct ViewPrivate
type ViewPrivate struct {
	P unsafe.Pointer
}

func ViewPrivateGetType() gi.GType {
	ret := _I.GetGType(76, "ViewPrivate")
	return ret
}

// gtk_source_completion_error_quark
// container is nil
func CompletionErrorQuark() (result uint32) {
	iv, err := _I.Get(365, "completion_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// gtk_source_encoding_get_all
// container is nil
func EncodingGetAll() (result glib.SList) {
	iv, err := _I.Get(366, "encoding_get_all", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_encoding_get_current
// container is nil
func EncodingGetCurrent() (result Encoding) {
	iv, err := _I.Get(367, "encoding_get_current", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_encoding_get_default_candidates
// container is nil
func EncodingGetDefaultCandidates() (result glib.SList) {
	iv, err := _I.Get(368, "encoding_get_default_candidates", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_encoding_get_from_charset
// container is nil
func EncodingGetFromCharset(charset string) (result Encoding) {
	iv, err := _I.Get(369, "encoding_get_from_charset", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_charset := gi.CString(charset)
	arg_charset := gi.NewStringArgument(c_charset)
	args := []gi.Argument{arg_charset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_charset)
	result.P = ret.Pointer()
	return
}

// gtk_source_encoding_get_utf8
// container is nil
func EncodingGetUtf8() (result Encoding) {
	iv, err := _I.Get(370, "encoding_get_utf8", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// gtk_source_file_loader_error_quark
// container is nil
func FileLoaderErrorQuark() (result uint32) {
	iv, err := _I.Get(371, "file_loader_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// gtk_source_file_saver_error_quark
// container is nil
func FileSaverErrorQuark() (result uint32) {
	iv, err := _I.Get(372, "file_saver_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// gtk_source_finalize
// container is nil
func Finalize() {
	iv, err := _I.Get(373, "finalize", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// gtk_source_init
// container is nil
func Init() {
	iv, err := _I.Get(374, "init", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// gtk_source_utils_escape_search_text
// container is nil
func UtilsEscapeSearchText(text string) (result string) {
	iv, err := _I.Get(375, "utils_escape_search_text", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_text := gi.NewStringArgument(c_text)
	args := []gi.Argument{arg_text}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_text)
	result = ret.String().Take()
	return
}

// gtk_source_utils_unescape_search_text
// container is nil
func UtilsUnescapeSearchText(text string) (result string) {
	iv, err := _I.Get(376, "utils_unescape_search_text", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_text := gi.NewStringArgument(c_text)
	args := []gi.Argument{arg_text}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_text)
	result = ret.String().Take()
	return
}

// constants
const ()