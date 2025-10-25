package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Images
// image パッケージは、以下の Image インタフェースを定義している
// ```
// package image

// type Image interface {
//     ColorModel() color.Model
//     Bounds() Rectangle
//     At(x, y int) color.Color
// }
// ```
// @ref: https://pkg.go.dev/image#Image

type Image struct{}

func (Image) ColorModel() color.Model {
	// ColorModel は RGBA カラーモデルを返す
	return color.RGBAModel
}

func (Image) Bounds() image.Rectangle {
	// Bounds は画像の範囲を返す（今回は固定サイズ 256x256）
	return image.Rect(0, 0, 256, 256)
	// x0 int, y0 int, x1 int, y1 int
}

// At は指定された (x, y) のピクセルの色を返す
func (Image) At(x, y int) color.Color {
	v := uint8((x ^ y)) // パターン生成好きに変えてOK
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}