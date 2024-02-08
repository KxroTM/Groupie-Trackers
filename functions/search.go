package Groupie_Trackers

import (
	"fmt"
	"strings"
)

func SearchByMember(input string) []Artist {
	data := ArtistData()
	var result []Artist
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i].Members); j++ {
			if strings.Contains(strings.ToLower(data[i].Members[j]), strings.ToLower(input)) {
				result = append(result, data[i])
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

func FinalResearch(input string) {
	result := SearchByMember(input)
	result2 := SearchByName(input)

	if len(result) != 0 {
		fmt.Println(result)
	} else if result2.Name != "" {
		fmt.Println(result2)
	}
}
