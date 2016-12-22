var ip = location.host;
var ws = new WebSocket("ws://"+ip+"/ws");

// Map info
var bsize = 20;
var psize = 2;
var ppsize = psize*2;
var ppos = bsize/2;
var layout = [];
var xy;
var player;
var plsize = 5;

var c = document.getElementById("myCanvas");
var ctx = c.getContext("2d");

// Draw player on canvas
function drawPlayer() {
    ctx.beginPath();

    ctx.fillStyle = 'black';
    ctx.fillRect(1*bsize, 1*bsize, bsize, bsize);
    
    ctx.strokeStyle = 'yellow';
    ctx.fillStyle = 'yellow';
    ctx.arc(player%(xy[0]*bsize), player/(xy[0]*bsize), plsize, 0, 2*Math.PI);
    ctx.fill();
    ctx.stroke();
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
	if (layout[i] == 120) {
	    ctx.fillStyle = 'black';
	    ctx.fillRect(x*bsize, y*bsize, bsize, bsize);
	} else if (layout[i] == 46) {
	    ctx.beginPath();
	    ctx.strokeStyle = 'green';
	    ctx.fillStyle = 'green';
	    ctx.arc(x*bsize+ppos, y*bsize+ppos, psize, 0, 2*Math.PI);
	    ctx.fill();
	} else if (layout[i] == 42) {
	    ctx.beginPath();
	    ctx.strokeStyle = 'green';
	    ctx.fillStyle = 'green';
	    ctx.arc(x*bsize+ppos, y*bsize+ppos, ppsize, 0, 2*Math.PI);
	    ctx.fill();
	}
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

// Get map information
function getMapInfo(data) {
    var str = data.substring(data.lastIndexOf("{")+1, data.lastIndexOf("["));
    xy = str.split(" ");
    str = data.substring(data.lastIndexOf("[")+1, data.lastIndexOf("]"));
    layout = str.split(" ");
}

ws.onmessage = function (event) {
    if (layout.length == 0) {
    	getMapInfo(event.data);
	drawMap();
    } else {
    	player = event.datao;
	drawPlayer();
    }
    // DEBUG
    document.getElementById('log').innerHTML = player;
    // END DEBUG
}

