package box

// (c) Christian Maurer   v. 130302 - license see murus.go

import (
  . "murus/str"; "murus/kbd"
  "murus/col"; "murus/scr"
)
const (
  pack = "box"
  space = ' '
)
type
  Imp struct {
       width,
       start uint
      cF, cB col.Colour
overwritable,
   graphical,
 transparent,
   numerical,
 TRnumerical,
   usesMouse bool
       index uint
     command kbd.Comm
       depth uint
             }
var
  edited bool = true


func New () *Imp {
//
  return &Imp { width: scr.NColumns(), cF: col.ScreenF, cB: col.ScreenB, command: kbd.None }
}


func (B *Imp) Wd (n uint) {
//
  B.width = n
}


func (B *Imp) SetNumerical () {
//
  B.numerical = true
}


func (B *Imp) SetTransparent (t bool) {
//
  B.transparent = t
}


func (B *Imp) UseMouse () {
//
  B.usesMouse = true
}


func (x *Imp) Colours (f, b col.Colour) {
//
  x.cF, x.cB = f, b
}


func (x *Imp) ColoursScreen () {
//
  x.cF, x.cB = col.ScreenF, col.ScreenB
}


func (x *Imp) ColourF (f col.Colour) {
//
  x.cF, x.cB = f, col.ScreenB
}


func (x *Imp) ColourB (b col.Colour) {
//
  x.cF, x.cB = col.ScreenF, b
}


func (B *Imp) Write (s string, l, c uint) {
//
  if l >= scr.NLines () { return }
  if c >= scr.NColumns() { return }
//  Wd (&s, s)
  n, b:= uint(len (s)), B.width
  if c + b > scr.NColumns() { B.width = scr.NColumns() - c }
  if B.width == 0 { B.width = n }
  if B.width > n { B.width = n }
  if B.width < n { Norm (&s, B.width) }
//  Norm (&s, B.width)
  if B.numerical || B.TRnumerical { Move (&s, false) }
  scr.Lock()
  scr.Colours (B.cF, B.cB)
  if B.transparent { scr.SwitchTransparence (true) }
  scr.Write (s, l, c)
  if B.transparent { scr.SwitchTransparence (false) }
  scr.Unlock()
  B.width = b
}


func (B *Imp) WriteGr (s string, x, y int) {
//
  if uint(y) >= scr.NY() { return }
  if uint(x) >= scr.NX() - scr.NX1() { return }
  n, b:= uint(len (s)), B.width
  if B.width == 0 { B.width = n }
  if uint(x) + B.width * scr.NX1() > scr.NX() {
    B.width = (scr.NX() - uint(x)) / scr.NX1()
  }
  if B.width > n { B.width = n }
  if B.width < n { Norm (&s, B.width) }
  if B.numerical || B.TRnumerical { Move (&s, false) }
  scr.Lock()
  scr.Colours (B.cF, B.cB)
  if B.transparent { scr.SwitchTransparence (true) }
  scr.WriteGr (s, x, y)
  if B.transparent { scr.SwitchTransparence (false) }
  scr.Unlock()
  B.width = b
}


func (B *Imp) Clr (L, C uint) {
//
  if B.width == 0 { return }
  scr.Lock()
  scr.Colours (col.ScreenF, col.ScreenB)
  scr.WriteGr (Clr (B.width), int(scr.NX1 () * C), int(scr.NY1() * L))
  scr.Unlock()
}


func (B *Imp) Start (C uint) {
//
  B.start = C
}


func (B *Imp) write (Text string, x, y uint) {
//
  scr.Lock ()
  scr.Colours (B.cF, B.cB)
  if B.transparent { scr.SwitchTransparence (true) }
  y1:= B.width
  if y1 > uint(len (Text)) { y1 = uint(len (Text)) }
  for x1:= B.index; x1 < y1; x1++ {
    if B.graphical {
      scr.Write1Gr (Text[x1], int(x + scr.NX1() * x1), int(y))
    } else {
      scr.Write1 (Text[x1], y / scr.NY1(), x / scr.NX1() + x1)
    }
  }
  if B.transparent { scr.SwitchTransparence (false) }
  scr.Unlock ()
}


