package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	"main/package"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pong",
		Bounds: pixel.R(0, 0, 1850, 1050),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	for !win.Closed() {
		if _package.PointEngine >= 2 || _package.PointPlayer >= 2 {
			win.SetClosed(true)
		} else {
			win.Clear(colornames.Black)

			_package.Rectangle(imd, _package.Rectx[0], _package.Recty[0])
			_package.Rectangle(imd, _package.Rectx[1], _package.Recty[1])

			_package.BallRemove(imd, _package.Ballx, _package.Bally)
			movingRectangle0(imd, win, 0)

			_package.Ballx = _package.Ballx + _package.Vectorx
			_package.Bally = _package.Bally + _package.Vectory
			_package.RectangleColision(0)
			_package.RectangleColision(1)
			_package.WallColision()
			engine(imd)
			_package.Ball(imd, _package.Ballx, _package.Bally)
			go pointTracker(win, imd)

			imd.Draw(win)
			win.Update()
		}
	}
}

func movingRectangle0(imd *imdraw.IMDraw, win *pixelgl.Window, rectangleNumber int) {
	if win.Pressed(pixelgl.KeyW) && _package.Recty[rectangleNumber] < 1050 {
		_package.RectangleRemove(imd, _package.Rectx[rectangleNumber], _package.Recty[rectangleNumber])
		_package.Recty[rectangleNumber] += 10
		_package.Rectangle(imd, _package.Rectx[rectangleNumber], _package.Recty[rectangleNumber])
	}
	if win.Pressed(pixelgl.KeyS) && _package.Recty[rectangleNumber] > 175 {
		_package.RectangleRemove(imd, _package.Rectx[rectangleNumber], _package.Recty[rectangleNumber])
		_package.Recty[rectangleNumber] -= 10
		_package.Rectangle(imd, _package.Rectx[rectangleNumber], _package.Recty[rectangleNumber])
	}
}

func engine(imd *imdraw.IMDraw) {
	if _package.Recty[1]-75 < _package.Bally {

		_package.RectangleRemove(imd, _package.Rectx[1], _package.Recty[1])
		_package.Recty[1] += _package.EngineSpeed
		_package.Rectangle(imd, _package.Rectx[1], _package.Recty[1])

	} else {
		_package.RectangleRemove(imd, _package.Rectx[1], _package.Recty[1])

		_package.Recty[1] -= _package.EngineSpeed
		_package.Rectangle(imd, _package.Rectx[1], _package.Recty[1])

	}

}
func pointTracker(win *pixelgl.Window, imd *imdraw.IMDraw) {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	score := text.New(pixel.V(875, 900), atlas)

	fmt.Fprintln(score, _package.PointPlayer, "		", _package.PointEngine)
	score.Draw(win, pixel.IM)
	if _package.Ballx < _package.Rectx[0]-100 {
		_package.BallRemove(imd, _package.Ballx, _package.Bally)
		_package.PointEngine++
		_package.Ballx = 920
		_package.Bally = 525

	} else if _package.Ballx > _package.Rectx[1]+100 {
		_package.BallRemove(imd, _package.Ballx, _package.Bally)
		_package.PointPlayer++
		_package.Ballx = 920
		_package.Bally = 525

	}
}
