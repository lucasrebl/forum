const onClickSubmit = () => {
    if (inputConform(document.getElementById("name").value, document.getElementById("email").value, document.getElementById("password").value, document.getElementById("confirm-password").value)) {
        console.log(name.value)
        fetch("/tryRegister", {
            method: "POST",
            header: {
                "content-type": "application/json"
            },
            body: JSON.stringify({
                name: document.getElementById("name").value,
                email: document.getElementById("email").value,
                password: document.getElementById("password").value
            })
        })
        .then(response => {
            if (response.ok){
                
                window.location.href = "/";
            }
        })
    } else {
        document.getElementById("errorMessage").innerText = "Un des champs n'est pas réspécté"
    }
}

function inputConform(name, email, password, confirmPassword) {
    var emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    if (!emailRegex.test(email)) {
        return false
    } else if (name.length <= 4) {
        return false
    } else if (confirmPassword != password) {
        return false
    } else {
        return true
    }
}