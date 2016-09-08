package main

import "math"

func LU_num_flops(N int)float64 {
  Nd := float64(N)
  
  return 2.0 * Nd * Nd * Nd / 3.0
}

func LU_factor(M,N int, A[][]float64,pivot []int)int {
  var minMN int
  if M < N {
    minMN = M
  } else {
    minMN = N
  }
  
  for j:=0;j<minMN;j++ {
    jp := j
    t := math.Abs(A[j][j])
    for i:=j+1;i<M;i++ {
      ab := math.Abs(A[i][j])
      if ab > t {
        jp = i
        t = ab
      }
    }
    
    pivot[j] = jp;
    
    if A[jp][j] == 0 {
      return 1
    }
    
    if jp != j {
      tA := A[j]
      A[j] = A[jp]
      A[jp] = tA
    }
    
    if j < M-1 {
      recp := 1.0/A[j][j]
      for k:=j+1;k<M;k++ {
        A[k][j] *= recp
      }
    }
    
    if j < minMN-1 {
      for ii:=j+1;ii<M;ii++ {
        Aii := A[ii]
        Aj := A[j]
        AiiJ := Aii[j]
        
        for jj := j+1; jj<N;jj++ {
          Aii[jj] -= AiiJ * Aj[jj]
        }
      }
    }
    
  }
  
  return 0
}


