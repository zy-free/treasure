package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func Test_go_Panic(t *testing.T) {
	GoSafe(func() {
		panic("123")
	})
}

func Test_go_range(t *testing.T) {
	s := []string{"a", "b", "c", "d"}
	for _, v := range s {
		GoSafe(func() {
			fmt.Println(v)
		})
	}
	for _, v := range s {
		f := func(a string) {
			fmt.Println(v)
		}
		GoSafe(func() {
			f(v)
		})
	}
	time.Sleep(100 * time.Millisecond)
}
