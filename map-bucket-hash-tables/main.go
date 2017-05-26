package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	const distribution = 12
	resp := getResponse("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	defer resp.Body.Close()

	scanner := getScanner(resp)
	// fmt.Printf("scanner: %+v \n", scanner)

	buckets := getBuckets(distribution)
	fmt.Println(buckets)

	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, distribution)
		// fmt.Println(n, " - ", word)
		buckets[n][word]++
	}

	for i, val := range buckets {
		fmt.Println(i, " - ", len(val))
	}

	fmt.Println("6th bucket: ", buckets[6])
}

func getResponse(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func getScanner(resp *http.Response) *bufio.Scanner {
	scanner := bufio.NewScanner(resp.Body)

	scanner.Split(bufio.ScanWords)

	return scanner
}

func getBuckets(distribution int) map[int]map[string]int {
	// create slice if slice of string to hold slice of words
	buckets := make(map[int]map[string]int, distribution)
	// create maps to hold words
	for i := 0; i < distribution; i++ {
		buckets[i] = make(map[string]int)
	}

	return buckets

}

func HashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
