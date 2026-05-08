package kata
​
import (
  "math"
)
​
func Beeramid(bonus int, price float64) int {
  p := float64(price)
  cans := int(math.Floor(float64(float64(bonus) / p)))
  count := 1.0
  pow := 2.0
  res := 0
​
  for cans > 0 {
    //remove one possible pyramid from cans
    cans -= int(math.Pow(count, pow))
​
    //break early to avoid extra can
    if cans < 0 {
      break
    }
​
    //calculate next pyramid
    count++
​
    //count of pyramids built
    res++
  }
  return res
}
​