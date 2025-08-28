package icon

import (
	"github.com/a-h/templ"
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

func (i *Icon) SetAttributes(attributes templ.Attributes) *Icon {
	i.Attributes = attributes
	return i
}

func (i *Icon) SetSize(size string) *Icon {
	switch size {
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
