var time = 0;
var nstep = 0;
var step = 6;
var inter = 34;

var frames = 0;
var sframe = 0;

function move() {
    var tmp = Date.now();
    nstep = 6/1000*(tmp - time);
    frames++;
    sframe += (tmp - time);
    if (sframe > 1000) {
	document.getElementById('log').innerHTML = frames;
	frames = 0;
	sframe = 0;
    }
    time = tmp;
    
    switch (parseInt(dir)) {
    case 38:
	var x = pxy[0];
	var y = pxy[1] - nstep;
	break;
    case 39:
	var x = pxy[0] + nstep;
	var y = pxy[1];
	break;
    case 40:
	var x = pxy[0];
	var y = pxy[1] + nstep;
	break;
    case 37:
	var x = pxy[0] - nstep;
	var y = pxy[1];
	break;
    }

    // document.getElementById('log').innerHTML = Math.ceil(x)+", "+Math.ceil(y);

    // if (layout[Math.floor(x)+Math.floor(y)*xy[0]] != 120 && layout[Math.ceil(x)+Math.ceil(y)*xy[0]] != 120) {
	pxy[0] = x;
	pxy[1] = y;
    drawPlayer()
    // }
}
