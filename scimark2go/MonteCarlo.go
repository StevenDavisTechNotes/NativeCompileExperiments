package main

const (
  SEED = 113
)

func MonteCarlo_numFlops(NumSamples int)float64 {
  return float64(NumSamples)*4.0
}

func MonteCarlo_integrate(NumSamples int)float64 {
  R := new_Random_seed(SEED)
  
  under_curve := 0
  
  for count := 0; count < NumSamples;count++ {
    x := Random_nextDouble(R)
    y := Random_nextDouble(R)
    
    if x*x + y*y <= 1.0 {
      under_curve++
    }
  }
  
  return (float64(under_curve)/float64(NumSamples))*4.0
}
