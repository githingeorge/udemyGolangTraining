package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}
	words := make(map[string]string)

	sc := bufio.NewScanner(resp.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k, _ := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}

	// bs, _ := ioutil.ReadAll(resp.Body)
	// str := string(bs)
	// fmt.Println(str)
}
