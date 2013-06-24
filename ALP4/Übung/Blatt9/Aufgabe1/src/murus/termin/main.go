package main

// (c) Christian Maurer   v. 130526 - license see murus.go

// TODO Markierung der Termine mit dem eingegebenen Wort funzt noch nicht.

import (
  . "murus/obj"; "murus/ker"; "murus/kbd"
//  "murus/font"
  "murus/prt"; "murus/col"; "murus/scr"; "murus/errh"
  "murus/day"
  "murus/termin/dayattr"; "murus/termin/cal"
)
const (
  l1 = 24; c1 = 35
)
var (
  actualDay *day.Imp
  period day.Period
  help []string
  l0, c0, dc uint
)


func setWeekdayColours (a Any) {
//
  a.(*day.Imp).SetColours (day.WeekdayNameF, day.WeekdayNameB)
}


func pos (D *day.Imp) (uint, uint) {
//
  switch period {
  case day.Yearly:
    return D.PosInYear ()
  case day.Monthly:
    return D.PosInMonth (true, 1, 3, dc)
  case day.Weekly:
    return D.PosInWeek (false, dc)
  }
  return 0, 0
}


func writeAll (D *day.Imp) {
//
  switch period { case day.Monthly, day.Weekly:
    if D.Empty () { return }
  default:
    return
  }
  D1:= day.New ()
  D1.Copy (D)
  D1.SetBeginning (period)
  cal.SetFormat (period)
  for D.Equiv (D1, period) {
    l, c:= pos (D1)
    cal.Seek (D1)
    cal.SetFormat (period) // weil Seek über Define <- Clone das Format mitkopiert
    cal.WriteDay (l0 + l, c0 + c)
    D1.Inc (day.Daily)
  }
}


func clearAll (D *day.Imp) {
//
  switch period { case day.Monthly, day.Weekly:
    if D.Empty () { return }
  default:
    return
  }
  D1:= day.New ()
  D1.Copy (D)
  D1.SetBeginning (period)
  cal.SetFormat (period)
  for D.Equiv (D1, period) {
    l, c:= pos (D1)
errh.Error ("Tag Nr.", D1.OrdDay())
    cal.ClearDay (D1, l0 + l, c0 + c)
    D1.Inc (day.Daily)
  }
}


func write () {
//
  switch period { case day.Yearly:
    actualDay.WriteYear (l0, c0)
  case day.Monthly:
    actualDay.SetFormat (day.Dd_mm_)
    actualDay.WriteMonth (true, 1, 3, dc, l0, c0)
  case day.Weekly:
    actualDay.WriteWeek (false, dc, l0, c0 + 3)
  }
}

