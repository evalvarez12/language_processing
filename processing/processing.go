package processing

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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

func Process(data []byte) []string {
	del_list := []int{}

	words := strings.Split(string(data), " ")

	for i := range words {
		for _, sym := range symbols {
			words[i] = strings.Replace(words[i], sym, "", -1)
		}

		words[i] = strings.ToLower(words[i])

		if words[i] == "" {
			del_list = append(del_list, i)
		}

		if _, err := strconv.Atoi(words[i]); err == nil {
			words[i] = "#"
		}

	}

	for _, i := range del_list {
		if i == 0 {
			fmt.Println(i)
			words = words[1:]
		} else if i == len(words) {
			words = words[:i-1]
		} else {
			words = append(words[:i], words[i+1:]...)
		}
	}

	return words
	// 		fmt.Printf("%s\r\n", string(dat))

}
