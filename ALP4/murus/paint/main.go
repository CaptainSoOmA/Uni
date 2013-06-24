package main // paint

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  "murus/ker"; "murus/env"
//  . "murus/obj"
  "murus/str"; "murus/kbd"
  "murus/col"; "murus/scr"; "murus/box"
  "murus/img"; "murus/sel"
)
type
  figure byte; const (line = iota; rectangle; circle; ellipse; Nfigure)
type
  format byte; const (border = iota; inv; full)


func paint (f figure, fo format, x0, y0, x, y int) {
//
  var a, b uint
  if x >= x0 { a = uint(x - x0) } else { a = uint(x0 - x) }
  if y >= y0 { b = uint(y - y0) } else { b = uint(y0 - y) }
  switch f { case line:
    switch fo { case border:
      scr.Line (x0, y0, x, y)
    case inv:
      scr.LineInv (x0, y0, x, y)
    case full:
      // scr.StrichAusgeben (x0, y0, x, y, Staerke)
  }
  case rectangle:
    switch fo { case border:
      scr.Rectangle (x0, y0, x, y)
    case inv:
      scr.RectangleInv (x0, y0, x, y)
    case full:
      scr.RectangleFull (x0, y0, x, y)
    }
  case circle:
    if b > a { a = b }
    switch fo { case border:
      scr.Circle (x0, y0, a)
    case inv:
      scr.CircleInv (x0, y0, a)
    case full:
      scr.CircleFull (x0, y0, a)
    }
  case ellipse:
    switch fo { case border:
/*
      scr.Ellipse ((x0 + x) / 2, (y0 + y) / 2, a, b)
    case invers:
      scr.EllipseInv ((x0 + x) / 2, (y + y0) / 2, a, b)
    case full:
      scr.EllipseFull ((x0 + x) / 2, (y0 + y) / 2, a, b) 
*/
      scr.Ellipse (x0, y0, a, b)
    case inv:
      scr.EllipseInv (x0, y0, a, b)
    case full:
      scr.EllipseFull (x0, y0, a, b)
    }
  }
//  scr.Colour (paintColour), (* Wörkeraund: Bildschirm....Invertieren f *)
}


func main () {
//
  if ! scr.MouseEx () { return }
  var symbol [Nfigure]byte
  symbol[line] =  'S' // "Strecke"
  symbol[rectangle] = 'R' // "Rechteck"
  symbol[circle] =    'K' // "Kreis"
  symbol[ellipse] =  'E' // "Ellipse"
  X, Y:= 0, 0
  X1, Y1:= scr.NX(), scr.NY()
//  Farbe, Papier:= col.LightWhite, col.Black
  Farbe, Papier:= col.Black, col.LightWhite
  col.ScreenF, col.ScreenB = Farbe, Papier
  scr.Cls()
  paintColour:= Farbe
  scr.Colour (paintColour)
//  Staerke = 3
  bx:= box.New()
  bx.Wd (20)
  bx.Colours (Papier, Farbe)
  Name:= env.Par (1)
  if str.Empty (Name) { Name = "temp" }
  scr.Save (0, 0, 20, 1)
  for {
    bx.Edit (&Name, 0, 0)
    if ! str.Empty (Name) {
      str.RemSpaces (&Name)
      break
    }
  }
  scr.Restore (0, 0, 20, 1)
  img.Get (Name, uint(X), uint(Y))
  scr.MouseCursor (true)
  Figur:= figure(rectangle)
  var x, y, x0, y0 int
  loop: for {
    scr.Colour (paintColour)
    Zeichen, Kommando, T:= kbd.Read()
    switch Kommando { case kbd.None:
      x, y = scr.MousePosGr ()
      scr.SwitchTransparence (true)
      scr.Write1Gr (Zeichen, x, y - int(scr.NY1()))
//    scr.WarpMouse (x + scr.NX1(), y)
    case kbd.Esc:
      break loop
    case kbd.Back:
      switch T { case 0: x, y = scr.MousePosGr ()
        x -= int(scr.NX1())
        scr.Colour (Papier)
        scr.Write1Gr (' ', x, y - int(scr.NY1()))
//        scr.RectangleFull (x, y - scr.NY1(), x + scr.NX1(), y)
//        scr.WarpMouseGr (x, y)
        scr.Colour (paintColour)
      default:
        scr.Cls()
      }
/*
    case kbd.Ins:
      img.Write (X, Y, X1, Y1 - 16, Name)
      box.Edit (Feld, Name, scr.Zeilenzahl () - 1, 0)
      img.Get (X, Y, Name)
*/
    case kbd.Help:
      paintColour = sel.Colour ()
//    case kbd.LookFor:
//      Staerke = Strichstaerken.Staerke()
    case kbd.Enter:
      if T > 0 {
        x0, y0 = scr.MousePosGr ()
//        scr.Fill1 (x0, y0)
      }
    case kbd.PrintScr:
      img.Print (uint(X), uint(Y), X1, Y1 - 16)
    case kbd.Tab:
      if T == 0 {
        if Figur + 1 < Nfigure { Figur ++ } else { Figur = figure(0) }
      } else {
        if Figur > 0 { Figur -- } else { Figur = figure(Nfigure - 1) }
      }
      scr.Colours (col.White, Papier)
      scr.Write1 (symbol [Figur], scr.NY() - 1, 0)
    case kbd.Here:
      x0, y0 = scr.MousePosGr ()
      scr.CircleFull (x0, y0, 3 / 2)
    case kbd.Pull:
      x, y = scr.MousePosGr ()
      scr.Line (x0, y0, x, y)
      x0, y0 = x, y
    case kbd.Hither:
      x, y = scr.MousePosGr ()
      scr.Line (x0, y0, x, y)
    case kbd.There:
      x0, y0 = scr.MousePosGr ()
      x, y = x0, y0
    case kbd.Push:
      paint (Figur, inv, x0, y0, x, y)
      x, y = scr.MousePosGr ()
      paint (Figur, inv, x0, y0, x, y)
    case kbd.Thither:
      paint (Figur, inv, x0, y0, x, y)
//      scr.Colour (paintColour)
      x, y = scr.MousePosGr ()
      paint (Figur, border, x0, y0, x, y)
      x0, y0 = x, y
    case kbd.This:
      x0, y0 = scr.MousePosGr ()
      x, y = x0, y0
    case kbd.Move:
      scr.LineInv (x0, y0, x, y)
      x, y = scr.MousePosGr ()
      scr.LineInv (x0, y0, x, y)
    case kbd.Thus:
      scr.LineInv (x0, y0, x, y)
      x, y = scr.MousePosGr ()
      scr.Line (x0, y0, x, y)
      x0, y0 = x, y
    }
  }
  scr.Save (0, 0, 20, 1)
  for {
    bx.Edit (&Name, 0, 0)
// TODO make sure, that "Name" is x{x|y} where x is letter, y is digit
    if ! str.Empty (Name) {
      str.RemSpaces (&Name)
      break
    }
  }
  scr.Restore (0, 0, 20, 1)
  img.Put (Name, uint(X), uint(Y), X1, Y1)
  ker.Terminate()
}
