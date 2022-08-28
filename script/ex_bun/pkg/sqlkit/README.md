# pg_flags
- 基於 cobra 跟 viper 做一個讀取設定的機制
- 基於內部的 default.yaml 進行預設值設定
- 如果有環境變數，則會覆蓋預設值
- 如果 cobra 有輸入 flags, 則以 flags 為主(方便在 pod 中進行手動測試)
- 可以用 `SetEnvKeyReplacer` 去取代 namespace
- yaml 參數可以預設小寫，但 ENV 一定是大寫
- 可以用 `SetConfigName` 配合加上多個資料夾來搜尋設定，或是直接用 `SetConfigFile` 來指定設定檔

## TODO
- 可以做一個 ConfigUnmarshal(p , config_path, name_space) 來做一個可以切換 namespace 的設定 loader
- 再用此 config 作為預設settings 在 cobra 的 flag set 做 binding
