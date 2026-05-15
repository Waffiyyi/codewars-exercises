package kata
​
func Josephus(items []interface{}, k int) []interface{} {
  res := []any{}
  
  count := 1
  i := 0
  
  for len(items) > 0 {
    if count == k {
      res=append(res, items[i])
      items = append(items[:i], items[i+1:]...)
      count = 1
      if i >= len(items){
      i = 0
     }
      continue
    }
    
    
    count++
    i++
    
    if i >= len(items){
      i = 0
    }
  }
  return res
}