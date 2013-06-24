package netz

// (c) Christian Maurer   v. 130309 - license see murus.go

import (
  "murus/ker"
  . "murus/obj"
  "murus/kbd"; "murus/str"
  "murus/col"; "murus/scr"
  "murus/nat"; "murus/errh"
  "murus/img"
  "murus/gra"
  . "murus/sub/linie"; "murus/sub/gleis"; "murus/sub/bahnhof"
)
const (
  pack = "murus/sub/netz" // 292 Bahnhöfe
  tU = 5 // stark vereinfacht: universelle Zeit zum Umsteigen // TODO
)
var (
  _bhf *bahnhof.Imp = bahnhof.New()
  _gleis *gleis.Imp = gleis.New()
  netz *gra.Imp = gra.New (false, _bhf, _gleis)
  help []string
)


// Eff.: Aktuelle Ecke ist (L, n), postaktuelle Ecke ist die, die vorher aktuell war.
func ins (l Linie, nr uint, n, n1 string, b byte, y, x float64) {
//
  _bhf.Def (l, nr, n, n1, b, y, x)
  netz.Ins (_bhf)
}


// Eff.: Aktuelle Ecke ist (L, n), postaktuelle Ecke ist die, die vorher aktuell war.
func ins1 (l Linie, nr uint, n, n1 string, b byte, y, x float64, t uint) {
//
  ins (l, nr, n, n1, b, y, x)
  _gleis.Def (l, t)
  netz.Edge1 (_gleis)
}


// Eff.: s. ins
func insert (l Linie, n uint, l0 Linie, n0, u uint) {
//
  var bhf *bahnhof.Imp
  if netz.ExPred (func (a Any) bool { b:= a.(*bahnhof.Imp); return b.Linie() == l0 && b.Nummer() == n0 }) {
    bhf = netz.Get ().(*bahnhof.Imp)
  } else {
    errh.Error2 ("fehlt: Linie", uint(l0), "Bhf", n0); ker.Stop (pack, 1)
  }
// aktuell: (l0, n0)
  bhf1:= bhf.Clone ().(*bahnhof.Imp)
  bhf1.Numerieren (l, n)
  bhf1.Umstieg ()
  netz.Ins (bhf1)
  edge (l0, n0, l, n, u)
// postaktuell: (l0, n0), aktuell: (l, n)
}


// Eff.: s. ins1
func insert1 (l Linie, n uint, l0 Linie, n0, u, t uint) {
//
// aktuell: vorherige aktuelle
  netz.Reposition ()
  var bhf *bahnhof.Imp
// aktuell: vorherige aktuelle, postaktuell: alt.
  if netz.ExPred (func (a Any) bool { b:= a.(*bahnhof.Imp); return b.Linie() == l0 && b.Nummer() == n0 }) {
   // aktuell: (L0, n0), postaktuell: vorherige aktuelle
    bhf = netz.Get ().(*bahnhof.Imp)
  } else {
    errh.Error2 ("fehlt: Linie", uint(l0), "Bhf", n0); ker.Stop (pack, 2)
  }
  netz.Reposition ()
// aktuell: vorherige aktuelle, postaktuell: (l0, n0)
  bhf1:= bhf.Clone ().(*bahnhof.Imp)
  bhf1.Numerieren (l, n)
  bhf1.Umstieg ()
  netz.Ins (bhf1)
// aktuell: (l, n), postaktuell: vorherige aktuelle
  _gleis.Def (l, t)
  netz.Edge1 (_gleis)
// verbunden: vorherige aktuelle mit (l, n), aktuell: (l, n).
  edge (l0, n0, l, n, u)
// postaktuell: (l0, n0), aktuell: (l, n)
}


const (
  imBahnhof = true
  mitFußweg = false
)


// Aktuelle Ecke ist (L1, n1), postaktuelle Ecke ist (L, n).
func edge1 (l Linie, n uint, l1 Linie, n1, t uint, direkt bool) {
//
  if netz.ExPred2 (func (a  Any) bool { b :=  a.(*bahnhof.Imp); return  b.Linie() == l  &&  b.Nummer() == n },
                   func (a1 Any) bool { b1:= a1.(*bahnhof.Imp); return b1.Linie() == l1 && b1.Nummer() == n1 }) {
    b, b1:= netz.Get2 (); bhf, bhf1:= b.(*bahnhof.Imp), b1.(*bahnhof.Imp)
    bhf.Umstieg (); bhf1.Umstieg ()
    netz.Put2 (bhf, bhf1)
    if direkt {
      l = bhf.Linie ()
    } else {
      l = Fußweg
    }
    _gleis.Def (l, t)
    netz.Edge1 (_gleis)
  } else {
    if ! netz.ExPred (func (a Any) bool { b:= a.(*bahnhof.Imp); return b.Linie() == l && b.Nummer() == n }) {
      errh.Error2 ("fehlt Linie", uint(l),  "Bhf", n)
    }
    if ! netz.ExPred (func (a Any) bool { b:= a.(*bahnhof.Imp); return b.Linie() == l1 && b.Nummer() == n1 }) {
      errh.Error2 ("fehlt Linie", uint(l1), "Bhf", n1)
    }
    ker.Stop (pack, 3)
  }
}


func edge (l Linie, n uint, l1 Linie, n1, t uint) {
//
  edge1 (l, n, l1, n1, t, imBahnhof)
}


func wr (b Any, a bool) {
//
  b.(*bahnhof.Imp).Write (a)
}


func wr3 (g, b, b1 Any, a bool) {
//
//  b.(*bahnhof.Imp).Write (a)
  g.(*gleis.Imp).Write (a)
  b.(*bahnhof.Imp).Write1 (b1.(*bahnhof.Imp))
}

func write (blinkend bool) {
//
  scr.Buf (true)
  netz.Trav3Cond (wr3)
// ker.Sleep (3)
  netz.TravCond (func (a Any, b bool) { a.(*bahnhof.Imp).Write (b) })
  scr.Buf (false)
}


