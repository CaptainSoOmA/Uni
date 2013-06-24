package ker

// (c) Christian Maurer   v. 121118 - license see murus.go

func Binomial (n, k uint) uint {
//
  if n < k { return 0 }
  if n - k < k { k = n - k }
  b:= uint(1)
  for i:= uint(1); i <= k; i++ {
    b *= n
    b /= i
    n --
  }
  return b
}
