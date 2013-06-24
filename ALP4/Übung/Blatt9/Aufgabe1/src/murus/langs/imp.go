package langs

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  . "murus/obj"
  "murus/col"; "murus/box"; "murus/errh"
  "murus/font"; "murus/pbox"
  "murus/bnat"; "murus/enum"
//  "murus/subject"
)
const (
  min = 2
  max = 4
)
const ( // Format
  Short = iota
  Long
  NFormats
)
type
  Imp struct {
        lang [max]*enum.Imp
    from, to [max]*bnat.Imp
      cF, cB col.Colour
           f Format
             }
var (
  bx *box.Imp = box.New ()
  pbx *pbox.Imp = pbox.New ()
  lLa, cLa, lFr, cFr, lTo, cTo [NFormats][max]uint
)


func (x *Imp) imp (Y Object) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}



func New () *Imp {
//
  x:= new (Imp)
  for n:= uint(0); n < max; n++ {
    x.lang[n] = enum.New (enum.Subject)
    x.from[n], x.to[n] = bnat.New (11), bnat.New (11)
  }
  x.cF, x.cB = col.ScreenF, col.ScreenB
  return x
}


func (x *Imp) Empty () bool {
//
  for n:= uint(0); n < max; n++ {
    if ! x.lang[n].Empty () { return false }
  }
  return true
}


func (x *Imp) Clr () {
//
  for n:= uint(0); n < max; n++ {
    x.lang[n].Clr ()
    x.from[n].Clr ()
    x.to[n].Clr ()
  }
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= x.imp (Y)
  for n:= uint(0); n < max; n++ {
    if ! x.lang[n].Eq (y.lang[n]) {
      return false
    }
    if ! x.from[n].Eq (y.from[n]) || ! x.to[n].Eq (y.to[n]) {
      return false
    }
  }
  return true
}


func (x *Imp) Less (Y Object) bool {
//
  return false
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  for n:= uint(0); n < max; n++ {
    x.lang[n].Copy (y.lang[n])
    x.from[n].Copy (y.from[n])
    x.to[n].Copy (y.to[n])
  }
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.Copy (x)
  return y
}


func (x *Imp) Num (lg []*enum.Imp, from, to[]uint) uint {
//
  n:= uint(0)
  for {
   if x.lang[n].Empty () {
      break
    } else if n < max - 1 {
      n ++
    } else {
      break
    }
  }
  if n == 0 { return 0 }
  lg = make ([]*enum.Imp, n)
  from, to = make ([]uint, n), make ([]uint, n)
  for i:= uint(0); i < n; i++ {
    lg[i].Copy (x.lang[i])
    from[i], to[i] = x.from[i].Val (), x.to[i].Val ()
  }
  return n
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  for n:= uint(0); n < max; n++ {
    x.lang[n].SetColours (f, b)
    x.from[n].SetColours (f, b)
    x.to[n].SetColours (f, b)
  }
}


func (x *Imp) writeMask (l, c uint) {
//
  bx.ColoursScreen ()
  switch x.f { case Short:
/*        1         2         3         4
012345678901234567890123456789012345678901
_ (__-__)  e (__-__)  g (__-__)  _ (__-__) */
    bx.Wd (7)
    for n:= uint(0); n < max; n++ {
      bx.Write ("(  -  )", l, c + cFr[x.f][n] - 1)
    }
  case Long:
/*        1         2         3
0123456789012345678901234567890123456789
___________ von Klasse __ bis Klasse __ */
    bx.Wd (10)
    for n:= uint(0); n < max; n++ {
      bx.Write ("von Klasse", l + lFr[x.f][n], c + cFr[x.f][n] - 11)
      bx.Write ("bis Klasse", l + lTo[x.f][n], c + cTo[x.f][n] - 11)
    }
  }
}


func (x *Imp) SetFormat (f Format) {
//
  x.f = f
  for n:= uint(0); n < max; n++ {
    if x.f == Short {
      x.lang[n].SetFormat (enum.Short)
    } else {
      x.lang[n].SetFormat (enum.Long)
    }
  }
}


func (x *Imp) Write (l, c uint) {
//
  x.writeMask (l, c)
  for n:= uint(0); n < max; n++ {
    x.lang[n].Write (l + lLa[x.f][n], c + cLa[x.f][n])
    x.from[n].Write (l + lFr[x.f][n], c + cFr[x.f][n])
    x.to[n].Write (l + lTo[x.f][n], c + cTo[x.f][n])
  }
}


func (x *Imp) multiple (n *uint) bool {
//
  for i:= uint(1); i < max; i++ {
    for k:= i + 1; k < i; k++ {
      if ! x.lang[i].Empty () && x.lang[k].Eq (x.lang[i]) {
        *n = k
        return true
      }
    }
  }
  *n = 0
  return false
}


