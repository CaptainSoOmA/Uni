package adj

// (c) Christian Maurer v. 130115 - license see murus.go

import (
  "murus/ker"
  . "murus/obj"; "murus/col"; "murus/box"
)
const
  pack = "adj"
type
  Imp struct {
           d,     // number of rows == number of columns
           n,     // Codelen of nodes
           e uint //            edges
        node []Any
   emptyEdge Any
        edge [][]Any
        f, g col.Colour
             }
var (
  bx *box.Imp = box.New()
  st [2]string = [2]string { " ", "*" }
)


func (x *Imp) imp (Y Object) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  if ! TypeEq (x.node[0], y.node[0]) { TypeNotEqPanic (x.node[0], y.node[0]) }
  return y
}


func New (n []Any, a Any) *Imp { // a Stringer ?
//
  x:= new (Imp)
  x.d = uint(len (n))
  x.node = make ([]Any, x.d)
  for i:= uint(0); i < x.d; i++ {
    x.node[i] = Clone (n[i])
  }
  x.n = Codelen (x.node[0])
  x.emptyEdge = Clone (a)
  x.edge = make ([][]Any, x.d)
  x.e = Codelen (x.emptyEdge)
  for i:= uint(0); i < x.d; i++ {
    x.edge[i] = make ([]Any, x.d)
    for k:= uint(0); k < x.d; k++ {
      x.edge[i][k] = Clone (x.emptyEdge)
    }
  }
  x.f, x.g = col.ScreenF, col.ScreenB
  return x
}


func (x *Imp) Empty () bool {
//
  for i:= uint(0); i < x.d; i++ {
    for k:= uint(0); k < x.d; k++ {
      if ! Eq (x.edge[i][k], x.emptyEdge) {
        return false
      }
    }
  }
  return true
}


func (x *Imp) Clr () {
//
  for i:= uint(0); i < x.d; i++ {
    for k:= uint(0); k < x.d; k++ {
      x.edge[i][k] = Clone (x.emptyEdge)
    }
  }
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= x.imp (Y)
  for i:= uint(0); i < x.d; i++ {
    if ! Eq (x.node[i], y.node[i]) {
      return false
    }
    for k:= uint(0); k < x.d; k++ {
      if ! Eq (x.edge[i][k], y.edge[i][k]) {
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


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.d, x.n, x.e = y.d, y.n, y.e
  for i:= uint(0); i < x.d; i++ {
    x.node[i] = Clone (y.node[i])
    for k:= uint(0); k < x.d; k++ {
      x.edge[i][k] = Clone (y.edge[i][k])
    }
  }
}


func (x *Imp) Clone () Object {
//
  y:= New (x.node, x.emptyEdge)
  y.Copy (x)
  return y
}


func (x *Imp) Codelen () uint {
//
  return /* 4 + */ x.d * (x.n + x.d * x.e)
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  a:= uint(0)
//  copy (b[0:4], Encode (x.d))
//  a += 4
  for i:= uint(0); i < x.d; i++ {
    copy (b[a:a+x.n], Encode (x.node[i]))
    a += x.n
    for k:= uint(0); k < x.d; k++ {
      copy (b[a:a+x.e], Encode (x.edge[i][k]))
      a += x.e
    }
  }
  return b
}


func (x *Imp) Decode (b []byte) {
//
  a:= uint(0)
//  x.d = Decode (uint(0), b[0:4]).(uint)
//  a += 4
  for i:= uint(0); i < x.d; i++ {
    Decode (x.node[i], b[a:a+x.n])
    a += x.n
    for k:= uint(0); k < x.d; k++ {
      x.edge[i][k] = Decode (x.edge[i][k], b[a:a+x.e])
      a += x.e
    }
  }
}


func (x *Imp) SetColours (f, g col.Colour) {
//
  x.f, x.g = f, g
}


func (x *Imp) Write (l, c uint) {
//
  bx.Colours (x.f, x.g)
  for i:= uint(0); i < x.d; i++ {
    for k:= uint(0); k < x.d; k++ {
      b:= 1
      if Eq (x.edge[i][k], x.emptyEdge) {
        b = 0
      }
      bx.Write (st[b], l + i, c + 2 * k)
    }
  }
}


func (x *Imp) Ok () bool {
//
  for i:= uint(0); i < x.d; i++ {
    if ! Eq (x.edge[i][i], x.emptyEdge) {
      return false
    }
  }
  return true
}


func (x *Imp) Loop () uint {
//
  for i:= uint(0); i < x.d; i++ {
    if Eq (x.edge[i][i], x.emptyEdge) {
      return i
    }
  }
  return x.d
}


func (x *Imp) Num () uint {
//
  return x.d
}


func (x *Imp) Node (i uint) Any {
//
  if i >= x.d { return nil }
  return Clone (x.node[i])
}


func (x *Imp) Edge (i, k uint) Any {
//
  if i >= x.d || k >= x.d {
    return nil
  }
  return Clone (x.edge[i][k])
}


func (x *Imp) Symmetric () bool {
//
  for i:= uint(0); i < x.d; i++ {
    for k:= uint(0); k < x.d; k++ {
      if ! Eq (x.edge[i][k], x.edge[k][i]) {
        return false
      }
    }
  }
  return true
}


func (x *Imp) Directed () bool {
//
  for i:= uint(0); i < x.d; i++ {
    for k:= uint(0); k < x.d; k++ {
      if i != k || Eq (x.edge[i][k], x.edge[k][i]) {
        return false
      }
    }
  }
  return true
}


func (x *Imp) Add (Y AdjacencyMatrix) {
//
  y, ok:= Y.(*Imp)
  if ! ok {
    return
  }
  for i:= uint(0); i < x.d; i++ {
    for k:= uint(0); k < x.d; k++ {
      if ! Eq (y.edge[i][k], x.emptyEdge) && ! Eq (x.edge[i][k], y.edge[i][k]) {
        if ! Eq (x.edge[i][k], x.emptyEdge) { ker.Stop (pack, 123456) }
        x.edge[i][k] = Clone (y.edge[i][k])
      }
    }
  }
}


func (x *Imp) Invert () {
//
  for i:= uint(0); i < x.d; i++ {
    for k:= uint(0); k < x.d; k++ {
      if i != k {
        x.edge[i][k], x.edge[k][i] = x.edge[k][i], x.edge[i][k]
      }
    }
  }
}


func (x *Imp) Put (i, k uint, a Any) {
//
  if ! TypeEq (a, x.emptyEdge) { return }
  if i >= x.d || k >= x.d { return }
  x.edge[i][k] = Clone (a)
}


func (x *Imp) Del (i, k uint) {
//
  if i >= x.d || k >= x.d { return }
  x.edge[i][k] = Clone (x.emptyEdge)
}


func (x *Imp) Full () bool{
//
  for i:= uint(0); i < x.d; i++ {
    f:= false
    for k:= uint(0); k < x.d; k++ {
      f = f || ! Eq (x.edge[i][k], x.emptyEdge)
    }
    if ! f { return false }
  }
  return true
}


func init () {
//
  bx.Wd (1)
  var _ AdjacencyMatrix = New ([]Any {1, 2}, false)
}
