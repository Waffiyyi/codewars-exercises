public class PigLatin {
    public static String pigIt(String s) {
        StringBuilder res = new StringBuilder();
        String[] words = s.split(" ", -1);
        String rot = "ay";
        String punct = ".,:;!?";
​
        for (int i = 0; i < words.length; i++) {
            String word = words[i];
​
            if (word.length() == 1 && punct.contains(word)) {
                res.append(word);
                continue; 
            }
​
            String first = word.substring(0, 1);
            String pig = word.substring(1) + first + rot;
​
            res.append(pig);
​
            if (i < words.length - 1) {
                res.append(' ');
            }
        }
​
        return res.toString();
    }
}