package day

// (c) Christian Maurer   v. 130526 - license see murus.go

import (
  . "murus/ker"; . "murus/obj"; "murus/rand"; "murus/str"
  "murus/kbd"
  "murus/col"; "murus/scr"; "murus/box"; "murus/errh"
  "murus/font"; "murus/pbox"
  "murus/nat"
)
const (
  pack = "day"
  emptyYear = uint(1879)
  startYear = emptyYear + 1
  limitYear = uint(2020) // Yy: 0..20 -> 2000..2020; 21..99 -> 1921..1999
  endYear   = uint(2058) // emptyYear + 179, // got to change that, if I am 113 years old
                         // 179 Jahre < MAX (uint16) Tage < 180 Jahre
  maxLength = 18 // length[Dd_M_yyyy] == 2 + 1 + 10 (== len("Donnerstag")) + 1 + 4
  maxmonth = uint(12)
  maxday = uint(31)
  maxCode = uint16(65379) // 31.12.2058
)
type
  Imp struct {
         day,
       month,
        year uint
         fmt Format
      cF, cB col.Colour
        font font.Font
             }
var (
  today Calendarday = New()
  currentCentury uint
  todayCode uint16
  nameMonth = [13]string { "xxxxxxxxx",
    "Januar   ", "Februar  ", "März     ", "April    ", "Mai      ", "Juni     ",
    "Juli     ", "August   ", "September", "Oktober  ", "November ", "Dezember " }
  WdText = [NWeekdays]string { "Montag    ", "Dienstag  ", "Mittwoch  ",
                 "Donnerstag", "Freitag   ", "Sonnabend ", "Sonntag   " }
  WdShorttext = [NWeekdays]string { "Mo", "Di", "Mi", "Do", "Fr", "Sa", "So" }
  length = []uint {
    Dd:          2,
    Dd_mm_:      6,
    Dd_mm_yy:    8,
    Yymmdd:      6,
    Yyyymmdd:    8,
    Dd_mm_yyyy: 10,
    Dd_M:       maxLength - 5,
    Dd_M_yyyy:  maxLength, // see above
    Yy:          2,
    Yyyy:        4,
    Wd:          2,
    WD:         10,
    Mmm:         3,
    M:           9,
    Myyyy:      14,
    Wn:          2,
    WN:          8,
    WNyyyy:     13,
    Qu:          6,
  }
  bx *box.Imp = box.New()
  pbx *pbox.Imp = pbox.New()
  actualDay, actualMonth, actualYear = maxday, maxmonth, emptyYear
  Codeyear = emptyYear
  yearcode = uint16(0)
  actualHolidayYear = emptyYear
  holiday [1+maxday][1+maxmonth]bool
  actualSummer = emptyYear
  op Op = attribute
//  carnival *Imp
)


func (x *Imp) imp (Y Any) *Imp {
//
  y, ok:= Y.(*Imp)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}


func New() *Imp {
//
  return &Imp { maxday, maxmonth, emptyYear, Dd_mm_yy, col.ScreenF, col.ScreenB, font.Roman }
}


func New3 (d, m, y uint) *Imp {
//
  x:= New ()
  x.Set (d, m, y)
  return x
}


func (D *Imp) Randomize() {
//
  const n = uint16(14976) // Code of 1.1.1921
  D.Decode (Encode (n + uint16(rand.Natural (uint(todayCode + uint16(1) - n)))))
}


func (x *Imp) Empty () bool {
//
  return x.year == emptyYear
}


func (D *Imp) Clr () {
//
  D.day, D.month, D.year = maxday, maxmonth, emptyYear
}


func (D *Imp) SetMin () {
//
  D.day, D.month, D.year = 1, 1, startYear
}


func (D *Imp) SetMax () {
//
  D.day, D.month, D.year = maxday, maxmonth, endYear
}


func (D *Imp) Actualize () {
//
  D.day, D.month, D.year = ActualizeDate ()
}


func (D *Imp) Set (Tag, Monat, Jahr uint) bool {
//
  if D.defined (Tag, Monat, Jahr) {
    return true
  }
  D.Clr ()
  return false
}


func (x *Imp) Copy (Y Object) {
//
  y:= x.imp (Y)
  x.day, x.month, x.year = y.day, y.month, y.year
  x.fmt = y.fmt
  x.cF, x.cB = y.cF, y.cB
}


func (x *Imp) Clone () Object {
//
  y:= New ()
  y.day, y.month, y.year = x.day, x.month, x.year
  y.fmt = x.fmt
  y.cF, y.cB = x.cF, x.cB
  return y
}


func (x *Imp) Eq (Y Object) bool {
//
  y:= x.imp (Y)
  if x.year == emptyYear { return y.year == emptyYear }
  if y.year == emptyYear { return false }
  return x.day == y.day && x.month == y.month && x.year == y.year
}


func (x *Imp) Less (Y Object) bool {
//
  y:= x.imp (Y)
  if x.year == emptyYear {
    return y.year != emptyYear
  }
  if y.year == emptyYear { return false }
  if x.year == y.year {
    if x.month == y.month {
      return x.day < y.day
    } else {
      return x.month < y.month
    }
  }
  return x.year < y.year
}


func (x *Imp) LessInYear (Y Calendarday) bool {
//
  y:= x.imp (Y)
  if x.year == emptyYear {
    return x.year != emptyYear
  }
  if x.year == emptyYear { return false }
  if x.month == y.month {
    if x.day == y.day {
      return x.year < y.year
    } else {
      return x.day < y.day
    }
  }
  return x.month < y.month
}


func (x *Imp) Equiv (Y Calendarday, p Period) bool {
//
  y:= x.imp (Y)
  if x.year == emptyYear {
    return y.year == emptyYear
  }
  if y.year == emptyYear { return false }
  switch p { case Daily:
    return x.Eq (y)
  case Weekly:
    c:= internalCode (*x)
    w:= (c + 2) % 7
    c1:= internalCode (*y)
    w1:= (c1 + 2) % 7
    if c <= c1 {
      if c1 - c < 7 {
        return w <= w1
      }
    } else if c - c1 < 7 {
      return w > w1
    }
    break
  case Monthly:
    return x.month == y.month && x.year == y.year
  case Quarterly:
    return ((x.month - 1) / 3 == (y.month - 1) / 3) && x.year == y.year
  case HalfYearly:
    return ((x.month - 1) / 6 == (y.month - 1) / 6) && x.year == y.year
  case Yearly:
    return x.year == y.year
  case Decadic:
    return x.year % 10 == y.year % 10
  }
  return false
}


