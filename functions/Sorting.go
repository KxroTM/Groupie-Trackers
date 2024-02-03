package functions

import "fmt"

type AlbumDate struct {
	day   int
	month int
	year  int
}

// Sort Artist by Date

func ArtistbyCreationDateRange(allArtists AllArtists, startdate, enddate int) AllArtists { // Range Version
	var artistList AllArtists

	for i := 0; i < len(allArtists); i++ {
		if int64(startdate) <= allArtists[i].CreationDate && allArtists[i].CreationDate <= int64(enddate) {
			artistList = append(artistList, allArtists[i])
		}
	}
	return artistList
}

func ArtistbyCreationDateCheck(allArtists AllArtists, dates []int) AllArtists { // Check Version
	var artistList AllArtists

	for i := 0; i < len(allArtists); i++ {
		if IsNumberinSlice(int(allArtists[i].CreationDate), dates) {
			artistList = append(artistList, allArtists[i])
		}
	}
	return artistList
}

// Sort Artist by date of First Album

func ArtistbyFirstAlbumDateRange(allArtists AllArtists, startdate, enddate string) AllArtists { // Range Version
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

func ArtistbyFirstAlbumDateCheck(allArtists AllArtists, dates []string) AllArtists { // Check Version
	var artistList AllArtists

	for i := 0; i < len(allArtists); i++ {
		artistParts, err := DateStringToIntSlice(allArtists[i].FirstAlbum)
		if err != nil {
			fmt.Println("Error:", err)
		}
		for j := 0; j < len(dates); j++ {
			datesParts, err := DateStringToIntSlice(dates[j])
			if err != nil {
				fmt.Println("Error:", err)
			}
			if artistParts[2] >= datesParts[2] && artistParts[1] == datesParts[1] && artistParts[0] == datesParts[0] {
				artistList = append(artistList, allArtists[i])
			}
		}

	}
	return artistList
}

func ArtistbyNumberofMemberRange(allArtists AllArtists, startnumber, endnumber int) AllArtists {
	var artistList AllArtists

	for i := 0; i < len(allArtists); i++ {
		numMembers := len(allArtists[i].Members)

		// Check if the number of members is within the specified range
		if numMembers >= startnumber && numMembers <= endnumber {
			artistList = append(artistList, allArtists[i])
		}
	}

	return artistList
}

func ArtistbyNumberofMemberCheck(allArtists AllArtists, listnumber []int) AllArtists {
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
