package fanout

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestFanout_Do(t *testing.T) {
	ca := New("cache", Worker(1), Buffer(1024))
	var run bool
	ca.Do(context.Background(), func(c context.Context) {
		run = true
		panic("error")
	})
	time.Sleep(time.Millisecond * 50)
	t.Log("not panic")
	if !run {
		t.Fatal("expect run be true")
	}
}

func TestFanout(t *testing.T) {
	workers := make(map[string]*Fanout)
	s := []string{"test1", "test2", "test1", "test3"}
	for _, v := range s {
		if _, ok := workers[v]; !ok {
			workers[v] = New("v", Worker(1), Buffer(1024))
		}
		workers[v].Do(context.Background(), func(ctx context.Context) {
			fmt.Println("1")
		})
	}

	time.Sleep(time.Millisecond * 50)

}

func TestFanout_Close(t *testing.T) {
	ca := New("cache", Worker(1), Buffer(1024))
	ca.Close()
	err := ca.Do(context.Background(), func(c context.Context) {})
	if err == nil {
		t.Fatal("expect get err")
	}
}
