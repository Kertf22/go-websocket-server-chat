import WebSocket from 'ws';

const ws = new WebSocket('ws://localhost:8080');

ws.on("open", (am) => {
    // setInterval(() => {
    //     ws.send('ping' + new Date().getTime());
    // }, 1000);

    ws.on('message', socket => {
        const message = socket.toString()
        console.log('message', message)
    });

    
})


// get user input
const stdin = process.openStdin();
stdin.addListener("data", function(d) {
    const message = d.toString().trim();
    if(message) {
        ws.send(message);
    }
});