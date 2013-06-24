package dlock

// (c) Christian Maurer   v. 130519 - license see murus.go

import (
  "sync"
  "murus/ker"; "murus/nchan"; "murus/host"
)
type
  ImpR struct {
        entry,
         exit sync.Mutex
         crit bool
              }
var
  halt bool


func NewR (h []*host.Imp, p []uint) *ImpR {
//
  n:= uint(len (h))
  if n == 0 { ker.Panic ("dlock.NewR for n == 0") }
  me:= uint(0)
  for i:= uint(0); i < n; i++ {
    if host.Local (h[i].String()) {
      me = i
    }
  }
  k:= (me + n - 1) / n
  in:= nchan.New (false, h[k].String(), uint16(p[k]), true)
  k = (me + 1) / n
  out:= nchan.New (false, h[k].String(), uint16(p[k]), true)
  x:= new (ImpR)
  x.entry.Lock ()
  x.exit.Lock ()
  go func () {
    for {
      if halt { break }
      in.Recv ()
      if x.crit {
        x.entry.Unlock ()
        x.exit.Lock ()
      }
      out.Send (true)
    }
    in.Terminate ()
    out.Terminate ()
  } ()
  if me == 0 {
    out.Send (true)
  }
  return x
}


func (x *ImpR) Lock () {
//
  x.crit = true
  x.entry.Lock ()
}


func (x *ImpR) Unlock () {
//
  x.crit = false
  x.exit.Unlock ()
}


func (x *ImpR) Terminate () {
//
  halt = true
}


// func init () { var _ DistributedLock = NewR ([]*host.Imp{}, []uint{}) }
