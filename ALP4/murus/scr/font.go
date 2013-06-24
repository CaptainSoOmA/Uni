package scr

// (c) Christian Maurer   v. 120909 - license see murus.go

import (
  "murus/xker"; "murus/font"
)
var
  actualFontsize font.Size


func actFontsize () font.Size {
//
  return actualFontsize
}

/*
func correspondingFont (f font.Size) fonts {
//
  switch f { case font.Tiny:
    return tiny
  case font.Small:
    return small
  case font.Normal:
    break
  case font.Big:
    return big
  case Huge:
    return huge
  }
  return normal
}
*/

func swFontsize (f font.Size) {
//
  actualFontsize = f
  actualCharheight = charheight (actualFontsize)
  actualCharwidth = charwidth (actualFontsize)
  nLines = nY [mode] / actualCharheight
  nColumns = nX [mode] / actualCharwidth
}


func switchFontsize (f font.Size) {
//
  swFontsize (f)
  if underX {
    xker.SetFontsize (actualCharheight)
  }
}
