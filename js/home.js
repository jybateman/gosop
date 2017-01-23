// Map info
var bsize = 20;
var psize = 2;
var ppsize = psize*2;
var ppos = bsize/2;
var layout = [];
var xy;

var player;
var pxy = [];
var plsize = 7
var dir = "39";


var c = document.getElementById("myCanvas");
var ctx = c.getContext("2d");

// Draw player on canvas
function drawPlayer() {
    var x = Math.floor(pxy[0]);
    var y = Math.floor(pxy[1]);

    drawSquare(x+1, y, layout[x+1+y*xy[0]]);
    drawSquare(x-1, y, layout[x-1+y*xy[0]]);
    drawSquare(x, y+1, layout[x+(y+1)*xy[0]]);
    drawSquare(x, y-1, layout[x+(y-1)*xy[0]]);

    drawSquare(x, y, layout[x+y*xy[0]]);
    layout[x+y*xy[0]] = 32;
    
    ctx.beginPath();
    ctx.strokeStyle = 'black';
    ctx.fillStyle = 'yellow';
    ctx.arc(pxy[0]*bsize+ppos, pxy[1]*bsize+ppos, plsize, 0, 2*Math.PI);

    ctx.fill();
    ctx.stroke();
}

// Draw square b at x and y 
function drawSquare(x, y, b) {
    ctx.fillStyle = 'white';
    ctx.fillRect(x*bsize, y*bsize, bsize, bsize);
    switch (parseInt(b)) {
    case 120:
	ctx.fillStyle = 'black';
	ctx.fillRect(x*bsize, y*bsize, bsize, bsize);
	break;
    case 46:
	ctx.beginPath();
	ctx.strokeStyle = 'green';
	ctx.fillStyle = 'green';
	ctx.arc(x*bsize+ppos, y*bsize+ppos, psize, 0, 2*Math.PI);
	ctx.fill();
	break;
    case 42:
	ctx.beginPath();
	ctx.strokeStyle = 'green';
	ctx.fillStyle = 'green';
	ctx.arc(x*bsize+ppos, y*bsize+ppos, ppsize, 0, 2*Math.PI);
	ctx.fill();
    }
}

// Draw map on canvas
function drawMap() {
    ctx.canvas.width  = xy[0] * bsize;
    ctx.canvas.height = xy[1] * bsize;
    ctx.rect(0, 0, xy[0] * bsize, xy[1] * bsize);
    for (var i = 0, y = 0, x = 0, len = layout.length; i < len; i++) {
	// DEBUG
	// ctx.rect(x*bsize, y*bsize, bsize, bsize);
	// ctx.stroke();
	// END DEBUG
	drawSquare(x, y, layout[i])
	if (((i+1)%xy[0]) == 0) {
	    x = 0;
	    y++;
	} else {
	    x++;
	}
    }
    ctx.fill();
    ctx.stroke();
}

// Get map information from data
function getMapInfo(data) {
    var str = data.substring(data.lastIndexOf("{")+1, data.lastIndexOf("["));
    xy = str.split(" ");
    str = data.substring(data.lastIndexOf("[")+1, data.lastIndexOf("]"));
    layout = str.split(" ");
    ctx.canvas.width  = xy[0] * bsize;
    ctx.canvas.height = xy[1] * bsize;
}
