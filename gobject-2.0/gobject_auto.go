package gobject

/*
#cgo pkg-config: gobject-2.0
#include <glib-object.h>
extern void myGObjectBaseFinalizeFunc(GTypeClass* g_class);
static void* getPointer_myGObjectBaseFinalizeFunc() {
return (void*)(myGObjectBaseFinalizeFunc);
}
extern void myGObjectBaseInitFunc(GTypeClass* g_class);
static void* getPointer_myGObjectBaseInitFunc() {
return (void*)(myGObjectBaseInitFunc);
}
extern void myGObjectBindingTransformFunc(GBinding* binding, GValue* from_value, GValue* to_value, gpointer user_data);
static void* getPointer_myGObjectBindingTransformFunc() {
return (void*)(myGObjectBindingTransformFunc);
}
extern void myGObjectBoxedCopyFunc(gpointer boxed);
static void* getPointer_myGObjectBoxedCopyFunc() {
return (void*)(myGObjectBoxedCopyFunc);
}
extern void myGObjectBoxedFreeFunc(gpointer boxed);
static void* getPointer_myGObjectBoxedFreeFunc() {
return (void*)(myGObjectBoxedFreeFunc);
}
extern void myGObjectCallback();
static void* getPointer_myGObjectCallback() {
return (void*)(myGObjectCallback);
}
extern void myGObjectClassFinalizeFunc(GTypeClass* g_class, gpointer class_data);
static void* getPointer_myGObjectClassFinalizeFunc() {
return (void*)(myGObjectClassFinalizeFunc);
}
extern void myGObjectClassInitFunc(GTypeClass* g_class, gpointer class_data);
static void* getPointer_myGObjectClassInitFunc() {
return (void*)(myGObjectClassInitFunc);
}
extern void myGObjectClosureMarshal(GClosure* closure, GValue* return_value, guint32 n_param_values, gpointer param_values, gpointer invocation_hint, gpointer marshal_data);
static void* getPointer_myGObjectClosureMarshal() {
return (void*)(myGObjectClosureMarshal);
}
extern void myGObjectClosureNotify(gpointer data, GClosure* closure);
static void* getPointer_myGObjectClosureNotify() {
return (void*)(myGObjectClosureNotify);
}
extern void myGObjectInstanceInitFunc(GTypeInstance* instance, GTypeClass* g_class);
static void* getPointer_myGObjectInstanceInitFunc() {
return (void*)(myGObjectInstanceInitFunc);
}
extern void myGObjectInterfaceFinalizeFunc(GTypeInterface* g_iface, gpointer iface_data);
static void* getPointer_myGObjectInterfaceFinalizeFunc() {
return (void*)(myGObjectInterfaceFinalizeFunc);
}
extern void myGObjectInterfaceInitFunc(GTypeInterface* g_iface, gpointer iface_data);
static void* getPointer_myGObjectInterfaceInitFunc() {
return (void*)(myGObjectInterfaceInitFunc);
}
extern void myGObjectObjectFinalizeFunc(GObject* object);
static void* getPointer_myGObjectObjectFinalizeFunc() {
return (void*)(myGObjectObjectFinalizeFunc);
}
extern void myGObjectObjectGetPropertyFunc(GObject* object, guint32 property_id, GValue* value, GParamSpec* pspec);
static void* getPointer_myGObjectObjectGetPropertyFunc() {
return (void*)(myGObjectObjectGetPropertyFunc);
}
extern void myGObjectObjectSetPropertyFunc(GObject* object, guint32 property_id, GValue* value, GParamSpec* pspec);
static void* getPointer_myGObjectObjectSetPropertyFunc() {
return (void*)(myGObjectObjectSetPropertyFunc);
}
extern void myGObjectSignalAccumulator(GSignalInvocationHint* ihint, GValue* return_accu, GValue* handler_return, gpointer data);
static void* getPointer_myGObjectSignalAccumulator() {
return (void*)(myGObjectSignalAccumulator);
}
extern void myGObjectSignalEmissionHook(GSignalInvocationHint* ihint, guint32 n_param_values, gpointer param_values, gpointer data);
static void* getPointer_myGObjectSignalEmissionHook() {
return (void*)(myGObjectSignalEmissionHook);
}
extern void myGObjectToggleNotify(gpointer data, GObject* object, gboolean is_last_ref);
static void* getPointer_myGObjectToggleNotify() {
return (void*)(myGObjectToggleNotify);
}
extern void myGObjectTypeClassCacheFunc(gpointer cache_data, GTypeClass* g_class);
static void* getPointer_myGObjectTypeClassCacheFunc() {
return (void*)(myGObjectTypeClassCacheFunc);
}
extern void myGObjectTypeInterfaceCheckFunc(gpointer check_data, GTypeInterface* g_iface);
static void* getPointer_myGObjectTypeInterfaceCheckFunc() {
return (void*)(myGObjectTypeInterfaceCheckFunc);
}
extern void myGObjectTypePluginCompleteInterfaceInfo(GTypePlugin* plugin, gpointer instance_type, gpointer interface_type, GInterfaceInfo* info);
static void* getPointer_myGObjectTypePluginCompleteInterfaceInfo() {
return (void*)(myGObjectTypePluginCompleteInterfaceInfo);
}
extern void myGObjectTypePluginCompleteTypeInfo(GTypePlugin* plugin, gpointer g_type, GTypeInfo* info, GTypeValueTable* value_table);
static void* getPointer_myGObjectTypePluginCompleteTypeInfo() {
return (void*)(myGObjectTypePluginCompleteTypeInfo);
}
extern void myGObjectTypePluginUnuse(GTypePlugin* plugin);
static void* getPointer_myGObjectTypePluginUnuse() {
return (void*)(myGObjectTypePluginUnuse);
}
extern void myGObjectTypePluginUse(GTypePlugin* plugin);
static void* getPointer_myGObjectTypePluginUse() {
return (void*)(myGObjectTypePluginUse);
}
extern void myGObjectValueTransform(GValue* src_value, GValue* dest_value);
static void* getPointer_myGObjectValueTransform() {
return (void*)(myGObjectValueTransform);
}
extern void myGObjectWeakNotify(gpointer data, GObject* where_the_object_was);
static void* getPointer_myGObjectWeakNotify() {
return (void*)(myGObjectWeakNotify);
}
*/
import "C"
import "github.com/electricface/go-gir/glib-2.0"
import "github.com/electricface/go-gir3/gi"
import "log"
import "unsafe"

var _I = gi.NewInvokerCache("GObject")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GObject", "2.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

type BaseFinalizeFuncStruct struct {
	F_g_class TypeClass
}

func GetPointer_myBaseFinalizeFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectBaseFinalizeFunc())
}

//export myGObjectBaseFinalizeFunc
func myGObjectBaseFinalizeFunc(g_class *C.GTypeClass) {
	// TODO: not found user_data
}

type BaseInitFuncStruct struct {
	F_g_class TypeClass
}

func GetPointer_myBaseInitFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectBaseInitFunc())
}

//export myGObjectBaseInitFunc
func myGObjectBaseInitFunc(g_class *C.GTypeClass) {
	// TODO: not found user_data
}

// Object Binding
type Binding struct {
	Object
}

func WrapBinding(p unsafe.Pointer) (r Binding) { r.P = p; return }

type IBinding interface{ P_Binding() unsafe.Pointer }

func (v Binding) P_Binding() unsafe.Pointer { return v.P }
func BindingGetType() gi.GType {
	ret := _I.GetGType(0, "Binding")
	return ret
}

// g_binding_get_flags
// container is not nil, container is Binding
// is method
func (v Binding) GetFlags() (result BindingFlags) {
	iv, err := _I.Get(0, "Binding", "get_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = BindingFlags(ret.Int())
	return
}

// g_binding_get_source
// container is not nil, container is Binding
// is method
func (v Binding) GetSource() (result Object) {
	iv, err := _I.Get(1, "Binding", "get_source")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_binding_get_source_property
// container is not nil, container is Binding
// is method
func (v Binding) GetSourceProperty() (result string) {
	iv, err := _I.Get(2, "Binding", "get_source_property")
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

// g_binding_get_target
// container is not nil, container is Binding
// is method
func (v Binding) GetTarget() (result Object) {
	iv, err := _I.Get(3, "Binding", "get_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_binding_get_target_property
// container is not nil, container is Binding
// is method
func (v Binding) GetTargetProperty() (result string) {
	iv, err := _I.Get(4, "Binding", "get_target_property")
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

// g_binding_unbind
// container is not nil, container is Binding
// is method
func (v Binding) Unbind() {
	iv, err := _I.Get(5, "Binding", "unbind")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// Flags BindingFlags
type BindingFlags int

const (
	BindingFlagsDefault       BindingFlags = 0
	BindingFlagsBidirectional BindingFlags = 1
	BindingFlagsSyncCreate    BindingFlags = 2
	BindingFlagsInvertBoolean BindingFlags = 4
)

func BindingFlagsGetType() gi.GType {
	ret := _I.GetGType(1, "BindingFlags")
	return ret
}

type BindingTransformFuncStruct struct {
	F_binding    Binding
	F_from_value Value
	F_to_value   Value
}

func GetPointer_myBindingTransformFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectBindingTransformFunc())
}

//export myGObjectBindingTransformFunc
func myGObjectBindingTransformFunc(binding *C.GBinding, from_value *C.GValue, to_value *C.GValue, user_data C.gpointer) {
	fn := gi.GetFunc(uint(uintptr(user_data)))
	args := &BindingTransformFuncStruct{
		F_binding:    WrapBinding(unsafe.Pointer(binding)),
		F_from_value: Value{P: unsafe.Pointer(from_value)},
		F_to_value:   Value{P: unsafe.Pointer(to_value)},
	}
	fn(args)
}

type BoxedCopyFuncStruct struct {
	F_boxed unsafe.Pointer
}

func GetPointer_myBoxedCopyFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectBoxedCopyFunc())
}

//export myGObjectBoxedCopyFunc
func myGObjectBoxedCopyFunc(boxed C.gpointer) {
	// TODO: not found user_data
}

type BoxedFreeFuncStruct struct {
	F_boxed unsafe.Pointer
}

func GetPointer_myBoxedFreeFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectBoxedFreeFunc())
}

//export myGObjectBoxedFreeFunc
func myGObjectBoxedFreeFunc(boxed C.gpointer) {
	// TODO: not found user_data
}

// Struct CClosure
type CClosure struct {
	P unsafe.Pointer
}

const SizeOfStructCClosure = 72

func CClosureGetType() gi.GType {
	ret := _I.GetGType(2, "CClosure")
	return ret
}

// g_cclosure_marshal_BOOLEAN__BOXED_BOXED
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalBoolean_BoxedBoxed1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(6, "CClosure", "marshal_BOOLEAN__BOXED_BOXED")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_BOOLEAN__FLAGS
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalBoolean_Flags1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(7, "CClosure", "marshal_BOOLEAN__FLAGS")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_STRING__OBJECT_POINTER
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalString_ObjectPointer1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(8, "CClosure", "marshal_STRING__OBJECT_POINTER")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__BOOLEAN
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Boolean1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(9, "CClosure", "marshal_VOID__BOOLEAN")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__BOXED
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Boxed1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(10, "CClosure", "marshal_VOID__BOXED")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__CHAR
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Char1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(11, "CClosure", "marshal_VOID__CHAR")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__DOUBLE
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Double1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(12, "CClosure", "marshal_VOID__DOUBLE")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__ENUM
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Enum1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(13, "CClosure", "marshal_VOID__ENUM")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__FLAGS
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Flags1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(14, "CClosure", "marshal_VOID__FLAGS")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__FLOAT
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Float1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(15, "CClosure", "marshal_VOID__FLOAT")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__INT
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Int1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(16, "CClosure", "marshal_VOID__INT")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__LONG
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Long1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(17, "CClosure", "marshal_VOID__LONG")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__OBJECT
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Object1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(18, "CClosure", "marshal_VOID__OBJECT")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__PARAM
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Param1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(19, "CClosure", "marshal_VOID__PARAM")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__POINTER
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Pointer1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(20, "CClosure", "marshal_VOID__POINTER")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__STRING
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_String1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(21, "CClosure", "marshal_VOID__STRING")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__UCHAR
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Uchar1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(22, "CClosure", "marshal_VOID__UCHAR")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__UINT
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Uint1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(23, "CClosure", "marshal_VOID__UINT")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__UINT_POINTER
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_UintPointer1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(24, "CClosure", "marshal_VOID__UINT_POINTER")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__ULONG
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Ulong1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(25, "CClosure", "marshal_VOID__ULONG")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__VARIANT
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Variant1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(26, "CClosure", "marshal_VOID__VARIANT")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__VOID
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalVoid_Void1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(27, "CClosure", "marshal_VOID__VOID")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_generic
// container is not nil, container is CClosure
// is method
// arg0Type tag: interface, isPtr: true
func CClosureMarshalGeneric1(closure Closure, return_gvalue Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(28, "CClosure", "marshal_generic")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_gvalue := gi.NewPointerArgument(return_gvalue.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_gvalue, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

type CallbackStruct struct {
}

func GetPointer_myCallback() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectCallback())
}

//export myGObjectCallback
func myGObjectCallback() {
	// TODO: not found user_data
}

type ClassFinalizeFuncStruct struct {
	F_g_class    TypeClass
	F_class_data unsafe.Pointer
}

func GetPointer_myClassFinalizeFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectClassFinalizeFunc())
}

//export myGObjectClassFinalizeFunc
func myGObjectClassFinalizeFunc(g_class *C.GTypeClass, class_data C.gpointer) {
	// TODO: not found user_data
}

type ClassInitFuncStruct struct {
	F_g_class    TypeClass
	F_class_data unsafe.Pointer
}

func GetPointer_myClassInitFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectClassInitFunc())
}

//export myGObjectClassInitFunc
func myGObjectClassInitFunc(g_class *C.GTypeClass, class_data C.gpointer) {
	// TODO: not found user_data
}

// Struct Closure
type Closure struct {
	P unsafe.Pointer
}

const SizeOfStructClosure = 64

func ClosureGetType() gi.GType {
	ret := _I.GetGType(3, "Closure")
	return ret
}

// g_closure_new_object
// container is not nil, container is Closure
// is constructor
func NewClosureObject(sizeof_closure uint32, object IObject) (result Closure) {
	iv, err := _I.Get(29, "Closure", "new_object")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if object != nil {
		tmp = object.P_Object()
	}
	arg_sizeof_closure := gi.NewUint32Argument(sizeof_closure)
	arg_object := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_sizeof_closure, arg_object}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_closure_new_simple
// container is not nil, container is Closure
// is constructor
func NewClosureSimple(sizeof_closure uint32, data unsafe.Pointer) (result Closure) {
	iv, err := _I.Get(30, "Closure", "new_simple")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_sizeof_closure := gi.NewUint32Argument(sizeof_closure)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_sizeof_closure, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_closure_invalidate
// container is not nil, container is Closure
// is method
func (v Closure) Invalidate() {
	iv, err := _I.Get(31, "Closure", "invalidate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_closure_invoke
// container is not nil, container is Closure
// is method
// arg 2 param_values lenArgIdx 1
func (v Closure) Invoke(return_value Value, n_param_values uint32, param_values unsafe.Pointer, invocation_hint unsafe.Pointer) {
	iv, err := _I.Get(32, "Closure", "invoke")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	args := []gi.Argument{arg_v, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint}
	iv.Call(args, nil, nil)
}

// g_closure_ref
// container is not nil, container is Closure
// is method
func (v Closure) Ref() (result Closure) {
	iv, err := _I.Get(33, "Closure", "ref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_closure_sink
// container is not nil, container is Closure
// is method
func (v Closure) Sink() {
	iv, err := _I.Get(34, "Closure", "sink")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_closure_unref
// container is not nil, container is Closure
// is method
func (v Closure) Unref() {
	iv, err := _I.Get(35, "Closure", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

type ClosureMarshalStruct struct {
	F_closure         Closure
	F_return_value    Value
	F_n_param_values  uint32
	F_param_values    unsafe.Pointer
	F_invocation_hint unsafe.Pointer
	F_marshal_data    unsafe.Pointer
}

func GetPointer_myClosureMarshal() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectClosureMarshal())
}

//export myGObjectClosureMarshal
func myGObjectClosureMarshal(closure *C.GClosure, return_value *C.GValue, n_param_values C.guint32, param_values C.gpointer, invocation_hint C.gpointer, marshal_data C.gpointer) {
	// TODO: not found user_data
}

type ClosureNotifyStruct struct {
	F_data    unsafe.Pointer
	F_closure Closure
}

func GetPointer_myClosureNotify() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectClosureNotify())
}

//export myGObjectClosureNotify
func myGObjectClosureNotify(data C.gpointer, closure *C.GClosure) {
	// TODO: not found user_data
}

// Struct ClosureNotifyData
type ClosureNotifyData struct {
	P unsafe.Pointer
}

const SizeOfStructClosureNotifyData = 16

func ClosureNotifyDataGetType() gi.GType {
	ret := _I.GetGType(4, "ClosureNotifyData")
	return ret
}

// Flags ConnectFlags
type ConnectFlags int

const (
	ConnectFlagsAfter   ConnectFlags = 1
	ConnectFlagsSwapped ConnectFlags = 2
)

func ConnectFlagsGetType() gi.GType {
	ret := _I.GetGType(5, "ConnectFlags")
	return ret
}

// Struct EnumClass
type EnumClass struct {
	P unsafe.Pointer
}

const SizeOfStructEnumClass = 32

func EnumClassGetType() gi.GType {
	ret := _I.GetGType(6, "EnumClass")
	return ret
}

// Struct EnumValue
type EnumValue struct {
	P unsafe.Pointer
}

const SizeOfStructEnumValue = 24

func EnumValueGetType() gi.GType {
	ret := _I.GetGType(7, "EnumValue")
	return ret
}

// Struct FlagsClass
type FlagsClass struct {
	P unsafe.Pointer
}

const SizeOfStructFlagsClass = 24

func FlagsClassGetType() gi.GType {
	ret := _I.GetGType(8, "FlagsClass")
	return ret
}

// Struct FlagsValue
type FlagsValue struct {
	P unsafe.Pointer
}

const SizeOfStructFlagsValue = 24

func FlagsValueGetType() gi.GType {
	ret := _I.GetGType(9, "FlagsValue")
	return ret
}

// Object InitiallyUnowned
type InitiallyUnowned struct {
	Object
}

func WrapInitiallyUnowned(p unsafe.Pointer) (r InitiallyUnowned) { r.P = p; return }

type IInitiallyUnowned interface{ P_InitiallyUnowned() unsafe.Pointer }

func (v InitiallyUnowned) P_InitiallyUnowned() unsafe.Pointer { return v.P }
func InitiallyUnownedGetType() gi.GType {
	ret := _I.GetGType(10, "InitiallyUnowned")
	return ret
}

// ignore GType struct InitiallyUnownedClass
type InstanceInitFuncStruct struct {
	F_instance TypeInstance
	F_g_class  TypeClass
}

func GetPointer_myInstanceInitFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectInstanceInitFunc())
}

//export myGObjectInstanceInitFunc
func myGObjectInstanceInitFunc(instance *C.GTypeInstance, g_class *C.GTypeClass) {
	// TODO: not found user_data
}

type InterfaceFinalizeFuncStruct struct {
	F_g_iface    TypeInterface
	F_iface_data unsafe.Pointer
}

func GetPointer_myInterfaceFinalizeFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectInterfaceFinalizeFunc())
}

