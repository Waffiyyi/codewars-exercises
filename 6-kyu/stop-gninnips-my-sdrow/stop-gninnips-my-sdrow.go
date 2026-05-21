package kata
​
import "strings"
​
func SpinWords(str string) string {
  var res strings.Builder
  
  words := strings.Split(str, " ")
  
  for i, word := range words {
    if len(word) >= 5 {
      res.WriteString(reverseWord(word))
    } else {
      res.WriteString(word)
    }
    
    if i < len(words) - 1 {
      res.WriteRune(' ')
    }
  }
  return res.String()
}
​
func reverseWord(s string) string {
  var b strings.Builder
  
  for i := len(s) - 1; i >= 0; i--{
    b.WriteByte(s[i])
  }
  return b.String()
}