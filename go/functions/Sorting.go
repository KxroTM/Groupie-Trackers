package functions

import (
	"fmt"
	"sort"
	"strings"
)

// Sort Artist by Date

func ArtistbyCreationDateRange(allArtists AllArtists, startdate, enddate float64) AllArtists {
	var artistList AllArtists

	for i := 0; i < len(allArtists); i++ {
		if startdate <= allArtists[i].CreationDate && allArtists[i].CreationDate <= enddate {
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

func ArtistbyFirstAlbumDateRange(allArtists AllArtists, startyear, endyear float64) AllArtists { // Range Version
	var artistList AllArtists

	for i := 0; i < len(allArtists); i++ {
		artistYear, err := DateStringToYear(allArtists[i].FirstAlbum)
		if err != nil {
			fmt.Println("Error:", err)
		}
		if artistYear >= startyear && artistYear <= endyear {
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

func ArtistbyNumberofMemberRange(allArtists AllArtists, startnumber, endnumber int) AllArtists { // Range Version
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

func ArtistbyNumberofMemberCheck(allArtists AllArtists, listnumber []int) AllArtists { // Check Version
	var artistList AllArtists

	for i := 0; i < len(allArtists); i++ {
		if IsNumberinSlice(len(allArtists[i].Members), listnumber) {
			artistList = append(artistList, allArtists[i])
		}
	}
	return artistList
}

func ArtistbyLocations(allArtists AllArtists, allLocations AllLocations, locations []string) AllArtists {
	var artistList AllArtists

	for i := 0; i < len(allLocations.Index); i++ {

		for j := 0; j < len(locations); j++ {

			if IsStringInSlice(locations[j], allLocations.Index[i].Locations) {
				artistList = append(artistList, allArtists[i])
			}
		}
	}
	return artistList
}

func ArtistbyCountry(allArtists AllArtists, allLocations AllLocations, country []string) AllArtists {
	var artistList AllArtists

	for i := 0; i < len(allLocations.Index); i++ {

		for j := 0; j < len(country); j++ {

			for k := 0; k < len(allLocations.Index[i].Locations); k++ {

				if strings.Split(allLocations.Index[i].Locations[k], "-")[len(strings.Split(allLocations.Index[i].Locations[k], "-"))-1] == country[j] {
					if i == 0 {
						artistList = append(artistList, allArtists[i])
					} else {
						artistList = append(artistList, allArtists[i-1])
					}
					break
				}
			}
		}

	}
	return artistList
}

// Sort Artist by date of First Album (Ascending)
func SortByFirstAlbumDescending(allArtists AllArtists) AllArtists {
	sort.Slice(allArtists, func(i, j int) bool {
		return parseDate(allArtists[i].FirstAlbum) > parseDate(allArtists[j].FirstAlbum)
	})
	return allArtists
}

// Sort Artist by date of First Album (Ascending)
func SortByFirstAlbumAscending(allArtists AllArtists) AllArtists {
	sort.Slice(allArtists, func(i, j int) bool {
		return parseDate(allArtists[i].FirstAlbum) < parseDate(allArtists[j].FirstAlbum)
	})
	return allArtists
}