//export myGObjectInterfaceFinalizeFunc
func myGObjectInterfaceFinalizeFunc(g_iface *C.GTypeInterface, iface_data C.gpointer) {
	// TODO: not found user_data
}

// Struct InterfaceInfo
type InterfaceInfo struct {
	P unsafe.Pointer
}

const SizeOfStructInterfaceInfo = 24

func InterfaceInfoGetType() gi.GType {
	ret := _I.GetGType(11, "InterfaceInfo")
	return ret
}

type InterfaceInitFuncStruct struct {
	F_g_iface    TypeInterface
	F_iface_data unsafe.Pointer
}

func GetPointer_myInterfaceInitFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectInterfaceInitFunc())
}

//export myGObjectInterfaceInitFunc
func myGObjectInterfaceInitFunc(g_iface *C.GTypeInterface, iface_data C.gpointer) {
	// TODO: not found user_data
}

// Object Object
type Object struct {
	P unsafe.Pointer
}
type IObject interface{ P_Object() unsafe.Pointer }

func (v Object) P_Object() unsafe.Pointer { return v.P }
func ObjectGetType() gi.GType {
	ret := _I.GetGType(12, "Object")
	return ret
}

// g_object_newv
// container is not nil, container is Object
// is constructor
// arg 2 parameters lenArgIdx 1
func NewObjectv(object_type gi.GType, n_parameters uint32, parameters unsafe.Pointer) (result Object) {
	iv, err := _I.Get(36, "Object", "newv")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_object_type := gi.NewUintArgument(uint(object_type))
	arg_n_parameters := gi.NewUint32Argument(n_parameters)
	arg_parameters := gi.NewPointerArgument(parameters)
	args := []gi.Argument{arg_object_type, arg_n_parameters, arg_parameters}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_object_compat_control
// container is not nil, container is Object
// is method
// arg0Type tag: guint64, isPtr: false
func ObjectCompatControl1(what uint64, data unsafe.Pointer) (result uint64) {
	iv, err := _I.Get(37, "Object", "compat_control")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_what := gi.NewUint64Argument(what)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_what, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_object_interface_find_property
// container is not nil, container is Object
// is method
// arg0Type tag: interface, isPtr: true
func ObjectInterfaceFindProperty1(g_iface TypeInterface, property_name string) (result ParamSpec) {
	iv, err := _I.Get(38, "Object", "interface_find_property")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_property_name := gi.CString(property_name)
	arg_g_iface := gi.NewPointerArgument(g_iface.P)
	arg_property_name := gi.NewStringArgument(c_property_name)
	args := []gi.Argument{arg_g_iface, arg_property_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_property_name)
	result.P = ret.Pointer()
	return
}

// g_object_interface_install_property
// container is not nil, container is Object
// is method
// arg0Type tag: interface, isPtr: true
func ObjectInterfaceInstallProperty1(g_iface TypeInterface, pspec IParamSpec) {
	iv, err := _I.Get(39, "Object", "interface_install_property")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pspec != nil {
		tmp = pspec.P_ParamSpec()
	}
	arg_g_iface := gi.NewPointerArgument(g_iface.P)
	arg_pspec := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_g_iface, arg_pspec}
	iv.Call(args, nil, nil)
}

// g_object_interface_list_properties
// container is not nil, container is Object
// is method
// arg0Type tag: interface, isPtr: true
// ret lenArgIdx 1
func ObjectInterfaceListProperties1(g_iface TypeInterface) (result gi.PointerArray) {
	iv, err := _I.Get(40, "Object", "interface_list_properties")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_g_iface := gi.NewPointerArgument(g_iface.P)
	arg_n_properties_p := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_g_iface, arg_n_properties_p}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_properties_p uint32
	_ = n_properties_p
	n_properties_p = outArgs[0].Uint32()
	result = gi.PointerArray{P: ret.Pointer(), Len: int(n_properties_p)}
	return
}

// g_object_bind_property
// container is not nil, container is Object
// is method
func (v Object) BindProperty(source_property string, target IObject, target_property string, flags BindingFlags) (result Binding) {
	iv, err := _I.Get(41, "Object", "bind_property")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_source_property := gi.CString(source_property)
	var tmp unsafe.Pointer
	if target != nil {
		tmp = target.P_Object()
	}
	c_target_property := gi.CString(target_property)
	arg_v := gi.NewPointerArgument(v.P)
	arg_source_property := gi.NewStringArgument(c_source_property)
	arg_target := gi.NewPointerArgument(tmp)
	arg_target_property := gi.NewStringArgument(c_target_property)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_source_property, arg_target, arg_target_property, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_source_property)
	gi.Free(c_target_property)
	result.P = ret.Pointer()
	return
}

// g_object_bind_property_with_closures
// container is not nil, container is Object
// is method
func (v Object) BindPropertyFull(source_property string, target IObject, target_property string, flags BindingFlags, transform_to Closure, transform_from Closure) (result Binding) {
	iv, err := _I.Get(42, "Object", "bind_property_full")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_source_property := gi.CString(source_property)
	var tmp unsafe.Pointer
	if target != nil {
		tmp = target.P_Object()
	}
	c_target_property := gi.CString(target_property)
	arg_v := gi.NewPointerArgument(v.P)
	arg_source_property := gi.NewStringArgument(c_source_property)
	arg_target := gi.NewPointerArgument(tmp)
	arg_target_property := gi.NewStringArgument(c_target_property)
	arg_flags := gi.NewIntArgument(int(flags))
	arg_transform_to := gi.NewPointerArgument(transform_to.P)
	arg_transform_from := gi.NewPointerArgument(transform_from.P)
	args := []gi.Argument{arg_v, arg_source_property, arg_target, arg_target_property, arg_flags, arg_transform_to, arg_transform_from}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_source_property)
	gi.Free(c_target_property)
	result.P = ret.Pointer()
	return
}

