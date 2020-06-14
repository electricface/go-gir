package cairo

/*
#cgo pkg-config: cairo-gobject
*/
import "C"
import "github.com/electricface/go-gir3/gi"
import "log"
import "unsafe"

var _I = gi.NewInvokerCache("cairo")
var _ unsafe.Pointer
var _ *log.Logger

func init() {
	repo := gi.DefaultRepository()
	_, err := repo.Require("cairo", "1.0", gi.REPOSITORY_LOAD_FLAG_LAZY)
	if err != nil {
		panic(err)
	}
}

// Struct Context
type Context struct {
	P unsafe.Pointer
}

func ContextGetType() gi.GType {
	ret := _I.GetGType(0, "Context")
	return ret
}

// Struct Device
type Device struct {
	P unsafe.Pointer
}

func DeviceGetType() gi.GType {
	ret := _I.GetGType(1, "Device")
	return ret
}

// Struct Surface
type Surface struct {
	P unsafe.Pointer
}

func SurfaceGetType() gi.GType {
	ret := _I.GetGType(2, "Surface")
	return ret
}

// Struct Matrix
type Matrix struct {
	P unsafe.Pointer
}

func MatrixGetType() gi.GType {
	ret := _I.GetGType(3, "Matrix")
	return ret
}

// Struct Pattern
type Pattern struct {
	P unsafe.Pointer
}

func PatternGetType() gi.GType {
	ret := _I.GetGType(4, "Pattern")
	return ret
}

// Struct Region
type Region struct {
	P unsafe.Pointer
}

func RegionGetType() gi.GType {
	ret := _I.GetGType(5, "Region")
	return ret
}

// Enum Status
type StatusEnum int

const (
	StatusSuccess                 StatusEnum = 0
	StatusNoMemory                StatusEnum = 1
	StatusInvalidRestore          StatusEnum = 2
	StatusInvalidPopGroup         StatusEnum = 3
	StatusNoCurrentPoint          StatusEnum = 4
	StatusInvalidMatrix           StatusEnum = 5
	StatusInvalidStatus           StatusEnum = 6
	StatusNullPointer             StatusEnum = 7
	StatusInvalidString           StatusEnum = 8
	StatusInvalidPathData         StatusEnum = 9
	StatusReadError               StatusEnum = 10
	StatusWriteError              StatusEnum = 11
	StatusSurfaceFinished         StatusEnum = 12
	StatusSurfaceTypeMismatch     StatusEnum = 13
	StatusPatternTypeMismatch     StatusEnum = 14
	StatusInvalidContent          StatusEnum = 15
	StatusInvalidFormat           StatusEnum = 16
	StatusInvalidVisual           StatusEnum = 17
	StatusFileNotFound            StatusEnum = 18
	StatusInvalidDash             StatusEnum = 19
	StatusInvalidDscComment       StatusEnum = 20
	StatusInvalidIndex            StatusEnum = 21
	StatusClipNotRepresentable    StatusEnum = 22
	StatusTempFileError           StatusEnum = 23
	StatusInvalidStride           StatusEnum = 24
	StatusFontTypeMismatch        StatusEnum = 25
	StatusUserFontImmutable       StatusEnum = 26
	StatusUserFontError           StatusEnum = 27
	StatusNegativeCount           StatusEnum = 28
	StatusInvalidClusters         StatusEnum = 29
	StatusInvalidSlant            StatusEnum = 30
	StatusInvalidWeight           StatusEnum = 31
	StatusInvalidSize             StatusEnum = 32
	StatusUserFontNotImplemented  StatusEnum = 33
	StatusDeviceTypeMismatch      StatusEnum = 34
	StatusDeviceError             StatusEnum = 35
	StatusInvalidMeshConstruction StatusEnum = 36
	StatusDeviceFinished          StatusEnum = 37
	StatusJbig2GlobalMissing      StatusEnum = 38
)

func StatusGetType() gi.GType {
	ret := _I.GetGType(6, "Status")
	return ret
}

// Enum Content
type ContentEnum int

const (
	ContentColor      ContentEnum = 4096
	ContentAlpha      ContentEnum = 8192
	ContentColorAlpha ContentEnum = 12288
)

