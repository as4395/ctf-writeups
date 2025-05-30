// arpspoof_detect.go
package main

import (
    "fmt"
    "os/exec"
    "strings"
)

func main() {
    out, err := exec.Command("arp", "-a").Output()
    if err != nil {
        fmt.Println("Error running arp command:", err)
        return
    }

    lines := strings.Split(string(out), "\n")
    macMap := make(map[string]string)

    for _, line := range lines {
        parts := strings.Fields(line)
        if len(parts) >= 3 {
            ip := parts[0]
            mac := parts[1]

            if prevIP, exists := macMap[mac]; exists && prevIP != ip {
                fmt.Printf("Possible ARP spoofing detected: %s also used by %s\n", mac, ip)
            }
            macMap[mac] = ip
        }
    }
}
