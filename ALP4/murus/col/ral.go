package col

// (c) Christian Maurer   v. 120819 - license see murus.go

var // RAL-Farben
  Grünbeige, Beige, Sandgelb, Signalgelb, Goldgelb, // 1000 .. 1004
  Honiggelb, Maisgelb, Narzissengelb, Braunbeige, Zitronengelb, // 1005 .. 1007, 1011 .. 1012
  Perlweiß, Elfenbein, Hellelfenbein, Schwefelgelb, Safrangelb, // 1013 .. 1017
  Zinkgelb, Graubeige, Olivgelb, Rapsgelb, // 1018 .. 1021
  Verkehrsgelb, Ockergelb, Leuchtgelb, Currygelb, Melonengelb, // 1023 .. 1024, 1026 .. 1028
  Ginstergelb, Dahliengelb, Pastelgelb, // 1032 .. 1034

  Gelborange, Rotorange, Blutorange, Pastellorange, Reinorange, Leuchtorange, // 2000 .. 2005
  Leuchthellorange, Hellrotorange, Verkehrsorange, // 2007 .. 2009
  Signalorange, Tieforange, Lachsorange, // 2010 .. 2012

  Feuerrot, Signalrot, Karminrot, Rubinrot, Purpurrot, Weinrot, // 3000 .. 3005
  Schwarzrot, Oxidrot, Braunrot, Beigerot, Tomatenrot, // 3007, 3009, 3011 .. 3013
  Altrosa, Hellrosa, Korallenrot, Rose, Erdbeerrot, // 3014 .. 3018
  Verkehrsrot, Lachsrot, Leuchtrot, // 3020, 3022, 3024
  Leuchthellrot, Himbeerrot, Orientrot, // 3026, 3027, 3031

  Rotlila, Rotmagenta, Erikamagenta, Bordeauxmagenta, Blaulila, // 4001 .. 4005
  Verkehrspurpur, Purpurmagenta, Signalviolett, Pastelviolett, Telemagenta, // 4006 .. 4010

  Violettblau, Grünblau, Ultramarinblau, Saphirblau, Schwarzblau, Signalblau, // 5000 .. 5005
  Brillantblau, Graublau, Azurblau, Enzianblau, Stahlblau, // 5007 .. 5011
  Lichtblau, Kobaltblau, Taubenblau, Himmelblau, // 5012 .. 5015
  Verkehrsblau, Türkisblau, Capriblau, Ozeanblau, // 5017 .. 5020
  Wasserblau, Nachtblau, Fernblau, Pastellblau, // 5021 .. 5024

  Patinagrün, Smaragdgrün, Laubgrün, Olivgrün, Blaugrün, Moosgrün, // 6000 .. 6005
  Grauoliv, Flaschengrün, Braungrün, Tannengrün, Grasgrün, Resedagrün, // 6006 .. 6011
  Schwarzgrün, Schilfgrün, Gelboliv, Schwarzoliv, Cyangrün, Maigrün, // 6012 .. 6017
  Gelbgrün, Weißgrün, Chromoxidgrün, Blassgrün, Braunoliv, // 6018 .. 6022
  Verkehrsgrün, Farngrün, Opalgrün, Lichtgrün, Kieferngrün, Minzgrün, // 6024 .. 6029
  Signalgrün, Minttürkis, Pasteltürkis, // 6032 .. 6034

  Fehgrau, Silbergrau, Olivgrau, Moosgrau, Signalgrau, Mausgrau, Beigegrau, // 7000 .. 7006
  Khakigrau, Grüngrau, Zeltgrau, Eisengrau, Basaltgrau, Braungrau, // 7008 .. 7013
  Schiefergrau, Anthrazitgrau, // 7015 .. 7016
  Schwarzgrau, Umbragrau, Betongrau, Graphitgrau, Granitgrau, // 7021 .. 7024, 7026
  Steingrau, Blaugrau, Kieselgrau, Zementgrau, Gelbgrau, Lichtgrau, // 7030 .. 7035
  Platingrau, Staubgrau, Achatgrau, Quarzgrau, Fenstergrau, // 7036 .. 7040
  VerkehrsgrauA, VerkehrsgrauB, Seidengrau, Telegrau1, Telegrau2, Telegrau4, // 7042 .. 7047

  Grünbraun, Ockerbraun, Signalbraun, Lehmbraun, Kupferbraun, // 8000 .. 8004
  Rehbraun, Olivbraun, Nussbraun, Rotbraun, // 8007 .. 8008, 8011 .. 8012
  Sepiabraun, Kastanienbraun, Mahagonibraun, Schokoladenbraun, // 8014 .. 8017
  Graubraun, Schwarzbraun, Orangebraun, Beigebraun, Blassbraun, Terrabraun, // 8019, 8022 .. 8025, 8028

  Cremeweiß, Grauweiß, Signalweiß, Signalschwarz, Tiefschwarz, // 9001 .. 9005
  Aluminiumweiß, Aluminiumgrau, Reinweiß, Graphitschwarz, // 9006 .. 9007, 9010 .. 9011
  Verkehrweiß, Verkehrschwarz, Papyrusweiß Colour // 9016 .. 9018


