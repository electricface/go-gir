package glib

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
extern void myChildWatchFunc(gint32 pid, gint32 status, gpointer user_data);
static void* getPointer_myChildWatchFunc() {
return (void*)(myChildWatchFunc);
}
extern void myCompareDataFunc(gpointer a, gpointer b, gpointer user_data);
static void* getPointer_myCompareDataFunc() {
return (void*)(myCompareDataFunc);
}
extern void myDataForeachFunc(guint32 key_id, gpointer data, gpointer user_data);
static void* getPointer_myDataForeachFunc() {
return (void*)(myDataForeachFunc);
}
extern void myDuplicateFunc(gpointer data, gpointer user_data);
static void* getPointer_myDuplicateFunc() {
return (void*)(myDuplicateFunc);
}
extern void myFunc(gpointer data, gpointer user_data);
static void* getPointer_myFunc() {
return (void*)(myFunc);
}
extern void myHFunc(gpointer key, gpointer value, gpointer user_data);
static void* getPointer_myHFunc() {
return (void*)(myHFunc);
}
extern void myHRFunc(gpointer key, gpointer value, gpointer user_data);
static void* getPointer_myHRFunc() {
return (void*)(myHRFunc);
}
extern void myLogFunc(gchar* log_domain, GLogLevelFlags log_level, gchar* message, gpointer user_data);
static void* getPointer_myLogFunc() {
return (void*)(myLogFunc);
}
extern void myLogWriterFunc(GLogLevelFlags log_level, gpointer fields, guint64 n_fields, gpointer user_data);
static void* getPointer_myLogWriterFunc() {
return (void*)(myLogWriterFunc);
}
extern void myRegexEvalCallback(gpointer match_info, gpointer result, gpointer user_data);
static void* getPointer_myRegexEvalCallback() {
return (void*)(myRegexEvalCallback);
}
extern void mySourceFunc(gpointer user_data);
static void* getPointer_mySourceFunc() {
return (void*)(mySourceFunc);
}
extern void mySpawnChildSetupFunc(gpointer user_data);
static void* getPointer_mySpawnChildSetupFunc() {
return (void*)(mySpawnChildSetupFunc);
}
extern void myTestDataFunc(gpointer user_data);
static void* getPointer_myTestDataFunc() {
return (void*)(myTestDataFunc);
}
extern void myTestFixtureFunc(gpointer fixture, gpointer user_data);
static void* getPointer_myTestFixtureFunc() {
return (void*)(myTestFixtureFunc);
}
extern void myTestLogFatalFunc(gchar* log_domain, GLogLevelFlags log_level, gchar* message, gpointer user_data);
static void* getPointer_myTestLogFatalFunc() {
return (void*)(myTestLogFatalFunc);
}
extern void myUnixFDSourceFunc(gint32 fd, GIOCondition condition, gpointer user_data);
static void* getPointer_myUnixFDSourceFunc() {
return (void*)(myUnixFDSourceFunc);
}
*/
import "C"
import "github.com/electricface/go-gir3/gi"
import "log"
import "unsafe"

var _I = gi.NewInvokerCache("GLib")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GLib", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Struct Array
type Array struct {
	P unsafe.Pointer
}

const SizeOfStructArray = 16

func ArrayGetType() gi.GType {
	ret := _I.GetGType(0, "Array")
	return ret
}

// Flags AsciiType
type AsciiTypeFlags int

const (
	AsciiTypeAlnum  AsciiTypeFlags = 1
	AsciiTypeAlpha  AsciiTypeFlags = 2
	AsciiTypeCntrl  AsciiTypeFlags = 4
	AsciiTypeDigit  AsciiTypeFlags = 8
	AsciiTypeGraph  AsciiTypeFlags = 16
	AsciiTypeLower  AsciiTypeFlags = 32
	AsciiTypePrint  AsciiTypeFlags = 64
	AsciiTypePunct  AsciiTypeFlags = 128
	AsciiTypeSpace  AsciiTypeFlags = 256
	AsciiTypeUpper  AsciiTypeFlags = 512
	AsciiTypeXdigit AsciiTypeFlags = 1024
)

func AsciiTypeGetType() gi.GType {
	ret := _I.GetGType(1, "AsciiType")
	return ret
}

// Struct AsyncQueue
type AsyncQueue struct {
	P unsafe.Pointer
}

func AsyncQueueGetType() gi.GType {
	ret := _I.GetGType(2, "AsyncQueue")
	return ret
}

// g_async_queue_length
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) Length() (result int32) {
	iv, err := _I.Get(0, "AsyncQueue", "length")
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

// g_async_queue_length_unlocked
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) LengthUnlocked() (result int32) {
	iv, err := _I.Get(1, "AsyncQueue", "length_unlocked")
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

// g_async_queue_lock
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) Lock() {
	iv, err := _I.Get(2, "AsyncQueue", "lock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_async_queue_pop
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) Pop() (result unsafe.Pointer) {
	iv, err := _I.Get(3, "AsyncQueue", "pop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_async_queue_pop_unlocked
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) PopUnlocked() (result unsafe.Pointer) {
	iv, err := _I.Get(4, "AsyncQueue", "pop_unlocked")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_async_queue_push
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) Push(data unsafe.Pointer) {
	iv, err := _I.Get(5, "AsyncQueue", "push")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	iv.Call(args, nil, nil)
}

// g_async_queue_push_front
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) PushFront(item unsafe.Pointer) {
	iv, err := _I.Get(6, "AsyncQueue", "push_front")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_item := gi.NewPointerArgument(item)
	args := []gi.Argument{arg_v, arg_item}
	iv.Call(args, nil, nil)
}

// g_async_queue_push_front_unlocked
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) PushFrontUnlocked(item unsafe.Pointer) {
	iv, err := _I.Get(7, "AsyncQueue", "push_front_unlocked")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_item := gi.NewPointerArgument(item)
	args := []gi.Argument{arg_v, arg_item}
	iv.Call(args, nil, nil)
}

// g_async_queue_push_unlocked
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) PushUnlocked(data unsafe.Pointer) {
	iv, err := _I.Get(8, "AsyncQueue", "push_unlocked")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	iv.Call(args, nil, nil)
}

// g_async_queue_ref_unlocked
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) RefUnlocked() {
	iv, err := _I.Get(9, "AsyncQueue", "ref_unlocked")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_async_queue_remove
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) Remove(item unsafe.Pointer) (result bool) {
	iv, err := _I.Get(10, "AsyncQueue", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_item := gi.NewPointerArgument(item)
	args := []gi.Argument{arg_v, arg_item}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_async_queue_remove_unlocked
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) RemoveUnlocked(item unsafe.Pointer) (result bool) {
	iv, err := _I.Get(11, "AsyncQueue", "remove_unlocked")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_item := gi.NewPointerArgument(item)
	args := []gi.Argument{arg_v, arg_item}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_async_queue_timed_pop
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) TimedPop(end_time TimeVal) (result unsafe.Pointer) {
	iv, err := _I.Get(12, "AsyncQueue", "timed_pop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_end_time := gi.NewPointerArgument(end_time.P)
	args := []gi.Argument{arg_v, arg_end_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_async_queue_timed_pop_unlocked
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) TimedPopUnlocked(end_time TimeVal) (result unsafe.Pointer) {
	iv, err := _I.Get(13, "AsyncQueue", "timed_pop_unlocked")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_end_time := gi.NewPointerArgument(end_time.P)
	args := []gi.Argument{arg_v, arg_end_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_async_queue_timeout_pop
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) TimeoutPop(timeout uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(14, "AsyncQueue", "timeout_pop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timeout := gi.NewUint64Argument(timeout)
	args := []gi.Argument{arg_v, arg_timeout}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_async_queue_timeout_pop_unlocked
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) TimeoutPopUnlocked(timeout uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(15, "AsyncQueue", "timeout_pop_unlocked")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timeout := gi.NewUint64Argument(timeout)
	args := []gi.Argument{arg_v, arg_timeout}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_async_queue_try_pop
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) TryPop() (result unsafe.Pointer) {
	iv, err := _I.Get(16, "AsyncQueue", "try_pop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_async_queue_try_pop_unlocked
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) TryPopUnlocked() (result unsafe.Pointer) {
	iv, err := _I.Get(17, "AsyncQueue", "try_pop_unlocked")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_async_queue_unlock
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) Unlock() {
	iv, err := _I.Get(18, "AsyncQueue", "unlock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_async_queue_unref
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) Unref() {
	iv, err := _I.Get(19, "AsyncQueue", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_async_queue_unref_and_unlock
// container is not nil, container is AsyncQueue
// is method
func (v AsyncQueue) UnrefAndUnlock() {
	iv, err := _I.Get(20, "AsyncQueue", "unref_and_unlock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct BookmarkFile
type BookmarkFile struct {
	P unsafe.Pointer
}

func BookmarkFileGetType() gi.GType {
	ret := _I.GetGType(3, "BookmarkFile")
	return ret
}

// g_bookmark_file_add_application
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) AddApplication(uri string, name string, exec string) {
	iv, err := _I.Get(21, "BookmarkFile", "add_application")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	c_name := gi.CString(name)
	c_exec := gi.CString(exec)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_name := gi.NewStringArgument(c_name)
	arg_exec := gi.NewStringArgument(c_exec)
	args := []gi.Argument{arg_v, arg_uri, arg_name, arg_exec}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
	gi.Free(c_name)
	gi.Free(c_exec)
}

// g_bookmark_file_add_group
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) AddGroup(uri string, group string) {
	iv, err := _I.Get(22, "BookmarkFile", "add_group")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	c_group := gi.CString(group)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_group := gi.NewStringArgument(c_group)
	args := []gi.Argument{arg_v, arg_uri, arg_group}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
	gi.Free(c_group)
}

// g_bookmark_file_free
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) Free() {
	iv, err := _I.Get(23, "BookmarkFile", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_bookmark_file_get_added
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) GetAdded(uri string) (result int64, err error) {
	iv, err := _I.Get(24, "BookmarkFile", "get_added")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int64()
	return
}

// g_bookmark_file_get_app_info
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) GetAppInfo(uri string, name string) (result bool, exec string, count uint32, stamp int64, err error) {
	iv, err := _I.Get(25, "BookmarkFile", "get_app_info")
	if err != nil {
		return
	}
	var outArgs [4]gi.Argument
	c_uri := gi.CString(uri)
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_name := gi.NewStringArgument(c_name)
	arg_exec := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_count := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_stamp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_uri, arg_name, arg_exec, arg_count, arg_stamp, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	gi.Free(c_name)
	err = gi.ToError(outArgs[3].Pointer())
	exec = outArgs[0].String().Take()
	count = outArgs[1].Uint32()
	stamp = outArgs[2].Int64()
	result = ret.Bool()
	return
}

// g_bookmark_file_get_applications
// container is not nil, container is BookmarkFile
// is method
// ret lenArgIdx 1
func (v BookmarkFile) GetApplications(uri string) (result gi.CStrArray, err error) {
	iv, err := _I.Get(26, "BookmarkFile", "get_applications")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_uri, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: int(length)}
	return
}

// g_bookmark_file_get_description
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) GetDescription(uri string) (result string, err error) {
	iv, err := _I.Get(27, "BookmarkFile", "get_description")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_bookmark_file_get_groups
// container is not nil, container is BookmarkFile
// is method
// ret lenArgIdx 1
func (v BookmarkFile) GetGroups(uri string) (result gi.CStrArray, err error) {
	iv, err := _I.Get(28, "BookmarkFile", "get_groups")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_uri, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: int(length)}
	return
}

// g_bookmark_file_get_icon
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) GetIcon(uri string) (result bool, href string, mime_type string, err error) {
	iv, err := _I.Get(29, "BookmarkFile", "get_icon")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_href := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_mime_type := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_uri, arg_href, arg_mime_type, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[2].Pointer())
	href = outArgs[0].String().Take()
	mime_type = outArgs[1].String().Take()
	result = ret.Bool()
	return
}

// g_bookmark_file_get_is_private
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) GetIsPrivate(uri string) (result bool, err error) {
	iv, err := _I.Get(30, "BookmarkFile", "get_is_private")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_get_mime_type
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) GetMimeType(uri string) (result string, err error) {
	iv, err := _I.Get(31, "BookmarkFile", "get_mime_type")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_bookmark_file_get_modified
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) GetModified(uri string) (result int64, err error) {
	iv, err := _I.Get(32, "BookmarkFile", "get_modified")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int64()
	return
}

// g_bookmark_file_get_size
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) GetSize() (result int32) {
	iv, err := _I.Get(33, "BookmarkFile", "get_size")
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

// g_bookmark_file_get_title
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) GetTitle(uri string) (result string, err error) {
	iv, err := _I.Get(34, "BookmarkFile", "get_title")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_bookmark_file_get_uris
// container is not nil, container is BookmarkFile
// is method
// ret lenArgIdx 0
func (v BookmarkFile) GetUris() (result gi.CStrArray) {
	iv, err := _I.Get(35, "BookmarkFile", "get_uris")
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
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: int(length)}
	return
}

// g_bookmark_file_get_visited
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) GetVisited(uri string) (result int64, err error) {
	iv, err := _I.Get(36, "BookmarkFile", "get_visited")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int64()
	return
}

// g_bookmark_file_has_application
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) HasApplication(uri string, name string) (result bool, err error) {
	iv, err := _I.Get(37, "BookmarkFile", "has_application")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_name := gi.NewStringArgument(c_name)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_name, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	gi.Free(c_name)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_has_group
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) HasGroup(uri string, group string) (result bool, err error) {
	iv, err := _I.Get(38, "BookmarkFile", "has_group")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	c_group := gi.CString(group)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_group := gi.NewStringArgument(c_group)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_group, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	gi.Free(c_group)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_has_item
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) HasItem(uri string) (result bool) {
	iv, err := _I.Get(39, "BookmarkFile", "has_item")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	args := []gi.Argument{arg_v, arg_uri}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_uri)
	result = ret.Bool()
	return
}

// g_bookmark_file_load_from_data
// container is not nil, container is BookmarkFile
// is method
// arg 0 data lenArgIdx 1
func (v BookmarkFile) LoadFromData(data gi.Uint8Array, length uint64) (result bool, err error) {
	iv, err := _I.Get(40, "BookmarkFile", "load_from_data")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_length := gi.NewUint64Argument(length)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_data, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_load_from_data_dirs
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) LoadFromDataDirs(file string) (result bool, full_path string, err error) {
	iv, err := _I.Get(41, "BookmarkFile", "load_from_data_dirs")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_file := gi.CString(file)
	arg_v := gi.NewPointerArgument(v.P)
	arg_file := gi.NewStringArgument(c_file)
	arg_full_path := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_file, arg_full_path, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_file)
	err = gi.ToError(outArgs[1].Pointer())
	full_path = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// g_bookmark_file_load_from_file
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) LoadFromFile(filename string) (result bool, err error) {
	iv, err := _I.Get(42, "BookmarkFile", "load_from_file")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_v := gi.NewPointerArgument(v.P)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_filename, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_move_item
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) MoveItem(old_uri string, new_uri string) (result bool, err error) {
	iv, err := _I.Get(43, "BookmarkFile", "move_item")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_old_uri := gi.CString(old_uri)
	c_new_uri := gi.CString(new_uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_old_uri := gi.NewStringArgument(c_old_uri)
	arg_new_uri := gi.NewStringArgument(c_new_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_old_uri, arg_new_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_old_uri)
	gi.Free(c_new_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_remove_application
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) RemoveApplication(uri string, name string) (result bool, err error) {
	iv, err := _I.Get(44, "BookmarkFile", "remove_application")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_name := gi.NewStringArgument(c_name)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_name, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	gi.Free(c_name)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_remove_group
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) RemoveGroup(uri string, group string) (result bool, err error) {
	iv, err := _I.Get(45, "BookmarkFile", "remove_group")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	c_group := gi.CString(group)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_group := gi.NewStringArgument(c_group)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_group, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	gi.Free(c_group)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_remove_item
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) RemoveItem(uri string) (result bool, err error) {
	iv, err := _I.Get(46, "BookmarkFile", "remove_item")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_set_added
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) SetAdded(uri string, added int64) {
	iv, err := _I.Get(47, "BookmarkFile", "set_added")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_added := gi.NewInt64Argument(added)
	args := []gi.Argument{arg_v, arg_uri, arg_added}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
}

// g_bookmark_file_set_app_info
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) SetAppInfo(uri string, name string, exec string, count int32, stamp int64) (result bool, err error) {
	iv, err := _I.Get(48, "BookmarkFile", "set_app_info")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_uri := gi.CString(uri)
	c_name := gi.CString(name)
	c_exec := gi.CString(exec)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_name := gi.NewStringArgument(c_name)
	arg_exec := gi.NewStringArgument(c_exec)
	arg_count := gi.NewInt32Argument(count)
	arg_stamp := gi.NewInt64Argument(stamp)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_uri, arg_name, arg_exec, arg_count, arg_stamp, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	gi.Free(c_name)
	gi.Free(c_exec)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_bookmark_file_set_description
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) SetDescription(uri string, description string) {
	iv, err := _I.Get(49, "BookmarkFile", "set_description")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	c_description := gi.CString(description)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_description := gi.NewStringArgument(c_description)
	args := []gi.Argument{arg_v, arg_uri, arg_description}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
	gi.Free(c_description)
}

// g_bookmark_file_set_groups
// container is not nil, container is BookmarkFile
// is method
// arg 1 groups lenArgIdx 2
func (v BookmarkFile) SetGroups(uri string, groups gi.CStrArray, length uint64) {
	iv, err := _I.Get(50, "BookmarkFile", "set_groups")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_groups := gi.NewPointerArgument(groups.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_v, arg_uri, arg_groups, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
}

// g_bookmark_file_set_icon
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) SetIcon(uri string, href string, mime_type string) {
	iv, err := _I.Get(51, "BookmarkFile", "set_icon")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	c_href := gi.CString(href)
	c_mime_type := gi.CString(mime_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_href := gi.NewStringArgument(c_href)
	arg_mime_type := gi.NewStringArgument(c_mime_type)
	args := []gi.Argument{arg_v, arg_uri, arg_href, arg_mime_type}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
	gi.Free(c_href)
	gi.Free(c_mime_type)
}

// g_bookmark_file_set_is_private
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) SetIsPrivate(uri string, is_private bool) {
	iv, err := _I.Get(52, "BookmarkFile", "set_is_private")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_is_private := gi.NewBoolArgument(is_private)
	args := []gi.Argument{arg_v, arg_uri, arg_is_private}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
}

// g_bookmark_file_set_mime_type
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) SetMimeType(uri string, mime_type string) {
	iv, err := _I.Get(53, "BookmarkFile", "set_mime_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	c_mime_type := gi.CString(mime_type)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_mime_type := gi.NewStringArgument(c_mime_type)
	args := []gi.Argument{arg_v, arg_uri, arg_mime_type}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
	gi.Free(c_mime_type)
}

// g_bookmark_file_set_modified
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) SetModified(uri string, modified int64) {
	iv, err := _I.Get(54, "BookmarkFile", "set_modified")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_modified := gi.NewInt64Argument(modified)
	args := []gi.Argument{arg_v, arg_uri, arg_modified}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
}

// g_bookmark_file_set_title
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) SetTitle(uri string, title string) {
	iv, err := _I.Get(55, "BookmarkFile", "set_title")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	c_title := gi.CString(title)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_title := gi.NewStringArgument(c_title)
	args := []gi.Argument{arg_v, arg_uri, arg_title}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
	gi.Free(c_title)
}

// g_bookmark_file_set_visited
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) SetVisited(uri string, visited int64) {
	iv, err := _I.Get(56, "BookmarkFile", "set_visited")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_v := gi.NewPointerArgument(v.P)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_visited := gi.NewInt64Argument(visited)
	args := []gi.Argument{arg_v, arg_uri, arg_visited}
	iv.Call(args, nil, nil)
	gi.Free(c_uri)
}

// g_bookmark_file_to_data
// container is not nil, container is BookmarkFile
// is method
// ret lenArgIdx 0
func (v BookmarkFile) ToData() (result gi.Uint8Array, err error) {
	iv, err := _I.Get(57, "BookmarkFile", "to_data")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(length)}
	return
}

// g_bookmark_file_to_file
// container is not nil, container is BookmarkFile
// is method
func (v BookmarkFile) ToFile(filename string) (result bool, err error) {
	iv, err := _I.Get(58, "BookmarkFile", "to_file")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_v := gi.NewPointerArgument(v.P)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_filename, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// Enum BookmarkFileError
type BookmarkFileErrorEnum int

const (
	BookmarkFileErrorInvalidUri       BookmarkFileErrorEnum = 0
	BookmarkFileErrorInvalidValue     BookmarkFileErrorEnum = 1
	BookmarkFileErrorAppNotRegistered BookmarkFileErrorEnum = 2
	BookmarkFileErrorUriNotFound      BookmarkFileErrorEnum = 3
	BookmarkFileErrorRead             BookmarkFileErrorEnum = 4
	BookmarkFileErrorUnknownEncoding  BookmarkFileErrorEnum = 5
	BookmarkFileErrorWrite            BookmarkFileErrorEnum = 6
	BookmarkFileErrorFileNotFound     BookmarkFileErrorEnum = 7
)

func BookmarkFileErrorGetType() gi.GType {
	ret := _I.GetGType(4, "BookmarkFileError")
	return ret
}

// Struct ByteArray
type ByteArray struct {
	P unsafe.Pointer
}

const SizeOfStructByteArray = 16

func ByteArrayGetType() gi.GType {
	ret := _I.GetGType(5, "ByteArray")
	return ret
}

// g_byte_array_free
// container is not nil, container is ByteArray
// is method
// arg0Type tag: array, isPtr: true
func ByteArrayFree1(array ByteArray, free_segment bool) (result uint8) {
	iv, err := _I.Get(60, "ByteArray", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_array := gi.NewPointerArgument(array.P)
	arg_free_segment := gi.NewBoolArgument(free_segment)
	args := []gi.Argument{arg_array, arg_free_segment}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_byte_array_free_to_bytes
// container is not nil, container is ByteArray
// is method
// arg0Type tag: array, isPtr: true
func ByteArrayFreeToBytes1(array ByteArray) (result Bytes) {
	iv, err := _I.Get(61, "ByteArray", "free_to_bytes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_array := gi.NewPointerArgument(array.P)
	args := []gi.Argument{arg_array}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_byte_array_new_take
// container is not nil, container is ByteArray
// is method
// arg0Type tag: array, isPtr: true
// arg 0 data lenArgIdx 1
func ByteArrayNewTake1(data gi.Uint8Array, len1 uint64) (result ByteArray) {
	iv, err := _I.Get(63, "ByteArray", "new_take")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data.P)
	arg_len1 := gi.NewUint64Argument(len1)
	args := []gi.Argument{arg_data, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_byte_array_unref
// container is not nil, container is ByteArray
// is method
// arg0Type tag: array, isPtr: true
func ByteArrayUnref1(array ByteArray) {
	iv, err := _I.Get(64, "ByteArray", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_array := gi.NewPointerArgument(array.P)
	args := []gi.Argument{arg_array}
	iv.Call(args, nil, nil)
}

// Struct Bytes
type Bytes struct {
	P unsafe.Pointer
}

func BytesGetType() gi.GType {
	ret := _I.GetGType(6, "Bytes")
	return ret
}

// g_bytes_new
// container is not nil, container is Bytes
// is constructor
// arg 0 data lenArgIdx 1
func NewBytes(data gi.Uint8Array, size uint64) (result Bytes) {
	iv, err := _I.Get(65, "Bytes", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data.P)
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_data, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_bytes_new_take
// container is not nil, container is Bytes
// is constructor
// arg 0 data lenArgIdx 1
func NewBytesTake(data gi.Uint8Array, size uint64) (result Bytes) {
	iv, err := _I.Get(66, "Bytes", "new_take")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data.P)
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_data, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_bytes_compare
// container is not nil, container is Bytes
// is method
func (v Bytes) Compare(bytes2 Bytes) (result int32) {
	iv, err := _I.Get(67, "Bytes", "compare")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_bytes2 := gi.NewPointerArgument(bytes2.P)
	args := []gi.Argument{arg_v, arg_bytes2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_bytes_equal
// container is not nil, container is Bytes
// is method
func (v Bytes) Equal(bytes2 Bytes) (result bool) {
	iv, err := _I.Get(68, "Bytes", "equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_bytes2 := gi.NewPointerArgument(bytes2.P)
	args := []gi.Argument{arg_v, arg_bytes2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_bytes_get_data
// container is not nil, container is Bytes
// is method
// ret lenArgIdx 0
func (v Bytes) GetData() (result gi.Uint8Array) {
	iv, err := _I.Get(69, "Bytes", "get_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var size uint64
	_ = size
	size = outArgs[0].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(size)}
	return
}

// g_bytes_get_size
// container is not nil, container is Bytes
// is method
func (v Bytes) GetSize() (result uint64) {
	iv, err := _I.Get(70, "Bytes", "get_size")
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

// g_bytes_hash
// container is not nil, container is Bytes
// is method
func (v Bytes) Hash() (result uint32) {
	iv, err := _I.Get(71, "Bytes", "hash")
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

// g_bytes_new_from_bytes
// container is not nil, container is Bytes
// is method
func (v Bytes) NewFromBytes(offset uint64, length uint64) (result Bytes) {
	iv, err := _I.Get(72, "Bytes", "new_from_bytes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewUint64Argument(offset)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_v, arg_offset, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_bytes_ref
// container is not nil, container is Bytes
// is method
func (v Bytes) Ref() (result Bytes) {
	iv, err := _I.Get(73, "Bytes", "ref")
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

// g_bytes_unref
// container is not nil, container is Bytes
// is method
func (v Bytes) Unref() {
	iv, err := _I.Get(74, "Bytes", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_bytes_unref_to_array
// container is not nil, container is Bytes
// is method
func (v Bytes) UnrefToArray() (result ByteArray) {
	iv, err := _I.Get(75, "Bytes", "unref_to_array")
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

// g_bytes_unref_to_data
// container is not nil, container is Bytes
// is method
// ret lenArgIdx 0
func (v Bytes) UnrefToData() (result gi.Uint8Array) {
	iv, err := _I.Get(76, "Bytes", "unref_to_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var size uint64
	_ = size
	size = outArgs[0].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(size)}
	return
}

// Struct Checksum
type Checksum struct {
	P unsafe.Pointer
}

func ChecksumGetType() gi.GType {
	ret := _I.GetGType(7, "Checksum")
	return ret
}

// g_checksum_new
// container is not nil, container is Checksum
// is constructor
func NewChecksum(checksum_type ChecksumTypeEnum) (result Checksum) {
	iv, err := _I.Get(77, "Checksum", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_checksum_type := gi.NewIntArgument(int(checksum_type))
	args := []gi.Argument{arg_checksum_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_checksum_copy
// container is not nil, container is Checksum
// is method
func (v Checksum) Copy() (result Checksum) {
	iv, err := _I.Get(78, "Checksum", "copy")
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

// g_checksum_free
// container is not nil, container is Checksum
// is method
func (v Checksum) Free() {
	iv, err := _I.Get(79, "Checksum", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_checksum_get_string
// container is not nil, container is Checksum
// is method
func (v Checksum) GetString() (result string) {
	iv, err := _I.Get(80, "Checksum", "get_string")
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

// g_checksum_reset
// container is not nil, container is Checksum
// is method
func (v Checksum) Reset() {
	iv, err := _I.Get(81, "Checksum", "reset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_checksum_update
// container is not nil, container is Checksum
// is method
// arg 0 data lenArgIdx 1
func (v Checksum) Update(data gi.Uint8Array, length int64) {
	iv, err := _I.Get(82, "Checksum", "update")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_v, arg_data, arg_length}
	iv.Call(args, nil, nil)
}

// g_checksum_type_get_length
// container is not nil, container is Checksum
// is method
// arg0Type tag: interface, isPtr: false
func ChecksumTypeGetLength1(checksum_type ChecksumTypeEnum) (result int64) {
	iv, err := _I.Get(83, "Checksum", "type_get_length")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_checksum_type := gi.NewIntArgument(int(checksum_type))
	args := []gi.Argument{arg_checksum_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// Enum ChecksumType
type ChecksumTypeEnum int

const (
	ChecksumTypeMd5    ChecksumTypeEnum = 0
	ChecksumTypeSha1   ChecksumTypeEnum = 1
	ChecksumTypeSha256 ChecksumTypeEnum = 2
	ChecksumTypeSha512 ChecksumTypeEnum = 3
	ChecksumTypeSha384 ChecksumTypeEnum = 4
)

func ChecksumTypeGetType() gi.GType {
	ret := _I.GetGType(8, "ChecksumType")
	return ret
}

type ChildWatchFuncStruct struct {
	F_pid    int32
	F_status int32
}

//export myChildWatchFunc
func myChildWatchFunc(pid C.gint32, status C.gint32, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := ChildWatchFuncStruct{
		F_pid:    int32(pid),
		F_status: int32(status),
	}
	fn(args)
}

// ignore callback ClearHandleFunc
type CompareDataFuncStruct struct {
	F_a unsafe.Pointer
	F_b unsafe.Pointer
}

//export myCompareDataFunc
func myCompareDataFunc(a C.gpointer, b C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := CompareDataFuncStruct{
		F_a: unsafe.Pointer(a),
		F_b: unsafe.Pointer(b),
	}
	fn(args)
}

// ignore callback CompareFunc
// Struct Cond
type Cond struct {
	P unsafe.Pointer
}

const SizeOfStructCond = 16

func CondGetType() gi.GType {
	ret := _I.GetGType(9, "Cond")
	return ret
}

// g_cond_broadcast
// container is not nil, container is Cond
// is method
func (v Cond) Broadcast() {
	iv, err := _I.Get(84, "Cond", "broadcast")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_cond_clear
// container is not nil, container is Cond
// is method
func (v Cond) Clear() {
	iv, err := _I.Get(85, "Cond", "clear")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_cond_init
// container is not nil, container is Cond
// is method
func (v Cond) Init() {
	iv, err := _I.Get(86, "Cond", "init")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_cond_signal
// container is not nil, container is Cond
// is method
func (v Cond) Signal() {
	iv, err := _I.Get(87, "Cond", "signal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_cond_wait
// container is not nil, container is Cond
// is method
func (v Cond) Wait(mutex Mutex) {
	iv, err := _I.Get(88, "Cond", "wait")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mutex := gi.NewPointerArgument(mutex.P)
	args := []gi.Argument{arg_v, arg_mutex}
	iv.Call(args, nil, nil)
}

// g_cond_wait_until
// container is not nil, container is Cond
// is method
func (v Cond) WaitUntil(mutex Mutex, end_time int64) (result bool) {
	iv, err := _I.Get(89, "Cond", "wait_until")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_mutex := gi.NewPointerArgument(mutex.P)
	arg_end_time := gi.NewInt64Argument(end_time)
	args := []gi.Argument{arg_v, arg_mutex, arg_end_time}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Enum ConvertError
type ConvertErrorEnum int

const (
	ConvertErrorNoConversion    ConvertErrorEnum = 0
	ConvertErrorIllegalSequence ConvertErrorEnum = 1
	ConvertErrorFailed          ConvertErrorEnum = 2
	ConvertErrorPartialInput    ConvertErrorEnum = 3
	ConvertErrorBadUri          ConvertErrorEnum = 4
	ConvertErrorNotAbsolutePath ConvertErrorEnum = 5
	ConvertErrorNoMemory        ConvertErrorEnum = 6
	ConvertErrorEmbeddedNul     ConvertErrorEnum = 7
)

func ConvertErrorGetType() gi.GType {
	ret := _I.GetGType(10, "ConvertError")
	return ret
}

// ignore callback CopyFunc
// Struct Data
type Data struct {
	P unsafe.Pointer
}

func DataGetType() gi.GType {
	ret := _I.GetGType(11, "Data")
	return ret
}

type DataForeachFuncStruct struct {
	F_key_id uint32
	F_data   unsafe.Pointer
}

//export myDataForeachFunc
func myDataForeachFunc(key_id C.guint32, data C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := DataForeachFuncStruct{
		F_key_id: uint32(key_id),
		F_data:   unsafe.Pointer(data),
	}
	fn(args)
}

// Struct Date
type Date struct {
	P unsafe.Pointer
}

const SizeOfStructDate = 24

func DateGetType() gi.GType {
	ret := _I.GetGType(12, "Date")
	return ret
}

// g_date_new
// container is not nil, container is Date
// is constructor
func NewDate() (result Date) {
	iv, err := _I.Get(90, "Date", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_new_dmy
// container is not nil, container is Date
// is constructor
func NewDateDmy(day uint8, month DateMonthEnum, year uint16) (result Date) {
	iv, err := _I.Get(91, "Date", "new_dmy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_day := gi.NewUint8Argument(day)
	arg_month := gi.NewIntArgument(int(month))
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_day, arg_month, arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_new_julian
// container is not nil, container is Date
// is constructor
func NewDateJulian(julian_day uint32) (result Date) {
	iv, err := _I.Get(92, "Date", "new_julian")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_julian_day := gi.NewUint32Argument(julian_day)
	args := []gi.Argument{arg_julian_day}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_add_days
// container is not nil, container is Date
// is method
func (v Date) AddDays(n_days uint32) {
	iv, err := _I.Get(93, "Date", "add_days")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_days := gi.NewUint32Argument(n_days)
	args := []gi.Argument{arg_v, arg_n_days}
	iv.Call(args, nil, nil)
}

// g_date_add_months
// container is not nil, container is Date
// is method
func (v Date) AddMonths(n_months uint32) {
	iv, err := _I.Get(94, "Date", "add_months")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_months := gi.NewUint32Argument(n_months)
	args := []gi.Argument{arg_v, arg_n_months}
	iv.Call(args, nil, nil)
}

// g_date_add_years
// container is not nil, container is Date
// is method
func (v Date) AddYears(n_years uint32) {
	iv, err := _I.Get(95, "Date", "add_years")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_years := gi.NewUint32Argument(n_years)
	args := []gi.Argument{arg_v, arg_n_years}
	iv.Call(args, nil, nil)
}

// g_date_clamp
// container is not nil, container is Date
// is method
func (v Date) Clamp(min_date Date, max_date Date) {
	iv, err := _I.Get(96, "Date", "clamp")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_min_date := gi.NewPointerArgument(min_date.P)
	arg_max_date := gi.NewPointerArgument(max_date.P)
	args := []gi.Argument{arg_v, arg_min_date, arg_max_date}
	iv.Call(args, nil, nil)
}

// g_date_clear
// container is not nil, container is Date
// is method
func (v Date) Clear(n_dates uint32) {
	iv, err := _I.Get(97, "Date", "clear")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_dates := gi.NewUint32Argument(n_dates)
	args := []gi.Argument{arg_v, arg_n_dates}
	iv.Call(args, nil, nil)
}

// g_date_compare
// container is not nil, container is Date
// is method
func (v Date) Compare(rhs Date) (result int32) {
	iv, err := _I.Get(98, "Date", "compare")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_rhs := gi.NewPointerArgument(rhs.P)
	args := []gi.Argument{arg_v, arg_rhs}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_copy
// container is not nil, container is Date
// is method
func (v Date) Copy() (result Date) {
	iv, err := _I.Get(99, "Date", "copy")
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

// g_date_days_between
// container is not nil, container is Date
// is method
func (v Date) DaysBetween(date2 Date) (result int32) {
	iv, err := _I.Get(100, "Date", "days_between")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_date2 := gi.NewPointerArgument(date2.P)
	args := []gi.Argument{arg_v, arg_date2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_free
// container is not nil, container is Date
// is method
func (v Date) Free() {
	iv, err := _I.Get(101, "Date", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_date_get_day
// container is not nil, container is Date
// is method
func (v Date) GetDay() (result uint8) {
	iv, err := _I.Get(102, "Date", "get_day")
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

// g_date_get_day_of_year
// container is not nil, container is Date
// is method
func (v Date) GetDayOfYear() (result uint32) {
	iv, err := _I.Get(103, "Date", "get_day_of_year")
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

// g_date_get_iso8601_week_of_year
// container is not nil, container is Date
// is method
func (v Date) GetIso8601WeekOfYear() (result uint32) {
	iv, err := _I.Get(104, "Date", "get_iso8601_week_of_year")
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

// g_date_get_julian
// container is not nil, container is Date
// is method
func (v Date) GetJulian() (result uint32) {
	iv, err := _I.Get(105, "Date", "get_julian")
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

// g_date_get_monday_week_of_year
// container is not nil, container is Date
// is method
func (v Date) GetMondayWeekOfYear() (result uint32) {
	iv, err := _I.Get(106, "Date", "get_monday_week_of_year")
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

// g_date_get_month
// container is not nil, container is Date
// is method
func (v Date) GetMonth() (result DateMonthEnum) {
	iv, err := _I.Get(107, "Date", "get_month")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = DateMonthEnum(ret.Int())
	return
}

// g_date_get_sunday_week_of_year
// container is not nil, container is Date
// is method
func (v Date) GetSundayWeekOfYear() (result uint32) {
	iv, err := _I.Get(108, "Date", "get_sunday_week_of_year")
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

// g_date_get_weekday
// container is not nil, container is Date
// is method
func (v Date) GetWeekday() (result DateWeekdayEnum) {
	iv, err := _I.Get(109, "Date", "get_weekday")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = DateWeekdayEnum(ret.Int())
	return
}

// g_date_get_year
// container is not nil, container is Date
// is method
func (v Date) GetYear() (result uint16) {
	iv, err := _I.Get(110, "Date", "get_year")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint16()
	return
}

// g_date_is_first_of_month
// container is not nil, container is Date
// is method
func (v Date) IsFirstOfMonth() (result bool) {
	iv, err := _I.Get(111, "Date", "is_first_of_month")
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

// g_date_is_last_of_month
// container is not nil, container is Date
// is method
func (v Date) IsLastOfMonth() (result bool) {
	iv, err := _I.Get(112, "Date", "is_last_of_month")
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

// g_date_order
// container is not nil, container is Date
// is method
func (v Date) Order(date2 Date) {
	iv, err := _I.Get(113, "Date", "order")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_date2 := gi.NewPointerArgument(date2.P)
	args := []gi.Argument{arg_v, arg_date2}
	iv.Call(args, nil, nil)
}

// g_date_set_day
// container is not nil, container is Date
// is method
func (v Date) SetDay(day uint8) {
	iv, err := _I.Get(114, "Date", "set_day")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_day := gi.NewUint8Argument(day)
	args := []gi.Argument{arg_v, arg_day}
	iv.Call(args, nil, nil)
}

// g_date_set_dmy
// container is not nil, container is Date
// is method
func (v Date) SetDmy(day uint8, month DateMonthEnum, y uint16) {
	iv, err := _I.Get(115, "Date", "set_dmy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_day := gi.NewUint8Argument(day)
	arg_month := gi.NewIntArgument(int(month))
	arg_y := gi.NewUint16Argument(y)
	args := []gi.Argument{arg_v, arg_day, arg_month, arg_y}
	iv.Call(args, nil, nil)
}

// g_date_set_julian
// container is not nil, container is Date
// is method
func (v Date) SetJulian(julian_date uint32) {
	iv, err := _I.Get(116, "Date", "set_julian")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_julian_date := gi.NewUint32Argument(julian_date)
	args := []gi.Argument{arg_v, arg_julian_date}
	iv.Call(args, nil, nil)
}

// g_date_set_month
// container is not nil, container is Date
// is method
func (v Date) SetMonth(month DateMonthEnum) {
	iv, err := _I.Get(117, "Date", "set_month")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_month := gi.NewIntArgument(int(month))
	args := []gi.Argument{arg_v, arg_month}
	iv.Call(args, nil, nil)
}

// g_date_set_parse
// container is not nil, container is Date
// is method
func (v Date) SetParse(str string) {
	iv, err := _I.Get(118, "Date", "set_parse")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_v := gi.NewPointerArgument(v.P)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_v, arg_str}
	iv.Call(args, nil, nil)
	gi.Free(c_str)
}

// g_date_set_time
// container is not nil, container is Date
// is method
func (v Date) SetTime(time_ int32) {
	iv, err := _I.Get(119, "Date", "set_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_time_ := gi.NewInt32Argument(time_)
	args := []gi.Argument{arg_v, arg_time_}
	iv.Call(args, nil, nil)
}

// g_date_set_time_t
// container is not nil, container is Date
// is method
func (v Date) SetTimeT(timet int64) {
	iv, err := _I.Get(120, "Date", "set_time_t")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timet := gi.NewInt64Argument(timet)
	args := []gi.Argument{arg_v, arg_timet}
	iv.Call(args, nil, nil)
}

// g_date_set_time_val
// container is not nil, container is Date
// is method
func (v Date) SetTimeVal(timeval TimeVal) {
	iv, err := _I.Get(121, "Date", "set_time_val")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timeval := gi.NewPointerArgument(timeval.P)
	args := []gi.Argument{arg_v, arg_timeval}
	iv.Call(args, nil, nil)
}

// g_date_set_year
// container is not nil, container is Date
// is method
func (v Date) SetYear(year uint16) {
	iv, err := _I.Get(122, "Date", "set_year")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_v, arg_year}
	iv.Call(args, nil, nil)
}

// g_date_subtract_days
// container is not nil, container is Date
// is method
func (v Date) SubtractDays(n_days uint32) {
	iv, err := _I.Get(123, "Date", "subtract_days")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_days := gi.NewUint32Argument(n_days)
	args := []gi.Argument{arg_v, arg_n_days}
	iv.Call(args, nil, nil)
}

// g_date_subtract_months
// container is not nil, container is Date
// is method
func (v Date) SubtractMonths(n_months uint32) {
	iv, err := _I.Get(124, "Date", "subtract_months")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_months := gi.NewUint32Argument(n_months)
	args := []gi.Argument{arg_v, arg_n_months}
	iv.Call(args, nil, nil)
}

// g_date_subtract_years
// container is not nil, container is Date
// is method
func (v Date) SubtractYears(n_years uint32) {
	iv, err := _I.Get(125, "Date", "subtract_years")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_years := gi.NewUint32Argument(n_years)
	args := []gi.Argument{arg_v, arg_n_years}
	iv.Call(args, nil, nil)
}

// g_date_to_struct_tm
// container is not nil, container is Date
// is method
func (v Date) ToStructTm(tm unsafe.Pointer) {
	iv, err := _I.Get(126, "Date", "to_struct_tm")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tm := gi.NewPointerArgument(tm)
	args := []gi.Argument{arg_v, arg_tm}
	iv.Call(args, nil, nil)
}

// g_date_valid
// container is not nil, container is Date
// is method
func (v Date) Valid() (result bool) {
	iv, err := _I.Get(127, "Date", "valid")
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

// g_date_get_days_in_month
// container is not nil, container is Date
// is method
// arg0Type tag: interface, isPtr: false
func DateGetDaysInMonth1(month DateMonthEnum, year uint16) (result uint8) {
	iv, err := _I.Get(128, "Date", "get_days_in_month")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_month := gi.NewIntArgument(int(month))
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_month, arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_date_get_monday_weeks_in_year
// container is not nil, container is Date
// is method
// arg0Type tag: guint16, isPtr: false
func DateGetMondayWeeksInYear1(year uint16) (result uint8) {
	iv, err := _I.Get(129, "Date", "get_monday_weeks_in_year")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_date_get_sunday_weeks_in_year
// container is not nil, container is Date
// is method
// arg0Type tag: guint16, isPtr: false
func DateGetSundayWeeksInYear1(year uint16) (result uint8) {
	iv, err := _I.Get(130, "Date", "get_sunday_weeks_in_year")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_date_is_leap_year
// container is not nil, container is Date
// is method
// arg0Type tag: guint16, isPtr: false
func DateIsLeapYear1(year uint16) (result bool) {
	iv, err := _I.Get(131, "Date", "is_leap_year")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_strftime
// container is not nil, container is Date
// is method
// arg0Type tag: utf8, isPtr: true
func DateStrftime1(s string, slen uint64, format string, date Date) (result uint64) {
	iv, err := _I.Get(132, "Date", "strftime")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_s := gi.CString(s)
	c_format := gi.CString(format)
	arg_s := gi.NewStringArgument(c_s)
	arg_slen := gi.NewUint64Argument(slen)
	arg_format := gi.NewStringArgument(c_format)
	arg_date := gi.NewPointerArgument(date.P)
	args := []gi.Argument{arg_s, arg_slen, arg_format, arg_date}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_s)
	gi.Free(c_format)
	result = ret.Uint64()
	return
}

// g_date_valid_day
// container is not nil, container is Date
// is method
// arg0Type tag: guint8, isPtr: false
func DateValidDay1(day uint8) (result bool) {
	iv, err := _I.Get(133, "Date", "valid_day")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_day := gi.NewUint8Argument(day)
	args := []gi.Argument{arg_day}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_dmy
// container is not nil, container is Date
// is method
// arg0Type tag: guint8, isPtr: false
func DateValidDmy1(day uint8, month DateMonthEnum, year uint16) (result bool) {
	iv, err := _I.Get(134, "Date", "valid_dmy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_day := gi.NewUint8Argument(day)
	arg_month := gi.NewIntArgument(int(month))
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_day, arg_month, arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_julian
// container is not nil, container is Date
// is method
// arg0Type tag: guint32, isPtr: false
func DateValidJulian1(julian_date uint32) (result bool) {
	iv, err := _I.Get(135, "Date", "valid_julian")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_julian_date := gi.NewUint32Argument(julian_date)
	args := []gi.Argument{arg_julian_date}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_month
// container is not nil, container is Date
// is method
// arg0Type tag: interface, isPtr: false
func DateValidMonth1(month DateMonthEnum) (result bool) {
	iv, err := _I.Get(136, "Date", "valid_month")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_month := gi.NewIntArgument(int(month))
	args := []gi.Argument{arg_month}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_weekday
// container is not nil, container is Date
// is method
// arg0Type tag: interface, isPtr: false
func DateValidWeekday1(weekday DateWeekdayEnum) (result bool) {
	iv, err := _I.Get(137, "Date", "valid_weekday")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_weekday := gi.NewIntArgument(int(weekday))
	args := []gi.Argument{arg_weekday}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_year
// container is not nil, container is Date
// is method
// arg0Type tag: guint16, isPtr: false
func DateValidYear1(year uint16) (result bool) {
	iv, err := _I.Get(138, "Date", "valid_year")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Enum DateDMY
type DateDMYEnum int

const (
	DateDMYDay   DateDMYEnum = 0
	DateDMYMonth DateDMYEnum = 1
	DateDMYYear  DateDMYEnum = 2
)

func DateDMYGetType() gi.GType {
	ret := _I.GetGType(13, "DateDMY")
	return ret
}

// Enum DateMonth
type DateMonthEnum int

const (
	DateMonthBadMonth  DateMonthEnum = 0
	DateMonthJanuary   DateMonthEnum = 1
	DateMonthFebruary  DateMonthEnum = 2
	DateMonthMarch     DateMonthEnum = 3
	DateMonthApril     DateMonthEnum = 4
	DateMonthMay       DateMonthEnum = 5
	DateMonthJune      DateMonthEnum = 6
	DateMonthJuly      DateMonthEnum = 7
	DateMonthAugust    DateMonthEnum = 8
	DateMonthSeptember DateMonthEnum = 9
	DateMonthOctober   DateMonthEnum = 10
	DateMonthNovember  DateMonthEnum = 11
	DateMonthDecember  DateMonthEnum = 12
)

func DateMonthGetType() gi.GType {
	ret := _I.GetGType(14, "DateMonth")
	return ret
}

// Struct DateTime
type DateTime struct {
	P unsafe.Pointer
}

func DateTimeGetType() gi.GType {
	ret := _I.GetGType(15, "DateTime")
	return ret
}

// g_date_time_new
// container is not nil, container is DateTime
// is constructor
func NewDateTime(tz TimeZone, year int32, month int32, day int32, hour int32, minute int32, seconds float64) (result DateTime) {
	iv, err := _I.Get(139, "DateTime", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_tz := gi.NewPointerArgument(tz.P)
	arg_year := gi.NewInt32Argument(year)
	arg_month := gi.NewInt32Argument(month)
	arg_day := gi.NewInt32Argument(day)
	arg_hour := gi.NewInt32Argument(hour)
	arg_minute := gi.NewInt32Argument(minute)
	arg_seconds := gi.NewDoubleArgument(seconds)
	args := []gi.Argument{arg_tz, arg_year, arg_month, arg_day, arg_hour, arg_minute, arg_seconds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_from_iso8601
// container is not nil, container is DateTime
// is constructor
func NewDateTimeFromIso8601(text string, default_tz TimeZone) (result DateTime) {
	iv, err := _I.Get(140, "DateTime", "new_from_iso8601")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_text := gi.NewStringArgument(c_text)
	arg_default_tz := gi.NewPointerArgument(default_tz.P)
	args := []gi.Argument{arg_text, arg_default_tz}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_text)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_from_timeval_local
// container is not nil, container is DateTime
// is constructor
func NewDateTimeFromTimevalLocal(tv TimeVal) (result DateTime) {
	iv, err := _I.Get(141, "DateTime", "new_from_timeval_local")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_tv := gi.NewPointerArgument(tv.P)
	args := []gi.Argument{arg_tv}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_from_timeval_utc
// container is not nil, container is DateTime
// is constructor
func NewDateTimeFromTimevalUtc(tv TimeVal) (result DateTime) {
	iv, err := _I.Get(142, "DateTime", "new_from_timeval_utc")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_tv := gi.NewPointerArgument(tv.P)
	args := []gi.Argument{arg_tv}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_from_unix_local
// container is not nil, container is DateTime
// is constructor
func NewDateTimeFromUnixLocal(t int64) (result DateTime) {
	iv, err := _I.Get(143, "DateTime", "new_from_unix_local")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_t := gi.NewInt64Argument(t)
	args := []gi.Argument{arg_t}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_from_unix_utc
// container is not nil, container is DateTime
// is constructor
func NewDateTimeFromUnixUtc(t int64) (result DateTime) {
	iv, err := _I.Get(144, "DateTime", "new_from_unix_utc")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_t := gi.NewInt64Argument(t)
	args := []gi.Argument{arg_t}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_local
// container is not nil, container is DateTime
// is constructor
func NewDateTimeLocal(year int32, month int32, day int32, hour int32, minute int32, seconds float64) (result DateTime) {
	iv, err := _I.Get(145, "DateTime", "new_local")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewInt32Argument(year)
	arg_month := gi.NewInt32Argument(month)
	arg_day := gi.NewInt32Argument(day)
	arg_hour := gi.NewInt32Argument(hour)
	arg_minute := gi.NewInt32Argument(minute)
	arg_seconds := gi.NewDoubleArgument(seconds)
	args := []gi.Argument{arg_year, arg_month, arg_day, arg_hour, arg_minute, arg_seconds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_now
// container is not nil, container is DateTime
// is constructor
func NewDateTimeNow(tz TimeZone) (result DateTime) {
	iv, err := _I.Get(146, "DateTime", "new_now")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_tz := gi.NewPointerArgument(tz.P)
	args := []gi.Argument{arg_tz}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_now_local
// container is not nil, container is DateTime
// is constructor
func NewDateTimeNowLocal() (result DateTime) {
	iv, err := _I.Get(147, "DateTime", "new_now_local")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_now_utc
// container is not nil, container is DateTime
// is constructor
func NewDateTimeNowUtc() (result DateTime) {
	iv, err := _I.Get(148, "DateTime", "new_now_utc")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_new_utc
// container is not nil, container is DateTime
// is constructor
func NewDateTimeUtc(year int32, month int32, day int32, hour int32, minute int32, seconds float64) (result DateTime) {
	iv, err := _I.Get(149, "DateTime", "new_utc")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewInt32Argument(year)
	arg_month := gi.NewInt32Argument(month)
	arg_day := gi.NewInt32Argument(day)
	arg_hour := gi.NewInt32Argument(hour)
	arg_minute := gi.NewInt32Argument(minute)
	arg_seconds := gi.NewDoubleArgument(seconds)
	args := []gi.Argument{arg_year, arg_month, arg_day, arg_hour, arg_minute, arg_seconds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add
// container is not nil, container is DateTime
// is method
func (v DateTime) Add(timespan int64) (result DateTime) {
	iv, err := _I.Get(150, "DateTime", "add")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timespan := gi.NewInt64Argument(timespan)
	args := []gi.Argument{arg_v, arg_timespan}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_days
// container is not nil, container is DateTime
// is method
func (v DateTime) AddDays(days int32) (result DateTime) {
	iv, err := _I.Get(151, "DateTime", "add_days")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_days := gi.NewInt32Argument(days)
	args := []gi.Argument{arg_v, arg_days}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_full
// container is not nil, container is DateTime
// is method
func (v DateTime) AddFull(years int32, months int32, days int32, hours int32, minutes int32, seconds float64) (result DateTime) {
	iv, err := _I.Get(152, "DateTime", "add_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_years := gi.NewInt32Argument(years)
	arg_months := gi.NewInt32Argument(months)
	arg_days := gi.NewInt32Argument(days)
	arg_hours := gi.NewInt32Argument(hours)
	arg_minutes := gi.NewInt32Argument(minutes)
	arg_seconds := gi.NewDoubleArgument(seconds)
	args := []gi.Argument{arg_v, arg_years, arg_months, arg_days, arg_hours, arg_minutes, arg_seconds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_hours
// container is not nil, container is DateTime
// is method
func (v DateTime) AddHours(hours int32) (result DateTime) {
	iv, err := _I.Get(153, "DateTime", "add_hours")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_hours := gi.NewInt32Argument(hours)
	args := []gi.Argument{arg_v, arg_hours}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_minutes
// container is not nil, container is DateTime
// is method
func (v DateTime) AddMinutes(minutes int32) (result DateTime) {
	iv, err := _I.Get(154, "DateTime", "add_minutes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_minutes := gi.NewInt32Argument(minutes)
	args := []gi.Argument{arg_v, arg_minutes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_months
// container is not nil, container is DateTime
// is method
func (v DateTime) AddMonths(months int32) (result DateTime) {
	iv, err := _I.Get(155, "DateTime", "add_months")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_months := gi.NewInt32Argument(months)
	args := []gi.Argument{arg_v, arg_months}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_seconds
// container is not nil, container is DateTime
// is method
func (v DateTime) AddSeconds(seconds float64) (result DateTime) {
	iv, err := _I.Get(156, "DateTime", "add_seconds")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_seconds := gi.NewDoubleArgument(seconds)
	args := []gi.Argument{arg_v, arg_seconds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_weeks
// container is not nil, container is DateTime
// is method
func (v DateTime) AddWeeks(weeks int32) (result DateTime) {
	iv, err := _I.Get(157, "DateTime", "add_weeks")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_weeks := gi.NewInt32Argument(weeks)
	args := []gi.Argument{arg_v, arg_weeks}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_add_years
// container is not nil, container is DateTime
// is method
func (v DateTime) AddYears(years int32) (result DateTime) {
	iv, err := _I.Get(158, "DateTime", "add_years")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_years := gi.NewInt32Argument(years)
	args := []gi.Argument{arg_v, arg_years}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_difference
// container is not nil, container is DateTime
// is method
func (v DateTime) Difference(begin DateTime) (result int64) {
	iv, err := _I.Get(159, "DateTime", "difference")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_begin := gi.NewPointerArgument(begin.P)
	args := []gi.Argument{arg_v, arg_begin}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_date_time_format
// container is not nil, container is DateTime
// is method
func (v DateTime) Format(format string) (result string) {
	iv, err := _I.Get(160, "DateTime", "format")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_format := gi.CString(format)
	arg_v := gi.NewPointerArgument(v.P)
	arg_format := gi.NewStringArgument(c_format)
	args := []gi.Argument{arg_v, arg_format}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_format)
	result = ret.String().Take()
	return
}

// g_date_time_get_day_of_month
// container is not nil, container is DateTime
// is method
func (v DateTime) GetDayOfMonth() (result int32) {
	iv, err := _I.Get(161, "DateTime", "get_day_of_month")
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

// g_date_time_get_day_of_week
// container is not nil, container is DateTime
// is method
func (v DateTime) GetDayOfWeek() (result int32) {
	iv, err := _I.Get(162, "DateTime", "get_day_of_week")
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

// g_date_time_get_day_of_year
// container is not nil, container is DateTime
// is method
func (v DateTime) GetDayOfYear() (result int32) {
	iv, err := _I.Get(163, "DateTime", "get_day_of_year")
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

// g_date_time_get_hour
// container is not nil, container is DateTime
// is method
func (v DateTime) GetHour() (result int32) {
	iv, err := _I.Get(164, "DateTime", "get_hour")
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

// g_date_time_get_microsecond
// container is not nil, container is DateTime
// is method
func (v DateTime) GetMicrosecond() (result int32) {
	iv, err := _I.Get(165, "DateTime", "get_microsecond")
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

// g_date_time_get_minute
// container is not nil, container is DateTime
// is method
func (v DateTime) GetMinute() (result int32) {
	iv, err := _I.Get(166, "DateTime", "get_minute")
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

// g_date_time_get_month
// container is not nil, container is DateTime
// is method
func (v DateTime) GetMonth() (result int32) {
	iv, err := _I.Get(167, "DateTime", "get_month")
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

// g_date_time_get_second
// container is not nil, container is DateTime
// is method
func (v DateTime) GetSecond() (result int32) {
	iv, err := _I.Get(168, "DateTime", "get_second")
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

// g_date_time_get_seconds
// container is not nil, container is DateTime
// is method
func (v DateTime) GetSeconds() (result float64) {
	iv, err := _I.Get(169, "DateTime", "get_seconds")
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

// g_date_time_get_timezone
// container is not nil, container is DateTime
// is method
func (v DateTime) GetTimezone() (result TimeZone) {
	iv, err := _I.Get(170, "DateTime", "get_timezone")
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

// g_date_time_get_timezone_abbreviation
// container is not nil, container is DateTime
// is method
func (v DateTime) GetTimezoneAbbreviation() (result string) {
	iv, err := _I.Get(171, "DateTime", "get_timezone_abbreviation")
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

// g_date_time_get_utc_offset
// container is not nil, container is DateTime
// is method
func (v DateTime) GetUtcOffset() (result int64) {
	iv, err := _I.Get(172, "DateTime", "get_utc_offset")
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

// g_date_time_get_week_numbering_year
// container is not nil, container is DateTime
// is method
func (v DateTime) GetWeekNumberingYear() (result int32) {
	iv, err := _I.Get(173, "DateTime", "get_week_numbering_year")
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

// g_date_time_get_week_of_year
// container is not nil, container is DateTime
// is method
func (v DateTime) GetWeekOfYear() (result int32) {
	iv, err := _I.Get(174, "DateTime", "get_week_of_year")
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

// g_date_time_get_year
// container is not nil, container is DateTime
// is method
func (v DateTime) GetYear() (result int32) {
	iv, err := _I.Get(175, "DateTime", "get_year")
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

// g_date_time_get_ymd
// container is not nil, container is DateTime
// is method
func (v DateTime) GetYmd() (year int32, month int32, day int32) {
	iv, err := _I.Get(176, "DateTime", "get_ymd")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_year := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_month := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_day := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_year, arg_month, arg_day}
	iv.Call(args, nil, &outArgs[0])
	year = outArgs[0].Int32()
	month = outArgs[1].Int32()
	day = outArgs[2].Int32()
	return
}

// g_date_time_is_daylight_savings
// container is not nil, container is DateTime
// is method
func (v DateTime) IsDaylightSavings() (result bool) {
	iv, err := _I.Get(177, "DateTime", "is_daylight_savings")
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

// g_date_time_ref
// container is not nil, container is DateTime
// is method
func (v DateTime) Ref() (result DateTime) {
	iv, err := _I.Get(178, "DateTime", "ref")
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

// g_date_time_to_local
// container is not nil, container is DateTime
// is method
func (v DateTime) ToLocal() (result DateTime) {
	iv, err := _I.Get(179, "DateTime", "to_local")
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

// g_date_time_to_timeval
// container is not nil, container is DateTime
// is method
func (v DateTime) ToTimeval(tv TimeVal) (result bool) {
	iv, err := _I.Get(180, "DateTime", "to_timeval")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tv := gi.NewPointerArgument(tv.P)
	args := []gi.Argument{arg_v, arg_tv}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_time_to_timezone
// container is not nil, container is DateTime
// is method
func (v DateTime) ToTimezone(tz TimeZone) (result DateTime) {
	iv, err := _I.Get(181, "DateTime", "to_timezone")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tz := gi.NewPointerArgument(tz.P)
	args := []gi.Argument{arg_v, arg_tz}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_date_time_to_unix
// container is not nil, container is DateTime
// is method
func (v DateTime) ToUnix() (result int64) {
	iv, err := _I.Get(182, "DateTime", "to_unix")
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

// g_date_time_to_utc
// container is not nil, container is DateTime
// is method
func (v DateTime) ToUtc() (result DateTime) {
	iv, err := _I.Get(183, "DateTime", "to_utc")
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

// g_date_time_unref
// container is not nil, container is DateTime
// is method
func (v DateTime) Unref() {
	iv, err := _I.Get(184, "DateTime", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_date_time_compare
// container is not nil, container is DateTime
// is method
// arg0Type tag: void, isPtr: true
func DateTimeCompare1(dt1 unsafe.Pointer, dt2 unsafe.Pointer) (result int32) {
	iv, err := _I.Get(185, "DateTime", "compare")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dt1 := gi.NewPointerArgument(dt1)
	arg_dt2 := gi.NewPointerArgument(dt2)
	args := []gi.Argument{arg_dt1, arg_dt2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_time_equal
// container is not nil, container is DateTime
// is method
// arg0Type tag: void, isPtr: true
func DateTimeEqual1(dt1 unsafe.Pointer, dt2 unsafe.Pointer) (result bool) {
	iv, err := _I.Get(186, "DateTime", "equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dt1 := gi.NewPointerArgument(dt1)
	arg_dt2 := gi.NewPointerArgument(dt2)
	args := []gi.Argument{arg_dt1, arg_dt2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_time_hash
// container is not nil, container is DateTime
// is method
// arg0Type tag: void, isPtr: true
func DateTimeHash1(datetime unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(187, "DateTime", "hash")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_datetime := gi.NewPointerArgument(datetime)
	args := []gi.Argument{arg_datetime}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// Enum DateWeekday
type DateWeekdayEnum int

const (
	DateWeekdayBadWeekday DateWeekdayEnum = 0
	DateWeekdayMonday     DateWeekdayEnum = 1
	DateWeekdayTuesday    DateWeekdayEnum = 2
	DateWeekdayWednesday  DateWeekdayEnum = 3
	DateWeekdayThursday   DateWeekdayEnum = 4
	DateWeekdayFriday     DateWeekdayEnum = 5
	DateWeekdaySaturday   DateWeekdayEnum = 6
	DateWeekdaySunday     DateWeekdayEnum = 7
)

func DateWeekdayGetType() gi.GType {
	ret := _I.GetGType(16, "DateWeekday")
	return ret
}

// Struct DebugKey
type DebugKey struct {
	P unsafe.Pointer
}

const SizeOfStructDebugKey = 16

func DebugKeyGetType() gi.GType {
	ret := _I.GetGType(17, "DebugKey")
	return ret
}

// ignore callback DestroyNotify
// Struct Dir
type Dir struct {
	P unsafe.Pointer
}

func DirGetType() gi.GType {
	ret := _I.GetGType(18, "Dir")
	return ret
}

// g_dir_close
// container is not nil, container is Dir
// is method
func (v Dir) Close() {
	iv, err := _I.Get(188, "Dir", "close")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_dir_read_name
// container is not nil, container is Dir
// is method
func (v Dir) ReadName() (result string) {
	iv, err := _I.Get(189, "Dir", "read_name")
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

// g_dir_rewind
// container is not nil, container is Dir
// is method
func (v Dir) Rewind() {
	iv, err := _I.Get(190, "Dir", "rewind")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_dir_make_tmp
// container is not nil, container is Dir
// is method
// arg0Type tag: filename, isPtr: true
func DirMakeTmp1(tmpl string) (result string, err error) {
	iv, err := _I.Get(191, "Dir", "make_tmp")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_tmpl := gi.CString(tmpl)
	arg_tmpl := gi.NewStringArgument(c_tmpl)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_tmpl, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tmpl)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// Union DoubleIEEE754
type DoubleIEEE754 struct {
	P unsafe.Pointer
}

const SizeOfUnionDoubleIEEE754 = 8

func DoubleIEEE754GetType() gi.GType {
	ret := _I.GetGType(19, "DoubleIEEE754")
	return ret
}

type DuplicateFuncStruct struct {
	F_data unsafe.Pointer
}

//export myDuplicateFunc
func myDuplicateFunc(data C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := DuplicateFuncStruct{
		F_data: unsafe.Pointer(data),
	}
	fn(args)
}

// ignore callback EqualFunc
// Struct Error
type Error struct {
	P unsafe.Pointer
}

const SizeOfStructError = 16

func ErrorGetType() gi.GType {
	ret := _I.GetGType(20, "Error")
	return ret
}

// g_error_new_literal
// container is not nil, container is Error
// is constructor
func NewErrorLiteral(domain uint32, code int32, message string) (result Error) {
	iv, err := _I.Get(192, "Error", "new_literal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_message := gi.CString(message)
	arg_domain := gi.NewUint32Argument(domain)
	arg_code := gi.NewInt32Argument(code)
	arg_message := gi.NewStringArgument(c_message)
	args := []gi.Argument{arg_domain, arg_code, arg_message}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_message)
	result.P = ret.Pointer()
	return
}

// g_error_copy
// container is not nil, container is Error
// is method
func (v Error) Copy() (result Error) {
	iv, err := _I.Get(193, "Error", "copy")
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

// g_error_free
// container is not nil, container is Error
// is method
func (v Error) Free() {
	iv, err := _I.Get(194, "Error", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_error_matches
// container is not nil, container is Error
// is method
func (v Error) Matches(domain uint32, code int32) (result bool) {
	iv, err := _I.Get(195, "Error", "matches")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_domain := gi.NewUint32Argument(domain)
	arg_code := gi.NewInt32Argument(code)
	args := []gi.Argument{arg_v, arg_domain, arg_code}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Enum ErrorType
type ErrorTypeEnum int

const (
	ErrorTypeUnknown           ErrorTypeEnum = 0
	ErrorTypeUnexpEof          ErrorTypeEnum = 1
	ErrorTypeUnexpEofInString  ErrorTypeEnum = 2
	ErrorTypeUnexpEofInComment ErrorTypeEnum = 3
	ErrorTypeNonDigitInConst   ErrorTypeEnum = 4
	ErrorTypeDigitRadix        ErrorTypeEnum = 5
	ErrorTypeFloatRadix        ErrorTypeEnum = 6
	ErrorTypeFloatMalformed    ErrorTypeEnum = 7
)

func ErrorTypeGetType() gi.GType {
	ret := _I.GetGType(21, "ErrorType")
	return ret
}

// Enum FileError
type FileErrorEnum int

const (
	FileErrorExist       FileErrorEnum = 0
	FileErrorIsdir       FileErrorEnum = 1
	FileErrorAcces       FileErrorEnum = 2
	FileErrorNametoolong FileErrorEnum = 3
	FileErrorNoent       FileErrorEnum = 4
	FileErrorNotdir      FileErrorEnum = 5
	FileErrorNxio        FileErrorEnum = 6
	FileErrorNodev       FileErrorEnum = 7
	FileErrorRofs        FileErrorEnum = 8
	FileErrorTxtbsy      FileErrorEnum = 9
	FileErrorFault       FileErrorEnum = 10
	FileErrorLoop        FileErrorEnum = 11
	FileErrorNospc       FileErrorEnum = 12
	FileErrorNomem       FileErrorEnum = 13
	FileErrorMfile       FileErrorEnum = 14
	FileErrorNfile       FileErrorEnum = 15
	FileErrorBadf        FileErrorEnum = 16
	FileErrorInval       FileErrorEnum = 17
	FileErrorPipe        FileErrorEnum = 18
	FileErrorAgain       FileErrorEnum = 19
	FileErrorIntr        FileErrorEnum = 20
	FileErrorIo          FileErrorEnum = 21
	FileErrorPerm        FileErrorEnum = 22
	FileErrorNosys       FileErrorEnum = 23
	FileErrorFailed      FileErrorEnum = 24
)

func FileErrorGetType() gi.GType {
	ret := _I.GetGType(22, "FileError")
	return ret
}

// Flags FileTest
type FileTestFlags int

const (
	FileTestIsRegular    FileTestFlags = 1
	FileTestIsSymlink    FileTestFlags = 2
	FileTestIsDir        FileTestFlags = 4
	FileTestIsExecutable FileTestFlags = 8
	FileTestExists       FileTestFlags = 16
)

func FileTestGetType() gi.GType {
	ret := _I.GetGType(23, "FileTest")
	return ret
}

// Union FloatIEEE754
type FloatIEEE754 struct {
	P unsafe.Pointer
}

const SizeOfUnionFloatIEEE754 = 4

func FloatIEEE754GetType() gi.GType {
	ret := _I.GetGType(24, "FloatIEEE754")
	return ret
}

// Flags FormatSizeFlags
type FormatSizeFlags int

const (
	FormatSizeFlagsDefault    FormatSizeFlags = 0
	FormatSizeFlagsLongFormat FormatSizeFlags = 1
	FormatSizeFlagsIecUnits   FormatSizeFlags = 2
	FormatSizeFlagsBits       FormatSizeFlags = 4
)

func FormatSizeFlagsGetType() gi.GType {
	ret := _I.GetGType(25, "FormatSizeFlags")
	return ret
}

// ignore callback FreeFunc
type FuncStruct struct {
	F_data unsafe.Pointer
}

//export myFunc
func myFunc(data C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := FuncStruct{
		F_data: unsafe.Pointer(data),
	}
	fn(args)
}

type HFuncStruct struct {
	F_key   unsafe.Pointer
	F_value unsafe.Pointer
}

//export myHFunc
func myHFunc(key C.gpointer, value C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := HFuncStruct{
		F_key:   unsafe.Pointer(key),
		F_value: unsafe.Pointer(value),
	}
	fn(args)
}

type HRFuncStruct struct {
	F_key   unsafe.Pointer
	F_value unsafe.Pointer
}

//export myHRFunc
func myHRFunc(key C.gpointer, value C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := HRFuncStruct{
		F_key:   unsafe.Pointer(key),
		F_value: unsafe.Pointer(value),
	}
	fn(args)
}

// ignore callback HashFunc
// Struct HashTable
type HashTable struct {
	P unsafe.Pointer
}

func HashTableGetType() gi.GType {
	ret := _I.GetGType(26, "HashTable")
	return ret
}

// g_hash_table_add
// container is not nil, container is HashTable
// is method
// arg0Type tag: ghash, isPtr: true
func HashTableAdd1(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(196, "HashTable", "add")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_contains
// container is not nil, container is HashTable
// is method
// arg0Type tag: ghash, isPtr: true
func HashTableContains1(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(197, "HashTable", "contains")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_destroy
// container is not nil, container is HashTable
// is method
// arg0Type tag: ghash, isPtr: true
func HashTableDestroy1(hash_table HashTable) {
	iv, err := _I.Get(198, "HashTable", "destroy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hash_table_insert
// container is not nil, container is HashTable
// is method
// arg0Type tag: ghash, isPtr: true
func HashTableInsert1(hash_table HashTable, key unsafe.Pointer, value unsafe.Pointer) (result bool) {
	iv, err := _I.Get(199, "HashTable", "insert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_hash_table, arg_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_lookup
// container is not nil, container is HashTable
// is method
// arg0Type tag: ghash, isPtr: true
func HashTableLookup1(hash_table HashTable, key unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(200, "HashTable", "lookup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_hash_table_lookup_extended
// container is not nil, container is HashTable
// is method
// arg0Type tag: ghash, isPtr: true
func HashTableLookupExtended1(hash_table HashTable, lookup_key unsafe.Pointer) (result bool, orig_key unsafe.Pointer, value unsafe.Pointer) {
	iv, err := _I.Get(201, "HashTable", "lookup_extended")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_lookup_key := gi.NewPointerArgument(lookup_key)
	arg_orig_key := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_hash_table, arg_lookup_key, arg_orig_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	orig_key = outArgs[0].Pointer()
	value = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// g_hash_table_remove
// container is not nil, container is HashTable
// is method
// arg0Type tag: ghash, isPtr: true
func HashTableRemove1(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(202, "HashTable", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_remove_all
// container is not nil, container is HashTable
// is method
// arg0Type tag: ghash, isPtr: true
func HashTableRemoveAll1(hash_table HashTable) {
	iv, err := _I.Get(203, "HashTable", "remove_all")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hash_table_replace
// container is not nil, container is HashTable
// is method
// arg0Type tag: ghash, isPtr: true
func HashTableReplace1(hash_table HashTable, key unsafe.Pointer, value unsafe.Pointer) (result bool) {
	iv, err := _I.Get(204, "HashTable", "replace")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_hash_table, arg_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_size
// container is not nil, container is HashTable
// is method
// arg0Type tag: ghash, isPtr: true
func HashTableSize1(hash_table HashTable) (result uint32) {
	iv, err := _I.Get(205, "HashTable", "size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_hash_table_steal
// container is not nil, container is HashTable
// is method
// arg0Type tag: ghash, isPtr: true
func HashTableSteal1(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(206, "HashTable", "steal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_steal_all
// container is not nil, container is HashTable
// is method
// arg0Type tag: ghash, isPtr: true
func HashTableStealAll1(hash_table HashTable) {
	iv, err := _I.Get(207, "HashTable", "steal_all")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hash_table_steal_extended
// container is not nil, container is HashTable
// is method
// arg0Type tag: ghash, isPtr: true
func HashTableStealExtended1(hash_table HashTable, lookup_key unsafe.Pointer) (result bool, stolen_key unsafe.Pointer, stolen_value unsafe.Pointer) {
	iv, err := _I.Get(208, "HashTable", "steal_extended")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_lookup_key := gi.NewPointerArgument(lookup_key)
	arg_stolen_key := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_stolen_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_hash_table, arg_lookup_key, arg_stolen_key, arg_stolen_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	stolen_key = outArgs[0].Pointer()
	stolen_value = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// g_hash_table_unref
// container is not nil, container is HashTable
// is method
// arg0Type tag: ghash, isPtr: true
func HashTableUnref1(hash_table HashTable) {
	iv, err := _I.Get(209, "HashTable", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// Struct HashTableIter
type HashTableIter struct {
	P unsafe.Pointer
}

const SizeOfStructHashTableIter = 40

func HashTableIterGetType() gi.GType {
	ret := _I.GetGType(27, "HashTableIter")
	return ret
}

// g_hash_table_iter_init
// container is not nil, container is HashTableIter
// is method
func (v HashTableIter) Init(hash_table HashTable) {
	iv, err := _I.Get(210, "HashTableIter", "init")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_v, arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hash_table_iter_next
// container is not nil, container is HashTableIter
// is method
func (v HashTableIter) Next() (result bool, key unsafe.Pointer, value unsafe.Pointer) {
	iv, err := _I.Get(211, "HashTableIter", "next")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	key = outArgs[0].Pointer()
	value = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// g_hash_table_iter_remove
// container is not nil, container is HashTableIter
// is method
func (v HashTableIter) Remove() {
	iv, err := _I.Get(212, "HashTableIter", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_hash_table_iter_replace
// container is not nil, container is HashTableIter
// is method
func (v HashTableIter) Replace(value unsafe.Pointer) {
	iv, err := _I.Get(213, "HashTableIter", "replace")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, nil)
}

// g_hash_table_iter_steal
// container is not nil, container is HashTableIter
// is method
func (v HashTableIter) Steal() {
	iv, err := _I.Get(214, "HashTableIter", "steal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct Hmac
type Hmac struct {
	P unsafe.Pointer
}

func HmacGetType() gi.GType {
	ret := _I.GetGType(28, "Hmac")
	return ret
}

// g_hmac_get_digest
// container is not nil, container is Hmac
// is method
// arg 0 buffer lenArgIdx 1
func (v Hmac) GetDigest(buffer gi.Uint8Array, digest_len int /*TODO:TYPE*/) {
	iv, err := _I.Get(215, "Hmac", "get_digest")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer := gi.NewPointerArgument(buffer.P)
	args := []gi.Argument{arg_v, arg_buffer}
	iv.Call(args, nil, &outArgs[0])
}

// g_hmac_get_string
// container is not nil, container is Hmac
// is method
func (v Hmac) GetString() (result string) {
	iv, err := _I.Get(216, "Hmac", "get_string")
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

// g_hmac_unref
// container is not nil, container is Hmac
// is method
func (v Hmac) Unref() {
	iv, err := _I.Get(217, "Hmac", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_hmac_update
// container is not nil, container is Hmac
// is method
// arg 0 data lenArgIdx 1
func (v Hmac) Update(data gi.Uint8Array, length int64) {
	iv, err := _I.Get(218, "Hmac", "update")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_v, arg_data, arg_length}
	iv.Call(args, nil, nil)
}

// Struct Hook
type Hook struct {
	P unsafe.Pointer
}

const SizeOfStructHook = 64

func HookGetType() gi.GType {
	ret := _I.GetGType(29, "Hook")
	return ret
}

// g_hook_compare_ids
// container is not nil, container is Hook
// is method
func (v Hook) CompareIds(sibling Hook) (result int32) {
	iv, err := _I.Get(219, "Hook", "compare_ids")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_sibling := gi.NewPointerArgument(sibling.P)
	args := []gi.Argument{arg_v, arg_sibling}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_hook_destroy
// container is not nil, container is Hook
// is method
// arg0Type tag: interface, isPtr: true
func HookDestroy1(hook_list HookList, hook_id uint64) (result bool) {
	iv, err := _I.Get(220, "Hook", "destroy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook_id := gi.NewUint64Argument(hook_id)
	args := []gi.Argument{arg_hook_list, arg_hook_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hook_destroy_link
// container is not nil, container is Hook
// is method
// arg0Type tag: interface, isPtr: true
func HookDestroyLink1(hook_list HookList, hook Hook) {
	iv, err := _I.Get(221, "Hook", "destroy_link")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_free
// container is not nil, container is Hook
// is method
// arg0Type tag: interface, isPtr: true
func HookFree1(hook_list HookList, hook Hook) {
	iv, err := _I.Get(222, "Hook", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_insert_before
// container is not nil, container is Hook
// is method
// arg0Type tag: interface, isPtr: true
func HookInsertBefore1(hook_list HookList, sibling Hook, hook Hook) {
	iv, err := _I.Get(223, "Hook", "insert_before")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_sibling := gi.NewPointerArgument(sibling.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_sibling, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_prepend
// container is not nil, container is Hook
// is method
// arg0Type tag: interface, isPtr: true
func HookPrepend1(hook_list HookList, hook Hook) {
	iv, err := _I.Get(224, "Hook", "prepend")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_unref
// container is not nil, container is Hook
// is method
// arg0Type tag: interface, isPtr: true
func HookUnref1(hook_list HookList, hook Hook) {
	iv, err := _I.Get(225, "Hook", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

// ignore callback HookCheckFunc
// ignore callback HookCheckMarshaller
// ignore callback HookCompareFunc
// ignore callback HookFinalizeFunc
// ignore callback HookFindFunc
// Flags HookFlagMask
type HookFlagMaskFlags int

const (
	HookFlagMaskActive HookFlagMaskFlags = 1
	HookFlagMaskInCall HookFlagMaskFlags = 2
	HookFlagMaskMask   HookFlagMaskFlags = 15
)

func HookFlagMaskGetType() gi.GType {
	ret := _I.GetGType(30, "HookFlagMask")
	return ret
}

// ignore callback HookFunc
// Struct HookList
type HookList struct {
	P unsafe.Pointer
}

const SizeOfStructHookList = 56

func HookListGetType() gi.GType {
	ret := _I.GetGType(31, "HookList")
	return ret
}

// g_hook_list_clear
// container is not nil, container is HookList
// is method
func (v HookList) Clear() {
	iv, err := _I.Get(226, "HookList", "clear")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_hook_list_init
// container is not nil, container is HookList
// is method
func (v HookList) Init(hook_size uint32) {
	iv, err := _I.Get(227, "HookList", "init")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_hook_size := gi.NewUint32Argument(hook_size)
	args := []gi.Argument{arg_v, arg_hook_size}
	iv.Call(args, nil, nil)
}

// g_hook_list_invoke
// container is not nil, container is HookList
// is method
func (v HookList) Invoke(may_recurse bool) {
	iv, err := _I.Get(228, "HookList", "invoke")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_may_recurse := gi.NewBoolArgument(may_recurse)
	args := []gi.Argument{arg_v, arg_may_recurse}
	iv.Call(args, nil, nil)
}

// g_hook_list_invoke_check
// container is not nil, container is HookList
// is method
func (v HookList) InvokeCheck(may_recurse bool) {
	iv, err := _I.Get(229, "HookList", "invoke_check")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_may_recurse := gi.NewBoolArgument(may_recurse)
	args := []gi.Argument{arg_v, arg_may_recurse}
	iv.Call(args, nil, nil)
}

// ignore callback HookMarshaller
// Struct IOChannel
type IOChannel struct {
	P unsafe.Pointer
}

const SizeOfStructIOChannel = 136

func IOChannelGetType() gi.GType {
	ret := _I.GetGType(32, "IOChannel")
	return ret
}

// g_io_channel_new_file
// container is not nil, container is IOChannel
// is constructor
func NewIOChannelFile(filename string, mode string) (result IOChannel, err error) {
	iv, err := _I.Get(230, "IOChannel", "new_file")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	c_mode := gi.CString(mode)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_mode := gi.NewStringArgument(c_mode)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_filename, arg_mode, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	gi.Free(c_mode)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// g_io_channel_unix_new
// container is not nil, container is IOChannel
// is constructor
func IOChannelUnixNew(fd int32) (result IOChannel) {
	iv, err := _I.Get(231, "IOChannel", "unix_new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_fd := gi.NewInt32Argument(fd)
	args := []gi.Argument{arg_fd}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_io_channel_close
// container is not nil, container is IOChannel
// is method
func (v IOChannel) Close() {
	iv, err := _I.Get(232, "IOChannel", "close")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_io_channel_flush
// container is not nil, container is IOChannel
// is method
func (v IOChannel) Flush() (result IOStatusEnum, err error) {
	iv, err := _I.Get(233, "IOChannel", "flush")
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
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_get_buffer_condition
// container is not nil, container is IOChannel
// is method
func (v IOChannel) GetBufferCondition() (result IOConditionFlags) {
	iv, err := _I.Get(234, "IOChannel", "get_buffer_condition")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = IOConditionFlags(ret.Int())
	return
}

// g_io_channel_get_buffer_size
// container is not nil, container is IOChannel
// is method
func (v IOChannel) GetBufferSize() (result uint64) {
	iv, err := _I.Get(235, "IOChannel", "get_buffer_size")
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

// g_io_channel_get_buffered
// container is not nil, container is IOChannel
// is method
func (v IOChannel) GetBuffered() (result bool) {
	iv, err := _I.Get(236, "IOChannel", "get_buffered")
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

// g_io_channel_get_close_on_unref
// container is not nil, container is IOChannel
// is method
func (v IOChannel) GetCloseOnUnref() (result bool) {
	iv, err := _I.Get(237, "IOChannel", "get_close_on_unref")
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

// g_io_channel_get_encoding
// container is not nil, container is IOChannel
// is method
func (v IOChannel) GetEncoding() (result string) {
	iv, err := _I.Get(238, "IOChannel", "get_encoding")
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

// g_io_channel_get_flags
// container is not nil, container is IOChannel
// is method
func (v IOChannel) GetFlags() (result IOFlags) {
	iv, err := _I.Get(239, "IOChannel", "get_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = IOFlags(ret.Int())
	return
}

// g_io_channel_get_line_term
// container is not nil, container is IOChannel
// is method
func (v IOChannel) GetLineTerm(length int32) (result string) {
	iv, err := _I.Get(240, "IOChannel", "get_line_term")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_v, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_io_channel_init
// container is not nil, container is IOChannel
// is method
func (v IOChannel) Init() {
	iv, err := _I.Get(241, "IOChannel", "init")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_io_channel_read
// container is not nil, container is IOChannel
// is method
func (v IOChannel) Read(buf string, count uint64, bytes_read uint64) (result IOErrorEnum) {
	iv, err := _I.Get(242, "IOChannel", "read")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_buf := gi.CString(buf)
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewStringArgument(c_buf)
	arg_count := gi.NewUint64Argument(count)
	arg_bytes_read := gi.NewUint64Argument(bytes_read)
	args := []gi.Argument{arg_v, arg_buf, arg_count, arg_bytes_read}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_buf)
	result = IOErrorEnum(ret.Int())
	return
}

// g_io_channel_read_chars
// container is not nil, container is IOChannel
// is method
// arg 0 buf lenArgIdx 1
func (v IOChannel) ReadChars(buf gi.Uint8Array, count uint64) (result IOStatusEnum, bytes_read uint64, err error) {
	iv, err := _I.Get(243, "IOChannel", "read_chars")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_count := gi.NewUint64Argument(count)
	arg_bytes_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_buf, arg_count, arg_bytes_read, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	bytes_read = outArgs[0].Uint64()
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_read_line
// container is not nil, container is IOChannel
// is method
func (v IOChannel) ReadLine() (result IOStatusEnum, str_return string, length uint64, terminator_pos uint64, err error) {
	iv, err := _I.Get(244, "IOChannel", "read_line")
	if err != nil {
		return
	}
	var outArgs [4]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_str_return := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_terminator_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_v, arg_str_return, arg_length, arg_terminator_pos, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[3].Pointer())
	str_return = outArgs[0].String().Take()
	length = outArgs[1].Uint64()
	terminator_pos = outArgs[2].Uint64()
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_read_line_string
// container is not nil, container is IOChannel
// is method
func (v IOChannel) ReadLineString(buffer String, terminator_pos uint64) (result IOStatusEnum, err error) {
	iv, err := _I.Get(245, "IOChannel", "read_line_string")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffer := gi.NewPointerArgument(buffer.P)
	arg_terminator_pos := gi.NewUint64Argument(terminator_pos)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_buffer, arg_terminator_pos, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_read_to_end
// container is not nil, container is IOChannel
// is method
// arg 0 str_return lenArgIdx 1
func (v IOChannel) ReadToEnd() (result IOStatusEnum, str_return gi.Uint8Array, err error) {
	iv, err := _I.Get(246, "IOChannel", "read_to_end")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_str_return := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_v, arg_str_return, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var length uint64
	_ = length
	err = gi.ToError(outArgs[2].Pointer())
	str_return.P = outArgs[0].Pointer()
	length = outArgs[1].Uint64()
	result = IOStatusEnum(ret.Int())
	str_return.Len = int(length)
	return
}

// g_io_channel_read_unichar
// container is not nil, container is IOChannel
// is method
func (v IOChannel) ReadUnichar() (result IOStatusEnum, thechar rune, err error) {
	iv, err := _I.Get(247, "IOChannel", "read_unichar")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_thechar := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_thechar, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	thechar = rune(outArgs[0].Uint32())
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_ref
// container is not nil, container is IOChannel
// is method
func (v IOChannel) Ref() (result IOChannel) {
	iv, err := _I.Get(248, "IOChannel", "ref")
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

// g_io_channel_seek
// container is not nil, container is IOChannel
// is method
func (v IOChannel) Seek(offset int64, type1 SeekTypeEnum) (result IOErrorEnum) {
	iv, err := _I.Get(249, "IOChannel", "seek")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewInt64Argument(offset)
	arg_type1 := gi.NewIntArgument(int(type1))
	args := []gi.Argument{arg_v, arg_offset, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = IOErrorEnum(ret.Int())
	return
}

// g_io_channel_seek_position
// container is not nil, container is IOChannel
// is method
func (v IOChannel) SeekPosition(offset int64, type1 SeekTypeEnum) (result IOStatusEnum, err error) {
	iv, err := _I.Get(250, "IOChannel", "seek_position")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_offset := gi.NewInt64Argument(offset)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_offset, arg_type1, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_set_buffer_size
// container is not nil, container is IOChannel
// is method
func (v IOChannel) SetBufferSize(size uint64) {
	iv, err := _I.Get(251, "IOChannel", "set_buffer_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_v, arg_size}
	iv.Call(args, nil, nil)
}

// g_io_channel_set_buffered
// container is not nil, container is IOChannel
// is method
func (v IOChannel) SetBuffered(buffered bool) {
	iv, err := _I.Get(252, "IOChannel", "set_buffered")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_buffered := gi.NewBoolArgument(buffered)
	args := []gi.Argument{arg_v, arg_buffered}
	iv.Call(args, nil, nil)
}

// g_io_channel_set_close_on_unref
// container is not nil, container is IOChannel
// is method
func (v IOChannel) SetCloseOnUnref(do_close bool) {
	iv, err := _I.Get(253, "IOChannel", "set_close_on_unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_do_close := gi.NewBoolArgument(do_close)
	args := []gi.Argument{arg_v, arg_do_close}
	iv.Call(args, nil, nil)
}

// g_io_channel_set_encoding
// container is not nil, container is IOChannel
// is method
func (v IOChannel) SetEncoding(encoding string) (result IOStatusEnum, err error) {
	iv, err := _I.Get(254, "IOChannel", "set_encoding")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_encoding := gi.CString(encoding)
	arg_v := gi.NewPointerArgument(v.P)
	arg_encoding := gi.NewStringArgument(c_encoding)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_encoding, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_encoding)
	err = gi.ToError(outArgs[0].Pointer())
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_set_flags
// container is not nil, container is IOChannel
// is method
func (v IOChannel) SetFlags(flags IOFlags) (result IOStatusEnum, err error) {
	iv, err := _I.Get(255, "IOChannel", "set_flags")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_set_line_term
// container is not nil, container is IOChannel
// is method
func (v IOChannel) SetLineTerm(line_term string, length int32) {
	iv, err := _I.Get(256, "IOChannel", "set_line_term")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_line_term := gi.CString(line_term)
	arg_v := gi.NewPointerArgument(v.P)
	arg_line_term := gi.NewStringArgument(c_line_term)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_v, arg_line_term, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_line_term)
}

// g_io_channel_shutdown
// container is not nil, container is IOChannel
// is method
func (v IOChannel) Shutdown(flush bool) (result IOStatusEnum, err error) {
	iv, err := _I.Get(257, "IOChannel", "shutdown")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_flush := gi.NewBoolArgument(flush)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_flush, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_unix_get_fd
// container is not nil, container is IOChannel
// is method
func (v IOChannel) UnixGetFd() (result int32) {
	iv, err := _I.Get(258, "IOChannel", "unix_get_fd")
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

// g_io_channel_unref
// container is not nil, container is IOChannel
// is method
func (v IOChannel) Unref() {
	iv, err := _I.Get(259, "IOChannel", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_io_channel_write
// container is not nil, container is IOChannel
// is method
func (v IOChannel) Write(buf string, count uint64, bytes_written uint64) (result IOErrorEnum) {
	iv, err := _I.Get(260, "IOChannel", "write")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_buf := gi.CString(buf)
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewStringArgument(c_buf)
	arg_count := gi.NewUint64Argument(count)
	arg_bytes_written := gi.NewUint64Argument(bytes_written)
	args := []gi.Argument{arg_v, arg_buf, arg_count, arg_bytes_written}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_buf)
	result = IOErrorEnum(ret.Int())
	return
}

// g_io_channel_write_chars
// container is not nil, container is IOChannel
// is method
func (v IOChannel) WriteChars(buf gi.Uint8Array, count int64) (result IOStatusEnum, bytes_written uint64, err error) {
	iv, err := _I.Get(261, "IOChannel", "write_chars")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_buf := gi.NewPointerArgument(buf.P)
	arg_count := gi.NewInt64Argument(count)
	arg_bytes_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_buf, arg_count, arg_bytes_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	bytes_written = outArgs[0].Uint64()
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_write_unichar
// container is not nil, container is IOChannel
// is method
func (v IOChannel) WriteUnichar(thechar rune) (result IOStatusEnum, err error) {
	iv, err := _I.Get(262, "IOChannel", "write_unichar")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_thechar := gi.NewUint32Argument(uint32(thechar))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_thechar, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = IOStatusEnum(ret.Int())
	return
}

// g_io_channel_error_from_errno
// container is not nil, container is IOChannel
// is method
// arg0Type tag: gint32, isPtr: false
func IOChannelErrorFromErrno1(en int32) (result IOChannelErrorEnum) {
	iv, err := _I.Get(263, "IOChannel", "error_from_errno")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_en := gi.NewInt32Argument(en)
	args := []gi.Argument{arg_en}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = IOChannelErrorEnum(ret.Int())
	return
}

// Enum IOChannelError
type IOChannelErrorEnum int

const (
	IOChannelErrorFbig     IOChannelErrorEnum = 0
	IOChannelErrorInval    IOChannelErrorEnum = 1
	IOChannelErrorIo       IOChannelErrorEnum = 2
	IOChannelErrorIsdir    IOChannelErrorEnum = 3
	IOChannelErrorNospc    IOChannelErrorEnum = 4
	IOChannelErrorNxio     IOChannelErrorEnum = 5
	IOChannelErrorOverflow IOChannelErrorEnum = 6
	IOChannelErrorPipe     IOChannelErrorEnum = 7
	IOChannelErrorFailed   IOChannelErrorEnum = 8
)

func IOChannelErrorGetType() gi.GType {
	ret := _I.GetGType(33, "IOChannelError")
	return ret
}

// Flags IOCondition
type IOConditionFlags int

const (
	IOConditionIn   IOConditionFlags = 1
	IOConditionOut  IOConditionFlags = 4
	IOConditionPri  IOConditionFlags = 2
	IOConditionErr  IOConditionFlags = 8
	IOConditionHup  IOConditionFlags = 16
	IOConditionNval IOConditionFlags = 32
)

func IOConditionGetType() gi.GType {
	ret := _I.GetGType(34, "IOCondition")
	return ret
}

// Enum IOError
type IOErrorEnum int

const (
	IOErrorNone    IOErrorEnum = 0
	IOErrorAgain   IOErrorEnum = 1
	IOErrorInval   IOErrorEnum = 2
	IOErrorUnknown IOErrorEnum = 3
)

func IOErrorGetType() gi.GType {
	ret := _I.GetGType(35, "IOError")
	return ret
}

// Flags IOFlags
type IOFlags int

const (
	IOFlagsAppend      IOFlags = 1
	IOFlagsNonblock    IOFlags = 2
	IOFlagsIsReadable  IOFlags = 4
	IOFlagsIsWritable  IOFlags = 8
	IOFlagsIsWriteable IOFlags = 8
	IOFlagsIsSeekable  IOFlags = 16
	IOFlagsMask        IOFlags = 31
	IOFlagsGetMask     IOFlags = 31
	IOFlagsSetMask     IOFlags = 3
)

func IOFlagsGetType() gi.GType {
	ret := _I.GetGType(36, "IOFlags")
	return ret
}

// ignore callback IOFunc
// Struct IOFuncs
type IOFuncs struct {
	P unsafe.Pointer
}

const SizeOfStructIOFuncs = 64

func IOFuncsGetType() gi.GType {
	ret := _I.GetGType(37, "IOFuncs")
	return ret
}

// Enum IOStatus
type IOStatusEnum int

const (
	IOStatusError  IOStatusEnum = 0
	IOStatusNormal IOStatusEnum = 1
	IOStatusEof    IOStatusEnum = 2
	IOStatusAgain  IOStatusEnum = 3
)

func IOStatusGetType() gi.GType {
	ret := _I.GetGType(38, "IOStatus")
	return ret
}

// Struct KeyFile
type KeyFile struct {
	P unsafe.Pointer
}

func KeyFileGetType() gi.GType {
	ret := _I.GetGType(39, "KeyFile")
	return ret
}

// g_key_file_new
// container is not nil, container is KeyFile
// is constructor
func NewKeyFile() (result KeyFile) {
	iv, err := _I.Get(265, "KeyFile", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_key_file_get_boolean
// container is not nil, container is KeyFile
// is method
func (v KeyFile) GetBoolean(group_name string, key string) (result bool, err error) {
	iv, err := _I.Get(266, "KeyFile", "get_boolean")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_get_boolean_list
// container is not nil, container is KeyFile
// is method
// ret lenArgIdx 2
func (v KeyFile) GetBooleanList(group_name string, key string) (result gi.BoolArray, err error) {
	iv, err := _I.Get(267, "KeyFile", "get_boolean_list")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.BoolArray{P: ret.Pointer(), Len: int(length)}
	return
}

// g_key_file_get_comment
// container is not nil, container is KeyFile
// is method
func (v KeyFile) GetComment(group_name string, key string) (result string, err error) {
	iv, err := _I.Get(268, "KeyFile", "get_comment")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_key_file_get_double
// container is not nil, container is KeyFile
// is method
func (v KeyFile) GetDouble(group_name string, key string) (result float64, err error) {
	iv, err := _I.Get(269, "KeyFile", "get_double")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Double()
	return
}

// g_key_file_get_double_list
// container is not nil, container is KeyFile
// is method
// ret lenArgIdx 2
func (v KeyFile) GetDoubleList(group_name string, key string) (result gi.DoubleArray, err error) {
	iv, err := _I.Get(270, "KeyFile", "get_double_list")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.DoubleArray{P: ret.Pointer(), Len: int(length)}
	return
}

// g_key_file_get_groups
// container is not nil, container is KeyFile
// is method
func (v KeyFile) GetGroups() (result gi.CStrArray, length uint64) {
	iv, err := _I.Get(271, "KeyFile", "get_groups")
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
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_key_file_get_int64
// container is not nil, container is KeyFile
// is method
func (v KeyFile) GetInt64(group_name string, key string) (result int64, err error) {
	iv, err := _I.Get(272, "KeyFile", "get_int64")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int64()
	return
}

// g_key_file_get_integer
// container is not nil, container is KeyFile
// is method
func (v KeyFile) GetInteger(group_name string, key string) (result int32, err error) {
	iv, err := _I.Get(273, "KeyFile", "get_integer")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Int32()
	return
}

// g_key_file_get_integer_list
// container is not nil, container is KeyFile
// is method
// ret lenArgIdx 2
func (v KeyFile) GetIntegerList(group_name string, key string) (result gi.Int32Array, err error) {
	iv, err := _I.Get(274, "KeyFile", "get_integer_list")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.Int32Array{P: ret.Pointer(), Len: int(length)}
	return
}

// g_key_file_get_keys
// container is not nil, container is KeyFile
// is method
func (v KeyFile) GetKeys(group_name string) (result gi.CStrArray, length uint64, err error) {
	iv, err := _I.Get(275, "KeyFile", "get_keys")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_group_name := gi.CString(group_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_group_name, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_key_file_get_locale_for_key
// container is not nil, container is KeyFile
// is method
func (v KeyFile) GetLocaleForKey(group_name string, key string, locale string) (result string) {
	iv, err := _I.Get(276, "KeyFile", "get_locale_for_key")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_locale := gi.CString(locale)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_locale := gi.NewStringArgument(c_locale)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_locale}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_locale)
	result = ret.String().Take()
	return
}

// g_key_file_get_locale_string
// container is not nil, container is KeyFile
// is method
func (v KeyFile) GetLocaleString(group_name string, key string, locale string) (result string, err error) {
	iv, err := _I.Get(277, "KeyFile", "get_locale_string")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_locale := gi.CString(locale)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_locale := gi.NewStringArgument(c_locale)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_locale, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_locale)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_key_file_get_locale_string_list
// container is not nil, container is KeyFile
// is method
// ret lenArgIdx 3
func (v KeyFile) GetLocaleStringList(group_name string, key string, locale string) (result gi.CStrArray, err error) {
	iv, err := _I.Get(278, "KeyFile", "get_locale_string_list")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_locale := gi.CString(locale)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_locale := gi.NewStringArgument(c_locale)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_locale, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_locale)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_key_file_get_start_group
// container is not nil, container is KeyFile
// is method
func (v KeyFile) GetStartGroup() (result string) {
	iv, err := _I.Get(279, "KeyFile", "get_start_group")
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

// g_key_file_get_string
// container is not nil, container is KeyFile
// is method
func (v KeyFile) GetString(group_name string, key string) (result string, err error) {
	iv, err := _I.Get(280, "KeyFile", "get_string")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_key_file_get_string_list
// container is not nil, container is KeyFile
// is method
// ret lenArgIdx 2
func (v KeyFile) GetStringList(group_name string, key string) (result gi.CStrArray, err error) {
	iv, err := _I.Get(281, "KeyFile", "get_string_list")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_key_file_get_uint64
// container is not nil, container is KeyFile
// is method
func (v KeyFile) GetUint64(group_name string, key string) (result uint64, err error) {
	iv, err := _I.Get(282, "KeyFile", "get_uint64")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Uint64()
	return
}

// g_key_file_get_value
// container is not nil, container is KeyFile
// is method
func (v KeyFile) GetValue(group_name string, key string) (result string, err error) {
	iv, err := _I.Get(283, "KeyFile", "get_value")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_key_file_has_group
// container is not nil, container is KeyFile
// is method
func (v KeyFile) HasGroup(group_name string) (result bool) {
	iv, err := _I.Get(284, "KeyFile", "has_group")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	args := []gi.Argument{arg_v, arg_group_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_group_name)
	result = ret.Bool()
	return
}

// g_key_file_load_from_bytes
// container is not nil, container is KeyFile
// is method
func (v KeyFile) LoadFromBytes(bytes Bytes, flags KeyFileFlags) (result bool, err error) {
	iv, err := _I.Get(285, "KeyFile", "load_from_bytes")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_bytes := gi.NewPointerArgument(bytes.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_bytes, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_load_from_data
// container is not nil, container is KeyFile
// is method
func (v KeyFile) LoadFromData(data string, length uint64, flags KeyFileFlags) (result bool, err error) {
	iv, err := _I.Get(286, "KeyFile", "load_from_data")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_data := gi.CString(data)
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewStringArgument(c_data)
	arg_length := gi.NewUint64Argument(length)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_data, arg_length, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_data)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_load_from_data_dirs
// container is not nil, container is KeyFile
// is method
func (v KeyFile) LoadFromDataDirs(file string, flags KeyFileFlags) (result bool, full_path string, err error) {
	iv, err := _I.Get(287, "KeyFile", "load_from_data_dirs")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_file := gi.CString(file)
	arg_v := gi.NewPointerArgument(v.P)
	arg_file := gi.NewStringArgument(c_file)
	arg_full_path := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_file, arg_full_path, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_file)
	err = gi.ToError(outArgs[1].Pointer())
	full_path = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// g_key_file_load_from_dirs
// container is not nil, container is KeyFile
// is method
func (v KeyFile) LoadFromDirs(file string, search_dirs gi.CStrArray, flags KeyFileFlags) (result bool, full_path string, err error) {
	iv, err := _I.Get(288, "KeyFile", "load_from_dirs")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_file := gi.CString(file)
	arg_v := gi.NewPointerArgument(v.P)
	arg_file := gi.NewStringArgument(c_file)
	arg_search_dirs := gi.NewPointerArgument(search_dirs.P)
	arg_full_path := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_file, arg_search_dirs, arg_full_path, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_file)
	err = gi.ToError(outArgs[1].Pointer())
	full_path = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// g_key_file_load_from_file
// container is not nil, container is KeyFile
// is method
func (v KeyFile) LoadFromFile(file string, flags KeyFileFlags) (result bool, err error) {
	iv, err := _I.Get(289, "KeyFile", "load_from_file")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_file := gi.CString(file)
	arg_v := gi.NewPointerArgument(v.P)
	arg_file := gi.NewStringArgument(c_file)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_file, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_file)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_remove_comment
// container is not nil, container is KeyFile
// is method
func (v KeyFile) RemoveComment(group_name string, key string) (result bool, err error) {
	iv, err := _I.Get(290, "KeyFile", "remove_comment")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_remove_group
// container is not nil, container is KeyFile
// is method
func (v KeyFile) RemoveGroup(group_name string) (result bool, err error) {
	iv, err := _I.Get(291, "KeyFile", "remove_group")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_remove_key
// container is not nil, container is KeyFile
// is method
func (v KeyFile) RemoveKey(group_name string, key string) (result bool, err error) {
	iv, err := _I.Get(292, "KeyFile", "remove_key")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_save_to_file
// container is not nil, container is KeyFile
// is method
func (v KeyFile) SaveToFile(filename string) (result bool, err error) {
	iv, err := _I.Get(293, "KeyFile", "save_to_file")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_v := gi.NewPointerArgument(v.P)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_filename, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_set_boolean
// container is not nil, container is KeyFile
// is method
func (v KeyFile) SetBoolean(group_name string, key string, value bool) {
	iv, err := _I.Get(294, "KeyFile", "set_boolean")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewBoolArgument(value)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_boolean_list
// container is not nil, container is KeyFile
// is method
// arg 2 list lenArgIdx 3
func (v KeyFile) SetBooleanList(group_name string, key string, list gi.BoolArray, length uint64) {
	iv, err := _I.Get(295, "KeyFile", "set_boolean_list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_list := gi.NewPointerArgument(list.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_list, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_comment
// container is not nil, container is KeyFile
// is method
func (v KeyFile) SetComment(group_name string, key string, comment string) (result bool, err error) {
	iv, err := _I.Get(296, "KeyFile", "set_comment")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_comment := gi.CString(comment)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_comment := gi.NewStringArgument(c_comment)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_comment, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_comment)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_key_file_set_double
// container is not nil, container is KeyFile
// is method
func (v KeyFile) SetDouble(group_name string, key string, value float64) {
	iv, err := _I.Get(297, "KeyFile", "set_double")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewDoubleArgument(value)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_double_list
// container is not nil, container is KeyFile
// is method
// arg 2 list lenArgIdx 3
func (v KeyFile) SetDoubleList(group_name string, key string, list gi.DoubleArray, length uint64) {
	iv, err := _I.Get(298, "KeyFile", "set_double_list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_list := gi.NewPointerArgument(list.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_list, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_int64
// container is not nil, container is KeyFile
// is method
func (v KeyFile) SetInt64(group_name string, key string, value int64) {
	iv, err := _I.Get(299, "KeyFile", "set_int64")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewInt64Argument(value)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_integer
// container is not nil, container is KeyFile
// is method
func (v KeyFile) SetInteger(group_name string, key string, value int32) {
	iv, err := _I.Get(300, "KeyFile", "set_integer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewInt32Argument(value)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_integer_list
// container is not nil, container is KeyFile
// is method
// arg 2 list lenArgIdx 3
func (v KeyFile) SetIntegerList(group_name string, key string, list gi.Int32Array, length uint64) {
	iv, err := _I.Get(301, "KeyFile", "set_integer_list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_list := gi.NewPointerArgument(list.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_list, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_list_separator
// container is not nil, container is KeyFile
// is method
func (v KeyFile) SetListSeparator(separator int8) {
	iv, err := _I.Get(302, "KeyFile", "set_list_separator")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_separator := gi.NewInt8Argument(separator)
	args := []gi.Argument{arg_v, arg_separator}
	iv.Call(args, nil, nil)
}

// g_key_file_set_locale_string
// container is not nil, container is KeyFile
// is method
func (v KeyFile) SetLocaleString(group_name string, key string, locale string, string string) {
	iv, err := _I.Get(303, "KeyFile", "set_locale_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_locale := gi.CString(locale)
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_locale := gi.NewStringArgument(c_locale)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_locale, arg_string}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_locale)
	gi.Free(c_string)
}

// g_key_file_set_locale_string_list
// container is not nil, container is KeyFile
// is method
// arg 3 list lenArgIdx 4
func (v KeyFile) SetLocaleStringList(group_name string, key string, locale string, list gi.CStrArray, length uint64) {
	iv, err := _I.Get(304, "KeyFile", "set_locale_string_list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_locale := gi.CString(locale)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_locale := gi.NewStringArgument(c_locale)
	arg_list := gi.NewPointerArgument(list.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_locale, arg_list, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_locale)
}

// g_key_file_set_string
// container is not nil, container is KeyFile
// is method
func (v KeyFile) SetString(group_name string, key string, string string) {
	iv, err := _I.Get(305, "KeyFile", "set_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_string}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_string)
}

// g_key_file_set_string_list
// container is not nil, container is KeyFile
// is method
// arg 2 list lenArgIdx 3
func (v KeyFile) SetStringList(group_name string, key string, list gi.CStrArray, length uint64) {
	iv, err := _I.Get(306, "KeyFile", "set_string_list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_list := gi.NewPointerArgument(list.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_list, arg_length}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_uint64
// container is not nil, container is KeyFile
// is method
func (v KeyFile) SetUint64(group_name string, key string, value uint64) {
	iv, err := _I.Get(307, "KeyFile", "set_uint64")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewUint64Argument(value)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
}

// g_key_file_set_value
// container is not nil, container is KeyFile
// is method
func (v KeyFile) SetValue(group_name string, key string, value string) {
	iv, err := _I.Get(308, "KeyFile", "set_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_group_name := gi.CString(group_name)
	c_key := gi.CString(key)
	c_value := gi.CString(value)
	arg_v := gi.NewPointerArgument(v.P)
	arg_group_name := gi.NewStringArgument(c_group_name)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewStringArgument(c_value)
	args := []gi.Argument{arg_v, arg_group_name, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_group_name)
	gi.Free(c_key)
	gi.Free(c_value)
}

// g_key_file_to_data
// container is not nil, container is KeyFile
// is method
func (v KeyFile) ToData() (result string, length uint64, err error) {
	iv, err := _I.Get(309, "KeyFile", "to_data")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	length = outArgs[0].Uint64()
	result = ret.String().Take()
	return
}

// g_key_file_unref
// container is not nil, container is KeyFile
// is method
func (v KeyFile) Unref() {
	iv, err := _I.Get(310, "KeyFile", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum KeyFileError
type KeyFileErrorEnum int

const (
	KeyFileErrorUnknownEncoding KeyFileErrorEnum = 0
	KeyFileErrorParse           KeyFileErrorEnum = 1
	KeyFileErrorNotFound        KeyFileErrorEnum = 2
	KeyFileErrorKeyNotFound     KeyFileErrorEnum = 3
	KeyFileErrorGroupNotFound   KeyFileErrorEnum = 4
	KeyFileErrorInvalidValue    KeyFileErrorEnum = 5
)

func KeyFileErrorGetType() gi.GType {
	ret := _I.GetGType(40, "KeyFileError")
	return ret
}

// Flags KeyFileFlags
type KeyFileFlags int

const (
	KeyFileFlagsNone             KeyFileFlags = 0
	KeyFileFlagsKeepComments     KeyFileFlags = 1
	KeyFileFlagsKeepTranslations KeyFileFlags = 2
)

func KeyFileFlagsGetType() gi.GType {
	ret := _I.GetGType(41, "KeyFileFlags")
	return ret
}

// Struct List
type List struct {
	P unsafe.Pointer
}

const SizeOfStructList = 24

func ListGetType() gi.GType {
	ret := _I.GetGType(42, "List")
	return ret
}

// Struct LogField
type LogField struct {
	P unsafe.Pointer
}

const SizeOfStructLogField = 24

func LogFieldGetType() gi.GType {
	ret := _I.GetGType(43, "LogField")
	return ret
}

type LogFuncStruct struct {
	F_log_domain string
	F_log_level  LogLevelFlags
	F_message    string
}

//export myLogFunc
func myLogFunc(log_domain *C.gchar, log_level C.GLogLevelFlags, message *C.gchar, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := LogFuncStruct{
		F_log_domain: gi.GoString(unsafe.Pointer(log_domain)),
		F_log_level:  LogLevelFlags(log_level),
		F_message:    gi.GoString(unsafe.Pointer(message)),
	}
	fn(args)
}

// Flags LogLevelFlags
type LogLevelFlags int

const (
	LogLevelFlagsFlagRecursion LogLevelFlags = 1
	LogLevelFlagsFlagFatal     LogLevelFlags = 2
	LogLevelFlagsLevelError    LogLevelFlags = 4
	LogLevelFlagsLevelCritical LogLevelFlags = 8
	LogLevelFlagsLevelWarning  LogLevelFlags = 16
	LogLevelFlagsLevelMessage  LogLevelFlags = 32
	LogLevelFlagsLevelInfo     LogLevelFlags = 64
	LogLevelFlagsLevelDebug    LogLevelFlags = 128
	LogLevelFlagsLevelMask     LogLevelFlags = -4
)

func LogLevelFlagsGetType() gi.GType {
	ret := _I.GetGType(44, "LogLevelFlags")
	return ret
}

type LogWriterFuncStruct struct {
	F_log_level LogLevelFlags
	F_fields    unsafe.Pointer /*TODO_CB tag: array, isPtr: true*/
	F_n_fields  uint64
}

//export myLogWriterFunc
func myLogWriterFunc(log_level C.GLogLevelFlags, fields C.gpointer, n_fields C.guint64, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := LogWriterFuncStruct{
		F_log_level: LogLevelFlags(log_level),
		F_fields:    unsafe.Pointer(fields),
		F_n_fields:  uint64(n_fields),
	}
	fn(args)
}

// Enum LogWriterOutput
type LogWriterOutputEnum int

const (
	LogWriterOutputHandled   LogWriterOutputEnum = 1
	LogWriterOutputUnhandled LogWriterOutputEnum = 0
)

func LogWriterOutputGetType() gi.GType {
	ret := _I.GetGType(45, "LogWriterOutput")
	return ret
}

// Struct MainContext
type MainContext struct {
	P unsafe.Pointer
}

func MainContextGetType() gi.GType {
	ret := _I.GetGType(46, "MainContext")
	return ret
}

// g_main_context_new
// container is not nil, container is MainContext
// is constructor
func NewMainContext() (result MainContext) {
	iv, err := _I.Get(312, "MainContext", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_context_acquire
// container is not nil, container is MainContext
// is method
func (v MainContext) Acquire() (result bool) {
	iv, err := _I.Get(313, "MainContext", "acquire")
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

// g_main_context_add_poll
// container is not nil, container is MainContext
// is method
func (v MainContext) AddPoll(fd PollFD, priority int32) {
	iv, err := _I.Get(314, "MainContext", "add_poll")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	arg_priority := gi.NewInt32Argument(priority)
	args := []gi.Argument{arg_v, arg_fd, arg_priority}
	iv.Call(args, nil, nil)
}

// g_main_context_check
// container is not nil, container is MainContext
// is method
// arg 1 fds lenArgIdx 2
func (v MainContext) Check(max_priority int32, fds unsafe.Pointer, n_fds int32) (result bool) {
	iv, err := _I.Get(315, "MainContext", "check")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_max_priority := gi.NewInt32Argument(max_priority)
	arg_fds := gi.NewPointerArgument(fds)
	arg_n_fds := gi.NewInt32Argument(n_fds)
	args := []gi.Argument{arg_v, arg_max_priority, arg_fds, arg_n_fds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_main_context_dispatch
// container is not nil, container is MainContext
// is method
func (v MainContext) Dispatch() {
	iv, err := _I.Get(316, "MainContext", "dispatch")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_main_context_find_source_by_funcs_user_data
// container is not nil, container is MainContext
// is method
func (v MainContext) FindSourceByFuncsUserData(funcs SourceFuncs, user_data unsafe.Pointer) (result Source) {
	iv, err := _I.Get(317, "MainContext", "find_source_by_funcs_user_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_funcs := gi.NewPointerArgument(funcs.P)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_funcs, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_context_find_source_by_id
// container is not nil, container is MainContext
// is method
func (v MainContext) FindSourceById(source_id uint32) (result Source) {
	iv, err := _I.Get(318, "MainContext", "find_source_by_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_source_id := gi.NewUint32Argument(source_id)
	args := []gi.Argument{arg_v, arg_source_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_context_find_source_by_user_data
// container is not nil, container is MainContext
// is method
func (v MainContext) FindSourceByUserData(user_data unsafe.Pointer) (result Source) {
	iv, err := _I.Get(319, "MainContext", "find_source_by_user_data")
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

// g_main_context_invoke_full
// container is not nil, container is MainContext
// is method
func (v MainContext) InvokeFull(priority int32, function int /*TODO_TYPE isPtr: false, tag: interface*/, data unsafe.Pointer, notify int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(320, "MainContext", "invoke_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_priority := gi.NewInt32Argument(priority)
	arg_function := gi.NewIntArgument(function) /*TODO*/
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewIntArgument(notify) /*TODO*/
	args := []gi.Argument{arg_v, arg_priority, arg_function, arg_data, arg_notify}
	iv.Call(args, nil, nil)
}

// g_main_context_is_owner
// container is not nil, container is MainContext
// is method
func (v MainContext) IsOwner() (result bool) {
	iv, err := _I.Get(321, "MainContext", "is_owner")
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

// g_main_context_iteration
// container is not nil, container is MainContext
// is method
func (v MainContext) Iteration(may_block bool) (result bool) {
	iv, err := _I.Get(322, "MainContext", "iteration")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_may_block := gi.NewBoolArgument(may_block)
	args := []gi.Argument{arg_v, arg_may_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_main_context_pending
// container is not nil, container is MainContext
// is method
func (v MainContext) Pending() (result bool) {
	iv, err := _I.Get(323, "MainContext", "pending")
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

// g_main_context_pop_thread_default
// container is not nil, container is MainContext
// is method
func (v MainContext) PopThreadDefault() {
	iv, err := _I.Get(324, "MainContext", "pop_thread_default")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_main_context_prepare
// container is not nil, container is MainContext
// is method
func (v MainContext) Prepare(priority int32) (result bool) {
	iv, err := _I.Get(325, "MainContext", "prepare")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_priority := gi.NewInt32Argument(priority)
	args := []gi.Argument{arg_v, arg_priority}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_main_context_push_thread_default
// container is not nil, container is MainContext
// is method
func (v MainContext) PushThreadDefault() {
	iv, err := _I.Get(326, "MainContext", "push_thread_default")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_main_context_query
// container is not nil, container is MainContext
// is method
// arg 2 fds lenArgIdx 3
func (v MainContext) Query(max_priority int32, fds unsafe.Pointer, n_fds int32) (result int32, timeout_ int32) {
	iv, err := _I.Get(327, "MainContext", "query")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_max_priority := gi.NewInt32Argument(max_priority)
	arg_timeout_ := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_fds := gi.NewPointerArgument(fds)
	arg_n_fds := gi.NewInt32Argument(n_fds)
	args := []gi.Argument{arg_v, arg_max_priority, arg_timeout_, arg_fds, arg_n_fds}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	timeout_ = outArgs[0].Int32()
	result = ret.Int32()
	return
}

// g_main_context_ref
// container is not nil, container is MainContext
// is method
func (v MainContext) Ref() (result MainContext) {
	iv, err := _I.Get(328, "MainContext", "ref")
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

// g_main_context_release
// container is not nil, container is MainContext
// is method
func (v MainContext) Release() {
	iv, err := _I.Get(329, "MainContext", "release")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_main_context_remove_poll
// container is not nil, container is MainContext
// is method
func (v MainContext) RemovePoll(fd PollFD) {
	iv, err := _I.Get(330, "MainContext", "remove_poll")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_fd}
	iv.Call(args, nil, nil)
}

// g_main_context_unref
// container is not nil, container is MainContext
// is method
func (v MainContext) Unref() {
	iv, err := _I.Get(331, "MainContext", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_main_context_wait
// container is not nil, container is MainContext
// is method
func (v MainContext) Wait(cond Cond, mutex Mutex) (result bool) {
	iv, err := _I.Get(332, "MainContext", "wait")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_cond := gi.NewPointerArgument(cond.P)
	arg_mutex := gi.NewPointerArgument(mutex.P)
	args := []gi.Argument{arg_v, arg_cond, arg_mutex}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_main_context_wakeup
// container is not nil, container is MainContext
// is method
func (v MainContext) Wakeup() {
	iv, err := _I.Get(333, "MainContext", "wakeup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct MainLoop
type MainLoop struct {
	P unsafe.Pointer
}

func MainLoopGetType() gi.GType {
	ret := _I.GetGType(47, "MainLoop")
	return ret
}

// g_main_loop_new
// container is not nil, container is MainLoop
// is constructor
func NewMainLoop(context MainContext, is_running bool) (result MainLoop) {
	iv, err := _I.Get(337, "MainLoop", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_context := gi.NewPointerArgument(context.P)
	arg_is_running := gi.NewBoolArgument(is_running)
	args := []gi.Argument{arg_context, arg_is_running}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_loop_get_context
// container is not nil, container is MainLoop
// is method
func (v MainLoop) GetContext() (result MainContext) {
	iv, err := _I.Get(338, "MainLoop", "get_context")
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

// g_main_loop_is_running
// container is not nil, container is MainLoop
// is method
func (v MainLoop) IsRunning() (result bool) {
	iv, err := _I.Get(339, "MainLoop", "is_running")
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

// g_main_loop_quit
// container is not nil, container is MainLoop
// is method
func (v MainLoop) Quit() {
	iv, err := _I.Get(340, "MainLoop", "quit")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_main_loop_ref
// container is not nil, container is MainLoop
// is method
func (v MainLoop) Ref() (result MainLoop) {
	iv, err := _I.Get(341, "MainLoop", "ref")
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

// g_main_loop_run
// container is not nil, container is MainLoop
// is method
func (v MainLoop) Run() {
	iv, err := _I.Get(342, "MainLoop", "run")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_main_loop_unref
// container is not nil, container is MainLoop
// is method
func (v MainLoop) Unref() {
	iv, err := _I.Get(343, "MainLoop", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct MappedFile
type MappedFile struct {
	P unsafe.Pointer
}

func MappedFileGetType() gi.GType {
	ret := _I.GetGType(48, "MappedFile")
	return ret
}

// g_mapped_file_new
// container is not nil, container is MappedFile
// is constructor
func NewMappedFile(filename string, writable bool) (result MappedFile, err error) {
	iv, err := _I.Get(344, "MappedFile", "new")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_writable := gi.NewBoolArgument(writable)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_filename, arg_writable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// g_mapped_file_new_from_fd
// container is not nil, container is MappedFile
// is constructor
func NewMappedFileFromFd(fd int32, writable bool) (result MappedFile, err error) {
	iv, err := _I.Get(345, "MappedFile", "new_from_fd")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_fd := gi.NewInt32Argument(fd)
	arg_writable := gi.NewBoolArgument(writable)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_fd, arg_writable, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// g_mapped_file_free
// container is not nil, container is MappedFile
// is method
func (v MappedFile) Free() {
	iv, err := _I.Get(346, "MappedFile", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_mapped_file_get_bytes
// container is not nil, container is MappedFile
// is method
func (v MappedFile) GetBytes() (result Bytes) {
	iv, err := _I.Get(347, "MappedFile", "get_bytes")
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

// g_mapped_file_get_contents
// container is not nil, container is MappedFile
// is method
func (v MappedFile) GetContents() (result string) {
	iv, err := _I.Get(348, "MappedFile", "get_contents")
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

// g_mapped_file_get_length
// container is not nil, container is MappedFile
// is method
func (v MappedFile) GetLength() (result uint64) {
	iv, err := _I.Get(349, "MappedFile", "get_length")
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

// g_mapped_file_ref
// container is not nil, container is MappedFile
// is method
func (v MappedFile) Ref() (result MappedFile) {
	iv, err := _I.Get(350, "MappedFile", "ref")
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

// g_mapped_file_unref
// container is not nil, container is MappedFile
// is method
func (v MappedFile) Unref() {
	iv, err := _I.Get(351, "MappedFile", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Flags MarkupCollectType
type MarkupCollectTypeFlags int

const (
	MarkupCollectTypeInvalid  MarkupCollectTypeFlags = 0
	MarkupCollectTypeString   MarkupCollectTypeFlags = 1
	MarkupCollectTypeStrdup   MarkupCollectTypeFlags = 2
	MarkupCollectTypeBoolean  MarkupCollectTypeFlags = 3
	MarkupCollectTypeTristate MarkupCollectTypeFlags = 4
	MarkupCollectTypeOptional MarkupCollectTypeFlags = 65536
)

func MarkupCollectTypeGetType() gi.GType {
	ret := _I.GetGType(49, "MarkupCollectType")
	return ret
}

// Enum MarkupError
type MarkupErrorEnum int

const (
	MarkupErrorBadUtf8          MarkupErrorEnum = 0
	MarkupErrorEmpty            MarkupErrorEnum = 1
	MarkupErrorParse            MarkupErrorEnum = 2
	MarkupErrorUnknownElement   MarkupErrorEnum = 3
	MarkupErrorUnknownAttribute MarkupErrorEnum = 4
	MarkupErrorInvalidContent   MarkupErrorEnum = 5
	MarkupErrorMissingAttribute MarkupErrorEnum = 6
)

func MarkupErrorGetType() gi.GType {
	ret := _I.GetGType(50, "MarkupError")
	return ret
}

// Struct MarkupParseContext
type MarkupParseContext struct {
	P unsafe.Pointer
}

func MarkupParseContextGetType() gi.GType {
	ret := _I.GetGType(51, "MarkupParseContext")
	return ret
}

// g_markup_parse_context_new
// container is not nil, container is MarkupParseContext
// is constructor
func NewMarkupParseContext(parser MarkupParser, flags MarkupParseFlags, user_data unsafe.Pointer, user_data_dnotify int /*TODO_TYPE isPtr: false, tag: interface*/) (result MarkupParseContext) {
	iv, err := _I.Get(352, "MarkupParseContext", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_parser := gi.NewPointerArgument(parser.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_user_data_dnotify := gi.NewIntArgument(user_data_dnotify) /*TODO*/
	args := []gi.Argument{arg_parser, arg_flags, arg_user_data, arg_user_data_dnotify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_markup_parse_context_end_parse
// container is not nil, container is MarkupParseContext
// is method
func (v MarkupParseContext) EndParse() (result bool, err error) {
	iv, err := _I.Get(353, "MarkupParseContext", "end_parse")
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

// g_markup_parse_context_free
// container is not nil, container is MarkupParseContext
// is method
func (v MarkupParseContext) Free() {
	iv, err := _I.Get(354, "MarkupParseContext", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_markup_parse_context_get_element
// container is not nil, container is MarkupParseContext
// is method
func (v MarkupParseContext) GetElement() (result string) {
	iv, err := _I.Get(355, "MarkupParseContext", "get_element")
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

// g_markup_parse_context_get_position
// container is not nil, container is MarkupParseContext
// is method
func (v MarkupParseContext) GetPosition(line_number int32, char_number int32) {
	iv, err := _I.Get(356, "MarkupParseContext", "get_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_line_number := gi.NewInt32Argument(line_number)
	arg_char_number := gi.NewInt32Argument(char_number)
	args := []gi.Argument{arg_v, arg_line_number, arg_char_number}
	iv.Call(args, nil, nil)
}

// g_markup_parse_context_get_user_data
// container is not nil, container is MarkupParseContext
// is method
func (v MarkupParseContext) GetUserData() (result unsafe.Pointer) {
	iv, err := _I.Get(357, "MarkupParseContext", "get_user_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_markup_parse_context_parse
// container is not nil, container is MarkupParseContext
// is method
func (v MarkupParseContext) Parse(text string, text_len int64) (result bool, err error) {
	iv, err := _I.Get(358, "MarkupParseContext", "parse")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_text_len := gi.NewInt64Argument(text_len)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_text, arg_text_len, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_text)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_markup_parse_context_pop
// container is not nil, container is MarkupParseContext
// is method
func (v MarkupParseContext) Pop() (result unsafe.Pointer) {
	iv, err := _I.Get(359, "MarkupParseContext", "pop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_markup_parse_context_push
// container is not nil, container is MarkupParseContext
// is method
func (v MarkupParseContext) Push(parser MarkupParser, user_data unsafe.Pointer) {
	iv, err := _I.Get(360, "MarkupParseContext", "push")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_parser := gi.NewPointerArgument(parser.P)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_parser, arg_user_data}
	iv.Call(args, nil, nil)
}

// g_markup_parse_context_ref
// container is not nil, container is MarkupParseContext
// is method
func (v MarkupParseContext) Ref() (result MarkupParseContext) {
	iv, err := _I.Get(361, "MarkupParseContext", "ref")
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

// g_markup_parse_context_unref
// container is not nil, container is MarkupParseContext
// is method
func (v MarkupParseContext) Unref() {
	iv, err := _I.Get(362, "MarkupParseContext", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Flags MarkupParseFlags
type MarkupParseFlags int

const (
	MarkupParseFlagsDoNotUseThisUnsupportedFlag MarkupParseFlags = 1
	MarkupParseFlagsTreatCdataAsText            MarkupParseFlags = 2
	MarkupParseFlagsPrefixErrorPosition         MarkupParseFlags = 4
	MarkupParseFlagsIgnoreQualified             MarkupParseFlags = 8
)

func MarkupParseFlagsGetType() gi.GType {
	ret := _I.GetGType(52, "MarkupParseFlags")
	return ret
}

// Struct MarkupParser
type MarkupParser struct {
	P unsafe.Pointer
}

const SizeOfStructMarkupParser = 40

func MarkupParserGetType() gi.GType {
	ret := _I.GetGType(53, "MarkupParser")
	return ret
}

// Struct MatchInfo
type MatchInfo struct {
	P unsafe.Pointer
}

func MatchInfoGetType() gi.GType {
	ret := _I.GetGType(54, "MatchInfo")
	return ret
}

// g_match_info_expand_references
// container is not nil, container is MatchInfo
// is method
func (v MatchInfo) ExpandReferences(string_to_expand string) (result string, err error) {
	iv, err := _I.Get(363, "MatchInfo", "expand_references")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_string_to_expand := gi.CString(string_to_expand)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string_to_expand := gi.NewStringArgument(c_string_to_expand)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_string_to_expand, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_string_to_expand)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_match_info_fetch
// container is not nil, container is MatchInfo
// is method
func (v MatchInfo) Fetch(match_num int32) (result string) {
	iv, err := _I.Get(364, "MatchInfo", "fetch")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_match_num := gi.NewInt32Argument(match_num)
	args := []gi.Argument{arg_v, arg_match_num}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_match_info_fetch_all
// container is not nil, container is MatchInfo
// is method
func (v MatchInfo) FetchAll() (result gi.CStrArray) {
	iv, err := _I.Get(365, "MatchInfo", "fetch_all")
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

// g_match_info_fetch_named
// container is not nil, container is MatchInfo
// is method
func (v MatchInfo) FetchNamed(name string) (result string) {
	iv, err := _I.Get(366, "MatchInfo", "fetch_named")
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

// g_match_info_fetch_named_pos
// container is not nil, container is MatchInfo
// is method
func (v MatchInfo) FetchNamedPos(name string) (result bool, start_pos int32, end_pos int32) {
	iv, err := _I.Get(367, "MatchInfo", "fetch_named_pos")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	arg_start_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_name, arg_start_pos, arg_end_pos}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_name)
	start_pos = outArgs[0].Int32()
	end_pos = outArgs[1].Int32()
	result = ret.Bool()
	return
}

// g_match_info_fetch_pos
// container is not nil, container is MatchInfo
// is method
func (v MatchInfo) FetchPos(match_num int32) (result bool, start_pos int32, end_pos int32) {
	iv, err := _I.Get(368, "MatchInfo", "fetch_pos")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_match_num := gi.NewInt32Argument(match_num)
	arg_start_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_end_pos := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_match_num, arg_start_pos, arg_end_pos}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	start_pos = outArgs[0].Int32()
	end_pos = outArgs[1].Int32()
	result = ret.Bool()
	return
}

// g_match_info_free
// container is not nil, container is MatchInfo
// is method
func (v MatchInfo) Free() {
	iv, err := _I.Get(369, "MatchInfo", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_match_info_get_match_count
// container is not nil, container is MatchInfo
// is method
func (v MatchInfo) GetMatchCount() (result int32) {
	iv, err := _I.Get(370, "MatchInfo", "get_match_count")
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

// g_match_info_get_regex
// container is not nil, container is MatchInfo
// is method
func (v MatchInfo) GetRegex() (result Regex) {
	iv, err := _I.Get(371, "MatchInfo", "get_regex")
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

// g_match_info_get_string
// container is not nil, container is MatchInfo
// is method
func (v MatchInfo) GetString() (result string) {
	iv, err := _I.Get(372, "MatchInfo", "get_string")
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

// g_match_info_is_partial_match
// container is not nil, container is MatchInfo
// is method
func (v MatchInfo) IsPartialMatch() (result bool) {
	iv, err := _I.Get(373, "MatchInfo", "is_partial_match")
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

// g_match_info_matches
// container is not nil, container is MatchInfo
// is method
func (v MatchInfo) Matches() (result bool) {
	iv, err := _I.Get(374, "MatchInfo", "matches")
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

// g_match_info_next
// container is not nil, container is MatchInfo
// is method
func (v MatchInfo) Next() (result bool, err error) {
	iv, err := _I.Get(375, "MatchInfo", "next")
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

// g_match_info_ref
// container is not nil, container is MatchInfo
// is method
func (v MatchInfo) Ref() (result MatchInfo) {
	iv, err := _I.Get(376, "MatchInfo", "ref")
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

// g_match_info_unref
// container is not nil, container is MatchInfo
// is method
func (v MatchInfo) Unref() {
	iv, err := _I.Get(377, "MatchInfo", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct MemVTable
type MemVTable struct {
	P unsafe.Pointer
}

const SizeOfStructMemVTable = 48

func MemVTableGetType() gi.GType {
	ret := _I.GetGType(55, "MemVTable")
	return ret
}

// Union Mutex
type Mutex struct {
	P unsafe.Pointer
}

const SizeOfUnionMutex = 8

func MutexGetType() gi.GType {
	ret := _I.GetGType(56, "Mutex")
	return ret
}

// g_mutex_clear
// container is not nil, container is Mutex
// is method
func (v Mutex) Clear() {
	iv, err := _I.Get(378, "Mutex", "clear")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_mutex_init
// container is not nil, container is Mutex
// is method
func (v Mutex) Init() {
	iv, err := _I.Get(379, "Mutex", "init")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_mutex_lock
// container is not nil, container is Mutex
// is method
func (v Mutex) Lock() {
	iv, err := _I.Get(380, "Mutex", "lock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_mutex_trylock
// container is not nil, container is Mutex
// is method
func (v Mutex) Trylock() (result bool) {
	iv, err := _I.Get(381, "Mutex", "trylock")
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

// g_mutex_unlock
// container is not nil, container is Mutex
// is method
func (v Mutex) Unlock() {
	iv, err := _I.Get(382, "Mutex", "unlock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct Node
type Node struct {
	P unsafe.Pointer
}

const SizeOfStructNode = 40

func NodeGetType() gi.GType {
	ret := _I.GetGType(57, "Node")
	return ret
}

// g_node_child_index
// container is not nil, container is Node
// is method
func (v Node) ChildIndex(data unsafe.Pointer) (result int32) {
	iv, err := _I.Get(383, "Node", "child_index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_node_child_position
// container is not nil, container is Node
// is method
func (v Node) ChildPosition(child Node) (result int32) {
	iv, err := _I.Get(384, "Node", "child_position")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_child := gi.NewPointerArgument(child.P)
	args := []gi.Argument{arg_v, arg_child}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_node_depth
// container is not nil, container is Node
// is method
func (v Node) Depth() (result uint32) {
	iv, err := _I.Get(385, "Node", "depth")
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

// g_node_destroy
// container is not nil, container is Node
// is method
func (v Node) Destroy() {
	iv, err := _I.Get(386, "Node", "destroy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_node_is_ancestor
// container is not nil, container is Node
// is method
func (v Node) IsAncestor(descendant Node) (result bool) {
	iv, err := _I.Get(387, "Node", "is_ancestor")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_descendant := gi.NewPointerArgument(descendant.P)
	args := []gi.Argument{arg_v, arg_descendant}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_node_max_height
// container is not nil, container is Node
// is method
func (v Node) MaxHeight() (result uint32) {
	iv, err := _I.Get(388, "Node", "max_height")
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

// g_node_n_children
// container is not nil, container is Node
// is method
func (v Node) NChildren() (result uint32) {
	iv, err := _I.Get(389, "Node", "n_children")
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

// g_node_n_nodes
// container is not nil, container is Node
// is method
func (v Node) NNodes(flags TraverseFlags) (result uint32) {
	iv, err := _I.Get(390, "Node", "n_nodes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_node_reverse_children
// container is not nil, container is Node
// is method
func (v Node) ReverseChildren() {
	iv, err := _I.Get(391, "Node", "reverse_children")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_node_unlink
// container is not nil, container is Node
// is method
func (v Node) Unlink() {
	iv, err := _I.Get(392, "Node", "unlink")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// ignore callback NodeForeachFunc
// ignore callback NodeTraverseFunc
// Enum NormalizeMode
type NormalizeModeEnum int

const (
	NormalizeModeDefault        NormalizeModeEnum = 0
	NormalizeModeNfd            NormalizeModeEnum = 0
	NormalizeModeDefaultCompose NormalizeModeEnum = 1
	NormalizeModeNfc            NormalizeModeEnum = 1
	NormalizeModeAll            NormalizeModeEnum = 2
	NormalizeModeNfkd           NormalizeModeEnum = 2
	NormalizeModeAllCompose     NormalizeModeEnum = 3
	NormalizeModeNfkc           NormalizeModeEnum = 3
)

func NormalizeModeGetType() gi.GType {
	ret := _I.GetGType(58, "NormalizeMode")
	return ret
}

// Enum NumberParserError
type NumberParserErrorEnum int

const (
	NumberParserErrorInvalid     NumberParserErrorEnum = 0
	NumberParserErrorOutOfBounds NumberParserErrorEnum = 1
)

func NumberParserErrorGetType() gi.GType {
	ret := _I.GetGType(59, "NumberParserError")
	return ret
}

// Struct Once
type Once struct {
	P unsafe.Pointer
}

const SizeOfStructOnce = 16

func OnceGetType() gi.GType {
	ret := _I.GetGType(60, "Once")
	return ret
}

// g_once_init_enter
// container is not nil, container is Once
// is method
// arg0Type tag: void, isPtr: true
func OnceInitEnter1(location unsafe.Pointer) (result bool) {
	iv, err := _I.Get(393, "Once", "init_enter")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_location := gi.NewPointerArgument(location)
	args := []gi.Argument{arg_location}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_once_init_leave
// container is not nil, container is Once
// is method
// arg0Type tag: void, isPtr: true
func OnceInitLeave1(location unsafe.Pointer, result uint64) {
	iv, err := _I.Get(394, "Once", "init_leave")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_location := gi.NewPointerArgument(location)
	arg_result := gi.NewUint64Argument(result)
	args := []gi.Argument{arg_location, arg_result}
	iv.Call(args, nil, nil)
}

// Enum OnceStatus
type OnceStatusEnum int

const (
	OnceStatusNotcalled OnceStatusEnum = 0
	OnceStatusProgress  OnceStatusEnum = 1
	OnceStatusReady     OnceStatusEnum = 2
)

func OnceStatusGetType() gi.GType {
	ret := _I.GetGType(61, "OnceStatus")
	return ret
}

// Enum OptionArg
type OptionArgEnum int

const (
	OptionArgNone          OptionArgEnum = 0
	OptionArgString        OptionArgEnum = 1
	OptionArgInt           OptionArgEnum = 2
	OptionArgCallback      OptionArgEnum = 3
	OptionArgFilename      OptionArgEnum = 4
	OptionArgStringArray   OptionArgEnum = 5
	OptionArgFilenameArray OptionArgEnum = 6
	OptionArgDouble        OptionArgEnum = 7
	OptionArgInt64         OptionArgEnum = 8
)

func OptionArgGetType() gi.GType {
	ret := _I.GetGType(62, "OptionArg")
	return ret
}

// ignore callback OptionArgFunc
// Struct OptionContext
type OptionContext struct {
	P unsafe.Pointer
}

func OptionContextGetType() gi.GType {
	ret := _I.GetGType(63, "OptionContext")
	return ret
}

// g_option_context_add_group
// container is not nil, container is OptionContext
// is method
func (v OptionContext) AddGroup(group OptionGroup) {
	iv, err := _I.Get(395, "OptionContext", "add_group")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_group := gi.NewPointerArgument(group.P)
	args := []gi.Argument{arg_v, arg_group}
	iv.Call(args, nil, nil)
}

// g_option_context_add_main_entries
// container is not nil, container is OptionContext
// is method
func (v OptionContext) AddMainEntries(entries OptionEntry, translation_domain string) {
	iv, err := _I.Get(396, "OptionContext", "add_main_entries")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_translation_domain := gi.CString(translation_domain)
	arg_v := gi.NewPointerArgument(v.P)
	arg_entries := gi.NewPointerArgument(entries.P)
	arg_translation_domain := gi.NewStringArgument(c_translation_domain)
	args := []gi.Argument{arg_v, arg_entries, arg_translation_domain}
	iv.Call(args, nil, nil)
	gi.Free(c_translation_domain)
}

// g_option_context_free
// container is not nil, container is OptionContext
// is method
func (v OptionContext) Free() {
	iv, err := _I.Get(397, "OptionContext", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_option_context_get_description
// container is not nil, container is OptionContext
// is method
func (v OptionContext) GetDescription() (result string) {
	iv, err := _I.Get(398, "OptionContext", "get_description")
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

// g_option_context_get_help
// container is not nil, container is OptionContext
// is method
func (v OptionContext) GetHelp(main_help bool, group OptionGroup) (result string) {
	iv, err := _I.Get(399, "OptionContext", "get_help")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_main_help := gi.NewBoolArgument(main_help)
	arg_group := gi.NewPointerArgument(group.P)
	args := []gi.Argument{arg_v, arg_main_help, arg_group}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_option_context_get_help_enabled
// container is not nil, container is OptionContext
// is method
func (v OptionContext) GetHelpEnabled() (result bool) {
	iv, err := _I.Get(400, "OptionContext", "get_help_enabled")
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

// g_option_context_get_ignore_unknown_options
// container is not nil, container is OptionContext
// is method
func (v OptionContext) GetIgnoreUnknownOptions() (result bool) {
	iv, err := _I.Get(401, "OptionContext", "get_ignore_unknown_options")
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

// g_option_context_get_main_group
// container is not nil, container is OptionContext
// is method
func (v OptionContext) GetMainGroup() (result OptionGroup) {
	iv, err := _I.Get(402, "OptionContext", "get_main_group")
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

// g_option_context_get_strict_posix
// container is not nil, container is OptionContext
// is method
func (v OptionContext) GetStrictPosix() (result bool) {
	iv, err := _I.Get(403, "OptionContext", "get_strict_posix")
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

// g_option_context_get_summary
// container is not nil, container is OptionContext
// is method
func (v OptionContext) GetSummary() (result string) {
	iv, err := _I.Get(404, "OptionContext", "get_summary")
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

// g_option_context_parse
// container is not nil, container is OptionContext
// is method
// arg 1 argv lenArgIdx 0
func (v OptionContext) Parse(argc int /*TODO:TYPE*/, argv int /*TODO:TYPE*/) (result bool, err error) {
	iv, err := _I.Get(405, "OptionContext", "parse")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_option_context_parse_strv
// container is not nil, container is OptionContext
// is method
func (v OptionContext) ParseStrv(arguments int /*TODO:TYPE*/) (result bool, err error) {
	iv, err := _I.Get(406, "OptionContext", "parse_strv")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_option_context_set_description
// container is not nil, container is OptionContext
// is method
func (v OptionContext) SetDescription(description string) {
	iv, err := _I.Get(407, "OptionContext", "set_description")
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

// g_option_context_set_help_enabled
// container is not nil, container is OptionContext
// is method
func (v OptionContext) SetHelpEnabled(help_enabled bool) {
	iv, err := _I.Get(408, "OptionContext", "set_help_enabled")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_help_enabled := gi.NewBoolArgument(help_enabled)
	args := []gi.Argument{arg_v, arg_help_enabled}
	iv.Call(args, nil, nil)
}

// g_option_context_set_ignore_unknown_options
// container is not nil, container is OptionContext
// is method
func (v OptionContext) SetIgnoreUnknownOptions(ignore_unknown bool) {
	iv, err := _I.Get(409, "OptionContext", "set_ignore_unknown_options")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ignore_unknown := gi.NewBoolArgument(ignore_unknown)
	args := []gi.Argument{arg_v, arg_ignore_unknown}
	iv.Call(args, nil, nil)
}

// g_option_context_set_main_group
// container is not nil, container is OptionContext
// is method
func (v OptionContext) SetMainGroup(group OptionGroup) {
	iv, err := _I.Get(410, "OptionContext", "set_main_group")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_group := gi.NewPointerArgument(group.P)
	args := []gi.Argument{arg_v, arg_group}
	iv.Call(args, nil, nil)
}

// g_option_context_set_strict_posix
// container is not nil, container is OptionContext
// is method
func (v OptionContext) SetStrictPosix(strict_posix bool) {
	iv, err := _I.Get(411, "OptionContext", "set_strict_posix")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_strict_posix := gi.NewBoolArgument(strict_posix)
	args := []gi.Argument{arg_v, arg_strict_posix}
	iv.Call(args, nil, nil)
}

// g_option_context_set_summary
// container is not nil, container is OptionContext
// is method
func (v OptionContext) SetSummary(summary string) {
	iv, err := _I.Get(412, "OptionContext", "set_summary")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_summary := gi.CString(summary)
	arg_v := gi.NewPointerArgument(v.P)
	arg_summary := gi.NewStringArgument(c_summary)
	args := []gi.Argument{arg_v, arg_summary}
	iv.Call(args, nil, nil)
	gi.Free(c_summary)
}

// g_option_context_set_translate_func
// container is not nil, container is OptionContext
// is method
func (v OptionContext) SetTranslateFunc(func1 int /*TODO_TYPE isPtr: false, tag: interface*/, data unsafe.Pointer, destroy_notify int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(413, "OptionContext", "set_translate_func")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewIntArgument(func1) /*TODO*/
	arg_data := gi.NewPointerArgument(data)
	arg_destroy_notify := gi.NewIntArgument(destroy_notify) /*TODO*/
	args := []gi.Argument{arg_v, arg_func1, arg_data, arg_destroy_notify}
	iv.Call(args, nil, nil)
}

// g_option_context_set_translation_domain
// container is not nil, container is OptionContext
// is method
func (v OptionContext) SetTranslationDomain(domain string) {
	iv, err := _I.Get(414, "OptionContext", "set_translation_domain")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	arg_v := gi.NewPointerArgument(v.P)
	arg_domain := gi.NewStringArgument(c_domain)
	args := []gi.Argument{arg_v, arg_domain}
	iv.Call(args, nil, nil)
	gi.Free(c_domain)
}

// Struct OptionEntry
type OptionEntry struct {
	P unsafe.Pointer
}

const SizeOfStructOptionEntry = 48

func OptionEntryGetType() gi.GType {
	ret := _I.GetGType(64, "OptionEntry")
	return ret
}

// Enum OptionError
type OptionErrorEnum int

const (
	OptionErrorUnknownOption OptionErrorEnum = 0
	OptionErrorBadValue      OptionErrorEnum = 1
	OptionErrorFailed        OptionErrorEnum = 2
)

func OptionErrorGetType() gi.GType {
	ret := _I.GetGType(65, "OptionError")
	return ret
}

// ignore callback OptionErrorFunc
// Flags OptionFlags
type OptionFlags int

const (
	OptionFlagsNone        OptionFlags = 0
	OptionFlagsHidden      OptionFlags = 1
	OptionFlagsInMain      OptionFlags = 2
	OptionFlagsReverse     OptionFlags = 4
	OptionFlagsNoArg       OptionFlags = 8
	OptionFlagsFilename    OptionFlags = 16
	OptionFlagsOptionalArg OptionFlags = 32
	OptionFlagsNoalias     OptionFlags = 64
)

func OptionFlagsGetType() gi.GType {
	ret := _I.GetGType(66, "OptionFlags")
	return ret
}

// Struct OptionGroup
type OptionGroup struct {
	P unsafe.Pointer
}

func OptionGroupGetType() gi.GType {
	ret := _I.GetGType(67, "OptionGroup")
	return ret
}

// g_option_group_new
// container is not nil, container is OptionGroup
// is constructor
func NewOptionGroup(name string, description string, help_description string, user_data unsafe.Pointer, destroy int /*TODO_TYPE isPtr: false, tag: interface*/) (result OptionGroup) {
	iv, err := _I.Get(415, "OptionGroup", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_description := gi.CString(description)
	c_help_description := gi.CString(help_description)
	arg_name := gi.NewStringArgument(c_name)
	arg_description := gi.NewStringArgument(c_description)
	arg_help_description := gi.NewStringArgument(c_help_description)
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_destroy := gi.NewIntArgument(destroy) /*TODO*/
	args := []gi.Argument{arg_name, arg_description, arg_help_description, arg_user_data, arg_destroy}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_description)
	gi.Free(c_help_description)
	result.P = ret.Pointer()
	return
}

// g_option_group_add_entries
// container is not nil, container is OptionGroup
// is method
func (v OptionGroup) AddEntries(entries OptionEntry) {
	iv, err := _I.Get(416, "OptionGroup", "add_entries")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_entries := gi.NewPointerArgument(entries.P)
	args := []gi.Argument{arg_v, arg_entries}
	iv.Call(args, nil, nil)
}

// g_option_group_free
// container is not nil, container is OptionGroup
// is method
func (v OptionGroup) Free() {
	iv, err := _I.Get(417, "OptionGroup", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_option_group_ref
// container is not nil, container is OptionGroup
// is method
func (v OptionGroup) Ref() (result OptionGroup) {
	iv, err := _I.Get(418, "OptionGroup", "ref")
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

// g_option_group_set_translate_func
// container is not nil, container is OptionGroup
// is method
func (v OptionGroup) SetTranslateFunc(func1 int /*TODO_TYPE isPtr: false, tag: interface*/, data unsafe.Pointer, destroy_notify int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(419, "OptionGroup", "set_translate_func")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewIntArgument(func1) /*TODO*/
	arg_data := gi.NewPointerArgument(data)
	arg_destroy_notify := gi.NewIntArgument(destroy_notify) /*TODO*/
	args := []gi.Argument{arg_v, arg_func1, arg_data, arg_destroy_notify}
	iv.Call(args, nil, nil)
}

// g_option_group_set_translation_domain
// container is not nil, container is OptionGroup
// is method
func (v OptionGroup) SetTranslationDomain(domain string) {
	iv, err := _I.Get(420, "OptionGroup", "set_translation_domain")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	arg_v := gi.NewPointerArgument(v.P)
	arg_domain := gi.NewStringArgument(c_domain)
	args := []gi.Argument{arg_v, arg_domain}
	iv.Call(args, nil, nil)
	gi.Free(c_domain)
}

// g_option_group_unref
// container is not nil, container is OptionGroup
// is method
func (v OptionGroup) Unref() {
	iv, err := _I.Get(421, "OptionGroup", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// ignore callback OptionParseFunc
// Struct PatternSpec
type PatternSpec struct {
	P unsafe.Pointer
}

func PatternSpecGetType() gi.GType {
	ret := _I.GetGType(68, "PatternSpec")
	return ret
}

// g_pattern_spec_equal
// container is not nil, container is PatternSpec
// is method
func (v PatternSpec) Equal(pspec2 PatternSpec) (result bool) {
	iv, err := _I.Get(422, "PatternSpec", "equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pspec2 := gi.NewPointerArgument(pspec2.P)
	args := []gi.Argument{arg_v, arg_pspec2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_pattern_spec_free
// container is not nil, container is PatternSpec
// is method
func (v PatternSpec) Free() {
	iv, err := _I.Get(423, "PatternSpec", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct PollFD
type PollFD struct {
	P unsafe.Pointer
}

const SizeOfStructPollFD = 8

func PollFDGetType() gi.GType {
	ret := _I.GetGType(69, "PollFD")
	return ret
}

// ignore callback PollFunc
// ignore callback PrintFunc
// Struct Private
type Private struct {
	P unsafe.Pointer
}

const SizeOfStructPrivate = 32

func PrivateGetType() gi.GType {
	ret := _I.GetGType(70, "Private")
	return ret
}

// g_private_get
// container is not nil, container is Private
// is method
func (v Private) Get() (result unsafe.Pointer) {
	iv, err := _I.Get(424, "Private", "get")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_private_replace
// container is not nil, container is Private
// is method
func (v Private) Replace(value unsafe.Pointer) {
	iv, err := _I.Get(425, "Private", "replace")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, nil)
}

// g_private_set
// container is not nil, container is Private
// is method
func (v Private) Set(value unsafe.Pointer) {
	iv, err := _I.Get(426, "Private", "set")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, nil)
}

// Struct PtrArray
type PtrArray struct {
	P unsafe.Pointer
}

const SizeOfStructPtrArray = 16

func PtrArrayGetType() gi.GType {
	ret := _I.GetGType(71, "PtrArray")
	return ret
}

// Struct Queue
type Queue struct {
	P unsafe.Pointer
}

const SizeOfStructQueue = 24

func QueueGetType() gi.GType {
	ret := _I.GetGType(72, "Queue")
	return ret
}

// g_queue_clear
// container is not nil, container is Queue
// is method
func (v Queue) Clear() {
	iv, err := _I.Get(427, "Queue", "clear")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_queue_free
// container is not nil, container is Queue
// is method
func (v Queue) Free() {
	iv, err := _I.Get(428, "Queue", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_queue_free_full
// container is not nil, container is Queue
// is method
func (v Queue) FreeFull(free_func int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(429, "Queue", "free_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_free_func := gi.NewIntArgument(free_func) /*TODO*/
	args := []gi.Argument{arg_v, arg_free_func}
	iv.Call(args, nil, nil)
}

// g_queue_get_length
// container is not nil, container is Queue
// is method
func (v Queue) GetLength() (result uint32) {
	iv, err := _I.Get(430, "Queue", "get_length")
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

// g_queue_index
// container is not nil, container is Queue
// is method
func (v Queue) Index(data unsafe.Pointer) (result int32) {
	iv, err := _I.Get(431, "Queue", "index")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_queue_init
// container is not nil, container is Queue
// is method
func (v Queue) Init() {
	iv, err := _I.Get(432, "Queue", "init")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_queue_is_empty
// container is not nil, container is Queue
// is method
func (v Queue) IsEmpty() (result bool) {
	iv, err := _I.Get(433, "Queue", "is_empty")
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

// g_queue_peek_head
// container is not nil, container is Queue
// is method
func (v Queue) PeekHead() (result unsafe.Pointer) {
	iv, err := _I.Get(434, "Queue", "peek_head")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_queue_peek_nth
// container is not nil, container is Queue
// is method
func (v Queue) PeekNth(n uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(435, "Queue", "peek_nth")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n := gi.NewUint32Argument(n)
	args := []gi.Argument{arg_v, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_queue_peek_tail
// container is not nil, container is Queue
// is method
func (v Queue) PeekTail() (result unsafe.Pointer) {
	iv, err := _I.Get(436, "Queue", "peek_tail")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_queue_pop_head
// container is not nil, container is Queue
// is method
func (v Queue) PopHead() (result unsafe.Pointer) {
	iv, err := _I.Get(437, "Queue", "pop_head")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_queue_pop_nth
// container is not nil, container is Queue
// is method
func (v Queue) PopNth(n uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(438, "Queue", "pop_nth")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n := gi.NewUint32Argument(n)
	args := []gi.Argument{arg_v, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_queue_pop_tail
// container is not nil, container is Queue
// is method
func (v Queue) PopTail() (result unsafe.Pointer) {
	iv, err := _I.Get(439, "Queue", "pop_tail")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_queue_push_head
// container is not nil, container is Queue
// is method
func (v Queue) PushHead(data unsafe.Pointer) {
	iv, err := _I.Get(440, "Queue", "push_head")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	iv.Call(args, nil, nil)
}

// g_queue_push_nth
// container is not nil, container is Queue
// is method
func (v Queue) PushNth(data unsafe.Pointer, n int32) {
	iv, err := _I.Get(441, "Queue", "push_nth")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	arg_n := gi.NewInt32Argument(n)
	args := []gi.Argument{arg_v, arg_data, arg_n}
	iv.Call(args, nil, nil)
}

// g_queue_push_tail
// container is not nil, container is Queue
// is method
func (v Queue) PushTail(data unsafe.Pointer) {
	iv, err := _I.Get(442, "Queue", "push_tail")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	iv.Call(args, nil, nil)
}

// g_queue_remove
// container is not nil, container is Queue
// is method
func (v Queue) Remove(data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(443, "Queue", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_queue_remove_all
// container is not nil, container is Queue
// is method
func (v Queue) RemoveAll(data unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(444, "Queue", "remove_all")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_queue_reverse
// container is not nil, container is Queue
// is method
func (v Queue) Reverse() {
	iv, err := _I.Get(445, "Queue", "reverse")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct RWLock
type RWLock struct {
	P unsafe.Pointer
}

const SizeOfStructRWLock = 16

func RWLockGetType() gi.GType {
	ret := _I.GetGType(73, "RWLock")
	return ret
}

// g_rw_lock_clear
// container is not nil, container is RWLock
// is method
func (v RWLock) Clear() {
	iv, err := _I.Get(446, "RWLock", "clear")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rw_lock_init
// container is not nil, container is RWLock
// is method
func (v RWLock) Init() {
	iv, err := _I.Get(447, "RWLock", "init")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rw_lock_reader_lock
// container is not nil, container is RWLock
// is method
func (v RWLock) ReaderLock() {
	iv, err := _I.Get(448, "RWLock", "reader_lock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rw_lock_reader_trylock
// container is not nil, container is RWLock
// is method
func (v RWLock) ReaderTrylock() (result bool) {
	iv, err := _I.Get(449, "RWLock", "reader_trylock")
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

// g_rw_lock_reader_unlock
// container is not nil, container is RWLock
// is method
func (v RWLock) ReaderUnlock() {
	iv, err := _I.Get(450, "RWLock", "reader_unlock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rw_lock_writer_lock
// container is not nil, container is RWLock
// is method
func (v RWLock) WriterLock() {
	iv, err := _I.Get(451, "RWLock", "writer_lock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rw_lock_writer_trylock
// container is not nil, container is RWLock
// is method
func (v RWLock) WriterTrylock() (result bool) {
	iv, err := _I.Get(452, "RWLock", "writer_trylock")
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

// g_rw_lock_writer_unlock
// container is not nil, container is RWLock
// is method
func (v RWLock) WriterUnlock() {
	iv, err := _I.Get(453, "RWLock", "writer_unlock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct Rand
type Rand struct {
	P unsafe.Pointer
}

func RandGetType() gi.GType {
	ret := _I.GetGType(74, "Rand")
	return ret
}

// g_rand_double
// container is not nil, container is Rand
// is method
func (v Rand) Double() (result float64) {
	iv, err := _I.Get(454, "Rand", "double")
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

// g_rand_double_range
// container is not nil, container is Rand
// is method
func (v Rand) DoubleRange(begin float64, end float64) (result float64) {
	iv, err := _I.Get(455, "Rand", "double_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_begin := gi.NewDoubleArgument(begin)
	arg_end := gi.NewDoubleArgument(end)
	args := []gi.Argument{arg_v, arg_begin, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// g_rand_free
// container is not nil, container is Rand
// is method
func (v Rand) Free() {
	iv, err := _I.Get(456, "Rand", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rand_int
// container is not nil, container is Rand
// is method
func (v Rand) Int() (result uint32) {
	iv, err := _I.Get(457, "Rand", "int")
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

// g_rand_int_range
// container is not nil, container is Rand
// is method
func (v Rand) IntRange(begin int32, end int32) (result int32) {
	iv, err := _I.Get(458, "Rand", "int_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_begin := gi.NewInt32Argument(begin)
	arg_end := gi.NewInt32Argument(end)
	args := []gi.Argument{arg_v, arg_begin, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_rand_set_seed
// container is not nil, container is Rand
// is method
func (v Rand) SetSeed(seed uint32) {
	iv, err := _I.Get(459, "Rand", "set_seed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_seed := gi.NewUint32Argument(seed)
	args := []gi.Argument{arg_v, arg_seed}
	iv.Call(args, nil, nil)
}

// g_rand_set_seed_array
// container is not nil, container is Rand
// is method
func (v Rand) SetSeedArray(seed uint32, seed_length uint32) {
	iv, err := _I.Get(460, "Rand", "set_seed_array")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_seed := gi.NewUint32Argument(seed)
	arg_seed_length := gi.NewUint32Argument(seed_length)
	args := []gi.Argument{arg_v, arg_seed, arg_seed_length}
	iv.Call(args, nil, nil)
}

// Struct RecMutex
type RecMutex struct {
	P unsafe.Pointer
}

const SizeOfStructRecMutex = 16

func RecMutexGetType() gi.GType {
	ret := _I.GetGType(75, "RecMutex")
	return ret
}

// g_rec_mutex_clear
// container is not nil, container is RecMutex
// is method
func (v RecMutex) Clear() {
	iv, err := _I.Get(461, "RecMutex", "clear")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rec_mutex_init
// container is not nil, container is RecMutex
// is method
func (v RecMutex) Init() {
	iv, err := _I.Get(462, "RecMutex", "init")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rec_mutex_lock
// container is not nil, container is RecMutex
// is method
func (v RecMutex) Lock() {
	iv, err := _I.Get(463, "RecMutex", "lock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_rec_mutex_trylock
// container is not nil, container is RecMutex
// is method
func (v RecMutex) Trylock() (result bool) {
	iv, err := _I.Get(464, "RecMutex", "trylock")
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

// g_rec_mutex_unlock
// container is not nil, container is RecMutex
// is method
func (v RecMutex) Unlock() {
	iv, err := _I.Get(465, "RecMutex", "unlock")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct Regex
type Regex struct {
	P unsafe.Pointer
}

func RegexGetType() gi.GType {
	ret := _I.GetGType(76, "Regex")
	return ret
}

// g_regex_new
// container is not nil, container is Regex
// is constructor
func NewRegex(pattern string, compile_options RegexCompileFlags, match_options RegexMatchFlags) (result Regex, err error) {
	iv, err := _I.Get(466, "Regex", "new")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_pattern := gi.CString(pattern)
	arg_pattern := gi.NewStringArgument(c_pattern)
	arg_compile_options := gi.NewIntArgument(int(compile_options))
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_pattern, arg_compile_options, arg_match_options, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_pattern)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// g_regex_get_capture_count
// container is not nil, container is Regex
// is method
func (v Regex) GetCaptureCount() (result int32) {
	iv, err := _I.Get(467, "Regex", "get_capture_count")
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

// g_regex_get_compile_flags
// container is not nil, container is Regex
// is method
func (v Regex) GetCompileFlags() (result RegexCompileFlags) {
	iv, err := _I.Get(468, "Regex", "get_compile_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = RegexCompileFlags(ret.Int())
	return
}

// g_regex_get_has_cr_or_lf
// container is not nil, container is Regex
// is method
func (v Regex) GetHasCrOrLf() (result bool) {
	iv, err := _I.Get(469, "Regex", "get_has_cr_or_lf")
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

// g_regex_get_match_flags
// container is not nil, container is Regex
// is method
func (v Regex) GetMatchFlags() (result RegexMatchFlags) {
	iv, err := _I.Get(470, "Regex", "get_match_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = RegexMatchFlags(ret.Int())
	return
}

// g_regex_get_max_backref
// container is not nil, container is Regex
// is method
func (v Regex) GetMaxBackref() (result int32) {
	iv, err := _I.Get(471, "Regex", "get_max_backref")
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

// g_regex_get_max_lookbehind
// container is not nil, container is Regex
// is method
func (v Regex) GetMaxLookbehind() (result int32) {
	iv, err := _I.Get(472, "Regex", "get_max_lookbehind")
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

// g_regex_get_pattern
// container is not nil, container is Regex
// is method
func (v Regex) GetPattern() (result string) {
	iv, err := _I.Get(473, "Regex", "get_pattern")
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

// g_regex_get_string_number
// container is not nil, container is Regex
// is method
func (v Regex) GetStringNumber(name string) (result int32) {
	iv, err := _I.Get(474, "Regex", "get_string_number")
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
	result = ret.Int32()
	return
}

// g_regex_match
// container is not nil, container is Regex
// is method
func (v Regex) Match(string string, match_options RegexMatchFlags) (result bool, match_info MatchInfo) {
	iv, err := _I.Get(475, "Regex", "match")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewStringArgument(c_string)
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_match_info := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_string, arg_match_options, arg_match_info}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_string)
	match_info.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// g_regex_match_all
// container is not nil, container is Regex
// is method
func (v Regex) MatchAll(string string, match_options RegexMatchFlags) (result bool, match_info MatchInfo) {
	iv, err := _I.Get(476, "Regex", "match_all")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewStringArgument(c_string)
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_match_info := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_string, arg_match_options, arg_match_info}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_string)
	match_info.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// g_regex_match_all_full
// container is not nil, container is Regex
// is method
// arg 0 string lenArgIdx 1
func (v Regex) MatchAllFull(string gi.CStrArray, string_len int64, start_position int32, match_options RegexMatchFlags) (result bool, match_info MatchInfo, err error) {
	iv, err := _I.Get(477, "Regex", "match_all_full")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewPointerArgument(string.P)
	arg_string_len := gi.NewInt64Argument(string_len)
	arg_start_position := gi.NewInt32Argument(start_position)
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_match_info := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_string, arg_string_len, arg_start_position, arg_match_options, arg_match_info, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	match_info.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// g_regex_match_full
// container is not nil, container is Regex
// is method
// arg 0 string lenArgIdx 1
func (v Regex) MatchFull(string gi.CStrArray, string_len int64, start_position int32, match_options RegexMatchFlags) (result bool, match_info MatchInfo, err error) {
	iv, err := _I.Get(478, "Regex", "match_full")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewPointerArgument(string.P)
	arg_string_len := gi.NewInt64Argument(string_len)
	arg_start_position := gi.NewInt32Argument(start_position)
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_match_info := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_v, arg_string, arg_string_len, arg_start_position, arg_match_options, arg_match_info, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[1].Pointer())
	match_info.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// g_regex_ref
// container is not nil, container is Regex
// is method
func (v Regex) Ref() (result Regex) {
	iv, err := _I.Get(479, "Regex", "ref")
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

// g_regex_replace
// container is not nil, container is Regex
// is method
// arg 0 string lenArgIdx 1
func (v Regex) Replace(string gi.CStrArray, string_len int64, start_position int32, replacement string, match_options RegexMatchFlags) (result string, err error) {
	iv, err := _I.Get(480, "Regex", "replace")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_replacement := gi.CString(replacement)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewPointerArgument(string.P)
	arg_string_len := gi.NewInt64Argument(string_len)
	arg_start_position := gi.NewInt32Argument(start_position)
	arg_replacement := gi.NewStringArgument(c_replacement)
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_string, arg_string_len, arg_start_position, arg_replacement, arg_match_options, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_replacement)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_regex_replace_literal
// container is not nil, container is Regex
// is method
// arg 0 string lenArgIdx 1
func (v Regex) ReplaceLiteral(string gi.CStrArray, string_len int64, start_position int32, replacement string, match_options RegexMatchFlags) (result string, err error) {
	iv, err := _I.Get(481, "Regex", "replace_literal")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_replacement := gi.CString(replacement)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewPointerArgument(string.P)
	arg_string_len := gi.NewInt64Argument(string_len)
	arg_start_position := gi.NewInt32Argument(start_position)
	arg_replacement := gi.NewStringArgument(c_replacement)
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_string, arg_string_len, arg_start_position, arg_replacement, arg_match_options, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_replacement)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_regex_split
// container is not nil, container is Regex
// is method
func (v Regex) Split(string string, match_options RegexMatchFlags) (result gi.CStrArray) {
	iv, err := _I.Get(482, "Regex", "split")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewStringArgument(c_string)
	arg_match_options := gi.NewIntArgument(int(match_options))
	args := []gi.Argument{arg_v, arg_string, arg_match_options}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_regex_split_full
// container is not nil, container is Regex
// is method
// arg 0 string lenArgIdx 1
func (v Regex) SplitFull(string gi.CStrArray, string_len int64, start_position int32, match_options RegexMatchFlags, max_tokens int32) (result gi.CStrArray, err error) {
	iv, err := _I.Get(483, "Regex", "split_full")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewPointerArgument(string.P)
	arg_string_len := gi.NewInt64Argument(string_len)
	arg_start_position := gi.NewInt32Argument(start_position)
	arg_match_options := gi.NewIntArgument(int(match_options))
	arg_max_tokens := gi.NewInt32Argument(max_tokens)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_string, arg_string_len, arg_start_position, arg_match_options, arg_max_tokens, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_regex_unref
// container is not nil, container is Regex
// is method
func (v Regex) Unref() {
	iv, err := _I.Get(484, "Regex", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_regex_check_replacement
// container is not nil, container is Regex
// is method
// arg0Type tag: utf8, isPtr: true
func RegexCheckReplacement1(replacement string) (result bool, has_references bool, err error) {
	iv, err := _I.Get(485, "Regex", "check_replacement")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_replacement := gi.CString(replacement)
	arg_replacement := gi.NewStringArgument(c_replacement)
	arg_has_references := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_replacement, arg_has_references, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_replacement)
	err = gi.ToError(outArgs[1].Pointer())
	has_references = outArgs[0].Bool()
	result = ret.Bool()
	return
}

// g_regex_escape_nul
// container is not nil, container is Regex
// is method
// arg0Type tag: utf8, isPtr: true
func RegexEscapeNul1(string string, length int32) (result string) {
	iv, err := _I.Get(487, "Regex", "escape_nul")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_string, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_regex_escape_string
// container is not nil, container is Regex
// is method
// arg0Type tag: array, isPtr: true
// arg 0 string lenArgIdx 1
func RegexEscapeString1(string gi.CStrArray, length int32) (result string) {
	iv, err := _I.Get(488, "Regex", "escape_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_string := gi.NewPointerArgument(string.P)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_string, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_regex_match_simple
// container is not nil, container is Regex
// is method
// arg0Type tag: utf8, isPtr: true
func RegexMatchSimple1(pattern string, string string, compile_options RegexCompileFlags, match_options RegexMatchFlags) (result bool) {
	iv, err := _I.Get(489, "Regex", "match_simple")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_pattern := gi.CString(pattern)
	c_string := gi.CString(string)
	arg_pattern := gi.NewStringArgument(c_pattern)
	arg_string := gi.NewStringArgument(c_string)
	arg_compile_options := gi.NewIntArgument(int(compile_options))
	arg_match_options := gi.NewIntArgument(int(match_options))
	args := []gi.Argument{arg_pattern, arg_string, arg_compile_options, arg_match_options}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_pattern)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_regex_split_simple
// container is not nil, container is Regex
// is method
// arg0Type tag: utf8, isPtr: true
func RegexSplitSimple1(pattern string, string string, compile_options RegexCompileFlags, match_options RegexMatchFlags) (result gi.CStrArray) {
	iv, err := _I.Get(490, "Regex", "split_simple")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_pattern := gi.CString(pattern)
	c_string := gi.CString(string)
	arg_pattern := gi.NewStringArgument(c_pattern)
	arg_string := gi.NewStringArgument(c_string)
	arg_compile_options := gi.NewIntArgument(int(compile_options))
	arg_match_options := gi.NewIntArgument(int(match_options))
	args := []gi.Argument{arg_pattern, arg_string, arg_compile_options, arg_match_options}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_pattern)
	gi.Free(c_string)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// Flags RegexCompileFlags
type RegexCompileFlags int

const (
	RegexCompileFlagsCaseless         RegexCompileFlags = 1
	RegexCompileFlagsMultiline        RegexCompileFlags = 2
	RegexCompileFlagsDotall           RegexCompileFlags = 4
	RegexCompileFlagsExtended         RegexCompileFlags = 8
	RegexCompileFlagsAnchored         RegexCompileFlags = 16
	RegexCompileFlagsDollarEndonly    RegexCompileFlags = 32
	RegexCompileFlagsUngreedy         RegexCompileFlags = 512
	RegexCompileFlagsRaw              RegexCompileFlags = 2048
	RegexCompileFlagsNoAutoCapture    RegexCompileFlags = 4096
	RegexCompileFlagsOptimize         RegexCompileFlags = 8192
	RegexCompileFlagsFirstline        RegexCompileFlags = 262144
	RegexCompileFlagsDupnames         RegexCompileFlags = 524288
	RegexCompileFlagsNewlineCr        RegexCompileFlags = 1048576
	RegexCompileFlagsNewlineLf        RegexCompileFlags = 2097152
	RegexCompileFlagsNewlineCrlf      RegexCompileFlags = 3145728
	RegexCompileFlagsNewlineAnycrlf   RegexCompileFlags = 5242880
	RegexCompileFlagsBsrAnycrlf       RegexCompileFlags = 8388608
	RegexCompileFlagsJavascriptCompat RegexCompileFlags = 33554432
)

func RegexCompileFlagsGetType() gi.GType {
	ret := _I.GetGType(77, "RegexCompileFlags")
	return ret
}

// Enum RegexError
type RegexErrorEnum int

const (
	RegexErrorCompile                                  RegexErrorEnum = 0
	RegexErrorOptimize                                 RegexErrorEnum = 1
	RegexErrorReplace                                  RegexErrorEnum = 2
	RegexErrorMatch                                    RegexErrorEnum = 3
	RegexErrorInternal                                 RegexErrorEnum = 4
	RegexErrorStrayBackslash                           RegexErrorEnum = 101
	RegexErrorMissingControlChar                       RegexErrorEnum = 102
	RegexErrorUnrecognizedEscape                       RegexErrorEnum = 103
	RegexErrorQuantifiersOutOfOrder                    RegexErrorEnum = 104
	RegexErrorQuantifierTooBig                         RegexErrorEnum = 105
	RegexErrorUnterminatedCharacterClass               RegexErrorEnum = 106
	RegexErrorInvalidEscapeInCharacterClass            RegexErrorEnum = 107
	RegexErrorRangeOutOfOrder                          RegexErrorEnum = 108
	RegexErrorNothingToRepeat                          RegexErrorEnum = 109
	RegexErrorUnrecognizedCharacter                    RegexErrorEnum = 112
	RegexErrorPosixNamedClassOutsideClass              RegexErrorEnum = 113
	RegexErrorUnmatchedParenthesis                     RegexErrorEnum = 114
	RegexErrorInexistentSubpatternReference            RegexErrorEnum = 115
	RegexErrorUnterminatedComment                      RegexErrorEnum = 118
	RegexErrorExpressionTooLarge                       RegexErrorEnum = 120
	RegexErrorMemoryError                              RegexErrorEnum = 121
	RegexErrorVariableLengthLookbehind                 RegexErrorEnum = 125
	RegexErrorMalformedCondition                       RegexErrorEnum = 126
	RegexErrorTooManyConditionalBranches               RegexErrorEnum = 127
	RegexErrorAssertionExpected                        RegexErrorEnum = 128
	RegexErrorUnknownPosixClassName                    RegexErrorEnum = 130
	RegexErrorPosixCollatingElementsNotSupported       RegexErrorEnum = 131
	RegexErrorHexCodeTooLarge                          RegexErrorEnum = 134
	RegexErrorInvalidCondition                         RegexErrorEnum = 135
	RegexErrorSingleByteMatchInLookbehind              RegexErrorEnum = 136
	RegexErrorInfiniteLoop                             RegexErrorEnum = 140
	RegexErrorMissingSubpatternNameTerminator          RegexErrorEnum = 142
	RegexErrorDuplicateSubpatternName                  RegexErrorEnum = 143
	RegexErrorMalformedProperty                        RegexErrorEnum = 146
	RegexErrorUnknownProperty                          RegexErrorEnum = 147
	RegexErrorSubpatternNameTooLong                    RegexErrorEnum = 148
	RegexErrorTooManySubpatterns                       RegexErrorEnum = 149
	RegexErrorInvalidOctalValue                        RegexErrorEnum = 151
	RegexErrorTooManyBranchesInDefine                  RegexErrorEnum = 154
	RegexErrorDefineRepetion                           RegexErrorEnum = 155
	RegexErrorInconsistentNewlineOptions               RegexErrorEnum = 156
	RegexErrorMissingBackReference                     RegexErrorEnum = 157
	RegexErrorInvalidRelativeReference                 RegexErrorEnum = 158
	RegexErrorBacktrackingControlVerbArgumentForbidden RegexErrorEnum = 159
	RegexErrorUnknownBacktrackingControlVerb           RegexErrorEnum = 160
	RegexErrorNumberTooBig                             RegexErrorEnum = 161
	RegexErrorMissingSubpatternName                    RegexErrorEnum = 162
	RegexErrorMissingDigit                             RegexErrorEnum = 163
	RegexErrorInvalidDataCharacter                     RegexErrorEnum = 164
	RegexErrorExtraSubpatternName                      RegexErrorEnum = 165
	RegexErrorBacktrackingControlVerbArgumentRequired  RegexErrorEnum = 166
	RegexErrorInvalidControlChar                       RegexErrorEnum = 168
	RegexErrorMissingName                              RegexErrorEnum = 169
	RegexErrorNotSupportedInClass                      RegexErrorEnum = 171
	RegexErrorTooManyForwardReferences                 RegexErrorEnum = 172
	RegexErrorNameTooLong                              RegexErrorEnum = 175
	RegexErrorCharacterValueTooLarge                   RegexErrorEnum = 176
)

func RegexErrorGetType() gi.GType {
	ret := _I.GetGType(78, "RegexError")
	return ret
}

type RegexEvalCallbackStruct struct {
	F_match_info unsafe.Pointer /*TODO_CB tag: interface, isPtr: true*/
	F_result     unsafe.Pointer /*TODO_CB tag: interface, isPtr: true*/
}

//export myRegexEvalCallback
func myRegexEvalCallback(match_info C.gpointer, result C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := RegexEvalCallbackStruct{
		F_match_info: unsafe.Pointer(match_info),
		F_result:     unsafe.Pointer(result),
	}
	fn(args)
}

// Flags RegexMatchFlags
type RegexMatchFlags int

const (
	RegexMatchFlagsAnchored        RegexMatchFlags = 16
	RegexMatchFlagsNotbol          RegexMatchFlags = 128
	RegexMatchFlagsNoteol          RegexMatchFlags = 256
	RegexMatchFlagsNotempty        RegexMatchFlags = 1024
	RegexMatchFlagsPartial         RegexMatchFlags = 32768
	RegexMatchFlagsNewlineCr       RegexMatchFlags = 1048576
	RegexMatchFlagsNewlineLf       RegexMatchFlags = 2097152
	RegexMatchFlagsNewlineCrlf     RegexMatchFlags = 3145728
	RegexMatchFlagsNewlineAny      RegexMatchFlags = 4194304
	RegexMatchFlagsNewlineAnycrlf  RegexMatchFlags = 5242880
	RegexMatchFlagsBsrAnycrlf      RegexMatchFlags = 8388608
	RegexMatchFlagsBsrAny          RegexMatchFlags = 16777216
	RegexMatchFlagsPartialSoft     RegexMatchFlags = 32768
	RegexMatchFlagsPartialHard     RegexMatchFlags = 134217728
	RegexMatchFlagsNotemptyAtstart RegexMatchFlags = 268435456
)

func RegexMatchFlagsGetType() gi.GType {
	ret := _I.GetGType(79, "RegexMatchFlags")
	return ret
}

// Struct SList
type SList struct {
	P unsafe.Pointer
}

const SizeOfStructSList = 16

func SListGetType() gi.GType {
	ret := _I.GetGType(80, "SList")
	return ret
}

// Struct Scanner
type Scanner struct {
	P unsafe.Pointer
}

const SizeOfStructScanner = 144

func ScannerGetType() gi.GType {
	ret := _I.GetGType(81, "Scanner")
	return ret
}

// g_scanner_cur_line
// container is not nil, container is Scanner
// is method
func (v Scanner) CurLine() (result uint32) {
	iv, err := _I.Get(491, "Scanner", "cur_line")
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

// g_scanner_cur_position
// container is not nil, container is Scanner
// is method
func (v Scanner) CurPosition() (result uint32) {
	iv, err := _I.Get(492, "Scanner", "cur_position")
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

// g_scanner_cur_token
// container is not nil, container is Scanner
// is method
func (v Scanner) CurToken() (result TokenTypeEnum) {
	iv, err := _I.Get(493, "Scanner", "cur_token")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TokenTypeEnum(ret.Int())
	return
}

// g_scanner_destroy
// container is not nil, container is Scanner
// is method
func (v Scanner) Destroy() {
	iv, err := _I.Get(494, "Scanner", "destroy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_scanner_eof
// container is not nil, container is Scanner
// is method
func (v Scanner) Eof() (result bool) {
	iv, err := _I.Get(495, "Scanner", "eof")
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

// g_scanner_get_next_token
// container is not nil, container is Scanner
// is method
func (v Scanner) GetNextToken() (result TokenTypeEnum) {
	iv, err := _I.Get(496, "Scanner", "get_next_token")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TokenTypeEnum(ret.Int())
	return
}

// g_scanner_input_file
// container is not nil, container is Scanner
// is method
func (v Scanner) InputFile(input_fd int32) {
	iv, err := _I.Get(497, "Scanner", "input_file")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_input_fd := gi.NewInt32Argument(input_fd)
	args := []gi.Argument{arg_v, arg_input_fd}
	iv.Call(args, nil, nil)
}

// g_scanner_input_text
// container is not nil, container is Scanner
// is method
func (v Scanner) InputText(text string, text_len uint32) {
	iv, err := _I.Get(498, "Scanner", "input_text")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_v := gi.NewPointerArgument(v.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_text_len := gi.NewUint32Argument(text_len)
	args := []gi.Argument{arg_v, arg_text, arg_text_len}
	iv.Call(args, nil, nil)
	gi.Free(c_text)
}

// g_scanner_lookup_symbol
// container is not nil, container is Scanner
// is method
func (v Scanner) LookupSymbol(symbol string) (result unsafe.Pointer) {
	iv, err := _I.Get(499, "Scanner", "lookup_symbol")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_symbol := gi.CString(symbol)
	arg_v := gi.NewPointerArgument(v.P)
	arg_symbol := gi.NewStringArgument(c_symbol)
	args := []gi.Argument{arg_v, arg_symbol}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_symbol)
	result = ret.Pointer()
	return
}

// g_scanner_peek_next_token
// container is not nil, container is Scanner
// is method
func (v Scanner) PeekNextToken() (result TokenTypeEnum) {
	iv, err := _I.Get(500, "Scanner", "peek_next_token")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = TokenTypeEnum(ret.Int())
	return
}

// g_scanner_scope_add_symbol
// container is not nil, container is Scanner
// is method
func (v Scanner) ScopeAddSymbol(scope_id uint32, symbol string, value unsafe.Pointer) {
	iv, err := _I.Get(501, "Scanner", "scope_add_symbol")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_symbol := gi.CString(symbol)
	arg_v := gi.NewPointerArgument(v.P)
	arg_scope_id := gi.NewUint32Argument(scope_id)
	arg_symbol := gi.NewStringArgument(c_symbol)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_v, arg_scope_id, arg_symbol, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_symbol)
}

// g_scanner_scope_lookup_symbol
// container is not nil, container is Scanner
// is method
func (v Scanner) ScopeLookupSymbol(scope_id uint32, symbol string) (result unsafe.Pointer) {
	iv, err := _I.Get(502, "Scanner", "scope_lookup_symbol")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_symbol := gi.CString(symbol)
	arg_v := gi.NewPointerArgument(v.P)
	arg_scope_id := gi.NewUint32Argument(scope_id)
	arg_symbol := gi.NewStringArgument(c_symbol)
	args := []gi.Argument{arg_v, arg_scope_id, arg_symbol}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_symbol)
	result = ret.Pointer()
	return
}

// g_scanner_scope_remove_symbol
// container is not nil, container is Scanner
// is method
func (v Scanner) ScopeRemoveSymbol(scope_id uint32, symbol string) {
	iv, err := _I.Get(503, "Scanner", "scope_remove_symbol")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_symbol := gi.CString(symbol)
	arg_v := gi.NewPointerArgument(v.P)
	arg_scope_id := gi.NewUint32Argument(scope_id)
	arg_symbol := gi.NewStringArgument(c_symbol)
	args := []gi.Argument{arg_v, arg_scope_id, arg_symbol}
	iv.Call(args, nil, nil)
	gi.Free(c_symbol)
}

// g_scanner_set_scope
// container is not nil, container is Scanner
// is method
func (v Scanner) SetScope(scope_id uint32) (result uint32) {
	iv, err := _I.Get(504, "Scanner", "set_scope")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_scope_id := gi.NewUint32Argument(scope_id)
	args := []gi.Argument{arg_v, arg_scope_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_scanner_sync_file_offset
// container is not nil, container is Scanner
// is method
func (v Scanner) SyncFileOffset() {
	iv, err := _I.Get(505, "Scanner", "sync_file_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_scanner_unexp_token
// container is not nil, container is Scanner
// is method
func (v Scanner) UnexpToken(expected_token TokenTypeEnum, identifier_spec string, symbol_spec string, symbol_name string, message string, is_error int32) {
	iv, err := _I.Get(506, "Scanner", "unexp_token")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_identifier_spec := gi.CString(identifier_spec)
	c_symbol_spec := gi.CString(symbol_spec)
	c_symbol_name := gi.CString(symbol_name)
	c_message := gi.CString(message)
	arg_v := gi.NewPointerArgument(v.P)
	arg_expected_token := gi.NewIntArgument(int(expected_token))
	arg_identifier_spec := gi.NewStringArgument(c_identifier_spec)
	arg_symbol_spec := gi.NewStringArgument(c_symbol_spec)
	arg_symbol_name := gi.NewStringArgument(c_symbol_name)
	arg_message := gi.NewStringArgument(c_message)
	arg_is_error := gi.NewInt32Argument(is_error)
	args := []gi.Argument{arg_v, arg_expected_token, arg_identifier_spec, arg_symbol_spec, arg_symbol_name, arg_message, arg_is_error}
	iv.Call(args, nil, nil)
	gi.Free(c_identifier_spec)
	gi.Free(c_symbol_spec)
	gi.Free(c_symbol_name)
	gi.Free(c_message)
}

// Struct ScannerConfig
type ScannerConfig struct {
	P unsafe.Pointer
}

const SizeOfStructScannerConfig = 128

func ScannerConfigGetType() gi.GType {
	ret := _I.GetGType(82, "ScannerConfig")
	return ret
}

// ignore callback ScannerMsgFunc
// Enum SeekType
type SeekTypeEnum int

const (
	SeekTypeCur SeekTypeEnum = 0
	SeekTypeSet SeekTypeEnum = 1
	SeekTypeEnd SeekTypeEnum = 2
)

func SeekTypeGetType() gi.GType {
	ret := _I.GetGType(83, "SeekType")
	return ret
}

// Struct Sequence
type Sequence struct {
	P unsafe.Pointer
}

func SequenceGetType() gi.GType {
	ret := _I.GetGType(84, "Sequence")
	return ret
}

// g_sequence_append
// container is not nil, container is Sequence
// is method
func (v Sequence) Append(data unsafe.Pointer) (result SequenceIter) {
	iv, err := _I.Get(507, "Sequence", "append")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_free
// container is not nil, container is Sequence
// is method
func (v Sequence) Free() {
	iv, err := _I.Get(508, "Sequence", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_sequence_get_begin_iter
// container is not nil, container is Sequence
// is method
func (v Sequence) GetBeginIter() (result SequenceIter) {
	iv, err := _I.Get(509, "Sequence", "get_begin_iter")
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

// g_sequence_get_end_iter
// container is not nil, container is Sequence
// is method
func (v Sequence) GetEndIter() (result SequenceIter) {
	iv, err := _I.Get(510, "Sequence", "get_end_iter")
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

// g_sequence_get_iter_at_pos
// container is not nil, container is Sequence
// is method
func (v Sequence) GetIterAtPos(pos int32) (result SequenceIter) {
	iv, err := _I.Get(511, "Sequence", "get_iter_at_pos")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewInt32Argument(pos)
	args := []gi.Argument{arg_v, arg_pos}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_get_length
// container is not nil, container is Sequence
// is method
func (v Sequence) GetLength() (result int32) {
	iv, err := _I.Get(512, "Sequence", "get_length")
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

// g_sequence_is_empty
// container is not nil, container is Sequence
// is method
func (v Sequence) IsEmpty() (result bool) {
	iv, err := _I.Get(513, "Sequence", "is_empty")
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

// g_sequence_prepend
// container is not nil, container is Sequence
// is method
func (v Sequence) Prepend(data unsafe.Pointer) (result SequenceIter) {
	iv, err := _I.Get(514, "Sequence", "prepend")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_get
// container is not nil, container is Sequence
// is method
// arg0Type tag: interface, isPtr: true
func SequenceGet1(iter SequenceIter) (result unsafe.Pointer) {
	iv, err := _I.Get(515, "Sequence", "get")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	args := []gi.Argument{arg_iter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_sequence_insert_before
// container is not nil, container is Sequence
// is method
// arg0Type tag: interface, isPtr: true
func SequenceInsertBefore1(iter SequenceIter, data unsafe.Pointer) (result SequenceIter) {
	iv, err := _I.Get(516, "Sequence", "insert_before")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_iter, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_move
// container is not nil, container is Sequence
// is method
// arg0Type tag: interface, isPtr: true
func SequenceMove1(src SequenceIter, dest SequenceIter) {
	iv, err := _I.Get(517, "Sequence", "move")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_src := gi.NewPointerArgument(src.P)
	arg_dest := gi.NewPointerArgument(dest.P)
	args := []gi.Argument{arg_src, arg_dest}
	iv.Call(args, nil, nil)
}

// g_sequence_move_range
// container is not nil, container is Sequence
// is method
// arg0Type tag: interface, isPtr: true
func SequenceMoveRange1(dest SequenceIter, begin SequenceIter, end SequenceIter) {
	iv, err := _I.Get(518, "Sequence", "move_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_begin := gi.NewPointerArgument(begin.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_dest, arg_begin, arg_end}
	iv.Call(args, nil, nil)
}

// g_sequence_range_get_midpoint
// container is not nil, container is Sequence
// is method
// arg0Type tag: interface, isPtr: true
func SequenceRangeGetMidpoint1(begin SequenceIter, end SequenceIter) (result SequenceIter) {
	iv, err := _I.Get(519, "Sequence", "range_get_midpoint")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_begin := gi.NewPointerArgument(begin.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_begin, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_remove
// container is not nil, container is Sequence
// is method
// arg0Type tag: interface, isPtr: true
func SequenceRemove1(iter SequenceIter) {
	iv, err := _I.Get(520, "Sequence", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	args := []gi.Argument{arg_iter}
	iv.Call(args, nil, nil)
}

// g_sequence_remove_range
// container is not nil, container is Sequence
// is method
// arg0Type tag: interface, isPtr: true
func SequenceRemoveRange1(begin SequenceIter, end SequenceIter) {
	iv, err := _I.Get(521, "Sequence", "remove_range")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_begin := gi.NewPointerArgument(begin.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_begin, arg_end}
	iv.Call(args, nil, nil)
}

// g_sequence_set
// container is not nil, container is Sequence
// is method
// arg0Type tag: interface, isPtr: true
func SequenceSet1(iter SequenceIter, data unsafe.Pointer) {
	iv, err := _I.Get(522, "Sequence", "set")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_iter, arg_data}
	iv.Call(args, nil, nil)
}

// g_sequence_swap
// container is not nil, container is Sequence
// is method
// arg0Type tag: interface, isPtr: true
func SequenceSwap1(a SequenceIter, b SequenceIter) {
	iv, err := _I.Get(523, "Sequence", "swap")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_a := gi.NewPointerArgument(a.P)
	arg_b := gi.NewPointerArgument(b.P)
	args := []gi.Argument{arg_a, arg_b}
	iv.Call(args, nil, nil)
}

// Struct SequenceIter
type SequenceIter struct {
	P unsafe.Pointer
}

func SequenceIterGetType() gi.GType {
	ret := _I.GetGType(85, "SequenceIter")
	return ret
}

// g_sequence_iter_compare
// container is not nil, container is SequenceIter
// is method
func (v SequenceIter) Compare(b SequenceIter) (result int32) {
	iv, err := _I.Get(524, "SequenceIter", "compare")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_b := gi.NewPointerArgument(b.P)
	args := []gi.Argument{arg_v, arg_b}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_sequence_iter_get_position
// container is not nil, container is SequenceIter
// is method
func (v SequenceIter) GetPosition() (result int32) {
	iv, err := _I.Get(525, "SequenceIter", "get_position")
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

// g_sequence_iter_get_sequence
// container is not nil, container is SequenceIter
// is method
func (v SequenceIter) GetSequence() (result Sequence) {
	iv, err := _I.Get(526, "SequenceIter", "get_sequence")
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

// g_sequence_iter_is_begin
// container is not nil, container is SequenceIter
// is method
func (v SequenceIter) IsBegin() (result bool) {
	iv, err := _I.Get(527, "SequenceIter", "is_begin")
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

// g_sequence_iter_is_end
// container is not nil, container is SequenceIter
// is method
func (v SequenceIter) IsEnd() (result bool) {
	iv, err := _I.Get(528, "SequenceIter", "is_end")
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

// g_sequence_iter_move
// container is not nil, container is SequenceIter
// is method
func (v SequenceIter) Move(delta int32) (result SequenceIter) {
	iv, err := _I.Get(529, "SequenceIter", "move")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_delta := gi.NewInt32Argument(delta)
	args := []gi.Argument{arg_v, arg_delta}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_iter_next
// container is not nil, container is SequenceIter
// is method
func (v SequenceIter) Next() (result SequenceIter) {
	iv, err := _I.Get(530, "SequenceIter", "next")
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

// g_sequence_iter_prev
// container is not nil, container is SequenceIter
// is method
func (v SequenceIter) Prev() (result SequenceIter) {
	iv, err := _I.Get(531, "SequenceIter", "prev")
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

// ignore callback SequenceIterCompareFunc
// Enum ShellError
type ShellErrorEnum int

const (
	ShellErrorBadQuoting  ShellErrorEnum = 0
	ShellErrorEmptyString ShellErrorEnum = 1
	ShellErrorFailed      ShellErrorEnum = 2
)

func ShellErrorGetType() gi.GType {
	ret := _I.GetGType(86, "ShellError")
	return ret
}

// Enum SliceConfig
type SliceConfigEnum int

const (
	SliceConfigAlwaysMalloc      SliceConfigEnum = 1
	SliceConfigBypassMagazines   SliceConfigEnum = 2
	SliceConfigWorkingSetMsecs   SliceConfigEnum = 3
	SliceConfigColorIncrement    SliceConfigEnum = 4
	SliceConfigChunkSizes        SliceConfigEnum = 5
	SliceConfigContentionCounter SliceConfigEnum = 6
)

func SliceConfigGetType() gi.GType {
	ret := _I.GetGType(87, "SliceConfig")
	return ret
}

// Struct Source
type Source struct {
	P unsafe.Pointer
}

const SizeOfStructSource = 96

func SourceGetType() gi.GType {
	ret := _I.GetGType(88, "Source")
	return ret
}

// g_source_new
// container is not nil, container is Source
// is constructor
func NewSource(source_funcs SourceFuncs, struct_size uint32) (result Source) {
	iv, err := _I.Get(532, "Source", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_source_funcs := gi.NewPointerArgument(source_funcs.P)
	arg_struct_size := gi.NewUint32Argument(struct_size)
	args := []gi.Argument{arg_source_funcs, arg_struct_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_source_add_child_source
// container is not nil, container is Source
// is method
func (v Source) AddChildSource(child_source Source) {
	iv, err := _I.Get(533, "Source", "add_child_source")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_child_source := gi.NewPointerArgument(child_source.P)
	args := []gi.Argument{arg_v, arg_child_source}
	iv.Call(args, nil, nil)
}

// g_source_add_poll
// container is not nil, container is Source
// is method
func (v Source) AddPoll(fd PollFD) {
	iv, err := _I.Get(534, "Source", "add_poll")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_fd}
	iv.Call(args, nil, nil)
}

// g_source_add_unix_fd
// container is not nil, container is Source
// is method
func (v Source) AddUnixFd(fd int32, events IOConditionFlags) (result unsafe.Pointer) {
	iv, err := _I.Get(535, "Source", "add_unix_fd")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewInt32Argument(fd)
	arg_events := gi.NewIntArgument(int(events))
	args := []gi.Argument{arg_v, arg_fd, arg_events}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_source_attach
// container is not nil, container is Source
// is method
func (v Source) Attach(context MainContext) (result uint32) {
	iv, err := _I.Get(536, "Source", "attach")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_context := gi.NewPointerArgument(context.P)
	args := []gi.Argument{arg_v, arg_context}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_source_destroy
// container is not nil, container is Source
// is method
func (v Source) Destroy() {
	iv, err := _I.Get(537, "Source", "destroy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_source_get_can_recurse
// container is not nil, container is Source
// is method
func (v Source) GetCanRecurse() (result bool) {
	iv, err := _I.Get(538, "Source", "get_can_recurse")
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

// g_source_get_context
// container is not nil, container is Source
// is method
func (v Source) GetContext() (result MainContext) {
	iv, err := _I.Get(539, "Source", "get_context")
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

// g_source_get_current_time
// container is not nil, container is Source
// is method
func (v Source) GetCurrentTime(timeval TimeVal) {
	iv, err := _I.Get(540, "Source", "get_current_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_timeval := gi.NewPointerArgument(timeval.P)
	args := []gi.Argument{arg_v, arg_timeval}
	iv.Call(args, nil, nil)
}

// g_source_get_id
// container is not nil, container is Source
// is method
func (v Source) GetId() (result uint32) {
	iv, err := _I.Get(541, "Source", "get_id")
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

// g_source_get_name
// container is not nil, container is Source
// is method
func (v Source) GetName() (result string) {
	iv, err := _I.Get(542, "Source", "get_name")
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

// g_source_get_priority
// container is not nil, container is Source
// is method
func (v Source) GetPriority() (result int32) {
	iv, err := _I.Get(543, "Source", "get_priority")
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

// g_source_get_ready_time
// container is not nil, container is Source
// is method
func (v Source) GetReadyTime() (result int64) {
	iv, err := _I.Get(544, "Source", "get_ready_time")
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

// g_source_get_time
// container is not nil, container is Source
// is method
func (v Source) GetTime() (result int64) {
	iv, err := _I.Get(545, "Source", "get_time")
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

// g_source_is_destroyed
// container is not nil, container is Source
// is method
func (v Source) IsDestroyed() (result bool) {
	iv, err := _I.Get(546, "Source", "is_destroyed")
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

// g_source_modify_unix_fd
// container is not nil, container is Source
// is method
func (v Source) ModifyUnixFd(tag unsafe.Pointer, new_events IOConditionFlags) {
	iv, err := _I.Get(547, "Source", "modify_unix_fd")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewPointerArgument(tag)
	arg_new_events := gi.NewIntArgument(int(new_events))
	args := []gi.Argument{arg_v, arg_tag, arg_new_events}
	iv.Call(args, nil, nil)
}

// g_source_query_unix_fd
// container is not nil, container is Source
// is method
func (v Source) QueryUnixFd(tag unsafe.Pointer) (result IOConditionFlags) {
	iv, err := _I.Get(548, "Source", "query_unix_fd")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewPointerArgument(tag)
	args := []gi.Argument{arg_v, arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = IOConditionFlags(ret.Int())
	return
}

// g_source_ref
// container is not nil, container is Source
// is method
func (v Source) Ref() (result Source) {
	iv, err := _I.Get(549, "Source", "ref")
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

// g_source_remove_child_source
// container is not nil, container is Source
// is method
func (v Source) RemoveChildSource(child_source Source) {
	iv, err := _I.Get(550, "Source", "remove_child_source")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_child_source := gi.NewPointerArgument(child_source.P)
	args := []gi.Argument{arg_v, arg_child_source}
	iv.Call(args, nil, nil)
}

// g_source_remove_poll
// container is not nil, container is Source
// is method
func (v Source) RemovePoll(fd PollFD) {
	iv, err := _I.Get(551, "Source", "remove_poll")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_fd := gi.NewPointerArgument(fd.P)
	args := []gi.Argument{arg_v, arg_fd}
	iv.Call(args, nil, nil)
}

// g_source_remove_unix_fd
// container is not nil, container is Source
// is method
func (v Source) RemoveUnixFd(tag unsafe.Pointer) {
	iv, err := _I.Get(552, "Source", "remove_unix_fd")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewPointerArgument(tag)
	args := []gi.Argument{arg_v, arg_tag}
	iv.Call(args, nil, nil)
}

// g_source_set_callback
// container is not nil, container is Source
// is method
func (v Source) SetCallback(func1 int /*TODO_TYPE isPtr: false, tag: interface*/, data unsafe.Pointer, notify int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(553, "Source", "set_callback")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_func1 := gi.NewIntArgument(func1) /*TODO*/
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewIntArgument(notify) /*TODO*/
	args := []gi.Argument{arg_v, arg_func1, arg_data, arg_notify}
	iv.Call(args, nil, nil)
}

// g_source_set_callback_indirect
// container is not nil, container is Source
// is method
func (v Source) SetCallbackIndirect(callback_data unsafe.Pointer, callback_funcs SourceCallbackFuncs) {
	iv, err := _I.Get(554, "Source", "set_callback_indirect")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_callback_data := gi.NewPointerArgument(callback_data)
	arg_callback_funcs := gi.NewPointerArgument(callback_funcs.P)
	args := []gi.Argument{arg_v, arg_callback_data, arg_callback_funcs}
	iv.Call(args, nil, nil)
}

// g_source_set_can_recurse
// container is not nil, container is Source
// is method
func (v Source) SetCanRecurse(can_recurse bool) {
	iv, err := _I.Get(555, "Source", "set_can_recurse")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_can_recurse := gi.NewBoolArgument(can_recurse)
	args := []gi.Argument{arg_v, arg_can_recurse}
	iv.Call(args, nil, nil)
}

// g_source_set_funcs
// container is not nil, container is Source
// is method
func (v Source) SetFuncs(funcs SourceFuncs) {
	iv, err := _I.Get(556, "Source", "set_funcs")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_funcs := gi.NewPointerArgument(funcs.P)
	args := []gi.Argument{arg_v, arg_funcs}
	iv.Call(args, nil, nil)
}

// g_source_set_name
// container is not nil, container is Source
// is method
func (v Source) SetName(name string) {
	iv, err := _I.Get(557, "Source", "set_name")
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

// g_source_set_priority
// container is not nil, container is Source
// is method
func (v Source) SetPriority(priority int32) {
	iv, err := _I.Get(558, "Source", "set_priority")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_priority := gi.NewInt32Argument(priority)
	args := []gi.Argument{arg_v, arg_priority}
	iv.Call(args, nil, nil)
}

// g_source_set_ready_time
// container is not nil, container is Source
// is method
func (v Source) SetReadyTime(ready_time int64) {
	iv, err := _I.Get(559, "Source", "set_ready_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_ready_time := gi.NewInt64Argument(ready_time)
	args := []gi.Argument{arg_v, arg_ready_time}
	iv.Call(args, nil, nil)
}

// g_source_unref
// container is not nil, container is Source
// is method
func (v Source) Unref() {
	iv, err := _I.Get(560, "Source", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_source_remove
// container is not nil, container is Source
// is method
// arg0Type tag: guint32, isPtr: false
func SourceRemove1(tag uint32) (result bool) {
	iv, err := _I.Get(561, "Source", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_tag := gi.NewUint32Argument(tag)
	args := []gi.Argument{arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_source_remove_by_funcs_user_data
// container is not nil, container is Source
// is method
// arg0Type tag: interface, isPtr: true
func SourceRemoveByFuncsUserData1(funcs SourceFuncs, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(562, "Source", "remove_by_funcs_user_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_funcs := gi.NewPointerArgument(funcs.P)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_funcs, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_source_remove_by_user_data
// container is not nil, container is Source
// is method
// arg0Type tag: void, isPtr: true
func SourceRemoveByUserData1(user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(563, "Source", "remove_by_user_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_source_set_name_by_id
// container is not nil, container is Source
// is method
// arg0Type tag: guint32, isPtr: false
func SourceSetNameById1(tag uint32, name string) {
	iv, err := _I.Get(564, "Source", "set_name_by_id")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_tag := gi.NewUint32Argument(tag)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_tag, arg_name}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
}

// Struct SourceCallbackFuncs
type SourceCallbackFuncs struct {
	P unsafe.Pointer
}

const SizeOfStructSourceCallbackFuncs = 24

func SourceCallbackFuncsGetType() gi.GType {
	ret := _I.GetGType(89, "SourceCallbackFuncs")
	return ret
}

// ignore callback SourceDummyMarshal
type SourceFuncStruct struct {
}

//export mySourceFunc
func mySourceFunc(user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := SourceFuncStruct{}
	fn(args)
}

// Struct SourceFuncs
type SourceFuncs struct {
	P unsafe.Pointer
}

const SizeOfStructSourceFuncs = 48

func SourceFuncsGetType() gi.GType {
	ret := _I.GetGType(90, "SourceFuncs")
	return ret
}

// Struct SourcePrivate
type SourcePrivate struct {
	P unsafe.Pointer
}

func SourcePrivateGetType() gi.GType {
	ret := _I.GetGType(91, "SourcePrivate")
	return ret
}

type SpawnChildSetupFuncStruct struct {
}

//export mySpawnChildSetupFunc
func mySpawnChildSetupFunc(user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := SpawnChildSetupFuncStruct{}
	fn(args)
}

// Enum SpawnError
type SpawnErrorEnum int

const (
	SpawnErrorFork        SpawnErrorEnum = 0
	SpawnErrorRead        SpawnErrorEnum = 1
	SpawnErrorChdir       SpawnErrorEnum = 2
	SpawnErrorAcces       SpawnErrorEnum = 3
	SpawnErrorPerm        SpawnErrorEnum = 4
	SpawnErrorTooBig      SpawnErrorEnum = 5
	SpawnError2big        SpawnErrorEnum = 5
	SpawnErrorNoexec      SpawnErrorEnum = 6
	SpawnErrorNametoolong SpawnErrorEnum = 7
	SpawnErrorNoent       SpawnErrorEnum = 8
	SpawnErrorNomem       SpawnErrorEnum = 9
	SpawnErrorNotdir      SpawnErrorEnum = 10
	SpawnErrorLoop        SpawnErrorEnum = 11
	SpawnErrorTxtbusy     SpawnErrorEnum = 12
	SpawnErrorIo          SpawnErrorEnum = 13
	SpawnErrorNfile       SpawnErrorEnum = 14
	SpawnErrorMfile       SpawnErrorEnum = 15
	SpawnErrorInval       SpawnErrorEnum = 16
	SpawnErrorIsdir       SpawnErrorEnum = 17
	SpawnErrorLibbad      SpawnErrorEnum = 18
	SpawnErrorFailed      SpawnErrorEnum = 19
)

func SpawnErrorGetType() gi.GType {
	ret := _I.GetGType(92, "SpawnError")
	return ret
}

// Flags SpawnFlags
type SpawnFlags int

const (
	SpawnFlagsDefault              SpawnFlags = 0
	SpawnFlagsLeaveDescriptorsOpen SpawnFlags = 1
	SpawnFlagsDoNotReapChild       SpawnFlags = 2
	SpawnFlagsSearchPath           SpawnFlags = 4
	SpawnFlagsStdoutToDevNull      SpawnFlags = 8
	SpawnFlagsStderrToDevNull      SpawnFlags = 16
	SpawnFlagsChildInheritsStdin   SpawnFlags = 32
	SpawnFlagsFileAndArgvZero      SpawnFlags = 64
	SpawnFlagsSearchPathFromEnvp   SpawnFlags = 128
	SpawnFlagsCloexecPipes         SpawnFlags = 256
)

func SpawnFlagsGetType() gi.GType {
	ret := _I.GetGType(93, "SpawnFlags")
	return ret
}

// Struct StatBuf
type StatBuf struct {
	P unsafe.Pointer
}

func StatBufGetType() gi.GType {
	ret := _I.GetGType(94, "StatBuf")
	return ret
}

// Struct String
type String struct {
	P unsafe.Pointer
}

const SizeOfStructString = 24

func StringGetType() gi.GType {
	ret := _I.GetGType(95, "String")
	return ret
}

// g_string_append
// container is not nil, container is String
// is method
func (v String) Append(val string) (result String) {
	iv, err := _I.Get(565, "String", "append")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewStringArgument(c_val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_append_c
// container is not nil, container is String
// is method
func (v String) AppendC(c int8) (result String) {
	iv, err := _I.Get(566, "String", "append_c")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_c := gi.NewInt8Argument(c)
	args := []gi.Argument{arg_v, arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_append_len
// container is not nil, container is String
// is method
func (v String) AppendLen(val string, len1 int64) (result String) {
	iv, err := _I.Get(567, "String", "append_len")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewStringArgument(c_val)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_v, arg_val, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_append_unichar
// container is not nil, container is String
// is method
func (v String) AppendUnichar(wc rune) (result String) {
	iv, err := _I.Get(568, "String", "append_unichar")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_wc := gi.NewUint32Argument(uint32(wc))
	args := []gi.Argument{arg_v, arg_wc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_append_uri_escaped
// container is not nil, container is String
// is method
func (v String) AppendUriEscaped(unescaped string, reserved_chars_allowed string, allow_utf8 bool) (result String) {
	iv, err := _I.Get(569, "String", "append_uri_escaped")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_unescaped := gi.CString(unescaped)
	c_reserved_chars_allowed := gi.CString(reserved_chars_allowed)
	arg_v := gi.NewPointerArgument(v.P)
	arg_unescaped := gi.NewStringArgument(c_unescaped)
	arg_reserved_chars_allowed := gi.NewStringArgument(c_reserved_chars_allowed)
	arg_allow_utf8 := gi.NewBoolArgument(allow_utf8)
	args := []gi.Argument{arg_v, arg_unescaped, arg_reserved_chars_allowed, arg_allow_utf8}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_unescaped)
	gi.Free(c_reserved_chars_allowed)
	result.P = ret.Pointer()
	return
}

// g_string_ascii_down
// container is not nil, container is String
// is method
func (v String) AsciiDown() (result String) {
	iv, err := _I.Get(570, "String", "ascii_down")
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

// g_string_ascii_up
// container is not nil, container is String
// is method
func (v String) AsciiUp() (result String) {
	iv, err := _I.Get(571, "String", "ascii_up")
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

// g_string_assign
// container is not nil, container is String
// is method
func (v String) Assign(rval string) (result String) {
	iv, err := _I.Get(572, "String", "assign")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_rval := gi.CString(rval)
	arg_v := gi.NewPointerArgument(v.P)
	arg_rval := gi.NewStringArgument(c_rval)
	args := []gi.Argument{arg_v, arg_rval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_rval)
	result.P = ret.Pointer()
	return
}

// g_string_down
// container is not nil, container is String
// is method
func (v String) Down() (result String) {
	iv, err := _I.Get(573, "String", "down")
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

// g_string_equal
// container is not nil, container is String
// is method
func (v String) Equal(v2 String) (result bool) {
	iv, err := _I.Get(574, "String", "equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v2 := gi.NewPointerArgument(v2.P)
	args := []gi.Argument{arg_v, arg_v2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_string_erase
// container is not nil, container is String
// is method
func (v String) Erase(pos int64, len1 int64) (result String) {
	iv, err := _I.Get(575, "String", "erase")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewInt64Argument(pos)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_v, arg_pos, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_free
// container is not nil, container is String
// is method
func (v String) Free(free_segment bool) (result string) {
	iv, err := _I.Get(576, "String", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_free_segment := gi.NewBoolArgument(free_segment)
	args := []gi.Argument{arg_v, arg_free_segment}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_string_free_to_bytes
// container is not nil, container is String
// is method
func (v String) FreeToBytes() (result Bytes) {
	iv, err := _I.Get(577, "String", "free_to_bytes")
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

// g_string_hash
// container is not nil, container is String
// is method
func (v String) Hash() (result uint32) {
	iv, err := _I.Get(578, "String", "hash")
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

// g_string_insert
// container is not nil, container is String
// is method
func (v String) Insert(pos int64, val string) (result String) {
	iv, err := _I.Get(579, "String", "insert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewInt64Argument(pos)
	arg_val := gi.NewStringArgument(c_val)
	args := []gi.Argument{arg_v, arg_pos, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_insert_c
// container is not nil, container is String
// is method
func (v String) InsertC(pos int64, c int8) (result String) {
	iv, err := _I.Get(580, "String", "insert_c")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewInt64Argument(pos)
	arg_c := gi.NewInt8Argument(c)
	args := []gi.Argument{arg_v, arg_pos, arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_insert_len
// container is not nil, container is String
// is method
func (v String) InsertLen(pos int64, val string, len1 int64) (result String) {
	iv, err := _I.Get(581, "String", "insert_len")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewInt64Argument(pos)
	arg_val := gi.NewStringArgument(c_val)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_v, arg_pos, arg_val, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_insert_unichar
// container is not nil, container is String
// is method
func (v String) InsertUnichar(pos int64, wc rune) (result String) {
	iv, err := _I.Get(582, "String", "insert_unichar")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewInt64Argument(pos)
	arg_wc := gi.NewUint32Argument(uint32(wc))
	args := []gi.Argument{arg_v, arg_pos, arg_wc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_overwrite
// container is not nil, container is String
// is method
func (v String) Overwrite(pos uint64, val string) (result String) {
	iv, err := _I.Get(583, "String", "overwrite")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewUint64Argument(pos)
	arg_val := gi.NewStringArgument(c_val)
	args := []gi.Argument{arg_v, arg_pos, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_overwrite_len
// container is not nil, container is String
// is method
func (v String) OverwriteLen(pos uint64, val string, len1 int64) (result String) {
	iv, err := _I.Get(584, "String", "overwrite_len")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_pos := gi.NewUint64Argument(pos)
	arg_val := gi.NewStringArgument(c_val)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_v, arg_pos, arg_val, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_prepend
// container is not nil, container is String
// is method
func (v String) Prepend(val string) (result String) {
	iv, err := _I.Get(585, "String", "prepend")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewStringArgument(c_val)
	args := []gi.Argument{arg_v, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_prepend_c
// container is not nil, container is String
// is method
func (v String) PrependC(c int8) (result String) {
	iv, err := _I.Get(586, "String", "prepend_c")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_c := gi.NewInt8Argument(c)
	args := []gi.Argument{arg_v, arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_prepend_len
// container is not nil, container is String
// is method
func (v String) PrependLen(val string, len1 int64) (result String) {
	iv, err := _I.Get(587, "String", "prepend_len")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_val := gi.CString(val)
	arg_v := gi.NewPointerArgument(v.P)
	arg_val := gi.NewStringArgument(c_val)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_v, arg_val, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_val)
	result.P = ret.Pointer()
	return
}

// g_string_prepend_unichar
// container is not nil, container is String
// is method
func (v String) PrependUnichar(wc rune) (result String) {
	iv, err := _I.Get(588, "String", "prepend_unichar")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_wc := gi.NewUint32Argument(uint32(wc))
	args := []gi.Argument{arg_v, arg_wc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_set_size
// container is not nil, container is String
// is method
func (v String) SetSize(len1 uint64) (result String) {
	iv, err := _I.Get(589, "String", "set_size")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_len1 := gi.NewUint64Argument(len1)
	args := []gi.Argument{arg_v, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_truncate
// container is not nil, container is String
// is method
func (v String) Truncate(len1 uint64) (result String) {
	iv, err := _I.Get(590, "String", "truncate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_len1 := gi.NewUint64Argument(len1)
	args := []gi.Argument{arg_v, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_string_up
// container is not nil, container is String
// is method
func (v String) Up() (result String) {
	iv, err := _I.Get(591, "String", "up")
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

// Struct StringChunk
type StringChunk struct {
	P unsafe.Pointer
}

func StringChunkGetType() gi.GType {
	ret := _I.GetGType(96, "StringChunk")
	return ret
}

// g_string_chunk_clear
// container is not nil, container is StringChunk
// is method
func (v StringChunk) Clear() {
	iv, err := _I.Get(592, "StringChunk", "clear")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_string_chunk_free
// container is not nil, container is StringChunk
// is method
func (v StringChunk) Free() {
	iv, err := _I.Get(593, "StringChunk", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_string_chunk_insert
// container is not nil, container is StringChunk
// is method
func (v StringChunk) Insert(string string) (result string) {
	iv, err := _I.Get(594, "StringChunk", "insert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_v, arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_string_chunk_insert_const
// container is not nil, container is StringChunk
// is method
func (v StringChunk) InsertConst(string string) (result string) {
	iv, err := _I.Get(595, "StringChunk", "insert_const")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_v, arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_string_chunk_insert_len
// container is not nil, container is StringChunk
// is method
func (v StringChunk) InsertLen(string string, len1 int64) (result string) {
	iv, err := _I.Get(596, "StringChunk", "insert_len")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_string := gi.NewStringArgument(c_string)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_v, arg_string, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// Struct TestCase
type TestCase struct {
	P unsafe.Pointer
}

func TestCaseGetType() gi.GType {
	ret := _I.GetGType(97, "TestCase")
	return ret
}

// Struct TestConfig
type TestConfig struct {
	P unsafe.Pointer
}

const SizeOfStructTestConfig = 24

func TestConfigGetType() gi.GType {
	ret := _I.GetGType(98, "TestConfig")
	return ret
}

type TestDataFuncStruct struct {
}

//export myTestDataFunc
func myTestDataFunc(user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := TestDataFuncStruct{}
	fn(args)
}

// Enum TestFileType
type TestFileTypeEnum int

const (
	TestFileTypeDist  TestFileTypeEnum = 0
	TestFileTypeBuilt TestFileTypeEnum = 1
)

func TestFileTypeGetType() gi.GType {
	ret := _I.GetGType(99, "TestFileType")
	return ret
}

type TestFixtureFuncStruct struct {
	F_fixture unsafe.Pointer
}

//export myTestFixtureFunc
func myTestFixtureFunc(fixture C.gpointer, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := TestFixtureFuncStruct{
		F_fixture: unsafe.Pointer(fixture),
	}
	fn(args)
}

// ignore callback TestFunc
// Struct TestLogBuffer
type TestLogBuffer struct {
	P unsafe.Pointer
}

const SizeOfStructTestLogBuffer = 16

func TestLogBufferGetType() gi.GType {
	ret := _I.GetGType(100, "TestLogBuffer")
	return ret
}

// g_test_log_buffer_free
// container is not nil, container is TestLogBuffer
// is method
func (v TestLogBuffer) Free() {
	iv, err := _I.Get(597, "TestLogBuffer", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_test_log_buffer_push
// container is not nil, container is TestLogBuffer
// is method
func (v TestLogBuffer) Push(n_bytes uint32, bytes uint8) {
	iv, err := _I.Get(598, "TestLogBuffer", "push")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_bytes := gi.NewUint32Argument(n_bytes)
	arg_bytes := gi.NewUint8Argument(bytes)
	args := []gi.Argument{arg_v, arg_n_bytes, arg_bytes}
	iv.Call(args, nil, nil)
}

type TestLogFatalFuncStruct struct {
	F_log_domain string
	F_log_level  LogLevelFlags
	F_message    string
}

//export myTestLogFatalFunc
func myTestLogFatalFunc(log_domain *C.gchar, log_level C.GLogLevelFlags, message *C.gchar, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := TestLogFatalFuncStruct{
		F_log_domain: gi.GoString(unsafe.Pointer(log_domain)),
		F_log_level:  LogLevelFlags(log_level),
		F_message:    gi.GoString(unsafe.Pointer(message)),
	}
	fn(args)
}

// Struct TestLogMsg
type TestLogMsg struct {
	P unsafe.Pointer
}

const SizeOfStructTestLogMsg = 32

func TestLogMsgGetType() gi.GType {
	ret := _I.GetGType(101, "TestLogMsg")
	return ret
}

// g_test_log_msg_free
// container is not nil, container is TestLogMsg
// is method
func (v TestLogMsg) Free() {
	iv, err := _I.Get(599, "TestLogMsg", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum TestLogType
type TestLogTypeEnum int

const (
	TestLogTypeNone        TestLogTypeEnum = 0
	TestLogTypeError       TestLogTypeEnum = 1
	TestLogTypeStartBinary TestLogTypeEnum = 2
	TestLogTypeListCase    TestLogTypeEnum = 3
	TestLogTypeSkipCase    TestLogTypeEnum = 4
	TestLogTypeStartCase   TestLogTypeEnum = 5
	TestLogTypeStopCase    TestLogTypeEnum = 6
	TestLogTypeMinResult   TestLogTypeEnum = 7
	TestLogTypeMaxResult   TestLogTypeEnum = 8
	TestLogTypeMessage     TestLogTypeEnum = 9
	TestLogTypeStartSuite  TestLogTypeEnum = 10
	TestLogTypeStopSuite   TestLogTypeEnum = 11
)

func TestLogTypeGetType() gi.GType {
	ret := _I.GetGType(102, "TestLogType")
	return ret
}

// Enum TestResult
type TestResultEnum int

const (
	TestResultSuccess    TestResultEnum = 0
	TestResultSkipped    TestResultEnum = 1
	TestResultFailure    TestResultEnum = 2
	TestResultIncomplete TestResultEnum = 3
)

func TestResultGetType() gi.GType {
	ret := _I.GetGType(103, "TestResult")
	return ret
}

// Flags TestSubprocessFlags
type TestSubprocessFlags int

const (
	TestSubprocessFlagsStdin  TestSubprocessFlags = 1
	TestSubprocessFlagsStdout TestSubprocessFlags = 2
	TestSubprocessFlagsStderr TestSubprocessFlags = 4
)

func TestSubprocessFlagsGetType() gi.GType {
	ret := _I.GetGType(104, "TestSubprocessFlags")
	return ret
}

// Struct TestSuite
type TestSuite struct {
	P unsafe.Pointer
}

func TestSuiteGetType() gi.GType {
	ret := _I.GetGType(105, "TestSuite")
	return ret
}

// g_test_suite_add
// container is not nil, container is TestSuite
// is method
func (v TestSuite) Add(test_case TestCase) {
	iv, err := _I.Get(600, "TestSuite", "add")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_test_case := gi.NewPointerArgument(test_case.P)
	args := []gi.Argument{arg_v, arg_test_case}
	iv.Call(args, nil, nil)
}

// g_test_suite_add_suite
// container is not nil, container is TestSuite
// is method
func (v TestSuite) AddSuite(nestedsuite TestSuite) {
	iv, err := _I.Get(601, "TestSuite", "add_suite")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_nestedsuite := gi.NewPointerArgument(nestedsuite.P)
	args := []gi.Argument{arg_v, arg_nestedsuite}
	iv.Call(args, nil, nil)
}

// Flags TestTrapFlags
type TestTrapFlags int

const (
	TestTrapFlagsSilenceStdout TestTrapFlags = 128
	TestTrapFlagsSilenceStderr TestTrapFlags = 256
	TestTrapFlagsInheritStdin  TestTrapFlags = 512
)

func TestTrapFlagsGetType() gi.GType {
	ret := _I.GetGType(106, "TestTrapFlags")
	return ret
}

// Struct Thread
type Thread struct {
	P unsafe.Pointer
}

func ThreadGetType() gi.GType {
	ret := _I.GetGType(107, "Thread")
	return ret
}

// g_thread_join
// container is not nil, container is Thread
// is method
func (v Thread) Join() (result unsafe.Pointer) {
	iv, err := _I.Get(602, "Thread", "join")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_thread_ref
// container is not nil, container is Thread
// is method
func (v Thread) Ref() (result Thread) {
	iv, err := _I.Get(603, "Thread", "ref")
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

// g_thread_unref
// container is not nil, container is Thread
// is method
func (v Thread) Unref() {
	iv, err := _I.Get(604, "Thread", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_thread_exit
// container is not nil, container is Thread
// is method
// arg0Type tag: void, isPtr: true
func ThreadExit1(retval unsafe.Pointer) {
	iv, err := _I.Get(606, "Thread", "exit")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_retval := gi.NewPointerArgument(retval)
	args := []gi.Argument{arg_retval}
	iv.Call(args, nil, nil)
}

// Enum ThreadError
type ThreadErrorEnum int

const (
	ThreadErrorThreadErrorAgain ThreadErrorEnum = 0
)

func ThreadErrorGetType() gi.GType {
	ret := _I.GetGType(108, "ThreadError")
	return ret
}

// ignore callback ThreadFunc
// Struct ThreadPool
type ThreadPool struct {
	P unsafe.Pointer
}

const SizeOfStructThreadPool = 24

func ThreadPoolGetType() gi.GType {
	ret := _I.GetGType(109, "ThreadPool")
	return ret
}

// g_thread_pool_free
// container is not nil, container is ThreadPool
// is method
func (v ThreadPool) Free(immediate bool, wait_ bool) {
	iv, err := _I.Get(609, "ThreadPool", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_immediate := gi.NewBoolArgument(immediate)
	arg_wait_ := gi.NewBoolArgument(wait_)
	args := []gi.Argument{arg_v, arg_immediate, arg_wait_}
	iv.Call(args, nil, nil)
}

// g_thread_pool_get_max_threads
// container is not nil, container is ThreadPool
// is method
func (v ThreadPool) GetMaxThreads() (result int32) {
	iv, err := _I.Get(610, "ThreadPool", "get_max_threads")
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

// g_thread_pool_get_num_threads
// container is not nil, container is ThreadPool
// is method
func (v ThreadPool) GetNumThreads() (result uint32) {
	iv, err := _I.Get(611, "ThreadPool", "get_num_threads")
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

// g_thread_pool_move_to_front
// container is not nil, container is ThreadPool
// is method
func (v ThreadPool) MoveToFront(data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(612, "ThreadPool", "move_to_front")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_thread_pool_push
// container is not nil, container is ThreadPool
// is method
func (v ThreadPool) Push(data unsafe.Pointer) (result bool, err error) {
	iv, err := _I.Get(613, "ThreadPool", "push")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_data, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_thread_pool_set_max_threads
// container is not nil, container is ThreadPool
// is method
func (v ThreadPool) SetMaxThreads(max_threads int32) (result bool, err error) {
	iv, err := _I.Get(614, "ThreadPool", "set_max_threads")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_max_threads := gi.NewInt32Argument(max_threads)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_max_threads, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_thread_pool_unprocessed
// container is not nil, container is ThreadPool
// is method
func (v ThreadPool) Unprocessed() (result uint32) {
	iv, err := _I.Get(615, "ThreadPool", "unprocessed")
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

// g_thread_pool_set_max_idle_time
// container is not nil, container is ThreadPool
// is method
// arg0Type tag: guint32, isPtr: false
func ThreadPoolSetMaxIdleTime1(interval uint32) {
	iv, err := _I.Get(619, "ThreadPool", "set_max_idle_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_interval := gi.NewUint32Argument(interval)
	args := []gi.Argument{arg_interval}
	iv.Call(args, nil, nil)
}

// g_thread_pool_set_max_unused_threads
// container is not nil, container is ThreadPool
// is method
// arg0Type tag: gint32, isPtr: false
func ThreadPoolSetMaxUnusedThreads1(max_threads int32) {
	iv, err := _I.Get(620, "ThreadPool", "set_max_unused_threads")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_max_threads := gi.NewInt32Argument(max_threads)
	args := []gi.Argument{arg_max_threads}
	iv.Call(args, nil, nil)
}

// Enum TimeType
type TimeTypeEnum int

const (
	TimeTypeStandard  TimeTypeEnum = 0
	TimeTypeDaylight  TimeTypeEnum = 1
	TimeTypeUniversal TimeTypeEnum = 2
)

func TimeTypeGetType() gi.GType {
	ret := _I.GetGType(110, "TimeType")
	return ret
}

// Struct TimeVal
type TimeVal struct {
	P unsafe.Pointer
}

const SizeOfStructTimeVal = 16

func TimeValGetType() gi.GType {
	ret := _I.GetGType(111, "TimeVal")
	return ret
}

// g_time_val_add
// container is not nil, container is TimeVal
// is method
func (v TimeVal) Add(microseconds int64) {
	iv, err := _I.Get(622, "TimeVal", "add")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_microseconds := gi.NewInt64Argument(microseconds)
	args := []gi.Argument{arg_v, arg_microseconds}
	iv.Call(args, nil, nil)
}

// g_time_val_to_iso8601
// container is not nil, container is TimeVal
// is method
func (v TimeVal) ToIso8601() (result string) {
	iv, err := _I.Get(623, "TimeVal", "to_iso8601")
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

// g_time_val_from_iso8601
// container is not nil, container is TimeVal
// is method
// arg0Type tag: utf8, isPtr: true
func TimeValFromIso86011(iso_date string, time_ TimeVal) (result bool) {
	iv, err := _I.Get(624, "TimeVal", "from_iso8601")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_iso_date := gi.CString(iso_date)
	arg_iso_date := gi.NewStringArgument(c_iso_date)
	arg_time_ := gi.NewPointerArgument(time_.P)
	args := []gi.Argument{arg_iso_date, arg_time_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_iso_date)
	result = ret.Bool()
	return
}

// Struct TimeZone
type TimeZone struct {
	P unsafe.Pointer
}

func TimeZoneGetType() gi.GType {
	ret := _I.GetGType(112, "TimeZone")
	return ret
}

// g_time_zone_new
// container is not nil, container is TimeZone
// is constructor
func NewTimeZone(identifier string) (result TimeZone) {
	iv, err := _I.Get(625, "TimeZone", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_identifier := gi.CString(identifier)
	arg_identifier := gi.NewStringArgument(c_identifier)
	args := []gi.Argument{arg_identifier}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_identifier)
	result.P = ret.Pointer()
	return
}

// g_time_zone_new_local
// container is not nil, container is TimeZone
// is constructor
func NewTimeZoneLocal() (result TimeZone) {
	iv, err := _I.Get(626, "TimeZone", "new_local")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_time_zone_new_offset
// container is not nil, container is TimeZone
// is constructor
func NewTimeZoneOffset(seconds int32) (result TimeZone) {
	iv, err := _I.Get(627, "TimeZone", "new_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_seconds := gi.NewInt32Argument(seconds)
	args := []gi.Argument{arg_seconds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_time_zone_new_utc
// container is not nil, container is TimeZone
// is constructor
func NewTimeZoneUtc() (result TimeZone) {
	iv, err := _I.Get(628, "TimeZone", "new_utc")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_time_zone_adjust_time
// container is not nil, container is TimeZone
// is method
func (v TimeZone) AdjustTime(type1 TimeTypeEnum, time_ int64) (result int32) {
	iv, err := _I.Get(629, "TimeZone", "adjust_time")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_time_ := gi.NewInt64Argument(time_)
	args := []gi.Argument{arg_v, arg_type1, arg_time_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_time_zone_find_interval
// container is not nil, container is TimeZone
// is method
func (v TimeZone) FindInterval(type1 TimeTypeEnum, time_ int64) (result int32) {
	iv, err := _I.Get(630, "TimeZone", "find_interval")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_time_ := gi.NewInt64Argument(time_)
	args := []gi.Argument{arg_v, arg_type1, arg_time_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_time_zone_get_abbreviation
// container is not nil, container is TimeZone
// is method
func (v TimeZone) GetAbbreviation(interval int32) (result string) {
	iv, err := _I.Get(631, "TimeZone", "get_abbreviation")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interval := gi.NewInt32Argument(interval)
	args := []gi.Argument{arg_v, arg_interval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_time_zone_get_identifier
// container is not nil, container is TimeZone
// is method
func (v TimeZone) GetIdentifier() (result string) {
	iv, err := _I.Get(632, "TimeZone", "get_identifier")
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

// g_time_zone_get_offset
// container is not nil, container is TimeZone
// is method
func (v TimeZone) GetOffset(interval int32) (result int32) {
	iv, err := _I.Get(633, "TimeZone", "get_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interval := gi.NewInt32Argument(interval)
	args := []gi.Argument{arg_v, arg_interval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_time_zone_is_dst
// container is not nil, container is TimeZone
// is method
func (v TimeZone) IsDst(interval int32) (result bool) {
	iv, err := _I.Get(634, "TimeZone", "is_dst")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_interval := gi.NewInt32Argument(interval)
	args := []gi.Argument{arg_v, arg_interval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_time_zone_ref
// container is not nil, container is TimeZone
// is method
func (v TimeZone) Ref() (result TimeZone) {
	iv, err := _I.Get(635, "TimeZone", "ref")
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

// g_time_zone_unref
// container is not nil, container is TimeZone
// is method
func (v TimeZone) Unref() {
	iv, err := _I.Get(636, "TimeZone", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Struct Timer
type Timer struct {
	P unsafe.Pointer
}

func TimerGetType() gi.GType {
	ret := _I.GetGType(113, "Timer")
	return ret
}

// g_timer_continue
// container is not nil, container is Timer
// is method
func (v Timer) Continue() {
	iv, err := _I.Get(637, "Timer", "continue")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_timer_destroy
// container is not nil, container is Timer
// is method
func (v Timer) Destroy() {
	iv, err := _I.Get(638, "Timer", "destroy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_timer_elapsed
// container is not nil, container is Timer
// is method
func (v Timer) Elapsed(microseconds uint64) (result float64) {
	iv, err := _I.Get(639, "Timer", "elapsed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_microseconds := gi.NewUint64Argument(microseconds)
	args := []gi.Argument{arg_v, arg_microseconds}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// g_timer_reset
// container is not nil, container is Timer
// is method
func (v Timer) Reset() {
	iv, err := _I.Get(640, "Timer", "reset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_timer_start
// container is not nil, container is Timer
// is method
func (v Timer) Start() {
	iv, err := _I.Get(641, "Timer", "start")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_timer_stop
// container is not nil, container is Timer
// is method
func (v Timer) Stop() {
	iv, err := _I.Get(642, "Timer", "stop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum TokenType
type TokenTypeEnum int

const (
	TokenTypeEof            TokenTypeEnum = 0
	TokenTypeLeftParen      TokenTypeEnum = 40
	TokenTypeRightParen     TokenTypeEnum = 41
	TokenTypeLeftCurly      TokenTypeEnum = 123
	TokenTypeRightCurly     TokenTypeEnum = 125
	TokenTypeLeftBrace      TokenTypeEnum = 91
	TokenTypeRightBrace     TokenTypeEnum = 93
	TokenTypeEqualSign      TokenTypeEnum = 61
	TokenTypeComma          TokenTypeEnum = 44
	TokenTypeNone           TokenTypeEnum = 256
	TokenTypeError          TokenTypeEnum = 257
	TokenTypeChar           TokenTypeEnum = 258
	TokenTypeBinary         TokenTypeEnum = 259
	TokenTypeOctal          TokenTypeEnum = 260
	TokenTypeInt            TokenTypeEnum = 261
	TokenTypeHex            TokenTypeEnum = 262
	TokenTypeFloat          TokenTypeEnum = 263
	TokenTypeString         TokenTypeEnum = 264
	TokenTypeSymbol         TokenTypeEnum = 265
	TokenTypeIdentifier     TokenTypeEnum = 266
	TokenTypeIdentifierNull TokenTypeEnum = 267
	TokenTypeCommentSingle  TokenTypeEnum = 268
	TokenTypeCommentMulti   TokenTypeEnum = 269
)

func TokenTypeGetType() gi.GType {
	ret := _I.GetGType(114, "TokenType")
	return ret
}

// Union TokenValue
type TokenValue struct {
	P unsafe.Pointer
}

const SizeOfUnionTokenValue = 8

func TokenValueGetType() gi.GType {
	ret := _I.GetGType(115, "TokenValue")
	return ret
}

// ignore callback TranslateFunc
// Struct TrashStack
type TrashStack struct {
	P unsafe.Pointer
}

const SizeOfStructTrashStack = 8

func TrashStackGetType() gi.GType {
	ret := _I.GetGType(116, "TrashStack")
	return ret
}

// g_trash_stack_height
// container is not nil, container is TrashStack
// is method
// arg0Type tag: interface, isPtr: true
func (v TrashStack) Height() (result uint32) {
	iv, err := _I.Get(643, "TrashStack", "height")
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

// g_trash_stack_peek
// container is not nil, container is TrashStack
// is method
// arg0Type tag: interface, isPtr: true
func (v TrashStack) Peek() (result unsafe.Pointer) {
	iv, err := _I.Get(644, "TrashStack", "peek")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_trash_stack_pop
// container is not nil, container is TrashStack
// is method
// arg0Type tag: interface, isPtr: true
func (v TrashStack) Pop() (result unsafe.Pointer) {
	iv, err := _I.Get(645, "TrashStack", "pop")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_trash_stack_push
// container is not nil, container is TrashStack
// is method
// arg0Type tag: interface, isPtr: true
func (v TrashStack) Push(data_p unsafe.Pointer) {
	iv, err := _I.Get(646, "TrashStack", "push")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data_p := gi.NewPointerArgument(data_p)
	args := []gi.Argument{arg_v, arg_data_p}
	iv.Call(args, nil, nil)
}

// Flags TraverseFlags
type TraverseFlags int

const (
	TraverseFlagsLeaves    TraverseFlags = 1
	TraverseFlagsNonLeaves TraverseFlags = 2
	TraverseFlagsAll       TraverseFlags = 3
	TraverseFlagsMask      TraverseFlags = 3
	TraverseFlagsLeafs     TraverseFlags = 1
	TraverseFlagsNonLeafs  TraverseFlags = 2
)

func TraverseFlagsGetType() gi.GType {
	ret := _I.GetGType(117, "TraverseFlags")
	return ret
}

// ignore callback TraverseFunc
// Enum TraverseType
type TraverseTypeEnum int

const (
	TraverseTypeInOrder    TraverseTypeEnum = 0
	TraverseTypePreOrder   TraverseTypeEnum = 1
	TraverseTypePostOrder  TraverseTypeEnum = 2
	TraverseTypeLevelOrder TraverseTypeEnum = 3
)

func TraverseTypeGetType() gi.GType {
	ret := _I.GetGType(118, "TraverseType")
	return ret
}

// Struct Tree
type Tree struct {
	P unsafe.Pointer
}

func TreeGetType() gi.GType {
	ret := _I.GetGType(119, "Tree")
	return ret
}

// g_tree_destroy
// container is not nil, container is Tree
// is method
func (v Tree) Destroy() {
	iv, err := _I.Get(647, "Tree", "destroy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_tree_height
// container is not nil, container is Tree
// is method
func (v Tree) Height() (result int32) {
	iv, err := _I.Get(648, "Tree", "height")
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

// g_tree_insert
// container is not nil, container is Tree
// is method
func (v Tree) Insert(key unsafe.Pointer, value unsafe.Pointer) {
	iv, err := _I.Get(649, "Tree", "insert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewPointerArgument(key)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_v, arg_key, arg_value}
	iv.Call(args, nil, nil)
}

// g_tree_lookup
// container is not nil, container is Tree
// is method
func (v Tree) Lookup(key unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(650, "Tree", "lookup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_v, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_tree_lookup_extended
// container is not nil, container is Tree
// is method
func (v Tree) LookupExtended(lookup_key unsafe.Pointer, orig_key unsafe.Pointer, value unsafe.Pointer) (result bool) {
	iv, err := _I.Get(651, "Tree", "lookup_extended")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_lookup_key := gi.NewPointerArgument(lookup_key)
	arg_orig_key := gi.NewPointerArgument(orig_key)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_v, arg_lookup_key, arg_orig_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_tree_nnodes
// container is not nil, container is Tree
// is method
func (v Tree) Nnodes() (result int32) {
	iv, err := _I.Get(652, "Tree", "nnodes")
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

// g_tree_remove
// container is not nil, container is Tree
// is method
func (v Tree) Remove(key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(653, "Tree", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_v, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_tree_replace
// container is not nil, container is Tree
// is method
func (v Tree) Replace(key unsafe.Pointer, value unsafe.Pointer) {
	iv, err := _I.Get(654, "Tree", "replace")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewPointerArgument(key)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_v, arg_key, arg_value}
	iv.Call(args, nil, nil)
}

// g_tree_steal
// container is not nil, container is Tree
// is method
func (v Tree) Steal(key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(655, "Tree", "steal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_v, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_tree_unref
// container is not nil, container is Tree
// is method
func (v Tree) Unref() {
	iv, err := _I.Get(656, "Tree", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum UnicodeBreakType
type UnicodeBreakTypeEnum int

const (
	UnicodeBreakTypeMandatory                  UnicodeBreakTypeEnum = 0
	UnicodeBreakTypeCarriageReturn             UnicodeBreakTypeEnum = 1
	UnicodeBreakTypeLineFeed                   UnicodeBreakTypeEnum = 2
	UnicodeBreakTypeCombiningMark              UnicodeBreakTypeEnum = 3
	UnicodeBreakTypeSurrogate                  UnicodeBreakTypeEnum = 4
	UnicodeBreakTypeZeroWidthSpace             UnicodeBreakTypeEnum = 5
	UnicodeBreakTypeInseparable                UnicodeBreakTypeEnum = 6
	UnicodeBreakTypeNonBreakingGlue            UnicodeBreakTypeEnum = 7
	UnicodeBreakTypeContingent                 UnicodeBreakTypeEnum = 8
	UnicodeBreakTypeSpace                      UnicodeBreakTypeEnum = 9
	UnicodeBreakTypeAfter                      UnicodeBreakTypeEnum = 10
	UnicodeBreakTypeBefore                     UnicodeBreakTypeEnum = 11
	UnicodeBreakTypeBeforeAndAfter             UnicodeBreakTypeEnum = 12
	UnicodeBreakTypeHyphen                     UnicodeBreakTypeEnum = 13
	UnicodeBreakTypeNonStarter                 UnicodeBreakTypeEnum = 14
	UnicodeBreakTypeOpenPunctuation            UnicodeBreakTypeEnum = 15
	UnicodeBreakTypeClosePunctuation           UnicodeBreakTypeEnum = 16
	UnicodeBreakTypeQuotation                  UnicodeBreakTypeEnum = 17
	UnicodeBreakTypeExclamation                UnicodeBreakTypeEnum = 18
	UnicodeBreakTypeIdeographic                UnicodeBreakTypeEnum = 19
	UnicodeBreakTypeNumeric                    UnicodeBreakTypeEnum = 20
	UnicodeBreakTypeInfixSeparator             UnicodeBreakTypeEnum = 21
	UnicodeBreakTypeSymbol                     UnicodeBreakTypeEnum = 22
	UnicodeBreakTypeAlphabetic                 UnicodeBreakTypeEnum = 23
	UnicodeBreakTypePrefix                     UnicodeBreakTypeEnum = 24
	UnicodeBreakTypePostfix                    UnicodeBreakTypeEnum = 25
	UnicodeBreakTypeComplexContext             UnicodeBreakTypeEnum = 26
	UnicodeBreakTypeAmbiguous                  UnicodeBreakTypeEnum = 27
	UnicodeBreakTypeUnknown                    UnicodeBreakTypeEnum = 28
	UnicodeBreakTypeNextLine                   UnicodeBreakTypeEnum = 29
	UnicodeBreakTypeWordJoiner                 UnicodeBreakTypeEnum = 30
	UnicodeBreakTypeHangulLJamo                UnicodeBreakTypeEnum = 31
	UnicodeBreakTypeHangulVJamo                UnicodeBreakTypeEnum = 32
	UnicodeBreakTypeHangulTJamo                UnicodeBreakTypeEnum = 33
	UnicodeBreakTypeHangulLvSyllable           UnicodeBreakTypeEnum = 34
	UnicodeBreakTypeHangulLvtSyllable          UnicodeBreakTypeEnum = 35
	UnicodeBreakTypeCloseParanthesis           UnicodeBreakTypeEnum = 36
	UnicodeBreakTypeConditionalJapaneseStarter UnicodeBreakTypeEnum = 37
	UnicodeBreakTypeHebrewLetter               UnicodeBreakTypeEnum = 38
	UnicodeBreakTypeRegionalIndicator          UnicodeBreakTypeEnum = 39
	UnicodeBreakTypeEmojiBase                  UnicodeBreakTypeEnum = 40
	UnicodeBreakTypeEmojiModifier              UnicodeBreakTypeEnum = 41
	UnicodeBreakTypeZeroWidthJoiner            UnicodeBreakTypeEnum = 42
)

func UnicodeBreakTypeGetType() gi.GType {
	ret := _I.GetGType(120, "UnicodeBreakType")
	return ret
}

// Enum UnicodeScript
type UnicodeScriptEnum int

const (
	UnicodeScriptInvalidCode           UnicodeScriptEnum = -1
	UnicodeScriptCommon                UnicodeScriptEnum = 0
	UnicodeScriptInherited             UnicodeScriptEnum = 1
	UnicodeScriptArabic                UnicodeScriptEnum = 2
	UnicodeScriptArmenian              UnicodeScriptEnum = 3
	UnicodeScriptBengali               UnicodeScriptEnum = 4
	UnicodeScriptBopomofo              UnicodeScriptEnum = 5
	UnicodeScriptCherokee              UnicodeScriptEnum = 6
	UnicodeScriptCoptic                UnicodeScriptEnum = 7
	UnicodeScriptCyrillic              UnicodeScriptEnum = 8
	UnicodeScriptDeseret               UnicodeScriptEnum = 9
	UnicodeScriptDevanagari            UnicodeScriptEnum = 10
	UnicodeScriptEthiopic              UnicodeScriptEnum = 11
	UnicodeScriptGeorgian              UnicodeScriptEnum = 12
	UnicodeScriptGothic                UnicodeScriptEnum = 13
	UnicodeScriptGreek                 UnicodeScriptEnum = 14
	UnicodeScriptGujarati              UnicodeScriptEnum = 15
	UnicodeScriptGurmukhi              UnicodeScriptEnum = 16
	UnicodeScriptHan                   UnicodeScriptEnum = 17
	UnicodeScriptHangul                UnicodeScriptEnum = 18
	UnicodeScriptHebrew                UnicodeScriptEnum = 19
	UnicodeScriptHiragana              UnicodeScriptEnum = 20
	UnicodeScriptKannada               UnicodeScriptEnum = 21
	UnicodeScriptKatakana              UnicodeScriptEnum = 22
	UnicodeScriptKhmer                 UnicodeScriptEnum = 23
	UnicodeScriptLao                   UnicodeScriptEnum = 24
	UnicodeScriptLatin                 UnicodeScriptEnum = 25
	UnicodeScriptMalayalam             UnicodeScriptEnum = 26
	UnicodeScriptMongolian             UnicodeScriptEnum = 27
	UnicodeScriptMyanmar               UnicodeScriptEnum = 28
	UnicodeScriptOgham                 UnicodeScriptEnum = 29
	UnicodeScriptOldItalic             UnicodeScriptEnum = 30
	UnicodeScriptOriya                 UnicodeScriptEnum = 31
	UnicodeScriptRunic                 UnicodeScriptEnum = 32
	UnicodeScriptSinhala               UnicodeScriptEnum = 33
	UnicodeScriptSyriac                UnicodeScriptEnum = 34
	UnicodeScriptTamil                 UnicodeScriptEnum = 35
	UnicodeScriptTelugu                UnicodeScriptEnum = 36
	UnicodeScriptThaana                UnicodeScriptEnum = 37
	UnicodeScriptThai                  UnicodeScriptEnum = 38
	UnicodeScriptTibetan               UnicodeScriptEnum = 39
	UnicodeScriptCanadianAboriginal    UnicodeScriptEnum = 40
	UnicodeScriptYi                    UnicodeScriptEnum = 41
	UnicodeScriptTagalog               UnicodeScriptEnum = 42
	UnicodeScriptHanunoo               UnicodeScriptEnum = 43
	UnicodeScriptBuhid                 UnicodeScriptEnum = 44
	UnicodeScriptTagbanwa              UnicodeScriptEnum = 45
	UnicodeScriptBraille               UnicodeScriptEnum = 46
	UnicodeScriptCypriot               UnicodeScriptEnum = 47
	UnicodeScriptLimbu                 UnicodeScriptEnum = 48
	UnicodeScriptOsmanya               UnicodeScriptEnum = 49
	UnicodeScriptShavian               UnicodeScriptEnum = 50
	UnicodeScriptLinearB               UnicodeScriptEnum = 51
	UnicodeScriptTaiLe                 UnicodeScriptEnum = 52
	UnicodeScriptUgaritic              UnicodeScriptEnum = 53
	UnicodeScriptNewTaiLue             UnicodeScriptEnum = 54
	UnicodeScriptBuginese              UnicodeScriptEnum = 55
	UnicodeScriptGlagolitic            UnicodeScriptEnum = 56
	UnicodeScriptTifinagh              UnicodeScriptEnum = 57
	UnicodeScriptSylotiNagri           UnicodeScriptEnum = 58
	UnicodeScriptOldPersian            UnicodeScriptEnum = 59
	UnicodeScriptKharoshthi            UnicodeScriptEnum = 60
	UnicodeScriptUnknown               UnicodeScriptEnum = 61
	UnicodeScriptBalinese              UnicodeScriptEnum = 62
	UnicodeScriptCuneiform             UnicodeScriptEnum = 63
	UnicodeScriptPhoenician            UnicodeScriptEnum = 64
	UnicodeScriptPhagsPa               UnicodeScriptEnum = 65
	UnicodeScriptNko                   UnicodeScriptEnum = 66
	UnicodeScriptKayahLi               UnicodeScriptEnum = 67
	UnicodeScriptLepcha                UnicodeScriptEnum = 68
	UnicodeScriptRejang                UnicodeScriptEnum = 69
	UnicodeScriptSundanese             UnicodeScriptEnum = 70
	UnicodeScriptSaurashtra            UnicodeScriptEnum = 71
	UnicodeScriptCham                  UnicodeScriptEnum = 72
	UnicodeScriptOlChiki               UnicodeScriptEnum = 73
	UnicodeScriptVai                   UnicodeScriptEnum = 74
	UnicodeScriptCarian                UnicodeScriptEnum = 75
	UnicodeScriptLycian                UnicodeScriptEnum = 76
	UnicodeScriptLydian                UnicodeScriptEnum = 77
	UnicodeScriptAvestan               UnicodeScriptEnum = 78
	UnicodeScriptBamum                 UnicodeScriptEnum = 79
	UnicodeScriptEgyptianHieroglyphs   UnicodeScriptEnum = 80
	UnicodeScriptImperialAramaic       UnicodeScriptEnum = 81
	UnicodeScriptInscriptionalPahlavi  UnicodeScriptEnum = 82
	UnicodeScriptInscriptionalParthian UnicodeScriptEnum = 83
	UnicodeScriptJavanese              UnicodeScriptEnum = 84
	UnicodeScriptKaithi                UnicodeScriptEnum = 85
	UnicodeScriptLisu                  UnicodeScriptEnum = 86
	UnicodeScriptMeeteiMayek           UnicodeScriptEnum = 87
	UnicodeScriptOldSouthArabian       UnicodeScriptEnum = 88
	UnicodeScriptOldTurkic             UnicodeScriptEnum = 89
	UnicodeScriptSamaritan             UnicodeScriptEnum = 90
	UnicodeScriptTaiTham               UnicodeScriptEnum = 91
	UnicodeScriptTaiViet               UnicodeScriptEnum = 92
	UnicodeScriptBatak                 UnicodeScriptEnum = 93
	UnicodeScriptBrahmi                UnicodeScriptEnum = 94
	UnicodeScriptMandaic               UnicodeScriptEnum = 95
	UnicodeScriptChakma                UnicodeScriptEnum = 96
	UnicodeScriptMeroiticCursive       UnicodeScriptEnum = 97
	UnicodeScriptMeroiticHieroglyphs   UnicodeScriptEnum = 98
	UnicodeScriptMiao                  UnicodeScriptEnum = 99
	UnicodeScriptSharada               UnicodeScriptEnum = 100
	UnicodeScriptSoraSompeng           UnicodeScriptEnum = 101
	UnicodeScriptTakri                 UnicodeScriptEnum = 102
	UnicodeScriptBassaVah              UnicodeScriptEnum = 103
	UnicodeScriptCaucasianAlbanian     UnicodeScriptEnum = 104
	UnicodeScriptDuployan              UnicodeScriptEnum = 105
	UnicodeScriptElbasan               UnicodeScriptEnum = 106
	UnicodeScriptGrantha               UnicodeScriptEnum = 107
	UnicodeScriptKhojki                UnicodeScriptEnum = 108
	UnicodeScriptKhudawadi             UnicodeScriptEnum = 109
	UnicodeScriptLinearA               UnicodeScriptEnum = 110
	UnicodeScriptMahajani              UnicodeScriptEnum = 111
	UnicodeScriptManichaean            UnicodeScriptEnum = 112
	UnicodeScriptMendeKikakui          UnicodeScriptEnum = 113
	UnicodeScriptModi                  UnicodeScriptEnum = 114
	UnicodeScriptMro                   UnicodeScriptEnum = 115
	UnicodeScriptNabataean             UnicodeScriptEnum = 116
	UnicodeScriptOldNorthArabian       UnicodeScriptEnum = 117
	UnicodeScriptOldPermic             UnicodeScriptEnum = 118
	UnicodeScriptPahawhHmong           UnicodeScriptEnum = 119
	UnicodeScriptPalmyrene             UnicodeScriptEnum = 120
	UnicodeScriptPauCinHau             UnicodeScriptEnum = 121
	UnicodeScriptPsalterPahlavi        UnicodeScriptEnum = 122
	UnicodeScriptSiddham               UnicodeScriptEnum = 123
	UnicodeScriptTirhuta               UnicodeScriptEnum = 124
	UnicodeScriptWarangCiti            UnicodeScriptEnum = 125
	UnicodeScriptAhom                  UnicodeScriptEnum = 126
	UnicodeScriptAnatolianHieroglyphs  UnicodeScriptEnum = 127
	UnicodeScriptHatran                UnicodeScriptEnum = 128
	UnicodeScriptMultani               UnicodeScriptEnum = 129
	UnicodeScriptOldHungarian          UnicodeScriptEnum = 130
	UnicodeScriptSignwriting           UnicodeScriptEnum = 131
	UnicodeScriptAdlam                 UnicodeScriptEnum = 132
	UnicodeScriptBhaiksuki             UnicodeScriptEnum = 133
	UnicodeScriptMarchen               UnicodeScriptEnum = 134
	UnicodeScriptNewa                  UnicodeScriptEnum = 135
	UnicodeScriptOsage                 UnicodeScriptEnum = 136
	UnicodeScriptTangut                UnicodeScriptEnum = 137
	UnicodeScriptMasaramGondi          UnicodeScriptEnum = 138
	UnicodeScriptNushu                 UnicodeScriptEnum = 139
	UnicodeScriptSoyombo               UnicodeScriptEnum = 140
	UnicodeScriptZanabazarSquare       UnicodeScriptEnum = 141
	UnicodeScriptDogra                 UnicodeScriptEnum = 142
	UnicodeScriptGunjalaGondi          UnicodeScriptEnum = 143
	UnicodeScriptHanifiRohingya        UnicodeScriptEnum = 144
	UnicodeScriptMakasar               UnicodeScriptEnum = 145
	UnicodeScriptMedefaidrin           UnicodeScriptEnum = 146
	UnicodeScriptOldSogdian            UnicodeScriptEnum = 147
	UnicodeScriptSogdian               UnicodeScriptEnum = 148
)

func UnicodeScriptGetType() gi.GType {
	ret := _I.GetGType(121, "UnicodeScript")
	return ret
}

// Enum UnicodeType
type UnicodeTypeEnum int

const (
	UnicodeTypeControl            UnicodeTypeEnum = 0
	UnicodeTypeFormat             UnicodeTypeEnum = 1
	UnicodeTypeUnassigned         UnicodeTypeEnum = 2
	UnicodeTypePrivateUse         UnicodeTypeEnum = 3
	UnicodeTypeSurrogate          UnicodeTypeEnum = 4
	UnicodeTypeLowercaseLetter    UnicodeTypeEnum = 5
	UnicodeTypeModifierLetter     UnicodeTypeEnum = 6
	UnicodeTypeOtherLetter        UnicodeTypeEnum = 7
	UnicodeTypeTitlecaseLetter    UnicodeTypeEnum = 8
	UnicodeTypeUppercaseLetter    UnicodeTypeEnum = 9
	UnicodeTypeSpacingMark        UnicodeTypeEnum = 10
	UnicodeTypeEnclosingMark      UnicodeTypeEnum = 11
	UnicodeTypeNonSpacingMark     UnicodeTypeEnum = 12
	UnicodeTypeDecimalNumber      UnicodeTypeEnum = 13
	UnicodeTypeLetterNumber       UnicodeTypeEnum = 14
	UnicodeTypeOtherNumber        UnicodeTypeEnum = 15
	UnicodeTypeConnectPunctuation UnicodeTypeEnum = 16
	UnicodeTypeDashPunctuation    UnicodeTypeEnum = 17
	UnicodeTypeClosePunctuation   UnicodeTypeEnum = 18
	UnicodeTypeFinalPunctuation   UnicodeTypeEnum = 19
	UnicodeTypeInitialPunctuation UnicodeTypeEnum = 20
	UnicodeTypeOtherPunctuation   UnicodeTypeEnum = 21
	UnicodeTypeOpenPunctuation    UnicodeTypeEnum = 22
	UnicodeTypeCurrencySymbol     UnicodeTypeEnum = 23
	UnicodeTypeModifierSymbol     UnicodeTypeEnum = 24
	UnicodeTypeMathSymbol         UnicodeTypeEnum = 25
	UnicodeTypeOtherSymbol        UnicodeTypeEnum = 26
	UnicodeTypeLineSeparator      UnicodeTypeEnum = 27
	UnicodeTypeParagraphSeparator UnicodeTypeEnum = 28
	UnicodeTypeSpaceSeparator     UnicodeTypeEnum = 29
)

func UnicodeTypeGetType() gi.GType {
	ret := _I.GetGType(122, "UnicodeType")
	return ret
}

type UnixFDSourceFuncStruct struct {
	F_fd        int32
	F_condition IOConditionFlags
}

//export myUnixFDSourceFunc
func myUnixFDSourceFunc(fd C.gint32, condition C.GIOCondition, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := UnixFDSourceFuncStruct{
		F_fd:        int32(fd),
		F_condition: IOConditionFlags(condition),
	}
	fn(args)
}

// Enum UserDirectory
type UserDirectoryEnum int

const (
	UserDirectoryDirectoryDesktop     UserDirectoryEnum = 0
	UserDirectoryDirectoryDocuments   UserDirectoryEnum = 1
	UserDirectoryDirectoryDownload    UserDirectoryEnum = 2
	UserDirectoryDirectoryMusic       UserDirectoryEnum = 3
	UserDirectoryDirectoryPictures    UserDirectoryEnum = 4
	UserDirectoryDirectoryPublicShare UserDirectoryEnum = 5
	UserDirectoryDirectoryTemplates   UserDirectoryEnum = 6
	UserDirectoryDirectoryVideos      UserDirectoryEnum = 7
	UserDirectoryNDirectories         UserDirectoryEnum = 8
)

func UserDirectoryGetType() gi.GType {
	ret := _I.GetGType(123, "UserDirectory")
	return ret
}

// Struct Variant
type Variant struct {
	P unsafe.Pointer
}

func VariantGetType() gi.GType {
	ret := _I.GetGType(124, "Variant")
	return ret
}

// g_variant_new_array
// container is not nil, container is Variant
// is constructor
// arg 1 children lenArgIdx 2
func NewVariantArray(child_type VariantType, children gi.PointerArray, n_children uint64) (result Variant) {
	iv, err := _I.Get(657, "Variant", "new_array")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_child_type := gi.NewPointerArgument(child_type.P)
	arg_children := gi.NewPointerArgument(children.P)
	arg_n_children := gi.NewUint64Argument(n_children)
	args := []gi.Argument{arg_child_type, arg_children, arg_n_children}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_boolean
// container is not nil, container is Variant
// is constructor
func NewVariantBoolean(value bool) (result Variant) {
	iv, err := _I.Get(658, "Variant", "new_boolean")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewBoolArgument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_byte
// container is not nil, container is Variant
// is constructor
func NewVariantByte(value uint8) (result Variant) {
	iv, err := _I.Get(659, "Variant", "new_byte")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewUint8Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_bytestring
// container is not nil, container is Variant
// is constructor
func NewVariantBytestring(string gi.Uint8Array) (result Variant) {
	iv, err := _I.Get(660, "Variant", "new_bytestring")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_string := gi.NewPointerArgument(string.P)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_bytestring_array
// container is not nil, container is Variant
// is constructor
// arg 0 strv lenArgIdx 1
func NewVariantBytestringArray(strv gi.CStrArray, length int64) (result Variant) {
	iv, err := _I.Get(661, "Variant", "new_bytestring_array")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_strv := gi.NewPointerArgument(strv.P)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_strv, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_dict_entry
// container is not nil, container is Variant
// is constructor
func NewVariantDictEntry(key Variant, value Variant) (result Variant) {
	iv, err := _I.Get(662, "Variant", "new_dict_entry")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_key := gi.NewPointerArgument(key.P)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_double
// container is not nil, container is Variant
// is constructor
func NewVariantDouble(value float64) (result Variant) {
	iv, err := _I.Get(663, "Variant", "new_double")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewDoubleArgument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_fixed_array
// container is not nil, container is Variant
// is constructor
func NewVariantFixedArray(element_type VariantType, elements unsafe.Pointer, n_elements uint64, element_size uint64) (result Variant) {
	iv, err := _I.Get(664, "Variant", "new_fixed_array")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_element_type := gi.NewPointerArgument(element_type.P)
	arg_elements := gi.NewPointerArgument(elements)
	arg_n_elements := gi.NewUint64Argument(n_elements)
	arg_element_size := gi.NewUint64Argument(element_size)
	args := []gi.Argument{arg_element_type, arg_elements, arg_n_elements, arg_element_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_from_bytes
// container is not nil, container is Variant
// is constructor
func NewVariantFromBytes(type1 VariantType, bytes Bytes, trusted bool) (result Variant) {
	iv, err := _I.Get(665, "Variant", "new_from_bytes")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewPointerArgument(type1.P)
	arg_bytes := gi.NewPointerArgument(bytes.P)
	arg_trusted := gi.NewBoolArgument(trusted)
	args := []gi.Argument{arg_type1, arg_bytes, arg_trusted}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_from_data
// container is not nil, container is Variant
// is constructor
// arg 1 data lenArgIdx 2
func NewVariantFromData(type1 VariantType, data gi.Uint8Array, size uint64, trusted bool, notify int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer) (result Variant) {
	iv, err := _I.Get(666, "Variant", "new_from_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewPointerArgument(type1.P)
	arg_data := gi.NewPointerArgument(data.P)
	arg_size := gi.NewUint64Argument(size)
	arg_trusted := gi.NewBoolArgument(trusted)
	arg_notify := gi.NewIntArgument(notify) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_type1, arg_data, arg_size, arg_trusted, arg_notify, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_handle
// container is not nil, container is Variant
// is constructor
func NewVariantHandle(value int32) (result Variant) {
	iv, err := _I.Get(667, "Variant", "new_handle")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewInt32Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_int16
// container is not nil, container is Variant
// is constructor
func NewVariantInt16(value int16) (result Variant) {
	iv, err := _I.Get(668, "Variant", "new_int16")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewInt16Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_int32
// container is not nil, container is Variant
// is constructor
func NewVariantInt32(value int32) (result Variant) {
	iv, err := _I.Get(669, "Variant", "new_int32")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewInt32Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_int64
// container is not nil, container is Variant
// is constructor
func NewVariantInt64(value int64) (result Variant) {
	iv, err := _I.Get(670, "Variant", "new_int64")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewInt64Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_maybe
// container is not nil, container is Variant
// is constructor
func NewVariantMaybe(child_type VariantType, child Variant) (result Variant) {
	iv, err := _I.Get(671, "Variant", "new_maybe")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_child_type := gi.NewPointerArgument(child_type.P)
	arg_child := gi.NewPointerArgument(child.P)
	args := []gi.Argument{arg_child_type, arg_child}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_object_path
// container is not nil, container is Variant
// is constructor
func NewVariantObjectPath(object_path string) (result Variant) {
	iv, err := _I.Get(672, "Variant", "new_object_path")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_object_path := gi.CString(object_path)
	arg_object_path := gi.NewStringArgument(c_object_path)
	args := []gi.Argument{arg_object_path}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_object_path)
	result.P = ret.Pointer()
	return
}

// g_variant_new_objv
// container is not nil, container is Variant
// is constructor
// arg 0 strv lenArgIdx 1
func NewVariantObjv(strv gi.CStrArray, length int64) (result Variant) {
	iv, err := _I.Get(673, "Variant", "new_objv")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_strv := gi.NewPointerArgument(strv.P)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_strv, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_signature
// container is not nil, container is Variant
// is constructor
func NewVariantSignature(signature string) (result Variant) {
	iv, err := _I.Get(674, "Variant", "new_signature")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_signature := gi.CString(signature)
	arg_signature := gi.NewStringArgument(c_signature)
	args := []gi.Argument{arg_signature}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_signature)
	result.P = ret.Pointer()
	return
}

// g_variant_new_string
// container is not nil, container is Variant
// is constructor
func NewVariantString(string string) (result Variant) {
	iv, err := _I.Get(675, "Variant", "new_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result.P = ret.Pointer()
	return
}

// g_variant_new_strv
// container is not nil, container is Variant
// is constructor
// arg 0 strv lenArgIdx 1
func NewVariantStrv(strv gi.CStrArray, length int64) (result Variant) {
	iv, err := _I.Get(676, "Variant", "new_strv")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_strv := gi.NewPointerArgument(strv.P)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_strv, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_tuple
// container is not nil, container is Variant
// is constructor
// arg 0 children lenArgIdx 1
func NewVariantTuple(children gi.PointerArray, n_children uint64) (result Variant) {
	iv, err := _I.Get(677, "Variant", "new_tuple")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_children := gi.NewPointerArgument(children.P)
	arg_n_children := gi.NewUint64Argument(n_children)
	args := []gi.Argument{arg_children, arg_n_children}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_uint16
// container is not nil, container is Variant
// is constructor
func NewVariantUint16(value uint16) (result Variant) {
	iv, err := _I.Get(678, "Variant", "new_uint16")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewUint16Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_uint32
// container is not nil, container is Variant
// is constructor
func NewVariantUint32(value uint32) (result Variant) {
	iv, err := _I.Get(679, "Variant", "new_uint32")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewUint32Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_uint64
// container is not nil, container is Variant
// is constructor
func NewVariantUint64(value uint64) (result Variant) {
	iv, err := _I.Get(680, "Variant", "new_uint64")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewUint64Argument(value)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_new_variant
// container is not nil, container is Variant
// is constructor
func NewVariantVariant(value Variant) (result Variant) {
	iv, err := _I.Get(681, "Variant", "new_variant")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_byteswap
// container is not nil, container is Variant
// is method
func (v Variant) Byteswap() (result Variant) {
	iv, err := _I.Get(682, "Variant", "byteswap")
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

// g_variant_check_format_string
// container is not nil, container is Variant
// is method
func (v Variant) CheckFormatString(format_string string, copy_only bool) (result bool) {
	iv, err := _I.Get(683, "Variant", "check_format_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_format_string := gi.CString(format_string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_format_string := gi.NewStringArgument(c_format_string)
	arg_copy_only := gi.NewBoolArgument(copy_only)
	args := []gi.Argument{arg_v, arg_format_string, arg_copy_only}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_format_string)
	result = ret.Bool()
	return
}

// g_variant_classify
// container is not nil, container is Variant
// is method
func (v Variant) Classify() (result VariantClassEnum) {
	iv, err := _I.Get(684, "Variant", "classify")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = VariantClassEnum(ret.Int())
	return
}

// g_variant_compare
// container is not nil, container is Variant
// is method
func (v Variant) Compare(two Variant) (result int32) {
	iv, err := _I.Get(685, "Variant", "compare")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_two := gi.NewPointerArgument(two.P)
	args := []gi.Argument{arg_v, arg_two}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_variant_dup_bytestring
// container is not nil, container is Variant
// is method
// ret lenArgIdx 0
func (v Variant) DupBytestring() (result gi.Uint8Array) {
	iv, err := _I.Get(686, "Variant", "dup_bytestring")
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
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(length)}
	return
}

// g_variant_dup_bytestring_array
// container is not nil, container is Variant
// is method
// ret lenArgIdx 0
func (v Variant) DupBytestringArray() (result gi.CStrArray) {
	iv, err := _I.Get(687, "Variant", "dup_bytestring_array")
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
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: int(length)}
	return
}

// g_variant_dup_objv
// container is not nil, container is Variant
// is method
// ret lenArgIdx 0
func (v Variant) DupObjv() (result gi.CStrArray) {
	iv, err := _I.Get(688, "Variant", "dup_objv")
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
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_variant_dup_string
// container is not nil, container is Variant
// is method
func (v Variant) DupString() (result string, length uint64) {
	iv, err := _I.Get(689, "Variant", "dup_string")
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
	length = outArgs[0].Uint64()
	result = ret.String().Take()
	return
}

// g_variant_dup_strv
// container is not nil, container is Variant
// is method
// ret lenArgIdx 0
func (v Variant) DupStrv() (result gi.CStrArray) {
	iv, err := _I.Get(690, "Variant", "dup_strv")
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
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_variant_equal
// container is not nil, container is Variant
// is method
func (v Variant) Equal(two Variant) (result bool) {
	iv, err := _I.Get(691, "Variant", "equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_two := gi.NewPointerArgument(two.P)
	args := []gi.Argument{arg_v, arg_two}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_get_boolean
// container is not nil, container is Variant
// is method
func (v Variant) GetBoolean() (result bool) {
	iv, err := _I.Get(692, "Variant", "get_boolean")
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

// g_variant_get_byte
// container is not nil, container is Variant
// is method
func (v Variant) GetByte() (result uint8) {
	iv, err := _I.Get(693, "Variant", "get_byte")
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

// g_variant_get_bytestring
// container is not nil, container is Variant
// is method
func (v Variant) GetBytestring() (result gi.Uint8Array) {
	iv, err := _I.Get(694, "Variant", "get_bytestring")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(0)}
	return
}

// g_variant_get_bytestring_array
// container is not nil, container is Variant
// is method
// ret lenArgIdx 0
func (v Variant) GetBytestringArray() (result gi.CStrArray) {
	iv, err := _I.Get(695, "Variant", "get_bytestring_array")
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
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: int(length)}
	return
}

// g_variant_get_child_value
// container is not nil, container is Variant
// is method
func (v Variant) GetChildValue(index_ uint64) (result Variant) {
	iv, err := _I.Get(696, "Variant", "get_child_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewUint64Argument(index_)
	args := []gi.Argument{arg_v, arg_index_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_get_data
// container is not nil, container is Variant
// is method
func (v Variant) GetData() (result unsafe.Pointer) {
	iv, err := _I.Get(697, "Variant", "get_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_variant_get_data_as_bytes
// container is not nil, container is Variant
// is method
func (v Variant) GetDataAsBytes() (result Bytes) {
	iv, err := _I.Get(698, "Variant", "get_data_as_bytes")
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

// g_variant_get_double
// container is not nil, container is Variant
// is method
func (v Variant) GetDouble() (result float64) {
	iv, err := _I.Get(699, "Variant", "get_double")
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

// g_variant_get_handle
// container is not nil, container is Variant
// is method
func (v Variant) GetHandle() (result int32) {
	iv, err := _I.Get(700, "Variant", "get_handle")
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

// g_variant_get_int16
// container is not nil, container is Variant
// is method
func (v Variant) GetInt16() (result int16) {
	iv, err := _I.Get(701, "Variant", "get_int16")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int16()
	return
}

// g_variant_get_int32
// container is not nil, container is Variant
// is method
func (v Variant) GetInt32() (result int32) {
	iv, err := _I.Get(702, "Variant", "get_int32")
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

// g_variant_get_int64
// container is not nil, container is Variant
// is method
func (v Variant) GetInt64() (result int64) {
	iv, err := _I.Get(703, "Variant", "get_int64")
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

// g_variant_get_maybe
// container is not nil, container is Variant
// is method
func (v Variant) GetMaybe() (result Variant) {
	iv, err := _I.Get(704, "Variant", "get_maybe")
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

// g_variant_get_normal_form
// container is not nil, container is Variant
// is method
func (v Variant) GetNormalForm() (result Variant) {
	iv, err := _I.Get(705, "Variant", "get_normal_form")
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

// g_variant_get_objv
// container is not nil, container is Variant
// is method
// ret lenArgIdx 0
func (v Variant) GetObjv() (result gi.CStrArray) {
	iv, err := _I.Get(706, "Variant", "get_objv")
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
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_variant_get_size
// container is not nil, container is Variant
// is method
func (v Variant) GetSize() (result uint64) {
	iv, err := _I.Get(707, "Variant", "get_size")
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

// g_variant_get_string
// container is not nil, container is Variant
// is method
func (v Variant) GetString() (result string, length uint64) {
	iv, err := _I.Get(708, "Variant", "get_string")
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
	length = outArgs[0].Uint64()
	result = ret.String().Take()
	return
}

// g_variant_get_strv
// container is not nil, container is Variant
// is method
// ret lenArgIdx 0
func (v Variant) GetStrv() (result gi.CStrArray) {
	iv, err := _I.Get(709, "Variant", "get_strv")
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
	var length uint64
	_ = length
	length = outArgs[0].Uint64()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_variant_get_type
// container is not nil, container is Variant
// is method
func (v Variant) GetType() (result VariantType) {
	iv, err := _I.Get(710, "Variant", "get_type")
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

// g_variant_get_type_string
// container is not nil, container is Variant
// is method
func (v Variant) GetTypeString() (result string) {
	iv, err := _I.Get(711, "Variant", "get_type_string")
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

// g_variant_get_uint16
// container is not nil, container is Variant
// is method
func (v Variant) GetUint16() (result uint16) {
	iv, err := _I.Get(712, "Variant", "get_uint16")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint16()
	return
}

// g_variant_get_uint32
// container is not nil, container is Variant
// is method
func (v Variant) GetUint32() (result uint32) {
	iv, err := _I.Get(713, "Variant", "get_uint32")
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

// g_variant_get_uint64
// container is not nil, container is Variant
// is method
func (v Variant) GetUint64() (result uint64) {
	iv, err := _I.Get(714, "Variant", "get_uint64")
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

// g_variant_get_variant
// container is not nil, container is Variant
// is method
func (v Variant) GetVariant() (result Variant) {
	iv, err := _I.Get(715, "Variant", "get_variant")
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

// g_variant_hash
// container is not nil, container is Variant
// is method
func (v Variant) Hash() (result uint32) {
	iv, err := _I.Get(716, "Variant", "hash")
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

// g_variant_is_container
// container is not nil, container is Variant
// is method
func (v Variant) IsContainer() (result bool) {
	iv, err := _I.Get(717, "Variant", "is_container")
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

// g_variant_is_floating
// container is not nil, container is Variant
// is method
func (v Variant) IsFloating() (result bool) {
	iv, err := _I.Get(718, "Variant", "is_floating")
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

// g_variant_is_normal_form
// container is not nil, container is Variant
// is method
func (v Variant) IsNormalForm() (result bool) {
	iv, err := _I.Get(719, "Variant", "is_normal_form")
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

// g_variant_is_of_type
// container is not nil, container is Variant
// is method
func (v Variant) IsOfType(type1 VariantType) (result bool) {
	iv, err := _I.Get(720, "Variant", "is_of_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewPointerArgument(type1.P)
	args := []gi.Argument{arg_v, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_lookup_value
// container is not nil, container is Variant
// is method
func (v Variant) LookupValue(key string, expected_type VariantType) (result Variant) {
	iv, err := _I.Get(721, "Variant", "lookup_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	arg_expected_type := gi.NewPointerArgument(expected_type.P)
	args := []gi.Argument{arg_v, arg_key, arg_expected_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_key)
	result.P = ret.Pointer()
	return
}

// g_variant_n_children
// container is not nil, container is Variant
// is method
func (v Variant) NChildren() (result uint64) {
	iv, err := _I.Get(722, "Variant", "n_children")
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

// g_variant_print
// container is not nil, container is Variant
// is method
func (v Variant) Print(type_annotate bool) (result string) {
	iv, err := _I.Get(723, "Variant", "print")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type_annotate := gi.NewBoolArgument(type_annotate)
	args := []gi.Argument{arg_v, arg_type_annotate}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_variant_ref
// container is not nil, container is Variant
// is method
func (v Variant) Ref() (result Variant) {
	iv, err := _I.Get(724, "Variant", "ref")
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

// g_variant_ref_sink
// container is not nil, container is Variant
// is method
func (v Variant) RefSink() (result Variant) {
	iv, err := _I.Get(725, "Variant", "ref_sink")
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

// g_variant_store
// container is not nil, container is Variant
// is method
func (v Variant) Store(data unsafe.Pointer) {
	iv, err := _I.Get(726, "Variant", "store")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_data}
	iv.Call(args, nil, nil)
}

// g_variant_take_ref
// container is not nil, container is Variant
// is method
func (v Variant) TakeRef() (result Variant) {
	iv, err := _I.Get(727, "Variant", "take_ref")
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

// g_variant_unref
// container is not nil, container is Variant
// is method
func (v Variant) Unref() {
	iv, err := _I.Get(728, "Variant", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_variant_is_object_path
// container is not nil, container is Variant
// is method
// arg0Type tag: utf8, isPtr: true
func VariantIsObjectPath1(string string) (result bool) {
	iv, err := _I.Get(729, "Variant", "is_object_path")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_variant_is_signature
// container is not nil, container is Variant
// is method
// arg0Type tag: utf8, isPtr: true
func VariantIsSignature1(string string) (result bool) {
	iv, err := _I.Get(730, "Variant", "is_signature")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_variant_parse
// container is not nil, container is Variant
// is method
// arg0Type tag: interface, isPtr: true
func VariantParse1(type1 VariantType, text string, limit string, endptr string) (result Variant, err error) {
	iv, err := _I.Get(731, "Variant", "parse")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_text := gi.CString(text)
	c_limit := gi.CString(limit)
	c_endptr := gi.CString(endptr)
	arg_type1 := gi.NewPointerArgument(type1.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_limit := gi.NewStringArgument(c_limit)
	arg_endptr := gi.NewStringArgument(c_endptr)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_type1, arg_text, arg_limit, arg_endptr, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_text)
	gi.Free(c_limit)
	gi.Free(c_endptr)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// g_variant_parse_error_print_context
// container is not nil, container is Variant
// is method
// arg0Type tag: error, isPtr: true
func VariantParseErrorPrintContext1(error Error, source_str string) (result string) {
	iv, err := _I.Get(732, "Variant", "parse_error_print_context")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_source_str := gi.CString(source_str)
	arg_error := gi.NewPointerArgument(error.P)
	arg_source_str := gi.NewStringArgument(c_source_str)
	args := []gi.Argument{arg_error, arg_source_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_source_str)
	result = ret.String().Take()
	return
}

// Struct VariantBuilder
type VariantBuilder struct {
	P unsafe.Pointer
}

func VariantBuilderGetType() gi.GType {
	ret := _I.GetGType(125, "VariantBuilder")
	return ret
}

// g_variant_builder_new
// container is not nil, container is VariantBuilder
// is constructor
func NewVariantBuilder(type1 VariantType) (result VariantBuilder) {
	iv, err := _I.Get(735, "VariantBuilder", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewPointerArgument(type1.P)
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_builder_add_value
// container is not nil, container is VariantBuilder
// is method
func (v VariantBuilder) AddValue(value Variant) {
	iv, err := _I.Get(736, "VariantBuilder", "add_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_value}
	iv.Call(args, nil, nil)
}

// g_variant_builder_close
// container is not nil, container is VariantBuilder
// is method
func (v VariantBuilder) Close() {
	iv, err := _I.Get(737, "VariantBuilder", "close")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_variant_builder_end
// container is not nil, container is VariantBuilder
// is method
func (v VariantBuilder) End() (result Variant) {
	iv, err := _I.Get(738, "VariantBuilder", "end")
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

// g_variant_builder_open
// container is not nil, container is VariantBuilder
// is method
func (v VariantBuilder) Open(type1 VariantType) {
	iv, err := _I.Get(739, "VariantBuilder", "open")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewPointerArgument(type1.P)
	args := []gi.Argument{arg_v, arg_type1}
	iv.Call(args, nil, nil)
}

// g_variant_builder_ref
// container is not nil, container is VariantBuilder
// is method
func (v VariantBuilder) Ref() (result VariantBuilder) {
	iv, err := _I.Get(740, "VariantBuilder", "ref")
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

// g_variant_builder_unref
// container is not nil, container is VariantBuilder
// is method
func (v VariantBuilder) Unref() {
	iv, err := _I.Get(741, "VariantBuilder", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum VariantClass
type VariantClassEnum int

const (
	VariantClassBoolean    VariantClassEnum = 98
	VariantClassByte       VariantClassEnum = 121
	VariantClassInt16      VariantClassEnum = 110
	VariantClassUint16     VariantClassEnum = 113
	VariantClassInt32      VariantClassEnum = 105
	VariantClassUint32     VariantClassEnum = 117
	VariantClassInt64      VariantClassEnum = 120
	VariantClassUint64     VariantClassEnum = 116
	VariantClassHandle     VariantClassEnum = 104
	VariantClassDouble     VariantClassEnum = 100
	VariantClassString     VariantClassEnum = 115
	VariantClassObjectPath VariantClassEnum = 111
	VariantClassSignature  VariantClassEnum = 103
	VariantClassVariant    VariantClassEnum = 118
	VariantClassMaybe      VariantClassEnum = 109
	VariantClassArray      VariantClassEnum = 97
	VariantClassTuple      VariantClassEnum = 40
	VariantClassDictEntry  VariantClassEnum = 123
)

func VariantClassGetType() gi.GType {
	ret := _I.GetGType(126, "VariantClass")
	return ret
}

// Struct VariantDict
type VariantDict struct {
	P unsafe.Pointer
}

func VariantDictGetType() gi.GType {
	ret := _I.GetGType(127, "VariantDict")
	return ret
}

// g_variant_dict_new
// container is not nil, container is VariantDict
// is constructor
func NewVariantDict(from_asv Variant) (result VariantDict) {
	iv, err := _I.Get(742, "VariantDict", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_from_asv := gi.NewPointerArgument(from_asv.P)
	args := []gi.Argument{arg_from_asv}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_dict_clear
// container is not nil, container is VariantDict
// is method
func (v VariantDict) Clear() {
	iv, err := _I.Get(743, "VariantDict", "clear")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_variant_dict_contains
// container is not nil, container is VariantDict
// is method
func (v VariantDict) Contains(key string) (result bool) {
	iv, err := _I.Get(744, "VariantDict", "contains")
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

// g_variant_dict_end
// container is not nil, container is VariantDict
// is method
func (v VariantDict) End() (result Variant) {
	iv, err := _I.Get(745, "VariantDict", "end")
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

// g_variant_dict_insert_value
// container is not nil, container is VariantDict
// is method
func (v VariantDict) InsertValue(key string, value Variant) {
	iv, err := _I.Get(746, "VariantDict", "insert_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_key, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_key)
}

// g_variant_dict_lookup_value
// container is not nil, container is VariantDict
// is method
func (v VariantDict) LookupValue(key string, expected_type VariantType) (result Variant) {
	iv, err := _I.Get(747, "VariantDict", "lookup_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	arg_expected_type := gi.NewPointerArgument(expected_type.P)
	args := []gi.Argument{arg_v, arg_key, arg_expected_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_key)
	result.P = ret.Pointer()
	return
}

// g_variant_dict_ref
// container is not nil, container is VariantDict
// is method
func (v VariantDict) Ref() (result VariantDict) {
	iv, err := _I.Get(748, "VariantDict", "ref")
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

// g_variant_dict_remove
// container is not nil, container is VariantDict
// is method
func (v VariantDict) Remove(key string) (result bool) {
	iv, err := _I.Get(749, "VariantDict", "remove")
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

// g_variant_dict_unref
// container is not nil, container is VariantDict
// is method
func (v VariantDict) Unref() {
	iv, err := _I.Get(750, "VariantDict", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Enum VariantParseError
type VariantParseErrorEnum int

const (
	VariantParseErrorFailed                     VariantParseErrorEnum = 0
	VariantParseErrorBasicTypeExpected          VariantParseErrorEnum = 1
	VariantParseErrorCannotInferType            VariantParseErrorEnum = 2
	VariantParseErrorDefiniteTypeExpected       VariantParseErrorEnum = 3
	VariantParseErrorInputNotAtEnd              VariantParseErrorEnum = 4
	VariantParseErrorInvalidCharacter           VariantParseErrorEnum = 5
	VariantParseErrorInvalidFormatString        VariantParseErrorEnum = 6
	VariantParseErrorInvalidObjectPath          VariantParseErrorEnum = 7
	VariantParseErrorInvalidSignature           VariantParseErrorEnum = 8
	VariantParseErrorInvalidTypeString          VariantParseErrorEnum = 9
	VariantParseErrorNoCommonType               VariantParseErrorEnum = 10
	VariantParseErrorNumberOutOfRange           VariantParseErrorEnum = 11
	VariantParseErrorNumberTooBig               VariantParseErrorEnum = 12
	VariantParseErrorTypeError                  VariantParseErrorEnum = 13
	VariantParseErrorUnexpectedToken            VariantParseErrorEnum = 14
	VariantParseErrorUnknownKeyword             VariantParseErrorEnum = 15
	VariantParseErrorUnterminatedStringConstant VariantParseErrorEnum = 16
	VariantParseErrorValueExpected              VariantParseErrorEnum = 17
)

func VariantParseErrorGetType() gi.GType {
	ret := _I.GetGType(128, "VariantParseError")
	return ret
}

// Struct VariantType
type VariantType struct {
	P unsafe.Pointer
}

func VariantTypeGetType() gi.GType {
	ret := _I.GetGType(129, "VariantType")
	return ret
}

// g_variant_type_new
// container is not nil, container is VariantType
// is constructor
func NewVariantType(type_string string) (result VariantType) {
	iv, err := _I.Get(751, "VariantType", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_type_string := gi.CString(type_string)
	arg_type_string := gi.NewStringArgument(c_type_string)
	args := []gi.Argument{arg_type_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_type_string)
	result.P = ret.Pointer()
	return
}

// g_variant_type_new_array
// container is not nil, container is VariantType
// is constructor
func NewVariantTypeArray(element VariantType) (result VariantType) {
	iv, err := _I.Get(752, "VariantType", "new_array")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_element := gi.NewPointerArgument(element.P)
	args := []gi.Argument{arg_element}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_type_new_dict_entry
// container is not nil, container is VariantType
// is constructor
func NewVariantTypeDictEntry(key VariantType, value VariantType) (result VariantType) {
	iv, err := _I.Get(753, "VariantType", "new_dict_entry")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_key := gi.NewPointerArgument(key.P)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_type_new_maybe
// container is not nil, container is VariantType
// is constructor
func NewVariantTypeMaybe(element VariantType) (result VariantType) {
	iv, err := _I.Get(754, "VariantType", "new_maybe")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_element := gi.NewPointerArgument(element.P)
	args := []gi.Argument{arg_element}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_type_new_tuple
// container is not nil, container is VariantType
// is constructor
// arg 0 items lenArgIdx 1
func NewVariantTypeTuple(items gi.PointerArray, length int32) (result VariantType) {
	iv, err := _I.Get(755, "VariantType", "new_tuple")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_items := gi.NewPointerArgument(items.P)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_items, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_variant_type_copy
// container is not nil, container is VariantType
// is method
func (v VariantType) Copy() (result VariantType) {
	iv, err := _I.Get(756, "VariantType", "copy")
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

// g_variant_type_dup_string
// container is not nil, container is VariantType
// is method
func (v VariantType) DupString() (result string) {
	iv, err := _I.Get(757, "VariantType", "dup_string")
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

// g_variant_type_element
// container is not nil, container is VariantType
// is method
func (v VariantType) Element() (result VariantType) {
	iv, err := _I.Get(758, "VariantType", "element")
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

// g_variant_type_equal
// container is not nil, container is VariantType
// is method
func (v VariantType) Equal(type2 VariantType) (result bool) {
	iv, err := _I.Get(759, "VariantType", "equal")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type2 := gi.NewPointerArgument(type2.P)
	args := []gi.Argument{arg_v, arg_type2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_type_first
// container is not nil, container is VariantType
// is method
func (v VariantType) First() (result VariantType) {
	iv, err := _I.Get(760, "VariantType", "first")
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

// g_variant_type_free
// container is not nil, container is VariantType
// is method
func (v VariantType) Free() {
	iv, err := _I.Get(761, "VariantType", "free")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_variant_type_get_string_length
// container is not nil, container is VariantType
// is method
func (v VariantType) GetStringLength() (result uint64) {
	iv, err := _I.Get(762, "VariantType", "get_string_length")
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

// g_variant_type_hash
// container is not nil, container is VariantType
// is method
func (v VariantType) Hash() (result uint32) {
	iv, err := _I.Get(763, "VariantType", "hash")
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

// g_variant_type_is_array
// container is not nil, container is VariantType
// is method
func (v VariantType) IsArray() (result bool) {
	iv, err := _I.Get(764, "VariantType", "is_array")
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

// g_variant_type_is_basic
// container is not nil, container is VariantType
// is method
func (v VariantType) IsBasic() (result bool) {
	iv, err := _I.Get(765, "VariantType", "is_basic")
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

// g_variant_type_is_container
// container is not nil, container is VariantType
// is method
func (v VariantType) IsContainer() (result bool) {
	iv, err := _I.Get(766, "VariantType", "is_container")
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

// g_variant_type_is_definite
// container is not nil, container is VariantType
// is method
func (v VariantType) IsDefinite() (result bool) {
	iv, err := _I.Get(767, "VariantType", "is_definite")
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

// g_variant_type_is_dict_entry
// container is not nil, container is VariantType
// is method
func (v VariantType) IsDictEntry() (result bool) {
	iv, err := _I.Get(768, "VariantType", "is_dict_entry")
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

// g_variant_type_is_maybe
// container is not nil, container is VariantType
// is method
func (v VariantType) IsMaybe() (result bool) {
	iv, err := _I.Get(769, "VariantType", "is_maybe")
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

// g_variant_type_is_subtype_of
// container is not nil, container is VariantType
// is method
func (v VariantType) IsSubtypeOf(supertype VariantType) (result bool) {
	iv, err := _I.Get(770, "VariantType", "is_subtype_of")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_supertype := gi.NewPointerArgument(supertype.P)
	args := []gi.Argument{arg_v, arg_supertype}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_variant_type_is_tuple
// container is not nil, container is VariantType
// is method
func (v VariantType) IsTuple() (result bool) {
	iv, err := _I.Get(771, "VariantType", "is_tuple")
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

// g_variant_type_is_variant
// container is not nil, container is VariantType
// is method
func (v VariantType) IsVariant() (result bool) {
	iv, err := _I.Get(772, "VariantType", "is_variant")
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

// g_variant_type_key
// container is not nil, container is VariantType
// is method
func (v VariantType) Key() (result VariantType) {
	iv, err := _I.Get(773, "VariantType", "key")
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

// g_variant_type_n_items
// container is not nil, container is VariantType
// is method
func (v VariantType) NItems() (result uint64) {
	iv, err := _I.Get(774, "VariantType", "n_items")
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

// g_variant_type_next
// container is not nil, container is VariantType
// is method
func (v VariantType) Next() (result VariantType) {
	iv, err := _I.Get(775, "VariantType", "next")
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

// g_variant_type_value
// container is not nil, container is VariantType
// is method
func (v VariantType) Value() (result VariantType) {
	iv, err := _I.Get(776, "VariantType", "value")
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

// g_variant_type_checked_
// container is not nil, container is VariantType
// is method
// arg0Type tag: utf8, isPtr: true
func VariantTypeChecked_1(arg0 string) (result VariantType) {
	iv, err := _I.Get(777, "VariantType", "checked_")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg0 := gi.CString(arg0)
	arg_arg0 := gi.NewStringArgument(c_arg0)
	args := []gi.Argument{arg_arg0}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_arg0)
	result.P = ret.Pointer()
	return
}

// g_variant_type_string_get_depth_
// container is not nil, container is VariantType
// is method
// arg0Type tag: utf8, isPtr: true
func VariantTypeStringGetDepth_1(type_string string) (result uint64) {
	iv, err := _I.Get(778, "VariantType", "string_get_depth_")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_type_string := gi.CString(type_string)
	arg_type_string := gi.NewStringArgument(c_type_string)
	args := []gi.Argument{arg_type_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_type_string)
	result = ret.Uint64()
	return
}

// g_variant_type_string_is_valid
// container is not nil, container is VariantType
// is method
// arg0Type tag: utf8, isPtr: true
func VariantTypeStringIsValid1(type_string string) (result bool) {
	iv, err := _I.Get(779, "VariantType", "string_is_valid")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_type_string := gi.CString(type_string)
	arg_type_string := gi.NewStringArgument(c_type_string)
	args := []gi.Argument{arg_type_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_type_string)
	result = ret.Bool()
	return
}

// g_variant_type_string_scan
// container is not nil, container is VariantType
// is method
// arg0Type tag: utf8, isPtr: true
func VariantTypeStringScan1(string string, limit string) (result bool, endptr string) {
	iv, err := _I.Get(780, "VariantType", "string_scan")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_string := gi.CString(string)
	c_limit := gi.CString(limit)
	arg_string := gi.NewStringArgument(c_string)
	arg_limit := gi.NewStringArgument(c_limit)
	arg_endptr := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_string, arg_limit, arg_endptr}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_string)
	gi.Free(c_limit)
	endptr = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// ignore callback VoidFunc
// g_access
// container is nil
func Access(filename string, mode int32) (result int32) {
	iv, err := _I.Get(781, "access", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_mode := gi.NewInt32Argument(mode)
	args := []gi.Argument{arg_filename, arg_mode}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	result = ret.Int32()
	return
}

// g_ascii_digit_value
// container is nil
func AsciiDigitValue(c int8) (result int32) {
	iv, err := _I.Get(782, "ascii_digit_value", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewInt8Argument(c)
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_ascii_dtostr
// container is nil
func AsciiDtostr(buffer string, buf_len int32, d float64) (result string) {
	iv, err := _I.Get(783, "ascii_dtostr", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_buffer := gi.CString(buffer)
	arg_buffer := gi.NewStringArgument(c_buffer)
	arg_buf_len := gi.NewInt32Argument(buf_len)
	arg_d := gi.NewDoubleArgument(d)
	args := []gi.Argument{arg_buffer, arg_buf_len, arg_d}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_buffer)
	result = ret.String().Take()
	return
}

// g_ascii_formatd
// container is nil
func AsciiFormatd(buffer string, buf_len int32, format string, d float64) (result string) {
	iv, err := _I.Get(784, "ascii_formatd", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_buffer := gi.CString(buffer)
	c_format := gi.CString(format)
	arg_buffer := gi.NewStringArgument(c_buffer)
	arg_buf_len := gi.NewInt32Argument(buf_len)
	arg_format := gi.NewStringArgument(c_format)
	arg_d := gi.NewDoubleArgument(d)
	args := []gi.Argument{arg_buffer, arg_buf_len, arg_format, arg_d}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_buffer)
	gi.Free(c_format)
	result = ret.String().Take()
	return
}

// g_ascii_strcasecmp
// container is nil
func AsciiStrcasecmp(s1 string, s2 string) (result int32) {
	iv, err := _I.Get(785, "ascii_strcasecmp", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_s1 := gi.CString(s1)
	c_s2 := gi.CString(s2)
	arg_s1 := gi.NewStringArgument(c_s1)
	arg_s2 := gi.NewStringArgument(c_s2)
	args := []gi.Argument{arg_s1, arg_s2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_s1)
	gi.Free(c_s2)
	result = ret.Int32()
	return
}

// g_ascii_strdown
// container is nil
func AsciiStrdown(str string, len1 int64) (result string) {
	iv, err := _I.Get(786, "ascii_strdown", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_ascii_string_to_signed
// container is nil
func AsciiStringToSigned(str string, base uint32, min int64, max int64) (result bool, out_num int64, err error) {
	iv, err := _I.Get(787, "ascii_string_to_signed", "")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_base := gi.NewUint32Argument(base)
	arg_min := gi.NewInt64Argument(min)
	arg_max := gi.NewInt64Argument(max)
	arg_out_num := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_str, arg_base, arg_min, arg_max, arg_out_num, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_str)
	err = gi.ToError(outArgs[1].Pointer())
	out_num = outArgs[0].Int64()
	result = ret.Bool()
	return
}

// g_ascii_string_to_unsigned
// container is nil
func AsciiStringToUnsigned(str string, base uint32, min uint64, max uint64) (result bool, out_num uint64, err error) {
	iv, err := _I.Get(788, "ascii_string_to_unsigned", "")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_base := gi.NewUint32Argument(base)
	arg_min := gi.NewUint64Argument(min)
	arg_max := gi.NewUint64Argument(max)
	arg_out_num := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_str, arg_base, arg_min, arg_max, arg_out_num, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_str)
	err = gi.ToError(outArgs[1].Pointer())
	out_num = outArgs[0].Uint64()
	result = ret.Bool()
	return
}

// g_ascii_strncasecmp
// container is nil
func AsciiStrncasecmp(s1 string, s2 string, n uint64) (result int32) {
	iv, err := _I.Get(789, "ascii_strncasecmp", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_s1 := gi.CString(s1)
	c_s2 := gi.CString(s2)
	arg_s1 := gi.NewStringArgument(c_s1)
	arg_s2 := gi.NewStringArgument(c_s2)
	arg_n := gi.NewUint64Argument(n)
	args := []gi.Argument{arg_s1, arg_s2, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_s1)
	gi.Free(c_s2)
	result = ret.Int32()
	return
}

// g_ascii_strtod
// container is nil
func AsciiStrtod(nptr string) (result float64, endptr string) {
	iv, err := _I.Get(790, "ascii_strtod", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_nptr := gi.CString(nptr)
	arg_nptr := gi.NewStringArgument(c_nptr)
	arg_endptr := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_nptr, arg_endptr}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_nptr)
	endptr = outArgs[0].String().Take()
	result = ret.Double()
	return
}

// g_ascii_strtoll
// container is nil
func AsciiStrtoll(nptr string, base uint32) (result int64, endptr string) {
	iv, err := _I.Get(791, "ascii_strtoll", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_nptr := gi.CString(nptr)
	arg_nptr := gi.NewStringArgument(c_nptr)
	arg_endptr := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_base := gi.NewUint32Argument(base)
	args := []gi.Argument{arg_nptr, arg_endptr, arg_base}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_nptr)
	endptr = outArgs[0].String().Take()
	result = ret.Int64()
	return
}

// g_ascii_strtoull
// container is nil
func AsciiStrtoull(nptr string, base uint32) (result uint64, endptr string) {
	iv, err := _I.Get(792, "ascii_strtoull", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_nptr := gi.CString(nptr)
	arg_nptr := gi.NewStringArgument(c_nptr)
	arg_endptr := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_base := gi.NewUint32Argument(base)
	args := []gi.Argument{arg_nptr, arg_endptr, arg_base}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_nptr)
	endptr = outArgs[0].String().Take()
	result = ret.Uint64()
	return
}

// g_ascii_strup
// container is nil
func AsciiStrup(str string, len1 int64) (result string) {
	iv, err := _I.Get(793, "ascii_strup", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_ascii_tolower
// container is nil
func AsciiTolower(c int8) (result int8) {
	iv, err := _I.Get(794, "ascii_tolower", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewInt8Argument(c)
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int8()
	return
}

// g_ascii_toupper
// container is nil
func AsciiToupper(c int8) (result int8) {
	iv, err := _I.Get(795, "ascii_toupper", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewInt8Argument(c)
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int8()
	return
}

// g_ascii_xdigit_value
// container is nil
func AsciiXdigitValue(c int8) (result int32) {
	iv, err := _I.Get(796, "ascii_xdigit_value", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewInt8Argument(c)
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_assert_warning
// container is nil
func AssertWarning(log_domain string, file string, line int32, pretty_function string, expression string) {
	iv, err := _I.Get(797, "assert_warning", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_log_domain := gi.CString(log_domain)
	c_file := gi.CString(file)
	c_pretty_function := gi.CString(pretty_function)
	c_expression := gi.CString(expression)
	arg_log_domain := gi.NewStringArgument(c_log_domain)
	arg_file := gi.NewStringArgument(c_file)
	arg_line := gi.NewInt32Argument(line)
	arg_pretty_function := gi.NewStringArgument(c_pretty_function)
	arg_expression := gi.NewStringArgument(c_expression)
	args := []gi.Argument{arg_log_domain, arg_file, arg_line, arg_pretty_function, arg_expression}
	iv.Call(args, nil, nil)
	gi.Free(c_log_domain)
	gi.Free(c_file)
	gi.Free(c_pretty_function)
	gi.Free(c_expression)
}

// g_assertion_message
// container is nil
func AssertionMessage(domain string, file string, line int32, func1 string, message string) {
	iv, err := _I.Get(798, "assertion_message", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_file := gi.CString(file)
	c_func1 := gi.CString(func1)
	c_message := gi.CString(message)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_file := gi.NewStringArgument(c_file)
	arg_line := gi.NewInt32Argument(line)
	arg_func1 := gi.NewStringArgument(c_func1)
	arg_message := gi.NewStringArgument(c_message)
	args := []gi.Argument{arg_domain, arg_file, arg_line, arg_func1, arg_message}
	iv.Call(args, nil, nil)
	gi.Free(c_domain)
	gi.Free(c_file)
	gi.Free(c_func1)
	gi.Free(c_message)
}

// g_assertion_message_cmpstr
// container is nil
func AssertionMessageCmpstr(domain string, file string, line int32, func1 string, expr string, arg1 string, cmp string, arg2 string) {
	iv, err := _I.Get(799, "assertion_message_cmpstr", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_file := gi.CString(file)
	c_func1 := gi.CString(func1)
	c_expr := gi.CString(expr)
	c_arg1 := gi.CString(arg1)
	c_cmp := gi.CString(cmp)
	c_arg2 := gi.CString(arg2)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_file := gi.NewStringArgument(c_file)
	arg_line := gi.NewInt32Argument(line)
	arg_func1 := gi.NewStringArgument(c_func1)
	arg_expr := gi.NewStringArgument(c_expr)
	arg_arg1 := gi.NewStringArgument(c_arg1)
	arg_cmp := gi.NewStringArgument(c_cmp)
	arg_arg2 := gi.NewStringArgument(c_arg2)
	args := []gi.Argument{arg_domain, arg_file, arg_line, arg_func1, arg_expr, arg_arg1, arg_cmp, arg_arg2}
	iv.Call(args, nil, nil)
	gi.Free(c_domain)
	gi.Free(c_file)
	gi.Free(c_func1)
	gi.Free(c_expr)
	gi.Free(c_arg1)
	gi.Free(c_cmp)
	gi.Free(c_arg2)
}

// g_assertion_message_error
// container is nil
func AssertionMessageError(domain string, file string, line int32, func1 string, expr string, error Error, error_domain uint32, error_code int32) {
	iv, err := _I.Get(800, "assertion_message_error", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_file := gi.CString(file)
	c_func1 := gi.CString(func1)
	c_expr := gi.CString(expr)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_file := gi.NewStringArgument(c_file)
	arg_line := gi.NewInt32Argument(line)
	arg_func1 := gi.NewStringArgument(c_func1)
	arg_expr := gi.NewStringArgument(c_expr)
	arg_error := gi.NewPointerArgument(error.P)
	arg_error_domain := gi.NewUint32Argument(error_domain)
	arg_error_code := gi.NewInt32Argument(error_code)
	args := []gi.Argument{arg_domain, arg_file, arg_line, arg_func1, arg_expr, arg_error, arg_error_domain, arg_error_code}
	iv.Call(args, nil, nil)
	gi.Free(c_domain)
	gi.Free(c_file)
	gi.Free(c_func1)
	gi.Free(c_expr)
}

// g_atexit
// container is nil
func Atexit(func1 int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(801, "atexit", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_func1 := gi.NewIntArgument(func1) /*TODO*/
	args := []gi.Argument{arg_func1}
	iv.Call(args, nil, nil)
}

// g_atomic_int_add
// container is nil
func AtomicIntAdd(atomic int32, val int32) (result int32) {
	iv, err := _I.Get(802, "atomic_int_add", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewInt32Argument(atomic)
	arg_val := gi.NewInt32Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_atomic_int_and
// container is nil
func AtomicIntAnd(atomic uint32, val uint32) (result uint32) {
	iv, err := _I.Get(803, "atomic_int_and", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewUint32Argument(atomic)
	arg_val := gi.NewUint32Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_atomic_int_compare_and_exchange
// container is nil
func AtomicIntCompareAndExchange(atomic int32, oldval int32, newval int32) (result bool) {
	iv, err := _I.Get(804, "atomic_int_compare_and_exchange", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewInt32Argument(atomic)
	arg_oldval := gi.NewInt32Argument(oldval)
	arg_newval := gi.NewInt32Argument(newval)
	args := []gi.Argument{arg_atomic, arg_oldval, arg_newval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_atomic_int_dec_and_test
// container is nil
func AtomicIntDecAndTest(atomic int32) (result bool) {
	iv, err := _I.Get(805, "atomic_int_dec_and_test", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewInt32Argument(atomic)
	args := []gi.Argument{arg_atomic}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_atomic_int_exchange_and_add
// container is nil
func AtomicIntExchangeAndAdd(atomic int32, val int32) (result int32) {
	iv, err := _I.Get(806, "atomic_int_exchange_and_add", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewInt32Argument(atomic)
	arg_val := gi.NewInt32Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_atomic_int_get
// container is nil
func AtomicIntGet(atomic int32) (result int32) {
	iv, err := _I.Get(807, "atomic_int_get", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewInt32Argument(atomic)
	args := []gi.Argument{arg_atomic}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_atomic_int_inc
// container is nil
func AtomicIntInc(atomic int32) {
	iv, err := _I.Get(808, "atomic_int_inc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewInt32Argument(atomic)
	args := []gi.Argument{arg_atomic}
	iv.Call(args, nil, nil)
}

// g_atomic_int_or
// container is nil
func AtomicIntOr(atomic uint32, val uint32) (result uint32) {
	iv, err := _I.Get(809, "atomic_int_or", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewUint32Argument(atomic)
	arg_val := gi.NewUint32Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_atomic_int_set
// container is nil
func AtomicIntSet(atomic int32, newval int32) {
	iv, err := _I.Get(810, "atomic_int_set", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewInt32Argument(atomic)
	arg_newval := gi.NewInt32Argument(newval)
	args := []gi.Argument{arg_atomic, arg_newval}
	iv.Call(args, nil, nil)
}

// g_atomic_int_xor
// container is nil
func AtomicIntXor(atomic uint32, val uint32) (result uint32) {
	iv, err := _I.Get(811, "atomic_int_xor", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewUint32Argument(atomic)
	arg_val := gi.NewUint32Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_atomic_pointer_add
// container is nil
func AtomicPointerAdd(atomic unsafe.Pointer, val int64) (result int64) {
	iv, err := _I.Get(812, "atomic_pointer_add", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewPointerArgument(atomic)
	arg_val := gi.NewInt64Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_atomic_pointer_and
// container is nil
func AtomicPointerAnd(atomic unsafe.Pointer, val uint64) (result uint64) {
	iv, err := _I.Get(813, "atomic_pointer_and", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewPointerArgument(atomic)
	arg_val := gi.NewUint64Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_atomic_pointer_compare_and_exchange
// container is nil
func AtomicPointerCompareAndExchange(atomic unsafe.Pointer, oldval unsafe.Pointer, newval unsafe.Pointer) (result bool) {
	iv, err := _I.Get(814, "atomic_pointer_compare_and_exchange", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewPointerArgument(atomic)
	arg_oldval := gi.NewPointerArgument(oldval)
	arg_newval := gi.NewPointerArgument(newval)
	args := []gi.Argument{arg_atomic, arg_oldval, arg_newval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_atomic_pointer_get
// container is nil
func AtomicPointerGet(atomic unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(815, "atomic_pointer_get", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewPointerArgument(atomic)
	args := []gi.Argument{arg_atomic}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_atomic_pointer_or
// container is nil
func AtomicPointerOr(atomic unsafe.Pointer, val uint64) (result uint64) {
	iv, err := _I.Get(816, "atomic_pointer_or", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewPointerArgument(atomic)
	arg_val := gi.NewUint64Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_atomic_pointer_set
// container is nil
func AtomicPointerSet(atomic unsafe.Pointer, newval unsafe.Pointer) {
	iv, err := _I.Get(817, "atomic_pointer_set", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewPointerArgument(atomic)
	arg_newval := gi.NewPointerArgument(newval)
	args := []gi.Argument{arg_atomic, arg_newval}
	iv.Call(args, nil, nil)
}

// g_atomic_pointer_xor
// container is nil
func AtomicPointerXor(atomic unsafe.Pointer, val uint64) (result uint64) {
	iv, err := _I.Get(818, "atomic_pointer_xor", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_atomic := gi.NewPointerArgument(atomic)
	arg_val := gi.NewUint64Argument(val)
	args := []gi.Argument{arg_atomic, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_atomic_rc_box_acquire
// container is nil
func AtomicRcBoxAcquire(mem_block unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(819, "atomic_rc_box_acquire", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_mem_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_atomic_rc_box_alloc
// container is nil
func AtomicRcBoxAlloc(block_size uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(820, "atomic_rc_box_alloc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	args := []gi.Argument{arg_block_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_atomic_rc_box_alloc0
// container is nil
func AtomicRcBoxAlloc0(block_size uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(821, "atomic_rc_box_alloc0", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	args := []gi.Argument{arg_block_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_atomic_rc_box_dup
// container is nil
func AtomicRcBoxDup(block_size uint64, mem_block unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(822, "atomic_rc_box_dup", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_block_size, arg_mem_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_atomic_rc_box_get_size
// container is nil
func AtomicRcBoxGetSize(mem_block unsafe.Pointer) (result uint64) {
	iv, err := _I.Get(823, "atomic_rc_box_get_size", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_mem_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_atomic_rc_box_release
// container is nil
func AtomicRcBoxRelease(mem_block unsafe.Pointer) {
	iv, err := _I.Get(824, "atomic_rc_box_release", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_mem_block}
	iv.Call(args, nil, nil)
}

// g_atomic_rc_box_release_full
// container is nil
func AtomicRcBoxReleaseFull(mem_block unsafe.Pointer, clear_func int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(825, "atomic_rc_box_release_full", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	arg_clear_func := gi.NewIntArgument(clear_func) /*TODO*/
	args := []gi.Argument{arg_mem_block, arg_clear_func}
	iv.Call(args, nil, nil)
}

// g_atomic_ref_count_compare
// container is nil
func AtomicRefCountCompare(arc int32, val int32) (result bool) {
	iv, err := _I.Get(826, "atomic_ref_count_compare", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_arc := gi.NewInt32Argument(arc)
	arg_val := gi.NewInt32Argument(val)
	args := []gi.Argument{arg_arc, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_atomic_ref_count_dec
// container is nil
func AtomicRefCountDec(arc int32) (result bool) {
	iv, err := _I.Get(827, "atomic_ref_count_dec", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_arc := gi.NewInt32Argument(arc)
	args := []gi.Argument{arg_arc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_atomic_ref_count_inc
// container is nil
func AtomicRefCountInc(arc int32) {
	iv, err := _I.Get(828, "atomic_ref_count_inc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_arc := gi.NewInt32Argument(arc)
	args := []gi.Argument{arg_arc}
	iv.Call(args, nil, nil)
}

// g_atomic_ref_count_init
// container is nil
func AtomicRefCountInit(arc int32) {
	iv, err := _I.Get(829, "atomic_ref_count_init", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_arc := gi.NewInt32Argument(arc)
	args := []gi.Argument{arg_arc}
	iv.Call(args, nil, nil)
}

// g_base64_decode
// container is nil
// ret lenArgIdx 1
func Base64Decode(text string) (result gi.Uint8Array) {
	iv, err := _I.Get(830, "base64_decode", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_text := gi.CString(text)
	arg_text := gi.NewStringArgument(c_text)
	arg_out_len := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_text, arg_out_len}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_text)
	var out_len uint64
	_ = out_len
	out_len = outArgs[0].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(out_len)}
	return
}

// g_base64_decode_inplace
// container is nil
// arg 0 text lenArgIdx 1
func Base64DecodeInplace(text int /*TODO:TYPE*/, out_len int /*TODO:TYPE*/) (result uint8) {
	iv, err := _I.Get(831, "base64_decode_inplace", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	var ret gi.Argument
	iv.Call(nil, &ret, &outArgs[0])
	result = ret.Uint8()
	return
}

// g_base64_encode
// container is nil
// arg 0 data lenArgIdx 1
func Base64Encode(data gi.Uint8Array, len1 uint64) (result string) {
	iv, err := _I.Get(832, "base64_encode", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data.P)
	arg_len1 := gi.NewUint64Argument(len1)
	args := []gi.Argument{arg_data, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_base64_encode_close
// container is nil
func Base64EncodeClose(break_lines bool, state int /*TODO:TYPE*/, save int /*TODO:TYPE*/) (result uint64, out gi.Uint8Array) {
	iv, err := _I.Get(833, "base64_encode_close", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_break_lines := gi.NewBoolArgument(break_lines)
	arg_out := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_break_lines, arg_out}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	out.P = outArgs[0].Pointer()
	result = ret.Uint64()
	return
}

// g_base64_encode_step
// container is nil
// arg 0 in lenArgIdx 1
func Base64EncodeStep(in gi.Uint8Array, len1 uint64, break_lines bool, state int /*TODO:TYPE*/, save int /*TODO:TYPE*/) (result uint64, out gi.Uint8Array) {
	iv, err := _I.Get(834, "base64_encode_step", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [3]gi.Argument
	arg_in := gi.NewPointerArgument(in.P)
	arg_len1 := gi.NewUint64Argument(len1)
	arg_break_lines := gi.NewBoolArgument(break_lines)
	arg_out := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_in, arg_len1, arg_break_lines, arg_out}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	out.P = outArgs[0].Pointer()
	result = ret.Uint64()
	return
}

// g_basename
// container is nil
func Basename(file_name string) (result string) {
	iv, err := _I.Get(835, "basename", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_file_name := gi.CString(file_name)
	arg_file_name := gi.NewStringArgument(c_file_name)
	args := []gi.Argument{arg_file_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_file_name)
	result = ret.String().Take()
	return
}

// g_bit_lock
// container is nil
func BitLock(address int32, lock_bit int32) {
	iv, err := _I.Get(836, "bit_lock", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_address := gi.NewInt32Argument(address)
	arg_lock_bit := gi.NewInt32Argument(lock_bit)
	args := []gi.Argument{arg_address, arg_lock_bit}
	iv.Call(args, nil, nil)
}

// g_bit_nth_lsf
// container is nil
func BitNthLsf(mask uint64, nth_bit int32) (result int32) {
	iv, err := _I.Get(837, "bit_nth_lsf", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mask := gi.NewUint64Argument(mask)
	arg_nth_bit := gi.NewInt32Argument(nth_bit)
	args := []gi.Argument{arg_mask, arg_nth_bit}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_bit_nth_msf
// container is nil
func BitNthMsf(mask uint64, nth_bit int32) (result int32) {
	iv, err := _I.Get(838, "bit_nth_msf", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mask := gi.NewUint64Argument(mask)
	arg_nth_bit := gi.NewInt32Argument(nth_bit)
	args := []gi.Argument{arg_mask, arg_nth_bit}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_bit_storage
// container is nil
func BitStorage(number uint64) (result uint32) {
	iv, err := _I.Get(839, "bit_storage", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_number := gi.NewUint64Argument(number)
	args := []gi.Argument{arg_number}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_bit_trylock
// container is nil
func BitTrylock(address int32, lock_bit int32) (result bool) {
	iv, err := _I.Get(840, "bit_trylock", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_address := gi.NewInt32Argument(address)
	arg_lock_bit := gi.NewInt32Argument(lock_bit)
	args := []gi.Argument{arg_address, arg_lock_bit}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_bit_unlock
// container is nil
func BitUnlock(address int32, lock_bit int32) {
	iv, err := _I.Get(841, "bit_unlock", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_address := gi.NewInt32Argument(address)
	arg_lock_bit := gi.NewInt32Argument(lock_bit)
	args := []gi.Argument{arg_address, arg_lock_bit}
	iv.Call(args, nil, nil)
}

// g_bookmark_file_error_quark
// container is nil
func BookmarkFileErrorQuark() (result uint32) {
	iv, err := _I.Get(842, "bookmark_file_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_build_filenamev
// container is nil
func BuildFilenamev(args gi.CStrArray) (result string) {
	iv, err := _I.Get(843, "build_filenamev", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_args := gi.NewPointerArgument(args.P)
	args1 := []gi.Argument{arg_args}
	var ret gi.Argument
	iv.Call(args1, &ret, nil)
	result = ret.String().Take()
	return
}

// g_build_pathv
// container is nil
func BuildPathv(separator string, args gi.CStrArray) (result string) {
	iv, err := _I.Get(844, "build_pathv", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_separator := gi.CString(separator)
	arg_separator := gi.NewStringArgument(c_separator)
	arg_args := gi.NewPointerArgument(args.P)
	args1 := []gi.Argument{arg_separator, arg_args}
	var ret gi.Argument
	iv.Call(args1, &ret, nil)
	gi.Free(c_separator)
	result = ret.String().Take()
	return
}

// g_byte_array_free
// container is nil
func ByteArrayFree(array ByteArray, free_segment bool) (result uint8) {
	iv, err := _I.Get(845, "byte_array_free", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_array := gi.NewPointerArgument(array.P)
	arg_free_segment := gi.NewBoolArgument(free_segment)
	args := []gi.Argument{arg_array, arg_free_segment}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_byte_array_free_to_bytes
// container is nil
func ByteArrayFreeToBytes(array ByteArray) (result Bytes) {
	iv, err := _I.Get(846, "byte_array_free_to_bytes", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_array := gi.NewPointerArgument(array.P)
	args := []gi.Argument{arg_array}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_byte_array_new
// container is nil
func ByteArrayNew() (result ByteArray) {
	iv, err := _I.Get(847, "byte_array_new", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_byte_array_new_take
// container is nil
// arg 0 data lenArgIdx 1
func ByteArrayNewTake(data gi.Uint8Array, len1 uint64) (result ByteArray) {
	iv, err := _I.Get(848, "byte_array_new_take", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data.P)
	arg_len1 := gi.NewUint64Argument(len1)
	args := []gi.Argument{arg_data, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_byte_array_unref
// container is nil
func ByteArrayUnref(array ByteArray) {
	iv, err := _I.Get(849, "byte_array_unref", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_array := gi.NewPointerArgument(array.P)
	args := []gi.Argument{arg_array}
	iv.Call(args, nil, nil)
}

// g_canonicalize_filename
// container is nil
func CanonicalizeFilename(filename string, relative_to string) (result string) {
	iv, err := _I.Get(850, "canonicalize_filename", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	c_relative_to := gi.CString(relative_to)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_relative_to := gi.NewStringArgument(c_relative_to)
	args := []gi.Argument{arg_filename, arg_relative_to}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	gi.Free(c_relative_to)
	result = ret.String().Take()
	return
}

// g_chdir
// container is nil
func Chdir(path string) (result int32) {
	iv, err := _I.Get(851, "chdir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_path := gi.CString(path)
	arg_path := gi.NewStringArgument(c_path)
	args := []gi.Argument{arg_path}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_path)
	result = ret.Int32()
	return
}

// glib_check_version
// container is nil
func CheckVersion(required_major uint32, required_minor uint32, required_micro uint32) (result string) {
	iv, err := _I.Get(852, "check_version", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_required_major := gi.NewUint32Argument(required_major)
	arg_required_minor := gi.NewUint32Argument(required_minor)
	arg_required_micro := gi.NewUint32Argument(required_micro)
	args := []gi.Argument{arg_required_major, arg_required_minor, arg_required_micro}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_checksum_type_get_length
// container is nil
func ChecksumTypeGetLength(checksum_type ChecksumTypeEnum) (result int64) {
	iv, err := _I.Get(853, "checksum_type_get_length", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_checksum_type := gi.NewIntArgument(int(checksum_type))
	args := []gi.Argument{arg_checksum_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_child_watch_add_full
// container is nil
func ChildWatchAdd(priority int32, pid int32, function int /*TODO_TYPE isPtr: false, tag: interface*/, data unsafe.Pointer, notify int /*TODO_TYPE isPtr: false, tag: interface*/) (result uint32) {
	iv, err := _I.Get(854, "child_watch_add", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_priority := gi.NewInt32Argument(priority)
	arg_pid := gi.NewInt32Argument(pid)
	arg_function := gi.NewIntArgument(function) /*TODO*/
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewIntArgument(notify) /*TODO*/
	args := []gi.Argument{arg_priority, arg_pid, arg_function, arg_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_child_watch_source_new
// container is nil
func ChildWatchSourceNew(pid int32) (result Source) {
	iv, err := _I.Get(855, "child_watch_source_new", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_pid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_clear_error
// container is nil
func ClearError() (err error) {
	iv, err := _I.Get(856, "clear_error", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_err}
	iv.Call(args, nil, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	return
}

// g_close
// container is nil
func Close(fd int32) (result bool, err error) {
	iv, err := _I.Get(857, "close", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_fd := gi.NewInt32Argument(fd)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_fd, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_compute_checksum_for_bytes
// container is nil
func ComputeChecksumForBytes(checksum_type ChecksumTypeEnum, data Bytes) (result string) {
	iv, err := _I.Get(858, "compute_checksum_for_bytes", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_checksum_type := gi.NewIntArgument(int(checksum_type))
	arg_data := gi.NewPointerArgument(data.P)
	args := []gi.Argument{arg_checksum_type, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_compute_checksum_for_data
// container is nil
// arg 1 data lenArgIdx 2
func ComputeChecksumForData(checksum_type ChecksumTypeEnum, data gi.Uint8Array, length uint64) (result string) {
	iv, err := _I.Get(859, "compute_checksum_for_data", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_checksum_type := gi.NewIntArgument(int(checksum_type))
	arg_data := gi.NewPointerArgument(data.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_checksum_type, arg_data, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_compute_checksum_for_string
// container is nil
func ComputeChecksumForString(checksum_type ChecksumTypeEnum, str string, length int64) (result string) {
	iv, err := _I.Get(860, "compute_checksum_for_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_checksum_type := gi.NewIntArgument(int(checksum_type))
	arg_str := gi.NewStringArgument(c_str)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_checksum_type, arg_str, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_compute_hmac_for_bytes
// container is nil
func ComputeHmacForBytes(digest_type ChecksumTypeEnum, key Bytes, data Bytes) (result string) {
	iv, err := _I.Get(861, "compute_hmac_for_bytes", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_digest_type := gi.NewIntArgument(int(digest_type))
	arg_key := gi.NewPointerArgument(key.P)
	arg_data := gi.NewPointerArgument(data.P)
	args := []gi.Argument{arg_digest_type, arg_key, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_compute_hmac_for_data
// container is nil
// arg 1 key lenArgIdx 2
// arg 3 data lenArgIdx 4
func ComputeHmacForData(digest_type ChecksumTypeEnum, key gi.Uint8Array, key_len uint64, data gi.Uint8Array, length uint64) (result string) {
	iv, err := _I.Get(862, "compute_hmac_for_data", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_digest_type := gi.NewIntArgument(int(digest_type))
	arg_key := gi.NewPointerArgument(key.P)
	arg_key_len := gi.NewUint64Argument(key_len)
	arg_data := gi.NewPointerArgument(data.P)
	arg_length := gi.NewUint64Argument(length)
	args := []gi.Argument{arg_digest_type, arg_key, arg_key_len, arg_data, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_compute_hmac_for_string
// container is nil
// arg 1 key lenArgIdx 2
func ComputeHmacForString(digest_type ChecksumTypeEnum, key gi.Uint8Array, key_len uint64, str string, length int64) (result string) {
	iv, err := _I.Get(863, "compute_hmac_for_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_digest_type := gi.NewIntArgument(int(digest_type))
	arg_key := gi.NewPointerArgument(key.P)
	arg_key_len := gi.NewUint64Argument(key_len)
	arg_str := gi.NewStringArgument(c_str)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_digest_type, arg_key, arg_key_len, arg_str, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_convert
// container is nil
// arg 0 str lenArgIdx 1
// ret lenArgIdx 5
func Convert(str gi.Uint8Array, len1 int64, to_codeset string, from_codeset string) (result gi.Uint8Array, bytes_read uint64, err error) {
	iv, err := _I.Get(864, "convert", "")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_to_codeset := gi.CString(to_codeset)
	c_from_codeset := gi.CString(from_codeset)
	arg_str := gi.NewPointerArgument(str.P)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_to_codeset := gi.NewStringArgument(c_to_codeset)
	arg_from_codeset := gi.NewStringArgument(c_from_codeset)
	arg_bytes_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_bytes_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_to_codeset, arg_from_codeset, arg_bytes_read, arg_bytes_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_to_codeset)
	gi.Free(c_from_codeset)
	var bytes_written uint64
	_ = bytes_written
	err = gi.ToError(outArgs[2].Pointer())
	bytes_read = outArgs[0].Uint64()
	bytes_written = outArgs[1].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(bytes_written)}
	return
}

// g_convert_error_quark
// container is nil
func ConvertErrorQuark() (result uint32) {
	iv, err := _I.Get(865, "convert_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_convert_with_fallback
// container is nil
// arg 0 str lenArgIdx 1
// ret lenArgIdx 6
func ConvertWithFallback(str gi.Uint8Array, len1 int64, to_codeset string, from_codeset string, fallback string) (result gi.Uint8Array, bytes_read uint64, err error) {
	iv, err := _I.Get(866, "convert_with_fallback", "")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_to_codeset := gi.CString(to_codeset)
	c_from_codeset := gi.CString(from_codeset)
	c_fallback := gi.CString(fallback)
	arg_str := gi.NewPointerArgument(str.P)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_to_codeset := gi.NewStringArgument(c_to_codeset)
	arg_from_codeset := gi.NewStringArgument(c_from_codeset)
	arg_fallback := gi.NewStringArgument(c_fallback)
	arg_bytes_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_bytes_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_to_codeset, arg_from_codeset, arg_fallback, arg_bytes_read, arg_bytes_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_to_codeset)
	gi.Free(c_from_codeset)
	gi.Free(c_fallback)
	var bytes_written uint64
	_ = bytes_written
	err = gi.ToError(outArgs[2].Pointer())
	bytes_read = outArgs[0].Uint64()
	bytes_written = outArgs[1].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(bytes_written)}
	return
}

// g_datalist_foreach
// container is nil
func DatalistForeach(datalist Data, func1 int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(867, "datalist_foreach", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_datalist := gi.NewPointerArgument(datalist.P)
	arg_func1 := gi.NewIntArgument(func1) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_datalist, arg_func1, arg_user_data}
	iv.Call(args, nil, nil)
}

// g_datalist_get_data
// container is nil
func DatalistGetData(datalist Data, key string) (result unsafe.Pointer) {
	iv, err := _I.Get(868, "datalist_get_data", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	arg_datalist := gi.NewPointerArgument(datalist.P)
	arg_key := gi.NewStringArgument(c_key)
	args := []gi.Argument{arg_datalist, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_key)
	result = ret.Pointer()
	return
}

// g_datalist_get_flags
// container is nil
func DatalistGetFlags(datalist Data) (result uint32) {
	iv, err := _I.Get(869, "datalist_get_flags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_datalist := gi.NewPointerArgument(datalist.P)
	args := []gi.Argument{arg_datalist}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_datalist_id_get_data
// container is nil
func DatalistIdGetData(datalist Data, key_id uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(870, "datalist_id_get_data", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_datalist := gi.NewPointerArgument(datalist.P)
	arg_key_id := gi.NewUint32Argument(key_id)
	args := []gi.Argument{arg_datalist, arg_key_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_datalist_set_flags
// container is nil
func DatalistSetFlags(datalist Data, flags uint32) {
	iv, err := _I.Get(871, "datalist_set_flags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_datalist := gi.NewPointerArgument(datalist.P)
	arg_flags := gi.NewUint32Argument(flags)
	args := []gi.Argument{arg_datalist, arg_flags}
	iv.Call(args, nil, nil)
}

// g_datalist_unset_flags
// container is nil
func DatalistUnsetFlags(datalist Data, flags uint32) {
	iv, err := _I.Get(872, "datalist_unset_flags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_datalist := gi.NewPointerArgument(datalist.P)
	arg_flags := gi.NewUint32Argument(flags)
	args := []gi.Argument{arg_datalist, arg_flags}
	iv.Call(args, nil, nil)
}

// g_dataset_destroy
// container is nil
func DatasetDestroy(dataset_location unsafe.Pointer) {
	iv, err := _I.Get(873, "dataset_destroy", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dataset_location := gi.NewPointerArgument(dataset_location)
	args := []gi.Argument{arg_dataset_location}
	iv.Call(args, nil, nil)
}

// g_dataset_foreach
// container is nil
func DatasetForeach(dataset_location unsafe.Pointer, func1 int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer) {
	iv, err := _I.Get(874, "dataset_foreach", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dataset_location := gi.NewPointerArgument(dataset_location)
	arg_func1 := gi.NewIntArgument(func1) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_dataset_location, arg_func1, arg_user_data}
	iv.Call(args, nil, nil)
}

// g_dataset_id_get_data
// container is nil
func DatasetIdGetData(dataset_location unsafe.Pointer, key_id uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(875, "dataset_id_get_data", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dataset_location := gi.NewPointerArgument(dataset_location)
	arg_key_id := gi.NewUint32Argument(key_id)
	args := []gi.Argument{arg_dataset_location, arg_key_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_date_get_days_in_month
// container is nil
func DateGetDaysInMonth(month DateMonthEnum, year uint16) (result uint8) {
	iv, err := _I.Get(876, "date_get_days_in_month", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_month := gi.NewIntArgument(int(month))
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_month, arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_date_get_monday_weeks_in_year
// container is nil
func DateGetMondayWeeksInYear(year uint16) (result uint8) {
	iv, err := _I.Get(877, "date_get_monday_weeks_in_year", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_date_get_sunday_weeks_in_year
// container is nil
func DateGetSundayWeeksInYear(year uint16) (result uint8) {
	iv, err := _I.Get(878, "date_get_sunday_weeks_in_year", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint8()
	return
}

// g_date_is_leap_year
// container is nil
func DateIsLeapYear(year uint16) (result bool) {
	iv, err := _I.Get(879, "date_is_leap_year", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_strftime
// container is nil
func DateStrftime(s string, slen uint64, format string, date Date) (result uint64) {
	iv, err := _I.Get(880, "date_strftime", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_s := gi.CString(s)
	c_format := gi.CString(format)
	arg_s := gi.NewStringArgument(c_s)
	arg_slen := gi.NewUint64Argument(slen)
	arg_format := gi.NewStringArgument(c_format)
	arg_date := gi.NewPointerArgument(date.P)
	args := []gi.Argument{arg_s, arg_slen, arg_format, arg_date}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_s)
	gi.Free(c_format)
	result = ret.Uint64()
	return
}

// g_date_time_compare
// container is nil
func DateTimeCompare(dt1 unsafe.Pointer, dt2 unsafe.Pointer) (result int32) {
	iv, err := _I.Get(881, "date_time_compare", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dt1 := gi.NewPointerArgument(dt1)
	arg_dt2 := gi.NewPointerArgument(dt2)
	args := []gi.Argument{arg_dt1, arg_dt2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_date_time_equal
// container is nil
func DateTimeEqual(dt1 unsafe.Pointer, dt2 unsafe.Pointer) (result bool) {
	iv, err := _I.Get(882, "date_time_equal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dt1 := gi.NewPointerArgument(dt1)
	arg_dt2 := gi.NewPointerArgument(dt2)
	args := []gi.Argument{arg_dt1, arg_dt2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_time_hash
// container is nil
func DateTimeHash(datetime unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(883, "date_time_hash", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_datetime := gi.NewPointerArgument(datetime)
	args := []gi.Argument{arg_datetime}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_date_valid_day
// container is nil
func DateValidDay(day uint8) (result bool) {
	iv, err := _I.Get(884, "date_valid_day", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_day := gi.NewUint8Argument(day)
	args := []gi.Argument{arg_day}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_dmy
// container is nil
func DateValidDmy(day uint8, month DateMonthEnum, year uint16) (result bool) {
	iv, err := _I.Get(885, "date_valid_dmy", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_day := gi.NewUint8Argument(day)
	arg_month := gi.NewIntArgument(int(month))
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_day, arg_month, arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_julian
// container is nil
func DateValidJulian(julian_date uint32) (result bool) {
	iv, err := _I.Get(886, "date_valid_julian", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_julian_date := gi.NewUint32Argument(julian_date)
	args := []gi.Argument{arg_julian_date}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_month
// container is nil
func DateValidMonth(month DateMonthEnum) (result bool) {
	iv, err := _I.Get(887, "date_valid_month", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_month := gi.NewIntArgument(int(month))
	args := []gi.Argument{arg_month}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_weekday
// container is nil
func DateValidWeekday(weekday DateWeekdayEnum) (result bool) {
	iv, err := _I.Get(888, "date_valid_weekday", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_weekday := gi.NewIntArgument(int(weekday))
	args := []gi.Argument{arg_weekday}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_date_valid_year
// container is nil
func DateValidYear(year uint16) (result bool) {
	iv, err := _I.Get(889, "date_valid_year", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_year := gi.NewUint16Argument(year)
	args := []gi.Argument{arg_year}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_dcgettext
// container is nil
func Dcgettext(domain string, msgid string, category int32) (result string) {
	iv, err := _I.Get(890, "dcgettext", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_msgid := gi.CString(msgid)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_msgid := gi.NewStringArgument(c_msgid)
	arg_category := gi.NewInt32Argument(category)
	args := []gi.Argument{arg_domain, arg_msgid, arg_category}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_domain)
	gi.Free(c_msgid)
	result = ret.String().Take()
	return
}

// g_dgettext
// container is nil
func Dgettext(domain string, msgid string) (result string) {
	iv, err := _I.Get(891, "dgettext", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_msgid := gi.CString(msgid)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_msgid := gi.NewStringArgument(c_msgid)
	args := []gi.Argument{arg_domain, arg_msgid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_domain)
	gi.Free(c_msgid)
	result = ret.String().Take()
	return
}

// g_dir_make_tmp
// container is nil
func DirMakeTmp(tmpl string) (result string, err error) {
	iv, err := _I.Get(892, "dir_make_tmp", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_tmpl := gi.CString(tmpl)
	arg_tmpl := gi.NewStringArgument(c_tmpl)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_tmpl, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tmpl)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_direct_equal
// container is nil
func DirectEqual(v1 unsafe.Pointer, v2 unsafe.Pointer) (result bool) {
	iv, err := _I.Get(893, "direct_equal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v1 := gi.NewPointerArgument(v1)
	arg_v2 := gi.NewPointerArgument(v2)
	args := []gi.Argument{arg_v1, arg_v2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_direct_hash
// container is nil
func DirectHash(v unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(894, "direct_hash", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_dngettext
// container is nil
func Dngettext(domain string, msgid string, msgid_plural string, n uint64) (result string) {
	iv, err := _I.Get(895, "dngettext", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_msgid := gi.CString(msgid)
	c_msgid_plural := gi.CString(msgid_plural)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_msgid := gi.NewStringArgument(c_msgid)
	arg_msgid_plural := gi.NewStringArgument(c_msgid_plural)
	arg_n := gi.NewUint64Argument(n)
	args := []gi.Argument{arg_domain, arg_msgid, arg_msgid_plural, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_domain)
	gi.Free(c_msgid)
	gi.Free(c_msgid_plural)
	result = ret.String().Take()
	return
}

// g_double_equal
// container is nil
func DoubleEqual(v1 unsafe.Pointer, v2 unsafe.Pointer) (result bool) {
	iv, err := _I.Get(896, "double_equal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v1 := gi.NewPointerArgument(v1)
	arg_v2 := gi.NewPointerArgument(v2)
	args := []gi.Argument{arg_v1, arg_v2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_double_hash
// container is nil
func DoubleHash(v unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(897, "double_hash", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_dpgettext
// container is nil
func Dpgettext(domain string, msgctxtid string, msgidoffset uint64) (result string) {
	iv, err := _I.Get(898, "dpgettext", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_msgctxtid := gi.CString(msgctxtid)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_msgctxtid := gi.NewStringArgument(c_msgctxtid)
	arg_msgidoffset := gi.NewUint64Argument(msgidoffset)
	args := []gi.Argument{arg_domain, arg_msgctxtid, arg_msgidoffset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_domain)
	gi.Free(c_msgctxtid)
	result = ret.String().Take()
	return
}

// g_dpgettext2
// container is nil
func Dpgettext2(domain string, context string, msgid string) (result string) {
	iv, err := _I.Get(899, "dpgettext2", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_context := gi.CString(context)
	c_msgid := gi.CString(msgid)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_context := gi.NewStringArgument(c_context)
	arg_msgid := gi.NewStringArgument(c_msgid)
	args := []gi.Argument{arg_domain, arg_context, arg_msgid}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_domain)
	gi.Free(c_context)
	gi.Free(c_msgid)
	result = ret.String().Take()
	return
}

// g_environ_getenv
// container is nil
func EnvironGetenv(envp gi.CStrArray, variable string) (result string) {
	iv, err := _I.Get(900, "environ_getenv", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_variable := gi.CString(variable)
	arg_envp := gi.NewPointerArgument(envp.P)
	arg_variable := gi.NewStringArgument(c_variable)
	args := []gi.Argument{arg_envp, arg_variable}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_variable)
	result = ret.String().Take()
	return
}

// g_environ_setenv
// container is nil
func EnvironSetenv(envp gi.CStrArray, variable string, value string, overwrite bool) (result gi.CStrArray) {
	iv, err := _I.Get(901, "environ_setenv", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_variable := gi.CString(variable)
	c_value := gi.CString(value)
	arg_envp := gi.NewPointerArgument(envp.P)
	arg_variable := gi.NewStringArgument(c_variable)
	arg_value := gi.NewStringArgument(c_value)
	arg_overwrite := gi.NewBoolArgument(overwrite)
	args := []gi.Argument{arg_envp, arg_variable, arg_value, arg_overwrite}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_variable)
	gi.Free(c_value)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_environ_unsetenv
// container is nil
func EnvironUnsetenv(envp gi.CStrArray, variable string) (result gi.CStrArray) {
	iv, err := _I.Get(902, "environ_unsetenv", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_variable := gi.CString(variable)
	arg_envp := gi.NewPointerArgument(envp.P)
	arg_variable := gi.NewStringArgument(c_variable)
	args := []gi.Argument{arg_envp, arg_variable}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_variable)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_file_error_from_errno
// container is nil
func FileErrorFromErrno(err_no int32) (result FileErrorEnum) {
	iv, err := _I.Get(903, "file_error_from_errno", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_err_no := gi.NewInt32Argument(err_no)
	args := []gi.Argument{arg_err_no}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = FileErrorEnum(ret.Int())
	return
}

// g_file_error_quark
// container is nil
func FileErrorQuark() (result uint32) {
	iv, err := _I.Get(904, "file_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_file_get_contents
// container is nil
// arg 1 contents lenArgIdx 2
func FileGetContents(filename string) (result bool, contents gi.Uint8Array, err error) {
	iv, err := _I.Get(905, "file_get_contents", "")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_contents := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_length := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_filename, arg_contents, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	var length uint64
	_ = length
	err = gi.ToError(outArgs[2].Pointer())
	contents.P = outArgs[0].Pointer()
	length = outArgs[1].Uint64()
	result = ret.Bool()
	contents.Len = int(length)
	return
}

// g_file_open_tmp
// container is nil
func FileOpenTmp(tmpl string) (result int32, name_used string, err error) {
	iv, err := _I.Get(906, "file_open_tmp", "")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_tmpl := gi.CString(tmpl)
	arg_tmpl := gi.NewStringArgument(c_tmpl)
	arg_name_used := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_tmpl, arg_name_used, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_tmpl)
	err = gi.ToError(outArgs[1].Pointer())
	name_used = outArgs[0].String().Take()
	result = ret.Int32()
	return
}

// g_file_read_link
// container is nil
func FileReadLink(filename string) (result string, err error) {
	iv, err := _I.Get(907, "file_read_link", "")
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
	result = ret.String().Take()
	return
}

// g_file_set_contents
// container is nil
// arg 1 contents lenArgIdx 2
func FileSetContents(filename string, contents gi.Uint8Array, length int64) (result bool, err error) {
	iv, err := _I.Get(908, "file_set_contents", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_contents := gi.NewPointerArgument(contents.P)
	arg_length := gi.NewInt64Argument(length)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_filename, arg_contents, arg_length, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_file_test
// container is nil
func FileTest(filename string, test FileTestFlags) (result bool) {
	iv, err := _I.Get(909, "file_test", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_test := gi.NewIntArgument(int(test))
	args := []gi.Argument{arg_filename, arg_test}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	result = ret.Bool()
	return
}

// g_filename_display_basename
// container is nil
func FilenameDisplayBasename(filename string) (result string) {
	iv, err := _I.Get(910, "filename_display_basename", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	args := []gi.Argument{arg_filename}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	result = ret.String().Take()
	return
}

// g_filename_display_name
// container is nil
func FilenameDisplayName(filename string) (result string) {
	iv, err := _I.Get(911, "filename_display_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	args := []gi.Argument{arg_filename}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	result = ret.String().Take()
	return
}

// g_filename_from_uri
// container is nil
func FilenameFromUri(uri string) (result string, hostname string, err error) {
	iv, err := _I.Get(912, "filename_from_uri", "")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_uri := gi.CString(uri)
	arg_uri := gi.NewStringArgument(c_uri)
	arg_hostname := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_uri, arg_hostname, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_uri)
	err = gi.ToError(outArgs[1].Pointer())
	hostname = outArgs[0].String().Take()
	result = ret.String().Take()
	return
}

// g_filename_from_utf8
// container is nil
func FilenameFromUtf8(utf8string string, len1 int64) (result string, bytes_read uint64, bytes_written uint64, err error) {
	iv, err := _I.Get(913, "filename_from_utf8", "")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_utf8string := gi.CString(utf8string)
	arg_utf8string := gi.NewStringArgument(c_utf8string)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_bytes_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_bytes_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_utf8string, arg_len1, arg_bytes_read, arg_bytes_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_utf8string)
	err = gi.ToError(outArgs[2].Pointer())
	bytes_read = outArgs[0].Uint64()
	bytes_written = outArgs[1].Uint64()
	result = ret.String().Take()
	return
}

// g_filename_to_uri
// container is nil
func FilenameToUri(filename string, hostname string) (result string, err error) {
	iv, err := _I.Get(914, "filename_to_uri", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_filename := gi.CString(filename)
	c_hostname := gi.CString(hostname)
	arg_filename := gi.NewStringArgument(c_filename)
	arg_hostname := gi.NewStringArgument(c_hostname)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_filename, arg_hostname, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_filename)
	gi.Free(c_hostname)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_filename_to_utf8
// container is nil
func FilenameToUtf8(opsysstring string, len1 int64) (result string, bytes_read uint64, bytes_written uint64, err error) {
	iv, err := _I.Get(915, "filename_to_utf8", "")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_opsysstring := gi.CString(opsysstring)
	arg_opsysstring := gi.NewStringArgument(c_opsysstring)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_bytes_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_bytes_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_opsysstring, arg_len1, arg_bytes_read, arg_bytes_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_opsysstring)
	err = gi.ToError(outArgs[2].Pointer())
	bytes_read = outArgs[0].Uint64()
	bytes_written = outArgs[1].Uint64()
	result = ret.String().Take()
	return
}

// g_find_program_in_path
// container is nil
func FindProgramInPath(program string) (result string) {
	iv, err := _I.Get(916, "find_program_in_path", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_program := gi.CString(program)
	arg_program := gi.NewStringArgument(c_program)
	args := []gi.Argument{arg_program}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_program)
	result = ret.String().Take()
	return
}

// g_format_size
// container is nil
func FormatSize(size uint64) (result string) {
	iv, err := _I.Get(917, "format_size", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_size := gi.NewUint64Argument(size)
	args := []gi.Argument{arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_format_size_for_display
// container is nil
func FormatSizeForDisplay(size int64) (result string) {
	iv, err := _I.Get(918, "format_size_for_display", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_size := gi.NewInt64Argument(size)
	args := []gi.Argument{arg_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_format_size_full
// container is nil
func FormatSizeFull(size uint64, flags FormatSizeFlags) (result string) {
	iv, err := _I.Get(919, "format_size_full", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_size := gi.NewUint64Argument(size)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_size, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_free
// container is nil
func Free(mem unsafe.Pointer) {
	iv, err := _I.Get(920, "free", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem := gi.NewPointerArgument(mem)
	args := []gi.Argument{arg_mem}
	iv.Call(args, nil, nil)
}

// g_get_application_name
// container is nil
func GetApplicationName() (result string) {
	iv, err := _I.Get(921, "get_application_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_get_charset
// container is nil
func GetCharset() (result bool, charset string) {
	iv, err := _I.Get(922, "get_charset", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_charset := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_charset}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	charset = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// g_get_codeset
// container is nil
func GetCodeset() (result string) {
	iv, err := _I.Get(923, "get_codeset", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_get_current_dir
// container is nil
func GetCurrentDir() (result string) {
	iv, err := _I.Get(924, "get_current_dir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_get_current_time
// container is nil
func GetCurrentTime(result TimeVal) {
	iv, err := _I.Get(925, "get_current_time", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_result := gi.NewPointerArgument(result.P)
	args := []gi.Argument{arg_result}
	iv.Call(args, nil, nil)
}

// g_get_environ
// container is nil
func GetEnviron() (result gi.CStrArray) {
	iv, err := _I.Get(926, "get_environ", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_get_filename_charsets
// container is nil
func GetFilenameCharsets() (result bool, filename_charsets gi.CStrArray) {
	iv, err := _I.Get(927, "get_filename_charsets", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_filename_charsets := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_filename_charsets}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	filename_charsets.P = outArgs[0].Pointer()
	result = ret.Bool()
	return
}

// g_get_home_dir
// container is nil
func GetHomeDir() (result string) {
	iv, err := _I.Get(928, "get_home_dir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_get_host_name
// container is nil
func GetHostName() (result string) {
	iv, err := _I.Get(929, "get_host_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_get_language_names
// container is nil
func GetLanguageNames() (result gi.CStrArray) {
	iv, err := _I.Get(930, "get_language_names", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_get_language_names_with_category
// container is nil
func GetLanguageNamesWithCategory(category_name string) (result gi.CStrArray) {
	iv, err := _I.Get(931, "get_language_names_with_category", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_category_name := gi.CString(category_name)
	arg_category_name := gi.NewStringArgument(c_category_name)
	args := []gi.Argument{arg_category_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_category_name)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_get_locale_variants
// container is nil
func GetLocaleVariants(locale string) (result gi.CStrArray) {
	iv, err := _I.Get(932, "get_locale_variants", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_locale := gi.CString(locale)
	arg_locale := gi.NewStringArgument(c_locale)
	args := []gi.Argument{arg_locale}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_locale)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_get_monotonic_time
// container is nil
func GetMonotonicTime() (result int64) {
	iv, err := _I.Get(933, "get_monotonic_time", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int64()
	return
}

// g_get_num_processors
// container is nil
func GetNumProcessors() (result uint32) {
	iv, err := _I.Get(934, "get_num_processors", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_get_prgname
// container is nil
func GetPrgname() (result string) {
	iv, err := _I.Get(935, "get_prgname", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_get_real_name
// container is nil
func GetRealName() (result string) {
	iv, err := _I.Get(936, "get_real_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_get_real_time
// container is nil
func GetRealTime() (result int64) {
	iv, err := _I.Get(937, "get_real_time", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int64()
	return
}

// g_get_system_config_dirs
// container is nil
func GetSystemConfigDirs() (result gi.CStrArray) {
	iv, err := _I.Get(938, "get_system_config_dirs", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_get_system_data_dirs
// container is nil
func GetSystemDataDirs() (result gi.CStrArray) {
	iv, err := _I.Get(939, "get_system_data_dirs", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_get_tmp_dir
// container is nil
func GetTmpDir() (result string) {
	iv, err := _I.Get(940, "get_tmp_dir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_get_user_cache_dir
// container is nil
func GetUserCacheDir() (result string) {
	iv, err := _I.Get(941, "get_user_cache_dir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_get_user_config_dir
// container is nil
func GetUserConfigDir() (result string) {
	iv, err := _I.Get(942, "get_user_config_dir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_get_user_data_dir
// container is nil
func GetUserDataDir() (result string) {
	iv, err := _I.Get(943, "get_user_data_dir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_get_user_name
// container is nil
func GetUserName() (result string) {
	iv, err := _I.Get(944, "get_user_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_get_user_runtime_dir
// container is nil
func GetUserRuntimeDir() (result string) {
	iv, err := _I.Get(945, "get_user_runtime_dir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_get_user_special_dir
// container is nil
func GetUserSpecialDir(directory UserDirectoryEnum) (result string) {
	iv, err := _I.Get(946, "get_user_special_dir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_directory := gi.NewIntArgument(int(directory))
	args := []gi.Argument{arg_directory}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_getenv
// container is nil
func Getenv(variable string) (result string) {
	iv, err := _I.Get(947, "getenv", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_variable := gi.CString(variable)
	arg_variable := gi.NewStringArgument(c_variable)
	args := []gi.Argument{arg_variable}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_variable)
	result = ret.String().Take()
	return
}

// g_hash_table_add
// container is nil
func HashTableAdd(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(948, "hash_table_add", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_contains
// container is nil
func HashTableContains(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(949, "hash_table_contains", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_destroy
// container is nil
func HashTableDestroy(hash_table HashTable) {
	iv, err := _I.Get(950, "hash_table_destroy", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hash_table_insert
// container is nil
func HashTableInsert(hash_table HashTable, key unsafe.Pointer, value unsafe.Pointer) (result bool) {
	iv, err := _I.Get(951, "hash_table_insert", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_hash_table, arg_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_lookup
// container is nil
func HashTableLookup(hash_table HashTable, key unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(952, "hash_table_lookup", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_hash_table_lookup_extended
// container is nil
func HashTableLookupExtended(hash_table HashTable, lookup_key unsafe.Pointer) (result bool, orig_key unsafe.Pointer, value unsafe.Pointer) {
	iv, err := _I.Get(953, "hash_table_lookup_extended", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_lookup_key := gi.NewPointerArgument(lookup_key)
	arg_orig_key := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_hash_table, arg_lookup_key, arg_orig_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	orig_key = outArgs[0].Pointer()
	value = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// g_hash_table_remove
// container is nil
func HashTableRemove(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(954, "hash_table_remove", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_remove_all
// container is nil
func HashTableRemoveAll(hash_table HashTable) {
	iv, err := _I.Get(955, "hash_table_remove_all", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hash_table_replace
// container is nil
func HashTableReplace(hash_table HashTable, key unsafe.Pointer, value unsafe.Pointer) (result bool) {
	iv, err := _I.Get(956, "hash_table_replace", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	arg_value := gi.NewPointerArgument(value)
	args := []gi.Argument{arg_hash_table, arg_key, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_size
// container is nil
func HashTableSize(hash_table HashTable) (result uint32) {
	iv, err := _I.Get(957, "hash_table_size", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_hash_table_steal
// container is nil
func HashTableSteal(hash_table HashTable, key unsafe.Pointer) (result bool) {
	iv, err := _I.Get(958, "hash_table_steal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_key := gi.NewPointerArgument(key)
	args := []gi.Argument{arg_hash_table, arg_key}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hash_table_steal_all
// container is nil
func HashTableStealAll(hash_table HashTable) {
	iv, err := _I.Get(959, "hash_table_steal_all", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hash_table_steal_extended
// container is nil
func HashTableStealExtended(hash_table HashTable, lookup_key unsafe.Pointer) (result bool, stolen_key unsafe.Pointer, stolen_value unsafe.Pointer) {
	iv, err := _I.Get(960, "hash_table_steal_extended", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	arg_lookup_key := gi.NewPointerArgument(lookup_key)
	arg_stolen_key := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_stolen_value := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_hash_table, arg_lookup_key, arg_stolen_key, arg_stolen_value}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	stolen_key = outArgs[0].Pointer()
	stolen_value = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// g_hash_table_unref
// container is nil
func HashTableUnref(hash_table HashTable) {
	iv, err := _I.Get(961, "hash_table_unref", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hash_table := gi.NewPointerArgument(hash_table.P)
	args := []gi.Argument{arg_hash_table}
	iv.Call(args, nil, nil)
}

// g_hook_destroy
// container is nil
func HookDestroy(hook_list HookList, hook_id uint64) (result bool) {
	iv, err := _I.Get(962, "hook_destroy", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook_id := gi.NewUint64Argument(hook_id)
	args := []gi.Argument{arg_hook_list, arg_hook_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_hook_destroy_link
// container is nil
func HookDestroyLink(hook_list HookList, hook Hook) {
	iv, err := _I.Get(963, "hook_destroy_link", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_free
// container is nil
func HookFree(hook_list HookList, hook Hook) {
	iv, err := _I.Get(964, "hook_free", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_insert_before
// container is nil
func HookInsertBefore(hook_list HookList, sibling Hook, hook Hook) {
	iv, err := _I.Get(965, "hook_insert_before", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_sibling := gi.NewPointerArgument(sibling.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_sibling, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_prepend
// container is nil
func HookPrepend(hook_list HookList, hook Hook) {
	iv, err := _I.Get(966, "hook_prepend", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hook_unref
// container is nil
func HookUnref(hook_list HookList, hook Hook) {
	iv, err := _I.Get(967, "hook_unref", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_hook_list := gi.NewPointerArgument(hook_list.P)
	arg_hook := gi.NewPointerArgument(hook.P)
	args := []gi.Argument{arg_hook_list, arg_hook}
	iv.Call(args, nil, nil)
}

// g_hostname_is_ascii_encoded
// container is nil
func HostnameIsAsciiEncoded(hostname string) (result bool) {
	iv, err := _I.Get(968, "hostname_is_ascii_encoded", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_hostname := gi.CString(hostname)
	arg_hostname := gi.NewStringArgument(c_hostname)
	args := []gi.Argument{arg_hostname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_hostname)
	result = ret.Bool()
	return
}

// g_hostname_is_ip_address
// container is nil
func HostnameIsIpAddress(hostname string) (result bool) {
	iv, err := _I.Get(969, "hostname_is_ip_address", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_hostname := gi.CString(hostname)
	arg_hostname := gi.NewStringArgument(c_hostname)
	args := []gi.Argument{arg_hostname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_hostname)
	result = ret.Bool()
	return
}

// g_hostname_is_non_ascii
// container is nil
func HostnameIsNonAscii(hostname string) (result bool) {
	iv, err := _I.Get(970, "hostname_is_non_ascii", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_hostname := gi.CString(hostname)
	arg_hostname := gi.NewStringArgument(c_hostname)
	args := []gi.Argument{arg_hostname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_hostname)
	result = ret.Bool()
	return
}

// g_hostname_to_ascii
// container is nil
func HostnameToAscii(hostname string) (result string) {
	iv, err := _I.Get(971, "hostname_to_ascii", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_hostname := gi.CString(hostname)
	arg_hostname := gi.NewStringArgument(c_hostname)
	args := []gi.Argument{arg_hostname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_hostname)
	result = ret.String().Take()
	return
}

// g_hostname_to_unicode
// container is nil
func HostnameToUnicode(hostname string) (result string) {
	iv, err := _I.Get(972, "hostname_to_unicode", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_hostname := gi.CString(hostname)
	arg_hostname := gi.NewStringArgument(c_hostname)
	args := []gi.Argument{arg_hostname}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_hostname)
	result = ret.String().Take()
	return
}

// g_idle_add_full
// container is nil
func IdleAdd(priority int32, function int /*TODO_TYPE isPtr: false, tag: interface*/, data unsafe.Pointer, notify int /*TODO_TYPE isPtr: false, tag: interface*/) (result uint32) {
	iv, err := _I.Get(973, "idle_add", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_priority := gi.NewInt32Argument(priority)
	arg_function := gi.NewIntArgument(function) /*TODO*/
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewIntArgument(notify) /*TODO*/
	args := []gi.Argument{arg_priority, arg_function, arg_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_idle_remove_by_data
// container is nil
func IdleRemoveByData(data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(974, "idle_remove_by_data", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_idle_source_new
// container is nil
func IdleSourceNew() (result Source) {
	iv, err := _I.Get(975, "idle_source_new", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_int64_equal
// container is nil
func Int64Equal(v1 unsafe.Pointer, v2 unsafe.Pointer) (result bool) {
	iv, err := _I.Get(976, "int64_equal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v1 := gi.NewPointerArgument(v1)
	arg_v2 := gi.NewPointerArgument(v2)
	args := []gi.Argument{arg_v1, arg_v2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_int64_hash
// container is nil
func Int64Hash(v unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(977, "int64_hash", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_int_equal
// container is nil
func IntEqual(v1 unsafe.Pointer, v2 unsafe.Pointer) (result bool) {
	iv, err := _I.Get(978, "int_equal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v1 := gi.NewPointerArgument(v1)
	arg_v2 := gi.NewPointerArgument(v2)
	args := []gi.Argument{arg_v1, arg_v2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_int_hash
// container is nil
func IntHash(v unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(979, "int_hash", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_intern_static_string
// container is nil
func InternStaticString(string string) (result string) {
	iv, err := _I.Get(980, "intern_static_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_intern_string
// container is nil
func InternString(string string) (result string) {
	iv, err := _I.Get(981, "intern_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_io_add_watch_full
// container is nil
func IoAddWatch(channel IOChannel, priority int32, condition IOConditionFlags, func1 int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer, notify int /*TODO_TYPE isPtr: false, tag: interface*/) (result uint32) {
	iv, err := _I.Get(982, "io_add_watch", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_channel := gi.NewPointerArgument(channel.P)
	arg_priority := gi.NewInt32Argument(priority)
	arg_condition := gi.NewIntArgument(int(condition))
	arg_func1 := gi.NewIntArgument(func1) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewIntArgument(notify) /*TODO*/
	args := []gi.Argument{arg_channel, arg_priority, arg_condition, arg_func1, arg_user_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_io_channel_error_from_errno
// container is nil
func IoChannelErrorFromErrno(en int32) (result IOChannelErrorEnum) {
	iv, err := _I.Get(983, "io_channel_error_from_errno", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_en := gi.NewInt32Argument(en)
	args := []gi.Argument{arg_en}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = IOChannelErrorEnum(ret.Int())
	return
}

// g_io_channel_error_quark
// container is nil
func IoChannelErrorQuark() (result uint32) {
	iv, err := _I.Get(984, "io_channel_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_io_create_watch
// container is nil
func IoCreateWatch(channel IOChannel, condition IOConditionFlags) (result Source) {
	iv, err := _I.Get(985, "io_create_watch", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_channel := gi.NewPointerArgument(channel.P)
	arg_condition := gi.NewIntArgument(int(condition))
	args := []gi.Argument{arg_channel, arg_condition}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_key_file_error_quark
// container is nil
func KeyFileErrorQuark() (result uint32) {
	iv, err := _I.Get(986, "key_file_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_listenv
// container is nil
func Listenv() (result gi.CStrArray) {
	iv, err := _I.Get(987, "listenv", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_locale_from_utf8
// container is nil
// ret lenArgIdx 3
func LocaleFromUtf8(utf8string string, len1 int64) (result gi.Uint8Array, bytes_read uint64, err error) {
	iv, err := _I.Get(988, "locale_from_utf8", "")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_utf8string := gi.CString(utf8string)
	arg_utf8string := gi.NewStringArgument(c_utf8string)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_bytes_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_bytes_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_utf8string, arg_len1, arg_bytes_read, arg_bytes_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_utf8string)
	var bytes_written uint64
	_ = bytes_written
	err = gi.ToError(outArgs[2].Pointer())
	bytes_read = outArgs[0].Uint64()
	bytes_written = outArgs[1].Uint64()
	result = gi.Uint8Array{P: ret.Pointer(), Len: int(bytes_written)}
	return
}

// g_locale_to_utf8
// container is nil
// arg 0 opsysstring lenArgIdx 1
func LocaleToUtf8(opsysstring gi.Uint8Array, len1 int64) (result string, bytes_read uint64, bytes_written uint64, err error) {
	iv, err := _I.Get(989, "locale_to_utf8", "")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_opsysstring := gi.NewPointerArgument(opsysstring.P)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_bytes_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_bytes_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_opsysstring, arg_len1, arg_bytes_read, arg_bytes_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	bytes_read = outArgs[0].Uint64()
	bytes_written = outArgs[1].Uint64()
	result = ret.String().Take()
	return
}

// g_log_default_handler
// container is nil
func LogDefaultHandler(log_domain string, log_level LogLevelFlags, message string, unused_data unsafe.Pointer) {
	iv, err := _I.Get(990, "log_default_handler", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_log_domain := gi.CString(log_domain)
	c_message := gi.CString(message)
	arg_log_domain := gi.NewStringArgument(c_log_domain)
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_message := gi.NewStringArgument(c_message)
	arg_unused_data := gi.NewPointerArgument(unused_data)
	args := []gi.Argument{arg_log_domain, arg_log_level, arg_message, arg_unused_data}
	iv.Call(args, nil, nil)
	gi.Free(c_log_domain)
	gi.Free(c_message)
}

// g_log_remove_handler
// container is nil
func LogRemoveHandler(log_domain string, handler_id uint32) {
	iv, err := _I.Get(991, "log_remove_handler", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_log_domain := gi.CString(log_domain)
	arg_log_domain := gi.NewStringArgument(c_log_domain)
	arg_handler_id := gi.NewUint32Argument(handler_id)
	args := []gi.Argument{arg_log_domain, arg_handler_id}
	iv.Call(args, nil, nil)
	gi.Free(c_log_domain)
}

// g_log_set_always_fatal
// container is nil
func LogSetAlwaysFatal(fatal_mask LogLevelFlags) (result LogLevelFlags) {
	iv, err := _I.Get(992, "log_set_always_fatal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_fatal_mask := gi.NewIntArgument(int(fatal_mask))
	args := []gi.Argument{arg_fatal_mask}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = LogLevelFlags(ret.Int())
	return
}

// g_log_set_fatal_mask
// container is nil
func LogSetFatalMask(log_domain string, fatal_mask LogLevelFlags) (result LogLevelFlags) {
	iv, err := _I.Get(993, "log_set_fatal_mask", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_log_domain := gi.CString(log_domain)
	arg_log_domain := gi.NewStringArgument(c_log_domain)
	arg_fatal_mask := gi.NewIntArgument(int(fatal_mask))
	args := []gi.Argument{arg_log_domain, arg_fatal_mask}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_log_domain)
	result = LogLevelFlags(ret.Int())
	return
}

// g_log_set_handler_full
// container is nil
func LogSetHandler(log_domain string, log_levels LogLevelFlags, log_func int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer, destroy int /*TODO_TYPE isPtr: false, tag: interface*/) (result uint32) {
	iv, err := _I.Get(994, "log_set_handler", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_log_domain := gi.CString(log_domain)
	arg_log_domain := gi.NewStringArgument(c_log_domain)
	arg_log_levels := gi.NewIntArgument(int(log_levels))
	arg_log_func := gi.NewIntArgument(log_func) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_destroy := gi.NewIntArgument(destroy) /*TODO*/
	args := []gi.Argument{arg_log_domain, arg_log_levels, arg_log_func, arg_user_data, arg_destroy}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_log_domain)
	result = ret.Uint32()
	return
}

// g_log_set_writer_func
// container is nil
func LogSetWriterFunc(func1 int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer, user_data_free int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(995, "log_set_writer_func", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_func1 := gi.NewIntArgument(func1) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_user_data_free := gi.NewIntArgument(user_data_free) /*TODO*/
	args := []gi.Argument{arg_func1, arg_user_data, arg_user_data_free}
	iv.Call(args, nil, nil)
}

// g_log_structured_array
// container is nil
// arg 1 fields lenArgIdx 2
func LogStructuredArray(log_level LogLevelFlags, fields unsafe.Pointer, n_fields uint64) {
	iv, err := _I.Get(996, "log_structured_array", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_fields := gi.NewPointerArgument(fields)
	arg_n_fields := gi.NewUint64Argument(n_fields)
	args := []gi.Argument{arg_log_level, arg_fields, arg_n_fields}
	iv.Call(args, nil, nil)
}

// g_log_variant
// container is nil
func LogVariant(log_domain string, log_level LogLevelFlags, fields Variant) {
	iv, err := _I.Get(997, "log_variant", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_log_domain := gi.CString(log_domain)
	arg_log_domain := gi.NewStringArgument(c_log_domain)
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_fields := gi.NewPointerArgument(fields.P)
	args := []gi.Argument{arg_log_domain, arg_log_level, arg_fields}
	iv.Call(args, nil, nil)
	gi.Free(c_log_domain)
}

// g_log_writer_default
// container is nil
// arg 1 fields lenArgIdx 2
func LogWriterDefault(log_level LogLevelFlags, fields unsafe.Pointer, n_fields uint64, user_data unsafe.Pointer) (result LogWriterOutputEnum) {
	iv, err := _I.Get(998, "log_writer_default", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_fields := gi.NewPointerArgument(fields)
	arg_n_fields := gi.NewUint64Argument(n_fields)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_log_level, arg_fields, arg_n_fields, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = LogWriterOutputEnum(ret.Int())
	return
}

// g_log_writer_format_fields
// container is nil
// arg 1 fields lenArgIdx 2
func LogWriterFormatFields(log_level LogLevelFlags, fields unsafe.Pointer, n_fields uint64, use_color bool) (result string) {
	iv, err := _I.Get(999, "log_writer_format_fields", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_fields := gi.NewPointerArgument(fields)
	arg_n_fields := gi.NewUint64Argument(n_fields)
	arg_use_color := gi.NewBoolArgument(use_color)
	args := []gi.Argument{arg_log_level, arg_fields, arg_n_fields, arg_use_color}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_log_writer_is_journald
// container is nil
func LogWriterIsJournald(output_fd int32) (result bool) {
	iv, err := _I.Get(1000, "log_writer_is_journald", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_output_fd := gi.NewInt32Argument(output_fd)
	args := []gi.Argument{arg_output_fd}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_log_writer_journald
// container is nil
// arg 1 fields lenArgIdx 2
func LogWriterJournald(log_level LogLevelFlags, fields unsafe.Pointer, n_fields uint64, user_data unsafe.Pointer) (result LogWriterOutputEnum) {
	iv, err := _I.Get(1001, "log_writer_journald", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_fields := gi.NewPointerArgument(fields)
	arg_n_fields := gi.NewUint64Argument(n_fields)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_log_level, arg_fields, arg_n_fields, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = LogWriterOutputEnum(ret.Int())
	return
}

// g_log_writer_standard_streams
// container is nil
// arg 1 fields lenArgIdx 2
func LogWriterStandardStreams(log_level LogLevelFlags, fields unsafe.Pointer, n_fields uint64, user_data unsafe.Pointer) (result LogWriterOutputEnum) {
	iv, err := _I.Get(1002, "log_writer_standard_streams", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_fields := gi.NewPointerArgument(fields)
	arg_n_fields := gi.NewUint64Argument(n_fields)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_log_level, arg_fields, arg_n_fields, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = LogWriterOutputEnum(ret.Int())
	return
}

// g_log_writer_supports_color
// container is nil
func LogWriterSupportsColor(output_fd int32) (result bool) {
	iv, err := _I.Get(1003, "log_writer_supports_color", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_output_fd := gi.NewInt32Argument(output_fd)
	args := []gi.Argument{arg_output_fd}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_main_context_default
// container is nil
func MainContextDefault() (result MainContext) {
	iv, err := _I.Get(1004, "main_context_default", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_context_get_thread_default
// container is nil
func MainContextGetThreadDefault() (result MainContext) {
	iv, err := _I.Get(1005, "main_context_get_thread_default", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_context_ref_thread_default
// container is nil
func MainContextRefThreadDefault() (result MainContext) {
	iv, err := _I.Get(1006, "main_context_ref_thread_default", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_current_source
// container is nil
func MainCurrentSource() (result Source) {
	iv, err := _I.Get(1007, "main_current_source", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_main_depth
// container is nil
func MainDepth() (result int32) {
	iv, err := _I.Get(1008, "main_depth", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int32()
	return
}

// g_malloc
// container is nil
func Malloc(n_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1009, "malloc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_bytes := gi.NewUint64Argument(n_bytes)
	args := []gi.Argument{arg_n_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_malloc0
// container is nil
func Malloc0(n_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1010, "malloc0", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_bytes := gi.NewUint64Argument(n_bytes)
	args := []gi.Argument{arg_n_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_malloc0_n
// container is nil
func Malloc0N(n_blocks uint64, n_block_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1011, "malloc0_n", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_blocks := gi.NewUint64Argument(n_blocks)
	arg_n_block_bytes := gi.NewUint64Argument(n_block_bytes)
	args := []gi.Argument{arg_n_blocks, arg_n_block_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_malloc_n
// container is nil
func MallocN(n_blocks uint64, n_block_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1012, "malloc_n", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_blocks := gi.NewUint64Argument(n_blocks)
	arg_n_block_bytes := gi.NewUint64Argument(n_block_bytes)
	args := []gi.Argument{arg_n_blocks, arg_n_block_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_markup_error_quark
// container is nil
func MarkupErrorQuark() (result uint32) {
	iv, err := _I.Get(1013, "markup_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_markup_escape_text
// container is nil
func MarkupEscapeText(text string, length int64) (result string) {
	iv, err := _I.Get(1014, "markup_escape_text", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_text := gi.CString(text)
	arg_text := gi.NewStringArgument(c_text)
	arg_length := gi.NewInt64Argument(length)
	args := []gi.Argument{arg_text, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_text)
	result = ret.String().Take()
	return
}

// g_mem_is_system_malloc
// container is nil
func MemIsSystemMalloc() (result bool) {
	iv, err := _I.Get(1015, "mem_is_system_malloc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// g_mem_profile
// container is nil
func MemProfile() {
	iv, err := _I.Get(1016, "mem_profile", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// g_mem_set_vtable
// container is nil
func MemSetVtable(vtable MemVTable) {
	iv, err := _I.Get(1017, "mem_set_vtable", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_vtable := gi.NewPointerArgument(vtable.P)
	args := []gi.Argument{arg_vtable}
	iv.Call(args, nil, nil)
}

// g_memdup
// container is nil
func Memdup(mem unsafe.Pointer, byte_size uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(1018, "memdup", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem := gi.NewPointerArgument(mem)
	arg_byte_size := gi.NewUint32Argument(byte_size)
	args := []gi.Argument{arg_mem, arg_byte_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_mkdir_with_parents
// container is nil
func MkdirWithParents(pathname string, mode int32) (result int32) {
	iv, err := _I.Get(1019, "mkdir_with_parents", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_pathname := gi.CString(pathname)
	arg_pathname := gi.NewStringArgument(c_pathname)
	arg_mode := gi.NewInt32Argument(mode)
	args := []gi.Argument{arg_pathname, arg_mode}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_pathname)
	result = ret.Int32()
	return
}

// g_nullify_pointer
// container is nil
func NullifyPointer(nullify_location unsafe.Pointer) {
	iv, err := _I.Get(1020, "nullify_pointer", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_nullify_location := gi.NewPointerArgument(nullify_location)
	args := []gi.Argument{arg_nullify_location}
	iv.Call(args, nil, nil)
}

// g_number_parser_error_quark
// container is nil
func NumberParserErrorQuark() (result uint32) {
	iv, err := _I.Get(1021, "number_parser_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_on_error_query
// container is nil
func OnErrorQuery(prg_name string) {
	iv, err := _I.Get(1022, "on_error_query", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_prg_name := gi.CString(prg_name)
	arg_prg_name := gi.NewStringArgument(c_prg_name)
	args := []gi.Argument{arg_prg_name}
	iv.Call(args, nil, nil)
	gi.Free(c_prg_name)
}

// g_on_error_stack_trace
// container is nil
func OnErrorStackTrace(prg_name string) {
	iv, err := _I.Get(1023, "on_error_stack_trace", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_prg_name := gi.CString(prg_name)
	arg_prg_name := gi.NewStringArgument(c_prg_name)
	args := []gi.Argument{arg_prg_name}
	iv.Call(args, nil, nil)
	gi.Free(c_prg_name)
}

// g_once_init_enter
// container is nil
func OnceInitEnter(location unsafe.Pointer) (result bool) {
	iv, err := _I.Get(1024, "once_init_enter", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_location := gi.NewPointerArgument(location)
	args := []gi.Argument{arg_location}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_once_init_leave
// container is nil
func OnceInitLeave(location unsafe.Pointer, result uint64) {
	iv, err := _I.Get(1025, "once_init_leave", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_location := gi.NewPointerArgument(location)
	arg_result := gi.NewUint64Argument(result)
	args := []gi.Argument{arg_location, arg_result}
	iv.Call(args, nil, nil)
}

// g_option_error_quark
// container is nil
func OptionErrorQuark() (result uint32) {
	iv, err := _I.Get(1026, "option_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_parse_debug_string
// container is nil
// arg 1 keys lenArgIdx 2
func ParseDebugString(string string, keys unsafe.Pointer, nkeys uint32) (result uint32) {
	iv, err := _I.Get(1027, "parse_debug_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	arg_keys := gi.NewPointerArgument(keys)
	arg_nkeys := gi.NewUint32Argument(nkeys)
	args := []gi.Argument{arg_string, arg_keys, arg_nkeys}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Uint32()
	return
}

// g_path_get_basename
// container is nil
func PathGetBasename(file_name string) (result string) {
	iv, err := _I.Get(1028, "path_get_basename", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_file_name := gi.CString(file_name)
	arg_file_name := gi.NewStringArgument(c_file_name)
	args := []gi.Argument{arg_file_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_file_name)
	result = ret.String().Take()
	return
}

// g_path_get_dirname
// container is nil
func PathGetDirname(file_name string) (result string) {
	iv, err := _I.Get(1029, "path_get_dirname", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_file_name := gi.CString(file_name)
	arg_file_name := gi.NewStringArgument(c_file_name)
	args := []gi.Argument{arg_file_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_file_name)
	result = ret.String().Take()
	return
}

// g_path_is_absolute
// container is nil
func PathIsAbsolute(file_name string) (result bool) {
	iv, err := _I.Get(1030, "path_is_absolute", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_file_name := gi.CString(file_name)
	arg_file_name := gi.NewStringArgument(c_file_name)
	args := []gi.Argument{arg_file_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_file_name)
	result = ret.Bool()
	return
}

// g_path_skip_root
// container is nil
func PathSkipRoot(file_name string) (result string) {
	iv, err := _I.Get(1031, "path_skip_root", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_file_name := gi.CString(file_name)
	arg_file_name := gi.NewStringArgument(c_file_name)
	args := []gi.Argument{arg_file_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_file_name)
	result = ret.String().Take()
	return
}

// g_pattern_match
// container is nil
func PatternMatch(pspec PatternSpec, string_length uint32, string string, string_reversed string) (result bool) {
	iv, err := _I.Get(1032, "pattern_match", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	c_string_reversed := gi.CString(string_reversed)
	arg_pspec := gi.NewPointerArgument(pspec.P)
	arg_string_length := gi.NewUint32Argument(string_length)
	arg_string := gi.NewStringArgument(c_string)
	arg_string_reversed := gi.NewStringArgument(c_string_reversed)
	args := []gi.Argument{arg_pspec, arg_string_length, arg_string, arg_string_reversed}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	gi.Free(c_string_reversed)
	result = ret.Bool()
	return
}

// g_pattern_match_simple
// container is nil
func PatternMatchSimple(pattern string, string string) (result bool) {
	iv, err := _I.Get(1033, "pattern_match_simple", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_pattern := gi.CString(pattern)
	c_string := gi.CString(string)
	arg_pattern := gi.NewStringArgument(c_pattern)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_pattern, arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_pattern)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_pattern_match_string
// container is nil
func PatternMatchString(pspec PatternSpec, string string) (result bool) {
	iv, err := _I.Get(1034, "pattern_match_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_pspec := gi.NewPointerArgument(pspec.P)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_pspec, arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_pointer_bit_lock
// container is nil
func PointerBitLock(address unsafe.Pointer, lock_bit int32) {
	iv, err := _I.Get(1035, "pointer_bit_lock", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_address := gi.NewPointerArgument(address)
	arg_lock_bit := gi.NewInt32Argument(lock_bit)
	args := []gi.Argument{arg_address, arg_lock_bit}
	iv.Call(args, nil, nil)
}

// g_pointer_bit_trylock
// container is nil
func PointerBitTrylock(address unsafe.Pointer, lock_bit int32) (result bool) {
	iv, err := _I.Get(1036, "pointer_bit_trylock", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_address := gi.NewPointerArgument(address)
	arg_lock_bit := gi.NewInt32Argument(lock_bit)
	args := []gi.Argument{arg_address, arg_lock_bit}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_pointer_bit_unlock
// container is nil
func PointerBitUnlock(address unsafe.Pointer, lock_bit int32) {
	iv, err := _I.Get(1037, "pointer_bit_unlock", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_address := gi.NewPointerArgument(address)
	arg_lock_bit := gi.NewInt32Argument(lock_bit)
	args := []gi.Argument{arg_address, arg_lock_bit}
	iv.Call(args, nil, nil)
}

// g_poll
// container is nil
func Poll(fds PollFD, nfds uint32, timeout int32) (result int32) {
	iv, err := _I.Get(1038, "poll", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_fds := gi.NewPointerArgument(fds.P)
	arg_nfds := gi.NewUint32Argument(nfds)
	arg_timeout := gi.NewInt32Argument(timeout)
	args := []gi.Argument{arg_fds, arg_nfds, arg_timeout}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_propagate_error
// container is nil
func PropagateError(src Error) (dest Error) {
	iv, err := _I.Get(1039, "propagate_error", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_dest := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_src := gi.NewPointerArgument(src.P)
	args := []gi.Argument{arg_dest, arg_src}
	iv.Call(args, nil, &outArgs[0])
	dest.P = outArgs[0].Pointer()
	return
}

// g_quark_from_static_string
// container is nil
func QuarkFromStaticString(string string) (result uint32) {
	iv, err := _I.Get(1040, "quark_from_static_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Uint32()
	return
}

// g_quark_from_string
// container is nil
func QuarkFromString(string string) (result uint32) {
	iv, err := _I.Get(1041, "quark_from_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Uint32()
	return
}

// g_quark_to_string
// container is nil
func QuarkToString(quark uint32) (result string) {
	iv, err := _I.Get(1042, "quark_to_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_quark := gi.NewUint32Argument(quark)
	args := []gi.Argument{arg_quark}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_quark_try_string
// container is nil
func QuarkTryString(string string) (result uint32) {
	iv, err := _I.Get(1043, "quark_try_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Uint32()
	return
}

// g_random_double
// container is nil
func RandomDouble() (result float64) {
	iv, err := _I.Get(1044, "random_double", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Double()
	return
}

// g_random_double_range
// container is nil
func RandomDoubleRange(begin float64, end float64) (result float64) {
	iv, err := _I.Get(1045, "random_double_range", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_begin := gi.NewDoubleArgument(begin)
	arg_end := gi.NewDoubleArgument(end)
	args := []gi.Argument{arg_begin, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// g_random_int
// container is nil
func RandomInt() (result uint32) {
	iv, err := _I.Get(1046, "random_int", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_random_int_range
// container is nil
func RandomIntRange(begin int32, end int32) (result int32) {
	iv, err := _I.Get(1047, "random_int_range", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_begin := gi.NewInt32Argument(begin)
	arg_end := gi.NewInt32Argument(end)
	args := []gi.Argument{arg_begin, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_random_set_seed
// container is nil
func RandomSetSeed(seed uint32) {
	iv, err := _I.Get(1048, "random_set_seed", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_seed := gi.NewUint32Argument(seed)
	args := []gi.Argument{arg_seed}
	iv.Call(args, nil, nil)
}

// g_rc_box_acquire
// container is nil
func RcBoxAcquire(mem_block unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(1049, "rc_box_acquire", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_mem_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_rc_box_alloc
// container is nil
func RcBoxAlloc(block_size uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1050, "rc_box_alloc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	args := []gi.Argument{arg_block_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_rc_box_alloc0
// container is nil
func RcBoxAlloc0(block_size uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1051, "rc_box_alloc0", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	args := []gi.Argument{arg_block_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_rc_box_dup
// container is nil
func RcBoxDup(block_size uint64, mem_block unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(1052, "rc_box_dup", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_block_size, arg_mem_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_rc_box_get_size
// container is nil
func RcBoxGetSize(mem_block unsafe.Pointer) (result uint64) {
	iv, err := _I.Get(1053, "rc_box_get_size", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_mem_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_rc_box_release
// container is nil
func RcBoxRelease(mem_block unsafe.Pointer) {
	iv, err := _I.Get(1054, "rc_box_release", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_mem_block}
	iv.Call(args, nil, nil)
}

// g_rc_box_release_full
// container is nil
func RcBoxReleaseFull(mem_block unsafe.Pointer, clear_func int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(1055, "rc_box_release_full", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem_block := gi.NewPointerArgument(mem_block)
	arg_clear_func := gi.NewIntArgument(clear_func) /*TODO*/
	args := []gi.Argument{arg_mem_block, arg_clear_func}
	iv.Call(args, nil, nil)
}

// g_realloc
// container is nil
func Realloc(mem unsafe.Pointer, n_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1056, "realloc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem := gi.NewPointerArgument(mem)
	arg_n_bytes := gi.NewUint64Argument(n_bytes)
	args := []gi.Argument{arg_mem, arg_n_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_realloc_n
// container is nil
func ReallocN(mem unsafe.Pointer, n_blocks uint64, n_block_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1057, "realloc_n", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem := gi.NewPointerArgument(mem)
	arg_n_blocks := gi.NewUint64Argument(n_blocks)
	arg_n_block_bytes := gi.NewUint64Argument(n_block_bytes)
	args := []gi.Argument{arg_mem, arg_n_blocks, arg_n_block_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_ref_count_compare
// container is nil
func RefCountCompare(rc int32, val int32) (result bool) {
	iv, err := _I.Get(1058, "ref_count_compare", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_rc := gi.NewInt32Argument(rc)
	arg_val := gi.NewInt32Argument(val)
	args := []gi.Argument{arg_rc, arg_val}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_ref_count_dec
// container is nil
func RefCountDec(rc int32) (result bool) {
	iv, err := _I.Get(1059, "ref_count_dec", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_rc := gi.NewInt32Argument(rc)
	args := []gi.Argument{arg_rc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_ref_count_inc
// container is nil
func RefCountInc(rc int32) {
	iv, err := _I.Get(1060, "ref_count_inc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_rc := gi.NewInt32Argument(rc)
	args := []gi.Argument{arg_rc}
	iv.Call(args, nil, nil)
}

// g_ref_count_init
// container is nil
func RefCountInit(rc int32) {
	iv, err := _I.Get(1061, "ref_count_init", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_rc := gi.NewInt32Argument(rc)
	args := []gi.Argument{arg_rc}
	iv.Call(args, nil, nil)
}

// g_ref_string_acquire
// container is nil
func RefStringAcquire(str string) (result string) {
	iv, err := _I.Get(1062, "ref_string_acquire", "")
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

// g_ref_string_length
// container is nil
func RefStringLength(str string) (result uint64) {
	iv, err := _I.Get(1063, "ref_string_length", "")
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
	result = ret.Uint64()
	return
}

// g_ref_string_new
// container is nil
func RefStringNew(str string) (result string) {
	iv, err := _I.Get(1064, "ref_string_new", "")
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

// g_ref_string_new_intern
// container is nil
func RefStringNewIntern(str string) (result string) {
	iv, err := _I.Get(1065, "ref_string_new_intern", "")
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

// g_ref_string_new_len
// container is nil
func RefStringNewLen(str string, len1 int64) (result string) {
	iv, err := _I.Get(1066, "ref_string_new_len", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_ref_string_release
// container is nil
func RefStringRelease(str string) {
	iv, err := _I.Get(1067, "ref_string_release", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_str}
	iv.Call(args, nil, nil)
	gi.Free(c_str)
}

// g_regex_check_replacement
// container is nil
func RegexCheckReplacement(replacement string) (result bool, has_references bool, err error) {
	iv, err := _I.Get(1068, "regex_check_replacement", "")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_replacement := gi.CString(replacement)
	arg_replacement := gi.NewStringArgument(c_replacement)
	arg_has_references := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_replacement, arg_has_references, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_replacement)
	err = gi.ToError(outArgs[1].Pointer())
	has_references = outArgs[0].Bool()
	result = ret.Bool()
	return
}

// g_regex_error_quark
// container is nil
func RegexErrorQuark() (result uint32) {
	iv, err := _I.Get(1069, "regex_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_regex_escape_nul
// container is nil
func RegexEscapeNul(string string, length int32) (result string) {
	iv, err := _I.Get(1070, "regex_escape_nul", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_string, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_regex_escape_string
// container is nil
// arg 0 string lenArgIdx 1
func RegexEscapeString(string gi.CStrArray, length int32) (result string) {
	iv, err := _I.Get(1071, "regex_escape_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_string := gi.NewPointerArgument(string.P)
	arg_length := gi.NewInt32Argument(length)
	args := []gi.Argument{arg_string, arg_length}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_regex_match_simple
// container is nil
func RegexMatchSimple(pattern string, string string, compile_options RegexCompileFlags, match_options RegexMatchFlags) (result bool) {
	iv, err := _I.Get(1072, "regex_match_simple", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_pattern := gi.CString(pattern)
	c_string := gi.CString(string)
	arg_pattern := gi.NewStringArgument(c_pattern)
	arg_string := gi.NewStringArgument(c_string)
	arg_compile_options := gi.NewIntArgument(int(compile_options))
	arg_match_options := gi.NewIntArgument(int(match_options))
	args := []gi.Argument{arg_pattern, arg_string, arg_compile_options, arg_match_options}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_pattern)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_regex_split_simple
// container is nil
func RegexSplitSimple(pattern string, string string, compile_options RegexCompileFlags, match_options RegexMatchFlags) (result gi.CStrArray) {
	iv, err := _I.Get(1073, "regex_split_simple", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_pattern := gi.CString(pattern)
	c_string := gi.CString(string)
	arg_pattern := gi.NewStringArgument(c_pattern)
	arg_string := gi.NewStringArgument(c_string)
	arg_compile_options := gi.NewIntArgument(int(compile_options))
	arg_match_options := gi.NewIntArgument(int(match_options))
	args := []gi.Argument{arg_pattern, arg_string, arg_compile_options, arg_match_options}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_pattern)
	gi.Free(c_string)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_reload_user_special_dirs_cache
// container is nil
func ReloadUserSpecialDirsCache() {
	iv, err := _I.Get(1074, "reload_user_special_dirs_cache", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// g_rmdir
// container is nil
func Rmdir(filename string) (result int32) {
	iv, err := _I.Get(1075, "rmdir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	args := []gi.Argument{arg_filename}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	result = ret.Int32()
	return
}

// g_sequence_get
// container is nil
func SequenceGet(iter SequenceIter) (result unsafe.Pointer) {
	iv, err := _I.Get(1076, "sequence_get", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	args := []gi.Argument{arg_iter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_sequence_insert_before
// container is nil
func SequenceInsertBefore(iter SequenceIter, data unsafe.Pointer) (result SequenceIter) {
	iv, err := _I.Get(1077, "sequence_insert_before", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_iter, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_move
// container is nil
func SequenceMove(src SequenceIter, dest SequenceIter) {
	iv, err := _I.Get(1078, "sequence_move", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_src := gi.NewPointerArgument(src.P)
	arg_dest := gi.NewPointerArgument(dest.P)
	args := []gi.Argument{arg_src, arg_dest}
	iv.Call(args, nil, nil)
}

// g_sequence_move_range
// container is nil
func SequenceMoveRange(dest SequenceIter, begin SequenceIter, end SequenceIter) {
	iv, err := _I.Get(1079, "sequence_move_range", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dest := gi.NewPointerArgument(dest.P)
	arg_begin := gi.NewPointerArgument(begin.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_dest, arg_begin, arg_end}
	iv.Call(args, nil, nil)
}

// g_sequence_range_get_midpoint
// container is nil
func SequenceRangeGetMidpoint(begin SequenceIter, end SequenceIter) (result SequenceIter) {
	iv, err := _I.Get(1080, "sequence_range_get_midpoint", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_begin := gi.NewPointerArgument(begin.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_begin, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_sequence_remove
// container is nil
func SequenceRemove(iter SequenceIter) {
	iv, err := _I.Get(1081, "sequence_remove", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	args := []gi.Argument{arg_iter}
	iv.Call(args, nil, nil)
}

// g_sequence_remove_range
// container is nil
func SequenceRemoveRange(begin SequenceIter, end SequenceIter) {
	iv, err := _I.Get(1082, "sequence_remove_range", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_begin := gi.NewPointerArgument(begin.P)
	arg_end := gi.NewPointerArgument(end.P)
	args := []gi.Argument{arg_begin, arg_end}
	iv.Call(args, nil, nil)
}

// g_sequence_set
// container is nil
func SequenceSet(iter SequenceIter, data unsafe.Pointer) {
	iv, err := _I.Get(1083, "sequence_set", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iter := gi.NewPointerArgument(iter.P)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_iter, arg_data}
	iv.Call(args, nil, nil)
}

// g_sequence_swap
// container is nil
func SequenceSwap(a SequenceIter, b SequenceIter) {
	iv, err := _I.Get(1084, "sequence_swap", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_a := gi.NewPointerArgument(a.P)
	arg_b := gi.NewPointerArgument(b.P)
	args := []gi.Argument{arg_a, arg_b}
	iv.Call(args, nil, nil)
}

// g_set_application_name
// container is nil
func SetApplicationName(application_name string) {
	iv, err := _I.Get(1085, "set_application_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_application_name := gi.CString(application_name)
	arg_application_name := gi.NewStringArgument(c_application_name)
	args := []gi.Argument{arg_application_name}
	iv.Call(args, nil, nil)
	gi.Free(c_application_name)
}

// g_set_error_literal
// container is nil
func SetErrorLiteral(domain uint32, code int32, message string) (err1 Error) {
	iv, err := _I.Get(1086, "set_error_literal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_message := gi.CString(message)
	arg_err1 := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_domain := gi.NewUint32Argument(domain)
	arg_code := gi.NewInt32Argument(code)
	arg_message := gi.NewStringArgument(c_message)
	args := []gi.Argument{arg_err1, arg_domain, arg_code, arg_message}
	iv.Call(args, nil, &outArgs[0])
	gi.Free(c_message)
	err1.P = outArgs[0].Pointer()
	return
}

// g_set_prgname
// container is nil
func SetPrgname(prgname string) {
	iv, err := _I.Get(1087, "set_prgname", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_prgname := gi.CString(prgname)
	arg_prgname := gi.NewStringArgument(c_prgname)
	args := []gi.Argument{arg_prgname}
	iv.Call(args, nil, nil)
	gi.Free(c_prgname)
}

// g_setenv
// container is nil
func Setenv(variable string, value string, overwrite bool) (result bool) {
	iv, err := _I.Get(1088, "setenv", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_variable := gi.CString(variable)
	c_value := gi.CString(value)
	arg_variable := gi.NewStringArgument(c_variable)
	arg_value := gi.NewStringArgument(c_value)
	arg_overwrite := gi.NewBoolArgument(overwrite)
	args := []gi.Argument{arg_variable, arg_value, arg_overwrite}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_variable)
	gi.Free(c_value)
	result = ret.Bool()
	return
}

// g_shell_error_quark
// container is nil
func ShellErrorQuark() (result uint32) {
	iv, err := _I.Get(1089, "shell_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_shell_parse_argv
// container is nil
// arg 2 argvp lenArgIdx 1
func ShellParseArgv(command_line string) (result bool, argvp gi.CStrArray, err error) {
	iv, err := _I.Get(1090, "shell_parse_argv", "")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_command_line := gi.CString(command_line)
	arg_command_line := gi.NewStringArgument(c_command_line)
	arg_argcp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_argvp := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_command_line, arg_argcp, arg_argvp, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_command_line)
	var argcp int32
	_ = argcp
	err = gi.ToError(outArgs[2].Pointer())
	argcp = outArgs[0].Int32()
	argvp.P = outArgs[1].Pointer()
	result = ret.Bool()
	return
}

// g_shell_quote
// container is nil
func ShellQuote(unquoted_string string) (result string) {
	iv, err := _I.Get(1091, "shell_quote", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_unquoted_string := gi.CString(unquoted_string)
	arg_unquoted_string := gi.NewStringArgument(c_unquoted_string)
	args := []gi.Argument{arg_unquoted_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_unquoted_string)
	result = ret.String().Take()
	return
}

// g_shell_unquote
// container is nil
func ShellUnquote(quoted_string string) (result string, err error) {
	iv, err := _I.Get(1092, "shell_unquote", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_quoted_string := gi.CString(quoted_string)
	arg_quoted_string := gi.NewStringArgument(c_quoted_string)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_quoted_string, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_quoted_string)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.String().Take()
	return
}

// g_slice_alloc
// container is nil
func SliceAlloc(block_size uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1093, "slice_alloc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	args := []gi.Argument{arg_block_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_slice_alloc0
// container is nil
func SliceAlloc0(block_size uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1094, "slice_alloc0", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	args := []gi.Argument{arg_block_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_slice_copy
// container is nil
func SliceCopy(block_size uint64, mem_block unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(1095, "slice_copy", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_block_size, arg_mem_block}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_slice_free1
// container is nil
func SliceFree1(block_size uint64, mem_block unsafe.Pointer) {
	iv, err := _I.Get(1096, "slice_free1", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	arg_mem_block := gi.NewPointerArgument(mem_block)
	args := []gi.Argument{arg_block_size, arg_mem_block}
	iv.Call(args, nil, nil)
}

// g_slice_free_chain_with_offset
// container is nil
func SliceFreeChainWithOffset(block_size uint64, mem_chain unsafe.Pointer, next_offset uint64) {
	iv, err := _I.Get(1097, "slice_free_chain_with_offset", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_block_size := gi.NewUint64Argument(block_size)
	arg_mem_chain := gi.NewPointerArgument(mem_chain)
	arg_next_offset := gi.NewUint64Argument(next_offset)
	args := []gi.Argument{arg_block_size, arg_mem_chain, arg_next_offset}
	iv.Call(args, nil, nil)
}

// g_slice_get_config
// container is nil
func SliceGetConfig(ckey SliceConfigEnum) (result int64) {
	iv, err := _I.Get(1098, "slice_get_config", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ckey := gi.NewIntArgument(int(ckey))
	args := []gi.Argument{arg_ckey}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_slice_get_config_state
// container is nil
func SliceGetConfigState(ckey SliceConfigEnum, address int64, n_values uint32) (result int64) {
	iv, err := _I.Get(1099, "slice_get_config_state", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ckey := gi.NewIntArgument(int(ckey))
	arg_address := gi.NewInt64Argument(address)
	arg_n_values := gi.NewUint32Argument(n_values)
	args := []gi.Argument{arg_ckey, arg_address, arg_n_values}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int64()
	return
}

// g_slice_set_config
// container is nil
func SliceSetConfig(ckey SliceConfigEnum, value int64) {
	iv, err := _I.Get(1100, "slice_set_config", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ckey := gi.NewIntArgument(int(ckey))
	arg_value := gi.NewInt64Argument(value)
	args := []gi.Argument{arg_ckey, arg_value}
	iv.Call(args, nil, nil)
}

// g_source_remove
// container is nil
func SourceRemove(tag uint32) (result bool) {
	iv, err := _I.Get(1101, "source_remove", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_tag := gi.NewUint32Argument(tag)
	args := []gi.Argument{arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_source_remove_by_funcs_user_data
// container is nil
func SourceRemoveByFuncsUserData(funcs SourceFuncs, user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(1102, "source_remove_by_funcs_user_data", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_funcs := gi.NewPointerArgument(funcs.P)
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_funcs, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_source_remove_by_user_data
// container is nil
func SourceRemoveByUserData(user_data unsafe.Pointer) (result bool) {
	iv, err := _I.Get(1103, "source_remove_by_user_data", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_source_set_name_by_id
// container is nil
func SourceSetNameById(tag uint32, name string) {
	iv, err := _I.Get(1104, "source_set_name_by_id", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_tag := gi.NewUint32Argument(tag)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_tag, arg_name}
	iv.Call(args, nil, nil)
	gi.Free(c_name)
}

// g_spaced_primes_closest
// container is nil
func SpacedPrimesClosest(num uint32) (result uint32) {
	iv, err := _I.Get(1105, "spaced_primes_closest", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_num := gi.NewUint32Argument(num)
	args := []gi.Argument{arg_num}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_spawn_async
// container is nil
func SpawnAsync(working_directory string, argv gi.CStrArray, envp gi.CStrArray, flags SpawnFlags, child_setup int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer) (result bool, child_pid int32, err error) {
	iv, err := _I.Get(1106, "spawn_async", "")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_working_directory := gi.CString(working_directory)
	arg_working_directory := gi.NewStringArgument(c_working_directory)
	arg_argv := gi.NewPointerArgument(argv.P)
	arg_envp := gi.NewPointerArgument(envp.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_child_setup := gi.NewIntArgument(child_setup) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_child_pid := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_working_directory, arg_argv, arg_envp, arg_flags, arg_child_setup, arg_user_data, arg_child_pid, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_working_directory)
	err = gi.ToError(outArgs[1].Pointer())
	child_pid = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// g_spawn_async_with_fds
// container is nil
func SpawnAsyncWithFds(working_directory string, argv gi.CStrArray, envp gi.CStrArray, flags SpawnFlags, child_setup int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer, stdin_fd int32, stdout_fd int32, stderr_fd int32) (result bool, child_pid int32, err error) {
	iv, err := _I.Get(1107, "spawn_async_with_fds", "")
	if err != nil {
		return
	}
	var outArgs [2]gi.Argument
	c_working_directory := gi.CString(working_directory)
	arg_working_directory := gi.NewStringArgument(c_working_directory)
	arg_argv := gi.NewPointerArgument(argv.P)
	arg_envp := gi.NewPointerArgument(envp.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_child_setup := gi.NewIntArgument(child_setup) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_child_pid := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_stdin_fd := gi.NewInt32Argument(stdin_fd)
	arg_stdout_fd := gi.NewInt32Argument(stdout_fd)
	arg_stderr_fd := gi.NewInt32Argument(stderr_fd)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	args := []gi.Argument{arg_working_directory, arg_argv, arg_envp, arg_flags, arg_child_setup, arg_user_data, arg_child_pid, arg_stdin_fd, arg_stdout_fd, arg_stderr_fd, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_working_directory)
	err = gi.ToError(outArgs[1].Pointer())
	child_pid = outArgs[0].Int32()
	result = ret.Bool()
	return
}

// g_spawn_async_with_pipes
// container is nil
func SpawnAsyncWithPipes(working_directory string, argv gi.CStrArray, envp gi.CStrArray, flags SpawnFlags, child_setup int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer) (result bool, child_pid int32, standard_input int32, standard_output int32, standard_error int32, err error) {
	iv, err := _I.Get(1108, "spawn_async_with_pipes", "")
	if err != nil {
		return
	}
	var outArgs [5]gi.Argument
	c_working_directory := gi.CString(working_directory)
	arg_working_directory := gi.NewStringArgument(c_working_directory)
	arg_argv := gi.NewPointerArgument(argv.P)
	arg_envp := gi.NewPointerArgument(envp.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_child_setup := gi.NewIntArgument(child_setup) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_child_pid := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_standard_input := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_standard_output := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_standard_error := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[4]))
	args := []gi.Argument{arg_working_directory, arg_argv, arg_envp, arg_flags, arg_child_setup, arg_user_data, arg_child_pid, arg_standard_input, arg_standard_output, arg_standard_error, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_working_directory)
	err = gi.ToError(outArgs[4].Pointer())
	child_pid = outArgs[0].Int32()
	standard_input = outArgs[1].Int32()
	standard_output = outArgs[2].Int32()
	standard_error = outArgs[3].Int32()
	result = ret.Bool()
	return
}

// g_spawn_check_exit_status
// container is nil
func SpawnCheckExitStatus(exit_status int32) (result bool, err error) {
	iv, err := _I.Get(1109, "spawn_check_exit_status", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_exit_status := gi.NewInt32Argument(exit_status)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_exit_status, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_spawn_close_pid
// container is nil
func SpawnClosePid(pid int32) {
	iv, err := _I.Get(1110, "spawn_close_pid", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_pid := gi.NewInt32Argument(pid)
	args := []gi.Argument{arg_pid}
	iv.Call(args, nil, nil)
}

// g_spawn_command_line_async
// container is nil
func SpawnCommandLineAsync(command_line string) (result bool, err error) {
	iv, err := _I.Get(1111, "spawn_command_line_async", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_command_line := gi.CString(command_line)
	arg_command_line := gi.NewStringArgument(c_command_line)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_command_line, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_command_line)
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_spawn_command_line_sync
// container is nil
func SpawnCommandLineSync(command_line string) (result bool, standard_output gi.Uint8Array, standard_error gi.Uint8Array, exit_status int32, err error) {
	iv, err := _I.Get(1112, "spawn_command_line_sync", "")
	if err != nil {
		return
	}
	var outArgs [4]gi.Argument
	c_command_line := gi.CString(command_line)
	arg_command_line := gi.NewStringArgument(c_command_line)
	arg_standard_output := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_standard_error := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_exit_status := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_command_line, arg_standard_output, arg_standard_error, arg_exit_status, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_command_line)
	err = gi.ToError(outArgs[3].Pointer())
	standard_output.P = outArgs[0].Pointer()
	standard_error.P = outArgs[1].Pointer()
	exit_status = outArgs[2].Int32()
	result = ret.Bool()
	return
}

// g_spawn_error_quark
// container is nil
func SpawnErrorQuark() (result uint32) {
	iv, err := _I.Get(1113, "spawn_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_spawn_exit_error_quark
// container is nil
func SpawnExitErrorQuark() (result uint32) {
	iv, err := _I.Get(1114, "spawn_exit_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_spawn_sync
// container is nil
func SpawnSync(working_directory string, argv gi.CStrArray, envp gi.CStrArray, flags SpawnFlags, child_setup int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer) (result bool, standard_output gi.Uint8Array, standard_error gi.Uint8Array, exit_status int32, err error) {
	iv, err := _I.Get(1115, "spawn_sync", "")
	if err != nil {
		return
	}
	var outArgs [4]gi.Argument
	c_working_directory := gi.CString(working_directory)
	arg_working_directory := gi.NewStringArgument(c_working_directory)
	arg_argv := gi.NewPointerArgument(argv.P)
	arg_envp := gi.NewPointerArgument(envp.P)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_child_setup := gi.NewIntArgument(child_setup) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_standard_output := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_standard_error := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_exit_status := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[3]))
	args := []gi.Argument{arg_working_directory, arg_argv, arg_envp, arg_flags, arg_child_setup, arg_user_data, arg_standard_output, arg_standard_error, arg_exit_status, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_working_directory)
	err = gi.ToError(outArgs[3].Pointer())
	standard_output.P = outArgs[0].Pointer()
	standard_error.P = outArgs[1].Pointer()
	exit_status = outArgs[2].Int32()
	result = ret.Bool()
	return
}

// g_stpcpy
// container is nil
func Stpcpy(dest string, src string) (result string) {
	iv, err := _I.Get(1116, "stpcpy", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_dest := gi.CString(dest)
	c_src := gi.CString(src)
	arg_dest := gi.NewStringArgument(c_dest)
	arg_src := gi.NewStringArgument(c_src)
	args := []gi.Argument{arg_dest, arg_src}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_dest)
	gi.Free(c_src)
	result = ret.String().Take()
	return
}

// g_str_equal
// container is nil
func StrEqual(v1 unsafe.Pointer, v2 unsafe.Pointer) (result bool) {
	iv, err := _I.Get(1117, "str_equal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v1 := gi.NewPointerArgument(v1)
	arg_v2 := gi.NewPointerArgument(v2)
	args := []gi.Argument{arg_v1, arg_v2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_str_has_prefix
// container is nil
func StrHasPrefix(str string, prefix string) (result bool) {
	iv, err := _I.Get(1118, "str_has_prefix", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	c_prefix := gi.CString(prefix)
	arg_str := gi.NewStringArgument(c_str)
	arg_prefix := gi.NewStringArgument(c_prefix)
	args := []gi.Argument{arg_str, arg_prefix}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	gi.Free(c_prefix)
	result = ret.Bool()
	return
}

// g_str_has_suffix
// container is nil
func StrHasSuffix(str string, suffix string) (result bool) {
	iv, err := _I.Get(1119, "str_has_suffix", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	c_suffix := gi.CString(suffix)
	arg_str := gi.NewStringArgument(c_str)
	arg_suffix := gi.NewStringArgument(c_suffix)
	args := []gi.Argument{arg_str, arg_suffix}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	gi.Free(c_suffix)
	result = ret.Bool()
	return
}

// g_str_hash
// container is nil
func StrHash(v unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(1120, "str_hash", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_str_is_ascii
// container is nil
func StrIsAscii(str string) (result bool) {
	iv, err := _I.Get(1121, "str_is_ascii", "")
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
	result = ret.Bool()
	return
}

// g_str_match_string
// container is nil
func StrMatchString(search_term string, potential_hit string, accept_alternates bool) (result bool) {
	iv, err := _I.Get(1122, "str_match_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_search_term := gi.CString(search_term)
	c_potential_hit := gi.CString(potential_hit)
	arg_search_term := gi.NewStringArgument(c_search_term)
	arg_potential_hit := gi.NewStringArgument(c_potential_hit)
	arg_accept_alternates := gi.NewBoolArgument(accept_alternates)
	args := []gi.Argument{arg_search_term, arg_potential_hit, arg_accept_alternates}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_search_term)
	gi.Free(c_potential_hit)
	result = ret.Bool()
	return
}

// g_str_to_ascii
// container is nil
func StrToAscii(str string, from_locale string) (result string) {
	iv, err := _I.Get(1123, "str_to_ascii", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	c_from_locale := gi.CString(from_locale)
	arg_str := gi.NewStringArgument(c_str)
	arg_from_locale := gi.NewStringArgument(c_from_locale)
	args := []gi.Argument{arg_str, arg_from_locale}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	gi.Free(c_from_locale)
	result = ret.String().Take()
	return
}

// g_str_tokenize_and_fold
// container is nil
func StrTokenizeAndFold(string string, translit_locale string) (result gi.CStrArray, ascii_alternates gi.CStrArray) {
	iv, err := _I.Get(1124, "str_tokenize_and_fold", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_string := gi.CString(string)
	c_translit_locale := gi.CString(translit_locale)
	arg_string := gi.NewStringArgument(c_string)
	arg_translit_locale := gi.NewStringArgument(c_translit_locale)
	arg_ascii_alternates := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_string, arg_translit_locale, arg_ascii_alternates}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_string)
	gi.Free(c_translit_locale)
	ascii_alternates.P = outArgs[0].Pointer()
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_strcanon
// container is nil
func Strcanon(string string, valid_chars string, substitutor int8) (result string) {
	iv, err := _I.Get(1125, "strcanon", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	c_valid_chars := gi.CString(valid_chars)
	arg_string := gi.NewStringArgument(c_string)
	arg_valid_chars := gi.NewStringArgument(c_valid_chars)
	arg_substitutor := gi.NewInt8Argument(substitutor)
	args := []gi.Argument{arg_string, arg_valid_chars, arg_substitutor}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	gi.Free(c_valid_chars)
	result = ret.String().Take()
	return
}

// g_strcasecmp
// container is nil
func Strcasecmp(s1 string, s2 string) (result int32) {
	iv, err := _I.Get(1126, "strcasecmp", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_s1 := gi.CString(s1)
	c_s2 := gi.CString(s2)
	arg_s1 := gi.NewStringArgument(c_s1)
	arg_s2 := gi.NewStringArgument(c_s2)
	args := []gi.Argument{arg_s1, arg_s2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_s1)
	gi.Free(c_s2)
	result = ret.Int32()
	return
}

// g_strchomp
// container is nil
func Strchomp(string string) (result string) {
	iv, err := _I.Get(1127, "strchomp", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_strchug
// container is nil
func Strchug(string string) (result string) {
	iv, err := _I.Get(1128, "strchug", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_strcmp0
// container is nil
func Strcmp0(str1 string, str2 string) (result int32) {
	iv, err := _I.Get(1129, "strcmp0", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str1 := gi.CString(str1)
	c_str2 := gi.CString(str2)
	arg_str1 := gi.NewStringArgument(c_str1)
	arg_str2 := gi.NewStringArgument(c_str2)
	args := []gi.Argument{arg_str1, arg_str2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str1)
	gi.Free(c_str2)
	result = ret.Int32()
	return
}

// g_strcompress
// container is nil
func Strcompress(source string) (result string) {
	iv, err := _I.Get(1130, "strcompress", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_source := gi.CString(source)
	arg_source := gi.NewStringArgument(c_source)
	args := []gi.Argument{arg_source}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_source)
	result = ret.String().Take()
	return
}

// g_strdelimit
// container is nil
func Strdelimit(string string, delimiters string, new_delimiter int8) (result string) {
	iv, err := _I.Get(1131, "strdelimit", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	c_delimiters := gi.CString(delimiters)
	arg_string := gi.NewStringArgument(c_string)
	arg_delimiters := gi.NewStringArgument(c_delimiters)
	arg_new_delimiter := gi.NewInt8Argument(new_delimiter)
	args := []gi.Argument{arg_string, arg_delimiters, arg_new_delimiter}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	gi.Free(c_delimiters)
	result = ret.String().Take()
	return
}

// g_strdown
// container is nil
func Strdown(string string) (result string) {
	iv, err := _I.Get(1132, "strdown", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_strdup
// container is nil
func Strdup(str string) (result string) {
	iv, err := _I.Get(1133, "strdup", "")
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

// g_strerror
// container is nil
func Strerror(errnum int32) (result string) {
	iv, err := _I.Get(1134, "strerror", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_errnum := gi.NewInt32Argument(errnum)
	args := []gi.Argument{arg_errnum}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_strescape
// container is nil
func Strescape(source string, exceptions string) (result string) {
	iv, err := _I.Get(1135, "strescape", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_source := gi.CString(source)
	c_exceptions := gi.CString(exceptions)
	arg_source := gi.NewStringArgument(c_source)
	arg_exceptions := gi.NewStringArgument(c_exceptions)
	args := []gi.Argument{arg_source, arg_exceptions}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_source)
	gi.Free(c_exceptions)
	result = ret.String().Take()
	return
}

// g_strfreev
// container is nil
func Strfreev(str_array string) {
	iv, err := _I.Get(1136, "strfreev", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str_array := gi.CString(str_array)
	arg_str_array := gi.NewStringArgument(c_str_array)
	args := []gi.Argument{arg_str_array}
	iv.Call(args, nil, nil)
	gi.Free(c_str_array)
}

// g_string_new
// container is nil
func StringNew(init string) (result String) {
	iv, err := _I.Get(1137, "string_new", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_init := gi.CString(init)
	arg_init := gi.NewStringArgument(c_init)
	args := []gi.Argument{arg_init}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_init)
	result.P = ret.Pointer()
	return
}

// g_string_new_len
// container is nil
func StringNewLen(init string, len1 int64) (result String) {
	iv, err := _I.Get(1138, "string_new_len", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_init := gi.CString(init)
	arg_init := gi.NewStringArgument(c_init)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_init, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_init)
	result.P = ret.Pointer()
	return
}

// g_string_sized_new
// container is nil
func StringSizedNew(dfl_size uint64) (result String) {
	iv, err := _I.Get(1139, "string_sized_new", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_dfl_size := gi.NewUint64Argument(dfl_size)
	args := []gi.Argument{arg_dfl_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_strip_context
// container is nil
func StripContext(msgid string, msgval string) (result string) {
	iv, err := _I.Get(1140, "strip_context", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_msgid := gi.CString(msgid)
	c_msgval := gi.CString(msgval)
	arg_msgid := gi.NewStringArgument(c_msgid)
	arg_msgval := gi.NewStringArgument(c_msgval)
	args := []gi.Argument{arg_msgid, arg_msgval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_msgid)
	gi.Free(c_msgval)
	result = ret.String().Take()
	return
}

// g_strjoinv
// container is nil
func Strjoinv(separator string, str_array string) (result string) {
	iv, err := _I.Get(1141, "strjoinv", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_separator := gi.CString(separator)
	c_str_array := gi.CString(str_array)
	arg_separator := gi.NewStringArgument(c_separator)
	arg_str_array := gi.NewStringArgument(c_str_array)
	args := []gi.Argument{arg_separator, arg_str_array}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_separator)
	gi.Free(c_str_array)
	result = ret.String().Take()
	return
}

// g_strlcat
// container is nil
func Strlcat(dest string, src string, dest_size uint64) (result uint64) {
	iv, err := _I.Get(1142, "strlcat", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_dest := gi.CString(dest)
	c_src := gi.CString(src)
	arg_dest := gi.NewStringArgument(c_dest)
	arg_src := gi.NewStringArgument(c_src)
	arg_dest_size := gi.NewUint64Argument(dest_size)
	args := []gi.Argument{arg_dest, arg_src, arg_dest_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_dest)
	gi.Free(c_src)
	result = ret.Uint64()
	return
}

// g_strlcpy
// container is nil
func Strlcpy(dest string, src string, dest_size uint64) (result uint64) {
	iv, err := _I.Get(1143, "strlcpy", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_dest := gi.CString(dest)
	c_src := gi.CString(src)
	arg_dest := gi.NewStringArgument(c_dest)
	arg_src := gi.NewStringArgument(c_src)
	arg_dest_size := gi.NewUint64Argument(dest_size)
	args := []gi.Argument{arg_dest, arg_src, arg_dest_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_dest)
	gi.Free(c_src)
	result = ret.Uint64()
	return
}

// g_strncasecmp
// container is nil
func Strncasecmp(s1 string, s2 string, n uint32) (result int32) {
	iv, err := _I.Get(1144, "strncasecmp", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_s1 := gi.CString(s1)
	c_s2 := gi.CString(s2)
	arg_s1 := gi.NewStringArgument(c_s1)
	arg_s2 := gi.NewStringArgument(c_s2)
	arg_n := gi.NewUint32Argument(n)
	args := []gi.Argument{arg_s1, arg_s2, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_s1)
	gi.Free(c_s2)
	result = ret.Int32()
	return
}

// g_strndup
// container is nil
func Strndup(str string, n uint64) (result string) {
	iv, err := _I.Get(1145, "strndup", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_n := gi.NewUint64Argument(n)
	args := []gi.Argument{arg_str, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_strnfill
// container is nil
func Strnfill(length uint64, fill_char int8) (result string) {
	iv, err := _I.Get(1146, "strnfill", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_length := gi.NewUint64Argument(length)
	arg_fill_char := gi.NewInt8Argument(fill_char)
	args := []gi.Argument{arg_length, arg_fill_char}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_strreverse
// container is nil
func Strreverse(string string) (result string) {
	iv, err := _I.Get(1147, "strreverse", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_strrstr
// container is nil
func Strrstr(haystack string, needle string) (result string) {
	iv, err := _I.Get(1148, "strrstr", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_haystack := gi.CString(haystack)
	c_needle := gi.CString(needle)
	arg_haystack := gi.NewStringArgument(c_haystack)
	arg_needle := gi.NewStringArgument(c_needle)
	args := []gi.Argument{arg_haystack, arg_needle}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_haystack)
	gi.Free(c_needle)
	result = ret.String().Take()
	return
}

// g_strrstr_len
// container is nil
func StrrstrLen(haystack string, haystack_len int64, needle string) (result string) {
	iv, err := _I.Get(1149, "strrstr_len", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_haystack := gi.CString(haystack)
	c_needle := gi.CString(needle)
	arg_haystack := gi.NewStringArgument(c_haystack)
	arg_haystack_len := gi.NewInt64Argument(haystack_len)
	arg_needle := gi.NewStringArgument(c_needle)
	args := []gi.Argument{arg_haystack, arg_haystack_len, arg_needle}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_haystack)
	gi.Free(c_needle)
	result = ret.String().Take()
	return
}

// g_strsignal
// container is nil
func Strsignal(signum int32) (result string) {
	iv, err := _I.Get(1150, "strsignal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_signum := gi.NewInt32Argument(signum)
	args := []gi.Argument{arg_signum}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_strstr_len
// container is nil
func StrstrLen(haystack string, haystack_len int64, needle string) (result string) {
	iv, err := _I.Get(1151, "strstr_len", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_haystack := gi.CString(haystack)
	c_needle := gi.CString(needle)
	arg_haystack := gi.NewStringArgument(c_haystack)
	arg_haystack_len := gi.NewInt64Argument(haystack_len)
	arg_needle := gi.NewStringArgument(c_needle)
	args := []gi.Argument{arg_haystack, arg_haystack_len, arg_needle}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_haystack)
	gi.Free(c_needle)
	result = ret.String().Take()
	return
}

// g_strtod
// container is nil
func Strtod(nptr string) (result float64, endptr string) {
	iv, err := _I.Get(1152, "strtod", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_nptr := gi.CString(nptr)
	arg_nptr := gi.NewStringArgument(c_nptr)
	arg_endptr := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_nptr, arg_endptr}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_nptr)
	endptr = outArgs[0].String().Take()
	result = ret.Double()
	return
}

// g_strup
// container is nil
func Strup(string string) (result string) {
	iv, err := _I.Get(1153, "strup", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.String().Take()
	return
}

// g_strv_contains
// container is nil
func StrvContains(strv string, str string) (result bool) {
	iv, err := _I.Get(1154, "strv_contains", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_strv := gi.CString(strv)
	c_str := gi.CString(str)
	arg_strv := gi.NewStringArgument(c_strv)
	arg_str := gi.NewStringArgument(c_str)
	args := []gi.Argument{arg_strv, arg_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_strv)
	gi.Free(c_str)
	result = ret.Bool()
	return
}

// g_strv_get_type
// container is nil
func StrvGetType() (result gi.GType) {
	iv, err := _I.Get(1155, "strv_get_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// g_strv_length
// container is nil
func StrvLength(str_array string) (result uint32) {
	iv, err := _I.Get(1156, "strv_length", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str_array := gi.CString(str_array)
	arg_str_array := gi.NewStringArgument(c_str_array)
	args := []gi.Argument{arg_str_array}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str_array)
	result = ret.Uint32()
	return
}

// g_test_add_data_func
// container is nil
func TestAddDataFunc(testpath string, test_data unsafe.Pointer, test_func int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(1157, "test_add_data_func", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_testpath := gi.CString(testpath)
	arg_testpath := gi.NewStringArgument(c_testpath)
	arg_test_data := gi.NewPointerArgument(test_data)
	arg_test_func := gi.NewIntArgument(test_func) /*TODO*/
	args := []gi.Argument{arg_testpath, arg_test_data, arg_test_func}
	iv.Call(args, nil, nil)
	gi.Free(c_testpath)
}

// g_test_add_data_func_full
// container is nil
func TestAddDataFuncFull(testpath string, test_data unsafe.Pointer, test_func int /*TODO_TYPE isPtr: false, tag: interface*/, data_free_func int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(1158, "test_add_data_func_full", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_testpath := gi.CString(testpath)
	arg_testpath := gi.NewStringArgument(c_testpath)
	arg_test_data := gi.NewPointerArgument(test_data)
	arg_test_func := gi.NewIntArgument(test_func)           /*TODO*/
	arg_data_free_func := gi.NewIntArgument(data_free_func) /*TODO*/
	args := []gi.Argument{arg_testpath, arg_test_data, arg_test_func, arg_data_free_func}
	iv.Call(args, nil, nil)
	gi.Free(c_testpath)
}

// g_test_add_func
// container is nil
func TestAddFunc(testpath string, test_func int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(1159, "test_add_func", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_testpath := gi.CString(testpath)
	arg_testpath := gi.NewStringArgument(c_testpath)
	arg_test_func := gi.NewIntArgument(test_func) /*TODO*/
	args := []gi.Argument{arg_testpath, arg_test_func}
	iv.Call(args, nil, nil)
	gi.Free(c_testpath)
}

// g_test_assert_expected_messages_internal
// container is nil
func TestAssertExpectedMessagesInternal(domain string, file string, line int32, func1 string) {
	iv, err := _I.Get(1160, "test_assert_expected_messages_internal", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_file := gi.CString(file)
	c_func1 := gi.CString(func1)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_file := gi.NewStringArgument(c_file)
	arg_line := gi.NewInt32Argument(line)
	arg_func1 := gi.NewStringArgument(c_func1)
	args := []gi.Argument{arg_domain, arg_file, arg_line, arg_func1}
	iv.Call(args, nil, nil)
	gi.Free(c_domain)
	gi.Free(c_file)
	gi.Free(c_func1)
}

// g_test_bug
// container is nil
func TestBug(bug_uri_snippet string) {
	iv, err := _I.Get(1161, "test_bug", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_bug_uri_snippet := gi.CString(bug_uri_snippet)
	arg_bug_uri_snippet := gi.NewStringArgument(c_bug_uri_snippet)
	args := []gi.Argument{arg_bug_uri_snippet}
	iv.Call(args, nil, nil)
	gi.Free(c_bug_uri_snippet)
}

// g_test_bug_base
// container is nil
func TestBugBase(uri_pattern string) {
	iv, err := _I.Get(1162, "test_bug_base", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri_pattern := gi.CString(uri_pattern)
	arg_uri_pattern := gi.NewStringArgument(c_uri_pattern)
	args := []gi.Argument{arg_uri_pattern}
	iv.Call(args, nil, nil)
	gi.Free(c_uri_pattern)
}

// g_test_expect_message
// container is nil
func TestExpectMessage(log_domain string, log_level LogLevelFlags, pattern string) {
	iv, err := _I.Get(1163, "test_expect_message", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_log_domain := gi.CString(log_domain)
	c_pattern := gi.CString(pattern)
	arg_log_domain := gi.NewStringArgument(c_log_domain)
	arg_log_level := gi.NewIntArgument(int(log_level))
	arg_pattern := gi.NewStringArgument(c_pattern)
	args := []gi.Argument{arg_log_domain, arg_log_level, arg_pattern}
	iv.Call(args, nil, nil)
	gi.Free(c_log_domain)
	gi.Free(c_pattern)
}

// g_test_fail
// container is nil
func TestFail() {
	iv, err := _I.Get(1164, "test_fail", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// g_test_failed
// container is nil
func TestFailed() (result bool) {
	iv, err := _I.Get(1165, "test_failed", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// g_test_get_dir
// container is nil
func TestGetDir(file_type TestFileTypeEnum) (result string) {
	iv, err := _I.Get(1166, "test_get_dir", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_file_type := gi.NewIntArgument(int(file_type))
	args := []gi.Argument{arg_file_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_test_incomplete
// container is nil
func TestIncomplete(msg string) {
	iv, err := _I.Get(1167, "test_incomplete", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_msg := gi.CString(msg)
	arg_msg := gi.NewStringArgument(c_msg)
	args := []gi.Argument{arg_msg}
	iv.Call(args, nil, nil)
	gi.Free(c_msg)
}

// g_test_log_type_name
// container is nil
func TestLogTypeName(log_type TestLogTypeEnum) (result string) {
	iv, err := _I.Get(1168, "test_log_type_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_log_type := gi.NewIntArgument(int(log_type))
	args := []gi.Argument{arg_log_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_test_queue_destroy
// container is nil
func TestQueueDestroy(destroy_func int /*TODO_TYPE isPtr: false, tag: interface*/, destroy_data unsafe.Pointer) {
	iv, err := _I.Get(1169, "test_queue_destroy", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_destroy_func := gi.NewIntArgument(destroy_func) /*TODO*/
	arg_destroy_data := gi.NewPointerArgument(destroy_data)
	args := []gi.Argument{arg_destroy_func, arg_destroy_data}
	iv.Call(args, nil, nil)
}

// g_test_queue_free
// container is nil
func TestQueueFree(gfree_pointer unsafe.Pointer) {
	iv, err := _I.Get(1170, "test_queue_free", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_gfree_pointer := gi.NewPointerArgument(gfree_pointer)
	args := []gi.Argument{arg_gfree_pointer}
	iv.Call(args, nil, nil)
}

// g_test_rand_double
// container is nil
func TestRandDouble() (result float64) {
	iv, err := _I.Get(1171, "test_rand_double", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Double()
	return
}

// g_test_rand_double_range
// container is nil
func TestRandDoubleRange(range_start float64, range_end float64) (result float64) {
	iv, err := _I.Get(1172, "test_rand_double_range", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_range_start := gi.NewDoubleArgument(range_start)
	arg_range_end := gi.NewDoubleArgument(range_end)
	args := []gi.Argument{arg_range_start, arg_range_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Double()
	return
}

// g_test_rand_int
// container is nil
func TestRandInt() (result int32) {
	iv, err := _I.Get(1173, "test_rand_int", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int32()
	return
}

// g_test_rand_int_range
// container is nil
func TestRandIntRange(begin int32, end int32) (result int32) {
	iv, err := _I.Get(1174, "test_rand_int_range", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_begin := gi.NewInt32Argument(begin)
	arg_end := gi.NewInt32Argument(end)
	args := []gi.Argument{arg_begin, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_test_run
// container is nil
func TestRun() (result int32) {
	iv, err := _I.Get(1175, "test_run", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int32()
	return
}

// g_test_run_suite
// container is nil
func TestRunSuite(suite TestSuite) (result int32) {
	iv, err := _I.Get(1176, "test_run_suite", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_suite := gi.NewPointerArgument(suite.P)
	args := []gi.Argument{arg_suite}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_test_set_nonfatal_assertions
// container is nil
func TestSetNonfatalAssertions() {
	iv, err := _I.Get(1177, "test_set_nonfatal_assertions", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// g_test_skip
// container is nil
func TestSkip(msg string) {
	iv, err := _I.Get(1178, "test_skip", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_msg := gi.CString(msg)
	arg_msg := gi.NewStringArgument(c_msg)
	args := []gi.Argument{arg_msg}
	iv.Call(args, nil, nil)
	gi.Free(c_msg)
}

// g_test_subprocess
// container is nil
func TestSubprocess() (result bool) {
	iv, err := _I.Get(1179, "test_subprocess", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// g_test_timer_elapsed
// container is nil
func TestTimerElapsed() (result float64) {
	iv, err := _I.Get(1180, "test_timer_elapsed", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Double()
	return
}

// g_test_timer_last
// container is nil
func TestTimerLast() (result float64) {
	iv, err := _I.Get(1181, "test_timer_last", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Double()
	return
}

// g_test_timer_start
// container is nil
func TestTimerStart() {
	iv, err := _I.Get(1182, "test_timer_start", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// g_test_trap_assertions
// container is nil
func TestTrapAssertions(domain string, file string, line int32, func1 string, assertion_flags uint64, pattern string) {
	iv, err := _I.Get(1183, "test_trap_assertions", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_domain := gi.CString(domain)
	c_file := gi.CString(file)
	c_func1 := gi.CString(func1)
	c_pattern := gi.CString(pattern)
	arg_domain := gi.NewStringArgument(c_domain)
	arg_file := gi.NewStringArgument(c_file)
	arg_line := gi.NewInt32Argument(line)
	arg_func1 := gi.NewStringArgument(c_func1)
	arg_assertion_flags := gi.NewUint64Argument(assertion_flags)
	arg_pattern := gi.NewStringArgument(c_pattern)
	args := []gi.Argument{arg_domain, arg_file, arg_line, arg_func1, arg_assertion_flags, arg_pattern}
	iv.Call(args, nil, nil)
	gi.Free(c_domain)
	gi.Free(c_file)
	gi.Free(c_func1)
	gi.Free(c_pattern)
}

// g_test_trap_fork
// container is nil
func TestTrapFork(usec_timeout uint64, test_trap_flags TestTrapFlags) (result bool) {
	iv, err := _I.Get(1184, "test_trap_fork", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_usec_timeout := gi.NewUint64Argument(usec_timeout)
	arg_test_trap_flags := gi.NewIntArgument(int(test_trap_flags))
	args := []gi.Argument{arg_usec_timeout, arg_test_trap_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_test_trap_has_passed
// container is nil
func TestTrapHasPassed() (result bool) {
	iv, err := _I.Get(1185, "test_trap_has_passed", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// g_test_trap_reached_timeout
// container is nil
func TestTrapReachedTimeout() (result bool) {
	iv, err := _I.Get(1186, "test_trap_reached_timeout", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Bool()
	return
}

// g_test_trap_subprocess
// container is nil
func TestTrapSubprocess(test_path string, usec_timeout uint64, test_flags TestSubprocessFlags) {
	iv, err := _I.Get(1187, "test_trap_subprocess", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_test_path := gi.CString(test_path)
	arg_test_path := gi.NewStringArgument(c_test_path)
	arg_usec_timeout := gi.NewUint64Argument(usec_timeout)
	arg_test_flags := gi.NewIntArgument(int(test_flags))
	args := []gi.Argument{arg_test_path, arg_usec_timeout, arg_test_flags}
	iv.Call(args, nil, nil)
	gi.Free(c_test_path)
}

// g_thread_error_quark
// container is nil
func ThreadErrorQuark() (result uint32) {
	iv, err := _I.Get(1188, "thread_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_thread_exit
// container is nil
func ThreadExit(retval unsafe.Pointer) {
	iv, err := _I.Get(1189, "thread_exit", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_retval := gi.NewPointerArgument(retval)
	args := []gi.Argument{arg_retval}
	iv.Call(args, nil, nil)
}

// g_thread_pool_get_max_idle_time
// container is nil
func ThreadPoolGetMaxIdleTime() (result uint32) {
	iv, err := _I.Get(1190, "thread_pool_get_max_idle_time", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_thread_pool_get_max_unused_threads
// container is nil
func ThreadPoolGetMaxUnusedThreads() (result int32) {
	iv, err := _I.Get(1191, "thread_pool_get_max_unused_threads", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Int32()
	return
}

// g_thread_pool_get_num_unused_threads
// container is nil
func ThreadPoolGetNumUnusedThreads() (result uint32) {
	iv, err := _I.Get(1192, "thread_pool_get_num_unused_threads", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_thread_pool_set_max_idle_time
// container is nil
func ThreadPoolSetMaxIdleTime(interval uint32) {
	iv, err := _I.Get(1193, "thread_pool_set_max_idle_time", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_interval := gi.NewUint32Argument(interval)
	args := []gi.Argument{arg_interval}
	iv.Call(args, nil, nil)
}

// g_thread_pool_set_max_unused_threads
// container is nil
func ThreadPoolSetMaxUnusedThreads(max_threads int32) {
	iv, err := _I.Get(1194, "thread_pool_set_max_unused_threads", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_max_threads := gi.NewInt32Argument(max_threads)
	args := []gi.Argument{arg_max_threads}
	iv.Call(args, nil, nil)
}

// g_thread_pool_stop_unused_threads
// container is nil
func ThreadPoolStopUnusedThreads() {
	iv, err := _I.Get(1195, "thread_pool_stop_unused_threads", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// g_thread_self
// container is nil
func ThreadSelf() (result Thread) {
	iv, err := _I.Get(1196, "thread_self", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_thread_yield
// container is nil
func ThreadYield() {
	iv, err := _I.Get(1197, "thread_yield", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// g_time_val_from_iso8601
// container is nil
func TimeValFromIso8601(iso_date string, time_ TimeVal) (result bool) {
	iv, err := _I.Get(1198, "time_val_from_iso8601", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_iso_date := gi.CString(iso_date)
	arg_iso_date := gi.NewStringArgument(c_iso_date)
	arg_time_ := gi.NewPointerArgument(time_.P)
	args := []gi.Argument{arg_iso_date, arg_time_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_iso_date)
	result = ret.Bool()
	return
}

// g_timeout_add_full
// container is nil
func TimeoutAdd(priority int32, interval uint32, function int /*TODO_TYPE isPtr: false, tag: interface*/, data unsafe.Pointer, notify int /*TODO_TYPE isPtr: false, tag: interface*/) (result uint32) {
	iv, err := _I.Get(1199, "timeout_add", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_priority := gi.NewInt32Argument(priority)
	arg_interval := gi.NewUint32Argument(interval)
	arg_function := gi.NewIntArgument(function) /*TODO*/
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewIntArgument(notify) /*TODO*/
	args := []gi.Argument{arg_priority, arg_interval, arg_function, arg_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_timeout_add_seconds_full
// container is nil
func TimeoutAddSeconds(priority int32, interval uint32, function int /*TODO_TYPE isPtr: false, tag: interface*/, data unsafe.Pointer, notify int /*TODO_TYPE isPtr: false, tag: interface*/) (result uint32) {
	iv, err := _I.Get(1200, "timeout_add_seconds", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_priority := gi.NewInt32Argument(priority)
	arg_interval := gi.NewUint32Argument(interval)
	arg_function := gi.NewIntArgument(function) /*TODO*/
	arg_data := gi.NewPointerArgument(data)
	arg_notify := gi.NewIntArgument(notify) /*TODO*/
	args := []gi.Argument{arg_priority, arg_interval, arg_function, arg_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_timeout_source_new
// container is nil
func TimeoutSourceNew(interval uint32) (result Source) {
	iv, err := _I.Get(1201, "timeout_source_new", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_interval := gi.NewUint32Argument(interval)
	args := []gi.Argument{arg_interval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_timeout_source_new_seconds
// container is nil
func TimeoutSourceNewSeconds(interval uint32) (result Source) {
	iv, err := _I.Get(1202, "timeout_source_new_seconds", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_interval := gi.NewUint32Argument(interval)
	args := []gi.Argument{arg_interval}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_trash_stack_height
// container is nil
func TrashStackHeight(stack_p TrashStack) (result uint32) {
	iv, err := _I.Get(1203, "trash_stack_height", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_stack_p := gi.NewPointerArgument(stack_p.P)
	args := []gi.Argument{arg_stack_p}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_trash_stack_peek
// container is nil
func TrashStackPeek(stack_p TrashStack) (result unsafe.Pointer) {
	iv, err := _I.Get(1204, "trash_stack_peek", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_stack_p := gi.NewPointerArgument(stack_p.P)
	args := []gi.Argument{arg_stack_p}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_trash_stack_pop
// container is nil
func TrashStackPop(stack_p TrashStack) (result unsafe.Pointer) {
	iv, err := _I.Get(1205, "trash_stack_pop", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_stack_p := gi.NewPointerArgument(stack_p.P)
	args := []gi.Argument{arg_stack_p}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_trash_stack_push
// container is nil
func TrashStackPush(stack_p TrashStack, data_p unsafe.Pointer) {
	iv, err := _I.Get(1206, "trash_stack_push", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_stack_p := gi.NewPointerArgument(stack_p.P)
	arg_data_p := gi.NewPointerArgument(data_p)
	args := []gi.Argument{arg_stack_p, arg_data_p}
	iv.Call(args, nil, nil)
}

// g_try_malloc
// container is nil
func TryMalloc(n_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1207, "try_malloc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_bytes := gi.NewUint64Argument(n_bytes)
	args := []gi.Argument{arg_n_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_try_malloc0
// container is nil
func TryMalloc0(n_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1208, "try_malloc0", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_bytes := gi.NewUint64Argument(n_bytes)
	args := []gi.Argument{arg_n_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_try_malloc0_n
// container is nil
func TryMalloc0N(n_blocks uint64, n_block_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1209, "try_malloc0_n", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_blocks := gi.NewUint64Argument(n_blocks)
	arg_n_block_bytes := gi.NewUint64Argument(n_block_bytes)
	args := []gi.Argument{arg_n_blocks, arg_n_block_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_try_malloc_n
// container is nil
func TryMallocN(n_blocks uint64, n_block_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1210, "try_malloc_n", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_blocks := gi.NewUint64Argument(n_blocks)
	arg_n_block_bytes := gi.NewUint64Argument(n_block_bytes)
	args := []gi.Argument{arg_n_blocks, arg_n_block_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_try_realloc
// container is nil
func TryRealloc(mem unsafe.Pointer, n_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1211, "try_realloc", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem := gi.NewPointerArgument(mem)
	arg_n_bytes := gi.NewUint64Argument(n_bytes)
	args := []gi.Argument{arg_mem, arg_n_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_try_realloc_n
// container is nil
func TryReallocN(mem unsafe.Pointer, n_blocks uint64, n_block_bytes uint64) (result unsafe.Pointer) {
	iv, err := _I.Get(1212, "try_realloc_n", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_mem := gi.NewPointerArgument(mem)
	arg_n_blocks := gi.NewUint64Argument(n_blocks)
	arg_n_block_bytes := gi.NewUint64Argument(n_block_bytes)
	args := []gi.Argument{arg_mem, arg_n_blocks, arg_n_block_bytes}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_ucs4_to_utf16
// container is nil
func Ucs4ToUtf16(str rune, len1 int64) (result uint16, items_read int64, items_written int64, err error) {
	iv, err := _I.Get(1213, "ucs4_to_utf16", "")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_str := gi.NewUint32Argument(uint32(str))
	arg_len1 := gi.NewInt64Argument(len1)
	arg_items_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_items_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_items_read, arg_items_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	items_read = outArgs[0].Int64()
	items_written = outArgs[1].Int64()
	result = ret.Uint16()
	return
}

// g_ucs4_to_utf8
// container is nil
func Ucs4ToUtf8(str rune, len1 int64) (result string, items_read int64, items_written int64, err error) {
	iv, err := _I.Get(1214, "ucs4_to_utf8", "")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_str := gi.NewUint32Argument(uint32(str))
	arg_len1 := gi.NewInt64Argument(len1)
	arg_items_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_items_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_items_read, arg_items_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	items_read = outArgs[0].Int64()
	items_written = outArgs[1].Int64()
	result = ret.String().Take()
	return
}

// g_unichar_break_type
// container is nil
func UnicharBreakType(c rune) (result UnicodeBreakTypeEnum) {
	iv, err := _I.Get(1215, "unichar_break_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = UnicodeBreakTypeEnum(ret.Int())
	return
}

// g_unichar_combining_class
// container is nil
func UnicharCombiningClass(uc rune) (result int32) {
	iv, err := _I.Get(1216, "unichar_combining_class", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_uc := gi.NewUint32Argument(uint32(uc))
	args := []gi.Argument{arg_uc}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_unichar_compose
// container is nil
func UnicharCompose(a rune, b rune, ch rune) (result bool) {
	iv, err := _I.Get(1217, "unichar_compose", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_a := gi.NewUint32Argument(uint32(a))
	arg_b := gi.NewUint32Argument(uint32(b))
	arg_ch := gi.NewUint32Argument(uint32(ch))
	args := []gi.Argument{arg_a, arg_b, arg_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_decompose
// container is nil
func UnicharDecompose(ch rune, a rune, b rune) (result bool) {
	iv, err := _I.Get(1218, "unichar_decompose", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewUint32Argument(uint32(ch))
	arg_a := gi.NewUint32Argument(uint32(a))
	arg_b := gi.NewUint32Argument(uint32(b))
	args := []gi.Argument{arg_ch, arg_a, arg_b}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_digit_value
// container is nil
func UnicharDigitValue(c rune) (result int32) {
	iv, err := _I.Get(1219, "unichar_digit_value", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_unichar_fully_decompose
// container is nil
func UnicharFullyDecompose(ch rune, compat bool, result rune, result_len uint64) (result1 uint64) {
	iv, err := _I.Get(1220, "unichar_fully_decompose", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewUint32Argument(uint32(ch))
	arg_compat := gi.NewBoolArgument(compat)
	arg_result := gi.NewUint32Argument(uint32(result))
	arg_result_len := gi.NewUint64Argument(result_len)
	args := []gi.Argument{arg_ch, arg_compat, arg_result, arg_result_len}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result1 = ret.Uint64()
	return
}

// g_unichar_get_mirror_char
// container is nil
func UnicharGetMirrorChar(ch rune, mirrored_ch rune) (result bool) {
	iv, err := _I.Get(1221, "unichar_get_mirror_char", "")
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

// g_unichar_get_script
// container is nil
func UnicharGetScript(ch rune) (result UnicodeScriptEnum) {
	iv, err := _I.Get(1222, "unichar_get_script", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewUint32Argument(uint32(ch))
	args := []gi.Argument{arg_ch}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = UnicodeScriptEnum(ret.Int())
	return
}

// g_unichar_isalnum
// container is nil
func UnicharIsalnum(c rune) (result bool) {
	iv, err := _I.Get(1223, "unichar_isalnum", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isalpha
// container is nil
func UnicharIsalpha(c rune) (result bool) {
	iv, err := _I.Get(1224, "unichar_isalpha", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_iscntrl
// container is nil
func UnicharIscntrl(c rune) (result bool) {
	iv, err := _I.Get(1225, "unichar_iscntrl", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isdefined
// container is nil
func UnicharIsdefined(c rune) (result bool) {
	iv, err := _I.Get(1226, "unichar_isdefined", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isdigit
// container is nil
func UnicharIsdigit(c rune) (result bool) {
	iv, err := _I.Get(1227, "unichar_isdigit", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isgraph
// container is nil
func UnicharIsgraph(c rune) (result bool) {
	iv, err := _I.Get(1228, "unichar_isgraph", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_islower
// container is nil
func UnicharIslower(c rune) (result bool) {
	iv, err := _I.Get(1229, "unichar_islower", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_ismark
// container is nil
func UnicharIsmark(c rune) (result bool) {
	iv, err := _I.Get(1230, "unichar_ismark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isprint
// container is nil
func UnicharIsprint(c rune) (result bool) {
	iv, err := _I.Get(1231, "unichar_isprint", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_ispunct
// container is nil
func UnicharIspunct(c rune) (result bool) {
	iv, err := _I.Get(1232, "unichar_ispunct", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isspace
// container is nil
func UnicharIsspace(c rune) (result bool) {
	iv, err := _I.Get(1233, "unichar_isspace", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_istitle
// container is nil
func UnicharIstitle(c rune) (result bool) {
	iv, err := _I.Get(1234, "unichar_istitle", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isupper
// container is nil
func UnicharIsupper(c rune) (result bool) {
	iv, err := _I.Get(1235, "unichar_isupper", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_iswide
// container is nil
func UnicharIswide(c rune) (result bool) {
	iv, err := _I.Get(1236, "unichar_iswide", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_iswide_cjk
// container is nil
func UnicharIswideCjk(c rune) (result bool) {
	iv, err := _I.Get(1237, "unichar_iswide_cjk", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_isxdigit
// container is nil
func UnicharIsxdigit(c rune) (result bool) {
	iv, err := _I.Get(1238, "unichar_isxdigit", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_iszerowidth
// container is nil
func UnicharIszerowidth(c rune) (result bool) {
	iv, err := _I.Get(1239, "unichar_iszerowidth", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_unichar_to_utf8
// container is nil
func UnicharToUtf8(c rune) (result int32, outbuf string) {
	iv, err := _I.Get(1240, "unichar_to_utf8", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_c := gi.NewUint32Argument(uint32(c))
	arg_outbuf := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_c, arg_outbuf}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	outbuf = outArgs[0].String().Take()
	result = ret.Int32()
	return
}

// g_unichar_tolower
// container is nil
func UnicharTolower(c rune) (result rune) {
	iv, err := _I.Get(1241, "unichar_tolower", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = rune(ret.Uint32())
	return
}

// g_unichar_totitle
// container is nil
func UnicharTotitle(c rune) (result rune) {
	iv, err := _I.Get(1242, "unichar_totitle", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = rune(ret.Uint32())
	return
}

// g_unichar_toupper
// container is nil
func UnicharToupper(c rune) (result rune) {
	iv, err := _I.Get(1243, "unichar_toupper", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = rune(ret.Uint32())
	return
}

// g_unichar_type
// container is nil
func UnicharType(c rune) (result UnicodeTypeEnum) {
	iv, err := _I.Get(1244, "unichar_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = UnicodeTypeEnum(ret.Int())
	return
}

// g_unichar_validate
// container is nil
func UnicharValidate(ch rune) (result bool) {
	iv, err := _I.Get(1245, "unichar_validate", "")
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

// g_unichar_xdigit_value
// container is nil
func UnicharXdigitValue(c rune) (result int32) {
	iv, err := _I.Get(1246, "unichar_xdigit_value", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_unicode_canonical_decomposition
// container is nil
func UnicodeCanonicalDecomposition(ch rune, result_len uint64) (result rune) {
	iv, err := _I.Get(1247, "unicode_canonical_decomposition", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ch := gi.NewUint32Argument(uint32(ch))
	arg_result_len := gi.NewUint64Argument(result_len)
	args := []gi.Argument{arg_ch, arg_result_len}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = rune(ret.Uint32())
	return
}

// g_unicode_canonical_ordering
// container is nil
func UnicodeCanonicalOrdering(string rune, len1 uint64) {
	iv, err := _I.Get(1248, "unicode_canonical_ordering", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_string := gi.NewUint32Argument(uint32(string))
	arg_len1 := gi.NewUint64Argument(len1)
	args := []gi.Argument{arg_string, arg_len1}
	iv.Call(args, nil, nil)
}

// g_unicode_script_from_iso15924
// container is nil
func UnicodeScriptFromIso15924(iso15924 uint32) (result UnicodeScriptEnum) {
	iv, err := _I.Get(1249, "unicode_script_from_iso15924", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_iso15924 := gi.NewUint32Argument(iso15924)
	args := []gi.Argument{arg_iso15924}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = UnicodeScriptEnum(ret.Int())
	return
}

// g_unicode_script_to_iso15924
// container is nil
func UnicodeScriptToIso15924(script UnicodeScriptEnum) (result uint32) {
	iv, err := _I.Get(1250, "unicode_script_to_iso15924", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_script := gi.NewIntArgument(int(script))
	args := []gi.Argument{arg_script}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_unix_error_quark
// container is nil
func UnixErrorQuark() (result uint32) {
	iv, err := _I.Get(1251, "unix_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_unix_fd_add_full
// container is nil
func UnixFdAddFull(priority int32, fd int32, condition IOConditionFlags, function int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer, notify int /*TODO_TYPE isPtr: false, tag: interface*/) (result uint32) {
	iv, err := _I.Get(1252, "unix_fd_add_full", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_priority := gi.NewInt32Argument(priority)
	arg_fd := gi.NewInt32Argument(fd)
	arg_condition := gi.NewIntArgument(int(condition))
	arg_function := gi.NewIntArgument(function) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewIntArgument(notify) /*TODO*/
	args := []gi.Argument{arg_priority, arg_fd, arg_condition, arg_function, arg_user_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_unix_fd_source_new
// container is nil
func UnixFdSourceNew(fd int32, condition IOConditionFlags) (result Source) {
	iv, err := _I.Get(1253, "unix_fd_source_new", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_fd := gi.NewInt32Argument(fd)
	arg_condition := gi.NewIntArgument(int(condition))
	args := []gi.Argument{arg_fd, arg_condition}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_unix_open_pipe
// container is nil
func UnixOpenPipe(fds int32, flags int32) (result bool, err error) {
	iv, err := _I.Get(1254, "unix_open_pipe", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_fds := gi.NewInt32Argument(fds)
	arg_flags := gi.NewInt32Argument(flags)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_fds, arg_flags, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_unix_set_fd_nonblocking
// container is nil
func UnixSetFdNonblocking(fd int32, nonblock bool) (result bool, err error) {
	iv, err := _I.Get(1255, "unix_set_fd_nonblocking", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	arg_fd := gi.NewInt32Argument(fd)
	arg_nonblock := gi.NewBoolArgument(nonblock)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_fd, arg_nonblock, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[0].Pointer())
	result = ret.Bool()
	return
}

// g_unix_signal_add_full
// container is nil
func UnixSignalAdd(priority int32, signum int32, handler int /*TODO_TYPE isPtr: false, tag: interface*/, user_data unsafe.Pointer, notify int /*TODO_TYPE isPtr: false, tag: interface*/) (result uint32) {
	iv, err := _I.Get(1256, "unix_signal_add", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_priority := gi.NewInt32Argument(priority)
	arg_signum := gi.NewInt32Argument(signum)
	arg_handler := gi.NewIntArgument(handler) /*TODO*/
	arg_user_data := gi.NewPointerArgument(user_data)
	arg_notify := gi.NewIntArgument(notify) /*TODO*/
	args := []gi.Argument{arg_priority, arg_signum, arg_handler, arg_user_data, arg_notify}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_unix_signal_source_new
// container is nil
func UnixSignalSourceNew(signum int32) (result Source) {
	iv, err := _I.Get(1257, "unix_signal_source_new", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_signum := gi.NewInt32Argument(signum)
	args := []gi.Argument{arg_signum}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_unlink
// container is nil
func Unlink(filename string) (result int32) {
	iv, err := _I.Get(1258, "unlink", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_filename := gi.CString(filename)
	arg_filename := gi.NewStringArgument(c_filename)
	args := []gi.Argument{arg_filename}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_filename)
	result = ret.Int32()
	return
}

// g_unsetenv
// container is nil
func Unsetenv(variable string) {
	iv, err := _I.Get(1259, "unsetenv", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_variable := gi.CString(variable)
	arg_variable := gi.NewStringArgument(c_variable)
	args := []gi.Argument{arg_variable}
	iv.Call(args, nil, nil)
	gi.Free(c_variable)
}

// g_uri_escape_string
// container is nil
func UriEscapeString(unescaped string, reserved_chars_allowed string, allow_utf8 bool) (result string) {
	iv, err := _I.Get(1260, "uri_escape_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_unescaped := gi.CString(unescaped)
	c_reserved_chars_allowed := gi.CString(reserved_chars_allowed)
	arg_unescaped := gi.NewStringArgument(c_unescaped)
	arg_reserved_chars_allowed := gi.NewStringArgument(c_reserved_chars_allowed)
	arg_allow_utf8 := gi.NewBoolArgument(allow_utf8)
	args := []gi.Argument{arg_unescaped, arg_reserved_chars_allowed, arg_allow_utf8}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_unescaped)
	gi.Free(c_reserved_chars_allowed)
	result = ret.String().Take()
	return
}

// g_uri_list_extract_uris
// container is nil
func UriListExtractUris(uri_list string) (result gi.CStrArray) {
	iv, err := _I.Get(1261, "uri_list_extract_uris", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri_list := gi.CString(uri_list)
	arg_uri_list := gi.NewStringArgument(c_uri_list)
	args := []gi.Argument{arg_uri_list}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_uri_list)
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_uri_parse_scheme
// container is nil
func UriParseScheme(uri string) (result string) {
	iv, err := _I.Get(1262, "uri_parse_scheme", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_uri := gi.CString(uri)
	arg_uri := gi.NewStringArgument(c_uri)
	args := []gi.Argument{arg_uri}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_uri)
	result = ret.String().Take()
	return
}

// g_uri_unescape_segment
// container is nil
func UriUnescapeSegment(escaped_string string, escaped_string_end string, illegal_characters string) (result string) {
	iv, err := _I.Get(1263, "uri_unescape_segment", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_escaped_string := gi.CString(escaped_string)
	c_escaped_string_end := gi.CString(escaped_string_end)
	c_illegal_characters := gi.CString(illegal_characters)
	arg_escaped_string := gi.NewStringArgument(c_escaped_string)
	arg_escaped_string_end := gi.NewStringArgument(c_escaped_string_end)
	arg_illegal_characters := gi.NewStringArgument(c_illegal_characters)
	args := []gi.Argument{arg_escaped_string, arg_escaped_string_end, arg_illegal_characters}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_escaped_string)
	gi.Free(c_escaped_string_end)
	gi.Free(c_illegal_characters)
	result = ret.String().Take()
	return
}

// g_uri_unescape_string
// container is nil
func UriUnescapeString(escaped_string string, illegal_characters string) (result string) {
	iv, err := _I.Get(1264, "uri_unescape_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_escaped_string := gi.CString(escaped_string)
	c_illegal_characters := gi.CString(illegal_characters)
	arg_escaped_string := gi.NewStringArgument(c_escaped_string)
	arg_illegal_characters := gi.NewStringArgument(c_illegal_characters)
	args := []gi.Argument{arg_escaped_string, arg_illegal_characters}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_escaped_string)
	gi.Free(c_illegal_characters)
	result = ret.String().Take()
	return
}

// g_usleep
// container is nil
func Usleep(microseconds uint64) {
	iv, err := _I.Get(1265, "usleep", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_microseconds := gi.NewUint64Argument(microseconds)
	args := []gi.Argument{arg_microseconds}
	iv.Call(args, nil, nil)
}

// g_utf16_to_ucs4
// container is nil
func Utf16ToUcs4(str uint16, len1 int64) (result rune, items_read int64, items_written int64, err error) {
	iv, err := _I.Get(1266, "utf16_to_ucs4", "")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_str := gi.NewUint16Argument(str)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_items_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_items_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_items_read, arg_items_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	items_read = outArgs[0].Int64()
	items_written = outArgs[1].Int64()
	result = rune(ret.Uint32())
	return
}

// g_utf16_to_utf8
// container is nil
func Utf16ToUtf8(str uint16, len1 int64) (result string, items_read int64, items_written int64, err error) {
	iv, err := _I.Get(1267, "utf16_to_utf8", "")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	arg_str := gi.NewUint16Argument(str)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_items_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_items_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_items_read, arg_items_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	err = gi.ToError(outArgs[2].Pointer())
	items_read = outArgs[0].Int64()
	items_written = outArgs[1].Int64()
	result = ret.String().Take()
	return
}

// g_utf8_casefold
// container is nil
func Utf8Casefold(str string, len1 int64) (result string) {
	iv, err := _I.Get(1268, "utf8_casefold", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_collate
// container is nil
func Utf8Collate(str1 string, str2 string) (result int32) {
	iv, err := _I.Get(1269, "utf8_collate", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str1 := gi.CString(str1)
	c_str2 := gi.CString(str2)
	arg_str1 := gi.NewStringArgument(c_str1)
	arg_str2 := gi.NewStringArgument(c_str2)
	args := []gi.Argument{arg_str1, arg_str2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str1)
	gi.Free(c_str2)
	result = ret.Int32()
	return
}

// g_utf8_collate_key
// container is nil
func Utf8CollateKey(str string, len1 int64) (result string) {
	iv, err := _I.Get(1270, "utf8_collate_key", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_collate_key_for_filename
// container is nil
func Utf8CollateKeyForFilename(str string, len1 int64) (result string) {
	iv, err := _I.Get(1271, "utf8_collate_key_for_filename", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_find_next_char
// container is nil
func Utf8FindNextChar(p string, end string) (result string) {
	iv, err := _I.Get(1272, "utf8_find_next_char", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_p := gi.CString(p)
	c_end := gi.CString(end)
	arg_p := gi.NewStringArgument(c_p)
	arg_end := gi.NewStringArgument(c_end)
	args := []gi.Argument{arg_p, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_p)
	gi.Free(c_end)
	result = ret.String().Take()
	return
}

// g_utf8_find_prev_char
// container is nil
func Utf8FindPrevChar(str string, p string) (result string) {
	iv, err := _I.Get(1273, "utf8_find_prev_char", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	c_p := gi.CString(p)
	arg_str := gi.NewStringArgument(c_str)
	arg_p := gi.NewStringArgument(c_p)
	args := []gi.Argument{arg_str, arg_p}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	gi.Free(c_p)
	result = ret.String().Take()
	return
}

// g_utf8_get_char
// container is nil
func Utf8GetChar(p string) (result rune) {
	iv, err := _I.Get(1274, "utf8_get_char", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_p := gi.CString(p)
	arg_p := gi.NewStringArgument(c_p)
	args := []gi.Argument{arg_p}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_p)
	result = rune(ret.Uint32())
	return
}

// g_utf8_get_char_validated
// container is nil
func Utf8GetCharValidated(p string, max_len int64) (result rune) {
	iv, err := _I.Get(1275, "utf8_get_char_validated", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_p := gi.CString(p)
	arg_p := gi.NewStringArgument(c_p)
	arg_max_len := gi.NewInt64Argument(max_len)
	args := []gi.Argument{arg_p, arg_max_len}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_p)
	result = rune(ret.Uint32())
	return
}

// g_utf8_make_valid
// container is nil
func Utf8MakeValid(str string, len1 int64) (result string) {
	iv, err := _I.Get(1276, "utf8_make_valid", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_normalize
// container is nil
func Utf8Normalize(str string, len1 int64, mode NormalizeModeEnum) (result string) {
	iv, err := _I.Get(1277, "utf8_normalize", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_mode := gi.NewIntArgument(int(mode))
	args := []gi.Argument{arg_str, arg_len1, arg_mode}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_offset_to_pointer
// container is nil
func Utf8OffsetToPointer(str string, offset int64) (result string) {
	iv, err := _I.Get(1278, "utf8_offset_to_pointer", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_offset := gi.NewInt64Argument(offset)
	args := []gi.Argument{arg_str, arg_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_pointer_to_offset
// container is nil
func Utf8PointerToOffset(str string, pos string) (result int64) {
	iv, err := _I.Get(1279, "utf8_pointer_to_offset", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	c_pos := gi.CString(pos)
	arg_str := gi.NewStringArgument(c_str)
	arg_pos := gi.NewStringArgument(c_pos)
	args := []gi.Argument{arg_str, arg_pos}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	gi.Free(c_pos)
	result = ret.Int64()
	return
}

// g_utf8_prev_char
// container is nil
func Utf8PrevChar(p string) (result string) {
	iv, err := _I.Get(1280, "utf8_prev_char", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_p := gi.CString(p)
	arg_p := gi.NewStringArgument(c_p)
	args := []gi.Argument{arg_p}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_p)
	result = ret.String().Take()
	return
}

// g_utf8_strchr
// container is nil
func Utf8Strchr(p string, len1 int64, c rune) (result string) {
	iv, err := _I.Get(1281, "utf8_strchr", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_p := gi.CString(p)
	arg_p := gi.NewStringArgument(c_p)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_p, arg_len1, arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_p)
	result = ret.String().Take()
	return
}

// g_utf8_strdown
// container is nil
func Utf8Strdown(str string, len1 int64) (result string) {
	iv, err := _I.Get(1282, "utf8_strdown", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_strlen
// container is nil
func Utf8Strlen(p string, max int64) (result int64) {
	iv, err := _I.Get(1283, "utf8_strlen", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_p := gi.CString(p)
	arg_p := gi.NewStringArgument(c_p)
	arg_max := gi.NewInt64Argument(max)
	args := []gi.Argument{arg_p, arg_max}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_p)
	result = ret.Int64()
	return
}

// g_utf8_strncpy
// container is nil
func Utf8Strncpy(dest string, src string, n uint64) (result string) {
	iv, err := _I.Get(1284, "utf8_strncpy", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_dest := gi.CString(dest)
	c_src := gi.CString(src)
	arg_dest := gi.NewStringArgument(c_dest)
	arg_src := gi.NewStringArgument(c_src)
	arg_n := gi.NewUint64Argument(n)
	args := []gi.Argument{arg_dest, arg_src, arg_n}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_dest)
	gi.Free(c_src)
	result = ret.String().Take()
	return
}

// g_utf8_strrchr
// container is nil
func Utf8Strrchr(p string, len1 int64, c rune) (result string) {
	iv, err := _I.Get(1285, "utf8_strrchr", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_p := gi.CString(p)
	arg_p := gi.NewStringArgument(c_p)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_c := gi.NewUint32Argument(uint32(c))
	args := []gi.Argument{arg_p, arg_len1, arg_c}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_p)
	result = ret.String().Take()
	return
}

// g_utf8_strreverse
// container is nil
func Utf8Strreverse(str string, len1 int64) (result string) {
	iv, err := _I.Get(1286, "utf8_strreverse", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_strup
// container is nil
func Utf8Strup(str string, len1 int64) (result string) {
	iv, err := _I.Get(1287, "utf8_strup", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	args := []gi.Argument{arg_str, arg_len1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_substring
// container is nil
func Utf8Substring(str string, start_pos int64, end_pos int64) (result string) {
	iv, err := _I.Get(1288, "utf8_substring", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_start_pos := gi.NewInt64Argument(start_pos)
	arg_end_pos := gi.NewInt64Argument(end_pos)
	args := []gi.Argument{arg_str, arg_start_pos, arg_end_pos}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_str)
	result = ret.String().Take()
	return
}

// g_utf8_to_ucs4
// container is nil
func Utf8ToUcs4(str string, len1 int64) (result rune, items_read int64, items_written int64, err error) {
	iv, err := _I.Get(1289, "utf8_to_ucs4", "")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_items_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_items_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_items_read, arg_items_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_str)
	err = gi.ToError(outArgs[2].Pointer())
	items_read = outArgs[0].Int64()
	items_written = outArgs[1].Int64()
	result = rune(ret.Uint32())
	return
}

// g_utf8_to_ucs4_fast
// container is nil
func Utf8ToUcs4Fast(str string, len1 int64) (result rune, items_written int64) {
	iv, err := _I.Get(1290, "utf8_to_ucs4_fast", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_items_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_str, arg_len1, arg_items_written}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_str)
	items_written = outArgs[0].Int64()
	result = rune(ret.Uint32())
	return
}

// g_utf8_to_utf16
// container is nil
func Utf8ToUtf16(str string, len1 int64) (result uint16, items_read int64, items_written int64, err error) {
	iv, err := _I.Get(1291, "utf8_to_utf16", "")
	if err != nil {
		return
	}
	var outArgs [3]gi.Argument
	c_str := gi.CString(str)
	arg_str := gi.NewStringArgument(c_str)
	arg_len1 := gi.NewInt64Argument(len1)
	arg_items_read := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_items_written := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[2]))
	args := []gi.Argument{arg_str, arg_len1, arg_items_read, arg_items_written, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_str)
	err = gi.ToError(outArgs[2].Pointer())
	items_read = outArgs[0].Int64()
	items_written = outArgs[1].Int64()
	result = ret.Uint16()
	return
}

// g_utf8_validate
// container is nil
// arg 0 str lenArgIdx 1
func Utf8Validate(str gi.Uint8Array, max_len int64) (result bool, end string) {
	iv, err := _I.Get(1292, "utf8_validate", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_str := gi.NewPointerArgument(str.P)
	arg_max_len := gi.NewInt64Argument(max_len)
	arg_end := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_str, arg_max_len, arg_end}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	end = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// g_uuid_string_is_valid
// container is nil
func UuidStringIsValid(str string) (result bool) {
	iv, err := _I.Get(1293, "uuid_string_is_valid", "")
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
	result = ret.Bool()
	return
}

// g_uuid_string_random
// container is nil
func UuidStringRandom() (result string) {
	iv, err := _I.Get(1294, "uuid_string_random", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.String().Take()
	return
}

// g_variant_get_gtype
// container is nil
func VariantGetGtype() (result gi.GType) {
	iv, err := _I.Get(1295, "variant_get_gtype", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// g_variant_is_object_path
// container is nil
func VariantIsObjectPath(string string) (result bool) {
	iv, err := _I.Get(1296, "variant_is_object_path", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_variant_is_signature
// container is nil
func VariantIsSignature(string string) (result bool) {
	iv, err := _I.Get(1297, "variant_is_signature", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_string := gi.CString(string)
	arg_string := gi.NewStringArgument(c_string)
	args := []gi.Argument{arg_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_string)
	result = ret.Bool()
	return
}

// g_variant_parse
// container is nil
func VariantParse(type1 VariantType, text string, limit string, endptr string) (result Variant, err error) {
	iv, err := _I.Get(1298, "variant_parse", "")
	if err != nil {
		return
	}
	var outArgs [1]gi.Argument
	c_text := gi.CString(text)
	c_limit := gi.CString(limit)
	c_endptr := gi.CString(endptr)
	arg_type1 := gi.NewPointerArgument(type1.P)
	arg_text := gi.NewStringArgument(c_text)
	arg_limit := gi.NewStringArgument(c_limit)
	arg_endptr := gi.NewStringArgument(c_endptr)
	arg_err := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_type1, arg_text, arg_limit, arg_endptr, arg_err}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_text)
	gi.Free(c_limit)
	gi.Free(c_endptr)
	err = gi.ToError(outArgs[0].Pointer())
	result.P = ret.Pointer()
	return
}

// g_variant_parse_error_print_context
// container is nil
func VariantParseErrorPrintContext(error Error, source_str string) (result string) {
	iv, err := _I.Get(1299, "variant_parse_error_print_context", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_source_str := gi.CString(source_str)
	arg_error := gi.NewPointerArgument(error.P)
	arg_source_str := gi.NewStringArgument(c_source_str)
	args := []gi.Argument{arg_error, arg_source_str}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_source_str)
	result = ret.String().Take()
	return
}

// g_variant_parse_error_quark
// container is nil
func VariantParseErrorQuark() (result uint32) {
	iv, err := _I.Get(1300, "variant_parse_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_variant_parser_get_error_quark
// container is nil
func VariantParserGetErrorQuark() (result uint32) {
	iv, err := _I.Get(1301, "variant_parser_get_error_quark", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_variant_type_checked_
// container is nil
func VariantTypeChecked_(arg0 string) (result VariantType) {
	iv, err := _I.Get(1302, "variant_type_checked_", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_arg0 := gi.CString(arg0)
	arg_arg0 := gi.NewStringArgument(c_arg0)
	args := []gi.Argument{arg_arg0}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_arg0)
	result.P = ret.Pointer()
	return
}

// g_variant_type_string_get_depth_
// container is nil
func VariantTypeStringGetDepth_(type_string string) (result uint64) {
	iv, err := _I.Get(1303, "variant_type_string_get_depth_", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_type_string := gi.CString(type_string)
	arg_type_string := gi.NewStringArgument(c_type_string)
	args := []gi.Argument{arg_type_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_type_string)
	result = ret.Uint64()
	return
}

// g_variant_type_string_is_valid
// container is nil
func VariantTypeStringIsValid(type_string string) (result bool) {
	iv, err := _I.Get(1304, "variant_type_string_is_valid", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_type_string := gi.CString(type_string)
	arg_type_string := gi.NewStringArgument(c_type_string)
	args := []gi.Argument{arg_type_string}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_type_string)
	result = ret.Bool()
	return
}

// g_variant_type_string_scan
// container is nil
func VariantTypeStringScan(string string, limit string) (result bool, endptr string) {
	iv, err := _I.Get(1305, "variant_type_string_scan", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	c_string := gi.CString(string)
	c_limit := gi.CString(limit)
	arg_string := gi.NewStringArgument(c_string)
	arg_limit := gi.NewStringArgument(c_limit)
	arg_endptr := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_string, arg_limit, arg_endptr}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_string)
	gi.Free(c_limit)
	endptr = outArgs[0].String().Take()
	result = ret.Bool()
	return
}

// constants
const (
	ANALYZER_ANALYZING                         = 1
	ASCII_DTOSTR_BUF_SIZE                      = 39
	BIG_ENDIAN                                 = 4321
	CSET_A_2_Z                                 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CSET_DIGITS                                = "0123456789"
	CSET_a_2_z                                 = "abcdefghijklmnopqrstuvwxyz"
	DATALIST_FLAGS_MASK                        = 3
	DATE_BAD_DAY                               = 0
	DATE_BAD_JULIAN                            = 0
	DATE_BAD_YEAR                              = 0
	DIR_SEPARATOR                              = 47
	DIR_SEPARATOR_S                            = "/"
	E                                          = 2.718282
	GINT16_FORMAT                              = "hi"
	GINT16_MODIFIER                            = "h"
	GINT32_FORMAT                              = "i"
	GINT32_MODIFIER                            = ""
	GINT64_FORMAT                              = "li"
	GINT64_MODIFIER                            = "l"
	GINTPTR_FORMAT                             = "li"
	GINTPTR_MODIFIER                           = "l"
	GNUC_FUNCTION                              = ""
	GNUC_PRETTY_FUNCTION                       = ""
	GSIZE_FORMAT                               = "lu"
	GSIZE_MODIFIER                             = "l"
	GSSIZE_FORMAT                              = "li"
	GSSIZE_MODIFIER                            = "l"
	GUINT16_FORMAT                             = "hu"
	GUINT32_FORMAT                             = "u"
	GUINT64_FORMAT                             = "lu"
	GUINTPTR_FORMAT                            = "lu"
	HAVE_GINT64                                = 1
	HAVE_GNUC_VARARGS                          = 1
	HAVE_GNUC_VISIBILITY                       = 1
	HAVE_GROWING_STACK                         = 0
	HAVE_ISO_VARARGS                           = 1
	HOOK_FLAG_USER_SHIFT                       = 4
	IEEE754_DOUBLE_BIAS                        = 1023
	IEEE754_FLOAT_BIAS                         = 127
	KEY_FILE_DESKTOP_ACTION_GROUP_PREFIX       = "Desktop Action"
	KEY_FILE_DESKTOP_GROUP                     = "Desktop Entry"
	KEY_FILE_DESKTOP_KEY_ACTIONS               = "Actions"
	KEY_FILE_DESKTOP_KEY_CATEGORIES            = "Categories"
	KEY_FILE_DESKTOP_KEY_COMMENT               = "Comment"
	KEY_FILE_DESKTOP_KEY_DBUS_ACTIVATABLE      = "DBusActivatable"
	KEY_FILE_DESKTOP_KEY_EXEC                  = "Exec"
	KEY_FILE_DESKTOP_KEY_FULLNAME              = "X-GNOME-FullName"
	KEY_FILE_DESKTOP_KEY_GENERIC_NAME          = "GenericName"
	KEY_FILE_DESKTOP_KEY_GETTEXT_DOMAIN        = "X-GNOME-Gettext-Domain"
	KEY_FILE_DESKTOP_KEY_HIDDEN                = "Hidden"
	KEY_FILE_DESKTOP_KEY_ICON                  = "Icon"
	KEY_FILE_DESKTOP_KEY_KEYWORDS              = "Keywords"
	KEY_FILE_DESKTOP_KEY_MIME_TYPE             = "MimeType"
	KEY_FILE_DESKTOP_KEY_NAME                  = "Name"
	KEY_FILE_DESKTOP_KEY_NOT_SHOW_IN           = "NotShowIn"
	KEY_FILE_DESKTOP_KEY_NO_DISPLAY            = "NoDisplay"
	KEY_FILE_DESKTOP_KEY_ONLY_SHOW_IN          = "OnlyShowIn"
	KEY_FILE_DESKTOP_KEY_PATH                  = "Path"
	KEY_FILE_DESKTOP_KEY_STARTUP_NOTIFY        = "StartupNotify"
	KEY_FILE_DESKTOP_KEY_STARTUP_WM_CLASS      = "StartupWMClass"
	KEY_FILE_DESKTOP_KEY_TERMINAL              = "Terminal"
	KEY_FILE_DESKTOP_KEY_TRY_EXEC              = "TryExec"
	KEY_FILE_DESKTOP_KEY_TYPE                  = "Type"
	KEY_FILE_DESKTOP_KEY_URL                   = "URL"
	KEY_FILE_DESKTOP_KEY_VERSION               = "Version"
	KEY_FILE_DESKTOP_TYPE_APPLICATION          = "Application"
	KEY_FILE_DESKTOP_TYPE_DIRECTORY            = "Directory"
	KEY_FILE_DESKTOP_TYPE_LINK                 = "Link"
	LITTLE_ENDIAN                              = 1234
	LN10                                       = 2.302585
	LN2                                        = 0.693147
	LOG_2_BASE_10                              = 0.30103
	LOG_DOMAIN                                 = 0
	LOG_FATAL_MASK                             = 5
	LOG_LEVEL_USER_SHIFT                       = 8
	MAJOR_VERSION                              = 2
	MAXINT16                                   = 32767
	MAXINT32                                   = 2147483647
	MAXINT64                                   = 9223372036854775807
	MAXINT8                                    = 127
	MAXUINT16                                  = 65535
	MAXUINT32                                  = 4294967295
	MAXUINT64                                  = 18446744073709551615
	MAXUINT8                                   = 255
	MICRO_VERSION                              = 3
	MININT16                                   = -32768
	MININT32                                   = -2147483648
	MININT64                                   = -9223372036854775808
	MININT8                                    = -128
	MINOR_VERSION                              = 58
	MODULE_SUFFIX                              = "so"
	OPTION_REMAINING                           = ""
	PDP_ENDIAN                                 = 3412
	PI                                         = 3.141593
	PID_FORMAT                                 = "i"
	PI_2                                       = 1.570796
	PI_4                                       = 0.785398
	POLLFD_FORMAT                              = "%d"
	PRIORITY_DEFAULT                           = 0
	PRIORITY_DEFAULT_IDLE                      = 200
	PRIORITY_HIGH                              = -100
	PRIORITY_HIGH_IDLE                         = 100
	PRIORITY_LOW                               = 300
	SEARCHPATH_SEPARATOR                       = 58
	SEARCHPATH_SEPARATOR_S                     = ":"
	SIZEOF_LONG                                = 8
	SIZEOF_SIZE_T                              = 8
	SIZEOF_SSIZE_T                             = 8
	SIZEOF_VOID_P                              = 8
	SOURCE_CONTINUE                            = true
	SOURCE_REMOVE                              = false
	SQRT2                                      = 1.414214
	STR_DELIMITERS                             = "_-|> <."
	SYSDEF_AF_INET                             = 2
	SYSDEF_AF_INET6                            = 10
	SYSDEF_AF_UNIX                             = 1
	SYSDEF_MSG_DONTROUTE                       = 4
	SYSDEF_MSG_OOB                             = 1
	SYSDEF_MSG_PEEK                            = 2
	TIME_SPAN_DAY                              = 86400000000
	TIME_SPAN_HOUR                             = 3600000000
	TIME_SPAN_MILLISECOND                      = 1000
	TIME_SPAN_MINUTE                           = 60000000
	TIME_SPAN_SECOND                           = 1000000
	UNICHAR_MAX_DECOMPOSITION_LENGTH           = 18
	URI_RESERVED_CHARS_GENERIC_DELIMITERS      = ":/?#[]@"
	URI_RESERVED_CHARS_SUBCOMPONENT_DELIMITERS = "!$&'()*+,;="
	USEC_PER_SEC                               = 1000000
	VA_COPY_AS_ARRAY                           = 1
	VERSION_MIN_REQUIRED                       = 2
	WIN32_MSG_HANDLE                           = 19981206
)
