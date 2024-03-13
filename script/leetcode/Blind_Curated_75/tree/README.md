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

# Construct Binary Tree from Preorder and Inorder Traversal
[link](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

## intro
給一個 preorder 和 inorder，輸出一個 tree

## plan
- preorder: root -> left -> right
- inorder: left -> root -> right, 所以可以用 inorder 的 val 建立 map 存 inorder 的 index
- 建立 recursive func(preorder, inorderMap, currIdx, leftIdx, rightIdx), 並把 preorder 的 list 傳入, leftIdx 跟 rightIdx 是指此範圍的 inorder values, 如果 leftIdx > rightIdx 就回傳 nil


# binary tree maximum path sum
[link](https://leetcode.com/problems/binary-tree-maximum-path-sum/)

## intro
給一個 BST, 從裡面找出最大的 path sum， path 就是一個不重複走的線

## plan
- 一樣會是一個 recursive func(maxGain), 計算以這個 node 為 root 的 maxSum
  - 當 node 為 nil 時直接 return 0(等於不走)
  - 計算 currMaxSum
    - leftGain 為 max(maxGain(leftNode), 0), 如果 0 就是不走左邊
    - rightGain 同理
    - currMaxSum 可以是 leftGain + node.Val + rightGain, 如果 > maxPathSum 就更新
  - return MaxBranch, 如果是要 return branch 給parent 接，就只能在 leftGain/rjghtGain 選一條線, MaxBranch = node.val + max(leftGain, rjghtGain)