func (B *Imp) done (Text *string, x, y uint) bool {
//
  switch B.command { case kbd.Enter, kbd.Esc:
    return true
  case kbd.Left:
    if B.depth == 0 {
      if B.index > 0 {
        B.index--
      }
    } else {
      return true
    }
  case kbd.Right:
    if B.depth == 0 {
      if B.index < B.width - 1 {
        B.index++
      }
    } else {
      return true
    }
  case kbd.Down, kbd.Up:
    return true
  case kbd.Pos1:
    if B.depth == 0 {
      B.index = 0
    } else {
      return true
    }
  case kbd.End:
    if B.depth == 0 {
      B.index = B.width
      for {
        if B.index == 0 { break }
        if (*Text) [B.index-1] == space {
          B.index--
        } else {
          break
        }
      }
    } else {
      return true
    }
  case kbd.Tab:
    return true
  case kbd.Back:
    switch B.depth { case 0:
      if B.index > 0 {
        B.index--
        Rem (Text, B.index, 1)
        *Text = *Text + " "
      }
    case 1:
      B.index = 0
      *Text = Clr (B.width)
      if B.overwritable {
        B.overwritable = ! B.overwritable
      }
    default:
      return true
    }
    B.write (*Text, x, y)
  case kbd.Del:
    switch B.depth { case 0:
      if B.index < ProperLen (*Text) {
        Rem (Text, B.index, 1)
        *Text = *Text + " "
      }
    case 1:
      if B.overwritable {
        B.index = 0
        *Text = Clr (B.width)
      } else {
        return true
      }
    default:
      return true
    }
    B.write (*Text, x, y)
  case kbd.Ins:
    if B.depth == 0 {
      B.overwritable = ! B.overwritable
    } else {
      return true
    }
  case kbd.Help, kbd.LookFor, kbd.Act, kbd.Cfg, kbd.Mark, kbd.Demark, kbd.Paste, kbd.Deposit:
    return true
  case kbd.Black, kbd.Red, kbd.Green, kbd.Blue:
    return true
  case kbd.PrintScr, kbd.Roll, kbd.Pause:
    return true
  case kbd.Go:
    ;
  case kbd.Here, kbd.There, kbd.This:
    ; // return true
  case kbd.Pull, kbd.Push, kbd.Move:
    ;
  case kbd.Hither, kbd.Thither, kbd.Thus:
    ;
  case kbd.Navigate:
    ;
  }
  return false
}


func (B *Imp) possible (Text *string, x, y uint) bool {
//
  if B.index < B.width {
    if B.overwritable { return true }
    if (*Text)[B.width - 1] == space {
      if ! B.overwritable { // move Text one to the right and write again
        InsSpace (Text, B.index) // -> this operation
        *Text = (*Text)[0:B.width]  // -> to str
        B.write (*Text, x, y)
      }
      return true
    }
  } else { // B.index >= B.width
    // editNumber
  }
  return false
}


func (B *Imp) editText (imGraphikmodus bool, Text *string, x, y uint) {
//
  var char byte
  var cursorshape scr.Shape
  B.graphical = imGraphikmodus
// if B.usesMouse { scr.SwitchMouseCursor (true) }
  Norm (Text, B.width)
  B.overwritable = ! Empty (*Text)
  B.index = 0
  B.write (*Text, x, y)
  B.overwritable = ! Empty (*Text)
  B.write (*Text, x, y)
  if B.start > 0 && B.start < B.width {
    B.index = B.start
    B.start = 0
  } else {
    B.index = 0
  }
  for {
    if B.overwritable {
      cursorshape = scr.Block
    } else {
      cursorshape = scr.Understroke
    }
    if B.graphical {
      scr.WarpGr (x + scr.NX1 () * B.index, y, cursorshape)
    } else {
      scr.Warp (y / scr.NY1 (), x / scr.NX1 () + B.index, cursorshape)
    }
    for {
      char, B.command, B.depth = kbd.Read ()
      if B.command < kbd.Go {
        break
      }
    }
    edited = char != byte(0)
    if B.graphical {
      scr.WarpGr (x + scr.NX1 () * B.index, y, scr.Off)
    } else {
      scr.Warp (y / scr.NY1 (), x / scr.NX1 () + B.index, scr.Off)
    }
    if B.command == kbd.None {
      if B.index == B.width {
        // see editNumber
      } else {
        if B.possible (Text, x, y) {
          Replace (Text, B.index, char)
          scr.Lock()
          scr.Colours (B.cF, B.cB)
          if B.graphical {
            scr.Write1Gr (char, int(x + scr.NX1 () * B.index), int(y))
          } else {
            scr.Write1 (char, y / scr.NY1 (), x / scr.NX1 () + B.index)
          }
          scr.Unlock()
          B.index++
        }
      }
    } else {
      if B.done (Text, x, y) {
        break
      }
    }
  }
// if B.usesMouse { scr.SwitchMouseCursor (false) }
}


