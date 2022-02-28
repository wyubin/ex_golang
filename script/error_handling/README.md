# Ref
[Ref](https://earthly.dev/blog/golang-errors/)

- 是目前認為比較好的實做方式，是直接用原生的 errors 去實做
[Go Error Best Practices](https://levelup.gitconnected.com/go-error-best-practice-f0864c5c2385)

[Golang Error Handling — Best Practice in 2020](https://itnext.io/golang-error-handling-best-practice-a36f47b0b94c)

[汎型寫法可以參考一下](https://go.dev/doc/tutorial/generics)


## 心得
- 主要學到 error 的封裝方式
- msg 不同時，也就當作不同 error object
- 可以用 panic 去啟動 error，也可以用 recover 去將得到的 error 去處理？(應該要用 transaction或是 db commit 去做？)
