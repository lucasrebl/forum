fetch("/tokenVerification", {
    method: "GET",
    headers: {
        "content-type": "application/json"
    },
})
.then(response => response.json())
.then(data => {
    if (data.result) {
        console.log(data.message);  // "Cookie UserSession existe, valeur : [valeur du cookie]"

        //Toggle front
        document.getElementById('myBtn2').innerHTML = 'deconnexion'
        document.getElementById('email').style.visibility = 'hidden'
        document.getElementById('password').style.visibility = 'hidden'
        document.getElementById('errorMessage').style.visibility = 'hidden'
        document.getElementById('send2').style.visibility = 'hidden'
        document.getElementById('send3').style.visibility = 'visible'
        document.getElementById('ButtonCreatePost').style.display = 'grid'
        document.getElementById('ButtonCreateAcc').style.display = 'none'

        //

    } else {
        console.log(data.message);  // "Le cookie UserSession n'existe pas"

        //Toggle front
        
        //
    }
})
.catch(err => console.error(err));
