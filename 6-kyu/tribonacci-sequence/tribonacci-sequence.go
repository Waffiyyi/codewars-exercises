package kata
​
func Tribonacci(signature [3]float64, n int) []float64 {
  if n == 0 {
    return []float64{}
  }
  res := make([]float64, n + len(signature))
  track := 0
  i, j, k := signature[track], signature[track+1], signature[track+2]
  res[0], res[1], res[2] = i, j, k
  next := i + j + k
​
  for in := 3; in < n; in++ {
    res[in] = next
    track++
    i, j, k = res[track], res[track+1], res[track+2]
    next = i + j + k
  }
  return res[:n]
}