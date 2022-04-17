# 目的
- 測試 golang 一些新功能，主要是 go work
- 可以基於 [go doc](https://go.dev/doc/tutorial/workspaces)跟[What are go workspaces and how do I use them?](https://dev.to/gophers/what-are-go-workspaces-and-how-do-i-use-them-1643)
- 目前大部分的用法都是把目前常用的repo clone 一份在 local ，提供自己修改以測試，但先前的用法是在 go.mod 裡面做replace
```text
module test1.18

go 1.18

require golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect

replace golang.org/x/sync => ./local/sync
```
- 但在某些情況下，可能會想要用另外自己維護的，而且可能沒必要跟其他人同步的，就可以放在 go.work

# workflow
- 先開局
```shell
go mod init example.com/ex_workfile
# 會產生需要用的 go.mod
# 建立簡單的 main.gp function，要先 go mod tidy 這樣才會下載相關module 來引用，並加到 go.mod
go mod tidy
go run example.com/ex_workfile
# 跟 go run main.go 結果相同
```
- 開始測試 work，go.work 是建立在 模組外的 workspace
```shell
go work init ./ex_workfile/
# 會建立 go.work 並指定使用 ./ex_workfile
git clone https://go.googlesource.com/example
# clone module 原始碼，並新增 function 
go work use ./example
# 指定使用 example，如此，當在這個workspace 中 import 到 golang.org/x/example，就會以 ./example 做取代
go run example.com/ex_workfile
# 就會 run 出對應結果，但目前go.work 似乎尚未在 vscode 支援，而用處也有限，先不用好了
```