func (x *Imp) IsBeginning (p Period) bool {
//
  switch p { case Daily:
    break
  case Weekly:
    return x.Weekday (Daily) == Monday
  case Monthly:
    return x.day == 1
  case Quarterly:
    return x.day == 1 && 3 * (x.month - 1) / 3 + 1 == x.month
  case HalfYearly:
    return x.day == 1 && 6 * (x.month - 1) / 6 + 1 == x.month
  case Yearly:
    return uint8 (x.day) * uint8 (x.month) == 1
  case Decadic:
    return uint8 (x.day) * uint8 (x.month) == 1 && x.year % 10 == 0
  }
  return true
}


func (x *Imp) SetBeginning (p Period) {
//
  if x.year == emptyYear { return }
  switch p { case Daily:
    return
  case Weekly:
    for w:= x.Weekday (Daily); w > Monday; w-- {
      x.Dec (Daily)
    }
  case Monthly:
    x.day = 1
  case Quarterly:
    x.day, x.month = 1, 3 * (x.month - 1) / 3 + 1
  case HalfYearly:
    x.day, x.month = 1, 6 * (x.month - 1) / 6 + 1
  case Yearly:
    x.day, x.month = 1, 1
  case Decadic:
    if 10 * x.year / 10 > emptyYear {
      x.day, x.month, x.year = 1, 1, 10 * x.year / 10
    }
  }
}


func (x *Imp) SetEnd (p Period) {
//
  if x.year == emptyYear { return }
  switch p { case Daily:
    return
  case Weekly:
    for w:= x.Weekday (Daily); w < Sunday; w++ {
      x.Inc (Daily)
    }
  case Monthly:
    x.day = daysInMonth (*x)
  case Quarterly:
    x.month = 3 * (((x.month - 1) / 3) + 1)
    x.day = daysInMonth (*x)
  case HalfYearly:
    x.month = 6 * (((x.month - 1) / 6) + 1)
    x.day = daysInMonth (*x)
  case Yearly:
    x.day = maxday
    x.month = maxmonth
  case Decadic:
    if x.year + 9 <= endYear + x.year % 10 {
      x.day = maxday
      x.month = maxmonth
      x.year += 9 - x.year % 10
    }
  }
}


func isLeapYear (y uint) bool {
//
/*
  if y % 400 == 0 {
    return true
  else if y % 100 != 0 {
    for the range of this implementation we only need:
*/
  if y == 1900 {
    return false
  }
  return (y % 4) == 0 // emptyYear: false
}


func daysInMonth (x Imp) uint {
//
  if x.year == actualYear {
    if x.month == actualMonth {
      return actualDay
    } else {
      actualMonth = x.month
    }
  } else {
    actualMonth, actualYear = x.month, x.year
  }
  if x.month == 2 {
    actualDay = 28
    if isLeapYear (x.year) { actualDay ++ }
  } else if x.month / 8 == x.month % 2 { // Fingerknöchelprinzip!
    actualDay = 30 // maxday - 1
  } else {
    actualDay = maxday
  }
  return actualDay
}


func daysInYear (y uint) uint16 {
//
  d:= uint16(365)
  if isLeapYear (y) {
    d ++
  }
  return d
}


func internalCode (x Imp) uint16 {
//
  if x.year == emptyYear { return 0 }
  code:= uint16(x.day)
  for x.month > 1 {
    x.month --
    code += uint16(daysInMonth (x))
  }
  if x.year != Codeyear {
    Codeyear = x.year
    yearcode = 0
    for x.year > startYear {
      x.year --
      yearcode += daysInYear (x.year)
    }
  }
  code += yearcode
  return code
}


func (x *Imp) Actual () bool {
//
  return internalCode (*x) == todayCode
}


func (x *Imp) Elapsed () bool {
//
  if x.year == emptyYear {
    return false
  }
  return internalCode (*x) < todayCode
}


func (x *Imp) Distance (Y Calendarday) uint {
//
  y:= x.imp (Y)
  if x.year == emptyYear || y.year == emptyYear { return MaxNat }
  c:= internalCode (*x)
  c1:= internalCode (*y)
  if c < c1 {
    return uint(c1 - c)
  }
  return uint(c - c1)
}


func (x *Imp) NumberOfDays () uint {
//
  if x.year == emptyYear { return 0 }
  return uint(daysInYear (x.year))
}


func (x *Imp) OrdDay () uint {
//
  n:= uint(0)
  if x.year != emptyYear {
    y:= new (Imp)
    y.day, y.month, y.year = 1, 1, x.year
    for y.month < x.month {
      n += uint(daysInMonth (*y))
      y.month ++
    }
    n += uint(x.day)
  }
  return n
}


func (x *Imp) Inc (p Period) {
//
  if x.year == emptyYear { return }
  t:= daysInMonth (*x)
  d, m, y:= x.day, x.month, x.year
  switch p { case Daily:
    if d < t {
      d ++
    } else {
      d = 1
      if m < maxmonth {
        m ++
      } else if y < endYear {
        y ++
        m = 1
      } else {
        return
      }
    }
  case Weekly:
    if d + 7 <= t {
      d += 7
    } else if m < maxmonth {
      m ++
      d -= t - 7
    } else if y < endYear {
      y ++
      m = 1
      d -= 24
    } else {
      return
    }
  case Monthly:
    if m < maxmonth {
      m++
    } else if y < endYear {
      m = 1
      y ++
    } else {
      return
    }
  case Quarterly:
    if m < 10 {
      m += 3
    } else if y < endYear {
      m -= 9
      y ++
    } else {
      return
    }
  case HalfYearly:
    if m < 7 {
      m += 6
    } else if y < endYear {
      m -= 6
      y ++
    } else {
      return
    }
  case Yearly:
    if y < endYear {
      y ++
    } else {
      return
    }
  case Decadic:
    if y <= endYear - 10 {
      y += 10
    } else {
      return
    }
  }
  x.day, x.month, x.year = d, m, y
  t = daysInMonth (*x)
  if x.day > t { x.day = t }
}


