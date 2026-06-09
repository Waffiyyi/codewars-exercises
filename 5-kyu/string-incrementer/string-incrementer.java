import java.math.BigInteger;
​
public class Kata {
  public static String incrementString(String s) {
        StringBuilder numInStr = new StringBuilder();
​
        // Iterate backwards to find only the trailing digits
        for (int i = s.length() - 1; i >= 0; i--) {
            char ch = s.charAt(i);
            if (ch >= '0' && ch <= '9') {
                numInStr.append(ch);
            } else {
                break;
            }
        }
​
        if (numInStr.length() == 0) {
            return s + "1";
        }
​
        int l = numInStr.length();
        
        // Extract the non-number prefix
        String str = s.substring(0, s.length() - l);
​
       BigInteger num = new BigInteger(numInStr.reverse().toString());
num = num.add(BigInteger.ONE);
​
String incremented = num.toString();
​
while (incremented.length() < l) {
    incremented = "0" + incremented;
}
​
return str + incremented;
    }
}