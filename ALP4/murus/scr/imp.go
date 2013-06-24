package scr

// (c) Christian Maurer   v. 130126 - license see murus.go

import (
  "murus/xker"; "murus/col"
)
const (
  pack = "scr"
  esc1 = "\x1b["
  clearScreen = esc1 + "H" + esc1 + "J"
)
type
  pointtype byte; const (pt = iota; ptinv; onpt)
var (
  XX, YY = uint(0), uint(0)
  nLines, nColumns uint
  actualCharwidth, actualCharheight uint = 8, 16 // charwidth, charheight (normal)
  fbmem, fbcop, emptyBackground, archive []byte
)


func natord (x, y, x1, y1 *uint) {
//
  if *x > *x1 { *x, *x1 = *x1, *x }
  if *y > *y1 { *y, *y1 = *y1, *y }
}


func rectangleOk (x, y, x1, y1 *uint) bool {
//
  if ! underX {
    if ! visible {
      return false
    }
  }
  natord (x, y, x1, y1)
  if *x >= nX [mode] || *y >= nY [mode] {
    return false
  }
  if *x1 >= nX [mode] { *x1 = nX [mode] - 1 }
  if *y1 >= nY [mode] { *y1 = nY [mode] - 1 }
  return true
}


func cc (n uint) []byte {
//
  b:= make([]byte, colourdepth)
  for i:= 0; i < int(colourdepth); i++ {
    b[i] = byte(n)
    n >>= 8
  }
  return b
}

/*
// Pre: len(B) == colourdepth
func cd (B []byte) uint {
//
  n:= uint(0)
  for i:= int(colourdepth) - 1; i >= 0; i-- {
    n = n * 256 + uint(B[i])
  }
  return n
}
*/

func cls () {
//
  Colours (col.ScreenF, col.ScreenB)
  if underX {
    xker.Clr (0, 0, nX [mode], nY [mode])
    return
  }
  if ! visible { return }
  B:= cc (col.Code (col.ScreenB))
  l:= int(colourdepth) * int(nX [mode])
  a:= 0
  for y:= 0; y < int(nY [mode]); y++ {
    for x:= 0; x < int(nX [mode]); x++ {
      copy (emptyBackground[a:a+int(colourdepth)], B)
      a += int(colourdepth)
    }
  }
  a = 0
  for y:= 0; y < int(nY [mode]); y++ {
    copy (fbmem[a:a+l], emptyBackground)
    copy (fbcop[a:a+l], emptyBackground)
    a += int(XX) * int(colourdepth)
  }
  setColours (col.ScreenB)
}


func clear (L, C, W, H uint) {
//
  Colours (col.ScreenF, col.ScreenB)
  x:= C * actualCharwidth
  y:= L * actualCharheight
  if underX {
    xker.Clr (x, y, W * actualCharwidth, H * actualCharheight)
  } else {
    ClearGr (x, y, x + W * actualCharwidth - 0 /* 0 ? */, y + H * actualCharheight - 1)
  }
}


func clearGr (x, y, x1, y1 uint) {
//
  if ! rectangleOk (&x, &y, &x1, &y1) { return }
  if underX {
    xker.Clr (x, y, x1 - x + 1, y1 - y + 1)
  } else {
    if ! visible { return }
    dx:= int(x1 - x) * int(colourdepth)
    dy:= int(nX [maxMode]) * int(colourdepth)
    a:= int(y * nX [maxMode] + x) * int(colourdepth)
    for z:= uint(0); z <= y1 - y; z++ {
      copy (fbmem[a:a+dx], emptyBackground[0:dx])
      a += dy
    }
  }
}


func invert (L, C, W, H uint) {
//
  x:= actualCharwidth * C
  y:= actualCharheight * L
  invertGr (x, y, x + actualCharwidth * W - 1, y + actualCharheight * H - 1)
}


func invertGr (x, y, x1, y1 uint) {
//
  if ! rectangleOk (&x, &y, &x1, &y1) {
    return
  }
  if underX {
    xker.Invert (x, y, x1, y1)
  } else {
    for z:= y; z <= y1; z++ {
      invertHorizontal (x, z, x1)
    }
  }
}


func ar (s, d, d1 []byte, x0, y0, x1, y1 uint) {
//
  if underX { return }
  if ! rectangleOk (&x0, &y0, &x1, &y1) { return }
  a:= int(XX * y0 + x0) * int(colourdepth)
  da:= int(x1 - x0 + 1) * int(colourdepth)
  i:= int(nX [maxMode]) * int(colourdepth)
  for k:= y0; k <= y1; k++ {
    copy (d[a:a+da], s[a:a+da])
    if d1 != nil {
      copy (d1[a:a+da], s[a:a+da])
    }
    a += i
  }
}


var
  only2Buf bool


func buf (on bool) {
//
  if on == only2Buf { return }
  only2Buf = on
  if underX {
    xker.Buf (on)
    return
  }
  a:= 0
  da:= int(colourdepth) * int(nX [maxMode])
  b:= int(colourdepth) * int(nX [mode])
  for y:= 0; y < int(nY [mode]); y++ {
    if on {
      copy (fbcop[a:a+b], emptyBackground[0:b])
    } else {
      copy (fbmem[a:a+b], fbcop[a:a+b])
    }
    a += da
  }
}


func save (L, C, W, H uint) {
//
  x:= actualCharwidth * C
  y:= actualCharheight * L
  SaveGr (x, y, x + actualCharwidth * W, y + actualCharheight * H)
}


func saveGr (x, y, x1, y1 uint) {
//
  if underX {
    xker.Save (x, y, x1 - x + 1, y1 - y + 1)
  } else {
    natord (&x, &y, &x1, &y1)
    if ! visible { return }
    if x + y == 0 && x1 == nX [mode] && y1 == nY [mode] {
      c:= int(colourdepth) * int(nX [maxMode] * nY [mode])
      copy (archive[0:c], fbmem[0:c])
    } else {
      if mouseOn { MouseCursor (false) }
        ar (fbcop /* fbmem */, archive, nil, x, y, x1, y1)
      if mouseOn { MouseCursor (true) }
    }
  }
}


func restore (L, C, W, H uint) {
//
  x:= actualCharwidth * C
  y:= actualCharheight * L
  RestoreGr (x, y, x + actualCharwidth * W, y + actualCharheight * H)
}


func restoreGr (x, y, x1, y1 uint) {
//
  if underX {
    xker.Restore (x, y, x1 - x + 1, y1 - y + 1)
  } else {
    natord (&x, &y, &x1, &y1)
    if ! visible { return }
    if x + y == 0 && x1 == nX [mode] && y1 == nY [mode] {
      c:= int(colourdepth) * int(nX [maxMode] * nY [mode])
      copy (fbmem[0:c], archive[0:c])
      copy (fbcop[0:c], archive[0:c])
    } else {
      ar (archive, fbmem, fbcop, x, y, x1, y1)
    }
  }
}


func init () {
//
  defaultMode = VGA
  initModes ()
  underX = xker.Active ()
  if underX {
    initX ()
  } else {
    initConsole ()
  }
//  col.Init()
  setColours = nixfaerben
}
