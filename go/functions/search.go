package functions

import (
	"strings"
)

func Search(data AllArtists, input string) []Artist {
	var result []Artist
	var in = false
	for i := 0; i < len(data); i++ {
		if strings.Contains(strings.ToLower(data[i].Name), strings.ToLower(input)) {
			result = append(result, data[i])
		}
		for j := 0; j < len(data[i].Members); j++ {
			if strings.Contains(strings.ToLower(data[i].Members[j]), strings.ToLower(input)) {
				for k := 0; k < len(result); k++ {
					if result[k].ID == data[i].ID {
						in = true
					}
				}
				if !in {
					result = append(result, data[i])
				}
				in = false
			}
		}
	}
	return result
}

func SearchByLocation(input string) []Artist {
	data := LocationsData()
	datas := ArtistData()
	var result []Artist
	var in = false
	for i := 0; i < len(data.Index); i++ {
		for j := 0; j < len(data.Index[i].Locations); j++ {
			if strings.Contains(strings.ToLower(data.Index[i].Locations[j]), strings.ToLower(input)) {
				for k := 0; k < len(result); k++ {
					if result[k].ID == data.Index[i].ID {
						in = true
					}
				}
				if !in {
					result = append(result, datas[i])
				}
				in = false
			}
		}
	}
	return result
}
