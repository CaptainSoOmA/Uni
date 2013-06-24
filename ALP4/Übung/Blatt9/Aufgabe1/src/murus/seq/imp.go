package seq

// (c) Christian Maurer   v. 130128 - license see murus.go

import
  . "murus/obj"
const
  pack = "seq"
type (
  cell struct {
         head Any
         next,
         prev *cell
       }
  Imp struct {
         num,
         pos uint
      anchor,
      actual *cell
     ordered bool
             }
)
var
  cluint = Codelen (uint(0))


func (x *Imp) check (a Any) {
//
  if ! TypeEq (x.anchor.head, a) {
    TypeNotEqPanic (x, a)
  }
}


func (x *Imp) imp (a Any) *Imp {
//
  y, ok:= a.(*Imp)
  if ! ok { TypePanic() }
  x.check (y.anchor.head)
  return y
}


func checked (a Any) bool {
//
  return AtomicOrObject (a)
}


func New (a Any) *Imp {
//
  PanicIfNotOk (a)
  x:= new (Imp)
  x.anchor = new (cell)
  x.anchor.head = Clone (a)
  x.anchor.next, x.anchor.prev = x.anchor, x.anchor
  x.actual = x.anchor
  return x
}


func (x *Imp) Empty () bool {
//
  return x.anchor.next == x.anchor
}


func (x *Imp) remove () {
//
  a:= x.actual.next
  x.actual.prev.next = a
  a.prev = x.actual.prev
  x.actual.prev, x.actual.next = nil, nil
  x.actual = a
}


func (x *Imp) Clr () {
//
  x.actual = x.anchor.next
  for x.actual != x.anchor {
    x.remove ()
  }
  x.num, x.pos = 0, 0
}


func (x *Imp) ins (a Any) {
//
  c:= new (cell)
  c.head = Clone (a)
  c.next, c.prev = x.actual, x.actual.prev
  x.actual.prev.next = c
  x.actual.prev = c
}


func (x *Imp) Copy (o Object) {
//
  y:= x.imp (o)
  x.Clr ()
  x.anchor.head = Clone (y.anchor.head)
  for l:= y.anchor.next; l != y.anchor; l = l.next {
    x.ins (l.head)
  }
  x.num, x.pos = y.num, y.num
}


func (x *Imp) Clone () Object {
//
  var y *Imp
  y.Copy (x)
  return y
}


func (x *Imp) e (y *Imp, r Rel) bool {
//
  if x.num != y.num { return false }
  for l, l1:= x.anchor.next, y.anchor.next; l != x.anchor; l, l1 = l.next, l1.next {
    if ! r (l.head, l1.head) {
      return false
    }
  }
  return true
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.e (x.imp (Y), Eq)
}


func (x *Imp) Less (Y Object) bool {
//
  y:= x.imp (Y)
  if x.num >= y.num { return false }
  l:= x.anchor.next
  l1:= y.anchor.next
  for l != x.anchor {
    for {
      if l1 == y.anchor {
        return false
      }
      if Eq (l.head, l1.head) {
        l1 = l1.next
        break
      }
      l1 = l1.next
    }
    l = l.next
  }
  return true
}


func (x *Imp) Equiv (Y Iterator, r Rel) bool {
//
  return x.e (x.imp (Y), r)
}


func (x *Imp) Num () uint {
//
  return x.num
}


func (x *Imp) NumPred (p Pred) uint {
//
  n:= uint(0);
  for l:= x.anchor.next; l != x.anchor; l = l.next {
    if p (l.head) {
      n++
    }
  }
  return n
}


func (x *Imp) Ex (a Any) bool {
//
  x.check (a)
  p:= uint(0)
  for l:= x.anchor.next; l != x.anchor; l = l.next {
    if Eq (l.head, a) {
      x.actual = l
      x.pos = p
      return true
    }
    p++
  }
  return false
}


func (x *Imp) Step (forward bool) {
//
  if forward {
    if x.actual != x.anchor {
      x.actual = x.actual.next
      x.pos ++
    }
  } else if x.actual != x.anchor.next {
    x.actual = x.actual.prev
    x.pos --
  }
}


