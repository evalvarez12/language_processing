package processing

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Symbols you want to remove of the text
var symbols = [...]string{",", ".", ";", ":", "'", `"`, "(", ")", `/`, "&", `<br`, `>`, "?", "!", `\`}

func ProcessFile(file string) []string {
	dat, err := ioutil.ReadFile(file)
	if err == nil {
		return Process(dat)
	} else {
		fmt.Printf("%s\r\n", err)
		return []string{}
	}
}

// func Process(data []byte) []string {
// 	del_list := []int{}
// 	add_list := []int{}
//
// 	// Words are separated by a space
// 	words := strings.Split(string(data), " ")
//
// 	for i := range words {
// 		for _, sym := range symbols {
// 			words[i] = strings.Replace(words[i], sym, "", -1)
// 		}
//
// 		words[i] = strings.ToLower(words[i])
//
// 		if words[i] == "" {
// 			del_list = append(del_list, i)
// 		}
//
// 		if strings.Contains(words[i], "-") {
// 			add_list = append(add_list, i)
// 		}
//
// 		if _, err := strconv.Atoi(words[i]); err == nil {
// 			words[i] = "#"
// 		}
//
// 	}
//
// 	for _, i := range del_list {
// 		if i == 0 {
// 			fmt.Println(i)
// 			words = words[1:]
// 		} else if i == len(words) {
// 			words = words[:i-1]
// 		} else {
// 			words = append(words[:i], words[i+1:]...)
// 		}
// 	}
//
// 	for _, i := range add_list {
// 		words_extra := strings.Split(words[i], "-")
// 		if i == 0 {
// 			words = append(words_extra, words[i+1:]...)
// 		} else if i == len(words) {
// 			words = append(words[:i-1], words_extra...)
// 		} else {
// 			words = append(append(words[:i], words_extra...), words[i+1:]...)
// 		}
// 	}
//
// 	return words
// 	// 		fmt.Printf("%s\r\n", string(dat))
//
// }

func Process(data []byte) []string {

	// Words are separated by a space
	words := strings.Split(string(data), " ")

	processed_words := []string{}

	for _, i := range words {
		processed_words = append(processed_words, ProcessWord(i)...)
	}

	return processed_words
}

func ProcessWord(word string) []string {

	if strings.Contains(word, "-") {
		extra_words := strings.Split(word, "-")
		words := []string{}
		for _,i := range extra_words {
			words = append(words,ProcessWord(i)...)
		}
		return words
	} else {
		for _, sym := range symbols {
			word = strings.Replace(word, sym, "", -1)
		}
		
		if strings.Contains(word, "\t") {
			word = ""
		}
		
		if word == "" {
			return []string{}
		} else {

			word = strings.ToLower(word)

			if _, err := strconv.Atoi(word); err == nil {
				word = "#"
			}

			return []string{word}
		}
	}
}
