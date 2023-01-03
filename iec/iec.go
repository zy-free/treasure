package iec

import "fmt"

// ByteCountSI 以1000作为基数
func ByteCountSI(b int64) string {
    const unit = 1000
    if b < unit {
        return fmt.Sprintf("%dB", b)
    }
    div, exp := int64(unit), 0
    for n := b / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.2f%cB",
        float64(b)/float64(div), "KMGTPE"[exp])
}

// ByteCountIEC 以1024作为基数
func ByteCountIEC(b int64) string {
    const unit = 1024
    if b < unit {
        return fmt.Sprintf("%dB", b)
    }
    div, exp := int64(unit), 0
    for n := b / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.2f%cB",
        float64(b)/float64(div), "KMGTPE"[exp])
}
