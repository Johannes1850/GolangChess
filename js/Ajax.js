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

function adjustEvalBar(data) {
    console.log(data);
    let bar = document.getElementById("evalBar");
    let progress = document.getElementById("deepEval");
    let width = bar.clientWidth;
    let newWidth = (progress.clientWidth/2) + (progress.clientWidth/2)*parseFloat(data);
    if (newWidth > width) {let id = setInterval(frame1, 10);}
    if (newWidth < width) {let id = setInterval(frame2, 10);}
    function frame1() {
        if (newWidth <= bar.clientWidth) {clearInterval(id)}
        else {
            let width = bar.clientWidth;
            let currentWidth = width+1;
            bar.style.width=currentWidth.toString()+"px";
            let evalWidth = (currentWidth-progress.clientWidth/2)/progress.clientWidth/2;
            bar.innerHTML=precise_round(evalWidth*4, 2).toString();
        }
    }
    function frame2() {
        if (newWidth >= bar.clientWidth) {clearInterval(id)}
        else {
            let width = bar.clientWidth;
            let currentWidth = width-1;
            bar.style.width=currentWidth.toString()+"px";
            let evalWidth = (currentWidth-progress.clientWidth/2)/progress.clientWidth/2;
            bar.innerHTML=precise_round(evalWidth*4, 2).toString();
        }
    }
}

function precise_round(num,decimals) {
    let sign = num >= 0 ? 1 : -1;
    return (Math.round((num*Math.pow(10,decimals)) + (sign*0.001)) / Math.pow(10,decimals)).toFixed(decimals);
}