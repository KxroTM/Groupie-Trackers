package functions

func ArtistbyCreationDate(allArtists AllArtistsData, date int) AllArtistsData {
	var artistList []Artist

	for i := 0; i < len(allArtists); i++ {
		if allArtists[i].CreationDate == int64(date) {
			artistList = append(artistList, allArtists[i])
		}
	}
	return artistList
}

func ArtistbyFirstAlbumDate(allArtists AllArtistsData, date string) AllArtistsData {
	var artistList []Artist
	for i := 0; i < len(allArtists); i++ {
		if allArtists[i].FirstAlbum == date {
			artistList = append(artistList, allArtists[i])
		}
	}
	return artistList
}

func ArtistbyNumberofMember(allArtists AllArtistsData, nomber int) AllArtistsData {
	var artistList []Artist
	for i := 0; i < len(allArtists); i++ {
		if len(allArtists[i].Members) == nomber {
			artistList = append(artistList, allArtists[i])
		}
	}
	return artistList
}
