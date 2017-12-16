/*
The purpose of this program is to download all statistical data from a given URL and process
the data in somewhat useful manner. It is an exercise taken from O'Reilly - Doing Data Science at page 37.
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

const (
	downloadpath = "download"
)

var wg sync.WaitGroup

func main() {
	baseurl := flag.String("baseurl", "", "baseurl")
	sprintf := flag.String("sprintf", "", "spritnf")
	ntimes := flag.Int("n", 0, "n times")
	flag.Parse()

	if *baseurl == "" {
		fmt.Println("Missing baseurl!")
		os.Exit(1)
	}
	if *sprintf == "" {
		fmt.Println("Missing sprintf!")
		os.Exit(1)
	}
	if *ntimes == 0 {
		fmt.Println("Missing ntimes!")
		os.Exit(1)
	}

	urls := buildURLs(*baseurl, *sprintf, *ntimes)
	// urls := buildURLs("http://stat.columbia.edu/~rachel/datasets/", "nyt%d.csv", *ntimes)
	if err := download(urls, downloadpath); err != nil {
		log.Fatal(err)
	}
}

// BuildURLs is creating urls
func buildURLs(url, sprintf string, ntimes int) (urls []string) {
	for i := 1; i < ntimes+1; i++ {
		url := url + fmt.Sprintf(sprintf, i)
		// fmt.Println("builing url:", url)
		urls = append(urls, url)
	}
	return urls
}

// Download file from urls and save output to destitation folder concurrently
func download(urls []string, dest string) error {
	_, err := os.Stat(dest)
	if err != nil {
		log.Fatalf("cannot stat desination folder: %v\n", err)
	}
	for _, url := range urls {
		wg.Add(1)
		go fetch(url, dest)
	}
	wg.Wait()
	return nil
}
func fetch(url, dest string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Printf("downloading file %s\n", url)
	out, err := os.Create(dest + "/" + filepath.Base(url))
	if err != nil {
		fmt.Printf("Unable to create file: %v\n", err)
	}
	defer out.Close()
	io.Copy(out, resp.Body)
	fmt.Printf("%s done!\n", filepath.Base(url))
}