func (x *Imp) Inc1 (n uint) {
//
  if x.year == emptyYear { return }
  d, m, y:= x.day, x.month, x.year
  for n > 0 {
    x.Inc (Daily)
    if x.Empty () {
      x.day, x.month, x.year = d, m, y
      return
    }
    n --
  }
}


func (x *Imp) Dec (p Period) {
//
  if x.year == emptyYear { return }
  t:= daysInMonth (*x)
  d, m, y:= x.day, x.month, x.year
  switch p { case Daily:
    if d > 1 {
      d --
    } else if m > 1 {
      m --
      d = daysInMonth (*x)
    } else {
      y --
      m = maxmonth
      d = maxday
      if y == emptyYear {
        return
      }
    }
  case Weekly:
    if d > 7 {
      d -= 7
    } else if m > 1 {
      m --
      t = daysInMonth (*x)
      d += t - 7
    } else {
      y --
      m = maxmonth
      d += 24
      if y == emptyYear { return }
    }
  case Monthly:
    if m > 1 {
      m --
    } else {
      y --
      m = maxmonth
      if y == emptyYear { return }
    }
  case Quarterly:
    if m > 3 {
      m -= 3
    } else {
      y --
      m += 9
      if y == emptyYear { return }
    }
  case HalfYearly:
    if m > 6 {
      m -= 6
    } else {
      y --
      m += 6
      if y == emptyYear { return }
    }
  case Yearly:
    y --
    if y == emptyYear { return }
  case Decadic:
    if y > emptyYear + 10 {
      y -= 10
    } else {
      return
    }
  }
  x.day, x.month, x.year = d, m, y
  t = daysInMonth (*x)
  if x.day > t { x.day = t }
}


func (x *Imp) Change (c kbd.Comm, d uint) {
//
  if x.year == emptyYear { return }
  if Period(d) >= NPeriods { d = uint(NPeriods) - 1 }
  p:= NPeriods
  if Period(d) < p { p = Period(d) }
  switch c { case kbd.Enter, kbd.Esc:
    return
  case kbd.Right, kbd.Down:
    x.Inc (p)
  case kbd.Left, kbd.Up:
    x.Dec (p)
  case kbd.Pos1:
    x.SetBeginning (p)
  case kbd.End:
    x.SetEnd (p)
  case kbd.Here:
    x.Actualize ()
    x.SetBeginning (p)
  }
}


func weekday (x Imp) Weekday {
//
  return Weekday((internalCode (x) + uint16(Wednesday)) % 7) // The day with code 0, 31.12.1879, was a Wednesday
}


func (x *Imp) Weekday (p Period) Weekday {
//
  var d Imp
  d.day, d.month, d.year = x.day, x.month, x.year
  switch p { case Daily:
    ;
  case Weekly:
    return Monday
  case Monthly:
    d.day = 1
  case Quarterly:
    d.day, d.month = 1, 3 * (d.month - 1) / 3 + 1
  case HalfYearly:
    d.day, d.month = 1, 6 * (d.month - 1) / 6 + 1
  case Yearly:
    d.day, d.month = 1, 1
  case Decadic:
    if 10 * d.year / 10 > emptyYear {
      d.day, d.month, d.year = 1, 1, 10 * d.year / 10
    }
  }
  return Weekday(weekday (d))
}


func computeHolidays () { // Quelle: S. Deschauer, Die Osterfestberechnung. DdM 14 (1986), 68-84
//
  var D Imp
  D.day, D.month, D.year = 1, 1, actualHolidayYear
  Wochentag:= weekday (D)
  for m:= uint(1); m <= maxmonth; m++ {
    D.month = uint(m)
    for t:= 1; uint(t) <= daysInMonth (D); t++ {
      holiday [t][m] = Wochentag == Sunday
      if Wochentag == Sunday {
        Wochentag = Monday
      } else {
        Wochentag++
      }
    }
  }
  holiday [1][1] = true // Neujahr
  if actualHolidayYear >= 1890 { // Tag der Arbeit
    holiday [1][5] = true
  }
  if actualHolidayYear > 1953 { // Tag der deutschen Einheit
    if actualHolidayYear < 1990 { // 17.6.1990 ein Sonntag
      holiday [17][6] = true
    } else {
      holiday [3][10] = true
    }
  }
  holiday [25][12] = true // Weihnachten
  holiday [26][12] = true
/* >>> 1583..6199:
  s = J / 100 - J / 400 - 2
   für 1900..2099: 13, 1800..1899: 12
  m = (J - 100 * (J / 4200)) / 300 - 2
    für 1800..2000: 4
  M = (15 + s - m) % 30
    für 1900, 2000: 24, 1800..1899: 23
  N = (6 + s) % 7
    für 1900..2099: 5, 1800..1899: 4
>>> für 1800..2099 reicht also: */
  mm:= uint(24)
  nn:= uint(5)
  if actualHolidayYear < 1900 {
    mm--
    nn--
  }
  d:= (mm + 19 * (actualHolidayYear % 19)) % 30
  e:= (nn + 2 * (actualHolidayYear % 4) + 4 * (actualHolidayYear % 7) + 6 * d) % 7
  t:= 22 + d + e
  if e == 6 { // Sonntag
    if d == 29 || actualHolidayYear % 19 >= 11 && d == 28 {
      t -= 7
    }
  }
  // 22 <= t <= 56, t. März ist Ostersonntag
  if t <= 30 { // Ostermontag
    holiday [t + 1][3] = true
  } else {
    holiday [t - 30][4] = true
  }
/*
  var f, ft uint
  carnival.year = actualHolidayYear
  // carnival = 7 Wochen vor Osterdienstag = 48 Tage vor Ostermontag
  // 2. Februar <= carnival <= 8. Mürz
  ft = t
  carnival.month = 3, // März
  if ft <= 48 {
    ft += 28
    if isLeapYear (actualHolidayYear) {
      ft++
    }
    carnival.month-- // Februar
  }
  ft -= 48
  carnival.day = ft
*/
  if t <= 33 { // Karfreitag
    holiday [t - 2][3] = true
  } else {
    holiday [t - 33][4] = true
  }
  t -= 11 // Ostersonntag + 50 Tage - April - Mai: 11 <= t <= 46
  if t <= maxday { // Pfingstmontag
    holiday [t][5] = true
  } else {
    holiday [t - maxday][6] = true
  }
  t -= 11 // Pfingstmontag - 11 Tage: 0 <= t <= 35
  if t == 0 { // Himmelfahrt
    holiday [30][4] = true
  } else if t <= maxday {
    holiday [t][5] = true
  } else {
    holiday [t - maxday][6] = true
  }
  if actualHolidayYear <= 1994 { // Bußtag
    D.day = 20 // wenn das ein Mo ist, ist der 22. Bußtag
    D.month = 11
    D.year = actualHolidayYear
    D.day -= uint(weekday (D))
    D.day += 2 // Spanne von Montag bis Mittwoch, s.o.
    holiday [D.day][11] = true
  }
}


