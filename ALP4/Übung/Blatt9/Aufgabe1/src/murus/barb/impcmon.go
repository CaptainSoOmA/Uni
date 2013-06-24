package barb

// (c) Christian Maurer   v. 130424 - license see murus.go

import (
  . "murus/obj"; "murus/mon"
)
type
  ImpCMon struct {
                 mon.Monitor
                 }


func NewCMon () *ImpCMon {
//
  var n uint
  do:= func (a Any, i uint) Any {
         if i == customer {
           n ++
         } else { // i == barber
           n --
         }
         return 0
       }
  c:= func (a Any, i uint) bool {
        if i == customer {
          return true
        }
        return n > 0 // i == barber
      }
  return &ImpCMon { mon.NewC (2, do, c) }
}


func (x *ImpCMon) Customer () {
//
  x.F (nil, customer)
}


func (x *ImpCMon) Barber () {
//
  x.F (nil, barber)
}
