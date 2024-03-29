# Longest Substring Without Repeating Characters
[link](https://leetcode.com/problems/longest-substring-without-repeating-characters/?envType=list&envId=xoqag3yj)
- 提供一個 string，輸出最長不重複字串的長度
- 不需要輸出 substr, 所以不用存達到最長時的結果(如果需要其實也可以放一個 list 來存，然後 return list[len-1])

## intro
```shell
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

## plan
- 通常的解法是用 left/right 兩個 int 紀錄 substr(s) 的 start/end, 用 map[rune]int(charIdx) 紀錄每個 char 所出現的位置
- 然後 iter right 時，
    - 如果 s[right] 已經出現在 charIdx 中, 代表已有 s[right] 重複，就要把 left 更新到 charIdx[s[right]] + 1
    - 如果 s[right] 沒有出現在 charIdx 中，就檢查是否已超過 maxLength 並更新 maxLength
    - 更新 charIdx[s[right]] = right
- 最後 return maxLength

# Longest Palindromic Substring
- [link](https://leetcode.com/problems/longest-palindromic-substring)

## intro
給一串 string，輸出其中的最長 palindrome substring
```shell
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
```

## plan
- 比較好的解法是用 DP 去解，也就是需要從最小範圍去紀錄是否為Palindromic 再往外延伸, 先預設 每個 i/j index 對應都 false, 如果為 Palindromic 就 true
- sliding windows 中(idxEnd, idxStart), 兩層 for
    - 當 s[j] == s[i], then
        - i-j < 2(像是 `aa` or `aba`) 就直接 dp[i][j] = true
        - or dp[i-1][j+1] true(也就是內面那層確定是 Palindromic)，就 dp[i][j] = true 來延伸出去
    - 如果有延伸再確認 是否比maxlength 長，就可以紀錄現在最長的 maxSubstring

# valid_parentheses
[link](https://leetcode.com/problems/valid-parentheses)
有幾種括號是對應的 '(', ')', '{', '}', '[' and ']', 依據輸入的字串來確認是否有對應的括號

## intro
```shell
Input: s = "()[]{}"
Output: true
Input: s = "(]"
Output: false
```

## plan
- 用一個map[byte]byte 來紀錄 括號的對應
- 用 stack(後進先出) 紀錄應該後括號出現的順序，
    - 如果是前括號，就append對應的後括號
    - 如果不是前括號
        - 遇到 stack 為空，就代表沒有對應的括號，回傳 false
        - 就 pop 一個括號，來確認 current string 是否為對應的括號
最後結束 loop 後，也要確認 stack 為空，回傳 true

# Group Anagrams
將提供的 []string 分組, 擁有相同set 的 char 的 string 就是同一組

## intro
```shell
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
```

# plan
- 寫一個 func 把 string 轉成一個 sorted char 的string 作為 unique key
- 用 sortedChar map[string][]string 去紀錄 相同 sorted char 的 string
- 輸出 sortedChar 的 value

# minimun window substring
[link](https://leetcode.com/problems/minimum-window-substring)

## intro
針對一個字串 s，給定一個較小字串 t，輸出一個在 s 可以找到的最小 window會包含所有的 t char
```shell
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
```

## plan
- 如果只是從兩邊 narrow down 直到區域中字串完全包含 t，那可能會找到錯誤區段，所以的確還是需要 iteration 去找真正最小的 substring
- 用 byte2count 來存 t char 的出現次數, 並且init minlen 跟 start_idx 來標記出現最小 sunstring 的資料
- count 來紀錄目前在 s 還需要掃的 t char
- 最外面的 iter 是 把 end 往後，當 curr byte 是存在 t char 時，就減少 byte2count 及 count
 - 直到 count ==0 時，就開始移動 start, 直到 currbyte 有在 byte2count 時, 再將 byte2count 及 count ++
 - 然後 end 會再開始往前移動直到 len(s)-1

# Decode Ways
[link](https://leetcode.com/problems/decode-ways)

## intro
decode 可以是 string mapping 的方式，定義 A-Z 的 encode 是 1-26 的字串，給定一串數字，輸出有幾種decode 方法

```shell
Input: s = "12"
Output: 2
```

## plan
- decode 的方法數也是可以前後加起來的，所以一樣可以用 DP 來解
- 例外: numstr[0] == "0" 是不會有 code 的, return 0
- dp[n] 是存解到 numstr[:n] 的方法
- codon map[string]struct{}, 紀錄能夠作為 codon 的字串
- 有點像每次只能走一步或兩步的 stair 解法，但每次還需要確認一步或兩步是不是在 mapping 中

# Valid Palindrome
[link](https://leetcode.com/problems/valid-palindrome)

## intro
給一個字串，若除去 Alphanumeric 以外的字元後，字串是否回文
```shell
Input: s = "A man, a plan, a canal: Panama"
Output: true
```

## plan
- 可以先將本來字串做 lower case
- 定兩個 pointer 分別是 left/right 是開始及結束
    - 如果不是 alphanumeric 就移動 pointer
    - 再來比較是否相等, 如果不是就直接 return false
- 最後 return true


# Word break
[link](https://leetcode.com/problems/word-break)

## intro
給定一個 string 來判斷是否能用 word list 去完整 break 這個 string
```shell
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true

Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false
```

## plan
- 先 init maxDictWord 來紀錄可用的最長 word
- 用 DP(len(string)+1) 去記錄在 string 的 j:i 位置能不能 break(也就是存在於 dictWord), 如果存在 wordDict 就可以在 DP[i] = true 代表這邊是可以斷點的，所以當 DP[len(string)] = true 就代表 string 可以被完整切成多個 word

- iter
    - for loop i from 1 to len(string)
        - for loop j from i-1 to i-1-maxDictWord, 也就是一次檢查一個word
            - 如果 DP[j] == true && string[j:i] in wordDict 就代表這裡可以是斷點, DP[i] = true, 然後 break
以上的 iter 等於 用一個 可變的window(j:i) 去看 string 來紀錄累積的斷點直到最後
- 最後 return DP[len(string)]


# valid anagram
[link](https://leetcode.com/problems/valid-anagram)

## intro
輸入兩個 string, 測試兩者是否為 anagram
```shell
Input: s = "anagram", t = "nagaram"
Output: true
```

## plan
簡單用 char2len map[rune]int 用兩個 forloop 就可以解
- 第一個 for 把 a 的 rune ++, b 的 rune --
- 第二個 for 把 char2len 中如果不是 0 就 return false
- return true

# encode and decode string
[link](https://leetcode.com/problems/encode-and-decode-strings)

## intro
實作 encode 跟 decode func, 能夠將一個 []str 轉成單一 string 再轉回同樣的 []str

## plan
直接看實作的解法

# longest Repeating character replacement
[link](https://leetcode.com/problems/longest-repeating-character-replacement/)

## intro
給一個s都大寫的字串，還有一個k表示最多可以 replace 幾個 char, 求最長的同一字母的長度
```shell
Input: s = "ABAB", k = 2
Output: 4

Input: s = "AABABBA", k = 1
Output: 4
```

## plan
- 是一個 window size 的問題，所以用雙指針 idxStart, idxEnd 來移動 windows, 每次移動 idxEnd 都要紀錄每個字母的 countChar, 並計算最大 maxRepeat
- 當 (idxEnd-idxStart+1) - maxRepeat > k, 就要移動 idxStart 並紀錄 countChar
- 然後 idxEnd++ 直到 == len

# palindrome substring
[link](https://leetcode.com/problems/palindromic-substrings)

## intro
給一串 string，輸出其中的所有 palindrome substring
```shell
Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
```


## plan
可以參考 Longest Palindromic Substring 的作法，以 windows scan 配合 dp 的方法來輸出每個window的 palindrome substring
- init [][]bool(dp) 來存 j到i是否為 palindrome
- 用 res []string 來紀錄所有的 palindrome substring
- iter idxEnd 時，每一個對應的 char 都可以塞進 res
    - iter idxStart 時，當 s[idxStart] == s[idxEnd] && (idxEnd-idxStart == 1 || dp[idxStart+1][idxEnd-1]) 代表是 palindrome
        - s[idxStart:(idxEnd+1)]就塞進 res
        - dp[idxStart][idxEnd] = true
