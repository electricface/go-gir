package g

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
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _ gi.GType
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
	ret := _I.GetGType1(130, "GObject", "Binding")
	return ret
}

// g_binding_get_flags
//
// [ result ] trans: nothing
//
func (v Binding) GetFlags() (result BindingFlags) {
	iv, err := _I.Get1(1306, "GObject", "Binding", "get_flags")
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
//
// [ result ] trans: nothing
//
func (v Binding) GetSource() (result Object) {
	iv, err := _I.Get1(1307, "GObject", "Binding", "get_source")
	if err != nil {
		log.Println("WARN:", err)
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
//
// [ result ] trans: nothing
//
func (v Binding) GetSourceProperty() (result string) {
	iv, err := _I.Get1(1308, "GObject", "Binding", "get_source_property")
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

// g_binding_get_target
//
// [ result ] trans: nothing
//
func (v Binding) GetTarget() (result Object) {
	iv, err := _I.Get1(1309, "GObject", "Binding", "get_target")
	if err != nil {
		log.Println("WARN:", err)
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
//
// [ result ] trans: nothing
//
func (v Binding) GetTargetProperty() (result string) {
	iv, err := _I.Get1(1310, "GObject", "Binding", "get_target_property")
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

// g_binding_unbind
//
func (v Binding) Unbind() {
	iv, err := _I.Get1(1311, "GObject", "Binding", "unbind")
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
	ret := _I.GetGType1(131, "GObject", "BindingFlags")
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
	ret := _I.GetGType1(132, "GObject", "CClosure")
	return ret
}

// g_cclosure_marshal_BOOLEAN__BOXED_BOXED
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalBoolean_BoxedBoxed1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1312, "GObject", "CClosure", "marshal_BOOLEAN__BOXED_BOXED")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalBoolean_Flags1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1313, "GObject", "CClosure", "marshal_BOOLEAN__FLAGS")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalString_ObjectPointer1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1314, "GObject", "CClosure", "marshal_STRING__OBJECT_POINTER")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Boolean1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1315, "GObject", "CClosure", "marshal_VOID__BOOLEAN")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Boxed1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1316, "GObject", "CClosure", "marshal_VOID__BOXED")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Char1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1317, "GObject", "CClosure", "marshal_VOID__CHAR")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Double1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1318, "GObject", "CClosure", "marshal_VOID__DOUBLE")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Enum1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1319, "GObject", "CClosure", "marshal_VOID__ENUM")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Flags1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1320, "GObject", "CClosure", "marshal_VOID__FLAGS")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Float1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1321, "GObject", "CClosure", "marshal_VOID__FLOAT")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Int1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1322, "GObject", "CClosure", "marshal_VOID__INT")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Long1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1323, "GObject", "CClosure", "marshal_VOID__LONG")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Object1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1324, "GObject", "CClosure", "marshal_VOID__OBJECT")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Param1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1325, "GObject", "CClosure", "marshal_VOID__PARAM")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Pointer1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1326, "GObject", "CClosure", "marshal_VOID__POINTER")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_String1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1327, "GObject", "CClosure", "marshal_VOID__STRING")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Uchar1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1328, "GObject", "CClosure", "marshal_VOID__UCHAR")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Uint1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1329, "GObject", "CClosure", "marshal_VOID__UINT")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_UintPointer1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1330, "GObject", "CClosure", "marshal_VOID__UINT_POINTER")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Ulong1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1331, "GObject", "CClosure", "marshal_VOID__ULONG")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Variant1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1332, "GObject", "CClosure", "marshal_VOID__VARIANT")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalVoid_Void1(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1333, "GObject", "CClosure", "marshal_VOID__VOID")
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
//
// [ closure ] trans: nothing
//
// [ return_gvalue ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CClosureMarshalGeneric1(closure Closure, return_gvalue Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1334, "GObject", "CClosure", "marshal_generic")
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
	ret := _I.GetGType1(133, "GObject", "Closure")
	return ret
}

