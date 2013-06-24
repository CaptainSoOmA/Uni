package linie

// (c) Christian Maurer   v. 130526 - license see murus.go

import ( // "murus/env";
  "murus/font"; "murus/col"; "murus/scr"; "murus/v" // ; "murus/errh"
)


func init () {
//
  v.Want (13, 1, 19)
  Text[Fußweg] = "F"; Farbe[Fußweg] = col.Colour3 (170, 170, 170)
  Text[U1] =  "U1";  Farbe[U1]  = col.Colour3 ( 83, 177,  71)
  Text[U2] =  "U2";  Farbe[U2]  = col.Colour3 (241,  89,  35)
  Text[U3] =  "U3";  Farbe[U3]  = col.Colour3 ( 22, 166, 150)
  Text[U4] =  "U4";  Farbe[U4]  = col.Colour3 (255, 212,   2)
  Text[U5] =  "U5";  Farbe[U5]  = col.Colour3 (131,  90,  67)
  Text[U55] = "U55"; Farbe[U55] = col.Colour3 (131,  90,  67)
  Text[U6] =  "U6";  Farbe[U6]  = col.Colour3 (129, 114, 173)
  Text[U7] =  "U7";  Farbe[U7]  = col.Colour3 (  6, 158, 211)
  Text[U8] =  "U8";  Farbe[U8]  = col.Colour3 (  0,  97, 159)
  Text[U9] =  "U9";  Farbe[U9]  = col.Colour3 (234, 133,  28)
  Text[S1] =  "S1";  Farbe[S1]  = col.Colour3 (229,  76, 255) // ab hier alte Farben
  Text[S2] =  "S2";  Farbe[S2]  = col.Colour3 (  0, 179,   0)
  Text[S25] = "S25"; Farbe[S25] = Farbe[S2]
  Text[S3] =  "S3";  Farbe[S3]  = col.Colour3 (  0, 115, 242)
  Text[S41] = "S41"; Farbe[S41] = col.Colour3 (54,  38, 208)
  Text[S42] = "S42"; Farbe[S42] = col.Colour3 (91,  76, 208)
  Text[S4] =  "S4";  Farbe[S4]  =  Farbe[S41]
  Text[S45] = "S45"; Farbe[S45] = col.Colour3 (92, 128,  38)
  Text[S46] = "S46"; Farbe[S47] = Farbe[S45]
  Text[S47] = "S47"; Farbe[S47] = Farbe[S45]
  Text[S5] =  "S5";  Farbe[S5]  = col.Colour3 (254,  89,   0)
  Text[S7] =  "S7";  Farbe[S7]  = col.Colour3 (102,  88, 243)
  Text[S75] = "S75"; Farbe[S75] = Farbe[S7]
  Text[S8] =  "S8";  Farbe[S8]  = col.Colour3 ( 75, 243,   0)
  Text[S85] = "S85"; Farbe[S85] = Farbe[S8]
  Text[S9] =  "S9";  Farbe[S9]  = col.Colour3 (127,   0,  77)
/*
  if true { // zu Testzwecken
    T, M, sch:= env.Par (1); scr.Mode (), false
    if len (T) == 1 {
      sch = true
      switch T[0] {
      case 'v': M = VGA
      case 'x': M = XGA
      case 'w': M = WXGA
      case 's': M = SXGA
      case 'u': M = UXGA
      default: sch = false
      }
    }
    if sch && scr.Switchable (M) {
      scr.Switch (M)
    }
  }
*/
  M:= scr.MaxMode()
  scr.Switch (M)
//  f, l, b:= v.Colours(); t:= ""; errh.MurusLicense ("sunetz", "19. Januar 2013", "Christian Maurer", f, l, b, &t)
  if M <= scr.XGA {
    scr.SwitchFontsize (font.Tiny)
  } else {
    scr.SwitchFontsize (font.Small)
  }
  col.ScreenF, col.ScreenB = col.Black, col.LightWhite
  scr.Cls()
}