// Precondition: n > 0, len (S) >= n - 1.
// Returns true, if S contains a character != ' ' in a position < n.
func leftNotEmpty (S string, n uint) bool {
//
  if n == 0 || len (S) + 1 < int(n) { return false }
  for i:= 0; i < int(n) - 1; i++ {
    if S[i] != ' ' { return true }
  }
  return false
}


type
  stati byte; const (
  start = iota
  bp // before '.'
  ap // after '.'
  ee // after 'E', i.e. in exponent
)
var
  status stati


func getStatus (Text *string) {
//
  var p uint
  if Contains (*Text, 'E', &p) {
    status = ee
  } else if Contains (*Text, '.', &p) {
    status = ap
  } else if Empty (*Text) {
    status = start
  } else {
    status = bp
  }
}


func (B *Imp) doneNumerical (Text *string, x, y uint) bool {
//
  switch B.command { case kbd.Enter, kbd.Esc:
    return true
/*
  case kbd.Left:
    if B.depth == 0 {
      if B.index > 0 && leftNotEmpty (Text, B.index) {
        B.index--
      }
    } else {
      return true
    }
  case kbd.Right:
    if B.depth == 0 {
      if B.index < B.width - 1 {
        B.index++
      }
    }
    return true
  case kbd.Down, kbd.Up: 
    return |
  case kbd.Pos1:
    if B.depth == 0 {
      B.index = 0
      for {
        if B.index == B.width { break }
        if Text [B.index] != space { break }
        B.index++
      }
    } else {
      return true
    }
  case kbd.End:
    if B.depth == 0 {
      B.index = B.width
    } else {
      return true
    }
*/
  case kbd.Back:
    switch B.depth { case 0:
      if B.overwritable {
        if B.index == 0 {
        } else {
          Rem (Text, B.index - 1, 1)
          *Text += " "
        }
      } else if B.index < B.width {
        Rem (Text, B.index, 1)
        *Text += " "
        B.index++
      } else if B.index == B.width {
        Rem (Text, B.width - 1, 1)
        *Text = " " + *Text
      }
    case 1:
      *Text = Clr (B.width)
      status = start
      B.index = B.width
    default:
      return true
    }
    getStatus (Text)
    if B.index < B.width {
      B.write (*Text, x, y)
    } else {
      i:= B.index
      B.index = 0
      B.write (*Text, x, y)
      B.index = i
    }
  case kbd.Del:
    switch B.depth { case 0:
      if B.overwritable {
        if B.index == 0 {
        } else {
          Rem (Text, B.index - 1, 1)
          *Text += " "
        }
      } else if B.index < B.width {
        Rem (Text, B.index, 1)
        *Text += " "
        B.index++
      } else if B.index == B.width {
        Rem (Text, B.width - 1, 1)
        *Text = " " + *Text
      }
    case 1:
      *Text = Clr (B.width)
      B.index = B.width
    default:
      return true
    }
    if B.index < B.width {
      B.write (*Text, x, y)
    } else {
      i:= B.index
      B.index = 0
      B.write (*Text, x, y)
      B.index = i
    }
/*
  case kbd.Ins:
    if B.depth == 0 {
      if B.overwritable {
        B.overwritable = false
      } else if i < B.width {
        B.overwritable = true
      }
    } else {
      return true
    }
*/
  case kbd.Go:
    ;
  case kbd.Here, kbd.There, kbd.This:
    return true
  case kbd.Pull, kbd.Push, kbd.Move:
    ;
  case kbd.Hither, kbd.Thither, kbd.Thus:
    ;
  default:
    return true
  }
  return false
}


