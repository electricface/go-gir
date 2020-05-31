package cairo

import "github.com/electricface/go-gir3/gi"
import "log"
import "unsafe"

var _I = gi.NewInvokerCache("cairo")
var _ unsafe.Pointer

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
	StatusNoMemory                           = 1
	StatusInvalidRestore                     = 2
	StatusInvalidPopGroup                    = 3
	StatusNoCurrentPoint                     = 4
	StatusInvalidMatrix                      = 5
	StatusInvalidStatus                      = 6
	StatusNullPointer                        = 7
	StatusInvalidString                      = 8
	StatusInvalidPathData                    = 9
	StatusReadError                          = 10
	StatusWriteError                         = 11
	StatusSurfaceFinished                    = 12
	StatusSurfaceTypeMismatch                = 13
	StatusPatternTypeMismatch                = 14
	StatusInvalidContent                     = 15
	StatusInvalidFormat                      = 16
	StatusInvalidVisual                      = 17
	StatusFileNotFound                       = 18
	StatusInvalidDash                        = 19
	StatusInvalidDscComment                  = 20
	StatusInvalidIndex                       = 21
	StatusClipNotRepresentable               = 22
	StatusTempFileError                      = 23
	StatusInvalidStride                      = 24
	StatusFontTypeMismatch                   = 25
	StatusUserFontImmutable                  = 26
	StatusUserFontError                      = 27
	StatusNegativeCount                      = 28
	StatusInvalidClusters                    = 29
	StatusInvalidSlant                       = 30
	StatusInvalidWeight                      = 31
	StatusInvalidSize                        = 32
	StatusUserFontNotImplemented             = 33
	StatusDeviceTypeMismatch                 = 34
	StatusDeviceError                        = 35
	StatusInvalidMeshConstruction            = 36
	StatusDeviceFinished                     = 37
	StatusJbig2GlobalMissing                 = 38
)

// Enum Content
type ContentEnum int

const (
	ContentColor      ContentEnum = 4096
	ContentAlpha                  = 8192
	ContentColorAlpha             = 12288
)

// Enum Operator
type OperatorEnum int

const (
	OperatorClear         OperatorEnum = 0
	OperatorSource                     = 1
	OperatorOver                       = 2
	OperatorIn                         = 3
	OperatorOut                        = 4
	OperatorAtop                       = 5
	OperatorDest                       = 6
	OperatorDestOver                   = 7
	OperatorDestIn                     = 8
	OperatorDestOut                    = 9
	OperatorDestAtop                   = 10
	OperatorXor                        = 11
	OperatorAdd                        = 12
	OperatorSaturate                   = 13
	OperatorMultiply                   = 14
	OperatorScreen                     = 15
	OperatorOverlay                    = 16
	OperatorDarken                     = 17
	OperatorLighten                    = 18
	OperatorColorDodge                 = 19
	OperatorColorBurn                  = 20
	OperatorHardLight                  = 21
	OperatorSoftLight                  = 22
	OperatorDifference                 = 23
	OperatorExclusion                  = 24
	OperatorHslHue                     = 25
	OperatorHslSaturation              = 26
	OperatorHslColor                   = 27
	OperatorHslLuminosity              = 28
)

// Enum Antialias
type AntialiasEnum int

const (
	AntialiasDefault  AntialiasEnum = 0
	AntialiasNone                   = 1
	AntialiasGray                   = 2
	AntialiasSubpixel               = 3
	AntialiasFast                   = 4
	AntialiasGood                   = 5
	AntialiasBest                   = 6
)

// Enum FillRule
type FillRuleEnum int

const (
	FillRuleWinding FillRuleEnum = 0
	FillRuleEvenOdd              = 1
)

// Enum LineCap
type LineCapEnum int

const (
	LineCapButt   LineCapEnum = 0
	LineCapRound              = 1
	LineCapSquare             = 2
)

// Enum LineJoin
type LineJoinEnum int

const (
	LineJoinMiter LineJoinEnum = 0
	LineJoinRound              = 1
	LineJoinBevel              = 2
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
	FontSlantItalic                = 1
	FontSlantOblique               = 2
)

// Enum FontWeight
type FontWeightEnum int

const (
	FontWeightNormal FontWeightEnum = 0
	FontWeightBold                  = 1
)

// Enum SubpixelOrder
type SubpixelOrderEnum int

