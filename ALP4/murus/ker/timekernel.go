package ker

// (c) Christian Maurer   v. 120417 - license see murus.go

import
  "time"


func Sleep (s uint) {
//
  time.Sleep (time.Duration(s) * 1e9)
}


func Msleep (s uint) {
//
  time.Sleep (time.Duration(s) * 1e6)
}


func Usleep (s uint) {
//
  time.Sleep (time.Duration(s) * 1e3)
}


func actualizeTimeDate () (h, m, s, d, mo, y uint) {
//
  T:= time.Now()
  hh, mm, ss:= T.Clock()
  h, m, s = uint(hh), uint(mm), uint(ss)
  d = uint(T.Day())
  mo = uint(T.Month())
  y = uint (T.Year())
  return
}


func ActualizeTime () (Hour, Min, Sec uint) {
//
  Hour, Min, Sec, _, _, _ = actualizeTimeDate ()
  return
}


func ActualizeDate () (Day, Month, Year uint) {
//
  _, _, _, Day, Month, Year = actualizeTimeDate ()
  return
}


func SecondsSinceUnix () (s uint, us uint64) {
//
  t:= time.Now()
  s, us = uint(t.Unix()), uint64(t.UnixNano())
  return
}
