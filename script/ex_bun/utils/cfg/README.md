# intro
- viper 可以讀複雜的 yaml 或是 json 設定檔，但會使用 mapstructure tag 在 struct
- 如果要同時兼容 linux env 則只能設定 單層 struct，但可以分開 Unmarshul就好
- 如果有 env 預設會讀 env 而不是 file 給的

based on [Load config from file & environment variables in Golang with Viper](https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d)

[A clean way to pass configs in a Go application](https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64)