func (D *Imp) IsHoliday () bool {
//
  if D.year == emptyYear { return false }
  if D.year != actualHolidayYear {
    actualHolidayYear = D.year
    computeHolidays()
  }
  return holiday [D.day][D.month]
}


func (x *Imp) SetEaster () {
//
  if x.year == emptyYear { return }
  if x.year != actualHolidayYear {
    actualHolidayYear = x.year
    computeHolidays()
  }
  x.day = 24
  x.month = 3 // earliest possible Eastermonday
  w:= x.Weekday (Daily)
  if w != Monday {
    x.day += 7 - uint(w)
  } // the first monday after
  for ! holiday [x.day][x.month] {
    x.day += 7
    if x.day > maxday {
      x.day -= maxday
      x.month = 4
    }
  } // Eastermonday
  if x.day > 1 {
    x.day --
  } else {
    x.day = maxday
    x.month = 3
  }
}


func (x *Imp) LastSunday (a bool) Calendarday {
//
  y:= New()
  if x.year == emptyYear {
    return y
  }
  y.year = x.year
  if a { // October
    y.month = 10
  } else {
    y.month = 3
  }
  y.day = maxday
  wd:= weekday (*y)
  if wd != Sunday {
    y.day -= 1 + uint(wd)
  }
  return y
}


func (x *Imp) Normal () bool {
//
  if x.year == emptyYear { return false }
  oct, mar:= x.LastSunday (true), x.LastSunday (false)
  return x.Less (mar) || oct.Eq (x) || oct.Less (x)
}


func (x *Imp) Normal1 () bool {
//
  if x.year == emptyYear { return false }
  oct, mar:= x.LastSunday (true), x.LastSunday (false)
  if x.Normal () {
    return x.Eq (oct)
  }
  return x.Eq (mar)
}


func (x *Imp) IsWorkday () bool {
//
  switch x.Weekday (Daily) { case Saturday, Sunday:
    return false
  }
  return ! x.IsHoliday ()
}


func (x *Imp) NWorkdays (Y Calendarday) uint {
//
  y:= x.imp (Y)
  if x.Empty() || y.Empty() { return 0 } // MaxNat ?
  a:= uint(0)
  if x.Less (y) {
    z:= x.Clone ().(*Imp)
    for {
      if z.IsWorkday () {
        a++
      }
      z.Inc (Daily)
      if y.Less (z) { break }
    }
  } else {
    z:= y.Clone ().(*Imp)
    for {
      if z.IsWorkday () {
        a++
      }
      z.Inc (Daily)
      if x.Less (z) { break }
    }
  }
  return a
}


func (x *Imp) SetFormat (f Format) {
//
  if f < NFormats {
    x.fmt = f
  }
}


func (x *Imp) SetColours (f, b col.Colour) {
//
  x.cF, x.cB = f, b
}


func (x *Imp) Weeknumber () uint {
//
  if x.year == emptyYear { return 0 }
  const Stichtag = Thursday // DIN 8601 (1975), entspricht ISO-Entwurf
  var y Imp
  y.day, y.month, y.year = 1, 1, x.year
  n:= uint(internalCode (*x) - internalCode (y))
  wd:= weekday (y)
  n += uint(wd)
  if wd <= Stichtag { n += 7 }
  return n / 7
}


func (x *Imp) String () string {
//
  if x.year == emptyYear {
    return str.Clr (length [x.fmt])
  }
  const mitNullen = true
  s:= ""
  switch x.fmt { case Dd, Dd_mm_, Dd_mm_yy, Dd_mm_yyyy, Dd_M, Dd_M_yyyy:
    if x.day == 0 { Panic ("day.String: x.day == 0") }
    s = nat.StringFmt (x.day, 2, mitNullen)
    if x.fmt == Dd { return s }
    s += "."
    switch x.fmt { case Dd_M, Dd_M_yyyy:
      s += " " + nameMonth [x.month]
      str.RemSpaces (&s)
      if x.fmt == Dd_M { return s }
      s += " "
    default:
      s += nat.StringFmt (x.month, 2, mitNullen) + "."
    }
    switch x.fmt { case Dd_mm_:
      ;
    case Dd_mm_yy:
      s += nat.StringFmt (x.year, 2, true)
    default: // Dd_mm_yyyy, Dd_M_yyyy:
      s += nat.StringFmt (x.year, 4, false)
    }
  case Yymmdd:
    s = nat.StringFmt (x.year % 100, 2, true) +
        nat.StringFmt (x.month, 2, true) +
        nat.StringFmt (x.day, 2, true)
  case Yyyymmdd:
    s = nat.StringFmt (x.year, 4, true) +
        nat.StringFmt (x.month, 2, true) +
        nat.StringFmt (x.day, 2, true)
  case Yy:
    s = nat.StringFmt (x.year, 2, true)
  case Yyyy:
    s = nat.StringFmt (x.year, 4, false)
  case Wd:
    s = WdShorttext [x.Weekday (Daily)]
  case WD:
    s = WdText [x.Weekday (Daily)]
  case Mmm, M:
    s = nameMonth [x.month]
    if x.fmt == Mmm { s = str.Part (s, 0, 3) }
  case Myyyy:
    s = nameMonth [x.month] + " " +
        nat.StringFmt (x.year, 4, false)
  case Wn, WN, WNyyyy:
    s = nat.StringFmt (x.Weeknumber(), 2, false)
    if x.fmt > Wn { s += ".Woche" }
    if x.fmt == WNyyyy {
      s += " " + nat.StringFmt (x.year, 4, false)
    }
  case Qu:
    switch (x.month - 1) / 3 { case 0:
      s = "  I"
    case 1:
      s = " II"
    case 2:
      s = "III"
    case 3:
      s = " IV"
    }
    s += "/" + nat.StringFmt (x.year, 2, true)
  }
  return s
}


