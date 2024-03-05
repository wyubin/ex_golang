# twoSum
[link](https://leetcode.com/problems/two-sum/description/)
```golang
twoSum(nums []int, target int) []int
```
- 輸入一個整數的 array 及目標整數，輸出其中兩個整數加起來值為目標整數的 index，不然就輸出空array
- 可用一個 hash ， iter nums 存每個 element 的 info
  - key 為 target - element value, val 為 element index
  - 所以只要後面的 element value map 到 hash 的 key 就可以輸出對應數字的 element index

# Container With Most Water
[link](https://leetcode.com/problems/container-with-most-water/description/?envType=list&envId=xoqag3yj)

## intro
輸入 []int array為多個牆面的高度，每個牆的距離為 1，任選兩面牆為 boundary 來算面積，輸出最大面積

## plan
### try
- 直接用 for loop 來做，但是要找出最大面積，所以要用兩個 pointer 來記錄，一個是 left pointer，一個是 right pointer, 再用距離來計算面積
- 這個解法有兩個 for loop...

### solution
要直接從最大寬度開始確認，直接設定兩邊 pointer，然後往中間移動，確認面積，如果面積比目前最大面積還大，就更新最大面積。
另外在往內移動兩邊 pointer 的條件是只移動高度低的，因為只有提高高度低的才可能增加面積，因為 boundary 的最小值才能形成面積，

# 3sum
[link](https://leetcode.com/problems/3sum/)

## intro
輸入一個 []int, 確認有沒有三數相加為 0, int 可以重複

## plan
### think
暴力法就是把 3sum 當作有參數的 2sum 去處理，但是時間複雜度會比 2sum 要慢很多
但 twosum 是不重複，但 3sum 是可以有重複 item 的，所以無法用 2sum 去解

### solution
- 因為複雜度比2sum 高很多，基本上只能兩個 loop
- 可以先將 []int 排序，用 i 做第一個int 的 idx, j/k 分別是第二個/第三個int 的 idx
- i 不動，j 從 i+1 開始，k 從最後一個數開始, 如果 sun > 0，代表要減少，k-- else j++, 同時 j < k

# Search in Rotated Sorted Array
[link](https://leetcode.com/problems/search-in-rotated-sorted-array/)

## intro
輸入由一個從小到大排序的 array with possibly rotated at an unknown pivot index(ex [4,5,6,7,0,1,2])，輸入 target，輸出 target 在 array 中的 index，沒有則輸出 -1

## plan
- 以 low, high 兩個 idx 去做 binary search(不斷用 mid 去確認是不是 target)
- 此數列的型態會至少有一邊是遞增的，特性是 left <= right
    - 如果 low 半邊是遞增的
        - 確認 target 在 low 跟 mid 之間的話 -> high = mid -1 else low = mid + 1
    - 不然 high 半邊一定是遞增的
        - 確認 target 在 mid 跟 high 之間的話 -> low = mid + 1 else high = mid -1 
