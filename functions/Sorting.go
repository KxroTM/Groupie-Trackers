package functions

import "fmt"

type AlbumDate struct {
	day   int
	month int
	year  int
}

func ArtistbyCreationDate(allArtists AllArtists, startdate, enddate int) AllArtists {
	var artistList AllArtists

	for i := 0; i < len(allArtists); i++ {
		if int64(startdate) <= allArtists[i].CreationDate && allArtists[i].CreationDate <= int64(enddate) {
			artistList = append(artistList, allArtists[i])
		}
	}
	return artistList
}

func ArtistbyFirstAlbumDate(allArtists AllArtists, startdate, enddate string) AllArtists {
	var artistList AllArtists

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

func ArtistbyNumberofMember(allArtists AllArtists, listnumber []int) AllArtists {
	var artistList AllArtists

	for i := 0; i < len(allArtists); i++ {
		if IsNumberinSlice(len(allArtists[i].Members), listnumber) {
			artistList = append(artistList, allArtists[i])
		}
	}
	return artistList
}

func ArtistbyLoactions(allArtists AllArtists, allLocations AllLocations, locations string) AllArtists {
	var artistList AllArtists

	for i := 0; i < len(allLocations.Index); i++ {
		if IsStringInSlice(locations, allLocations.Index[i].Locations) {
			artistList = append(artistList, allArtists[i])
		}
	}
	return artistList
}
