const url = "http://localhost";
const port = "8080";
const endPoint = url + ":" + port

var user

function getParamValue(name) {
    let location = decodeURI(window.location.toString());
    let index = location.indexOf('?') + 1;
    let subs = location.substring(index, location.length);
    let splitted = subs.split('&');

    for (let item of splitted) {
        let s = item.split('=');
        let pName = s[0];
        let pValue = s[1];
        if (pName == name)
            return pValue;
    }
}

async function onLoad() {
    let id = getParamValue("id");
    
    await fetch(endPoint + `/user/${id}`).then(
        response => {
            return response.json()
        }
    ).then(
        data => {
            user = data
        }
    )
}

document.addEventListener("DOMContentLoaded", () => {
    onLoad();
})
