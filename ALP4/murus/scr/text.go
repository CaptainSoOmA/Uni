package scr

// (c) Christian Maurer   v. 130303 - license see murus.go

import (
  "murus/ker"; "murus/xker"; "murus/z"
  "murus/col"
)
var
  transparent bool


func write (S string, L, C uint) {
//
  n:= len (S)
  if C + uint(n) > nColumns {
    n = int(nColumns - C)
  }
  if n == 0 { return }
  z.ToHellWithUTF8 (&S)
  n = len (S)
  if underX {
    xker.Write (S, int(C * actualCharwidth), int(L * actualCharheight), transparent)
  } else {
    if ! visible { return }
    for s:= 0; s < n; s++ {
      write1 (S [s], L, C + uint(s))
    }
  }
}


func write1 (B byte, L, C uint) {
//
  if L >= nLines || C >= nColumns {
    return
  }
//  if B < 32 || B > 128 { return }
  if underX {
    Write (string (B), L, C)
  } else {
    if ! visible { return }
    cf:= col.CodeF
    lb:= actualLinewidth
    actualLinewidth = Thin
    for Z:= uint(0); Z < actualCharheight; Z++ {
      for S:= uint(0); S < actualCharwidth; S++ {
        if pointed (actualFontsize, B, Z, S) {
          col.CodeF = cf
        } else {
          col.CodeF = col.CodeB
        }
        point (int(actualCharwidth * C + S), int(actualCharheight * L + Z))
      }
    }
    actualLinewidth = lb
    col.CodeF = cf
  }
}


func writeNat (n, Z, S uint) {
//
  t:= "00"
  if n > 0 {
    const M = 10
    B:= make ([]byte, M)
    for i:= 0; i < M; i++ {
      B[M - 1 - i] = byte('0' + n % 10)
      n = n / M
    }
    s:= 0
    for s < M && B[s] == '0' {
      s++
    }
    t = ""
    if s == M - 1 { s = M - 2 }
    for i:= s; i < M - int(n); i++ {
      t += string(B[i])
    }
  }
  write (t, Z, S)
}


func write1Gr (s byte, x, y int) {
//
  if underX {
    t:= string (s)
    z.ToHellWithUTF8 (&t)
    xker.Write (t, x, y, transparent)
    return
  }
  if ! visible ||
     x < 0 || x >= int (nX [mode] - actualCharwidth) ||
     y < 0 || y >= int (nY [mode] - actualCharheight) {
    return
  }
  CF:= col.CodeF
  LB:= actualLinewidth
  actualLinewidth = Thin
  for Y:= uint(0); Y < actualCharheight; Y++ {
    for X:= uint(0); X < actualCharwidth; X++ {
      if pointed (actualFontsize, s, Y, X) {
        col.CodeF = CF
        point (x + int(X), y + int(Y))
      } else if ! transparent {
        col.CodeF = col.CodeB
        point (x + int(X), y + int(Y))
      }
    }
  }
  col.CodeF = CF
  actualLinewidth = LB
}


func writeGr (s string, x, y int) {
//
  z.ToHellWithUTF8 (&s)
  n:= len (s)
  if n == 0 { return }
  if underX {
    xker.Write (s, x, y, transparent)
    return
  }
  if ! visible || x < 0 || y < 0 { return }
  if n == 0 { ker.Stop (pack, 3) }
  for i:= 0; i < n; i++ {
    write1Gr (s[i], x + i * int(actualCharwidth), y)
  }
}


func write1InvGr (Z byte, x, y int) {
//
  if underX {
    t:= string (Z)
    z.ToHellWithUTF8 (&t)
    xker.WriteInvert (t, x, y, transparent)
    return
  }
  if ! visible || x < 0 || x >= int(nX [mode] - actualCharwidth) ||
                           y < 0 || y >= int(nY [mode] - actualCharheight) {
    return
  }
  for Y:= uint(0); Y < actualCharheight; Y++ {
    for X:= uint(0); X < actualCharwidth; X++ {
      if pointed (actualFontsize, Z, Y, X) {
        pointInv (x + int(X), y + int(Y))
      } else if ! transparent {
        pointInv (x + int(X), y + int(Y))
      }
    }
  }
}


func writeInvGr (s string, x, y int) {
//
  z.ToHellWithUTF8 (&s)
  n:= len (s)
  if n == 0 {
    return
  }
  if underX {
    xker.WriteInvert (s, x, y, transparent)
    return
  }
  if ! visible || x < 0 || y < 0 {
    return
  }
  for i:= 0; i < n; i++ {
    write1InvGr (s[i], x + i * int(actualCharwidth), y)
  }
}
