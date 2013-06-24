package main

// (c) Christian Maurer   v. 130302 - license see murus.go

import (
  "murus/ker"; . "murus/obj"; "murus/font"
  "murus/str"; "murus/spc"; "murus/kbd"
  "murus/col"; "murus/scr"; "murus/errh"; "murus/sel" // ; "murus/img"
  "murus/seq"; "murus/stk"; "murus/pseq" // ; "murus/files"
  "murus/fig2"; "murus/psp"
)
var (
  figure *fig2.Imp
  figures, copiedFigures, cuttedFigures, depot, image *seq.Imp
  stack *stk.Imp
  actColour, paperColour col.Colour
)


func bgColour (c col.Colour) {
//
  col.ScreenB = c
  scr.Cls()
  write()
}


func names () (string, string) {
//
  scr.SwitchFontsize (font.Normal)
  return sel.Names ("Bild:", "epn", 32, 0, 0, paperColour, actColour)
}


func load (f *seq.Imp, N string) {
//
  n:= pseq.Length (N)
  if n == 0 { return }
  buf:= make ([]byte, n)
  file:= pseq.New (buf)
  file.Name (N)
  buf = file.Get().([]byte)
  file.Terminate()
  f.Decode (buf)
}


func store (f *seq.Imp, N string) {
//
  pseq.Erase (N)
  n:= f.Codelen()
  buf:= make ([]byte, n)
  buf= f.Encode()
  file:= pseq.New (buf)
  file.Name (N)
  file.Put (buf)
  file.Terminate()
}


func write() {
//
  scr.Buf (true)
  figures.Trav (func (a Any) { a.(*fig2.Imp).Write() })
  scr.Buf (false)
}


func print (n string) {
//
  f:= psp.New ()
  f.Name (n)
  figures.Trav (func (a Any) { a.(*fig2.Imp).Print(f) })
  f.Terminate ()
}


// Liefert genau dann true, wenn eine figure inzident mit (x, y) ist.
// In diesem Fall ist f die erste dieser figureen und die figures ist auf f positioniert,
// andernfalls ist f unverändert.
func underMouse (x, y uint) bool {
//
  if figures.ExPred (func (a Any) bool { return a.(*fig2.Imp).On (int(x), int(y), 7) }, true) {
    figure = figures.Get().(*fig2.Imp)
    return true
  }
  return false
}


func im (a Any) { // for multiple use
//
  if a.(*fig2.Imp).Marked() { a.(*fig2.Imp).Invert() }
}


func invMarked() {
//
  figures.Trav (im)
}


func showMarked() {
//
  figures.Trav (im)
  ker.Msleep (250)
  figures.Trav (im)
}


func copyMarked() {
//
  if ! figures.ExPred (func (a Any) bool { return a.(*fig2.Imp).Marked() }, true) { return }
  figures.Filter (copiedFigures, func (a Any) bool { return a.(*fig2.Imp).Marked() })
//  figures.Trav (func (a Any) { a.(*fig2.Imp).Demark() })
  copiedFigures.Trav (func (a Any) { a.(*fig2.Imp).Mark (false) })
  x, y:= figures.Get().(*fig2.Imp).Pos()
  xm, ym:= scr.MousePosGr()
  copiedFigures.Trav (func (a Any) { a.(*fig2.Imp).Move (xm - x, ym - y) })
  copiedFigures.Trav (func (a Any) { a.(*fig2.Imp).Write() })
  figures.Join (copiedFigures)
}


func delMarked() {
//
  figures.Cut (depot, func (a Any) bool { return a.(*fig2.Imp).Marked() })
}


func getMarked() {
//
  depot.Trav (func (a Any) { a.(*fig2.Imp).Write() })
  figures.Join (depot)
}

/*
func inject (c col.Colour) {
//
  if figures.Empty() { return }
  figure.Clr()
  figure.SetColour (c)
  figures.Seek (0)
  figures.Ins (figure)
}


func extract (c col.Colour) {
//
  if figures.Empty(){ return }
  figures.Seek (0)
  H:= figures.Get().(*fig2.Imp).Colour()
  figures.Del()
  if ! col.Eq (c, H) {
    c = H
    col.ScreenB = c
    scr.Cls()
  }
  for figures.ExPred ( func (a Any) bool { return a.(*fig2.Imp).Empty()}, true) {
    figures.Del()
  }
}
*/

func cutMarked() {
//
  cuttedFigures.Cut (figures, func (a Any) bool { return a.(*fig2.Imp).Marked() })
  write()
  cuttedFigures.Trav (im)
}


func moveMarked (dx, dy int) {
//
  cuttedFigures.Trav (im)
  cuttedFigures.Trav (func (a Any) { a.(*fig2.Imp).Move (dx, dy) })
  cuttedFigures.Trav (im)
}


func joinMarked() {
//
  cuttedFigures.Trav (im)
  cuttedFigures.Trav (func (a Any) { a.(*fig2.Imp).Write() })
  figures.Join (cuttedFigures)
}


