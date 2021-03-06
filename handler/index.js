
function uuidv4() {
    return ([1e7]+-1e3+-4e3+-8e3+-1e11).replace(/[018]/g, c =>
        (c ^ crypto.getRandomValues(new Uint8Array(1))[0] & 15 >> c / 4).toString(16)
    );
}

function detailDataFormat(detailData) {
    const urlParams = new URLSearchParams(window.location.search);

    if (urlParams.get('source_id')) {
        return {
            "sourceId": urlParams.get('source_id') ? parseInt(urlParams.get('source_id')) : null,
            "utmSource": urlParams.get('utm_source'),
            "utmMedium": urlParams.get('utm_medium'),
            "tags": urlParams.get('tags') ? urlParams.get('tags').split(',') : [],
        }
    }

    return {}
}

if (!sessionStorage.getItem('uuid')){
    sessionStorage.setItem('uuid', uuidv4())
}

const socket = new WebSocket("ws://localhost:8000/ws")
const uuid = sessionStorage.getItem('uuid')
const ecommerceHashInput = document.getElementById('eh')
const detailData = detailDataFormat(new URLSearchParams(window.location.search))
let ecommerceHash = null

if (ecommerceHashInput) {
    ecommerceHash = ecommerceHashInput.value
}

const data = {
    uuid: uuid,
    ecommerceHash: ecommerceHash,
    referrer: document.referrer,
    cookie: document.cookie,
    userAgent: window.navigator.userAgent,
    query: window.location.search,
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
    detailData: detailData,
}

socket.onopen = () => {
    socket.send(JSON.stringify(data))
}