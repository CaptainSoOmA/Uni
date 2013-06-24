package str

// (c) Christian Maurer   v. 130303 - license see murus.go

import
  "murus/z"


func isUTF8 (S *string) bool {
//
  return z.DevilsDung (S)
}


func clr (n uint) string {
//
  if n == 0 {
    return ""
  }
  return const_ (' ', n)
}


func set (S *string, T string) {
//
  *S = T
  z.ToHellWithUTF8 (S)
}


func lat1 (s *string) {
//
  z.ToHellWithUTF8 (s)
}


func utf8 (s *string) {
//
  for i:= len (*s) - 1; i >= 0; i-- {
    c:= (*s)[i]
    if z.IsLatin1 (c) {
      *s = (*s)[:i] + string(c) + (*s)[i+1:]
    }
  }
}


func lit (s string) bool {
//
  i:= len (s)
  if i == 0 { return false } // (?) TODO
  c:= s[0]
  if ! (c < 'A' || 'Z' < c) && (c < 'a' || c > 'z') {
    return false
  }
  return true
}


func replace (s *string, p uint, c byte) {
//
  n:= len (*s)
  if int(p) >= n { return }
  t:= string(c)
  *s = (*s)[:p] + t + (*s)[p+1:]
  z.ToHellWithUTF8 (s)
}


func empty (s string) bool {
//
  for i:= 0; i < len (s); i++ {
    if s[i] != ' ' {
      return false
    }
  }
  return true
}


func const_ (c byte, n uint) string {
//
  s:= ""
  for i:= uint(0); i < n; i++ {
    t:= string(c)
    z.ToHellWithUTF8 (&t)
    s += t
  }
  return s
}


func properLen (s string) uint {
//
  n:= len (s)
  for {
    if n == 0 {
      break
    }
    if s [n-1] == ' ' {
      n --
    } else {
      break
    }
  }
  return uint(n)
}


func quasiEq (S, T string) bool {
//
  n, n1:= properLen (S), properLen (T)
  if n != n1 { return false }
  if n == 0 { return true }
  return S[0:n] == T[0:n]
}


func toUpper (s *string) {
//
  n:= len (*s)
  b:= make ([]byte, n)
  for i:= 0; i < n; i++ {
    b[i] = z.Cap ((*s)[i])
  }
  *s = string(b)
}


func toLower (s *string) {
//
  n:= len (*s)
  b:= make ([]byte, n)
  for i:= 0; i < n; i++ {
    b[i] = z.Lower ((*s)[i])
  }
  *s = string(b)
}


func toUpper0 (s *string) {
//
  if len (*s) == 0 { return }
  *s = string(z.Cap ((*s)[0])) + (*s)[1:]
}


func toLower0 (s *string) {
//
  if len (*s) == 0 { return }
  *s = string(z.Lower ((*s)[0])) + (*s)[1:]
}


func cap0 (s string) bool {
//
  if s == "" { return false }
  return s[0] == z.Cap (s[0])
}


func equiv (s, t string) bool {
//
  n:= len (s)
  if len (t) != n {
    return false
  }
  for i:= 0; i < n; i++ {
    if ! z.Equiv (s[i], t[i]) {
      return false
    }
  }
  return true
}


func quasiEquiv (S, T string) bool {
//
  toUpper (&S)
  toUpper (&T)
  return quasiEq (S, T)
}


func less (s, t string) bool {
//
  n, n1:= len (s), len (t)
  i:= 0
  for {
    if i == n {
      return n < n1
    }
    if i == n1 {
      return false
    }
    if z.Less (s[i], t[i]) {
      return true
    }
    if z.Less (t[i], s[i]) {
      return false
    }
    i++
  }
  return false
}


func quasiLess (s, t string) bool {
//
  s1, t1:= s, t
  toUpper (&s1)
  toUpper (&t1)
  return less (s1, t1)
}


func contains (s string, b byte, n *uint) bool {
//
  for *n = uint(0); *n < uint(len (s)); *n++ {
    if s[*n] == b {
      return true
    }
  }
  *n ++
  return false
}


func pos (s string, b byte) uint {
//
  n:= uint(len (s))
  for i:= uint(0); i < n; i++ {
    if s[i] == b {
      return i
    }
  }
  return n
}


func quasiContains (s string, b byte, n *uint) bool {
//
  for *n = uint(0); *n < uint(len (s)); *n++ {
//    if z.Cap (s[*n]) == z.Cap (b) {
    if z.Equiv (s[*n], b) {
      return true
    }
  }
  *n ++
  return false
}


func isPart (S, T string, p *uint) bool {
//
  *p = 0
  k:= len (S)
  if k == 0 {
    return true
  }
  n:= len (T)
  if k > n {
    *p = uint(n)
    return false
  }
  var i int
  for {
    i = 0
    for {
      if i == k {
        return true
      }
      if S[i] != T[int(*p) + i] {
        break
      } else {
        i ++
      }
    }
    if int(*p) + k < n {
      *p ++
    } else {
      *p = uint(n)
      break
    }
  }
  return false
}


func isEquivPart (s, t string, p *uint) bool {
//
  toUpper (&s)
  toUpper (&t)
  return isPart (s, t, p)
}


