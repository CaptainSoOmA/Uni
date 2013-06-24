package enum

// (c) Christian Maurer   v. 130117 - license see murus.go

import
  . "murus/obj"
const ( // Format
  Short = iota
  Long
  NFormats
)
type
  Enum byte; const (
  Title = Enum(iota); Composer; RecordLabel; AudioMedium; SparsCode
  Religion; Subject
  LexicalCategory
  Casus; Genus; Persona; Numerus; Tempus; Modus; GenusVerbi
  NEnums
)
const ( undefLexCat = Enum(iota)
  Substantiv; Adjektiv; Pronomen; Numerale; Verb; Adverb; Präposition; Konjunktion; Interjektion
)
const ( undefCasus = Enum(iota)
  Nominativ; Genitiv; Dativ; Akkusativ; Ablativ
)
const ( undefGenus = Enum(iota)
  masc; fem; neut
)
// const ( undefPersona = Enum(iota); Erste; Zweite; Dritte )
const ( undefNumerus = Enum(iota)
  Singular; Plural
)
const ( undefTempus = Enum(iota)
  Präsens; Imperfekt; FuturI; Perfekt; Plusquamperfekt; FuturII
)
const ( undefModus = Enum(iota)
  Indikativ; Konjunktiv
)
const ( undefGenusVerbi = Enum(iota)
  Aktiv; Passiv
)
type
  Enumerator interface { // A set of at most 256 enumerated elements,
                         // represented by strings of len <= 64.
                         // The 0-th element is "empty", represented by spaces.

  Formatter
  Editor
  Stringer
  Printer

// Returns the type of x.
  Typ () byte

// Returns the number of elements of Enum (common for all elements).
  Num () uint

// Returns the order number of x.
  Ord () uint

// Returns the width of the string representation of x (common for all elements).
  Wd () uint

// Returns true, iff there is an n-th element in Enum.
// In this case x is that element, otherwise x is empty.
  Set (n uint) bool
}