func init () {
//
  Grünbeige =        colour3 (214, 199, 148) // RAL 1000
  Beige =            colour3 (217, 186, 140) // RAL 1001
  Sandgelb =         colour3 (214, 176, 117) // RAL 1002
  Signalgelb =       colour3 (252, 163,  41) // RAL 1003
  Goldgelb =         colour3 (227, 150,  36) // RAL 1004
  Honiggelb =        colour3 (201, 135,  33) // RAL 1005
  Maisgelb =         colour3 (224, 130,  31) // RAL 1006
  Narzissengelb =    colour3 (227, 122,  31) // RAL 1007
  Braunbeige =       colour3 (173, 122,  79) // RAL 1011
  Zitronengelb =     colour3 (227, 184,  56) // RAL 1012
  Perlweiß =         colour3 (255, 245, 227) // RAL 1013
  Elfenbein =        colour3 (240, 214, 171) // RAL 1014
  Hellelfenbein =    colour3 (252, 235, 204) // RAL 1015
  Schwefelgelb =     colour3 (255, 245,  66) // RAL 1016
  Safrangelb =       colour3 (255, 171,  89) // RAL 1017
  Zinkgelb =         colour3 (255, 214,  77) // RAL 1018
  Graubeige =        colour3 (163, 140, 122) // RAL 1019
  Olivgelb =         colour3 (156, 143,  97) // RAL 1020
  Rapsgelb =         colour3 (252, 189,  31) // RAL 1021
  Verkehrsgelb =     colour3 (252, 184,  33) // RAL 1023
  Ockergelb =        colour3 (181, 140,  79) // RAL 1024
  Leuchtgelb =       colour3 (255, 255,  10) // RAL 1026
  Currygelb =        colour3 (153, 117,  33) // RAL 1027
  Melonengelb =      colour3 (255, 140,  26) // RAL 1028
  Ginstergelb =      colour3 (227, 163,  41) // RAL 1032
  Dahliengelb =      colour3 (255, 148,  54) // RAL 1033
  Pastelgelb =       colour3 (247, 153,  92) // RAL 1034

  Gelborange =       colour3 (224,  94,  31) // RAL 2000
  Rotorange =        colour3 (186,  46,  33) // RAL 2001
  Blutorange =       colour3 (204,  36,  28) // RAL 2002
  Pastellorange =    colour3 (255,  99,  54) // RAL 2003
  Reinorange =       colour3 (242,  59,  28) // RAL 2004
  Leuchtorange =     colour3 (252,  28,  20) // RAL 2005
  Leuchthellorange = colour3 (255, 117,  33) // RAL 2007
  Hellrotorange =    colour3 (250,  79,  41) // RAL 2008
  Verkehrsorange =   colour3 (235,  59,  28) // RAL 2009
  Signalorange =     colour3 (212,  69,  41) // RAL 2010
  Tieforange =       colour3 (237,  92,  41) // RAL 2011
  Lachsorange =      colour3 (222,  82,  71) // RAL 2012

  Feuerrot =         colour3 (171,  31,  28) // RAL 3000
  Signalrot =        colour3 (163,  23,  26) // RAL 3001
  Karminrot =        colour3 (163,  26,  26) // RAL 3002
  Rubinrot =         colour3 (138,  18,  20) // RAL 3003
  Purpurrot =        colour3 (105,  15,  20) // RAL 3004
  Weinrot =          colour3 ( 79,  18,  26) // RAL 3005
  Schwarzrot =       colour3 ( 46,  18,  26) // RAL 3007
  Oxidrot =          colour3 ( 94,  33,  33) // RAL 3009
  Braunrot =         colour3 (120,  20,  23) // RAL 3011
  Beigerot =         colour3 (204, 130, 115) // RAL 3012
  Tomatenrot =       colour3 (150,  31,  28) // RAL 3013
  Altrosa =          colour3 (217, 102, 117) // RAL 3014
  Hellrosa =         colour3 (232, 156, 181) // RAL 3015
  Korallenrot =      colour3 (166,  36,  38) // RAL 3016
  Rose =             colour3 (209,  54,  84) // RAL 3017
  Erdbeerrot =       colour3 (207,  41,  66) // RAL 3018
  Verkehrsrot =      colour3 (199,  23,  18) // RAL 3020
  Lachsrot =         colour3 (217,  89,  79) // RAL 3022
  Leuchtrot =        colour3 (252,  10,  28) // RAL 3024
  Leuchthellrot =    colour3 (252,  20,  20) // RAL 3026
  Himbeerrot =       colour3 (181,  18,  51) // RAL 3027
  Orientrot =        colour3 (166,  28,  46) // RAL 3031

  Rotlila =          colour3 (130,  64, 128) // RAL 4001
  Rotmagenta =       colour3 (143,  38,  64) // RAL 4002
  Erikamagenta =     colour3 (201,  56, 140) // RAL 4003
  Bordeauxmagenta =  colour3 ( 92,   8,  43) // RAL 4004
  Blaulila =         colour3 ( 99,  61, 156) // RAL 4005
  Verkehrspurpur =   colour3 (145,  15, 102) // RAL 4006
  Purpurmagenta =    colour3 ( 56,  10,  46) // RAL 4007
  Signalviolett =    colour3 (125,  31, 122) // RAL 4008
  Pastelviolett =    colour3 (158, 115, 148) // RAL 4009
  Telemagenta =      colour3 (191,  23, 115) // RAL 4010

  Violettblau =      colour3 ( 23,  51, 107) // RAL 5000
  Grünblau =         colour3 ( 10,  51,  84) // RAL 5001
  Ultramarinblau =   colour3 (  0,  15, 117) // RAL 5002
  Saphirblau =       colour3 (  0,  23,  69) // RAL 5003
  Schwarzblau =      colour3 (  3,  13,  31) // RAL 5004
  Signalblau =       colour3 (  0,  46, 122) // RAL 5005
  Brillantblau =     colour3 ( 38,  79, 135) // RAL 5007
  Graublau =         colour3 ( 26,  41,  56) // RAL 5008
  Azurblau =         colour3 ( 23,  69, 112) // RAL 5009
  Enzianblau =       colour3 (  0,  43, 112) // RAL 5010
  Stahlblau =        colour3 (  3,  20,  46) // RAL 5011
  Lichtblau =        colour3 ( 41, 115, 184) // RAL 5012
  Kobaltblau =       colour3 (  0,  18,  69) // RAL 5013
  Taubenblau =       colour3 ( 77, 105, 153) // RAL 5014
  Himmelblau =       colour3 ( 23,  97, 171) // RAL 5015
  Verkehrsblau =     colour3 (  0,  59, 128) // RAL 5017
  Türkisblau =       colour3 ( 56, 148, 130) // RAL 5018
  Capriblau =        colour3 ( 10,  66, 120) // RAL 5019
  Ozeanblau =        colour3 (  5,  51,  51) // RAL 5020
  Wasserblau =       colour3 ( 26, 122,  99) // RAL 5021
  Nachtblau =        colour3 (  0,   8,  79) // RAL 5022
  Fernblau =         colour3 ( 46,  82, 143) // RAL 5023
  Pastellblau =      colour3 ( 87, 140, 181) // RAL 5024

  Patinagrün =       colour3 ( 51, 120,  84) // RAL 6000
  Smaragdgrün =      colour3 ( 38, 102,  41) // RAL 6001
  Laubgrün =         colour3 ( 38,  87,  33) // RAL 6002
  Olivgrün =         colour3 ( 61,  69,  46) // RAL 6003
  Blaugrün =         colour3 ( 13,  59,  46) // RAL 6004
  Moosgrün =         colour3 ( 10,  56,  31) // RAL 6005
  Grauoliv =         colour3 ( 41,  43,  36) // RAL 6006
  Flaschengrün =     colour3 ( 28,  38,  23) // RAL 6007
  Braungrün =        colour3 ( 33,  33,  26) // RAL 6008
  Tannengrün =       colour3 ( 23,  41,  28) // RAL 6009
  Grasgrün =         colour3 ( 54, 105,  38) // RAL 6010
  Resedagrün =       colour3 ( 94, 125,  79) // RAL 6011
  Schwarzgrün =      colour3 ( 31,  46,  43) // RAL 6012
  Schilfgrün =       colour3 (117, 115,  79) // RAL 6013
  Gelboliv =         colour3 ( 51,  48,  38) // RAL 6014
  Schwarzoliv =      colour3 ( 41,  43,  38) // RAL 6015
  Cyangrün =         colour3 ( 15, 112,  51) // RAL 6016
  Maigrün =          colour3 ( 64, 130,  54) // RAL 6017
  Gelbgrün =         colour3 ( 79, 168,  51) // RAL 6018
  Weißgrün =         colour3 (191, 227, 186) // RAL 6019
  Chromoxidgrün =    colour3 ( 38,  56,  41) // RAL 6020
  Blassgrün =        colour3 (133, 166, 122) // RAL 6021
  Braunoliv =        colour3 ( 43,  38,  28) // RAL 6022
  Verkehrsgrün =     colour3 ( 36, 145,  64) // RAL 6024
  Farngrün =         colour3 ( 74, 110,  51) // RAL 6025
  Opalgrün =         colour3 ( 10,  92,  51) // RAL 6026
  Lichtgrün =        colour3 (125, 204, 189) // RAL 6027
  Kieferngrün =      colour3 ( 38,  74,  51) // RAL 6028
  Minzgrün =         colour3 ( 18, 120,  38) // RAL 6029
  Signalgrün =       colour3 ( 41, 138,  64) // RAL 6032
  Minttürkis =       colour3 ( 66, 140, 120) // RAL 6033
  Pasteltürkis =     colour3 (125, 189, 181) // RAL 6034

  Fehgrau =          colour3 (115, 133, 145) // RAL 7000
  Silbergrau =       colour3 (135, 148, 166) // RAL 7001
  Olivgrau =         colour3 (122, 117,  97) // RAL 7002
  Moosgrau =         colour3 (112, 112,  97) // RAL 7003
  Signalgrau =       colour3 (156, 156, 166) // RAL 7004
  Mausgrau =         colour3 ( 97, 105, 105) // RAL 7005
  Beigegrau =        colour3 (107,  97,  87) // RAL 7006
  Khakigrau =        colour3 (105,  84,  56) // RAL 7008
  Grüngrau =         colour3 ( 77,  82,  74) // RAL 7009
  Zeltgrau =         colour3 ( 74,  79,  74) // RAL 7010
  Eisengrau =        colour3 ( 64,  74,  84) // RAL 7011
  Basaltgrau =       colour3 ( 74,  84,  89) // RAL 7012
  Braungrau =        colour3 ( 71,  66,  56) // RAL 7013
  Schiefergrau =     colour3 ( 61,  66,  82) // RAL 7015
  Anthrazitgrau =    colour3 ( 38,  46,  56) // RAL 7016
  Schwarzgrau =      colour3 ( 26,  33,  41) // RAL 7021
  Umbragrau =        colour3 ( 61,  61,  59) // RAL 7022
  Betongrau =        colour3 (122, 125, 117) // RAL 7023
  Graphitgrau =      colour3 ( 48,  56,  69) // RAL 7024
  Granitgrau =       colour3 ( 38,  51,  56) // RAL 7026
  Steingrau =        colour3 (145, 143, 135) // RAL 7030
  Blaugrau =         colour3 ( 77,  92, 107) // RAL 7031
  Kieselgrau =       colour3 (189, 186, 171) // RAL 7032
  Zementgrau =       colour3 (122, 130, 117) // RAL 7033
  Gelbgrau =         colour3 (143, 135, 112) // RAL 7034
  Lichtgrau =        colour3 (212, 217, 219) // RAL 7035
  Platingrau =       colour3 (158, 150, 156) // RAL 7036
  Staubgrau =        colour3 (122, 125, 128) // RAL 7037
  Achatgrau =        colour3 (186, 189, 186) // RAL 7038
  Quarzgrau =        colour3 ( 97,  94,  89) // RAL 7039
  Fenstergrau =      colour3 (158, 163, 176) // RAL 7040
  VerkehrsgrauA =    colour3 (143, 150, 153) // RAL 7042
  VerkehrsgrauB =    colour3 ( 64,  69,  69) // RAL 7043
  Seidengrau =       colour3 (194, 191, 184) // RAL 7044
  Telegrau1 =        colour3 (143, 148, 158) // RAL 7045
  Telegrau2 =        colour3 (120, 130, 140) // RAL 7046
  Telegrau4 =        colour3 (217, 214, 219) // RAL 7047

  Grünbraun =        colour3 (125,  92,  56) // RAL 8000
  Ockerbraun =       colour3 (145,  82,  46) // RAL 8001
  Signalbraun =      colour3 (110,  59,  48) // RAL 8002
  Lehmbraun =        colour3 (115,  59,  36) // RAL 8003
  Kupferbraun =      colour3 (133,  56,  43) // RAL 8004
  Rehbraun =         colour3 ( 94,  51,  31) // RAL 8007
  Olivbraun =        colour3 ( 99,  61,  36) // RAL 8008
  Nussbraun =        colour3 ( 71,  38,  28) // RAL 8011
  Rotbraun =         colour3 ( 84,  31,  31) // RAL 8012
  Sepiabraun =       colour3 ( 56,  38,  28) // RAL 8014
  Kastanienbraun =   colour3 ( 77,  31,  28) // RAL 8015
  Mahagonibraun =    colour3 ( 61,  31,  28) // RAL 8016
  Schokoladenbraun = colour3 ( 46,  28,  28) // RAL 8017
  Graubraun =        colour3 ( 43,  38,  41) // RAL 8019
  Schwarzbraun =     colour3 ( 13,   8,  13) // RAL 8022
  Orangebraun =      colour3 (156,  69,  41) // RAL 8023
  Beigebraun =       colour3 (110,  64,  48) // RAL 8024
  Blassbraun =       colour3 (102,  74,  61) // RAL 8025
  Terrabraun =       colour3 ( 64,  46,  33) // RAL 8028

  Cremeweiß =        colour3 (255, 252, 240) // RAL 9001
  Grauweiß =         colour3 (240, 237, 230) // RAL 9002
  Signalweiß =       colour3 (255, 255, 255) // RAL 9003
  Signalschwarz =    colour3 ( 28,  28,  33) // RAL 9004
  Tiefschwarz =      colour3 (  3,   5,  10) // RAL 9005
  Aluminiumweiß =    colour3 (166, 171, 181) // RAL 9006
  Aluminiumgrau =    colour3 (125, 122, 120) // RAL 9007
  Reinweiß =         colour3 (250, 255, 255) // RAL 9010
  Graphitschwarz =   colour3 ( 13,  18,  26) // RAL 9011
  Verkehrweiß =      colour3 (252, 255, 255) // RAL 9016
  Verkehrschwarz =   colour3 ( 20,  23,  28) // RAL 9017
  Papyrusweiß =      colour3 (219, 227, 222) // RAL 9019
}