func (x *Imp) Day () uint {
//
  if x.year == emptyYear { return 0 }
  return x.day
}


func (x *Imp) Month () uint {
//
  if x.year == emptyYear { return 0 }
  return x.month
}


func (x *Imp) Year () uint {
//
  if x.year == emptyYear { return 0 }
  return x.year
}


func (x *Imp) Write (l, c uint) {
//
//  switch x.fmt { case W, M, Dd_M_yyyy:
//    var e Imp; e.day, e.month, e.year = maxday, maxmonth, emptyYear
//    e.fmt = x.fmt
//    e.cF, e.cB = x.cF, x.cB
//    e.Write (l, c)
//  }
//   default:
  bx.Wd (length [x.fmt])
  bx.Colours (x.cF, x.cB)
  bx.Write (x.String(), l, c)
}


func (x *Imp) PosInWeek (vertical bool, a uint) (uint, uint) {
//
  if x.year == emptyYear { return 0, 0 }
  l:= uint(0)
  S:= a * uint(weekday (*x))
  if vertical { l = S; S = 0 }
  return l, S
}


func (x *Imp) WriteWeek (vertical bool, a, l, c uint) {
//
  if x.year == emptyYear { return }
// oldF, oldB:= x.cF, x.cB
  if vertical {
    if a == 0 { a = 1 }
  } else {
    if a == 0 { a = length [x.fmt] + 1 }
  }
  y:= x.Clone ().(*Imp)
  y.SetBeginning (Weekly)
  l1, c1:= uint(0), uint(0)
  for i:= 0; i <= 6; i++ {
    if vertical {
      l1 = a * uint(i)
    } else {
      c1 = a * uint(i)
    }
    op (y)
    y.Write (l + l1, c + c1)
    y.Inc (Daily)
  }
// x.SetColours (oldF, oldB)
}

/*
func (x *Imp) changeWithMouseInWeek (vertical bool, a, l, c uint) {
//
  x.changeWithMouse (Weekly, vertical, a, 0, 0, l, c)
}
*/

func (x *Imp) PosInMonth (vertical bool, n, z, s uint) (uint, uint) {
//
  if x.year == emptyYear { return 0, 0 }
  if n == 0 { n = 1 }
  n = 7 * n
  if vertical {
    if z == 0 { z = 1 }
  } else {
    if s == 0 { s = length [x.fmt] + 1 }
  }
  d:= *x
  d.day = 1
  i:= uint (weekday (d)) + x.day - 1
  if vertical {
    return z * (i % n), s * (i / n)
  }
  return z * (i / n), s * (i % n)
}


func (x *Imp) WriteMonth (vertical bool, n, l0, c0, l, c uint) {
//
  if x.year == emptyYear { return }
  t:= daysInMonth (*x)
  if n == 0 { n = 1 }
  n = 7 * n
  if vertical {
    if l0 == 0 { l0 = 1 }
  } else {
    if c0 == 0 { c0 = length [x.fmt] + 1 }
  }
  max:= int((maxday - 2) / n + 2)
  max *= int(n)
  y:= x.Clone ().(*Imp)
  y.day = 1
  w:= int(weekday (*y))
  var e Imp
  e.day, e.month, e.year = maxday, maxmonth, emptyYear
  e.fmt = x.fmt
  e.SetColours (col.Blue, WeekdayB)
  var l1, c1 uint
  for i:= 0; i < max; i++ {
    if vertical {
      l1, c1 = l0 * (uint(i) % n), c0 * (uint(i) / n)
    } else {
      l1, c1 = l0 * (uint(i) / n), c0 * (uint(i) % n)
    }
    if i < w || i >= w + int(t) {
      e.Write (l + l1, c + c1)
    } else {
      op (y)
      y.Write (l + l1, c + c1)
      if y.day < t {
        y.day ++
      }
    }
  }
}

/*
func (x *Imp) changeWithMouseInMonth (vertical bool, n, l, c, l1, c1 uint) {
//
  x.changeWithMouse (Monthly, vertical, n, l, c, l1, c1)
}
*/


func (x *Imp) PrintMonth (vertical bool, n, z, s, l, S uint) {
//
  if x.year == emptyYear { return }
  if n == 0 { n = 1 }
  n = 7 * n
  if vertical {
    if z == 0 { z = 1 }
  } else {
    if s == 0 { s = length [x.fmt] + 1 }
  }
  max:= (maxday - 2) / n + 2
  max = n * max
  y:= x.Clone().(*Imp)
  y.day = 1
  W:= uint(weekday (*y))
  t:= daysInMonth (*x)
  for i:= uint(0); i < max; i++ {
    var l1, S1 uint
    if vertical {
      l1 = z * (i % n)
      S1 = s * (i / n)
    } else {
      l1 = z * (i / n)
      S1 = s * (i % n)
    }
    if i < W || i >= W + t {
      // pbx.Clr (l + l1, S + S1) // TODO
    } else {
      op (y)
      if y.IsHoliday() {
        y.SetFont (font.Bold)
      } else {
        y.SetFont (font.Roman)
      }
      y.Print (l + l1, S + S1)
      if y.day < t {
        y.day++
      }
    }
  }
}


