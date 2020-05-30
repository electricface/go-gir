package gudev

import "github.com/electricface/go-gir/gobject-2.0"
import "github.com/electricface/go-gir3/gi"
import "log"
import "unsafe"

var _I = gi.NewInvokerCache("GUdev")
var _ unsafe.Pointer

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("GUdev", "1.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Object Client
type Client struct {
	gobject.Object
}

// g_udev_client_new
// container is not nil, container is Client
// is constructor
func NewClient(subsystems int /*TODO_TYPE isPtr: true, tag: array*/) (result Client) {
	iv, err := _I.Get(0, "Client", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_subsystems := gi.NewIntArgument(subsystems) /*TODO*/
	args := []gi.Argument{arg_subsystems}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_udev_client_query_by_device_file
// container is not nil, container is Client
// is method
func (v Client) QueryByDeviceFile(device_file string) (result Device) {
	iv, err := _I.Get(1, "Client", "query_by_device_file")
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
	result.P = ret.Pointer()
	gi.Free(c_device_file)
	return
}

// g_udev_client_query_by_device_number
// container is not nil, container is Client
// is method
func (v Client) QueryByDeviceNumber(type1 int /*TODO_TYPE isPtr: false, tag: interface*/, number uint64) (result Device) {
	iv, err := _I.Get(2, "Client", "query_by_device_number")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_v := gi.NewPointerArgument(v.P)
	arg_type1 := gi.NewIntArgument(type1) /*TODO*/
	arg_number := gi.NewUint64Argument(number)
	args := []gi.Argument{arg_v, arg_type1, arg_number}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_udev_client_query_by_subsystem
// container is not nil, container is Client
// is method
func (v Client) QueryBySubsystem(subsystem string) (result int /*TODO_TYPE isPtr: true, tag: glist*/) {
	iv, err := _I.Get(3, "Client", "query_by_subsystem")
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
	result = ret.Int() /*TODO*/
	gi.Free(c_subsystem)
	return
}

// g_udev_client_query_by_subsystem_and_name
// container is not nil, container is Client
// is method
func (v Client) QueryBySubsystemAndName(subsystem string, name string) (result Device) {
	iv, err := _I.Get(4, "Client", "query_by_subsystem_and_name")
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
	result.P = ret.Pointer()
	gi.Free(c_subsystem)
	gi.Free(c_name)
	return
}

// g_udev_client_query_by_sysfs_path
// container is not nil, container is Client
// is method
func (v Client) QueryBySysfsPath(sysfs_path string) (result Device) {
	iv, err := _I.Get(5, "Client", "query_by_sysfs_path")
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
	result.P = ret.Pointer()
	gi.Free(c_sysfs_path)
	return
}

// ignore GType struct ClientClass
// Struct ClientPrivate
type ClientPrivate struct {
	P unsafe.Pointer
}

// Object Device
type Device struct {
	gobject.Object
}

// g_udev_device_get_action
// container is not nil, container is Device
// is method
func (v Device) GetAction() (result string) {
	iv, err := _I.Get(6, "Device", "get_action")
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

// g_udev_device_get_device_file
// container is not nil, container is Device
// is method
func (v Device) GetDeviceFile() (result string) {
	iv, err := _I.Get(7, "Device", "get_device_file")
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

// g_udev_device_get_device_file_symlinks
// container is not nil, container is Device
// is method
func (v Device) GetDeviceFileSymlinks() (result int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(8, "Device", "get_device_file_symlinks")
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

// g_udev_device_get_device_number
// container is not nil, container is Device
// is method
func (v Device) GetDeviceNumber() (result uint64) {
	iv, err := _I.Get(9, "Device", "get_device_number")
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
// container is not nil, container is Device
// is method
func (v Device) GetDeviceType() (result int /*TODO_TYPE isPtr: false, tag: interface*/) {
	iv, err := _I.Get(10, "Device", "get_device_type")
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

// g_udev_device_get_devtype
// container is not nil, container is Device
// is method
func (v Device) GetDevtype() (result string) {
	iv, err := _I.Get(11, "Device", "get_devtype")
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

// g_udev_device_get_driver
// container is not nil, container is Device
// is method
func (v Device) GetDriver() (result string) {
	iv, err := _I.Get(12, "Device", "get_driver")
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

// g_udev_device_get_is_initialized
// container is not nil, container is Device
// is method
func (v Device) GetIsInitialized() (result bool) {
	iv, err := _I.Get(13, "Device", "get_is_initialized")
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
// container is not nil, container is Device
// is method
func (v Device) GetName() (result string) {
	iv, err := _I.Get(14, "Device", "get_name")
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

// g_udev_device_get_number
// container is not nil, container is Device
// is method
func (v Device) GetNumber() (result string) {
	iv, err := _I.Get(15, "Device", "get_number")
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

// g_udev_device_get_parent
// container is not nil, container is Device
// is method
func (v Device) GetParent() (result Device) {
	iv, err := _I.Get(16, "Device", "get_parent")
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
// container is not nil, container is Device
// is method
func (v Device) GetParentWithSubsystem(subsystem string, devtype string) (result Device) {
	iv, err := _I.Get(17, "Device", "get_parent_with_subsystem")
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
	result.P = ret.Pointer()
	gi.Free(c_subsystem)
	gi.Free(c_devtype)
	return
}

// g_udev_device_get_property
// container is not nil, container is Device
// is method
func (v Device) GetProperty(key string) (result string) {
	iv, err := _I.Get(18, "Device", "get_property")
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
	result = ret.String().Take()
	gi.Free(c_key)
	return
}

// g_udev_device_get_property_as_boolean
// container is not nil, container is Device
// is method
func (v Device) GetPropertyAsBoolean(key string) (result bool) {
	iv, err := _I.Get(19, "Device", "get_property_as_boolean")
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
	result = ret.Bool()
	gi.Free(c_key)
	return
}

// g_udev_device_get_property_as_double
// container is not nil, container is Device
// is method
func (v Device) GetPropertyAsDouble(key string) (result float64) {
	iv, err := _I.Get(20, "Device", "get_property_as_double")
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
	result = ret.Double()
	gi.Free(c_key)
	return
}

// g_udev_device_get_property_as_int
// container is not nil, container is Device
// is method
func (v Device) GetPropertyAsInt(key string) (result int32) {
	iv, err := _I.Get(21, "Device", "get_property_as_int")
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
	result = ret.Int32()
	gi.Free(c_key)
	return
}

// g_udev_device_get_property_as_strv
// container is not nil, container is Device
// is method
func (v Device) GetPropertyAsStrv(key string) (result int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(22, "Device", "get_property_as_strv")
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
	result = ret.Int() /*TODO*/
	gi.Free(c_key)
	return
}

// g_udev_device_get_property_as_uint64
// container is not nil, container is Device
// is method
func (v Device) GetPropertyAsUint64(key string) (result uint64) {
	iv, err := _I.Get(23, "Device", "get_property_as_uint64")
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
	result = ret.Uint64()
	gi.Free(c_key)
	return
}

// g_udev_device_get_property_keys
// container is not nil, container is Device
// is method
func (v Device) GetPropertyKeys() (result int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(24, "Device", "get_property_keys")
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

// g_udev_device_get_seqnum
// container is not nil, container is Device
// is method
func (v Device) GetSeqnum() (result uint64) {
	iv, err := _I.Get(25, "Device", "get_seqnum")
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
// container is not nil, container is Device
// is method
func (v Device) GetSubsystem() (result string) {
	iv, err := _I.Get(26, "Device", "get_subsystem")
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

// g_udev_device_get_sysfs_attr
// container is not nil, container is Device
// is method
func (v Device) GetSysfsAttr(name string) (result string) {
	iv, err := _I.Get(27, "Device", "get_sysfs_attr")
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
	result = ret.String().Take()
	gi.Free(c_name)
	return
}

// g_udev_device_get_sysfs_attr_as_boolean
// container is not nil, container is Device
// is method
func (v Device) GetSysfsAttrAsBoolean(name string) (result bool) {
	iv, err := _I.Get(28, "Device", "get_sysfs_attr_as_boolean")
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
	result = ret.Bool()
	gi.Free(c_name)
	return
}

// g_udev_device_get_sysfs_attr_as_double
// container is not nil, container is Device
// is method
func (v Device) GetSysfsAttrAsDouble(name string) (result float64) {
	iv, err := _I.Get(29, "Device", "get_sysfs_attr_as_double")
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
	result = ret.Double()
	gi.Free(c_name)
	return
}

// g_udev_device_get_sysfs_attr_as_int
// container is not nil, container is Device
// is method
func (v Device) GetSysfsAttrAsInt(name string) (result int32) {
	iv, err := _I.Get(30, "Device", "get_sysfs_attr_as_int")
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
	result = ret.Int32()
	gi.Free(c_name)
	return
}

// g_udev_device_get_sysfs_attr_as_strv
// container is not nil, container is Device
// is method
func (v Device) GetSysfsAttrAsStrv(name string) (result int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(31, "Device", "get_sysfs_attr_as_strv")
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
	result = ret.Int() /*TODO*/
	gi.Free(c_name)
	return
}

// g_udev_device_get_sysfs_attr_as_uint64
// container is not nil, container is Device
// is method
func (v Device) GetSysfsAttrAsUint64(name string) (result uint64) {
	iv, err := _I.Get(32, "Device", "get_sysfs_attr_as_uint64")
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
	result = ret.Uint64()
	gi.Free(c_name)
	return
}

// g_udev_device_get_sysfs_attr_keys
// container is not nil, container is Device
// is method
func (v Device) GetSysfsAttrKeys() (result int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(33, "Device", "get_sysfs_attr_keys")
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

// g_udev_device_get_sysfs_path
// container is not nil, container is Device
// is method
func (v Device) GetSysfsPath() (result string) {
	iv, err := _I.Get(34, "Device", "get_sysfs_path")
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

// g_udev_device_get_tags
// container is not nil, container is Device
// is method
func (v Device) GetTags() (result int /*TODO_TYPE isPtr: true, tag: array*/) {
	iv, err := _I.Get(35, "Device", "get_tags")
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

// g_udev_device_get_usec_since_initialized
// container is not nil, container is Device
// is method
func (v Device) GetUsecSinceInitialized() (result uint64) {
	iv, err := _I.Get(36, "Device", "get_usec_since_initialized")
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
// container is not nil, container is Device
// is method
func (v Device) HasProperty(key string) (result bool) {
	iv, err := _I.Get(37, "Device", "has_property")
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
	result = ret.Bool()
	gi.Free(c_key)
	return
}

// g_udev_device_has_sysfs_attr
// container is not nil, container is Device
// is method
func (v Device) HasSysfsAttr(key string) (result bool) {
	iv, err := _I.Get(38, "Device", "has_sysfs_attr")
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
	result = ret.Bool()
	gi.Free(c_key)
	return
}

// ignore GType struct DeviceClass
// Struct DevicePrivate
type DevicePrivate struct {
	P unsafe.Pointer
}
type DeviceTypeEnum int

const (
	DeviceTypeNone  DeviceTypeEnum = 0
	DeviceTypeBlock                = 98
	DeviceTypeChar                 = 99
)

// Object Enumerator
type Enumerator struct {
	gobject.Object
}

// g_udev_enumerator_new
// container is not nil, container is Enumerator
// is constructor
func NewEnumerator(client Client) (result Enumerator) {
	iv, err := _I.Get(39, "Enumerator", "new")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	arg_client := gi.NewPointerArgument(client.P)
	args := []gi.Argument{arg_client}
	var ret gi.Argument
	iv.Call(args, &ret, nil)
	result.P = ret.Pointer()
	return
}

// g_udev_enumerator_add_match_is_initialized
// container is not nil, container is Enumerator
// is method
func (v Enumerator) AddMatchIsInitialized() (result Enumerator) {
	iv, err := _I.Get(40, "Enumerator", "add_match_is_initialized")
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
// container is not nil, container is Enumerator
// is method
func (v Enumerator) AddMatchName(name string) (result Enumerator) {
	iv, err := _I.Get(41, "Enumerator", "add_match_name")
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
	result.P = ret.Pointer()
	gi.Free(c_name)
	return
}

// g_udev_enumerator_add_match_property
// container is not nil, container is Enumerator
// is method
func (v Enumerator) AddMatchProperty(name string, value string) (result Enumerator) {
	iv, err := _I.Get(42, "Enumerator", "add_match_property")
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
	result.P = ret.Pointer()
	gi.Free(c_name)
	gi.Free(c_value)
	return
}

// g_udev_enumerator_add_match_subsystem
// container is not nil, container is Enumerator
// is method
func (v Enumerator) AddMatchSubsystem(subsystem string) (result Enumerator) {
	iv, err := _I.Get(43, "Enumerator", "add_match_subsystem")
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
	result.P = ret.Pointer()
	gi.Free(c_subsystem)
	return
}

// g_udev_enumerator_add_match_sysfs_attr
// container is not nil, container is Enumerator
// is method
func (v Enumerator) AddMatchSysfsAttr(name string, value string) (result Enumerator) {
	iv, err := _I.Get(44, "Enumerator", "add_match_sysfs_attr")
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
	result.P = ret.Pointer()
	gi.Free(c_name)
	gi.Free(c_value)
	return
}

// g_udev_enumerator_add_match_tag
// container is not nil, container is Enumerator
// is method
func (v Enumerator) AddMatchTag(tag string) (result Enumerator) {
	iv, err := _I.Get(45, "Enumerator", "add_match_tag")
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
	result.P = ret.Pointer()
	gi.Free(c_tag)
	return
}

// g_udev_enumerator_add_nomatch_subsystem
// container is not nil, container is Enumerator
// is method
func (v Enumerator) AddNomatchSubsystem(subsystem string) (result Enumerator) {
	iv, err := _I.Get(46, "Enumerator", "add_nomatch_subsystem")
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
	result.P = ret.Pointer()
	gi.Free(c_subsystem)
	return
}

// g_udev_enumerator_add_nomatch_sysfs_attr
// container is not nil, container is Enumerator
// is method
func (v Enumerator) AddNomatchSysfsAttr(name string, value string) (result Enumerator) {
	iv, err := _I.Get(47, "Enumerator", "add_nomatch_sysfs_attr")
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
	result.P = ret.Pointer()
	gi.Free(c_name)
	gi.Free(c_value)
	return
}

// g_udev_enumerator_add_sysfs_path
// container is not nil, container is Enumerator
// is method
func (v Enumerator) AddSysfsPath(sysfs_path string) (result Enumerator) {
	iv, err := _I.Get(48, "Enumerator", "add_sysfs_path")
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
	result.P = ret.Pointer()
	gi.Free(c_sysfs_path)
	return
}

// g_udev_enumerator_execute
// container is not nil, container is Enumerator
// is method
func (v Enumerator) Execute() (result int /*TODO_TYPE isPtr: true, tag: glist*/) {
	iv, err := _I.Get(49, "Enumerator", "execute")
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

// ignore GType struct EnumeratorClass
// Struct EnumeratorPrivate
type EnumeratorPrivate struct {
	P unsafe.Pointer
}

// constants
const ()
