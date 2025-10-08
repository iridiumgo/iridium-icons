package icon

import (
	"github.com/a-h/templ"
)

type Size string

const (
	SizeXXS Size = "xxs"
	SizeXS  Size = "xs"
	SizeSM  Size = "sm"
	SizeMD  Size = "md"
	SizeLG  Size = "lg"
	SizeXL  Size = "xl"
	SizeXXL Size = "xxl"
)

type Icon struct {
	Width          string
	Height         string
	ViewBox        string
	Fill           string
	Stroke         string
	StrokeWidth    string
	StrokeLineCap  string
	StrokeLineJoin string
	Attributes     templ.Attributes
	componentFn    func(i *Icon) templ.Component
}

func NewIcon(svgFn func(i *Icon) templ.Component) *Icon {
	return &Icon{
		Width:          "16",
		Height:         "16",
		ViewBox:        "0 0 24 24",
		Fill:           "none",
		Stroke:         "currentColor",
		StrokeWidth:    "2",
		StrokeLineCap:  "round",
		StrokeLineJoin: "round",
		Attributes:     templ.Attributes{},
		componentFn:    svgFn,
	}
}

func (i *Icon) Copy() *Icon {
	return &Icon{
		Width:          i.Width,
		Height:         i.Height,
		ViewBox:        i.ViewBox,
		Fill:           i.Fill,
		Stroke:         i.Stroke,
		StrokeWidth:    i.StrokeWidth,
		StrokeLineCap:  i.StrokeLineCap,
		StrokeLineJoin: i.StrokeLineJoin,
		Attributes:     DeepCopyMap(i.Attributes),
		componentFn:    i.componentFn,
	}
}

func (i *Icon) SetAttributes(attributes templ.Attributes) *Icon {
	cpy := i.Copy()
	cpy.Attributes = attributes
	return cpy
}

func (i *Icon) SetSize(size string) *Icon {
	cpy := i.Copy()
	switch size {
	case "xxs":
		cpy.Width, cpy.Height = "6", "6"
	case "xs":
		cpy.Width, cpy.Height = "12", "12"
	case "sm":
		cpy.Width, cpy.Height = "16", "16"
	case "md":
		cpy.Width, cpy.Height = "20", "20"
	case "lg":
		cpy.Width, cpy.Height = "24", "24"
	case "xl":
		cpy.Width, cpy.Height = "32", "32"
	case "xxl":
		cpy.Width, cpy.Height = "46", "46"
	default:
		cpy.Width, cpy.Height = "24", "24"
	}
	return cpy
}

func (i *Icon) SetStroke(stroke string) *Icon {
	cpy := i.Copy()
	cpy.Stroke = stroke
	return i
}

func (i *Icon) SetStrokeWidth(width string) *Icon {
	cpy := i.Copy()
	cpy.StrokeWidth = width
	return i
}

func (i *Icon) Component() templ.Component {
	if i.componentFn == nil {
		return templ.NopComponent
	}
	return i.componentFn(i)
}