const (
  monthsHorizontally = 4
  leftMargin = uint(5)) // mindestens 3, höchstens 5


func (x *Imp) shift (l, c *uint) {
//
  *l += (7 + 1) * ((x.month - 1) / monthsHorizontally)
  *c += (7 - 1) * (2 + 1) * ((x.month - 1) % monthsHorizontally)
}


func (x *Imp) PosInYear() (uint, uint) {
//
  l, c:= x.PosInMonth (true, 1, 1, 3)
  x.shift (&l, &c)
  l ++
  c += leftMargin
  return l, c
}


func (x *Imp) changeWithMouse (p Period, vertical bool, a, l, c, l0, c0 uint) {
//
  if ! scr.MouseEx() { return }
  switch p { case Daily, Decadic: return; default: }
  lm, cm:= scr.MousePos ()
//  if p == Yearly { SM += leftMargin }
  y:= x.Clone ().(*Imp)
  y.SetBeginning (p)
  A:= y.Clone ().(*Imp)
  y.SetFormat (Dd_mm_yy)
  n:= length [x.fmt]
  var lpos, cpos uint
  for {
    n = length [x.fmt]
    if ! y.Equiv (A, p) {
      break
    }
    switch p { case Weekly:
      lpos, cpos = y.PosInWeek (vertical, a)
      lpos += l0; cpos += c0
    case Monthly:
      lpos, cpos = y.PosInMonth (vertical, a, l, c)
      lpos += l0; cpos += c0
    case Quarterly:
      errh.Error ("in " + pack + " not yet implemented", 3)
      return
    case HalfYearly:
      errh.Error ("in " + pack + " not yet implemented", 6)
      return
    case Yearly:
      n = 2
      lpos, cpos = y.PosInYear()
    default:
    }
    if lm == lpos && cpos <= cm && cm < cpos + n {
      y.Copy (x)
      break
    } else {
      y.Inc (Daily)
    }
  }
}


func (x *Imp) writeYearmask (l, c uint) {
//
  const X = 80
// scr.Clr (l, S, X, 25 - 1)
//  y.day = 1
//  y.year = x.year
  y:= x.Clone ().(*Imp)
  y.fmt = Yyyy
  y.SetColours (YearnumberF, YearnumberB)
  y.Write (l, c)
  y.Write (l, c + 80 - 4)
  bx.Colours (MonthF, MonthB)
  bx.Wd (X - 4 /* - c */)
  T1:= str.Clr (X - 4 /* - c */)
  bx.Write (T1, l, c + 4)
  bx.Wd (X /* - c */)
  T1 = str.Clr (X /* - c */)
  bx.Write (T1, l + 8, c)
  bx.Write (T1, l + 16, c)
  y.fmt = M
  var l1, c1 uint
  for m:= uint(1); m <= maxmonth; m++ {
    y.month = m
    l1, c1 = 0, leftMargin
    y.shift (&l1, &c1)
    y.SetColours (MonthF, MonthB)
    y.Write (l + l1, c + c1 + 3)
  }
  bx.Colours (WeekdayNameF, WeekdayNameB)
  y.fmt = Wd
  c2:= leftMargin + monthsHorizontally * 6 * 3
                // 6 Spalten pro Monat     tt-Format + 1 Zwischenraum
  bx.Wd (2) // len (WdShorttext)
  for m:= uint(1); m <= maxmonth; m += monthsHorizontally {
    for w:= Monday; w <= Sunday; w++ {
      for i:= uint(0); i <= 2; i++ {
        l1 = 1 + 8 * i + uint(w)
        bx.Write (WdShorttext [w], l + l1, c + 1)
        bx.Write (WdShorttext [w], l + l1, c + c2)
      }
    }
  }
}


func (x *Imp) WriteYear (l, c uint) {
//
  if x.year == emptyYear { return }
  var l1, c1 uint
  x.writeYearmask (l, c)
  y:= x.Clone ().(*Imp)
  y.fmt = Dd
  for m:= uint(1); m <= maxmonth; m++ {
    y.day, y.month = 1, uint(m)
    l1, c1 = 1, leftMargin
    y.shift (&l1, &c1)
    y.WriteMonth (true, 1, 1, 3, l + l1, c + c1)
  }
}


func (x *Imp) EditInYear (l, c uint) {
//
  for {
    c, _:= kbd.Command ()
    if c == kbd.Here {
      x.changeWithMouse (Yearly, false, 0, 0, 0, 0, 0)
      break
    }
  }
}

func (x *Imp) printYearMask (l, c uint) {
//
  const X = 80
  y:= x.Clone ().(*Imp)
  y.fmt = Yyyy
  y.font = font.Bold
  y.Print (l, c)
  y.Print (l, X - 4)
  pbx.Print (str.Clr (X - 4 - c), l, c + 4)
  pbx.Print (str.Clr (X - c), l + 8, c)
  pbx.Print (str.Clr (X - c), l + 16, c)
  y.fmt = M
  y.font = font.Italic
  var l1, c1 uint
  for m:= uint(1); m <= maxmonth; m++ {
    y.month = m
    l1, c1 = 0, leftMargin
    y.shift (&l1, &c1)
    y.Print (l + l1, c + c1 + 3)
  }
  y.fmt = Wd
  c2:= leftMargin + monthsHorizontally * 6 * 3
                // 6 Spalten pro Monat     tt-Format + 1 Zwischenraum
  pbx.SetFont (font.Italic)
  for m:= uint(1); m <= maxmonth; m += monthsHorizontally {
    for w:= Monday; w <= Sunday; w++ {
      for i:= uint(0); i <= 2; i++ {
        l1 = 1 + 8 * i + uint(w)
        pbx.Print (WdShorttext [w], l + l1, c + 1)
        pbx.Print (WdShorttext [w], l + l1, c + c2)
      }
    }
  }
}


