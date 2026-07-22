package kata
​
func FindEvenIndex(arr []int) int {
  sum := 0
  
  for _, v := range arr {
    sum += v
  }
  
  toLeft := 0
  toRight := 0
  
  for i := 0; i < len(arr); i++ {
      toRight = sum - toLeft - arr[i]
      if toLeft == toRight {
      return i
    }
    toLeft += arr[i]
  }
  return -1
}
​
​