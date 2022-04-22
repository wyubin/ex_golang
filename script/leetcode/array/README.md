# twoSum
```golang
twoSum(nums []int, target int) []int
```
- 輸入一個整數的 array 及目標整數，輸出其中兩個整數加起來值為目標整數的 index，不然就輸出空array
- 可用一個 hash ， iter nums 存每個 element 的 info
  - key 為 target - element value, val 為 element index
  - 所以只要後面的 element value map 到 hash 的 key 就可以輸出對應數字的 element index
  
# best_income_buy_and_sell
-  int array 為市場價格，是否可算出在時候買入及賣出會有的最大收益 
```golang
maxProfit(prices []int) int
```
- 一樣 iter prices, 初始 valBuy 買入價格是 -1
  - 如果到最後或是下一個價格會降
    - 就要看目前是不是買入的狀態，要賣出，而且把收益加上去
  - else 則確定是不是非買入，就買入等漲價
- 最後輸出收益

# contains_duplicate
- 看 array 裡面有沒有重複值，通常都是用 hash ，省記憶體又容易理解
- golang 可以把value 以 struct{} 指定，省記憶體

# intersection_two_array
- 計算兩個 array 裡面有多少相同的值並輸出
```golang
intersect(nums1 []int, nums2 []int) []int
```
- 先從小排到大
- 用短的去比長的 iter 短的
  - while val1 不大於 val2，就iter ind2 去看有沒有 val1==val2 然後 append 並 break 
    - 當 num2 到最後了也要 break

