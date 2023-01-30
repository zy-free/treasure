package ringbuf

import (
	"testing"
)

func TestRing(t *testing.T) {
	var(
		err error
		num = uint64(5)
	)

	ring := NewRing(num)

	for i := 0; i < 8; i++ {
		buf, err := ring.Set()
		if err != nil {
			t.Fatal("Set fail")
		}
		*buf = i
		ring.SetAdv()
	}

	_, err = ring.Set()
	if err == nil {
		t.Fatal("Set full fail")
	}

	for i := 0; i < 8; i++ {
		buf, err := ring.Get()

		if err != nil || (*buf).(int) != i {
			t.Fatal("Get fail")
		}

		ring.GetAdv()
	}

	_, err = ring.Get()
	if err == nil {
		t.Fatal("Get empty fail")
	}
}