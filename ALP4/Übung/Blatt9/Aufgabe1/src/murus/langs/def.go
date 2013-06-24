package langs

// (c) Christian Maurer   v. 130115 - license see murus.go

import (
  . "murus/obj"
//  "murus/enum"
)
type
  LanguageSequence interface {

  Formatter
  Editor
  Printer

//  Num (l[]*subject.Imp, v, b[]uint) uint
}
