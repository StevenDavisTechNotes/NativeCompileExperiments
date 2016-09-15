package main

import (
  "fmt"
  "flag"
  "os"
)

const (
  RESOLUTION_DEFAULT = 2.0  /* secs (normally 2.0) */
  RANDOM_SEED = 101010

  /* default: small (cache-contained) problem sizes */

  FFT_SIZE = 1024  /* must be a power of two */
  SOR_SIZE =100 /* NxN grid */
  SPARSE_SIZE_M = 1000
  SPARSE_SIZE_nz = 5000
  LU_SIZE = 100

                                           /* large (out-of-cache) problem sizes */
  LG_FFT_SIZE = 1048576  /* must be a power of two */
  LG_SOR_SIZE =1000  /*  NxN grid  */
  LG_SPARSE_SIZE_M = 100000
  LG_SPARSE_SIZE_nz =1000000
  LG_LU_SIZE = 1000
  /* tiny problem sizes (used to mainly to preload network classes     */
  /*                     for applet, so that network download times    */
  /*                     are factored out of benchmark.)               */
  /*                                                                   */
  TINY_FFT_SIZE = 16  /* must be a power of two */
  TINY_SOR_SIZE =10 /* NxN grid */
  TINY_SPARSE_SIZE_M = 10
  TINY_SPARSE_SIZE_N = 10
  TINY_SPARSE_SIZE_nz = 50
  TINY_LU_SIZE = 10
)
var large bool
var min_time float64
func init() {
  flag.BoolVar(&large,"large",false,"large values")
  flag.Float64Var(&min_time,"mtime",RESOLUTION_DEFAULT,"minimum time")
}

func main(){

  var (
    FFT_size int = FFT_SIZE
    SOR_size int =  SOR_SIZE
    Sparse_size_M int = SPARSE_SIZE_M
    Sparse_size_nz int = SPARSE_SIZE_nz
    LU_size int = LU_SIZE
    NumTimes = 50

    res[6]float64
  )
  R := new_Random_seed(RANDOM_SEED)
  flag.Parse()
  if large {
      FFT_size = LG_FFT_SIZE
      SOR_size = LG_SOR_SIZE
      Sparse_size_M = LG_SPARSE_SIZE_M
      Sparse_size_nz = LG_SPARSE_SIZE_nz
      LU_size = LG_LU_SIZE
  }
  print_banner()
  fmt.Printf("Using %10.2f seconds min time per kernel.\n", min_time)

  f, err := os.OpenFile("../../ResultLog.txt", os.O_APPEND|os.O_WRONLY, 0600)
  if err != nil {
    panic(err)
  }

  defer f.Close()

  for iTime := 0; iTime < NumTimes; iTime++ {
    res[1] = kernel_measureFFT( FFT_size, min_time, R)
    res[2] = kernel_measure_SOR(SOR_size,min_time, R)
    res[3] = kernel_measureMonteCarlo(min_time, R)
    res[4] = kernel_measureSparseMatMult( Sparse_size_M,
                Sparse_size_nz, min_time, R)
    res[5] = kernel_measureLU( LU_size, min_time, R)

    res[0] = (res[1]+res[2]+res[3]+res[4]+res[5])/5

    if _, err = f.WriteString(fmt.Sprintf("UbuntuVM,GCCGO,%.2f\n", res[0])); err != nil {
      panic(err)
    }

  fmt.Printf("Composite Score:        %8.2f\n" ,res[0]);
  fmt.Printf("FFT             Mflops: %8.2f    (N=%d)\n", res[1], FFT_size)
  fmt.Printf("SOR             Mflops: %8.2f    (%d x %d)\n",
                                  res[2], SOR_size, SOR_size)
  fmt.Printf("MonteCarlo:     Mflops: %8.2f\n", res[3])
  fmt.Printf("Sparse matmult  Mflops: %8.2f    (N=%d, nz=%d)\n", res[4],
                                          Sparse_size_M, Sparse_size_nz)
  fmt.Printf("LU              Mflops: %8.2f    (M=%d, N=%d)\n", res[5],
                                  LU_size, LU_size)
  }
}

func print_banner(){
 fmt.Println("**                                                              **")
 fmt.Println("** SciMark2 Numeric Benchmark, see http://math.nist.gov/scimark **")
 fmt.Println("** for details. (Results can be submitted to pozo@nist.gov)     **")
 fmt.Println("**                                                              **")
}

