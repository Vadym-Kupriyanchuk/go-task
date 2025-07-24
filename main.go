package main

import (
	"bufio"
	"fmt"
	"net/http"
)

/*
Task:
Input:
	- A text file located on s server (https://siloh.pluto.tv/dev/task/input.txt), each line containing a single URL to a file(binary).
	- Assume no authentication is needed for downloads.
Functionality:
	- Download all files concurrently, (*optimizing for memory and bandwidth usage)
	- Concatenate the contents of the downloaded files in the input order.
	- Save the final output to output.mp4 on local disk.
Constraints:
	- * The number of URLs can be very high (e.g., 100k), so: avoid memory overload.
	- * Be efficient with IO and goroutines – no spawning 100k goroutines at once.
	- Ensure the download and concatenation preserves the order of URLs from the input list.
	- Handle transient download failures with retries and timeouts.
	- Log errors but continue processing the rest of the files.


/*
example of the input file content:
https://siloh.pluto.tv/dev/jupiter/avc3/video/avc3/init.mp4
https://siloh.pluto.tv/dev/jupiter/avc3/video/avc3/seg-1.m4s
https://siloh.pluto.tv/dev/jupiter/avc3/video/avc3/seg-2.m4s
***
*/

func getURLsFromFile(fileURL string) ([]string, error) {
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching URL %s: %v", fileURL, err)
	}

	scanner := bufio.NewScanner(resp.Body)
	var urls []string
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading from URL %s: %v", fileURL, err)
	}

	return urls, nil
}

func main() {
	fileURL := "https://siloh.pluto.tv/dev/task/input.txt"
	urls, err := getURLsFromFile(fileURL)
	if err != nil {
		fmt.Printf("Error fetching URLs: %v\n", err)
		return
	}

	fmt.Println("Fetched URLs:")
	for _, url := range urls {
		fmt.Println(url)
	}

	// Here you would implement the logic to download the files concurrently,
	// concatenate them, and save to output.mp4.
	// This is a placeholder for the actual implementation.
	fmt.Println("Next steps: Implement concurrent downloading and concatenation.")
}
