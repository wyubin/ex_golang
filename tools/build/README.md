# intro
紀錄 build binary 需要注意的事
- 可以用 `-ldflags="-s -w"` 來減少檔案大小，主要是把 debug 的模組去掉
- `CGO_ENABLED=0` 要加在 multistage 的使用中，以避免可能會有 cgo 的狀況
