242. Valid Anagram
Easy


Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

 

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
 

Constraints:

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.
 

Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

SOLUTION

func isAnagram(s string, t string) bool {
    // Check if the lengths are different
    if len(s) != len(t) {
        return false
    }

    // Create maps to store the frequency of each character in both strings
    sFreq := make(map[rune]int)
    tFreq := make(map[rune]int)

    // Update frequency for string s
    for _, char := range s {
        sFreq[char]++
    }

    // Update frequency for string t
    for _, char := range t {
        tFreq[char]++
    }

    // Compare the frequency of each character in both strings
    for char, freq := range sFreq {
        if tFreq[char] != freq {
            return false
        }
    }

    // If all characters have the same frequency, the strings are anagrams
    return true
}

func main() {
    // Example usage
    result1 := isAnagram("anagram", "nagaram")
    fmt.Println(result1) // Output: true

    result2 := isAnagram("rat", "car")
    fmt.Println(result2) // Output: false
}


