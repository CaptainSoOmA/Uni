package scr

// (c) Christian Maurer   v. 120909 - license see murus.go

import (
  "murus/obj"; "murus/xker"
)
var
  clz = int(obj.Codelen (uint(0)))


func codelen (w, h uint) uint {
//
  n:= 4 * uint (clz)
  if underX {
    n += xker.Codelen (w, h)
  } else {
    n += w * h * colourdepth
  }
  return n
}


func enc (x, y, w, h uint) []byte {
//
  e:= make ([]byte, 4 * clz)
  copy (e[     :1*clz], obj.Encode (x))
  copy (e[1*clz:2*clz], obj.Encode (y))
  copy (e[2*clz:3*clz], obj.Encode (w))
  copy (e[3*clz:4*clz], obj.Encode (h))
  return e
}


func dec (b []byte) (x, y, w, h uint) {
//
  x = obj.Decode (uint(0), b[     :  clz]).(uint)
  y = obj.Decode (uint(0), b[1*clz:2*clz]).(uint)
  w = obj.Decode (uint(0), b[2*clz:3*clz]).(uint)
  h = obj.Decode (uint(0), b[3*clz:4*clz]).(uint)
  return
}


func encode (x, y, w, h uint) []byte {
//
  if w == 0 || h == 0 || x + w > nX [mode] || y + h > nY [mode] {
    return []byte(nil)
  }
  e:= make ([]byte, Codelen (w, h))
  j:= 4 * uint(clz)
  copy (e[0:j], enc (x, y, w, h))
  if underX {
    xker.Encode (x, y, w, h, e[j:])
  } else {
    i:= (XX * y + x) * colourdepth
    di:= XX * colourdepth
    dj:= w * colourdepth
    for k:= y; k < y + h; k++ {
      copy (e[j:j+dj], fbcop[i:i+dj])
      i += di
      j += dj
    }
  }
  return e
}


func decode (b []byte) {
//
  if b == nil { return }
  j:= uint(4 * clz)
  x, y, w, h:= dec (b[:j])
  if underX {
    xker.Decode (x, y, w, h, b[j:])
    return
  }
  if ! visible { return }
  i:= (XX * y + x) * colourdepth
  di:= XX * colourdepth
  dj:= w * colourdepth
  for k:= uint(0); k < h; k++ {
    copy (fbmem[i:i+dj], b[j:j+dj])
    copy (fbcop[i:i+dj], b[j:j+dj])
    i += di
    j += dj
  }
}
