package barb

// (c) Christian Maurer   v. 130424 - license see murus.go

import (
  . "murus/obj"
  "murus/mon"
)
type
  ImpMon struct {
                mon.Monitor
                }


func NewMon () *ImpMon {
//
  var x mon.Monitor
  var n uint
//  barberFree:= true
  do:= func (a Any, i uint) Any {
         if i == customer {
//           for ! barberFree {
//             x.Wait (barber)
//           }
//           barberFree = false
           n ++
           x.Signal (customer)
         } else { // i == barber
//           barberFree = true
//           x.Signal (barber)
           for n == 0 {
             x.Wait (customer)
           }
           n --
         }
         return 0
       }
  x = mon.New (2, do)
  return &ImpMon { x }
}


func (x *ImpMon) Customer () {
//
  x.F (nil, customer)
}


func (x *ImpMon) Barber () {
//
  x.F (nil, barber)
}
