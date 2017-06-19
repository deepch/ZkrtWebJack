package ts

import (
	"testing"
	"../../av"
	"github.com/stretchr/testify/assert"
)

type TestWriter struct {
	buf   []byte
	count int
}

//Write write p to w.buf
func (w *TestWriter) Write(p []byte) (int, error) {
	w.count++
	w.buf = p
	return len(p), nil
}

func TestTSEncoder(t *testing.T) {
	at := assert.New(t)
	m := NewMuxer()

	w := &TestWriter{}
	data := []byte{0xaf, 0x01, 0x21, 0x19, 0xd3, 0x40, 0x7d, 0x0b, 0x6d, 0x44, 0xae, 0x81,
		       0x08, 0x00, 0x89, 0xa0, 0x3e, 0x85, 0xb6, 0x92, 0x57, 0x04, 0x80, 0x00, 0x5b, 0xb7,
		       0x78, 0x00, 0x84, 0x00, 0x00, 0x00, 0x00, 0x00, 0x38, 0x30, 0x00, 0x06, 0x00, 0x38,
	}
	p := av.Packet{
		IsVideo: false,
		Data:    data,
	}
	err := m.Mux(&p, w)
	at.Equal(err, nil)
	at.Equal(w.count, 1)
	at.Equal(w.buf, []byte{0x47, 0x41, 0x01, 0x31, 0x81, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			       0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			       0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			       0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			       0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			       0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			       0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			       0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			       0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
			       0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x01, 0xc0, 0x00, 0x30,
			       0x80, 0x80, 0x05, 0x21, 0x00, 0x01, 0x00, 0x01, 0xaf, 0x01, 0x21, 0x19, 0xd3, 0x40, 0x7d,
			       0x0b, 0x6d, 0x44, 0xae, 0x81, 0x08, 0x00, 0x89, 0xa0, 0x3e, 0x85, 0xb6, 0x92, 0x57, 0x04,
			       0x80, 0x00, 0x5b, 0xb7, 0x78, 0x00, 0x84, 0x00, 0x00, 0x00, 0x00, 0x00, 0x38, 0x30, 0x00,
			       0x06, 0x00, 0x38})
}
