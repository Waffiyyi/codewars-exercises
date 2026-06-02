package kata
​
import "strings"
​
func toWeirdCase(str string) string {
   var res strings.Builder
    i := 0
   newWord := false
   for _, ch := range str {
     if ch == ' ' {
       res.WriteRune(ch)
       newWord = true
       continue
     }
     if newWord && i % 2 != 0 {
       i++
     }
     if i % 2 != 0 {
       res.WriteString(strings.ToLower(string(ch)))
     } else {
       if newWord {
         newWord = !newWord
       }
       res.WriteString(strings.ToUpper(string(ch)))
     }
     i++
   }
    return res.String()
}