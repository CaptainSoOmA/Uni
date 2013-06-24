package species

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  . "murus/obj"
  "murus/col"; "murus/scr"
)
type
  numberOf func (*Imp) uint
const (
  max = 8
)
const ( // Eco
  plant = iota
  hare
  fox
  nSpec // > nSpec of Life
)
const ( // Life
  nothing = iota
  cell
)
const
  plantNothing = iota
const (
  unmarked = iota
  marked
  nMarks
)
type (
  image [Height]string
  Imp struct {
          nr,      // < nSpec
        mark uint8 // 0..1
             }
     representation struct {
                   word string
                    img image
                 cf, cb [nMarks]col.Colour
                        }
)
var (
  system System
  rep [nSpec]representation
  format Format
  Vergleichsart [nSpec]*Imp
)


func sys (s System) {
// My icons of hare and fox are dreadful - I ask for help !
  system = s
  var p image
  switch system { case Eco:
    rep [plant].word = "Pflanzen"
    rep [plant].cf [unmarked], rep [plant].cb [unmarked] = col.Green, col.Black
    rep [plant].cf [marked],   rep [plant].cb [marked]   = col.Black, col.LightGreen
    p = image {
      "*     *   *     ",
      " *   *     *   *",
      "  * *     *   * ",
      "   *       *  * ",
      " * *  *    * *  ",
      "  ** *      *   ",
      "   **   *  **   ",
      "  ***   * *  *  ",
      "    *    *   *  ",
      "    *    *   *  ",
      "   ***   *    * ",
      "  *  *  *      *",
      "  *   **    * * ",
      " *     *     *  ",
      "*       *   * * ",
      "       *   *   *" }
    for i:= 0; i < Height; i++ {
      rep [plant].img [i] = p [i]
    }
    rep [hare].word = "Hasen"
    rep [hare].cf [unmarked], rep [hare].cb [unmarked] = col.Yellow, col.Black
    rep [hare].cf [marked],   rep [hare].cb [marked]   = col.Black,  col.LightYellow
    p = image {
      "     *     *    ",
      "     *     *    ",
      "     **   **    ",
      "     * * * *    ",
      "     * *** *    ",
      "     * * * *    ",
      "      *   *     ",
      "     * ***  *   ",
      "    *        *  ",
      "   * * * * * *  ",
      "   * **   ** *  ",
      "    * *   * *   ",
      "     *     *    ",
      "     **  **     ",
      "    * *  * *    ",
      "     **  **     " }
    for i:= 0; i < Height; i++ {
      rep [hare].img [i] = p [i]
    }
    rep [fox].word = "Füchse"
    rep [fox].cf [unmarked], rep [fox].cb [unmarked] = col.Brown, col.Black
    rep [fox].cf [marked],   rep [fox].cb [marked]   = col.Black, col.LightBrown
    p = image {
      "                ",
      "                ",
      "                ",
      "                ",
      "      *         ",
      "     **         ",
      "   ****         ",
      "  *****         ",
      "    **          ",
      "    ********    ",
      "   **********   ",
      "   ********  *  ",
      "    *      *  * ",
      "    *       *   ",
      "    *       *   ",
      "                "}
    for i:= 0; i < Height; i++ {
      rep [fox].img [i] = p [i]
    }
    NNeighbours = 4
    Suffix = "eco"
  case Life:
    rep [nothing].word = "--------"
    rep [nothing].cf [unmarked], rep [nothing].cb [unmarked] = col.Gray,       col.Black
    rep [nothing].cf [marked],   rep [nothing].cb [marked]   = col.LightWhite, col.Black
    p = image {
      "                ",
      "                ",
      "                ",
      "                ",
      "                ",
      "       *        ",
      "      * *       ",
      "     *   *      ",
      "      * *       ",
      "       *        ",
      "                ",
      "                ",
      "                ",
      "                ",
      "                ",
      "                "}
    for i:= 0; i < Height; i++ {
      rep [nothing].img [i] = p [i]
    }
    rep [cell].word = "Zellen"
    rep [cell].cf [unmarked], rep [cell].cb [unmarked] = col.LightMagenta, col.Black
    rep [cell].cf [marked],   rep [cell].cb [marked]   = col.LightWhite,   col.Black
    p = image {
      "                ",
      "                ",
      "      ****      ",
      "    **    **    ",
      "   *        *   ",
      "  *          *  ",
      " *            * ",
      " *            * ",
      " *            * ",
      " *            * ",
      "  *          *  ",
      "   *        *   ",
      "    **    **    ",
      "      ****      ",
      "                ",
      "                "}
    for i:= 0; i < Height; i++ {
      rep [cell].img [i] = p [i]
    }
    NNeighbours = 8
    Suffix = "life"
  }
}


