package scr

// (c) Christian Maurer   v. 121217 - license see murus.go

import (
  "math"
  "murus/ker"; "murus/xker"
  "murus/col"
)


func iok (x, y int) bool {
//
  if x < 0 || y < 0 {
    return false
  }
  return uint(x) < nX [mode] && uint(y) < nY [mode]
}


func point (X, Y int) {
//
  if underX {
    xker.Point (X, Y, true)
    return
  }
  if ! visible || ! iok (X, Y) {
    return
  }
  x:= uint(X)
  y:= uint(Y)
  if (x >= nX [mode]) || (y >= nY [mode]) {
    return
  }
  a:= (int(XX) * Y + X) * int(colourdepth)
  c:= cc (col.CodeF)
  if ! only2Buf {
    copy (fbmem[a:a+int(colourdepth)], c)
/* TODO
    if actualLinewidth > Thin && x + 1 < nX [mode] && y + 1 < nY [mode] {
      if x + 1 < nY [mode] {
        M += colourdepth++
        copy (M, &Colours.CodeF, colourdepth)
      }
      if y + 1 < nX [mode] {
        M += (XX - 1) * colourdepth
        copy (M, &Colours.CodeF, colourdepth)
      }
      if actualLinewidth == Thicker {
        M, colourdepth++
        copy (M, &Colours.CodeF, colourdepth)
      } else { // Yetthicker
        if x > 0 && y > 0 {
          M -= XX * 2 * colourdepth
          copy (M, &Colours.CodeF, colourdepth)
          M += (XX - 1) * colourdepth
          copy (M, &Colours.CodeF, colourdepth)
        }
      }
    }
*/
  }
  copy (fbcop[a:a+int(colourdepth)], c)
/* TODO
  if actualLinewidth > Thin &&
    x + 1 < nX [mode] && y + 1 < nY [mode] {
    M += colourdepth
    copy (M, &Colours.CodeF), colourdepth)
    M += (XX - 1) * colourdepth
    copy (M, &Colours.CodeF, colourdepth)
    if actualLinewidth == Thicker {
      M, colourdepth++
      copy (M, &Colours.CodeF, colourdepth)
    } else { // Yetthicker
      if x > 0 && y > 0 {
        M -= XX * 2 * colourdepth
        copy (M, &Colours.CodeF, colourdepth)
        M += (XX - 1) * colourdepth
        copy (M, &Colours.CodeF, colourdepth)
      }
    }
  }
*/
}


func ok2 (x, y []int) bool {
//
  return len (x) == len (y)
}


func ok4 (x, y, x1, y1 []int) bool {
//
  return len (x) == len (y) &&
         len (x1) == len (y1) &&
         len (x) == len (x1)
}


func pointset (x, y []int) {
//
  if ! ok2 (x, y) { return }
  n:= uint(len (x))
  if n == 0 { return }
  if underX {
    xker.Points (x, y, true)
    return
  }
  if ! visible { return }
  for i:= 0; i < int(n); i++ {
    point (x [i], y [i])
  }
}


func pointsetInv (x, y []int) {
//
  if ! ok2 (x, y) { return }
  n:= uint(len (x))
  if n == 0 { return }
  if underX {
    xker.Points (x, y, true)
    return
  }
  if ! visible { return }
  for i:= 0; i < int(n); i++ {
    pointInv (x [i], y [i])
  }
}


func pointColour (x, y uint) col.Colour {
//
  if x >= nX [mode] || y >= nY [mode] {
    return col.ScreenB
  }
  if underX {
    return xker.Colour (x, y)
  }
  if ! visible {
    return col.ScreenB
  }
  a:= int(XX * y + x) * int(colourdepth)
  return col.P6Colour (fbcop [a:a+int(colourdepth)])
}


func pointInv (X, Y int) {
//
  if underX {
    xker.Point (X, Y, false)
    return
  }
  if ! iok (X, Y) { return }
  x, y:= uint(X), uint(Y)
  AF:= col.ActualF
  CF:= col.CodeF
  CB:= col.CodeB
  C:= PointColour (x, y)
  col.Invert (&C)
  Colour (C)
  Point (int(x), int(y))
  col.ActualF = AF
  col.CodeF = CF
  col.CodeB = CB
}


