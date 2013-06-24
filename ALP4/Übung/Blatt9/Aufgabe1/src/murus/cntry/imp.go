package cntry

// (c) Christian Maurer   v. 130309 - license see murus.go

import (
  . "murus/obj"; "murus/str"
  "murus/kbd"
  "murus/col"; "murus/scr"; "murus/box"; "murus/errh"; "murus/nat"
  "murus/font"; "murus/pbox"
  "murus/sel"
)
const (
  length = 22
  pack = "cntry"
)
type
  country byte; const (keineAngabe = iota
/* Europa */  Albanien; Andorra; Belgien; BosnienHerzegowina; Bulgarien
              Dänemark; Deutschland; Estland; Finnland; Frankreich
              Griechenland; Großbritannien; Irland; Island; Italien; Kroatien
              Lettland; Liechtenstein; Litauen; Luxemburg; Malta; Mazedonien
              Moldau; Monaco; Montenegro; Niederlande; Norwegen; Österreich
              Polen; Portugal; Rumänien; Russland; SanMarino; Schweden
              Schweiz; Serbien; Slowakei; Slowenien; Spanien; Tschechien
              Türkei; Ukraine; Ungarn; Vatikan; Weißrussland; Zypern
/* Afrika */  Ägypten; ÄquatorialGuinea; Äthiopien; Algerien; Angola; Benin
              Botsuana; BurkinaFaso; Burundi; Dschibuti; Elfenbeinküste
              Eritrea; Gabun; Gambia; Ghana; Guinea; GuineaBissau; Kamerun
              KapVerde; Kenia; Komoren; Kongo; KongoDemRep; Lesotho; Liberia
              Libyen; Madagaskar; Malawi; Mali; Marokko; Mauretanien; Mauritius
              Mosambik; Namibia; Niger; Nigeria; Ruanda; Sambia; SanTomePrincipe
              Senegal; Seychellen; SierraLeone; Simbabwe; Somalia; Südafrika
              Südsudan; Sudan; Swasiland; Tansania; Togo; Tschad; Tunesien
              Uganda; Zentralafrika
/* Amerika */ Antigua; Argentinien; Bahamas; Barbados; Belize; Bolivien
              Brasilien; Chile; CostaRica; Dominica; DominikanRep; Ecuador
              ElSalvador; Grenada; Guatemala; Guyana; Haiti; Honduras; Jamaika
              Kanada; StKittsNevis; Kolumbien; Kuba; StLucia; Mexiko; Nikaragua
              Panama; Paraguay; Peru; Suriname; TrinidadTobago; Uruguay; USA
              Venezuela; StVincent
/* Asien */   Afghanistan; Armenien; Aserbaidschan; Bahrain; Bangladesch
              Bhutan; Brunei; China; Georgien; Indien; Indonesien; Irak; Iran
              Israel; Japan; Jemen; Jordanien; Kambodscha; Kasachstan; Katar
              Kirgisistan; Kuwait; Laos; Libanon; Malaysia; Malediven; Mongolei
              Myanmar; Nepal; Nordkorea; Oman; Osttimor; Pakistan; Palästina
              Philippinen; SaudiArabien; Singapur; SriLanka; Südkorea; Syrien
              Taiwan; Thailand; Tadschikistan; Turkmenistan; Usbekistan
              VerArabEmirate; Vietnam
/* Australien und Ozeanien */
              Australien; Cookinseln; Fidschi; Kiribati; Marshallinseln
              Mikronesien; Nauru; Neuseeland; Niue; Palau; PapuaNeuguinea
              Salomonen; Samoa; Tonga; Tuvalu; Vanuatu
              noNations)
type (
  attribut struct {
              iso,       // len 4
              tld,       // len 3
             name string
           prefix uint16
              kfz,
              ioc string // len 3
                  }
  Imp struct {
         cnt country
         att attribut
         fmt Format
      cF, cB col.Colour
          fo font.Font
            }
  Codes [2]byte // tld without trailing 0
)
var (
  bx *box.Imp = box.New ()
  Font uint
  pbx *pbox.Imp = pbox.New ()
  Folge []attribut
)