func (x *Imp) Jump (forward bool) {
//
  if x.num > 0 {
    if forward {
      x.actual = x.anchor.prev
      x.pos = x.num - 1
    } else {
      x.actual = x.anchor.next
      x.pos = 0
    }
  }
}


func (x *Imp) Eoc (forward bool) bool {
//
  if x.actual == x.anchor {
    return false
  }
  if forward {
    return (x.actual.next == x.anchor)
  }
  return (x.actual.prev == x.anchor)
}


func (x *Imp) Offc () bool {
//
  return x.actual == x.anchor
}


func (x *Imp) Pos () uint {
//
  return x.pos
}


func (x *Imp) Seek (i uint) {
//
  if i == 0 {
    x.actual = x.anchor.next
    x.pos = 0
    return
  }
  if i >= x.num {
    x.actual = x.anchor
    x.pos = x.num
    return
  }
  if i == 1 {
    x.actual = x.anchor.next.next
    x.pos = 1
    return
  }
  if i + 1 == x.num {
    x.actual = x.anchor.prev
    x.pos = x.num - 1
    return
  }
  for x.pos < i {
    x.actual = x.actual.next
    x.pos ++
  }
  for x.pos > i {
    x.actual = x.actual.prev
    x.pos --
  }
}


func (x *Imp) Get () Any {
//
  if x.actual == x.anchor {
    return nil
  }
  return Clone (x.actual.head)
}


func (x *Imp) Put (a Any) {
//
  x.check (a)
  if x.actual == x.anchor {
    x.ins (a)
    x.actual = x.actual.prev
    x.pos = x.num
    x.num ++
  } else {
    x.actual.head = Clone (a)
  }
}


func (x *Imp) insert (a Any) {
//
  x.ins (a)
  x.num ++
  x.pos ++
}


func (x *Imp) Ins (a Any) {
//
  x.check (a)
  if x.ordered {
    x.actual = x.anchor.next
    x.pos = 0
    for x.actual != x.anchor {
      if Less (x.actual.head, a) {
        x.actual = x.actual.next
        x.pos ++
      } else {
        if Less (a, x.actual.head) {
          break
        } else { // already there
          return
        }
      }
    }
  }
  x.insert (a)
}


func (x *Imp) InsRel (a Any, r Rel) {
//
  x.check (a)
  x.actual, x.pos = x.anchor.next, 0
  for x.actual != x.anchor {
    if r (x.actual.head, a) {
      x.actual = x.actual.next
      x.pos ++
    } else {
      break
    }
  }
  x.insert (a)
}


func (x *Imp) Del () Any {
//
  if x.actual == x.anchor {
    return nil
  }
  defer x.remove ()
  x.num --
  return Clone (x.actual.head)
}


func (x *Imp) ExPred (p Pred, forward bool) bool {
//
//  s, i:= x.actual, x.pos
  if x.num == 0 {
    return false
  }
  s, i:= x.anchor.next, uint(0)
  if ! forward {
    s, i = x.anchor.prev, x.num - 1
  }
  for s != x.anchor {
    if p (s.head) {
      x.actual, x.pos = s, i
      return true
    }
    if forward {
      s = s.next
      i++
    } else {
      s = s.prev
      i--
    }
/*
    if p (s.head) {
      x.actual, x.pos = s, i
      return true
    }
*/
  }
  return false
}


func (x *Imp) StepPred (p Pred, forward bool) bool {
//
  s, i:= x.actual, x.pos
  for {
    if forward {
      s = s.next
      i++
    } else {
      s = s.prev
      i--
    }
    if s == x.anchor { break }
    if p (s.head) {
      x.actual = s
      x.pos = i
      return true
    }
  }
  return false
}


func (x *Imp) All (p Pred) bool {
//
  for l:= x.anchor.next; l != x.anchor; l = l.next {
    if ! p (l.head) {
      return false
    }
  }
  return true
}


func (x *Imp) Ordered () bool {
//
  l:= x.anchor.next
  if l == x.anchor { return true }
  for l.next != x.anchor {
    if Less (l.head, l.next.head) {
      l = l.next
    } else {
      return false
    }
  }
  return true
}


