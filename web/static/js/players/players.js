import * as vframe from "../framework/vframe.js"

export const AddPlayersToMap = (players) => {
  const map = document.getElementById("map")
  const blockSize = 64
  return players.map((player)=>{
      const playerNode = vframe.vNode("div",
      {
        id:`player${player.id}`,
        class: "players",
        style: `top: ${player.x * blockSize}px; left: ${player.y * blockSize}px; background-position: ${player.currentFrame}px 0`

      },[])
      vframe.mount(playerNode, map)
      return playerNode
  })
}

export const LoadPlayersLives = (players) => {
  const playersInfo =  players.map(player => {
     const lives = Array.from({ length: player.lives }, () =>
      vframe.vNode("div", { class: "life_img"}, [])
    )

    return vframe.vNode("div", {id: `player${player.id}_info`,class: "player_info"},
      [
                vframe.vNode("h4", {}, player.username),
                vframe.vNode("div", {class:"players", id: `player${player.id}_img`},[]),
                vframe.vNode("div", {class:"lives"},lives)
      ])
  })
 return vframe.vNode("div", {id: "players_status"}, playersInfo)
}


export const RemovePlayersFromMap = (players) => {
  players.map((player) => {
    document.getElementById(`player${player.id}`).remove() 
  }) 
}
