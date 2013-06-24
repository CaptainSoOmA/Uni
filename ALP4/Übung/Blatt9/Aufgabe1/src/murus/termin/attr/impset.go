package attr

// (c) Christian Maurer   v. 130511 - license see murus.go

import (
  "murus/scr"
)
type
  Set struct {
           m [NAttrs]bool
             }


func NewSet () *Set {
//
  return new (Set)
}


func (x *Set) Empty () bool {
//
  for a:= Attr(1); a < NAttrs; a++ {
    if x.m[a] { return false }
  }
  return true
}


func (x *Set) Clr () {
//
  for a:= Attr(1); a < NAttrs; a++ {
    x.m[a] = false
  }
}


func (x *Set) Copy (Y AttrSet) {
//
  for a:= Attr(1); a < NAttrs; a++ {
    x.m[a] = Y.(*Set).m[a]
  }
}


func (x *Set) Ins (A Attribute) {
//
  a:= A.(*Imp)
  x.m[a.a] = true
}


func (x *Set) Write (l, c uint) {
//
  t:= ""
  if x.Empty () {
    scr.Clr (l, c, uint(NAttrs), 1)
  } else {
    for a:= Attr(1); a < NAttrs; a++ {
      if x.m[a] {
        t += string(txt[a][0])
      } else {
        t += " "
      }
    }
    setbx.Write (t, l, c)
  }
}
