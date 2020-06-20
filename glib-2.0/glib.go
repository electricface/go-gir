package glib

/*
#cgo pkg-config: glib-2.0
#include <glib.h>


extern void myGFunc(gpointer data, gpointer user_data);

static void* getPointer_myGFunc() {
	return (void*)(myGFunc);
}
*/
import "C"

import (
	gi "github.com/electricface/go-gir3/gi-lite"
	"unsafe"
)

type GFuncStruct struct {
	Data unsafe.Pointer
}

//export myGFunc
func myGFunc(data C.gpointer, user_data C.gpointer) {
	call := gi.GetFunc(uint(uintptr(user_data)))
	args := &GFuncStruct{
		Data: unsafe.Pointer(data),
	}
	call(args)
}

//func getPointer_myGFunc() unsafe.Pointer {
//	return C.getPointer_myGFunc()
//}

func (v List) p() *C.GList {
	return (*C.GList)(v.P)
}

func NewList() List {
	ret := C.g_list_alloc()
	return wrapList(ret)
}

func (v List) ForEach(fn func(item unsafe.Pointer)) {
	for l := v.p(); l != nil; l = l.next {
		fn(unsafe.Pointer(l.data))
	}
}

func (v List) ForEachC(fn func(args interface{})) {
	fnId := gi.RegisterFunc(fn)
	C.g_list_foreach(v.p(), C.GFunc(C.getPointer_myGFunc()), C.gpointer(fnId))
	// todo unregister func fnId
}

func (v *List) FullFree(fn func(item unsafe.Pointer)) {
	v.ForEach(fn)
	v.Free()
	v.P = nil
}

func wrapList(p *C.GList) List {
	return List{P: unsafe.Pointer(p)}
}

func (v List) Next() List {
	native := v.p()
	return wrapList(native.next)
}

func (v List) Previous() List {
	native := v.p()
	return wrapList(native.prev)
}

func (v *List) Free() {
	C.g_list_free(v.p())
	v.P = nil
}

func (v List) Data() unsafe.Pointer {
	native := v.p()
	return unsafe.Pointer(native.data)
}

func (v List) Length() int {
	return int(C.g_list_length(v.p()))
}

func (v List) NthData(n uint) unsafe.Pointer {
	data := C.g_list_nth_data(v.p(), C.guint(n))
	return unsafe.Pointer(data)
}

func (v List) Nth(n uint) List {
	list := C.g_list_nth(v.p(), C.guint(n))
	return wrapList(list)
}

func (v List) Append(data unsafe.Pointer) List {
	list := C.g_list_append(v.p(), C.gpointer(data))
	return wrapList(list)
}

func (v List) Prepend(data unsafe.Pointer) List {
	list := C.g_list_prepend(v.p(), C.gpointer(data))
	return wrapList(list)
}

func (v List) Insert(data unsafe.Pointer, position int) List {
	list := C.g_list_insert(v.p(), C.gpointer(data), C.gint(position))
	return wrapList(list)
}

func (v List) InsertBefore(sibling List, data unsafe.Pointer) List {
	list := C.g_list_insert_before(v.p(), sibling.p(), C.gpointer(data))
	return wrapList(list)
}

func (v SList) p() *C.GSList {
	return (*C.GSList)(v.P)
}

// g_slist_alloc
//
// [ result ] trans: everything
func NewSList() SList {
	ret := C.g_slist_alloc()
	return wrapSList(ret)
}

func wrapSList(p *C.GSList) SList {
	return SList{P: unsafe.Pointer(p)}
}

func (v SList) Append(data unsafe.Pointer) SList {
	list := C.g_slist_append(v.p(), C.gpointer(data))
	return wrapSList(list)
}

func (v SList) Prepend(data unsafe.Pointer) SList {
	list := C.g_slist_prepend(v.p(), C.gpointer(data))
	return wrapSList(list)
}

func (v SList) Insert(data unsafe.Pointer, position int) SList {
	list := C.g_slist_insert(v.p(), C.gpointer(data), C.gint(position))
	return wrapSList(list)
}

func (v SList) InsertBefore(sibling SList, data unsafe.Pointer) SList {
	list := C.g_slist_insert_before(v.p(), sibling.p(), C.gpointer(data))
	return wrapSList(list)
}

func (v SList) Remove(data unsafe.Pointer) SList {
	list := C.g_slist_remove(v.p(), C.gconstpointer(data))
	return wrapSList(list)
}

func (v SList) RemoveLink(link_ SList) SList {
	list := C.g_slist_remove_link(v.p(), link_.p())
	return wrapSList(list)
}

func (v SList) DeleteLink(link_ SList) SList {
	list := C.g_slist_delete_link(v.p(), link_.p())
	return wrapSList(list)
}

func (v SList) RemoveAll(data unsafe.Pointer) SList {
	list := C.g_slist_remove_all(v.p(), C.gconstpointer(data))
	return wrapSList(list)
}

func (v SList) Free() {
	C.g_slist_free(v.p())
}

// TODO g_slist_free_full

func (v SList) Free1() {
	C.g_slist_free_1(v.p())
}

func (v SList) Length() uint {
	return uint(C.g_slist_length(v.p()))
}

func (v SList) Copy() SList {
	list := C.g_slist_copy(v.p())
	return wrapSList(list)
}

// TODO g_slist_deep_copy

func (v SList) Reverse() SList {
	list := C.g_slist_reverse(v.p())
	return wrapSList(list)
}

// TODO: g_slist_insert_sorted_with_data

// TODO: g_slist_sort

// TODO: g_slist_sort_with_data

func (v SList) Concat(list2 SList) SList {
	list := C.g_slist_concat(v.p(), list2.p())
	return wrapSList(list)
}

// TODO: g_slist_foreach
func (v SList) ForEach(fn func(item unsafe.Pointer)) {
	for l := v.p(); l != nil; l = l.next {
		fn(unsafe.Pointer(l.data))
	}
}

func (v SList) Last() SList {
	list := C.g_slist_last(v.p())
	return wrapSList(list)
}

func (v SList) Next() SList {
	native := v.p()
	return wrapSList(native.next)
}

func (v SList) Nth(n uint) SList {
	list := C.g_slist_nth(v.p(), C.guint(n))
	return wrapSList(list)
}

func (v SList) NthData(n uint) unsafe.Pointer {
	ptr := C.g_slist_nth_data(v.p(), C.guint(n))
	return unsafe.Pointer(ptr)
}

func (v SList) Find(data unsafe.Pointer) SList {
	list := C.g_slist_find(v.p(), C.gconstpointer(data))
	return wrapSList(list)
}

// TODO: g_slist_find_custom

func (v SList) Position(llist SList) int {
	return int(C.g_slist_position(v.p(), llist.p()))
}

func (v SList) Index(data unsafe.Pointer) int {
	return int(C.g_slist_index(v.p(), C.gconstpointer(data)))
}
