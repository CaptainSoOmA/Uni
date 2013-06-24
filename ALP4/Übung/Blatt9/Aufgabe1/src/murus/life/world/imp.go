package world

// (c) Christian Maurer   v. 130309 - license see murus.go

import (
  "os"
  "murus/ker"; . "murus/obj"; "murus/str"
  "murus/kbd"
  "murus/col"; "murus/scr"
  "murus/nat"; "murus/errh"
  "murus/pseq"; "murus/files"
  "murus/life/species"
)
const (
  y0 = 2
  xMax = 64 // s. WVGApp
  yMax = 36
  lenName = 8
  worldHint = "Namen der Welt eingeben (nur Buchstaben/Ziffern)  Programmende: leere Eingabe"
)
type
  Imp struct {
        name string
        spec [yMax][xMax]*species.Imp
        line,
      column uint
             }
var (
  ny,     // <= yMax - 2 // first two lines remain free for world and informations
  nx uint // <= xMax
  suffix string
  shadow *Imp
  help [1]string
)


func sys (s species.System) {
//
  ny = scr.NY () / species.Height - 2
  nx = scr.NX () / species.Width
  species.Sys (s)
  suffix = species.Suffix
  shadow = New ()
}


func New () *Imp {
//
  w:= new (Imp)
  w.name = str.Clr (lenName)
  w.name = "welt"
  for y:= uint(0); y < ny; y++ {
    for x:= uint(0); x < nx; x++ {
      w.spec [y][x] = species.New ()
    }
  }
  w.line, w.column = ny / 2, nx / 2
  return w
}


func (w *Imp) Name (name string) {
//
  w.name = name
  file:= pseq.New (w)
  file.Name (w.name + "." + suffix)
  if file.Empty () {
    w.Clr ()
  } else {
    w.Copy (file.Get ().(*Imp))
  }
  w.Write (0, 0)
  file.Terminate ()
}


func (w *Imp) terminate () {
//
  file:= pseq.New (w)
  file.Name (w.name + "." + suffix)
  file.Clr ()
  file.Put (w)
  file.Terminate ()
}


func (w *Imp) Rename (name string) {
//
  file:= pseq.New (w)
  file.Name (w.name + "." + suffix)
  w.name = name
  file.Rename (w.name + "." + suffix)
  file.Terminate ()
}


func (w *Imp) Empty () bool {
//
  for Z:= uint(0); Z < ny; Z++ {
    for S:= uint(0); S < nx; S++ {
      if ! w.spec [Z][S].Empty() {
        return false
      }
    }
  }
  return true
}


func (w *Imp) Clr () {
//
  for y:= uint(0); y < ny; y++ {
    for x:= uint(0); x < nx; x++ {
      w.spec [y][x].Clr ()
    }
  }
}


func (w *Imp) Eq (X Object) bool {
//
  v, ok:= X.(*Imp)
  if ! ok { return false }
  for y:= uint(0); y < ny; y++ {
    for x:= uint(0); x < nx; x++ {
      if ! w.spec [y][x].Eq (v.spec [y][x]) {
        return false
      }
    }
  }
  return true
}


func (x *Imp) Less (Y Object) bool {
//
  return false
}


func (w *Imp) Copy (X Object) {
//
  v, ok:= X.(*Imp)
  if ! ok { return }
  for y:= uint(0); y < ny; y++ {
    for x:= uint(0); x < nx; x++ {
      w.spec [y][x].Copy (v.spec [y][x])
    }
  }
  w.line, w.column = v.line, v.column
}


func (w *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (w)
  return y
}


func (w *Imp) SetColours (f, b col.Colour) { // fake
//
}


func (w *Imp) number (s *species.Imp) uint {
//
  n:= uint(0)
  for y:= uint(0); y < ny; y++ {
    for x:= uint(0); x < nx; x++ {
      if w.spec [y][x].Eq (s) {
        n++
      }
    }
  }
  return n
}


func (w *Imp) writeNumbers () {
//
  spec:= species.New ()
  w.spec[0][0].SetFormat (species.Long)
  for a:= uint(0); a < spec.Number (); a++ {
    c:= 27 + 20 * a
    spec.Write (y0 - 1, c)
    nat.Write (w.number (spec), y0 - 1, c + 1 + lenName + 2)
    spec.Inc ()
  }
  w.spec[0][0].SetFormat (species.Short)
}


