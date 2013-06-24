package kbd

// (c) Christian Maurer   v. 130312 - license see murus.go

// >>> implements only a german keyboard !

import (
  "os"
  "murus/ker"; "murus/spc"; "murus/xker"
  "murus/z"; "murus/mouse"; "murus/navi"
)
const (
  pack = "kbd"
  PIPE_BUF = 256
// raw key codes
// alphanumeric keyboard: control keys
  shiftL = 42; shiftR = 54; shiftLock = 58
  ctrlL = 29; ctrlR = 97
  altL = 56; altR = 100
  escape = 1; backspace = 14; tab = 15;
  enter = 28;
// alphanumeric keyboard:
  f1 = 59; f2 = 60; f3 = 61; f4 = 62; f5 = 63; f6 = 64; f7 = 65; f8 = 66; f9 = 67; f10 = 68
  f11 = 87; f12 = 88
// numeric keypad:
  numEnter = 96
  num0 = 82; num1 = 79; num2 = 80; num3 = 81; num4 = 75; num5 = 76; num6 = 77; num7 = 71; num8 = 72; num9 = 73; numSep = 83
  numMinus = 74; numPlus = 78; numTimes = 55; numDiv = 98
  left = 105; right = 106; up = 103; down = 108
// ? = 101; ? = 112
// ? = 117; ? = 118
// ? = 120; ? = 121; ? = 122; ? = 123; ? = 124
// special keypad:
  pageUp = 104; pageDown = 109; pos1 = 102; end = 107;
  insert = 110; delete = 111
// ? = 89; ? = 90; ? = 91; ? = 92; ? = 93; ? = 94; ? = 95
  print_ = 99; roll = 70; pause = 119
  numOnOff = 69;
  onOff = 113; lower = 114; louder = 115
  doofL = 125; doofM = 126; doofR = 127
  noKeycodes = 128
//  toolbox   = 501 // 501 % 256 = 245, % 128 = 117
  pageRight = 158 // only under X
  pageLeft  = 159 // only under X
  off       = 128
// combinations:
  shiftLoff = shiftL + off
  shiftRoff = shiftR + off
  shiftLockoff = shiftLock + off
  ctrlLoff  = ctrlL + off
  ctrlRoff  = ctrlR + off
  altLoff   = altL + off
  altRoff   = altR + off
  doofLoff  = doofL + off // d(o,o)f
  doofMoff  = doofM + off
  doofRoff  = doofR + off
  function  = 143
// combinations with Fn:
// ? = 466 // Fn F1
  lockComputer = 152 // Fn F2/F3
// TODO translations
  akkuUndStromverbrauchVerwalten = 236 // Fn F3
  bereitschaftsmodusAktivieren = 142 // Fn F4
  drahtloseVerbindungenVerwalten = 238 // Fn F5
// ? = 471 // Fn F6
  anzeigeeinstellungenAendern = 227 // Fn F7
  einstellungenVonEingabeeinheitenAendern = 192 // Fn F8
  externeEinheitenVerwalten = 194 // Fn F9 // große Tastatur: 502 -> 246
// ? = 476 // Fn F11
  hibernationsmodusAktivieren = 205 // Fn F12
  lighter = 225 // Fn Pos1
  darker = 224 // Fn Ende
  toggleThinkLight = 228
  fnNum = 140 // Fn Num
  anzeigeninhaltVergroessern = 372 // 372 % 256 = 116 // Fn Space
)
var (
  bb,       // key
  bB,       // key + Shift
  aa []byte // key + AltGr
//  aA []byte // key + Shift + AltGr
  kK [noKeycodes]Comm
  keypipe chan byte
  mousepipe chan mouse.Command
  navipipe chan navi.Command
  lastbyte byte
  lastcommand Comm
  lastdepth uint
  shift, shiftFix, ctrl, alt, altGr /* , numOnOff */, fn, lBut, mBut, rBut bool
)


func isAlpha (n uint) bool {
//
  switch n { case 41, 2, 3, 4, 5, 6, 7, 8, 9,10,11,12,13, // ^ 1 2 3 4 5 6 7 8 9 0 ß '
                   16,17,18,19,20,21,22,23,24,25,26,27,   //  Q W E R T Z U I O P Ü +
                    30,31,32,33,34,35,36,37,38,39,40,43,  //   A S D F G H J K L Ö Ä #
                   86,44,45,46,47,48,49,50,51,52,53,      //  < Y X C V B N M , . -
                                  57,                     // space
                  numMinus, numPlus, numTimes, numDiv:    // keypad
    return true
  }
  return false
}


