function receiveCalcProgression() {
    $.ajax({
        url: 'calcProgress',
        type: 'post',
        success: function (data) {
            incrLoadingBar(data);
        }
    });
}

function incrLoadingBar(data) {
    let bar = document.getElementById("bar");
    let progress = document.getElementById("progress");
    let id = setInterval(frame, 5);
    function frame() {
        if (bar.clientWidth >= progress.clientWidth*parseInt(data)/100) {clearInterval(id)}
        else {
            let width = bar.clientWidth;
            let newWidth = ((width+3)/progress.offsetWidth)*100;
            bar.style.width=newWidth.toString()+"%";
        }
    }
}