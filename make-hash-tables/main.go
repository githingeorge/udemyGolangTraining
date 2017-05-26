package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	// n := HashBucket("Go", 12)
	// fmt.Println(n)
	// scanning()
	int_slice_plus_plus()
}

func getWords() {
	resp, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bs)
}

func HashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}

func scanning() {
	const input = "It is not the critic who counts; not the man who points out how the strong man stumbles, or where the doer of deeds could have done them better. The credit belongs to the man who is actually in the arena, whose face is marred by dust and sweat and blood; who strives valiantly; who errs, who comes short again and again, because there is no effort without error and shortcoming; but who does actually strive to do the deeds; who knows great enthusiasms, the great devotions; who spends himself in a worthy cause; who at the best knows in the end the triumph of high achievement, and who at the worst, if he fails, at least fails while daring greatly, so that his place shall never be with those cold and timid souls who neither know victory nor defeat. "

	scanner := bufio.NewScanner(strings.NewReader(input))
	// set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// Count the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}
}

func int_slice_plus_plus() {
	buckets := make([]int, 1)
	fmt.Println(buckets[0])
	buckets[0] = 42
	fmt.Println(buckets[0])
	buckets[0]++
	fmt.Println(buckets[0])

}
