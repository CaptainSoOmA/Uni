package bytes

// (c) Christian Maurer   v. 120910 - license see murus.go

// >>> Just for fun, most likely completely worthless

import
  . "murus/obj"
type
  Imp struct {
           s []byte
             }


func New (n uint) *Imp {
//
  return &Imp { make ([]byte, n) }
}


func (x *Imp) Empty () bool {
//
  for _, a:= range (x.s) {
    if a != byte(0) {
      return false
    }
  }
  return true
}


func (x *Imp) Clr () {
//
  for i:= 0; i < len (x.s); i++ {
    x.s[i] = byte(0)
  }
}


func (x *Imp) Copy (Y Object) {
//
  y, ok:= Y.(*Imp)
  if ! ok || len (y.s) != len (x.s) { return }
  Copy (x.s, y.s)
}


func (x *Imp) Clone () Object {
//
  y:= New (uint(len (x.s)))
  Copy (y.s, x.s)
  return y
}


func (x *Imp) Eq (Y Object) bool {
//
  if Y == nil { return false }
  y, ok:= Y.(*Imp)
  if ! ok || len (y.s) != len (x.s) { return false }
  for i, a:= range (y.s) {
    if x.s[i] != a {
      return false
    }
  }
  return true
}


func (x *Imp) Codelen () uint {
//
  return uint(len (x.s))
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, len (x.s))
  copy (b, x.s)
  return b
}


func (x *Imp) Decode (b []byte) {
//
  if len (b) == len (x.s) {
    copy (x.s, b)
  } else {
    x.Clr ()
  }
}
