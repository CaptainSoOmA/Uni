package cntry

// (c) Christian Maurer   v. 130115 - license see murus.go

import
  . "murus/obj"
const ( // Format
  Tld = iota // Top-Level-Domain, z.B. "de",          "it"
  Long       // Bezeichnung,      z.B. "Deutschland", "Italien"
  Tel        // Telefon-Vorwahl,  z.B. 49,            39
  Car        // Kfz-Kennzeichen,  z.B. "D",           "I"
  Ioc        // IOC-Code,         z.B. "GER",         "ITA"
  NFormats
)
type
  Continent byte; const (
  Europa = iota
  Afrika
  Amerika
  Asien
  Ozeanien
  NContinents
)
type
  Country interface {

  Editor
  Formatter
  Stringer
  Printer

  InContinent (c Continent) bool
}
