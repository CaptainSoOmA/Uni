package img

// (c) Christian Maurer   v. 121011 - license see murus.go

import (
  "os/exec"
  "murus/ker"; "murus/str"; "murus/scr"; "murus/prt"
  "murus/errh"
  "murus/pseq"
)
const
  suffix = ".ppm"
var
  toPrint bool


func put (n string, x, y, w, h uint) {
//
  if str.Empty (n) { return }
  str.RemSpaces (&n)
  filename:= n + suffix
  if scr.UnderX() { errh.Hint ("bitte etwas Geduld ...") }
  buf:= scr.P6Encode (x, y, w, h)
  if scr.UnderX() { errh.DelHint () }
  file:= pseq.New (buf)
  file.Name (filename)
  file.Clr ()
  file.Put (buf)
  file.Terminate ()
  if ! toPrint {
    exec.Command ("pnmtopng", filename + suffix, ">", n, ".png").Run()
    ker.Msleep (100)
    exec.Command ("rm", filename)
  }
}


func size_(n string) (uint, uint) {
//
  w, h:= uint(0), uint(0)
  if str.Empty (n) {
    return w, h
  }
  str.RemSpaces (&n)
  filename:= n + suffix
  l:= pseq.Length (filename)
  if l == 0 {
    return w, h
  }
  buf:= make ([]byte, l)
  file:= pseq.New (buf)
  file.Name (filename)
  buf= file.Get ().([]byte)
  file.Terminate ()
  w, h = scr.P6Size (buf)
  return w, h
}


func get (n string, x, y uint) {
//
  const tst = true
  if str.Empty (n) { return }
  str.RemSpaces (&n)
  filename:= n + suffix
  l:= pseq.Length (filename)
  if l == 0 { return }
  buf:= make ([]byte, l)
  file:= pseq.New (buf)
  file.Name (filename)
  buf = file.Get ().([]byte)
  file.Terminate ()
  scr.P6Decode (x, y, buf)
}


func print_ (x, y, w, h uint) {
//
  toPrint = true
  filename:= "tmp"
  Put (filename, x, y, w, h)
  toPrint = false
  exec.Command (prt.PrintCommand, "-o", "landscape", "-o", "scaling=99", filename + suffix).Run()
}


func print1 () {
//
  print_ (0, 0, scr.NX(), scr.NY())
}
