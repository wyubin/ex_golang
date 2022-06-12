# Ref
[Ref](https://earthly.dev/blog/golang-errors/)

- 是目前認為比較好的實做方式，是直接用原生的 errors 去實做
[Go Error Best Practices](https://levelup.gitconnected.com/go-error-best-practice-f0864c5c2385)

[Golang Error Handling — Best Practice in 2020](https://itnext.io/golang-error-handling-best-practice-a36f47b0b94c)

[汎型寫法可以參考一下](https://go.dev/doc/tutorial/generics)

[Building APIs with Go — Part 3 Instrumentation and Error Handling](https://fernando-bandeira.medium.com/building-apis-with-go-part-3-instrumentation-and-error-handling-daba9385e3ec) - 可以參考一下裡面包error 的方式，還蠻符合一般使用的


## 心得
- 主要學到 error 的封裝方式
- msg 不同時，也就當作不同 error object
- 可以用 panic 去啟動 error，也可以用 recover 去將得到的 error 去處理？(應該要用 transaction或是 db commit 去做？)

## consider 20220612
- 以error 可用在哪裡的角度來輸出 type error
  - 像是可用在 http 等輸出的 error 就是 StatusError
- new error 可以輸入一個 error 及 code(http code)來輸出一個具有 code 跟 caller 的 StatusError 封裝
