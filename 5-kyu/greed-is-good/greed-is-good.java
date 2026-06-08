import java.util.HashMap;
import java.util.Map;
​
public class Greed{
  public static int greedy(int[] dice) {
        int score = 0;
        Map<Integer, Integer> freq = new HashMap<>();
​
        // Populate the frequency map
        for (int v : dice) {
            freq.put(v, freq.getOrDefault(v, 0) + 1);
        }
​
        // Calculate the score
        for (int num : dice) {
            if (num == 5) {
                while (freq.getOrDefault(num, 0) != 0) {
                    int v = freq.get(num);
                    if (v >= 3) {
                        score += 500;
                        freq.put(num, v - 3);
                    } else {
                        score += 50;
                        freq.put(num, v - 1);
                    }
                }
            } else if (num == 1) {
                while (freq.getOrDefault(num, 0) != 0) {