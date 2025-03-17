 # Reddit Image Scraper Documentation
 
 ## Overview
 
 This Go program allows you to scrape and download images from a specified Reddit subreddit. It utilizes the Colly library for web scraping and makes the process faster and more efficient by using concurrency with goroutines.
 
 ### Features:
 - **Scrapes Images from Reddit**: Fetches image URLs from posts within a subreddit and downloads them.
 - **Concurrency**: Uses Go's goroutines to download multiple images simultaneously, making the scraping process faster.
 - **Efficient**: The program handles multiple pages within a subreddit, following the "next" button to continue scraping and downloading images.
 - **Customizable**: You can specify any subreddit to scrape images from.
 
 ## Dependencies:
 - **Colly** (`github.com/gocolly/colly`) for web scraping.
 - Go standard libraries:
   - `net/http` for HTTP requests.
   - `io` for copying data to files.
   - `os` for file handling.
   - `strconv`, `strings`, and `sync` for concurrency and data handling.
 
 ## How to Use:
 
 1. **Install Dependencies**: Install Colly by running:
    ```bash
    go get github.com/gocolly/colly
    ```
 
 2. **Run the Program**:
    ```bash
    go run main.go
    ```
 
 3. **Enter Subreddit URL**: The program will prompt you to enter a subreddit URL (e.g., `golang`, `funny`, etc.).
 
 4. **Download Images**: The program will automatically start scraping images from the specified subreddit and save them locally with filenames like `image1.jpg`, `image2.jpg`, etc.
 
 ## Concurrency and Performance:
 - The program uses **goroutines** to download images concurrently, which speeds up the scraping and downloading process.
 - It ensures efficient image downloading by managing concurrency with **sync.WaitGroup** and **sync.RWMutex**.
