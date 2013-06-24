package scr

// (c) Christian Maurer   v. 130312 - license see murus.go

import (
  "syscall"
  "murus/ker"; "murus/col"; "murus/font"
)
var (
  visible bool // only for console switching
  consoleShape Shape
)


func consoleOn () {
//
  ker.ActivateConsole ()
  n:= nY [mode] * nX [maxMode] * uint(colourdepth)
  copy (fbmem[:n], fbcop[:n])
  visible = true
  warp (actualCharheight * blinkY, actualCharwidth * blinkX, consoleShape)
}


func consoleOff () {
//
  visible = false
  consoleShape = blinkShape
  warp (actualCharheight * blinkY, actualCharwidth * blinkX, Off)
  ker.DeactivateConsole ()
}


func terminateConsole () {
//
// TODO wait (blink())
// TODO terminate (blink())
  terminated = true
  ker.Msleep (250) // provisorial
  cursorShape = Off
  print (esc1 + "H" + esc1 + "J")
  print (esc1 + "?25h" + esc1 + "?0c")
}


func initConsole () {
//
  colbits:= uint(0)
  XX, YY, colbits, fbmem = ker.Framebuffer ()
  if fbmem == nil {
//    terminateConsole ()
    panic ("framebuffer was not initialized! (is /dev/fb0 crw-rw-rw ?)") // Speicherzugriffsfehler
  }
  maxMode = Mode(0)
  for {
    if XX == nX[maxMode] && YY == nY[maxMode] {
      mode = maxMode
      break
    } else if maxMode + 1 < NModes {
      maxMode ++
    } else {
      mode = defaultMode
      XX, YY = nX[mode], nY[mode]
      break
    }
  }
  switch colbits { case 4, 8:
    colourdepth = 1
  case 15, 16:
    colourdepth = 2
  case 24:
    colourdepth = 3
  case 32:
    colourdepth = 4
  default:
    ker.Stop (pack, 100 + colbits)
  }
  col.SetColourDepth (colbits)
  fbmemsize:= nY [maxMode] * nX [maxMode] * uint (colourdepth)
  if uint(len (fbmem)) != fbmemsize {
    terminateConsole ()
    println ("len(fbmem) == ", len(fbmem), " != fbmemsize == ", fbmemsize)
    panic ("len(fbmem) not ok !")
  }
  swFontsize (font.Normal)
  fbcop = make ([]byte, fbmemsize)
  archive = make ([]byte, fbmemsize)
//  for n:= 0; n < int(len(fbmem)) - int(colourdepth); n += int(colourdepth) {
//    copy (archive[n:n+int(colourdepth)], cc (col.EncodeF))
//  }
  emptyBackground = make ([]byte, fbmemsize)
  print (esc1 + "2J" + esc1 + "?1c" + esc1 + "?25l")
  Colours (col.ScreenF, col.ScreenB)
  initMouse ()
  visible = true
  go blink ()
  ker.InitConsole ()
  ker.SetAction (syscall.SIGUSR1, consoleOff)
  ker.SetAction (syscall.SIGUSR2, consoleOn)
  ker.InstallTerm (terminateConsole)
  go ker.CatchSignals ()
  initConsoleFonts()
}
