function move() {
    pxy[0] += 0.5;
    ctx.beginPath();
    ctx.strokeStyle = 'black';
    ctx.fillStyle = 'yellow';
    ctx.arc(pxy[0]*bsize+ppos, pxy[1]*bsize+ppos, plsize, 0, 2*Math.PI);
    ctx.fill();
    ctx.stroke();
}

document.body.onkeydown = function(e) {
    ws.send(e.keyCode);
    // window.setInterval(move, 2000);
}
