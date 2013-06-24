package sel

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  "murus/kbd"
  "murus/col"; "murus/scr"
)
const (
  M = 128 // max < M
  H = 32
)
var (
  X, Y uint
  pattern [M]col.Colour
  pattern16 [16]col.Colour
  max uint
)


func write (FZ, B, x, y uint) {
//
  f:= col.ActualF
  for i:= uint(0); i < FZ; i++ {
    switch FZ { case 16:
      scr.Colour (pattern16 [i])
    default:
      scr.Colour (pattern [i])
    }
    scr.RectangleFull (int(x + i * B), int(y), int(x + i * B + B - 1), int(y + H - 1))
  }
  col.ActualF = f
}


func define (FZ, B uint, C *col.Colour) {
//
  xi, yi:= scr.MousePosGr ()
  x, y:= uint(xi), uint(yi)
  x -= X
  x = x / B
  if x < FZ && Y <= y && y < Y + H {
    if FZ == 16 {
      *C = pattern16 [x]
    } else {
      *C = pattern [x]
    }
    scr.Colour (*C)
  } else {
    *C = col.ScreenB
  }
}


func colour (FZ, B uint) col.Colour {
//
  MausAn:= scr.MouseCursorOn ()
  if ! MausAn {
    scr.MouseCursor (true)
  }
  xm, ym:= scr.MousePosGr ()
  X, Y = uint(xm), uint(ym)
  M:= FZ * B / 2
  if X >= scr.NY() - M { X = scr.NX() - M }
  if X >= M { X -= M } else { X = 0 }
  if Y >= H { Y -= H } else { Y = 0 }
  scr.SaveGr (X, Y, X + 2 * FZ * B, Y + H)
  write (FZ, B, X, Y)
  clicked:= false
  C:= col.ScreenF
  loop: for {
    scr.MouseCursor (true)
    K, _:= kbd.Command ()
    switch K { case kbd.Esc, kbd.Back, kbd.There, kbd.This:
      break loop
    case kbd.Here:
      define (FZ, B, &C)
      clicked = true
    case kbd.Hither:
      if clicked { break loop }
    }
  }
  scr.RestoreGr (X, Y, X + 2 * FZ * B, Y + H)
  if ! MausAn {
    scr.MouseCursor (false)
  }
  return C
}


var
  n int


func use (C col.Colour) {
//
  if n > M { n = M }
  pattern [n] = C
  n++
}


