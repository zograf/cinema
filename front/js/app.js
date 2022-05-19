const url = "http://localhost";
const port = "8080";
const endPoint = url + ":" + port

function get() {
    req = new XMLHttpRequest() ;

    req.onreadystatechange = function() {
        if (this.readyState==4 && this.status == 200) {
            console.log(req.responseText)
        } 
    }

    req.open("GET", url + ":" + port);
    req.send();
}

var form = document.getElementById("submit-form");
form.addEventListener("submit", (e) => {
    e.preventDefault();
    login();
});

function login() {

    let username = document.getElementById("username").value;
    let password = document.getElementById("password").value;
    fetch(endPoint + `/login/${username}/${password}`).then(
        response => response.json()
    ).then(
        data => {
            if (Object.keys(data).length === 0) {
                alert("Login failed")
            } else {
                window.location.replace("home.html?id=" + data.id)
            }
        }
    )
}
