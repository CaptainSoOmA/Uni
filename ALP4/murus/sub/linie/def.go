package linie

// (c) Christian Maurer   v. 130118 - license see murus.go

import
  "murus/col"
type
  Linie byte; const (
  Fußweg = Linie(iota)
  U1; U2; U3; U4; U5; U55; U6; U7; U8; U9
  S1; S2; S25; S3; S41; S42; S4; S45; S46; S47; S5; S6; S7; S75; S8; S85; S9
  NLinien)
var (
  Text [NLinien]string
  Farbe [NLinien]col.Colour
)
