package scr

// (c) Christian Maurer   v. 130102 - license see murus.go

import (
  "murus/xker"; "murus/col"
)
const
  t6 = 64
var
  p6t []byte = []byte("P6\n# murus   (c) 1986-2013   Dr. Christian Maurer\n____ ____\n255\n")


func p6Codelen (w, h uint) uint {
//
  return t6 + w * h * col.P6
}


func p6txt (n, k uint) {
//
  for i:= 0; i < 4; i++ {
    p6t[int(k) + 3 - i] = '0' + byte(n % 10)
    n = n / 10
  }
}


func p6number (b []byte) (uint, int) {
//
  i:= 0
  for '0' <= b[i] && b[i] <= '9' {
    i ++
  }
  n:= uint(0)
  for j:= 0; j < i; j++ {
    n = 10 * n + uint(b[j] - '0')
  }
  return n, i
}


func p6dec (b []byte) (uint, uint, uint, uint) {
//
  w, h, fix:= uint(0), uint(0), uint(0)
  if ! visible { return w, h, fix, 0 }
  p6:= string (b[0:2])
  if p6 != "P6" { return w, h, fix, 0 }
  i, di:= 3, 0
  if b[i] == '#' { // ignore comment
    for {
      i ++
      if b[i] < ' ' { break }
    }
  }
  i ++ // ignore LF
  w, di = p6number (b[i:])
  i += 1 + di // ignore LF or space
  h, di = p6number (b[i:])
  i += 1 + di
  fix, di = p6number (b[i:])
  i += 1 + di
  return w, h, fix, uint(i)
}


func p6Size (b []byte) (uint, uint) {
//
  w, h, fix, _:= p6dec (b)
  if fix != 255 {
    w, h = uint(0), uint(0)
  }
  return w, h
}


func p6Enc (b []byte) []byte {
//
  i:= 4 * clz
  _, _, w, h:= dec (b[0:i])
  cl:= p6Codelen (w, h)
  p:= make ([]byte, cl)
  p6txt (w, 50)
  p6txt (h, 55)
  j:= t6
  copy (p[0:j], p6t)
  if w == 0 || h == 0 {
    return p
  }
  if underX {
    xker.P6Encode (w, h, b[i:], p[j:])
  } else {
    di, dj:= int(colourdepth), col.P6
    for y:= uint(0); y < h; y++ {
      for x:= uint(0); x < w; x++ {
        col.P6Encode (b[i:i+di], p[j:j+dj])
        i += di
        j += dj
      }
    }
  }
  return p
}


func p6Encode (x, y, w, h uint) []byte {
//
  return p6Enc (encode (x, y, w, h))
}


func p6Dec (x, y uint, p []byte) []byte {
//
  w, h, fix, n:= p6dec (p)
  if w == 0 || h == 0 || fix != 255 {
    return nil
  }
  i:= 4 * clz
  b:= make ([]byte, Codelen (w, h))
  copy (b[:i], enc (x, y, w, h))
  if underX {
    xker.P6Decode (x, y, w, h, p[n:], b[i:])
  } else {
    j:= int(n)
    di, dj:= int(colourdepth), col.P6
    var c col.Colour
    for y:= uint(0); y < h; y++ {
      for x:= uint(0); x < w; x++ {
        col.Decode (&c, p[j:j+dj])
        copy (b[i:i+di], cc (col.Code (c)))
        i += di
        j += dj
      }
    }
  }
  return b
}


func p6Decode (x, y uint, P6 []byte) {
//
  w, h:= p6Size (P6)
  if x + w > NX() || y + h > NY() {
    return
  }
  b:= p6Dec (x, y, P6)
  decode (b)
}
