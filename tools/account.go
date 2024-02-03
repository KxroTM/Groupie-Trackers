package Groupie_Trackers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type AccountData struct {
	Account []struct {
		Username string `json:"Unsername"`
		Password string `json:"Password"`
	} `json:"account"`
}

type Account struct {
	Username string `json:"Unsername"`
	Password string `json:"Password"`
}

var dataAccount AccountData

// func main() {
// 	//Appeler la fonction login sur la page de connection
// 	login("chems", "mdp")
// 	//Appeler la fonction register si la fonction login returne false
// 	register("test", "mdp", "mauvaismdp")
// }

// Return true if the username is in the DB, false if not
func findAccount(username string) bool {
	data, err := os.ReadFile("db.json")
	if err != nil {
		log.Println("Erreur lors de la lecture du fichier :", err)
		return false
	}

	err = json.Unmarshal(data, &dataAccount)
	if err != nil {
		log.Println("Erreur lors de la conversion JSON :", err)
		return false
	}

	utilisateurTrouve := false

	for _, compte := range dataAccount.Account {
		if compte.Username == username {
			utilisateurTrouve = true
			break
		}
	}

	return utilisateurTrouve
}

// Create the struct with the user's information
func userBuild(username string) *Account {
	for i := 0; i < len(dataAccount.Account); i++ {
		if dataAccount.Account[i].Username == username {
			user := &Account{
				Username: dataAccount.Account[i].Username,
				Password: dataAccount.Account[i].Password,
			}
			return user
		}
	}
	return nil
}

// Create a new user and update the db
func createUser(username string, password string) {
	hashedPassword := hashPasswordSHA256(password)

	newUser := Account{
		Username: username,
		Password: hashedPassword,
	}

	dataAccount.Account = append(dataAccount.Account, newUser)
	updateDB()
}

// Update the db
func updateDB() {
	data, err := json.Marshal(dataAccount)
	if err != nil {
		log.Println("Erreur lors de la conversion JSON :", err)
		return
	}

	err = os.WriteFile("db.json", data, 0644)
	if err != nil {
		log.Println("Erreur lors de l'écriture dans le fichier :", err)
		return
	}
}

// Return true if the password is the good one, false if not
func checkPassword(password, hashedPassword string) bool {
	return hashedPassword == hashPasswordSHA256(password)
}

// Function for retry if the register's conditions are not respected
func try() (username, password, password2 string) {
	var newUsername string
	var newPassword string
	var newPassword2 string
	fmt.Println("Veuillez choisir un pseudo")
	fmt.Scan(&newUsername)
	fmt.Println("Veuillez choisir un mot de passe")
	fmt.Scan(&newPassword)
	fmt.Println("Veuillez confirmer votre mot de passe")
	fmt.Scan(&newPassword2)
	return newUsername, newPassword, newPassword2
}

// Function for retry if the login's conditions are not respected
func retry() (username, password string) {
	var newUsername string
	var newPassword string
	fmt.Println("Pseudo : ")
	fmt.Scan(&newUsername)
	fmt.Println("Mot de passe : ")
	fmt.Scan(&newPassword)
	return newUsername, newPassword
}

// Function for login
func Login(username, password string) bool {
	if findAccount(username) {
		user := userBuild(username)
		if checkPassword(password, user.Password) {
			log.Printf("Connecté en tant que %s", username)
			return true
		} else {
			log.Printf("Mauvais mot de passe !")
			return Login(retry())
		}
	}
	return false
}

// Function for register
func Register(username, password, passwordcheck string) {
	if password == passwordcheck && !findAccount(username) {
		createUser(username, password)
		log.Printf("Utilisateur créé : %s", username)
	} else {
		if findAccount(username) {
			fmt.Println("Compte déjà existant")
			Register(try())
		} else {
			fmt.Println("Veuillez saisir le même mot de passe !")
			Register(try())
		}
	}
}

// Function for hash a password
func hashPasswordSHA256(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
