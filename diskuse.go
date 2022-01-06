// Package diskuse provides human-readable pretty-printing of file sizes.
package diskuse

import (
	"fmt"
	"math"

	"github.com/kklash/rounders"
)

const (
	// Binary file byte-unit sizes.
	Byte     int64 = 1
	Kilobyte int64 = 1024
	Megabyte int64 = 1024 * 1024
	Gigabyte int64 = 1024 * 1024 * 1024
	Terabyte int64 = 1024 * 1024 * 1024 * 1024
	Petabyte int64 = 1024 * 1024 * 1024 * 1024 * 1024
	Exabyte  int64 = 1024 * 1024 * 1024 * 1024 * 1024 * 1024
)

var byteUnits = []string{
	"b",
	"KiB",
	"MiB",
	"GiB",
	"TiB",
	"PiB",
	"EiB",
	"ZiB",
	"YiB",
}

// ReadableFileSize returns a stringification of the given file size with an appended byte
// unit size suffix, such as "GiB" or "MiB", depending on the scale of the given file size.
// The size is rounded to three significant figures.
func ReadableFileSize(size int64) string {
	var (
		scale     float64 = 1
		suffix    string  = "b"
		sizeFloat         = float64(size)
	)

	for i := len(byteUnits) - 1; i >= 0; i-- {
		unitSize := math.Pow(1024, float64(i))
		if sizeFloat >= unitSize {
			scale = unitSize
			suffix = byteUnits[i]
			break
		}
	}

	scaledSize := rounders.RoundToSigFigs(sizeFloat/scale, 3)

	return fmt.Sprintf("%v %s", scaledSize, suffix)
}
