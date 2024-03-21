package functions

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
)

type AccountData struct {
	Account []struct {
		Username  string   `json:"Username"`
		Password  string   `json:"Password"`
		Favorites []string `json:"Favorites"`
	} `json:"account"`
}

type Account struct {
	Username  string   `json:"Username"`
	Password  string   `json:"Password"`
	Favorites []string `json:"Favorites"`
}

var dataAccount AccountData

// func main() {
// 	//Appeler la fonction login sur la page de connection
// 	login("chems", "mdp")
// 	//Appeler la fonction register si la fonction login returne false
// 	register("test", "mdp", "mauvaismdp")
// }

// Function for loading the all db
func LoadDb() {
	data, err := os.ReadFile("./database/account.json")
	if err != nil {
		log.Println("Erreur lors de la lecture du fichier :", err)
	}

	err = json.Unmarshal(data, &dataAccount)
	if err != nil {
		log.Println("Erreur lors de la conversion JSON :", err)
	}
}

// Return true if the username is in the DB, false if not
func findAccount(username string) bool {
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
				Username:  dataAccount.Account[i].Username,
				Password:  dataAccount.Account[i].Password,
				Favorites: dataAccount.Account[i].Favorites,
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

	err = os.WriteFile("./database/account.json", data, 0644)
	if err != nil {
		log.Println("Erreur lors de l'écriture dans le fichier :", err)
		return
	}
}

// Return true if the password is the good one, false if not
func checkPassword(password, hashedPassword string) bool {
	return hashedPassword == hashPasswordSHA256(password)
}

// Function for login
func Login(username, password string) bool {
	if findAccount(username) {
		user := userBuild(username)
		if checkPassword(password, user.Password) {
			log.Printf("Connecté en tant que %s", username)
			return true
		}
	}
	return false
}

// Function for register
func Register(username, password, passwordcheck string) bool {
	if password == passwordcheck && !findAccount(username) && username != "" && password != "" && passwordcheck != "" {
		createUser(username, password)
		log.Printf("Utilisateur créé : %s", username)
		return true
	}
	return false
}

// Function for hash a password
func hashPasswordSHA256(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
