package gudev

/*
#cgo pkg-config: gudev-1.0
#include <gudev/gudev.h>
*/
import "C"
import "github.com/electricface/go-gir/g-2.0"
import "log"
import "unsafe"
import gi "github.com/electricface/go-gir3/gi-lite"

var _I = gi.NewInvokerCache("GUdev")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GUdev", "1.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Object Client
type Client struct {
	g.Object
}

func WrapClient(p unsafe.Pointer) (r Client) { r.P = p; return }

type IClient interface{ P_Client() unsafe.Pointer }

func (v Client) P_Client() unsafe.Pointer { return v.P }
func ClientGetType() gi.GType {
	ret := _I.GetGType(0, "Client")
	return ret
}

// g_udev_client_new
//
// [ subsystems ] trans: nothing
//
// [ result ] trans: everything
//
func NewClient(subsystems gi.CStrArray) (result Client) {
	iv, err := _I.Get(0, "Client", "new", 0, 0, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_subsystems := gi.NewPointerArgument(subsystems.P)
	args := []gi.Argument{arg_subsystems}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_udev_client_query_by_device_file
//
// [ device_file ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) QueryByDeviceFile(device_file string) (result Device) {
	iv, err := _I.Get(1, "Client", "query_by_device_file", 0, 1, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_device_file := gi.CString(device_file)
	arg_v := gi.NewPointerArgument(v.P)
	arg_device_file := gi.NewStringArgument(c_device_file)
	args := []gi.Argument{arg_v, arg_device_file}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_device_file)
	result.P = ret.Pointer()
	return
}

// g_udev_client_query_by_device_number
//
// [ type1 ] trans: nothing
//
// [ number ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) QueryByDeviceNumber(type1 DeviceTypeEnum, number uint64) (result Device) {
	iv, err := _I.Get(2, "Client", "query_by_device_number", 0, 2, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(int(type1))
	arg_number := gi.NewUint64Argument(number)
	args := []gi.Argument{arg_v, arg_type1, arg_number}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_udev_client_query_by_subsystem
//
// [ subsystem ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) QueryBySubsystem(subsystem string) (result g.List) {
	iv, err := _I.Get(3, "Client", "query_by_subsystem", 0, 3, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_subsystem := gi.CString(subsystem)
	arg_v := gi.NewPointerArgument(v.P)
	arg_subsystem := gi.NewStringArgument(c_subsystem)
	args := []gi.Argument{arg_v, arg_subsystem}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_subsystem)
	result.P = ret.Pointer()
	return
}

// g_udev_client_query_by_subsystem_and_name
//
// [ subsystem ] trans: nothing
//
// [ name ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) QueryBySubsystemAndName(subsystem string, name string) (result Device) {
	iv, err := _I.Get(4, "Client", "query_by_subsystem_and_name", 0, 4, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_subsystem := gi.CString(subsystem)
	c_name := gi.CString(name)
	arg_v := gi.NewPointerArgument(v.P)
	arg_subsystem := gi.NewStringArgument(c_subsystem)
	arg_name := gi.NewStringArgument(c_name)
	args := []gi.Argument{arg_v, arg_subsystem, arg_name}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_subsystem)
	gi.Free(c_name)
	result.P = ret.Pointer()
	return
}

// g_udev_client_query_by_sysfs_path
//
// [ sysfs_path ] trans: nothing
//
// [ result ] trans: everything
//
func (v Client) QueryBySysfsPath(sysfs_path string) (result Device) {
	iv, err := _I.Get(5, "Client", "query_by_sysfs_path", 0, 5, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_sysfs_path := gi.CString(sysfs_path)
	arg_v := gi.NewPointerArgument(v.P)
	arg_sysfs_path := gi.NewStringArgument(c_sysfs_path)
	args := []gi.Argument{arg_v, arg_sysfs_path}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_sysfs_path)
	result.P = ret.Pointer()
	return
}

// ignore GType struct ClientClass

