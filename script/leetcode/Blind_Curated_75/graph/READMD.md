# Clone Graph
[link](https://leetcode.com/problems/clone-graph/description/)

## intro
複製一個 graph，但是內容是完全一樣的

## plan
- init
    - 用一個 map node2clone 來存 node 和 clone 的關係
    - visited 用來記錄已經走過的 node
    - 用一個 queue(channel) 來存 node
    - rootNodeClone 是最後return 的 node
- bfs: while queue is not empty
    - currNode := <-queue
    - iter currNode.neighbors
        - 如果沒在 visited 就 clone node 並放入 queue, 新增 node2clone
        - 總是 append node2clone[child] 到 node2clone[currNode].neighbors

- return rootNodeClone
