package host

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  "os"; "net"
  . "murus/ker"; . "murus/obj"
  "murus/col"; "murus/box"; "murus/errh"
)
const
  pack = "host"
type
  Imp struct {
          ip net.IP // []net.IP
        name []string // maximal length max see some header-file in /usr/include[/...]
         fmt Format
      fg, bg col.Colour
      marked bool
             }
var
  localHost = New ()
var
  ll [NFormats]uint = [NFormats]uint { 32, 39 } // 32: nackte Willkür
var
  bx = box.New ()


func (x *Imp) imp (Y Object) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func New () *Imp {
//
  x:= new (Imp)
  x.fmt = Hostname
  x.fg, x.bg = col.ScreenF, col.ScreenB
  return x
}


func (x *Imp) Empty () bool {
//
  for _, b:= range (x.ip) {
    if b != 0 {
      return false
    }
  }
  return true
}


func (x *Imp) Clr () {
//
  x.ip = nil
  x.name = []string { "" }
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.ip.Equal (x.imp (Y).ip)
}


func l (a, b []byte, r Rel) bool {
//
  if len (a) != len (b) || len (a) == 0 {
    return false
  }
  if a[0] == b[0] {
    return r (a[1:], b[1:])
  }
  return r (a[0], b[0])
}


func (x *Imp) Less (Y Object) bool {
//
  return l (x.ip, x.imp (Y).ip, Less)
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.name = y.name
  if len (x.ip) != len (y.ip) {
    x.ip = make ([]byte, len (y.ip))
  }
  for i, b:= range (y.ip) {
    x.ip[i] = b
  }
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Codelen () uint {
//
  return uint(len (x.ip))
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  copy (b, x.ip)
  return b
}


func (x *Imp) Decode (b []byte) {
//
  copy (x.ip, b)
  if h, err:= net.LookupAddr (x.String ()); err == nil {
    x.name = h
  } else {
//    x.Clr ()
    x.name = []string { "" }
  }
}


func (x *Imp) SetFormat (f Format) {
//
  if f < NFormats {
    x.fmt = f
  }
}


func (x *Imp) String () string {
//
  if x.fmt == Hostname {
    return x.name [0]
  }
  return x.ip.String ()
}


func (x *Imp) Defined (s string) bool {
//
  x.Clr ()
  if x.fmt == Hostname {
    if ip, err:= net.LookupHost (s); err == nil {
      x.ip = net.ParseIP (ip[0]) // default
      for _, n:= range (ip) { // get first IPv4-number, if such exists
        if len (n) == 4 {
          x.ip = net.ParseIP (n)
          break
        }
      }
      if h, err1:= net.LookupAddr (s); err1 == nil {
        x.name = h
      } else {
        x.name = []string { s }
      }
      return true
    }
  } else { // x.fmt == IPnumber
    if h, err:= net.LookupAddr (s); err == nil {
      x.ip = net.ParseIP (s)
      x.name = h
      return true
    }
  }
  return false
}


func (x *Imp) Equiv (s string) bool {
//
  b:= net.ParseIP (s)
  if b != nil {
    return x.ip.Equal (b)
  }
  for i:= 0; i < len (x.name); i++ {
    if x.name[i] == s {
      return true
    }
  }
  return false
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.fg, x.bg = f, b
}


func (x *Imp) Write (l, c uint) {
//
  bx.Wd (ll [x.fmt])
  bx.Colours (x.fg, x.bg)
  bx.Write (x.String(), l, c)
}


func (x *Imp) Edit (l, c uint) {
//
  x.Write (l, c)
  s:= x.String ()
  for {
    bx.Edit (&s, l, c)
    if x.Defined (s) {
       break
    } else {
      errh.Error ("falsche Eingabe", 0)
    }
  }
}


func (x *Imp) Mark (m bool) {
//
  x.marked = m
}


func (x *Imp) Marked () bool {
//
  return x.marked
}


func (x *Imp) Sonstewas () bool {
//
  return l (localHost.ip, x.ip, Less)
}


func (x *Imp) Number () uint32 {
// deprecated
  return (((uint32(x.ip[0]) * 256) + uint32(x.ip[1])) * 256 + uint32(x.ip[2])) * 256 + uint32(x.ip[3])
}


func (x *Imp) IP () []byte {
//
  return x.ip
}

/*
func (x *Imp) Local () bool {
//
  return x.Eq (localHost)
}
*/

func local (s string) bool {
//
  for i:= 0; i < len (localHost.name); i++ {
    if localHost.name[i] == s {
      return true
    }
  }
  return false
}


func init () {
//
  localname, err:= os.Hostname ()
  if err != nil || ! localHost.Defined (localname) { Stop (pack, 1) }
//  var _ Host = New()
}
