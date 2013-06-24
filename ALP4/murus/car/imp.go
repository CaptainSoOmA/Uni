package car

// (c) Christian Maurer   v. 120420 - license see murus.go

import (
  "murus/col"
  "murus/scr"
)
var
  car = [...]string {
   "                         *      ",
   "                         *      ",
   "      ************       *      ",
   "     ***************     *      ",
   "    ***      *      *    *      ",
   "   ***       *       *   *      ",
   "  ***        *        *  *      ",
   " ****        *         * *      ",
   "******************************* ",
   "**************  ************** *",
   "* **************************** *",
   "* ***************************** ",
   "******************************* ",
   " *****************************  ",
   "     *****          *****       ",
   "      ***            ***        " }


func draw (right bool, c col.Colour, X, Y int) {
//
  scr.Lock()
  for y:= 0; y < H; y++ {
    for x:= 0; x < W; x++ {
      if car [y][x] == '*' {
        scr.Colour (c)
      } else {
        scr.Colour (col.ScreenB)
      }
      if right {
        scr.Point (X + x, Y + y)
      } else {
        scr.Point (X + W - 1 - x, Y + y)
      }
    }
  }
  scr.Unlock()
}
