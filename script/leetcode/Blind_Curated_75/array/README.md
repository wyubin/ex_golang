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


## plan
### try
應該直接用一個 for loop 來做，但是要找出最大面積，所以要用兩個 pointer 來記錄，一個是 left pointer，一個是 right pointer
