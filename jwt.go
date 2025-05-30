// jwt_decoder.go
package main

import (
    "encoding/base64"
    "fmt"
    "strings"
)

func main() {
    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiYWRtaW4ifQ.signature"

    parts := strings.Split(token, ".")
    if len(parts) < 2 {
        fmt.Println("Invalid JWT token")
        return
    }

    for i, part := range parts[:2] {
        decoded, err := base64.RawURLEncoding.DecodeString(part)
        if err != nil {
            fmt.Printf("Failed to decode part %d: %v\n", i, err)
            continue
        }
        fmt.Printf("Part %d: %s\n", i, decoded)
    }
}
