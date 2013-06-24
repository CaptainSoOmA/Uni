package conn

// (c) Christian Maurer   v. 130309 - license see murus.go

import (
  . "murus/obj"; "murus/str"; "murus/nat"
  "murus/host"
)
type
  Imp struct {
             *host.Imp
        port uint16
             Format
             }
var
  cl uint


func (x *Imp) imp (Y Object) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func New () *Imp {
//
  return &Imp { host.New (), 0, 0 }
}


func (x *Imp) Empty () bool {
//
  return x.Imp.Empty ()
}


func (x *Imp) Clr () {
//
  x.Imp.Clr ()
  x.port = 0
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= x.imp (Y)
  return x.Imp.Eq (y.Imp) &&
         x.port == y.port
}


func (x *Imp) Less (Y Object) bool {
//
  y:= x.imp (Y)
  if x.Imp.Eq (y.Imp) {
    return x.port < y.port
  }
  return x.Imp.Less (y.Imp)
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.Imp.Copy (y.Imp)
  x.port = y.port
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Codelen () uint {
//
  return x.Imp.Codelen() + 2
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  copy (b[:cl], x.Imp.Encode ())
  copy (b[cl:], Encode (x.port))
  return b
}


func (x *Imp) Decode (b []byte) {
//
  x.Imp.Decode (b[:cl])
  x.port = Decode (x.port, b[cl:]).(uint16)
}


func (x *Imp) SetFormat (f Format) {
//
  if f < host.NFormats {
    x.Format = f
  }
}


const
  separator = ':'


func (x *Imp) Defined (s string) bool {
//
  n, i, p:= uint(len (s)), uint(0), uint(0)
  if str.Contains (s, separator, &i) && i < n &&
     nat.Defined (&p, str.Part (s, i + 1, n - (i + 1))) {
    return x.Defined2 (str.Part (s, 0, i), p)
  }
  return false
}


func (x *Imp) String () string {
//
  x.Imp.SetFormat (x.Format)
  return x.Imp.String () + string (separator) + nat.String (uint(x.port))
}


func (x *Imp) Defined2 (s string, p uint) bool {
//
  if x.Imp.Defined (s) && p < 1<<16 {
    x.port = uint16(p)
    return true
  }
  return false
}

/*
func (x *Imp) HostPort () (string, uint) {
//
  return x.Imp.String (), uint(x.port))
}
*/

func init () {
//
  var c *Imp = New ()
  cl = c.Imp.Codelen()
}
