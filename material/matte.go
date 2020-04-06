package material

import "image/color"

type Matte struct {
	Material
	Color color.NRGBA
}

func (m Matte) GetColor() color.NRGBA {
	return m.Color
}
