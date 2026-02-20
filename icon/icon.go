package icon

import (
	"maps"

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
	StrokeStr      string
	StrokeWidthStr string
	StrokeLineCap  string
	StrokeLineJoin string
	AttributesMap  templ.Attributes
	componentFn    func(i *Icon) templ.Component
}

func NewIcon(svgFn func(i *Icon) templ.Component) *Icon {
	return &Icon{
		Width:          "16",
		Height:         "16",
		ViewBox:        "0 0 24 24",
		Fill:           "none",
		StrokeStr:      "currentColor",
		StrokeWidthStr: "2",
		StrokeLineCap:  "round",
		StrokeLineJoin: "round",
		AttributesMap:  templ.Attributes{},
		componentFn:    svgFn,
	}
}

func (i *Icon) Copy() *Icon {
	return &Icon{
		Width:          i.Width,
		Height:         i.Height,
		ViewBox:        i.ViewBox,
		Fill:           i.Fill,
		StrokeStr:      i.StrokeStr,
		StrokeWidthStr: i.StrokeWidthStr,
		StrokeLineCap:  i.StrokeLineCap,
		StrokeLineJoin: i.StrokeLineJoin,
		AttributesMap:  DeepCopyMap(i.AttributesMap),
		componentFn:    i.componentFn,
	}
}

func (i *Icon) Attributes(attributes templ.Attributes) *Icon {
	cpy := i.Copy()
	cpy.AttributesMap = attributes
	return cpy
}

func (i *Icon) Size(size string) *Icon {
	cpy := i.Copy()
	cpy.applySize(Size(size))
	return cpy
}

func (i *Icon) Stroke(stroke string) *Icon {
	cpy := i.Copy()
	cpy.StrokeStr = stroke
	return i
}

func (i *Icon) StrokeWidth(width string) *Icon {
	cpy := i.Copy()
	cpy.StrokeWidthStr = width
	return i
}

func (i *Icon) Class(className string) *Icon {
	cpy := i.Copy()
	if current, ok := cpy.AttributesMap["class"].(string); ok {
		cpy.AttributesMap["class"] = current + " " + className
	} else {
		cpy.AttributesMap["class"] = className
	}
	return cpy
}

// GLOBALS
func (i *Icon) GlobalSize(s Size) *Icon {
	i.applySize(s)
	return i
}

func (i *Icon) GlobalStroke(c string) *Icon {
	i.StrokeStr = c
	return i
}

func (i *Icon) GlobalStrokeWidth(w string) *Icon {
	i.StrokeWidthStr = w
	return i
}

func (i *Icon) GlobalAttributes(attrs templ.Attributes) *Icon {
	maps.Copy(i.AttributesMap, attrs)
	return i
}

func (i *Icon) GlobalClass(className string) *Icon {
	if current, ok := i.AttributesMap["class"].(string); ok {
		i.AttributesMap["class"] = current + " " + className
	} else {
		i.AttributesMap["class"] = className
	}
	return i
}

// helpers

func (i *Icon) applySize(s Size) {
	switch s {
	case SizeXXS:
		i.Width, i.Height = "6", "6"
	case SizeXS:
		i.Width, i.Height = "12", "12"
	case SizeSM:
		i.Width, i.Height = "16", "16"
	case SizeMD:
		i.Width, i.Height = "24", "24"
	case SizeLG:
		i.Width, i.Height = "32", "32"
	case SizeXL:
		i.Width, i.Height = "48", "48"
	case SizeXXL:
		i.Width, i.Height = "64", "64"
	default:
		i.Width, i.Height = "24", "24"
	}
}

func (i *Icon) Component() templ.Component {
	if i.componentFn == nil {
		return templ.NopComponent
	}
	return i.componentFn(i)
}
