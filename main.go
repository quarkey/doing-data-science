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
	"log"
	"sync"
)

const (
	url          = "http://stat.columbia.edu/~rachel/datasets/"
	fileregex    = `(\d.)`
	downloadpath = "download"
)

var wg sync.WaitGroup

func main() {
	urls := buildURLs("http://stat.columbia.edu/~rachel/datasets/", "nyt%d.csv", 31)
	if err := Download(urls, downloadpath); err != nil {
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
