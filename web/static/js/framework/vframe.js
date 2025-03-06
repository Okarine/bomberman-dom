// Example of Virtual Node
const vNodeExample = {
    tag: "div",
    attrs: {
        class: 'container'
    },
    children: [
        {
            tag : 'h1',
            attrs: {
                title: 'this is a title'
            },
            children: 'Basics of JS'
        },
        {
            tag : 'p',
            attrs: {
                class: 'description'
            },
            children: 'Here we go!'
        }
    ]
}
// Creates a Virtual Node
export const vNode = (tag, attrs, children) => {
    return {
        tag,
        attrs,
        children
    }
}
// Mounting vNode to the Real DOM
export const mount = (vNode, container) => {
    const el = document.createElement(vNode.tag)
    
    for (const key in vNode.attrs){
        
        // Check if need to add event listener by starting name with "on"
        if (key.startsWith("on")){
            el.addEventListener(key.slice(2).toLowerCase(),vNode.attrs[key])
        }  
        else {
            el.setAttribute(key, vNode.attrs[key])
        }
    }
    if (typeof vNode.children === 'string'){
        el.textContent = vNode.children
    } else {
        vNode.children.forEach(child => {
          mount(child, el)  
        })
    }
    container.appendChild(el)
    //Write real element from DOM into  Virtual Node
    vNode.el = el
}

// Mounting vNode to the start of container in Real DOM
export const mountStart = (vNode, container) => {
    const el = document.createElement(vNode.tag)
    
    for (const key in vNode.attrs){
        
        // Check if need to add event listener by starting name with "on"
        if (key.startsWith("on")){
            el.addEventListener(key.slice(2).toLowerCase(),vNode.attrs[key])
        }  
        else {
            el.setAttribute(key, vNode.attrs[key])
        }
    }
    if (typeof vNode.children === 'string'){
        el.textContent = vNode.children
    } else {
        vNode.children.forEach(child => {
          mount(child, el)  
        })
    }
    container.prepend(el)
    //Write real element from DOM into  Virtual Node
    vNode.el = el
}


// Unmounting a vNode from the Real DOM
export const unmount = (vNode) => {
    // Find parent element and remove child element which written in vNode
    vNode.el.parentNode.removeChild(vNode.el)
}

// Compares 2 nodes and searching for diff
export const patch = (oldNode, newNode) => {
    // if tags not equal then need to do a replacement
    if (oldNode.tag != newNode.tag){
        mount(newNode, oldNode.el.parentNode)
        unmount(oldNode)
    } else {

        // Now newNode will contain real element which was in oldNode
        newNode.el = oldNode.el

        // Resetting all attributes
        while(newNode.el.attributes.length > 0){
            newNode.el.removeAttribute(newNode.el.attributes[0].name)
        }

        // Remove all event listeners
        for (const key in oldNode.attrs){
            if (key.startsWith("on")){
                newNode.el.removeEventListener(key.slice(2).toLowerCase(),oldNode.attrs[key])
            }
        }


        // Add new attributes
        for (const key in newNode.attrs){
            if (key.startsWith("on")){
                newNode.el.addEventListener(key.slice(2).toLowerCase(),newNode.attrs[key])
            } else {
                newNode.el.setAttribute(key, newNode.attrs[key])
            }
        }

        if (typeof newNode.children === 'string'){ // <h1 class="description">text</h1>  <h1 class="info">text</h1>
            newNode.el.textContent = newNode.children
        } else {
            // If old node had a children with string then need to remount with new children 
            if (typeof oldNode.children === 'string'){
                newNode.el.textContent = null
                newNode.children.forEach(child => {
                    mount(child, newNode.el)
                })
            } else {
                const commonLength = Math.min(oldNode.children.length, newNode.children.length)
                
                //Comparing children
                for (let i = 0; i < commonLength; i++){
                    patch(oldNode.children[i], newNode.children[i])
                }
                
                // If oldNode has more children than new then we need to remove others in the end
                if (oldNode.children.length > newNode.children.length){
                    oldNode.children.slice(newNode.children.length).forEach(child => {
                        unmount(child)
                    })
                }
                // If newNode has more children than old then we need to add others in the end
                else if (newNode.children.length > oldNode.children.length) {
                    newNode.children.slice(oldNode.children.length).forEach(child => {
                        mount(child, newNode.el)
                    })
                }
            }

        }
    }
}
const updateEventListeners = (oldNode, newNode) => {
    // Remove old event listeners
    for (const event in oldNode.events) {
        oldNode.el.removeEventListener(event, oldNode.events[event]);
    }

    // Add new event listeners
    for (const event in newNode.events) {
        newNode.el.addEventListener(event, newNode.events[event]);
    }
};

let activeEffect

const useEffect=(fn)=>{
    activeEffect = fn
    fn()
    activeEffect = null
}

class Dependecy{
    constructor(){
        this.subscribers = new Set()
    }
    depend(){
        if(activeEffect){
            this.subscribers.add(activeEffect)
        }
    }
    notify(){
        this.subscribers.forEach((subscriber) => subscriber())
    }
}

const reactive=(obj)=>{
    Object.keys(obj).forEach((key)=>{
        const dep = new Dependecy()
        let value = obj[key]
        Object.defineProperty(obj, key,{
            get(){
                dep.depend()
                return value
            },
            set(newValue){
                value = newValue
                dep.notify()
            }
        })
    })
    return obj
}





const Router = () => {
    const routes = {}
    let currentContainer

    const addRoute = (path, component) => {
        routes[path] = component
    };

    const navigateTo = (path) => {
        window.location.hash = path
        handleRouteChange()
    };

    const handleRouteChange = () => {
        const path = window.location.hash.slice(1)
        const component = routes[path]
        
        if (component && currentContainer) {
            const newVNode = component()
            if (newVNode){
                mount(newVNode, currentContainer)
            }
        }
    }


    window.addEventListener('hashchange', handleRouteChange)

    return {
        addRoute,
        navigateTo,
        
        setCurrentContainer: (container) => {
            currentContainer = container
        }
    }
}


