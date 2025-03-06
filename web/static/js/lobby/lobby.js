import * as vframe from "../framework/vframe.js"
import { ws } from "../websocket/ws.js"

export const AddUserPrompt = () => {

  return vframe.vNode("div", {id: "user_welcome"}, 
    [
      vframe.vNode("h2", {class:"welcome_msg"},"Enter your nickname"),
      vframe.vNode("input", {type:"text", id:"name", maxlength:"10", minlength:"1", size:"15", placeholder:"Type here",
        onkeypress: (event) => {
          if (event.key === 'Enter') {
            sendUsername()
          }
        }
      }, []),
      vframe.vNode("button", {id:"welcome_btn", onclick:()=>sendUsername()}, "Enter Game")
    ]
  )
}

export const AddLobby = () => {
  return vframe.vNode("div", {id:"lobby"},
    [
      vframe.vNode("div", {class:"lobby_head"},
        [
          vframe.vNode("h2", {}, "Lobby"),
          vframe.vNode("div", {id:"lobby_users"},[])
        ]
      )
    ]
  )
}
export let lobbyUsers = vframe.vNode("div",{id:"users"},[])

export let timer =  vframe.vNode("div", {id:"timer_start"},[])

export const UpdateTimer = (timeout) => {
  const newTimer = vframe.vNode("div", {id:"timer"}, [
    vframe.vNode("h3", {}, "The game will start in " + timeout + " seconds")
  ])
  vframe.patch(timer, newTimer)
  timer = newTimer
}


const sendUsername = () => {
  let usernameEl = document.getElementById("name")
  if (usernameEl === undefined) return
  const username = usernameEl.value.trim()
  usernameEl.value = ''

  if (username !== ''){
    const data = {
      type: "joinPlayer",
      username: username,
    }
    ws.send(JSON.stringify(data))
  }
}


export const JoinPlayers = (players) => {

  const playersNodes = players.map(player => {
    return vframe.vNode("div", {class:"user_info"},
      [
        vframe.vNode("img", {alt:"player", src:`/static/img/players/player${player.id}_lobby.png` }, []),
        vframe.vNode("div", {}, player.username),
      ]
    )
  })
  const newLobbyUsers = vframe.vNode("div", {id:"users"}, playersNodes)
  vframe.patch(lobbyUsers, newLobbyUsers)
  lobbyUsers = newLobbyUsers
}

