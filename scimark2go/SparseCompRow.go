package main

func SparseCompRow_num_flops(N,nz,num_iterations int)float64 {
  actual_nz := (nz/N)*N
  return float64(actual_nz)*2.0*float64(num_iterations)
}

func SparseCompRow_matmult(M int, y,val []float64, 
    row,col []int, x []float64, NUM_ITERATIONS int) {
    for reps := 0;reps < NUM_ITERATIONS;reps++ {
      for r := 0;r < M;r++ {
        sum := 0.0
        rowR := row[r]
        rowRp1 := row[r+1]
        for i:=rowR;i<rowRp1;i++ {
          sum += x[ col[i] ] * val[i]
        }
        y[r] = sum
      }
    }
}
