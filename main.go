package easywindow

import(
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/pixelgl"
  //"golang.com/x/image/colornames"
)
//これでpixelgl.Run(CreateWindowWithData)で作る
func CreateWindowWithData(title string, x1 float64, y1 float64, x2 float64, y2 float64, vsync bool) {
  pixelgl.NewWindow(pixelgl.WindowConfig{
    Title:title,
    Bounds:pixel.R(x1, y1, x2, y2),
    VSync:vsync,
  })

}
