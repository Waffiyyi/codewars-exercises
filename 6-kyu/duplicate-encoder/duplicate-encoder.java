import java.util.*;
​
public class DuplicateEncoder {
  static String encode(String word){
    word = word.toLowerCase();
    Map<Character, Integer> m = new HashMap<>();
    
    String res = "";
    
    for(Character ch: word.toCharArray()) {
      m.put(ch, m.getOrDefault(ch, 0) + 1);
    }
    
    for(int i = 0; i < word.length(); i++) {
      if(m.get(word.charAt(i)) > 1 ) {
        res += ")";
      } else {
        res += "(";
      }
    }
    return res;
  }
}
​