func (w *Imp) Write (l, c uint) { // l, c: fake
//
  for y:= uint(0); y < ny; y++ {
    for x:= uint(0); x < nx; x++ {
      w.spec [y][x].Write (y0 + y, 2 * x)
    }
  }
  w.writeNumbers ()
}


var
  y_, x_ uint


func nNeighbours (spec *species.Imp) uint {
//
  n:= uint(0)
  if x_ < nx - 1 {
    if shadow.spec [y_][x_ + 1].Eq (spec) { n++ }
  }
  if y_ > 0 {
    if shadow.spec [y_ - 1][x_].Eq (spec) { n++ }
  }
  if x_ > 0 {
    if shadow.spec [y_][x_ - 1].Eq (spec) { n++ }
  }
  if y_ < ny - 1 {
    if shadow.spec [y_ + 1][x_].Eq (spec) { n++ }
  }
  if species.NNeighbours == 4 { return n }
  if y_ > 0 && x_ + 1 < nx {
    if shadow.spec [y_ - 1][x_ + 1].Eq (spec) { n++ }
  }
  if y_ > 0 && x_ > 0 {
    if shadow.spec [y_ - 1][x_ - 1].Eq (spec) { n++ }
  }
  if y_ + 1 < ny && x_ > 0 {
    if shadow.spec [y_ + 1][x_ - 1].Eq (spec) { n++ }
  }
  if y_ + 1 < ny && x_ + 1 < nx {
    if shadow.spec [y_ + 1][x_ + 1].Eq (spec) { n++ }
  }
  return n
}


func (w *Imp) modify () {
//
  shadow.Copy (w)
  w.spec[0][0].SetFormat (species.Short)
  for y:= uint(0); y < ny; y++ {
    for x:= uint(0); x < nx; x++ {
      y_, x_ = y, x
      w.spec [y][x].Modify (nNeighbours)
      w.spec [y][x].Write (y0 + y, 2 * x)
    }
  }
}


func (w *Imp) Edit (l, c uint) { // l, c: fake
//
  s:= w.spec [w.line][w.column]
  sa:= s
  y, x:= scr.MousePos ()
  ya, xa:= y, x
  loop:
  for {
    s.Write (y0 + w.line, 2 * w.column)
    ya, xa = y, x
    sa = s
    c, _:= kbd.Command ()
    y, x = scr.MousePos ()
    if y0 <= y && y < ny + y0 && x < 2 * nx {
      w.line, w.column = y - y0, x / 2
      s = w.spec [w.line][w.column]
    }
    if ya >= y0 {
      sa.Mark (false)
      sa.Write (ya, 2 * (xa / 2))
    }
    switch c { case kbd.Esc:
      break loop
    case kbd.Enter:
      w.modify ()
      w.Write (0, 0)
    case kbd.Help:
      errh.WriteHelp (help[:])
    case kbd.Here:
      s.Inc ()
    case kbd.There:
      s.Dec ()
    }
    s.Mark (true)
    w.writeNumbers ()
  }
  w.terminate ()
}


func (w *Imp) Codelen () uint {
//
  return ny * nx * w.spec [0][0].Codelen () + 2 * 4
}


func (w *Imp) Encode () []byte {
//
  b:= make ([]byte, w.Codelen())
  i:= uint(0)
  di:= w.spec[0][0].Codelen()
  for y:= uint(0); y < ny; y++ {
    for x:= uint(0); x < nx; x++ {
      copy (b[i:i+di], w.spec [y][x].Encode ())
      i += di
    }
  }
  copy (b[i:i+4], Encode (w.line))
  i += 4
  copy (b[i:i+4], Encode (w.column))
  return b
}


func (w *Imp) Decode (b []byte) {
//
  i:= uint(0)
  di:= w.spec[0][0].Codelen()
  for y:= uint(0); y < ny; y++ {
    for x:= uint(0); x < nx; x++ {
      w.spec [y][x].Decode (b[i:i+di])
      i += di
    }
  }
  w.line = Decode (w.line, b[i:i+4]).(uint)
  i += 4
  w.column = Decode (w.column, b[i:i+4]).(uint)
}


func init () {
//
//  var x World = New(); if x == nil {} // ok
  if ! scr.Switchable (scr.WVGApp) {
    errh.Error ("Bildschirmauflösung kleiner als 1024 x 576", 0)
    ker.Terminate (); os.Exit (0)
  } else {
    scr.Switch (scr.WVGApp)
  }
  help[0] = "Hilfe ist noch nicht implementiert" // TODO
  files.Cd0()
}
