package sel

// (c) Christian Maurer   v. 130127 - license see murus.go

import (
  "murus/ker"; . "murus/obj"; "murus/env"; "murus/str"; "murus/kbd"
  "murus/col"; "murus/scr"; "murus/box"; "murus/errh"
  "murus/files"
)
const
  pack = "sel"
var
  bx, mbx box.Box = box.New(), box.New()


func select_ (write WritingCol, n, h, w uint, i *uint, l, c uint, f, b col.Colour) {
//
  if n == 0 { ker.Stop (pack, 1) }
  if n == 1 { *i = 0; return }
  if h == 0 { ker.Stop (pack, 2) }
  if h > n { h = n }
  if w == 0 { w = scr.NColumns () }
  if w > scr.NColumns () { w = scr.NColumns () }
  if c + w > scr.NColumns () { c = scr.NColumns () - w }
 // so, dass letzte Zeile frei bleibt
  if l + h >= scr.NLines () {
    h = scr.NLines () - l - 1
  }
  if *i >= n { *i = n - 1 }
  MouseOn:= scr.MouseCursorOn()
  var x, y int
  if MouseOn {
    scr.MouseCursor (false)
    x, y = scr.MousePosGr ()
  }
  scr.WarpMouse (l + *i, c)
  scr.Save (l, c, w, h)
  i0, n0:= uint(0), uint(0)
  if *i == 0 { n0 = 1 } // else { n0 = 0 }
  neu:= true
  loop: for {
    if *i < i0 {
      i0 = *i
      neu = true
    } else if *i > i0 + h - 1 {
      i0 = *i - (h - 1)
      neu = true
    } else {
      neu = *i != n0
    }
    if neu {
      neu = false
      var cF, cB col.Colour
      for j:= uint(0); j < h; j++ {
        if i0 + j == *i {
          cF, cB = f, b
        } else {
          cF, cB = b, f
        }
        write (i0 + j, l + j, c, cF, cB)
      }
    }
    n0 = *i
    C, d:= kbd.Command()
    switch C { case kbd.Esc, kbd.Thither:
      *i = n
      break loop
    case kbd.Enter, kbd.Hither:
      break loop
    case kbd.Left, kbd.Up:
      if d == 0 {
        if *i > 0 {
          *i --
        }
      } else {
        if *i >= 10 {
          *i -= 10
        }
      }
    case kbd.Right, kbd.Down:
      if d == 0 {
        if *i + 1 < n {
          *i ++
        }
      } else {
        if *i + 10 < n {
          *i += 10
        }
      }
    case kbd.Pos1:
      *i = 0
    case kbd.End:
      *i = n - 1
    case kbd.Go:
      _, yM:= scr.MousePosGr ()
      if uint(yM) <= l * scr.NY1() + scr.NY1() / 2 {
        if *i > 0 {
          *i --
        }
      } else if uint(yM) >= (l + h) * scr.NY1() {
        if *i < n - 1 {
          *i ++
        }
      } else {
        *i = i0 + uint(yM) / scr.NY1() - l
      }
/*
    case kbd.Help:
      errh.Hint (errh.zumAuswaehlen)
      kbd.Wait (true)
      errh.DelHint()
*/
    }
  }
  scr.Restore (l, c, w, h)
  if MouseOn {
    scr.MouseCursor (true)
    scr.WarpMouseGr (x, y)
  }
}


func select1 (Line []string, h, w uint, i *uint, l, c uint, f, b col.Colour) {
//
//  if len (Zeile) + 1 < h { h = len (Zeile) + 1 }
  bx.Wd (w)
  Select (func (k, l, c uint, f, b col.Colour) { bx.Colours (f, b); bx.Write (Line[k], l, c) }, h, h, w, i, l, c, f, b)
}


var
  ptSuffix string


func hasSuffix (a Any) bool {
//
  var p uint
  return str.IsPart (ptSuffix, a.(string), &p) &&
         p == str.ProperLen (a.(string)) - uint(len (ptSuffix))
}


func aus (n, l, c uint, f, b col.Colour) {
//
  N:= files.NamePred (hasSuffix, n)
  var p uint
  if str.IsPart (ptSuffix, N, &p) {
    N = str.Part (N, 0, p)
  }
  bx.Colours (f, b)
  bx.Write (N, l, c)
}


func names (mask, suffix string, n uint, l, c uint, f, b col.Colour) (string, string) {
//
  t, t1:= uint(len (mask)), uint(0)
  if t > 0 {
    t1 = 1 + t
  }
  scr.Save (l, c, t1 + n, 1)
  if t > 0 {
    mbx.Wd (t)
    mbx.ColoursScreen ()
    mbx.Write (mask, l, c)
  }
  bx.Wd (n)
  bx.Colours (f, b)
  ptSuffix = "." + suffix
  errh.Hint ("falls Dateien vorhanden, auswählen F2-, dann Pfeil-/Eingabetaste, ggf. Esc")
  name:= env.Par (1)
  if name == "" {
    name = str.Clr (n) // Wörkeraunt um Fehler in box/imp.go
  }
  var p uint
  if str.Contains (name, '.' , &p) {
    name = str.Part (name, 0, p)
  }
  bx.Edit (&name, l, c + t1)
  str.RemSpaces (&name)
  if str.Contains (name, '.', &p) {
    name = str.Part (name, 0, p)
  }
  filename:= name + ptSuffix
  a:= files.NumPred (hasSuffix)
  if a > 0 {
    var d uint
    switch kbd.LastCommand (&d) { case kbd.Esc:
      return "", "" // str.Clr (n), ""
    case kbd.Enter:
      // entered
    case kbd.LookFor:
      i:= uint(0)
      select_ (aus, a, a, n, &i, l, c + t1, b, f)
      if i == a {
        return "", "" // str.Clr (n), ""
      } else {
        str.Set (&filename, files.NamePred (hasSuffix, i))
      }
    }
  }
  errh.DelHint()
  str.RemSpaces (&filename)
  if str.Contains (filename, '.', &p) {
    name = str.Part (filename, 0, p)
  }
  scr.Restore (l, c, t1 + n, 1)
  return name, filename
}
