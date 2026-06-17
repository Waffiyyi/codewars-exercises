package kata
​
import (
  "sort"
  "strings"
)
​
func OrderWeight(s string) string {
  if s == "" {
    return ""
  }
​
  nums := strings.Fields(s)
​
  sort.Slice(nums, func(i, j int) bool {
    wi := weight(nums[i])
    wj := weight(nums[j])
​
    if wi == wj {
      return nums[i] < nums[j]
    }
​
    return wi < wj
  })
​
  return strings.Join(nums, " ")
}
​
func weight(s string) int {
  sum := 0
​
  for _, ch := range s {
    sum += int(ch - '0')
  }
​
  return sum
}