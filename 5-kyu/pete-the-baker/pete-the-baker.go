package kata
​
import (
  "math"
)
​
func Cakes(recipe, available map[string]int) int {
​
  var res float64
​
  for k := range recipe {
    if _, ok := available[k]; !ok {
      return 0
    }
  }
​
  i := 0
​
  for k, v := range available {
    if elem, ok := recipe[k]; !ok {
      continue
    } else {
      if res == 0 && i == 0 {
        res = math.Max(float64(math.Floor(float64(v/elem))), res)
      } else {
        res = math.Min(float64(math.Floor(float64(v/elem))), res)
      }
    }
    i++
  }
​
  return int(res)
}
​