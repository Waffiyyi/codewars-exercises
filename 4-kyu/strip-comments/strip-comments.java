public class StripComments {
​
public static String stripComments(String s, String[] commentSymbols) {
        // -1 limit ensures trailing empty lines aren't discarded
        String[] words = s.split("\n", -1);
        StringBuilder res = new StringBuilder();
​
        for (int i = 0; i < words.length; i++) {
            String word = words[i];
            int minIndex = -1;
​
            for (String symbol : commentSymbols) {
                int index = word.indexOf(symbol);
                if (index != -1) {
                    if (minIndex == -1 || index < minIndex) {
                        minIndex = index;
                    }
                }
            }
​
            if (minIndex != -1) {
                res.append(word.substring(0, minIndex).stripTrailing());
            } else {
                res.append(word.stripTrailing());
            }
​
            if (i < words.length - 1) {
                res.append('\n');
            }
        }
        
        return res.toString();
    }
  
}
​