func isF (n uint) bool {
//
  switch n { case f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12:
    return true
  }
  return false
}


func isCmd (n uint) bool {
//
  switch n { case backspace, tab, enter,
                  left, right, up, down, pageUp, pageDown, pos1, end,
                  insert, delete,
                  print_, roll, pause,
                  onOff, lower, louder,
                  numEnter:
    return true
  }
  return isF (n)
}


func isFunction (n uint) bool {
//
  switch n { case lockComputer, akkuUndStromverbrauchVerwalten,
                  bereitschaftsmodusAktivieren, drahtloseVerbindungenVerwalten,
                  anzeigeeinstellungenAendern, einstellungenVonEingabeeinheitenAendern,
                  externeEinheitenVerwalten, hibernationsmodusAktivieren,
                  lighter, darker, toggleThinkLight,
                  fnNum, anzeigeninhaltVergroessern:
    return true
  }
  return false
}


func isKeypad (n uint) bool {
//
  switch n { case num0, num1, num2, num3, num4, num5, num6, num7, num8, num9, numSep:
    return true
  }
  return false
}


func catch () {
//
  shift, ctrl, alt, altGr, fn = false, false, false, false, false
  var b byte
  defer ker.Terminate () // Hilft nix. Warum nicht ???
  for {
    ker.ReadTerminal (&b)
    switch b { // case 0:
      // ker.Stop (pack, 1) // Fn-key combination !
    case shiftL, shiftR, shiftLock:
      shift = true
    case ctrlL, doofL, ctrlR:
      ctrl = true
    case altL, doofM:
      alt = true
    case altR, doofR:
      altGr = true
    case shiftLoff, shiftRoff, shiftLockoff:
      shift = false
    case ctrlLoff, doofLoff, ctrlRoff:
      ctrl = false
    case altLoff, doofMoff:
      alt = false
    case altRoff, doofRoff:
      altGr = false
    case function:
      // println ("Fn-Key")
      fn = true
    default:
      if ctrl && // (alt || altGr) && b == pause ||
                 b == 46 { // 'C'
        ker.Terminate()
        os.Exit (1)
      } else if b < off && ctrl && (alt || altGr) {
        switch b { case left, right:
          ker.Console1 (b == right)
        case f1, f2, f3, f4, f5, f6, f7, f8, f9, f10:
          ker.Console (b - f1 + 1)
        case f11, f12:
          ker.Console (b - f11 + 11)
        case escape, backspace, tab, enter, roll, numEnter, pos1, up, pageUp, end, down, pageDown, insert, delete:
          keypipe <- b
        }
      } else {
        keypipe <- b
      }
    }
  }
}


func input (b *byte, c *Comm, d *uint) {
//
  var (
    b0 byte
    k, k1 uint
    mc mouse.Command
    m3c navi.Command
    ok bool
  )
  loop: for {
    *b, *c, *d = 0, None, 0
    select {
    case mc = <-mousepipe:
      *c, *d = Go + Comm (mc), 0
      if shift || ctrl { *d ++ }
      if alt || altGr { *d += 2 }
      break loop
    case m3c = <-navipipe:
      *c, *d = Go + Comm (m3c), 0
      if shift || ctrl { *d ++ }
      if alt || altGr { *d += 2 }
      break loop
    case b0, ok = <-keypipe:
      if ok {
        k = uint(b0)
      } else {
        ker.Stop (pack, 2)
      }
    }
//    if k == 0 { ker.Stop (pack, 3) }
    k1 = k
    k = k % off
    if shift || ctrl { *d ++ }
    if alt || altGr { *d += 2 }
    switch b0 { case pageUp, pageDown: *d += 2 }
    switch {
    case isAlpha (k):
      switch *d { case 0:
        *b = bb [k]
      case 1:
        *b = bB [k]
      case 2:
        *b = aa [k]
      default:
//        if altGr {
//          *b = aA [k]
//        } else {
        *b = aa [k]
//        }
      }
    case k == escape || k == numEnter || isCmd (k):
      *c = kK [k]
    case k == shiftLock:
      shift = true
    case isKeypad (k):
      if shift {
        *c = kK [k]
        switch k { case num9, num3: *d = 2 }
      } else {
        *b = bb [k]
      }
    case k == function:
      // println ("Fn-Key")
    case isFunction (k):
      // TODO
    default:
      switch k {
      case 0:
        ; // ignore
      case anzeigeninhaltVergroessern % 128, // == 116
           lockComputer % 128,
           akkuUndStromverbrauchVerwalten % 128,
           bereitschaftsmodusAktivieren % 128,
           drahtloseVerbindungenVerwalten % 128,
           anzeigeeinstellungenAendern % 128,
           einstellungenVonEingabeeinheitenAendern % 128,
           externeEinheitenVerwalten % 128,
           hibernationsmodusAktivieren % 128,
           lighter % 128,
           darker % 128,
           toggleThinkLight % 128,
           fnNum % 128:
        ; // nixtun
      default:
        ker.Stop (pack, 1000 + k)
      }
    }
    if k1 < off { // key pressed, not released
      if *b == 0 {
        if *c > None {
          break loop
        }
      } else {
        lastbyte = *b
        *c = None
        break loop
      }
    }
  }
  lastcommand = *c
  lastdepth = *d
}


