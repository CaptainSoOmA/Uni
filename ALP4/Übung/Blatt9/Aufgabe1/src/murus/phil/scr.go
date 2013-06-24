package phil

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  "sync"; "math"
  "murus/env"; "murus/col"; "murus/scr"; "murus/img"
)
var (
  colour [max]col.Colour
  notUsedColour col.Colour
  image [max+1]string
  thinkColour col.Colour
  mutex sync.Mutex
  x0, y0 uint
  x, y, x1, y1, xt, yt [max]int
)


func writePlate (p uint, i bool) {
//
  var r uint
  if nPhilos <= 7 {
    r = y0 / 8
  } else {
    r = y0 / 3
  }
  scr.Circle (xt[p], yt[p], r);
  if i {
    scr.Circle (xt[p], yt[p], (2 * r) / 1)
  }
}


func writeFork (p uint) {
//
  scr.Line (x[p], y[p], x1[p], y1[p])
}


func write (p uint) {
//
  if stat[p] == satisfied {
    scr.Colour (thinkColour)
    writePlate (p, false)
    scr.WriteGr (text[stat[p]], xt[p] - 12 * 4, yt[p] - 8)
    img.Get (image[p], uint(xt[p]) - 50, uint(yt[p]) - 50)
  } else {
    img.Get (image[max], uint(xt[p] - 50), uint(yt[p]) - 50)
    scr.Colour (colour[p])
    writePlate (p, false)
    scr.WriteGr (text[stat[p]], xt[p] - 12 * 4, yt[p] - 8)
  }
  switch stat[p] { case satisfied:
    writePlate (p, true)
    scr.Colour (notUsedColour)
    writeFork (left (p))
    writeFork (p)
  case hungry:
    ;
  case starving:
    ;
  case hasRightFork:
    writeFork (p)
  case hasLeftFork:
    writeFork (left (p))
  case dining:
    writePlate (p, true)
    writeFork (p)
    writeFork (left (p))
  }
}


func init () {
//
  colour = [max]col.Colour { col.LightCyan, col.LightYellow, col.Pink, col.LightGreen,
                             col.LightWhite, col.Orange, col.LightBlue, col.Yellow,
                             col.LightGreen, col.LightRed, col.LightMagenta, col.LightOrange }
  notUsedColour = col.DarkGray
  image = [max+1]string { "Anaxagoras", "Aristoteles", "Cicero", "Demokrit",
                          "Diogenes", "Epikur", "Heraklit", "Platon",
                          "Protagoras", "Pythagoras", "Sokrates", "Thales",
                          "Niemand" }
  for n:= 0; n <= max; n++ {
    image[n] = env.Val ("GOSRC") + "/murus/phil/pics/" + image[n]
  }
  const zweipi = float64(2 * math.Pi)
  text = [nStatuses]string { "    satt    ", "  hungrig   ", "sehr hungrig",
                             "sehr hungrig", "sehr hungrig", "  speisend  " }
  thinkColour = col.ScreenB
/*
  if p > 6 {
    if scr.UnderX () {
      scr.Fullscreen()
    }
  }
*/
  x0, y0 = scr.NX() / 2, scr.NY() / 2
  const f = float64(0.5)
  for p:= uint(0); p < nPhilos; p++ {
// Attention: the mathematical positiv sense is inverted on the scr, because lines count upwards !
// middlepoint of plates:
    w, r:= float64 (p) / float64(nPhilos), 0.75 * float64 (y0)
    xt[p] = int(math.Trunc (r * math.Cos (zweipi * w) + f)) + int(x0)
    yt[p] = int(math.Trunc (- r * math.Sin (zweipi * w) + f)) + int(y0)
// endpoints of the fork to the right (with the same number):
    w, r = w + 0.5 / float64 (nPhilos), float64 (y0)
    x[p] = int(math.Trunc (r * math.Cos (zweipi * w) + f)) + int(x0)
    y[p] = int(math.Trunc (- r * math.Sin (zweipi * w) + f)) + int(y0)
    r = r / 2.0
    x1[p] = int(math.Trunc (r * math.Cos (zweipi * w) + f)) + int(x0)
    y1[p] = int(math.Trunc (- r * math.Sin (zweipi * w) + f)) + int(y0)
//    stat[p] = hungry
  }
}
