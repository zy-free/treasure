package ringbuf

import "errors"

// Ring ring proto buffer.
type Ring struct {
	num  uint64
	mask uint64

	// write index
	wpIdx uint64
	// read index
	rpIdx uint64
	data  []interface{}
}

// NewRing new a ring buffer.
func NewRing(num uint64) *Ring {
	r := new(Ring)
	r.init(num)
	return r
}

func (r *Ring) init(num uint64) {
	// 2^N can easy to scale or reduce
	if num&(num-1) != 0 {
		for num&(num-1) != 0 {
			num &= num - 1
		}
		num = num << 1
	}
	r.data = make([]interface{}, num)
	r.num = num
	r.mask = r.num - 1
}

// Get get a proto from ring.
func (r *Ring) Get() (data *interface{}, err error) {
	if r.rpIdx == r.wpIdx {
		return nil, errors.New("ErrRingEmpty")
	}
	data = &r.data[r.rpIdx&r.mask]
	return
}

// GetAdv incr read index.
func (r *Ring) GetAdv() {
	r.rpIdx++
}

// Set get a proto to write.
func (r *Ring) Set() (data *interface{}, err error) {
	if r.wpIdx-r.rpIdx >= r.num {
		return nil, errors.New("ErrRingFull")
	}
	data = &r.data[r.wpIdx&r.mask]
	return
}

// SetAdv incr write index.
func (r *Ring) SetAdv() {
	r.wpIdx++
}

// Reset reset ring.
func (r *Ring) Reset() {
	r.rpIdx = 0
	r.wpIdx = 0
}