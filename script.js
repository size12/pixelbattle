var canvas = document.getElementById('canvas');
var ctx = canvas.getContext('2d');
ctx.canvas.width  = window.innerWidth;
ctx.canvas.height = window.innerHeight;

let ws = new WebSocket("ws://127.0.0.1:8080/websocket");

function coordinates(x, y) {
    return [Math.trunc(x / 10) * 10, Math.trunc(y / 10) * 10]
}

function draw(x, y) {
    ctx.fillRect(x, y, 10, 10);
}

function clear(x, y) {
    ctx.clearRect(x, y, 10, 10);
}

document.addEventListener("click", (event) => {
    let [x, y] = coordinates(event.offsetX, event.offsetY)
    //console.log("Click on:", x, y)

    if (ws.readyState == 1) {
        ws.send(JSON.stringify({"x" : x, "y" : y}))
    }
})


ws.onmessage = event => {
    //console.log(event.data)
    data = JSON.parse(event.data)
    if (data.event == "fill") {
        draw(data.x, data.y)
    } else {
        clear(data.x, data.y)
    }
}

ws.onerror = event => {
    console.log(event)
}