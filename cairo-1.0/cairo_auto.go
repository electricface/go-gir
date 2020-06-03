package cairo

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

// Struct Device
type Device struct {
	P unsafe.Pointer
}

// Struct Surface
type Surface struct {
	P unsafe.Pointer
}

// Struct Matrix
type Matrix struct {
	P unsafe.Pointer
}

// Struct Pattern
type Pattern struct {
	P unsafe.Pointer
}

// Struct Region
type Region struct {
	P unsafe.Pointer
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

// Enum Content
type ContentEnum int

const (
	ContentColor      ContentEnum = 4096
	ContentAlpha      ContentEnum = 8192
	ContentColorAlpha ContentEnum = 12288
)

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

// Enum FillRule
type FillRuleEnum int

const (
	FillRuleWinding FillRuleEnum = 0
	FillRuleEvenOdd FillRuleEnum = 1
)

// Enum LineCap
type LineCapEnum int

const (
	LineCapButt   LineCapEnum = 0
	LineCapRound  LineCapEnum = 1
	LineCapSquare LineCapEnum = 2
)

// Enum LineJoin
type LineJoinEnum int

const (
	LineJoinMiter LineJoinEnum = 0
	LineJoinRound LineJoinEnum = 1
	LineJoinBevel LineJoinEnum = 2
)

// Enum TextClusterFlags
type TextClusterFlagsEnum int

const (
	TextClusterFlagsBackward TextClusterFlagsEnum = 1
)

// Enum FontSlant
type FontSlantEnum int

const (
	FontSlantNormal  FontSlantEnum = 0
	FontSlantItalic  FontSlantEnum = 1
	FontSlantOblique FontSlantEnum = 2
)

// Enum FontWeight
type FontWeightEnum int

const (
	FontWeightNormal FontWeightEnum = 0
	FontWeightBold   FontWeightEnum = 1
)

// Enum SubpixelOrder
type SubpixelOrderEnum int

const (
	SubpixelOrderDefault SubpixelOrderEnum = 0
	SubpixelOrderRgb     SubpixelOrderEnum = 1
	SubpixelOrderBgr     SubpixelOrderEnum = 2
	SubpixelOrderVrgb    SubpixelOrderEnum = 3
	SubpixelOrderVbgr    SubpixelOrderEnum = 4
)

// Enum HintStyle
type HintStyleEnum int

const (
	HintStyleDefault HintStyleEnum = 0
	HintStyleNone    HintStyleEnum = 1
	HintStyleSlight  HintStyleEnum = 2
	HintStyleMedium  HintStyleEnum = 3
	HintStyleFull    HintStyleEnum = 4
)

// Enum HintMetrics
type HintMetricsEnum int

const (
	HintMetricsDefault HintMetricsEnum = 0
	HintMetricsOff     HintMetricsEnum = 1
	HintMetricsOn      HintMetricsEnum = 2
)

// Struct FontOptions
type FontOptions struct {
	P unsafe.Pointer
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

// Enum PathDataType
type PathDataTypeEnum int

const (
	PathDataTypeMoveTo    PathDataTypeEnum = 0
	PathDataTypeLineTo    PathDataTypeEnum = 1
	PathDataTypeCurveTo   PathDataTypeEnum = 2
	PathDataTypeClosePath PathDataTypeEnum = 3
)

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

// Enum Extend
type ExtendEnum int

const (
	ExtendNone    ExtendEnum = 0
	ExtendRepeat  ExtendEnum = 1
	ExtendReflect ExtendEnum = 2
	ExtendPad     ExtendEnum = 3
)

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

// Enum RegionOverlap
type RegionOverlapEnum int

const (
	RegionOverlapIn   RegionOverlapEnum = 0
	RegionOverlapOut  RegionOverlapEnum = 1
	RegionOverlapPart RegionOverlapEnum = 2
)

// Struct FontFace
type FontFace struct {
	P unsafe.Pointer
}

// Struct ScaledFont
type ScaledFont struct {
	P unsafe.Pointer
}

// Struct Path
type Path struct {
	P unsafe.Pointer
}

// Struct RectangleInt
type RectangleInt struct {
	P unsafe.Pointer
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