func (x *Imp) imp (Y Object) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func New () *Imp {
//
  x:= new (Imp)
  x.Clr ()
  x.fmt = Long
  x.cF, x.cB = col.ScreenF, col.ScreenB
  return x
}


func (x *Imp) InContinent (c Continent) bool {
//
  switch c { case Europa:
    return Albanien <= x.cnt && x.cnt <= Zypern
  case Afrika:
    return Ägypten <= x.cnt && x.cnt <= Zentralafrika
  case Amerika:
    return Antigua <= x.cnt && x.cnt <= StVincent
  case Asien:
    return Afghanistan <= x.cnt && x.cnt <= Vietnam
  case Ozeanien:
    return Australien <= x.cnt && x.cnt <= Vanuatu
  }
  return false
}


func (x *Imp) Empty () bool {
//
  return x.cnt == keineAngabe
}


func (x *Imp) Clr () {
//
  x.cnt = keineAngabe
  x.att.iso = "   "
  x.att.tld = "  "
  switch x.fmt { case Tld:
    x.att.name = str.Clr (2)
  case Long:
    x.att.name = str.Clr (length)
  }
  x.att.prefix = 0
  x.att.kfz = "   "
  x.att.ioc = "   "
}


func (x *Imp) Clone () Object {
//
  y:= New()
  y.Copy (x)
  return y
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.cnt = y.cnt
  x.att.iso = y.att.iso
  x.att.tld = y.att.tld
  x.att.name = y.att.name
  x.att.prefix = y.att.prefix
  x.att.kfz = y.att.kfz
  x.att.ioc = y.att.ioc
  x.fmt = y.fmt
  x.cF = y.cB
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.cnt == x.imp (Y).cnt
}


func (x *Imp) Less (Y Object) bool {
//
  return str.Less (x.att.name, x.imp (Y).att.name)
}


func (x *Imp) String () string {
//
  return x.att.name
}


func (x *Imp) Defined (t string) bool {
//
  str.RemSpaces (&t)
  if str.Empty (t) {
    x.Clr ()
    return true
  }
  for c:= country(keineAngabe); c < noNations; c++ {
    x.cnt = c
    x.att = Folge [c]
    switch x.fmt { case Tld:
      if t == x.att.tld { return true }
    case Long:
      var p uint
      if str.IsPart (t, x.att.name, &p) && p == 0 { return true }
    default:
      return false
    }
  }
  return false
}


func (x *Imp) SetColours (S, H col.Colour) {
//
  x.cF, x.cB = S, H
}


func (x *Imp) SetFormat (f Format) {
//
  if f < NFormats {
    x.fmt = f
    switch x.fmt { case Tld:
      bx.Wd (2)
    case Long:
      bx.Wd (length)
    default:
      bx.Wd (3)
    }
  }
}


func (x *Imp) Write (Z, S uint) {
//
  bx.Colours (x.cF, x.cB)
  switch x.fmt { case Tld:
    bx.Write (x.att.tld, Z, S)
  case Long:
    bx.Clr (Z, S)
    bx.Write (x.att.name, Z, S)
  case Tel:
    nat.SetColours (x.cF, x.cB)
    nat.Write (uint(x.att.prefix), Z, S + 3)
  case Car:
    bx.Write (x.att.kfz, Z, S)
  case Ioc:
    bx.Write (x.att.ioc, Z, S)
  }
}


func (x *Imp) Edit (Z, S uint) {
//
  bx.Colours (x.cF, x.cB)
  var l uint
  for {
    switch x.fmt { case Tld:
      l = 2
      bx.Edit (&x.att.tld, Z, S)
    case Long:
      l = length
      bx.Edit (&x.att.name, Z, S)
    default:
      return
    }
    var T uint
    if kbd.LastCommand (&T) == kbd.LookFor {
      n:= uint(x.cnt)
      sel.Select (func (P, Z, S uint, V, H col.Colour) { x.att = Folge [country(P)]; x.SetColours (V, H); x.Write (Z, S) },
                    noNations, scr.NLines(), l, &n, Z, S, x.cB, x.cF)
      if country(n) == noNations {
        n = uint(x.cnt) // Nation unverändert
      } else {
        x.cnt = country(n)
      }
      x.att = Folge [x.cnt]
      break
    } else {
      var ok bool
      switch x.fmt { case Tld:
        ok = x.Defined (x.att.tld)
      case Long:
        ok = x.Defined (x.att.name)
      }
      if ok {
        break
      } else {
        errh.Error ("kein Land", 0)
      }
    }
  }
  x.Write (Z, S)
}


