package main

import (
    "bufio"
    "fmt"
    "net/http"
    "os"
)

func main() {
    // Open the file containing the URLs
    file, err := os.Open("urls.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create a scanner to read each line of the file
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        url := scanner.Text()

        // Send a GET request to the URL
        resp, err := http.Get(url)
        if err != nil {
            fmt.Printf("Error checking %s: %v\n", url, err)
            continue
        }
        defer resp.Body.Close()

        // Check the status code
        if resp.StatusCode == 200 {
            fmt.Printf("%s is live\n", url)
        } else {
            fmt.Printf("%s returned status code %d\n", url, resp.StatusCode)
        }
    }

    // Check for any scanner errors
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}
