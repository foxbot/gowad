package gowad

import "testing"

var readIntData = map[int]([]byte){
	0:          []byte{0x0, 0x0, 0x0, 0x0},
	1:          []byte{0x0, 0x0, 0x0, 0x1},
	255:        []byte{0x0, 0x0, 0x0, 0xFF},
	48996426:   []byte{0x02, 0xEB, 0xA0, 0x4A},
	2147483647: []byte{0x7F, 0xFF, 0xFF, 0xFF},
}

func TestReadInt(t *testing.T) {
	for expected, data := range readIntData {
		result := readInt(data)
		if result != expected {
			t.Errorf("readInt failed (expected %d, got %d)", expected, result)
		}
	}
}
