package phil

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
 "murus/ker"; "murus/rand"
//  "murus/env"
//  "murus/str"; "murus/nat"
)
const (
  min =  5 // minimal and
  max = 12 // maximal number of philosophers
)
type
  status byte; const (
  satisfied = status(iota)
  hungry
  starving
  hasRightFork
  hasLeftFork
  dining
  nStatuses
)
const (
  lock = iota; unlock; nLock
)
var (
  nPhilos uint = 5 // min .. max
  stat [max]status
  text [nStatuses]string
)


func NPhilos () uint {
//
  return nPhilos
}


func setNumber (n uint) {
//
  if n < min {
    nPhilos = min
  } else if n < max {
    nPhilos = n
  } else {
    nPhilos = max
  }
}


func left (p uint) uint {
//
  return (p + nPhilos - 1) % nPhilos
}


func right (p uint) uint {
//
  return (p + 1) % nPhilos
}


func wait (t uint) {
//
  ker.Sleep (1 + (rand.Natural (3)))
}


func changeStatus (p uint, s status) {
//
  stat[p] = s
  write (p)
  switch s /* tat[p] */ { case satisfied, dining:
    // waiting time is defined in the main program
  default:
//    wait (1)
  }
}


func Eat (p uint) { // p < max
//
  switch p { case 0, 2:
    wait (3) // those two prefer to eat ...
  default:
    wait (2)
  }
//  changeStatus (p, satisfied)
}


func Think (p uint) { // p < max
//
  switch p { case 0, 2:
    wait (4) // ... instead to think
  default:
    wait (6)
  }
}

/*
func Codelen () uint {
//
  return nPhilos
}


func Code () []byte {
//
  b:= make ([]byte, Codelen())
  for p:= uint(0); p < nPhilos; p++ {
    b[p] = byte(stat[p])
  }
  return B
}


func Decode (b []byte) {
//
  for p:= uint(0); n < nPhilos; p++ {
    stat[p] = status(b[p])
  }
}
*/

/*
func init () {
//
  par:= env.Par (1)
  var p uint // min .. <= max
  if ! str.Empty (par) && nat.Defined (&p, par) {
    if p < min { p = min } else if p > max { p = max }
    setNumber (p)
  }
}
*/
