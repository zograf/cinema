const url = "http://localhost";
const port = "8080";

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
