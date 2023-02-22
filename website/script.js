const canvas = document.getElementById('canvas');
const ctx = canvas.getContext('2d');
const colors = ["white", "red", "green", "blue", "black", "yellow"]

ctx.canvas.width  = 500;
ctx.canvas.height = 500;
ctx.canvas.style.border = "2px black solid";


let ws = new WebSocket("ws://127.0.0.1:8080/draw");


function coordinates(x, y) {
    return [Math.trunc(x / 10) * 10, Math.trunc(y / 10) * 10]
}

function draw(x, y, color) {
    ctx.fillStyle = color
    ctx.fillRect(x, y, 10, 10);
}


document.addEventListener("click", (event) => {
    let [x, y] = coordinates(event.offsetX, event.offsetY)
    console.log("Click on:", x, y)

    let c = document.getElementById("color")

    draw(x, y, c.value)

    if (ws.readyState == 1) {
        ws.send(JSON.stringify({"x" : Math.trunc(y / 10), "y" : Math.trunc(x / 10), "color" : c.value}))
    }
})


ws.onmessage = event => {
    //console.log(event.data)
    let data = JSON.parse(event.data)
    // console.log(data)
    for (let i = 0; i < data.length; i++) {
        for (let j = 0; j < data.length; j++) {
            draw(i * 10, j * 10, colors[data[i][j]])
        }
    }
}

ws.onerror = event => {
    console.log(event)
}