package col

// (c) Christian Maurer   v. 130526 - license see murus.go

import
  "murus/rand"
const (
  m  = 1<<8
  m1 = m - 1
  light = byte(m1)
  m3 = m1 / 3
  m2 = 2 * m3
  codelen = 3
)


func colour3 (r, g, b uint) Colour {
//
  var c Colour
  c.R, c.G, c.B = byte(r % m), byte(g % m), byte(b % m)
  return c
}


func ansiEncode (c Colour) uint { // 0..15 // doch vielleicht Mist
//
  const (black = 1<<iota / 2; red; green; blue; light)
  r, g, b:= c.R, c.G, c.B
  n:= uint(black)
  if r >= m2 { r -= m2; n += red }
  if g >= m2 { g -= m2; n += green }
  if b >= m2 { b -= m2; n += blue }
  if r >= m3 && g >= m3 && b >= m3 { n += light }
  return n
}


func float (c Colour) (float32, float32, float32) {
//
  const m1f = float32(m1)
  return float32(c.R) / m1f, float32(c.G) / m1f, float32(c.B) / m1f
}


func longFloat (c Colour) (float64, float64, float64) {
//
  const m1f = float64(m1)
  return float64(c.R) / m1f, float64(c.G) / m1f, float64(c.B) / m1f
}


func colourRand () Colour {
//
  var c Colour
  c.R, c.G, c.B = byte (rand.Natural (m)), byte (rand.Natural (m)), byte (rand.Natural (m))
  return c
}


func changeRand (c *Colour) {
//
  const (N = 32; N2 = N / 2)
  n:= byte(rand.Natural (N))
  if n < N2 {
    if c.R >= n {
      c.R -= n
    } else {
      c.R = 0
    }
  } else { // n >= N2
    n -= N2
    if c.R + n <= m1 {
      c.R += n
    } else {
      c.R = light
    }
  }
  n = byte(rand.Natural (N))
  if n < N2 {
    if c.G >= n {
      c.G -= n
    } else {
      c.G = 0
    }
  } else {
    n -= N2
    if c.G + n <= m1 {
      c.G += n
    } else {
      c.G = light
    }
  }
  n = byte(rand.Natural (N))
  if n < N2 {
    if c.B >= n {
      c.B -= n
    } else {
      c.B = 0
    }
  } else {
    n -= N2
    if c.B + n <= m1 {
      c.B += n
    } else {
      c.B = light
    }
  }
}


func eq (c, c1 Colour) bool {
//
  return c.R == c1.R &&
         c.G == c1.G &&
         c.B == c1.B
}



func isBlack (c Colour) bool {
//
  return c.R == 0 &&
         c.G == 0 &&
         c.B == 0
}


func isLightWhite (c Colour) bool {
//
  return c.R == light &&
         c.G == light &&
         c.B == light
}


func invert (c *Colour) {
//
  c.R, c.G, c.B = m1 - c.R, m1 - c.G, m1 - c.B
}


func contrast (c *Colour) {
//
  const lightlimit = 352 // 320 352 384 416 448 480 512 <-- difficult problem,
                         // highly dependent of the intensity of green,
                         // and our eyes are particularly sensible for green !
  if c.G > 224 {
    *c = Black
  } else if int(c.R) + int(c.G) + int(c.B) < lightlimit {
    *c = LightWhite
  } else {
    *c = Black
  }
}


func ok (b byte) bool {
//
  if b < '9' {
    return true
  } else if 'A' <= b && b <= 'F' {
    return true
  }
  return false
}


func value (b byte) uint{
//
  if b < '9' {
    return uint(b - '0')
  } else if 'A' <= b && b <= 'F' {
    return uint(b - 'A' + 10)
  }
  return 0
}


func defined (c *Colour, s string) bool {
//
  if len (s) != 6 { return false }
  for i:= 0; i < 6; i++ {
    if ! ok (s[i]) { return false }
  }
  c.R = byte(16 * value (s[0]) + value (s[1]))
  c.G = byte(16 * value (s[2]) + value (s[3]))
  c.B = byte(16 * value (s[4]) + value (s[5]))
  return true
}


func change (c *Colour, rgb, d byte, lighter bool) {
//
  if rgb > 2 || d > 127 { return }
  switch rgb { case 0:
    if lighter {
      if c.R <= m1 - d {
        c.R += d
      }
    } else if c.R >= d {
      c.R -= d
    }
  case 1:
    if lighter {
      if c.G <= m1 - d {
        c.G += d
      }
    } else if c.G >= d {
      c.G -= d
    }
  case 2:
    if lighter {
      if c.B <= m1 - d {
        c.B += d
      }
    } else if c.B >= d {
      c.B -= d
    }
  }
}


func char (n uint) string {
//
  if n < 9 {
    return string (n + 48)
  } else if n < 15 {
    return string (n + 65)
  }
  return string (0)
}


func string_ (c Colour) string {
//
  s:= char (uint(c.R) / 16) + char (uint(c.G) % 16)
  s+= char (uint(c.G) / 16) + char (uint(c.G) % 16)
  s+= char (uint(c.B) / 16) + char (uint(c.B) % 16)
  return s
}


