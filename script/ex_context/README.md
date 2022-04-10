[小白也能看懂的context包详解：从入门到精通](https://segmentfault.com/a/1190000040917752)
[用 10 分鐘了解 Go 語言 context package 使用場景及介紹](https://blog.wu-boy.com/2020/05/understant-golang-context-in-10-minutes/)
[Context, 眾Goroutine手上的電話蟲](https://ithelp.ithome.com.tw/articles/10219405)

[Understanding Context in Golang](https://betterprogramming.pub/understanding-context-in-golang-7f574d9d94e0)

[Context in Go Programming](https://david-yappeter.medium.com/context-in-go-programming-part-1-3a8d470617d0)

- 可以用 WithValue 跟Value 來跟 context 拿值，就很容易達到只有在同一個 thread 拿資料的方法

- http.Request 有 Context() 可以輸出內含的 Context{}，也有 .WithContext(ctx) 把 ctx 存回去，以這種方式，就可以做 middleware 把想存的資料放在 context 裡面以便在 thread 裡面傳遞資料
