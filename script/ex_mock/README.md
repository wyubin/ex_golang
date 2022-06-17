[Unit Testing Using Mocking in Go](https://levelup.gitconnected.com/unit-testing-using-mocking-in-go-f281122f499f)

[gomock](https://github.com/golang/mock)

# start
- 通常是基於一個 interface 進行 mock，而這也是在設計之初，應該要針對想 mock 的 function 要有的基本想法，就是要掛在一個 interface上
- 執行以下指令產生 mock 的 module
```shell
mockgen -destination=Mocks/mock_car.go -package=Mocks -source=cars.go
```
會產生以下結構

./Mocks

└── mock_car.go

- 然後就可以開始寫測試了
