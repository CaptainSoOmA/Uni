package obj

// (c) Christian Maurer   v. 130116 - license see murus.go

import (
  "reflect"
  "strconv"
  "murus/ker"
)


func DivBy0Panic () {
//
  ker.Panic ("division by 0")
}


func TypePanic () {
//
  ker.Panic ("the type does not fit to the implementation")
}


func TypeNotEqPanic (a, b Any) {
//
  ker.Panic ("the types " + reflect.TypeOf (a).String () +
                  " and " + reflect.TypeOf (b).String () + " are not equal")
}


func NotCompatiblePanic () {
//
  ker.Panic ("the two involved types are not compatible")
}


func PanicIfNotOk (a Any) {
//
  if ! Atomic (a) && ! IsObject (a) {
    ker.Panic ("parameter is neither Atomic nor implements Object")
  }
}


func WrongUintParameterPanic (s string, a Any, n uint) {
//
  ker.Panic ("method " + s +
             " for object of type " + reflect.TypeOf (a).String () +
             " got wrong value for " + strconv.FormatUint (uint64(n), 10))
}
