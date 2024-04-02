package functions

import (
	"log"
	"strings"
)

// Function to add content to user's favorites
func AddToFavorites(username, content string) {
	user := UserBuild(username)
	if user != nil {
		for _, favorite := range user.Favorites {
			if favorite == content {
				log.Printf("%s est déjà dans les favoris de %s.", content, username)
				return
			}
		}
		user.Favorites = append(user.Favorites, content)
		for i := range dataAccount.Account {
			if dataAccount.Account[i].Username == username {
				dataAccount.Account[i] = *user
				break
			}
		}
		UpdateDB()
		log.Printf("%s ajouté aux favoris de %s", content, username)
	} else {
		log.Printf("Erreur %s n'est pas dans la base de donnée", username)
	}
}

// Function to delete content from user's favorites
func DeleteFavorite(username, content string) {
	user := UserBuild(username)
	if user != nil {
		found := false
		for i, favorite := range user.Favorites {
			if favorite == content {
				user.Favorites = append(user.Favorites[:i], user.Favorites[i+1:]...)
				found = true
				break
			}
		}
		if found {
			for i := range dataAccount.Account {
				if dataAccount.Account[i].Username == username {
					dataAccount.Account[i] = *user
					break
				}
			}
			UpdateDB()
			log.Printf("%s supprimé des favoris de %s", content, username)
		} else {
			log.Printf("%s n'est pas dans les favoris de %s.", content, username)
		}
	} else {
		log.Printf("Erreur %s n'est pas dans la base de donnée", username)
	}
}

// Function to delete all user's favorites
func DeleteAllFavorites(username string) {
	user := UserBuild(username)
	if user != nil {
		user.Favorites = []string{}
		for i := range dataAccount.Account {
			if dataAccount.Account[i].Username == username {
				dataAccount.Account[i] = *user
				break
			}
		}
		UpdateDB()
		log.Printf("Tous les favoris de %s ont été supprimés", username)
	} else {
		log.Printf("Erreur %s n'est pas dans la base de donnée", username)
	}
}

// Function to check if content is in user's favorites
func IsInFavorite(username, content string) bool {
	user := UserBuild(username)
	if user != nil {
		for _, favorite := range user.Favorites {
			if strings.Contains(favorite, content) {
				return true
			}
		}
	}
	return false
}