func (B *Imp) possibleNumerical (Text *string, x, y uint) bool {
//
  if B.index < B.width {
    panic ("uff") // return false
    if B.overwritable { return true }
    if (*Text)[B.width - 1] == ' ' {
      // if ! overwritable, shift Text one to the right and Write
      InsSpace (Text, B.index)
      B.write (*Text, x, y)
      return true
    }
  } else { // overwritable == false
    i:= uint(0)
    for {
      if i + 2 == B.width {
        break
      }
      if (*Text)[i] == '0' && (*Text)[i + 1] == '0' {
        Replace (Text, i, ' ')
      } else {
        break
      }
      i++
    }
    if (*Text)[0] == ' ' {
      if B.width > 1 {
        Rem (Text, 0, 1)
        *Text = *Text + " "
      }
      return true
    }
  }
  return false
}


func (B *Imp) editNumber (imGraphikmodus bool, Text *string, x, y uint) {
//
  var (
    char byte
    cursorshape scr.Shape
    temp uint
    firstTime bool
  )
  B.graphical = imGraphikmodus
//  if B.usesMouse { scr.SwitchMouseCursor (true) }
  Norm (Text, B.width)
  B.overwritable = ! Empty (*Text)
  Move (Text, false)
  B.index = 0
  B.write (*Text, x, y)
  B.index = B.width
  if B.TRnumerical {
    firstTime = true
    edited = false
    // Zahl beim ersten Lesen eines Zeichens zurücksetzen, s.u.
  } else {
    edited = true
  }
  for {
    getStatus (Text)
    if B.overwritable {
      cursorshape = scr.Block
    } else {
      cursorshape = scr.Understroke
    }
    if B.graphical {
      scr.WarpGr (x + scr.NX1() * B.index, y, cursorshape)
    } else {
      scr.Warp (y / scr.NY1(), x / scr.NX1() + B.index, cursorshape) // Off
    }
    for {
      char, B.command, B.depth = kbd.Read ()
      switch char { case 0: // Command
        break
      case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
        if B.TRnumerical {
          if firstTime {
            *Text = Clr (B.width)
            status = start
            firstTime = false
            edited = true
          }
        }
        if status == start {
          status = bp
          break
        } else if status == ee {
          if Contains (*Text, 'E', &temp) {
            if temp >= B.width - 3 { // not more than 2 digits after 'E'
              break
            }
          }
        } else {
          break
        }
      case '-':
        if B.TRnumerical {
          kbd.DepositCommand (kbd.None)
          kbd.DepositByte (char)
          return
        } else {
          if Empty (*Text) || (*Text)[B.width - 1] == 'E' {
            break
          }
        }
      case  '.', ',':
        if status == bp {
          status = ap
          break
        }
      case 'E':
        if B.numerical || B.TRnumerical {
          if status == ap && // noch Platz für zwei Zeichen
             (*Text)[0] == space && (*Text)[1] == space {
            status = ee
            if B.numerical {
              break
            } else {
              Rem (Text, B.width - 2, 2)
              *Text = *Text + "E+"
              char = 0
              break
            }
          }
        }
      case 'v':
        char = 0
        if B.TRnumerical { // || B.numerical {
          if status == bp || status == ap {
            temp = 0
            for (*Text)[temp] == space { temp++ }
            if (*Text)[temp] == '-' {
              Replace (Text, temp, '+')
              break
            } else if (*Text)[temp] == '+' {
              Replace (Text, temp, '-')
              break
            } else if temp > 0 {
              Replace (Text, temp - 1, '-')
              break
            }
          } else if status == ee {
            if Contains (*Text, 'E', &temp) {
              if (*Text)[temp + 1] == '-' {
                Replace (Text, temp + 1, '+')
                break
              } else if (*Text)[temp + 1] == '+' {
                Replace (Text, temp + 1, '-')
                break
              }
            }
          }
        }
      default:
        if B.TRnumerical {
   // >>> Besser wäre dies nur für den Fall, dass 'Zeichen' ein Funktionszeichen aus dem Zahlen-Modul ist:
          kbd.DepositCommand (kbd.None)
          kbd.DepositByte (char)
          return
        }
      }
    }
    if B.graphical {
      scr.WarpGr (x + scr.NX1() * B.index, y, scr.Off)
    } else {
      scr.Warp (y / scr.NY1(), x / scr.NX1() + B.index, scr.Off)
    }
    if B.command == kbd.None {
      if B.index == B.width {
        if B.overwritable {
          B.overwritable = false
        }
        if char == 0 { // change of sign or exponent
          temp = B.index
          B.index = 0
          B.write (*Text, x, y)
          B.index = temp
        } else if B.possibleNumerical (Text, x, y) {
          temp = B.index
          B.index = 0
          B.write (*Text, x, y)
          B.index = temp
          Replace (Text, B.index - 1, char)
          scr.Lock()
          scr.Colours (B.cF, B.cB)
          if B.graphical {
            scr.Write1Gr (char, int(x + scr.NX1() * (B.index - 1)), int(y))
          } else {
            scr.Write1 (char, y / scr.NY1(), x / scr.NX1() + B.index - 1)
          }
          scr.Unlock()
        } else {
        }
      } else {
        // see editText
      }
    } else {
      if B.doneNumerical (Text, x, y) {
        break
      }
    }
  }
// if B.usesMouse { scr.SwitchMouseCursor (false) }
}