func (x *Imp) PrintYear (l, c uint) {
//
  if x.year == emptyYear { return }
  x.printYearMask (l, c)
  y:= New ()
  y.fmt = Dd
  for m:= uint(1); m <= maxmonth; m++ {
    y.day, y.month = 1, m
    dl, dc:= uint(1), leftMargin
    y.shift (&dl, &dc)
    y.PrintMonth (true, 1, 1, 3, l + dl, c + dc)
  }
}


func isYear (y *uint) bool {
//
  if *y < uint(100) {
    *y += currentCentury
    if *y > limitYear {
      *y -= 100
    }
  }
  return startYear <= *y && *y <= endYear
}


func isMonth (Monat *uint, Wort string) bool {
//
  var T string
  n:= str.ProperLen (Wort)
  if n > 0 {
    for m:= uint(1); m <= maxmonth; m++ {
      T = str.Part (nameMonth [m], 0, n)
      if Wort == T { // str.QuasiEq (Wort, T) {
        *Monat = uint(m)
        return true
      }
    }
  }
  return false
}


func (x *Imp) defined (d, m, y uint) bool {
//
  if d == 0 || d > maxday { return false }
  if m == 0 || m > maxmonth { return false }
  if ! isYear (&y) { return false }
  x.day, x.month, x.year = d, m, y
  return x.day <= daysInMonth (*x)
}


func (x *Imp) Defined (s string) bool {
//
  if str.Empty (s) { x.Clr(); return true }
  var d Imp
  d.day, d.month, d.year = x.day, x.month, x.year
  var T string
  var l, p uint
  n, ss, P, L:= nat.DigitSequences (s)
  switch x.fmt {
  case Dd, // e.g. " 8"
       Dd_mm_, // e.g. " 8.10."
       Dd_mm_yy, // e.g. " 8.10.07"
       Dd_mm_yyyy: // e.g. " 8.10.2007" *):
    switch n {
    case 1:
      l = 2
    case 2, 3:
      l = L[0]
    default:
      return false
    } // see below
  case Dd_M, // e.g. "8. Oktober"
       Dd_M_yyyy: // e.g. "8. Oktober 2007"
    if x.fmt == Dd_M {
      if n != 1 { return false }
    } else {
      if n != 2 { return false }
    }
    if ! str.Contains (s, '.', &p) { return false }
    if x.fmt == Dd_M_yyyy {
//      l = str.ProperLen (s)
//      T = str.Part (s, p, l - p)
      T = ss[1]
      if ! nat.Defined (&d.year, T) {
        return false
      }
    }
    T = ss[0]
    str.Move (&T, true)
    if ! nat.Defined (&d.day, T) {
      return false
    }
    T = str.Part (s, p + 1, P[1] - p - 1)
    str.Move (&T, true)
    if ! isMonth (&d.month, T) { return false }
    return x.defined (d.day, d.month, d.year)
  case Yymmdd: // e.g. "090418"
    if ! nat.Defined (&d.year, str.Part (s, 0, 2)) { return false }
    if ! nat.Defined (&d.month, str.Part (s, 2, 2)) { return false }
    if ! nat.Defined (&d.day, str.Part (s, 4, 2)) { return false }
    return x.defined (d.day, d.month, d.year)
  case Yyyymmdd: // e.g. "20090418"
    if ! nat.Defined (&d.year, str.Part (s, 0, 4)) { return false }
    if ! nat.Defined (&d.month, str.Part (s, 4, 2)) { return false }
    if ! nat.Defined (&d.day, str.Part (s, 6, 2)) { return false }
    return x.defined (d.day, d.month, d.year)
  case Yy, // e.g. "08"
       Yyyy: // e.g. "2007"
    if n != 1 { return false }
    if nat.Defined (&d.year, ss[0]) {
      return x.defined (d.day, d.month, d.year)
    } else {
      return false
    }
  case Wd, // e.g. "Mo"
       WD: // e.g. "Monday"
    return false // Fall noch nicht erledigt
  case Mmm, // e.g. "Mon"
       M: // e.g. "Oktober"
    if ! isMonth (&d.month, s) { return false }
    return x.defined (d.day, d.month, d.year)
  case Myyyy: // e.g. "Oktober 2007"
    if n != 1 { return false }
    if ! nat.Defined (&d.year, ss[0]) {
      return false
    }
    if ! str.Contains (s, ' ', &p) { return false }
    if ! isMonth (&d.month, str.Part (s, 0, p)) { return false }
    return x.defined (d.day, d.month, d.year)
  case Wn, // e.g. "1" (.Woche)
       WN: // e.g. "1.Woche"
    if n != 1 { return false }
    if nat.Defined (&n, T) {
      if 0 < n && n <= 3 {
        d.day, d.month, d.year = 1, 1, x.year
        c:= internalCode (d)
        w:= weekday (d)
        if w > Thursday { c += 7 } // see Weeknumber
        if c < uint16(w) { return false }
        c -= uint16(w) // so c is a Monday
        d.Decode (Encode (uint(c) + 7 * n))
        if d.year == x.year {
          x.day = d.day
          x.month = d.month
          return true
        }
      }
      return false
    }
  case WNyyyy: // e.g. "1.Woche 2007"
    return false // not yet implemented
  case Qu: // e.g. "  I/06"
    if n != 1 { return false }
    if ! str.Contains (s, '/', &p) { return false }
    if ! nat.Defined (&d.year, ss[0]) {
      return false
    }
    T = str.Part (s, 0, p)
    str.Move (&T, true)
    n = str.ProperLen (T)
    if T [0] != 'I' { return false }
    switch n { case 1:
      d.month = 1
    case 2:
      switch T [1] { case 'I':
        d.month = 4
      case 'V':
        d.month = 10
      default:
        return false
      }
    case 3:
      if T [1] == 'I' && T [2] == 'I' { d.month = 7 }
    default:
      return false
    }
    return x.defined (d.day, d.month, d.year)
  }
  if ! nat.Defined (&d.day, str.Part (s, P[0], l)) { return false }
  if n == 1 {
    if L [0] > 8 { return false } // maximal "Dd_mm_yyyy"
    if L [0] > 2 {
      if ! nat.Defined (&d.month, str.Part (s, P [0] + 2, 2)) { return false }
    }
    if L [0] > 4 {
      if ! nat.Defined (&d.year, str.Part (s, P [0] + 4, L [0] - 4)) { return false }
    }
  } else { // n == 2, 3
    if ! nat.Defined (&d.month, ss[1]) { return false }
    if n == 2 && x.Empty() {
      d.year = today.(*Imp).year
    }
    if n == 3 {
      if ! nat.Defined (&d.year, ss[2]) { return false }
    }
  }
  return x.defined (d.day, d.month, d.year)
}