// Pre: x <= x1 < NX, y < NY.
func writeHorizontal (x, y, x1 uint) {
//
  if x == x1 { Point (int(x), int(y)); return }
  if x > x1 { x = x1; x1 = x }
  if x >= nX [mode] { return }
  if x1 >= nX [mode] { x1 = nX [mode] - 1 }
  a:= int(XX * y + x) * int(colourdepth)
  c:= cc (col.CodeF)
  for i:= x; i <= x1; i++ {
    if ! only2Buf {
      copy (fbmem[a:a+int(colourdepth)], c)
    }
    copy (fbcop[a:a+int(colourdepth)], c)
    a += int(colourdepth)
  }
/*
  if actualLinewidth > Thin {
    if y + 1 <= nY [mode] {
      a = (XX * (y + 1) + x) * colourdepth
      if ! only2Buf {
        M = fbmem + a
        for i:= x; x1 {
          copy (M, &Colours.CodeF, colourdepth)
          a += colourdepth
        }
      }
      M = fbcop + a
      for i:= x; x1 {
        copy (M, &Colours.CodeF, colourdepth)
        a += colourdepth
      }
    }
  }
  if actualLinewidth > Thicker {
    if y > 0 {
      a = (XX * (y - 1) + x) * colourdepth
      if ! only2Buf {
        M = fbmem + a
        for i:= x; i <= x1; i++ {
          copy (M, &Colours.CodeF, colourdepth)
          a += colourdepth
        }
      }
      M = fbcop + a
      for i:= x; i <= x1; i++ {
        copy (M, &Colours.CodeF, colourdepth)
        M += colourdepth
      }
    }
  }
*/
}


// Pre: x <= x1 < NX, y < NY.
func invertHorizontal (x, y, x1 uint) {
//
  for x <= x1 {
    PointInv (int(x), int(y))
    x++
  }
}


func horizontal (x, y, x1 uint, p pointtype) {
//
  if x > x1 { x = x1; x1 = x }
  switch p { case pt:
    writeHorizontal (x, y, x1)
  case ptinv:
    invertHorizontal (x, y, x1)
  case onpt:
    for x <= x1 {
      on (int(x), int(y))
      x++
    }
  }
}


// Pre: x < NX, y <= y1 < NY.
func writeVertical (x, y, y1 uint) {
//
  if y == y1 {
    Point (int(x), int(y))
    return
  }
  if y > y1 { y = y1; y1 = y }
  if y1 >= nY [mode] { y1 = nY [mode] - 1 }
  a:= int(XX * y + x) * int(colourdepth)
  c:= cc (col.CodeF)
  for i:= y; i <= y1; i++ {
    if ! only2Buf {
      copy (fbmem[a:a+int(colourdepth)], c)
    }
    copy (fbcop[a:a+int(colourdepth)], c)
    a += int(colourdepth) * int(XX)
  }
/*
  if actualLinewidth > Thin {
    if x + 1 < nX [mode] {
      a = (XX * y + x + 1) * colourdepth
      if ! only2Buf {
        M = fbmem + a
        for i:= y; y1 {
        copy (M, &Colours.CodeF, colourdepth)
          M += colourdepth * XX
        }
      }
      M = fbcop + a
      for i:= y; y1 {
        copy (M, &Colours.CodeF, colourdepth)
        M += colourdepth * XX
      }
    }
  }
  if actualLinewidth > Thicker {
    if x > 0 {
      a = (XX * y + x - 1) * colourdepth
      if ! only2Buf {
        M = fbmem + a
        for i:= y; y1 {
        copy (M, &Colours.CodeF, colourdepth)
          M += colourdepth * XX
        }
      }
      M = fbcop + a
      for i:= y; y1 {
        copy (M, &Colours.CodeF, colourdepth)
        M += colourdepth * XX
      }
    }
  }
*/
}


func vertical (x, y, y1 uint, p pointtype) {
//
  if y > y1 { y, y1 = y1, y }
  if p == pt {
    writeVertical (x, y, y1)
  } else {
    for y <= y1 {
      PointInv (int(x), int(y))
      y++
    }
  }
}