func (x *Imp) Sort () {
//
  x.ordered = true
  if x.Ordered () { return }
  l:= x.anchor.next
  if l == x.anchor { return }
  if l.next == x.anchor { return }
  x.anchor.next = l.next
  l.next.prev = x.anchor
  x.num --
  var y *Imp
  y = New (x.anchor.head)
  l1:= x.anchor.next
  var l2 *cell
  for l1 != x.anchor {
    l2 = l1.next
    if Less (l.head, l1.head) {
      l1.prev.next = l1.next
      l1.next.prev = l1.prev
      l1.next, l1.prev = y.anchor, y.anchor.prev
      l1.prev.next = l1
      y.anchor.prev = l1
      x.num --
      y.num ++
    }
    l1 = l2
  }
  x.Sort ()
  y.Sort ()
  l.next = y.anchor.next
  y.anchor.next = l
  l.prev = y.anchor
  l.next.prev = l
  y.num ++
  x.concatenate (y)
  x.actual = x.anchor
  x.pos = x.num
}


func (x *Imp) ExGeq (a Any) bool {
//
  if ! x.ordered {
    return false // TODO Panic ?
  }
  x.check (a)
  p:= uint(0)
  for l:= x.anchor.next; l != x.anchor; l, p = l.next, p + 1 {
    if Less (a, l.head) {
      x.actual = l
      x.pos = p
      return true
    }
  }
  return false
}


func (x *Imp) Trav (op Op) {
//
  for l:= x.anchor.next; l != x.anchor; l = l.next {
    op (l.head)
  }
}


func (x *Imp) TravCond (p Pred, op CondOp) {
//
  for l:= x.anchor.next; l != x.anchor; l = l.next {
    op (l.head, p (l.head))
  }
}


func (x *Imp) TravPred (p Pred, op Op) {
//
  for l:= x.anchor.next; l != x.anchor; l = l.next {
    if p (l.head) {
      op (l.head)
    }
  }
}


func (x *Imp) Filter (Y Iterator, p Pred) {
//
  y:= x.imp (Y)
  if y == x { return }
  y.Clr ()
  for l:= x.anchor.next; l != x.anchor; l = l.next {
    if p (l.head) {
      y.ins (l.head)
      y.num ++
    }
  }
  y.pos = x.num
}


func (x *Imp) Split (Y Iterator) {
//
  y:= x.imp (Y)
  if y == x { return }
  y.Clr ()
  if x.actual == x.anchor { return }
  y.anchor.next, y.anchor.prev = x.actual, x.anchor.prev
  x.anchor.prev.next = y.anchor
  x.anchor.prev = x.actual.prev
  x.actual.prev.next = x.anchor
  x.actual.prev = y.anchor
  x.actual = x.anchor
  y.actual = y.anchor.next
  y.num = x.num - x.pos
  x.num = x.pos
  x.pos = x.num
}


func (x *Imp) Cut (Y Iterator, p Pred) {
//
  y:= x.imp (Y)
  if y == x { return }
  y.Clr ()
  l:= x.anchor.next
  var l1 *cell
  for l != x.anchor {
    l1 = l.next
    if p (l.head) {
      l.prev.next = l.next
      l.next.prev = l.prev
      l.next, l.prev = y.anchor, y.anchor.prev
      l.prev.next = l
      y.anchor.prev = l
      x.num --
      y.num ++
    }
    l = l1
  }
  x.actual = x.anchor
  x.pos = x.num
  y.actual = y.anchor
  y.pos = y.num
}


func (x *Imp) ClrPred (p Pred) {
//
  l:= x.anchor.next
  for l != x.anchor {
    a:= l
    l:= l.next
    if p (a.head) {
      a.prev.next = a.next
      a.next.prev = a.prev
      a.prev, a.next = nil, nil
      if x.actual == a {
        x.actual = l
        x.pos ++
      }
      x.num --
    }
  }
}