func read () (byte, Comm, uint) {
//
  var (b byte; c Comm; d uint)
  if underX {
    inputX (&b, &c, &d)
  } else {
    input (&b, &c, &d)
  }
  return b, c, d
}


func mouseEx () bool {
//
  if underX {
    return true
  }
  return mouse.Ex ()
}


func byte_ () byte {
//
  b:= byte(0)
  for {
    b, _, _ = Read ()
    if b != 0 {
      break
    }
  }
  return b
}


func command () (Comm, uint) {
//
  var ( c Comm; d uint )
  for {
    _, c, d = Read ()
//    if b == 0 { break }
    if c != None { break }
  }
  return c, d
}


func readNavi () (spc.GridCoord, spc.GridCoord) {
//
  return navi.Read ()
}


func lastByte () byte {
//
  return lastbyte
}


func lastCommand (d *uint) Comm {
//
  *d = lastdepth
  return lastcommand
}


func depositCommand (c Comm) {
//
  lastcommand = c
}


func depositByte (b byte) {
//
  lastbyte = b
}


func wait (b bool) {
//
  c0, d0:= lastcommand, lastdepth
  var c Comm
  for {
    c, _ = Command ()
    if b {
      if c == Enter /* || c == Here */ { break }
    } else {
      if c == Esc || c == Back /* || c == There */ { break }
    }
  }
  lastcommand, lastdepth = c0, d0
}


func confirmed (w bool) bool {
//
  c0, d0:= lastcommand, lastdepth
  var ( c Comm; d, dmin uint )
  if w {
    dmin = 1
  } else {
    dmin = 0
  }
  var b bool
  for {
    c, d = Command ()
    if c == Enter {
      if d >= dmin {
        b = true
        break
      }
    } else if c == Esc {
      if d >= dmin {
        b = false
        break
      }
    }
  }
  lastcommand, lastdepth = c0, d0
  return b
}


var
  text [NComms]string


func string_ (c Comm) string {
//
  if c < NComms {
    return text [c]
  }
  return "häh ???"
}

/*
func Control (n uint, i *uint) {
//
  var d uint
  switch LastCommand (&d) {
  case Esc:
    break loop
  case Enter:
    if d == 0 {
      if *i + 1 < n {
        *i ++
      } else {
        break loop
      }
    } else {
      break loop
    }
  case Down:
    if *i + 1 < n {
      *i ++
    } else {
      *i = 0
    }
  case Up:
    if *i > 0 {
      *i --
    } else {
      *i = n - 1
    }
  case Pos1:
    *i = 0
  case End:
    *i = n - 1
  }
}
*/