// Pre: 0 <= x <= x1 < NColumns, 0 <= y != y1 < NLines.
func bresenham (x, y, x1, y1 int, p pointtype) {
//
  var A application
  switch p { case pt:
    A = Point
  case ptinv:
    A = PointInv
  case onpt:
    A = on
  }
  dx:= x1 - x
  Fehler:= 0
  var dy int
  if y <= y1 { // Steigung positiv
    dy = y1 - y
    if dy <= dx { // Steigung <= 45 Grad
      for {
        A (x, y)
        if x == x1 {
          break
        }
        x++
        Fehler += 2 * dy
        if Fehler > dx {
          y++
          Fehler -= 2 * dx
        }
      }
    } else { // Steigung > 45 Grad
      for {
        A (x, y)
        if y == y1 {
          break
        }
        y++
        Fehler += 2 * dx
        if Fehler > dy {
          x++
          Fehler -= 2 * dy
        }
      }
    }
  } else { // Steigung negativ
    dy = y - y1
    if dy <= dx { // Steigung >= -45 Grad
      for {
        A (x, y)
        if x == x1 {
          break
        }
        x++
        Fehler += 2 * dy
        if Fehler > dx {
          y--
          Fehler -= 2 * dx
        }
      }
    } else { // Steigung < -45 Grad
      for {
        A (x, y)
        if y == y1 {
          break
        }
        y--
        Fehler += 2 * dx
        if Fehler > dy {
          x++
          Fehler -= 2 * dy
        }
      }
    }
  }
}


// Pre: 0 <= x <= x1 < xx, y != y1, 0 <= y, y1 < yy.
func bresenhamInf (xx, yy, x, y, x1, y1 int, p pointtype) {
//
  var A application
  switch p { case pt:
    A = Point
  case ptinv:
    A = PointInv
  case onpt:
    A = on
  }
  dx:= x1 - x
  Fehler:= 0
  x0:= x
  y0:= y
  var dy int
  if y <= y1 { // Steigung positiv
    dy = y1 - y
    if dy <= dx { // Steigung Less Eq 45 Grad
      for {
        A (x, y)
        if x == xx - 1 || y == yy - 1 {
          break
        }
        x++
        Fehler += 2 * dy
        if Fehler > dx {
          y++
          Fehler -= 2 * dx
        }
      }
      x = x0
      y = y0
      Fehler = 0
      for {
        Fehler += 2 * dy
        if Fehler > dx {
          y--
          Fehler -= 2 * dx
        }
        A (x, y)
        if x == 0 || y == 0 {
          break
        }
        x--
      }
    } else { // Steigung > 45 Grad
      for {
        A (x, y)
        if y == yy - 1 || x == xx - 1 {
          break
        }
        y++
        Fehler += 2 * dx
        if Fehler > dy {
          x++
          Fehler -= 2 * dy
        }
      }
      x = x0
      y = y0
      Fehler = 0
      for {
        Fehler += 2 * dx
        if Fehler > dy {
          x--
          Fehler -= 2 * dy
        }
        A (x, y)
        if x == 0 || y == 0 {
          break
        }
        y--
      }
    }
  } else {
    dy = y - y1 // Steigung negativ
    if dy <= dx { // Steigung >= -45 Grad
      for {
        A (x, y)
        if (x == xx - 1) || (y == 0) {
          break
        }
        x++
        Fehler += 2 * dy
        if Fehler > dx {
          y--
          Fehler -= 2 * dx
        }
      }
      x = x0
      y = y0
      Fehler = 0
      for {
        A (x, y)
        if x == 0 || y == yy - 1 {
        break
        }
        x--
        Fehler += 2 * dy
        if Fehler > dx {
          y++
          Fehler -= 2 * dx
        }
      }
    } else { // Steigung < -45 Grad
      for {
        A (x, y)
        if x == xx - 1 || y == 0 {
          break
        }
        y--
        Fehler += 2 * dx
        if Fehler > dy {
          x++
          Fehler -= 2 * dy
        }
      }
      x = x0
      y = y0
      Fehler = 0
      for {
        A (x, y)
        if x == 0 || y == yy - 1 {
          break
        }
        y++
        Fehler += 2 * dx
        if Fehler > dy {
          x--
          Fehler -= 2 * dy
        }
      }
    }
  }
}


func nat (x, y int) bool {
//
  return x >= 0 && y >= 0
}


func _line (x, y, x1, y1 int, p pointtype) {
//
  if x1 < x {
    x, x1 = x1, x
    y, y1 = y1, y
  }
  if underX {
    xker.Line (x, y, x1, y1, p == pt)
    return
  }
  if ! visible { return }
  if ! iok (x, y) || ! iok (x1, y1) { return }
  if y == y1 {
    horizontal (uint(x), uint(y), uint(x1), p)
    return
  }
  if x == x1 {
    vertical (uint(x), uint(y), uint(y1), p)
    return
  }
  bresenham (x, y, x1, y1, p)
}


func line (x, y, x1, y1 int) {
//
  if underX {
    xker.Line (x, y, x1, y1, true)
    return
  }
  if iok (x, y) && iok (x1, y1) {
    _line (x, y, x1, y1, pt)
  }
}


