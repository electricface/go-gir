package gobject

/*
#cgo pkg-config: glib-2.0 gobject-2.0
#include <glib.h>
#include <glib-object.h>
#include <stdlib.h>
extern void	goMarshal(GClosure *, GValue *, guint, GValue *, gpointer, GValue *);

static GClosure *
_g_closure_new() {
	GClosure	*closure;

	closure = g_closure_new_simple(sizeof(GClosure), NULL);
	g_closure_set_marshal(closure, (GClosureMarshal)(goMarshal));
	return closure;
}

extern void	removeClosure(gpointer, GClosure *);

static void
_g_closure_add_finalize_notifier(GClosure *closure) {
	g_closure_add_finalize_notifier(closure, NULL, removeClosure);
}

*/
import "C"
import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime/debug"
	"sync"
	"unsafe"

	"github.com/electricface/go-gir/glib-2.0"
	"github.com/electricface/go-gir/util"
)

type closureContext struct {
	rf reflect.Value
}

var (
	errNilPtr = errors.New("cgo returned unexpected nil pointer")

	closures = struct {
		sync.RWMutex
		m map[*C.GClosure]closureContext
	}{
		m: make(map[*C.GClosure]closureContext),
	}

	signals = make(map[SignalHandle]Closure)
)

// removeClosure removes a closure from the internal closures map.  This is
// needed to prevent a leak where Go code can access the closure context
// (along with rf and userdata) even after an object has been destroyed and
// the GClosure is invalidated and will never run.
//
//export removeClosure
func removeClosure(_ C.gpointer, closure *C.GClosure) {
	closures.Lock()
	delete(closures.m, closure)
	closures.Unlock()
}

//export goMarshal
func goMarshal(closure *C.GClosure, retValue *C.GValue,
	nParams C.guint, params *C.GValue,
	invocationHint C.gpointer, marshalData *C.GValue) {

	// Get the context associated with this callback closure.
	closures.RLock()
	cc := closures.m[closure]
	closures.RUnlock()

	var args []interface{}
	nGLibParams := int(nParams)
	if nGLibParams > 0 {
		args = make([]interface{}, nGLibParams)
		gValues := gValueSlice(params, nGLibParams)
		for i := 0; i < nGLibParams; i++ {
			v := Value{unsafe.Pointer(&gValues[i])}
			val, _ := v.Get()
			args[i] = val
		}
	}

	defer func() {
		err := recover()
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "func panic with error:", err)
			debug.PrintStack()
		}
	}()

	switch fn := cc.rf.Interface().(type) {
	case func():
		// 无参数，无返回值
		fn()
	case func() interface{}:
		// 无参数，有返回值
		ret := fn()
		gRetValue := Value{unsafe.Pointer(retValue)}
		err := gRetValue.Set(ret)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cannot save callback return value: %v", err)
		}

	case func(args []interface{}):
		// 有参数，无返回值
		fn(args)

	case func(args []interface{}) interface{}:
		// 有参数，有返回值
		ret := fn(args)
		gRetValue := Value{unsafe.Pointer(retValue)}
		err := gRetValue.Set(ret)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cannot save callback return value: %v", err)
		}
	default:
		_, _ = fmt.Fprintf(os.Stderr, "invalid func type")
	}
}

// gValueSlice converts a C array of GValues to a Go slice.
func gValueSlice(values *C.GValue, nValues int) (slice []C.GValue) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	header.Cap = nValues
	header.Len = nValues
	header.Data = uintptr(unsafe.Pointer(values))
	return
}

type Type uint

type CanGetTypeAndGValueGetter interface {
	GetType() Type
	GetGValueGetter() GValueGetter
}

var cgt CanGetTypeAndGValueGetter
var cgtIfc = reflect.TypeOf(&cgt).Elem()

// ClosureNew creates a new GClosure and adds its callback function
// to the internally-maintained map. It's exported for visibility to other
// gotk3 packages and shouldn't be used in application code.
func ClosureNew(f interface{}) Closure {
	// Create a reflect.Value from f.  This is called when the
	// returned GClosure runs.
	rf := reflect.ValueOf(f)

	// Create closure context which points to the reflected func.
	cc := closureContext{rf: rf}

	// Closures can only be created from funcs.
	if rf.Type().Kind() != reflect.Func {
		panic("f is not a func")
	}

	fType := rf.Type()
	for i := 0; i < fType.NumIn(); i++ {
		argType := fType.In(i)
		//fmt.Printf("i: %d, argType: %v\n", i, argType)
		if argType.Implements(cgtIfc) {
			argEmptyValue := reflect.New(argType)
			cctm := argEmptyValue.Interface().(CanGetTypeAndGValueGetter)
			registerGValueGetter(cctm.GetType(), cctm.GetGValueGetter())
		}
	}

	c := C._g_closure_new()
	C._g_closure_add_finalize_notifier(c)

	// Associate the GClosure with rf.  rf will be looked up in this
	// map by the closure when the closure runs.
	closures.Lock()
	closures.m[c] = cc
	closures.Unlock()

	return Closure{unsafe.Pointer(c)}
}

type SignalHandle uint

func (v Closure) native() *C.GClosure {
	return (*C.GClosure)(v.P)
}

func (v Object) connectClosure(after bool, detailedSignal string, f interface{}) SignalHandle {
	cstr := C.CString(detailedSignal)
	defer C.free(unsafe.Pointer(cstr))

	closure := ClosureNew(f)

	c := C.g_signal_connect_closure(C.gpointer(v.P),
		(*C.gchar)(cstr), closure.native(), C.gboolean(util.Bool2Int(after)))
	handle := SignalHandle(c)

	// Map the signal handle to the closure.
	signals[handle] = closure

	return handle
}

//func SourceSetClosure(src glib.Source, closure Closure) {
//	C.g_source_set_closure((*C.GSource)(src.Ptr), closure.native())
//}

type SourceFunc func() bool

func IdleAdd(f SourceFunc) uint {
	src := glib.IdleSourceNew()
	id := setupSourceFunc(src, f)
	src.Unref()
	return id
}

func TimeoutAdd(interval uint, f SourceFunc) uint {
	src := glib.TimeoutSourceNew(uint32(interval))
	id := setupSourceFunc(src, f)
	src.Unref()
	return id
}

func TimeoutAddSeconds(interval uint, f SourceFunc) uint {
	src := glib.TimeoutSourceNewSeconds(uint32(interval))
	id := setupSourceFunc(src, f)
	src.Unref()
	return id
}

func setupSourceFunc(src glib.Source, f SourceFunc) uint {
	closure := ClosureNew(f)
	SourceSetClosure(src, closure)
	return uint(src.Attach(glib.MainContextDefault()))
}
