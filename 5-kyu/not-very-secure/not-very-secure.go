package kata
​
func alphanumeric(str string) bool {
  if len(str) == 0 {
    return false
  }
   for _, ch := range str {
     if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
       continue
     } else if ch >= '0' && ch <= '9' {
         continue
       } else {
         return false
       }
     }
  return true
}