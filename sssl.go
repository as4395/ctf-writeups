// ssl_checker.go
package main

import (
    "crypto/tls"
    "fmt"
    "time"
)

func main() {
    domain := "google.com:443"
    conn, err := tls.Dial("tcp", domain, nil)
    if err != nil {
        fmt.Println("Failed to connect:", err)
        return
    }
    defer conn.Close()

    cert := conn.ConnectionState().PeerCertificates[0]
    fmt.Println("Issuer:", cert.Issuer)
    fmt.Println("Expires on:", cert.NotAfter)
    daysLeft := time.Until(cert.NotAfter).Hours() / 24
    fmt.Printf("Days left: %.0f\n", daysLeft)
}
