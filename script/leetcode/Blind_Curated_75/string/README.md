# Longest Substring Without Repeating Characters
[link](https://leetcode.com/problems/longest-substring-without-repeating-characters/?envType=list&envId=xoqag3yj)
- 提供一個 string，輸出最長不重複字串的長度
- 不需要輸出 substr, 所以不用存達到最長時的結果(如果需要其實也可以放一個 list 來存，然後 return list[len-1])

## plan
- 通常的解法是用 left/right 兩個 int 紀錄 substr(s) 的 start/end, 用 map[rune]int(charIdx) 紀錄每個 char 所出現的位置
- 然後 iter right 時，
    - 如果 s[right] 已經出現在 charIdx 中, 代表已有 s[right] 重複，就要把 left 更新到 charIdx[s[right]] + 1
    - 如果 s[right] 沒有出現在 charIdx 中，就檢查是否已超過 maxLength 並更新 maxLength
    - 更新 charIdx[s[right]] = right
- 最後 return maxLength

# Longest Palindromic Substring
- [link](https://leetcode.com/problems/longest-palindromic-substring)

## plan
- 比較好的解法是用 DP 去解，也就是需要從最小範圍去紀錄是否為Palindromic 再往外延伸, 先預設 每個 i/j index 對應都 false, 如果為 Palindromic 就 true
    - 當 s[j] == s[i], then
        - i-j < 2(像是 `aa` or `aba`) 就直接 dp[i][j] = true
        - or dp[i-1][j+1] true(也就是內面那層確定是 Palindromic)，就 dp[i][j] = true 來延伸出去
    - 如果有延伸再確認 是否比maxlength 長，就可以紀錄現在最長的 maxSubstring