func (x *Imp) Edit (l, c uint) {
//
  x.Write (l, c)
  n:= uint(0)
  loop_n:
  for {
    i:= uint(0)
    loop_i: for {
      weg:= false
      switch i { case 0: // lang
        for {
          for {
            x.lang[n].Edit (l + lLa[x.f][n], c + cLa[x.f][n])
            if x.lang[n].Ord() >= 2 && // Englisch
               x.lang[n].Ord() <= 12 { // Griechisch
              break
            } else {
              errh.Error ("keine Fremdsprache", 0)
            }
          }
          if x.lang[n].Empty () {
            if n > 1 {
              weg = true
              break
            } else {
              errh.Error2 ("", n + 1, ". Fremdsprache fehlt", 0)
            }
          } else {
            break
          }
        }
      case 1: // from
        if weg {
          x.from[n].Clr ()
          x.from[n].Write (l + lFr[x.f][n], c + cFr[x.f][n])
        } else {
          x.from[n].Edit (l + lFr[x.f][n], c + cFr[x.f][n])
        }
      case 2: // to
        if weg {
          x.to[n].Clr ()
          x.to[n].Write (l + lTo[x.f][n], c + cTo[x.f][n])
          } else {
            for {
              x.to[n].Edit (l + lTo[x.f][n], c + cTo[x.f][n])
              if x.to[n].Empty() || x.to[n].Val () == 0 && x.to[n].Val () >= 12 {
                errh.Error ("geht nich", x.to[n].Val ())
            } else {
              break
            }
          }
        }
      }
      if i < 2 {
        i ++
      } else {
        if x.from[n].Eq (x.to[n]) || x.from[n].Less (x.to[n]) {
          break loop_i
        } else {
          i = 1
        }
      }
    } // loop_i
    if n + 1 < max {
      n ++
    } else {
      k:= uint(0)
      if x.multiple (&k) {
        errh.Error2 ("Die", k + 1, ". Fremdsprache kommt mehrfach vor", 0)
      } else {
        break loop_n
      }
    }
  }
}


func (x *Imp) printMask (l, c uint) {
//
  switch x.f { case Short:
    for n:= uint(0); n < max; n++ {
      pbx.Print ("(  -  )", l, c + cFr[x.f][n] - 1)
    }
  case Long:
    for n:= uint(0); n < max; n++ {
      pbx.Print ("von Klasse", l + lFr[x.f][n], c + cFr[x.f][n] - 11)
      pbx.Print ("bis Klasse", l + lTo[x.f][n], c + cTo[x.f][n] - 11)
    }
  }
}


func (x *Imp) SetFont (f font.Font) {
//
  for n:= uint(0); n < max; n++ {
    x.lang[n].SetFont (f)
    x.from[n].SetFont (f)
    x.to[n].SetFont (f)
  }
}


func (x *Imp) Print (l, c uint) {
//
  x.printMask (l, c)
  for n:= uint(0); n < max; n++ {
    x.lang[n].Print (l + lLa[x.f][n], c + cLa[x.f][n])
    x.from[n].Print (l + lFr[x.f][n], c + cFr[x.f][n])
    x.to[n].Print (l + lTo[x.f][n], c + cTo[x.f][n])
  }
}


func (x *Imp) Codelen () uint {
//
  return max * (x.lang[0].Codelen() + 1)
}


func (x *Imp) Encode ()[]byte {
//
  b:= make ([]byte, x.Codelen())
  i:= uint(0)
  for n:= uint(0); n < max; n++ {
    a:= x.lang[n].Codelen()
    copy (b[i:i+a], x.lang[n].Encode())
    i += a
    c:= byte (x.from[n].Val () + 16 * x.to[n].Val ())
    copy (b[i:i+1], Encode (c))
    i ++
  }
  return b
}


func (x *Imp) Decode (b[]byte) {
//
  i:= uint(0)
  for n:= uint(0); n < max; n++ {
    a:= x.lang[n].Codelen()
    x.lang[n].Decode (b[i:i+a])
    i += a
    c:= uint(Decode (byte(0), b[i:i+1]).(byte))
    x.from[n].Set (c % 16)
    x.to[n].Set (c / 16)
    i ++
  }
}


func init () {
//
  var _ LanguageSequence = New ()
  for n:= uint(0); n < max; n++ {
    lLa[Short][n] = 0; cLa[Short][n] = 11 * n
    lFr[Short][n] = 0; cFr[Short][n] = 11 * n + 3
    lTo[Short][n] = 0; cTo[Short][n] = 11 * n + 6
    lLa[Long][n] = n; cLa[Long][n] = 0
    lFr[Long][n] = n; cFr[Long][n] = 23
    lTo[Long][n] = n; cTo[Long][n] = 37
  }
}