// Struct ClientPrivate
type ClientPrivate struct {
	P unsafe.Pointer
}

func ClientPrivateGetType() gi.GType {
	ret := _I.GetGType(1, "ClientPrivate")
	return ret
}

// Object Device
type Device struct {
	g.Object
}

func WrapDevice(p unsafe.Pointer) (r Device) { r.P = p; return }

type IDevice interface{ P_Device() unsafe.Pointer }

func (v Device) P_Device() unsafe.Pointer { return v.P }
func DeviceGetType() gi.GType {
	ret := _I.GetGType(2, "Device")
	return ret
}

// g_udev_device_get_action
//
// [ result ] trans: nothing
//
func (v Device) GetAction() (result string) {
	iv, err := _I.Get(6, "Device", "get_action", 3, 0, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_device_file
//
// [ result ] trans: nothing
//
func (v Device) GetDeviceFile() (result string) {
	iv, err := _I.Get(7, "Device", "get_device_file", 3, 1, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_device_file_symlinks
//
// [ result ] trans: nothing
//
func (v Device) GetDeviceFileSymlinks() (result gi.CStrArray) {
	iv, err := _I.Get(8, "Device", "get_device_file_symlinks", 3, 2, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_device_number
//
// [ result ] trans: nothing
//
func (v Device) GetDeviceNumber() (result uint64) {
	iv, err := _I.Get(9, "Device", "get_device_number", 3, 3, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_device_type
//
// [ result ] trans: nothing
//
func (v Device) GetDeviceType() (result DeviceTypeEnum) {
	iv, err := _I.Get(10, "Device", "get_device_type", 3, 4, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	args := []gi.Argument{arg_v}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result = DeviceTypeEnum(ret.Int())
	return
}

// g_udev_device_get_devtype
//
// [ result ] trans: nothing
//
func (v Device) GetDevtype() (result string) {
	iv, err := _I.Get(11, "Device", "get_devtype", 3, 5, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_driver
//
// [ result ] trans: nothing
//
func (v Device) GetDriver() (result string) {
	iv, err := _I.Get(12, "Device", "get_driver", 3, 6, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_is_initialized
//
// [ result ] trans: nothing
//
func (v Device) GetIsInitialized() (result bool) {
	iv, err := _I.Get(13, "Device", "get_is_initialized", 3, 7, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_name
//
// [ result ] trans: nothing
//
func (v Device) GetName() (result string) {
	iv, err := _I.Get(14, "Device", "get_name", 3, 8, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_number
//
// [ result ] trans: nothing
//
func (v Device) GetNumber() (result string) {
	iv, err := _I.Get(15, "Device", "get_number", 3, 9, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_parent
//
// [ result ] trans: everything
//
func (v Device) GetParent() (result Device) {
	iv, err := _I.Get(16, "Device", "get_parent", 3, 10, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_parent_with_subsystem
//
// [ subsystem ] trans: nothing
//
// [ devtype ] trans: nothing
//
// [ result ] trans: everything
//
func (v Device) GetParentWithSubsystem(subsystem string, devtype string) (result Device) {
	iv, err := _I.Get(17, "Device", "get_parent_with_subsystem", 3, 11, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_subsystem := gi.CString(subsystem)
	c_devtype := gi.CString(devtype)
	arg_v := gi.NewPointerArgument(v.P)
	arg_subsystem := gi.NewStringArgument(c_subsystem)
	arg_devtype := gi.NewStringArgument(c_devtype)
	args := []gi.Argument{arg_v, arg_subsystem, arg_devtype}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_subsystem)
	gi.Free(c_devtype)
	result.P = ret.Pointer()
	return
}

// g_udev_device_get_property
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) GetProperty(key string) (result string) {
	iv, err := _I.Get(18, "Device", "get_property", 3, 12, gi.INFO_TYPE_OBJECT, 0)
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
	result = ret.String().Copy()
	return
}

// g_udev_device_get_property_as_boolean
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) GetPropertyAsBoolean(key string) (result bool) {
	iv, err := _I.Get(19, "Device", "get_property_as_boolean", 3, 13, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_property_as_double
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) GetPropertyAsDouble(key string) (result float64) {
	iv, err := _I.Get(20, "Device", "get_property_as_double", 3, 14, gi.INFO_TYPE_OBJECT, 0)
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
	result = ret.Double()
	return
}

// g_udev_device_get_property_as_int
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) GetPropertyAsInt(key string) (result int32) {
	iv, err := _I.Get(21, "Device", "get_property_as_int", 3, 15, gi.INFO_TYPE_OBJECT, 0)
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
	result = ret.Int32()
	return
}

// g_udev_device_get_property_as_strv
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) GetPropertyAsStrv(key string) (result gi.CStrArray) {
	iv, err := _I.Get(22, "Device", "get_property_as_strv", 3, 16, gi.INFO_TYPE_OBJECT, 0)
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
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_udev_device_get_property_as_uint64
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) GetPropertyAsUint64(key string) (result uint64) {
	iv, err := _I.Get(23, "Device", "get_property_as_uint64", 3, 17, gi.INFO_TYPE_OBJECT, 0)
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
	result = ret.Uint64()
	return
}

// g_udev_device_get_property_keys
//
// [ result ] trans: nothing
//
func (v Device) GetPropertyKeys() (result gi.CStrArray) {
	iv, err := _I.Get(24, "Device", "get_property_keys", 3, 18, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_seqnum
//
// [ result ] trans: nothing
//
func (v Device) GetSeqnum() (result uint64) {
	iv, err := _I.Get(25, "Device", "get_seqnum", 3, 19, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_subsystem
//
// [ result ] trans: nothing
//
func (v Device) GetSubsystem() (result string) {
	iv, err := _I.Get(26, "Device", "get_subsystem", 3, 20, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_sysfs_attr
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) GetSysfsAttr(name string) (result string) {
	iv, err := _I.Get(27, "Device", "get_sysfs_attr", 3, 21, gi.INFO_TYPE_OBJECT, 0)
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
	result = ret.String().Copy()
	return
}

// g_udev_device_get_sysfs_attr_as_boolean
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) GetSysfsAttrAsBoolean(name string) (result bool) {
	iv, err := _I.Get(28, "Device", "get_sysfs_attr_as_boolean", 3, 22, gi.INFO_TYPE_OBJECT, 0)
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
	result = ret.Bool()
	return
}

// g_udev_device_get_sysfs_attr_as_double
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) GetSysfsAttrAsDouble(name string) (result float64) {
	iv, err := _I.Get(29, "Device", "get_sysfs_attr_as_double", 3, 23, gi.INFO_TYPE_OBJECT, 0)
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
	result = ret.Double()
	return
}

// g_udev_device_get_sysfs_attr_as_int
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) GetSysfsAttrAsInt(name string) (result int32) {
	iv, err := _I.Get(30, "Device", "get_sysfs_attr_as_int", 3, 24, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_sysfs_attr_as_strv
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) GetSysfsAttrAsStrv(name string) (result gi.CStrArray) {
	iv, err := _I.Get(31, "Device", "get_sysfs_attr_as_strv", 3, 25, gi.INFO_TYPE_OBJECT, 0)
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
	result = gi.CStrArray{P: ret.Pointer(), Len: -1}
	result.SetLenZT()
	return
}

// g_udev_device_get_sysfs_attr_as_uint64
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) GetSysfsAttrAsUint64(name string) (result uint64) {
	iv, err := _I.Get(32, "Device", "get_sysfs_attr_as_uint64", 3, 26, gi.INFO_TYPE_OBJECT, 0)
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
	result = ret.Uint64()
	return
}

// g_udev_device_get_sysfs_attr_keys
//
// [ result ] trans: nothing
//
func (v Device) GetSysfsAttrKeys() (result gi.CStrArray) {
	iv, err := _I.Get(33, "Device", "get_sysfs_attr_keys", 3, 27, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_sysfs_path
//
// [ result ] trans: nothing
//
func (v Device) GetSysfsPath() (result string) {
	iv, err := _I.Get(34, "Device", "get_sysfs_path", 3, 28, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_tags
//
// [ result ] trans: nothing
//
func (v Device) GetTags() (result gi.CStrArray) {
	iv, err := _I.Get(35, "Device", "get_tags", 3, 29, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_get_usec_since_initialized
//
// [ result ] trans: nothing
//
func (v Device) GetUsecSinceInitialized() (result uint64) {
	iv, err := _I.Get(36, "Device", "get_usec_since_initialized", 3, 30, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_has_property
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) HasProperty(key string) (result bool) {
	iv, err := _I.Get(37, "Device", "has_property", 3, 31, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_device_has_sysfs_attr
//
// [ key ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Device) HasSysfsAttr(key string) (result bool) {
	iv, err := _I.Get(38, "Device", "has_sysfs_attr", 3, 32, gi.INFO_TYPE_OBJECT, 0)
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

// ignore GType struct DeviceClass

// Struct DevicePrivate
type DevicePrivate struct {
	P unsafe.Pointer
}

func DevicePrivateGetType() gi.GType {
	ret := _I.GetGType(3, "DevicePrivate")
	return ret
}

// Enum DeviceType
type DeviceTypeEnum int

const (
	DeviceTypeNone  DeviceTypeEnum = 0
	DeviceTypeBlock DeviceTypeEnum = 98
	DeviceTypeChar  DeviceTypeEnum = 99
)

func DeviceTypeGetType() gi.GType {
	ret := _I.GetGType(4, "DeviceType")
	return ret
}

// Object Enumerator
type Enumerator struct {
	g.Object
}

func WrapEnumerator(p unsafe.Pointer) (r Enumerator) { r.P = p; return }

type IEnumerator interface{ P_Enumerator() unsafe.Pointer }

func (v Enumerator) P_Enumerator() unsafe.Pointer { return v.P }
func EnumeratorGetType() gi.GType {
	ret := _I.GetGType(5, "Enumerator")
	return ret
}

// g_udev_enumerator_new
//
// [ client ] trans: nothing
//
// [ result ] trans: everything
//
func NewEnumerator(client IClient) (result Enumerator) {
	iv, err := _I.Get(39, "Enumerator", "new", 7, 0, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	var tmp unsafe.Pointer
	if client != nil {
		tmp = client.P_Client()
	}
	arg_client := gi.NewPointerArgument(tmp)
	args := []gi.Argument{arg_client}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_udev_enumerator_add_match_is_initialized
//
// [ result ] trans: nothing
//
func (v Enumerator) AddMatchIsInitialized() (result Enumerator) {
	iv, err := _I.Get(40, "Enumerator", "add_match_is_initialized", 7, 1, gi.INFO_TYPE_OBJECT, 0)
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

// g_udev_enumerator_add_match_name
//
// [ name ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Enumerator) AddMatchName(name string) (result Enumerator) {
	iv, err := _I.Get(41, "Enumerator", "add_match_name", 7, 2, gi.INFO_TYPE_OBJECT, 0)
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
	result.P = ret.Pointer()
	return
}

// g_udev_enumerator_add_match_property
//
// [ name ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Enumerator) AddMatchProperty(name string, value string) (result Enumerator) {
	iv, err := _I.Get(42, "Enumerator", "add_match_property", 7, 3, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_value := gi.CString(value)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	arg_value := gi.NewStringArgument(c_value)
	args := []gi.Argument{arg_v, arg_name, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_value)
	result.P = ret.Pointer()
	return
}

// g_udev_enumerator_add_match_subsystem
//
// [ subsystem ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Enumerator) AddMatchSubsystem(subsystem string) (result Enumerator) {
	iv, err := _I.Get(43, "Enumerator", "add_match_subsystem", 7, 4, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_subsystem := gi.CString(subsystem)
	arg_v := gi.NewPointerArgument(v.P)
	arg_subsystem := gi.NewStringArgument(c_subsystem)
	args := []gi.Argument{arg_v, arg_subsystem}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_subsystem)
	result.P = ret.Pointer()
	return
}

// g_udev_enumerator_add_match_sysfs_attr
//
// [ name ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Enumerator) AddMatchSysfsAttr(name string, value string) (result Enumerator) {
	iv, err := _I.Get(44, "Enumerator", "add_match_sysfs_attr", 7, 5, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_value := gi.CString(value)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	arg_value := gi.NewStringArgument(c_value)
	args := []gi.Argument{arg_v, arg_name, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_value)
	result.P = ret.Pointer()
	return
}

// g_udev_enumerator_add_match_tag
//
// [ tag ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Enumerator) AddMatchTag(tag string) (result Enumerator) {
	iv, err := _I.Get(45, "Enumerator", "add_match_tag", 7, 6, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_tag := gi.CString(tag)
	arg_v := gi.NewPointerArgument(v.P)
	arg_tag := gi.NewStringArgument(c_tag)
	args := []gi.Argument{arg_v, arg_tag}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_tag)
	result.P = ret.Pointer()
	return
}

// g_udev_enumerator_add_nomatch_subsystem
//
// [ subsystem ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Enumerator) AddNomatchSubsystem(subsystem string) (result Enumerator) {
	iv, err := _I.Get(46, "Enumerator", "add_nomatch_subsystem", 7, 7, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_subsystem := gi.CString(subsystem)
	arg_v := gi.NewPointerArgument(v.P)
	arg_subsystem := gi.NewStringArgument(c_subsystem)
	args := []gi.Argument{arg_v, arg_subsystem}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_subsystem)
	result.P = ret.Pointer()
	return
}

// g_udev_enumerator_add_nomatch_sysfs_attr
//
// [ name ] trans: nothing
//
// [ value ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Enumerator) AddNomatchSysfsAttr(name string, value string) (result Enumerator) {
	iv, err := _I.Get(47, "Enumerator", "add_nomatch_sysfs_attr", 7, 8, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_name := gi.CString(name)
	c_value := gi.CString(value)
	arg_v := gi.NewPointerArgument(v.P)
	arg_name := gi.NewStringArgument(c_name)
	arg_value := gi.NewStringArgument(c_value)
	args := []gi.Argument{arg_v, arg_name, arg_value}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_name)
	gi.Free(c_value)
	result.P = ret.Pointer()
	return
}

// g_udev_enumerator_add_sysfs_path
//
// [ sysfs_path ] trans: nothing
//
// [ result ] trans: nothing
//
func (v Enumerator) AddSysfsPath(sysfs_path string) (result Enumerator) {
	iv, err := _I.Get(48, "Enumerator", "add_sysfs_path", 7, 9, gi.INFO_TYPE_OBJECT, 0)
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	c_sysfs_path := gi.CString(sysfs_path)
	arg_v := gi.NewPointerArgument(v.P)
	arg_sysfs_path := gi.NewStringArgument(c_sysfs_path)
	args := []gi.Argument{arg_v, arg_sysfs_path}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	gi.Free(c_sysfs_path)
	result.P = ret.Pointer()
	return
}

// g_udev_enumerator_execute
//
// [ result ] trans: everything
//
func (v Enumerator) Execute() (result g.List) {
	iv, err := _I.Get(49, "Enumerator", "execute", 7, 10, gi.INFO_TYPE_OBJECT, 0)
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

// ignore GType struct EnumeratorClass

// Struct EnumeratorPrivate
type EnumeratorPrivate struct {
	P unsafe.Pointer
}

func EnumeratorPrivateGetType() gi.GType {
	ret := _I.GetGType(6, "EnumeratorPrivate")
	return ret
}

// constants
const ()
