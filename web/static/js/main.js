import { ws } from "./websocket/ws.js"
import * as vframe from "./framework/vframe.js"

import { AddUserPrompt, AddLobby, JoinPlayers, lobbyUsers, timer, UpdateTimer } from "./lobby/lobby.js"
import { AddChat } from "./chat/chat.js"
import { AddErrorMsg } from "./error/error.js"
import { AddPlayerToStorage } from "./player/storage.js"
import { GetMessage } from "./chat/chat.js"
import { AddGameContainer, GameOver, WinGame } from "./game/game.js"
import { LoadPlayersLives } from "./players/players.js"
import { AddMap } from "./map/map.js"
import { AddPlayersToMap, RemovePlayersFromMap} from "./players/players.js"
import { handlePlayerMovement, handlePlayerBomb, UpdatePlayerPosition, DamagePlayer } from "./player/player.js"

import { AddBombToMap, ExplodeBomb, AddExplosion } from "./bomb/bomb.js"

import { AddPowerUpsToMap,RemovePowerUp } from "./powerups/powerups.js"


const root = document.getElementById('root')
const userPrompt = AddUserPrompt()
const lobby = AddLobby()
const chat = AddChat()
const gameContainer = AddGameContainer()


let playersLives = vframe.vNode("div", {id: "players_status"}, [])
let currentMap = vframe.vNode("div",{id: "map"},[])

vframe.mount(userPrompt, root)

ws.onmessage = (event) => {
  const data = JSON.parse(event.data)
  switch (data.type){
    case "joinedPlayer":
      vframe.unmount(userPrompt)
      vframe.mount(chat, root)
      vframe.mount(lobby, root)
      const oldLobbyUsers = document.getElementById("lobby_users")
      vframe.mount(lobbyUsers, oldLobbyUsers)
      vframe.mount(timer, root)
      AddPlayerToStorage(JSON.stringify(data.player))
      break
    case "joinedPlayers":
      JoinPlayers(data.players)
      break
    case "timer":
      UpdateTimer(data.timer)
      break
    case "message":
      GetMessage(data.message)
      break
    case "startGame":
      vframe.unmount(timer)
      vframe.unmount(lobbyUsers)
      vframe.unmount(lobby)
      vframe.mount(gameContainer, root)
      const gameContainerBlock = document.getElementById("game_container")
      
      playersLives = LoadPlayersLives(data.players)
      let gameMap = AddMap(data.map.blocks)
      
      currentMap = gameMap

      vframe.mount(playersLives,gameContainerBlock)
      vframe.mount(gameMap,gameContainerBlock)
      AddPlayersToMap(data.players)

      handlePlayerMovement()
      handlePlayerBomb()
      break
    case "updatePlayerPosition":
      UpdatePlayerPosition(data.player)
      break
    case "addedBomb":
      AddBombToMap(data.bomb)
      break
    case "explotion":
      ExplodeBomb(data.bomb)
      let newGameMap = AddMap(data.map.blocks)
      AddExplosion(data.explotion)
      const filteredPlayers = data.players.filter(player => player.lives >= 1)
      

      let newPlayersLives = LoadPlayersLives(filteredPlayers)
      vframe.patch(playersLives,newPlayersLives)
      vframe.patch(currentMap, newGameMap)
      playersLives = newPlayersLives
      currentMap = newGameMap
      
      const removingPlayers = data.players.filter(player  => player.lives < 1)
      RemovePlayersFromMap(removingPlayers)


      AddPowerUpsToMap(data.powerup)
      break

    case "damagedPlayer":
      DamagePlayer(data.player)
      break

    case "removePowerup":
      RemovePowerUp(data.singlepowerup)
      break
    case "gameOver":
      GameOver()
      break
    case "winGame":
      WinGame()
      break
    case "error":
      const errorData = data.error
      console.log(errorData)
      vframe.unmount(userPrompt)
      const errorNode = AddErrorMsg(errorData.error)
      vframe.mount(errorNode,root)
      break
  }
}