func (x *Imp) SetFont (f font.Font) {
//
  x.fo = f
}


func (x *Imp) Print (l, c uint) {
//
  pbx.SetFont (x.fo)
  switch x.fmt { case Tld:
    pbx.Print (x.att.tld, l, c)
  case Long:
    pbx.Print (x.att.name, l, c)
  }
}


func (x *Imp) Codelen () uint {
//
  return 2
}


func (x *Imp) Encode () []byte {
//
  b:= make ([]byte, x.Codelen())
  b[0] = x.att.tld [0]
  b[1] = x.att.tld [1]
  return b
}


func (x *Imp) Decode (b []byte) {
//
  t:= string(b)
  if ! x.Defined (t) {
    x.Clr ()
  }
}


func def (n country, N string, I string, D string, V uint16, K, O string) {
//
  Folge [n].tld = D
  str.Set (&Folge [n].name, N)
  Folge [n].iso = I
  Folge [n].prefix = V
  Folge [n].kfz = K
  Folge [n].ioc = O
}


func init () {
//
  bx.Wd (length)
  Folge = make ([]attribut, noNations)
                                                      // iso   tld   Tel    kfz    ioc
  def (keineAngabe,        "                      ", "   ", "  ", 0   , "   ", "   ")

  def (Afghanistan,        "Afghanistan",            "AFG", "af", 93  , "AFG", "AFG")
  def (Ägypten,            "Ägypten",                "EGY", "eg", 20  , "ET ", "EGY")
  def (Albanien,           "Albanien",               "ALB", "al", 355 , "AL ", "ALB")
  def (Algerien,           "Algerien",               "DZA", "dz", 213 , "DZ ", "ALG")
  def (Andorra,            "Andorra",                "AND", "ad", 376 , "AND", "AND")
  def (Angola,             "Angola",                 "AGO", "ao", 244 , "ANG", "ANG")
  def (Antigua,            "Antigua und Barbuda",    "ATG", "ag", 1268, "AG ", "ANT")
  def (ÄquatorialGuinea,   "Äquatorial-Guinea",      "GNQ", "gq", 240 , "   ", "GEQ")
  def (Argentinien,        "Argentinien",            "ARG", "ar", 54  , "RA ", "ARG")
  def (Armenien,           "Armenien",               "ARM", "am", 374 , "AR ", "ARM")
  def (Aserbaidschan,      "Aserbaidschan",          "AZE", "az", 994 , "AZ ", "AZE")
  def (Äthiopien,          "Äthiopien",              "ETH", "et", 251 , "ETH", "ETH")
  def (Australien,         "Australien",             "AUS", "au", 61  , "AUS", "AUS")

  def (Bahamas,            "Bahamas",                "BHS", "bs", 1242, "BS ", "BAH")
  def (Bahrain,            "Bahrain",                "BHR", "bh", 973 , "BRN", "BRN")
  def (Bangladesch,        "Bangladesch",            "BGD", "bd", 880 , "BD ", "BAN")
  def (Barbados,           "Barbados",               "BRB", "bb", 1246, "BDS", "BAR")
  def (Belgien,            "Belgien",                "BEL", "be", 32  , "B  ", "BEL")
  def (Belize,             "Belize",                 "BLZ", "bz", 501 , "BZ ", "BIZ")
  def (Benin,              "Benin",                  "BEN", "bj", 229 , "DY ", "BEN")
  def (Bhutan,             "Bhutan",                 "BTN", "bt", 975 , "BHT", "BHU")
  def (Bolivien,           "Bolivien",               "BOL", "bo", 591 , "BOL", "BOL")
  def (BosnienHerzegowina, "Bosnien u. Herzegowina", "BIH", "ba", 387 , "BIH", "BIH")
  def (Botsuana,           "Botsuana",               "BWA", "bw", 267 , "RB ", "BOT")
  def (Brasilien,          "Brasilien",              "BRA", "br", 55  , "BR ", "BRA")
  def (Brunei,             "Brunei Darussalam",      "BRN", "bn", 673 , "BRU", "BRU")
  def (Bulgarien,          "Bulgarien",              "BGR", "bg", 359 , "BG ", "BUL")
  def (BurkinaFaso,        "Burkina Faso",           "BFA", "bf", 226 , "BF ", "BUR")
  def (Burundi,            "Burundi",                "BDI", "bi", 257 , "RU ", "BDI")

  def (Chile,              "Chile",                  "CHL", "cl", 56  , "RCH", "CHI")
  def (China,              "China",                  "CHN", "cn", 86  , "RC ", "CHN")
  def (Cookinseln,         "Cookinseln",             "COK", "ck", 682 , "NZ ", "COK")
  def (CostaRica,          "Costa Rica",             "CRI", "cr", 506 , "CR ", "CRC")

  def (Dänemark,           "Dänemark",               "DNK", "dk", 45  , "DK ", "DEN")
  def (Deutschland,        "Deutschland",            "DEU", "de", 49  , "D  ", "GER")
  def (Dominica,           "Dominica",               "DMA", "dm", 1767, "WD ", "DMA")
  def (DominikanRep,       "Dominikan. Republik",    "DOM", "do", 1809, "DOM", "DOM")
  def (Dschibuti,          "Dschibuti",              "DJI", "dj", 253 , "DSC", "DJI")

  def (Ecuador,            "Ecuador",                "ECU", "ec", 593 , "EC ", "ECU")
  def (ElSalvador,         "El Salvador",            "SLV", "sv", 503 , "ES ", "ESA")
  def (Elfenbeinküste,     "Elfenbeinküste",         "CIV", "ci", 225 , "CI ", "CIV")
  def (Eritrea,            "Eritrea",                "ERI", "er", 291 , "ER ", "ERI")
  def (Estland,            "Estland",                "EST", "ee", 372 , "EST", "EST")

  def (Fidschi,            "Fidschi",                "FIJ", "fj", 679 , "FJI", "FIJ")
  def (Finnland,           "Finnland",               "FIN", "fi", 358 , "FIN", "FIN")
  def (Frankreich,         "Frankreich",             "FRA", "fr", 33  , "F  ", "FRA")

  def (Gabun,              "Gabun",                  "GAB", "ga", 241 , "G  ", "GAB")
  def (Gambia,             "Gambia",                 "GMB", "gm", 220 , "WAG", "GAM")
  def (Georgien,           "Georgien",               "GEO", "ge", 995 , "GE ", "GEO")
  def (Ghana,              "Ghana",                  "GHA", "gh", 233 , "GH ", "GHA")
  def (Grenada,            "Grenada",                "GRD", "gd", 1473, "WG ", "GRN")
  def (Griechenland,       "Griechenland",           "GRC", "gr", 30  , "GR ", "GRE")
  def (Großbritannien,     "Großbritannien",         "GBR", "uk", 44  , "GB ", "GBR") // auch gb
  def (Guatemala,          "Guatemala",              "GTM", "gt", 502 , "GCA", "GUA")
  def (Guinea,             "Guinea",                 "GIN", "gn", 224 , "RG ", "GUI")
  def (GuineaBissau,       "Guinea-Bissau",          "GNB", "gw", 245 , "GNB", "GBS")
  def (Guyana,             "Guyana",                 "GUY", "gy", 592 , "GUY", "GUY")

  def (Haiti,              "Haiti",                  "HTI", "ht", 509 , "RH ", "HAI")
  def (Honduras,           "Honduras",               "HND", "hn", 504 , "HN ", "HON")

  def (Indien,             "Indien",                 "IND", "in", 91  , "IND", "IND")
  def (Indonesien,         "Indonesien",             "IDN", "id", 62  , "RI ", "INA")
  def (Irak,               "Irak",                   "IRQ", "iq", 964 , "IRQ", "IRQ")
  def (Iran,               "Iran",                   "IRN", "ir", 98  , "IR ", "IRI")
  def (Irland,             "Irland",                 "IRL", "ie", 353 , "IRL", "IRL")
  def (Island,             "Island",                 "ISL", "is", 354 , "IS ", "ISL")
  def (Israel,             "Israel",                 "ISR", "il", 972 , "IL ", "ISR")
  def (Italien,            "Italien",                "ITA", "it", 39  , "I  ", "ITA")

  def (Jamaika,            "Jamaika",                "JAM", "jm", 1876, "JA ", "JAM")
  def (Japan,              "Japan",                  "JPN", "jp", 81  , "J  ", "JPN")
  def (Jemen,              "Jemen",                  "YEM", "ye", 967 , "YAR", "YEM")
  def (Jordanien,          "Jordanien",              "JOR", "jo", 962 , "JOR", "JOR")

  def (Kambodscha,         "Kambodscha",             "KHM", "kh", 855 , "K  ", "CAM")
  def (Kamerun,            "Kamerun",                "CMR", "cm", 237 , "TC ", "CMR")
  def (Kanada,             "Kanada",                 "CAN", "ca", 1   , "CDN", "CAN")
  def (KapVerde,           "Kap Verde",              "CPV", "cv", 238 , "CV ", "CPV")
  def (Kasachstan,         "Kasachstan",             "KAZ", "kz", 7   , "KZ ", "KAZ")
  def (Katar,              "Katar",                  "QAT", "qa", 974 , "Q  ", "QAT")
  def (Kenia,              "Kenia",                  "KEN", "ke", 254 , "EAK", "KEN")
  def (Kirgisistan,        "Kirgisistan",            "KGZ", "kg", 996 , "KS ", "KGZ")
  def (Kiribati,           "Kiribati",               "KIR", "ki", 686 , "KI ", "KIR")
  def (Kolumbien,          "Kolumbien",              "COL", "co", 57  , "CO ", "COL")
  def (Komoren,            "Komoren",                "COM", "km", 269 , "COM", "COM")
  def (Kongo,              "Kongo",                  "COG", "cg", 242 , "RCB", "CGO")
  def (KongoDemRep,        "Kongo, Dem.Rep.",        "COD", "cd", 243 , "CD ", "COD")
//def (Kosovo,             "Kosovo",                 "XXK", "  ", 381 , "   ", "   ")
  def (Kroatien,           "Kroatien",               "HRV", "hr", 385 , "HR ", "CRO")
  def (Kuba,               "Kuba",                   "CUB", "cu", 53  , "C  ", "CUB")
  def (Kuwait,             "Kuwait",                 "KWT", "kw", 965 , "KWT", "KUW")

  def (Laos,               "Laos",                   "LAO", "la", 856 , "LAO", "LAO")
  def (Lesotho,            "Lesotho",                "LSO", "ls", 266 , "LS ", "LES")
  def (Lettland,           "Lettland",               "LVA", "lv", 371 , "LV ", "LAT")
  def (Libanon,            "Libanon",                "LBN", "lb", 961 , "RL ", "LIB")
  def (Liberia,            "Liberia",                "LBR", "lr", 231 , "LB ", "LBR")
  def (Libyen,             "Libyen",                 "LBY", "ly", 218 , "LAR", "LBA")
  def (Liechtenstein,      "Liechtenstein",          "LIE", "li", 423 , "FL ", "LIE")
  def (Litauen,            "Litauen",                "LTU", "lt", 370 , "LT ", "LTU")
  def (Luxemburg,          "Luxemburg",              "LUX", "lu", 352 , "L  ", "LUX")

  def (Madagaskar,         "Madagaskar",             "MDG", "mg", 261 , "RM ", "MAD")
  def (Malawi,             "Malawi",                 "MWI", "mw", 265 , "MW ", "MAW")
  def (Malaysia,           "Malaysia",               "MYS", "my", 60  , "MAL", "MAS")
  def (Malediven,          "Malediven",              "MDV", "mv", 960 , "MV ", "MDV")
  def (Mali,               "Mali",                   "MLI", "ml", 223 , "RMM", "MLI")
  def (Malta,              "Malta",                  "MLT", "mt", 356 , "M  ", "MLT")
  def (Marokko,            "Marokko",                "MAR", "ma", 212 , "MA ", "MAR")
  def (Marshallinseln,     "Marshallinseln",         "MHL", "mh", 692 , "MH ", "MHL")
  def (Mauretanien,        "Mauretanien",            "MRT", "mr", 222 , "RIM", "MTN")
  def (Mauritius,          "Mauritius",              "MUS", "mu", 230 , "MS ", "MRI")
  def (Mazedonien,         "Mazedonien",             "MKD", "mk", 389 , "MK ", "MKD")
  def (Mexiko,             "Mexiko",                 "MEX", "mx", 52  , "MEX", "MEX")
  def (Mikronesien,        "Mikronesien",            "FSM", "fm", 691 , "FSM", "FSM")
  def (Moldau,             "Moldau",                 "MDA", "md", 373 , "MD ", "MDA")
  def (Monaco,             "Monaco",                 "MCO", "mc", 377 , "MC ", "MON")
  def (Mongolei,           "Mongolei",               "MNG", "mn", 976 , "MGL", "MGL")
  def (Montenegro,         "Montenegro",             "MNE", "me", 382 , "MNE", "MNE") // auch yu
  def (Mosambik,           "Mosambik",               "MOZ", "mz", 258 , "MOC", "MOZ")
  def (Myanmar,            "Myanmar",                "MMR", "mm", 95  , "MYA", "MYA")

  def (Namibia,            "Namibia",                "NAM", "na", 264 , "NAM", "NAM")
  def (Nauru,              "Nauru",                  "NRU", "nr", 674 , "NAU", "NRU")
  def (Nepal,              "Nepal",                  "NPL", "np", 977 , "NEP", "NEP")
  def (Neuseeland,         "Neuseeland",             "NZL", "nz", 64  , "NZ ", "NZL")
  def (Nikaragua,          "Nicaragua",              "NIC", "ni", 505 , "NIC", "NCA")
  def (Niederlande,        "Niederlande",            "NLD", "nl", 31  , "NL ", "NED")
  def (Niger,              "Niger",                  "NER", "ne", 227 , "NIG", "NIG")
  def (Nigeria,            "Nigeria",                "NGA", "ng", 234 , "WAN", "NGR")
  def (Niue,               "Niue",                   "NIU", "nu", 683 , "NZ ", "   ")
  def (Nordkorea,          "Nordkorea",              "PRK", "kp", 850 , "KP ", "PRK")
  def (Norwegen,           "Norwegen",               "NOR", "no", 47  , "N  ", "NOR")

  def (Oman,               "Oman",                   "OMN", "om", 968 , "OM ", "OMA")
  def (Österreich,         "Österreich",             "AUT", "at", 43  , "A  ", "AUT")
  def (Osttimor,           "Timor-Leste",            "TLS", "tl", 670 , "TL ", "TLS") // auch tp

  def (Pakistan,           "Pakistan",               "PAK", "pk", 92  , "PK ", "PAK")
  def (Palästina,          "Palästin.Autonomiegeb.", "PSE", "ps", 970 , "   ", "PLE")
  def (Palau,              "Palau",                  "PLW", "pw", 680 , "PAL", "PLW")
  def (Panama,             "Panama",                 "PAN", "pa", 507 , "PA ", "PAN")
  def (PapuaNeuguinea,     "Papua Neuguinea",        "PNG", "pg", 675 , "PNG", "PNG")
  def (Paraguay,           "Paraguay",               "PRY", "py", 595 , "PY ", "PAR")
  def (Peru,               "Peru",                   "PER", "pe", 51  , "PE ", "PER")
  def (Philippinen,        "Philippinen",            "PHL", "ph", 63  , "RP ", "PHI")
  def (Polen,              "Polen",                  "POL", "pl", 48  , "PL ", "POL")
  def (Portugal,           "Portugal",               "PRT", "pt", 351 , "P  ", "POR")

  def (Ruanda,             "Ruanda",                 "RWA", "rw", 250 , "RWA", "RWA")
  def (Rumänien,           "Rumänien",               "ROU", "ro", 40  , "R  ", "ROU")
  def (Russland,           "Russische Föderation",   "RUS", "ru", 7   , "RUS", "RUS") // auch su

  def (Salomonen,          "Salomonen",              "SLB", "sb", 677 , "SOL", "SOL")
  def (Sambia,             "Sambia",                 "ZMB", "zm", 260 , "Z  ", "ZAM")
  def (Samoa,              "Samoa",                  "WSM", "ws", 685 , "WS ", "SAM")
  def (SanMarino,          "San Marino",             "SMR", "sm", 378 , "RSM", "SMR")
  def (SanTomePrincipe,    "Sao Tome und Principe",  "STP", "st", 239 , "STP", "STP")
  def (SaudiArabien,       "Saudi-Arabien",          "SAU", "sa", 966 , "KSA", "KSA")
  def (Schweden,           "Schweden",               "SWE", "se", 46  , "S  ", "SWE")
  def (Schweiz,            "Schweiz",                "CHE", "ch", 41  , "CH ", "SUI")
  def (Senegal,            "Senegal",                "SEN", "sn", 221 , "SN ", "SEN")
  def (Serbien,            "Serbien",                "SRB", "rs", 381 , "SRB", "SRB") // auch yu
  def (Seychellen,         "Seychellen",             "SYC", "sc", 248 , "SY ", "SEY")
  def (SierraLeone,        "Sierra Leone",           "SLE", "sl", 232 , "WAL", "SLE")
  def (Simbabwe,           "Simbabwe",               "ZWE", "zw", 263 , "ZW ", "ZIM")
  def (Singapur,           "Singapur",               "SGP", "sg", 65  , "SGP", "SIN")
  def (Slowakei,           "Slowakei",               "SVK", "sk", 421 , "SK ", "SVK")
  def (Slowenien,          "Slowenien",              "SVN", "si", 386 , "SLO", "SLO")
  def (Somalia,            "Somalia",                "SOM", "so", 252 , "SP ", "SOM")
  def (Spanien,            "Spanien",                "ESP", "es", 34  , "E  ", "ESP")
  def (SriLanka,           "Sri Lanka",              "LKA", "lk", 93  , "CL ", "SRI")
  def (StKittsNevis,       "St. Kitts und Nevis",    "KNA", "kn", 1869, "KAN", "SKN")
  def (StLucia,            "St. Lucia",              "LCA", "lc", 1758, "WL ", "LCA")
  def (StVincent,          "St. Vincent Grenadinen", "VCT", "vc", 1784, "WV ", "VIN")
  def (Südafrika,          "Südafrika",              "ZAF", "za", 27  , "ZA ", "RSA")
  def (Sudan,              "Sudan",                  "SDN", "sd", 249 , "SUD", "SUD")
  def (Südsudan,           "Südsudan",               "   ", "ss", 292 , "SSD", "   ")
  def (Südkorea,           "Südkorea",               "KOR", "kr", 82  , "ROK", "KOR")
  def (Suriname,           "Suriname",               "SUR", "sr", 597 , "SME", "SUR")
  def (Swasiland,          "Swasiland",              "SWZ", "sz", 268 , "SD ", "SWZ")
  def (Syrien,             "Syrien",                 "SYR", "sy", 963 , "SYR", "SYR")

  def (Tadschikistan,      "Tadschikistan",          "TJK", "tj", 992 , "TJ ", "TJK")
  def (Taiwan,             "Taiwan",                 "TWN", "tw", 886 , "RC ", "TPE")
  def (Tansania,           "Tansania",               "TZA", "tz", 255 , "EAT", "TAN")
  def (Thailand,           "Thailand",               "THA", "th", 66  , "THA", "THA")
  def (Togo,               "Togo",                   "TGO", "tg", 228 , "RT ", "TOG")
  def (Tonga,              "Tonga",                  "TON", "to", 676 , "TON", "TGA")
  def (TrinidadTobago,     "Trinidad und Tobago",    "TTO", "tt", 1868, "TT ", "TRI")
  def (Tschad,             "Tschad",                 "TCD", "td", 235 , "TCD", "CHA")
  def (Tschechien,         "Tschechische Republik",  "CZE", "cz", 420 , "CZ ", "CZE")
  def (Tunesien,           "Tunesien",               "TUN", "tn", 216 , "TN ", "TUN")
  def (Türkei,             "Türkei",                 "TUR", "tr", 90  , "TR ", "TUR")
  def (Turkmenistan,       "Turkmenistan",           "TKM", "tm", 993 , "TM ", "TKM")
  def (Tuvalu,             "Tuvalu",                 "TUV", "tv", 688 , "TUV", "TUV")

  def (Uganda,             "Uganda",                 "UGA", "ug", 256 , "EAU", "UGA")
  def (Ukraine,            "Ukraine",                "UKR", "ua", 380 , "UA ", "UKR")
  def (Ungarn,             "Ungarn",                 "HUN", "hu", 36  , "H  ", "HUN")
  def (Uruguay,            "Uruguay",                "URY", "uy", 598 , "ROU", "URU")
  def (USA,                "Ver. Staaten v.Amerika", "USA", "us", 1   , "USA", "USA")
  def (Usbekistan,         "Usbekistan",             "UZB", "uz", 998 , "UZB", "UZB")

  def (Vanuatu,            "Vanuatu",                "VUT", "vu", 678 , "VU ", "VAN")
  def (Vatikan,            "Vatikanstadt",           "VAT", "va", 379 , "V  ", "   ")
  def (Venezuela,          "Venezuela",              "VEN", "ve", 58  , "YV ", "VEN")
  def (VerArabEmirate,     "Ver. Arabische Emirate", "ARE", "ae", 971 , "UAE", "UAE")
  def (Vietnam,            "Vietnam",                "VNM", "vn", 84  , "VN ", "VIE")

  def (Weißrussland,       "Weißrussland",           "BLR", "by", 375 , "BY ", "BLR")

  def (Zentralafrika,      "Zentralafrikan. Rep.",   "CAF", "cf", 236 , "RZA", "CAF")
  def (Zypern,             "Zypern",                 "CYP", "cy", 357 , "CY ", "CYP")
/*
                           "Anguilla",               "   ", "ai"
                           "Niederl. Antillen",      "   ", "an"
                           "Antarktis",              "   ", "aq"
                           "Amerikan. Samoa",        "   ", "as"
                           "Aruba",                  "   ", "aw"
                           "Åland",                  "   ", "ax"
                           "Bermuda",                "   ", "bm"
                           "Bouvet Island no",       "   ", "bv"
                           "Cocos (Keeling) Insel",  "   ", "cc"
                           "Christmas Insel",        "   ", "cx"
                           "Europäische Union",      "   ", "eu"
                           "Falkland Insel",         "   ", "fk"
                           "Faröer",                 "   ", "fo"
                           "Franz. Guiana",          "   ", "gf"
                           "Guernse",                "   ", "gg"
                           "Gibraltar",              "   ", "gi"
                           "Grönland",               "   ", "gl"
                           "Guadeloupe etc.",        "   ", "gp"
                      "SouthGeorgia+Sandwich Insel", "   ", "gs"
                           "Guam",                   "   ", "gu"
                           "HongKong",               "   ", "hk"
                           "Heard u.McDonald Insel", "   ", "hm"
                           "Isle of Man",            "   ", "im"
                           "Brit.Terr.im Ind.Ozean", "   ", "io"
                           "Jersey",                 "   ", "je"
                           "Cayman Islands",         "   ", "ky"
                           "Macau",                  "   ", "mo"
                           "Nord Mariana Insel",     "   ", "mp"
                           "Martinique",             "   ", "mq" // fr
                           "Montserrat",             "   ", "ms"
                           "Neu Kaledonien",         "   ", "nc" // fr
                           "Norfolk Island",         "   ", "nf"
                           "Franz. Polynesien",      "   ", "pf" // fr
                           "Saint-Pierre u.Miquelon","   ", "pm"
                           "Pitcairn Islands",       "   ", "pn"
                           "Puerto Rico",            "   ", "pr"
                           "Réunion",                "   ", "re" // fr
                           "Saint Helena",           "   ", "sh"
                           "Svalbard+JanMayen Insel","   ", "sj" // no
                           "Turks and Caicos Insel", "   ", "tc"
                           "Franz.südl.u.antarkt.L." "   ", "tf"
                           "Tokelau",                "   ", "tk"
                           "British Virgin Islands", "   ", "vg"
                           "Virgin Islands",         "   ", "vi"
                           "Wallis and Futuna",      "   ", "wf" // fr
                           "Mayotte",                "   ", "yt" // fr
*/

}