func ContentGetType() gi.GType {
	ret := _I.GetGType(7, "Content")
	return ret
}

// Enum Operator
type OperatorEnum int

const (
	OperatorClear         OperatorEnum = 0
	OperatorSource        OperatorEnum = 1
	OperatorOver          OperatorEnum = 2
	OperatorIn            OperatorEnum = 3
	OperatorOut           OperatorEnum = 4
	OperatorAtop          OperatorEnum = 5
	OperatorDest          OperatorEnum = 6
	OperatorDestOver      OperatorEnum = 7
	OperatorDestIn        OperatorEnum = 8
	OperatorDestOut       OperatorEnum = 9
	OperatorDestAtop      OperatorEnum = 10
	OperatorXor           OperatorEnum = 11
	OperatorAdd           OperatorEnum = 12
	OperatorSaturate      OperatorEnum = 13
	OperatorMultiply      OperatorEnum = 14
	OperatorScreen        OperatorEnum = 15
	OperatorOverlay       OperatorEnum = 16
	OperatorDarken        OperatorEnum = 17
	OperatorLighten       OperatorEnum = 18
	OperatorColorDodge    OperatorEnum = 19
	OperatorColorBurn     OperatorEnum = 20
	OperatorHardLight     OperatorEnum = 21
	OperatorSoftLight     OperatorEnum = 22
	OperatorDifference    OperatorEnum = 23
	OperatorExclusion     OperatorEnum = 24
	OperatorHslHue        OperatorEnum = 25
	OperatorHslSaturation OperatorEnum = 26
	OperatorHslColor      OperatorEnum = 27
	OperatorHslLuminosity OperatorEnum = 28
)

func OperatorGetType() gi.GType {
	ret := _I.GetGType(8, "Operator")
	return ret
}

// Enum Antialias
type AntialiasEnum int

const (
	AntialiasDefault  AntialiasEnum = 0
	AntialiasNone     AntialiasEnum = 1
	AntialiasGray     AntialiasEnum = 2
	AntialiasSubpixel AntialiasEnum = 3
	AntialiasFast     AntialiasEnum = 4
	AntialiasGood     AntialiasEnum = 5
	AntialiasBest     AntialiasEnum = 6
)

func AntialiasGetType() gi.GType {
	ret := _I.GetGType(9, "Antialias")
	return ret
}

// Enum FillRule
type FillRuleEnum int

const (
	FillRuleWinding FillRuleEnum = 0
	FillRuleEvenOdd FillRuleEnum = 1
)

func FillRuleGetType() gi.GType {
	ret := _I.GetGType(10, "FillRule")
	return ret
}

// Enum LineCap
type LineCapEnum int

const (
	LineCapButt   LineCapEnum = 0
	LineCapRound  LineCapEnum = 1
	LineCapSquare LineCapEnum = 2
)

func LineCapGetType() gi.GType {
	ret := _I.GetGType(11, "LineCap")
	return ret
}

// Enum LineJoin
type LineJoinEnum int

const (
	LineJoinMiter LineJoinEnum = 0
	LineJoinRound LineJoinEnum = 1
	LineJoinBevel LineJoinEnum = 2
)

func LineJoinGetType() gi.GType {
	ret := _I.GetGType(12, "LineJoin")
	return ret
}

// Enum TextClusterFlags
type TextClusterFlagsEnum int

const (
	TextClusterFlagsBackward TextClusterFlagsEnum = 1
)

func TextClusterFlagsGetType() gi.GType {
	ret := _I.GetGType(13, "TextClusterFlags")
	return ret
}

// Enum FontSlant
type FontSlantEnum int

const (
	FontSlantNormal  FontSlantEnum = 0
	FontSlantItalic  FontSlantEnum = 1
	FontSlantOblique FontSlantEnum = 2
)

func FontSlantGetType() gi.GType {
	ret := _I.GetGType(14, "FontSlant")
	return ret
}

// Enum FontWeight
type FontWeightEnum int

const (
	FontWeightNormal FontWeightEnum = 0
	FontWeightBold   FontWeightEnum = 1
)

func FontWeightGetType() gi.GType {
	ret := _I.GetGType(15, "FontWeight")
	return ret
}

// Enum SubpixelOrder
type SubpixelOrderEnum int

