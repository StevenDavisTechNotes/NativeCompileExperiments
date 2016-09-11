package main
import (
  "time"
)

func kernel_measureFFT(N int, mintime float64, R *Random) float64{
  twoN := 2*N
  x := RandomVector(twoN,R)
  
  var cycles int = 1
  var result float64 = 0
  var start time.Time
  for {
    start = time.Now()
    for i:=0;i<cycles;i++ {
      FFT_transform(twoN,x)
      FFT_inverse(twoN, x)
    }
    end := time.Since(start)
    if end.Seconds() >= mintime {
      break
    }
    cycles *= 2
  }
  result = FFT_numFlops(N)*float64(cycles)/time.Since(start).Seconds()*1e-6
  return result
}

func kernel_measure_SOR(N int, min_time float64, R *Random)float64 {
  G := RandomMatrix(N,N,R)
  result := 0.0
  var cycles int
  var start time.Time
  for cycles = 1;;cycles*=2{	
    start = time.Now()
    SOR_execute(N,N,1.25,G,cycles)
    end := time.Since(start)
    if end.Seconds() >= min_time {
      break
    }
  }
  result = SOR_num_flops(N,N,cycles)/time.Since(start).Seconds()*1e-6
  return result
}

func kernel_measureMonteCarlo(min_time float64, R *Random)float64 {
  var result float64
  
  var cycles int
  var start time.Time
  for cycles = 1;;cycles*=2 {
    start = time.Now()
    MonteCarlo_integrate(cycles)
    end := time.Since(start)
    
    if end.Seconds() >= min_time {
      break
    }
  }
  result = MonteCarlo_numFlops(cycles)/time.Since(start).Seconds()*1e-6
  return result
}

func kernel_measureSparseMatMult(N,nz int, min_time float64, R *Random)float64 {
  x := RandomVector(N,R)
  y := make([]float64,N)
  
  result := 0.0
  
  nr := nz/N
  anz := nr * N
  
  val := RandomVector(anz, R)
  col := make([]int,nz)
  row := make([]int,N+1)
  
  var cycles int
  for r:=0;r<N;r++ {
    rowr := row[r]
    step := r/nr
    
    row[r+1] = rowr + nr
    
    if step < 1 {
      step = 1
    }
    
    for i:=0;i<nr;i++ {
      col[rowr+i]=i*step
    }
  }
  var start time.Time
  for cycles = 1;;cycles*=2 {
    start = time.Now()
    SparseCompRow_matmult(N,y,val,row,col,x,cycles)
    end:= time.Since(start)
    
    if end.Seconds() >= min_time {
      break
    }
  }
  
  result = SparseCompRow_num_flops(N, nz, cycles)/time.Since(start).Seconds()*1e-6
  return result
}

func kernel_measureLU(N int, min_time float64, R *Random)float64 {
  result := 0.0
  var cycles int
  var start time.Time
  
  A := RandomMatrix(N,N,R)
  lu := new_Array2D_double(N,N)
  pivot := make([]int,N)
  
  for cycles=1;;cycles*=2 {
    start = time.Now()
    for i:=0;i<cycles;i++ {
      Array2D_double_copy(N, N, lu, A)
      LU_factor(N, N, lu, pivot)
    }
    
    end := time.Since(start)
    
    if end.Seconds() >= min_time {
      break
    }
  }
  
  result = LU_num_flops(N) * float64(cycles) / time.Since(start).Seconds() * 1e-6
  
  return result
}