func lineInv (x, y, x1, y1 int) {
//
  if underX {
    xker.Line (x, y, x1, y1, false)
    return
  }
  if iok (x, y) && iok (x1, y1) {
    _line (x, y, x1, y1, ptinv)
  }
}


func between (i, k, m, t int) bool {
//
  return i <= m + t && m <= k + t ||
         k <= m + t && m <= i + t
}


var (
  xxx, yyy, ttt int
  incident bool
)


func onLine (x, y, x1, y1, a, b int, t uint) bool {
//
  if x1 < x {
    x, x1 = x1, x
    y, y1 = y1, y
  }
  if ! (between (x, x1, a, int(t)) && between (y, y1, b, int(t))) {
    return false
  }
  if x == x1 {
    return between (a, a, x, int(t))
  }
  if y == y1 {
    return between (b, b, y, int(t))
  }
  xxx, yyy, ttt, incident = a, b, int(t * t), false
  bresenham (x, y, x1, y1, onpt)
  return incident
}


func _lines (x, y, x1, y1 []int, p pointtype) {
//
  if ! ok4 (x, y, x1, y1) { return }
  if underX {
    xker.Segments (x, y, x1, y1, p == pt)
    return
  }
  if ! visible { return }
  for i:= 0; i < len (x); i++ {
    if x[i] < 0 || y[i] < 0 ||
      x[i] >= int(nX [mode]) || y[i] >= int(nY [mode]) {
      return
    }
  }
  for i:= 0; i < len (x); i++ {
    if iok (x[i], y[i]) && iok (x1[i], y1[i]) {
      _line (x[i], y[i], x1[i], y1[i], p)
    }
  }
}


func lines (x, y, x1, y1 []int) {
//
  if ! ok4 (x, y, x1, y1) { return }
  if len (x) == 0 { return }
  _lines (x, y, x1, y1, pt)
}



func linesInv (x, y, x1, y1 []int) {
//
  if ! ok4 (x, y, x1, y1) { return }
  if len (x) == 0 { return }
  _lines (x, y, x1, y1, ptinv)
}


func onLines (x, y, x1, y1 []int, a, b int, t uint) bool {
//
  if ! ok2 (x, y) { return false }
  if len (x) == 0 { return x[0] == a && y[0] == b }
  for i:= 0; i < len (x); i++ {
    if onLine (x[i], y[i], x1[i], y1[i], a, b, t) {
      return true
    }
  }
  return false
}


func _segs (x, y []int, p pointtype) {
//
  if ! ok2 (x, y) {
    return
  }
  if underX {
    xker.Lines (x, y, p == pt)
    return
  }
  if ! visible { return }
  n:= len (x)
  for i:= 0; i < n; i++ {
    if ! iok (x[i], y[i]) {
      return
    }
  }
  if n == 0 {
    if p == pt {
      point (x[0], y[0])
    } else {
      pointInv (x[0], y[0])
    }
  } else {
    for i:= 1; i < len (x); i++ {
      _line (x[i-1], y[i-1], x[i], y[i], p)
    }
  }
}


func segments (x, y []int) {
//
  if len (x) == 0 { return }
  _segs (x, y, pt)
}


func segmentsInv (x, y []int) {
//
  if ! ok2 (x, y) { return }
  if len (x) == 0 { return }
  _segs (x, y, ptinv)
// TODO the following is inefficient under X, should be smoothed:
  if len (x) > 1 {
    for i:= 1; i < len (x); i++ {
      pointInv (x[i], y[i])
    }
  }
}


func onSegments (x, y []int, a, b int, t uint) bool {
//
  if ! ok2 (x, y) { return false }
  if len (x) == 0 { return x[0] == a && y[0] == b }
  for i:= 1; i < len (x); i++ {
    if onLine (x[i-1], y[i-1], x[i], y[i], a, b, t) {
      return true
    }
  }
  return false
}


func on (x, y int) {
//
  incident = incident || ((x - xxx) * (x - xxx) + (y - yyy) * (y - yyy) <= ttt)
}


func _infLine (x, y, x1, y1 int, p pointtype) {
//
  if x == x1 && y == y1 { return }
  if x1 < x {
    x, x1 = x1, x
    y, y1 = y1, y
  }
  if underX {
    // TODO what ?
    return
  }
  if ! visible { return }
  if y == y1 {
    horizontal (0, uint(y), XX - 1, p)
    return
  }
  if x == x1 {
    vertical (uint(x), 0, YY, p)
    return
  }
  bresenhamInf (int(XX), int(YY), x, y, x1, y1, p)
}


