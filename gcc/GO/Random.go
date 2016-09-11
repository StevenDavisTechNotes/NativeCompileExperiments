package main

const (
  MDIG = 32
  ONE = 1
  m1 = (ONE << (MDIG-2)) + ((ONE << (MDIG-2) )-ONE)
  m2 = ONE << MDIG/2
)

var dm1 float64
type Random struct {
  m[17]int
  seed int
  i int                                /* originally = 4 */
  j int                                /* originally =  16 */
  haveRange bool            /* = false; */
  left float64                          /*= 0.0; */
  right float64                         /* = 1.0; */
  width float64                         /* = 1.0; */
}

func new_Random_seed(seed int) *Random {
  R := new(Random)
  initialize(R,seed)
  R.left = 0.0
  R.right = 1.0
  R.width = 1.0
  R.haveRange = false /*false*/
                  
  return R
}

func new_Random(seed int, left,right float64) *Random {
  R := new(Random)
  
  initialize(R,seed)
  R.left = left
  R.right = right
  R.width = right - left
  R.haveRange = true          /* true */
  return R
}

func Random_nextDouble(R *Random)float64 {
  var k int
  
  var i,j = R.i,R.j
  m := R.m[:]
  
  k = m[i]-m[j]
  if k<0 { k += m1  }
  R.m[j] = k
  
  if i == 0 { i = 16 
  } else { i-- }
  R.i = i
  
  if j == 0 { j = 16 
  } else { j-- }
  R.j = j
  
  if R.haveRange {
    return R.left + dm1 * float64(k) * R.width
  } else {
    return dm1 * float64(k)
  }
  
  return 0  
}

func initialize(R *Random, seed int) {
  var jseed, k0, k1, j0, j1, iloop int

  dm1  = 1.0 / float64(m1)
  R.seed = seed
  if seed < 0 {
    seed = -seed
  }
  if seed < m1 {
    jseed = seed
  } else {
    jseed = m1
  }
  
  if jseed % 2 == 0 {
    jseed--
  }
  k0 = 9069 % m2
  k1 = 9069 / m2
  j0 = jseed % m2
  j1 = jseed / m2
  for iloop = 0;iloop < 17; iloop++ {
    jseed = j0 * k0
    j1 = (jseed / m2 + j0 * k1 + j1 * k0) % (m2 / 2)
    j0 = jseed % m2
    R.m[iloop] = j0 + m2 * j1
  }
  R.i = 4
  R.j = 16
}

func RandomVector(N int, R *Random) []float64 {
  x := make([]float64,N)
  for i := range x {
    x[i] = Random_nextDouble(R)
  }
  return x
}

func RandomMatrix(M,N int, R *Random) [][]float64 {
  a := make([][]float64,M)
  for i := range a {
    a[i] = make([]float64,N)
    for j := range a[i] {
      a[i][j] = Random_nextDouble(R)
    }
  }
  return a
}
