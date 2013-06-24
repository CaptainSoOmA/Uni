package font

// (c) Christian Maurer   v. 120909 - license see murus.go

var (
  code = [NFonts]string { "r", "b", "s", "i" }
  size = [NSizes]string { "t", "s", "n", "b", "h" }
)


func Code (f Font, s Size) string {
//
  x, y:= "?", "?"
  if f < NFonts {
    x = code[f]
  }
  if s < NSizes {
    y = size[s]
  }
  return x + y
}


func init () {
//
  Name = []string { "winzig", "klein ", "normal", "groß  ", "riesig" }
}
