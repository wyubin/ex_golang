# remove-nth-node-from-end-of-list
[link](https://leetcode.com/problems/remove-nth-node-from-end-of-list/)

# intro
輸入一個 linked list，刪除從後面數第 n 個 node，並回傳新的 linked list

## plan
基本上只是用這個題目來了解 linked list 的基本操作
### solution
- 先計算要在第幾個 node 把 next.next 接到 next
- 然後就是另一個 loop, 來 triverse 到第幾個 node 來做事

# merge_sorted_list
[link](https://leetcode.com/problems/merge-two-sorted-lists/)

# intro
有兩個 sorted linked list，輸出合併後的 linked list

## plan
### solution
一樣是利用 linked list 的 Next 特性 比較 val 後把較小的 node 接到 current node 的 next 然後往前移動

# ext
## merge_k_sorted_list
[link](https://leetcode.com/problems/merge-k-sorted-lists/)
基本上就是for loop 的 merge two linked list
