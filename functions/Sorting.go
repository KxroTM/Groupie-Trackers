package functions

import "fmt"

type AlbumDate struct {
	day   int
	month int
	year  int
}

func ArtistbyCreationDate(allArtists AllArtistsData, startdate, enddate int) AllArtistsData {
	var artistList AllArtistsData

	for i := 0; i < len(allArtists); i++ {
		if int64(startdate) <= allArtists[i].CreationDate && allArtists[i].CreationDate <= int64(enddate) {
			artistList = append(artistList, allArtists[i])
		}
	}
	return artistList
}

func ArtistbyFirstAlbumDate(allArtists AllArtistsData, startdate, enddate string) AllArtistsData {
	var artistList AllArtistsData

	startParts, err := DateStringToIntSlice(startdate)
	if err != nil {
		fmt.Println("Error:", err)
	}

	endParts, err := DateStringToIntSlice(startdate)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for i := 0; i < len(allArtists); i++ {
		artistParts, err := DateStringToIntSlice(allArtists[i].FirstAlbum)
		if err != nil {
			fmt.Println("Error:", err)
		}
		if artistParts[2] >= startParts[2] && artistParts[1] >= startParts[1] && artistParts[0] >= startParts[0] &&
			artistParts[2] <= endParts[2] && artistParts[1] <= endParts[1] && artistParts[0] <= endParts[0] {
			artistList = append(artistList, allArtists[i])
		}

	}
	return artistList
}

func ArtistbyNumberofMember(allArtists AllArtistsData, listnumber []int) AllArtistsData {
	var artistList AllArtistsData

	for i := 0; i < len(allArtists); i++ {
		if IsNumberinSlice(len(allArtists[i].Members), listnumber) {
			artistList = append(artistList, allArtists[i])
		}
	}
	return artistList
}
