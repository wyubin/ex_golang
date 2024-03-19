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


# graph valid tree
[link](https://leetcode.com/problems/clone-graph/description/)

## intro
給一個 graph，輸出是否有滿足條件的 tree(所有 node 都有相互連結, 沒有 cycle)

## plan
藉由 DFS 確認是否有 cycle，沒有就回傳 true，有就回傳 false, 藉由 visitNodes 來確認是否所有 node 都被訪問過

# number of connected components in an undirected graph
[link](https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph)

## intro
給一個有向圖，node 數量為 n，還有 edges，回傳有幾個 connected component

## plan
init edges 並建立 visited map, 用dfs 去triverse graph, 每次 triverse 完如果還有 node 還沒 visited, 就可以 +1 繼續 triverse
- def 的 stop 條件
    - visited[node] == true 就直接return 不用 recursive
    - 然後就針對 childs 做 dfs