func infLine (x, y, x1, y1 int) {
//
  _infLine (x, y, x1, y1, pt)
}


func infLineInv (x, y, x1, y1 int) {
//
  _infLine (x, y, x1, y1, ptinv)
}


func onInfLine (x, y, x1, y1, a, b int, t uint) bool {
//
  if x1 < x {
    x, x1 = x1, x
    y, y1 = y1, y
  }
  xxx, yyy, ttt, incident = a, b, int(t * t), false
  bresenhamInf (int(XX), int(YY), x, y, x1, y1, onpt)
  return incident
}


func intord (x, y, x1, y1 *int) {
//
  if *x > *x1 { *x, *x1 = *x1, *x }
  if *y > *y1 { *y, *y1 = *y1, *y }
}


func rectang (x, y, x1, y1 uint, p pointtype) {
//
  if underX { return }
  if ! rectangleOk (&x, &y, &x1, &y1) { return }
  if x == x1 {
    if y == y1 {
      if p == pt {
        Point (int(x), int(y))
      } else {
        PointInv (int(x), int(y))
      }
    } else {
      vertical (x, y, y1, p)
    }
    return
  }
  horizontal (x, y, x1, p)
  if y == y1 {
    return
  }
  horizontal (x, y1, x1, p)
  vertical (x, y, y1, p)
  vertical (x1, y, y1, p)
}


func rectangle (x, y, x1, y1 int) {
//
  if underX {
    intord (&x, &y, &x1, &y1)
    xker.Rectangle (x, y, x1 - x, y1 - y, true, false)
    return
  }
  if ! iok (x, y) || ! iok (x1, y1) {
    return
  }
  rectang (uint(x), uint(y), uint(x1), uint(y1), pt)
}


func rectangleInv (x, y, x1, y1 int) {
//
  if underX {
    intord (&x, &y, &x1, &y1)
    xker.Rectangle (x, y, x1 - x, y1 - y, false, false)
    return
  }
  if ! iok (x, y) || ! iok (x1, y1) { return }
  rectang (uint(x), uint(y), uint(x1), uint(y1), ptinv)
  pointInv (x, y)
  pointInv (x1, y)
  pointInv (x, y1)
  pointInv (x1, y1)
}


func rectangleFull (x, y, x1, y1 int) {
//
  intord (&x, &y, &x1, &y1)
  if underX {
    xker.Rectangle (x, y, x1 - x + 1, y1 - y + 1, true, true)
    return
  }
  if ! visible { return }
  if ! iok (x, y) || ! iok (x1, y1) { return }
  if uint(x1) >= nX [mode] {
    x1 = int(nX [mode]) - 1
  }
  if uint(y1) >= nY [mode] {
    y1 = int(nY [mode]) - 1
  }
  for z:= uint(y); z <= uint(y1); z++ {
    writeHorizontal (uint(x), z, uint(x1))
  }
}


func rectangleFullInv (x, y, x1, y1 int) {
//
  invertGr (uint(x), uint(y), uint(x1), uint(y1))
}


func onRectangle (x, y, x1, y1, a, b int, t uint) bool {
//
  if ! (between (x, x1, a, int(t)) && between (y, y1, b, int(t))) {
    return false
  }
  return between (a, a, x, int(t)) || between (a, a, x1, int(t)) ||
         between (b, b, y, int(t)) || between (b, b, y1, int(t))
}


func inRectangle (x, y, x1, y1, a, b int) bool {
//
  return between (x, x1, a, 0) && between (y, y1, b, 0)
}


func polygon (x, y []int) {
//
  if ! ok2 (x, y) { return }
  _segs (x, y, pt)
  n:= len (x)
  if n > 1 {
    _line (x[n-1], y[n-1], x[0], y[0], pt)
  }
}


func polygonInv (x, y []int) {
//
  if ! ok2 (x, y) { return }
  _segs (x, y, ptinv)
  n:= len (x)
  if n > 1 {
    _line (x[n-1], y[n-1], x[0], y[0], ptinv)
    pointInv (x[0], y[0])
    pointInv (x[n-1], y[n-1])
  }
}

