package kata
​
import (
  "fmt"
  "math"
)
​
func RGB(r, g, b int) string {
​
  r1, r2 := getHexDigits(r)
  g1, g2 := getHexDigits(g)
  b1, b2 := getHexDigits(b)
​
  hex := getHexCode(r1) + getHexCode(r2) + getHexCode(g1) + getHexCode(g2) + getHexCode(b1) + getHexCode(b2)
  return hex
}
​
func getHexDigits(n int) (int, int) {
  if n == 0 || n < 0 {
    return 0, 0
  }
​
  if n > 255 {
    n = 255
  }
  
  firstDigit := math.Floor(float64(n / 16))
  secondDigit := n % 16
  return int(firstDigit), secondDigit
}
​
func getHexCode(n int) string {
  hex := map[int]rune{10: 'A', 11: 'B', 12: 'C', 13: 'D', 14: 'E', 15: 'F'}
  res := ""
  if n >= 0 && n <= 9 {
    res += fmt.Sprintf("%d", n)
  } else {
    res += fmt.Sprintf("%c", hex[n])
  }
  return res
}
​