func encode (c Colour) []byte {
//
  b:= make ([]byte, 3)
  b[0], b[1], b[2] = c.R, c.G, c.B
  return b
}


func decode (c *Colour, b []byte) {
//
  if len (b) == 3 {
    c.R, c.G, c.B = b[0], b[1], b[2]
  } else {
    *c = LightWhite
  }
}


func setColourDepth (depth uint) {
//
  switch depth { case 4, 8, 15, 16, 24, 32:
    bitColourdepth = depth
  default:
    bitColourdepth = 0
  }
}


func number () uint{
//
  switch bitColourdepth { case 4:
    return 16
  case 8:
    return m
  case 15:
    return 128 * m
  case 16:
    return m * m
  case 24, 32:
    return m * m * m
  }
  return 0
}


func code (c Colour) uint {
//
  switch bitColourdepth { case 4:
    return ansiEncode (c)
  case 8:
    return 8 * (4 * (uint(c.R) / 32) + uint(c.G) / 32) + uint(c.B) / 64 // direct colour
  case 15:
    return 32 * (32 * (uint(c.R) / 8) + uint(c.G) / 8) + uint(c.B) / 8
  case 16:
    return 64 * (32 * (uint(c.R) / 8) + uint(c.G) / 4) + uint(c.B) / 8
  case 24:
    return m * (m * uint(c.R) + uint(c.G)) + uint(c.B)
  case 32:
    return /* m * (m * (m * uint(c.alpha) + */ m * (m * uint(c.R) + uint(c.G)) + uint(c.B)
  }
  return 0
}


func p6Encode (A, P []byte) {
//
  switch bitColourdepth { case 4:
    ; // TODO
  case 8:
    ; // TODO
  case 15:
    ; // TODO
  case 16: // TODO: might be nonsense, has to be checked !
    P[0] = A[1] & 1 << 5
    P[1] = (A[0] & 1 << 3) >> 3 + A[1] >> 5
    P[2] = A[0] >> 3
  case 24:
    P[0] = A[2]
    P[1] = A[1]
    P[2] = A[0]
  case 32: // TODO: might be nonsense, has to be checked !
    P[0] = A[3]
    P[1] = A[2]
    P[2] = A[1] // leaving alpha out (?)
  default:
    for i:= 0; i < P6; i++ {
      P[i] = byte(0)
    }
  }
}


func p6Colour (A []byte) Colour {
//
  P:= make ([]byte, P6)
  p6Encode (A, P)
  var c Colour
  c.R, c.G, c.B = A[0], A[1], A[2]
  return c
}


func actualize (f, b Colour) {
//
  ActualF, ActualB = f, b
  CodeF, CodeB = code (ActualF), code (ActualB)
}


func init_ () {
//
  actualize (ScreenF, ScreenB) // Pre: Code is defined
}


