package cr

// (c) Christian Maurer   v. 130409 - license see murus.go

import (
  "murus/ker"
  . "murus/obj"
  "murus/cs"
)
const
  pack = "cr"
type (
  status struct {
            max []uint // indexed over process classes
         number,
          class uint
                }
  Imp struct {
        stat []status // indexed over resources
      nC, nR uint
             *cs.Imp
             }
)


func New (nc, nr uint) *Imp {
//
  x:= new (Imp)
  x.nC, x.nR = nc, nr
  x.stat = make ([]status, x.nC)
  for r:= uint(0); r < x.nR; r++ {
    x.stat[r].max = make ([]uint, x.nC)
    for c:= uint(0); c < x.nC; c++ {
      x.stat[r].max[c] = ker.MaxNat
    }
  }
  c:= func (k uint) bool {
        var b bool
        for r:= uint(0); r < x.nR; r++ {
          b = b ||
              x.stat[r].number == 0 ||
              x.stat[r].class == k && x.stat[r].number < x.stat[r].max[k]
        }
        return b
      }
  e:= func (a Any, k uint) {
        for r:= uint(0); r < x.nR; r++ {
          if x.stat[r].number == 0 || x.stat[r].class == k {
            x.stat[r].class = k
            x.stat[r].number ++
            n:= a.(*uint)
            *n = r
            return
          }
        }
        ker.Panic (pack + ".New() error")
      }
  l:= func (a Any, k uint) {
        for r:= uint(0); r < x.nR; r++ {
          if x.stat[r].class == k && x.stat[r].number > 0 {
            x.stat[r].number --
          }
        }
      }
  x.Imp = cs.New (x.nC, c, e, l)
  return x
}


func (x *Imp) Limit (m [][]uint) {
//
  for c:= uint(0); c < x.nC; c++ {
    for r:= uint(0); r < x.nR; r++ {
      x.stat[r].max[c] = m[c][r]
    }
  }
}


func (x *Imp) Enter (k uint) uint {
//
  var r uint
  x.Imp.Enter (k, &r)
  return r
}


func (x *Imp) Leave (k uint) {
//
  x.Imp.Leave (k, 0)
}
