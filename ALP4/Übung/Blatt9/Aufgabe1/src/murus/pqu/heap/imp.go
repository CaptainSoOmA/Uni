package heap

// (c) Christian Maurer   v. 120909 - license see murus.go

import
  . "murus/obj"
type
  Imp struct {
        top Any
 left, right *Imp
             }


func New () *Imp {
//
  return nil
}


// Pre: n > 0.
// Returns the greatest power of 2 <= n.
func g (n uint) uint {
//
  if n == 1 {
    return n
  }
  return 2 * g (n / 2)
}


// Pre: n > 0.
// Returns true, iff the last node of x with n nodes is contained
// in the left subheap of x, and in this case the number of nodes
// in the left, otherwise in the right subheap of x.
func f (n uint) (bool, uint) {
//
  left:= true
  if n == 1 {
    return left, 0
  }
  a:= g (n)
  b:= n - a
  left = b < a / 2
  if left {
    b += a / 2
  }
  return left, b
}


func (x *Imp) Ins (a Any, n uint) Heap {
//
  if n == 1 {
    x = new (Imp)
    x.top = Clone (a)
    x.left, x.right = nil, nil
  } else {
    left, k:= f (n)
    if left {
      x.left = x.left.Ins (a, k).(*Imp)
    } else {
      x.right = x.right.Ins (a, k).(*Imp)
    }
  }
  return x
}


func (x *Imp) swap (l bool) {
//
  if l {
    if x.left != nil {
      if Less (x.left.top, x.top) {
        x.top, x.left.top = x.left.top, x.top
      }
    }
  } else if x.right != nil {
    if Less (x.right.top, x.top) {
      x.top, x.right.top = x.right.top, x.top
    }
  }
}


func (x *Imp) Lift (n uint) {
//
  if n > 0 {
    left, k:= f (n)
    if left {
      x.left.Lift (k)
    } else {
      x.right.Lift (k)
    }
    x.swap (left)
  }
}


// Pre: n == number of objects in x > 0.
// Returns the former pointer to the n-th node of x,
// and this pointer is now nil.
func (x *Imp) last (n uint) *Imp {
//
  switch n { case 1:
    return x
  case 2:
    y:= x.left
    x.left = nil
    return y
  case 3:
    y:= x.right
    x.right = nil
    return y
  }
  left, k:= f (n)
  if left {
    return x.left.last (k)
  }
  return x.right.last (k)
}


func (x *Imp) Del (n uint) (Heap, Any) {
//
  y:= x.last (n)
  switch n { case 1:
    y = nil
  case 2:
    // see above
  case 3:
    y.left = x.left
  default:
    y.left = x.left
    y.right = x.right
  }
  return y, x.top
}


func (x *Imp) Sift (n uint) {
//
  if x.left != nil {
    if x.right == nil {
      if Less (x.left.top, x.top) {
        x.swap (true)
      }
    } else { // x.left != nil && x.right != nil
      if Less (x.top, x.left.top) && Less (x.top, x.right.top) {
        return
      }
      if Less (x.left.top, x.right.top) {
        x.swap (true)
        x.left.Sift (n)
      } else {
        x.swap (false)
        x.right.Sift (n)
      }
    }
  }
}


func (x *Imp) Get () Any {
//
  if x == nil {
    return nil
  }
  return Clone (x.top)
}
