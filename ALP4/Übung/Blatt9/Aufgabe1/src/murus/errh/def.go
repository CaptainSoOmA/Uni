package errh

// (c) Christian Maurer   v. 130224 - license see murus.go

import
  "murus/col"
var (
  ToWait, ToContinue, ToContinueOrNot, ToCancel, ToScroll,
  ToSelect, ToChange, ToSwitch, ToSelectWithPrint, ToPrint string
)

// s is written to the last line of the screen
// resp. to the screen starting at position (line, column) == (l, c).
func Hint (s string) { hint(s) }
func HintPos (s string, l, c uint) { hintPos(s,l,c) }
// The hints are deleted, the former content of the screen is restored.
func DelHint () { delHint() }
func DelHintPos (s string, l, c uint) { delHintPos(s,l,c) }

// s is written to the last line of the screen.
// The calling process is blocked, until Escape or Backspace is pressed;
// then the former content of the last line of the screen is restored.
func Error (s string, n uint) { error(s,n) }
func Error2 (s string, n uint, s1 string, n1 uint) { error2(s,n,s1,n1) }

// s is written to the screen, starting at position (line, column) == (l, c).
// The calling process is blocked, until Escape or Backspace is pressed;
// then the former content of the screen, starting at (l, c), is restored.
func ErrorPos (s string, n, l, c uint) { errorPos(s,n,l,c) }
func Error2Pos (s string, n uint, s1 string, n1 uint, l, c uint) { error2Pos(s,n,s1,n1,l,c) }

// The calling process is blocked, until the user has confirmed by some action.
func Confirmed () bool { return confirmed() }

// TODO Spec
//func WriteLicense (p, v, a string, f, l, b col.Colour, g []string, t *string) { writeLicense(p,v,a,f,l,b,g,t) }
func MurusLicense (p, v, a string, f, l, b col.Colour, t *string) { murusLicense(p,v,a,f,l,b,t) }
func WriteHeadline (p, v, a string, f, b col.Colour) { writeHeadline(p,v,a,f,b) }

// h is written to the center of the screen.
// The calling process is blocked, until Enter, Esc, Back or a mouse button is pressed;
// then the former content of the screen is restored.
func WriteHelp (h []string) { writeHelp(h) }

// TODO Spec
func WriteHelp1 () { writeHelp1() }
