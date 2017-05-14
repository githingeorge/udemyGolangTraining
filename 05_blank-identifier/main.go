package main

import "net/http"
import "log"
import "io/ioutil"
import "fmt"

func main() {
	resp, err := http.Get("http://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	page, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", page)
}
