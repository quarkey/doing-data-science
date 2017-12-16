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
	"time"
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
		urls = append(urls, url)
	}
	return urls
}

// download download files from urls and save output to destitation concurrently
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
	start := time.Now()
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Printf("[%s] downloading file %s ...\n", start.Format("15:04:05"), url)
	out, err := os.Create(dest + "/" + filepath.Base(url))
	if err != nil {
		fmt.Printf("Unable to create file: %v\n", err)
	}
	defer out.Close()
	io.Copy(out, resp.Body)
	fmt.Printf("[%s] %s downloaded in %s!\n", time.Now().Format("15:04:05"), filepath.Base(url), time.Since(start))
}
