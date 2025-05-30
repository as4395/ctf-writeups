// subfinder.go
package main

import (
    "fmt"
    "net"
)

func main() {
    domain := "example.com"
    wordlist := []string{"www", "mail", "ftp", "dev", "test"}

    for _, sub := range wordlist {
        fqdn := sub + "." + domain
        ips, err := net.LookupHost(fqdn)
        if err == nil {
            fmt.Printf("[FOUND] %s -> %v\n", fqdn, ips)
        }
    }
}
