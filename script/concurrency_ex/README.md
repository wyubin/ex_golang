# intro

[go-concurrency-guide](https://github.com/luk4z7/go-concurrency-guide)

[Concurrent API Patterns in Go](https://marksalpeter.com/concurrent-api-patterns-in-go-52fcb5a9c681)

[errGroup 用法及適用情境](https://blog.kennycoder.io/2021/10/03/Golang-errGroup-%E7%94%A8%E6%B3%95%E5%8F%8A%E9%81%A9%E7%94%A8%E6%83%85%E5%A2%83/)
- 主要用在 concurrency 中，如果有遇到錯誤，也可以主動停止其他 go routine 的方式

# concurrency
希望建立一個 concurrency 的 tools 來跑 multithread 的 job
## spec
- 用 go routine 跑設定的 function
- 輸入一個 []args 來作為 func 的參數
- 要能夠設定 max threads, 如果為0則用 []args 長度來做設定
- 要能夠拿回 function 的 return 跟 error
interface:
- `Run`: 直接載入 Task func 來進行 go routine
- `Wait`: 這裡才做 wait
func:
- `NewWorkerPool`: new 一個新的 workerpool, 如果 threads<=0 則用 `runtime.GOMAXPROCS(0)`

看起來可以用 [semaphore](https://pkg.go.dev/golang.org/x/sync@v0.2.0/semaphore) 實做就好
