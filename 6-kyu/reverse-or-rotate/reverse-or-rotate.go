package kata
​
import (
  "strconv"
  "strings"
)
​
func Revrot(s string, n int) string {
    if (n <= 0 || n > len(s)) || s == "" {
      return ""
    }
  
    strNums := chunk(s,n)
  
    var res strings.Builder
  
  for i := 0; i < len(strNums); i++ {
    if len(strNums[i]) == n {
      if shouldRev(strNums[i]) {
        res.WriteString(reverse(strNums[i]))
      } else {
        res.WriteString(rotate(strNums[i]))
      }
    }
  }
  return res.String()
}
​
func chunk(s string, n int) []string {
  res := []string{}
  
  for i := 0; i < len(s); i+=n {
    end := i + n
    
    if end > len(s) {
      end = len(s)
    }
    res = append(res, s[i:end])
  }
  return res
}
​
func reverse(s string) string {