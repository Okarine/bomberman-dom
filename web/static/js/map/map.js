import * as vframe from "../framework/vframe.js"
export const AddMap = (blocks) => {
  let rowId = -1
  const map =  blocks.map(row => {
    rowId++
    let colId = -1
    const cols = row.map((col)=>{
      colId++
      return vframe.vNode("div",{ class: col, "data-col": colId},[]) 

    })
    return vframe.vNode("div",{class:"row", "data-row": rowId},cols)
  })
 return vframe.vNode("div", {id: "map"}, map)
}

