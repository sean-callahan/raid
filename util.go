package raid

import "fmt"

// Byte sizes in base-10
const (
	Byte     = 1
	Kilobyte = 1000 * Byte
	Megabyte = 1000 * Kilobyte
	Gigabyte = 1000 * Megabyte
	Terabyte = 1000 * Gigabyte
	Petabyte = 1000 * Terabyte
	Exabyte  = 1000 * Petabyte
)

// ByteSize returns a string representation of the smallest multiple
// of `b` bytes with its appropriate suffix.
func ByteSize(b uint64) string {
	var unit string
	bytes := float64(b)

	switch {
	case bytes >= Exabyte:
		unit = "EB"
		bytes /= Exabyte
		break
	case bytes >= Petabyte:
		unit = "PB"
		bytes /= Petabyte
		break
	case bytes >= Terabyte:
		unit = "TB"
		bytes /= Terabyte
		break
	case bytes >= Gigabyte:
		unit = "GB"
		bytes /= Gigabyte
		break
	case bytes >= Megabyte:
		unit = "MB"
		bytes /= Megabyte
		break
	case bytes >= Kilobyte:
		unit = "KB"
		bytes /= Kilobyte
		break
	case bytes >= Byte:
		unit = "bytes"
		break
	default:
		return "0 bytes"
	}

	return fmt.Sprintf("%g %s", bytes, unit)
}

// Percent converts a decimal to a percentage.
func Percent(val float64) string {
	return fmt.Sprintf("%g%%", val*100)
}
