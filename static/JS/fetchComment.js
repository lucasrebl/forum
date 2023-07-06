const addComment = () => {
    // Récupérer les valeurs des champs du commentaire
    const commentContent = document.getElementById("commentInput").value;

    // Vérifier si le commentaire est vide
    if (commentContent.trim() === "") {
        alert("Veuillez entrer un commentaire.");
        return;
    }

    // Créer un objet de commentaire


    // Effectuer une requête POST pour envoyer le commentaire
    fetch("/addComment", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            content: document.getElementById("commentInput").value
        })
    })
        .then(response => {
            if (!response.ok) {
                throw new Error("Erreur lors de l'envoi du commentaire. Statut : " + response.status);
            }
            return response.json();
        })
        .then(data => {
            // Le commentaire a été ajouté avec succès, mettre à jour la page pour afficher le nouveau commentaire
            // Vous pouvez ajouter ici le code pour mettre à jour la liste de commentaires affichée sur la page
            console.log("Commentaire ajouté avec succès.");
            document.getElementById("commentInput").value = ""; // Réinitialiser le champ de saisie du commentaire
        })
        .catch(error => {
            console.error("Erreur lors de l'envoi du commentaire :", error);
        });
}

// Gérer le clic sur le bouton "Envoyer"
const submitButton = document.getElementById("submitComment");
submitButton.addEventListener("click", addComment);
