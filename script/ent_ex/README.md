# intro
想學一下 fb 的 ent framework，看起來是一個好用的 orm

# resource
[entgo](https://entgo.io/docs/getting-started/)

# 策略
- 先建一個基本的 gin service 有一個 user 的 crud，再依照 ent 實做出功能
- 再進一步使用像是 migration/transaction 等功能

# 安裝
- 在 project 資料夾下要安裝 ent
```shell 
go get entgo.io/ent/cmd/ent

```

# db 設定
## sqlite
sqlite 在設定上面是直接在 DSN動手腳，可以參考 [git:go-sqlite3](https://github.com/mattn/go-sqlite3#connection-string)

# 建立 orm 架構
## 基本架構
- 先 init 一個殼
```shell
go run entgo.io/ent/cmd/ent init User
## 可以同時 init 多個 models
go run entgo.io/ent/cmd/ent init Car Group
```
就會產生ent/schema/user.go 來進行 orm 的建構
填入需要的欄位 
```golang
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
	}
}
```
- 再產生相關的 orm script，就會產生除了 schema 以外的相關 structure script
```shell
go run entgo.io/ent/cmd/ent generate --feature sql/upsert ./ent/schema/
```