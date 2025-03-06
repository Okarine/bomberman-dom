import * as vframe from "../framework/vframe.js"

export const AddErrorMsg = (error) => {

  return vframe.vNode("div", {id: "error"}, 
    [
      vframe.vNode("h2", {class:"error_msg"}, error)
    ]
  )
}