/*
func isDigit (b byte) bool {
//
  return '0' <= b && b <= '9'
}


func (B *Imp) editNumber1 (imGraphikmodus bool, s *string, x, y uint) {
//
  for uint(len (*s)) < B.width { *s = " " + *s }; B.Write (*s, x, y); for *s != "" && (*s)[0] == ' ' { *s = (*s)[1:] }; if *s == " " { *s = "" }
  var char byte
  if B.graphical {
    scr.WarpGr (x + scr.NX1() * B.width, y, scr.Understroke)
  } else {
    scr.Warp (y / scr.NY1(), x / scr.NX1() + B.width, scr.Understroke)
  }
  loop: for {
    l:= uint(len (*s))
    char, B.command, B.depth = kbd.Read ()
    switch B.command {
    case kbd.None:
      if isDigit (char) && l < B.width {
        *s += string(char)
      }
    case kbd.Esc:
      break loop
    case kbd.Enter:
      break loop
    case kbd.Back, kbd.Del:
      if l > 0 {
        *s = (*s)[:l-1]
      }
    }
    for uint(len (*s)) < B.width { *s = " " + *s }; B.Write (*s, x, y); for *s != "" && (*s)[0] == ' ' { *s = (*s)[1:] }; if *s == " " { *s = "" }
  }
  if B.graphical {
    scr.WarpGr (x + scr.NX1() * B.width, y, scr.Off)
  } else {
    scr.Warp (y / scr.NY1(), x / scr.NX1() + B.width, scr.Off)
  }
}
*/

func (B *Imp) Edit (s *string, l, c uint) {
//
  if l >= scr.NLines() { return }
  if c >= scr.NColumns() { return }
  n, b:= uint(len (*s)), B.width
  if c + b > scr.NColumns() { B.width = scr.NColumns() - c }
  if B.width == 0 { B.width = n }
//  if B.width > n { B.width = n }
  if B.width < n { Norm (s, B.width) }
  B.graphical = false
  if B.numerical || B.TRnumerical {
    B.editNumber (false, s, scr.NX1() * c, scr.NY1() * l)
  } else {
    B.editText (false, s, scr.NX1() * c, scr.NY1() * l)
  }
  B.width = b
}


func (B *Imp) EditGr (s *string, x, y uint) {
//
  if y >= scr.NY() { return }
  if x >= scr.NX() - scr.NX1() { return }
  n, b:= uint(len (*s)), B.width
  if x + B.width * scr.NX1() > scr.NX() {
    B.width = (scr.NX() - uint(x)) / scr.NX1()
  }
  if B.width == 0 { B.width = n }
  if B.width < n { Norm (s, B.width) }
//  if B.width > n { B.width = n }
  B.graphical = true
  if B.numerical || B.TRnumerical {
    B.editNumber (true, s, x, y)
  } else {
    B.editText (true, s, x, y)
  }
  B.width = b
}


func Edited () bool {
//
  return edited
}
