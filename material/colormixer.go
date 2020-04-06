package material

import "image/color"

type ColorMixer struct {
	R, G, B, A int
	n          int
}

func (cm *ColorMixer) Add(c color.NRGBA) {
	cm.R += int(c.R)
	cm.G += int(c.G)
	cm.B += int(c.B)
	cm.A += int(c.A)
	cm.n += 1
}

func (cm *ColorMixer) ToNRGBA() color.NRGBA {
	return color.NRGBA{
		uint8(cm.R / cm.n),
		uint8(cm.G / cm.n),
		uint8(cm.B / cm.n),
		uint8(cm.A / cm.n)}
}
