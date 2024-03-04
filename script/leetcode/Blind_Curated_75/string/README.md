# Longest Substring Without Repeating Characters
- 提供一個 string，輸出最長不重複字串的長度
- 不需要輸出 substr, 所以不用存達到最長時的結果(如果需要其實也可以放一個 list 來存，然後 return list[len-1])

## plan
- 通常的解法是用 left/right 兩個 int 紀錄 substr(s) 的 start/end, 用 map[rune]int(charIdx) 紀錄每個 char 所出現的位置
- 然後 iter right 時，
    - 如果 s[right] 已經出現在 charIdx 中, 代表已有 s[right] 重複，就要把 left 更新到 charIdx[s[right]] + 1
    - 如果 s[right] 沒有出現在 charIdx 中，就檢查是否已超過 maxLength 並更新 maxLength
    - 更新 charIdx[s[right]] = right
- 最後 return maxLength