// g_closure_new_object
//
// [ sizeof_closure ] trans: nothing
//
// [ object ] trans: nothing
//
// [ result ] trans: everything
//
func NewClosureObject(sizeof_closure uint32, object IObject) (result Closure) {
	iv, err := _I.Get1(1335, "GObject", "Closure", "new_object")
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
//
// [ sizeof_closure ] trans: nothing
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func NewClosureSimple(sizeof_closure uint32, data unsafe.Pointer) (result Closure) {
	iv, err := _I.Get1(1336, "GObject", "Closure", "new_simple")
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
//
func (v Closure) Invalidate() {
	iv, err := _I.Get1(1337, "GObject", "Closure", "invalidate")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_closure_invoke
//
// [ return_value ] trans: nothing, dir: out
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
func (v Closure) Invoke(return_value Value, n_param_values uint32, param_values unsafe.Pointer, invocation_hint unsafe.Pointer) {
	iv, err := _I.Get1(1338, "GObject", "Closure", "invoke")
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
//
// [ result ] trans: nothing
//
func (v Closure) Ref() (result Closure) {
	iv, err := _I.Get1(1339, "GObject", "Closure", "ref")
	if err != nil {
		log.Println("WARN:", err)
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
//
func (v Closure) Sink() {
	iv, err := _I.Get1(1340, "GObject", "Closure", "sink")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_closure_unref
//
func (v Closure) Unref() {
	iv, err := _I.Get1(1341, "GObject", "Closure", "unref")
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
	ret := _I.GetGType1(134, "GObject", "ClosureNotifyData")
	return ret
}

// Flags ConnectFlags
type ConnectFlags int

const (
	ConnectFlagsAfter   ConnectFlags = 1
	ConnectFlagsSwapped ConnectFlags = 2
)

func ConnectFlagsGetType() gi.GType {
	ret := _I.GetGType1(135, "GObject", "ConnectFlags")
	return ret
}

// Struct EnumClass
type EnumClass struct {
	P unsafe.Pointer
}

const SizeOfStructEnumClass = 32

func EnumClassGetType() gi.GType {
	ret := _I.GetGType1(136, "GObject", "EnumClass")
	return ret
}

// Struct EnumValue
type EnumValue struct {
	P unsafe.Pointer
}

const SizeOfStructEnumValue = 24

func EnumValueGetType() gi.GType {
	ret := _I.GetGType1(137, "GObject", "EnumValue")
	return ret
}

// Struct FlagsClass
type FlagsClass struct {
	P unsafe.Pointer
}

const SizeOfStructFlagsClass = 24

func FlagsClassGetType() gi.GType {
	ret := _I.GetGType1(138, "GObject", "FlagsClass")
	return ret
}

// Struct FlagsValue
type FlagsValue struct {
	P unsafe.Pointer
}

const SizeOfStructFlagsValue = 24

func FlagsValueGetType() gi.GType {
	ret := _I.GetGType1(139, "GObject", "FlagsValue")
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
	ret := _I.GetGType1(140, "GObject", "InitiallyUnowned")
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
	ret := _I.GetGType1(141, "GObject", "InterfaceInfo")
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
	ret := _I.GetGType1(142, "GObject", "Object")
	return ret
}

// Deprecated
//
// g_object_newv
//
// [ object_type ] trans: nothing
//
// [ n_parameters ] trans: nothing
//
// [ parameters ] trans: nothing
//
// [ result ] trans: everything
//
func NewObjectv(object_type gi.GType, n_parameters uint32, parameters unsafe.Pointer) (result Object) {
	iv, err := _I.Get1(1342, "GObject", "Object", "newv")
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
//
// [ what ] trans: nothing
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectCompatControl1(what uint64, data unsafe.Pointer) (result uint64) {
	iv, err := _I.Get1(1343, "GObject", "Object", "compat_control")
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
//
// [ g_iface ] trans: nothing
//
// [ property_name ] trans: nothing
//
// [ result ] trans: nothing
//
func ObjectInterfaceFindProperty1(g_iface TypeInterface, property_name string) (result ParamSpec) {
	iv, err := _I.Get1(1344, "GObject", "Object", "interface_find_property")
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
//
// [ g_iface ] trans: nothing
//
// [ pspec ] trans: nothing
//
func ObjectInterfaceInstallProperty1(g_iface TypeInterface, pspec IParamSpec) {
	iv, err := _I.Get1(1345, "GObject", "Object", "interface_install_property")
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
//
// [ g_iface ] trans: nothing
//
// [ n_properties_p ] trans: everything, dir: out
//
// [ result ] trans: container
//
func ObjectInterfaceListProperties1(g_iface TypeInterface) (result gi.PointerArray) {
	iv, err := _I.Get1(1346, "GObject", "Object", "interface_list_properties")
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
//
// [ source_property ] trans: nothing
//
// [ target ] trans: nothing
//
// [ target_property ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) BindProperty(source_property string, target IObject, target_property string, flags BindingFlags) (result Binding) {
	iv, err := _I.Get1(1347, "GObject", "Object", "bind_property")
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
//
// [ source_property ] trans: nothing
//
// [ target ] trans: nothing
//
// [ target_property ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ transform_to ] trans: nothing
//
// [ transform_from ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) BindPropertyFull(source_property string, target IObject, target_property string, flags BindingFlags, transform_to Closure, transform_from Closure) (result Binding) {
	iv, err := _I.Get1(1348, "GObject", "Object", "bind_property_full")
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
//
func (v Object) ForceFloating() {
	iv, err := _I.Get1(1349, "GObject", "Object", "force_floating")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_object_freeze_notify
//
func (v Object) FreezeNotify() {
	iv, err := _I.Get1(1350, "GObject", "Object", "freeze_notify")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_object_get_data
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) GetData(key string) (result unsafe.Pointer) {
	iv, err := _I.Get1(1351, "GObject", "Object", "get_data")
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
//
// [ property_name ] trans: nothing
//
// [ value ] trans: nothing
//
func (v Object) GetProperty(property_name string, value Value) {
	iv, err := _I.Get1(1352, "GObject", "Object", "get_property")
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
//
// [ quark ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Object) GetQdata(quark uint32) (result unsafe.Pointer) {
	iv, err := _I.Get1(1353, "GObject", "Object", "get_qdata")
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
//
// [ n_properties ] trans: nothing
//
// [ names ] trans: nothing
//
// [ values ] trans: nothing
//
func (v Object) Getv(n_properties uint32, names gi.CStrArray, values unsafe.Pointer) {
	iv, err := _I.Get1(1354, "GObject", "Object", "getv")
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
//
// [ result ] trans: nothing
//
func (v Object) IsFloating() (result bool) {
	iv, err := _I.Get1(1355, "GObject", "Object", "is_floating")
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
//
// [ property_name ] trans: nothing
//
func (v Object) Notify(property_name string) {
	iv, err := _I.Get1(1356, "GObject", "Object", "notify")
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
//
// [ pspec ] trans: nothing
//
func (v Object) NotifyByPspec(pspec IParamSpec) {
	iv, err := _I.Get1(1357, "GObject", "Object", "notify_by_pspec")
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
//
// [ result ] trans: nothing
//
func (v Object) Ref() (result Object) {
	iv, err := _I.Get1(1358, "GObject", "Object", "ref")
	if err != nil {
		log.Println("WARN:", err)
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
//
// [ result ] trans: nothing
//
func (v Object) RefSink() (result Object) {
	iv, err := _I.Get1(1359, "GObject", "Object", "ref_sink")
	if err != nil {
		log.Println("WARN:", err)
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
//
func (v Object) RunDispose() {
	iv, err := _I.Get1(1360, "GObject", "Object", "run_dispose")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_object_set_data
//
// [ key ] trans: nothing
//
// [ data ] trans: nothing
//
func (v Object) SetData(key string, data unsafe.Pointer) {
	iv, err := _I.Get1(1361, "GObject", "Object", "set_data")
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
//
// [ property_name ] trans: nothing
//
// [ value ] trans: nothing
//
func (v Object) SetProperty(property_name string, value Value) {
	iv, err := _I.Get1(1362, "GObject", "Object", "set_property")
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
//
// [ key ] trans: nothing
//
// [ result ] trans: everything
//
func (v Object) StealData(key string) (result unsafe.Pointer) {
	iv, err := _I.Get1(1363, "GObject", "Object", "steal_data")
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
//
// [ quark ] trans: nothing
//
// [ result ] trans: everything
//
func (v Object) StealQdata(quark uint32) (result unsafe.Pointer) {
	iv, err := _I.Get1(1364, "GObject", "Object", "steal_qdata")
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
//
func (v Object) ThawNotify() {
	iv, err := _I.Get1(1365, "GObject", "Object", "thaw_notify")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_object_unref
//
func (v Object) Unref() {
	iv, err := _I.Get1(1366, "GObject", "Object", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_object_watch_closure
//
// [ closure ] trans: nothing
//
func (v Object) WatchClosure(closure Closure) {
	iv, err := _I.Get1(1367, "GObject", "Object", "watch_closure")
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
	ret := _I.GetGType1(143, "GObject", "ObjectConstructParam")
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
	ret := _I.GetGType1(144, "GObject", "ParamFlags")
	return ret
}

// Object ParamSpec
type ParamSpec struct {
	P unsafe.Pointer
}
type IParamSpec interface{ P_ParamSpec() unsafe.Pointer }

func (v ParamSpec) P_ParamSpec() unsafe.Pointer { return v.P }
func ParamSpecGetType() gi.GType {
	ret := _I.GetGType1(145, "GObject", "ParamSpec")
	return ret
}

// g_param_spec_get_blurb
//
// [ result ] trans: nothing
//
func (v ParamSpec) GetBlurb() (result string) {
	iv, err := _I.Get1(1368, "GObject", "ParamSpec", "get_blurb")
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

// g_param_spec_get_default_value
//
// [ result ] trans: nothing
//
func (v ParamSpec) GetDefaultValue() (result Value) {
	iv, err := _I.Get1(1369, "GObject", "ParamSpec", "get_default_value")
	if err != nil {
		log.Println("WARN:", err)
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
//
// [ result ] trans: nothing
//
func (v ParamSpec) GetName() (result string) {
	iv, err := _I.Get1(1370, "GObject", "ParamSpec", "get_name")
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

// g_param_spec_get_name_quark
//
// [ result ] trans: nothing
//
func (v ParamSpec) GetNameQuark() (result uint32) {
	iv, err := _I.Get1(1371, "GObject", "ParamSpec", "get_name_quark")
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
//
// [ result ] trans: nothing
//
func (v ParamSpec) GetNick() (result string) {
	iv, err := _I.Get1(1372, "GObject", "ParamSpec", "get_nick")
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

// g_param_spec_get_qdata
//
// [ quark ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ParamSpec) GetQdata(quark uint32) (result unsafe.Pointer) {
	iv, err := _I.Get1(1373, "GObject", "ParamSpec", "get_qdata")
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
//
// [ result ] trans: nothing
//
func (v ParamSpec) GetRedirectTarget() (result ParamSpec) {
	iv, err := _I.Get1(1374, "GObject", "ParamSpec", "get_redirect_target")
	if err != nil {
		log.Println("WARN:", err)
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
//
// [ quark ] trans: nothing
//
// [ data ] trans: nothing
//
func (v ParamSpec) SetQdata(quark uint32, data unsafe.Pointer) {
	iv, err := _I.Get1(1375, "GObject", "ParamSpec", "set_qdata")
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
//
func (v ParamSpec) Sink() {
	iv, err := _I.Get1(1376, "GObject", "ParamSpec", "sink")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_param_spec_steal_qdata
//
// [ quark ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ParamSpec) StealQdata(quark uint32) (result unsafe.Pointer) {
	iv, err := _I.Get1(1377, "GObject", "ParamSpec", "steal_qdata")
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
	ret := _I.GetGType1(146, "GObject", "ParamSpecBoolean")
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
	ret := _I.GetGType1(147, "GObject", "ParamSpecBoxed")
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
	ret := _I.GetGType1(148, "GObject", "ParamSpecChar")
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
	ret := _I.GetGType1(149, "GObject", "ParamSpecDouble")
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
	ret := _I.GetGType1(150, "GObject", "ParamSpecEnum")
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
	ret := _I.GetGType1(151, "GObject", "ParamSpecFlags")
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
	ret := _I.GetGType1(152, "GObject", "ParamSpecFloat")
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
	ret := _I.GetGType1(153, "GObject", "ParamSpecGType")
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
	ret := _I.GetGType1(154, "GObject", "ParamSpecInt")
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
	ret := _I.GetGType1(155, "GObject", "ParamSpecInt64")
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
	ret := _I.GetGType1(156, "GObject", "ParamSpecLong")
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
	ret := _I.GetGType1(157, "GObject", "ParamSpecObject")
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
	ret := _I.GetGType1(158, "GObject", "ParamSpecOverride")
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
	ret := _I.GetGType1(159, "GObject", "ParamSpecParam")
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
	ret := _I.GetGType1(160, "GObject", "ParamSpecPointer")
	return ret
}

// Struct ParamSpecPool
type ParamSpecPool struct {
	P unsafe.Pointer
}

func ParamSpecPoolGetType() gi.GType {
	ret := _I.GetGType1(161, "GObject", "ParamSpecPool")
	return ret
}

// g_param_spec_pool_insert
//
// [ pspec ] trans: nothing
//
// [ owner_type ] trans: nothing
//
func (v ParamSpecPool) Insert(pspec IParamSpec, owner_type gi.GType) {
	iv, err := _I.Get1(1378, "GObject", "ParamSpecPool", "insert")
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
//
// [ owner_type ] trans: nothing
//
// [ n_pspecs_p ] trans: everything, dir: out
//
// [ result ] trans: container
//
func (v ParamSpecPool) List(owner_type gi.GType) (result gi.PointerArray) {
	iv, err := _I.Get1(1379, "GObject", "ParamSpecPool", "list")
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
//
// [ owner_type ] trans: nothing
//
// [ result ] trans: container
//
func (v ParamSpecPool) ListOwned(owner_type gi.GType) (result List) {
	iv, err := _I.Get1(1380, "GObject", "ParamSpecPool", "list_owned")
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
//
// [ param_name ] trans: nothing
//
// [ owner_type ] trans: nothing
//
// [ walk_ancestors ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ParamSpecPool) Lookup(param_name string, owner_type gi.GType, walk_ancestors bool) (result ParamSpec) {
	iv, err := _I.Get1(1381, "GObject", "ParamSpecPool", "lookup")
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
//
// [ pspec ] trans: nothing
//
func (v ParamSpecPool) Remove(pspec IParamSpec) {
	iv, err := _I.Get1(1382, "GObject", "ParamSpecPool", "remove")
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
//
// [ type_prefixing ] trans: nothing
//
// [ result ] trans: nothing
//
func ParamSpecPoolNew1(type_prefixing bool) (result ParamSpecPool) {
	iv, err := _I.Get1(1383, "GObject", "ParamSpecPool", "new")
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
	ret := _I.GetGType1(162, "GObject", "ParamSpecString")
	return ret
}

// Struct ParamSpecTypeInfo
type ParamSpecTypeInfo struct {
	P unsafe.Pointer
}

const SizeOfStructParamSpecTypeInfo = 56

func ParamSpecTypeInfoGetType() gi.GType {
	ret := _I.GetGType1(163, "GObject", "ParamSpecTypeInfo")
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
	ret := _I.GetGType1(164, "GObject", "ParamSpecUChar")
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
	ret := _I.GetGType1(165, "GObject", "ParamSpecUInt")
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
	ret := _I.GetGType1(166, "GObject", "ParamSpecUInt64")
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
	ret := _I.GetGType1(167, "GObject", "ParamSpecULong")
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
	ret := _I.GetGType1(168, "GObject", "ParamSpecUnichar")
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
	ret := _I.GetGType1(169, "GObject", "ParamSpecValueArray")
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
	ret := _I.GetGType1(170, "GObject", "ParamSpecVariant")
	return ret
}

// Deprecated
//
// Struct Parameter
type Parameter struct {
	P unsafe.Pointer
}

const SizeOfStructParameter = 32

func ParameterGetType() gi.GType {
	ret := _I.GetGType1(171, "GObject", "Parameter")
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
	ret := _I.GetGType1(172, "GObject", "SignalFlags")
	return ret
}

// Struct SignalInvocationHint
type SignalInvocationHint struct {
	P unsafe.Pointer
}

const SizeOfStructSignalInvocationHint = 12

func SignalInvocationHintGetType() gi.GType {
	ret := _I.GetGType1(173, "GObject", "SignalInvocationHint")
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
	ret := _I.GetGType1(174, "GObject", "SignalMatchType")
	return ret
}

// Struct SignalQuery
type SignalQuery struct {
	P unsafe.Pointer
}

const SizeOfStructSignalQuery = 56

func SignalQueryGetType() gi.GType {
	ret := _I.GetGType1(175, "GObject", "SignalQuery")
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
	ret := _I.GetGType1(176, "GObject", "TypeCValue")
	return ret
}

// Struct TypeClass
type TypeClass struct {
	P unsafe.Pointer
}

const SizeOfStructTypeClass = 8

func TypeClassGetType() gi.GType {
	ret := _I.GetGType1(177, "GObject", "TypeClass")
	return ret
}

// Deprecated
//
// g_type_class_add_private
//
// [ private_size ] trans: nothing
//
func (v TypeClass) AddPrivate(private_size uint64) {
	iv, err := _I.Get1(1384, "GObject", "TypeClass", "add_private")
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
//
// [ private_type ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TypeClass) GetPrivate(private_type gi.GType) (result unsafe.Pointer) {
	iv, err := _I.Get1(1385, "GObject", "TypeClass", "get_private")
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
//
// [ result ] trans: nothing
//
func (v TypeClass) PeekParent() (result TypeClass) {
	iv, err := _I.Get1(1386, "GObject", "TypeClass", "peek_parent")
	if err != nil {
		log.Println("WARN:", err)
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
//
func (v TypeClass) Unref() {
	iv, err := _I.Get1(1387, "GObject", "TypeClass", "unref")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_type_class_adjust_private_offset
//
// [ g_class ] trans: nothing
//
// [ private_size_or_offset ] trans: nothing
//
func TypeClassAdjustPrivateOffset1(g_class unsafe.Pointer, private_size_or_offset int32) {
	iv, err := _I.Get1(1388, "GObject", "TypeClass", "adjust_private_offset")
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
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeClassPeek1(type1 gi.GType) (result TypeClass) {
	iv, err := _I.Get1(1389, "GObject", "TypeClass", "peek")
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
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeClassPeekStatic1(type1 gi.GType) (result TypeClass) {
	iv, err := _I.Get1(1390, "GObject", "TypeClass", "peek_static")
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
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeClassRef1(type1 gi.GType) (result TypeClass) {
	iv, err := _I.Get1(1391, "GObject", "TypeClass", "ref")
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

// Deprecated
//
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
	ret := _I.GetGType1(178, "GObject", "TypeDebugFlags")
	return ret
}

// Flags TypeFlags
type TypeFlags int

const (
	TypeFlagsAbstract      TypeFlags = 16
	TypeFlagsValueAbstract TypeFlags = 32
)

func TypeFlagsGetType() gi.GType {
	ret := _I.GetGType1(179, "GObject", "TypeFlags")
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
	ret := _I.GetGType1(180, "GObject", "TypeFundamentalFlags")
	return ret
}

// Struct TypeFundamentalInfo
type TypeFundamentalInfo struct {
	P unsafe.Pointer
}

const SizeOfStructTypeFundamentalInfo = 4

func TypeFundamentalInfoGetType() gi.GType {
	ret := _I.GetGType1(181, "GObject", "TypeFundamentalInfo")
	return ret
}

// Struct TypeInfo
type TypeInfo struct {
	P unsafe.Pointer
}

const SizeOfStructTypeInfo = 72

func TypeInfoGetType() gi.GType {
	ret := _I.GetGType1(182, "GObject", "TypeInfo")
	return ret
}

// Struct TypeInstance
type TypeInstance struct {
	P unsafe.Pointer
}

const SizeOfStructTypeInstance = 8

func TypeInstanceGetType() gi.GType {
	ret := _I.GetGType1(183, "GObject", "TypeInstance")
	return ret
}

// g_type_instance_get_private
//
// [ private_type ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TypeInstance) GetPrivate(private_type gi.GType) (result unsafe.Pointer) {
	iv, err := _I.Get1(1392, "GObject", "TypeInstance", "get_private")
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
	ret := _I.GetGType1(184, "GObject", "TypeInterface")
	return ret
}

// g_type_interface_peek_parent
//
// [ result ] trans: nothing
//
func (v TypeInterface) PeekParent() (result TypeInterface) {
	iv, err := _I.Get1(1393, "GObject", "TypeInterface", "peek_parent")
	if err != nil {
		log.Println("WARN:", err)
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
//
// [ interface_type ] trans: nothing
//
// [ prerequisite_type ] trans: nothing
//
func TypeInterfaceAddPrerequisite1(interface_type gi.GType, prerequisite_type gi.GType) {
	iv, err := _I.Get1(1394, "GObject", "TypeInterface", "add_prerequisite")
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
//
// [ instance_type ] trans: nothing
//
// [ interface_type ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeInterfaceGetPlugin1(instance_type gi.GType, interface_type gi.GType) (result TypePlugin) {
	iv, err := _I.Get1(1395, "GObject", "TypeInterface", "get_plugin")
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
//
// [ instance_class ] trans: nothing
//
// [ iface_type ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeInterfacePeek1(instance_class TypeClass, iface_type gi.GType) (result TypeInterface) {
	iv, err := _I.Get1(1396, "GObject", "TypeInterface", "peek")
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
//
// [ interface_type ] trans: nothing
//
// [ n_prerequisites ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func TypeInterfacePrerequisites1(interface_type gi.GType) (result gi.GTypeArray) {
	iv, err := _I.Get1(1397, "GObject", "TypeInterface", "prerequisites")
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
	ret := _I.GetGType1(185, "GObject", "TypeModule")
	return ret
}

// g_type_module_add_interface
//
// [ instance_type ] trans: nothing
//
// [ interface_type ] trans: nothing
//
// [ interface_info ] trans: nothing
//
func (v TypeModule) AddInterface(instance_type gi.GType, interface_type gi.GType, interface_info InterfaceInfo) {
	iv, err := _I.Get1(1398, "GObject", "TypeModule", "add_interface")
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
//
// [ name ] trans: nothing
//
// [ const_static_values ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TypeModule) RegisterEnum(name string, const_static_values EnumValue) (result gi.GType) {
	iv, err := _I.Get1(1399, "GObject", "TypeModule", "register_enum")
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
//
// [ name ] trans: nothing
//
// [ const_static_values ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TypeModule) RegisterFlags(name string, const_static_values FlagsValue) (result gi.GType) {
	iv, err := _I.Get1(1400, "GObject", "TypeModule", "register_flags")
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
//
// [ parent_type ] trans: nothing
//
// [ type_name ] trans: nothing
//
// [ type_info ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func (v TypeModule) RegisterType(parent_type gi.GType, type_name string, type_info TypeInfo, flags TypeFlags) (result gi.GType) {
	iv, err := _I.Get1(1401, "GObject", "TypeModule", "register_type")
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
//
// [ name ] trans: nothing
//
func (v TypeModule) SetName(name string) {
	iv, err := _I.Get1(1402, "GObject", "TypeModule", "set_name")
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
//
func (v TypeModule) Unuse() {
	iv, err := _I.Get1(1403, "GObject", "TypeModule", "unuse")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_type_module_use
//
// [ result ] trans: nothing
//
func (v TypeModule) Use() (result bool) {
	iv, err := _I.Get1(1404, "GObject", "TypeModule", "use")
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
	ret := _I.GetGType1(186, "GObject", "TypePlugin")
	return ret
}

// g_type_plugin_complete_interface_info
//
// [ instance_type ] trans: nothing
//
// [ interface_type ] trans: nothing
//
// [ info ] trans: nothing
//
func (v *TypePluginIfc) CompleteInterfaceInfo(instance_type gi.GType, interface_type gi.GType, info InterfaceInfo) {
	iv, err := _I.Get1(1405, "GObject", "TypePlugin", "complete_interface_info")
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
//
// [ g_type ] trans: nothing
//
// [ info ] trans: nothing
//
// [ value_table ] trans: nothing
//
func (v *TypePluginIfc) CompleteTypeInfo(g_type gi.GType, info TypeInfo, value_table TypeValueTable) {
	iv, err := _I.Get1(1406, "GObject", "TypePlugin", "complete_type_info")
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
//
func (v *TypePluginIfc) Unuse() {
	iv, err := _I.Get1(1407, "GObject", "TypePlugin", "unuse")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(*(*unsafe.Pointer)(unsafe.Pointer(v)))
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_type_plugin_use
//
func (v *TypePluginIfc) Use() {
	iv, err := _I.Get1(1408, "GObject", "TypePlugin", "use")
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
	ret := _I.GetGType1(187, "GObject", "TypeQuery")
	return ret
}

// Struct TypeValueTable
type TypeValueTable struct {
	P unsafe.Pointer
}

const SizeOfStructTypeValueTable = 64

func TypeValueTableGetType() gi.GType {
	ret := _I.GetGType1(188, "GObject", "TypeValueTable")
	return ret
}

// Struct Value
type Value struct {
	P unsafe.Pointer
}

const SizeOfStructValue = 24

func ValueGetType() gi.GType {
	ret := _I.GetGType1(189, "GObject", "Value")
	return ret
}

// g_value_copy
//
// [ dest_value ] trans: nothing
//
func (v Value) Copy(dest_value Value) {
	iv, err := _I.Get1(1409, "GObject", "Value", "copy")
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
//
// [ result ] trans: everything
//
func (v Value) DupObject() (result Object) {
	iv, err := _I.Get1(1410, "GObject", "Value", "dup_object")
	if err != nil {
		log.Println("WARN:", err)
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
//
// [ result ] trans: everything
//
func (v Value) DupString() (result string) {
	iv, err := _I.Get1(1411, "GObject", "Value", "dup_string")
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
//
// [ result ] trans: everything
//
func (v Value) DupVariant() (result Variant) {
	iv, err := _I.Get1(1412, "GObject", "Value", "dup_variant")
	if err != nil {
		log.Println("WARN:", err)
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
//
// [ result ] trans: nothing
//
func (v Value) FitsPointer() (result bool) {
	iv, err := _I.Get1(1413, "GObject", "Value", "fits_pointer")
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
//
// [ result ] trans: nothing
//
func (v Value) GetBoolean() (result bool) {
	iv, err := _I.Get1(1414, "GObject", "Value", "get_boolean")
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
//
// [ result ] trans: nothing
//
func (v Value) GetBoxed() (result unsafe.Pointer) {
	iv, err := _I.Get1(1415, "GObject", "Value", "get_boxed")
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

// Deprecated
//
// g_value_get_char
//
// [ result ] trans: nothing
//
func (v Value) GetChar() (result int8) {
	iv, err := _I.Get1(1416, "GObject", "Value", "get_char")
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
//
// [ result ] trans: nothing
//
func (v Value) GetDouble() (result float64) {
	iv, err := _I.Get1(1417, "GObject", "Value", "get_double")
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
//
// [ result ] trans: nothing
//
func (v Value) GetEnum() (result int32) {
	iv, err := _I.Get1(1418, "GObject", "Value", "get_enum")
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
//
// [ result ] trans: nothing
//
func (v Value) GetFlags() (result uint32) {
	iv, err := _I.Get1(1419, "GObject", "Value", "get_flags")
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
//
// [ result ] trans: nothing
//
func (v Value) GetFloat() (result float32) {
	iv, err := _I.Get1(1420, "GObject", "Value", "get_float")
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
//
// [ result ] trans: nothing
//
func (v Value) GetGtype() (result gi.GType) {
	iv, err := _I.Get1(1421, "GObject", "Value", "get_gtype")
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
//
// [ result ] trans: nothing
//
func (v Value) GetInt() (result int32) {
	iv, err := _I.Get1(1422, "GObject", "Value", "get_int")
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
//
// [ result ] trans: nothing
//
func (v Value) GetInt64() (result int64) {
	iv, err := _I.Get1(1423, "GObject", "Value", "get_int64")
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
//
// [ result ] trans: nothing
//
func (v Value) GetLong() (result int64) {
	iv, err := _I.Get1(1424, "GObject", "Value", "get_long")
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
//
// [ result ] trans: nothing
//
func (v Value) GetObject() (result Object) {
	iv, err := _I.Get1(1425, "GObject", "Value", "get_object")
	if err != nil {
		log.Println("WARN:", err)
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
//
// [ result ] trans: nothing
//
func (v Value) GetParam() (result ParamSpec) {
	iv, err := _I.Get1(1426, "GObject", "Value", "get_param")
	if err != nil {
		log.Println("WARN:", err)
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
//
// [ result ] trans: nothing
//
func (v Value) GetPointer() (result unsafe.Pointer) {
	iv, err := _I.Get1(1427, "GObject", "Value", "get_pointer")
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
//
// [ result ] trans: nothing
//
func (v Value) GetSchar() (result int8) {
	iv, err := _I.Get1(1428, "GObject", "Value", "get_schar")
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
//
// [ result ] trans: nothing
//
func (v Value) GetString() (result string) {
	iv, err := _I.Get1(1429, "GObject", "Value", "get_string")
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

// g_value_get_uchar
//
// [ result ] trans: nothing
//
func (v Value) GetUchar() (result uint8) {
	iv, err := _I.Get1(1430, "GObject", "Value", "get_uchar")
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
//
// [ result ] trans: nothing
//
func (v Value) GetUint() (result uint32) {
	iv, err := _I.Get1(1431, "GObject", "Value", "get_uint")
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
//
// [ result ] trans: nothing
//
func (v Value) GetUint64() (result uint64) {
	iv, err := _I.Get1(1432, "GObject", "Value", "get_uint64")
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
//
// [ result ] trans: nothing
//
func (v Value) GetUlong() (result uint64) {
	iv, err := _I.Get1(1433, "GObject", "Value", "get_ulong")
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
//
// [ result ] trans: nothing
//
func (v Value) GetVariant() (result Variant) {
	iv, err := _I.Get1(1434, "GObject", "Value", "get_variant")
	if err != nil {
		log.Println("WARN:", err)
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
//
// [ g_type ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Value) Init(g_type gi.GType) (result Value) {
	iv, err := _I.Get1(1435, "GObject", "Value", "init")
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
//
// [ instance ] trans: nothing
//
func (v Value) InitFromInstance(instance TypeInstance) {
	iv, err := _I.Get1(1436, "GObject", "Value", "init_from_instance")
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
//
// [ result ] trans: nothing
//
func (v Value) PeekPointer() (result unsafe.Pointer) {
	iv, err := _I.Get1(1437, "GObject", "Value", "peek_pointer")
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
//
// [ result ] trans: everything
//
func (v Value) Reset() (result Value) {
	iv, err := _I.Get1(1438, "GObject", "Value", "reset")
	if err != nil {
		log.Println("WARN:", err)
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
//
// [ v_boolean ] trans: nothing
//
func (v Value) SetBoolean(v_boolean bool) {
	iv, err := _I.Get1(1439, "GObject", "Value", "set_boolean")
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
//
// [ v_boxed ] trans: nothing
//
func (v Value) SetBoxed(v_boxed unsafe.Pointer) {
	iv, err := _I.Get1(1440, "GObject", "Value", "set_boxed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_boxed := gi.NewPointerArgument(v_boxed)
	args := []gi.Argument{arg_v, arg_v_boxed}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// g_value_set_boxed_take_ownership
//
// [ v_boxed ] trans: nothing
//
func (v Value) SetBoxedTakeOwnership(v_boxed unsafe.Pointer) {
	iv, err := _I.Get1(1441, "GObject", "Value", "set_boxed_take_ownership")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_boxed := gi.NewPointerArgument(v_boxed)
	args := []gi.Argument{arg_v, arg_v_boxed}
	iv.Call(args, nil, nil)
}

// Deprecated
//
// g_value_set_char
//
// [ v_char ] trans: nothing
//
func (v Value) SetChar(v_char int8) {
	iv, err := _I.Get1(1442, "GObject", "Value", "set_char")
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
//
// [ v_double ] trans: nothing
//
func (v Value) SetDouble(v_double float64) {
	iv, err := _I.Get1(1443, "GObject", "Value", "set_double")
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
//
// [ v_enum ] trans: nothing
//
func (v Value) SetEnum(v_enum int32) {
	iv, err := _I.Get1(1444, "GObject", "Value", "set_enum")
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
//
// [ v_flags ] trans: nothing
//
func (v Value) SetFlags(v_flags uint32) {
	iv, err := _I.Get1(1445, "GObject", "Value", "set_flags")
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
//
// [ v_float ] trans: nothing
//
func (v Value) SetFloat(v_float float32) {
	iv, err := _I.Get1(1446, "GObject", "Value", "set_float")
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
//
// [ v_gtype ] trans: nothing
//
func (v Value) SetGtype(v_gtype gi.GType) {
	iv, err := _I.Get1(1447, "GObject", "Value", "set_gtype")
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
//
// [ instance ] trans: nothing
//
func (v Value) SetInstance(instance unsafe.Pointer) {
	iv, err := _I.Get1(1448, "GObject", "Value", "set_instance")
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
//
// [ v_int ] trans: nothing
//
func (v Value) SetInt(v_int int32) {
	iv, err := _I.Get1(1449, "GObject", "Value", "set_int")
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
//
// [ v_int64 ] trans: nothing
//
func (v Value) SetInt64(v_int64 int64) {
	iv, err := _I.Get1(1450, "GObject", "Value", "set_int64")
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
//
// [ v_long ] trans: nothing
//
func (v Value) SetLong(v_long int64) {
	iv, err := _I.Get1(1451, "GObject", "Value", "set_long")
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
//
// [ v_object ] trans: nothing
//
func (v Value) SetObject(v_object IObject) {
	iv, err := _I.Get1(1452, "GObject", "Value", "set_object")
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
//
// [ param ] trans: nothing
//
func (v Value) SetParam(param IParamSpec) {
	iv, err := _I.Get1(1453, "GObject", "Value", "set_param")
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
//
// [ v_pointer ] trans: nothing
//
func (v Value) SetPointer(v_pointer unsafe.Pointer) {
	iv, err := _I.Get1(1454, "GObject", "Value", "set_pointer")
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
//
// [ v_char ] trans: nothing
//
func (v Value) SetSchar(v_char int8) {
	iv, err := _I.Get1(1455, "GObject", "Value", "set_schar")
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
//
// [ v_boxed ] trans: nothing
//
func (v Value) SetStaticBoxed(v_boxed unsafe.Pointer) {
	iv, err := _I.Get1(1456, "GObject", "Value", "set_static_boxed")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_v_boxed := gi.NewPointerArgument(v_boxed)
	args := []gi.Argument{arg_v, arg_v_boxed}
	iv.Call(args, nil, nil)
}

// g_value_set_string
//
// [ v_string ] trans: nothing
//
func (v Value) SetString(v_string string) {
	iv, err := _I.Get1(1457, "GObject", "Value", "set_string")
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

// Deprecated
//
// g_value_set_uchar
//
// [ v_uchar ] trans: nothing
//
func (v Value) SetUchar(v_uchar uint8) {
	iv, err := _I.Get1(1458, "GObject", "Value", "set_uchar")
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
//
// [ v_uint ] trans: nothing
//
func (v Value) SetUint(v_uint uint32) {
	iv, err := _I.Get1(1459, "GObject", "Value", "set_uint")
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
//
// [ v_uint64 ] trans: nothing
//
func (v Value) SetUint64(v_uint64 uint64) {
	iv, err := _I.Get1(1460, "GObject", "Value", "set_uint64")
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
//
// [ v_ulong ] trans: nothing
//
func (v Value) SetUlong(v_ulong uint64) {
	iv, err := _I.Get1(1461, "GObject", "Value", "set_ulong")
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
//
// [ variant ] trans: nothing
//
func (v Value) SetVariant(variant Variant) {
	iv, err := _I.Get1(1462, "GObject", "Value", "set_variant")
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
//
// [ v_boxed ] trans: nothing
//
func (v Value) TakeBoxed(v_boxed unsafe.Pointer) {
	iv, err := _I.Get1(1463, "GObject", "Value", "take_boxed")
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
//
// [ v_string ] trans: nothing
//
func (v Value) TakeString(v_string string) {
	iv, err := _I.Get1(1464, "GObject", "Value", "take_string")
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
//
// [ variant ] trans: everything
//
func (v Value) TakeVariant(variant Variant) {
	iv, err := _I.Get1(1465, "GObject", "Value", "take_variant")
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
//
// [ dest_value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Value) Transform(dest_value Value) (result bool) {
	iv, err := _I.Get1(1466, "GObject", "Value", "transform")
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
//
func (v Value) Unset() {
	iv, err := _I.Get1(1467, "GObject", "Value", "unset")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	iv.Call(args, nil, nil)
}

// g_value_type_compatible
//
// [ src_type ] trans: nothing
//
// [ dest_type ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueTypeCompatible1(src_type gi.GType, dest_type gi.GType) (result bool) {
	iv, err := _I.Get1(1468, "GObject", "Value", "type_compatible")
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
//
// [ src_type ] trans: nothing
//
// [ dest_type ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueTypeTransformable1(src_type gi.GType, dest_type gi.GType) (result bool) {
	iv, err := _I.Get1(1469, "GObject", "Value", "type_transformable")
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
	ret := _I.GetGType1(190, "GObject", "ValueArray")
	return ret
}

// Deprecated
//
// g_value_array_new
//
// [ n_prealloced ] trans: nothing
//
// [ result ] trans: everything
//
func NewValueArray(n_prealloced uint32) (result ValueArray) {
	iv, err := _I.Get1(1470, "GObject", "ValueArray", "new")
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

// Deprecated
//
// g_value_array_append
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ValueArray) Append(value Value) (result ValueArray) {
	iv, err := _I.Get1(1471, "GObject", "ValueArray", "append")
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

// Deprecated
//
// g_value_array_copy
//
// [ result ] trans: everything
//
func (v ValueArray) Copy() (result ValueArray) {
	iv, err := _I.Get1(1472, "GObject", "ValueArray", "copy")
	if err != nil {
		log.Println("WARN:", err)
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
// g_value_array_get_nth
//
// [ index_ ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ValueArray) GetNth(index_ uint32) (result Value) {
	iv, err := _I.Get1(1473, "GObject", "ValueArray", "get_nth")
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

// Deprecated
//
// g_value_array_insert
//
// [ index_ ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ValueArray) Insert(index_ uint32, value Value) (result ValueArray) {
	iv, err := _I.Get1(1474, "GObject", "ValueArray", "insert")
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

// Deprecated
//
// g_value_array_prepend
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ValueArray) Prepend(value Value) (result ValueArray) {
	iv, err := _I.Get1(1475, "GObject", "ValueArray", "prepend")
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

// Deprecated
//
// g_value_array_remove
//
// [ index_ ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ValueArray) Remove(index_ uint32) (result ValueArray) {
	iv, err := _I.Get1(1476, "GObject", "ValueArray", "remove")
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

// Deprecated
//
// g_value_array_sort_with_data
//
// [ compare_func ] trans: nothing
//
// [ user_data ] trans: nothing
//
// [ result ] trans: nothing
//
func (v ValueArray) Sort(compare_func int /*TODO_TYPE CALLBACK*/, user_data unsafe.Pointer) (result ValueArray) {
	iv, err := _I.Get1(1477, "GObject", "ValueArray", "sort")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_compare_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myCompareDataFunc()))
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
	ret := _I.GetGType1(191, "GObject", "WeakRef")
	return ret
}

// Union _Value__data__union
type _Value__data__union struct {
	P unsafe.Pointer
}

const SizeOfUnion_Value__data__union = 8

func _Value__data__unionGetType() gi.GType {
	ret := _I.GetGType1(192, "GObject", "_Value__data__union")
	return ret
}

// g_boxed_copy
//
// [ boxed_type ] trans: nothing
//
// [ src_boxed ] trans: nothing
//
// [ result ] trans: everything
//
func BoxedCopy(boxed_type gi.GType, src_boxed unsafe.Pointer) (result unsafe.Pointer) {
	iv, err := _I.Get1(1478, "GObject", "boxed_copy", "")
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
//
// [ boxed_type ] trans: nothing
//
// [ boxed ] trans: nothing
//
func BoxedFree(boxed_type gi.GType, boxed unsafe.Pointer) {
	iv, err := _I.Get1(1479, "GObject", "boxed_free", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalBoolean_BoxedBoxed(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1480, "GObject", "cclosure_marshal_BOOLEAN__BOXED_BOXED", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalBoolean_Flags(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1481, "GObject", "cclosure_marshal_BOOLEAN__FLAGS", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalString_ObjectPointer(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1482, "GObject", "cclosure_marshal_STRING__OBJECT_POINTER", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Boolean(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1483, "GObject", "cclosure_marshal_VOID__BOOLEAN", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Boxed(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1484, "GObject", "cclosure_marshal_VOID__BOXED", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Char(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1485, "GObject", "cclosure_marshal_VOID__CHAR", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Double(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1486, "GObject", "cclosure_marshal_VOID__DOUBLE", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Enum(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1487, "GObject", "cclosure_marshal_VOID__ENUM", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Flags(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1488, "GObject", "cclosure_marshal_VOID__FLAGS", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Float(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1489, "GObject", "cclosure_marshal_VOID__FLOAT", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Int(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1490, "GObject", "cclosure_marshal_VOID__INT", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Long(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1491, "GObject", "cclosure_marshal_VOID__LONG", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Object(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1492, "GObject", "cclosure_marshal_VOID__OBJECT", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Param(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1493, "GObject", "cclosure_marshal_VOID__PARAM", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Pointer(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1494, "GObject", "cclosure_marshal_VOID__POINTER", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_String(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1495, "GObject", "cclosure_marshal_VOID__STRING", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Uchar(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1496, "GObject", "cclosure_marshal_VOID__UCHAR", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Uint(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1497, "GObject", "cclosure_marshal_VOID__UINT", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_UintPointer(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1498, "GObject", "cclosure_marshal_VOID__UINT_POINTER", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Ulong(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1499, "GObject", "cclosure_marshal_VOID__ULONG", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Variant(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1500, "GObject", "cclosure_marshal_VOID__VARIANT", "")
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
//
// [ closure ] trans: nothing
//
// [ return_value ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalVoid_Void(closure Closure, return_value Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1501, "GObject", "cclosure_marshal_VOID__VOID", "")
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
//
// [ closure ] trans: nothing
//
// [ return_gvalue ] trans: nothing
//
// [ n_param_values ] trans: nothing
//
// [ param_values ] trans: nothing
//
// [ invocation_hint ] trans: nothing
//
// [ marshal_data ] trans: nothing
//
func CclosureMarshalGeneric(closure Closure, return_gvalue Value, n_param_values uint32, param_values Value, invocation_hint unsafe.Pointer, marshal_data unsafe.Pointer) {
	iv, err := _I.Get1(1502, "GObject", "cclosure_marshal_generic", "")
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
//
// [ g_enum_type ] trans: nothing
//
// [ info ] trans: everything, dir: out
//
// [ const_values ] trans: nothing
//
func EnumCompleteTypeInfo(g_enum_type gi.GType, const_values EnumValue) (info int /*TODO_TYPE tag: ifc, biType: struct*/) {
	iv, err := _I.Get1(1503, "GObject", "enum_complete_type_info", "")
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
//
// [ enum_class ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func EnumGetValue(enum_class EnumClass, value int32) (result EnumValue) {
	iv, err := _I.Get1(1504, "GObject", "enum_get_value", "")
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
//
// [ enum_class ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func EnumGetValueByName(enum_class EnumClass, name string) (result EnumValue) {
	iv, err := _I.Get1(1505, "GObject", "enum_get_value_by_name", "")
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
//
// [ enum_class ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ result ] trans: nothing
//
func EnumGetValueByNick(enum_class EnumClass, nick string) (result EnumValue) {
	iv, err := _I.Get1(1506, "GObject", "enum_get_value_by_nick", "")
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
//
// [ name ] trans: nothing
//
// [ const_static_values ] trans: nothing
//
// [ result ] trans: nothing
//
func EnumRegisterStatic(name string, const_static_values EnumValue) (result gi.GType) {
	iv, err := _I.Get1(1507, "GObject", "enum_register_static", "")
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
//
// [ g_enum_type ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: everything
//
func EnumToString(g_enum_type gi.GType, value int32) (result string) {
	iv, err := _I.Get1(1508, "GObject", "enum_to_string", "")
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
//
// [ g_flags_type ] trans: nothing
//
// [ info ] trans: everything, dir: out
//
// [ const_values ] trans: nothing
//
func FlagsCompleteTypeInfo(g_flags_type gi.GType, const_values FlagsValue) (info int /*TODO_TYPE tag: ifc, biType: struct*/) {
	iv, err := _I.Get1(1509, "GObject", "flags_complete_type_info", "")
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
//
// [ flags_class ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func FlagsGetFirstValue(flags_class FlagsClass, value uint32) (result FlagsValue) {
	iv, err := _I.Get1(1510, "GObject", "flags_get_first_value", "")
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
//
// [ flags_class ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func FlagsGetValueByName(flags_class FlagsClass, name string) (result FlagsValue) {
	iv, err := _I.Get1(1511, "GObject", "flags_get_value_by_name", "")
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
//
// [ flags_class ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ result ] trans: nothing
//
func FlagsGetValueByNick(flags_class FlagsClass, nick string) (result FlagsValue) {
	iv, err := _I.Get1(1512, "GObject", "flags_get_value_by_nick", "")
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
//
// [ name ] trans: nothing
//
// [ const_static_values ] trans: nothing
//
// [ result ] trans: nothing
//
func FlagsRegisterStatic(name string, const_static_values FlagsValue) (result gi.GType) {
	iv, err := _I.Get1(1513, "GObject", "flags_register_static", "")
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
//
// [ flags_type ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: everything
//
func FlagsToString(flags_type gi.GType, value uint32) (result string) {
	iv, err := _I.Get1(1514, "GObject", "flags_to_string", "")
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
//
// [ result ] trans: nothing
//
func GtypeGetType() (result gi.GType) {
	iv, err := _I.Get1(1515, "GObject", "gtype_get_type", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ default_value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecBooleanF(name string, nick string, blurb string, default_value bool, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1516, "GObject", "param_spec_boolean", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ boxed_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecBoxedF(name string, nick string, blurb string, boxed_type gi.GType, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1517, "GObject", "param_spec_boxed", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ minimum ] trans: nothing
//
// [ maximum ] trans: nothing
//
// [ default_value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecCharF(name string, nick string, blurb string, minimum int8, maximum int8, default_value int8, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1518, "GObject", "param_spec_char", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ minimum ] trans: nothing
//
// [ maximum ] trans: nothing
//
// [ default_value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecDoubleF(name string, nick string, blurb string, minimum float64, maximum float64, default_value float64, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1519, "GObject", "param_spec_double", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ enum_type ] trans: nothing
//
// [ default_value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecEnumF(name string, nick string, blurb string, enum_type gi.GType, default_value int32, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1520, "GObject", "param_spec_enum", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ flags_type ] trans: nothing
//
// [ default_value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecFlagsF(name string, nick string, blurb string, flags_type gi.GType, default_value uint32, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1521, "GObject", "param_spec_flags", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ minimum ] trans: nothing
//
// [ maximum ] trans: nothing
//
// [ default_value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecFloatF(name string, nick string, blurb string, minimum float32, maximum float32, default_value float32, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1522, "GObject", "param_spec_float", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ is_a_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecGtype(name string, nick string, blurb string, is_a_type gi.GType, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1523, "GObject", "param_spec_gtype", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ minimum ] trans: nothing
//
// [ maximum ] trans: nothing
//
// [ default_value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecIntF(name string, nick string, blurb string, minimum int32, maximum int32, default_value int32, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1524, "GObject", "param_spec_int", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ minimum ] trans: nothing
//
// [ maximum ] trans: nothing
//
// [ default_value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecInt64F(name string, nick string, blurb string, minimum int64, maximum int64, default_value int64, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1525, "GObject", "param_spec_int64", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ minimum ] trans: nothing
//
// [ maximum ] trans: nothing
//
// [ default_value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecLongF(name string, nick string, blurb string, minimum int64, maximum int64, default_value int64, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1526, "GObject", "param_spec_long", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ object_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecObjectF(name string, nick string, blurb string, object_type gi.GType, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1527, "GObject", "param_spec_object", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ param_type ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecParamF(name string, nick string, blurb string, param_type gi.GType, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1528, "GObject", "param_spec_param", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecPointerF(name string, nick string, blurb string, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1529, "GObject", "param_spec_pointer", "")
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
//
// [ type_prefixing ] trans: nothing
//
// [ result ] trans: nothing
//
func ParamSpecPoolNew(type_prefixing bool) (result ParamSpecPool) {
	iv, err := _I.Get1(1530, "GObject", "param_spec_pool_new", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ default_value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecStringF(name string, nick string, blurb string, default_value string, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1531, "GObject", "param_spec_string", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ minimum ] trans: nothing
//
// [ maximum ] trans: nothing
//
// [ default_value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecUchar(name string, nick string, blurb string, minimum uint8, maximum uint8, default_value uint8, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1532, "GObject", "param_spec_uchar", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ minimum ] trans: nothing
//
// [ maximum ] trans: nothing
//
// [ default_value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecUint(name string, nick string, blurb string, minimum uint32, maximum uint32, default_value uint32, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1533, "GObject", "param_spec_uint", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ minimum ] trans: nothing
//
// [ maximum ] trans: nothing
//
// [ default_value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecUint64(name string, nick string, blurb string, minimum uint64, maximum uint64, default_value uint64, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1534, "GObject", "param_spec_uint64", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ minimum ] trans: nothing
//
// [ maximum ] trans: nothing
//
// [ default_value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecUlong(name string, nick string, blurb string, minimum uint64, maximum uint64, default_value uint64, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1535, "GObject", "param_spec_ulong", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ default_value ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecUnicharF(name string, nick string, blurb string, default_value rune, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1536, "GObject", "param_spec_unichar", "")
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
//
// [ name ] trans: nothing
//
// [ nick ] trans: nothing
//
// [ blurb ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ default_value ] trans: everything
//
// [ flags ] trans: nothing
//
// [ result ] trans: everything
//
func ParamSpecVariantF(name string, nick string, blurb string, type1 VariantType, default_value Variant, flags ParamFlags) (result ParamSpec) {
	iv, err := _I.Get1(1537, "GObject", "param_spec_variant", "")
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
//
// [ name ] trans: nothing
//
// [ pspec_info ] trans: nothing
//
// [ result ] trans: nothing
//
func ParamTypeRegisterStatic(name string, pspec_info ParamSpecTypeInfo) (result gi.GType) {
	iv, err := _I.Get1(1538, "GObject", "param_type_register_static", "")
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
//
// [ pspec ] trans: nothing
//
// [ src_value ] trans: nothing
//
// [ dest_value ] trans: nothing
//
// [ strict_validation ] trans: nothing
//
// [ result ] trans: nothing
//
func ParamValueConvert(pspec IParamSpec, src_value Value, dest_value Value, strict_validation bool) (result bool) {
	iv, err := _I.Get1(1539, "GObject", "param_value_convert", "")
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
//
// [ pspec ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ParamValueDefaults(pspec IParamSpec, value Value) (result bool) {
	iv, err := _I.Get1(1540, "GObject", "param_value_defaults", "")
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
//
// [ pspec ] trans: nothing
//
// [ value ] trans: nothing
//
func ParamValueSetDefault(pspec IParamSpec, value Value) {
	iv, err := _I.Get1(1541, "GObject", "param_value_set_default", "")
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
//
// [ pspec ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func ParamValueValidate(pspec IParamSpec, value Value) (result bool) {
	iv, err := _I.Get1(1542, "GObject", "param_value_validate", "")
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
//
// [ pspec ] trans: nothing
//
// [ value1 ] trans: nothing
//
// [ value2 ] trans: nothing
//
// [ result ] trans: nothing
//
func ParamValuesCmp(pspec IParamSpec, value1 Value, value2 Value) (result int32) {
	iv, err := _I.Get1(1543, "GObject", "param_values_cmp", "")
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
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func PointerTypeRegisterStatic(name string) (result gi.GType) {
	iv, err := _I.Get1(1544, "GObject", "pointer_type_register_static", "")
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
//
// [ ihint ] trans: nothing
//
// [ return_accu ] trans: nothing
//
// [ handler_return ] trans: nothing
//
// [ dummy ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalAccumulatorFirstWins(ihint SignalInvocationHint, return_accu Value, handler_return Value, dummy unsafe.Pointer) (result bool) {
	iv, err := _I.Get1(1545, "GObject", "signal_accumulator_first_wins", "")
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
//
// [ ihint ] trans: nothing
//
// [ return_accu ] trans: nothing
//
// [ handler_return ] trans: nothing
//
// [ dummy ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalAccumulatorTrueHandled(ihint SignalInvocationHint, return_accu Value, handler_return Value, dummy unsafe.Pointer) (result bool) {
	iv, err := _I.Get1(1546, "GObject", "signal_accumulator_true_handled", "")
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
//
// [ signal_id ] trans: nothing
//
// [ detail ] trans: nothing
//
// [ hook_func ] trans: nothing
//
// [ hook_data ] trans: nothing
//
// [ data_destroy ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalAddEmissionHook(signal_id uint32, detail uint32, hook_func int /*TODO_TYPE CALLBACK*/, hook_data unsafe.Pointer, data_destroy int /*TODO_TYPE CALLBACK*/) (result uint64) {
	iv, err := _I.Get1(1547, "GObject", "signal_add_emission_hook", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_signal_id := gi.NewUint32Argument(signal_id)
	arg_detail := gi.NewUint32Argument(detail)
	arg_hook_func := gi.NewPointerArgument(unsafe.Pointer(GetPointer_mySignalEmissionHook()))
	arg_hook_data := gi.NewPointerArgument(hook_data)
	arg_data_destroy := gi.NewPointerArgument(unsafe.Pointer(GetPointer_myDestroyNotify()))
	args := []gi.Argument{arg_signal_id, arg_detail, arg_hook_func, arg_hook_data, arg_data_destroy}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.Uint64()
	return
}

// g_signal_chain_from_overridden
//
// [ instance_and_params ] trans: nothing
//
// [ return_value ] trans: nothing
//
func SignalChainFromOverridden(instance_and_params unsafe.Pointer, return_value Value) {
	iv, err := _I.Get1(1548, "GObject", "signal_chain_from_overridden", "")
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
//
// [ instance ] trans: nothing
//
// [ detailed_signal ] trans: nothing
//
// [ closure ] trans: nothing
//
// [ after ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalConnectClosure(instance IObject, detailed_signal string, closure Closure, after bool) (result uint64) {
	iv, err := _I.Get1(1549, "GObject", "signal_connect_closure", "")
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
//
// [ instance ] trans: nothing
//
// [ signal_id ] trans: nothing
//
// [ detail ] trans: nothing
//
// [ closure ] trans: nothing
//
// [ after ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalConnectClosureById(instance IObject, signal_id uint32, detail uint32, closure Closure, after bool) (result uint64) {
	iv, err := _I.Get1(1550, "GObject", "signal_connect_closure_by_id", "")
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
//
// [ instance_and_params ] trans: nothing
//
// [ signal_id ] trans: nothing
//
// [ detail ] trans: nothing
//
// [ return_value ] trans: everything, dir: inout
//
func SignalEmitv(instance_and_params unsafe.Pointer, signal_id uint32, detail uint32, return_value int /*TODO:TYPE*/) {
	iv, err := _I.Get1(1551, "GObject", "signal_emitv", "")
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
//
// [ instance ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalGetInvocationHint(instance IObject) (result SignalInvocationHint) {
	iv, err := _I.Get1(1552, "GObject", "signal_get_invocation_hint", "")
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
//
// [ instance ] trans: nothing
//
// [ handler_id ] trans: nothing
//
func SignalHandlerBlock(instance IObject, handler_id uint64) {
	iv, err := _I.Get1(1553, "GObject", "signal_handler_block", "")
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
//
// [ instance ] trans: nothing
//
// [ handler_id ] trans: nothing
//
func SignalHandlerDisconnect(instance IObject, handler_id uint64) {
	iv, err := _I.Get1(1554, "GObject", "signal_handler_disconnect", "")
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
//
// [ instance ] trans: nothing
//
// [ mask ] trans: nothing
//
// [ signal_id ] trans: nothing
//
// [ detail ] trans: nothing
//
// [ closure ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalHandlerFind(instance IObject, mask SignalMatchTypeFlags, signal_id uint32, detail uint32, closure Closure, func1 unsafe.Pointer, data unsafe.Pointer) (result uint64) {
	iv, err := _I.Get1(1555, "GObject", "signal_handler_find", "")
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
//
// [ instance ] trans: nothing
//
// [ handler_id ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalHandlerIsConnected(instance IObject, handler_id uint64) (result bool) {
	iv, err := _I.Get1(1556, "GObject", "signal_handler_is_connected", "")
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
//
// [ instance ] trans: nothing
//
// [ handler_id ] trans: nothing
//
func SignalHandlerUnblock(instance IObject, handler_id uint64) {
	iv, err := _I.Get1(1557, "GObject", "signal_handler_unblock", "")
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
//
// [ instance ] trans: nothing
//
// [ mask ] trans: nothing
//
// [ signal_id ] trans: nothing
//
// [ detail ] trans: nothing
//
// [ closure ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalHandlersBlockMatched(instance IObject, mask SignalMatchTypeFlags, signal_id uint32, detail uint32, closure Closure, func1 unsafe.Pointer, data unsafe.Pointer) (result uint32) {
	iv, err := _I.Get1(1558, "GObject", "signal_handlers_block_matched", "")
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
//
// [ instance ] trans: nothing
//
func SignalHandlersDestroy(instance IObject) {
	iv, err := _I.Get1(1559, "GObject", "signal_handlers_destroy", "")
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
//
// [ instance ] trans: nothing
//
// [ mask ] trans: nothing
//
// [ signal_id ] trans: nothing
//
// [ detail ] trans: nothing
//
// [ closure ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalHandlersDisconnectMatched(instance IObject, mask SignalMatchTypeFlags, signal_id uint32, detail uint32, closure Closure, func1 unsafe.Pointer, data unsafe.Pointer) (result uint32) {
	iv, err := _I.Get1(1560, "GObject", "signal_handlers_disconnect_matched", "")
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
//
// [ instance ] trans: nothing
//
// [ mask ] trans: nothing
//
// [ signal_id ] trans: nothing
//
// [ detail ] trans: nothing
//
// [ closure ] trans: nothing
//
// [ func1 ] trans: nothing
//
// [ data ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalHandlersUnblockMatched(instance IObject, mask SignalMatchTypeFlags, signal_id uint32, detail uint32, closure Closure, func1 unsafe.Pointer, data unsafe.Pointer) (result uint32) {
	iv, err := _I.Get1(1561, "GObject", "signal_handlers_unblock_matched", "")
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
//
// [ instance ] trans: nothing
//
// [ signal_id ] trans: nothing
//
// [ detail ] trans: nothing
//
// [ may_be_blocked ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalHasHandlerPending(instance IObject, signal_id uint32, detail uint32, may_be_blocked bool) (result bool) {
	iv, err := _I.Get1(1562, "GObject", "signal_has_handler_pending", "")
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
//
// [ itype ] trans: nothing
//
// [ n_ids ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func SignalListIds(itype gi.GType) (result gi.Uint32Array) {
	iv, err := _I.Get1(1563, "GObject", "signal_list_ids", "")
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
//
// [ name ] trans: nothing
//
// [ itype ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalLookup(name string, itype gi.GType) (result uint32) {
	iv, err := _I.Get1(1564, "GObject", "signal_lookup", "")
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
//
// [ signal_id ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalName(signal_id uint32) (result string) {
	iv, err := _I.Get1(1565, "GObject", "signal_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_signal_id := gi.NewUint32Argument(signal_id)
	args := []gi.Argument{arg_signal_id}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_signal_override_class_closure
//
// [ signal_id ] trans: nothing
//
// [ instance_type ] trans: nothing
//
// [ class_closure ] trans: nothing
//
func SignalOverrideClassClosure(signal_id uint32, instance_type gi.GType, class_closure Closure) {
	iv, err := _I.Get1(1566, "GObject", "signal_override_class_closure", "")
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
//
// [ detailed_signal ] trans: nothing
//
// [ itype ] trans: nothing
//
// [ signal_id_p ] trans: everything, dir: out
//
// [ detail_p ] trans: everything, dir: out
//
// [ force_detail_quark ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalParseName(detailed_signal string, itype gi.GType, force_detail_quark bool) (result bool, signal_id_p uint32, detail_p uint32) {
	iv, err := _I.Get1(1567, "GObject", "signal_parse_name", "")
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
//
// [ signal_id ] trans: nothing
//
// [ query ] trans: nothing, dir: out
//
func SignalQueryF(signal_id uint32, query SignalQuery) {
	iv, err := _I.Get1(1568, "GObject", "signal_query", "")
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
//
// [ signal_id ] trans: nothing
//
// [ hook_id ] trans: nothing
//
func SignalRemoveEmissionHook(signal_id uint32, hook_id uint64) {
	iv, err := _I.Get1(1569, "GObject", "signal_remove_emission_hook", "")
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
//
// [ signal_id ] trans: nothing
//
// [ instance_type ] trans: nothing
//
// [ va_marshaller ] trans: nothing
//
func SignalSetVaMarshaller(signal_id uint32, instance_type gi.GType, va_marshaller int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get1(1570, "GObject", "signal_set_va_marshaller", "")
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
//
// [ instance ] trans: nothing
//
// [ signal_id ] trans: nothing
//
// [ detail ] trans: nothing
//
func SignalStopEmission(instance IObject, signal_id uint32, detail uint32) {
	iv, err := _I.Get1(1571, "GObject", "signal_stop_emission", "")
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
//
// [ instance ] trans: nothing
//
// [ detailed_signal ] trans: nothing
//
func SignalStopEmissionByName(instance IObject, detailed_signal string) {
	iv, err := _I.Get1(1572, "GObject", "signal_stop_emission_by_name", "")
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
//
// [ itype ] trans: nothing
//
// [ struct_offset ] trans: nothing
//
// [ result ] trans: nothing
//
func SignalTypeCclosureNew(itype gi.GType, struct_offset uint32) (result Closure) {
	iv, err := _I.Get1(1573, "GObject", "signal_type_cclosure_new", "")
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
//
// [ source ] trans: nothing
//
// [ closure ] trans: nothing
//
func SourceSetClosure(source Source, closure Closure) {
	iv, err := _I.Get1(1574, "GObject", "source_set_closure", "")
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
//
// [ source ] trans: nothing
//
func SourceSetDummyCallback(source Source) {
	iv, err := _I.Get1(1575, "GObject", "source_set_dummy_callback", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_source := gi.NewPointerArgument(source.P)
	args := []gi.Argument{arg_source}
	iv.Call(args, nil, nil)
}

// g_strdup_value_contents
//
// [ value ] trans: nothing
//
// [ result ] trans: everything
//
func StrdupValueContents(value Value) (result string) {
	iv, err := _I.Get1(1576, "GObject", "strdup_value_contents", "")
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
//
// [ class_type ] trans: nothing
//
// [ private_size ] trans: nothing
//
func TypeAddClassPrivate(class_type gi.GType, private_size uint64) {
	iv, err := _I.Get1(1577, "GObject", "type_add_class_private", "")
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
//
// [ class_type ] trans: nothing
//
// [ private_size ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeAddInstancePrivate(class_type gi.GType, private_size uint64) (result int32) {
	iv, err := _I.Get1(1578, "GObject", "type_add_instance_private", "")
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
//
// [ instance_type ] trans: nothing
//
// [ interface_type ] trans: nothing
//
// [ plugin ] trans: nothing
//
func TypeAddInterfaceDynamic(instance_type gi.GType, interface_type gi.GType, plugin ITypePlugin) {
	iv, err := _I.Get1(1579, "GObject", "type_add_interface_dynamic", "")
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
//
// [ instance_type ] trans: nothing
//
// [ interface_type ] trans: nothing
//
// [ info ] trans: nothing
//
func TypeAddInterfaceStatic(instance_type gi.GType, interface_type gi.GType, info InterfaceInfo) {
	iv, err := _I.Get1(1580, "GObject", "type_add_interface_static", "")
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
//
// [ g_class ] trans: nothing
//
// [ is_a_type ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeCheckClassIsA(g_class TypeClass, is_a_type gi.GType) (result bool) {
	iv, err := _I.Get1(1581, "GObject", "type_check_class_is_a", "")
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
//
// [ instance ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeCheckInstance(instance TypeInstance) (result bool) {
	iv, err := _I.Get1(1582, "GObject", "type_check_instance", "")
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
//
// [ instance ] trans: nothing
//
// [ iface_type ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeCheckInstanceIsA(instance TypeInstance, iface_type gi.GType) (result bool) {
	iv, err := _I.Get1(1583, "GObject", "type_check_instance_is_a", "")
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
//
// [ instance ] trans: nothing
//
// [ fundamental_type ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeCheckInstanceIsFundamentallyA(instance TypeInstance, fundamental_type gi.GType) (result bool) {
	iv, err := _I.Get1(1584, "GObject", "type_check_instance_is_fundamentally_a", "")
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
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeCheckIsValueType(type1 gi.GType) (result bool) {
	iv, err := _I.Get1(1585, "GObject", "type_check_is_value_type", "")
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
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeCheckValue(value Value) (result bool) {
	iv, err := _I.Get1(1586, "GObject", "type_check_value", "")
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
//
// [ value ] trans: nothing
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeCheckValueHolds(value Value, type1 gi.GType) (result bool) {
	iv, err := _I.Get1(1587, "GObject", "type_check_value_holds", "")
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
//
// [ type1 ] trans: nothing
//
// [ n_children ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func TypeChildren(type1 gi.GType) (result gi.GTypeArray) {
	iv, err := _I.Get1(1588, "GObject", "type_children", "")
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
//
// [ g_class ] trans: nothing
//
// [ private_size_or_offset ] trans: nothing
//
func TypeClassAdjustPrivateOffset(g_class unsafe.Pointer, private_size_or_offset int32) {
	iv, err := _I.Get1(1589, "GObject", "type_class_adjust_private_offset", "")
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
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeClassPeek(type1 gi.GType) (result TypeClass) {
	iv, err := _I.Get1(1590, "GObject", "type_class_peek", "")
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
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeClassPeekStatic(type1 gi.GType) (result TypeClass) {
	iv, err := _I.Get1(1591, "GObject", "type_class_peek_static", "")
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
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeClassRef(type1 gi.GType) (result TypeClass) {
	iv, err := _I.Get1(1592, "GObject", "type_class_ref", "")
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
//
// [ g_type ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeDefaultInterfacePeek(g_type gi.GType) (result TypeInterface) {
	iv, err := _I.Get1(1593, "GObject", "type_default_interface_peek", "")
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
//
// [ g_type ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeDefaultInterfaceRef(g_type gi.GType) (result TypeInterface) {
	iv, err := _I.Get1(1594, "GObject", "type_default_interface_ref", "")
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
//
// [ g_iface ] trans: nothing
//
func TypeDefaultInterfaceUnref(g_iface TypeInterface) {
	iv, err := _I.Get1(1595, "GObject", "type_default_interface_unref", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_g_iface := gi.NewPointerArgument(g_iface.P)
	args := []gi.Argument{arg_g_iface}
	iv.Call(args, nil, nil)
}

// g_type_depth
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeDepth(type1 gi.GType) (result uint32) {
	iv, err := _I.Get1(1596, "GObject", "type_depth", "")
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
//
// [ type1 ] trans: nothing
//
func TypeEnsure(type1 gi.GType) {
	iv, err := _I.Get1(1597, "GObject", "type_ensure", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	iv.Call(args, nil, nil)
}

// g_type_free_instance
//
// [ instance ] trans: nothing
//
func TypeFreeInstance(instance TypeInstance) {
	iv, err := _I.Get1(1598, "GObject", "type_free_instance", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_instance := gi.NewPointerArgument(instance.P)
	args := []gi.Argument{arg_instance}
	iv.Call(args, nil, nil)
}

// g_type_from_name
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeFromName(name string) (result gi.GType) {
	iv, err := _I.Get1(1599, "GObject", "type_from_name", "")
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
//
// [ type_id ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeFundamental(type_id gi.GType) (result gi.GType) {
	iv, err := _I.Get1(1600, "GObject", "type_fundamental", "")
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
//
// [ result ] trans: nothing
//
func TypeFundamentalNext() (result gi.GType) {
	iv, err := _I.Get1(1601, "GObject", "type_fundamental_next", "")
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
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeGetInstanceCount(type1 gi.GType) (result int32) {
	iv, err := _I.Get1(1602, "GObject", "type_get_instance_count", "")
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
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeGetPlugin(type1 gi.GType) (result TypePlugin) {
	iv, err := _I.Get1(1603, "GObject", "type_get_plugin", "")
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
//
// [ type1 ] trans: nothing
//
// [ quark ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeGetQdata(type1 gi.GType, quark uint32) (result unsafe.Pointer) {
	iv, err := _I.Get1(1604, "GObject", "type_get_qdata", "")
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
//
// [ result ] trans: nothing
//
func TypeGetTypeRegistrationSerial() (result uint32) {
	iv, err := _I.Get1(1605, "GObject", "type_get_type_registration_serial", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var ret gi.Argument
	iv.Call(nil, &ret, nil)
	result = ret.Uint32()
	return
}

// Deprecated
//
// g_type_init
//
func TypeInit() {
	iv, err := _I.Get1(1606, "GObject", "type_init", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// Deprecated
//
// g_type_init_with_debug_flags
//
// [ debug_flags ] trans: nothing
//
func TypeInitWithDebugFlags(debug_flags TypeDebugFlags) {
	iv, err := _I.Get1(1607, "GObject", "type_init_with_debug_flags", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_debug_flags := gi.NewIntArgument(int(debug_flags))
	args := []gi.Argument{arg_debug_flags}
	iv.Call(args, nil, nil)
}

// g_type_interface_add_prerequisite
//
// [ interface_type ] trans: nothing
//
// [ prerequisite_type ] trans: nothing
//
func TypeInterfaceAddPrerequisite(interface_type gi.GType, prerequisite_type gi.GType) {
	iv, err := _I.Get1(1608, "GObject", "type_interface_add_prerequisite", "")
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
//
// [ instance_type ] trans: nothing
//
// [ interface_type ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeInterfaceGetPlugin(instance_type gi.GType, interface_type gi.GType) (result TypePlugin) {
	iv, err := _I.Get1(1609, "GObject", "type_interface_get_plugin", "")
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
//
// [ instance_class ] trans: nothing
//
// [ iface_type ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeInterfacePeek(instance_class TypeClass, iface_type gi.GType) (result TypeInterface) {
	iv, err := _I.Get1(1610, "GObject", "type_interface_peek", "")
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
//
// [ interface_type ] trans: nothing
//
// [ n_prerequisites ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func TypeInterfacePrerequisites(interface_type gi.GType) (result gi.GTypeArray) {
	iv, err := _I.Get1(1611, "GObject", "type_interface_prerequisites", "")
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
//
// [ type1 ] trans: nothing
//
// [ n_interfaces ] trans: everything, dir: out
//
// [ result ] trans: everything
//
func TypeInterfaces(type1 gi.GType) (result gi.GTypeArray) {
	iv, err := _I.Get1(1612, "GObject", "type_interfaces", "")
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
//
// [ type1 ] trans: nothing
//
// [ is_a_type ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeIsA(type1 gi.GType, is_a_type gi.GType) (result bool) {
	iv, err := _I.Get1(1613, "GObject", "type_is_a", "")
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
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeName(type1 gi.GType) (result string) {
	iv, err := _I.Get1(1614, "GObject", "type_name", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_type1 := gi.NewUintArgument(uint(type1))
	args := []gi.Argument{arg_type1}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_type_name_from_class
//
// [ g_class ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeNameFromClass(g_class TypeClass) (result string) {
	iv, err := _I.Get1(1615, "GObject", "type_name_from_class", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_g_class := gi.NewPointerArgument(g_class.P)
	args := []gi.Argument{arg_g_class}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_type_name_from_instance
//
// [ instance ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeNameFromInstance(instance TypeInstance) (result string) {
	iv, err := _I.Get1(1616, "GObject", "type_name_from_instance", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_instance := gi.NewPointerArgument(instance.P)
	args := []gi.Argument{arg_instance}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = ret.String().Copy()
	return
}

// g_type_next_base
//
// [ leaf_type ] trans: nothing
//
// [ root_type ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeNextBase(leaf_type gi.GType, root_type gi.GType) (result gi.GType) {
	iv, err := _I.Get1(1617, "GObject", "type_next_base", "")
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
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeParent(type1 gi.GType) (result gi.GType) {
	iv, err := _I.Get1(1618, "GObject", "type_parent", "")
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
//
// [ type1 ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeQname(type1 gi.GType) (result uint32) {
	iv, err := _I.Get1(1619, "GObject", "type_qname", "")
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
//
// [ type1 ] trans: nothing
//
// [ query ] trans: nothing, dir: out
//
func TypeQueryF(type1 gi.GType, query TypeQuery) {
	iv, err := _I.Get1(1620, "GObject", "type_query", "")
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
//
// [ parent_type ] trans: nothing
//
// [ type_name ] trans: nothing
//
// [ plugin ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeRegisterDynamic(parent_type gi.GType, type_name string, plugin ITypePlugin, flags TypeFlags) (result gi.GType) {
	iv, err := _I.Get1(1621, "GObject", "type_register_dynamic", "")
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
//
// [ type_id ] trans: nothing
//
// [ type_name ] trans: nothing
//
// [ info ] trans: nothing
//
// [ finfo ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeRegisterFundamental(type_id gi.GType, type_name string, info TypeInfo, finfo TypeFundamentalInfo, flags TypeFlags) (result gi.GType) {
	iv, err := _I.Get1(1622, "GObject", "type_register_fundamental", "")
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
//
// [ parent_type ] trans: nothing
//
// [ type_name ] trans: nothing
//
// [ info ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeRegisterStatic(parent_type gi.GType, type_name string, info TypeInfo, flags TypeFlags) (result gi.GType) {
	iv, err := _I.Get1(1623, "GObject", "type_register_static", "")
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
//
// [ type1 ] trans: nothing
//
// [ quark ] trans: nothing
//
// [ data ] trans: nothing
//
func TypeSetQdata(type1 gi.GType, quark uint32, data unsafe.Pointer) {
	iv, err := _I.Get1(1624, "GObject", "type_set_qdata", "")
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
//
// [ type1 ] trans: nothing
//
// [ flags ] trans: nothing
//
// [ result ] trans: nothing
//
func TypeTestFlags(type1 gi.GType, flags uint32) (result bool) {
	iv, err := _I.Get1(1625, "GObject", "type_test_flags", "")
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
//
// [ src_type ] trans: nothing
//
// [ dest_type ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueTypeCompatible(src_type gi.GType, dest_type gi.GType) (result bool) {
	iv, err := _I.Get1(1626, "GObject", "value_type_compatible", "")
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
//
// [ src_type ] trans: nothing
//
// [ dest_type ] trans: nothing
//
// [ result ] trans: nothing
//
func ValueTypeTransformable(src_type gi.GType, dest_type gi.GType) (result bool) {
	iv, err := _I.Get1(1627, "GObject", "value_type_transformable", "")
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
