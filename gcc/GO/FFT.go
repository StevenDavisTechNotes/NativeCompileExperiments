package main

import (
  "fmt"
  "math"
)

const (
  PI = 3.1415926535897932
)

func int_log2(n int) uint {
  var log uint = 0
  
  for k:=1;k<n;k*=2 {
    log++
  }
  if n != (1 << log) {
      panic(fmt.Sprintf("FFT: Data length is not a power of 2!: %d ",n))
  }
  return log
}

func FFT_numFlops(N int) float64 {
  var Nd, logN float64 = float64(N),float64(int_log2(N))
  return (5.0*Nd-2)*logN+2.0*(Nd+1.0)
}

func FFT_transform_internal(N int, data []float64,direction int) {
  var n,bit,logn,dual int = N/2,0,0,1
  
  if n == 1 { return }
  
  logn = int(int_log2(n))
  
  if N == 0 { return }
  
  FFT_bitreverse(N, data)
  
  for bit = 0;bit < logn; bit++ {
    var w_real,w_imag float64 = 1,0
    var a,b int
    
    theta := 2.0 * float64(direction) * PI / (2.0 * float64(dual))
    s := math.Sin(theta)
    t := math.Sin(theta / 2.0)
    s2 := 2.0 * t * t
    
    for a,b = 0,0;b<n;b += 2 * dual {
      i := 2 * b
      j := 2*(b+dual)
      
      wd_real := data[j]
      wd_imag := data[j+1]
      
      data[j] = data[i] - wd_real
      data[j+1] = data[i+1] - wd_imag
      data[i] += wd_real
      data[i+1] += wd_imag
    }
    
    for a = 1;a < dual; a++ {
      {
        tmp_real := w_real - s * w_imag - s2 * w_real;
        tmp_imag := w_imag + s * w_real - s2 * w_imag;
        w_real = tmp_real;
        w_imag = tmp_imag;
      }
      for b = 0 ;b < n; b += 2 * dual {
        i := 2*(b+a)
        j := 2*(b+a+dual)
        
        z1_real := data[j]
        z1_imag := data[j+1]
        
        wd_real := w_real * z1_real - w_imag * z1_imag
        wd_imag := w_real * z1_imag + w_imag * z1_real
        
        data[j] = data[i] - wd_real
        data[j+1] = data[i+1] - wd_imag
        data[i] += wd_real
        data[i+1] += wd_imag
      }
    }
    dual *= 2
  }
}

func FFT_bitreverse(N int, data[]float64) {
  n := N/2
  nm1 := n-1
  var i,j int = 0,0
  
  for ;i < nm1; i++ {
    ii := i << 1
    jj := j << 1
    
    k := n >> 1
    
    if i < j {
      tmp_real := data[ii]
      tmp_imag := data[ii+1]
      
      data[ii] = data[jj]
      data[ii+1] = data[jj+1]
      
      data[jj] = tmp_real
      data[jj+1] = tmp_imag
    }
    
    for k<=j {
      j -= k
      k >>= 1
    }
    j += k
  }
}

func FFT_transform(N int, data[]float64){
  FFT_transform_internal(N,data,-1)
}

func FFT_inverse(N int, data[]float64) {
  n := N/2
  FFT_transform_internal(N,data, +1)
  
  norm := 1/float64(n)
  for i := 0; i<N; i++ {
    data[i] *= norm
  }
}
