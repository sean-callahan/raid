package raid

import (
	"fmt"
	"testing"
)

var tests = map[uint64]string{
	Byte:            "1 B",
	Kilobyte:        "1 KB",
	Megabyte:        "1 MB",
	Gigabyte:        "1 GB",
	Terabyte:        "1 TB",
	Petabyte:        "1 PB",
	Exabyte:         "1 EB",
	10 * Gigabyte:   "10 GB",
	1230 * Terabyte: "1.23 PB",
	1100 * Gigabyte: "1.1 TB",
	0:               "0",
	1337 * Petabyte: "1.337 EB",
}

func TestByteSize(t *testing.T) {
	for k, v := range tests {
		result := ByteSize(k)
		if result != v {
			t.Errorf("Unexpected output! Need [%s] -> Got [%s]", v, result)
		}
		fmt.Printf("%s %s\n", v, result)
	}
}
