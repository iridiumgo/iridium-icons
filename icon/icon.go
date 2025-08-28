package icon

import (
	"github.com/a-h/templ"
)

const (
	SizeXXS = "xxs"
	SizeXS  = "xs"
	SizeSM  = "sm"
	SizeMD  = "md"
	SizeLG  = "lg"
	SizeXL  = "xl"
	SizeXXL = "xxl"
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
		Attributes:     templ.Attributes{},
		componentFn:    i.componentFn,
	}
}

func (i *Icon) SetAttributes(attributes templ.Attributes) *Icon {
	i.Attributes = attributes
	return i
}

func (i *Icon) SetSize(size string) *Icon {
	switch size {
	case "xxs":
		i.Width, i.Height = "6", "6"
	case "xs":
		i.Width, i.Height = "12", "12"
	case "sm":
		i.Width, i.Height = "16", "16"
	case "md":
		i.Width, i.Height = "20", "20"
	case "lg":
		i.Width, i.Height = "24", "24"
	case "xl":
		i.Width, i.Height = "32", "32"
	case "xxl":
		i.Width, i.Height = "46", "46"
	default:
		i.Width, i.Height = "24", "24"
	}
	return i
}

func (i *Icon) SetStroke(stroke string) *Icon {
	i.Stroke = stroke
	return i
}

func (i *Icon) SetStrokeWidth(width string) *Icon {
	i.StrokeWidth = width
	return i
}

func (i *Icon) Component() templ.Component {
	if i.componentFn == nil {
		return templ.NopComponent
	}
	return i.componentFn(i)
}
