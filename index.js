
function uuidv4() {
    return ([1e7]+-1e3+-4e3+-8e3+-1e11).replace(/[018]/g, c =>
        (c ^ crypto.getRandomValues(new Uint8Array(1))[0] & 15 >> c / 4).toString(16)
    );
}

if (!sessionStorage.getItem('uuid')){
    sessionStorage.setItem('uuid', uuidv4())
}

const socket = new WebSocket("ws://localhost:8000/ws")
const uuid = sessionStorage.getItem('uuid')

const data = {
    uuid: uuid,
    referrer: document.referrer,
    cookie: document.cookie,
    userAgent: window.navigator.userAgent,
    screen: {
        availHeight: screen.availHeight,
        availWidth: screen.availWidth,
        height: screen.height,
        width: screen.width,
        colorDepth: screen.colorDepth,
        pixelDepth: screen.pixelDepth,
    },
    navigator: {
        hardwareConcurrency: navigator.hardwareConcurrency,
        language: navigator.language,
        languages: navigator.languages,
    },
}

socket.onopen = () => {
    socket.send(JSON.stringify(data))
}