package col

// (c) Christian Maurer   v. 130526 - license see murus.go

const
  P6 = 3
type
  Colour struct {
        R, G, B byte
                }

// Returns the colour defined by (r, g, b).
func Colour3 (r, g, b uint) Colour { return colour3(r,g,b) }

// Returns the rgb-values of c scaled to the range from 0 to 1.
func Float (c Colour) (float32, float32, float32) { return float(c) }
func LongFloat (c Colour) (float64, float64, float64) { return longFloat(c) }

// Returns a random colour.
func ColourRand () Colour { return colourRand() }

// Returns true, if c and c1 coincide in their rgb-values.
func Eq (c, c1 Colour) bool { return eq(c,c1) }

// Returns true, if c is, what the name of the func says.
func IsBlack (c Colour) bool { return isBlack(c) }
func IsLightwhite (c Colour) bool { return isLightWhite(c) }

// c is changed in a manner suggested by the name of the func.
func Invert (c *Colour) { invert(c) }
func Contrast (c *Colour) { contrast(c) }

// Returns true, iff s is a string of 3 values in sedecimal basis
// (with uppercase letters). In that case, c is the colour with
// the corresponding rgb-values; otherwise, nothing has happened.
func Defined (c *Colour, s string) bool { return defined(c,s) }

// Returns "rrggbb", where "rr", "gg" and "bb" are the rgb-values
// in sedecimal basis (with uppercase letters).
func String (c Colour) string { return string_(c) }

// TODO Spec
func Change (c *Colour, rgb, d byte, l bool) { change(c,rgb,d,l) }

// see murus/obj/coder.go
func Codelen () uint{ return codelen }
func Encode (c Colour) []byte { return encode(c) }
func Decode (c *Colour, b []byte) { decode(c,b) }

// Pre: d is one of 4, 8, 15, 16, 24 or 32.
// The internal colourdepth is d.
func SetColourDepth (d uint) { setColourDepth(d) }

// Returns the number of available colours,
// depending on the internal colour depth.
func Number () uint { return number() }

// TODO Spec
func Code (c Colour) uint { return code(c) }
func P6Encode (A, P []byte) { p6Encode(A,P) }
func P6Colour (A []byte) Colour { return p6Colour(A) }

// Pre: The internal colourdepth is defined (by a call to
//      SetColourdDepth) due to the hardware of the computer.
// ActualF/ActualB are set to f/b,
// CodeF/CodeB is set to the values of Code(f)/Code(b).
func Actualize (f, b Colour) { actualize(f,b) }

// see Actualize.
func Init () { actualize(ScreenF,ScreenB) }

var (
  bitColourdepth,
  CodeF, CodeB uint
  ScreenF, ScreenB, // at the beginning White and Black resp.
  ActualF, ActualB, // actual colours for write- and draw-operations,
                    // coincide at the beginning with the colours of the Screen
  HintF, HintB, ErrorF, ErrorB, // colours for hints and error reports
  Black, Brown, Red, LightRed, Yellow, Green, // standard colours
  Cyan, Blue, Magenta, LightMagenta, Gray, White, LightWhite,
// other colours:
  DarkBrown, Siena,
  LightBrown, Cream, LightCream,
  Carmine, Crimson, FlashRed, PompejiRed, SignalRed, CinnabarRed,
  Orange, LightOrange,
  DarkYellow, FlashYellow,
  BirchGreen, FlashGreen,
  PrussianBlue, FlashBlue,
  FlashCyan,
  FlashMagenta,
  Pink,
  Silver,
  MurusF, MurusB Colour
// RAL-Farben: ral.go
// colours of XConsortium: x.go
)
