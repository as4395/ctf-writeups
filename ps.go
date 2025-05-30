// sniffer.go
package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    defer conn.Close()

    buf := make([]byte, 4096)
    for {
        n, addr, err := conn.ReadFrom(buf)
        if err != nil {
            fmt.Println("Read error:", err)
            continue
        }
        fmt.Printf("Packet from %s: %x\n", addr, buf[:n])
    }
}