/*
func Date (m, d, y int) *Imp {
//
  y:= New()
  if y.Defined3 (m, d, y) { }
  return y
}
*/

func (x *Imp) Selected (cop CondOp) bool {
//
  loop: for {
    cop (x, true) // colour auffallend
    c, t:= kbd.Command ()
    cop (x, false) // colour normal
    switch c { case kbd.Enter:
      return true
    case kbd.Esc:
      break loop
    default:
      if x.fmt == Yy || x.fmt == Yyyy {
        if t == 0 { t = 3 } else { t = 5 }
      }
      x.Change (c, t)
    }
  }
  return false
}


func (x *Imp) Edit (l, c uint) {
//
  bx.Wd (length [x.fmt])
  bx.Colours (x.cF, x.cB)
  s:= x.String ()
  nErr:= 0
  for {
    bx.Edit (&s, l, c)
    if x.Defined (s) {
      x.Write (l, c)
      return
/*
      var t uint
      if kbd.LastCommand (&t) == kbd.LookFor && t == 0 {
        if x.year == emptyYear { x.Copy (today) }
        errh.Hint ("Datum ändern: Kursortasten, Datum auswählen: Enter, Eingabe stornieren: Esc")
        loop: for {
          cm, _:= kbd.Command ()
          x.Change (cm, t)
          x.Write (l, c)
          switch cm { case kbd.Enter, kbd.Here:
            errh.DelHint(); return
          case kbd.Back, kbd.There:
            errh.DelHint(); break loop
          }
        }
      } else {
        x.Write (l, c)
        return
      }
*/
    } else {
      nErr ++
      switch nErr { case 1:
        errh.ErrorPos ("Die Eingabe stellt kein Datum dar!", 0, l + 1, c)
      case 2:
        errh.ErrorPos ("Das ist auch kein Datum!", 0, l + 1, c)
      case 3:
        errh.ErrorPos ("Jetzt passen Sie doch mal auf!", 0, l + 1, c)
      case 4:
        errh.ErrorPos ("Können Sie kein Datum eingeben?", 0, l + 1, c)
      default:
        errh.ErrorPos ("Was soll der Quatsch?", 0, l + 1, c)
        x.Actualize ()
        s = x.String ()
      }
    }
  }
}


func (x *Imp) Codelen () uint {
//
  return Codelen(uint16(0))
}


func (x *Imp) Encode () []byte {
//
  B:= make ([]byte, Codelen(uint16(0)))
  copy (B, Encode (uint16(internalCode(*x))))
  return B
}


func (x *Imp) decode (n uint16) {
//
  var d uint16
  if n == 0 {
    x.year = emptyYear
  } else {
    x.year = startYear
    for {
      d = daysInYear (x.year)
      if n > d {
        x.year++
        n -= d
      } else {
        break
      }
    }
    x.month = 1
    for {
      d = uint16(daysInMonth (*x))
      if n > d {
        x.month++
        n -= d
      } else {
        break
      }
    }
    x.day = uint(n)
  }
}


func (x *Imp) Decode (B []byte) {
//
  c:= Decode (uint16(0), B).(uint16)
  if c <= uint16(maxCode) {
    x.decode (c)
  } else {
    x.Clr()
  }
}


func (x *Imp) SetFont (f font.Font) {
//
  x.font = f
}


func (x *Imp) Print (l, c uint) {
//
/*
  if x.IsHoliday () && x.fmt <= Dd_M_yyyy {
    x.font = font.Bold
  } else {
    x.font = font.Roman
  }
  if x.fmt == Yyyy || x.fmt == M {
    x.font = font.Italic
  }
*/
  pbx.SetFont (x.font)
  pbx.Print (x.String(), l, c)
}


func (x *Imp) SetAttribute (p Op) {
//
  op = p
}


func attribute (a Any) {
//
  x, ok:= a.(*Imp)
  if ! ok { TypePanic () }
  switch x.fmt { case Dd, Dd_mm_, Dd_mm_yy, Yymmdd, Yyyymmdd, Dd_mm_yyyy, Dd_M, Dd_M_yyyy:
    if x.IsHoliday () {
      x.SetColours (HolidayF, HolidayB)
      x.SetFont (font.Bold)
    } else {
      x.SetColours (WeekdayF, WeekdayB)
      x.SetFont (font.Roman)
    }
  case Yy, Yyyy:
    x.SetColours (YearnumberF, YearnumberB)
    x.SetFont (font.Bold)
  case Wd, WD:
    x.SetColours (WeekdayNameF, WeekdayNameB)
    x.SetFont (font.Italic)
//  case Mmm:
  case M, Myyyy:
    x.SetColours (MonthF, MonthB)
    x.SetFont (font.Italic)
  default:
    x.SetColours (WeekdayF, MonthB)
    x.SetFont (font.Slanted)
  }
}


func init () {
//
  WeekdayF, WeekdayB = col.ScreenF, col.ScreenB
  HolidayF, HolidayB = col.Red, WeekdayB
  YearnumberF, YearnumberB = col.LightWhite, col.Magenta
  WeekdayNameF, WeekdayNameB = col.Magenta, WeekdayB
  MonthF, MonthB = YearnumberF, YearnumberB
  today.Actualize ()
  currentCentury = 100 * (today.(*Imp).year / 100)
  todayCode = internalCode (*today.(*Imp))
//  carnival = New()
}
