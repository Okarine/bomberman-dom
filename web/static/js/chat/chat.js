import * as vframe from "../framework/vframe.js"
import { ws } from "../websocket/ws.js"

export const AddChat = () => {

  return vframe.vNode("div", {id:"chat_window"},
    [
      vframe.vNode("div", {id:"chat_box"},
        [
          vframe.vNode("div", {id:"chat_info"}, [vframe.vNode("span", {}, "Game Chat")]),
          vframe.vNode("div", {id:"messages"}, []),
          vframe.vNode("div", {id:"input_section"},
            [
              vframe.vNode("div", {id:"input_tex_section"}, 
                [
                  vframe.vNode("input", {type:"text", id:"message_text", placeholder:"Type message...", 
                    onkeypress: (event) => {
                      if (event.key === 'Enter') {
                        sendMessage()
                      }
                    }
                  }, []),
                ]),
              vframe.vNode("button", {id:"input_chat_btn", onclick:(e)=>sendMessage(e),}, "Send"),
            ]
          )
        ]
      )
    ]
  )
}

export const GetMessage = (data) => { 
  const userData = localStorage.getItem('player')
 
  const messageBlock = document.getElementById("messages")
  if (data.sender === userData.username) {
    const message = CurrentPlayerMessage(data)
    vframe.mountStart(message,messageBlock)
  } else {
    const message = EnemyMessage(data)
    vframe.mountStart(message,messageBlock)
  }
}

const CurrentPlayerMessage = (data) => {

  return vframe.vNode("div", {class:"sent_message"},
    [
      vframe.vNode("div", {class:"message_info"},
        [
          vframe.vNode("div", {class:"sent"}, data.sender)
        ]),
      vframe.vNode("div", {}, data.description)
    ]
  )
}

const EnemyMessage = (data) => {

  return vframe.vNode("div", {class:"received_message"},
    [
      vframe.vNode("div", {class:"message_info"},
        [
          vframe.vNode("div", {class:"received"}, data.sender)
        ]),
      vframe.vNode("div", {}, data.description)
    ]
  )
}

const sendMessage = () => {
  let messageEl = document.getElementById("message_text")
  if (messageEl === undefined) return 

  const message = messageEl.value.trim()

  messageEl.value = ""

  if (message !== ''){
    const data = {
      type: "sendMessage",
      message: message
    }
    ws.send(JSON.stringify(data))
  }


}
