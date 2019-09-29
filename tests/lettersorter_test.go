package tests

import (
	"testing"

	"github.com/IndBrony/lettersorter"
)

func TestSortString(t *testing.T) {
	basicSortStringTesting(t, "osama", "aaoms")
	basicSortStringTesting(t, "omama", "aaomm")
	basicSortStringTesting(t, "test", "estt")
}
func basicSortStringTesting(t *testing.T, input, expectedOutput string) {
	output := lettersorter.SortString(input)
	if output != expectedOutput {
		t.Errorf("Test Fail, expecting %v but got %v", expectedOutput, output)
	}
}
