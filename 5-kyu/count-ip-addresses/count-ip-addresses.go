package kata
​
import (
  "math"
  "strconv"
  "strings"
)
​
func IpsBetween(start, end string) uint32 {
    f := getIpValue(start)
    s := getIpValue(end)
  
    if f > s {
      return uint32(f - s)
    }
    return uint32(s - f)
}
​
func getIpValue(s string) int {
  ips := strings.Split(s, ".")
  res := 0
  exponent := 0
  base := 256
  
  for i:= len(ips) - 1; i >= 0; i-- {
    num, _ := strconv.Atoi(ips[i])
    calc := num * int(math.Pow(float64(base), float64(exponent)))
    res += int(calc)
    exponent++
  }
  return res
}