package scr

// (c) Christian Maurer   v. 120909 - license see murus.go

import (
  "runtime"; "sync"
  "murus/ker"
)
var (
  cursorShape,
  blinkShape Shape
  blinkX, blinkY uint
  blinking sync.Mutex
  terminated bool
)


func blink () {
//
  var
    shape Shape
  for {
    blinking.Lock()
    if cursorShape == Off {
      shape = blinkShape
    } else {
      shape = Off
    }
    switchCursor (blinkX, blinkY, shape)
    blinking.Unlock()
    if terminated {
      break
    }
    ker.Msleep (250)
  }
  runtime.Goexit ()
}


func switchCursor (x, y uint, s Shape) {
//
  const s0 = 2
  var y0, y1 uint
  if cursorShape == s { return }
  switch s { case Off:
    switch cursorShape { case Understroke:
      y0 = actualCharheight - s0
      y1 = actualCharheight - 1
    case Block:
      y0 = 0
      y1 = actualCharheight - 1
    }
  case Understroke:
    switch cursorShape { case Off:
      y0 = actualCharheight - s0
      y1 = actualCharheight - 1
    case Block:
      y0 = 0
      y1 = actualCharheight - s0 - 1
    }
  case Block:
    switch cursorShape { case Off:
      y0 = 0
      y1 = actualCharheight - 1
    case Understroke:
      y0 = 0
      y1 = actualCharheight - s0 - 1
    }
  }
  cursorShape = s
//  Lock() // weg ?
  InvertGr (x, y + y0, x + actualCharwidth - 1, y + y1)
//  Unlock() // weg ?
}


func warp  (L, C uint, s Shape) {
//
  warpGr (actualCharwidth * C, actualCharheight * L, s)
}


func warpGr (x, y uint, s Shape) {
//
  if x >= nX [mode] || y >= nY [mode] { return }
  blinking.Lock()
  blinkX = x
  blinkY = y
  blinkShape = s
  switchCursor (x, y, blinkShape)
  blinking.Unlock()
}
