package page

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  "murus/ker"
  . "murus/obj"
  "murus/col"; "murus/scr"; "murus/text"
)
const (
  pack = "pset/page"
  max = 2 * N
  cluint = uint(4) // hopefully not 8
)
type
  Imp struct {
      object Object
         len,     // Codelen of a content
         num uint // number of nonempty objects on the page
         pos [max+2]uint
     content [max+1]Object
             }


func imp (X Object) *Imp {
//
  x, ok:= X.(*Imp)
  if ! ok { NotCompatiblePanic () }
  return x
}


func object (a Any) Object {
//
  o, ok:= a.(Object)
  if ! ok { TypePanic () }
  return o
}


func New (a Any) *Imp {
//
  o:= object (a)
  x:= new (Imp)
  x.object = o.Clone ()
  x.len = o.Codelen ()
  for i:= 0; i <= max; i++ {
    x.content[i] = o.Clone ()
  }
  return x
}


func (x *Imp) Empty () bool {
//
  if x.num > 0 { return false }
  for i:= 0; i < max; i++ {
    if x.pos[i] > 0 { return false }
    if ! x.content[i].Empty() { return false }
  }
  if x.pos[max] > 0 { return false }
  return true
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= imp (Y)
  if x.num != y.num { return false }
  for i:= 0; i < max; i++ {
    if x.pos[i] != y.pos[i] { return false }
    if ! x.content[i].Eq (y.content[i]) { return false }
  }
  if x.pos[max] != y.pos[max] { return false }
  return true
}


func (x *Imp) Less (Y Object) bool {
//
  return false
}


func (x *Imp) Copy (Y Object) {
//
  y:= imp (Y)
  x.num = y.num
  for i:= 0; i < max; i++ {
    x.pos[i] = y.pos[i]
    x.content[i].Copy (y.content[i])
  }
  x.pos[max] = y.pos[max]
}


func (x *Imp) Clone () Object {
//
  y:= New (x.object)
  y.Copy (x)
  return y
}


func (x *Imp) Clr () {
//
  x.num = 0
  for i:= uint(0); i < max; i++ {
    x.pos[i] = 0
    x.content[i].Clr ()
  }
  x.pos[max] = 0
}


func (x *Imp) Codelen () uint {
//
  return cluint + max * (cluint + x.len) + cluint
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  j:= uint(0)
  a:= cluint
  copy (b[j:j+a], Encode (x.num))
  j += a
  for i:= 0; i < max; i++ {
    a = cluint
    copy (b[j:j+a], Encode (x.pos[i]))
    j += a
    a = x.len
    copy (b[j:j+a], x.content[i].Encode ())
    j += a
  }
  a = cluint
  copy (b[j:j+a], Encode (x.pos[max]))
  return b
}


func (x *Imp) Decode (b []byte) {
//
  j:= uint(0)
  a:= cluint
  x.num = Decode (x.num, b[j:j+a]).(uint)
  j += a
  for i:= 0; i < max; i++ {
    a = cluint
    x.pos[i] = Decode (x.pos[i], b[j:j+a]).(uint)
    j += a
    a = x.len
    x.content[i].Decode (b[j:j+a])
    j += a
  }
  a = cluint
  x.pos[max] = Decode (x.pos[max], b[j:j+a]).(uint)
}


func (x *Imp) PutNum (n uint) {
//
  if n > max { ker.Stop (pack, 1) }
  x.num = n
}


func (x *Imp) GetNum () uint {
//
  return x.num
}


func (x *Imp) PutPos (p, n uint) {
//
  if p > max + 1 { ker.Stop (pack, 2) }
  x.pos[p] = n
}


func (x *Imp) GetPos (p uint) uint {
//
  if p > max + 1 { ker.Stop (pack, 3) }
  return x.pos[p]
}


func (x *Imp) Put (p uint, o Object) {
//
  if p > max + 1 { ker.Stop (pack, 4) }
  x.content[p] = o.Clone ()
}


func (x *Imp) Get (p uint) Object {
//
  if p > max + 1 { ker.Stop (pack, 5) }
  return x.content[p].Clone ()
}


func (x *Imp) Oper (p uint, op Op) {
//
  op (x.content[p])
}


func (x *Imp) Ins (o Object, p, n uint) {
//
  if p < x.num {
    for i:= x.num; i >= p + 1; i-- {
      x.pos[i + 1] = x.pos[i]
      x.content[i] = x.content[i - 1]
    }
  }
  x.content[p] = o
  x.pos[p + 1] = n
  x.num ++
  if x.num < max {
    for i:= x.num; i < max; i++ {
      x.content[i] = x.object
      x.pos[i + 1] = 0
    }
  }
}


func (x *Imp) IncNum () {
//
  x.num ++
}


func (x *Imp) DecNum () {
//
  if x.num == 0 { ker.Stop (pack, 6) }
  x.num --
}


func (x *Imp) RotLeft () {
//
  for i:= uint(1); i < x.num; i++ {
    x.content[i - 1] = x.content[i]
    x.pos[i - 1] = x.pos[i]
  }
  x.content[x.num - 1] = x.object
  x.pos[x.num - 1] = x.pos[x.num]
  x.pos[x.num] = 0
  x.num --
}


func (x *Imp) RotRight () {
//
  x.pos[x.num + 1] = x.pos[x.num]
//  for i:= x.num - 1; i >= 0; i-- { // does not work, because for uint: 0-- == 2^32 - 1  !
  i:= x.num - 1
  for {
    x.content[i + 1], x.pos[i + 1] = x.content[i], x.pos[i]
    if i == 0 {
      break
    }
    i--
  }
}


func (x *Imp) Join (p uint) {
//
  if p < x.num {
    for i:= p; i < x.num; i++ {
      x.content[i - 1] = x.content[i]
      x.pos[i] = x.pos[i + 1]
    }
  }
  x.content[x.num - 1] = x.object
  x.pos[x.num] = 0
  x.num --
}


func (x *Imp) Del (p uint) {
//
  if p + 1 < x.num {
    for i:= p + 1; i < x.num; i++ {
      x.content[i - 1] = x.content[i]
      x.pos[i] = x.pos[i + 1]
    }
  }
  x.content[x.num - 1] = x.object
  x.pos[x.num] = 0
}


func (x *Imp) ClrLast () {
//
  x.content[x.num - 1] = x.object
  x.pos[x.num - 1] = x.pos[x.num] // ?
  x.pos[x.num] = 0
  x.num --
}


func (x *Imp) Write (l, c uint) {
//
  scr.Colours (col.White, col.Blue)
  scr.WriteNat (x.num, l, c)
  c += 4
  for i:= uint(0); i < max; i++ {
    scr.Colours (col.Yellow, col.Red)
    scr.WriteNat (x.pos[i], l, c)
    c += 4
    scr.Colours (col.White, col.Blue)
    scr.Write (x.content[i].(*text.Imp).String(), l, c)
    c += 10
  }
  scr.Colours (col.Yellow, col.Red)
  scr.WriteNat (x.pos[max], l, c)
}
