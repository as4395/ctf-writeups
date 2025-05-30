// dirbrute.go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    target := "http://example.com/"
    wordlist := []string{"admin", "login", "dashboard", "secret"}

    for _, dir := range wordlist {
        url := target + dir
        res, err := http.Get(url)
        if err != nil {
            fmt.Printf("Request failed for %s: %v\n", url, err)
            continue
        }
        if res.StatusCode == 200 {
            fmt.Printf("[FOUND] %s\n", url)
        }
        res.Body.Close()
    }
}
