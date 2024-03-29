# twoSum
[link](https://leetcode.com/problems/two-sum/description/)
```golang
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
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
```shell
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
```

## plan
### try
- 直接用 for loop 來做，但是要找出最大面積，所以要用兩個 pointer 來記錄，一個是 left pointer，一個是 right pointer, 再用距離來計算面積
- 這個解法有兩個 for loop...

### solution
- idxStart, idxEnd, 但在 sliding 時僅從比較小的值去 sliding, 就可以確保 sliding 方向正確，所以只需要一次 sliding

- 要直接從最大寬度開始確認，直接設定兩邊 pointer，然後往中間移動，確認面積，如果面積比目前最大面積還大，就更新最大面積。
- 另外在往內移動兩邊 pointer 的條件是只移動高度低的，因為只有提高高度低的才可能增加面積，因為 boundary 的最小值才能形成面積，

# 3sum
[link](https://leetcode.com/problems/3sum/)

## intro
輸入一個 []int, 確認有沒有三數相加為 0
```shell
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
```

## plan
### think
暴力法就是把 3sum 當作有參數的 2sum 去處理，但是時間複雜度會比 2sum 要慢很多
但 twosum 是不重複，但 3sum 是可以有重複 item 的，所以無法用 2sum 去解

### think2
先對 num 進行 sort, 就可以 for loop 針對每個數來做 2sum，但因為是可重複的，所以要改成 two pointer 的方式，當 sum > 0，要減少 idxEnd，當 sum < 0，要增加 idxStart

### solution
- 因為複雜度比2sum 高很多，基本上只能兩個 loop
- 可以先將 []int 排序，用 i 做第一個int 的 idx, j/k 分別是第二個/第三個int 的 idx
- i 不動，j 從 i+1 開始，k 從最後一個數開始, 如果 sun > 0，代表要減少，k-- else j++, 同時 j < k

# Search in Rotated Sorted Array
[link](https://leetcode.com/problems/search-in-rotated-sorted-array/)

## intro
輸入由一個從小到大排序的 array with possibly rotated at an unknown pivot index(ex [4,5,6,7,0,1,2])，輸入 target，輸出 target 在 array 中的 index，沒有則輸出 -1

```shell
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

## plan
- 以 low, high 兩個 idx 去做 binary search(不斷用 mid 去確認是不是 target), for idxStart <= idxEnd
- 此數列的型態會至少有一邊是遞增的，特性是 left <= right
    - 如果 low 半邊是遞增的
        - 確認 target 在 low 跟 mid 之間的話 -> high = mid -1 else low = mid + 1
    - 不然 high 半邊一定是遞增的
        - 確認 target 在 mid 跟 high 之間的話 -> low = mid + 1 else high = mid -1 

# Combination Sum
[link](https://leetcode.com/problems/combination-sum/)
```shell
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
```

## intro
會給一個 cadidates []int, 設定一個target int, 找出可能的組合其和為 target，返回所有可能的組合

## plan
看很多解法都使用 recursive，需要用 sub-problem 的想法去思考
主要是因為這算是多重組合，而且是可重複的組合，因此 recursive 可以有效把時間複雜度減低

### solution
- recursive 主程式 checkCombination(arrCandidates []int, currCombination *[]int, remind_target int, idxStart int, result *[][]int)
  - 在內部先檢查目前輸入是不是已達reminder == 0
    - 如果是，就把目前組合加入結果中
    - 如果 reminder < 0 => 要將目前組合 pop() 並 return
  - 其他就 for loop candidates 試著加一層組合到 currCombination 並用 checkCombination 去檢查

# Rotate Image
[link](https://leetcode.com/problems/rotate-image/)

# intro
將一個提供的 2d 矩陣做 往右 rotate
```shell
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
```

# plan
經過觀察及大部分的方法都是
- 先將 2d 矩陣 transpose
  - for loop (上三角跟下三角互換) 做 swap
- 再將 2d 矩陣左右翻轉(column swap)
  - 

# Maximum Subarray
[link](https://leetcode.com/problems/maximum-subarray/)

# intro
輸入一個 []int，輸出最大 subarray 的總和
```shell
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
```

# plan
概念是連續加總, 用單個 for loop就可以
- 需要一個 pointer 去紀錄最大加總, 另一個 pointer 紀錄 subarray 加總
- 當subarray 的加總 < 0 時，代表這個subarray 已經無法讓後續的加總更大，就需要從這邊斷尾，將 sumSubarr 回歸為 0，並往前移動 point

# Spiral Matrix
[link](https://leetcode.com/problems/spiral-matrix/)

# intro
將一個 2d 矩陣轉為 spiral matrix  的 1d array
```shell
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
```

# plan
- 可以單純用走的路徑去 append int, 只是要注意，每次的方向要轉彎時， boundary 就要內縮 1
  - left-right: top--
  - top-bottom: right--
  - right-left: bottom++
  - bottom-top: left++
- 直到 left > right 或 top > bottom 就可以 break

# Jump Game
[link](https://leetcode.com/problems/jump-game/)

## intro
輸入一個 []int，每個數字代表可以往前進的距離，確認是否能夠到達最後一個 index，如果能到達則輸出 true，否則輸出 false
```shell
Input: nums = [2,3,1,1,4]
Output: true
Input: nums = [3,2,1,0,4]
Output: false
```

## plan
有點像是 骨牌遊戲或是加油策略，只是這問題是一直線的，而不是多線
- 用一個 pointer 紀錄目前可以往前的最長步數
- iteration時
  - 將先前maxStep -1 並與目前的 value 比較，將較大的更新到 maxStep
  - 如果 maxStep == 0 且 index 小於 lastIndex，代表無法往前，回傳 false
- 最後回傳 true

# Merge Intervals
[link](https://leetcode.com/problems/merge-intervals/)

## intro
給一個含有 start, end []int 的 array, 輸出最後可合併完成的 [][]int
```shell
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
```


## plan
主要概念是先排序過後，後面的 interval 基本上頂多就跟前一個 overlap, 或是沒 overlap
- 先將 input array 基於其 start 排序
- 用一個 result [][]int 來記錄合併後的結果
- iteration時
  - 如果 currStart > result[-1][1], 就直接 append 到 result
  - 如果 currEnd > result[-1][1], 就 update currEnd 到 result[-1][1]

# Insert Interval
[link](https://leetcode.com/problems/insert-interval/)

## intro
給一個 沒有overlap 的 interval array，輸入一個 interval，輸出插入後的 interval array，並且要維持沒有 overlap 的狀態
```shell
Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
```
## plan
可以 follow merge intervals 的概念，先把 insert interval 丟進去，再 sort 一次，再 iter 確認overlap, 但就是 O(NlogN)

進一步:
在 iter 時就邊確認 currEnd >= insertStart

# Unique Paths
[link](https://leetcode.com/problems/unique-paths/)

## intro
基本上這類有連續效應的計算，通常會藉由 DP 來紀錄已經計算過得結果，而且，通常也都會用 recursive 來做 subjob 的 implement
```shell
Input: m = 3, n = 7
Output: 28
```


## plan
- 需要設定 dp, 是紀錄每一格到終點的可能路線
- 因為每一格往下及往右才是最短路線，所以 dp[i][j] = dp[i+1][j] + dp[i][j+1]
  - 但超過 dimesion 就直接回傳 0, 也就是沒這個方向的路線
  - 如果 dp[i][j] 存在(!=0) 就直接回傳
  - 如果沒有就把 dp[i+1][j] + dp[i][j+1] 存進去並 return

# Climbing Stairs
[leetcode link](https://leetcode.com/problems/climbing-stairs/)

## intro
給一個距離 int, 每次只能走一步或兩步，計算可以剛好走完的步數組合
```shell
Input: n = 3
Output: 3
```

## plan
因為有兩種步伐長度，而又是累積效應，所以一樣可以用dp 去做
- 可以觀察到從總距從 2 to 3, 除了原來直接把2的組合各 +1 之外，因為最後一步可以走 2, 所以也要加上從 1 直接加上兩步的組合
- 於是 dp[i] = dp[i-1] + dp[i-2]

# Set Matrix Zeroes
[leetcode link](https://leetcode.com/problems/set-matrix-zeroes/)

## intro
給一個矩陣，0 的值，對應的整個 column 跟 整個 row 都要是 0
```shell
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]
```

## plan
- 一次 for loop 依序去改看起來無法，應該是一次 iter 標記需要 column & row index 再第二次 for loop 去改

# Word Search
[link](https://leetcode.com/problems/word-search)

# intro
基於一個提供的 char matrix, input string 並從 matrix 中確認是否能找到連續字串，找到就回傳 true，沒有則回傳 false
```shell
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
```

# plan
算是從一個 entry point 後，開始從四周開始找node 的方式，很像 dfs, 一樣也是用 recursive 去實做
- dfs 有幾個參數
  - board: 二維matrix
  - word: 目標字串
  - i: curr row index
  - j: curr col index
  - k: 目前已經找到 word 的第幾個 index
  - 已經看過的座標 paths
- 在 dfs 的結束條件有
  - k == len(word) -> 回傳 true
  - 無法往下走 -> 回傳 false
    - i 超過 board 的 row 範圍
    - j 超過 board col idx 範圍
    - word[k] != board[r][c]
    - i,j 存在 paths 中
  - 如果以上過了代表可以往下走
    - 先將 i,j 記到 paths，然後再往上下左右走
    - 如果可以往下走就 return true
    - 不能到最後的話要 backtrace
- 有了 dfs 之後，iter board 的

# best time to buy and sell stock
[leetcode link](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)

## intro
給一個[]int, 計算最大利潤
```shell
Input: prices = [7,1,5,3,6,4]
Output: 5
```

## plan
- 用 buyIdx/sellIdx 兩個 point 紀錄買進買出的點, init 0
- 用maxProfit 紀錄最大利潤
- iter sellIdx
  - 如果 price[sellIdx] > price[buyIdx], 就更新maxProfix
  - 否則(虧錢)，就把 buyIdx 設為 sellIdx(reset)，再把 sellIdx + 1

# longest consequtive sequences
[leetcode link](https://leetcode.com/problems/longest-consecutive-sequence)

## intro
給一個 []int, 計算最長連續數字
```shell
Input: nums = [100,4,200,1,3,2]
Output: 4
```
## plan
先用 map 紀錄item使用狀況，並用 pre/next 來 extend 每次的 seed 直到不連續
- 先用一個 map[int]bool(num2left) 去紀錄已經用過的數字，先都填 true
- init maxLen
- 再 iter nums(with num2left == false)
  - init currLen, preNum, nextNum = 1, num-1, num+1
  - while num2left[preNum] == false, currLen++, num2left[preNum] = true, preNum--
  - while num2left[nextNum] == false, currLen++, num2left[nextNum] = true, nextNum++
  - if currLen > maxLen, maxLen = currLen

- return maxLen

# Maximum Product Subarray
[leetcode link](https://leetcode.com/problems/maximum-product-subarray)

## intro
給一個[]int, 計算subarray的最大乘積
```shell
Input: nums = [2,3,-2,4]
Output: 6
```

## plan
需要考慮的跟 max sum 不同, 因為都是整數，只是可能有負數，或是遇到 0 就要 reset(prevProd=1)再繼續。 另外連乘有正負號因此會有方向性，分別從正向跟順向去找 maxProduct 應該就可以了

# find minimum in rotated Sorted array
[leetcode link](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array)

## intro
給一個 rotated Sorted []int, 計算最小值
```shell
Input: nums = [3,4,5,1,2]
Output: 1
```

## plan
一樣用 `Search in Rotated Sorted Array` 的概念，用 binary search, 
- 如果 nums[mid] > nums[end], 代表minimun 在右邊，所以 start = mid + 1, 
- 反之, 代表minimun 在左邊，所以 end = mid
無限逼近 while start < end

最後 return nums[start]

# house robber
[leetcode link](https://leetcode.com/problems/house-robber)

## intro
給一個[]int, 是每個家庭所擁有的金錢, 賊不能一晚連續偷兩家，試問一晚最多能偷多少錢
```shell
Input: nums = [1,2,3,1]
Output: 4
```

## plan
- 又是堆疊問題，可以用 dp 暫存偷到第幾家時，拿到到最多錢
- recursive func 是偷到第幾家時，拿到最多錢, 參數有三個, nums, dp, index
 - 當 i<0 代表沒有家了，回傳 0
 - 偷到第 i 個家時，拿到最多錢 dp[i] = max(nums[i] + dp[i-2], dp[i-1]), 也就是說要看你是偷 i-1 會比較多錢還是偷 i-2 + nums[i] 會比較多錢

## ext
延伸，如果 house 是圓形排列，那就代表 1th 跟 last 是不能同時偷的，那計算方式需要改嗎？
不用，只要算出 rob(s[:-1]) 跟 rob(s[1:]) 是哪個比較多

# number of islands
[leetcode link](https://leetcode.com/problems/number-of-islands)

## intro
搜尋方法跟 word search 很像，只是要有一個 map 紀錄已經看過的點(x_y set), for loop 如果是 1 又沒看過就又進到 recursive 去 extend 看過的點 並 island++, 不然就 continue
```shell
Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
```

## plan


# course schedule
[leetcode link](https://leetcode.com/problems/course-schedule)

## intro
給定一個有向圖，如果能完成所有課程，回傳 true，否則回傳 false
```shell
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
```
## plan
### 概念
- 可以用 prer2course 來紀錄需要先修的課程，同時用 indegree 來記錄每個課程需要多少課程先修, 先從沒有 indegree 的課程開始來檢查, ，visited node 都在 indegree 減一, 如果 node indegree 為 0，就把它加入 queue，直到 queue 為空
- 如果可以修完的課程數不等於課程數，代表有課程還沒修完，回傳 false，否則回傳 true


# product of array except self
[leetcode link](https://leetcode.com/problems/product-of-array-except-self)

## intro
給一個 src []int, return 一個 []int為 src[i] 以外的其他數字的 product
```shell
Input: nums = [1,2,3,4]
Output: [24,12,8,6]
```

## plan
最佳化的方法很簡單，先 init prefix 跟 suffix 兩個 []int,
- prifix 從 i=1 累乘到 n-1
- suffix 從 i=n-2 累乘到 0
- return prefix * suffix

# meeting rooms
[leetcode link](https://leetcode.com/problems/meeting-rooms)

## intro
給一個 meeting period [][]int{{0,30}, {5,10}, {15,20}}, 計算是否有重疊的 meeting

## plan
先基於 start 排序, 再 for loop 檢查currEnd > nextStart，如果有就 return false，否則 true

## ext
接下去問，試著計算出所指定的 meeting period 需要最小幾個會議室

### plan
可以用一個 time2room map[int]int{} 來記錄會議時間如果 start 是 +1, 如果 end 是 -1, 下一個 loop 則是分別用 roomCount 計算目前有幾個會議室，maxRoom = max(roomCount, maxRoom)

# missing number
[leetcode link](https://leetcode.com/problems/missing-number)

## intro
給一個 []int, 長度為 n, 包含 0 到 n，回傳缺少的數字
```shell
Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
```

## plan
- 直接 sort 再iter 確認 n[idx] == idx, 如果是就回傳 idx
- 或是用累加公式， 1 + 2 + 3 + ... + n = n * (n + 1) / 2, 再減去 sum(nums)


# alien dictionary
[leetcode link](https://leetcode.com/problems/alien-dictionary)
[ref](https://www.cnblogs.com/grandyang/p/5250200.html)

## intro
給一個 []string, 按照字典序排列，回傳最短字典序的排列
```shell
Input: words = ["wrt","wrf","er","ett","rftt"]
Output: "wertf"
```

## plan
- 用一個 charOrder 26*26 的 [][]bool 來紀錄每個字母前一個字母是否有相連的線，用 iter words 兩兩相比，只要有 char 不同，就可以記錄 charOrder[char1][char2] = true，然後break 看下一對，就可以建立 graph
- 接下來用 dfs(charOrder, runeCurr, visited, charSeq) 來做recursive，用 visited 來記錄已經走過的點
  - 如果 !charOrder[runeCurr][runeCurr] 就直接 return(代表字典沒這個字, 不記順序)
  - 紀錄 visited[rune] = true, 然後找上一個 char(forloop charSeq)
    - 如果 i==runeCurr || !charOrder[i][runeCurr] 就 continue
    - 如果 visited[i] == true 就 false(看過但順序還沒寄？)
    - dfs(charOrder, i, visited, charSeq), 如果 dfs 回傳 true 就 可以寫接下去的順序

# find median from data stream
[leetcode link](https://leetcode.com/problems/find-median-from-data-stream)

## intro
試建立一個 class(MedianFinder), 有 AddNum 可持續增加int 到 []int, 有 FindMedian 可回傳 []int 中位數

## plan
- addNum 就是直接把 int 放進 []int
- findMedian 然後是用 sort.Ints() 來排序，然後回傳 []int 中位數

# longest increasing subsequence
[leetcode link](https://leetcode.com/problems/longest-increasing-subsequence)

## intro
給一個 []int, 回傳最長的 increasing subsequence 的長度
```shell
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
```

## plan
直接用 dp(n) 去紀錄在 idx 的最長 increasing subsequence 的長度，然後用 max(dp) return

# coin change
[leetcode link](https://leetcode.com/problems/coin-change)

## intro
給一個硬幣面值 coins []int, 給一個 amount int, 回傳最少需要多少硬幣可以組成 amount
```shell
Input: coins = [1, 2, 5], amount = 11
Output: 3
```


## plan
用 idxCoinAmountCount (dp) 去紀錄從 idxCoin 使用的話, 總額amount的組成硬幣數量的最小值 => minCoinCount(dp, coins, idxCoin, amount)
- 以下幾種狀況為 stop
  - 如果 amount == 0, return 0
  - 如果 amount < 0, return math.MaxUint32
  - 如果 idxCoin == 0, 則直接確認 amount % coins[0] == 0
    - if true, return amount / coins[0]
    - if false, return math.MaxUint32
  - 如果 dp[idxCoin][amount] != 0, return dp[idxCoin][amount]
- 開始計算 dp 值
  - 如果不使用 idxCoin, nextCoinCount = minCoinCount(dp, coins, idxCoin-1, amount)
  - 如果使用 idxCoin, currCoinCount = 1 + minCoinCount(dp, coins, idxCoin, amount-coins[idxCoin])
  - dp[idxCoin][amount] = min(currCoinCount, nextCoinCount)
- 最後 return dp[idxCoin][amount]

# top k frequent elements
[leetcode link](https://leetcode.com/problems/top-k-frequent-elements)

## intro
給一個 []int with repeat, k int, 回傳前 k 重複多的元素

## plan
不用想得太複雜，其實就是 map -> pair -> sort

# pacific Atlantic water flow
[leetcode link](https://leetcode.com/problems/pacific-atlantic-water-flow)

## intro
給一個 island 的高度matrix [][]int, 回傳所有 pacific ocean(上跟左) 和 atlantic ocean(右跟下) 的 coordinate pair
```shell
Input: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
Output: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
```


## plan
需要init 2個 dp, 一個是可以水流到 pac 的 dp，一個是可以水流到 atl 的 dp，然後 dfs, pac/atl 都是本身的visited map, 另外要有一個 []pair 來紀錄可以同時連 pac 跟 atl 的 coordinate

也是 dfs，只是這次會從海邊往裡面走
- 可以上下左右走，需要限制在
  - 下一格高度需要 >= 上一格高度
  - 不能超過 matrix 的邊界
- stop 條件
  - 如果 visited 就直接 return
- 如果 pac[i][j] && atl[i][j] 就 return

# non-overlap interval
[leetcode link](https://leetcode.com/problems/non-overlapping-intervals)

## intro
給一個 interval array，若有重疊就移除，回傳最少需要移除的 interval 的數量
```shell
Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1 
```

## plan
在 course schedule 那邊，是以 startIdx 進行排序再做計算，但這邊是希望可以移除最少的 interval, 所以以 append 順序來說，以 endIdx 來做排序會是最好的，才能塞入最多的interval
- 先對 endIdx 進行排序
- iter interval(from idx 1)
  - 如果 endIdx[i-1] > startIdx[i], 就將 endIdx[i] = endIdx[i-1], 並且將 minRemoveCount++ (這樣就不會影響到下一個 idx 的比較)
 