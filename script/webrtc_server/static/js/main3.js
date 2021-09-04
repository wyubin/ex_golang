const video = document.querySelector('video');
const constraints = {
  audio: false, // 會有嚴重回授問題，local 測試要關掉
  video: true
};

function handleSuccess(stream) {
  window.stream = stream; // 方便可以在瀏覽器console
  video.srcObject = stream;
}

function handleError(error) {
  console.log('navigator.MediaDevices.getUserMedia error: ', error.message, error.name);
}

function onCapture () {
  navigator.mediaDevices.getUserMedia(constraints)
    .then(handleSuccess)
    .catch(handleError);
}

function stopMedia () {
  if (window.stream) {
    const videoStreams = window.stream.getVideoTracks()

    videoStreams.forEach(stream => {
      stream.stop() // 停止所有media stream
    });

    // 釋放資源
    video.src = video.srcObject = null;
  }
}

onCapture ()

// 截圖
const screenshotButton = document.querySelector('#screenshot-button');
const img = document.querySelector('#screenshot img');
const canvas = document.querySelector('#screenshot canvas');
const to_img = document.querySelector('#to_img');

// 濾鏡
const filterCSS = document.querySelector("#filter");

screenshotButton.onclick = video.onclick = function () {
  canvas.width = video.videoWidth;
  canvas.height = video.videoHeight;
  // 濾鏡
  canvas.className = filterCSS.value;
  // 渲染
  canvas.getContext('2d').drawImage(video, 0, 0, canvas.width, canvas.height);

};
to_img.onclick = function(){
  img.src = canvas.toDataURL('image/png');
}

// 濾鏡切換
filterCSS.onchange = function () {
  video.className = filterCSS.value;
}
