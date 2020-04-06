package material

import "image/color"

type Light struct {
	Material
	Color color.NRGBA
}

func (m Light) GetColor() color.NRGBA {
	return m.Color
}