func netzAufbauen () {
//
  ins  (U1, 10, "Uhlandstr", "",                              'l', 52.5030, 13.3276)
  ins1 (U1, 11, "________Kurfürstendamm", "",                 'o', 52.5038, 13.3314, 1)
  ins1 (U1, 12, "_Wittenberg-", "____platz",                  '.', 52.5018, 13.3430, 2)
  ins1 (U1, 13, "Nollendorf_", "__platz",                     'u', 52.4994, 13.3535, 2)
  ins1 (U1, 14, "Kurfür-", "stenstr",                         'o', 52.5001, 13.3615, 1)
  ins1 (U1, 15, "_Gleis", "dreieck",                          'u', 52.4992, 13.3753, 2)
  ins1 (U1, 16, "Möckern", "brücke",                          'o', 52.4991, 13.3828, 2)
  ins1 (U1, 17, "__Hallesches", "______Tor",                  '.', 52.4977, 13.3908, 1)
  ins1 (U1, 18, "Prinzen-", "__str",                          'u', 52.4984, 13.4056, 2)
  ins1 (U1, 19, "___Kottbusser", "___Tor",                    '.', 52.4994, 13.4187, 2)
  ins1 (U1, 20, "__Görlitzer", "____Bhf",                     'u', 52.4993, 13.4282, 2)
  ins1 (U1, 21, "__________Schlesisches", "__________Tor",    'u', 52.5009, 13.4415, 2)
  ins1 (U1, 22, "WS", "",                                  '.', 52.5049, 13.4490, 2) // U Warschauer Str.

  ins  (U2, 10, "Ruhleben", "",                               'o', 52.5256, 13.2413)
  ins1 (U2, 11, "Olympia-", "__Stadion",                      'l', 52.5173, 13.2501, 3)
  ins1 (U2, 12, "__Neu-", "Westend",                          '.', 52.5162, 13.2592, 2)
  ins1 (U2, 13, "Theodor-Heuss-_______", "_____Platz",        '.', 52.5100, 13.2727, 2)
  ins1 (U2, 14, "Kaiserdamm", "",                             'u', 52.5100, 13.2818, 1)
  ins1 (U2, 15, "Sophie-Charlotte-____", "______Platz",       'o', 52.5108, 13.2954, 2)
  ins1 (U2, 16, "Bismarck", "str",                            '.', 52.5117, 13.3058, 1)
  ins1 (U2, 17, "____Deutsche", "_____Oper",                  'u', 52.5119, 13.3108, 1)
  ins1 (U2, 18, "__Ernst-Reuter", "_____Platz",               'o', 52.5120, 13.3218, 2)
  ins1 (U2, 19, "_____Zoologischer", "________Garten",        'o', 52.5070, 13.3324, 2)

  insert1 (U2, 20, U1, 12,  2, 2)
  insert1 (U2, 21, U1, 13, tU, 2)

//  ins1 (U2, 20, "", "",                                       '.', 52.5016, 13.3430, 2) // Wittenbergplatz
//  ins1 (U2, 21, "", "",                                       'u', 52.4992, 13.3535, 2) // Nollendorfplatz

  ins1 (U2, 22, "Bülow", "_Str",                              'u', 52.4975, 13.3635, 1)
  insert1 (U2, 23, U1, 15, tU, 2)
  ins1 (U2, 24, "MendBarth______", "____Park",                '.', 52.5039, 13.3749, 1)
  ins1 (U2, 25, "Potsdamer Pl_______", "",                    '.', 52.5093, 13.3779, 2)
  ins1 (U2, 26, "Mohrenstr", "",                              'l', 52.5117, 13.3845, 1)
  ins1 (U2, 27, "Stadt", "mitte",                             '.', 52.5104, 13.3899, 1)
  ins1 (U2, 28, "Hausvogtei", "___platz",                     '.', 52.5133, 13.3962, 2)
  ins1 (U2, 29, "Spittel", "_markt",                          'u', 52.5113, 13.4043, 1)
  ins1 (U2, 30, "Märkisches", "__Museum",                     '.', 52.5124, 13.4103, 1)
  ins1 (U2, 31, "Klosterstr", "",                             'o', 52.5172, 13.4125, 2)
  ins1 (U2, 32, "___Alexander-", "____platz",                 '.', 52.5216, 13.4113, 2)
  ins1 (U2, 33, "Rosa-Luxemburg-Platz", "",                   'r', 52.5276, 13.4105, 2)
  ins1 (U2, 34, "Senefelderplatz", "",                        'r', 52.5326, 13.4126, 1)
  ins1 (U2, 35, "Ebers-", "walder Str",                       'r', 52.5417, 13.4122, 2)
  ins1 (U2, 36, "Schönhauser Allee", "",                      'r', 52.5494, 13.4141, 2) // Fußweg ?
  ins1 (U2, 37, "Vinetastr", "",                              'r', 52.5595, 13.4132, 2)
  ins1 (U2, 38, "Pankow", "",                                 'r', 52.5678, 13.4120, 2)

//  edge  (U2, 20, U1, 12, tU)
//  edge  (U2, 21, U1, 13, tU)

  insert  (U3, 10, U1, 13, tU)
  insert1 (U3, 11, U1, 12, tU, 2)

//  ins  (U3, 10, "", "",                                       'u', 52.4996, 13.3535) // Nollendorfplatz 
//  ins1 (U3, 11, "", "",                                       '.', 52.5020, 13.3430, 2) // Wittenbergplatz 

  ins1 (U3, 12, "AugsburgerStr", "",                          'l', 52.5005, 13.3366, 2)
  ins1 (U3, 13, "Spichernstr", "",                            'o', 52.4963, 13.3308, 1)
  ins1 (U3, 14, "Hohenzollern__", "____platz",                '.', 52.4941, 13.3248, 2)
  ins1 (U3, 15, "Fehrbel-__", "liner Pl__",                   'u', 52.4905, 13.3146, 1)
  ins1 (U3, 16, "Heidelberger Platz", "",                     'l', 52.4798, 13.3122, 2) // Fußweg ?
  ins1 (U3, 17, "Rüdesheimer Platz", "",                      'l', 52.4730, 13.3145, 2)
  ins1 (U3, 18, "Breitenbachplatz", "",                       'l', 52.4667, 13.3084, 1)
  ins1 (U3, 19, "Podbielskiallee______", "",                  'u', 52.4641, 13.2958, 2)
  ins1 (U3, 20, "Dahlem Dorf__", "",                          '.', 52.4573, 13.2897, 1)
  ins1 (U3, 21, "Thielplatz", "",                             'r', 52.4510, 13.2818, 2)
  ins1 (U3, 22, "_Oskar-", "Helene-Heim",                     '.', 52.4504, 13.2690, 2)
  ins1 (U3, 23, "Onkel Toms_____", "Hütte",                   '.', 52.4499, 13.2531, 2)
  ins1 (U3, 24, "Krumme Lanke_", "",                          '.', 52.4435, 13.2415, 2)

  edge  (U3, 10, U1, 13, tU)
  edge  (U3, 10, U2, 21, tU)
  edge  (U3, 11, U1, 12, tU)
  edge  (U3, 11, U2, 20, tU)

  insert  (U4, 10, U1, 13, tU)
  ins1 (U4, 11, "Viktoria-", "___Luise-Platz",                'u', 52.4960, 13.3428, 2)
  ins1 (U4, 12, "_Bayeri-", "scher Pl",                       'u', 52.4885, 13.3401, 2)
  ins1 (U4, 13, "___Rathaus Schöneberg", "",                  'u', 52.4831, 13.3420, 1)
  ins1 (U4, 14, "____Innsbrucker", "______Platz",             'u', 52.4784, 13.3429, 2) // Fußweg ?

//  edge  (U4, 10, U2, 21, tU)
//  edge  (U4, 10, U3, 10, tU)

//      U5:
  ins  (U55, 10, "_Haupt-", "bahnhof",                        'l', 52.5252, 13.3691)
  ins1 (U55, 11, "Bundes tag___", "",                         '.', 52.5202, 13.3730, 2)
  ins1 (U55, 12, "Brandenburger Tor__________", "",           '.', 52.5165, 13.3812, 1)
//  ins1 (U55, 13, "Unter den__", "___Linden",                  '.', 52.5171, 13.3889, 1)
//  ins1 (U55, 14, "Museums-", "_insel",                        'u', 52.5261, 13.3979, 1)
//  ins1 (U55, 15, "Berliner Rathaus_", "",                     '.', 52.5185, 13.4078, 1)
  insert  (U5, 10, U2, 32, tU)
//             ^^ 16 etc. etc.
  ins1 (U5, 11, "___________Schilling-", "___________str",    'o', 52.5204, 13.4219, 2)
  ins1 (U5, 12, "_________Strausberger", "_________Platz",    'o', 52.5180, 13.4330, 2)
  ins1 (U5, 13, "Weber-", "wiese",                            'u', 52.5167, 13.4449, 1)
  ins1 (U5, 14, "___Frankfurter", "______Tor",                'o', 52.5159, 13.4532, 1)
  ins1 (U5, 15, "Samariter-_", "___str",                      'u', 52.5148, 13.4644, 2)
  ins1 (U5, 16, "___Frankfurter", "_____Allee",               'o', 52.5144, 13.4749, 1) // Fußweg ?
  ins1 (U5, 17, "___Magdalenen-", "_____str",                 '.', 52.5125, 13.4871, 2)
  ins1 (U5, 18, "______Lichten-_", "________berg",            '.', 52.5106, 13.4970, 1)
  ins1 (U5, 19, "Friedrichs-___", "___felde",                 'u', 52.5056, 13.5128, 2)
  ins1 (U5, 20, "_Tierpark", "",                              'o', 52.5056, 13.5235, 2)
  ins1 (U5, 31, "_Biesdorf-Süd", "",                          'r', 52.4997, 13.5464, 3)
  ins1 (U5, 32, "Elsterwerdaer Platz", "",                    'r', 52.5050, 13.5605, 2)
  ins1 (U5, 33, "Wuhletal", "",                               'u', 52.5127, 13.5752, 2)
  ins1 (U5, 34, "Kaulsdorf Nord_____", "",                    '.', 52.5210, 13.5890, 2)
  ins1 (U5, 35, "__________Neue Grottkauer Str", "",          '.', 52.5285, 13.5906, 2)
  ins1 (U5, 36, "Cottbusser Platz_____", "",                  '.', 52.5339, 13.5965, 1)
  ins1 (U5, 37, "Hellersdorf___________", "",                 'o', 52.5367, 13.6067, 2)
  ins1 (U5, 38, "Louis-Lewin-Str", "",                        'o', 52.5390, 13.6184, 1)
  ins1 (U5, 39, "Hönow", "",                                  'u', 52.5384, 13.6333, 2)

  ins  (U6, 10, "Alt Tegel", "",                              'l', 52.5896, 13.2837)
  ins1 (U6, 11, "Borsigwerke", "",                            'l', 52.5818, 13.2906, 2)
  ins1 (U6, 12, "Holzhauser Str", "",                         'l', 52.5757, 13.2961, 1)
  ins1 (U6, 13, "Otisstr", "",                                'l', 52.5710, 13.3030, 1)
  ins1 (U6, 14, "Scharnweberstr", "",                         'l', 52.5668, 13.3127, 2)
  ins1 (U6, 15, "_____Kurt-Schumacher-", "__________Platz",   'o', 52.5641, 13.3278, 1)
  ins1 (U6, 16, "Afrikanische Str", "",                       'l', 52.5602, 13.3348, 2)
  ins1 (U6, 17, "Rehberge_______", "",                        'u', 52.5568, 13.3410, 1)
  ins1 (U6, 18, "Seestr", "",                                 'l', 52.5500, 13.3528, 2)
  ins1 (U6, 19, "Leopoldplatz", "",                           'l', 52.5464, 13.3591, 1)
  ins1 (U6, 20, "_Wed ding", "",                              '.', 52.5430, 13.3651, 1) // Fußweg ?
  ins1 (U6, 21, "_Reinicken dorfer Str", "",                  '.', 52.5399, 13.3704, 1)
  ins1 (U6, 22, "Schwartz kopffstr", "",                      '.', 52.5355, 13.3769, 1)
  ins1 (U6, 23, "Zinnowitzer Str", "",                        'l', 52.5313, 13.3822, 1)
  ins1 (U6, 24, "Oranienburger___", "____Tor",                'o', 52.5256, 13.3874, 2)
  ins1 (U6, 25, "Friedrichstr", "",                           'u', 52.5201, 13.3880, 1)
//  insert1 (U6, 26, U5, 13, tU, 2)
  ins1 (U6, 26, "Französische", "___Str.",                    '.', 52.5148, 13.3892, 1)
//          ^^ -> 27 etc. etc.
  insert1 (U6, 27, U2, 27, tU, 2)
  ins1 (U6, 28, "Koch_str_", "",                              'o', 52.5059, 13.3906, 1)
  insert1 (U6, 29, U1, 17, tU, 1)
  ins1 (U6, 30, "Mehring", "_damm",                           'u', 52.4943, 13.3888, 2)
  ins1 (U6, 31, "Platz der", "Luftbrücke",                    '.', 52.4852, 13.3860, 2)
  ins1 (U6, 32, "Paradestr", "",                              'u', 52.4782, 13.3859, 1)
  ins1 (U6, 33, "Tempelhof", "",                              'l', 52.4699, 13.3856, 1) // Fußweg ?
  ins1 (U6, 34, "_____Alt- Tempelhof", "",                    '.', 52.4659, 13.3858, 2)
  ins1 (U6, 35, "__Kaiserin-", "_Augusta-Str",                '.', 52.4594, 13.3845, 1)
  ins1 (U6, 36, "Ullsteinstr","",                             'r', 52.4527, 13.3845, 1)
  ins1 (U6, 37, "Westphalweg","",                             'u', 52.4465, 13.3857, 2)
  ins1 (U6, 38, "Alt-Mariendorf", "",                         'u', 52.4392, 13.3881, 1)

  ins  (U7, 10, "Rathaus Spandau", "",                        'l', 52.5353, 13.1999)
  ins1 (U7, 11, "Altstadt", "Spandau",                        'o', 52.5390, 13.2056, 1)
  ins1 (U7, 12, "Zitadelle", "",                              'u', 52.5377, 13.2187, 1)
  ins1 (U7, 13, "Hasel-", "horst",                            'o', 52.5382, 13.2319, 2)
  ins1 (U7, 14, "Paulstern-", "___str",                       'u', 52.5379, 13.2476, 2)
  ins1 (U7, 15, "Rohr", "damm",                               'o', 52.5372, 13.2622, 1)
  ins1 (U7, 16, "Siemens", "_damm",                           'u', 52.5367, 13.2729, 2)
  ins1 (U7, 17, "Halem-", "_weg",                             'o', 52.5367, 13.2864, 1)
  ins1 (U7, 18, "Jakob-Kaiser-", "_Platz",                    'r', 52.5363, 13.2944, 1)
  ins1 (U7, 19, "Jungfern-", "_heide",                        '.', 52.5308, 13.3005, 2) // Fußweg ?
  ins1 (U7, 20, "Mierendorff-", "__Platz",                    'r', 52.5266, 13.3050, 1)
  ins1 (U7, 21, "Richard-Wagner-", "_____Platz",              'o', 52.5161, 13.3067, 2)
  insert1 (U7, 22, U2, 16, tU, 1)
  ins1 (U7, 23, "Wilmers________","dorferStr",                '.', 52.5067, 13.3067, 1)
  ins1 (U7, 24, "Adenauerplatz__", "",                        'u', 52.5000, 13.3071, 2)
  ins1 (U7, 25, "Konstanzer__", "_____Str",                   'u', 52.4945, 13.3091, 1)
  insert1 (U7, 26, U3, 15, tU, 1)
  ins1 (U7, 27, "Blisse-", "_str",                            'u', 52.4866, 13.3221, 2)
  ins1 (U7, 28, "Berliner", "Str",                            '.', 52.4874, 13.3310, 1)
  insert1 (U7, 29, U4, 12, tU, 1)
  ins1 (U7, 30, "Eisenacher", "______Str",                    'u', 52.4895, 13.3503, 1)
  ins1 (U7, 31, "_____Kleist", "______park",                  'u', 52.4905, 13.3604, 2)
  ins1 (U7, 32, "______Yorck", "_______str",                  'u', 52.4929, 13.3698, 1)
  ins1 (U7, 33, "Möckern", "brücke",                          'o', 52.4991, 13.3828, 2)
  insert1 (U7, 34, U6, 30, tU, 2)
  ins1 (U7, 35, "Gneisenau", "str",                           'r', 52.4913, 13.3958, 1)
  ins1 (U7, 36, "_Süd-", "_stern",                            'u', 52.4891, 13.4077, 2)
  ins1 (U7, 37, "Hermann-", "__platz",                        'r', 52.4864, 13.4243, 2)
  ins1 (U7, 38, "Rathaus Neukölln", "",                       'r', 52.4819, 13.4339, 1)
  ins1 (U7, 39, "____Karl Marx-Str", "",                      '.', 52.4765, 13.4392, 2)
  ins1 (U7, 40, "__Neukölln", "",                             'o', 52.4686, 13.4419, 1) // Fußweg ?
  ins1 (U7, 41, "Grenzallee", "",                             'u', 52.4631, 13.4445, 2)
  ins1 (U7, 42, "Blaschkoallee", "",                          'l', 52.4516, 13.4495, 1)
  ins1 (U7, 43, "Parchimer Allee", "",                        'l', 52.4457, 13.4450, 2)
  ins1 (U7, 44, "Britz Süd", "",                              'l', 52.4377, 13.4475, 1)
  ins1 (U7, 45, "Johannistaler Chaussee_____", "",            '.', 52.4294, 13.4533, 2)
  ins1 (U7, 46, "Lipschitzallee", "",                         'l', 52.4235, 13.4626, 1)
  ins1 (U7, 47, "Wutzky", "allee",                            '.', 52.4226, 13.4750, 2)
  ins1 (U7, 48, "Zwickauer", "__Damm",                        'r', 52.4232, 13.4838, 1)
  ins1 (U7, 49, "Rudow", "",                                  'l', 52.4158, 13.4966, 1)

  ins  (U8, 10, "Wittenau", "",                               'r', 52.5926, 13.3344)
  ins1 (U8, 11, "Rathaus", "Reinickendorf_",                  '.', 52.5906, 13.3256, 1)
  ins1 (U8, 12, "Karl-Bonhoeffer-", "_Nervenklinik",          'o', 52.5786, 13.3330, 2)
  ins1 (U8, 13, "Lindauer______", "__Allee",                  'u', 52.5753, 13.3396, 1)
  ins1 (U8, 14, "________Paracelsus-Bad", "",                 'u', 52.5741, 13.3494, 2)
  ins1 (U8, 15, "____Residenzstr", "",                        'u', 52.5706, 13.3608, 2)
  ins1 (U8, 16, "________Franz Naumann-Platz", "",            '.', 52.5646, 13.3639, 1)
  ins1 (U8, 17, "Osloer Str___", "",                          '.', 52.5571, 13.3733, 2)
  ins1 (U8, 18, "______Pankstr", "",                          'o', 52.5523, 13.3815, 2)
  ins1 (U8, 19, "Gesund-", "brunnen",                         'r', 52.5488, 13.3891, 1) // Fußweg ?  U 52.5495, 13.3866 - S 52.5488, 13.3907
  ins1 (U8, 20, "Volta-", "_str",                             'r', 52.5425, 13.3927, 2)
  ins1 (U8, 21, "Bernau er Str", "",                          '.', 52.5377, 13.3965, 1)
  ins1 (U8, 22, "_____Rosen-", "____thaler Pl",               'o', 52.5297, 13.4014, 2)
  ins1 (U8, 23, "Weinmeisterstr", "",                         'r', 52.5255, 13.4057, 1)
  insert1 (U8, 24, U2, 32, tU, 2)
  ins1 (U8, 25, "Jannowitz", "brücke",                        'r', 52.5151, 13.4186, 2)
  ins1 (U8, 26, "Heinrich-", "Heine-Platz",                   'r', 52.5104, 13.4160, 1)
  ins1 (U8, 27, "Moritzplatz", "",                            'r', 52.5036, 13.4107, 2)
  insert1 (U8, 28, U1, 19, tU, 2)
  ins1 (U8, 29, "________Schönleinstr", "",                   'u', 52.4933, 13.4224, 1)
  insert1 (U8, 30, U7, 37, tU, 2)
  ins1 (U8, 31, "Boddinstr", "",                              'l', 52.4803, 13.4253, 2)
  ins1 (U8, 32, "Leinestr", "",                               'l', 52.4734, 13.4281, 2)
  ins1 (U8, 33, "Hermannstr__", "",                           'u', 52.4685, 13.4307, 1) // Fußweg ?
  edge  (U8, 24, U5, 10, tU)

  insert  (U9, 10, U8, 17, tU)
  ins1 (U9, 11, "Nauener", "Platz",                           '.', 52.5515, 13.3674, 1)
  insert1 (U9, 12, U6, 19, tU, 1)
  ins1 (U9, 13, "Amrumer Str", "",                            'l', 52.5422, 13.3495, 2)
  ins1 (U9, 14, "_West-", "_hafen",                           '.', 52.5362, 13.3443, 1)
  ins1 (U9, 15, "Birkenstr", "",                              'u', 52.5323, 13.3413, 1)
  ins1 (U9, 16, "Turmstr", "",                                'l', 52.5259, 13.3430, 2)
  ins1 (U9, 17, "Hansaplatz__________", "",                   'o', 52.5179, 13.3423, 1)
  insert1 (U9, 18, U2, 19, tU, 2)
  insert1 (U9, 19, U1, 11, tU, 2)
  insert1 (U9, 20, U3, 13, tU, 1)
  ins1 (U9, 21, "Güntzelstr____", "",                         '.', 52.4907, 13.3310, 1)
  insert1 (U9, 22, U7, 28, tU, 1)
  ins1 (U9, 23, "Bundes", "platz",                            '.', 52.4781, 13.3286, 2)
  ins1 (U9, 24, "_Friedrich-", "Wilhelm-Platz__",             '.', 52.4724, 13.3282, 1)
  ins1 (U9, 25, "___Walther-Schreiber-", "____Platz",         '.', 52.4649, 13.3284, 1)
  ins1 (U9, 26, "Schloß str___", "",                          '.', 52.4609, 13.3249, 2)
  ins1 (U9, 27, "_Rathaus Steglitz", "",                      '.', 52.4551, 13.3195, 1)

  ins  (S1, 10, "Wannsee", "",                                'r', 52.4213, 13.1797)
  ins1 (S1, 11, "Nikolassee", "",                             'r', 52.4319, 13.1932, 2)
  ins1 (S1, 12, "__Schlachtensee", "",                        'o', 52.4400, 13.2150, 3)
  ins1 (S1, 13, "_Mexiko-", "__platz",                        '.', 52.4369, 13.2331, 2)
  ins1 (S1, 14, "Zehlendorf", "",                             'u', 52.4310, 13.2582, 2)
  ins1 (S1, 15, "Sundgauer Str______", "",                    '.', 52.4363, 13.2738, 2)
  ins1 (S1, 16, "Lichterfelde West________", "",              '.', 52.4432, 13.2934, 3)
  ins1 (S1, 17, "Botanischer Garten_____", "",                '.', 52.4480, 13.3073, 1)
  insert1 (S1, 18, U9, 27, tU, 3)
  ins1 (S1, 19, "__________Feuerbach-", "__________str",      'u', 52.4633, 13.3327, 1)
  ins1 (S1, 20, "__________Friedenau", "",                    'u', 52.4704, 13.3412, 2)
  ins1 (S1, 21, "_Schöneberg", "",                            'r', 52.4793, 13.3519, 2)
  ins1 (S1, 22, "_Julius-Leber-", "________Brücke",           'u', 52.4861, 13.3607, 2)
  ins1 (S1, 23, "YorckGroßgörschen", "_____str",              '.', 52.4923, 13.3679, 1)
  ins1 (S1, 24, "Anhalter Bhf", "",                           'r', 52.5045, 13.3823, 2)
  insert1 (S1, 25, U2, 25, tU, 2)
  insert1 (S1, 26, U55, 12, tU, 2)
  insert1 (S1, 27, U6, 25, tU, 3)
  ins1 (S1, 28, "Oranienbur gerStr____", "",                  '.', 52.5250, 13.3929, 2)
  ins1 (S1, 29, "Nordbhf", "",                                'o', 52.5317, 13.3890, 2)
  ins1 (S1, 30, "__Humboldt-", "____hain",                    '.', 52.5448, 13.3794, 3)
  insert1 (S1, 31, U8, 19, tU, 2)
  ins1 (S1, 32, "__Bornholmer", "_______Str",                 '.', 52.5545, 13.3980, 2)
  ins1 (S1, 33, "_________Wollank-", "__________str",         '.', 52.5652, 13.3924, 2)
  ins1 (S1, 34, "Schönholz", "",                              'r', 52.5712, 13.3815, 2)
  ins1 (S1, 35, "Wilhelmsruh", "",                            'r', 52.5818, 13.3621, 3)
  ins1 (S1, 36, "Wittenau_", "",                              'r', 52.5970, 13.3343, 3)
  ins1 (S1, 37, "Waidmannslust", "",                          'r', 52.6064, 13.3211, 2)
  ins1 (S1, 38, "Hermsdorf", "",                              'r', 52.6174, 13.3075, 3)
  ins1 (S1, 39, "Frohnau", "",                                'o', 52.6323, 13.2904, 2)
  ins1 (S1, 40, "Hohen Neuendorf", "",                        'r', 52.6686, 13.2870, 5)
  ins1 (S1, 41, "Birkenwerder", "",                           'r', 52.6912, 13.2889, 2)
  ins1 (S1, 42, "Borgsdorf", "",                              'r', 52.7157, 13.2764, 4)
  ins1 (S1, 43, "Lehnitz", "",                                'r', 52.7411, 13.2635, 3)
  ins1 (S1, 44, "Oranienburg", "",                            'r', 52.7536, 13.2496, 2)
  edge  (S1, 40, S1, 39, tU)
  edge1  (S1, 23, U7, 32, 10, mitFußweg)
  edge1  (S1, 36, U8, 10, 10, mitFußweg)
  edge  (S1, 40, S1, 39, tU)

  ins  (S2,  6, "Bernau", "",                                 'r', 52.6756, 13.5917)
  ins1 (S2,  7, "Bernau-Friedenstal", "",                     'r', 52.6683, 13.5646, 2)
  ins1 (S2,  8, "Zepernick", "",                              'r', 52.6599, 13.5342, 3)
  ins1 (S2,  9, "Röntgental", "",                             'r', 52.6487, 13.5136, 3)
  ins1 (S2, 10, "Buch", "",                                   'o', 52.6358, 13.4916, 3)
  ins1 (S2, 11, "Karow", "",                                  'r', 52.6153, 13.4695, 3)
  ins1 (S2, 12, "Blankenburg", "",                            'r', 52.5916, 13.4436, 3)
  ins1 (S2, 13, "Pankow-Heinersdorf", "",                     'r', 52.5782, 13.4297, 2)
  insert1 (S2, 14, U2, 38, tU, 1)
  insert1 (S2, 15, S1, 32, tU, 2)
  insert1 (S2, 16, U8, 19, tU, 3)
  insert1 (S2, 17, S1, 30, tU, 2)
  insert1 (S2, 18, S1, 29, tU, 3)
  insert1 (S2, 19, S1, 28, tU, 2)
  insert1 (S2, 21, U6, 25, tU, 2)
  insert1 (S2, 20, S1, 26, tU, 3)
  insert1 (S2, 22, U2, 25, tU, 2)
  insert1 (S2, 23, S1, 24, tU, 2)
  ins1 (S2, 24, "______Yorck", "_______str",                  'u', 52.4914, 13.3721, 2)
  ins1 (S2, 25, "Südkreuz", "",                               'u', 52.4758, 13.3654, 1)
  ins1 (S2, 26, "Priesterweg", "",                            'u', 52.4597, 13.3562, 3)
  ins1 (S2, 27, "Attilastr", "",                              'r', 52.4479, 13.3608, 2)
  ins1 (S2, 28, "Marienfelde", "",                            'r', 52.4241, 13.3747, 3)
  ins1 (S2, 29, "Buckower Chaussee", "",                      '.', 52.4103, 13.3829, 3)
  ins1 (S2, 30, "Schichauweg", "",                            'r', 52.3988, 13.3892, 2)
  ins1 (S2, 31, "Lichtenrade", "",                            'u', 52.3877, 13.3963, 3)
  ins1 (S2, 32, "Mahlow", "",                                 'r', 52.3609, 13.4076, 5)
  ins1 (S2, 33, "Blankenfelde", "",                           'u', 52.3369, 13.4161, 3)
  edge  (S2, 16, S1, 31, tU)
  edge  (S2, 21, S1, 27, tU)
  edge  (S2, 22, S1, 25, tU)
  edge1  (S2, 24, U7, 32, 5, mitFußweg)
  edge1  (S2, 24, S1, 23, 5, mitFußweg)

  ins  (S25,10, "Hennigsdorf", "",                            'o', 52.6384, 13.2056)
  ins1 (S25,11, "Heiligensee", "",                            'o', 52.6245, 13.2294, 2)
  ins1 (S25,12, "Schulzendorf____________", "",               'u', 52.6130, 13.2459, 3)
  ins1 (S25,13, "Tegel_", "",                                 'r', 52.5881, 13.2898, 4)
  ins1 (S25,14, "Eichborndamm", "",                           'u', 52.5777, 13.3169, 3)
  insert1 (S25,15, U8, 12, tU, 2)
  ins1 (S25,16, "_____Alt-_", "__Reinickendorf",              '.', 52.5778, 13.3511, 2)
  insert1 (S25,17, S1, 34, tU, 2)
  insert1 (S25,18, S1, 33, tU, 2)
  insert1 (S25,19, S1, 32, tU, 2)
  insert1 (S25,20, S1, 31, tU, 2)
  insert1 (S25,21, S1, 30, tU, 2)
  insert1 (S25,22, S1, 29, tU, 3)
  insert1 (S25,23, S1, 28, tU, 2)
  insert1 (S25,24, S1, 27, tU, 2)
  insert1 (S25,25, S1, 26, tU, 3)
  insert1 (S25,26, S1, 25, tU, 2)
  insert1 (S25,27, S1, 24, tU, 2)
  insert1 (S25,28, U7, 32, tU, 2)
  insert1 (S25,29, S2, 25, tU, 1)
  insert1 (S25,30, S2, 26, tU, 3)
  ins1 (S25,31, "Südende", "",                                'l', 52.4484, 13.3539, 2)
  ins1 (S25,32, "Lankwitz", "",                               'l', 52.4387, 13.3418, 2)
  ins1 (S25,33, "Lichterfelde Ost_________", "",              '.', 52.4300, 13.3284, 2)
  ins1 (S25,34, "Osdorfer Str_____", "",                      '.', 52.4193, 13.3146, 2)
  ins1 (S25,35, "Lichterfelde Süd", "",                       'u', 52.4102, 13.3087, 2)
  ins1 (S25,36, "Teltow Stadt", "",                           'u', 52.3969, 13.2766, 2)
  edge1  (S25,13, U6, 10, 5, mitFußweg)
  edge  (S25,19, S2, 15, tU)
  edge  (S25,20, U8, 19, tU)
  edge  (S25,20, S1, 31, tU)
  edge  (S25,20, S2, 16, tU)
  edge  (S25,24, U6, 25, tU)
  edge  (S25,24, S2, 21, tU)
  edge  (S25,26, U2, 25, tU)
  edge  (S25,26, S1, 25, tU)
  edge  (S25,26, S2, 22, tU)
  edge  (S25,27, S2, 23, tU)
  edge  (S25,28, S1, 23, tU)

  ins  (S3, 10, "Ostbahnhof", "",                             'r', 52.5103, 13.4348)
  ins1 (S3, 11, "Warschauer Str_______", "",                  '.', 52.5061, 13.4515, 3) // S
  ins1 (S3, 12, "Ostkreuz", "",                               'o', 52.5031, 13.4693, 2)
  ins1 (S3, 13, "Rummelsburg", "",                            'u', 52.5012, 13.4786, 1)
  ins1 (S3, 14, "_______BBhf Rummelsburg", "",                'u', 52.4933, 13.4979, 3)
  ins1 (S3, 15, "Karlshorst", "",                             'l', 52.4805, 13.5272, 3)
  ins1 (S3, 16, "Wuhlheide", "",                              'l', 52.4686, 13.5543, 2)
  ins1 (S3, 17, "Köpenick", "",                               'l', 52.4586, 13.5815, 3)
  ins1 (S3, 18, "Hirschgarten", "",                           'u', 52.4579, 13.6031, 2)
  ins1 (S3, 19, "____Friedrichshagen", "",                    'o', 52.4574, 13.6236, 2)
  ins1 (S3, 20, "Rahnsdorf ", "",                             'r', 52.4516, 13.6901, 5)
  ins1 (S3, 21, "Wilhelmshagen ", "",                         'l', 52.4386, 13.7223, 4)
  ins1 (S3, 22, "Erkner", "",                                 'u', 52.4293, 13.7507, 3)
  edge1  (S3, 11, U1, 22, 3, mitFußweg)

  insert  (S4,  0, S3, 12, tU)
  insert1 (S4,  1, U5, 16, tU, 3)
  ins1 (S4,  2, "Storkower Str", "",                          'r', 52.5238, 13.4646, 2)
  ins1 (S4,  3, "Landsberger Allee", "",                      'r', 52.5295, 13.4548, 2)
  ins1 (S4,  4, "Greifswalder Str", "",                       'r', 52.5402, 13.4392, 3)
  ins1 (S4,  5, "Prenzlauer Allee", "",                       'r', 52.5448, 13.4259, 2)
  insert1 (S4,  6, U2, 36, tU, 1)
  insert1 (S4,  7, U8, 19, tU, 3)
  insert1 (S4,  8, U6, 20, tU, 3)
  insert1 (S4,  9, U9, 14, tU, 2)
  ins1 (S4, 10, "Beussel-___", "___str",                      '.', 52.5344, 13.3292, 2)
  insert1 (S4, 11, U7, 19, tU, 3)
  ins1 (S4, 12, "Westend________", "",                        'o', 52.5179, 13.2847, 2)
  ins1 (S4, 13, "Messe___", "Nord",                           'u', 52.5077, 13.2835, 2)
  ins1 (S4, 14, "_West kreuz", "",                            '.', 52.5008, 13.2840, 4)
  ins1 (S4, 15, "Halensee____", "",                           'u', 52.4961, 13.2905, 2)
  ins1 (S4, 16, "Hohenzollern-___________", "_________damm",  'u', 52.4886, 13.3003, 2)
  insert1 (S4, 17, U3, 16, tU, 2)
  insert1 (S4, 18, U9, 23, tU, 2)
  insert1 (S4, 19, U4, 14, tU, 2)
  insert1 (S4, 20, S1, 21, tU, 2)
  insert1 (S4, 21, S2, 25, tU, 2)
  insert1 (S4, 22, U6, 33, tU, 3)
  insert1 (S4, 23, U8, 33, tU, 4)
  insert1 (S4, 24, U7, 40, tU, 1)
  ins1 (S4, 25, "Sonnenallee", "",                            'o', 52.4729, 13.4556, 2)
  ins1 (S4, 26, "Treptower Park", "",                         'u', 52.4938, 13.4618, 2)
  edge  (S4,  7, S1, 31, tU)
  edge  (S4,  7, S2, 16, tU)
  edge  (S4,  7, S25,20, tU)
  edge  (S4, 26, S4,  0, 2)
  edge1  (S4, 13, U2, 14, 10, mitFußweg)

  ins  (S4, 52, "Köllnische", "__Heide",                      'u', 52.4697, 13.4685)
  ins1 (S4, 53, "Baumschulenweg ", "",                        'r', 52.4669, 13.4908, 2)
  ins1 (S4, 54, "_Schöneweide", "",                           'l', 52.4549, 13.5096, 3)
  ins1 (S4, 55, "Bbhf Schöneweide", "",                       'l', 52.4467, 13.5238, 3)
  ins1 (S4, 56, "Adlershof", "",                              'l', 52.4346, 13.5418, 3)
  ins1 (S4, 57, "Altglienicke", "",                           'r', 52.4072, 13.5586, 5)
  ins1 (S4, 58, "Grünbergallee", "",                          'r', 52.3992, 13.5416, 2)
  ins1 (S4, 59, "________Flughafen Berlin Schönefeld", "",    '.', 52.3910, 13.5130, 3)
  edge  (S4, 52, S4, 24, tU)
  ins  (S4, 78, "Grünau", "",                                 'r', 52.4124, 13.5743)
  ins1 (S4, 79, "Eichwalde", "",                              'r', 52.3713, 13.6154, 5)
  ins1 (S4, 80, "Zeuthen", "",                                'r', 52.3488, 13.6274, 3)
  ins1 (S4, 81, "Wildau", "",                                 'r', 52.3201, 13.6341, 4)
  ins1 (S4, 82, "______Königs Wusterhausen", "",              '.', 52.2964, 13.6315, 3)
  edge  (S4, 78, S4, 56, 5)
  ins  (S4, 98, "Oberspree", "",                              'r', 52.4524, 13.5381)
  ins1 (S4, 99, "Spindlersfeld", "",                          'u', 52.4473, 13.5613, 3)
  edge  (S4, 98, S4, 54, 3)

  insert  (S5, 10, S4, 14, tU)
  ins1 (S5, 11, "Charlotten____", "______burg",               '.', 52.5020, 13.3000, 2)
  ins1 (S5, 12, "Savigny-__", "_platz",                       '.', 52.5052, 13.3192, 2)
  insert1 (S5, 13, U2, 19, tU, 3)
  ins1 (S5, 14, "Tiergarten", "",                             'r', 52.5144, 13.3365, 1)
  ins1 (S5, 15, "Bellevue", "",                               'o', 52.5200, 13.3481, 2)
//  ins1 (S5, 16, "_Haupt-", "bahnhof",                         'l', 52.5252, 13.3691, 2)
  insert1 (S5, 16, U55, 10, tU, 3)
  insert1 (S5, 17, U6, 25, tU, 3)
  ins1 (S5, 18, "Hackescher_________", "____Markt",           '.', 52.5226, 13.4023, 2)
  insert1 (S5, 19, U2, 32, tU, 2)
  insert1 (S5, 20, U8, 25, tU, 2)
  insert1 (S5, 21, S3, 10, tU, 2)
  insert1 (S5, 22, S3, 11, tU, 3)
  insert1 (S5, 23, S3, 12, tU, 2)
  ins1 (S5, 24, "_Nöldner-", "__platz",                       '.', 52.5038, 13.4853, 2)
  insert1 (S5, 25, U5, 18, tU, 2)
  ins1 (S5, 26, "_Friedrichsfelde", "_________Ost",           '.', 52.5141, 13.5201, 3)
  ins1 (S5, 27, "Biesdorf_", "",                              'o', 52.5130, 13.5560, 3)
  insert1 (S5, 28, U5, 33, tU, 2)
  ins1 (S5, 29, "_Kaulsdorf", "",                             'o', 52.5122, 13.5901, 2)
  ins1 (S5, 30, "Mahlsdorf", "",                              'u', 52.5122, 13.6113, 2)

//  ins1 (S5, 31, "Birkenstein", "",                            ".", 52.0000, 13.0000, 0)
//  ins1 (S5, 32, "Hoppegarten", "",                            ".", 52.0000, 13.0000, 0)
//  ins1 (S5, 33, "Neuenhagen", "",                             ".", 52.0000, 13.0000, 0)
//  ins1 (S5, 34, "Fredersdorf", "",                            ".", 52.0000, 13.0000, 0)
//  ins1 (S5, 35, "Strausberg", "",                             ".", 52.0000, 13.0000, 0)
//  ins1 (S5, 36, "Hegermühle", "",                             ".", 52.0000, 13.0000, 0)
//  ins1 (S5, 37, "Strausberg Stadt", "",                       ".", 52.0000, 13.0000, 0)
//  ins1 (S5, 38, "Strausberg Nord", "",                        ".", 52.0000, 13.0000, 0)

  edge1  (S5, 11, U7, 23, 10, mitFußweg)
  edge  (S5, 13, U9, 18, tU)
  edge  (S5, 17, S1, 27, tU)
  edge  (S5, 17, S2, 21, tU)
  edge  (S5, 17, S25,24, tU)
  edge  (S5, 19, U5, 10, tU)
  edge  (S5, 19, U8, 24, tU)
  edge  (S5, 23, S4,  0, tU)

  ins  (S7,  7, "Potsdam", "",                                'o', 52.3918, 13.0671)
  ins1 (S7,  8, "Babelsberg", "",                             'u', 52.3914, 13.0928, 4)
  ins1 (S7,  9, "Griebnitzsee", "",                           'r', 52.3945, 13.1274, 3)
  insert1 (S7, 10, S1, 10, tU, 5)
  insert1 (S7, 11, S1, 11, tU, 2)
  ins1 (S7, 12, "Grunewald", "",                              'l', 52.4882, 13.2610, 7)
  insert1 (S7, 13, S4, 14, tU, 4)
  insert1 (S7, 14, S5, 11, tU, 2)
  insert1 (S7, 15, S5, 12, tU, 2)
  insert1 (S7, 16, S5, 13, tU, 3)
  insert1 (S7, 17, S5, 14, tU, 1)
  insert1 (S7, 18, S5, 15, tU, 2)
  insert1 (S7, 19, S5, 16, tU, 2)
  insert1 (S7, 20, S5, 17, tU, 3)
  insert1 (S7, 21, S5, 18, tU, 2)
  insert1 (S7, 22, S5, 19, tU, 2)
  insert1 (S7, 23, S5, 20, tU, 2)
  insert1 (S7, 24, S5, 21, tU, 2)
  insert1 (S7, 25, S5, 22, tU, 3)
  insert1 (S7, 26, S5, 23, tU, 2)
  insert1 (S7, 27, S5, 24, tU, 2)
  insert1 (S7, 28, S5, 25, tU, 2)
  insert1 (S7, 29, S5, 26, tU, 3)
  ins1 (S7, 30, "Springpfuhl", "",                            'r', 52.5270, 13.5366, 4)
  ins1 (S7, 31, "Poelchaustr", "",                            'r', 52.5358, 13.5355, 2)
  ins1 (S7, 32, "Marzahn", "",                                'r', 52.5436, 13.5413, 1)
  ins1 (S7, 33, "Raoul-Wallenberg-Str", "",                   'r', 52.5507, 13.5476, 2)
  ins1 (S7, 34, "Mehrower Allee", "",                         'r', 52.5576, 13.5536, 2)
  ins1 (S7, 35, "Ahrensfelde", "",                            'o', 52.5713, 13.5656, 2)
  edge  (S7, 13, S5, 10, tU)
  edge  (S7, 16, U2, 19, tU)
  edge  (S7, 16, U9, 18, tU)
  edge  (S7, 20, U6, 25, tU)
  edge  (S7, 20, S1, 27, tU)
  edge  (S7, 20, S2, 21, tU)
  edge  (S7, 20, S25,24, tU)
  edge  (S7, 20, S5, 17, tU)
  edge  (S7, 22, U2, 32, tU)
  edge  (S7, 22, U5, 10, tU)
  edge  (S7, 22, U8, 24, tU)
  edge  (S7, 22, S5, 19, tU)
  edge  (S7, 23, U8, 25, tU)

  ins  (S75,50, "Spandau_", "",                               'u', 52.5348, 13.1963)
  ins1 (S75,51, "Stresow", "",                                'r', 52.5319, 13.2093, 2)
  ins1 (S75,52, "Pichelsberg____", "",                        'u', 52.5102, 13.2276, 3)
  ins1 (S75,53, "Olympia-", "stadion",                        '.', 52.5112, 13.2424, 2)
  ins1 (S75,54, "Heerstr", "",                                'u', 52.5083, 13.2587, 2)
  ins1 (S75,55, "Messe Süd", "",                              'u', 52.4987, 13.2700, 3)
  insert1 (S75,56, S4, 14, tU, 2)
  insert1 (S75,57, S7, 14, tU, 2)
  insert1 (S75,58, S7, 15, tU, 2)
  insert1 (S75,59, S7, 16, tU, 3)
  insert1 (S75,60, S7, 17, tU, 1)
  insert1 (S75,61, S7, 18, tU, 2)
  insert1 (S75,62, S7, 19, tU, 2)
  insert1 (S75,63, S7, 20, tU, 3)
  insert1 (S75,64, S7, 21, tU, 2)
  insert1 (S75,65, S7, 22, tU, 2)
  insert1 (S75,66, S7, 23, tU, 2)
  insert1 (S75,67, S7, 24, tU, 2)
  insert1 (S75,68, S7, 25, tU, 3)
  insert1 (S75,69, S7, 26, tU, 2)
  insert1 (S75,70, S7, 27, tU, 2)
  insert1 (S75,71, S7, 28, tU, 2)
  insert1 (S75,72, S7, 29, tU, 3)
  insert1 (S75,73, S7, 30, tU, 4)
  ins1 (S75,74, "Gehrenseestr", "",                           'l', 52.5565, 13.5248, 4)
  ins1 (S75,75, "Hohenschönhausen", "",                       'l', 52.5663, 13.5125, 2)
  ins1 (S75,76, "Wartenberg", "",                             'o', 52.5730, 13.5038, 2)
  edge1  (S75,50, U7, 10, 10, mitFußweg)
  edge  (S75,56, S5, 10, tU)
  edge  (S75,56, S7, 13, tU)
  edge  (S75,57, S5, 11, tU)
  edge  (S75,57, S7, 14, tU)
  edge  (S75,58, S5, 12, tU)
  edge  (S75,58, S7, 15, tU)
  edge  (S75,58, S75,58, tU)
  edge  (S75,59, U2, 19, tU)
  edge  (S75,59, U9, 18, tU)
  edge  (S75,59, S5, 13, tU)
  edge  (S75,59, S7, 16, tU)
  edge  (S75,63, U6, 25, tU)
  edge  (S75,63, S1, 27, tU)
  edge  (S75,63, S2, 21, tU)
  edge  (S75,63, S5, 17, tU)
  edge  (S75,65, U2, 32, tU)
  edge  (S75,65, U5, 10, tU)
  edge  (S75,65, U8, 24, tU)
  edge  (S75,65, S5, 19, tU)
  edge  (S75,66, U8, 25, tU)
  edge  (S75,66, S5, 20, tU)
  edge  (S75,67, S3, 10, tU)
  edge  (S75,67, S5, 21, tU)
  edge  (S75,68, U1, 22, tU)
  edge  (S75,68, S3, 11, tU)
  edge  (S75,69, S3, 12, tU)
  edge  (S75,69, S4,  0, tU)
  edge  (S75,69, S5, 23, tU)
  edge  (S75,71, U5, 18, tU)
  edge  (S75,71, S5, 25, tU)
  edge  (S75,72, S5, 26, tU)

  insert  (S8, 06, S1, 40, tU)
  ins1 (S8,  7, "Bergfelde", "",                              'r', 52.6702, 13.3201, 4)
  ins1 (S8,  8, "Schönfließ", "",                             'r', 52.6646, 13.3406, 2)
  ins1 (S8,  9, "Mühlenbeck-Mönchmühle", "",                  'r', 52.6548, 13.3859, 3)
  insert1 (S8, 10, S2, 12, tU, 10)
  insert1 (S8, 11, S2, 13, tU, 2)
  insert1 (S8, 12, U2, 38, tU, 1)
  insert1 (S8, 13, S1, 32, tU, 2)
  insert1 (S8, 14, U2, 36, tU, 1)
  insert1 (S8, 15, S4,  5, tU, 2)
  insert1 (S8, 16, S4,  4, tU, 2)
  insert1 (S8, 17, S4,  3, tU, 3)
  insert1 (S8, 18, S4,  2, tU, 2)
  insert1 (S8, 19, U5, 16, tU, 2)
  insert1 (S8, 20, S3, 12, tU, 3)
  insert1 (S8, 21, S4, 26, tU, 2)
  ins1 (S8, 22, "Plänterwald", "",                            'r', 52.4785, 13.4733, 1)
  insert1 (S8, 23, S4, 53, tU, 2)
  insert1 (S8, 24, S4, 54, tU, 3)
  insert1 (S8, 25, S4, 55, tU, 3)
  insert1 (S8, 26, S4, 56, tU, 3)
  insert1 (S8, 27, S4, 78, tU, 5)
  insert1 (S8, 28, S4, 79, tU, 5)
  insert1 (S8, 29, S4, 80, tU, 3)
  edge  (S8, 12, S2, 14, tU)
  edge  (S8, 13, S2, 15, tU)
  edge  (S8, 13, S25,19, tU)
  edge  (S8, 14, S4,  6, tU)
  edge  (S8, 19, S4,  1, tU)
  edge  (S8, 20, S4,  0, tU)
  edge  (S8, 20, S5, 23, tU)

  insert  (S85,50, S1, 37, tU)
  insert1 (S85,51, S1, 36, tU, 3)
  insert1 (S85,52, S1, 35, tU, 3)
  insert1 (S85,53, S1, 34, tU, 2)
  insert1 (S85,54, S1, 33, tU, 2)
  insert1 (S85,55, S1, 32, tU, 2)
  edge  (S85,54, S25,18, tU)
  edge  (S85,55, S2, 15, tU)
  edge  (S85,55, S25,19, tU)
  edge  (S85,55, S8, 13, tU)

  write (false)
}


