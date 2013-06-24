package scr

// (c) Christian Maurer   v. 130126 - license see murus.go

import (
  "murus/ker"; "murus/xker"; "murus/col"; "murus/font"
)
var
  underX bool


func terminateX () {
//
// TODO wait (blink ())
// TODO terminate (blink ())
  terminated = true
  ker.Msleep (250) // provisorial
}


func initX () {
//
  visible = true
  mode = defaultMode
  XX, YY, colourdepth = xker.Init (nX[mode], nY[mode], col.ScreenF, col.ScreenB)
  maxMode = Mode (0)
  for {
    if XX == nX[maxMode] && YY == nY[maxMode] {
      break
    } else if maxMode + 1 < NModes {
      maxMode ++
    } else {
      mode = defaultMode
      XX, YY = nX[mode], nY[mode]
      break
    }
  }
  xker.Colours (col.ScreenF, col.ScreenB) // !
  xker.Clr (0, 0, nX[mode], nY[mode])
  swFontsize (font.Normal)
  go blink ()
// ker.InstallTerm (terminateX ())
}
