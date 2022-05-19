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

async function getMovies(){
    return await fetch(endPoint + `/movies`).then(
        response => response.json()
    )
}

async function showMovies(){
    let moviesContainer = document.getElementById('all-movies');
    let movies = await getMovies();
    for(let movie of movies){
        let newCard = document.createElement('div');
        newCard.classList.add('movieCard');
        let movieImage = document.createElement('img');
        movieImage.setAttribute('src', movie.thumbnail);
        let movieName = document.createElement('h2');
        movieName.innerText = movie.name;
        let movieGenre = document.createElement('p');
        movieGenre.innerText = movie.genre;
        let movieMore = document.createElement('button');
        movieMore.innerText = 'See more';
        movieMore.setAttribute("key", movie.name);
        movieMore.addEventListener('click', function() {
            window.location.replace(`movie.html/${this.getAttribute('key')}`);
        })

        newCard.appendChild(movieImage);
        newCard.appendChild(movieName);
        newCard.appendChild(movieGenre);
        newCard.appendChild(movieMore);
        moviesContainer.appendChild(newCard);
    }
}

document.addEventListener("DOMContentLoaded", () => {
    onLoad();
    showMovies();
})
