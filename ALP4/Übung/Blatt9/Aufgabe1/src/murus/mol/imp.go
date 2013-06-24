package mol

// (c) Christian Maurer   v. 130304 - license see murus.go

import (
  . "murus/obj" // ; "murus/ker"
  "murus/kbd"; "murus/col"; "murus/font"
  "murus/atom"; "murus/masks"
)
const
  pack = "mol"
type
  Imp struct {
         num uint
        comp []*atom.Imp
        l, c []uint
        mask *masks.Imp
             }


func (x *Imp) imp (Y Object) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  if x.num != y.num { TypeNotEqPanic (x, Y) }
//  if x.num == 0 { ker.Panic ("mol error: x.num == 0") }
  for i:= uint(0); i < y.num; i++ {
    if x.comp[i].Type () != y.comp[i].Type () { TypeNotEqPanic (x.comp[i], y.comp[i]) }
  }
  return y
}


func New () *Imp {
//
  return new (Imp)
}


func (x *Imp) Num () uint {
//
  return x.num
}


func (x *Imp) Component (n uint) Any {
//
  if n >= x.num { WrongUintParameterPanic ("Component", x, n) }
  return x.comp[n]
}


func (x *Imp) Ins (a *atom.Imp, l, c uint) {
//
  x.comp = append (x.comp, a.Clone ().(*atom.Imp))
  x.l, x.c = append (x.l, l), append (x.c, c)
  x.num ++
}


func (x *Imp) Del (n uint) {
//
  if n >= x.num { return }
  for i:= uint(n); i + 1 < x.num; i++ {
    x.comp[i] = x.comp[i + 1]
    x.l[i], x.c[i] = x.l[i + 1], x.c[i + 1]
  }
  x.num --
  x.comp[x.num] = nil
}


func (x *Imp) SetMask (m *masks.Imp) {
//
  x.mask = m
}


func (x *Imp) Empty () bool {
//
  for i:= uint(0); i < x.num; i++ {
    if ! x.comp[i].Empty() {
      return false
    }
  }
  return true
}


func (x *Imp) Clr () {
//
  for i:= uint(0); i < x.num; i++ {
    x.comp[i].Clr ()
  }
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= x.imp (Y)
  for i:= uint(0); i < x.num; i++ {
    if ! x.comp[i].Eq (y.comp[i]) {
      return false
    }
  }
  return true
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
//  x.num = y.num
//  x.comp = make ([]*atom.Imp, x.num)
//  x.l, x.c = make ([]uint, x.num), make ([]uint, x.num)
  for i:= uint(0); i < y.num; i++ {
    x.comp[i].Copy (y.comp[i])
    x.l[i], x.c[i] = y.l[i], y.c[i]
  }
  x.mask = y.mask // x.mask.Copy (y.mask) // wegen Typverlust in piset.New geht das nicht
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) less (y *Imp, n uint) bool {
//
  if n > x.num { return false }
  if x.comp[n].Less (y.comp[n]) { return true }
  if x.comp[n].Eq (y.comp[n]) { return x.less (y, n + 1) }
  return false
}


func (x *Imp) Less (Y Object) bool {
//
  return x.less (x.imp (Y), 0)
}


func (x *Imp) SetColours (f, b col.Colour) {
//
//  x.cF, x.cB = f, b
}


func (x *Imp) Write (l, c uint) {
//
  if x.mask != nil {
    x.mask.Write (l, c)
  }
  for i:= uint(0); i < x.num; i++ {
    if x.l[i] < 512 {
      x.comp[i].Write (l + x.l[i], c + x.c[i])
    }
  }
}


func (x *Imp) Edit (l, c uint) {
//
  x.Write (l, c)
  i:= uint(0)
  loop: for {
    x.comp[i].Edit (l + x.l[i], c + x.c[i])
    var d uint
    switch kbd.LastCommand(&d) {
    case kbd.Esc:
      break loop
    case kbd.Enter:
      if d == 0 {
        if i + 1 < x.num {
          i ++
        } else {
          break loop
        }
      } else {
        break loop
      }
    case kbd.Down:
      if i + 1 < x.num {
        i ++
      } else {
        i = 0
      }
    case kbd.Up:
      if i > 0 {
        i --
      } else {
        i = x.num - 1
      }
    case kbd.Pos1:
      i = 0
    case kbd.End:
      i = x.num - 1
    }
  }
}


func (x *Imp) SetFont (f font.Font) {
//
  for i:= uint(0); i < x.num; i++ {
    x.SetFont (f)
  }
}


func (x *Imp) Print (l, c uint) {
//
//  x.masks.Print (l, c)
  for i:= uint(0); i < x.num; i++ {
    x.comp[i].Print (x.l[i], x.c[i])
  }
}


func (x *Imp) Codelen () uint {
//
  c:= uint(4)
  for k:= uint(0); k < x.num; k++ {
    c += x.comp[k].Codelen()
  }
  return c
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  i, a:= uint(0), uint(4)
  copy (b[i:i+a], Encode (x.num))
  i += a
  for k:= uint(0); k < x.num; k++ {
    a = x.comp[k].Codelen()
    copy (b[i:i+a], x.comp[k].Encode())
    i += a
  }
  return b
}


func (x *Imp) Decode (b []byte) {
//
  i, a:= uint(0), uint(4)
  x.num = Decode (uint(0), b[i:i+a]).(uint)
  i += a
  for k:= uint(0); k < x.num; k++ {
    a = x.comp[k].Codelen()
    x.comp[k].Decode (b[i:i+a])
    i += a
  }
}
