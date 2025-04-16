package sync

import "fmt"

type Color string

var (
	BerryRed   Color = "berry_red"
	Red        Color = "red"
	Orange     Color = "orange"
	Yellow     Color = "yellow"
	OliveGreen Color = "olive_green"
	LimeGreen  Color = "lime_green"
	Green      Color = "green"
	MintGreen  Color = "mint_green"
	Teal       Color = "teal"
	SkyBlue    Color = "sky_blue"
	LightBlue  Color = "light_blue"
	Blue       Color = "blue"
	Grape      Color = "grape"
	Violet     Color = "violet"
	Lavender   Color = "lavender"
	Magenta    Color = "magenta"
	Salmon     Color = "salmon"
	Charcoal   Color = "charcoal"
	Grey       Color = "grey"
	Taupe      Color = "taupe"
)

type rgb struct {
	r, g, b uint8
}

var colorMap = map[Color]struct {
	hex string
	rgb rgb
}{
	BerryRed:   {"#B8255F", rgb{0xB8, 0x25, 0x5F}},
	Red:        {"#DC4C3E", rgb{0xDC, 0x4C, 0x3E}},
	Orange:     {"#C77100", rgb{0xC7, 0x71, 0x00}},
	Yellow:     {"#B29104", rgb{0xB2, 0x91, 0x04}},
	OliveGreen: {"#949C31", rgb{0x94, 0x9C, 0x31}},
	LimeGreen:  {"#65A33A", rgb{0x65, 0xA3, 0x3A}},
	Green:      {"#369307", rgb{0x36, 0x93, 0x07}},
	MintGreen:  {"#42A393", rgb{0x42, 0xA3, 0x93}},
	Teal:       {"#148FAD", rgb{0x14, 0x8F, 0xAD}},
	SkyBlue:    {"#319DC0", rgb{0x31, 0x9D, 0xC0}},
	LightBlue:  {"#6988A4", rgb{0x69, 0x88, 0xA4}},
	Blue:       {"#4180FF", rgb{0x41, 0x80, 0xFF}},
	Grape:      {"#692EC2", rgb{0x69, 0x2E, 0xC2}},
	Violet:     {"#CA3FEE", rgb{0xCA, 0x3F, 0xEE}},
	Lavender:   {"#A4698C", rgb{0xA4, 0x69, 0x8C}},
	Magenta:    {"#E05095", rgb{0xE0, 0x50, 0x95}},
	Salmon:     {"#C9766F", rgb{0xC9, 0x76, 0x6F}},
	Charcoal:   {"#808080", rgb{0x80, 0x80, 0x80}},
	Grey:       {"#999999", rgb{0x99, 0x99, 0x99}},
	Taupe:      {"#8F7A69", rgb{0x8F, 0x7A, 0x69}},
}

func ParseColor(s string) (Color, error) {
	if _, ok := colorMap[Color(s)]; !ok {
		return "", fmt.Errorf("invalid color %s", s)
	}
	return Color(s), nil
}

func ListColors() []Color {
	colors := make([]Color, len(colorMap))
	i := 0
	for k := range colorMap {
		colors[i] = k
		i++
	}
	return colors
}

func (c Color) Hex() string {
	return colorMap[c].hex
}

func (c Color) RGB() (int, int, int) {
	rgb := colorMap[c].rgb
	return int(rgb.r), int(rgb.g), int(rgb.b)
}
