package utils

import (
    "log"
    "time"
)

func LogRequest(method, path string) {
    log.Printf("[%s] %s - %s", time.Now().Format(time.RFC3339), method, path)
}
