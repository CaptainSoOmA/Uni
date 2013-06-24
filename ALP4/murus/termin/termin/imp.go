package termin

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  . "murus/obj"
  "murus/font"; "murus/kbd"; "murus/col"
  "murus/text"; "murus/clk"
  "murus/termin/word"; "murus/termin/attr"
)
const
  max = 56
type
  Imp struct {
           c *clk.Imp
           a *attr.Imp
           w *word.Imp
           t *text.Imp
      marked bool
             }
var (
  actFmt Format
  cS, cH, cMS, cMH, ctF, ctB, ctMF, ctMB col.Colour
)


func New () *Imp {
//
  x:= new (Imp)
  x.c = clk.New ()
  x.a = attr.New ()
  x.w = word.New ()
  x.t = text.New (54) // 80 - 5 - 5 - 12 - 4 (NCols - clk - attr - word - spaces)
  x.t.SetColours (ctF, ctB)
  return x
}


func (x *Imp) Empty () bool {
//
  return x.c.Empty () &&
         x.a.Empty () &&
         x.w.Empty () &&
         x.t.Empty ()
}


func (x *Imp) Clr () {
//
  x.c.Clr ()
  x.a.Clr ()
  x.w.Clr ()
  x.t.Clr ()
  x.Mark (false)
}


func (x *Imp) Copy (X Object) {
//
  x1, ok:= X.(*Imp)
  if ! ok { return }
  x.c.Copy (x1.c)
  x.a.Copy (x1.a)
  x.w.Copy (x1.w)
  x.t.Copy (x1.t)
  x.marked = x1.marked
}


func (x *Imp) Clone () Object {
//
  x1:= New ()
  x1.Copy (x)
  return x1
}


func (x *Imp) Eq (X Object) bool {
//
// buggy ?
  x1, ok:= X.(*Imp)
  if ! ok { return false }
  if x.Empty () {
    return x1.Empty ()
  } else if x1.Empty () {
    return false
  }
  return x.c.Eq (x1.c) &&
         x.a.Eq (x1.a) &&
         x.w.Eq (x1.w) && x.t.Eq (x.t)
}


func (x *Imp) Less (X Object) bool {
//
// TODO buggy ?
  x1, ok:= X.(*Imp)
  if ! ok { return false }
  if x.Empty () {
    return false
  }
  if x1.Empty () {
    return true
  }
  if x.c.Eq (x1.c) {
    if x.a.Eq (x1.a) {
      if x.w.Eq (x1.w) {
        return x.t.Less (x1.t)
      }
      return x.w.Less (x1.w)
    }
    return x.a.Less (x1.a)
  }
  return x.c.Less (x1.c)
}


func (x *Imp) HasWord () bool {
//
  return x.w.Ok ()
}


func (x *Imp) Attribute () *attr.Imp {
//
  return x.a.Clone ().(*attr.Imp)
}


func (x *Imp) Mark (ja bool) {
//
  x.marked = ja
  x.a.Mark (ja)
  x.w.Mark (ja)
  if ja {
    x.t.SetColours (ctMF, ctMB)
  } else {
    x.t.SetColours (ctF, ctB)
  }
}


func (x *Imp) Marked () bool {
//
  return x.marked
}


var
  as, ws, ts, ds uint


func (x *Imp) SetFormat (f Format) {
//
  actFmt = f
  switch f { case Lang:
    ds = 1
    x.a.SetFormat (attr.Lang)
  case Kurz:
    ds = 1
    x.a.SetFormat (attr.Mittel)
//    attr.SetFormat (attr.Kurz)
  case GanzKurz:
    ds = 0
    x.a.SetFormat (attr.Kurz)
  }
  as = attr.Wd + 0 + ds
  ws = as + 5 + ds
  ts = ws + 12 + ds
}


func (x *Imp) SetColours (f, b col.Colour) {
//
// dummy to get the interface right
}


func (x *Imp) Write (z, s uint) {
//
  switch actFmt { case Lang, Kurz:
    if x.marked {
      x.c.SetColours (cMS, cMH)
    } else {
      x.c.SetColours (cS, cH)
    }
    x.c.Write (z, s)
  }
  switch actFmt { case Lang, Kurz:
    x.a.Write (z, s + as)
  }
  switch actFmt { case Lang:
    x.w.Write (z, s + ws)
    x.t.Write (z, s + ts)
  case Kurz:
    // x.w.Write (z + 1, s)
  case GanzKurz:
    // nichts mehr schreiben
  }
}


func (x *Imp) Edit (z, s uint) {
//
  x.Write (z, s)
  i:= 0
  loop:
  for {
    switch i { case 0:
      x.c.Edit (z, s)
    case 1:
      x.a.Edit (z, s + as)
    case 2:
      switch actFmt { case Lang:
        x.w.Edit (z, s + ws)
      default:
//       x.w.Edit (z + 1, s)
      }
    case 3:
      if actFmt == Lang {
        x.t.Edit (z, s + ts)
      }
    }
    var e uint
    c:= kbd.LastCommand (&e)
    switch c { case kbd.Help:
      ;
    case kbd.Enter: // kbd.Here:
      if e == 0 {
        if i < 3 { i++ } else { break loop }
      } else {
        break loop
      }
    case kbd.Esc: // , kbd.There:
      break loop
    case kbd.Pos1, kbd.End:
      break loop
    case kbd.Down, kbd.Right:
      if e > 0 { break loop }
      if i < 3 { i++ } else { break loop }
    case kbd.Up, kbd.Left:
      if e > 0 { break loop }
      if i > 0 { i-- } else { break loop }
    case kbd.Del:
      if e == 0 {
        x.Clr ()
        x.Write (z, s)
      } else {
        break loop
      }
    case kbd.Mark, kbd.Demark:
      x.Mark (c == kbd.Mark)
      x.Write ( z, s)
      break loop
/*
    case kbd.Deposit:
      break loop
    case kbd.Paste:
      break loop
    case kbd.Mark:
      break loop
    case kbd.LookFor:
      break loop
*/
    default:
      break loop
    }
  }
}


func (x *Imp) SetFont (f font.Font) {
//
}


func (x *Imp) Print (z, s uint) {
//
  x.c.Print (z, s)
  x.a.Print (z, s + 6)
  x.w.Print (z, s + 11)
  x.t.Print (z, s + 21)
}


func (x *Imp) Codelen () uint {
//
  return x.c.Codelen () +
         x.a.Codelen () +
         x.w.Codelen () +
         max
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  i:= uint(0)
  a:= x.c.Codelen()
  copy (b[i:i+a], x.c.Encode ())
  i += a
  a = x.a.Codelen ()
  copy (b[i:i+a], x.a.Encode ())
  i += a
  a = x.w.Codelen ()
  copy (b[i:i+a], x.w.Encode ())
  i += a
  a = x.t.Codelen ()
  copy (b[i:i+a], x.t.Encode ())
  return b
}


func (x *Imp) Decode (b []byte) {
//
  i:= uint(0)
  a:= x.c.Codelen()
  x.c.Decode (b[i:i+a])
  i += a
  a = x.a.Codelen ()
  x.a.Decode (b[i:i+a])
  i += a
  a = x.w.Codelen ()
  x.w.Decode (b[i:i+a])
  i += a
  a = x.t.Codelen ()
  x.t.Decode (b[i:i+a])
  x.marked = false
}


func init () {
//
  cS, cH =  col.Black, col.Yellow
  cMS, cMH = col.Yellow, col.Black
  ctF, ctB =  col.Black, col.LightWhite
  ctMF, ctMB = col.Black, col.White
  var _ Termin = New()
}