func gew () bool {
//
  var dummy *bahnhof.Imp
  for {
//    ok:= false
    c, _:= kbd.Command ()
    scr.MouseCursor (true)
    switch c {
    case kbd.Esc:
      return false
    case kbd.Enter, kbd.Back, kbd.Left, kbd.Right, kbd.Up, kbd.Down:
      dummy.SkalaEditieren ()
      write (false)
    case kbd.Help:
      errh.WriteHelp (help)
    case kbd.Hither:
      if netz.ExPred (func (a Any) bool { return a.(*bahnhof.Imp).UnterMaus() }) {
        return true
      }
    case kbd.There, kbd.Push, kbd.Thither:
      dummy.SkalaEditieren ()
      write (false)
/*
    case kbd.This:
      ok = netz.ExPred (func (a Any) bool { return a.(*bahnhof.Imp).UnterMaus() })
    case kbd.Push:
      x, y:= scr.MousePosGr ()
      if ok {
        bhf:= netz.Get ().(*bahnhof.Imp)
        bhf.Rescale (uint(x), uint(y))
        netz.Put (bhf)
        write (false)
      }
*/
    case kbd.PrintScr:
      errh.DelHint ()
      img.Print1 ()
    }
  }
  return false
}


func gewählt () bool {
//
  loop: for {
    errh.Hint ("Start auswählen     (Klick mit linker Maustaste)")
    if gew () { // Start aktuell
      netz.Get ().(*bahnhof.Imp).Write (true)
      netz.Position (true) // Start postaktuell
      write (false)
    } else {
      break
    }
    errh.Hint ("Ziel auswählen     (Klick mit linker Maustaste)")
    for {
      if gew () { // Ziel aktuell
        if ! netz.Positioned () {
          netz.Get ().(*bahnhof.Imp).Write (true)
          errh.DelHint()
          return true
        }
      } else {
        break loop
      }
    }
  }
  write (false)
  defer ker.Terminate ()
  return false
}


