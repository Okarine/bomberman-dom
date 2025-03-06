import { ws } from "../websocket/ws.js"

export const handlePlayerMovement = () => {

  let x = 0, y = 0 

  document.addEventListener("keydown", (e) => {
    let moved = false

    switch (e.key) {
      case "ArrowUp":
        x -= 1
        moved = true
        break
      case "ArrowDown":
        x += 1
        moved = true
        break
      case "ArrowLeft":
        y -= 1
        moved = true
        break
      case "ArrowRight":
        y += 1
        moved = true
        break
      
    }
    if (moved){
      const data =  {
        type: "updatePlayerPosition",
        coordinates : {
          x: x,
          y: y,
        }
      }
      ws.send(JSON.stringify(data))
      x = 0
      y = 0
    }
  })
}

export const UpdatePlayerPosition = (player) => {
  const blockSize = 64
  const currentPlayer = document.getElementById(`player${player.id}`) 
  currentPlayer.style.backgroundPosition = `${player.currentFrame}px 0`
  currentPlayer.style.top = `${player.x * blockSize}px`;
  currentPlayer.style.left = `${player.y * blockSize}px`;
}


export const handlePlayerBomb = () => {

  document.addEventListener("keydown", (e) => {

    switch (e.key) {
      case "Control":
      const data =  {
        type: "addedBomb",
      }
      ws.send(JSON.stringify(data))
      break
    }
  })
}


export const DamagePlayer = (player) => {
  const playerElement = document.getElementById(`player${player.id}`)

  playerElement.classList.add("damaged")

  const removeDamagedClass = () => {
    playerElement.classList.remove("damaged")
    playerElement.removeEventListener('animationend', removeDamagedClass)
  }
  playerElement.addEventListener('animationend', removeDamagedClass)
}
