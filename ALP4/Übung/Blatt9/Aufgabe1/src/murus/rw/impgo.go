package rw

// (c) Christian Maurer   v. 130408 - license see murus.go

// >>> 1st readers/writers problem (readers' preference)
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. TODO

import
  . "sync"
type
  ImpGo struct {
               RWMutex
               }


func NewGo () *ImpGo {
//
  return new (ImpGo)
}


func (x *ImpGo) ReaderIn () {
//
  x.RLock ()
}


func (x *ImpGo) ReaderOut () {
//
  x.RUnlock ()
}


func (x *ImpGo) WriterIn () {
//
  x.RWMutex.Lock ()
}


func (x *ImpGo) WriterOut () {
//
  x.RWMutex.Unlock ()
}