func init () {
//
//  Schwarzbraun1 =    colour3 ( 64,  42,   0)
//  Schokoladenbraun1 = colour3 (m3,  42,   0)
//  Kastanienbraun =   colour3 (106,  64,   0)

  DarkBrown =        colour3 (127,  m3,   0)
  Siena =            colour3 (149,  m3,  42)
//  Hellsiena =        colour3 (191, 127,  42)
//  Rotbraun1 =        colour3 ( m2,  64,  64)
  Brown =            colour3 ( m2, 127,   0)
//  Umbrabraun =       colour3 (149, 135,   0)
//  Olivbraun1 =       colour3 (127, 127,   0)
//  Hellolivbraun =    colour3 (170,  m2,  m3)
//  Mittelbraun =      colour3 (149, 106,   0)
  LightBrown =       colour3 (212, 149,  64)
//  Orangebraun1 =     colour3 (127, 106,  42)
//  Dunkelocker =      colour3 ( m2, 127,  21)
//  Ocker =            colour3 ( m1,  m2,  64)
//  Hellocker =        colour3 ( m1, 191, 106)
//  Weißbraun =        colour3 ( m1, 212, 149)
//  Rosabraun =        colour3 ( m1, 191, 149)
  Cream =            colour3 ( m1, 234, 191)
  LightCream =       colour3 ( m1, 249, 224)
//  Hellbeige =        colour3 (234, 212,  m2)
//  Beige2 =           colour3 (212, 191, 149)
//  Ganzhellbraun =    colour3 (206,  m2, 127)

//  Schwarzrot1 =      colour3 ( m3,   0,   0)
//  DarkRed =          colour3 (106,   0,   0) // X
  Carmine =          colour3 (191,  64,  64)
//  CarminRed1 =       colour3 (149,  42,  64)
  Crimson =          colour3 (160,   0,   0)
  Red =              colour3 ( m2,   0,   0)
  FlashRed =         colour3 ( m1,   0,   0)
  PompejiRed =       colour3 (191,  64,  64)
  SignalRed =        colour3 (204,  m3,  42)
  CinnabarRed =      colour3 (234,   0,   0)
  LightRed =         colour3 ( m1,  m3,  m3)
//  Ziegelrot =        colour3 (212, 127,  42)
//  Weißrot =          colour3 ( m1, 149, 127)

//  DarkOrange =       colour3 (234, 127,  64)
//  BlutOrange1 =      colour3 ( m1, 112,  m3)
  Orange =           colour3 ( m1, 149,  54)
  LightOrange =      colour3 ( m1,  m2,   0) // X
//  Dunkelrosa =       colour3 (234,   0, 127)
  Pink =             colour3 ( m1,   0,  m2)
//  Rosa =             colour3 ( m1,  m2,  m2)
//  Hellrosa =         colour3 ( m1, 191, 191)

  DarkYellow =       colour3 ( m1, 212,   0) // X
  FlashYellow =      colour3 ( m1,  m1,   0)
  Yellow =           colour3 ( m1,  m1,  m3)
//  LightYellow =      colour3 ( m1,  m1,  m2) // X
//  Sandgelb1 =        colour3 (234, 206, 127)

//  Weißgrün1 =        colour3 ( m2,  m1,  m2)
//  LightGreen =       colour3 (106,  m1, 106)
  FlashGreen =       colour3 (  0,  m1,   0)
//  Zitronengelb1 =    colour3 (191,  m1,  m3)
  BirchGreen =       colour3 ( 42, 156,  42)
//  Grasgrün1 =        colour3 (  0, 144,   0)
//  Hellchromgrün =    colour3 ( m3,  m2,   0)
  Green =            colour3 (  0,  m2,   0)
//  Hellolivgrün =     colour3 ( m2, 196,  m3)
//  Gelbgrün1 =        colour3 ( m2,  m1,  m3)
//  DarkGreen =        colour3 (  0, 127,   0)
//  Tiefdunkelgrün =   colour3 (  0, 106,   0)
//  Schwarzgrün1 =     colour3 (  0,  m3,   0)
//  F244 =             colour3 ( m3, 170,  m2)
//  Wiesengrün =       colour3 (106, 212, 106)

//  Tiefdunkeltürkis = colour3 (  0,  m3,  m3)
//  F024 =             colour3 (  0,  m3,  m2)
//  DarkCyan =         colour3 (  0, 127, 127) // X
  Cyan =             colour3 (  0,  m2,  m2)
  FlashCyan =        colour3 (  0,  m1,  m1)
//  LightCyan =        colour3 ( m3,  m1,  m1)
//  Weißtürkis =       colour3 (  m2,  m1,  m1)
//  F264 =             colour3 (  m3,  m1,  m2)
//  F042 =             colour3 (   0,  m2,  m3)
//  F064 =             colour3 (   0,  m1,  m2)

//  Tiefdunkelblau =   colour3 (  0,   0,  m3)
  PrussianBlue =     colour3 (  0, 106,  m2)
//  DarkBlue =         colour3 (  0,   0, 127) // X
  Blue =             colour3 (  0,   0,  m2)
//  Enzianblau1 =      colour3 (  0,   0, 212)
  FlashBlue =        colour3 (  0,   0,  m1)
//  LightBlue =        colour3 ( m3,  m3,  m1) // X
//  Himmelblau1 =      colour3 (  0,  m2,  m1)
//  Weißblau =         colour3 ( m2,  m2,  m1)

//  Tiefdunkellila =   colour3 ( m3,   0,  m3)
//  Ultramarinblau1 =  colour3 ( 63,   0, 149)
//  DarkMagenta =      colour3 (127,   0, 127) // X
  Magenta =          colour3 ( m2,   0,  m2)
  FlashMagenta =     colour3 ( m1,   0,  m1)
  LightMagenta =     colour3 ( m1,  m3,  m1)
//  Weißlila =         colour3 ( m1,  m2,  m1)
//  F204 =             colour3 ( m3,   0,  m2)
//  F206 =             colour3 ( m3,   0,  m1)

  Black =            colour3 (  0,   0,   0)
//  Anthrazit =        colour3 ( 42,  42,  42)
//  DarkGray =         colour3 ( 63,  63,  63) // X
  Gray =             colour3 ( m3,  m3,  m3)
//  LightGray =        colour3 (127, 127, 127) // X
  White =            colour3 ( m2,  m2,  m2)
  Silver =           colour3 (212, 212, 212)
  LightWhite =       colour3 ( m1,  m1,  m1)

//  F026 =             colour3 (  0,  m3,  m1)
//  F440 =             colour3 ( m2,  m2,   0)

//  F242 =             colour3 ( m3,  m2,  m3)
//  F224 =             colour3 ( m3,  m3,  m2)
//  F402 =             colour3 ( m2,   0,  m3)
//  F406 =             colour3 ( m2,   0,  m1)
//  F422 =             colour3 ( m2,  m3,  m3)
//  F424 =             colour3 ( m2,  m3,  m2)
//  F426 =             colour3 ( m2,  m3,  m1)
//  F624 =             colour3 ( m1,  m3,  m2)

  ScreenF, ScreenB = White, Black
  HintF, HintB = LightWhite,  Magenta
  ErrorF, ErrorB = FlashYellow, Red
  MurusF, MurusB = colour3 (  0,  16,  64), colour3 (231, 238, 255)
}
