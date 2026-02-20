package icon_test

import (
	"testing"

	"github.com/iridiumgo/iridium-icons/icon"
	"github.com/iridiumgo/iridium-icons/icon/icons"
)

func BeforeEach() {
	icons.Angry = icon.NewIcon(icons.AngryComponent)
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestIconsSettingChanged(t *testing.T) {
	BeforeEach()
	i := icons.Angry
	if i.Height != "16" && i.Width != "16" {
		t.Errorf(`i.Height != "16"  && i.Width != "16"`)
	}
	i.Size("xl")
	if i.Height != "32" && i.Width != "32" {
		t.Errorf(`i.Height != "32"  && i.Width != "32"`)
	}
}

func TestIconsCopy(t *testing.T) {
	BeforeEach()

	i1 := icons.Angry.Copy().Size("xl")
	i2 := icons.Angry

	if i1.Height != "32" && i1.Width != "32" {
		t.Errorf(`i1.Height != "32"  && i1.Width != "32"`)
	}
	if i2.Height == "32" && i2.Width == "32" {
		t.Errorf(`i2.Height == "32"  && i2.Width == "32"`)
	}
}

func TestIconCopyAfterGlobalChange(t *testing.T) {
	BeforeEach()

	i1 := icons.Angry.Size("xl")
	i2 := icons.Angry

	if i1.Height != "32" && i1.Width != "32" {
		t.Errorf(`i1.Height != "32"  && i1.Width != "32"`)
	}
	if i2.Height != "32" && i2.Width != "32" {
		t.Errorf(`i2.Height == "32"  && i2.Width == "32"`)
	}
}
