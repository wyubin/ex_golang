# intro
希望找出在 golang 中寫 commandline 工具的最佳解，目前看起來會基於 cobra 跟 viper 去做

# get
紀錄架構概念
- 主指令就是 binary 本身，通常會附加一些副指令讓 binary 有多種功能
- commandline 的執行架構會放在 cmd/ 下，由 cmd.go 來統整要放哪些副指令，其他 go file 則架構副指令
- 期間會用 viper 將需要export 到其他package 的參數做 binding, 因此藉由 viper package 就可以找到目前
- cobra 的 Flags/PersistentFlags
    - PersistentFlags 會只 return 這一層的參數
- 可以把 add_flag 的行為寫進 pkg 的 config 裏面，只要在 對應 subcmd 去 init 他就好
- 如果不需要 induce 環境變數或是多個 config 檔，看起來基本上是不需要 viper
- 主要是用 init 加上 flags

# example
```shell
# compile
go build -o bin/democtl main.go
# run
go run main.go
# run binary
bin/democtl add --title aa --numbers 1,2,3 --numbers 4
```

# ref
- [link](https://github.com/b-nova-techhub/jamctl/)
- [How do you elegantly build a CLI tool in Go in 15 minutes with Cobra?](https://b-nova.com/en/home/content/how-to-build-an-elegant-cli-tool-with-cobra-in-only-fifteen-minutes)
