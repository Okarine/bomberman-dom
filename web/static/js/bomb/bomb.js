import * as vframe from "../framework/vframe.js"

export const AddBombToMap = (bomb) => {
  const row = document.querySelector(`[data-row="${bomb.x}"]`)
  const col = row.querySelector(`[data-col="${bomb.y}"]`)
  const bombNode = vframe.vNode("div",{ class: "bomb"},[])
  vframe.mount(bombNode, col)
}

export const ExplodeBomb = (bomb) => {
   const row = document.querySelector(`[data-row="${bomb.x}"]`)
  const col = row.querySelector(`[data-col="${bomb.y}"]`)
  while (col.firstChild) {
    col.removeChild(col.firstChild)
  }
}


export const AddExplosion = (explosionRange) => {
  const blockSize = 64
  const map = document.getElementById("map")

  return explosionRange.map((block) => {
    const explosion = document.createElement('div');
    explosion.classList.add('explosion');
    explosion.style.top = `${block.x * blockSize}px`
    explosion.style.left = `${block.y * blockSize}px`
    map.appendChild(explosion);

    explosion.addEventListener('animationend', () => {
      map.removeChild(explosion);
    });

    requestAnimationFrame(() => {
      explosion.style.animationPlayState = 'running';
    });
  })

}
