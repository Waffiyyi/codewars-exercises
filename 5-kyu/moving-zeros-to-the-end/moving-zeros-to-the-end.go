package kata
​
func MoveZeros(arr []int) []int {
  j := 0
  
  for i, v := range arr {
    if v != 0 {
      temp := v
      arr[i] = arr[j]
      arr[j] = temp
      j++
    }
  }
  return arr
}