/* TODO
func polygonFull (x, y []int) {
//
  if ! ok2 (x, y) { return }
  var (
    A0, A []byte
    C POINTER to uint16
    i uint
    xx, yy uint
    xMin, yMin, xMax, yMax int
  )
  n:= len (x)
  if n < 2 { return }
  if underX {
    ALLOCATE (A0, (n + 1) * TSIZE (uint))
    A = A0
    for i:= 0; i <= n; i++ {
      C = A; C:= x [i]
      A += TSIZE (uint16)
      C = A; C:= y [i]
      A += TSIZE (uint16)
    }
//    xker.PolygonFuellen (A0, n + 1, true)
    DEALLOCATE (A0, (n + 1) * TSIZE (uint))
  }
  neu // ! underX
  for i:= 0; i < n; i++ {
    if x[i] < x[i+1] {
      bresenham (x[i], y[i], x[i+1], y[i+1], setzen)
    } else {
      bresenham (x[i+1], y[i+1], x[i], y[i], setzen)
    }
  }
  if x[0] < x[n] {
    bresenham (x[0], y[0], x[n], y[n], setzen)
  } else {
    bresenham (x[n], y[n], x[0], y[0], setzen)
  }
  xx = 0
  yy = 0
  xMin = nX [mode]
  yMin = nY [mode]
  xMax = 0
  yMax = 0
  for i:= 0; i <= int(n); i++ {
    xx += x[i]
    yy += y[i]
    if x[i] < xMin { xMin = x[i] }
    if y[i] < yMin { yMin = y[i] }
    if x[i] > xMax { xMax = x[i] }
    if y[i] > yMax { yMax = y[i] }
  }
// gebraucht wird ein Punkt im Inneren des Polygons,
// aber das folgende löst das Problem nicht allgemein:
  fuellenV (xx / (n + 1), yy / (n + 1))
  for yy:= uint(yMin); yy <= uint(yMax); yy++ {
    for xx:= uint(xMin); xx <= uint(xMax); xx+ü {
      if gesetzt (xx, yy) {
        point (xx, yy)
      }
    }
  }
}
*/

/* TODO
func polygon1Full (x, y []int, x0, y0 int, n uint) {
var (
  i, xx, yy: uint
  xMin, yMin, xMax, yMax: int
//
  if n < 2 { return }
  if underX {
    PolygonFuellen (x, y, n)
  }
  neu
  for i:= 0; n - 1 {
    if x[i] < x[i+1] {
      bresenham (x[i], y[i], x[i+1], y[i+1], setzen)
    } else {
      bresenham (x[i+1], y[i+1], x[i], y[i], setzen)
    }
  }
  if x[0] < x[n] {
    bresenham (x[0], y[0], x[n], y[n], setzen)
  } else {
    bresenham (x[n], y[n], x[0], y[0], setzen)
  }
  xMin = nX [mode]
  yMin = nY [mode]
  xMax = 0
  yMax = 0
  for i:= 0; n {
    if x[i] < xMin { xMin = x[i] }
    if y[i] < yMin { yMin = y[i] }
    if x[i] > xMax { xMax = x[i] }
    if y[i] > yMax { yMax = y[i] }
  }
  fuellenV (x0, y0)
  for yy = VAL (uint, yMin); VAL (uint, yMax) {
    for xx = VAL (uint, xMin); VAL (uint, xMax) {
      if gesetzt (xx, yy) {
        point (xx, yy)
      }
    }
  }
}
*/


// TODO func PolygonFullInv ()


func onPolygon (x, y []int, a, b int, t uint) bool {
//
  if ! ok2 (x, y) { return false }
  n:= len (x)
  if n == 0 { return x[0] == a && y[0] == b }
  for i:= 1; i < int(n); i++ {
    if onLine (x[i-1], y[i-1], x[i], y[i], a, b, t) {
      return true
    }
  }
  return onLine (x[n-1], y[n-1], x[0], y[0], a, b, t)
}


