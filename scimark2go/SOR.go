package main

func SOR_num_flops(M,N,num_iterations int)float64 {
  Md := float64(M)
  Nd := float64(N)
  num_iterD := float64(num_iterations)
  
  return (Md-1)*(Nd-1)*num_iterD*6.0
}

func SOR_execute(M,N int,omega float64,G[][]float64,num_iterations int) {
  omega_over_four := omega * 0.25
  one_minus_omega := 1 - omega
  
  Mm1 := M-1
  Nm1 := N-1
  var Gi,Gim1,Gip1 []float64
  
  for p:=0;p<num_iterations;p++ {
    for i:=1;i<Mm1;i++ {
      Gi = G[i][:]
      Gim1 = G[i-1][:]
      Gip1 = G[i+1][:]
      for j:=1;j<Nm1;j++ {
        Gi[j] = omega_over_four * (Gim1[j] + Gip1[j] + Gi[j-1] +
                                  Gi[j+1]) + one_minus_omega * Gi[j]
      }
    }
  }
}

