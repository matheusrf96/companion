
const socket = new WebSocket("ws://localhost:8000/ws")
console.log("Attempting WebSocket Connection...")

socket.onopen = () => {
    console.log("Successfully connected")
    socket.send("Msg from the client")
}

socket.onclose = (event) => {
    console.log("Socket closed connection", event)
}

socket.onmessage = (msg) => {
    console.log(msg)
}

socket.onerror = (err) => {
    console.error("Socket error", err)
}