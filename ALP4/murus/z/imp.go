package z

// (c) Christian Maurer   v. 130419 - license see murus.go

const
  delta = 'a' - 'A'


func isLatin1 (b byte) bool {
//
  switch b { case
    AE, OE, UE, Ae, Oe, Ue, Sz,
    Euro, Para, Degree, ToThe2, ToThe3, Mue, Copyright: // ,
    // Registered, Pound, Female, Male, PlusMinus, Times, Division, Negate:
    return true
  }
  return false
}


func isLowerUmlaut (b byte) bool {
//
  switch b { case Ae, Oe, Ue, Sz:
    return true
  }
  return false
}


func isCapUmlaut (b byte) bool {
//
  switch b { case AE, OE, UE:
    return true
  }
  return false
}


func opensHell (b byte) bool {
//
  return b == byte(194) ||
         b == byte(195)
}


func devilsDung (s *string) bool {
//
  n:= len (*s)
  if n == 0 {
    return false
  }
  for i:= 0; i < n; i++ {
    switch (*s)[i] { case 194, 195:
      return true
    }
  }
  return false
}


func toHellWithUTF8 (s *string) {
//
  n:= len (*s)
  if n == 0 { return }
  bs:= []byte(*s)
  i, k:= 0, 0
  var b byte
  for i < n {
    b = bs[i]
    switch b { case 194:
      i++
      b = bs[i]
    case 195:
      i++
      b = bs[i] + 64
    }
    bs[k] = b
    i++
    k++
  }
  if k == n {
    return
  } else if k < n {
    *s = string(bs[:k])
  }
}


func Equiv (a, b byte) bool {
//
  switch { case a < 'A':
    return a == b
  case a <= 'Z', 'a' <= a && a <= 'z',
     a == AE, a == OE, a == UE,
     a == Ae, a == Oe, a == Ue:
    // see below
  default:
    return a == b
  }
  return a & 31 == b & 31
}


func cap (b byte) byte {
//
  switch b {
  case Ae:
    return AE
  case Oe:
    return OE
  case Ue:
    return UE
  }
  if 'a' <= b && b <= 'z' {
    return b - delta
  }
  return b
}


func lower (b byte) byte {
//
  switch b {
  case AE:
    return Ae
  case OE:
    return Oe
  case UE:
    return Ue
  }
  if 'A' <= b && b <= 'Z' {
    return b + delta
  }
  return b
}


var (
  nr [256]byte
  in [256]bool
)


func Less (a, b byte) bool {
//
  if a == b {
    return false
  }
  if in[a] {
    if in[b] {
      return nr[a] < nr[b]
    } else {
      return true // Sonderzeichen hinter Buchstaben
    }
  } else {
    if in[b] {
      return false // s. o.
    }
  }
  return a < b // nach ASCII
}


func postscript (b byte) string {
//
  switch b {
  case AE:
    return "Adieresis"
  case OE:
    return "Odieresis"
  case UE:
    return "Udieresis"
  case Ae:
    return "adieresis"
  case Oe:
    return "odieresis"
  case Ue:
    return "udieresis"
  case Sz:
    return "germandbls"
  case Euro:
    return "Euro"
  case Para:
    return "section"
  case Degree:
    return "degree"
/*
  case ToThe2:
    return ""
  case ToThe3:
    return ""
*/
  case Mue:
    return "mu"
  case Copyright:
    return "copyright"
/*
  case Registered:
    return "registered"
  case Pound:
    return "sterling"
  case Female:
    return ""
  case Male:
    return ""
  case PlusMinus:
    return "plusminus"
  case Times:
    return "multiply"
  case Division:
    return ""
  case Negate:
    return ""
*/
  }
  return ""
}


func init () {
//
  ord:= []byte(" 0123456789Aa  BbCcDdEeFfGgHhIiJjKkLlMmNnOo  PpQqRrSs TtUu  VvWwXxYyZz")
//                           Ää                            Öö        ß    Üü
//              0         1         2         3         4         5         6
//              0123456789012345678901234567890123456789012345678901234567890123456789
  ord[13] = AE
  ord[14] = Ae
  ord[43] = OE
  ord[44] = Oe
  ord[53] = Sz
  ord[58] = UE
  ord[59] = Ue
//  for b:= byte(0); b < byte(len (ord)); b++ {
//    eingeordnet[ord[b]] = false
//  }
  for b:= byte(0); b < byte(len (ord)); b++ {
    nr[ord[b]] = b
    in[ord[b]] = true
  }
}