func circ (x, y int, r uint, filled bool, p pointtype) {
// Algorithmus von Bresenham (Fellner: Computer Grafik, 5.5)
//
  if ! visible { return }
  if x >= int(nX [mode]) ||
     y >= int(nY [mode]) ||
     r >= nX [mode] {
    return
  }
  var A application
  switch p { case pt:
    A = point
  case ptinv:
    A = pointInv
  case onpt:
    A = on
  }
  if r == 0 {
    A (x, y)
    return
  }
  x1, y1:= 0, int(r)
  Fehler:= 3
  Fehler -= 2 * int(r)
/*
  if filled {
    horizontal (x - r, y, x + r, b)
    point (x, y - r)
    point (x, y + r)
  } else {
    A (x - r, y    )
    A (x + r, y    )
    A (x    , y - r)
    A (x    , y + r)
  }
  x1++
  if Fehler >= 0 {
    y1--
    Fehler -= 4 * y1
  }
  Fehler += 6
*/
  y0:= y1 + 1
  for x1 <= y1 {
    if filled {
      horizontal (uint(x - y1), uint(y - x1), uint(x + y1), p)
      if x1 > 0 {
        horizontal (uint(x - y1), uint(y + x1), uint(x + y1), p)
      }
      if y1 < y0 { // not yet correct, but a bit better than the above code
        y0 = y1
        horizontal (uint(x - x1), uint(y - y1), uint(x + x1), p)
        horizontal (uint(x - x1), uint(y + y1), uint(x + x1), p)
      }
    } else {
      A (x - y1, y - x1)
      A (x + y1, y - x1)
      A (x - y1, y + x1)
      A (x + y1, y + x1)
      A (x - x1, y - y1)
      A (x + x1, y - y1)
      A (x - x1, y + y1)
      A (x + x1, y + y1)
    }
    x1++
    if Fehler >= 0 {
      y1--
      Fehler -= 4 * y1
    }
    Fehler += 4 * x1 + 2
  }
}


func circle (x, y int, r uint) {
//
  if underX {
    xker.Ellipse (x, y, int(r), int(r), true, false)
    return
  }
  if iok (x, y) {
    if uint(x) >= r && uint(y) >= r {
      circ (x, y, r, false, pt)
    }
  }
}


func circleInv (x, y int, r uint) {
//
  if underX {
    xker.Ellipse (x, y, int(r), int(r), false, false)
    return
  }
  if iok (x, y) {
    if uint(x) >= r && uint(y) >= r {
      circ (x, y, r, false, ptinv)
    }
  }
}


func circleFull (x, y int, r uint) {
//
  if underX {
    xker.Ellipse (x, y, int(r), int(r), true, true)
    return
  }
  if iok (x, y) {
    if uint(x) >= r && uint(y) >= r {
      circ (x, y, r, true, pt)
    }
  }
}


func circleFullInv (x, y int, r uint) {
//
  if underX {
    xker.Ellipse (x, y, int(r), int(r), false, true)
    return
  }
  if iok (x, y) {
    if uint(x) >= r && uint(y) >= r {
      circ (x, y, r, true, ptinv)
    }
  }
}


func onCircle (x, y int, r uint, a, b int, t uint) bool {
//
//  if ! between (x - int(r), x + int(r), a) { return false }
/*
  if r == 0 { return a == x && b == y }
  z = a * a + b * b
  if z > r * r { z = z - r * r } else { z = r * r - z }
*/
  xxx, yyy, ttt, incident = a, b, int(t * t), false
  circ (x, y, r, false, onpt)
  return incident
}


func ell (x, y int, a, b uint, filled bool, p pointtype) {
//
  if a == b {
    circ (x, y, a, filled, p)
    return
  }
  if ! visible ||
     x >= int(nX [mode]) || y >= int(nY [mode]) {
    return
  }
  var A application
  switch p { case pt:
    A = point
  case ptinv:
    A = pointInv
  case onpt:
    A = on
  }
  if a == 0 {
    if b == 0 {
      A (x, y)
    } else {
      vertical (uint(x), uint(y - int(b)), uint(y + int(b)), p)
    }
    return
  } else {
    if b == 0 {
      horizontal (uint(x - int(a)), uint(y), uint(x + int(a)), p)
      return
    }
  }
  a1, b1:= 2 * a * a, 2 * b * b
  i:= int (a * b * b)
  x2, y2:= int(2 * a * b * b), 0
  xi, x1:= x - int(a), x + int(a)
  yi, y1:= y, y
  var xl int
  if xi < 0 {
    xl = 0
  } else {
    xl = xi
  }
  if filled {
    horizontal (uint(xl), uint(y), uint(x1), p)
  } else {
    A (xl, y)
    A (int(x1), y)
  }
  var yo int
  if a == 0 {
    if y < int(b) {
      yo = 0
    } else {
      yo = y - int(b)
    }
    vertical (uint(xi), uint(yo), uint(y) + b, p)
    return
  }
  for { // a > uint(0) {
    if i > 0 {
      yi--
      y1++
      y2 += int(a1)
      i -= int(y2)
    }
    if i <= 0 {
      xi++
      x1--
      x2 -= int(b1)
      i += int(x2)
      a--
    }
    if xi < 0 {
      xl = 0
    } else {
      xl = xi
    }
    if yi < 0 {
      yo = 0
    } else {
      yo = yi
    }
    var xr int
    if x1 < int(nX [mode]) {
      xr = int(x1)
    } else {
      xr = int(nX [mode]) - 1
    }
    var yu int
    if y1 < int(nY [mode]) {
      yu = int(y1)
    } else {
      yu = int(nY [mode]) - 1
    }
    if filled {
      horizontal (uint(xl), uint(yo), uint(xr), p)
      horizontal (uint(xl), uint(yu), uint(xr), p)
    } else {
      A (xl, yo)
      A (xr, yo)
      A (xl, yu)
      A (xr, yu)
    }
    if a == uint(0) {
      break
    }
  }
}


