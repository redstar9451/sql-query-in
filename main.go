package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func sqlQueryAddQuotes() string {
	keywords := readKeywords()

	if len(keywords) == 0 {
		return "you don't input any keyword"
	}

	keyList := make([]string, len(keywords))
	i := 0
	for keyword := range keywords {
		keyList[i] = "\"" + keyword + "\""
		i++
	}
	res := "in ("
	res += strings.Join(keyList, ", ")
	res += ")"
	return res
}

func readKeywords() map[string]int {
	keywords := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		for _, word := range strings.Fields(line) {
			if strings.TrimSpace(word) == "" {
				continue
			}
			keywords[word] = 1
		}
	}
	return keywords
}

func main() {
	res := sqlQueryAddQuotes()
	fmt.Println(res)
}
