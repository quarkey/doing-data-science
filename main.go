/*

The purpose of this program is to download all statistical data from a given URL and process
the data in somewhat useful manner. It is an exercise taken from O'Reilly - Doing Data Science at page 37.

Exercice: EDA - Exploratory Data Analysis

There are 31 datasets named nyt1.csv, nyt2.csv and so on, which you can find here:
https://github.com/oreillymedia/doing_data_sience

1. age_groups at <18, 18-24, 25-34, 35-44, 45-54, 56-64 and 65+

2. plot the distributions of number impression and

3. extend your analysis cross days. visualize some metrics and distributions over time.

4. describe and interpet any patterns you find
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

const (
	url          = "http://stat.columbia.edu/~rachel/datasets/"
	fileregex    = `(\d.)`
	downloadpath = "download"
)

var wg sync.WaitGroup

func main() {
	urls := []string{
		"http://stat.columbia.edu/~rachel/datasets/nyt1.csv",
		"http://stat.columbia.edu/~rachel/datasets/nyt2.csv",
		"http://stat.columbia.edu/~rachel/datasets/nyt3.csv",
		"http://stat.columbia.edu/~rachel/datasets/nyt4.csv",
		"http://stat.columbia.edu/~rachel/datasets/nyt5.csv",
		"http://stat.columbia.edu/~rachel/datasets/nyt6.csv",
		"http://stat.columbia.edu/~rachel/datasets/nyt7.csv",
		"http://stat.columbia.edu/~rachel/datasets/nyt8.csv",
	}
	if err := Download(urls, downloadpath); err != nil {
		log.Fatal(err)
	}
}

// Download file from urls and save output to destitation folder concurrently
func Download(urls []string, dest string) error {
	_, err := os.Stat(dest)
	if err != nil {
		log.Fatalf("cannot stat desination folder: %v\n", err)
	}
	for _, url := range urls {
		wg.Add(1)
		go download(url, dest)
	}
	wg.Wait()
	return nil
}
func download(url, dest string) {
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
