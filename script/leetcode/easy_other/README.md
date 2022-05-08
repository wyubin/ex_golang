- 基於 [leetcode](https://leetcode.com/problemset/all/?difficulty=EASY&page=1&listId=wpwgkgt) 還沒解的去解
# Roman to Integer
- [leetcode link](https://leetcode.com/problems/roman-to-integer/)
- 解法基本上就是看如果 ind level 比下一個少就是減，不然預設是加

# Valid Parentheses
- [leetcode link](https://leetcode.com/problems/valid-parentheses/)
- 需要用一個 stack 去做，如果是 start 就append，如果是end，那 stack pop 就應該是對應的start，直到最後 stack empty 才是 true

# Merge Two Sorted Lists
- [leetcode link](https://leetcode.com/problems/merge-two-sorted-lists/)
- 因為 linkedList 主要是用 pointer 去串，所以更容易把 Next 指到下一個 item

# Implement strStr()
- [leetcode link](https://leetcode.com/problems/implement-strstr/)
- 拿 haystack 的部份字串做比對就可以

# Maximum Subarray
- [leetcode link](https://leetcode.com/problems/maximum-subarray/)
- 在這邊偏向用暴力解，遍歷每個連續區段找最大

# Sqrt(x)
- [leetcode link](https://leetcode.com/problems/sqrtx/)
- 如果是有大小順序的結果，都可以用 binary search 去實做

# Climbing Stairs
- [leetcode link](https://leetcode.com/problems/climbing-stairs/)
- 經典用 DP 的解法，一次可以爬1或是2步，也就是說 n 台階的方法會是 (n-1) 法走一步 + (n-2) 法走兩步

# Merge Sorted Array
- [leetcode link](https://leetcode.com/problems/merge-sorted-array/)
- 主要是 iter num1 elem 跟 num2 比，如果比較大就換過去，
- 其中再把 num2 做sorting
- iter 結束就把 num2 補到num1 後面

# Binary Tree Inorder Traversal
- [leetcode link](https://leetcode.com/problems/binary-tree-inorder-traversal/)
- 用 linkedlist 來 紀錄 有 left 的 node
- 如果沒有 left，就輸出現在的 node.Val
- 如果 還有 right 就走下去
- 如果連 right 都沒有，就 pop 開始輸出 right value

# Excel Sheet Column Number
- [leetcode link](https://leetcode.com/problems/excel-sheet-column-number/)
- 簡單就是對回去再做 pow 就好

# Happy Number
- [leetcode link](https://leetcode.com/problems/happy-number/)
- 直接按照規則去做

# Power of Three
- [leetcode link](https://leetcode.com/problems/power-of-three/)
- 認為 iter 去除 3 ，如果最後商是1跟餘數0就是 3平方
- 遇到餘數不是0就直接 false

# Missing Number
- [leetcode link](https://leetcode.com/problems/missing-number/)
- 做完 sorting 再依次確認是不是跟index 一樣

# Fizz Buzz
- [leetcode link](https://leetcode.com/problems/fizz-buzz/)
- 做整除就可以

# Symmetric Tree
- [leetcode link](https://leetcode.com/problems/symmetric-tree/)
- recursive, node left vs right
    - 如果都是 null 就是 true
    - 如果一邊 null 就 false
    - 如果val 不同就 false
    - 其他就繼續輸入

# Maximum Depth of Binary Tree
- [leetcode link](https://leetcode.com/problems/maximum-depth-of-binary-tree/)
- 用 stack 做 append, 如果有下一個node 就 append 回去再下一次兔出來看child，直到沒有 child 可看就知道 level

# Convert Sorted Array to Binary Search Tree
- [leetcode link](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/)
- fallow binary search, 用recursive 從中間拆開

# Pascal's Triangle
- [leetcode link](https://leetcode.com/problems/pascals-triangle/)
- 想成一個有i row 跟 j col 的 table 但 j 長度遞增
- 直接照著規格做

# Best Time to Buy and Sell Stock
- [leetcode link](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)
- 只能進出一次，就是找最低點進，然後找後面的最高點出

# Linked List Cycle
- [leetcode link](https://leetcode.com/problems/linked-list-cycle/)
- 如果都把 linked list 存成 hash 就可以檢查，但記憶體耗多
- 可以用fast/slow pointer 最後如果是同一個就確定是 cycle 

# Min Stack
- [leetcode link](https://leetcode.com/problems/min-stack/)
- 基本上可以用 array 做 push 的 container，然後實做那幾個功能就好

# Intersection of Two Linked Lists
- [leetcode link](https://leetcode.com/problems/intersection-of-two-linked-lists/)
- 可以先 next 兩個 list 的長度，然後把比較長的那個先 next 多餘的長度，再同時 next 並確認是不是走到同一個 node，如果是就輸出了

# Majority Element
- [leetcode link](https://leetcode.com/problems/majority-element/)
- 只有兩個數字，其中一個會比較多
- 可以用一個 counter ，指定第一個為 major，當遇到一次就 +1，如果遇到不同的就 -1

# Reverse Linked List
- [leetcode link](https://leetcode.com/problems/reverse-linked-list/)
- 感覺是丟進去 stack 再拿出來把 next 接起來
- 實做方式是用 swap 的方法，就像是箭頭對換
```golang
head, head.Next, p = head.Next, p, head
```

# Palindrome Linked List
- [leetcode link](https://leetcode.com/problems/palindrome-linked-list/)
- 用 fast 跟 slow 來定位 中間 node
- 從 slow 做 reverse
- 再跟從 head 的 val 比較就可以確認

# Delete Node in a Linked List
- [leetcode link](https://leetcode.com/problems/delete-node-in-a-linked-list/)
