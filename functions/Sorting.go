package functions

func ArtistListbyCreationDate(allArtists AllArtistsData, date int) AllArtistsData {
	var artistList []Artist

	for i := 0; i < len(allArtists); i++ {
		if allArtists[i].CreationDate == int64(date) {
			artistList = append(artistList, allArtists[i])
		}
	}
	return artistList
}
