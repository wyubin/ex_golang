# remove-nth-node-from-end-of-list
[link](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

# intro
輸入一個 linked list，刪除從後面數第 n 個 node，並回傳新的 linked list

## plan
基本上只是用這個題目來了解 linked list 的基本操作
### solution
- 先計算要在第幾個 node 把 next.next 接到 next
- 然後就是另一個 loop, 來 triverse 到第幾個 node 來做事
