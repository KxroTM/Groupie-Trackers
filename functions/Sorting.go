package functions

func ArtistbyCreationDate(allArtists AllArtistsData, startdate, enddate int) AllArtistsData {
	var artistList AllArtistsData

	for i := 0; i < len(allArtists); i++ {
		if int64(startdate) <= allArtists[i].CreationDate && allArtists[i].CreationDate <= int64(enddate) {
			artistList = append(artistList, allArtists[i])
		}
	}
	return artistList
}
