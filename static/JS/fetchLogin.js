const onClickSubmit = () => {
    fetch("/tryLogin", {
        method: "POST",
        header: {
            "content-type": "application/json"
        },
        body: JSON.stringify({
            email: document.getElementById("email").value,
            password: document.getElementById("password").value
        })
    })
    .then(response => response.json())
    .then(data => {
        console.log(data);
        if (data) {
            console.log("test")
            window.location.href = "/"
        } else {
            document.getElementById("errorMessage").innerText = "Email ou mdp incorect";
        }
    })
}