const (
	SubpixelOrderDefault SubpixelOrderEnum = 0
	SubpixelOrderRgb                       = 1
	SubpixelOrderBgr                       = 2
	SubpixelOrderVrgb                      = 3
	SubpixelOrderVbgr                      = 4
)

// Enum HintStyle
type HintStyleEnum int

const (
	HintStyleDefault HintStyleEnum = 0
	HintStyleNone                  = 1
	HintStyleSlight                = 2
	HintStyleMedium                = 3
	HintStyleFull                  = 4
)

// Enum HintMetrics
type HintMetricsEnum int

const (
	HintMetricsDefault HintMetricsEnum = 0
	HintMetricsOff                     = 1
	HintMetricsOn                      = 2
)

// Struct FontOptions
type FontOptions struct {
	P unsafe.Pointer
}

// Enum FontType
type FontTypeEnum int

const (
	FontTypeToy    FontTypeEnum = 0
	FontTypeFt                  = 1
	FontTypeWin32               = 2
	FontTypeQuartz              = 3
	FontTypeUser                = 4
)

// Enum PathDataType
type PathDataTypeEnum int

const (
	PathDataTypeMoveTo    PathDataTypeEnum = 0
	PathDataTypeLineTo                     = 1
	PathDataTypeCurveTo                    = 2
	PathDataTypeClosePath                  = 3
)

// Enum DeviceType
type DeviceTypeEnum int

const (
	DeviceTypeDrm     DeviceTypeEnum = 0
	DeviceTypeGl                     = 1
	DeviceTypeScript                 = 2
	DeviceTypeXcb                    = 3
	DeviceTypeXlib                   = 4
	DeviceTypeXml                    = 5
	DeviceTypeCogl                   = 6
	DeviceTypeWin32                  = 7
	DeviceTypeInvalid                = -1
)

// Enum SurfaceType
type SurfaceTypeEnum int

const (
	SurfaceTypeImage         SurfaceTypeEnum = 0
	SurfaceTypePdf                           = 1
	SurfaceTypePs                            = 2
	SurfaceTypeXlib                          = 3
	SurfaceTypeXcb                           = 4
	SurfaceTypeGlitz                         = 5
	SurfaceTypeQuartz                        = 6
	SurfaceTypeWin32                         = 7
	SurfaceTypeBeos                          = 8
	SurfaceTypeDirectfb                      = 9
	SurfaceTypeSvg                           = 10
	SurfaceTypeOs2                           = 11
	SurfaceTypeWin32Printing                 = 12
	SurfaceTypeQuartzImage                   = 13
	SurfaceTypeScript                        = 14
	SurfaceTypeQt                            = 15
	SurfaceTypeRecording                     = 16
	SurfaceTypeVg                            = 17
	SurfaceTypeGl                            = 18
	SurfaceTypeDrm                           = 19
	SurfaceTypeTee                           = 20
	SurfaceTypeXml                           = 21
	SurfaceTypeSkia                          = 22
	SurfaceTypeSubsurface                    = 23
	SurfaceTypeCogl                          = 24
)

// Enum Format
type FormatEnum int

const (
	FormatInvalid  FormatEnum = -1
	FormatArgb32              = 0
	FormatRgb24               = 1
	FormatA8                  = 2
	FormatA1                  = 3
	FormatRgb16565            = 4
	FormatRgb30               = 5
)

// Enum PatternType
type PatternTypeEnum int

const (
	PatternTypeSolid        PatternTypeEnum = 0
	PatternTypeSurface                      = 1
	PatternTypeLinear                       = 2
	PatternTypeRadial                       = 3
	PatternTypeMesh                         = 4
	PatternTypeRasterSource                 = 5
)

// Enum Extend
type ExtendEnum int

const (
	ExtendNone    ExtendEnum = 0
	ExtendRepeat             = 1
	ExtendReflect            = 2
	ExtendPad                = 3
)

// Enum Filter
type FilterEnum int

const (
	FilterFast     FilterEnum = 0
	FilterGood                = 1
	FilterBest                = 2
	FilterNearest             = 3
	FilterBilinear            = 4
	FilterGaussian            = 5
)

// Enum RegionOverlap
type RegionOverlapEnum int

const (
	RegionOverlapIn   RegionOverlapEnum = 0
	RegionOverlapOut                    = 1
	RegionOverlapPart                   = 2
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
