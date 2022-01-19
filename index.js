
const socket = new WebSocket("ws://localhost:8000/ws")

const data = {
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