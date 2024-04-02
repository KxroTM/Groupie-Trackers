package functions

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
)

// Struct for the db
type AccountData struct {
	Account []struct {
		Username  string   `json:"Username"`
		Password  string   `json:"Password"`
		Favorites []string `json:"Favorites"`
		Ppf       string   `json:"Ppf"`
		History   []string `json:"History"`
		Playlists Playlist `json:"Playlists"`
	} `json:"account"`
}

// Struct for the user
type Account struct {
	Username  string   `json:"Username"`
	Password  string   `json:"Password"`
	Favorites []string `json:"Favorites"`
	Ppf       string   `json:"Ppf"`
	History   []string `json:"History"`
	Playlists Playlist `json:"Playlists"`
}

// Struct for the playlist
type Playlist struct {
	Playlist []struct {
		Name  string   `json:"Name"`
		Songs []string `json:"Songs"`
	} `json:"Playlist"`
}

// Struct for the remember me
type Remember struct {
	Username string `json:"Username"`
}

// Variable for the db
var dataAccount AccountData

var UserRemember Remember

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

	data, err = os.ReadFile("./database/saveAccount.json")

	if err != nil {
		log.Println("Erreur lors de la lecture du fichier :", err)
	}

	err = json.Unmarshal(data, &UserRemember)
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
func UserBuild(username string) *Account {
	for i := 0; i < len(dataAccount.Account); i++ {
		if dataAccount.Account[i].Username == username {
			user := &Account{
				Username:  dataAccount.Account[i].Username,
				Password:  dataAccount.Account[i].Password,
				Favorites: dataAccount.Account[i].Favorites,
				Ppf:       dataAccount.Account[i].Ppf,
				History:   dataAccount.Account[i].History,
				Playlists: dataAccount.Account[i].Playlists,
			}
			return user
		}
	}
	return nil
}

// Create a new user and update the db
func createUser(username string, password string) {
	hashedPassword := HashPasswordSHA256(password)

	newUser := Account{
		Username: username,
		Password: hashedPassword,
	}

	dataAccount.Account = append(dataAccount.Account, newUser)
	UpdateDB()
}

// Update the db
func UpdateDB() {
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
	return hashedPassword == HashPasswordSHA256(password)
}

// Function for login
func Login(username, password string) bool {
	if findAccount(username) {
		user := UserBuild(username)
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

func ChangePassword(username, oldPassword, newPassword, newPasswordCheck string) bool {
	if newPassword == newPasswordCheck && findAccount(username) {
		user := UserBuild(username)
		if !checkPassword(oldPassword, user.Password) {
			return false
		}
		user.Password = HashPasswordSHA256(newPassword) // Hash the new password
		for i := range dataAccount.Account {
			if dataAccount.Account[i].Username == username {
				dataAccount.Account[i] = *user
				break
			}
		}
		UpdateDB()
		log.Printf("Mot de passe de %s modifié", username)
		return true
	}
	return false
}

// Function for hash a password
func HashPasswordSHA256(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

// Function for remember me
func RememberMe(username string) {

	UserRemember.Username = username

	data, err := json.Marshal(UserRemember)
	if err != nil {
		log.Println("Erreur lors de la conversion JSON :", err)
		return
	}

	err = os.WriteFile("./database/saveAccount.json", data, 0644)
	if err != nil {
		log.Println("Erreur lors de l'écriture dans le fichier :", err)
		return
	}
}

// Function for log out
func LogOut() {
	UserRemember.Username = ""

	data, err := json.Marshal(UserRemember)
	if err != nil {
		log.Println("Erreur lors de la conversion JSON :", err)
		return
	}

	err = os.WriteFile("./database/saveAccount.json", data, 0644)
	if err != nil {
		log.Println("Erreur lors de l'écriture dans le fichier :", err)
		return
	}
}

// Function for add a profile picture
func AddPpf(username, content string) {
	user := UserBuild(username)
	if user != nil {
		user.Ppf = content

		for i := range dataAccount.Account {
			if dataAccount.Account[i].Username == username {
				dataAccount.Account[i] = *user
				break
			}
		}
		UpdateDB()
		log.Printf("%s ajouté aux ppf de %s", content, username)
	} else {
		log.Printf("Erreur %s n'est pas dans la base de donnée", username)
	}
}

// Function for add an artist in the history list
func AddHistory(username, history string) {
	for i := range dataAccount.Account {
		if dataAccount.Account[i].Username == username {
			dataAccount.Account[i].History = append(dataAccount.Account[i].History, history)
			break
		}
	}
	UpdateDB()
}

// Function for create a playlist
func CreatePlaylist(username, playlistName string) {
	for i := range dataAccount.Account {
		if dataAccount.Account[i].Username == username {
			newPlaylist := struct {
				Name  string   `json:"Name"`
				Songs []string `json:"Songs"`
			}{
				Name:  playlistName,
				Songs: nil,
			}
			dataAccount.Account[i].Playlists.Playlist = append(dataAccount.Account[i].Playlists.Playlist, newPlaylist)
			break
		}
	}
	UpdateDB()
}

// Function for delete a playlist
func DeletePlaylist(username, playlistName string) {
	for i := range dataAccount.Account {
		if dataAccount.Account[i].Username == username {
			for j := range dataAccount.Account[i].Playlists.Playlist {
				if dataAccount.Account[i].Playlists.Playlist[j].Name == playlistName {
					dataAccount.Account[i].Playlists.Playlist = append(dataAccount.Account[i].Playlists.Playlist[:j], dataAccount.Account[i].Playlists.Playlist[j+1:]...)
					break
				}
			}
			break
		}
	}
	UpdateDB()
}

// Function for add a song in a playlist
func AddSongToPlaylist(username, playlistName, song string) {
	for i := range dataAccount.Account {
		if dataAccount.Account[i].Username == username {
			for j := range dataAccount.Account[i].Playlists.Playlist {
				if dataAccount.Account[i].Playlists.Playlist[j].Name == playlistName {
					dataAccount.Account[i].Playlists.Playlist[j].Songs = append(dataAccount.Account[i].Playlists.Playlist[j].Songs, song)
					break
				}
			}
			break
		}
	}
	UpdateDB()
}

// Function for delete a song in a playlist
func DeleteSongFromPlaylist(username, playlistName, song string) {
	for i := range dataAccount.Account {
		if dataAccount.Account[i].Username == username {
			for j := range dataAccount.Account[i].Playlists.Playlist {
				if dataAccount.Account[i].Playlists.Playlist[j].Name == playlistName {
					for k := range dataAccount.Account[i].Playlists.Playlist[j].Songs {
						if dataAccount.Account[i].Playlists.Playlist[j].Songs[k] == song {
							dataAccount.Account[i].Playlists.Playlist[j].Songs = append(dataAccount.Account[i].Playlists.Playlist[j].Songs[:k], dataAccount.Account[i].Playlists.Playlist[j].Songs[k+1:]...)
							break
						}
					}
					break
				}
			}
			break
		}
	}
	UpdateDB()
}

// Function for check if a song is in a playlist
func IsInPlaylist(username, playlistName, song string) bool {
	for i := range dataAccount.Account {
		if dataAccount.Account[i].Username == username {
			for j := range dataAccount.Account[i].Playlists.Playlist {
				if dataAccount.Account[i].Playlists.Playlist[j].Name == playlistName {
					for k := range dataAccount.Account[i].Playlists.Playlist[j].Songs {
						if dataAccount.Account[i].Playlists.Playlist[j].Songs[k] == song {
							return true
						}
					}
					break
				}
			}
			break
		}
	}
	return false
}