const (
	SubpixelOrderDefault SubpixelOrderEnum = 0
	SubpixelOrderRgb     SubpixelOrderEnum = 1
	SubpixelOrderBgr     SubpixelOrderEnum = 2
	SubpixelOrderVrgb    SubpixelOrderEnum = 3
	SubpixelOrderVbgr    SubpixelOrderEnum = 4
)

func SubpixelOrderGetType() gi.GType {
	ret := _I.GetGType(16, "SubpixelOrder")
	return ret
}

// Enum HintStyle
type HintStyleEnum int

const (
	HintStyleDefault HintStyleEnum = 0
	HintStyleNone    HintStyleEnum = 1
	HintStyleSlight  HintStyleEnum = 2
	HintStyleMedium  HintStyleEnum = 3
	HintStyleFull    HintStyleEnum = 4
)

func HintStyleGetType() gi.GType {
	ret := _I.GetGType(17, "HintStyle")
	return ret
}

// Enum HintMetrics
type HintMetricsEnum int

const (
	HintMetricsDefault HintMetricsEnum = 0
	HintMetricsOff     HintMetricsEnum = 1
	HintMetricsOn      HintMetricsEnum = 2
)

func HintMetricsGetType() gi.GType {
	ret := _I.GetGType(18, "HintMetrics")
	return ret
}

// Struct FontOptions
type FontOptions struct {
	P unsafe.Pointer
}

func FontOptionsGetType() gi.GType {
	ret := _I.GetGType(19, "FontOptions")
	return ret
}

// Enum FontType
type FontTypeEnum int

const (
	FontTypeToy    FontTypeEnum = 0
	FontTypeFt     FontTypeEnum = 1
	FontTypeWin32  FontTypeEnum = 2
	FontTypeQuartz FontTypeEnum = 3
	FontTypeUser   FontTypeEnum = 4
)

func FontTypeGetType() gi.GType {
	ret := _I.GetGType(20, "FontType")
	return ret
}

// Enum PathDataType
type PathDataTypeEnum int

const (
	PathDataTypeMoveTo    PathDataTypeEnum = 0
	PathDataTypeLineTo    PathDataTypeEnum = 1
	PathDataTypeCurveTo   PathDataTypeEnum = 2
	PathDataTypeClosePath PathDataTypeEnum = 3
)

func PathDataTypeGetType() gi.GType {
	ret := _I.GetGType(21, "PathDataType")
	return ret
}

// Enum DeviceType
type DeviceTypeEnum int

const (
	DeviceTypeDrm     DeviceTypeEnum = 0
	DeviceTypeGl      DeviceTypeEnum = 1
	DeviceTypeScript  DeviceTypeEnum = 2
	DeviceTypeXcb     DeviceTypeEnum = 3
	DeviceTypeXlib    DeviceTypeEnum = 4
	DeviceTypeXml     DeviceTypeEnum = 5
	DeviceTypeCogl    DeviceTypeEnum = 6
	DeviceTypeWin32   DeviceTypeEnum = 7
	DeviceTypeInvalid DeviceTypeEnum = -1
)

func DeviceTypeGetType() gi.GType {
	ret := _I.GetGType(22, "DeviceType")
	return ret
}

// Enum SurfaceType
type SurfaceTypeEnum int

const (
	SurfaceTypeImage         SurfaceTypeEnum = 0
	SurfaceTypePdf           SurfaceTypeEnum = 1
	SurfaceTypePs            SurfaceTypeEnum = 2
	SurfaceTypeXlib          SurfaceTypeEnum = 3
	SurfaceTypeXcb           SurfaceTypeEnum = 4
	SurfaceTypeGlitz         SurfaceTypeEnum = 5
	SurfaceTypeQuartz        SurfaceTypeEnum = 6
	SurfaceTypeWin32         SurfaceTypeEnum = 7
	SurfaceTypeBeos          SurfaceTypeEnum = 8
	SurfaceTypeDirectfb      SurfaceTypeEnum = 9
	SurfaceTypeSvg           SurfaceTypeEnum = 10
	SurfaceTypeOs2           SurfaceTypeEnum = 11
	SurfaceTypeWin32Printing SurfaceTypeEnum = 12
	SurfaceTypeQuartzImage   SurfaceTypeEnum = 13
	SurfaceTypeScript        SurfaceTypeEnum = 14
	SurfaceTypeQt            SurfaceTypeEnum = 15
	SurfaceTypeRecording     SurfaceTypeEnum = 16
	SurfaceTypeVg            SurfaceTypeEnum = 17
	SurfaceTypeGl            SurfaceTypeEnum = 18
	SurfaceTypeDrm           SurfaceTypeEnum = 19
	SurfaceTypeTee           SurfaceTypeEnum = 20
	SurfaceTypeXml           SurfaceTypeEnum = 21
	SurfaceTypeSkia          SurfaceTypeEnum = 22
	SurfaceTypeSubsurface    SurfaceTypeEnum = 23
	SurfaceTypeCogl          SurfaceTypeEnum = 24
)

