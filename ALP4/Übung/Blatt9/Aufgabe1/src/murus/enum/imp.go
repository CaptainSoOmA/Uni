package enum

// (c) Christian Maurer   v. 130117 - license see murus.go

import (
  . "murus/obj"
  "murus/str"
  "murus/enum/enumbase"
)
type
  Imp struct {
             *enumbase.Imp
             }
var
  l, s [NEnums][]string


func init () {
//
  l[Title] = []string { "",
                        "Dr.", "Dr.med.", "Dr.med.dent.", "Dr.rer.nat.", "Dr.phil.",
                        "Dr.iur.", "Dr.med.vet.", "Dr.rer.pol.", "Dr.-Ing.",
                        "Prof.Dr.", "Prof.Dr.med.", "Prof.Dr.-Ing." }
  s[Title] = l[Title]

  l[Composer] = []string { "",
                           "Monteverdi, Claudio (1567-1643)",
                           "Frescobaldi, Girolamo (1583-1643)",
                           "Schütz, Heinrich (1585-1672)",
                           "Lully, Jean Baptiste (1632-1687)",
                           "Buxtehude, Dietrich (1637-1707)",
                           "Corelli, Arcangelo (1653-1713)",
                           "Purcell, Henry (1659-1695)",
                           "Scarlatti, Alessandro (1659-1725)",
                           "Couperin, Francois (1668-1733)",
                           "Caldara, Antonio (1670-1736)",
                           "Albinoni, Tomaso (1671-1751)",
                           "Vivaldi, Antonio (1680-1743)",
                           "Telemann, Georg Philipp (1681-1767)",
                           "Rameau, Jean-Philippe (1683-1764)  ",
                           "Bach, Johann Sebastian (1685-1750) ",
                           "Händel, Georg Friedrich (1685-1759)",
                           "Pergolesi, Giovanni Battista (1710-1736)",
                           "Gluck, Christoph Willibald (1714-1787)",
                           "Bach, Philip Emanuel (1714-1788)", // 19, see 15
                           "Haydn, Joseph (1732-1809)",
                           "Boccherini, Luigi (1743-1805)",
                           "Mozart, Wolfgang Amadeus (1756-1791)",
                           "Cherubini, Luigi (1760-1842)",
                           "Beethoven, Ludwig van (1770-1827)",
                           "Spohr, Ludwig (1774-1859)",
                           "Paganini, Niccolo (1782-1840)",
                           "Weber, Carl Maria von (1786-1826)",
                           "Rossini, Gioacchino (1792-1868)",
                           "Schubert, Franz (1797-1828)",
                           "Lortzing, Albert (1801-1851)",
                           "Berlioz, Hector (1803-1869)",
                           "Glinka, Michael (1804-1857)",
                           "Mendelssohn-Bartholdy, Felix (1809-1874)",
                           "Schumann, Robert (1810-1856)",
                           "Chopin, Frederic (1810-1849)",
                           "Liszt, Franz (1811-1886)",
                           "Wagner, Richard (1813-1883)",
                           "Verdi, Giuseppe (1813-1901)",
                           "Franck, Cesar (1822-1890)",
                           "Lalo, Edouard (1823-1892)",
                           "Smetana, Friedrich (1824-1884)",
                           "Bruckner, Anton (1824-1896)",
                           "Strauss, Johann (1825-1899)",
                           "Brahms, Johannes (1833-1897)",
                           "Borodin, Alexander (1834-1887)",
                           "Saint-Saens, Camille (1835-1921)",
                           "Bizet, Georges (1836-1875)",
                           "Mussorgski, Modest (1839-1881)",
                           "Tschaikowskij, Peter (1840-1893)",
                           "Dvorak, Antonin (1841-1904)",
                           "Grieg, Edward (1843-1907)",
                           "Rimskij-Korssakow, Nikolai (1844-1908)",
                           "Janacek, Leos (1854-1928)",
                           "Mahler, Gustav (1860-1911)",
                           "Debussy, Claude (1862-1918)",
                           "Strauß, Richard (1864-1949)",
                           "Sibelius, Jean (1865-1957)",
                           "Pfitzner, Hans (1869-1949)",
                           "Scriabin, Alexander (1872-1915)",
                           "Reger, Max (1873-1916)",
                           "Rachmaninow, Sergej (1873-1947)",
                           "Schönberg, Arnold (1874-1951)",
                           "Ravel, Maurice (1875-1937)",
                           "Falla, Manuel de (1876-1946)",
                           "Bartok, Bela (1881-1945)",
                           "Strawinsky, Igor (1882-1971)",
                           "Webern, Anton von (1883-1945)",
                           "Berg, Alban (1885-1935)",
                           "Furtwängler, Wilhelm (1886-1954)",
                           "Prokofieff, Serge (1891-1953)",
                           "Honegger, Arthur (1892-1955)",
                           "Hindemith, Paul (1895-1963)",
                           "Orff, Carl (1895-1982)",
                           "Blacher, Boris (1903-1975)",
                           "Chatschaturian, Aram (1903-1978)",
                           "Schostakowitsch, Dimitri (1906-1975)",
                           "Fortner, Wolfgang (1907-1987)",
                           "Britten, Benjamin (1913-1976)",
                           "Boulez, Pierre (1925-)",
                           "Henze, Hans Werner (1926-2012)" }
  n:= uint(len (l[Composer]))
  s[Composer] = make ([]string, n)
  s[Composer][0] = ""
  for i:= uint(1); i < n; i++ {
    s[Composer][i] = str.Part (l[Composer][i], 0, str.Pos (l[Composer][i], ','))
  }
  s[Composer][19] = "Bach, Ph.E."

  l[RecordLabel] = []string { "",
                              "2001", "Angel", "BMG", "CBS",
                              "Decca", "Denon", "Deutsche Grammophon",
                              "EMI", "Erato", "Harmonia mundi",
                              "Melodia", "Philips", "Polygram",
                              "Sony", "Supraphon", "Teldec",
                              "UMG", "Warner", "Zyx" }
  s[RecordLabel] = l[RecordLabel]

  l[AudioMedium] = []string { "",
                              "Single Play record", "Long Play record",
                              "Composeract disk", "Digital versatile disc",
                              "Super Audio CD", "Blu-ray disc" }
  s[AudioMedium] = []string { "",
                              "SP", "LP", "CD", "DVD", "SACD", "BD" }

  l[SparsCode] = []string { "AAA",
                            "AAD", "ADD", "DAD", "DDD" }
  s[SparsCode] = l[SparsCode]

  l[Religion] = []string { "keine",
                           "evangelisch", "katholisch", "jüdisch", "muslimisch",
                           "hinduistisch", "buddhistisch", "andere" }
  s[Religion] = l[Religion]

  l[Subject] = []string { "keinFach",
                          "Deutsch", "Englisch", "Französisch", "Italienisch", "Spanisch",
                          "Polnisch", "Russisch", "Türkisch", "Japanisch", "Chinesisch",
                          "Latein", "Griechisch",
                          "Musik", "BildendeKunst", "DarstellendesSpiel",
                          "Politikwissenschaft", "Geschichte", "Geografie",
                          "Sozialwissenschaften", "Psychologie", "Philosophie", "Recht",
                          "Wirtschaftswissenschaft", "Pädagogik", "RechnungswesenUndControlling", "Wirtschaft",
                          "Mathematik", "Physik", "Chemie", "Biologie", "Informatik",
                          "Physiktechnik", "Physiklabortechnik", "Elektrotechnik", "RegenerativeEnergietechnik",
                          "Bautechnik", "Mechatronik", "Metalltechnik_Maschinenbau",
                          "Chemietechnik", "Chemielabortechnik",
                          "Biologietechnik", "Biologielabortechnik", "Biotechnologie", "AgrartechnikMitBiologie",
                          "Gesundheit", "Ernährung", "Medizintechnik", "Umwelttechnik",
                          "Wirtschaftsinformatik", "TechnischeInformatik", "Medizininformatik", "Informationstechnik",
                          "Medientechnik", "GestaltungsMedientechnik", "Gestaltung",
                          "Sport" }
  s[Subject] = []string { "  ",
                          "de", "e ", "f ", "i ", "s ",
                          "p ", "r ", "t ", "j ", "c ",
                          "l ", "g ",
                          "mu", "ku", "ds",
                          "pw", "ge", "gg",
                          "sw", "ps", "ph", "re",
                          "ww", "pa", "rc", "wi",
                          "ma", "ph", "ch", "bi", "in",
                          "pt", "pt", "et", "rt",
                          "bt", "me", "mm",
                          "ct", "c",
                          "bt", "bt", "bt", "ab",
                          "gs", "er", "me", "ut",
                          "wi", "ti", "mi", "it",
                          "mt", "gm", "gt",
                          "sp" }

  l[LexicalCategory] = []string { "",
                                 "Substantiv", "Adjektiv", "Pronomen", "Numerale",
                                 "Verb", "Adverb", "Präposition", "Konjunktion", "Interjektion" }
  s[LexicalCategory] = []string { "",
                                  "Subst.", "Adj.", "Pron.", "Num.",
                                  "Verb", "Adv.", "Präp.", "Konj.", "Interj." }

  l[Casus] = []string { "",
                        "Nominativ", "Genitiv", "Dativ", "Akkusativ", "Ablativ" }
  s[Casus] = []string { "",
                        "Nom.", "Gen.", "Dat.", "Akk.", "Abl." }

  l[Genus] = []string { "",
                        "masc.", "fem.", "neut." }
//                        "maskulinum", "femininum", "neutrum" }
  s[Genus] = []string { "",
                        "m.", "f.", "n." }

  l[Persona] = []string { "",
                          "1.", "2.", "3." }
  s[Persona] = l[Persona]

  l[Numerus] = []string { "",
                          "Sing.", "Plur." }
//                          "Singular", "Plural" }
  s[Numerus] = []string { "",
                          "Sg.", "Pl." }

  l[Tempus] = []string { "",
                         "Präsens", "Imperfekt", "FuturI", "Perfekt", "Plusquamp.", "FuturII" }
  s[Tempus] = []string { "",
                         "Präs.", "Impf.", "Fut.I", "Perf.", "Plusq.", "Fut.II" }

  l[Modus] = []string { "",
                        "Indikativ", "Konjunktiv" }
  s[Modus] = []string { "",
                        "Ind.", "Konj." }

  l[GenusVerbi] = []string { "",
                             "Aktiv", "Passiv" }
  s[GenusVerbi] = []string { "",
                             "Akt.", "Pass." }
}


func (x *Imp) imp (Y Object) *enumbase.Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok || x.Imp.Typ () != y.Imp.Typ () { TypeNotEqPanic (x.Imp, y.Imp) }
  return y.Imp
}


func New (e Enum) *Imp {
//
  if e >= NEnums { TypePanic() }
  return &Imp { enumbase.New (byte(e), [enumbase.NFormats][]string { s[e], l[e] }) }
}


func (x *Imp) Eq (Y Object) bool {
//
  return x.Imp.Eq (x.imp (Y))
}


func (x *Imp) Copy (Y Object) {
//
  x.Imp.Copy (x.imp (Y))
}


func (x *Imp) Less (Y Object) bool {
//
  return x.Imp.Less (x.imp (Y))
}


func init () {
//
  var _ Enumerator = New (Enum(0))
}