func editiert () bool {
//
  scr.Cls()
  l0, c0 = 0, 0; dc = 0
  switch period { case day.Decadic:
    actualDay.SetFormat (day.Yyyy)
    actualDay.SetColours (day.YearnumberF, day.YearnumberB)
  case day.Yearly:
    ;
  case day.HalfYearly, day.Quarterly:
    ker.Stop ("termin.main", 1)
  case day.Monthly:
    l0, c0 = 3, 5; dc = 12 // 11 dayattr + 1
    actualDay.SetAttribute (setWeekdayColours)
    actualDay.SetFormat (day.Wd)
    actualDay.WriteWeek (true, 3, l0, 2)
    actualDay.WriteWeek (true, 3, l0, 2 + 6 * dc + 3)
  case day.Weekly:
    l0, c0 = 2, 2; dc = 11 // 7 x 11 == 77 < 80
    actualDay.SetAttribute (setWeekdayColours)
    actualDay.SetFormat (day.Wd)
    actualDay.WriteWeek (false, dc, l0, c0)
  case day.Daily:
    cal.Edit (actualDay, l0, c0)
  }
  switch period { case day.Weekly, day.Monthly, day.Quarterly, day.Yearly:
    dayattr.WriteActual (l1, c1)
  }
  var startDate = day.New ()
  var Original = day.New ()
  loop: for {
    switch period { case day.Yearly:
      actualDay.SetFormat (day.Dd)
    case day.Quarterly, day.HalfYearly:
      errh.Error ("nicht erreichbarer Punkt", 3)
    case day.Weekly, day.Monthly:
      actualDay.SetFormat (day.Dd_mm_)
    }
    if ! actualDay.Equiv (startDate, period) {
//      clearAll (startDate) // alte Zusatzinformationen löschen
      startDate.Copy (actualDay)
      switch period { case day.Yearly:
        dayattr.WriteActual (l1, c1)
        actualDay.SetFormat (day.Dd)
      case day.Monthly, day.Weekly:
        if period == day.Monthly {
          actualDay.SetColours (day.MonthF, day.MonthB)
        } else {
          actualDay.SetColours (day.WeekdayNameF, day.WeekdayNameB)
        }
        actualDay.SetFormat (day.Yyyy)
        actualDay.SetColours (day.YearnumberF, day.YearnumberB)
        actualDay.Write (0, 0)
        actualDay.Write (0, 80 - 4)
        if period == day.Monthly {
          actualDay.SetFormat (day.M)
        } else {
          actualDay.SetFormat (day.WN)
        }
        actualDay.SetColours (day.MonthF, day.MonthB)
        actualDay.Write (0, 30)
        actualDay.SetFormat (day.Dd_mm_)
      }
      actualDay.SetAttribute (dayattr.Attrib)
      write ()
      writeAll (actualDay)
    }
    l, c:= pos (actualDay)
    dayattr.Attrib (actualDay)
    Original.Copy (actualDay)
    switch period { case day.Daily:
      ;
    case day.Weekly:
      actualDay.Edit (l0 + l, c0 + c + 3)
    case day.Monthly, day.Quarterly, day.HalfYearly, day.Yearly:
      actualDay.Edit (l0 + l, c0 + c)
    case day.Decadic:
      actualDay.Edit (0, 0)
    }
    if actualDay.Empty () {
      actualDay.Copy (Original)
    }
    var d uint
    C:= kbd.LastCommand (&d)
//    actualDay.Write (l0 + l, c0 + c)
    switch C {
    case kbd.Enter:
      for {
        if period == day.Daily {
          return false
        } else {
          period --
          if period == day.HalfYearly { period -- }
          if period == day.Quarterly { period -- }
        }
        if d == 0 { return false } else { d -- }
      }
    case kbd.Esc, kbd.Back:
      for {
        if period == day.Decadic {
          break loop
        } else {
          period ++
          if period == day.Quarterly { period ++ }
          if period == day.HalfYearly { period ++ }
        }
        if d == 0 { return false } else { d -- }
      }
    case kbd.Tab:
      dayattr.Change (d == 0)
      write ()
      dayattr.WriteActual (l1, c1)
      C = kbd.Enter // see above
    case kbd.Help:
      errh.WriteHelp (help)
    case kbd.LookFor:
      dayattr.Normalize ()
      dayattr.WriteActual (l1, c1)
      cal.EditWord (l1, c1 + 1 + 8)
      dayattr.Clr ()
      cal.LookFor (func (a Any) { dayattr.Actualize (a.(*day.Imp), true) })
      write ()
      if period == day.Weekly {
        writeAll (actualDay)
      }
      C = kbd.Enter // so "actualDay" is not influenced
    case kbd.Mark, kbd.Demark:
      dayattr.Actualize (actualDay, C == kbd.Mark)
      dayattr.Attrib (actualDay)
      actualDay.Write (l0 + l, c0 + c)
      C = kbd.Down
    case kbd.PrintScr: // TODO Zusatzinformationen
      switch period { case day.Yearly:
        actualDay.PrintYear (0, 0)
        prt.GoPrint()
/*
      case day.Monthly:
        actualDay.PrintMonth (true, 1, 3, dc, l0, c0)
        prt.Printout()
      case day.Daily:
        cal.Print (0, 0)
        prt.Printout()
*/
      }
/*  case kbd.Here: // , kbd.There, kbd.This:
      switch period { case day.Yearly:
        // day.mitMausImJahrAendern (actualDay)
      case day.Monthly:
        l0 = 3, c0 = 5; dc = 12
//        day.mitMausImMonatAendern (actualDay, true, 1, 3, dc, l0, c0)
      case day.Weekly:
        l0 = 2, c0 = 2; dc = 11
//        day.mitMausInWocheAendern (actualDay, false, dc, l0, c0)
      } */
    default:
      if period == day.Decadic {
        if d == 0 { d = 3 } else { d = 5 }
      }
    }
    switch C { case kbd.Here, kbd.There, kbd.This:
      if period == day.Yearly {
        switch C { case kbd.Here:
          period = day.Monthly
        case kbd.There:
          period = day.Weekly
        case kbd.This:
          period = day.Daily
        }
        C = kbd.Enter
        return false
      }
    default:
      actualDay.Change (C, d)
    }
  }
  return true
}

func main () {
//
  cF, cB:= col.Black, col.LightWhite
  col.ScreenF, col.ScreenB = cF, cB
  scr.Switch (scr.TXT)
/*
  if scr.Switchable (scr.WXGA) {
    scr.Switch (scr.WXGA)
    scr.SwitchFontsize (font.Huge)
  } else if scr.Switchable (scr.WSVGA) {
    scr.Switch (scr.WSVGA)
    scr.SwitchFontsize (font.Big)
  } else {
    scr.Cls ()
  }
*/
  help = make ([]string, 1)
  help[0] = "Bedienungsanleitung siehe Handbuch S. 18 - 23"
  actualDay = day.New ()
  actualDay.Actualize ()
  period = day.Yearly
  day.WeekdayF, day.WeekdayB = cF, cB
  day.HolidayB = cB
  day.WeekdayNameF, day.WeekdayNameB = col.Magenta, cB
//  errh.DocAvail = true
  scr.MouseCursor (true)
  for ! editiert() {  }
  cal.Terminate()
  ker.Terminate()
}
