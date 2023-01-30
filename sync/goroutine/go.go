package goroutine

import (
	"runtime"

	"log"
)

func GoSafe(fn func()) {
	go runSafe(fn)
}

func runSafe(fn func()) {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("safe go: panic recovered: %s\n%s", r, buf)
		}
	}()

	fn()
}
