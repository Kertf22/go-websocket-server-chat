
let ws;
console.log(ws)

connect("AnyUsername");

const connect = (username) => {
    ws = new WebSocket('ws://localhost:8080');

    ws.on("open", (am) => {
        // setInterval(() => {
        //     ws.send('ping' + new Date().getTime());
        // }, 1000);
        

        ws.send('username', username)
        ws.on('message', socket => {
            const message = socket.toString()
            console.log('message', message)
        });
    })
};

const readMessage = (message) => {

};

const SendMessage = (message) => {
    if (!ws || ws.readyState !== WebSocket.OPEN || !message) {
        return;
    }

    ws.send(message);
}

