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
