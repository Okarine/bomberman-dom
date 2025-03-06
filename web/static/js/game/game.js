import * as vframe from "../framework/vframe.js"

export const AddGameContainer = () => {
    return vframe.vNode("div", {id: "game_container"}, [])
}

export const GameOver = () => {
  document.getElementById("map").classList.add("gameover")
}

export const WinGame = () => {
  document.getElementById("map").classList.add("wingame")
}
