package _package

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

func Rectangle(imd *imdraw.IMDraw, x, yT float64) {
	imd.Color = pixel.RGB(0, 255, 0)
	imd.Push(pixel.V(x, yT))
	imd.Color = pixel.RGB(0, 0, 1)
	imd.Push(pixel.V(x, yT-150))
	imd.Rectangle(20)
}
func RectangleRemove(imd *imdraw.IMDraw, x, yT float64) {
	imd.Color = pixel.RGB(0, 0, 0)
	imd.Push(pixel.V(x, yT))
	imd.Color = pixel.RGB(0, 0, 0)
	imd.Push(pixel.V(x, yT-150))
	imd.Rectangle(20)
}
func Ball(imd *imdraw.IMDraw, x, y float64) {
	imd.Color = pixel.RGB(240, 230, 140)
	imd.Push(pixel.V(x, y))
	imd.Circle(5, 10)
}
func BallRemove(imd *imdraw.IMDraw, x, y float64) {
	imd.Color = pixel.RGB(0, 0, 0)
	imd.Push(pixel.V(x, y))
	imd.Circle(5, 10)
}
