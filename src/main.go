package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)



func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: main.exe https://boards.4chan.org/X/thread/XXXXX")
		return
	}

	link := os.Args[1]
	id := strings.Split(link, "/")
	dirName := id[len(id)-1]

	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil{
		log.Fatal(err)
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		value, exists := s.Attr("href")
		if exists == true && strings.Contains(value, "is2."){
			url := "https:" + value
			fn := strings.Split(url, "/")
			filename := fn[len(fn)-1]
			err := DownloadImage(url, filename, dirName)
			if err != nil {
				log.Fatal(err)
			}
		}

	})
}

func DownloadImage(url string, filename string, dir string) error {
	fmt.Println("Downloading image: " + filename)
	os.Mkdir(dir, 0755)
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	out, err := os.Create(dir + "/" + filename)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}