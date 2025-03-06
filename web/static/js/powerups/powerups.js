import * as vframe from "../framework/vframe.js"

export const AddPowerUpsToMap = (powerups) => {
  const map = document.getElementById("map")
  const blockSize = 64
  return powerups.map((powerup)=>{
      const powerupNode = vframe.vNode("div",
      {
        id:`${powerup.type}_powerup`,
        class: "powerups",
        style: `top: ${powerup.x * blockSize}px; left: ${powerup.y * blockSize}px; background-position: ${powerup.currentFrame}px 0`

      },[])
      vframe.mount(powerupNode, map)
      return powerupNode
  })
}

export const RemovePowerUp = (powerup) => {
  document.getElementById(`${powerup.type}_powerup`).remove()
}
