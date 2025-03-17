package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

const redditUrl = "https://old.reddit.com/r/"

func main() {
	var subredditUrl string
	fmt.Println("Enter the subreddit url to scrape images from:")
	_, err := fmt.Scan(&subredditUrl)
	if err != nil {
		fmt.Println("Error reading input")
		return
	}
	wg := &sync.WaitGroup{}
	DownloadedImageNumber := 0
	downloadMutex := &sync.RWMutex{}
	scrapeNextUrl(redditUrl+subredditUrl, downloadMutex, &DownloadedImageNumber, wg)
	wg.Wait()
}

func scrapeNextUrl(url string, mutex *sync.RWMutex, downloadPageNumber *int, wg *sync.WaitGroup) {
	wg.Add(1)
	go scrapeAllImg(mutex, downloadPageNumber, url, wg)
	c := colly.NewCollector()
	c.OnError(func(r *colly.Response, err error) {
		println("Error err %s", err.Error())
	})

	c.OnHTML(".next-button a", func(h *colly.HTMLElement) {
		NextUrl := h.Attr("href")
		if !strings.HasPrefix(NextUrl, "http") {
			NextUrl = h.Request.AbsoluteURL(NextUrl)
		}
		fmt.Print("NextUrl:")
		fmt.Println(NextUrl)
		scrapeNextUrl(NextUrl, mutex, downloadPageNumber, wg)
	})

	c.Visit(url)
}

func scrapeAllImg(mux *sync.RWMutex, DownloadedImageNumber *int, imgPageUrl string, wg *sync.WaitGroup) {
	defer wg.Done()
	c := colly.NewCollector()
	c.OnHTML(".thumbnail.invisible-when-pinned.may-blank.outbound", func(h *colly.HTMLElement) {
		imageUrl := h.Attr("href")
		if !strings.HasPrefix(imageUrl, "http") {
			imageUrl = h.Request.AbsoluteURL(imageUrl)
		}
		if strings.Contains(imageUrl, "i.redd") {
			mux.Lock()
			downloadImageNumber := *DownloadedImageNumber
			*DownloadedImageNumber++
			mux.Unlock()
			println("parsed img no: " + strconv.Itoa(downloadImageNumber))
			wg.Add(1)
			go downloadImage(imageUrl, fmt.Sprintf("image%d.jpg", downloadImageNumber), wg)
		}
	})

	c.Visit(imgPageUrl)
}

func downloadImage(imageURL, fileName string, wg *sync.WaitGroup) error {
	defer wg.Done()
	resp, err := http.Get(imageURL)
	if err != nil {
		return fmt.Errorf("failed to fetch the image: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download image, status code: %d", resp.StatusCode)
	}
	file, err := os.Create(fileName)

	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save the image: %v", err)
	}

	fmt.Printf("Image downloaded successfully: %s\n", fileName)
	return nil
}
