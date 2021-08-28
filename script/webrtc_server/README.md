# intro
用來練習怎麼用前端的 webRTC api 跟後端的 signal server 聯絡實做 webRTC

# ref
[WebRTC之旅：始](https://ithelp.ithome.com.tw/articles/10236473)
[WebRTC (1) - 穿透技術以及 go 的 WebSocket 資訊交換實作](http://www.prochainsci.com/2020/10/webrtc-1-go-websocket.html)
[From zero to fully functional video conference app using Go and webRTC](https://medium.com/@ramezemadaiesec/from-zero-to-fully-functional-video-conference-app-using-go-and-webrtc-7d073c9287da)
[WebRTC in 100 Seconds // Build a Video Chat app from Scratch](https://www.youtube.com/watch?v=WmR9IMUD_CY)

# 策略筆記
## 後端
- 做一個 signaling serveice
- 可以基於 [pion webRTC](https://github.com/pion/webrtc)

## 前端
- 先藉由 signalling server 來做兩台 client 的 protocal 交握，基於這個驗證再讓兩台 client 能夠直接做 media steamming
- 先建立一個簡單的 static server 來做前端的架構
-  getUserMedia 需要在 https, file://, localhost 這幾種狀況才行
- generate certificate
```shell
openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes \
  -keyout key.pem -out cert.pem -subj "/CN=example.com" \
  -addext "subjectAltName=DNS:example.com,DNS:www.example.net,IP:192.168.1.108"
```
### 建立前端網頁
#### basic
- 先基於([link](https://ithelp.ithome.com.tw/articles/10239176)) 建立了一個基本可以開 video 的頁面(index.html)
- [相容性處理](https://ithelp.ithome.com.tw/articles/10239258) 可以先不管，開始有問題再封裝就好
- [adaptor.js](https://github.com/webrtc/adapter) 可以幫忙處理相容性問題
```html
<script src="https://webrtc.github.io/adapter/adapter-latest.js"></script>
```
#### capture/stop
- 加上 capture/stop 功能 [link](https://ithelp.ithome.com.tw/articles/10239260) --- index1.html
  - 預設會 onCapture
  - stopMedia 會斷開 dom 跟 stream 的資源，但隨時可以接上，同時在 mac 上面也會把錄影燈關掉