func (x *Imp) concatenate (y *Imp) {
//
  if y.Empty () { return }
  if ! Eq (x.anchor.head, y.anchor.head) { return }
  x.anchor.prev.next = y.anchor.next
  y.anchor.next.prev = x.anchor.prev
  y.anchor.prev.next = x.anchor
  x.anchor.prev = y.anchor.prev
  x.num += y.num
  if x.actual == x.anchor {
    x.actual = y.anchor.next
  }
  y.Clr()
}


func (x *Imp) join (y *Imp) {
//
  if y.anchor == y.anchor.next {
    x.actual = x.anchor
    x.pos = x.num
    return
  }
  l:= x.anchor.next
  y.actual = y.anchor.next
  for {
    if y.actual == y.anchor { break }
    if l == x.anchor { break }
    if Less (y.actual.head, l.head) {
      y.anchor.next = y.actual.next
      y.actual.prev = l.prev
      l.prev.next = y.actual
      l.prev = y.actual
      y.actual.next = l
      y.actual = y.anchor.next
    } else {
      l = l.next
    }
  }
  if y.actual != y.anchor {
    l = l.prev // == x.anchor.prev
    l.next = y.actual
    y.actual.prev = l
    x.anchor.prev = y.anchor.prev
    y.anchor.prev.next = x.anchor
  }
  x.actual = x.anchor
  x.num += y.num
  x.pos = x.num
  y.anchor.next, y.anchor.prev = y.anchor, y.anchor
  y.actual = y.anchor
  y.num, y.pos = 0, 0
}


func (x *Imp) Join (Y Iterator) {
//
  y:= x.imp (Y)
  if x.ordered {
    x.join (y)
  } else {
    x.concatenate (y)
  }
}


// Not documented - destroys the order, if x is ordered !!!
func (x *Imp) Reverse () {
//
//  if x.ordered { return }
  l:= x.anchor
  l1:= l.next
  for l1 != x.anchor {
    l1 = l.next
    l.next = l.prev
    l.prev = l1
    l = l1
  }
}


// Not documented - destroys the order, if x is ordered !!!
func (x *Imp) Rotate (forward bool) {
//
//  if x.ordered { return }
  if x.anchor.next == x.anchor || x.anchor.next == x.anchor.prev {
    return
  }
  if forward {
    l:= x.anchor.prev
    l.prev.next = x.anchor
    x.anchor.prev = l.prev
    l.prev = x.anchor
    l.next = x.anchor.next
    x.anchor.next = l
    l.next.prev = l
  } else {
    l:= x.anchor.next
    l.next.prev = x.anchor
    x.anchor.next = l.next
    l.next, l.prev = x.anchor, x.anchor.prev
    x.anchor.prev = l
    l.prev.next = l
  }
}


func (x *Imp) MinCodelen () uint {
//
  return Codelen (x.num)
}


func (x *Imp) Codelen () uint {
//
  n:= cluint
  for l:= x.anchor.next; l != x.anchor; l = l.next {
    n += cluint + Codelen (l.head)
  }
  return n
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  i, a:= uint(0), cluint
  copy (b[i:a], Encode (x.num))
  i += cluint
  for l:= x.anchor.next; l != x.anchor; l = l.next {
    n:= Codelen (l.head)
    copy (b[i:i+a], Encode (n))
    i += a
    copy (b[i:i+n], Encode (l.head))
    i += n
  }
  return b
}


func (x *Imp) Decode (b []byte) {
//
  x.Clr ()
  i, a:= uint(0), cluint
  x.num = Decode (x.num, b[i:a]).(uint)
  i += a
  for j:= uint(0); j < x.num; j++ {
    n:= Decode (uint(0), b[i:i+a]).(uint)
    i += a
    x.ins (Decode (x.anchor.head, b[i:i+n]))
    i += n
  }
}


func (x *Imp) Slice () []Any {
//
  a:= make ([]Any, x.Num())
  for i, l:= 0, x.anchor.next; l != x.anchor; i, l = i+1, l.next {
    a[i] = Clone (l.head)
  }
  return a
}


func (x *Imp) Deslice (b []Any) {
//
  x.Clr()
  for _, a:= range b {
    x.Ins (a)
  }
}


func init () { var _ Sequence = New (0) }
