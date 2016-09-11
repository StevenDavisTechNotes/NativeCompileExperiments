package main

func new_Array2D_double(M,N int)[][]float64 {
  A := make([][]float64,M)
  
  for i := range A {
    A[i] = make([]float64,N)
  }
  
  return A
}

func Array2D_double_copy(M,N int, B,A [][]float64) {
  remainder := N & 3
  
  for i:=0;i<M;i++ {
    Bi := B[i]
    Ai := A[i]
    
    for j:=0;j<remainder;j++ {
      Bi[j] = Ai[j]
    }
    for j:=remainder;j<N;j+=4 {
      Bi[j] = Ai[j]
      Bi[j+1] = Ai[j+1]
      Bi[j+2] = Ai[j+2]
      Bi[j+3] = Ai[j+3]
    }
  }
}