func SurfaceTypeGetType() gi.GType {
	ret := _I.GetGType(23, "SurfaceType")
	return ret
}

// Enum Format
type FormatEnum int

const (
	FormatInvalid  FormatEnum = -1
	FormatArgb32   FormatEnum = 0
	FormatRgb24    FormatEnum = 1
	FormatA8       FormatEnum = 2
	FormatA1       FormatEnum = 3
	FormatRgb16565 FormatEnum = 4
	FormatRgb30    FormatEnum = 5
)

func FormatGetType() gi.GType {
	ret := _I.GetGType(24, "Format")
	return ret
}

// Enum PatternType
type PatternTypeEnum int

const (
	PatternTypeSolid        PatternTypeEnum = 0
	PatternTypeSurface      PatternTypeEnum = 1
	PatternTypeLinear       PatternTypeEnum = 2
	PatternTypeRadial       PatternTypeEnum = 3
	PatternTypeMesh         PatternTypeEnum = 4
	PatternTypeRasterSource PatternTypeEnum = 5
)

func PatternTypeGetType() gi.GType {
	ret := _I.GetGType(25, "PatternType")
	return ret
}

// Enum Extend
type ExtendEnum int

const (
	ExtendNone    ExtendEnum = 0
	ExtendRepeat  ExtendEnum = 1
	ExtendReflect ExtendEnum = 2
	ExtendPad     ExtendEnum = 3
)

func ExtendGetType() gi.GType {
	ret := _I.GetGType(26, "Extend")
	return ret
}

// Enum Filter
type FilterEnum int

const (
	FilterFast     FilterEnum = 0
	FilterGood     FilterEnum = 1
	FilterBest     FilterEnum = 2
	FilterNearest  FilterEnum = 3
	FilterBilinear FilterEnum = 4
	FilterGaussian FilterEnum = 5
)

func FilterGetType() gi.GType {
	ret := _I.GetGType(27, "Filter")
	return ret
}

// Enum RegionOverlap
type RegionOverlapEnum int

const (
	RegionOverlapIn   RegionOverlapEnum = 0
	RegionOverlapOut  RegionOverlapEnum = 1
	RegionOverlapPart RegionOverlapEnum = 2
)

func RegionOverlapGetType() gi.GType {
	ret := _I.GetGType(28, "RegionOverlap")
	return ret
}

// Struct FontFace
type FontFace struct {
	P unsafe.Pointer
}

func FontFaceGetType() gi.GType {
	ret := _I.GetGType(29, "FontFace")
	return ret
}

// Struct ScaledFont
type ScaledFont struct {
	P unsafe.Pointer
}

func ScaledFontGetType() gi.GType {
	ret := _I.GetGType(30, "ScaledFont")
	return ret
}

// Struct Path
type Path struct {
	P unsafe.Pointer
}

func PathGetType() gi.GType {
	ret := _I.GetGType(31, "Path")
	return ret
}

// Struct RectangleInt
type RectangleInt struct {
	P unsafe.Pointer
}

const SizeOfStructRectangleInt = 16

func RectangleIntGetType() gi.GType {
	ret := _I.GetGType(32, "RectangleInt")
	return ret
}

// cairo_image_surface_create
// container is nil
func ImageSurfaceCreate() {
	iv, err := _I.Get(0, "image_surface_create", "")
	if err != nil {
		log.Println("WARN:", err)
		return
	}
	iv.Call(nil, nil, nil)
}

// constants
const ()
