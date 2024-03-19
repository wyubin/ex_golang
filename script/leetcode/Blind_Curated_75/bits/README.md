# Reverse Bits
[link](https://leetcode.com/problems/reverse-bits/)

## intro
[try](https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827)
- bits 的 operator 有一些特別
    - `&` 是 AND, 當有一個是 0 就是 0, 所以 a & 1 == 1 的話，a 的 1 位就是 1
        - x = x & 0xF0 會把 x 的前 4 位變成 0
    - `|` 是 OR, 當有一個是 1 就是 1
        - x = x | 0x0F 會把 x 的後 4 位變成 1(update)
    - `<<` 是左移幾個位數，例如 1 << 3 會變成 1000
    - `>>` 是右移幾個位數，例如 1 >> 3 會變成 0000

## plan
如果熟悉 bits 操作就可以完成了，就是 for loop 把 n 從最右的0/1 更新到res 的從左到右 的位數

# number of 1 bits
[link](https://leetcode.com/problems/number-of-1-bits/)

## intro
給一個 uint (目前限制為 32bits), 計算有多少 1 bits

## plan
目前認為應該熟悉 bits 操作就可以用單個 for loop 完成這個操作
- while n != 0 中
    - 當 `n & 1 == 1` 代表 n 的最右邊是 1，所以 res++
    - n = n >> 1

# counting bits
[link](https://leetcode.com/problems/counting-bits/)

## intro
給一個 n int, 輸出 0 到 n 之間有多少 1 bits

## plan
產生 src []int 之後，用 for loop 計算每個 i 的 1 bits 就可以

# sum of two integers
[link](https://leetcode.com/problems/sum-of-two-integers/)

## intro
不用 +/- 來做兩個 int 的相加

## plan
用 bits 的操作是最直覺的