func isEquivPart0 (s, t string) bool {
//
  n:= len (s)
  for {
    if n == 0 {
      break
    }
    if s[n-1] == ' ' {
      n --
    } else {
      break
    }
  }
  if n == 0 { return true }
//  z.ToHellWithUTF8 (&s) // sicher ist sicher
//  z.ToHellWithUTF8 (&t)
  s = string (s[:n])
  toUpper (&s)
  t = string (t[:n])
  toUpper (&t)
  return s == t
}


func ins1 (s *string, c byte, p uint) {
//
  t:= string(c)
  z.ToHellWithUTF8 (&t)
  ins (s, t, p)
}


func ins (s *string, t string, p uint) {
//
  if len (t) == 0 || p > uint(len (*s)) { return }
  *s = (*s)[:p] + t + (*s)[p:]
}


func rem (s *string, p, n uint) {
//
  if n == 0 { return }
  l:= uint(len (*s))
  if p >= l { return }
  if p + n >= l {
    n = l - p
  }
  *s = (*s)[:p] + (*s)[p+n:]
}


func part (s string, p, n uint) string {
//
  if n == 0 {
    return ""
  }
  l:= uint(len (s))
  if p >= l {
    return s
  }
  if p + n > l {
    n = l - p
  }
  return s[p:p+n]
}


func norm (s *string, n uint) {
//
  if n == 0 {
    *s = ""
    return
  }
  k:= uint(len (*s))
  if k > n {
    *s = (*s)[:n]
    return
  }
  for i:= k; i < n; i++ { // k <= n
    *s += " "
  }
}


func remSpaces (s *string) {
//
  n:= len (*s)
  for {
    if n == 0 { break }
    if (*s)[n - 1] == ' ' {
      n --
    } else {
      break
    }
  }
  *s = (*s)[:n]
}


func remAllSpaces (s *string) {
//
  n:= len (*s)
  if n == 0 { return }
  b:= make ([]byte, n)
  i, j:= 0, 0
  loop: for j < n {
    if j == n { break }
    for (*s)[j] == ' ' {
      j++
      if j == n {
        break loop
      }
    }
    b[i] = (*s)[j]
    i ++
    j ++
  }
  *s = string(b[0:i])
}


func move (s *string, left bool) {
//
  l:= uint(len (*s))
  if l == 0 {
    return
  }
  if left {
    n:= uint(0)
    for n = 0; n < l; n++ {
      if (*s)[n] != ' ' {
        break
      }
    }
    if n == 0 || n == l {
      return
    }
    *s = (*s)[n:]
    for i:= uint(0); i < n; i++ {
      *s = *s + " "
    }
  } else {
    n:= l
    for n = l; n >= 1; n-- {
      if (*s)[n - 1] != ' ' {
        break
      }
    }
    *s = (*s)[:n]
    for i:= n; i < l; i++ {
      *s = " " + *s
    }
  }
}


func insSpace (s *string, p uint) {
//
  l:= uint(len (*s))
  if l == 0 || p > l { return }
  *s = (*s)[:p] + " " + (*s)[p:]
}


func shift (s *string, p uint) {
//
  l:= uint(len (*s))
  if l <= 1 || p + 1 >= l { return }
  if (*s)[l-1] != ' ' { return }
  *s = (*s)[0:p] + " " + (*s)[p:l-1]
}


func center (s *string, n uint) {
//
  if n == 0 {
    return
  }
  move (s, false)
  l:= ProperLen (*s)
  if n < l {
    *s = (*s)[:n]
    return
  }
  if l == n {
    return
  }
  if n == l + 1 {
    *s += " "
    return
  }
  k:= (n - l) / 2 // + (n - l) % 2
  *s = clr (k) + *s + clr (n - (l + k))
}


func remAllNondigits (s *string) {
//
  l:= uint(len (*s))
  if l == 0 { return }
  b:= make ([]byte, l)
  i, j:= uint(0), uint(0)
  loop: for j < l {
    if j == l { break }
    for (*s)[j] < '0' || (*s)[j] > '9' {
      j ++
      if j == l {
        break loop
      }
    }
    b[i] = (*s)[j]
    i ++
    j ++
  }
  *s = string(b[:i]) + clr (l - i)
}


func words (s string) (uint, []string, []uint) {
//
  z.ToHellWithUTF8 (&s)
  var t []string
  var p []uint
  l:= properLen (s)
  spaceBefore:= true
  n:= uint(0)
  for i:= uint(0); i < l; i++ {
    if s[i] == ' ' {
      spaceBefore = true
    } else {
      if spaceBefore {
        t = append (t, string(s[i]))
        p = append (p, i)
        n ++
        spaceBefore = false
      } else {
        t[n - 1] += string(s[i])
      }
    }
  }
  return n, t, p
}


func appendLF (s *string) {
//
  *s += "\n"
}


func appendLine (s *string, t string) {
//
  *s += (t + "\n")
}


func splitLine (s *string) string {
//
  l:= uint(len (*s))
  if l == 0 { return "" }
  n:= uint(0)
  for n = 0; n < l; n++ {
    if (*s)[n] == '\n' {
      break
    }
  }
  t:= (*s)[:n]
  n ++
  *s = (*s)[n:]
  return t
}
