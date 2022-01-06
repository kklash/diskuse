package diskuse_test

import (
	"fmt"
	"testing"

	"github.com/kklash/diskuse"
)

func TestReadableFileSize(t *testing.T) {
	fixtures := []struct {
		size   int64
		output string
	}{
		{0, "0 b"},
		{1, "1 b"},
		{10, "10 b"},
		{100, "100 b"},
		{1000, "1000 b"},
		{1024, "1 KiB"},
		{1025, "1 KiB"},
		{1027, "1 KiB"},
		{1030, "1.01 KiB"},
		{1100, "1.07 KiB"},
		{1024 * 999, "999 KiB"},
		{1024 * 1022, "1020 KiB"},
		{1024 * 1024, "1 MiB"},
		{5 * 1024 * 1024 * 1024, "5 GiB"},
		{5*1024*1024*1024 + 758629433, "5.71 GiB"},
		{0xffffffffffff, "256 TiB"},
	}

	for _, fixture := range fixtures {
		result := diskuse.ReadableFileSize(fixture.size)
		if result != fixture.output {
			t.Errorf("failed to format size in human readable form\nWanted %s\nGot    %s", fixture.output, result)
		}
	}
}

func ExampleReadableFileSize() {
	fmt.Println(diskuse.ReadableFileSize(234))
	fmt.Println(diskuse.ReadableFileSize(38812))
	fmt.Println(diskuse.ReadableFileSize(99299992))
	fmt.Println(diskuse.ReadableFileSize(711222222222222))
	// output:
	// 234 b
	// 37.9 KiB
	// 94.7 MiB
	// 647 TiB
}
