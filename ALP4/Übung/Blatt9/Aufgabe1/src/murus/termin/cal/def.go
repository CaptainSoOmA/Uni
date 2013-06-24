package cal

// (c) Christian Maurer   v. 130127 - license see murus.go

// SetFormat (p day.Period)

// Das aktuelle Kalenderblatt ist das vom Tag d.
// Seek (d *day.Imp)

// WriteDay (Z, S uint)
// ClearDay (d *day.Imp, Z, S uint)

// Die Folge der Kalenderblätter ist durch Editieren verändert, wobei bei d begonnen wird.
// d ist danach das Datum des zuletzt editierten Kalenderblattes.
// Edit (d *day.Imp, Z, S uint)

// Das aktuelle Stichwort ist das an Position (Z, S) editierte.
// EditWord (Z, S uint)

// Print (Z, S uint)

// Auf alle Daten aus der aktuellen Folge, an denen das aktuelle Stichwort vorkommt, ist B angewandt.
// LookFor (b Op)


// func Index (X Object) Object {
