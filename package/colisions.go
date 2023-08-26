package _package

func RectangleColision(rectangleIndex int) {
	if Ballx == Rectx[rectangleIndex] && Bally < Recty[rectangleIndex] && Bally > Recty[rectangleIndex]-150 {
		Vectorx = -Vectorx

	}
}
func WallColision() {
	if Bally > 1050 || Bally < 30 {
		Vectory = -Vectory
	}
}