func New () *Imp {
//
  x:= new (Imp)
  return x
}


func (x *Imp) Empty () bool {
//
  return x.nr == plantNothing
}


func (x *Imp) Clr () {
//
  x.nr = plantNothing
}


func (x *Imp) Eq (X Object) bool {
//
  y, ok:= X.(*Imp)
  if ! ok { return false }
  return x.nr == y.nr
}


func (x *Imp) Less (Y Object) bool {
//
  return false
}


func (x *Imp) Copy (X1 Object) {
//
  y, ok:= X1.(*Imp)
  if ! ok {
    return
  }
  x.nr = y.nr
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Inc () {
//
  if x.nr + 1 < uint8(x.Number()) {
    x.nr ++
  } else {
    x.nr = plantNothing
  }
}


func (x *Imp) Dec () {
//
  if x.nr > 0 {
    x.nr --
  } else {
    x.nr = plantNothing
  }
}


func (x *Imp) Number () uint {
//
  if system == Eco {
    return 3
  }
  return 2 // Life
}


func (x *Imp) Mark (b bool) {
//
  x.mark = unmarked
  if b {
    x.mark = marked
  }
}


func (x *Imp) Marked () bool {
//
  return x.mark == marked
}


func (X *Imp) SetFormat (f Format) {
//
  format = f
}


func (X *Imp) SetColours (f, b col.Colour) {
//
// fake to be Editor
}


func (X *Imp) Write (l, c uint) {
//
  r:= rep [X.nr]
  switch format { case Short:
    switch system { case Life:
      f, b:= r.cf [X.mark], r.cb [X.mark]
      if X.nr == nothing {
        f, b = b, f
      }
      x, y:= int (8 * c) + 8, int (16 * l) + 8
      scr.Colour (f)
      scr.Circle (x, y, 6)
      scr.Circle (x, y, 5)
      scr.Colour (b)
      scr.CircleFull (x, y, 2)
    case Eco:
      for y:= 0; y < Height; y++ {
        for x:= 0; x < Width; x++ {
          f, b:= r.cf [X.mark], r.cb [X.mark] // !
          if r.img [y][x] == ' ' {
            f, b = b, f
          }
          scr.Colours (f, b)
          scr.Point (Width * int(c) / 2 + x, Height * int(l) + y)
        }
      }
    }
  case Long:
    scr.Colour (r.cf [plantNothing])
    scr.Write (r.word, l, c)
  }
}


func (X *Imp) Edit (l, c uint) {
//
// fake to be Editor
}


func (x *Imp) Modify (no numberOf) {
//
  switch system { case Life:
    c:= no (Vergleichsart [cell])
    switch x.nr { case nothing:
      if c == 3 { x.nr = cell }
    case cell:
      if c < 2 || c > 3 { x.nr = nothing }
    }
  case Eco:
    h:= no (Vergleichsart [hare])
    f:= no (Vergleichsart [fox])
    switch x.nr { case plant:
      if h > 0 && h < 4 { x.nr = hare }
    case hare:
      if h == 4 { x.nr = plant }
      if f > 0 { x.nr = fox }
    case fox:
      if h == 0 { x.nr = plant }
    }
  }
}


func (x *Imp) Codelen () uint {
//
  return 1
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, 1)
  b[0] = x.nr
  return b
}


func (x *Imp) Decode (b []byte) {
//
  x.nr = b[0]
  x.mark = unmarked
}


func init () {
//
//  var sp Species = New(); if sp == nil {} // ok
  Sys (Life)
  for n:= uint8(0); n < nSpec; n++ {
    Vergleichsart [n] = New ()
    Vergleichsart [n].nr = n
  }
  format = Short
}
