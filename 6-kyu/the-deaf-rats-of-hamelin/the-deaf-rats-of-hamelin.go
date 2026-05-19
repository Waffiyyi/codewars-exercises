package kata
​
​
func CountDeafRats(town string) int {
      hasSeenP := false 
      astray := 0
      tail := '~'
      head := 'O'
      
  
      for i := 0; i < len(town); i+=2{
        if town[i] == ' ' {
      i=    i-1
          continue
        } 
        
        if rune(town[i]) == 'P' {
          hasSeenP = true
          i=i-1
          continue
        }
        
        if i + 1 < len(town){
          if !hasSeenP && (rune(town[i]) == head && rune(town[i+1]) == tail){
          astray++
          
       
        }else if hasSeenP && (rune(town[i]) == tail && rune(town[i+1]) == head){
          astray++
        }
        }
           
      }
          return astray
}
​
​