package scr

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  "murus/xker"; "murus/mouse"
  "murus/col"
)


func mouseEx () bool {
//
  if underX {
    return true // whoever drives X, has got a mouse
  }
  return mouse.Ex ()
}


func mouseDef (L, C, W, H uint) {
//
  x:= int(actualCharwidth * C)
  y:= int(actualCharheight * L)
  MouseDefGr (x, y, x + int(actualCharwidth * W) - 1, y + int(actualCharheight * H) - 1)
}

func mouseDefGr (x, y, w, h int) {
//
  if x < 0 || y < 0 || w <= 0 || h <= 0 { return }
  if underX {
    xker.MouseDef (x, y, w, h)
    return
  }
  mouse.Def (uint(x), uint(y), uint(w), uint(h))
}


func mousePos () (uint, uint) {
//
  x, y:= MousePosGr()
  return uint(y) / actualCharheight, uint(x) / actualCharwidth
}


func mousePosGr () (int, int) {
//
  if underX {
    return xker.MousePos ()
  }
  return mouse.Pos ()
}


func warpMouse (L, C uint) {
//
  WarpMouseGr (int(actualCharwidth * C), int(actualCharheight * L))
}


func warpMouseGr (x, y int) {
//
  if underX {
    xker.WarpMouse (x, y)
    return
  }
  mouse.Warp (uint(x), uint(y))
  MouseCursor (true)
}


const (
  pointerH = 18
  pointerW = 10
)
var (
  pointer [pointerH]string
  mouseOn bool
  xMouse, yMouse int
)


func initMouse () {
//
  pointer [0] = "# . . . . . . . . ."
  pointer [1] = "# # . . . . . . . ."
  pointer [2] = "# * # . . . . . . ."
  pointer [3] = "# * * # . . . . . ."
  pointer [4] = "# * * * # . . . . ."
  pointer [5] = "# * * * * # . . . ."
  pointer [6] = "# * * * * * # . . ."
  pointer [7] = "# * * * * * * # . ."
  pointer [8] = "# * * * * * * * # ."
  pointer [9] = "# * * * * * # # # #"
  pointer[10] = "# * * * * * # . . ."
  pointer[11] = "# * # # * * # . . ."
  pointer[12] = "# # . # * * * # . ."
  pointer[13] = "# . . # * * * # . ."
  pointer[14] = ". . . . # * * * # ."
  pointer[15] = ". . . . # * * * # ."
  pointer[16] = ". . . . . # * # . ."
  pointer[17] = ". . . . . # # . . ."
  mouseOn = false
  mouse.Def (0, 0, XX, YY)
  mouse.Warp (nX [mode] / 3, nY [mode] / 3)
}


func setMousepoint (x, y int) {
//
  if visible {
    if ! (x < int(nX [mode]) && y < int(nY [mode])) { return }
    a:= (int(XX) * y + x) * int(colourdepth)
    copy (fbmem[a:a+int(colourdepth)], cc (col.CodeF))
  }
}


func mousePointer (x, y int) {
//
  F:= col.ActualF
  CV:= col.CodeF
  Colour (col.Black)
  for z:= 0; z < pointerH; z++ {
    for s:= 0; s < pointerW; s++ {
      if pointer[z][2 * s] == '#' {
        setMousepoint (x + s, y + z)
      }
    }
  }
  Colour (col.LightWhite)
  for z:= 0; z < pointerH; z++ {
    for s:= 0; s < pointerW; s++ {
      if pointer[z][2 * s] == '*' {
        setMousepoint (x + s, y + z)
      }
    }
  }
  col.ActualF = F
  col.CodeF = CV
}


func mouseCursor (on bool) {
//
  mouseOn = on
  if underX {
    return // TODO
  }
  if ! mouse.Ex () {
    return
  }
  if mouseOn {
    ar (fbcop, fbmem, nil, uint(xMouse), uint(yMouse), uint(xMouse + pointerW - 1), uint(yMouse + pointerH - 1))
    xMouse, yMouse = mouse.Pos ()
    mousePointer (xMouse, yMouse)
  }
}


func mouseCursorOn () bool {
//
  if underX {
    return true // TODO
  }
  if ! mouse.Ex () {
    return false
  }
  return mouseOn
}


func underMouse (L, C, W, H uint) bool {
//
  if underX {
    // we assume, a mouse is mousing
  } else {
    if ! mouse.Ex () {
      return false
    }
  }
  l, c:= MousePos ()
  return L <= l && l < L + H &&
         C <= c && c < C + W
}


func underMouseGr (x, y, x1, y1 int, T uint) bool {
//
  if underX {
    // we assume, a mouse is mousing
  } else {
    if ! mouse.Ex () {
      return false
    }
  }
  xm, ym:= MousePosGr ()
  t:= int(T)
  intord (&x, &y, &x1, &y1)
  return x <= xm + t && xm <= x1 + t &&
         y <= ym + t && ym <= y1 + t
}
