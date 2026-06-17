package kata
​
​
func Xbonacci(signature []int, n int) []int {
  r := signature[:]
  for i := 0; i < n; i++ {
    l := len(r)
    r = append(r, sumN(r[i:l]))
  }
  return r[:n]
}
​
func sumN(nums []int) int {
  sum := 0
  for _, v := range nums {
    sum += v
  }
  return sum
}
​