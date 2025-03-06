const URL = "ws://localhost:8080/"

const ws = new WebSocket(URL)
export { ws }

ws.onopen=() => {
  console.log("Connected to websocket! ")
}

ws.onclose=() => {
  console.log("Disconnected from websocket! ")
}
