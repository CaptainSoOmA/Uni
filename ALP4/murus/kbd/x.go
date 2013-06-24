package kbd

// (c) Christian Maurer   v. 130312 - license see murus.go

import (
  "os"
  "murus/ker"; "murus/xker"
)
var (
  underX bool
  xpipe chan xker.Event
)


// Pre: xker.x.initialized == true
func catchX () {
//
  for xker.Eventpipe == nil {
    ker.Msleep (10)
  }
//  println ("keyboard.catchX: Eventpipe != nil")
  for p:= range xker.Eventpipe {
    xpipe <- p
  }
  close (xpipe)
}


func isSet (bit, x uint) bool {
//
  return x >> bit % 2 == 1
}


func inputX (B *byte, C *Comm, T *uint) {
//
  const (
    shiftBit     = 0
    shiftLockBit = 1
    ctrlBit      = 2
    altBit       = 3
    altGrBit     = 7
    mouseBitL    = 8
    mouseBitM    = 9
    mouseBitR   = 10
  )
  var e xker.Event
  var k uint
  ok:= false
  loop: for {
    *B, *C, *T = 0, None, 0
    e, ok = <-xpipe
    if ! ok {
      println ("xker.inputX: ! ok")
    }
    shift = isSet (shiftBit, e.S)
    shiftFix = isSet (shiftLockBit, e.S)
    ctrl = isSet (ctrlBit, e.S)
    alt = isSet (altBit, e.S)
    altGr = isSet (altGrBit, e.S)
    fn = false
    lBut = isSet (mouseBitL, e.S)
    mBut = isSet (mouseBitM, e.S)
    rBut = isSet (mouseBitR, e.S)
    if shift || ctrl {
      *T ++
    }
    if alt {
      *T += 2
    }
    switch e.T {
    case xker.KeyPress:
      if e.C < 9 {
        println ("oops, got keycode ", e.C, " < 9") // TODO ?
      } else {
        k = e.C - 8
        switch {
        case k == escape:
          *C = Esc
        case k == shiftL || k == shiftR:
          shift = true
        case k == ctrlL || k ==  ctrlR:
          ctrl = true
        case k == altL:
          alt = true
        case k == altR:
          altGr = true
        case isAlpha (k):
          if ctrl && (k == 46 || k == 16 ) { // Ctrl C, Ctrl Q
            /* terminateX (); */ os.Exit (0)
          }
          switch *T { case 0:
            *B = bb[k]
          case 1:
            *B = bB[k]
          case 2:
            *B = aa[k]
          }
        case isCmd (k):
          *C = kK[k]
//          if k == pageUp || e.C == pageDown {
          if k == pageUp || k == pageDown {
//          if e.C == pageUp + 8 || e.C == pageDown + 8 { // 112/117 -> 104/109
            *T += 2
          }
          if (e.C == left || e.C == right) && e.S == 64 { *T += 2 }
          if e.C == backspace && *T > 2 { *C = None; *T = 0 } // doesn't help: wm crashes
        case k == shiftLock:
          shift = ! isSet (shiftLockBit, e.S) // weg isser
        case k == numOnOff:
          ; // TODO
        case isKeypad (k):
          switch *T { case 0:
            *B = bb[k]
          default:
            *C = kK[k]
          }
        case isFunction (k):
          ; // TODO
        default:
          println ("not yet handled: keycode ", k, "/ state ", e.S)
        }
      }
      if *B > 0 || *C > 0 {
        break loop
      }
    case xker.KeyRelease:
      ;
    case xker.ButtonPress:
      if *T > 1 { *T = 1 } // because the bloody WM eats everything else up
      switch e.C { case 1:
//        lBut = true
        *C = Here
      case 2:
//        mBut = true
        *C = This
      case 3:
//        rBut = true
        *C = There
      case 4:
        *C = Up
      case 5:
        *C = Down
      case 6:
        println ("xker.ButtonPress: button ", e.C, "/ state ", e.S)
      case 7:
        println ("xker.ButtonPress: button ", e.C, "/ state ", e.S)
      default:
        println ("xker.ButtonPress not yet handled: button ", e.C ,"/ state ", e.S)
      }
      if *C > 0 {
        break loop
      }
    case xker.ButtonRelease:
      if *T > 1 { *T = 1 } // because the bloody WM eats everything else up
      ctrl = false
      alt = false
      altGr = false
      switch e.C { case 1:
        if lBut {
          lBut = false
          *C = Hither
        }
      case 2:
        if mBut {
          mBut = false
          *C = Thus
        }
      case 3:
        if rBut {
          rBut = false
          *C = Thither
        }
      case 4:
        *C = Up
      case 5:
        *C = Down
      case 6:
        println ("xker.ButtonRelease: button ", e.C, "/ state ", e.S)
      case 7:
        println ("xker.ButtonRelease: button ", e.C, "/ state ", e.S)
      default:
        println ("xker.ButtonRelease not yet handled: button ", e.C ,"/ state ", e.S)
      }
      if *C > 0 {
        break loop
      }
    case xker.MotionNotify:
      *T = 0
      if lBut {
        *C = Pull
      } else if mBut {
        *C = Move
      } else if rBut {
        *C = Push
      } else {
        *C = Go
      }
      break loop
    case xker.ClientMessage:
      ; // break loop // navi
    default:
      *B, *C, *T = 0, None, 0
      break loop
    }
  }
  lastbyte, lastcommand, lastdepth = *B, *C, *T
}
