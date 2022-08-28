# intro
- viper 可以讀複雜的 yaml 或是 json 設定檔，但會使用 mapstructure tag 在 struct
- 如果要同時兼容 linux env 則只能設定 單層 struct，但可以分開 Unmarshul就好
- 如果有 env 預設會讀 env 而不是 file 給的

# flags
cobra 的 Flags 則是在 command 加上參數設定的模組
可以先經由 cobra.command 把改設定的參數設定完之後, 再經由 viper 的 BindPFlags 吃進 viper, 然後再做 Unmarshal

在設定時，不管 env/config/flags 都先傳進 viper, 可以在某個時間點再一起 umarshal 到 struct 中

based on [Load config from file & environment variables in Golang with Viper](https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d)

[A clean way to pass configs in a Go application](https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64)
