package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	forum "forum/DB"
	"strconv"

	"text/template"

	"net/http"
	"regexp"

	//"strings"
	"time"
)

type PostWithDetailURL struct {
	forum.Post
	DetailURL string
	UserName  string `json:"username"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Comment struct {
	CommentID    int    `json:"commentid"`
	PostID       int    `json:"postid"`
	UserID       int    `json:"userid"`
	CreationDate string `json:"creationdate"`
	Content      string `json:"content"`
	UserName     string `json:"username"`
}

func PostDetails(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID du poste à afficher depuis les paramètres de requête
	postIDQuerry := r.URL.Query().Get("id")

	postID, _ := strconv.Atoi(postIDQuerry)

	if postIDQuerry == "" {
		http.Error(w, "ID du poste manquant", http.StatusBadRequest)
		return
	}

	//Utiliser postID pour récupérer le contenu du poste depuis la base de données
	post, err := forum.GetPostByIDFromDB(postID)
	if err != nil {
		// Gérer l'erreur
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(post)

	//Utiliser postID pour récupérer les commentaires associés depuis la base de données
	comments, err := forum.GetCommentsByPostIDFromDB(postIDQuerry)
	if err != nil {
		// Gérer l'erreur
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Ajouter les commentaires au poste
	post.Comments = comments

	if r.Method == http.MethodPost {
		// Récupérer le contenu du commentaire depuis le formulaire
		newComment := r.FormValue("comment")

		// Ajouter le commentaire à la base de données
		err = forum.AddCommentToPost(postID, newComment)
		if err != nil {
			// Gérer l'erreur
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Effectuer une redirection après la soumission du formulaire
		http.Redirect(w, r, "/poste?id="+postIDQuerry, http.StatusSeeOther)
		return
	}

	// Charger le template
	tmpl, err := template.ParseFiles("pages/poste.html")
	if err != nil {
		// Gérer l'erreur
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var formatedPost forum.FormatedPost

	formatedPost.Username = forum.GetUserInfoByIdDB(post.UserId).Name
	formatedPost.Title = post.Title
	formatedPost.Theme = post.Theme
	formatedPost.Content = post.Content
	formatedPost.Comment = post.Comments

	// Exécuter le template avec les données du poste
	err = tmpl.Execute(w, formatedPost)
	if err != nil {
		// Gérer l'erreur
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	// Récupére posts de la bdd
	posts, err := forum.GetPostsFromDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var filterIsActive bool

	filterIsActive = true

	if filterIsActive {
		posts = forum.Filter(posts, r.FormValue("startDate"), r.FormValue("endDate"))
	}

	// Convertir les posts en une liste de PostWithDetailURL avec le champ DetailURL et UserName
	postsWithDetailURL := make([]PostWithDetailURL, len(posts))
	for i, post := range posts {
		postsWithDetailURL[i] = PostWithDetailURL{
			Post:      post,
			DetailURL: "/post?id=" + strconv.Itoa(post.PostId),
			UserName:  forum.GetUserInfoByIdDB(post.UserId).Name,
		}
	}

	// Charger le template
	tmpl, err := template.ParseFiles("pages/accueil.html")
	if err != nil {
		// Gérer l'erreur
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Exécuter le template avec les données
	err = tmpl.Execute(w, postsWithDetailURL)
	if err != nil {
		// Gérer l'erreur
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "pages/register.html")
}

func TryLogin(w http.ResponseWriter, r *http.Request) {
	var user User
	errJson := json.NewDecoder(r.Body).Decode(&user)
	if errJson != nil {
		http.Error(w, errJson.Error(), http.StatusBadRequest)
	}

	//PrintDebug
	fmt.Println("User try to connect with :")
	fmt.Println("Name: ", user.Name, "| Email: ", user.Email, "| Password: ", user.Password)
	//

	errUserLogin := forum.GetUserLogins(user.Email, user.Password)
	if errUserLogin != nil {
		fmt.Println("Erreur lors de la connexion :", errUserLogin)
		jsonResponse, _ := json.Marshal(false) //failed return
		w.Write(jsonResponse)
		return
	} else {
		fmt.Println("Connexion Succes, creating cookie")

		//généré une clé
		token, errTokenGenerating := GenerateSessionToken()
		if errTokenGenerating != nil {
			fmt.Println("ereur de la génération du token")
		}

		forum.SetUserTokenDB(forum.GetUserInfoByEmailDB(user.Email).Id, token)

		// Créer un cookie de session
		cookieSession := &http.Cookie{
			Name:  "UserSession",
			Value: strconv.Itoa(forum.GetUserInfoByEmailDB(user.Email).Id) + "#" + token,
		}
		// Ecrire le cookie
		http.SetCookie(w, cookieSession)
		jsonResponse, _ := json.Marshal(true) //succes return
		w.Write(jsonResponse)
	}
}

func TokenVerification(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("UserSession")
	var message string
	var result bool
	if err != nil {
		if err == http.ErrNoCookie {
			// Le cookie n'existe pas
			message = "Le cookie UserSession n'existe pas"
			result = false
		} else {
			// Autre erreur
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if forum.UserTokenIsValid(cookie.Value) {

		// Le cookie existe et est valide
		message = "Cookie UserSession existe et est valide, valeur :" + cookie.Value
		result = true
	} else if !forum.UserTokenIsValid(cookie.Value) {
		result = false
		cookieSession := &http.Cookie{
			Name:   "UserSession",
			Value:  "",
			MaxAge: -1,
			Path:   "/",
		}
		http.SetCookie(w, cookieSession)
	}

	jsonResponse, _ := json.Marshal(map[string]interface{}{
		"message": message,
		"result":  result,
	})
	w.Write(jsonResponse)
}

func GenerateSessionToken() (string, error) {
	byteToken := make([]byte, 32)
	_, err := rand.Read(byteToken)
	if err != nil {
		return "", err
	}
	token := base64.StdEncoding.EncodeToString(byteToken)
	return token, nil
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Supprimer le cookie d'utilisateur en définissant une nouvelle expiration dans le passé
	cookie := &http.Cookie{
		Name:   "UserSession",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	// Définir le cookie dans la réponse pour le supprimer
	http.SetCookie(w, cookie)
}

func TryRegister(w http.ResponseWriter, r *http.Request) {

	var user User
	errJson := json.NewDecoder(r.Body).Decode(&user)
	if errJson != nil {
		http.Error(w, errJson.Error(), http.StatusBadRequest)
	}

	//PrintDebug
	fmt.Println("User try to create account with :")
	fmt.Println("Name: ", user.Name, "| Email: ", user.Email, "| Password: ", user.Password)
	//

	//Deuxième passe de vérification d'information
	if inputConform(user.Name, user.Email, user.Password) {
		errAddUser := forum.AddUserDB(user.Name, user.Email, user.Password)
		if errAddUser != nil {
			fmt.Println("Erreur lors de l'ajout de l'utilisateur :", errAddUser)
			jsonResponse, _ := json.Marshal(false) //failed return
			w.Write(jsonResponse)
			return
		} else {
			fmt.Println("Utilisateur ajouté avec succès.")
			jsonResponse, _ := json.Marshal(true) //Réponse pour la redirection
			w.Write(jsonResponse)
		}
	}
}
func TryAddPosts(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("UserSession") //récupération du cookie de session
	if err != nil {
		fmt.Println("TRYADDPOST -> Aucun cookie de session, vous devez vous connecter pour crée un post")
	}

	userId, _ := forum.SplitToken(cookie.Value)

	var newPost forum.Post
	errJson := json.NewDecoder(r.Body).Decode(&newPost)
	if errJson != nil {
		http.Error(w, errJson.Error(), http.StatusBadRequest)
		return
	}

	newPost.CreationDate = time.Now() //Utiliser la date actuelle

	fmt.Println("TRY ADD POST -> user id: ", userId, " | post id (0): ", newPost.PostId, " | content: ", newPost.Content, " | title: ", newPost.Title, " | date: ", newPost.CreationDate)

	errAddPost := forum.AddPostDB(newPost.Title, userId, newPost.Content, newPost.CreationDate, newPost.Theme)
	if errAddPost != nil {
		fmt.Println("Erreur lors de l'ajout du post :", errAddPost)
		return
	} else {
		fmt.Println("Post ajouté avec succès.")
	}

	// Obtenir les posts de la bdd
	postsFromDB, err := forum.GetPostsFromDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("test")
	fmt.Println(postsFromDB)
}
func addComment(w http.ResponseWriter, r *http.Request) {
	// Récupérer les données du commentaire depuis la requête POST
	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		// Erreur lors de la désérialisation JSON
		http.Error(w, "Erreur lors de la lecture des données du commentaire", http.StatusBadRequest)
		return
	}

	// Vérifier si les données du commentaire sont valides
	if comment.PostID <= 0 || comment.UserID <= 0 || comment.Content == "" || comment.CreationDate == "" {
		// Données du commentaire incomplètes ou invalides
		http.Error(w, "Données du commentaire invalides", http.StatusBadRequest)
		return
	}

	// Appeler la fonction AddCommentToPost pour ajouter le commentaire à la base de données
	err = forum.AddCommentToPost(comment.PostID, comment.Content)
	if err != nil {
		// Erreur lors de l'ajout du commentaire à la base de données
		http.Error(w, "Erreur lors de l'ajout du commentaire", http.StatusInternalServerError)
		return
	}

	fmt.Println("user id: ", comment.UserID)
	fmt.Println("post id: ", comment.PostID)
	fmt.Println("content: ", comment.Content)
	fmt.Println("date : ", comment.CreationDate)

	// Convertir comment.PostID en string
	postIDStr := strconv.Itoa(comment.PostID)

	// Obtenir les commentaires associés au poste
	commentsFromDB, err := forum.GetCommentsByPostIDFromDB(postIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Commentaires associés au poste :")
	fmt.Println(commentsFromDB)
}

// Fonction de verification d'information (Name, Email, Password)
func inputConform(name, email, password string) bool {
	emailRegex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

	if !emailRegex.MatchString(email) {
		fmt.Printf("inputConformGO BAD")
		return false
	} else if len(name) <= 4 {
		fmt.Println("inputConformGO BAD")
		return false
	} else {
		fmt.Println("inputConformGO OK")
		return true
	}
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	// Récupérer les données de votre table SQLite
	posts, err := forum.GetPostsFromDB() //apelle de la fonction qui recupere les donnée de la table
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir les données en JSON
	jsonData, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Définir le type de contenu de la réponse HTTP en JSON
	w.Header().Set("Content-Type", "application/json")

	// Envoyer les données JSON en réponse
	w.Write(jsonData)
	fmt.Println(jsonData)

}

func main() {

	forum.InitDB() //initialise la DB pour les tables si elles n'existe pas

	http.HandleFunc("/", HomePage)                           //Home
	http.HandleFunc("/tryLogin", TryLogin)                   //page back login
	http.HandleFunc("/logout", Logout)                       //page back login
	http.HandleFunc("/register", Register)                   //page user register
	http.HandleFunc("/tryRegister", TryRegister)             //page back register
	http.HandleFunc("/tryAddPosts", TryAddPosts)             //page back post
	http.HandleFunc("/tokenVerification", TokenVerification) //page back userTokenSession verification
	http.HandleFunc("/post", PostDetails)                    //page front du poste
	http.HandleFunc("/addComment", addComment)               //page back des commentaires

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("serveur lancer")
	http.ListenAndServe(":8080", nil)
}