func init () {
//
  //           0         1         2         3         4         5         6         7         8         9
  //           012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678
  bb = []byte("  1234567890 '  qwertzuiop +  asdfghjkl  ^ #yxcvbnm,.- +               789-456+1230,  <           /")
  //                       ß             ü            öä
  bb[12] = z.Sz
  bb[26] = z.Ue
  bb[39] = z.Oe
  bb[40] = z.Ae

  //           0         1         2         3         4         5         6         7         8         9
  //           012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678
  bB = []byte("  !  $%&/()=?`  QWERTZUIOP *  ASDFGHJKL    'YXCVBNM;:_ *               789-456+1230,  >           /")
  //               §                     Ü            ÖÄ°
  bB [4] = z.Para
  bB[26] = z.UE
  bB[39] = z.OE
  bB[40] = z.AE
  bB[41] = z.Degree

  //           0         1         2         3         4         5         6         7         8         9
  //           012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678
  aa = []byte("  !  $%&/()=?`  @WERTZUIOP ~  ASDFGHJKL    'YXCVBNM;:_ ~               {[]-456+123},  |           /")
  //              ²³             ¤       ü            öä     ¢   µ
  aa [3] = z.ToThe2
  aa [4] = z.ToThe3
  aa[18] = z.Euro
  aa[26] = z.Ue
  aa[39] = z.Oe
  aa[40] = z.Ae
  aa[46] = z.Copyright
  aa[50] = z.Mue

//  //           0         1         2         3         4         5         6         7         8         9
//  //           012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678
//  aA = []byte("  !  $%&/()=?`  @WERTZUIOP ~  ASDFGHJKL    'YXCVBNM;:_ ~               {[]-456+123},  |           /")
//  //                              ®             ª       ¬    ¢   º ÷                    ± ¤     £ ×
//  aA[19] = z.Registered
//  aA[33] = z.Female
//  aA[41] = z.Negate
//  aA[46] = z.Copyright
//  aA[50] = z.Male
//  aA[52] = z.Division
//  aA[73] = z.PlusMinus
//  aA[75] = z.Euro
//  aA[81] = z.Pound
//  aA[83] = z.Times

  for b:= 0; b < noKeycodes; b++ { kK [b] = Esc }
  kK [escape] = Esc
  kK [f1] = Help
  kK [f2] = LookFor
  kK [f3] = Act
  kK [f4] = Cfg
  kK [f5] = Mark
  kK [f6] = Demark
  kK [f7] = Paste
  kK [f8] = Deposit
  kK [f9] = Black
  kK [f10] = Red
  kK [f11] = Green
  kK [f12] = Blue
  kK [backspace] = Back
  kK [tab] = Tab
  kK [enter] = Enter
  kK [print_] = PrintScr
  kK [pos1] = Pos1
  kK [up] = Up
  kK [pageUp] = Up
  kK [left] = Left
  kK [right] = Right
  kK [end] = End
  kK [down] = Down
  kK [pageDown] = Down
  kK [insert] = Ins
  kK [delete] = Del
  kK [roll] = Roll
  kK [num7] = kK [pos1]
  kK [num8] = kK [up]
  kK [num9] = kK [pageUp]
  kK [num4] = kK [left]
  kK [num6] = kK [right]
  kK [num7] = kK [end]
  kK [num8] = kK [down]
  kK [num9] = kK [pageDown]
  kK [num0] = kK [insert]
  kK [numSep] = kK [delete]
  kK [numEnter] = kK [enter]
  kK [pause] = Pause
  kK [onOff] = OnOff
  kK [lower] = Lower
  kK [louder] = Louder

  text = [NComms]string {
    "       ", "Esc    ", "Enter  ", "<==    ",
    "<-     ", "->     ", "^      ", "_      ",
    "Pos1   ", "Ende   ",
    "Tab    ", "Entf   ", "Einfg  ",
    "F1     ", "F2     ", "F3     ", "F4     ", "F5     ", "F6     ", "F7     ", "F8     ",
    "F9     ", "F10    ", "F11    ", "F12    ",
    "drucke ", "rolle  ", "Pause  ",
    "an/aus ", "leiser ", "lauter ",
    "laufe  ",
    "hier   ", "ziehe  ", "hierhin",
    "dort   ", "schiebe", "dorthin",
    "da     ", "bewege ", "dahin  ",
    "Navigat" }

  lastbyte, lastcommand, lastdepth = 0, None, 0
  underX = xker.Active ()
  if underX {
    xpipe = make (chan xker.Event)
    go catchX ()
  } else {
    ker.InitTerminal ()
    ker.InstallTerm (func () { ker.TerminateTerminal() } )
    keypipe = make (chan byte, PIPE_BUF)
    if mouse.Ex () {
      mousepipe = mouse.Channel ()
    } else {
      mousepipe = nil
    }
    navipipe = navi.Channel ()
    go catch ()
  }
}