func ellipse (x, y int, a, b uint) {
//
  if underX {
    xker.Ellipse (x, y, int(a), int(b), true, false)
    return
  }
  if iok (x, y) {
    if uint(x) >= a && uint(y) >= b {
      ell (x, y, a, b, false, pt)
    }
  }
}


func ellipseInv (x, y int, a, b uint) {
//
  if underX {
    xker.Ellipse (x, y, int(a), int(b), false, false)
    return
  }
  if iok (x, y) {
    if uint(x) >= a && uint(y) >= b {
      ell (x, y, a, b, false, ptinv)
    }
  }
}


func ellipseFull (x, y int, a, b uint) {
//
  if underX {
    xker.Ellipse (x, y, int(a), int(b), true, true)
    return
  }
  if iok (x, y) {
    if uint(x) >= a && uint(y) >= b {
      ell (x, y, a, b, true, pt)
    }
  }
}


func ellipseFullInv (x, y int, a, b uint) {
//
  if underX {
    xker.Ellipse (x, y, int(a), int(b), false, true)
    return
  }
  if iok (x, y) {
    if uint(x) >= a && uint(y) >= b {
      ell (x, y, a, b, true, ptinv)
    }
  }
}


func onEllipse (x, y int, a, b uint, A, B int, t uint) bool {
//
  xxx, yyy, ttt, incident = A, B, int(t * t), false
  ell (x, y, a, b, false, onpt)
  return incident
}


func bezier (x, y []int, m, n, i uint) (int, int) {
//
  a, b:= float64 (i) / float64 (n), float64 (n - i) / float64 (n)
  t, t1:= make ([]float64, m), make ([]float64, m)
  t[0], t1[0] = 1.0, 1.0
  for k:= uint(1); k < m; k++ {
    t[k], t1[k] = a * t[k - 1], b * t1[k - 1]
  }
  for k:= uint(0); k < m; k++ {
    w:= float64(ker.Binomial (m - 1, k)) * t[k] * t1[m - 1 - k]
    a += w * float64 (x[k])
    b += w * float64 (y[k])
  }
  return int(a + 0.5), int(b + 0.5)
}


func cv (x, y []int, p pointtype) {
//
  m:= uint(len (x))
  if m == 0 || m != uint(len (y)) { return }
  var A application
  switch p { case pt:
    A = point
  case ptinv:
    A = pointInv
  case onpt:
    A = on
  }
  A (x[0], y[0])
  n:= uint(0)
  var dx, dy int
  for i:= uint(1); i < m; i++ {
    if x[i] < x[i-1] { dx = x[i-1] - x[i] } else { dx = x[i] - x[i-1] }
    if y[i] < y[i-1] { dy = y[i-1] - y[i] } else { dy = y[i] - y[i-1] }
    n += uint(math.Sqrt (float64 (dx * dx + dy * dy) + 0.5))
  }
  if n == 0 { return }
  X, Y:= make ([]int, n), make ([]int, n)
  for i:= uint(0); i < n; i++ {
    X[i], Y[i] = bezier (x, y, m, n, i)
  }
  if underX && p != onpt {
    xker.Points (X, Y, p == pt)
    return
  }
  for i:= uint(0); i < n; i++ {
    A (X[i], Y[i])
  }
}


func curve (x, y []int) {
//
  if ! ok2 (x, y) { return }
  cv (x, y, pt)
}


func curveInv (x, y []int) {
//
  if ! ok2 (x, y) { return }
  cv (x, y, ptinv)
}


func onCurve (X, Y []int, a, b int, t uint) bool {
//
  if ! ok2 (X, Y) { return false }
  xxx, yyy, ttt, incident = a, b, int(t * t), false
  cv (X, Y, onpt)
  return incident
}
