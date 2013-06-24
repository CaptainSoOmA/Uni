package species

// (c) Christian Maurer   v. 120909 - license see murus.go

import
  . "murus/obj"
type
  System byte; const (
  Eco = iota // Ecosystem: foxes, hares and plants
  Life       // Game of Life (John Conway)
  NSystems
) // The actual system at the beginning is Life.
const (
  Width  = 16; Height = 16 // pixel size of single species
)
const ( // format
  Short = iota // ein kleines Bild der Größe Width x Height
  Long         // ein kurzer Text
)
var (
  Suffix string
  NNeighbours uint
)

// The actual system is s.
func Sys (s System) { sys(s) }

type Species interface {

  Editor
  Formatter
  Marker
// TODO: Inc/Dec; Number() uint; type numberOf func (*Imp) uint; Modify (n MumberOf)
}