func writeMarked() {
//
  if figures.ExPred (func (a Any) bool { return a.(*fig2.Imp).Marked() }, true) {
    figures.Filter (image, func (a Any) bool { return a.(*fig2.Imp).Marked() })
    _, filename:= names ()
    store (image, filename)
    // copiedFigures.Clr()
  }
}


func readMarked() {
//
  figures.Trav (func (a Any) { a.(*fig2.Imp).Mark (false) })
  _, filename:= names ()
  load (image, filename)
  image.Trav (func (a Any) { a.(*fig2.Imp).Mark (true) })
  image.Trav (func (a Any) { a.(*fig2.Imp).Write() })
  figures.Join (image)
}


func control() {
//
  scr.SwitchFontsize (font.Normal)
  scr.Write ("          ", 1, 0)
  scr.WriteNat (figures.Num(), 1, 0)
  scr.WriteNat (figures.NumPred (func (a Any) bool { return a.(*fig2.Imp).Marked() }), 1, 5)
}


func kick (r spc.Direction, s int) {
//
  figures.Trav (func (a Any) { switch r {
                    case spc.Right: a.(*fig2.Imp).Move (s, 0)
                    case spc.Front: // free for future development
                    case spc.Top: a.(*fig2.Imp).Move (0, s)
                  } })
  write()
}


func push() {
//
  stack.Push (figure)
}


func top() *fig2.Imp {
//
  if stack.Empty() { return nil }
  f:= stack.Top().(*fig2.Imp)
  stack.Pop()
  f.Write()
  return f
}


func ins (f *fig2.Imp) {
//
  figures.Seek (figures.Num())
  figures.Ins (f)
  write()
}


func write1 (f *fig2.Imp) {
//
  figures.Put (f)
  write()
}


func deleted (x, y int) bool {
//
  if underMouse (uint(x), uint(y)) {
    figures.Del()
    write()
    return true
  }
  return false
}


func delLast() {
//
  if figures.Empty() { return }
  figures.Seek (figures.Num() - 1)
//  figure wieder auf den stack ?
  figures.Del()
  write()
}


func generate (f *fig2.Imp) {
//
  f.Clr()
  f.Edit() // interaktiv in fig2
  if ! f.Empty() {
    ins (f)
  }
//  inject()
//  store (figures, filename)
}


func change (x, y int) {
//
  if ! underMouse (uint(x), uint(y)) { return }
  figure.Edit() // interaktiv in fig2
  write1 (figure)
}


func setColours (x, y int, c col.Colour) {
//
  if ! underMouse (uint(x), uint(y)) { return }
  figure.SetColour (c)
  figure.Write()
  write1 (figure)
}


func mark (x, y int, m bool) {
//
  if ! underMouse (uint(x), uint(y)) { return }
  if ! figure.Marked() {
    figure.Mark (m)
    write1 (figure)
  }
}


func markAll (m bool) {
//
  figures.Trav (func (a Any) { a.(*fig2.Imp).Mark (m) })
  write()
}


func cut (x, y int) {
//
  if deleted (x, y) {
// scr.WriteNat (figures.Num(), 0, 0)
//    f.Write()
// errh.Hint ("cut => figure.Invert")
    figure.Invert()
  }
}


func move (dx, dy int) {
//
  figure.Invert()
  figure.Move (dx, dy)
  figure.Invert()
}


func join() {
//
  figure.Invert()
  ins (figure)
  figure.Write()
}