func suchen () {
//
/*
  const
    maxU = 8
  var (
    startlinie, ziellinie [maxU]Linie
    startnummer, zielnummer [maxU]uint
    n, imin, kmin uint
  )
  t1, t2:= netz.Get2 ()
  bhf1, bhf2:= t1.(*bahnhof.Imp), t2.(*bahnhof.Imp)
  startlinie [0], startnummer [0] = bhf1.Linie (), bhf1.Nummer ()
  ziellinie [0], zielnummer [0] = bhf2.Linie (), bhf2.Nummer ()
  ss, zz:= 1, 1
  netz.Trav (func (a Any) {
               bhf:= a.(*bahnhof.Imp)
               if bhf.Equiv (bhf1) {
                 startlinie [ss], startnummer [ss] = bhf.Linie (), bhf.Nummer ()
                 ss ++
               } else if bhf.Equiv (bhf2) {
                 ziellinie [zz], zielnummer [zz] = bhf.Linie (), bhf.Nummer ()
                 zz ++
               }
            })
  nmin:= uint(ker.MaxNat)
  for i:= 0; i < ss; i++ {
    for k:= 0; k < zz; k++ {
      l, n:= startlinie [i], startnummer [i]
      l1, n1:= ziellinie [k], zielnummer [k]
      if ! netz.ExPred2 (func (a Any) bool { b:= a.(*bahnhof.Imp); return b.Linie() == l && b.Nummer() == n },
                         func (a Any) bool { b:= a.(*bahnhof.Imp); return b.Linie() == l1 && b.Nummer() == n1 }) {
        ker.Stop (pack, 4)
      }
      netz.Actualize ()
      n = netz.LenAct ()
      if n < nmin {
        nmin = n
        imin, kmin = uint(i), uint(k)
      }
    }
  }
  l, n:= startlinie [imin], startnummer [imin]
  l1, n1:= ziellinie [kmin], zielnummer [kmin]
  if ! netz.ExPred2 (func (a Any) bool { b:= a.(*bahnhof.Imp); return b.Linie() == l && b.Nummer() == n },
                     func (a Any) bool { b:= a.(*bahnhof.Imp); return b.Linie() == l1 && b.Nummer() == n1 }) {
    ker.Stop (pack, 5)
  }
*/
  netz.Actualize ()
  write (true)
  scr.Colours (col.HintF, col.HintB)
  scr.SwitchTransparence (false)
  na:= netz.LenAct ()
  scr.Write ("kürzeste Verbindung " + nat.String (na) + " Minuten", 0, 0)
  scr.SwitchTransparence (true)
}


func init () {
//
  h:= [...]string { "Start/Ziel auswählen: linke Maustaste   ",
                    "                                        ",
                    "      größer/kleiner: Eingabe-/Rücktaste",
                    "         verschieben: rechte Maustaste  ",
                    "                      oder Pfeiltasten  ",
                    "                                        ",
                    "    Programm beenden: Esc               " }
  help = make ([]string, len (h))
  for i, l:= range (h) { str.Set (&help[i], l) }
  netzAufbauen ()
//  netz.Set (gra.Breadth); netz.Install (wr, wr3)
}
