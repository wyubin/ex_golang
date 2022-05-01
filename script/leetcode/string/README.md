# anagram_valid
- 確認兩個字串是否有相同字母及次數
```golang
func isAnagram(s string, t string) bool
```
- 用一個 `map[byte]int{}` 來存第一個字串的byte 及存在次數
- 第二個字串經確認 map 的次數沒有被減到 0或不存在就可以 -1 並continue

# common_prefix
- 確認多個字串的最長相同前綴
```golang
func longestCommonPrefix(strs []string) string
```
- 先把最短排前面
-逐個比較如果看到不同的就輸出 strs[0][:indCompare]

# palindrome
```golang
func isPalindrome(s string) bool
```
- 先從前後過濾掉不能算的部份再進行依次比對，只要都一樣就輸出 ture

# reverse_integer
- 主要是把 '-153' -> '-351'
```golang
func reverse(x int) int
```
- 確定正負數後，先轉成絕對值，轉成 []int 之後，再乘回反向的整數
```golang
for _, val := range arrInt {
		y += val * fold
		fold *= 10
	}
```
# reverse_string
- 這邊更單純是反轉字串
- 先轉成[]byte ，只要少於一半的 index 就用swap 做代換，完成再轉回 string

# uniq_chr_2
- 輸出字串中第一個非重複的單字，最簡單的方式就是用 map[byte]int{} 計算再 for loop string 去算