func main() {
//
  h:= [...]string {
// 0         1         2         3         4         5         6         7
// 012345678901234567890123456789012345678901234567890123456789012345678901234567
  "   Art (Farbe) für neue Figur auswählen: (Umschalt- +) Leer- oder Rolltaste   ",
//  "                 Schriftgröße auswählen: Alt- + Leer- oder Rolltaste          ",
  "                    neue Figur erzeugen: linke Maustaste,                     ",
  "                                         Streckenzüge, Polygone und Kurven    ",
  "                                         mit rechter Maustaste abschließen,   ",
  "                                         Texteingabe mit Eingabetaste beenden ",
  "                           Figur ändern: Alt- + linke Maustaste, die 'Punkte' ",
  "                                         mit rechter Maustaste verschieben,   ",
  "                                         mit linker Maustaste abschließen     ",
  "                           Figur färben: Umschalt-, Alt- + linke Maustaste    ",
  "             (alle) Figur(en) markieren: (Umschalt- +) F5-Taste oder          ",
  "                                         (Alt- +) mittlere Maustaste          ",
  "          (alle) Figur(en) entmarkieren: (Umschalt- +) F6-Taste oder          ",
  "                                         Umschalt- + (Alt- +) mittl. Maustaste",
  "           markierte Figur(en) kopieren: Umschalt- + linke Maustaste          ",
  "      (markierte) Figur(en) verschieben: (Umschalt- +) rechte Maustaste       ",
  "          (markierte) Figur(en) löschen: (Umschalt- +) Entfernungtaste oder   ",
  "                                         (Umschalt-,) Alt- + rechte Maustaste ",
  "     letzte gelöschte Figur zurückholen: Rücktaste (<-)                       ",
  "          letzte erzeugte Figur löschen: Umschalt- + Rücktaste (<-)           ",
  "letzte gelöschte mark. Fig. zurückholen: Alt- + Rücktaste (<-)                ",
  "Hintergrundfarbe umschalten (auswählen): (Umschalt- +) F4-Taste               ",
  " Figuren aus eBoard holen und markieren: F7-Taste                             ",
  "    markierte Figuren in eBoard ablegen: F8-Taste                             ",
  "                      eBoard ausdrucken: Drucktaste                           ",
  "                           ePen beenden: Abbruchtaste (Esc)                   ",
  "                                                                              ",
  "                  Steuerungstaste (Strg) wirkt wie Umschalttaste              "}
  help:= make ([]string, len (h))
  for i:= 0; i < len (h); i++ { str.Set (&help [i], h [i]) }
  scr.Switch (scr.SXGA)
  actColour = col.ScreenF
  paperColour = col.ScreenB
//  fo = scr.Normal
  figure = fig2.New()
  newFigure:= fig2.New()
  newFigure.SetColour (actColour)
  figures, copiedFigures, image = seq.New (newFigure), seq.New (newFigure), seq.New (newFigure)
  cuttedFigures, depot = seq.New (newFigure), seq.New (newFigure)
  stack = stk.New (newFigure)
  name, filename:= names ()
  if filename == "" { return }
  origname:= filename
  load (figures, filename)
//  extract (paperColour)
  write()
  var x0, y0 int // kbd.Push
  var movable bool
  var Schub int
  loop: for {
    scr.MouseCursor (true)
//    control(); write()
    c, cmd, d:= kbd.Read()
    if cmd == kbd.None {
      switch c { case ' ':
        cmd = kbd.Roll
      }
    }
    xm, ym:= scr.MousePosGr()
    switch d { case 0:
      Schub = 8
    case 1:
      Schub = int(scr.NY() / 40)
    default:
      Schub = int(scr.NX() / 8)
    }
    switch cmd { case kbd.Esc:
      break loop
    case kbd.Enter:
      // actualize ?
    case kbd.Back:
      switch d { case 0:
        ins (top())
      case 1:
        delLast()
      default:
        getMarked()
      }
    case kbd.Left:
      kick (spc.Right, Schub)
    case kbd.Right:
      Schub = -Schub
      kick (spc.Right, Schub)
    case kbd.Up:
      kick (spc.Top, Schub)
    case kbd.Down:
      Schub = -Schub
      kick (spc.Top, Schub)
    case kbd.Del:
      switch d { case 0:
        if deleted (xm, ym) { push() }
      default:
        delMarked()
      }
    case kbd.Tab:
      // free for use
    case kbd.Help:
      scr.SwitchFontsize (font.Normal)
      errh.WriteHelp (help)
    case kbd.LookFor:
      // wird für die Auswahl des eBoards verwendet - NEIN, sondern:
      load (figures, filename)
      write()
    case kbd.Act:
//      inject (paperColour)
      store (figures, filename)
    case kbd.Cfg:
      paperColour = sel.Colour()
      bgColour (paperColour)
    case kbd.Mark:
      switch d { case 0:
        mark (xm, ym, true)
      case 1:
        markAll (true)
      default:
        showMarked()
      }
    case kbd.Demark:
      switch d { case 0:
        mark (xm, ym, false)
      case 1:
        markAll (false)
      default:
        showMarked()
      }
    case kbd.Deposit:
      writeMarked()
    case kbd.Paste:
      readMarked()
    case kbd.PrintScr:
      print (name)
    case kbd.Roll:
      switch d { case 0:
        newFigure.Select()
      case 1:
        actColour = sel.Colour()
        newFigure.SetColour (actColour)
      case 2:
        scr.SwitchFontsize (font.Normal)
//        fo = sel.Size (actColour)
//        newFigure.SetFont (fo)
      }
    case kbd.Here:
      switch d { case 0:
        generate (newFigure)
      case 1:
        copyMarked()
      case 2:
        change (xm, ym)
      case 3:
        setColours (xm, ym, actColour)
      }
    case kbd.There:
      movable = underMouse (uint(xm), uint(ym)) && d <= 1
      x0, y0 = xm, ym // kbd.Push
      switch d { case 0:
        if movable { cut (xm, ym) }
      case 1:
        if movable { cutMarked() }
      case 2:
        if deleted (xm, ym) { push() }
      default:
        delMarked()
        write()
      }
    case kbd.Push:
      if movable {
        switch d { case 0:
          move (xm - x0, ym - y0)
        case 1:
          moveMarked (xm - x0, ym - y0)
        }
        x0, y0 = scr.MousePosGr()
      }
    case kbd.Thither:
      if movable {
        switch d { case 0:
          join()
        case 1:
          joinMarked()
        }
      }
    case kbd.This:
      switch d { case 0:
        mark (xm, ym, true)
      case 1:
        mark (xm, ym, false)
      case 2:
        markAll (true)
      case 3:
        markAll (false)
      }
      invMarked()
    case kbd.Thus:
      invMarked()
    }
  }
  markAll (false)
  _, filename = names ()
  if filename == "" {
    filename = origname
  }
//  inject (paperColour)
  store (figures, filename) // -> Terminieren
  ker.Terminate()
}