// g_object_force_floating
// container is not nil, container is Object
// is method
func (v Object) ForceFloating() {
	iv, err := _I.Get(43, "Object", "force_floating")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_object_freeze_notify
// container is not nil, container is Object
// is method
func (v Object) FreezeNotify() {
	iv, err := _I.Get(44, "Object", "freeze_notify")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_object_get_data
// container is not nil, container is Object
// is method
func (v Object) GetData(key string) (result unsafe.Pointer) {
	iv, err := _I.Get(45, "Object", "get_data")
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
	result = ret.Pointer()
	return
}

// g_object_get_property
// container is not nil, container is Object
// is method
func (v Object) GetProperty(property_name string, value Value) {
	iv, err := _I.Get(46, "Object", "get_property")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_property_name := gi.CString(property_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_property_name := gi.NewStringArgument(c_property_name)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_property_name, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_property_name)
}

// g_object_get_qdata
// container is not nil, container is Object
// is method
func (v Object) GetQdata(quark uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(47, "Object", "get_qdata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_quark := gi.NewUint32Argument(quark)
	args := []gi.Argument{arg_v, arg_quark}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_object_getv
// container is not nil, container is Object
// is method
// arg 1 names lenArgIdx 0
// arg 2 values lenArgIdx 0
func (v Object) Getv(n_properties uint32, names gi.CStrArray, values unsafe.Pointer) {
	iv, err := _I.Get(48, "Object", "getv")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_n_properties := gi.NewUint32Argument(n_properties)
	arg_names := gi.NewPointerArgument(names.P)
	arg_values := gi.NewPointerArgument(values)
	args := []gi.Argument{arg_v, arg_n_properties, arg_names, arg_values}
	iv.Call(args, nil, nil)
}

// g_object_is_floating
// container is not nil, container is Object
// is method
func (v Object) IsFloating() (result bool) {
	iv, err := _I.Get(49, "Object", "is_floating")
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

// g_object_notify
// container is not nil, container is Object
// is method
func (v Object) Notify(property_name string) {
	iv, err := _I.Get(50, "Object", "notify")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_property_name := gi.CString(property_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_property_name := gi.NewStringArgument(c_property_name)
	args := []gi.Argument{arg_v, arg_property_name}
	iv.Call(args, nil, nil)
	gi.Free(c_property_name)
}

// g_object_notify_by_pspec
// container is not nil, container is Object
// is method
func (v Object) NotifyByPspec(pspec IParamSpec) {
	iv, err := _I.Get(51, "Object", "notify_by_pspec")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pspec != nil {
		tmp = pspec.P_ParamSpec()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pspec := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_pspec}
	iv.Call(args, nil, nil)
}

// g_object_ref
// container is not nil, container is Object
// is method
func (v Object) Ref() (result Object) {
	iv, err := _I.Get(52, "Object", "ref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_object_ref_sink
// container is not nil, container is Object
// is method
func (v Object) RefSink() (result Object) {
	iv, err := _I.Get(53, "Object", "ref_sink")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_object_run_dispose
// container is not nil, container is Object
// is method
func (v Object) RunDispose() {
	iv, err := _I.Get(54, "Object", "run_dispose")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_object_set_data
// container is not nil, container is Object
// is method
func (v Object) SetData(key string, data unsafe.Pointer) {
	iv, err := _I.Get(55, "Object", "set_data")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_key := gi.CString(key)
	arg_v := gi.NewPointerArgument(v.P)
	arg_key := gi.NewStringArgument(c_key)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_key, arg_data}
	iv.Call(args, nil, nil)
	gi.Free(c_key)
}

// g_object_set_property
// container is not nil, container is Object
// is method
func (v Object) SetProperty(property_name string, value Value) {
	iv, err := _I.Get(56, "Object", "set_property")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_property_name := gi.CString(property_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_property_name := gi.NewStringArgument(c_property_name)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_property_name, arg_value}
	iv.Call(args, nil, nil)
	gi.Free(c_property_name)
}

// g_object_steal_data
// container is not nil, container is Object
// is method
func (v Object) StealData(key string) (result unsafe.Pointer) {
	iv, err := _I.Get(57, "Object", "steal_data")
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
	result = ret.Pointer()
	return
}

// g_object_steal_qdata
// container is not nil, container is Object
// is method
func (v Object) StealQdata(quark uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(58, "Object", "steal_qdata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_quark := gi.NewUint32Argument(quark)
	args := []gi.Argument{arg_v, arg_quark}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_object_thaw_notify
// container is not nil, container is Object
// is method
func (v Object) ThawNotify() {
	iv, err := _I.Get(59, "Object", "thaw_notify")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_object_unref
// container is not nil, container is Object
// is method
func (v Object) Unref() {
	iv, err := _I.Get(60, "Object", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_object_watch_closure
// container is not nil, container is Object
// is method
func (v Object) WatchClosure(closure Closure) {
	iv, err := _I.Get(61, "Object", "watch_closure")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_closure := gi.NewPointerArgument(closure.P)
	args := []gi.Argument{arg_v, arg_closure}
	iv.Call(args, nil, nil)
}

// ignore GType struct ObjectClass
// Struct ObjectConstructParam
type ObjectConstructParam struct {
	P unsafe.Pointer
}

const SizeOfStructObjectConstructParam = 16

func ObjectConstructParamGetType() gi.GType {
	ret := _I.GetGType(13, "ObjectConstructParam")
	return ret
}

type ObjectFinalizeFuncStruct struct {
	F_object Object
}

func GetPointer_myObjectFinalizeFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectObjectFinalizeFunc())
}

//export myGObjectObjectFinalizeFunc
func myGObjectObjectFinalizeFunc(object *C.GObject) {
	// TODO: not found user_data
}

type ObjectGetPropertyFuncStruct struct {
	F_object      Object
	F_property_id uint32
	F_value       Value
	F_pspec       ParamSpec
}

func GetPointer_myObjectGetPropertyFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectObjectGetPropertyFunc())
}

//export myGObjectObjectGetPropertyFunc
func myGObjectObjectGetPropertyFunc(object *C.GObject, property_id C.guint32, value *C.GValue, pspec *C.GParamSpec) {
	// TODO: not found user_data
}

type ObjectSetPropertyFuncStruct struct {
	F_object      Object
	F_property_id uint32
	F_value       Value
	F_pspec       ParamSpec
}

func GetPointer_myObjectSetPropertyFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectObjectSetPropertyFunc())
}

//export myGObjectObjectSetPropertyFunc
func myGObjectObjectSetPropertyFunc(object *C.GObject, property_id C.guint32, value *C.GValue, pspec *C.GParamSpec) {
	// TODO: not found user_data
}

// Flags ParamFlags
type ParamFlags int

const (
	ParamFlagsReadable       ParamFlags = 1
	ParamFlagsWritable       ParamFlags = 2
	ParamFlagsReadwrite      ParamFlags = 3
	ParamFlagsConstruct      ParamFlags = 4
	ParamFlagsConstructOnly  ParamFlags = 8
	ParamFlagsLaxValidation  ParamFlags = 16
	ParamFlagsStaticName     ParamFlags = 32
	ParamFlagsPrivate        ParamFlags = 32
	ParamFlagsStaticNick     ParamFlags = 64
	ParamFlagsStaticBlurb    ParamFlags = 128
	ParamFlagsExplicitNotify ParamFlags = 1073741824
	ParamFlagsDeprecated     ParamFlags = 2147483648
)

func ParamFlagsGetType() gi.GType {
	ret := _I.GetGType(14, "ParamFlags")
	return ret
}

// Object ParamSpec
type ParamSpec struct {
	P unsafe.Pointer
}
type IParamSpec interface{ P_ParamSpec() unsafe.Pointer }

func (v ParamSpec) P_ParamSpec() unsafe.Pointer { return v.P }
func ParamSpecGetType() gi.GType {
	ret := _I.GetGType(15, "ParamSpec")
	return ret
}

// g_param_spec_get_blurb
// container is not nil, container is ParamSpec
// is method
func (v ParamSpec) GetBlurb() (result string) {
	iv, err := _I.Get(62, "ParamSpec", "get_blurb")
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

// g_param_spec_get_default_value
// container is not nil, container is ParamSpec
// is method
func (v ParamSpec) GetDefaultValue() (result Value) {
	iv, err := _I.Get(63, "ParamSpec", "get_default_value")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_param_spec_get_name
// container is not nil, container is ParamSpec
// is method
func (v ParamSpec) GetName() (result string) {
	iv, err := _I.Get(64, "ParamSpec", "get_name")
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

// g_param_spec_get_name_quark
// container is not nil, container is ParamSpec
// is method
func (v ParamSpec) GetNameQuark() (result uint32) {
	iv, err := _I.Get(65, "ParamSpec", "get_name_quark")
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

// g_param_spec_get_nick
// container is not nil, container is ParamSpec
// is method
func (v ParamSpec) GetNick() (result string) {
	iv, err := _I.Get(66, "ParamSpec", "get_nick")
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

// g_param_spec_get_qdata
// container is not nil, container is ParamSpec
// is method
func (v ParamSpec) GetQdata(quark uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(67, "ParamSpec", "get_qdata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_quark := gi.NewUint32Argument(quark)
	args := []gi.Argument{arg_v, arg_quark}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_param_spec_get_redirect_target
// container is not nil, container is ParamSpec
// is method
func (v ParamSpec) GetRedirectTarget() (result ParamSpec) {
	iv, err := _I.Get(68, "ParamSpec", "get_redirect_target")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_param_spec_set_qdata
// container is not nil, container is ParamSpec
// is method
func (v ParamSpec) SetQdata(quark uint32, data unsafe.Pointer) {
	iv, err := _I.Get(69, "ParamSpec", "set_qdata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_quark := gi.NewUint32Argument(quark)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_v, arg_quark, arg_data}
	iv.Call(args, nil, nil)
}

// g_param_spec_sink
// container is not nil, container is ParamSpec
// is method
func (v ParamSpec) Sink() {
	iv, err := _I.Get(70, "ParamSpec", "sink")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_param_spec_steal_qdata
// container is not nil, container is ParamSpec
// is method
func (v ParamSpec) StealQdata(quark uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(71, "ParamSpec", "steal_qdata")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_quark := gi.NewUint32Argument(quark)
	args := []gi.Argument{arg_v, arg_quark}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// Object ParamSpecBoolean
type ParamSpecBoolean struct {
	ParamSpec
}

func WrapParamSpecBoolean(p unsafe.Pointer) (r ParamSpecBoolean) { r.P = p; return }

type IParamSpecBoolean interface{ P_ParamSpecBoolean() unsafe.Pointer }

func (v ParamSpecBoolean) P_ParamSpecBoolean() unsafe.Pointer { return v.P }
func ParamSpecBooleanGetType() gi.GType {
	ret := _I.GetGType(16, "ParamSpecBoolean")
	return ret
}

// Object ParamSpecBoxed
type ParamSpecBoxed struct {
	ParamSpec
}

func WrapParamSpecBoxed(p unsafe.Pointer) (r ParamSpecBoxed) { r.P = p; return }

type IParamSpecBoxed interface{ P_ParamSpecBoxed() unsafe.Pointer }

func (v ParamSpecBoxed) P_ParamSpecBoxed() unsafe.Pointer { return v.P }
func ParamSpecBoxedGetType() gi.GType {
	ret := _I.GetGType(17, "ParamSpecBoxed")
	return ret
}

// Object ParamSpecChar
type ParamSpecChar struct {
	ParamSpec
}

func WrapParamSpecChar(p unsafe.Pointer) (r ParamSpecChar) { r.P = p; return }

type IParamSpecChar interface{ P_ParamSpecChar() unsafe.Pointer }

func (v ParamSpecChar) P_ParamSpecChar() unsafe.Pointer { return v.P }
func ParamSpecCharGetType() gi.GType {
	ret := _I.GetGType(18, "ParamSpecChar")
	return ret
}

// ignore GType struct ParamSpecClass
// Object ParamSpecDouble
type ParamSpecDouble struct {
	ParamSpec
}

func WrapParamSpecDouble(p unsafe.Pointer) (r ParamSpecDouble) { r.P = p; return }

type IParamSpecDouble interface{ P_ParamSpecDouble() unsafe.Pointer }

func (v ParamSpecDouble) P_ParamSpecDouble() unsafe.Pointer { return v.P }
func ParamSpecDoubleGetType() gi.GType {
	ret := _I.GetGType(19, "ParamSpecDouble")
	return ret
}

// Object ParamSpecEnum
type ParamSpecEnum struct {
	ParamSpec
}

func WrapParamSpecEnum(p unsafe.Pointer) (r ParamSpecEnum) { r.P = p; return }

type IParamSpecEnum interface{ P_ParamSpecEnum() unsafe.Pointer }

func (v ParamSpecEnum) P_ParamSpecEnum() unsafe.Pointer { return v.P }
func ParamSpecEnumGetType() gi.GType {
	ret := _I.GetGType(20, "ParamSpecEnum")
	return ret
}

// Object ParamSpecFlags
type ParamSpecFlags struct {
	ParamSpec
}

func WrapParamSpecFlags(p unsafe.Pointer) (r ParamSpecFlags) { r.P = p; return }

type IParamSpecFlags interface{ P_ParamSpecFlags() unsafe.Pointer }

func (v ParamSpecFlags) P_ParamSpecFlags() unsafe.Pointer { return v.P }
func ParamSpecFlagsGetType() gi.GType {
	ret := _I.GetGType(21, "ParamSpecFlags")
	return ret
}

// Object ParamSpecFloat
type ParamSpecFloat struct {
	ParamSpec
}

func WrapParamSpecFloat(p unsafe.Pointer) (r ParamSpecFloat) { r.P = p; return }

type IParamSpecFloat interface{ P_ParamSpecFloat() unsafe.Pointer }

func (v ParamSpecFloat) P_ParamSpecFloat() unsafe.Pointer { return v.P }
func ParamSpecFloatGetType() gi.GType {
	ret := _I.GetGType(22, "ParamSpecFloat")
	return ret
}

// Object ParamSpecGType
type ParamSpecGType struct {
	ParamSpec
}

func WrapParamSpecGType(p unsafe.Pointer) (r ParamSpecGType) { r.P = p; return }

type IParamSpecGType interface{ P_ParamSpecGType() unsafe.Pointer }

func (v ParamSpecGType) P_ParamSpecGType() unsafe.Pointer { return v.P }
func ParamSpecGTypeGetType() gi.GType {
	ret := _I.GetGType(23, "ParamSpecGType")
	return ret
}

// Object ParamSpecInt
type ParamSpecInt struct {
	ParamSpec
}

func WrapParamSpecInt(p unsafe.Pointer) (r ParamSpecInt) { r.P = p; return }

type IParamSpecInt interface{ P_ParamSpecInt() unsafe.Pointer }

func (v ParamSpecInt) P_ParamSpecInt() unsafe.Pointer { return v.P }
func ParamSpecIntGetType() gi.GType {
	ret := _I.GetGType(24, "ParamSpecInt")
	return ret
}

// Object ParamSpecInt64
type ParamSpecInt64 struct {
	ParamSpec
}

func WrapParamSpecInt64(p unsafe.Pointer) (r ParamSpecInt64) { r.P = p; return }

type IParamSpecInt64 interface{ P_ParamSpecInt64() unsafe.Pointer }

func (v ParamSpecInt64) P_ParamSpecInt64() unsafe.Pointer { return v.P }
func ParamSpecInt64GetType() gi.GType {
	ret := _I.GetGType(25, "ParamSpecInt64")
	return ret
}

// Object ParamSpecLong
type ParamSpecLong struct {
	ParamSpec
}

func WrapParamSpecLong(p unsafe.Pointer) (r ParamSpecLong) { r.P = p; return }

type IParamSpecLong interface{ P_ParamSpecLong() unsafe.Pointer }

func (v ParamSpecLong) P_ParamSpecLong() unsafe.Pointer { return v.P }
func ParamSpecLongGetType() gi.GType {
	ret := _I.GetGType(26, "ParamSpecLong")
	return ret
}

// Object ParamSpecObject
type ParamSpecObject struct {
	ParamSpec
}

func WrapParamSpecObject(p unsafe.Pointer) (r ParamSpecObject) { r.P = p; return }

type IParamSpecObject interface{ P_ParamSpecObject() unsafe.Pointer }

func (v ParamSpecObject) P_ParamSpecObject() unsafe.Pointer { return v.P }
func ParamSpecObjectGetType() gi.GType {
	ret := _I.GetGType(27, "ParamSpecObject")
	return ret
}

// Object ParamSpecOverride
type ParamSpecOverride struct {
	ParamSpec
}

func WrapParamSpecOverride(p unsafe.Pointer) (r ParamSpecOverride) { r.P = p; return }

type IParamSpecOverride interface{ P_ParamSpecOverride() unsafe.Pointer }

func (v ParamSpecOverride) P_ParamSpecOverride() unsafe.Pointer { return v.P }
func ParamSpecOverrideGetType() gi.GType {
	ret := _I.GetGType(28, "ParamSpecOverride")
	return ret
}

// Object ParamSpecParam
type ParamSpecParam struct {
	ParamSpec
}

func WrapParamSpecParam(p unsafe.Pointer) (r ParamSpecParam) { r.P = p; return }

type IParamSpecParam interface{ P_ParamSpecParam() unsafe.Pointer }

func (v ParamSpecParam) P_ParamSpecParam() unsafe.Pointer { return v.P }
func ParamSpecParamGetType() gi.GType {
	ret := _I.GetGType(29, "ParamSpecParam")
	return ret
}

// Object ParamSpecPointer
type ParamSpecPointer struct {
	ParamSpec
}

func WrapParamSpecPointer(p unsafe.Pointer) (r ParamSpecPointer) { r.P = p; return }

type IParamSpecPointer interface{ P_ParamSpecPointer() unsafe.Pointer }

func (v ParamSpecPointer) P_ParamSpecPointer() unsafe.Pointer { return v.P }
func ParamSpecPointerGetType() gi.GType {
	ret := _I.GetGType(30, "ParamSpecPointer")
	return ret
}

// Struct ParamSpecPool
type ParamSpecPool struct {
	P unsafe.Pointer
}

func ParamSpecPoolGetType() gi.GType {
	ret := _I.GetGType(31, "ParamSpecPool")
	return ret
}

// g_param_spec_pool_insert
// container is not nil, container is ParamSpecPool
// is method
func (v ParamSpecPool) Insert(pspec IParamSpec, owner_type gi.GType) {
	iv, err := _I.Get(72, "ParamSpecPool", "insert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pspec != nil {
		tmp = pspec.P_ParamSpec()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pspec := gi.NewPointerArgument(tmp)
	arg_owner_type := gi.NewUintArgument(uint(owner_type))
	args := []gi.Argument{arg_v, arg_pspec, arg_owner_type}
	iv.Call(args, nil, nil)
}

// g_param_spec_pool_list
// container is not nil, container is ParamSpecPool
// is method
// ret lenArgIdx 1
func (v ParamSpecPool) List(owner_type gi.GType) (result gi.PointerArray) {
	iv, err := _I.Get(73, "ParamSpecPool", "list")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_v := gi.NewPointerArgument(v.P)
	arg_owner_type := gi.NewUintArgument(uint(owner_type))
	arg_n_pspecs_p := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_v, arg_owner_type, arg_n_pspecs_p}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_pspecs_p uint32
	_ = n_pspecs_p
	n_pspecs_p = outArgs[0].Uint32()
	result = gi.PointerArray{P: ret.Pointer(), Len: int(n_pspecs_p)}
	return
}

// g_param_spec_pool_list_owned
// container is not nil, container is ParamSpecPool
// is method
func (v ParamSpecPool) ListOwned(owner_type gi.GType) (result glib.List) {
	iv, err := _I.Get(74, "ParamSpecPool", "list_owned")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_owner_type := gi.NewUintArgument(uint(owner_type))
	args := []gi.Argument{arg_v, arg_owner_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_param_spec_pool_lookup
// container is not nil, container is ParamSpecPool
// is method
func (v ParamSpecPool) Lookup(param_name string, owner_type gi.GType, walk_ancestors bool) (result ParamSpec) {
	iv, err := _I.Get(75, "ParamSpecPool", "lookup")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_param_name := gi.CString(param_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_param_name := gi.NewStringArgument(c_param_name)
	arg_owner_type := gi.NewUintArgument(uint(owner_type))
	arg_walk_ancestors := gi.NewBoolArgument(walk_ancestors)
	args := []gi.Argument{arg_v, arg_param_name, arg_owner_type, arg_walk_ancestors}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_param_name)
	result.P = ret.Pointer()
	return
}

// g_param_spec_pool_remove
// container is not nil, container is ParamSpecPool
// is method
func (v ParamSpecPool) Remove(pspec IParamSpec) {
	iv, err := _I.Get(76, "ParamSpecPool", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pspec != nil {
		tmp = pspec.P_ParamSpec()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_pspec := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_pspec}
	iv.Call(args, nil, nil)
}

// g_param_spec_pool_new
// container is not nil, container is ParamSpecPool
// is method
// arg0Type tag: gboolean, isPtr: false
func ParamSpecPoolNew1(type_prefixing bool) (result ParamSpecPool) {
	iv, err := _I.Get(77, "ParamSpecPool", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type_prefixing := gi.NewBoolArgument(type_prefixing)
	args := []gi.Argument{arg_type_prefixing}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// Object ParamSpecString
type ParamSpecString struct {
	ParamSpec
}

func WrapParamSpecString(p unsafe.Pointer) (r ParamSpecString) { r.P = p; return }

type IParamSpecString interface{ P_ParamSpecString() unsafe.Pointer }

func (v ParamSpecString) P_ParamSpecString() unsafe.Pointer { return v.P }
func ParamSpecStringGetType() gi.GType {
	ret := _I.GetGType(32, "ParamSpecString")
	return ret
}

// Struct ParamSpecTypeInfo
type ParamSpecTypeInfo struct {
	P unsafe.Pointer
}

const SizeOfStructParamSpecTypeInfo = 56

func ParamSpecTypeInfoGetType() gi.GType {
	ret := _I.GetGType(33, "ParamSpecTypeInfo")
	return ret
}

// Object ParamSpecUChar
type ParamSpecUChar struct {
	ParamSpec
}

func WrapParamSpecUChar(p unsafe.Pointer) (r ParamSpecUChar) { r.P = p; return }

type IParamSpecUChar interface{ P_ParamSpecUChar() unsafe.Pointer }

func (v ParamSpecUChar) P_ParamSpecUChar() unsafe.Pointer { return v.P }
func ParamSpecUCharGetType() gi.GType {
	ret := _I.GetGType(34, "ParamSpecUChar")
	return ret
}

// Object ParamSpecUInt
type ParamSpecUInt struct {
	ParamSpec
}

func WrapParamSpecUInt(p unsafe.Pointer) (r ParamSpecUInt) { r.P = p; return }

type IParamSpecUInt interface{ P_ParamSpecUInt() unsafe.Pointer }

func (v ParamSpecUInt) P_ParamSpecUInt() unsafe.Pointer { return v.P }
func ParamSpecUIntGetType() gi.GType {
	ret := _I.GetGType(35, "ParamSpecUInt")
	return ret
}

// Object ParamSpecUInt64
type ParamSpecUInt64 struct {
	ParamSpec
}

func WrapParamSpecUInt64(p unsafe.Pointer) (r ParamSpecUInt64) { r.P = p; return }

type IParamSpecUInt64 interface{ P_ParamSpecUInt64() unsafe.Pointer }

func (v ParamSpecUInt64) P_ParamSpecUInt64() unsafe.Pointer { return v.P }
func ParamSpecUInt64GetType() gi.GType {
	ret := _I.GetGType(36, "ParamSpecUInt64")
	return ret
}

// Object ParamSpecULong
type ParamSpecULong struct {
	ParamSpec
}

func WrapParamSpecULong(p unsafe.Pointer) (r ParamSpecULong) { r.P = p; return }

type IParamSpecULong interface{ P_ParamSpecULong() unsafe.Pointer }

func (v ParamSpecULong) P_ParamSpecULong() unsafe.Pointer { return v.P }
func ParamSpecULongGetType() gi.GType {
	ret := _I.GetGType(37, "ParamSpecULong")
	return ret
}

// Object ParamSpecUnichar
type ParamSpecUnichar struct {
	ParamSpec
}

func WrapParamSpecUnichar(p unsafe.Pointer) (r ParamSpecUnichar) { r.P = p; return }

type IParamSpecUnichar interface{ P_ParamSpecUnichar() unsafe.Pointer }

func (v ParamSpecUnichar) P_ParamSpecUnichar() unsafe.Pointer { return v.P }
func ParamSpecUnicharGetType() gi.GType {
	ret := _I.GetGType(38, "ParamSpecUnichar")
	return ret
}

// Object ParamSpecValueArray
type ParamSpecValueArray struct {
	ParamSpec
}

func WrapParamSpecValueArray(p unsafe.Pointer) (r ParamSpecValueArray) { r.P = p; return }

type IParamSpecValueArray interface{ P_ParamSpecValueArray() unsafe.Pointer }

func (v ParamSpecValueArray) P_ParamSpecValueArray() unsafe.Pointer { return v.P }
func ParamSpecValueArrayGetType() gi.GType {
	ret := _I.GetGType(39, "ParamSpecValueArray")
	return ret
}

// Object ParamSpecVariant
type ParamSpecVariant struct {
	ParamSpec
}

func WrapParamSpecVariant(p unsafe.Pointer) (r ParamSpecVariant) { r.P = p; return }

type IParamSpecVariant interface{ P_ParamSpecVariant() unsafe.Pointer }

func (v ParamSpecVariant) P_ParamSpecVariant() unsafe.Pointer { return v.P }
func ParamSpecVariantGetType() gi.GType {
	ret := _I.GetGType(40, "ParamSpecVariant")
	return ret
}

// Struct Parameter
type Parameter struct {
	P unsafe.Pointer
}

const SizeOfStructParameter = 32

func ParameterGetType() gi.GType {
	ret := _I.GetGType(41, "Parameter")
	return ret
}

type SignalAccumulatorStruct struct {
	F_ihint          SignalInvocationHint
	F_return_accu    Value
	F_handler_return Value
	F_data           unsafe.Pointer
}

func GetPointer_mySignalAccumulator() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectSignalAccumulator())
}

//export myGObjectSignalAccumulator
func myGObjectSignalAccumulator(ihint *C.GSignalInvocationHint, return_accu *C.GValue, handler_return *C.GValue, data C.gpointer) {
	// TODO: not found user_data
}

type SignalEmissionHookStruct struct {
	F_ihint          SignalInvocationHint
	F_n_param_values uint32
	F_param_values   unsafe.Pointer
	F_data           unsafe.Pointer
}

func GetPointer_mySignalEmissionHook() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectSignalEmissionHook())
}

//export myGObjectSignalEmissionHook
func myGObjectSignalEmissionHook(ihint *C.GSignalInvocationHint, n_param_values C.guint32, param_values C.gpointer, data C.gpointer) {
	// TODO: not found user_data
}

// Flags SignalFlags
type SignalFlags int

const (
	SignalFlagsRunFirst    SignalFlags = 1
	SignalFlagsRunLast     SignalFlags = 2
	SignalFlagsRunCleanup  SignalFlags = 4
	SignalFlagsNoRecurse   SignalFlags = 8
	SignalFlagsDetailed    SignalFlags = 16
	SignalFlagsAction      SignalFlags = 32
	SignalFlagsNoHooks     SignalFlags = 64
	SignalFlagsMustCollect SignalFlags = 128
	SignalFlagsDeprecated  SignalFlags = 256
)

func SignalFlagsGetType() gi.GType {
	ret := _I.GetGType(42, "SignalFlags")
	return ret
}

// Struct SignalInvocationHint
type SignalInvocationHint struct {
	P unsafe.Pointer
}

const SizeOfStructSignalInvocationHint = 12

func SignalInvocationHintGetType() gi.GType {
	ret := _I.GetGType(43, "SignalInvocationHint")
	return ret
}

// Flags SignalMatchType
type SignalMatchTypeFlags int

const (
	SignalMatchTypeId        SignalMatchTypeFlags = 1
	SignalMatchTypeDetail    SignalMatchTypeFlags = 2
	SignalMatchTypeClosure   SignalMatchTypeFlags = 4
	SignalMatchTypeFunc      SignalMatchTypeFlags = 8
	SignalMatchTypeData      SignalMatchTypeFlags = 16
	SignalMatchTypeUnblocked SignalMatchTypeFlags = 32
)

func SignalMatchTypeGetType() gi.GType {
	ret := _I.GetGType(44, "SignalMatchType")
	return ret
}

// Struct SignalQuery
type SignalQuery struct {
	P unsafe.Pointer
}

const SizeOfStructSignalQuery = 56

func SignalQueryGetType() gi.GType {
	ret := _I.GetGType(45, "SignalQuery")
	return ret
}

type ToggleNotifyStruct struct {
	F_data        unsafe.Pointer
	F_object      Object
	F_is_last_ref bool
}

func GetPointer_myToggleNotify() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectToggleNotify())
}

//export myGObjectToggleNotify
func myGObjectToggleNotify(data C.gpointer, object *C.GObject, is_last_ref C.gboolean) {
	// TODO: not found user_data
}

// Union TypeCValue
type TypeCValue struct {
	P unsafe.Pointer
}

const SizeOfUnionTypeCValue = 8

func TypeCValueGetType() gi.GType {
	ret := _I.GetGType(46, "TypeCValue")
	return ret
}

// Struct TypeClass
type TypeClass struct {
	P unsafe.Pointer
}

const SizeOfStructTypeClass = 8

func TypeClassGetType() gi.GType {
	ret := _I.GetGType(47, "TypeClass")
	return ret
}

// g_type_class_add_private
// container is not nil, container is TypeClass
// is method
func (v TypeClass) AddPrivate(private_size uint64) {
	iv, err := _I.Get(78, "TypeClass", "add_private")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_private_size := gi.NewUint64Argument(private_size)
	args := []gi.Argument{arg_v, arg_private_size}
	iv.Call(args, nil, nil)
}

// g_type_class_get_private
// container is not nil, container is TypeClass
// is method
func (v TypeClass) GetPrivate(private_type gi.GType) (result unsafe.Pointer) {
	iv, err := _I.Get(79, "TypeClass", "get_private")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_private_type := gi.NewUintArgument(uint(private_type))
	args := []gi.Argument{arg_v, arg_private_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_type_class_peek_parent
// container is not nil, container is TypeClass
// is method
func (v TypeClass) PeekParent() (result TypeClass) {
	iv, err := _I.Get(80, "TypeClass", "peek_parent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_class_unref
// container is not nil, container is TypeClass
// is method
func (v TypeClass) Unref() {
	iv, err := _I.Get(81, "TypeClass", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_type_class_adjust_private_offset
// container is not nil, container is TypeClass
// is method
// arg0Type tag: void, isPtr: true
func TypeClassAdjustPrivateOffset1(g_class unsafe.Pointer, private_size_or_offset int32) {
	iv, err := _I.Get(82, "TypeClass", "adjust_private_offset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_g_class := gi.NewPointerArgument(g_class)
	arg_private_size_or_offset := gi.NewInt32Argument(private_size_or_offset)
	args := []gi.Argument{arg_g_class, arg_private_size_or_offset}
	iv.Call(args, nil, nil)
}

// g_type_class_peek
// container is not nil, container is TypeClass
// is method
// arg0Type tag: GType, isPtr: false
func TypeClassPeek1(type1 gi.GType) (result TypeClass) {
	iv, err := _I.Get(83, "TypeClass", "peek")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_class_peek_static
// container is not nil, container is TypeClass
// is method
// arg0Type tag: GType, isPtr: false
func TypeClassPeekStatic1(type1 gi.GType) (result TypeClass) {
	iv, err := _I.Get(84, "TypeClass", "peek_static")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_class_ref
// container is not nil, container is TypeClass
// is method
// arg0Type tag: GType, isPtr: false
func TypeClassRef1(type1 gi.GType) (result TypeClass) {
	iv, err := _I.Get(85, "TypeClass", "ref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

type TypeClassCacheFuncStruct struct {
	F_cache_data unsafe.Pointer
	F_g_class    TypeClass
}

func GetPointer_myTypeClassCacheFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectTypeClassCacheFunc())
}

//export myGObjectTypeClassCacheFunc
func myGObjectTypeClassCacheFunc(cache_data C.gpointer, g_class *C.GTypeClass) {
	// TODO: not found user_data
}

// Flags TypeDebugFlags
type TypeDebugFlags int

const (
	TypeDebugFlagsNone          TypeDebugFlags = 0
	TypeDebugFlagsObjects       TypeDebugFlags = 1
	TypeDebugFlagsSignals       TypeDebugFlags = 2
	TypeDebugFlagsInstanceCount TypeDebugFlags = 4
	TypeDebugFlagsMask          TypeDebugFlags = 7
)

func TypeDebugFlagsGetType() gi.GType {
	ret := _I.GetGType(48, "TypeDebugFlags")
	return ret
}

// Flags TypeFlags
type TypeFlags int

const (
	TypeFlagsAbstract      TypeFlags = 16
	TypeFlagsValueAbstract TypeFlags = 32
)

func TypeFlagsGetType() gi.GType {
	ret := _I.GetGType(49, "TypeFlags")
	return ret
}

// Flags TypeFundamentalFlags
type TypeFundamentalFlags int

const (
	TypeFundamentalFlagsClassed        TypeFundamentalFlags = 1
	TypeFundamentalFlagsInstantiatable TypeFundamentalFlags = 2
	TypeFundamentalFlagsDerivable      TypeFundamentalFlags = 4
	TypeFundamentalFlagsDeepDerivable  TypeFundamentalFlags = 8
)

func TypeFundamentalFlagsGetType() gi.GType {
	ret := _I.GetGType(50, "TypeFundamentalFlags")
	return ret
}

// Struct TypeFundamentalInfo
type TypeFundamentalInfo struct {
	P unsafe.Pointer
}

const SizeOfStructTypeFundamentalInfo = 4

func TypeFundamentalInfoGetType() gi.GType {
	ret := _I.GetGType(51, "TypeFundamentalInfo")
	return ret
}

// Struct TypeInfo
type TypeInfo struct {
	P unsafe.Pointer
}

const SizeOfStructTypeInfo = 72

func TypeInfoGetType() gi.GType {
	ret := _I.GetGType(52, "TypeInfo")
	return ret
}

// Struct TypeInstance
type TypeInstance struct {
	P unsafe.Pointer
}

const SizeOfStructTypeInstance = 8

func TypeInstanceGetType() gi.GType {
	ret := _I.GetGType(53, "TypeInstance")
	return ret
}

// g_type_instance_get_private
// container is not nil, container is TypeInstance
// is method
func (v TypeInstance) GetPrivate(private_type gi.GType) (result unsafe.Pointer) {
	iv, err := _I.Get(86, "TypeInstance", "get_private")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_private_type := gi.NewUintArgument(uint(private_type))
	args := []gi.Argument{arg_v, arg_private_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// Struct TypeInterface
type TypeInterface struct {
	P unsafe.Pointer
}

const SizeOfStructTypeInterface = 16

func TypeInterfaceGetType() gi.GType {
	ret := _I.GetGType(54, "TypeInterface")
	return ret
}

// g_type_interface_peek_parent
// container is not nil, container is TypeInterface
// is method
func (v TypeInterface) PeekParent() (result TypeInterface) {
	iv, err := _I.Get(87, "TypeInterface", "peek_parent")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_interface_add_prerequisite
// container is not nil, container is TypeInterface
// is method
// arg0Type tag: GType, isPtr: false
func TypeInterfaceAddPrerequisite1(interface_type gi.GType, prerequisite_type gi.GType) {
	iv, err := _I.Get(88, "TypeInterface", "add_prerequisite")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_interface_type := gi.NewUintArgument(uint(interface_type))
	arg_prerequisite_type := gi.NewUintArgument(uint(prerequisite_type))
	args := []gi.Argument{arg_interface_type, arg_prerequisite_type}
	iv.Call(args, nil, nil)
}

// g_type_interface_get_plugin
// container is not nil, container is TypeInterface
// is method
// arg0Type tag: GType, isPtr: false
func TypeInterfaceGetPlugin1(instance_type gi.GType, interface_type gi.GType) (result TypePlugin) {
	iv, err := _I.Get(89, "TypeInterface", "get_plugin")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_instance_type := gi.NewUintArgument(uint(instance_type))
	arg_interface_type := gi.NewUintArgument(uint(interface_type))
	args := []gi.Argument{arg_instance_type, arg_interface_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_interface_peek
// container is not nil, container is TypeInterface
// is method
// arg0Type tag: interface, isPtr: true
func TypeInterfacePeek1(instance_class TypeClass, iface_type gi.GType) (result TypeInterface) {
	iv, err := _I.Get(90, "TypeInterface", "peek")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_instance_class := gi.NewPointerArgument(instance_class.P)
	arg_iface_type := gi.NewUintArgument(uint(iface_type))
	args := []gi.Argument{arg_instance_class, arg_iface_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_interface_prerequisites
// container is not nil, container is TypeInterface
// is method
// arg0Type tag: GType, isPtr: false
// ret lenArgIdx 1
func TypeInterfacePrerequisites1(interface_type gi.GType) (result gi.GTypeArray) {
	iv, err := _I.Get(91, "TypeInterface", "prerequisites")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_interface_type := gi.NewUintArgument(uint(interface_type))
	arg_n_prerequisites := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_interface_type, arg_n_prerequisites}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_prerequisites uint32
	_ = n_prerequisites
	n_prerequisites = outArgs[0].Uint32()
	result = gi.GTypeArray{P: ret.Pointer(), Len: int(n_prerequisites)}
	return
}

type TypeInterfaceCheckFuncStruct struct {
	F_check_data unsafe.Pointer
	F_g_iface    TypeInterface
}

func GetPointer_myTypeInterfaceCheckFunc() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectTypeInterfaceCheckFunc())
}

//export myGObjectTypeInterfaceCheckFunc
func myGObjectTypeInterfaceCheckFunc(check_data C.gpointer, g_iface *C.GTypeInterface) {
	// TODO: not found user_data
}

// Object TypeModule
type TypeModule struct {
	TypePluginIfc
	Object
}

func WrapTypeModule(p unsafe.Pointer) (r TypeModule) { r.P = p; return }

type ITypeModule interface{ P_TypeModule() unsafe.Pointer }

func (v TypeModule) P_TypeModule() unsafe.Pointer { return v.P }
func (v TypeModule) P_TypePlugin() unsafe.Pointer { return v.P }
func TypeModuleGetType() gi.GType {
	ret := _I.GetGType(55, "TypeModule")
	return ret
}

// g_type_module_add_interface
// container is not nil, container is TypeModule
// is method
func (v TypeModule) AddInterface(instance_type gi.GType, interface_type gi.GType, interface_info InterfaceInfo) {
	iv, err := _I.Get(92, "TypeModule", "add_interface")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_instance_type := gi.NewUintArgument(uint(instance_type))
	arg_interface_type := gi.NewUintArgument(uint(interface_type))
	arg_interface_info := gi.NewPointerArgument(interface_info.P)
	args := []gi.Argument{arg_v, arg_instance_type, arg_interface_type, arg_interface_info}
	iv.Call(args, nil, nil)
}

// g_type_module_register_enum
// container is not nil, container is TypeModule
// is method
func (v TypeModule) RegisterEnum(name string, const_static_values EnumValue) (result gi.GType) {
	iv, err := _I.Get(93, "TypeModule", "register_enum")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	arg_const_static_values := gi.NewPointerArgument(const_static_values.P)
	args := []gi.Argument{arg_v, arg_name, arg_const_static_values}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = gi.GType(ret.Uint())
	return
}

// g_type_module_register_flags
// container is not nil, container is TypeModule
// is method
func (v TypeModule) RegisterFlags(name string, const_static_values FlagsValue) (result gi.GType) {
	iv, err := _I.Get(94, "TypeModule", "register_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	arg_const_static_values := gi.NewPointerArgument(const_static_values.P)
	args := []gi.Argument{arg_v, arg_name, arg_const_static_values}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = gi.GType(ret.Uint())
	return
}

// g_type_module_register_type
// container is not nil, container is TypeModule
// is method
func (v TypeModule) RegisterType(parent_type gi.GType, type_name string, type_info TypeInfo, flags TypeFlags) (result gi.GType) {
	iv, err := _I.Get(95, "TypeModule", "register_type")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_type_name := gi.CString(type_name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_parent_type := gi.NewUintArgument(uint(parent_type))
	arg_type_name := gi.NewStringArgument(c_type_name)
	arg_type_info := gi.NewPointerArgument(type_info.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_v, arg_parent_type, arg_type_name, arg_type_info, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_type_name)
	result = gi.GType(ret.Uint())
	return
}

// g_type_module_set_name
// container is not nil, container is TypeModule
// is method
func (v TypeModule) SetName(name string) {
	iv, err := _I.Get(96, "TypeModule", "set_name")
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

// g_type_module_unuse
// container is not nil, container is TypeModule
// is method
func (v TypeModule) Unuse() {
	iv, err := _I.Get(97, "TypeModule", "unuse")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_type_module_use
// container is not nil, container is TypeModule
// is method
func (v TypeModule) Use() (result bool) {
	iv, err := _I.Get(98, "TypeModule", "use")
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

// ignore GType struct TypeModuleClass
// Interface TypePlugin
type TypePlugin struct {
	TypePluginIfc
	P unsafe.Pointer
}
type TypePluginIfc struct{}
type ITypePlugin interface{ P_TypePlugin() unsafe.Pointer }

func (v TypePlugin) P_TypePlugin() unsafe.Pointer { return v.P }
func TypePluginGetType() gi.GType {
	ret := _I.GetGType(56, "TypePlugin")
	return ret
}

// g_type_plugin_complete_interface_info
// container is not nil, container is TypePlugin
// is method
func (v *TypePluginIfc) CompleteInterfaceInfo(instance_type gi.GType, interface_type gi.GType, info InterfaceInfo) {
	iv, err := _I.Get(99, "TypePlugin", "complete_interface_info")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_instance_type := gi.NewUintArgument(uint(instance_type))
	arg_interface_type := gi.NewUintArgument(uint(interface_type))
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_v, arg_instance_type, arg_interface_type, arg_info}
	iv.Call(args, nil, nil)
}

// g_type_plugin_complete_type_info
// container is not nil, container is TypePlugin
// is method
func (v *TypePluginIfc) CompleteTypeInfo(g_type gi.GType, info TypeInfo, value_table TypeValueTable) {
	iv, err := _I.Get(100, "TypePlugin", "complete_type_info")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	arg_g_type := gi.NewUintArgument(uint(g_type))
	arg_info := gi.NewPointerArgument(info.P)
	arg_value_table := gi.NewPointerArgument(value_table.P)
	args := []gi.Argument{arg_v, arg_g_type, arg_info, arg_value_table}
	iv.Call(args, nil, nil)
}

// g_type_plugin_unuse
// container is not nil, container is TypePlugin
// is method
func (v *TypePluginIfc) Unuse() {
	iv, err := _I.Get(101, "TypePlugin", "unuse")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_type_plugin_use
// container is not nil, container is TypePlugin
// is method
func (v *TypePluginIfc) Use() {
	iv, err := _I.Get(102, "TypePlugin", "use")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// ignore Class struct TypePluginClass, type of TypePlugin is interface
type TypePluginCompleteInterfaceInfoStruct struct {
	F_plugin         TypePlugin
	F_instance_type  unsafe.Pointer /*TODO_CB tag: GType, isPtr: false*/
	F_interface_type unsafe.Pointer /*TODO_CB tag: GType, isPtr: false*/
	F_info           InterfaceInfo
}

func GetPointer_myTypePluginCompleteInterfaceInfo() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectTypePluginCompleteInterfaceInfo())
}

//export myGObjectTypePluginCompleteInterfaceInfo
func myGObjectTypePluginCompleteInterfaceInfo(plugin *C.GTypePlugin, instance_type C.gpointer, interface_type C.gpointer, info *C.GInterfaceInfo) {
	// TODO: not found user_data
}

type TypePluginCompleteTypeInfoStruct struct {
	F_plugin      TypePlugin
	F_g_type      unsafe.Pointer /*TODO_CB tag: GType, isPtr: false*/
	F_info        TypeInfo
	F_value_table TypeValueTable
}

func GetPointer_myTypePluginCompleteTypeInfo() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectTypePluginCompleteTypeInfo())
}

//export myGObjectTypePluginCompleteTypeInfo
func myGObjectTypePluginCompleteTypeInfo(plugin *C.GTypePlugin, g_type C.gpointer, info *C.GTypeInfo, value_table *C.GTypeValueTable) {
	// TODO: not found user_data
}

type TypePluginUnuseStruct struct {
	F_plugin TypePlugin
}

func GetPointer_myTypePluginUnuse() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectTypePluginUnuse())
}

//export myGObjectTypePluginUnuse
func myGObjectTypePluginUnuse(plugin *C.GTypePlugin) {
	// TODO: not found user_data
}

type TypePluginUseStruct struct {
	F_plugin TypePlugin
}

func GetPointer_myTypePluginUse() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectTypePluginUse())
}

//export myGObjectTypePluginUse
func myGObjectTypePluginUse(plugin *C.GTypePlugin) {
	// TODO: not found user_data
}

// Struct TypeQuery
type TypeQuery struct {
	P unsafe.Pointer
}

const SizeOfStructTypeQuery = 24

func TypeQueryGetType() gi.GType {
	ret := _I.GetGType(57, "TypeQuery")
	return ret
}

// Struct TypeValueTable
type TypeValueTable struct {
	P unsafe.Pointer
}

const SizeOfStructTypeValueTable = 64

func TypeValueTableGetType() gi.GType {
	ret := _I.GetGType(58, "TypeValueTable")
	return ret
}

// Struct Value
type Value struct {
	P unsafe.Pointer
}

const SizeOfStructValue = 24

func ValueGetType() gi.GType {
	ret := _I.GetGType(59, "Value")
	return ret
}

// g_value_copy
// container is not nil, container is Value
// is method
func (v Value) Copy(dest_value Value) {
	iv, err := _I.Get(103, "Value", "copy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dest_value := gi.NewPointerArgument(dest_value.P)
	args := []gi.Argument{arg_v, arg_dest_value}
	iv.Call(args, nil, nil)
}

// g_value_dup_object
// container is not nil, container is Value
// is method
func (v Value) DupObject() (result Object) {
	iv, err := _I.Get(104, "Value", "dup_object")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_value_dup_string
// container is not nil, container is Value
// is method
func (v Value) DupString() (result string) {
	iv, err := _I.Get(105, "Value", "dup_string")
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

// g_value_dup_variant
// container is not nil, container is Value
// is method
func (v Value) DupVariant() (result glib.Variant) {
	iv, err := _I.Get(106, "Value", "dup_variant")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_value_fits_pointer
// container is not nil, container is Value
// is method
func (v Value) FitsPointer() (result bool) {
	iv, err := _I.Get(107, "Value", "fits_pointer")
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

// g_value_get_boolean
// container is not nil, container is Value
// is method
func (v Value) GetBoolean() (result bool) {
	iv, err := _I.Get(108, "Value", "get_boolean")
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

// g_value_get_boxed
// container is not nil, container is Value
// is method
func (v Value) GetBoxed() (result unsafe.Pointer) {
	iv, err := _I.Get(109, "Value", "get_boxed")
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

// g_value_get_char
// container is not nil, container is Value
// is method
func (v Value) GetChar() (result int8) {
	iv, err := _I.Get(110, "Value", "get_char")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int8()
	return
}

// g_value_get_double
// container is not nil, container is Value
// is method
func (v Value) GetDouble() (result float64) {
	iv, err := _I.Get(111, "Value", "get_double")
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

// g_value_get_enum
// container is not nil, container is Value
// is method
func (v Value) GetEnum() (result int32) {
	iv, err := _I.Get(112, "Value", "get_enum")
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

// g_value_get_flags
// container is not nil, container is Value
// is method
func (v Value) GetFlags() (result uint32) {
	iv, err := _I.Get(113, "Value", "get_flags")
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

// g_value_get_float
// container is not nil, container is Value
// is method
func (v Value) GetFloat() (result float32) {
	iv, err := _I.Get(114, "Value", "get_float")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Float()
	return
}

// g_value_get_gtype
// container is not nil, container is Value
// is method
func (v Value) GetGtype() (result gi.GType) {
	iv, err := _I.Get(115, "Value", "get_gtype")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// g_value_get_int
// container is not nil, container is Value
// is method
func (v Value) GetInt() (result int32) {
	iv, err := _I.Get(116, "Value", "get_int")
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

// g_value_get_int64
// container is not nil, container is Value
// is method
func (v Value) GetInt64() (result int64) {
	iv, err := _I.Get(117, "Value", "get_int64")
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

// g_value_get_long
// container is not nil, container is Value
// is method
func (v Value) GetLong() (result int64) {
	iv, err := _I.Get(118, "Value", "get_long")
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

// g_value_get_object
// container is not nil, container is Value
// is method
func (v Value) GetObject() (result Object) {
	iv, err := _I.Get(119, "Value", "get_object")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_value_get_param
// container is not nil, container is Value
// is method
func (v Value) GetParam() (result ParamSpec) {
	iv, err := _I.Get(120, "Value", "get_param")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_value_get_pointer
// container is not nil, container is Value
// is method
func (v Value) GetPointer() (result unsafe.Pointer) {
	iv, err := _I.Get(121, "Value", "get_pointer")
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

// g_value_get_schar
// container is not nil, container is Value
// is method
func (v Value) GetSchar() (result int8) {
	iv, err := _I.Get(122, "Value", "get_schar")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int8()
	return
}

// g_value_get_string
// container is not nil, container is Value
// is method
func (v Value) GetString() (result string) {
	iv, err := _I.Get(123, "Value", "get_string")
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

// g_value_get_uchar
// container is not nil, container is Value
// is method
func (v Value) GetUchar() (result uint8) {
	iv, err := _I.Get(124, "Value", "get_uchar")
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

// g_value_get_uint
// container is not nil, container is Value
// is method
func (v Value) GetUint() (result uint32) {
	iv, err := _I.Get(125, "Value", "get_uint")
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

// g_value_get_uint64
// container is not nil, container is Value
// is method
func (v Value) GetUint64() (result uint64) {
	iv, err := _I.Get(126, "Value", "get_uint64")
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

// g_value_get_ulong
// container is not nil, container is Value
// is method
func (v Value) GetUlong() (result uint64) {
	iv, err := _I.Get(127, "Value", "get_ulong")
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

// g_value_get_variant
// container is not nil, container is Value
// is method
func (v Value) GetVariant() (result glib.Variant) {
	iv, err := _I.Get(128, "Value", "get_variant")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_value_init
// container is not nil, container is Value
// is method
func (v Value) Init(g_type gi.GType) (result Value) {
	iv, err := _I.Get(129, "Value", "init")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_g_type := gi.NewUintArgument(uint(g_type))
	args := []gi.Argument{arg_v, arg_g_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_value_init_from_instance
// container is not nil, container is Value
// is method
func (v Value) InitFromInstance(instance TypeInstance) {
	iv, err := _I.Get(130, "Value", "init_from_instance")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_instance := gi.NewPointerArgument(instance.P)
	args := []gi.Argument{arg_v, arg_instance}
	iv.Call(args, nil, nil)
}

// g_value_peek_pointer
// container is not nil, container is Value
// is method
func (v Value) PeekPointer() (result unsafe.Pointer) {
	iv, err := _I.Get(131, "Value", "peek_pointer")
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

// g_value_reset
// container is not nil, container is Value
// is method
func (v Value) Reset() (result Value) {
	iv, err := _I.Get(132, "Value", "reset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_value_set_boolean
// container is not nil, container is Value
// is method
func (v Value) SetBoolean(v_boolean bool) {
	iv, err := _I.Get(133, "Value", "set_boolean")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_boolean := gi.NewBoolArgument(v_boolean)
	args := []gi.Argument{arg_v, arg_v_boolean}
	iv.Call(args, nil, nil)
}

// g_value_set_boxed
// container is not nil, container is Value
// is method
func (v Value) SetBoxed(v_boxed unsafe.Pointer) {
	iv, err := _I.Get(134, "Value", "set_boxed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_boxed := gi.NewPointerArgument(v_boxed)
	args := []gi.Argument{arg_v, arg_v_boxed}
	iv.Call(args, nil, nil)
}

// g_value_set_boxed_take_ownership
// container is not nil, container is Value
// is method
func (v Value) SetBoxedTakeOwnership(v_boxed unsafe.Pointer) {
	iv, err := _I.Get(135, "Value", "set_boxed_take_ownership")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_boxed := gi.NewPointerArgument(v_boxed)
	args := []gi.Argument{arg_v, arg_v_boxed}
	iv.Call(args, nil, nil)
}

// g_value_set_char
// container is not nil, container is Value
// is method
func (v Value) SetChar(v_char int8) {
	iv, err := _I.Get(136, "Value", "set_char")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_char := gi.NewInt8Argument(v_char)
	args := []gi.Argument{arg_v, arg_v_char}
	iv.Call(args, nil, nil)
}

// g_value_set_double
// container is not nil, container is Value
// is method
func (v Value) SetDouble(v_double float64) {
	iv, err := _I.Get(137, "Value", "set_double")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_double := gi.NewDoubleArgument(v_double)
	args := []gi.Argument{arg_v, arg_v_double}
	iv.Call(args, nil, nil)
}

// g_value_set_enum
// container is not nil, container is Value
// is method
func (v Value) SetEnum(v_enum int32) {
	iv, err := _I.Get(138, "Value", "set_enum")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_enum := gi.NewInt32Argument(v_enum)
	args := []gi.Argument{arg_v, arg_v_enum}
	iv.Call(args, nil, nil)
}

// g_value_set_flags
// container is not nil, container is Value
// is method
func (v Value) SetFlags(v_flags uint32) {
	iv, err := _I.Get(139, "Value", "set_flags")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_flags := gi.NewUint32Argument(v_flags)
	args := []gi.Argument{arg_v, arg_v_flags}
	iv.Call(args, nil, nil)
}

// g_value_set_float
// container is not nil, container is Value
// is method
func (v Value) SetFloat(v_float float32) {
	iv, err := _I.Get(140, "Value", "set_float")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_float := gi.NewFloatArgument(v_float)
	args := []gi.Argument{arg_v, arg_v_float}
	iv.Call(args, nil, nil)
}

// g_value_set_gtype
// container is not nil, container is Value
// is method
func (v Value) SetGtype(v_gtype gi.GType) {
	iv, err := _I.Get(141, "Value", "set_gtype")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_gtype := gi.NewUintArgument(uint(v_gtype))
	args := []gi.Argument{arg_v, arg_v_gtype}
	iv.Call(args, nil, nil)
}

// g_value_set_instance
// container is not nil, container is Value
// is method
func (v Value) SetInstance(instance unsafe.Pointer) {
	iv, err := _I.Get(142, "Value", "set_instance")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_instance := gi.NewPointerArgument(instance)
	args := []gi.Argument{arg_v, arg_instance}
	iv.Call(args, nil, nil)
}

// g_value_set_int
// container is not nil, container is Value
// is method
func (v Value) SetInt(v_int int32) {
	iv, err := _I.Get(143, "Value", "set_int")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_int := gi.NewInt32Argument(v_int)
	args := []gi.Argument{arg_v, arg_v_int}
	iv.Call(args, nil, nil)
}

// g_value_set_int64
// container is not nil, container is Value
// is method
func (v Value) SetInt64(v_int64 int64) {
	iv, err := _I.Get(144, "Value", "set_int64")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_int64 := gi.NewInt64Argument(v_int64)
	args := []gi.Argument{arg_v, arg_v_int64}
	iv.Call(args, nil, nil)
}

// g_value_set_long
// container is not nil, container is Value
// is method
func (v Value) SetLong(v_long int64) {
	iv, err := _I.Get(145, "Value", "set_long")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_long := gi.NewInt64Argument(v_long)
	args := []gi.Argument{arg_v, arg_v_long}
	iv.Call(args, nil, nil)
}

// g_value_set_object
// container is not nil, container is Value
// is method
func (v Value) SetObject(v_object IObject) {
	iv, err := _I.Get(146, "Value", "set_object")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if v_object != nil {
		tmp = v_object.P_Object()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_object := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_v_object}
	iv.Call(args, nil, nil)
}

// g_value_set_param
// container is not nil, container is Value
// is method
func (v Value) SetParam(param IParamSpec) {
	iv, err := _I.Get(147, "Value", "set_param")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if param != nil {
		tmp = param.P_ParamSpec()
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_param := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_v, arg_param}
	iv.Call(args, nil, nil)
}

// g_value_set_pointer
// container is not nil, container is Value
// is method
func (v Value) SetPointer(v_pointer unsafe.Pointer) {
	iv, err := _I.Get(148, "Value", "set_pointer")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_pointer := gi.NewPointerArgument(v_pointer)
	args := []gi.Argument{arg_v, arg_v_pointer}
	iv.Call(args, nil, nil)
}

// g_value_set_schar
// container is not nil, container is Value
// is method
func (v Value) SetSchar(v_char int8) {
	iv, err := _I.Get(149, "Value", "set_schar")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_char := gi.NewInt8Argument(v_char)
	args := []gi.Argument{arg_v, arg_v_char}
	iv.Call(args, nil, nil)
}

// g_value_set_static_boxed
// container is not nil, container is Value
// is method
func (v Value) SetStaticBoxed(v_boxed unsafe.Pointer) {
	iv, err := _I.Get(150, "Value", "set_static_boxed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_boxed := gi.NewPointerArgument(v_boxed)
	args := []gi.Argument{arg_v, arg_v_boxed}
	iv.Call(args, nil, nil)
}

// g_value_set_static_string
// container is not nil, container is Value
// is method
func (v Value) SetStaticString(v_string string) {
	iv, err := _I.Get(151, "Value", "set_static_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_v_string := gi.CString(v_string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_string := gi.NewStringArgument(c_v_string)
	args := []gi.Argument{arg_v, arg_v_string}
	iv.Call(args, nil, nil)
	gi.Free(c_v_string)
}

// g_value_set_string
// container is not nil, container is Value
// is method
func (v Value) SetString(v_string string) {
	iv, err := _I.Get(152, "Value", "set_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_v_string := gi.CString(v_string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_string := gi.NewStringArgument(c_v_string)
	args := []gi.Argument{arg_v, arg_v_string}
	iv.Call(args, nil, nil)
	gi.Free(c_v_string)
}

// g_value_set_string_take_ownership
// container is not nil, container is Value
// is method
func (v Value) SetStringTakeOwnership(v_string string) {
	iv, err := _I.Get(153, "Value", "set_string_take_ownership")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_v_string := gi.CString(v_string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_string := gi.NewStringArgument(c_v_string)
	args := []gi.Argument{arg_v, arg_v_string}
	iv.Call(args, nil, nil)
	gi.Free(c_v_string)
}

// g_value_set_uchar
// container is not nil, container is Value
// is method
func (v Value) SetUchar(v_uchar uint8) {
	iv, err := _I.Get(154, "Value", "set_uchar")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_uchar := gi.NewUint8Argument(v_uchar)
	args := []gi.Argument{arg_v, arg_v_uchar}
	iv.Call(args, nil, nil)
}

// g_value_set_uint
// container is not nil, container is Value
// is method
func (v Value) SetUint(v_uint uint32) {
	iv, err := _I.Get(155, "Value", "set_uint")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_uint := gi.NewUint32Argument(v_uint)
	args := []gi.Argument{arg_v, arg_v_uint}
	iv.Call(args, nil, nil)
}

// g_value_set_uint64
// container is not nil, container is Value
// is method
func (v Value) SetUint64(v_uint64 uint64) {
	iv, err := _I.Get(156, "Value", "set_uint64")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_uint64 := gi.NewUint64Argument(v_uint64)
	args := []gi.Argument{arg_v, arg_v_uint64}
	iv.Call(args, nil, nil)
}

// g_value_set_ulong
// container is not nil, container is Value
// is method
func (v Value) SetUlong(v_ulong uint64) {
	iv, err := _I.Get(157, "Value", "set_ulong")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_ulong := gi.NewUint64Argument(v_ulong)
	args := []gi.Argument{arg_v, arg_v_ulong}
	iv.Call(args, nil, nil)
}

// g_value_set_variant
// container is not nil, container is Value
// is method
func (v Value) SetVariant(variant glib.Variant) {
	iv, err := _I.Get(158, "Value", "set_variant")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_variant := gi.NewPointerArgument(variant.P)
	args := []gi.Argument{arg_v, arg_variant}
	iv.Call(args, nil, nil)
}

// g_value_take_boxed
// container is not nil, container is Value
// is method
func (v Value) TakeBoxed(v_boxed unsafe.Pointer) {
	iv, err := _I.Get(159, "Value", "take_boxed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_boxed := gi.NewPointerArgument(v_boxed)
	args := []gi.Argument{arg_v, arg_v_boxed}
	iv.Call(args, nil, nil)
}

// g_value_take_string
// container is not nil, container is Value
// is method
func (v Value) TakeString(v_string string) {
	iv, err := _I.Get(160, "Value", "take_string")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_v_string := gi.CString(v_string)
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_string := gi.NewStringArgument(c_v_string)
	args := []gi.Argument{arg_v, arg_v_string}
	iv.Call(args, nil, nil)
	gi.Free(c_v_string)
}

// g_value_take_variant
// container is not nil, container is Value
// is method
func (v Value) TakeVariant(variant glib.Variant) {
	iv, err := _I.Get(161, "Value", "take_variant")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_variant := gi.NewPointerArgument(variant.P)
	args := []gi.Argument{arg_v, arg_variant}
	iv.Call(args, nil, nil)
}

// g_value_transform
// container is not nil, container is Value
// is method
func (v Value) Transform(dest_value Value) (result bool) {
	iv, err := _I.Get(162, "Value", "transform")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_dest_value := gi.NewPointerArgument(dest_value.P)
	args := []gi.Argument{arg_v, arg_dest_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_value_unset
// container is not nil, container is Value
// is method
func (v Value) Unset() {
	iv, err := _I.Get(163, "Value", "unset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_value_type_compatible
// container is not nil, container is Value
// is method
// arg0Type tag: GType, isPtr: false
func ValueTypeCompatible1(src_type gi.GType, dest_type gi.GType) (result bool) {
	iv, err := _I.Get(164, "Value", "type_compatible")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_src_type := gi.NewUintArgument(uint(src_type))
	arg_dest_type := gi.NewUintArgument(uint(dest_type))
	args := []gi.Argument{arg_src_type, arg_dest_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_value_type_transformable
// container is not nil, container is Value
// is method
// arg0Type tag: GType, isPtr: false
func ValueTypeTransformable1(src_type gi.GType, dest_type gi.GType) (result bool) {
	iv, err := _I.Get(165, "Value", "type_transformable")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_src_type := gi.NewUintArgument(uint(src_type))
	arg_dest_type := gi.NewUintArgument(uint(dest_type))
	args := []gi.Argument{arg_src_type, arg_dest_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// Struct ValueArray
type ValueArray struct {
	P unsafe.Pointer
}

const SizeOfStructValueArray = 24

func ValueArrayGetType() gi.GType {
	ret := _I.GetGType(60, "ValueArray")
	return ret
}

// g_value_array_new
// container is not nil, container is ValueArray
// is constructor
func NewValueArray(n_prealloced uint32) (result ValueArray) {
	iv, err := _I.Get(166, "ValueArray", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_n_prealloced := gi.NewUint32Argument(n_prealloced)
	args := []gi.Argument{arg_n_prealloced}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_value_array_append
// container is not nil, container is ValueArray
// is method
func (v ValueArray) Append(value Value) (result ValueArray) {
	iv, err := _I.Get(167, "ValueArray", "append")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_value_array_copy
// container is not nil, container is ValueArray
// is method
func (v ValueArray) Copy() (result ValueArray) {
	iv, err := _I.Get(168, "ValueArray", "copy")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_value_array_get_nth
// container is not nil, container is ValueArray
// is method
func (v ValueArray) GetNth(index_ uint32) (result Value) {
	iv, err := _I.Get(169, "ValueArray", "get_nth")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewUint32Argument(index_)
	args := []gi.Argument{arg_v, arg_index_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_value_array_insert
// container is not nil, container is ValueArray
// is method
func (v ValueArray) Insert(index_ uint32, value Value) (result ValueArray) {
	iv, err := _I.Get(170, "ValueArray", "insert")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewUint32Argument(index_)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_index_, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_value_array_prepend
// container is not nil, container is ValueArray
// is method
func (v ValueArray) Prepend(value Value) (result ValueArray) {
	iv, err := _I.Get(171, "ValueArray", "prepend")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_v, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_value_array_remove
// container is not nil, container is ValueArray
// is method
func (v ValueArray) Remove(index_ uint32) (result ValueArray) {
	iv, err := _I.Get(172, "ValueArray", "remove")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_index_ := gi.NewUint32Argument(index_)
	args := []gi.Argument{arg_v, arg_index_}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_value_array_sort_with_data
// container is not nil, container is ValueArray
// is method
func (v ValueArray) Sort(compare_func int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result ValueArray) {
	iv, err := _I.Get(173, "ValueArray", "sort")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_compare_func := gi.NewPointerArgument(unsafe.Pointer(glib.GetPointer_myCompareDataFunc()))
	arg_user_data := gi.NewPointerArgument(user_data)
	args := []gi.Argument{arg_v, arg_compare_func, arg_user_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

type ValueTransformStruct struct {
	F_src_value  Value
	F_dest_value Value
}

func GetPointer_myValueTransform() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectValueTransform())
}

//export myGObjectValueTransform
func myGObjectValueTransform(src_value *C.GValue, dest_value *C.GValue) {
	// TODO: not found user_data
}

type WeakNotifyStruct struct {
	F_data                 unsafe.Pointer
	F_where_the_object_was Object
}

func GetPointer_myWeakNotify() unsafe.Pointer {
	return unsafe.Pointer(C.getPointer_myGObjectWeakNotify())
}

//export myGObjectWeakNotify
func myGObjectWeakNotify(data C.gpointer, where_the_object_was *C.GObject) {
	// TODO: not found user_data
}

// Struct WeakRef
type WeakRef struct {
	P unsafe.Pointer
}

func WeakRefGetType() gi.GType {
	ret := _I.GetGType(61, "WeakRef")
	return ret
}

// Union _Value__data__union
type _Value__data__union struct {
	P unsafe.Pointer
}

const SizeOfUnion_Value__data__union = 8

func _Value__data__unionGetType() gi.GType {
	ret := _I.GetGType(62, "_Value__data__union")
	return ret
}

// g_boxed_copy
// container is nil
func BoxedCopy(boxed_type gi.GType, src_boxed unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get(174, "boxed_copy", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_boxed_type := gi.NewUintArgument(uint(boxed_type))
	arg_src_boxed := gi.NewPointerArgument(src_boxed)
	args := []gi.Argument{arg_boxed_type, arg_src_boxed}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_boxed_free
// container is nil
func BoxedFree(boxed_type gi.GType, boxed unsafe.Pointer) {
	iv, err := _I.Get(175, "boxed_free", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_boxed_type := gi.NewUintArgument(uint(boxed_type))
	arg_boxed := gi.NewPointerArgument(boxed)
	args := []gi.Argument{arg_boxed_type, arg_boxed}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_BOOLEAN__BOXED_BOXED
// container is nil
func CclosureMarshalBoolean_BoxedBoxed(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(176, "cclosure_marshal_BOOLEAN__BOXED_BOXED", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_BOOLEAN__FLAGS
// container is nil
func CclosureMarshalBoolean_Flags(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(177, "cclosure_marshal_BOOLEAN__FLAGS", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_STRING__OBJECT_POINTER
// container is nil
func CclosureMarshalString_ObjectPointer(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(178, "cclosure_marshal_STRING__OBJECT_POINTER", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__BOOLEAN
// container is nil
func CclosureMarshalVoid_Boolean(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(179, "cclosure_marshal_VOID__BOOLEAN", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__BOXED
// container is nil
func CclosureMarshalVoid_Boxed(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(180, "cclosure_marshal_VOID__BOXED", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__CHAR
// container is nil
func CclosureMarshalVoid_Char(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(181, "cclosure_marshal_VOID__CHAR", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__DOUBLE
// container is nil
func CclosureMarshalVoid_Double(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(182, "cclosure_marshal_VOID__DOUBLE", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__ENUM
// container is nil
func CclosureMarshalVoid_Enum(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(183, "cclosure_marshal_VOID__ENUM", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__FLAGS
// container is nil
func CclosureMarshalVoid_Flags(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(184, "cclosure_marshal_VOID__FLAGS", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__FLOAT
// container is nil
func CclosureMarshalVoid_Float(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(185, "cclosure_marshal_VOID__FLOAT", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__INT
// container is nil
func CclosureMarshalVoid_Int(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(186, "cclosure_marshal_VOID__INT", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__LONG
// container is nil
func CclosureMarshalVoid_Long(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(187, "cclosure_marshal_VOID__LONG", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__OBJECT
// container is nil
func CclosureMarshalVoid_Object(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(188, "cclosure_marshal_VOID__OBJECT", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__PARAM
// container is nil
func CclosureMarshalVoid_Param(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(189, "cclosure_marshal_VOID__PARAM", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__POINTER
// container is nil
func CclosureMarshalVoid_Pointer(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(190, "cclosure_marshal_VOID__POINTER", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__STRING
// container is nil
func CclosureMarshalVoid_String(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(191, "cclosure_marshal_VOID__STRING", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__UCHAR
// container is nil
func CclosureMarshalVoid_Uchar(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(192, "cclosure_marshal_VOID__UCHAR", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__UINT
// container is nil
func CclosureMarshalVoid_Uint(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(193, "cclosure_marshal_VOID__UINT", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__UINT_POINTER
// container is nil
func CclosureMarshalVoid_UintPointer(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(194, "cclosure_marshal_VOID__UINT_POINTER", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__ULONG
// container is nil
func CclosureMarshalVoid_Ulong(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(195, "cclosure_marshal_VOID__ULONG", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__VARIANT
// container is nil
func CclosureMarshalVoid_Variant(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(196, "cclosure_marshal_VOID__VARIANT", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_VOID__VOID
// container is nil
func CclosureMarshalVoid_Void(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(197, "cclosure_marshal_VOID__VOID", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_value, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_cclosure_marshal_generic
// container is nil
func CclosureMarshalGeneric(closure Closure, return_gvalue Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get(198, "cclosure_marshal_generic", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_return_gvalue := gi.NewPointerArgument(return_gvalue.P)
	arg_n_param_values := gi.NewUint32Argument(n_param_values)
	arg_param_values := gi.NewPointerArgument(param_values.P)
	arg_invocation_hint := gi.NewPointerArgument(invocation_hint)
	arg_marshal_data := gi.NewPointerArgument(marshal_data)
	args := []gi.Argument{arg_closure, arg_return_gvalue, arg_n_param_values, arg_param_values, arg_invocation_hint, arg_marshal_data}
	iv.Call(args, nil, nil)
}

// g_enum_complete_type_info
// container is nil
func EnumCompleteTypeInfo(g_enum_type gi.GType, const_values EnumValue) (info int /*TODO_TYPE tag: ifc, biType: struct*/) {
	iv, err := _I.Get(199, "enum_complete_type_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_g_enum_type := gi.NewUintArgument(uint(g_enum_type))
	arg_info := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_const_values := gi.NewPointerArgument(const_values.P)
	args := []gi.Argument{arg_g_enum_type, arg_info, arg_const_values}
	iv.Call(args, nil, &outArgs[0])
	info = outArgs[0].Int() /*TODO*/
	return
}

// g_enum_get_value
// container is nil
func EnumGetValue(enum_class EnumClass, value int32) (result EnumValue) {
	iv, err := _I.Get(200, "enum_get_value", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_enum_class := gi.NewPointerArgument(enum_class.P)
	arg_value := gi.NewInt32Argument(value)
	args := []gi.Argument{arg_enum_class, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_enum_get_value_by_name
// container is nil
func EnumGetValueByName(enum_class EnumClass, name string) (result EnumValue) {
	iv, err := _I.Get(201, "enum_get_value_by_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_enum_class := gi.NewPointerArgument(enum_class.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_enum_class, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// g_enum_get_value_by_nick
// container is nil
func EnumGetValueByNick(enum_class EnumClass, nick string) (result EnumValue) {
	iv, err := _I.Get(202, "enum_get_value_by_nick", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_nick := gi.CString(nick)
	arg_enum_class := gi.NewPointerArgument(enum_class.P)
	arg_nick := gi.NewStringArgument(c_nick)
	args := []gi.Argument{arg_enum_class, arg_nick}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_nick)
	result.P = ret.Pointer()
	return
}

// g_enum_register_static
// container is nil
func EnumRegisterStatic(name string, const_static_values EnumValue) (result gi.GType) {
	iv, err := _I.Get(203, "enum_register_static", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	arg_const_static_values := gi.NewPointerArgument(const_static_values.P)
	args := []gi.Argument{arg_name, arg_const_static_values}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = gi.GType(ret.Uint())
	return
}

// g_enum_to_string
// container is nil
func EnumToString(g_enum_type gi.GType, value int32) (result string) {
	iv, err := _I.Get(204, "enum_to_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_g_enum_type := gi.NewUintArgument(uint(g_enum_type))
	arg_value := gi.NewInt32Argument(value)
	args := []gi.Argument{arg_g_enum_type, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_flags_complete_type_info
// container is nil
func FlagsCompleteTypeInfo(g_flags_type gi.GType, const_values FlagsValue) (info int /*TODO_TYPE tag: ifc, biType: struct*/) {
	iv, err := _I.Get(205, "flags_complete_type_info", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_g_flags_type := gi.NewUintArgument(uint(g_flags_type))
	arg_info := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_const_values := gi.NewPointerArgument(const_values.P)
	args := []gi.Argument{arg_g_flags_type, arg_info, arg_const_values}
	iv.Call(args, nil, &outArgs[0])
	info = outArgs[0].Int() /*TODO*/
	return
}

// g_flags_get_first_value
// container is nil
func FlagsGetFirstValue(flags_class FlagsClass, value uint32) (result FlagsValue) {
	iv, err := _I.Get(206, "flags_get_first_value", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_flags_class := gi.NewPointerArgument(flags_class.P)
	arg_value := gi.NewUint32Argument(value)
	args := []gi.Argument{arg_flags_class, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_flags_get_value_by_name
// container is nil
func FlagsGetValueByName(flags_class FlagsClass, name string) (result FlagsValue) {
	iv, err := _I.Get(207, "flags_get_value_by_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_flags_class := gi.NewPointerArgument(flags_class.P)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_flags_class, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// g_flags_get_value_by_nick
// container is nil
func FlagsGetValueByNick(flags_class FlagsClass, nick string) (result FlagsValue) {
	iv, err := _I.Get(208, "flags_get_value_by_nick", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_nick := gi.CString(nick)
	arg_flags_class := gi.NewPointerArgument(flags_class.P)
	arg_nick := gi.NewStringArgument(c_nick)
	args := []gi.Argument{arg_flags_class, arg_nick}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_nick)
	result.P = ret.Pointer()
	return
}

// g_flags_register_static
// container is nil
func FlagsRegisterStatic(name string, const_static_values FlagsValue) (result gi.GType) {
	iv, err := _I.Get(209, "flags_register_static", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	arg_const_static_values := gi.NewPointerArgument(const_static_values.P)
	args := []gi.Argument{arg_name, arg_const_static_values}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = gi.GType(ret.Uint())
	return
}

// g_flags_to_string
// container is nil
func FlagsToString(flags_type gi.GType, value uint32) (result string) {
	iv, err := _I.Get(210, "flags_to_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_flags_type := gi.NewUintArgument(uint(flags_type))
	arg_value := gi.NewUint32Argument(value)
	args := []gi.Argument{arg_flags_type, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_gtype_get_type
// container is nil
func GtypeGetType() (result gi.GType) {
	iv, err := _I.Get(211, "gtype_get_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// g_param_spec_boolean
// container is nil
func ParamSpecBooleanF(name string, nick string, blurb string, default_value bool, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(212, "param_spec_boolean", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_default_value := gi.NewBoolArgument(default_value)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_boxed
// container is nil
func ParamSpecBoxedF(name string, nick string, blurb string, boxed_type gi.GType, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(213, "param_spec_boxed", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_boxed_type := gi.NewUintArgument(uint(boxed_type))
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_boxed_type, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_char
// container is nil
func ParamSpecCharF(name string, nick string, blurb string, minimum int8, maximum int8, default_value int8, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(214, "param_spec_char", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_minimum := gi.NewInt8Argument(minimum)
	arg_maximum := gi.NewInt8Argument(maximum)
	arg_default_value := gi.NewInt8Argument(default_value)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_minimum, arg_maximum, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_double
// container is nil
func ParamSpecDoubleF(name string, nick string, blurb string, minimum float64, maximum float64, default_value float64, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(215, "param_spec_double", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_minimum := gi.NewDoubleArgument(minimum)
	arg_maximum := gi.NewDoubleArgument(maximum)
	arg_default_value := gi.NewDoubleArgument(default_value)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_minimum, arg_maximum, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_enum
// container is nil
func ParamSpecEnumF(name string, nick string, blurb string, enum_type gi.GType, default_value int32, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(216, "param_spec_enum", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_enum_type := gi.NewUintArgument(uint(enum_type))
	arg_default_value := gi.NewInt32Argument(default_value)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_enum_type, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_flags
// container is nil
func ParamSpecFlagsF(name string, nick string, blurb string, flags_type gi.GType, default_value uint32, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(217, "param_spec_flags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_flags_type := gi.NewUintArgument(uint(flags_type))
	arg_default_value := gi.NewUint32Argument(default_value)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_flags_type, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_float
// container is nil
func ParamSpecFloatF(name string, nick string, blurb string, minimum float32, maximum float32, default_value float32, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(218, "param_spec_float", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_minimum := gi.NewFloatArgument(minimum)
	arg_maximum := gi.NewFloatArgument(maximum)
	arg_default_value := gi.NewFloatArgument(default_value)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_minimum, arg_maximum, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_gtype
// container is nil
func ParamSpecGtype(name string, nick string, blurb string, is_a_type gi.GType, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(219, "param_spec_gtype", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_is_a_type := gi.NewUintArgument(uint(is_a_type))
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_is_a_type, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_int
// container is nil
func ParamSpecIntF(name string, nick string, blurb string, minimum int32, maximum int32, default_value int32, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(220, "param_spec_int", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_minimum := gi.NewInt32Argument(minimum)
	arg_maximum := gi.NewInt32Argument(maximum)
	arg_default_value := gi.NewInt32Argument(default_value)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_minimum, arg_maximum, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_int64
// container is nil
func ParamSpecInt64F(name string, nick string, blurb string, minimum int64, maximum int64, default_value int64, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(221, "param_spec_int64", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_minimum := gi.NewInt64Argument(minimum)
	arg_maximum := gi.NewInt64Argument(maximum)
	arg_default_value := gi.NewInt64Argument(default_value)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_minimum, arg_maximum, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_long
// container is nil
func ParamSpecLongF(name string, nick string, blurb string, minimum int64, maximum int64, default_value int64, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(222, "param_spec_long", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_minimum := gi.NewInt64Argument(minimum)
	arg_maximum := gi.NewInt64Argument(maximum)
	arg_default_value := gi.NewInt64Argument(default_value)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_minimum, arg_maximum, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_object
// container is nil
func ParamSpecObjectF(name string, nick string, blurb string, object_type gi.GType, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(223, "param_spec_object", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_object_type := gi.NewUintArgument(uint(object_type))
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_object_type, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_param
// container is nil
func ParamSpecParamF(name string, nick string, blurb string, param_type gi.GType, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(224, "param_spec_param", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_param_type := gi.NewUintArgument(uint(param_type))
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_param_type, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_pointer
// container is nil
func ParamSpecPointerF(name string, nick string, blurb string, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(225, "param_spec_pointer", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_pool_new
// container is nil
func ParamSpecPoolNew(type_prefixing bool) (result ParamSpecPool) {
	iv, err := _I.Get(226, "param_spec_pool_new", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type_prefixing := gi.NewBoolArgument(type_prefixing)
	args := []gi.Argument{arg_type_prefixing}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_param_spec_string
// container is nil
func ParamSpecStringF(name string, nick string, blurb string, default_value string, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(227, "param_spec_string", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	c_default_value := gi.CString(default_value)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_default_value := gi.NewStringArgument(c_default_value)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	gi.Free(c_default_value)
	result.P = ret.Pointer()
	return
}

// g_param_spec_uchar
// container is nil
func ParamSpecUchar(name string, nick string, blurb string, minimum uint8, maximum uint8, default_value uint8, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(228, "param_spec_uchar", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_minimum := gi.NewUint8Argument(minimum)
	arg_maximum := gi.NewUint8Argument(maximum)
	arg_default_value := gi.NewUint8Argument(default_value)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_minimum, arg_maximum, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_uint
// container is nil
func ParamSpecUint(name string, nick string, blurb string, minimum uint32, maximum uint32, default_value uint32, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(229, "param_spec_uint", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_minimum := gi.NewUint32Argument(minimum)
	arg_maximum := gi.NewUint32Argument(maximum)
	arg_default_value := gi.NewUint32Argument(default_value)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_minimum, arg_maximum, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_uint64
// container is nil
func ParamSpecUint64(name string, nick string, blurb string, minimum uint64, maximum uint64, default_value uint64, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(230, "param_spec_uint64", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_minimum := gi.NewUint64Argument(minimum)
	arg_maximum := gi.NewUint64Argument(maximum)
	arg_default_value := gi.NewUint64Argument(default_value)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_minimum, arg_maximum, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_ulong
// container is nil
func ParamSpecUlong(name string, nick string, blurb string, minimum uint64, maximum uint64, default_value uint64, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(231, "param_spec_ulong", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_minimum := gi.NewUint64Argument(minimum)
	arg_maximum := gi.NewUint64Argument(maximum)
	arg_default_value := gi.NewUint64Argument(default_value)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_minimum, arg_maximum, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_unichar
// container is nil
func ParamSpecUnicharF(name string, nick string, blurb string, default_value rune, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(232, "param_spec_unichar", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_default_value := gi.NewUint32Argument(uint32(default_value))
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_spec_variant
// container is nil
func ParamSpecVariantF(name string, nick string, blurb string, type1 glib.VariantType, default_value glib.Variant, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get(233, "param_spec_variant", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_nick := gi.CString(nick)
	c_blurb := gi.CString(blurb)
	arg_name := gi.NewStringArgument(c_name)
	arg_nick := gi.NewStringArgument(c_nick)
	arg_blurb := gi.NewStringArgument(c_blurb)
	arg_type1 := gi.NewPointerArgument(type1.P)
	arg_default_value := gi.NewPointerArgument(default_value.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_name, arg_nick, arg_blurb, arg_type1, arg_default_value, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_nick)
	gi.Free(c_blurb)
	result.P = ret.Pointer()
	return
}

// g_param_type_register_static
// container is nil
func ParamTypeRegisterStatic(name string, pspec_info ParamSpecTypeInfo) (result gi.GType) {
	iv, err := _I.Get(234, "param_type_register_static", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	arg_pspec_info := gi.NewPointerArgument(pspec_info.P)
	args := []gi.Argument{arg_name, arg_pspec_info}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = gi.GType(ret.Uint())
	return
}

// g_param_value_convert
// container is nil
func ParamValueConvert(pspec IParamSpec, src_value Value, dest_value Value, strict_validation bool) (result bool) {
	iv, err := _I.Get(235, "param_value_convert", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pspec != nil {
		tmp = pspec.P_ParamSpec()
	}
	arg_pspec := gi.NewPointerArgument(tmp)
	arg_src_value := gi.NewPointerArgument(src_value.P)
	arg_dest_value := gi.NewPointerArgument(dest_value.P)
	arg_strict_validation := gi.NewBoolArgument(strict_validation)
	args := []gi.Argument{arg_pspec, arg_src_value, arg_dest_value, arg_strict_validation}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_param_value_defaults
// container is nil
func ParamValueDefaults(pspec IParamSpec, value Value) (result bool) {
	iv, err := _I.Get(236, "param_value_defaults", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pspec != nil {
		tmp = pspec.P_ParamSpec()
	}
	arg_pspec := gi.NewPointerArgument(tmp)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_pspec, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_param_value_set_default
// container is nil
func ParamValueSetDefault(pspec IParamSpec, value Value) {
	iv, err := _I.Get(237, "param_value_set_default", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pspec != nil {
		tmp = pspec.P_ParamSpec()
	}
	arg_pspec := gi.NewPointerArgument(tmp)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_pspec, arg_value}
	iv.Call(args, nil, nil)
}

// g_param_value_validate
// container is nil
func ParamValueValidate(pspec IParamSpec, value Value) (result bool) {
	iv, err := _I.Get(238, "param_value_validate", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pspec != nil {
		tmp = pspec.P_ParamSpec()
	}
	arg_pspec := gi.NewPointerArgument(tmp)
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_pspec, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_param_values_cmp
// container is nil
func ParamValuesCmp(pspec IParamSpec, value1 Value, value2 Value) (result int32) {
	iv, err := _I.Get(239, "param_values_cmp", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if pspec != nil {
		tmp = pspec.P_ParamSpec()
	}
	arg_pspec := gi.NewPointerArgument(tmp)
	arg_value1 := gi.NewPointerArgument(value1.P)
	arg_value2 := gi.NewPointerArgument(value2.P)
	args := []gi.Argument{arg_pspec, arg_value1, arg_value2}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_pointer_type_register_static
// container is nil
func PointerTypeRegisterStatic(name string) (result gi.GType) {
	iv, err := _I.Get(240, "pointer_type_register_static", "")
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
	result = gi.GType(ret.Uint())
	return
}

// g_signal_accumulator_first_wins
// container is nil
func SignalAccumulatorFirstWins(ihint SignalInvocationHint, return_accu Value, handler_return Value, dummy unsafe.Pointer) (result bool) {
	iv, err := _I.Get(241, "signal_accumulator_first_wins", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ihint := gi.NewPointerArgument(ihint.P)
	arg_return_accu := gi.NewPointerArgument(return_accu.P)
	arg_handler_return := gi.NewPointerArgument(handler_return.P)
	arg_dummy := gi.NewPointerArgument(dummy)
	args := []gi.Argument{arg_ihint, arg_return_accu, arg_handler_return, arg_dummy}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_signal_accumulator_true_handled
// container is nil
func SignalAccumulatorTrueHandled(ihint SignalInvocationHint, return_accu Value, handler_return Value, dummy unsafe.Pointer) (result bool) {
	iv, err := _I.Get(242, "signal_accumulator_true_handled", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_ihint := gi.NewPointerArgument(ihint.P)
	arg_return_accu := gi.NewPointerArgument(return_accu.P)
	arg_handler_return := gi.NewPointerArgument(handler_return.P)
	arg_dummy := gi.NewPointerArgument(dummy)
	args := []gi.Argument{arg_ihint, arg_return_accu, arg_handler_return, arg_dummy}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_signal_add_emission_hook
// container is nil
func SignalAddEmissionHook(signal_id uint32, detail uint32, hook_func int /*TODO_TYPE CALLBACK*/, hook_data unsafe.Pointer, data_destroy int /*TODO_TYPE CALLBACK*/) (result uint64) {
	iv, err := _I.Get(243, "signal_add_emission_hook", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_signal_id := gi.NewUint32Argument(signal_id)
	arg_detail := gi.NewUint32Argument(detail)
	arg_hook_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySignalEmissionHook()))
	arg_hook_data := gi.NewPointerArgument(hook_data)
	arg_data_destroy := gi.NewPointerArgument(unsafe.Pointer(glib.GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_signal_id, arg_detail, arg_hook_func, arg_hook_data, arg_data_destroy}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_signal_chain_from_overridden
// container is nil
func SignalChainFromOverridden(instance_and_params unsafe.Pointer, return_value Value) {
	iv, err := _I.Get(244, "signal_chain_from_overridden", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_instance_and_params := gi.NewPointerArgument(instance_and_params)
	arg_return_value := gi.NewPointerArgument(return_value.P)
	args := []gi.Argument{arg_instance_and_params, arg_return_value}
	iv.Call(args, nil, nil)
}

// g_signal_connect_closure
// container is nil
func SignalConnectClosure(instance IObject, detailed_signal string, closure Closure, after bool) (result uint64) {
	iv, err := _I.Get(245, "signal_connect_closure", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if instance != nil {
		tmp = instance.P_Object()
	}
	c_detailed_signal := gi.CString(detailed_signal)
	arg_instance := gi.NewPointerArgument(tmp)
	arg_detailed_signal := gi.NewStringArgument(c_detailed_signal)
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_after := gi.NewBoolArgument(after)
	args := []gi.Argument{arg_instance, arg_detailed_signal, arg_closure, arg_after}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_detailed_signal)
	result = ret.Uint64()
	return
}

// g_signal_connect_closure_by_id
// container is nil
func SignalConnectClosureById(instance IObject, signal_id uint32, detail uint32, closure Closure, after bool) (result uint64) {
	iv, err := _I.Get(246, "signal_connect_closure_by_id", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if instance != nil {
		tmp = instance.P_Object()
	}
	arg_instance := gi.NewPointerArgument(tmp)
	arg_signal_id := gi.NewUint32Argument(signal_id)
	arg_detail := gi.NewUint32Argument(detail)
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_after := gi.NewBoolArgument(after)
	args := []gi.Argument{arg_instance, arg_signal_id, arg_detail, arg_closure, arg_after}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_signal_emitv
// container is nil
func SignalEmitv(instance_and_params unsafe.Pointer, signal_id uint32, detail uint32, return_value int /*TODO:TYPE*/) {
	iv, err := _I.Get(247, "signal_emitv", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_instance_and_params := gi.NewPointerArgument(instance_and_params)
	arg_signal_id := gi.NewUint32Argument(signal_id)
	arg_detail := gi.NewUint32Argument(detail)
	args := []gi.Argument{arg_instance_and_params, arg_signal_id, arg_detail}
	iv.Call(args, nil, &outArgs[0])
}

// g_signal_get_invocation_hint
// container is nil
func SignalGetInvocationHint(instance IObject) (result SignalInvocationHint) {
	iv, err := _I.Get(248, "signal_get_invocation_hint", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if instance != nil {
		tmp = instance.P_Object()
	}
	arg_instance := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_instance}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_signal_handler_block
// container is nil
func SignalHandlerBlock(instance IObject, handler_id uint64) {
	iv, err := _I.Get(249, "signal_handler_block", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if instance != nil {
		tmp = instance.P_Object()
	}
	arg_instance := gi.NewPointerArgument(tmp)
	arg_handler_id := gi.NewUint64Argument(handler_id)
	args := []gi.Argument{arg_instance, arg_handler_id}
	iv.Call(args, nil, nil)
}

// g_signal_handler_disconnect
// container is nil
func SignalHandlerDisconnect(instance IObject, handler_id uint64) {
	iv, err := _I.Get(250, "signal_handler_disconnect", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if instance != nil {
		tmp = instance.P_Object()
	}
	arg_instance := gi.NewPointerArgument(tmp)
	arg_handler_id := gi.NewUint64Argument(handler_id)
	args := []gi.Argument{arg_instance, arg_handler_id}
	iv.Call(args, nil, nil)
}

// g_signal_handler_find
// container is nil
func SignalHandlerFind(instance IObject, mask SignalMatchTypeFlags, signal_id uint32, detail uint32, closure Closure, func1 unsafe.Pointer, data unsafe.Pointer) (result uint64) {
	iv, err := _I.Get(251, "signal_handler_find", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if instance != nil {
		tmp = instance.P_Object()
	}
	arg_instance := gi.NewPointerArgument(tmp)
	arg_mask := gi.NewIntArgument(int(mask))
	arg_signal_id := gi.NewUint32Argument(signal_id)
	arg_detail := gi.NewUint32Argument(detail)
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_func1 := gi.NewPointerArgument(func1)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_instance, arg_mask, arg_signal_id, arg_detail, arg_closure, arg_func1, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_signal_handler_is_connected
// container is nil
func SignalHandlerIsConnected(instance IObject, handler_id uint64) (result bool) {
	iv, err := _I.Get(252, "signal_handler_is_connected", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if instance != nil {
		tmp = instance.P_Object()
	}
	arg_instance := gi.NewPointerArgument(tmp)
	arg_handler_id := gi.NewUint64Argument(handler_id)
	args := []gi.Argument{arg_instance, arg_handler_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_signal_handler_unblock
// container is nil
func SignalHandlerUnblock(instance IObject, handler_id uint64) {
	iv, err := _I.Get(253, "signal_handler_unblock", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if instance != nil {
		tmp = instance.P_Object()
	}
	arg_instance := gi.NewPointerArgument(tmp)
	arg_handler_id := gi.NewUint64Argument(handler_id)
	args := []gi.Argument{arg_instance, arg_handler_id}
	iv.Call(args, nil, nil)
}

// g_signal_handlers_block_matched
// container is nil
func SignalHandlersBlockMatched(instance IObject, mask SignalMatchTypeFlags, signal_id uint32, detail uint32, closure Closure, func1 unsafe.Pointer, data unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(254, "signal_handlers_block_matched", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if instance != nil {
		tmp = instance.P_Object()
	}
	arg_instance := gi.NewPointerArgument(tmp)
	arg_mask := gi.NewIntArgument(int(mask))
	arg_signal_id := gi.NewUint32Argument(signal_id)
	arg_detail := gi.NewUint32Argument(detail)
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_func1 := gi.NewPointerArgument(func1)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_instance, arg_mask, arg_signal_id, arg_detail, arg_closure, arg_func1, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_signal_handlers_destroy
// container is nil
func SignalHandlersDestroy(instance IObject) {
	iv, err := _I.Get(255, "signal_handlers_destroy", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if instance != nil {
		tmp = instance.P_Object()
	}
	arg_instance := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_instance}
	iv.Call(args, nil, nil)
}

// g_signal_handlers_disconnect_matched
// container is nil
func SignalHandlersDisconnectMatched(instance IObject, mask SignalMatchTypeFlags, signal_id uint32, detail uint32, closure Closure, func1 unsafe.Pointer, data unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(256, "signal_handlers_disconnect_matched", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if instance != nil {
		tmp = instance.P_Object()
	}
	arg_instance := gi.NewPointerArgument(tmp)
	arg_mask := gi.NewIntArgument(int(mask))
	arg_signal_id := gi.NewUint32Argument(signal_id)
	arg_detail := gi.NewUint32Argument(detail)
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_func1 := gi.NewPointerArgument(func1)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_instance, arg_mask, arg_signal_id, arg_detail, arg_closure, arg_func1, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_signal_handlers_unblock_matched
// container is nil
func SignalHandlersUnblockMatched(instance IObject, mask SignalMatchTypeFlags, signal_id uint32, detail uint32, closure Closure, func1 unsafe.Pointer, data unsafe.Pointer) (result uint32) {
	iv, err := _I.Get(257, "signal_handlers_unblock_matched", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if instance != nil {
		tmp = instance.P_Object()
	}
	arg_instance := gi.NewPointerArgument(tmp)
	arg_mask := gi.NewIntArgument(int(mask))
	arg_signal_id := gi.NewUint32Argument(signal_id)
	arg_detail := gi.NewUint32Argument(detail)
	arg_closure := gi.NewPointerArgument(closure.P)
	arg_func1 := gi.NewPointerArgument(func1)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_instance, arg_mask, arg_signal_id, arg_detail, arg_closure, arg_func1, arg_data}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_signal_has_handler_pending
// container is nil
func SignalHasHandlerPending(instance IObject, signal_id uint32, detail uint32, may_be_blocked bool) (result bool) {
	iv, err := _I.Get(258, "signal_has_handler_pending", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if instance != nil {
		tmp = instance.P_Object()
	}
	arg_instance := gi.NewPointerArgument(tmp)
	arg_signal_id := gi.NewUint32Argument(signal_id)
	arg_detail := gi.NewUint32Argument(detail)
	arg_may_be_blocked := gi.NewBoolArgument(may_be_blocked)
	args := []gi.Argument{arg_instance, arg_signal_id, arg_detail, arg_may_be_blocked}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_signal_list_ids
// container is nil
// ret lenArgIdx 1
func SignalListIds(itype gi.GType) (result gi.Uint32Array) {
	iv, err := _I.Get(259, "signal_list_ids", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_itype := gi.NewUintArgument(uint(itype))
	arg_n_ids := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_itype, arg_n_ids}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_ids uint32
	_ = n_ids
	n_ids = outArgs[0].Uint32()
	result = gi.Uint32Array{P: ret.Pointer(), Len: int(n_ids)}
	return
}

// g_signal_lookup
// container is nil
func SignalLookup(name string, itype gi.GType) (result uint32) {
	iv, err := _I.Get(260, "signal_lookup", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	arg_name := gi.NewStringArgument(c_name)
	arg_itype := gi.NewUintArgument(uint(itype))
	args := []gi.Argument{arg_name, arg_itype}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	result = ret.Uint32()
	return
}

// g_signal_name
// container is nil
func SignalName(signal_id uint32) (result string) {
	iv, err := _I.Get(261, "signal_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_signal_id := gi.NewUint32Argument(signal_id)
	args := []gi.Argument{arg_signal_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_signal_override_class_closure
// container is nil
func SignalOverrideClassClosure(signal_id uint32, instance_type gi.GType, class_closure Closure) {
	iv, err := _I.Get(262, "signal_override_class_closure", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_signal_id := gi.NewUint32Argument(signal_id)
	arg_instance_type := gi.NewUintArgument(uint(instance_type))
	arg_class_closure := gi.NewPointerArgument(class_closure.P)
	args := []gi.Argument{arg_signal_id, arg_instance_type, arg_class_closure}
	iv.Call(args, nil, nil)
}

// g_signal_parse_name
// container is nil
func SignalParseName(detailed_signal string, itype gi.GType, force_detail_quark bool) (result bool, signal_id_p uint32, detail_p uint32) {
	iv, err := _I.Get(263, "signal_parse_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [2]gi.Argument
	c_detailed_signal := gi.CString(detailed_signal)
	arg_detailed_signal := gi.NewStringArgument(c_detailed_signal)
	arg_itype := gi.NewUintArgument(uint(itype))
	arg_signal_id_p := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	arg_detail_p := gi.NewPointerArgument(unsafe.Pointer(&outArgs[1]))
	arg_force_detail_quark := gi.NewBoolArgument(force_detail_quark)
	args := []gi.Argument{arg_detailed_signal, arg_itype, arg_signal_id_p, arg_detail_p, arg_force_detail_quark}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	gi.Free(c_detailed_signal)
	signal_id_p = outArgs[0].Uint32()
	detail_p = outArgs[1].Uint32()
	result = ret.Bool()
	return
}

// g_signal_query
// container is nil
func SignalQueryF(signal_id uint32, query SignalQuery) {
	iv, err := _I.Get(264, "signal_query", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_signal_id := gi.NewUint32Argument(signal_id)
	arg_query := gi.NewPointerArgument(query.P)
	args := []gi.Argument{arg_signal_id, arg_query}
	iv.Call(args, nil, nil)
}

// g_signal_remove_emission_hook
// container is nil
func SignalRemoveEmissionHook(signal_id uint32, hook_id uint64) {
	iv, err := _I.Get(265, "signal_remove_emission_hook", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_signal_id := gi.NewUint32Argument(signal_id)
	arg_hook_id := gi.NewUint64Argument(hook_id)
	args := []gi.Argument{arg_signal_id, arg_hook_id}
	iv.Call(args, nil, nil)
}

// g_signal_set_va_marshaller
// container is nil
func SignalSetVaMarshaller(signal_id uint32, instance_type gi.GType, va_marshaller int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(266, "signal_set_va_marshaller", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_signal_id := gi.NewUint32Argument(signal_id)
	arg_instance_type := gi.NewUintArgument(uint(instance_type))
	arg_va_marshaller := gi.NewIntArgument(va_marshaller) /*TODO*/
	args := []gi.Argument{arg_signal_id, arg_instance_type, arg_va_marshaller}
	iv.Call(args, nil, nil)
}

// g_signal_stop_emission
// container is nil
func SignalStopEmission(instance IObject, signal_id uint32, detail uint32) {
	iv, err := _I.Get(267, "signal_stop_emission", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if instance != nil {
		tmp = instance.P_Object()
	}
	arg_instance := gi.NewPointerArgument(tmp)
	arg_signal_id := gi.NewUint32Argument(signal_id)
	arg_detail := gi.NewUint32Argument(detail)
	args := []gi.Argument{arg_instance, arg_signal_id, arg_detail}
	iv.Call(args, nil, nil)
}

// g_signal_stop_emission_by_name
// container is nil
func SignalStopEmissionByName(instance IObject, detailed_signal string) {
	iv, err := _I.Get(268, "signal_stop_emission_by_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if instance != nil {
		tmp = instance.P_Object()
	}
	c_detailed_signal := gi.CString(detailed_signal)
	arg_instance := gi.NewPointerArgument(tmp)
	arg_detailed_signal := gi.NewStringArgument(c_detailed_signal)
	args := []gi.Argument{arg_instance, arg_detailed_signal}
	iv.Call(args, nil, nil)
	gi.Free(c_detailed_signal)
}

// g_signal_type_cclosure_new
// container is nil
func SignalTypeCclosureNew(itype gi.GType, struct_offset uint32) (result Closure) {
	iv, err := _I.Get(269, "signal_type_cclosure_new", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_itype := gi.NewUintArgument(uint(itype))
	arg_struct_offset := gi.NewUint32Argument(struct_offset)
	args := []gi.Argument{arg_itype, arg_struct_offset}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_source_set_closure
// container is nil
func SourceSetClosure(source glib.Source, closure Closure) {
	iv, err := _I.Get(270, "source_set_closure", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_source := gi.NewPointerArgument(source.P)
	arg_closure := gi.NewPointerArgument(closure.P)
	args := []gi.Argument{arg_source, arg_closure}
	iv.Call(args, nil, nil)
}

// g_source_set_dummy_callback
// container is nil
func SourceSetDummyCallback(source glib.Source) {
	iv, err := _I.Get(271, "source_set_dummy_callback", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_source := gi.NewPointerArgument(source.P)
	args := []gi.Argument{arg_source}
	iv.Call(args, nil, nil)
}

// g_strdup_value_contents
// container is nil
func StrdupValueContents(value Value) (result string) {
	iv, err := _I.Get(272, "strdup_value_contents", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_type_add_class_private
// container is nil
func TypeAddClassPrivate(class_type gi.GType, private_size uint64) {
	iv, err := _I.Get(273, "type_add_class_private", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_class_type := gi.NewUintArgument(uint(class_type))
	arg_private_size := gi.NewUint64Argument(private_size)
	args := []gi.Argument{arg_class_type, arg_private_size}
	iv.Call(args, nil, nil)
}

// g_type_add_instance_private
// container is nil
func TypeAddInstancePrivate(class_type gi.GType, private_size uint64) (result int32) {
	iv, err := _I.Get(274, "type_add_instance_private", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_class_type := gi.NewUintArgument(uint(class_type))
	arg_private_size := gi.NewUint64Argument(private_size)
	args := []gi.Argument{arg_class_type, arg_private_size}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_type_add_interface_dynamic
// container is nil
func TypeAddInterfaceDynamic(instance_type gi.GType, interface_type gi.GType, plugin ITypePlugin) {
	iv, err := _I.Get(275, "type_add_interface_dynamic", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if plugin != nil {
		tmp = plugin.P_TypePlugin()
	}
	arg_instance_type := gi.NewUintArgument(uint(instance_type))
	arg_interface_type := gi.NewUintArgument(uint(interface_type))
	arg_plugin := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_instance_type, arg_interface_type, arg_plugin}
	iv.Call(args, nil, nil)
}

// g_type_add_interface_static
// container is nil
func TypeAddInterfaceStatic(instance_type gi.GType, interface_type gi.GType, info InterfaceInfo) {
	iv, err := _I.Get(276, "type_add_interface_static", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_instance_type := gi.NewUintArgument(uint(instance_type))
	arg_interface_type := gi.NewUintArgument(uint(interface_type))
	arg_info := gi.NewPointerArgument(info.P)
	args := []gi.Argument{arg_instance_type, arg_interface_type, arg_info}
	iv.Call(args, nil, nil)
}

// g_type_check_class_is_a
// container is nil
func TypeCheckClassIsA(g_class TypeClass, is_a_type gi.GType) (result bool) {
	iv, err := _I.Get(277, "type_check_class_is_a", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_g_class := gi.NewPointerArgument(g_class.P)
	arg_is_a_type := gi.NewUintArgument(uint(is_a_type))
	args := []gi.Argument{arg_g_class, arg_is_a_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_type_check_instance
// container is nil
func TypeCheckInstance(instance TypeInstance) (result bool) {
	iv, err := _I.Get(278, "type_check_instance", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_instance := gi.NewPointerArgument(instance.P)
	args := []gi.Argument{arg_instance}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_type_check_instance_is_a
// container is nil
func TypeCheckInstanceIsA(instance TypeInstance, iface_type gi.GType) (result bool) {
	iv, err := _I.Get(279, "type_check_instance_is_a", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_instance := gi.NewPointerArgument(instance.P)
	arg_iface_type := gi.NewUintArgument(uint(iface_type))
	args := []gi.Argument{arg_instance, arg_iface_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_type_check_instance_is_fundamentally_a
// container is nil
func TypeCheckInstanceIsFundamentallyA(instance TypeInstance, fundamental_type gi.GType) (result bool) {
	iv, err := _I.Get(280, "type_check_instance_is_fundamentally_a", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_instance := gi.NewPointerArgument(instance.P)
	arg_fundamental_type := gi.NewUintArgument(uint(fundamental_type))
	args := []gi.Argument{arg_instance, arg_fundamental_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_type_check_is_value_type
// container is nil
func TypeCheckIsValueType(type1 gi.GType) (result bool) {
	iv, err := _I.Get(281, "type_check_is_value_type", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_type_check_value
// container is nil
func TypeCheckValue(value Value) (result bool) {
	iv, err := _I.Get(282, "type_check_value", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	args := []gi.Argument{arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_type_check_value_holds
// container is nil
func TypeCheckValueHolds(value Value, type1 gi.GType) (result bool) {
	iv, err := _I.Get(283, "type_check_value_holds", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_value := gi.NewPointerArgument(value.P)
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_value, arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_type_children
// container is nil
// ret lenArgIdx 1
func TypeChildren(type1 gi.GType) (result gi.GTypeArray) {
	iv, err := _I.Get(284, "type_children", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_type1 := gi.NewUintArgument(uint(type1))
	arg_n_children := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_type1, arg_n_children}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_children uint32
	_ = n_children
	n_children = outArgs[0].Uint32()
	result = gi.GTypeArray{P: ret.Pointer(), Len: int(n_children)}
	return
}

// g_type_class_adjust_private_offset
// container is nil
func TypeClassAdjustPrivateOffset(g_class unsafe.Pointer, private_size_or_offset int32) {
	iv, err := _I.Get(285, "type_class_adjust_private_offset", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_g_class := gi.NewPointerArgument(g_class)
	arg_private_size_or_offset := gi.NewInt32Argument(private_size_or_offset)
	args := []gi.Argument{arg_g_class, arg_private_size_or_offset}
	iv.Call(args, nil, nil)
}

// g_type_class_peek
// container is nil
func TypeClassPeek(type1 gi.GType) (result TypeClass) {
	iv, err := _I.Get(286, "type_class_peek", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_class_peek_static
// container is nil
func TypeClassPeekStatic(type1 gi.GType) (result TypeClass) {
	iv, err := _I.Get(287, "type_class_peek_static", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_class_ref
// container is nil
func TypeClassRef(type1 gi.GType) (result TypeClass) {
	iv, err := _I.Get(288, "type_class_ref", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_default_interface_peek
// container is nil
func TypeDefaultInterfacePeek(g_type gi.GType) (result TypeInterface) {
	iv, err := _I.Get(289, "type_default_interface_peek", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_g_type := gi.NewUintArgument(uint(g_type))
	args := []gi.Argument{arg_g_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_default_interface_ref
// container is nil
func TypeDefaultInterfaceRef(g_type gi.GType) (result TypeInterface) {
	iv, err := _I.Get(290, "type_default_interface_ref", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_g_type := gi.NewUintArgument(uint(g_type))
	args := []gi.Argument{arg_g_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_default_interface_unref
// container is nil
func TypeDefaultInterfaceUnref(g_iface TypeInterface) {
	iv, err := _I.Get(291, "type_default_interface_unref", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_g_iface := gi.NewPointerArgument(g_iface.P)
	args := []gi.Argument{arg_g_iface}
	iv.Call(args, nil, nil)
}

// g_type_depth
// container is nil
func TypeDepth(type1 gi.GType) (result uint32) {
	iv, err := _I.Get(292, "type_depth", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_type_ensure
// container is nil
func TypeEnsure(type1 gi.GType) {
	iv, err := _I.Get(293, "type_ensure", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	iv.Call(args, nil, nil)
}

// g_type_free_instance
// container is nil
func TypeFreeInstance(instance TypeInstance) {
	iv, err := _I.Get(294, "type_free_instance", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_instance := gi.NewPointerArgument(instance.P)
	args := []gi.Argument{arg_instance}
	iv.Call(args, nil, nil)
}

// g_type_from_name
// container is nil
func TypeFromName(name string) (result gi.GType) {
	iv, err := _I.Get(295, "type_from_name", "")
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
	result = gi.GType(ret.Uint())
	return
}

// g_type_fundamental
// container is nil
func TypeFundamental(type_id gi.GType) (result gi.GType) {
	iv, err := _I.Get(296, "type_fundamental", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type_id := gi.NewUintArgument(uint(type_id))
	args := []gi.Argument{arg_type_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// g_type_fundamental_next
// container is nil
func TypeFundamentalNext() (result gi.GType) {
	iv, err := _I.Get(297, "type_fundamental_next", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// g_type_get_instance_count
// container is nil
func TypeGetInstanceCount(type1 gi.GType) (result int32) {
	iv, err := _I.Get(298, "type_get_instance_count", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Int32()
	return
}

// g_type_get_plugin
// container is nil
func TypeGetPlugin(type1 gi.GType) (result TypePlugin) {
	iv, err := _I.Get(299, "type_get_plugin", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_get_qdata
// container is nil
func TypeGetQdata(type1 gi.GType, quark uint32) (result unsafe.Pointer) {
	iv, err := _I.Get(300, "type_get_qdata", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	arg_quark := gi.NewUint32Argument(quark)
	args := []gi.Argument{arg_type1, arg_quark}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Pointer()
	return
}

// g_type_get_type_registration_serial
// container is nil
func TypeGetTypeRegistrationSerial() (result uint32) {
	iv, err := _I.Get(301, "type_get_type_registration_serial", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// g_type_init
// container is nil
func TypeInit() {
	iv, err := _I.Get(302, "type_init", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// g_type_init_with_debug_flags
// container is nil
func TypeInitWithDebugFlags(debug_flags TypeDebugFlags) {
	iv, err := _I.Get(303, "type_init_with_debug_flags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_debug_flags := gi.NewIntArgument(int(debug_flags))
	args := []gi.Argument{arg_debug_flags}
	iv.Call(args, nil, nil)
}

// g_type_interface_add_prerequisite
// container is nil
func TypeInterfaceAddPrerequisite(interface_type gi.GType, prerequisite_type gi.GType) {
	iv, err := _I.Get(304, "type_interface_add_prerequisite", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_interface_type := gi.NewUintArgument(uint(interface_type))
	arg_prerequisite_type := gi.NewUintArgument(uint(prerequisite_type))
	args := []gi.Argument{arg_interface_type, arg_prerequisite_type}
	iv.Call(args, nil, nil)
}

// g_type_interface_get_plugin
// container is nil
func TypeInterfaceGetPlugin(instance_type gi.GType, interface_type gi.GType) (result TypePlugin) {
	iv, err := _I.Get(305, "type_interface_get_plugin", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_instance_type := gi.NewUintArgument(uint(instance_type))
	arg_interface_type := gi.NewUintArgument(uint(interface_type))
	args := []gi.Argument{arg_instance_type, arg_interface_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_interface_peek
// container is nil
func TypeInterfacePeek(instance_class TypeClass, iface_type gi.GType) (result TypeInterface) {
	iv, err := _I.Get(306, "type_interface_peek", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_instance_class := gi.NewPointerArgument(instance_class.P)
	arg_iface_type := gi.NewUintArgument(uint(iface_type))
	args := []gi.Argument{arg_instance_class, arg_iface_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_type_interface_prerequisites
// container is nil
// ret lenArgIdx 1
func TypeInterfacePrerequisites(interface_type gi.GType) (result gi.GTypeArray) {
	iv, err := _I.Get(307, "type_interface_prerequisites", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_interface_type := gi.NewUintArgument(uint(interface_type))
	arg_n_prerequisites := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_interface_type, arg_n_prerequisites}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_prerequisites uint32
	_ = n_prerequisites
	n_prerequisites = outArgs[0].Uint32()
	result = gi.GTypeArray{P: ret.Pointer(), Len: int(n_prerequisites)}
	return
}

// g_type_interfaces
// container is nil
// ret lenArgIdx 1
func TypeInterfaces(type1 gi.GType) (result gi.GTypeArray) {
	iv, err := _I.Get(308, "type_interfaces", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var outArgs [1]gi.Argument
	arg_type1 := gi.NewUintArgument(uint(type1))
	arg_n_interfaces := gi.NewPointerArgument(unsafe.Pointer(&outArgs[0]))
	args := []gi.Argument{arg_type1, arg_n_interfaces}
	var ret gi.Argument
	iv.Call(args, &ret, &outArgs[0])
	var n_interfaces uint32
	_ = n_interfaces
	n_interfaces = outArgs[0].Uint32()
	result = gi.GTypeArray{P: ret.Pointer(), Len: int(n_interfaces)}
	return
}

// g_type_is_a
// container is nil
func TypeIsA(type1 gi.GType, is_a_type gi.GType) (result bool) {
	iv, err := _I.Get(309, "type_is_a", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	arg_is_a_type := gi.NewUintArgument(uint(is_a_type))
	args := []gi.Argument{arg_type1, arg_is_a_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_type_name
// container is nil
func TypeName(type1 gi.GType) (result string) {
	iv, err := _I.Get(310, "type_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_type_name_from_class
// container is nil
func TypeNameFromClass(g_class TypeClass) (result string) {
	iv, err := _I.Get(311, "type_name_from_class", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_g_class := gi.NewPointerArgument(g_class.P)
	args := []gi.Argument{arg_g_class}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_type_name_from_instance
// container is nil
func TypeNameFromInstance(instance TypeInstance) (result string) {
	iv, err := _I.Get(312, "type_name_from_instance", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_instance := gi.NewPointerArgument(instance.P)
	args := []gi.Argument{arg_instance}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Take()
	return
}

// g_type_next_base
// container is nil
func TypeNextBase(leaf_type gi.GType, root_type gi.GType) (result gi.GType) {
	iv, err := _I.Get(313, "type_next_base", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_leaf_type := gi.NewUintArgument(uint(leaf_type))
	arg_root_type := gi.NewUintArgument(uint(root_type))
	args := []gi.Argument{arg_leaf_type, arg_root_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// g_type_parent
// container is nil
func TypeParent(type1 gi.GType) (result gi.GType) {
	iv, err := _I.Get(314, "type_parent", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = gi.GType(ret.Uint())
	return
}

// g_type_qname
// container is nil
func TypeQname(type1 gi.GType) (result uint32) {
	iv, err := _I.Get(315, "type_qname", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint32()
	return
}

// g_type_query
// container is nil
func TypeQueryF(type1 gi.GType, query TypeQuery) {
	iv, err := _I.Get(316, "type_query", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	arg_query := gi.NewPointerArgument(query.P)
	args := []gi.Argument{arg_type1, arg_query}
	iv.Call(args, nil, nil)
}

// g_type_register_dynamic
// container is nil
func TypeRegisterDynamic(parent_type gi.GType, type_name string, plugin ITypePlugin, flags TypeFlags) (result gi.GType) {
	iv, err := _I.Get(317, "type_register_dynamic", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_type_name := gi.CString(type_name)
	var tmp unsafe.Pointer
	if plugin != nil {
		tmp = plugin.P_TypePlugin()
	}
	arg_parent_type := gi.NewUintArgument(uint(parent_type))
	arg_type_name := gi.NewStringArgument(c_type_name)
	arg_plugin := gi.NewPointerArgument(tmp)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_parent_type, arg_type_name, arg_plugin, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_type_name)
	result = gi.GType(ret.Uint())
	return
}

// g_type_register_fundamental
// container is nil
func TypeRegisterFundamental(type_id gi.GType, type_name string, info TypeInfo, finfo TypeFundamentalInfo, flags TypeFlags) (result gi.GType) {
	iv, err := _I.Get(318, "type_register_fundamental", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_type_name := gi.CString(type_name)
	arg_type_id := gi.NewUintArgument(uint(type_id))
	arg_type_name := gi.NewStringArgument(c_type_name)
	arg_info := gi.NewPointerArgument(info.P)
	arg_finfo := gi.NewPointerArgument(finfo.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_type_id, arg_type_name, arg_info, arg_finfo, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_type_name)
	result = gi.GType(ret.Uint())
	return
}

// g_type_register_static
// container is nil
func TypeRegisterStatic(parent_type gi.GType, type_name string, info TypeInfo, flags TypeFlags) (result gi.GType) {
	iv, err := _I.Get(319, "type_register_static", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_type_name := gi.CString(type_name)
	arg_parent_type := gi.NewUintArgument(uint(parent_type))
	arg_type_name := gi.NewStringArgument(c_type_name)
	arg_info := gi.NewPointerArgument(info.P)
	arg_flags := gi.NewIntArgument(int(flags))
	args := []gi.Argument{arg_parent_type, arg_type_name, arg_info, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_type_name)
	result = gi.GType(ret.Uint())
	return
}

// g_type_set_qdata
// container is nil
func TypeSetQdata(type1 gi.GType, quark uint32, data unsafe.Pointer) {
	iv, err := _I.Get(320, "type_set_qdata", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	arg_quark := gi.NewUint32Argument(quark)
	arg_data := gi.NewPointerArgument(data)
	args := []gi.Argument{arg_type1, arg_quark, arg_data}
	iv.Call(args, nil, nil)
}

// g_type_test_flags
// container is nil
func TypeTestFlags(type1 gi.GType, flags uint32) (result bool) {
	iv, err := _I.Get(321, "type_test_flags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	arg_flags := gi.NewUint32Argument(flags)
	args := []gi.Argument{arg_type1, arg_flags}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_value_type_compatible
// container is nil
func ValueTypeCompatible(src_type gi.GType, dest_type gi.GType) (result bool) {
	iv, err := _I.Get(322, "value_type_compatible", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_src_type := gi.NewUintArgument(uint(src_type))
	arg_dest_type := gi.NewUintArgument(uint(dest_type))
	args := []gi.Argument{arg_src_type, arg_dest_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// g_value_type_transformable
// container is nil
func ValueTypeTransformable(src_type gi.GType, dest_type gi.GType) (result bool) {
	iv, err := _I.Get(323, "value_type_transformable", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_src_type := gi.NewUintArgument(uint(src_type))
	arg_dest_type := gi.NewUintArgument(uint(dest_type))
	args := []gi.Argument{arg_src_type, arg_dest_type}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Bool()
	return
}

// constants
const (
	PARAM_MASK                      = 255
	PARAM_STATIC_STRINGS            = 224
	PARAM_USER_SHIFT                = 8
	SIGNAL_FLAGS_MASK               = 511
	SIGNAL_MATCH_MASK               = 63
	TYPE_FLAG_RESERVED_ID_BIT       = 1
	TYPE_FUNDAMENTAL_MAX            = 255
	TYPE_FUNDAMENTAL_SHIFT          = 2
	TYPE_RESERVED_BSE_FIRST         = 32
	TYPE_RESERVED_BSE_LAST          = 48
	TYPE_RESERVED_GLIB_FIRST        = 22
	TYPE_RESERVED_GLIB_LAST         = 31
	TYPE_RESERVED_USER_FIRST        = 49
	VALUE_COLLECT_FORMAT_MAX_LENGTH = 8
	VALUE_NOCOPY_CONTENTS           = 134217728
)