func init () {
//
  pattern16[ 0] = col.Black
  pattern16[ 1] = col.Brown
  pattern16[ 2] = col.Red
  pattern16[ 3] = col.LightRed
  pattern16[ 4] = col.Yellow
  pattern16[ 5] = col.LightGreen
  pattern16[ 6] = col.Green
  pattern16[ 7] = col.Cyan
  pattern16[ 8] = col.LightCyan
  pattern16[ 9] = col.LightBlue
  pattern16[10] = col.Blue
  pattern16[11] = col.Magenta
  pattern16[12] = col.LightMagenta
  pattern16[13] = col.LightWhite
  pattern16[14] = col.White
  pattern16[15] = col.Gray

  n = 0

////  use (col.Schwarzbraun1)
////  use (col.Schokoladenbraun1)
////  use (col.Kastanienbraun)
//  use (col.Darkbrown)
////  use (col.Siena)
////  use (col.Hellsiena)
////  use (col.Rotbraun1)
//  use (col.Brown)
////  use (col.Umbrabraun)
////  use (col.Olivbraun1)
////  use (col.Hellolivbraun)
////  use (col.Mittelbraun)
//  use (col.Lightbrown)
////  use (col.Orangebraun1)
////  use (col.Dunkelocker)
////  use (col.Ocker)
////  use (col.Hellocker)
////  use (col.Weißbraun)
////  use (col.Rosabraun)
////  use (col.Creme)
////  use (col.Hellcreme)
////  use (col.Hellbeige)
////  use (col.Beige2)
////  use (col.Ganzhellbraun)

////  use (col.Schwarzrot1)
//  use (col.Darkred)
////  use (col.Karminrot1)
////  use (col.Purpurrot1)
//  use (col.Red)
//  use (col.Flashred)
////  use (col.Pompejirot)
//  use (col.Signalred)
////  use (col.Zinnoberrot)
////  use (col.Grellrot)
//  use (col.Lightred)
////  use (col.Ziegelrot)
////  use (col.Weißrot)

//  use (col.Darkorange)
////  use (col.Blutorange1)
//  use (col.Orange)
//  use (col.Lightorange)
////  use (col.Dunkelrosa)
//  use (col.Pink)
////  use (col.Rosa)
////  use (col.Hellrosa)

//  use (col.Darkyellow)
//  use (col.Flashyellow)
//  use (col.Yellow)
//  use (col.Lightyellow)
////  use (col.Sandgelb1)

////  use (col.Weißgrün1)
//  use (col.Lightgreen)
//  use (col.Flashgreen)
////  use (col.Zitronengelb1)
////  use (col.Birkengrün)
////  use (col.Grasgrün1)
////  use (col.Hellchromgrün)
//  use (col.Green)
////  use (col.Hellolivgrün)
////  use (col.Gelbgrün1)
//  use (col.Darkgreen)
////  use (col.Tiefdunkelgrün)
////  use (col.Schwarzgrün1)
////  use (col.F244)
////  use (col.Wiesengrün)

////  use (col.Tiefdunkeltürkis)
////  use (col.F024)
//  use (col.Darkcyan)
//  use (col.Cyan)
//  use (col.Flashcyan)
//  use (col.Lightcyan)
////  use (col.Weißtürkis)
////  use (col.F264)
////  use (col.F042)
////  use (col.F064)

////  use (col.Tiefdunkelblau)
////  use (col.Preußischblau)
//  use (col.Darkblue)
//  use (col.Blue)
////  use (col.Enzianblau1)
//  use (col.Flashblue)
//  use (col.Lightblue)
////  use (col.Himmelblau1)
////  use (col.Weißblau)

////  use (col.Tiefdunkellila)
////  use (col.Ultramarinblau1)
//  use (col.Darkmagenta)
//  use (col.Magenta)
//  use (col.Flashmagenta)
//  use (col.Lightmagenta)
////  use (col.Weißlila)
////  use (col.F204)
////  use (col.F206)

//  use (col.Black)
////  use (col.Anthrazit)
//  use (col.Darkgray)
//  use (col.Gray)
//  use (col.Lightgray)
//  use (col.White)
//  use (col.Silver)
//  use (col.Lightwhite)

////  use (col.F026)
////  use (col.F440)

////  use (col.F242)
////  use (col.F224)
////  use (col.F402)
////  use (col.F406)
////  use (col.F422)
////  use (col.F424)
////  use (col.F426)
////  use (col.F624)

//  use (col.Grünbeige)
//  use (col.Beige)
//  use (col.Sandgelb)
  use (col.Signalgelb)
//  use (col.Goldgelb)
//  use (col.Honiggelb)
//  use (col.Maisgelb)
  use (col.Narzissengelb)
//  use (col.Braunbeige)
  use (col.Zitronengelb)
//  use (col.Perlweiß)
//  use (col.Elfenbein)
//  use (col.Hellelfenbein)
  use (col.Schwefelgelb)
  use (col.Safrangelb)
  use (col.Zinkgelb)
//  use (col.Graubeige)
//  use (col.Olivgelb)
//  use (col.Rapsgelb)
  use (col.Verkehrsgelb)
//  use (col.Ockergelb)
  use (col.Leuchtgelb)
//  use (col.Currygelb)
//  use (col.Melonengelb)
//  use (col.Ginstergelb)
//  use (col.Dahliengelb)
//  use (col.Pastelgelb)

  use (col.Gelborange)
//  use (col.Rotorange)
//  use (col.Blutorange)
//  use (col.Pastellorange)
//  use (col.Reinorange)
//  use (col.Leuchtorange)
//  use (col.Leuchthellorange)
//  use (col.Hellrotorange)
//  use (col.Verkehrsorange)
  use (col.Signalorange)
//  use (col.Tieforange)
//  use (col.Lachsorange)

//  use (col.Feuerrot)
  use (col.Signalrot)
//  use (col.Karminrot)
//  use (col.Rubinrot)
  use (col.Purpurrot)
//  use (col.Weinrot)
//  use (col.Schwarzrot)
//  use (col.Oxidrot)
//  use (col.Braunrot)
//  use (col.Beigerot)
//  use (col.Tomatenrot)
//  use (col.Altrosa)
//  use (col.Hellrosa)
//  use (col.Korallenrot)
//  use (col.Rose)
//  use (col.Erdbeerrot)
//  use (col.Verkehrsrot)
//  use (col.Lachsrot)
  use (col.Leuchtrot)
//  use (col.Leuchthellrot)
//  use (col.Himbeerrot)
//  use (col.Orientrot)

//  use (col.Rotlila)
//  use (col.Rotmagenta)
//  use (col.Erikamagenta)
//  use (col.Bordeauxmagenta)
//  use (col.Blaulila)
//  use (col.Verkehrspurpur)
//  use (col.Purpurmagenta)
  use (col.Signalviolett)
//  use (col.Pastelviolett)
//  use (col.Telemagenta)

  use (col.Violettblau)
//  use (col.Grünblau)
//  use (col.Ultramarinblau)
//  use (col.Saphirblau)
//  use (col.Schwarzblau)
  use (col.Signalblau)
//  use (col.Brillantblau)
//  use (col.Graublau)
//  use (col.Azurblau)
//  use (col.Enzianblau)
//  use (col.Stahlblau)
//  use (col.Lichtblau)
//  use (col.Kobaltblau)
//  use (col.Taubenblau)
//  use (col.Himmelblau)
//  use (col.Verkehrsblau)
//  use (col.Türkisblau)
  use (col.Capriblau)
//  use (col.Ozeanblau)
//  use (col.Wasserblau)
//  use (col.Nachtblau)
//  use (col.Fernblau)
//  use (col.Pastellblau)

//  use (col.Patinagrün)
//  use (col.Smaragdgrün)
//  use (col.Laubgrün)
//  use (col.Olivgrün)
  use (col.Blaugrün)
//  use (col.Moosgrün)
//  use (col.Grauoliv)
//  use (col.Flaschengrün)
//  use (col.Braungrün)
//  use (col.Tannengrün)
//  use (col.Grasgrün)
//  use (col.Resedagrün)
  use (col.Schwarzgrün)
//  use (col.Schilfgrün)
//  use (col.Gelboliv)
//  use (col.Schwarzoliv)
//  use (col.Cyangrün)
//  use (col.Maigrün)
//  use (col.Gelbgrün)
  use (col.Weißgrün)
//  use (col.Chromoxidgrün)
//  use (col.Blassgrün)
//  use (col.Braunoliv)
//  use (col.Verkehrsgrün)
//  use (col.Farngrün)
//  use (col.Opalgrün)
//  use (col.Lichtgrün)
//  use (col.Kieferngrün)
//  use (col.Minzgrün)
  use (col.Signalgrün)
//  use (col.Minttürkis)
//  use (col.Pasteltürkis)

//  use (col.Fehgrau)
//  use (col.Silbergrau)
//  use (col.Olivgrau)
//  use (col.Moosgrau)
  use (col.Signalgrau)
//  use (col.Mausgrau)
//  use (col.Beigegrau)
//  use (col.Khakigrau)
//  use (col.Grüngrau)
//  use (col.Zeltgrau)
//  use (col.Eisengrau)
//  use (col.Basaltgrau)
//  use (col.Braungrau)
//  use (col.Schiefergrau)
//  use (col.Anthrazitgrau)
//  use (col.Schwarzgrau)
//  use (col.Umbragrau)
//  use (col.Betongrau)
//  use (col.Graphitgrau)
//  use (col.Granitgrau)
//  use (col.Steingrau)
//  use (col.Blaugrau)
//  use (col.Kieselgrau)
//  use (col.Zementgrau)
//  use (col.Gelbgrau)
//  use (col.Lichtgrau)
//  use (col.Platingrau)
//  use (col.Staubgrau)
//  use (col.Achatgrau)
//  use (col.Quarzgrau)
//  use (col.Fenstergrau)
//  use (col.VerkehrsgrauA)
//  use (col.VerkehrsgrauB)
//  use (col.Seidengrau)
//  use (col.Telegrau1)
//  use (col.Telegrau2)
//  use (col.Telegrau4)

//  use (col.Grünbraun)
//  use (col.Ockerbraun)
  use (col.Signalbraun)
//  use (col.Lehmbraun)
//  use (col.Kupferbraun)
//  use (col.Rehbraun)
//  use (col.Olivbraun)
//  use (col.Nussbraun)
//  use (col.Rotbraun)
//  use (col.Sepiabraun)
//  use (col.Kastanienbraun)
//  use (col.Mahagonibraun)
  use (col.Schokoladenbraun)
//  use (col.Graubraun)
//  use (col.Schwarzbraun)
//  use (col.Orangebraun)
//  use (col.Beigebraun)
//  use (col.Blassbraun)
//  use (col.Terrabraun)

//  use (col.Cremeweiß)
//  use (col.Grauweiß)
  use (col.Signalweiß)
  use (col.Signalschwarz)
//  use (col.Tiefschwarz)
//  use (col.Aluminiumweiß)
//  use (col.Aluminiumgrau)
//  use (col.Reinweiß)
//  use (col.Graphitschwarz)
//  use (col.Verkehrweiß)
//  use (col.Verkehrschwarz)
//  use (col.Papyrusweiß)

//  use (col.Snow)
//  use (col.Snow1)
//  use (col.Snow2)
//  use (col.Snow3)
//  use (col.Snow4)
//  use (col.GhostWhite)
//  use (col.WhiteSmoke)
//  use (col.Gainsboro)
//  use (col.FloralWhite)
//  use (col.OldLace)
//  use (col.Linen)
//  use (col.AntiqueWhite)
//  use (col.AntiqueWhite1)
//  use (col.AntiqueWhite2)
//  use (col.AntiqueWhite3)
//  use (col.AntiqueWhite4)
//  use (col.PapayaWhip)
//  use (col.BlanchedAlmond)
//  use (col.Bisque)
//  use (col.Bisque1)
//  use (col.Bisque2)
//  use (col.Bisque3)
//  use (col.Bisque4)
//  use (col.PeachPuff)
//  use (col.PeachPuff1)
//  use (col.PeachPuff2)
//  use (col.PeachPuff3)
//  use (col.PeachPuff4)
//  use (col.NavajoWhite)
//  use (col.NavajoWhite1)
//  use (col.NavajoWhite2)
//  use (col.NavajoWhite3)
//  use (col.NavajoWhite4)
//  use (col.Moccasin)
//  use (col.Cornsilk)
//  use (col.Cornsilk1)
//  use (col.Cornsilk2)
//  use (col.Cornsilk3)
//  use (col.Cornsilk4)
//  use (col.Ivory)
//  use (col.Ivory1)
//  use (col.Ivory2)
//  use (col.Ivory3)
//  use (col.Ivory4)
//  use (col.LemonChiffon)
//  use (col.LemonChiffon1)
//  use (col.LemonChiffon2)
//  use (col.LemonChiffon3)
//  use (col.LemonChiffon4)
//  use (col.Seashell)
//  use (col.Seashell1)
//  use (col.Seashell2)
//  use (col.Seashell3)
//  use (col.Seashell4)
//  use (col.Honeydew)
//  use (col.Honeydew1)
//  use (col.Honeydew2)
//  use (col.Honeydew3)
//  use (col.Honeydew4)
//  use (col.MintCream)
//  use (col.Azure)
//  use (col.Azure1)
//  use (col.Azure2)
//  use (col.Azure3)
//  use (col.Azure4)
//  use (col.AliceBlue)
//  use (col.Lavender)
//  use (col.LavenderBlush)
//  use (col.LavenderBlush1)
//  use (col.LavenderBlush2)
//  use (col.LavenderBlush3)
//  use (col.LavenderBlush4)
//  use (col.MistyRose)
//  use (col.MistyRose1)
//  use (col.MistyRose2)
//  use (col.MistyRose3)
//  use (col.MistyRose4)
//  use (col.DarkSlateGray)
//  use (col.DarkSlateGray1)
//  use (col.DarkSlateGray2)
//  use (col.DarkSlateGray3)
//  use (col.DarkSlateGray4)
//  use (col.DimGray)
//  use (col.SlateGray)
//  use (col.SlateGray1)
//  use (col.SlateGray2)
//  use (col.SlateGray3)
//  use (col.SlateGray4)
//  use (col.LightSlateGray)
//  use (col.Gray1)
//  use (col.LightGray)
//  use (col.DarkGray)
//  use (col.MidnightBlue)
//  use (col.NavyBlue)
//  use (col.CornflowerBlue)
//  use (col.DarkSlateBlue)
//  use (col.SlateBlue)
//  use (col.SlateBlue1)
//  use (col.SlateBlue2)
//  use (col.SlateBlue3)
//  use (col.SlateBlue4)
//  use (col.MediumSlateBlue)
//  use (col.LightSlateBlue)
//  use (col.MediumBlue)
//  use (col.RoyalBlue)
//  use (col.RoyalBlue1)
//  use (col.RoyalBlue2)
//  use (col.RoyalBlue3)
//  use (col.RoyalBlue4)
//  use (col.Blue)
//  use (col.Blue1)
//  use (col.Blue2)
//  use (col.Blue3)
//  use (col.Blue4)
//  use (col.DarkBlue)
//  use (col.DodgerBlue)
//  use (col.DodgerBlue2)
//  use (col.DodgerBlue3)
//  use (col.DodgerBlue4)
//  use (col.DeepSkyBlue)
//  use (col.DeepSkyBlue2)
//  use (col.DeepSkyBlue3)
//  use (col.DeepSkyBlue4)
//  use (col.SkyBlue)
//  use (col.SkyBlue2)
//  use (col.SkyBlue3)
//  use (col.SkyBlue4)
//  use (col.LightSkyBlue)
//  use (col.LightSkyBlue1)
//  use (col.LightSkyBlue2)
//  use (col.LightSkyBlue3)
//  use (col.LightSkyBlue4)
//  use (col.SteelBlue)
//  use (col.SteelBlue1)
//  use (col.SteelBlue2)
//  use (col.SteelBlue3)
//  use (col.SteelBlue4)
//  use (col.LightSteelBlue)
//  use (col.LightSteelBlue1)
//  use (col.LightSteelBlue2)
//  use (col.LightSteelBlue3)
//  use (col.LightSteelBlue4)
//  use (col.LightBlue)
//  use (col.LightBlue1)
//  use (col.LightBlue2)
//  use (col.LightBlue3)
//  use (col.LightBlue4)
//  use (col.PowderBlue)
//  use (col.PaleTurquoise)
//  use (col.PaleTurquoise1)
//  use (col.PaleTurquoise2)
//  use (col.PaleTurquoise3)
//  use (col.PaleTurquoise4)
//  use (col.DarkTurquoise)
//  use (col.MediumTurquoise)
//  use (col.Turquoise)
//  use (col.Turquoise1)
//  use (col.Turquoise2)
//  use (col.Turquoise3)
//  use (col.Turquoise4)
//  use (col.Cyan)
//  use (col.Cyan1)
//  use (col.Cyan2)
//  use (col.Cyan3)
//  use (col.Cyan4)
//  use (col.DarkCyan)
//  use (col.LightCyan)
//  use (col.LightCyan2)
//  use (col.LightCyan3)
//  use (col.LightCyan4)
//  use (col.CadetBlue)
//  use (col.CadetBlue1)
//  use (col.CadetBlue2)
//  use (col.CadetBlue3)
//  use (col.CadetBlue4)
//  use (col.MediumAquamarine)
//  use (col.Aquamarine)
//  use (col.Aquamarine1)
//  use (col.Aquamarine2)
//  use (col.Aquamarine3)
//  use (col.Aquamarine4)
//  use (col.DarkGreen)
//  use (col.DarkOliveGreen)
//  use (col.DarkOliveGreen1)
//  use (col.DarkOliveGreen2)
//  use (col.DarkOliveGreen3)
//  use (col.DarkOliveGreen4)
//  use (col.DarkSeaGreen)
//  use (col.DarkSeaGreen1)
//  use (col.DarkSeaGreen2)
//  use (col.DarkSeaGreen3)
//  use (col.DarkSeaGreen4)
//  use (col.SeaGreen)
//  use (col.SeaGreen1)
//  use (col.SeaGreen2)
//  use (col.SeaGreen3)
//  use (col.SeaGreen4)
//  use (col.MediumSeaGreen)
//  use (col.LightSeaGreen)
//  use (col.LightGreen)
//  use (col.PaleGreen)
//  use (col.PaleGreen1)
//  use (col.PaleGreen2)
//  use (col.PaleGreen3)
//  use (col.PaleGreen4)
//  use (col.SpringGreen)
//  use (col.SpringGreen2)
//  use (col.SpringGreen3)
//  use (col.SpringGreen4)
//  use (col.LawnGreen)
//  use (col.Green1)
//  use (col.Green2)
//  use (col.Green3)
//  use (col.Green4)
//  use (col.Chartreuse)
//  use (col.Chartreuse2)
//  use (col.Chartreuse3)
//  use (col.Chartreuse4)
//  use (col.MediumSpringGreen)
//  use (col.GreenYellow)
//  use (col.LimeGreen)
//  use (col.YellowGreen)
//  use (col.ForestGreen)
//  use (col.OliveDrab)
//  use (col.OliveDrab1)
//  use (col.OliveDrab2)
//  use (col.OliveDrab3)
//  use (col.OliveDrab4)
//  use (col.DarkKhaki)
//  use (col.Khaki)
//  use (col.Khaki1)
//  use (col.Khaki2)
//  use (col.Khaki3)
//  use (col.Khaki4)
//  use (col.PaleGoldenrod)
//  use (col.LightGoldenrodYellow)
//  use (col.LightGoldenrod1)
//  use (col.LightGoldenrod2)
//  use (col.LightGoldenrod3)
//  use (col.LightGoldenrod4)
//  use (col.LightYellow)
//  use (col.LightYellow2)
//  use (col.LightYellow3)
//  use (col.LightYellow4)
//  use (col.Yellow1)
//  use (col.Yellow2)
//  use (col.Yellow3)
//  use (col.Yellow4)
//  use (col.Gold)
//  use (col.Gold1)
//  use (col.Gold2)
//  use (col.Gold3)
//  use (col.Gold4)
//  use (col.LightGoldenrod)
//  use (col.Goldenrod)
//  use (col.Goldenrod1)
//  use (col.Goldenrod2)
//  use (col.Goldenrod3)
//  use (col.Goldenrod4)
//  use (col.DarkGoldenrod)
//  use (col.DarkGoldenrod1)
//  use (col.DarkGoldenrod2)
//  use (col.DarkGoldenrod3)
//  use (col.DarkGoldenrod4)
//  use (col.RosyBrown)
//  use (col.RosyBrown1)
//  use (col.RosyBrown2)
//  use (col.RosyBrown3)
//  use (col.RosyBrown4)
//  use (col.IndianRed)
//  use (col.IndianRed1)
//  use (col.IndianRed2)
//  use (col.IndianRed3)
//  use (col.IndianRed4)
//  use (col.SaddleBrown)
//  use (col.Sienna)
//  use (col.Sienna1)
//  use (col.Sienna2)
//  use (col.Sienna3)
//  use (col.Sienna4)
//  use (col.Peru)
//  use (col.Burlywood)
//  use (col.Burlywood1)
//  use (col.Burlywood2)
//  use (col.Burlywood3)
//  use (col.Burlywood4)
//  use (col.Beige1)
//  use (col.Wheat)
//  use (col.Wheat1)
//  use (col.Wheat2)
//  use (col.Wheat3)
//  use (col.Wheat4)
//  use (col.SandyBrown)
//  use (col.Tan)
//  use (col.Tan1)
//  use (col.Tan2)
//  use (col.Tan3)
//  use (col.Tan4)
//  use (col.Chocolate)
//  use (col.Chocolate1)
//  use (col.Chocolate2)
//  use (col.Chocolate3)
//  use (col.Chocolate4)
//  use (col.Firebrick)
//  use (col.Firebrick1)
//  use (col.Firebrick2)
//  use (col.Firebrick3)
//  use (col.Firebrick4)
//  use (col.Brown0)
//  use (col.Brown1)
//  use (col.Brown2)
//  use (col.Brown3)
//  use (col.Brown4)
//  use (col.DarkSalmon)
//  use (col.Salmon)
//  use (col.Salmon1)
//  use (col.Salmon2)
//  use (col.Salmon3)
//  use (col.Salmon4)
//  use (col.LightSalmon)
//  use (col.LightSalmon1)
//  use (col.LightSalmon2)
//  use (col.LightSalmon3)
//  use (col.LightSalmon4)
//  use (col.Orange1)
//  use (col.Orange2)
//  use (col.Orange3)
//  use (col.Orange4)
//  use (col.DarkOrange)
//  use (col.DarkOrange1)
//  use (col.DarkOrange2)
//  use (col.DarkOrange3)
//  use (col.DarkOrange4)
//  use (col.Coral)
//  use (col.Coral1)
//  use (col.Coral2)
//  use (col.Coral3)
//  use (col.Coral4)
//  use (col.LightCoral)
//  use (col.Tomato)
//  use (col.Tomato2)
//  use (col.Tomato3)
//  use (col.Tomato4)
//  use (col.OrangeRed)
//  use (col.OrangeRed1)
//  use (col.OrangeRed2)
//  use (col.OrangeRed3)
//  use (col.OrangeRed4)
//  use (col.Red)
//  use (col.Red1)
//  use (col.Red2)
//  use (col.Red3)
//  use (col.Red4)
//  use (col.DarkRed)
//  use (col.HotPink)
//  use (col.HotPink1)
//  use (col.HotPink2)
//  use (col.HotPink3)
//  use (col.HotPink4)
//  use (col.DeepPink)
//  use (col.DeepPink2)
//  use (col.DeepPink3)
//  use (col.DeepPink4)
//  use (col.Pink0)
//  use (col.Pink1)
//  use (col.Pink2)
//  use (col.Pink3)
//  use (col.Pink4)
//  use (col.LightPink)
//  use (col.LightPink1)
//  use (col.LightPink2)
//  use (col.LightPink3)
//  use (col.LightPink4)
//  use (col.PaleVioletRed)
//  use (col.PaleVioletRed1)
//  use (col.PaleVioletRed2)
//  use (col.PaleVioletRed3)
//  use (col.PaleVioletRed4)
//  use (col.Maroon)
//  use (col.Maroon1)
//  use (col.Maroon2)
//  use (col.Maroon3)
//  use (col.Maroon4)
//  use (col.MediumVioletRed)
//  use (col.VioletRed)
//  use (col.VioletRed1)
//  use (col.VioletRed2)
//  use (col.VioletRed3)
//  use (col.VioletRed4)
//  use (col.Magenta0)
//  use (col.Magenta1)
//  use (col.Magenta2)
//  use (col.Magenta3)
//  use (col.Magenta4)
//  use (col.DarkMagenta)
//  use (col.Violet)
//  use (col.Plum)
//  use (col.Plum1)
//  use (col.Plum2)
//  use (col.Plum3)
//  use (col.Plum4)
//  use (col.Orchid)
//  use (col.Orchid1)
//  use (col.Orchid2)
//  use (col.Orchid3)
//  use (col.Orchid4)
//  use (col.MediumOrchid)
//  use (col.MediumOrchid1)
//  use (col.MediumOrchid2)
//  use (col.MediumOrchid3)
//  use (col.MediumOrchid4)
//  use (col.DarkOrchid)
//  use (col.DarkOrchid1)
//  use (col.DarkOrchid2)
//  use (col.DarkOrchid3)
//  use (col.DarkOrchid4)
//  use (col.DarkViolet)
//  use (col.BlueViolet)
//  use (col.Purple)
//  use (col.Purple1)
//  use (col.Purple2)
//  use (col.Purple3)
//  use (col.Purple4)
//  use (col.MediumPurple)
//  use (col.MediumPurple1)
//  use (col.MediumPurple2)
//  use (col.MediumPurple3)
//  use (col.MediumPurple4)
//  use (col.Thistle)
//  use (col.Thistle1)
//  use (col.Thistle2)
//  use (col.Thistle3)
//  use (col.Thistle4)

  max = uint(n)
}