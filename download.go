package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

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
