# Validate Binary Search Tree
[link](https://leetcode.com/problems/validate-binary-search-tree/)

## intro
給定一串 int 並轉成 Binay tree 並確認其是否為 valid BST, 若是則回傳 true，否則回傳 false
BST: left < root < right

## plan
- 建立一個 recursive func 確認 node 是否在一個 int 範圍內
  - 如果 node 為 null，回傳 true
  - 如果 node 的值不在範圍內，回傳 false
  - 其他則是接下去看 leftnode 跟rightnode
    - leftnode 範圍在 min < node.val < root
    - rightnode 範圍在 root < node.val < max
- 起始範圍直接跟題目指定一樣 (1, 10^4)

# Same Tree
[link](https://leetcode.com/problems/same-tree/)

## intro
判斷兩個 tree 是否相同

## plan
- 一樣 recursive func(sameNode), 同時輸入兩個tree 的 node 進行同步 retrival
- 如果有一邊 是 nil, 那就輸出兩邊是否同為 nil
- 否則就比較兩邊的 val, 如果 不同就 return false
- 否則接下去輸出兩個 node 的 left 跟 兩個 right node 的 sameNode

# Binary Tree Level Order Traversal
[link](https://leetcode.com/problems/binary-tree-level-order-traversal/)

## intro
給一個 tree，輸出其 level order 的結果

## plan
- 看起來一樣是操作，用一個 array 來 append 每次的 level 的 node values，然後把 array 輸出

# Maximum Depth of Binary Tree
[link](https://leetcode.com/problems/maximum-depth-of-binary-tree/)

## intro
給一個 tree，輸出其 depth

## plan
- recursive func(maxDepth), 輸入 tree 的 node，並把 node 的 depth 輸出
    - 如果 node 為 nil，回傳 0
    - 否則就 recursive 找比較長的 depth
    - +1 輸出
