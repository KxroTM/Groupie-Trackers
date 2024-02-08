package Groupie_Trackers

import (
	"fmt"
	"strings"
)

func SearchByMember(input string) []Artist {
	data := ArtistData()
	var result []Artist
	in := false
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i].Members); j++ {
			if strings.Contains(strings.ToLower(data[i].Members[j]), strings.ToLower(input)) {
				result = append(result, data[i])
			}
		}
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i].Name); j++ {
			if strings.Contains(strings.ToLower(data[i].Name), strings.ToLower(input)) {
				for k := 0; k < len(data); k++ {
					if data[i].ID == data[k].ID {
						in = true
					}
				}
				if !in {
					result = append(result, data[i])
					break
				}
			}
		}
	}
	return result
}

func SearchByName(input string) Artist {
	data := ArtistData()
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i].Name); j++ {
			if strings.Contains(strings.ToLower(data[i].Name), strings.ToLower(input)) {
				return data[i]
			}
		}
	}
	return Artist{}
}

func SearchByLocation(input string) []Artist {
	data := LocationsData()
	datas := ArtistData()
	var result []Artist
	for i := 0; i < len(data.Index); i++ {
		for j := 0; j < len(data.Index[i].Locations); j++ {
			if strings.Contains(strings.ToLower(data.Index[i].Locations[j]), strings.ToLower(input)) {
				result = append(result, datas[i])
			}
		}
	}
	return result
}

func FinalResearch(input string) {
	result := SearchByMember(input)
	result2 := SearchByName(input)

	if len(result) != 0 {
		fmt.Println(result)
	} else if result2.Name != "" {
		fmt.Println(result2)
	}
}
