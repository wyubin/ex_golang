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

# cycle linked list
[link](https://leetcode.com/problems/linked-list-cycle/)

## intro
輸入一個 linked list，判斷是否有 cycle

## plan
- 用兩個 pointer 來 triverse linked list, 一快一慢，如果沒遇到 nil, 當 fast.node == slow.node 代表有 cycle return true


## reorder list
[link](https://leetcode.com/problems/reorder-list/)

## intro
輸入一個 linked list，順序換成 n1->n5->n2->n4->n3 輸出新的 linked list

## plan
概念是用兩個 pointer 把 list([1,2,3,4,5]) 拆成前半([1,2,3])跟後半([4,5]), 把後半做 reverse, 然後在一個個合併，這邊會用到許多 linked list 的 operation
- split list
 - 直接用 fast/slow 的 point 直到 fast == nil, 那後半會是 slow.next
- reverse
 - 用 pre 跟 curr 做 iter 的 next 交換，中間用 tmpNode 暫存
- merge
 - head1 = head, head2 = pre(reverse list)
 - for head2 != nil 時
    - 暫存 tmpNode = head1.next
    - head1.next = head2
    - head2 = head2.next
    - head1 = tmpNode

# reverse linked list
[link](https://leetcode.com/problems/reverse-linked-list/)

## intro
輸入一個 linked list，回傳反轉後的 linked list

## plan
可以直接follow reorder list 的 reverse 部分實作
