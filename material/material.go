package material

import "image/color"

type Material interface {
	GetColor() color.NRGBA
}
