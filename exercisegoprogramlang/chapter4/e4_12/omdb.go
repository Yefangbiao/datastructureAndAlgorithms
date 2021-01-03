package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const apiKey = "837a1b8b"
const omdbURL = "http://www.omdbapi.com/"

type Movie struct {
	Title  string
	Year   string
	Poster string
}

func getPost(title string) {
	resp, err := http.Get(omdbURL + "?t=" + url.QueryEscape(title) + "&apikey=" + apiKey)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var m Movie
	if err := json.Unmarshal(data, &m); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("movie title: %s, be on %d", m.Title, m.Year)
	download(m.Poster)
}

func download(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	bcontent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	pos := strings.LastIndex(url, "/")
	if pos == -1 {
		fmt.Println("failed to parse the title of the poster")
		return
	}
	f, err := os.Create(url[pos+1:])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = f.Write(bcontent)
	if err != nil {
		fmt.Println(err)
	}
}

func SearchMovieByTitle(title string) {
	getPost(title)
}

func main() {
	var title string
	fmt.Scanln(&title)
	SearchMovieByTitle(title)
}
