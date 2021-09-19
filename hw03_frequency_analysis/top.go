package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Words struct {
	Word string
	Freq int
}

func Top10(s string) (returnList []string) {
	var tmpS []string
	//var wordS []Words
	returnCoint := 10

	tmpM := make(map[string]int)
	tmpS = strings.Fields(s)
	for _, value := range tmpS {
		tmpM[value]++
	}
	// do struct to sort
	wordS:=make([]Words,len(tmpM))
	for key, value := range tmpM {
		wordS = append(wordS, Words{
			Word: key,
			Freq: value,
		})
	}

	if len(wordS) < returnCoint {
		returnCoint = len(wordS)
	}

	sort.SliceStable(wordS, func(i, j int) bool {
		if wordS[i].Freq != wordS[j].Freq {
			return wordS[i].Freq > wordS[j].Freq
		}
		return wordS[i].Word < wordS[j].Word
	})

	for i := 0; i < returnCoint; i++ {
		returnList = append(returnList, wordS[i].Word)
	}
	return returnList
}
