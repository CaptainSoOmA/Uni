package menue

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  "murus/ker"; . "murus/obj"; "murus/str"
  "murus/col"; "murus/scr"; "murus/box"; "murus/errh"
  "murus/sel"
)
type
  Imp struct {
        text string
      isMenu,
   withTitle bool
   nextLevel []*Imp
     execute Stmt
     lastPos uint
       next *Imp
             }
var (
  bx *box.Imp = box.New ()
  menuheadF, menuheadB, menuF, menuB col.Colour
  depth uint
)


func (x *Imp) imp (Y Menue) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func New (s string) *Imp {
//
  x:= new (Imp)
  str.Set (&x.text, s)
  x.text += " "
  str.Norm (&x.text, scr.NColumns())
  x.isMenu = true
  x.execute = Null
  return x
}


func (x *Imp) Leaf (s Stmt, w bool) {
//
  if x.nextLevel != nil { return }
  x.withTitle = w
  x.isMenu = false
  x.execute = s
}


func (x *Imp) Ins (Y Menue) {
//
  y:= x.imp (Y)
  if y == nil || ! x.isMenu { return }
  n:= uint(len (x.nextLevel))
  if n >= scr.NLines() - 2 { return }
  x.nextLevel = append (x.nextLevel, y)
}


func (x *Imp) selected (l, c uint) bool {
//
  n:= uint(len (x.nextLevel))
  if n == 0 || ! x.isMenu { return false }
  if n == 1 { return true }
  bx.Colours (menuheadF, menuheadB)
  bx.Wd (scr.NColumns())
  bx.Write (x.text, l, c)
  errh.Hint (errh.ToSelect)
  i:= x.lastPos
  sel.Select (func (p, l, c uint, f, b col.Colour) {
                bx.Colours (f, b)
                bx.Write (x.nextLevel[p].text, l, c)
              }, n, scr.NLines() - 2, scr.NColumns(), &i, 2, 0, menuF, menuB)
  if i < n {
    x.lastPos = i
    x.next = x.nextLevel[i]
  }
  errh.DelHint()
  return i < n
}


func (x *Imp) Exec () {
//
  l, c:= uint(0), uint(0)
  depth ++
  if x.isMenu {
    for {
      if x.selected (l, c) {
        x.next.Exec ()
      } else {
        break
      }
    }
  } else {
    if x.withTitle {
      bx.Wd (scr.NColumns())
      bx.Colours (menuheadF, menuheadB)
      bx.Write (x.text, l, c)
    } else {
      scr.Clr (l, c, scr.NColumns(), 1)
    }
    x.execute ()
    scr.Cls ()
  }
  depth --
  if depth == 0 { ker.Terminate () }
}


func init () {
//
  menuheadF, menuheadB = col.LightWhite, col.Blue
  menuF, menuB = col.LightWhite, col.Red
}
